package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/byrnedo/typesafe-config/parse"
)

//A force is what one uses to connect to salesforce.
type force struct {
	Env   Env
	URL   string
	Token string
}

type Env struct {
	Path string
	Env  string
}

func getSalesforceLogin(env Env) (force, error) {
	svc := s3.New(session.New(), &aws.Config{Region: aws.String("eu-west-1")})

	d := s3manager.NewDownloaderWithClient(svc, func(d *s3manager.Downloader) {
		d.Concurrency = 1
	})

	w := &aws.WriteAtBuffer{}
	_, err := d.Download(w, &s3.GetObjectInput{
		Bucket: aws.String("membership-private"),
		Key:    aws.String(env.Path + "/membership-keys.conf"),
	})

	if err != nil {
		return force{}, err
	}

	c, err := parse.ParseBytes(w.Bytes())
	conf := c.GetConfig()
	if err != nil {
		return force{}, err
	}

	token, err := conf.GetString("touchpoint.backend.environments." + env.Env + ".salesforce.access_token")
	if err != nil {
		return force{}, err
	}
	url, err := conf.GetString("touchpoint.backend.environments." + env.Env + ".salesforce.instance_url")
	if err != nil {
		return force{}, err
	}
	return force{env, url, token}, nil
}
