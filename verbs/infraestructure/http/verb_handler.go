package http

import (
	"english-verbs/verbs/app"
	"github.com/go-chi/chi"
	"net/http"
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

// Routes returns post router with each endpoint.
func (srv *VerbHandler) Routes() http.Handler {

	r := chi.NewRouter()

	r.Get("/", srv.GetAll)

	return r
}
