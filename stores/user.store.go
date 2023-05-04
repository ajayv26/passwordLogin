package stores

import (
	"context"
	"database/sql"
	"jwt/models"
	"jwt/settings"

	_ "github.com/lib/pq"
)

func ListStores(ctx context.Context) ([]models.User, error) {
	queryStmnt := `SELECT * FROM users`

	rows, err := settings.DBClient.Query(queryStmnt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := []models.User{}
	for rows.Next() {
		user := models.User{}
		if err := rows.Scan(
			&user.ID,
			&user.Code,
			&user.FirstName,
			&user.LastName,
			&user.Email,
			&user.Phone,
			&user.PasswordHash,
			&user.IsAdmin,
			&user.IsArchived,
			&user.CreatedAt,
			&user.UpdatedAt,
		); err != nil {
			return nil, err
		}

		users = append(users, user)
	}
	return users, nil
}

func GetByIDStore(ctx context.Context, id int64) (*models.User, error) {
	queryStmnt := `SELECT * FROM users WHERE users.id=$1`

	arg := &models.User{}
	err := settings.DBClient.QueryRowContext(ctx, queryStmnt, id).Scan(
		&arg.ID,
		&arg.Code,
		&arg.FirstName,
		&arg.LastName,
		&arg.Email,
		&arg.Phone,
		&arg.PasswordHash,
		&arg.IsAdmin,
		&arg.IsArchived,
		&arg.CreatedAt,
		&arg.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return arg, nil
}
func GetByEmailStore(ctx context.Context, email string) (*models.User, error) {
	queryStmnt := `SELECT * FROM users WHERE users.email=$1`

	arg := &models.User{}
	err := settings.DBClient.QueryRowContext(ctx, queryStmnt, email).Scan(
		&arg.ID,
		&arg.Code,
		&arg.FirstName,
		&arg.LastName,
		&arg.Email,
		&arg.Phone,
		&arg.PasswordHash,
		&arg.IsAdmin,
		&arg.IsArchived,
		&arg.CreatedAt,
		&arg.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return arg, nil
}

func GetCountStore(ctx context.Context) (int64, error) {
	queryStmnt := ` SELECT COUNT(id) FROM users`
	var count int64
	err := settings.DBClient.QueryRowContext(ctx, queryStmnt).Scan(
		&count,
	)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func InsertStore(ctx context.Context, arg models.User) (*models.User, error) {
	queryStmt := `
	INSERT INTO
	users(
		code,
		first_name,
		last_name,
		email,
		phone,
		password_hash,
		is_admin,
		is_archived
	)VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	RETURNING *
	`
	row := settings.DBClient.QueryRowContext(ctx, queryStmt,
		&arg.Code,
		&arg.FirstName,
		&arg.LastName,
		&arg.Email,
		&arg.Phone,
		&arg.PasswordHash,
		&arg.IsAdmin,
		&arg.IsArchived,
	)
	user := &models.User{}

	err := row.Scan(
		&user.ID,
		&user.Code,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Phone,
		&user.PasswordHash,
		&user.IsAdmin,
		&user.IsArchived,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return user, nil

}

func DeleteUserStore(ctx context.Context, tx *sql.Tx, id int64) error {
	querStmnt := `
    DELETE FROM users WHERE id=$1  
   `
	_, err := tx.ExecContext(ctx, querStmnt, id)
	if err != nil {
		return err
	}
	return nil
}

func UpdateUserStore(ctx context.Context, tx *sql.Tx, arg models.User) error {
	queryStmt := `
	 Update users 
	 SET
	   first_name=$1,
	   last_name=$2,
	   email=$3,
	   phone=$4,
	   password_hash=$5
	   WHERE id=$6
	`
	_, err := tx.ExecContext(ctx, queryStmt,
		&arg.FirstName,
		&arg.LastName,
		&arg.Email,
		&arg.Phone,
		&arg.PasswordHash,
		&arg.ID,
	)
	if err != nil {
		return err
	}
	return nil
}
