package handlers

import (
	"encoding/json"
	"fmt"
	"jwt/models"
	"jwt/services"
	"jwt/stores"
	"net/http"

	"github.com/go-chi/render"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	req := models.LoginReq{}
	if decodeErr := json.NewDecoder(r.Body).Decode(&req); decodeErr != nil {
		err := fmt.Errorf("invalid json request")
		render.JSON(w, r, err.Error())
	}
	tx, err := stores.BeginTx(ctx)
	if err != nil {
		render.JSON(w, r, err.Error())
		return
	}
	defer stores.RollbackTx(ctx, tx)

	user, err := services.Login(ctx, tx, req)
	if err != nil {
		render.JSON(w, r, err.Error())
		return
	}
	if err := stores.CommitTx(ctx, tx); err != nil {
		render.JSON(w, r, err.Error())
		return
	}
	render.JSON(w, r, user)
}
