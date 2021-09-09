package main

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func GenerateProblemRandom(size int) ([][]int, int, int) {
	rand := rand.New(rand.NewSource(time.Now().Unix()))

	src := rand.Intn(size)
	dst := rand.Intn(size)

	// Fill in distances with random values for each Vertex and its adjacent vertices
	dist := make([][]int, size)
	for vId := range dist {
		dist[vId] = make([]int, size)
		for adjId := range dist[vId] {
			if vId == adjId {
				dist[vId][adjId] = 0 // Already is 0 but to be explicit
			} else {
				dist[vId][adjId] = rand.Intn(99) + 1
			}
		}
	}

	return dist, src, dst
}

func ReadGenerateProblem(r io.Reader) ([][]int, int, int, error) {
	s := bufio.NewScanner(r)
	size, src, dst := -1, -1, -1
	var dist [][]int
	for line := 1; s.Scan(); line++ {
		text := s.Text()
		switch line {
		case 1:
			n, err := strconv.Atoi(text)
			if err != nil {
				return nil, 0, 0, fmt.Errorf("invalid size: %w", err)
			}
			if n <= 0 {
				return nil, 0, 0, fmt.Errorf("invalid size '%d': size must > 0", n)
			}
			size = n
		case 2:
			n, err := strconv.Atoi(text)
			if err != nil {
				return nil, 0, 0, fmt.Errorf("invalid source: %w", err)
			}
			if n < 0 || n > size {
				return nil, 0, 0, fmt.Errorf("invalid source '%d': source must be in range 1-%d", n, size)
			}
			src = n - 1 // Convert NodeId to VertexId (zero based)
		case 3:
			n, err := strconv.Atoi(text)
			if err != nil {
				return nil, 0, 0, fmt.Errorf("invalid destination: %w", err)
			}
			if n < 0 || n > size {
				return nil, 0, 0, fmt.Errorf("invalid destination '%d': destination must be in range 1-%d", n, size)
			}
			dst = n - 1 // Convert NodeId to VertexId (zero based)
		default:
			parts := strings.Fields(text)
			if len(parts) != size {
				return nil, 0, 0, fmt.Errorf("invalid distance matrix: each row must be of length %d", size)
			}

			row := make([]int, len(parts))
			for i, part := range parts {
				n, err := strconv.Atoi(part)
				if err != nil {
					return nil, 0, 0, fmt.Errorf("invalid distance: %w", err)
				}
				if len(dist) == i {
					if n != 0 {
						return nil, 0, 0, fmt.Errorf("invalid distance '%d': distance from each vertex to itself must be 0", n)
					}
				} else if n <= 0 {
					return nil, 0, 0, fmt.Errorf("invalid distance '%d': distance from each vertex to an adjacent vertex must be greater than 0", n)
				}
				row[i] = n
			}
			dist = append(dist, row)
			// assert
		}
	}

	if err := s.Err(); err != nil {
		return nil, 0, 0, err
	}

	if len(dist) != size {
		return nil, 0, 0, fmt.Errorf("invalid adjacency matrix: expected %d rows but got %d", size, len(dist))
	}

	return dist, src, dst, nil
}
