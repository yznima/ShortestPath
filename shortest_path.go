package main

import "container/heap"

// FindShortestPath finds the shortest path from src to dst 
// vertex based on the provided dist,the adjacancy distance matrix.
// dist[i][j] represents the distance from vertices i -> j and 
// dist[j][i] represent the distance from vertices j -> i.
// Returns the path from the src vertex to dst vertex as.
// 
// NOTE: All the inputs and outputs in of this function are based on VertexId
// and not the NodeId. NodeId starts from 1 while VertexId starts from 0.
func FindShortestPath(dist [][]int, src, dst int) []int {
	if src == dst {
		return []int{src}
	}

	type State struct {
		visited bool
		parent  int
		vertex  *Vertex
	}

	// 1.Initiate the states with cost equal to distance from the source vertex
	states := make([]State, len(dist))
	pq := make(PriorityQueue, len(dist))
	for vId := range states {
		vertex := &Vertex{id: vId, cost: dist[src][vId]}

		states[vId] = State{
			visited: false,
			parent:  src,
			vertex:  vertex,
		}

		pq[vId] = vertex
	}
	heap.Init(&pq)

	for pq.Len() > 0 {
		// 2. Find the next not visited vertex with the least cost
		vertex := heap.Pop(&pq).(*Vertex)
		vId := vertex.id
		states[vId].visited = true

		for adjId, distance := range dist[vId] {
			if vId == adjId || states[adjId].visited {
				continue
			}

			// 3. Update the cost of all the adjacent vertices if there is a cheaper route
			newCost := vertex.cost + distance
			if newCost < states[adjId].vertex.cost {
				states[adjId].vertex.cost = newCost
				states[adjId].parent = vId
				heap.Fix(&pq, states[adjId].vertex.index)
			}
		}
	}

	// 4. Backtrack from the destination to source to find the cheapest route
	path := make([]int, 0, len(dist))
	for vId := dst; vId != src; vId = states[vId].parent {
		path = append(path, vId)
	}
	path = append(path, src)

	// For simplicity, reverse the path since it is backward right now
	for i := 0; i < len(path)/2; i++ {
		j := len(path) - 1 - i
		path[i], path[j] = path[j], path[i]
	}

	return path
}
