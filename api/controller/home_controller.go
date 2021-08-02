package controller

import (
	"net/http"

	"backend_go/api/responses"
)

func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "Hi, TelkomSigma, Welcome")
}
