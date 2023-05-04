package services

import (
	"context"
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"fmt"
	"jwt/models"
	"jwt/stores"
)

func ListUserService(ctx context.Context) ([]models.User, error) {
	res, err := stores.ListStores(ctx)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func GetBYIDService(ctx context.Context, id int64) (*models.User, error) {
	return stores.GetByIDStore(ctx, id)
}

func InsertService(ctx context.Context, tx *sql.Tx, req models.UserReq) (*models.User, error) {

	user := models.User{}
	count, err := stores.GetCountStore(ctx)
	if err != nil {
		return nil, err
	}

	user.FirstName = req.FirstName
	user.LastName = req.LastName
	user.Email = req.Email
	user.Phone = req.Phone
	user.PasswordHash = getMd5(req.Password)
	user.Code = fmt.Sprintf("U%05d", count+1)
	fmt.Println(user.Code)

	obj, err := stores.InsertStore(ctx, user)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

func DeleteService(ctx context.Context, tx *sql.Tx, id int64) error {

	err := stores.DeleteUserStore(ctx, tx, id)
	if err != nil {
		return err
	}
	return nil
}

func UpdateService(ctx context.Context, tx *sql.Tx, id int64, req models.UserReq) (*models.User, error) {
	user, err := stores.GetByIDStore(ctx, id)
	if err != nil {
		return nil, err
	}
	user.FirstName = req.FirstName
	user.LastName = req.LastName
	user.Email = req.Email
	user.Phone = req.Phone

	err = stores.UpdateUserStore(ctx, tx, *user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func getMd5(input string) string {
	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(input))
	return hex.EncodeToString(hash.Sum(nil))
}
