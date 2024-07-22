package helper

import (
	"database/sql"
)

func CommiOrRollBack(tx *sql.Tx) {
	err := recover()
	if err != nil {
		errorRollBack := tx.Rollback()
		PanicIfError(errorRollBack)
		panic(err)
	} else {
		errCommit := tx.Commit()
		PanicIfError(errCommit)
	}
}
