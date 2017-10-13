package main

import (
	"errors"
	"fmt"
	"math"
	"strings"
)

// WeightedDAG represents a Weighted Directed Acyclic Graph
type WeightedDAG map[string]map[string]float64

// Costs maps node name to the cost of traveling to it from the start node
type Costs map[string]float64

// Parents maps node name to its parent node name
type Parents map[string]string

func prepareGraph() WeightedDAG {
	graph := make(WeightedDAG)

	graph["book"] = map[string]float64{
		"disc":   5.0,
		"poster": 0.0,
	}

	graph["disc"] = map[string]float64{
		"drums":  20.0,
		"guitar": 15.0,
	}

	graph["poster"] = map[string]float64{
		"guitar": 30.0,
		"drums":  35.0,
	}

	graph["guitar"] = map[string]float64{
		"piano": 20.0,
		"bike":  5.0,
	}

	graph["drums"] = map[string]float64{
		"piano": 10.0,
		"bike":  10.0,
	}

	graph["piano"] = map[string]float64{}

	graph["bike"] = map[string]float64{}

	return graph
}

func findMinimumNode(costs Costs, processedNodes map[string]bool) (nodeName string) {
	minCost := math.Inf(1)
	for name := range costs {
		if _, ok := processedNodes[name]; costs[name] < minCost && ok == false {
			minCost = costs[name]
			nodeName = name
		}
	}
	return
}

func DijkstraPathSearch(graph WeightedDAG, sourceNodeName string, targetNodeName string) ([]string, error) {
	if len(graph) == 0 {
		return nil, errors.New("graph is empty")
	}
	var sourceNodeExists, targetNodeExists bool
	// setting initial parents to empty string
	var parents = make(Parents, len(graph))
	for nodeName := range graph {
		parents[nodeName] = ""

		// check that both source and target nodes present in the graph
		if nodeName == sourceNodeName {
			sourceNodeExists = true
		}
		if nodeName == targetNodeName {
			targetNodeExists = true
		}
	}

	if !(sourceNodeExists && targetNodeExists) {
		return nil, errors.New("one or both nodes do not belong to the graph")
	}

	// slice of already processed nodes
	var processedNodes = make(map[string]bool, len(graph))
	// mark source node as processed
	processedNodes[sourceNodeName] = true

	var costs = make(Costs)
	// assign initial cost to all nodes
	for nodeName := range graph {
		costs[nodeName] = math.Inf(1)
	}
	// set the cost of traveling to the the source node to zero
	costs[sourceNodeName] = 0.0

	currentNodeName := sourceNodeName
	for currentNodeName != "" {
		// get the cost of traveling to the current node
		currentNodeCost := costs[currentNodeName]
		// get all neighbors of the node
		neighbors := graph[currentNodeName]
		for node := range neighbors {
			// compute the cost of traveling to the neighbor from the current node
			neighborCost := currentNodeCost + neighbors[node]
			if neighborCost < costs[node] {
				// if it's cheaper to travel to the neighbor node from the current
				// then update its cost and set new parent of the neighbor node
				costs[node] = neighborCost
				parents[node] = currentNodeName
			}
		}
		// mark current node as processed so we don't process same node twice
		processedNodes[currentNodeName] = true
		// find next node with minimum traveling cost
		currentNodeName = findMinimumNode(costs, processedNodes)
	}

	// building result
	var cheapestPath []string
	nodeHop := targetNodeName

	for nodeHop != "" {
		// insert to the beginning of the slice
		cheapestPath = append([]string{nodeHop}, cheapestPath...)
		nodeHop = parents[nodeHop]
	}

	// if path does not start with the source node name
	// then there is no way between source and target nodes
	// the same comes if path does not end with target node name
	if len(cheapestPath) > 0 &&
		(cheapestPath[0] != sourceNodeName || cheapestPath[len(cheapestPath)-1] != targetNodeName) {
		return nil, errors.New("the is no path between source and target nodes")
	}
	return cheapestPath, nil
}

func main() {
	graph := prepareGraph()
	emptyGraph := make(WeightedDAG)
	type TestCase struct {
		Graph          *WeightedDAG
		SourceNode     string
		TargetNode     string
		ExpectedResult []string
		ErrorMsg       string
	}

	tests := []TestCase{
		{
			Graph:          &graph,
			SourceNode:     "book",
			TargetNode:     "piano",
			ExpectedResult: []string{"book", "disc", "drums", "piano"},
			ErrorMsg:       "",
		},
		{
			Graph:          &graph,
			SourceNode:     "book",
			TargetNode:     "bike",
			ExpectedResult: []string{"book", "disc", "guitar", "bike"},
			ErrorMsg:       "",
		},
		{
			Graph:          &graph,
			SourceNode:     "book",
			TargetNode:     "book",
			ExpectedResult: []string{"book"},
			ErrorMsg:       "",
		},
		{
			Graph:          &graph,
			SourceNode:     "book",
			TargetNode:     "poster",
			ExpectedResult: []string{"book", "poster"},
			ErrorMsg:       "",
		},
		{
			Graph:          &graph,
			SourceNode:     "poster",
			TargetNode:     "book",
			ExpectedResult: nil,
			ErrorMsg:       "the is no path between source and target nodes",
		},
		{
			Graph:          &emptyGraph,
			SourceNode:     "book",
			TargetNode:     "piano",
			ExpectedResult: nil,
			ErrorMsg:       "graph is empty",
		},
		{
			Graph:          &graph,
			SourceNode:     "",
			TargetNode:     "",
			ExpectedResult: nil,
			ErrorMsg:       "one or both nodes do not belong to the graph",
		},
		{
			Graph:          &graph,
			SourceNode:     "new book",
			TargetNode:     "piano",
			ExpectedResult: nil,
			ErrorMsg:       "one or both nodes do not belong to the graph",
		},
	}

	for _, testCase := range tests {
		cheapestPath, err := DijkstraPathSearch(*testCase.Graph, testCase.SourceNode, testCase.TargetNode)

		if err != nil && err.Error() != testCase.ErrorMsg {
			fmt.Printf("Got unexpected error: %s", err)
			return
		}

		if strings.Join(cheapestPath, "") != strings.Join(testCase.ExpectedResult, "") {
			fmt.Printf("Unexpected result: %s. Expected result: %s", cheapestPath, testCase.ExpectedResult)
			return
		}
	}

	fmt.Print("finished")
}
