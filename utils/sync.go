package utils

import (
	"codechallenge/logger"
	"database/sql"
	"errors"
	"fmt"
)

// GetDbTx returns a function that can be used to commit or rollback a transaction
func GetDbTx(db *sql.DB) (func(err *error), *sql.Tx, error) {
	tx, txError := db.Begin()
	if txError != nil {
		return func(err *error) {}, tx, txError
	}
	return func(resError *error) {
		var err error
		if recoverError := recover(); recoverError != nil {
			rollBackError := tx.Rollback()
			if rollBackError != nil {
				logger.Logger.Errorf("error rolling back transaction: %v\n", rollBackError)
			}
			if err != nil {
				err = errors.New("transaction failed: " + err.Error())
			} else {
				err = errors.New("transaction failed")
			}
			err = fmt.Errorf("%v : %v", recoverError, err)
		}
		if err != nil || *resError != nil {
			rollBackError := tx.Rollback()
			if rollBackError != nil {
				logger.Logger.Errorf("error rolling back transaction: %v\n", rollBackError)
			}
		} else {
			err = tx.Commit()
		}

		if *resError == nil {
			*resError = err
		}
	}, tx, nil
}
