package stores

import (
	"context"
	"database/sql"
	"jwt/models"
	"jwt/settings"

	"github.com/gofrs/uuid"
)

func InsertAuthToken(ctx context.Context, tx *sql.Tx, arg models.AuthToken) (*models.AuthToken, error) {
	queryStmnt := `
	INSERT INTO
	auth_tokens(
		user_id,
		token,
		expires_at
	)
	VALUES ($1, $2, $3)
	RETURNING *
	`
	row := tx.QueryRowContext(ctx, queryStmnt,
		&arg.UserID,
		&arg.Token,
		&arg.ExpiresAt,
	)
	obj := models.AuthToken{}

	if err := row.Scan(
		&obj.ID,
		&obj.UserID,
		&obj.Token,
		&obj.ExpiresAt,
		&obj.CreatedAt,
		&obj.UpdatedAt,
	); err != nil {
		return nil, err

	}
	return &obj, nil
}

func GetByToken(ctx context.Context, token uuid.UUID) (*models.AuthToken, error) {
	queryStmnt := `
	SELECT * FROM auth_tokens
	WHERE auth_token.token = $1
	`
	row := settings.DBClient.QueryRowContext(ctx, queryStmnt, token)
	obj := models.AuthToken{}
	if err := row.Scan(
		&obj.ID,
		&obj.UserID,
		&obj.Token,
		&obj.ExpiresAt,
		&obj.CreatedAt,
		&obj.UpdatedAt,
	); err != nil {
		return nil, err
	}
	return &obj, nil
}
