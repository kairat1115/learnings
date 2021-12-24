package context

import (
	"context"
	"fmt"
	"net/http"
)

type Store interface {
	Fetch(context.Context) (string, error)
}

func Server(store Store) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		data, err := store.Fetch(r.Context())
		if err != nil {
			return // todo: log error however you like
		}
		fmt.Fprint(rw, data)
	}
}
