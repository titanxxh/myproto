package trait

import (
	"fmt"
)

type (
	router struct {
		routees []int
	}
)

func (t *router) RouteMessage(key int) {
	fmt.Println(key)
}

func (t *router) init() {
	t = routerBuilder()
}

func routerBuilder() *router {
	return &router{}
}
