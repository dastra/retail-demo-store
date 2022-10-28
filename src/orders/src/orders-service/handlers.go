// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: MIT-0

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// Index Handler
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to the Orders Web Service")
}

// OrderIndex Handler
func OrderIndex(w http.ResponseWriter, r *http.Request) {
	setJsonContentTypeResponse(&w)
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(orders); err != nil {
		panic(err)
	}
}

// OrderIndexByUsername Handler
func OrderIndexByUsername(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	username := vars["username"]

	if err := json.NewEncoder(w).Encode(RepoFindOrdersByUsername(username)); err != nil {
		panic(err)
	}
}

// OrderShowByID Handler
func OrderShowByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	orderID := vars["orderID"]

	if err := json.NewEncoder(w).Encode(RepoFindOrderByID(orderID)); err != nil {
		panic(err)
	}
}

// OrderUpdate Func
func OrderUpdate(w http.ResponseWriter, r *http.Request) {
	var order Order
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &order); err != nil {
		setJsonContentTypeResponse(&w)
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	t := RepoUpdateOrder(order)
	setJsonContentTypeResponse(&w)
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(t); err != nil {
		panic(err)
	}
}

// OrderCreate Func
func OrderCreate(w http.ResponseWriter, r *http.Request) {
	var order Order
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &order); err != nil {
		setJsonContentTypeResponse(&w)
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	t := RepoCreateOrder(order)
	setJsonContentTypeResponse(&w)
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(t); err != nil {
		panic(err)
	}
}

func setJsonContentTypeResponse(w *http.ResponseWriter) {
	(*w).Header().Set("Content-Type", "application/json; charset=UTF-8")
}

// Sets the CORS headers
func setCorsHeaders(router *mux.Router) http.Handler {
	originsOK := handlers.AllowedOrigins([]string{"*"})
	methodsOK := handlers.AllowedMethods([]string{"GET", "POST", "PUT"})
	// Accept, Accept-Language, and Content-Language are always allowed.
	headersOK := handlers.AllowedHeaders([]string{"Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token",
		"Authorization", "X-Amzn-Trace-Id"})

	return handlers.CORS(originsOK, headersOK, methodsOK)(router)
}
