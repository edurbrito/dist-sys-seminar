import sha256 from 'crypto-js/sha256';

const DIFFICULTY = parseInt(process.env.difficulty?.toString() || '5');
const TARGET = parseInt(`0x${"0".repeat(DIFFICULTY)}${"f".repeat(64-DIFFICULTY)}`,16);

// Block class
export class Block {

    public hash: string;
    public previousHash: string;
    public nonce: number;
    public data: string;

    constructor(hash: string, previousHash: string, nonce: number, data: string) {
        this.hash = hash;
        this.previousHash = previousHash;
        this.nonce = nonce;
        this.data = data;
    }

    public prettyPrint() {
        return `Block:
        Hash: ${this.hash}
        Previous Hash: ${this.previousHash}
        Nonce: ${this.nonce}
        Data: ${this.data}
        `;
    }

    static isHashValid(hash: string) {
        return parseInt(hash, 16) < TARGET;
    }

    static calculateHash(previousHash: string, nonce: number, data: string): string {
        return sha256(previousHash + nonce + data).toString();
    }

    static genesisBlock(): Block {
        return new Block('0000003466f7289d52851720ad259c88f192604759e34ece83d6eb74fa03b5ed', '', 9865212, 'edurbrito');
    }
}
