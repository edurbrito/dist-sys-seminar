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
        return new Block('04ff496d58aae22254d59665a4a5e9dd2a011247b2414c5d8b8456a0353f5761', '', 112358, 'dddddd-edurbrito');
    }
}
