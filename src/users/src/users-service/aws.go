// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: MIT-0

package main

import (
	"github.com/aws/aws-xray-sdk-go/awsplugins/ecs"
	"github.com/aws/aws-xray-sdk-go/xray"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/pinpoint"
	"github.com/aws/aws-sdk-go/service/ssm"
)

// var sess *session.Session
var sess, err = session.NewSession(&aws.Config{})

var pinpoint_app_id = os.Getenv("PINPOINT_APP_ID")
var pinpoint_client = pinpoint.New(sess)
var ssm_client = ssm.New(sess)

// Connect Stuff
func init() {
	// X-Ray
	ecs.Init()
	_ = xray.Configure(xray.Config{})
}
