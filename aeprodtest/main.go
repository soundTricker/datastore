package aeprodtest

import (
	"net/http"

	"google.golang.org/appengine/v2"
	originaldatastore "google.golang.org/appengine/v2/datastore"
	"google.golang.org/appengine/v2/log"

	"go.mercari.io/datastore/aedatastore"
)

func init() {
	// Put Entity via original AppEngine Datastore
	http.HandleFunc("/api/test1", func(w http.ResponseWriter, r *http.Request) {
		ctx := appengine.NewContext(r)

		type Inner struct {
			A string
			B string
		}

		type Data struct {
			Slice []Inner
		}

		key := originaldatastore.NewIncompleteKey(ctx, "AETest", nil)
		_, err := originaldatastore.Put(ctx, key, &Data{
			Slice: []Inner{
				{A: "A1", B: "B1"},
				{A: "A2", B: "B2"},
				{A: "A3", B: "B3"},
			},
		})
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Warningf(ctx, "error: %v", err)
			return
		}

		w.WriteHeader(200)
	})

	// Put Entity via datastore
	http.HandleFunc("/api/test2", func(w http.ResponseWriter, r *http.Request) {
		ctx := appengine.NewContext(r)

		ds, err := aedatastore.FromContext(ctx)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Warningf(ctx, "error: %v", err)
			return
		}

		type Inner struct {
			A string
			B string
		}

		type Data struct {
			Slice []Inner
		}

		key := ds.IncompleteKey("AETest", nil)
		_, err = ds.Put(ctx, key, &Data{
			Slice: []Inner{
				{A: "A1", B: "B1"},
				{A: "A2", B: "B2"},
				{A: "A3", B: "B3"},
			},
		})
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Warningf(ctx, "error: %v", err)
			return
		}

		w.WriteHeader(200)
	})

	// Put Entity via datastore
	http.HandleFunc("/api/test3", func(w http.ResponseWriter, r *http.Request) {
		ctx := appengine.NewContext(r)

		ds, err := aedatastore.FromContext(ctx)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Warningf(ctx, "error: %v", err)
			return
		}

		type Inner struct {
			A string
			B string
		}

		type Data struct {
			Slice []Inner `datastore:",flatten"`
		}

		key := ds.IncompleteKey("AETest", nil)
		_, err = ds.Put(ctx, key, &Data{
			Slice: []Inner{
				{A: "A1", B: "B1"},
				{A: "A2", B: "B2"},
				{A: "A3", B: "B3"},
			},
		})
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Warningf(ctx, "error: %v", err)
			return
		}

		w.WriteHeader(200)
	})
}
