<template>
  <div class="flex flex-row w-3/5 h-3/5 items-center justify-center mr--5">
    <div class="flex flex-col items-center relative">
      <span class="block-data"
        :style="{ zIndex: zIndex * 2 + 1 }">{{ "#" + block.data.toUpperCase() }}</span>
      <svg
        version="1.1"
        id="Layer_1"
        xmlns="http://www.w3.org/2000/svg"
        xmlns:xlink="http://www.w3.org/1999/xlink"
        x="0px"
        y="0px"
        viewBox="0 0 512 512"
        xml:space="preserve"
        :style="{ fill: color, zIndex: zIndex * 2 }"
        class="block"
      >
        <path
          class="block-top-face-bg"
          d="M477.701,128c-2.009-3.48-4.918-6.439-8.518-8.518L267.637,3.119c-7.201-4.158-16.072-4.158-23.273,0
	L42.82,119.482c-3.6,2.079-6.509,5.039-8.518,8.518l221.699,128L477.701,128z"
        />
        <path
          class="block-top-face"
          d="M477.701,128c-2.009-3.48-4.918-6.439-8.518-8.518L267.637,3.119c-7.201-4.158-16.072-4.158-23.273,0
	L42.82,119.482c-3.6,2.079-6.509,5.039-8.518,8.518l221.699,128L477.701,128z"
        />
        <path
          d="M256.001,512c4.018,0,8.037-1.04,11.636-3.119l201.545-116.364
	c7.201-4.158,11.636-11.84,11.636-20.154V139.636c0-4.158-1.109-8.156-3.119-11.636L256.001,256V512z"
        />
        <path
          class="block-left-face-bg"
          d="M244.364,508.881c3.6,2.079,7.618,3.119,11.636,3.119V256L34.3,128
	c-2.009,3.48-3.119,7.478-3.119,11.636v232.727c0,8.315,4.436,15.998,11.636,20.154L244.364,508.881z"
        />
        <path
          class="block-left-face"
          d="M244.364,508.881c3.6,2.079,7.618,3.119,11.636,3.119V256L34.3,128
	c-2.009,3.48-3.119,7.478-3.119,11.636v232.727c0,8.315,4.436,15.998,11.636,20.154L244.364,508.881z"
        />
      </svg>
      <svg
        v-if="hasChain"
        xmlns="http://www.w3.org/2000/svg"
        viewBox="0 0 15 15"
        class="chain"
        :style="{ zIndex: zIndex * 2 - 1 }"
      >
        <path
          d="M0 7.5C0 7.22386 0.223858 7 0.5 7H3C3.27614 7 3.5 7.22386 3.5 7.5C3.5 7.77614 3.27614 8 3 8H0.5C0.223858 8 0 7.77614 0 7.5ZM5.75 7.5C5.75 7.22386 5.97386 7 6.25 7H8.75C9.02614 7 9.25 7.22386 9.25 7.5C9.25 7.77614 9.02614 8 8.75 8H6.25C5.97386 8 5.75 7.77614 5.75 7.5ZM12 7C11.7239 7 11.5 7.22386 11.5 7.5C11.5 7.77614 11.7239 8 12 8H14.5C14.7761 8 15 7.77614 15 7.5C15 7.22386 14.7761 7 14.5 7H12Z"
        />
      </svg>
    </div>
    <div class="block-hashes">
      <span class="block-hash">
        Block Hash: {{ block.hash.slice(-6).toUpperCase() }}
      </span>
      <span class="block-hash mt-1">
        Parent Hash: {{ block.previousHash.slice(-6).toUpperCase() }}
      </span>
    </div>
  </div>
</template>

<script>
export default {
  name: "Block",
  props: {
    uuid: {
      type: String,
      required: true,
    },
    block: {
      type: Object,
      required: true,
    },
    zIndex: {
      type: Number,
      default: 0,
    },
  },
  computed: {
    color() {
      return `#${this.block.data.slice(0, 6)}`;
    },
    hasChain() {
      return this.block.previousHash !== null && this.block.previousHash !== "";
    },
  },
};
</script>

<style>
.chain {
  width: 60%;
  height: 60%;
  transform: rotate(90deg);
  fill: #83fefe;
  margin-bottom: -25%;
  z-index: 1;
}

.block {
  margin-bottom: -1rem;
  z-index: 2;
}

.block-left-face {
  fill-opacity: 80%;
}

.block-left-face-bg {
  fill: white;
}

.block-top-face {
  fill-opacity: 60%;
}

.block-top-face-bg {
  fill: white;
}

.block-hashes {
  display: flex;
  flex-direction: column;
}

.block-hash {
  font-size: 2vmin;
  font-weight: 700;
}

.block-data {
  position: absolute;
  top: 25%;
  font-size: 2vmin;
  font-weight: 700;
}

.mr--5 {
  margin-right: -15vw;
}
</style>
