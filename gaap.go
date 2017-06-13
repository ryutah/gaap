package gaap

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Gaap is main struct of gaap.
type Gaap struct {
	router    *mux.Router
	Namespace string
}

// HandlerFunc defines a function to server HTTP requests.
type HandlerFunc func(ctx *Context)

// New create new gaap object.
func New() *Gaap {
	return &Gaap{router: mux.NewRouter()}
}

// WithNameSpace create new gaap object has given namespace.
func WithNameSpace(namespace string) *Gaap {
	return &Gaap{router: mux.NewRouter(), Namespace: namespace}
}

// Start is method of start server.
func (g *Gaap) Start() {
	http.Handle("/", g.router)
}

// GET registers a new GET route for a path with matching handler in the router
func (g *Gaap) GET(path string, handler HandlerFunc) {
	node := newRouterNode(handler)
	g.router.Handle(path, node).Methods(http.MethodGet)
}

// POST registers a new POST route for a path with matching handler in the router
func (g *Gaap) POST(path string, handler HandlerFunc) {
	node := newRouterNode(handler)
	g.router.Handle(path, node).Methods(http.MethodPost)
}

// PUT registers a new PUT route for a path with matching handler in the router
func (g *Gaap) PUT(path string, handler HandlerFunc) {
	node := newRouterNode(handler)
	g.router.Handle(path, node).Methods(http.MethodPut)
}

// DELETE registers a new DELETE route for a path with matching handler in the router
func (g *Gaap) DELETE(path string, handler HandlerFunc) {
	node := newRouterNode(handler)
	g.router.Handle(path, node).Methods(http.MethodDelete)
}

// Handle registers a new route for a path with matching handler in the router
func (g *Gaap) Handle(path string, handler HandlerFunc, methods ...string) {
	node := newRouterNode(handler)
	g.router.Handle(path, node).Methods(methods...)
}
