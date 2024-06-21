package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func main() {
	flag.Parse()
	awsSession := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	s3Client := s3.New(awsSession, &aws.Config{
		Region: aws.String("ap-northeast-1"),
	})
	req, _ := s3Client.PutObjectRequest(&s3.PutObjectInput{
		Bucket: aws.String(flag.Arg(0)),
		Key:    aws.String(flag.Arg(1)),
	})
	url, err := req.Presign(5 * time.Minute)
	if err != nil {
		panic(err)
	}

	fmt.Println(url)
}
