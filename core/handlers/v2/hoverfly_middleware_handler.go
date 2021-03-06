package v2

import (
	"encoding/json"
	"github.com/SpectoLabs/hoverfly/core/handlers"
	"github.com/codegangsta/negroni"
	"github.com/go-zoo/bone"
	"io/ioutil"
	"net/http"
)

type HoverflyMiddleware interface {
	GetMiddleware() string
	SetMiddleware(string) error
}

type HoverflyMiddlewareHandler struct {
	Hoverfly HoverflyMiddleware
}

func (this *HoverflyMiddlewareHandler) RegisterRoutes(mux *bone.Mux, am *handlers.AuthHandler) {
	mux.Get("/api/v2/hoverfly/middleware", negroni.New(
		negroni.HandlerFunc(am.RequireTokenAuthentication),
		negroni.HandlerFunc(this.Get),
	))

	mux.Put("/api/v2/hoverfly/middleware", negroni.New(
		negroni.HandlerFunc(am.RequireTokenAuthentication),
		negroni.HandlerFunc(this.Put),
	))
}

func (this *HoverflyMiddlewareHandler) Get(w http.ResponseWriter, req *http.Request, next http.HandlerFunc) {
	var middlewareView MiddlewareView
	middlewareView.Middleware = this.Hoverfly.GetMiddleware()

	middlewareBytes, _ := json.Marshal(middlewareView)

	handlers.WriteResponse(w, middlewareBytes)
}

func (this *HoverflyMiddlewareHandler) Put(w http.ResponseWriter, req *http.Request, next http.HandlerFunc) {
	defer req.Body.Close()

	var middlewareReq MiddlewareView

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		handlers.WriteErrorResponse(w, "Malformed JSON", 400)
		return
	}

	err = json.Unmarshal(body, &middlewareReq)
	if err != nil {
		handlers.WriteErrorResponse(w, "Malformed JSON", 400)
		return
	}

	err = this.Hoverfly.SetMiddleware(middlewareReq.Middleware)
	if err != nil {
		handlers.WriteErrorResponse(w, "Invalid middleware: "+err.Error(), 422)
		return
	}

	this.Get(w, req, next)
}
