// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: MIT-0

package main

import (
	"github.com/aws/aws-xray-sdk-go/awsplugins/ecs"
	"github.com/aws/aws-xray-sdk-go/xray"
)

// Initialize clients
func init() {
	// X-Ray
	ecs.Init()
	_ = xray.Configure(xray.Config{})
}
