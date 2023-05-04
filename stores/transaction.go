package stores

import (
	"context"
	"database/sql"
	"jwt/settings"
	"log"
)

func BeginTx(ctx context.Context) (*sql.Tx, error) {
	tx, err := settings.DBClient.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return nil, err
	}
	log.Println("Beign Transaction")

	return tx, nil

}

func CommitTx(ctx context.Context, tx *sql.Tx) error {
	err := tx.Commit()
	if err != nil {
		return err
	}
	log.Println("Commit Transaction")

	return nil
}

func RollbackTx(ctx context.Context, tx *sql.Tx) error {
	err := tx.Rollback()
	if err != nil {
		log.Println(err.Error())
		return nil
	}
	log.Println("Rollback Transaction")

	return nil
}
