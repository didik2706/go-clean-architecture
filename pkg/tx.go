package pkg

import (
	"database/sql"
	"fmt"
)

func CommitOrRollback(tx *sql.Tx) {
	err := recover()
	fmt.Println(err)
	if err != nil {
		errorRollback := tx.Rollback()
		PanicIfError(errorRollback)
		panic(err)
	} else {
		errorCommit := tx.Commit()
		PanicIfError(errorCommit)
	}
}