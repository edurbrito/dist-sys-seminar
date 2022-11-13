<template>
  <div class="flex flex-col items-center justify-center h-full w-full">
    <div class="id-container">
      <h1 class="id">My ID: #{{ uuid.toUpperCase() }}</h1>
      <div class="id-bg" :style="{ 'background-color': color }"></div>
    </div>
    <BlockComponent
      v-for="(block, index) in blockchain.mainChain"
      :key="block.hash"
      :uuid="uuid"
      :block="block"
      :z-index="blockchain.mainChain.length - index"
    />
  </div>
</template>

<script lang="ts">
import Vue from "vue";
import GUN from "gun";
import BlockComponent from "@/components/Block.vue";

import Miner from "~/src/miner.worker.ts";
import { Block } from "~/src/Block";
import { Blockchain } from "~/src/Blockchain";
const { v4: uuidv4 } = require("uuid");

export default Vue.extend({
  name: "IndexPage",
  components: {
    BlockComponent,
  },
  data() {
    return {
      uuid: uuidv4().slice(0, 6),
      currentBlock: Block.genesisBlock(),
      blockchain: new Blockchain(),
    };
  },
  computed: {
    color(): string {
      return "#" + this.uuid;
    },
  },
  beforeMount() {
    document.body.style.background =
      "radial-gradient(circle, #66bfbf 50%," + this.color + " 150%) fixed";
  },
  mounted() {
    // initialize the miner
    let miner = new Miner();

    // connect to the relay server(s)
    let peers = process.env.gunRelays?.split(",");
    let gun = GUN({ peers: peers });

    // get the current (first) block
    gun
      .get("gossip")
      .get("block")
      .once((data) => {
        if (data) {
          let block = JSON.parse(data);
          this.currentBlock = new Block(
            block.hash,
            block.previousHash,
            block.nonce,
            block.data
          );
        }
        this.blockchain.addBlock(this.currentBlock);
      });

    // get mined blocks from the miner
    let minerOnMessage = (e: any) => {
      let block = e.data;

      // gossip the block
      gun
        .get("gossip")
        .get("block")
        .put(
          JSON.stringify({
            previousHash: block.previousHash,
            nonce: block.nonce,
            data: block.data,
          })
        );
    };

    // listen for new blocks
    gun
      .get("gossip")
      .get("block")
      .on((data, key) => {
        // get the block
        let block = JSON.parse(data);
        // calculate the hash
        let hash = Block.calculateHash(
          block.previousHash,
          block.nonce,
          block.data
        );
        // validate the block
        if (Block.isHashValid(hash)) {
          // update the value
          this.currentBlock = new Block(
            hash,
            block.previousHash,
            block.nonce,
            block.data
          );
          // add the block to the blockchain
          this.blockchain.addBlock(this.currentBlock);
          // terminate the worker
          miner.terminate();
          // start a new worker
          miner = new Miner();
          // listen for mined blocks
          miner.onmessage = minerOnMessage;
          // mine on the new block
          miner.postMessage(["mine", this.currentBlock, this.uuid]);
        }
      });

    // listen for mined blocks
    miner.onmessage = minerOnMessage;
    // mine on the new block
    miner.postMessage(["mine", this.currentBlock, this.uuid]);
  },
});
</script>

<style>
.id-container {
  padding: 1rem;
  padding-left: 2rem;
  margin-bottom: 25vh;
  align-self: start;
  width: 100%;
  font-weight: 700;
  position: relative;
}

.id {
  position: absolute;
  color: #fff;
}

.id-bg {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 200%;
  opacity: 0.9;
  z-index: -1;
}
</style>
