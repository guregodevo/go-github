// Copyright 2013 The go-github AUTHORS. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package github

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestRepositoriesService_ListKeys(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/repos/o/r/keys", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testFormValues(t, r, values{"page": "2"})
		fmt.Fprint(w, `[{"id":1}]`)
	})

	opt := &ListOptions{Page: 2}
	ctx := context.Background()
	keys, _, err := client.Repositories.ListKeys(ctx, "o", "r", opt)
	if err != nil {
		t.Errorf("Repositories.ListKeys returned error: %v", err)
	}

	want := []*Key{{ID: Int64(1)}}
	if !reflect.DeepEqual(keys, want) {
		t.Errorf("Repositories.ListKeys returned %+v, want %+v", keys, want)
	}
}

func TestRepositoriesService_ListKeys_invalidOwner(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	ctx := context.Background()
	_, _, err := client.Repositories.ListKeys(ctx, "%", "%", nil)
	testURLParseError(t, err)
}

func TestRepositoriesService_GetKey(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/repos/o/r/keys/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{"id":1}`)
	})

	ctx := context.Background()
	key, _, err := client.Repositories.GetKey(ctx, "o", "r", 1)
	if err != nil {
		t.Errorf("Repositories.GetKey returned error: %v", err)
	}

	want := &Key{ID: Int64(1)}
	if !reflect.DeepEqual(key, want) {
		t.Errorf("Repositories.GetKey returned %+v, want %+v", key, want)
	}
}

func TestRepositoriesService_GetKey_invalidOwner(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	ctx := context.Background()
	_, _, err := client.Repositories.GetKey(ctx, "%", "%", 1)
	testURLParseError(t, err)
}

func TestRepositoriesService_CreateKey(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	input := &Key{Key: String("k"), Title: String("t")}

	mux.HandleFunc("/repos/o/r/keys", func(w http.ResponseWriter, r *http.Request) {
		v := new(Key)
		json.NewDecoder(r.Body).Decode(v)

		testMethod(t, r, "POST")
		if !reflect.DeepEqual(v, input) {
			t.Errorf("Request body = %+v, want %+v", v, input)
		}

		fmt.Fprint(w, `{"id":1}`)
	})

	ctx := context.Background()
	key, _, err := client.Repositories.CreateKey(ctx, "o", "r", input)
	if err != nil {
		t.Errorf("Repositories.GetKey returned error: %v", err)
	}

	want := &Key{ID: Int64(1)}
	if !reflect.DeepEqual(key, want) {
		t.Errorf("Repositories.GetKey returned %+v, want %+v", key, want)
	}
}

func TestRepositoriesService_CreateKey_invalidOwner(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	ctx := context.Background()
	_, _, err := client.Repositories.CreateKey(ctx, "%", "%", nil)
	testURLParseError(t, err)
}

func TestRepositoriesService_DeleteKey(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/repos/o/r/keys/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
	})

	ctx := context.Background()
	_, err := client.Repositories.DeleteKey(ctx, "o", "r", 1)
	if err != nil {
		t.Errorf("Repositories.DeleteKey returned error: %v", err)
	}
}

func TestRepositoriesService_DeleteKey_invalidOwner(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	ctx := context.Background()
	_, err := client.Repositories.DeleteKey(ctx, "%", "%", 1)
	testURLParseError(t, err)
}
