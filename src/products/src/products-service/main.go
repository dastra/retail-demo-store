// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: MIT-0

package main

import (
	"github.com/gorilla/handlers"
	"log"
	"net/http"
)

func main() {
	router := NewRouter()

	headersOK := handlers.AllowedHeaders([]string{"X-Amzn-Trace-Id"})
	originsOK := handlers.AllowedOrigins([]string{"*"})
	methodsOK := handlers.AllowedMethods([]string{"GET", "POST", "PUT"})

	log.Fatal(http.ListenAndServe(":80", handlers.CORS(originsOK, headersOK, methodsOK)(router)))
}
