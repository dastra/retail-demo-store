// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: MIT-0

package main

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/lambda"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"io"
	"io/ioutil"
	"net/http"
)

// Index Handler
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to the Carts Web Service")
}

// CartIndex Handler
func CartIndex(w http.ResponseWriter, r *http.Request) {
	setJsonContentTypeResponse(&w)
	w.WriteHeader(http.StatusOK)

	var values []Cart
	for _, value := range carts {
		values = append(values, value)
	}

	if err := json.NewEncoder(w).Encode(values); err != nil {
		panic(err)
	}
}

// CartShowByID Handler
func CartShowByID(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	cartID := vars["cartID"]

	if err := json.NewEncoder(w).Encode(RepoFindCartByID(cartID)); err != nil {
		panic(err)
	}
}

// CartUpdate Func
func CartUpdate(w http.ResponseWriter, r *http.Request) {

	var cart Cart
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &cart); err != nil {
		setJsonContentTypeResponse(&w)
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	vars := mux.Vars(r)
	cartID := vars["cartID"]

	t := RepoUpdateCart(cartID, cart)

	setJsonContentTypeResponse(&w)
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(t); err != nil {
		panic(err)
	}
}

// CartCreate Func
func CartCreate(w http.ResponseWriter, r *http.Request) {
	var cart Cart
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &cart); err != nil {
		setJsonContentTypeResponse(&w)
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	t := RepoCreateCart(cart)

	setJsonContentTypeResponse(&w)
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(t); err != nil {
		panic(err)
	}
}

// Sign a payload for Amazon Pay - delegates to a Lambda function for doing this.
func SignAmazonPayPayload(w http.ResponseWriter, r *http.Request) {
	awsSession := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	client := lambda.New(awsSession)

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	var requestBody map[string]interface{}
	json.Unmarshal(body, &requestBody)

	result, err := client.Invoke(&lambda.InvokeInput{FunctionName: aws.String("AmazonPaySigningLambda"), Payload: body})
	if err != nil {
		panic(err)
	}

	var responsePayload map[string]interface{}
	json.Unmarshal(result.Payload, &responsePayload)

	setJsonContentTypeResponse(&w)
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(responsePayload); err != nil {
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
