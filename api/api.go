package api

import (
	"context"
	"github.com/vitichkaa/githook/conf"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

type API struct {
	router *chi.Mux
	cnf    *conf.Conf

	//hc    *http.Client
	Hs *http.Server
}

func New(cnf conf.Conf) *API {
	api := &API{
		cnf: &cnf,
	}

	api.initRouter()
	return api
}

func (api *API) initRouter() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Route("/helloword/v1", func(r chi.Router) {
		r.Route("/testing", api.initTestingRoutes)

	})
	api.Hs = &http.Server{
		Addr:           api.cnf.Addr,
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20, // 1 Mb
	}
	api.router = r
}

// Start will launch http server and listen connection on the addr
func (api *API) Start() {
	log.Println("launching the Alteryx service at", api.cnf.Addr)
	err := http.ListenAndServe(api.cnf.Addr, api.router)
	if err != nil {
		log.Panic("ListenAndServe", err)
	}
}

// Stop will call Shutdown function
func (api *API) Stop() error {
	ctx, cnl := context.WithTimeout(context.Background(), 5*time.Second)
	defer cnl()
	return api.Hs.Shutdown(ctx)
}

func setContentType(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		h.ServeHTTP(w, r)
	})
}
