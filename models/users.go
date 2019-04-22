package models

import (
	"cloud.google.com/go/datastore"
	"fmt"
)

func main() {
	ctx := context.Background()

	// Create a datastore client. In a typical application, you would create
	// a single client which is reused for every datastore operation.
	dsClient, err := datastore.NewClient(ctx, "barclay-test-env")
	if err != nil {
		// Handle error.
	}

	k := datastore.NameKey("Entity", "stringID", nil)
	e := new(Entity)
	if err := dsClient.Get(ctx, k, e); err != nil {
		// Handle error.
		fmt.Println(err)
	}

	old := e.Value
	e.Value = "Hello World!"

	if _, err := dsClient.Put(ctx, k, e); err != nil {
		// Handle error.
		fmt.Println(err)
	}

	fmt.Printf("Updated value from %q to %q\n", old, e.Value)
}