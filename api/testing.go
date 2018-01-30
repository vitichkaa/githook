package api

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/go-chi/chi"
)

func (api *API) initTestingRoutes(r chi.Router) {
	r.Get("/hello", api.HelloServer)
	r.Get("/file", api.Download)
}

func (api *API) HelloServer(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello guest!\n")
}

func (api *API) Download(w http.ResponseWriter, r *http.Request) {
	file, _ := os.Open("список.txt")
	body, _ := ioutil.ReadAll(file)
	w.Header().Set("Content-Type", "applicaiton/octet-stream")
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", "список.txt"))
	w.Write(body)
}
