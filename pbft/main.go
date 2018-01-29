package main

import (
	"os"
	"github.com/byron1st/sampleConsensus/pbft/network"
)

func main() {
	nodeID := os.Args[1]
	server := network.NewServer(nodeID)

	server.Start()
}
