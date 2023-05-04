package services

import (
	"context"
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"fmt"
	"jwt/models"
	"jwt/stores"
	"time"

	"github.com/gofrs/uuid"
)

func Login(ctx context.Context, tx *sql.Tx, req models.LoginReq) (*models.Auther, error) {
	user, err := stores.GetByEmailStore(ctx, req.Email)
	if err != nil {
		return nil, err
	}
	passhash := GetMd5(req.Password)
	if passhash != user.PasswordHash {
		return nil, fmt.Errorf("password is invalid")
	}
	tokenUid, err := uuid.DefaultGenerator.NewV4()
	if err != nil {
		return nil, err
	}
	arg := models.AuthToken{}
	arg.UserID = user.ID
	arg.Token = tokenUid
	arg.ExpiresAt = time.Now().Add(24 * time.Hour)

	authToken, err := stores.InsertAuthToken(ctx, tx, arg)
	if err != nil {
		return nil, err
	}
	auther := &models.Auther{}
	auther.Name = fmt.Sprintf("%s %s", user.FirstName, user.LastName)
	auther.Token = authToken.Token
	return auther, nil
}

func getAuther(u *models.User) *models.Auther {
	name := fmt.Sprintf("%s %s", u.FirstName, u.LastName)
	return &models.Auther{
		ID:   u.ID,
		Name: name,
	}
}

func GetAuther(ctx context.Context, token string) (*models.Auther, error) {
	tokenUUID, err := uuid.FromString(token)
	if err != nil {
		return nil, err
	}

	authToken, err := stores.GetByToken(ctx, tokenUUID)
	if err != nil {
		return nil, err
	}
	if time.Now().UTC().Unix() > authToken.ExpiresAt.UTC().Unix() {
		return nil, fmt.Errorf("token is expired")
	}
	user, err := stores.GetByIDStore(ctx, authToken.ID)
	if err != nil {
		return nil, err
	}
	auther := models.Auther{
		ID:    user.ID,
		Name:  fmt.Sprintf("%s %s", user.FirstName, user.LastName),
		Token: authToken.Token,
	}
	return &auther, nil
}

func GetAutherByToken(ctx context.Context, token uuid.UUID) (*models.Auther, error) {
	authToken, err := stores.GetByToken(ctx, token)
	if err != nil {
		return nil, err
	}
	user, err := stores.GetByIDStore(ctx, authToken.ID)
	if err != nil {
		return nil, err
	}
	return getAuther(user), nil
}

func GetMd5(input string) string {
	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(input))
	return hex.EncodeToString(hash.Sum(nil))
}
