package config

import (
	"fmt"

	"github.com/boltdb/bolt"
)

// SaveEmail saves the EML string in the database with the given message ID as the key.
func SaveEmail(messageID, emlString, dbName string) error {
	// Open the BoltDB database.
	db, err := bolt.Open("emails.db", 0600, nil)
	if err != nil {
		return fmt.Errorf("failed to open database: %v", err)
	}
	defer db.Close()

	// Start a read-only transaction.
	err = db.View(func(tx *bolt.Tx) error {
		// Retrieve the bucket.
		bucket := tx.Bucket([]byte(dbName))
		if bucket == nil {
			// If the bucket does not exist, return nil.
			return nil
		}

		// Check if the message ID already exists in the bucket.
		existingValue := bucket.Get([]byte(messageID))
		if existingValue != nil {
			return fmt.Errorf("message ID %s already exists in the database", messageID)
		}

		return nil
	})
	if err != nil {
		return fmt.Errorf("failed: %v", err)
	}

	// Start a new read-write transaction to update the database.
	err = db.Update(func(tx *bolt.Tx) error {
		// Retrieve or create the bucket.
		bucket, err := tx.CreateBucketIfNotExists([]byte(dbName))
		if err != nil {
			return err
		}

		// Add the EML string to the bucket with the message ID as the key.
		err = bucket.Put([]byte(messageID), []byte(emlString))
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return fmt.Errorf("failed to save email: %v", err)
	}

	return nil
}

// RetrieveEmails retrieves all saved email objects from the specified bucket in the database.
func RetrieveEmails(dbName string) ([]string, error) {
	// Open the BoltDB database.
	db, err := bolt.Open("emails.db", 0600, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %v", err)
	}
	defer db.Close()

	var emails []string

	// Read-only transaction to retrieve emails from the specified bucket.
	err = db.View(func(tx *bolt.Tx) error {
		// Retrieve the bucket.
		bucket := tx.Bucket([]byte(dbName))
		if bucket == nil {
			// If the bucket does not exist, return nil.
			return nil
		}

		// Iterate over all key-value pairs in the bucket.
		if err := bucket.ForEach(func(k, v []byte) error {
			// Append the email value to the emails slice.
			emails = append(emails, string(v))
			return nil
		}); err != nil {
			return fmt.Errorf("failed to iterate over bucket: %v", err)
		}

		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve emails: %v", err)
	}

	return emails, nil
}

// RetrieveEmailsPaginated retrieves emails in a paginated fashion from the specified bucket.
func RetrieveEmailsPaginated(dbName string, pageNum int, pageSize int) ([]string, error) {
	// Open the BoltDB database.
	db, err := bolt.Open("emails.db", 0600, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %v", err)
	}
	defer db.Close()

	var emails []string

	// Calculate the starting index and the ending index for the items to be retrieved
	startIndex := (pageNum - 1) * pageSize
	endIndex := startIndex + pageSize

	// Initialize a counter to keep track of the current item index
	counter := 0

	err = db.View(func(tx *bolt.Tx) error {
		// Retrieve the bucket.
		bucket := tx.Bucket([]byte(dbName))
		if bucket == nil {
			return fmt.Errorf("bucket %s not found", dbName)
		}

		// Iterate over all items in the bucket.
		return bucket.ForEach(func(k, v []byte) error {
			if counter >= startIndex && counter < endIndex {
				// If the current item is within the desired range, add it to the results
				emails = append(emails, string(v))
			}
			// Increment the counter regardless of whether the item was added to the results
			counter++
			// If the counter has reached the end index, we can stop iterating
			if counter >= endIndex {
				return fmt.Errorf("stop iteration")
			}
			return nil
		})
	})

	// It's normal to encounter the "stop iteration" error; it's our signal to stop early.
	// We should not return it as an actual error to the caller.
	if err != nil && err.Error() == "stop iteration" {
		err = nil
	}

	if err != nil {
		return nil, fmt.Errorf("failed to retrieve emails: %v", err)
	}

	return emails, nil
}
