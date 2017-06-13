package gaap

import (
	"net/http"
)

type routeNode struct {
	handler HandlerFunc
}

func newRouterNode(handler HandlerFunc) *routeNode {
	return &routeNode{handler}
}

func (rn *routeNode) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := NewContext(w, r)
	rn.handler(ctx)
}
