package entity

import (
	"context"
	"encoding/json"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
)

// return http decoder based on request

func JsonDecoder(model interface{}) httptransport.DecodeRequestFunc {
	return func(_ context.Context, req *http.Request) (interface{}, error) {
		ent := model
		if err := json.NewDecoder(req.Body).Decode(ent); err != nil {
			return nil, err
		}
		return ent, nil
	}
}

func JsonEncoder() httptransport.EncodeResponseFunc {
	return func(c context.Context, w http.ResponseWriter, response interface{}) error {
		return json.NewEncoder(w).Encode(response)
	}
}
