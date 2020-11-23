package router

import "net/http"

//IRouter is interface that contains routes for api
type IRouter interface {
	SERVE(port string)
	GET(uri string, f func(w http.ResponseWriter, r *http.Request))
	POST(uri string, f func(w http.ResponseWriter, r *http.Request))
}
