package main

import (
	"context"
	"crypto/sha256"
	"fmt"
	"log"
	"math/rand"
	"strings"

	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/libp2p/go-libp2p/core/peerstore"

	ma "github.com/multiformats/go-multiaddr"
)

type Node struct {
	Host       host.Host
	Blockchain Blockchain
}

type Blockchain []Block

var currentVal int

const difficulty = 6

type Block struct {
	Val   int
	Nonce int
	Hash  string
}

func (b Block) String() string {
	return fmt.Sprintf("Val: %d, Nonce: %d, Hash: %s", b.Val, b.Nonce, b.Hash)
}

func newPeer(peers []string) {

	var logger = log.New(log.Writer(), "miner: ", log.Flags())

	// create the miner node
	node, err := createMinerNode(peers)
	if err != nil {
		logger.Printf("Error creating miner node: %v", err)
		return
	}

	logger.SetPrefix(fmt.Sprintf("miner %v: ", node.Host.ID()))

	logger.Println("info: miner node started")

	// start mining
	mine(node, logger)
}

func createMinerNode(peers []string) (Node, error) {
	// start a libp2p node that listens on a random local TCP port
	host, err := libp2p.New(
		libp2p.ListenAddrStrings("/ip4/127.0.0.1/tcp/0"),
	)

	if err != nil {
		return Node{}, fmt.Errorf("error: create node: %v", err)
	}

	node := Node{Host: host, Blockchain: Blockchain{}}

	if len(peers) == 0 {
		genesis := Block{Val: 0, Nonce: 0}
		genesis.Hash = calculateHash(genesis)
		node.Blockchain = append(node.Blockchain, genesis)
	} else {
		// TODO: connect to peers
		for _, p := range peers {
			paddr, err := ma.NewMultiaddr(p)
			if err != nil {
				return Node{}, fmt.Errorf("error: create peer address: %v", err)
			}
			info, err := peer.AddrInfoFromP2pAddr(paddr)
			if err != nil {
				return Node{}, fmt.Errorf("error: create peer info: %v", err)
			}

			node.Host.Peerstore().AddAddr(info.ID, paddr, peerstore.PermanentAddrTTL)
			if err := node.Host.Connect(context.Background(), *info); err != nil {
				return Node{}, fmt.Errorf("error: connect to peer: %v", err)
			}
		}
		// TODO: sync blockchain
	}

	return node, nil
}

func mine(node Node, logger *log.Logger) {
	for {
		currentVal = rand.Int()
		logger.Println("info: mining: ", currentVal)

		block := Block{Val: currentVal, Nonce: 0}
		block.Hash = calculateHash(block)
		for !strings.HasPrefix(block.Hash, strings.Repeat("0", difficulty)) {
			block.Nonce++
			block.Hash = calculateHash(block)
		}
		logger.Println("info: block mined: ", block.String())
		node.Blockchain = append(node.Blockchain, block)
	}
}

func calculateHash(b Block) string {
	blockData := fmt.Sprintf("%d%d", b.Val, b.Nonce)
	hashInBytes := sha256.Sum256([]byte(blockData))
	return fmt.Sprintf("%x", hashInBytes)
}
