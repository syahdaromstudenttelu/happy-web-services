package helper

import "database/sql"

func CommitOrRollback(tx *sql.Tx) {
	err := recover()
	if err != nil {
		errorRollback := tx.Rollback()
		DoPanicIfError(errorRollback)
		panic(err)
	} else {
		errorCommit := tx.Commit()
		DoPanicIfError(errorCommit)
	}
}
