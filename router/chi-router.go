package router

import "net/http"

type chiRouter struct {
}

var ()

//NewChiRouter implemented chi for routes
func NewChiRouter() IRouter {
	return &chiRouter{}
}

func (c *chiRouter) SERVE(port string) {

}

func (c *chiRouter) GET(uri string, f func(w http.ResponseWriter, r *http.Request)) {

}

func (c *chiRouter) POST(uri string, f func(w http.ResponseWriter, r *http.Request)) {

}
