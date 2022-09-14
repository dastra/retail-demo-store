// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: MIT-0

package main

import (
	"github.com/gorilla/handlers"
	"log"
	"net/http"
	"os"
)

func main() {
	port, exists := os.LookupEnv("PORT")
	if !exists {
		port = "80"
	}

	router := NewRouter()

	headersOK := handlers.AllowedHeaders([]string{"Content-Type", "X-Amzn-Trace-Id"})
	originsOK := handlers.AllowedOrigins([]string{os.Getenv("WEB_ROOT_URL")})
	methodsOK := handlers.AllowedMethods([]string{"GET", "POST", "PUT"})

	log.Fatal(http.ListenAndServe(":"+port, handlers.CORS(originsOK, headersOK, methodsOK)(router)))
}
