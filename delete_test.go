// Copyright 2016 Tim Shannon. All rights reserved.
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

package bolthold_test

import (
	"testing"
	"time"

	"github.com/timshannon/bolthold"
)

func TestDelete(t *testing.T) {
	testWrap(t, func(store *bolthold.Store, t *testing.T) {
		key := "testKey"
		data := &ItemTest{
			Name:    "Test Name",
			Created: time.Now(),
		}

		err := store.Insert(key, data)
		if err != nil {
			t.Fatalf("Error inserting data for delete test: %s", err)
		}

		result := &ItemTest{}

		err = store.Delete(key, result)
		if err != nil {
			t.Fatalf("Error deleting data from bolthold: %s", err)
		}

		err = store.Get(key, result)
		if err != bolthold.ErrNotFound {
			t.Fatalf("Data was not deleted from bolthold")
		}

	})
}