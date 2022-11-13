import { Block } from "./Block";

const ctx: Worker = self as any;

function mine(currentBlock: Block, identifier: string) {
    let nonce = 0;
    // create a new block
    while(true) {
            let hash = Block.calculateHash(currentBlock.hash, nonce, identifier);
            if(Block.isHashValid(hash)) {
                let block = new Block(hash, currentBlock.hash, nonce, identifier);
                currentBlock = block;
    
                // send the new block to the main thread
                ctx.postMessage(block);
                nonce = 0;
            }
            nonce++;
    }
}

// on message "mine" from the main thread
ctx.onmessage = (e) => {
    if (e.data[0] === "mine") {
        mine(e.data[1], e.data[2]);
    }
}