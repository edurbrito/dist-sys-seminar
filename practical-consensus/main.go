package main

func main() {
	peers := []string{}

	for i := 0; i < 3; i++ {
		peer, err := startNode(peers)
		if err != nil {
			panic(err)
		}
		peers = append(peers, peer)
	}

	select {}
}
