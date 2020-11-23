package router

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type muxRouter struct{}

var (
	muxDispatcher = mux.NewRouter()
)

//NewMuxRouter instancead a new Router
func NewMuxRouter() IRouter {
	return &muxRouter{}
}

func (m *muxRouter) SERVE(port string) {
	fmt.Printf("Mux HTTP server running on port %v", port)
	http.ListenAndServe(port, muxDispatcher)
}

func (m *muxRouter) GET(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	muxDispatcher.HandleFunc(uri, f).Methods(http.MethodGet)
}

func (m *muxRouter) POST(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	muxDispatcher.HandleFunc(uri, f).Methods(http.MethodPost)
}
