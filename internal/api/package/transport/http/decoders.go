package http

import (
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"gymshark-test/internal/api/package/transport"
	"gymshark-test/pkg/errors"
	"net/http"
	"strconv"
)

func emptyDecoder(_ context.Context, r *http.Request) (interface{}, error) {
	return nil, nil
}

func storePackageDecoder(_ context.Context, r *http.Request) (interface{}, error) {
	request := new(transport.StorePackageRequest)

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, errors.InvalidCharacter.SetDevMessage(err.Error())
	}

	return request, nil
}

func updatePackageDecoder(_ context.Context, r *http.Request) (interface{}, error) {
	request := new(transport.UpdatePackageRequest)

	vars := mux.Vars(r)
	idStr, ok := vars["id"]
	if !ok {
		return nil, errors.NotFound
	}

	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		return nil, errors.InvalidCharacter.SetDevMessage(err.Error())
	}

	request.ID = uint(id)

	if err = json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, errors.InvalidCharacter.SetDevMessage(err.Error())
	}

	return request, nil
}

func deletePackageDecoder(_ context.Context, r *http.Request) (interface{}, error) {
	request := new(transport.DeletePackageRequest)

	vars := mux.Vars(r)
	idStr, ok := vars["id"]
	if !ok {
		return nil, errors.NotFound
	}

	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		return nil, errors.InvalidCharacter.SetDevMessage(err.Error())
	}

	request.ID = uint(id)

	return request, nil
}

func packDecoder(_ context.Context, r *http.Request) (interface{}, error) {
	request := new(transport.PackRequest)

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, errors.InvalidCharacter.SetDevMessage(err.Error())
	}

	return request, nil
}
