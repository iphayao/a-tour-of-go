package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m = map[string]Vertex{
	"Bell Labs": Vertex{
		40.68, -74.40,
	},
	"Google": Vertex{
		37.42, -122.08,
	},
}

func main() {
	fmt.Println(m)
}
