package main

import "fmt"

// Looking for the shortest way between two locations
//
// Road map:
//
//          + ++++++++++  NY ++++++++++++++++++
//          +                                 +
//          +                                 +
//          +                                 +
//          +                                 +
//         SF +++++++++++ LA +++++++++++++ Moscow
//          +                                 +
//          +                                 +
//          +                                 +
//          +                                 +
//          ++++++++++ Istanbul +++++++ Ekaterinburg
//                      +   +                 +
//                      +   +                 +
//            ++++++++++    ++++++++++        +
//            +                      +        +
//            +                      +        +
//            +                      +        +
//       Dubrovnik                 Pula +++++++
//
//
//                  IsolatedLocation

type Route []string

type Destination struct {
	Hops Route
	Name string
}

type SimpleQueue struct {
	Nodes []Destination
}

func NewSimpleQueue() *SimpleQueue {
	return &SimpleQueue{}
}

func (sq *SimpleQueue) Enqueue(node Destination) {
	nodes := make([]Destination, len(sq.Nodes)+1)
	nodes[0] = node
	copy(nodes[1:], sq.Nodes)
	sq.Nodes = nodes
}

func (sq *SimpleQueue) Dequeue() Destination {
	if len(sq.Nodes) == 0 {
		return Destination{}
	}
	node := sq.Nodes[len(sq.Nodes)-1]
	nodes := make([]Destination, len(sq.Nodes)-1)
	copy(nodes, sq.Nodes[:len(sq.Nodes)-1])
	sq.Nodes = nodes
	return node
}

func (sq *SimpleQueue) Len() int {
	return len(sq.Nodes)
}

type RoadMap map[string][]string

func (rm RoadMap) GetPossibleDestinations(nodeName string) []string {
	return rm[nodeName]
}

func NewRoadmap() RoadMap {
	var routes = make(RoadMap, 9)

	routes["NY"] = []string{"SF", "Moscow"}
	routes["SF"] = []string{"NY", "LA", "Istanbul"}
	routes["LA"] = []string{"SF", "Moscow"}
	routes["Moscow"] = []string{"NY", "LA", "Ekaterinburg"}
	routes["Istanbul"] = []string{"SF", "Dubrovnik", "Ekaterinburg", "Pula"}
	routes["Ekaterinburg"] = []string{"Moscow", "Istanbul", "Pula"}
	routes["Dubrovnik"] = []string{"Istanbul"}
	routes["Pula"] = []string{"Istanbul", "Ekaterinburg"}
	// Isolated location with no in-out routes
	routes["IsolatedLocation"] = []string{}

	return routes
}

func BreadthSearch(start_location, finish_location string) Route {
	queue := NewSimpleQueue()
	roadMap := NewRoadmap()
	// in order to avoid cycling over the same locations we keep already visited locations in map[string]bool
	var visitedDestinations = make(map[string]bool)
	visitedDestinations[start_location] = true
	// check direct neighbors first
	queue.Enqueue(Destination{Name: start_location, Hops: Route{start_location}})

	// actual search
	for queue.Len() > 0 {
		location := queue.Dequeue()
		visitedDestinations[location.Name] = true

		// we found the shortest way to the destination location
		if location.Name == finish_location {
			return location.Hops
		}

		// if not found then we queue direct neighbors of the current location to be searched
		locDestinations := roadMap.GetPossibleDestinations(location.Name)
		if len(locDestinations) > 0 {
			for _, destName := range locDestinations {
				if _, ok := visitedDestinations[destName]; ok == false {
					// creating a new route by extending the parent route with current location (destName)
					route := make(Route, len(location.Hops)+1)
					copy(route, location.Hops)
					route[len(route)-1] = destName

					// queueing new Destination for further searching
					node := Destination{Name: destName, Hops: route}
					queue.Enqueue(node)
					visitedDestinations[destName] = true
				}
			}
		}
	}

	return nil
}

func routesEqual(aRoute Route, bRoute Route) bool {
	if len(aRoute) != len(bRoute) {
		return false
	}

	for ind := range aRoute {
		if aRoute[ind] != bRoute[ind] {
			return false
		}
	}

	return true
}

func runShortestRouteSearch(start_location, finish_location string, expectedRoute Route) {
	var res Route
	resultMessage := "The route between %v and %v is %v, and it is the shortest route: %v\n"
	res = BreadthSearch(start_location, finish_location)
	fmt.Printf(resultMessage, start_location, finish_location, res, routesEqual(res, expectedRoute))
}

func main() {
	runShortestRouteSearch("NY", "Dubrovnik", Route{"NY", "SF", "Istanbul", "Dubrovnik"})
	runShortestRouteSearch("SF", "Pula", Route{"SF", "Istanbul", "Pula"})
	runShortestRouteSearch("Moscow", "SF", Route{"Moscow", "NY", "SF"})
	runShortestRouteSearch("LA", "Dubrovnik", Route{"LA", "SF", "Istanbul", "Dubrovnik"})
	runShortestRouteSearch("LA", "Pula", Route{"LA", "SF", "Istanbul", "Pula"})
	runShortestRouteSearch("IsolatedLocation", "Pula", Route{})
	runShortestRouteSearch("NY", "IsolatedLocation", Route{})
}
