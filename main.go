package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	flags "github.com/jessevdk/go-flags"
)

type options struct {
	Profile string `short:"p" long:"profile" description:"aws profile" default:"default"`
}

var (
	Version  = "unset"
	Revision = "unset"
)

func main() {
	var opts options
	if _, err := flags.Parse(&opts); err != nil {
		return
	}

	sess := session.Must(session.NewSessionWithOptions(
		session.Options{
			Profile:           opts.Profile,
			SharedConfigState: session.SharedConfigEnable,
		},
	))
	svc := s3.New(sess)
	input := &s3.ListBucketsInput{}

	result, err := svc.ListBuckets(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			fmt.Println(err.Error())
		}
		return
	}

	for _, v := range result.Buckets {
		fmt.Println(aws.StringValue(v.Name))
	}

}
