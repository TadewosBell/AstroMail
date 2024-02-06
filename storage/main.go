package storage

import (
	"fmt"
	"log"

	"github.com/tidwall/buntdb"
)

func SaveCredentials(domain string, aws_id string, aws_secret string) {
	// Open the data.db file. It will be created if it doesn't exist.
	db, err := buntdb.Open("data.db")
	if err != nil {
		log.Fatal(err)
	}

	db_err := db.Update(func(tx *buntdb.Tx) error {
		_, _, err := tx.Set("aws_id", aws_id, nil)
		_, _, err = tx.Set("aws_secret", aws_secret, nil)
		_, _, err = tx.Set("domain", domain, nil)
		return err
	})

	if db_err != nil {
		log.Fatal(db_err)
	}

	db_err = db.View(func(tx *buntdb.Tx) error {
		val, err := tx.Get("aws_id")
		if err != nil {
			return err
		}
		fmt.Printf("value is %s\n", val)
		return nil
	})

	defer db.Close()

}
