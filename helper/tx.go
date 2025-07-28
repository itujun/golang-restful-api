package helper

import "database/sql"

func CommitOrRollback(tx *sql.Tx) {
	err := recover()
	if err != nil {
		// Rollback the transaction if there is an error
		errRollback := tx.Rollback()
		PanicIfError(errRollback)
		panic(err)
	} else {
		// Commit the transaction if no error
		errCommit := tx.Commit()
		PanicIfError(errCommit)
	}
}

