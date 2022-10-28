// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: MIT-0

package main

import (
	"log"
	"net/http"
)

func main() {
	router := NewRouter()
	handler := setCorsHeaders(router)

	log.Fatal(http.ListenAndServe(":80", handler))
}
