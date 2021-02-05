package main

import (
	"context"

	"github.com/grafana/grafana-plugin-sdk-go/data"

	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/defaults"
)

func credsProfile(ctx context.Context, profile string, query *Query) (*data.Frame, error) {

	sharedCreds := credentials.NewSharedCredentials(defaults.SharedCredentialsFilename(), profile)
	result, err := sharedCreds.Get()
	if err != nil {
		return nil, err
	}

	frame := data.NewFrame("response")
	frame.Fields = append(frame.Fields, data.NewField("AccessKeyId", nil, []*string{&result.AccessKeyID}))
	frame.Fields = append(frame.Fields, data.NewField("SecretAccessKey", nil, []*string{&result.SecretAccessKey}))
	frame.Fields = append(frame.Fields, data.NewField("SessionToken", nil, []*string{&result.SessionToken}))

	return frame, nil
}
