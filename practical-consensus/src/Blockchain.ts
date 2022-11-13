import { Block } from "./Block";

export class Blockchain {
    public activeBlocks:Map<string, Block> = new Map<string,Block>();
    public mainChain: Block[] = [];

    public addBlock(block: Block): void {
        let parentBlock = this.activeBlocks.get(block.previousHash);
        if (parentBlock !== undefined) {
            this.mainChain.unshift(parentBlock);
            this.activeBlocks.delete(block.previousHash);
        }
        
        if (this.mainChain.length == 0 || this.mainChain[0].hash === block.previousHash) {
            this.activeBlocks.set(block.hash, block);
        }
    }
}