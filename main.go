package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

var generate = flag.Bool("random", false, "Generates a random problem graph, source and destination vertices and solves it. Provided as a short and simple way to get started")
var file = flag.String("file", "", "Path to the file to read from. By default, the Stdin is used to read the input")
var help = flag.Bool("help", false, "Prints help message")

func main() {
	flag.Parse()

	var (
		dist [][]int
		src  int
		dst  int
	)
	if *generate {
		dist, src, dst = GenerateProblemRandom(6)
	} else if *help {
		flag.Usage()
		os.Exit(1)
	} else {
		var r io.Reader
		if *file == "" {
			r = os.Stdin
		} else {
			f, err := os.Open(*file)
			if err != nil {
				fmt.Println(err.Error())
				os.Exit(1)
			}

			r = bufio.NewReader(f)
		}

		var err error
		dist, src, dst, err = ReadGenerateProblem(r)
		if err != nil {
			fmt.Println("Input error:", err.Error())
			os.Exit(1)
		}
	}

	printAdjMatrix(dist)
	path := FindShortestPath(dist, src, dst)
	printFastestRout(path, dist, src, dst)
}

func printAdjMatrix(dist [][]int) {
	fmt.Println("Adjacency Distance Matrix")
	fmt.Println("---")
	for vId, _ := range dist {
		fmt.Printf("\t%s", nodeId(vId))
	}
	fmt.Println()

	for vId, _ := range dist {
		fmt.Printf("%s|\t", nodeId(vId))
		for adjId := range dist[vId] {
			fmt.Printf("%5d\t", dist[vId][adjId])
		}
		fmt.Println()
	}
	fmt.Println()
}

func printFastestRout(path []int, dist [][]int, src, dst int) {
	fmt.Printf("Fastest route between %s to %s\n", nodeId(src), nodeId(dst))
	fmt.Println("---")

	for i := 0; i < len(path)-1; i++ {
		thisV := path[i]
		nextV := path[i+1]
		fmt.Printf("%s --(%d)--> ", nodeId(thisV), dist[thisV][nextV])
	}
	fmt.Printf("%s\n", nodeId(path[len(path)-1]))
}

func nodeId(index int) string {
	return fmt.Sprintf("Node%d", index+1)
}
