package http

import (
	"english-verbs/verbs/app"
	"github.com/go-chi/chi"
	"net/http"
	"strconv"
)

type VerbHandler struct {
	Service app.VerbService
}

func (srv *VerbHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	verbs, err := srv.Service.GetAll(ctx)
	if err != nil {
		HTTPError(w, r, http.StatusNotFound, err.Error())
		return
	}

	JSON(w, r, http.StatusOK, verbs)
}

func (srv *VerbHandler) GetExam(w http.ResponseWriter, r *http.Request) {

	size := getExamSizeFromRequest(r)

	verbs, err := srv.Service.GetExam(size)

	if err != nil {
		HTTPError(w, r, http.StatusNotFound, err.Error())
		return
	}

	JSON(w, r, http.StatusOK, verbs)
}

func getExamSizeFromRequest(r *http.Request) int {

	var size = 10

	query := r.URL.Query()["size"]

	if len(query) == 1 {
		size, _ = strconv.Atoi(query[0])
	}

	return size
}

// Routes returns post router with each endpoint.
func (srv *VerbHandler) Routes() http.Handler {

	r := chi.NewRouter()

	r.Get("/", srv.GetAll)

	r.Get("/exam", srv.GetExam)

	return r
}
