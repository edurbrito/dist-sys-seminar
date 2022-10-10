package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/libp2p/go-libp2p/core/peerstore"

	pubsub "github.com/libp2p/go-libp2p-pubsub"
	ma "github.com/multiformats/go-multiaddr"
)

func startNode(peers []string) (string, error) {
	var logger = log.New(log.Writer(), "peer: ", log.Flags())

	node, err := createNode(peers)
	if err != nil {
		logger.Printf("error: create node: %v", err)
		return "", err
	}

	topic, sub, err := createPubSub(node, logger)
	if err != nil {
		logger.Printf("error: create pubsub: %v", err)
		return "", err
	}

	logger.SetPrefix(fmt.Sprintf("peer %v: ", node.ID().ShortString()))

	logger.Println("info: peer started")

	go subscribe(node, sub, logger)

	go publish(node, topic, logger)

	return getPeerMultiAddress(node), nil
}

func createNode(peers []string) (host.Host, error) {
	// start a libp2p node that listens on a random local TCP port
	node, err := libp2p.New(
		libp2p.ListenAddrStrings("/ip4/127.0.0.1/tcp/0"),
	)
	if err != nil {
		return nil, err
	}

	// add the peers to the peerstore
	for _, p := range peers {
		paddr, err := ma.NewMultiaddr(p)
		if err != nil {
			return nil, fmt.Errorf("error: create peer address: %v", err)
		}
		info, err := peer.AddrInfoFromP2pAddr(paddr)
		if err != nil {
			return nil, fmt.Errorf("error: create peer info: %v", err)
		}

		node.Peerstore().AddAddr(info.ID, paddr, peerstore.PermanentAddrTTL)
		if err := node.Connect(context.Background(), *info); err != nil {
			return nil, fmt.Errorf("error: connect to peer: %v", err)
		}
	}

	return node, nil
}

func createPubSub(node host.Host, logger *log.Logger) (*pubsub.Topic, *pubsub.Subscription, error) {
	ctx := context.Background()

	// create new pubsub topic
	pb, err := pubsub.NewGossipSub(ctx, node)
	if err != nil {
		logger.Printf("error: create topic: %v", err)
		return nil, nil, err
	}

	topic, err := pb.Join("test")
	if err != nil {
		logger.Printf("error: subscribe to topic: %v", err)
		return nil, nil, err
	}

	sub, err := topic.Subscribe()
	if err != nil {
		logger.Printf("error: subscribe to topic: %v", err)
		return nil, nil, err
	}

	return topic, sub, nil
}

type message struct {
	From string
	Data string
}

func subscribe(node host.Host, sub *pubsub.Subscription, logger *log.Logger) {
	ctx := context.Background()
	for {
		msg, err := sub.Next(ctx)
		if err != nil {
			logger.Printf("error: subscribe to topic: %v", err)
			continue
		}

		if msg.ReceivedFrom == node.ID() {
			continue
		}

		m := new(message)
		err = json.Unmarshal(msg.Data, m)
		if err != nil {
			logger.Printf("error: unmarshal message: %v", err)
			continue
		}

		logger.Printf("message: %s", m.Data)
	}
}

func publish(node host.Host, topic *pubsub.Topic, logger *log.Logger) {
	ctx := context.Background()
	for {
		time.Sleep(time.Duration(rand.Intn(6)+1) * time.Second)

		m := message{
			From: node.ID().ShortString(),
			Data: fmt.Sprintf("hello from %s", node.ID().ShortString()),
		}

		mBytes, err := json.Marshal(m)
		if err != nil {
			logger.Printf("error: marshal message: %v", err)
			continue
		}

		err = topic.Publish(ctx, mBytes)
		if err != nil {
			logger.Printf("error: publish message: %v", err)
			continue
		}
	}
}

func getPeerMultiAddress(ha host.Host) string {
	hostAddr, _ := ma.NewMultiaddr(fmt.Sprintf("/p2p/%s", ha.ID().Pretty()))
	addr := ha.Addrs()[0]
	return addr.Encapsulate(hostAddr).String()
}
