package handlers

import (
	"encoding/json"
	"fmt"
	"jwt/models"
	"jwt/services"
	"jwt/stores"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

func UserListHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	res, err := services.ListUserService(ctx)
	if err != nil {
		render.JSON(w, r, err)
	}
	render.JSON(w, r, res)
}

func UserGetByIDHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		render.JSON(w, r, err)
	}
	res, err := services.GetBYIDService(ctx, id)
	if err != nil {
		render.JSON(w, r, err)
	}
	render.JSON(w, r, res)
}

func UserInsertHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	req := models.UserReq{}
	if decodeErr := json.NewDecoder(r.Body).Decode(&req); decodeErr != nil {
		err := fmt.Errorf("invalid Json Request")
		render.JSON(w, r, err.Error())
		return
	}
	defer r.Body.Close()
	tx, err := stores.BeginTx(ctx)
	if err != nil {
		render.JSON(w, r, err)
		return
	}
	defer stores.RollbackTx(ctx, tx)

	res, err := services.InsertService(ctx, tx, req)
	if err != nil {
		render.JSON(w, r, err.Error())
		return
	}
	if err := stores.CommitTx(ctx, tx); err != nil {
		render.JSON(w, r, err)
		return
	}
	render.JSON(w, r, res)
}

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		render.JSON(w, r, err)
		return
	}

	tx, err := stores.BeginTx(ctx)
	if err != nil {
		render.JSON(w, r, err)
		return
	}
	defer stores.RollbackTx(ctx, tx)

	err = services.DeleteService(ctx, tx, id)
	if err != nil {
		render.JSON(w, r, err)
		return
	}
	if err := stores.CommitTx(ctx, tx); err != nil {
		render.JSON(w, r, err)
		return
	}
	render.JSON(w, r, "user deleted")
}

func UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	req := models.UserReq{}
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		render.JSON(w, r, err)
		return
	}
	if decodeERR := json.NewDecoder(r.Body).Decode(&req); decodeERR != nil {
		render.JSON(w, r, err)
		return
	}
	tx, err := stores.BeginTx(ctx)
	if err != nil {
		render.JSON(w, r, err)
	}
	defer stores.RollbackTx(ctx, tx)

	res, err := services.UpdateService(ctx, tx, id, req)
	if err != nil {
		render.JSON(w, r, err)
		return
	}
	if err := stores.CommitTx(ctx, tx); err != nil {
		render.JSON(w, r, err)
	}
	render.JSON(w, r, res)
}
