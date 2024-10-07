package utils

import (
	config2 "codechallenge/config"
	"context"
	"database/sql"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	_ "github.com/lib/pq"
	"log"
	"time"
)

func PostgresConnection(host, port, user, pass, database, sslmode string, maxOpenConns, maxIdleConns int, timeout time.Duration) (*sql.DB, error) {
	connString := PostgresURI(host, port, user, pass, database, sslmode)
	log.Println("postgres options -> " + connString)
	conn, err := sql.Open("postgres", connString)
	if err != nil {
		return nil, fmt.Errorf("error in openning postgres connection: %w", err)
	}

	conn.SetMaxOpenConns(maxOpenConns)
	conn.SetMaxIdleConns(maxIdleConns)

	dbContext, _ := context.WithTimeout(context.Background(), timeout)
	err = conn.PingContext(dbContext)
	if err != nil {
		return nil, fmt.Errorf("error in pinging postgres database: %w", err)
	}
	return conn, nil
}

func PostgresURI(host, port, user, pass, database, sslmode string) string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", user, pass, host, port, database, sslmode)
}

func LoadSQS(sqsAddress string) (*sqs.Client, error) {
	// Load the AWS configuration with a custom endpoint for LocalStack
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-west-2"))
	if err != nil {
		log.Fatal(err)
	}
	cfg.Credentials = aws.CredentialsProviderFunc(func(ctx context.Context) (aws.Credentials, error) {
		return aws.Credentials{
			AccessKeyID:     config2.AppConfig.Queue.AccessKey,
			SecretAccessKey: config2.AppConfig.Queue.SecretKey,
		}, nil
	})

	// Create SQS client
	client := sqs.NewFromConfig(cfg, func(o *sqs.Options) {
		o.BaseEndpoint = aws.String(sqsAddress) // Ensure we use localstack URL
	})

	// Create the bucket in localstack
	_, err = client.CreateQueue(context.Background(), &sqs.CreateQueueInput{
		QueueName: aws.String(config2.AppConfig.Queue.QueueAddress),
	})
	if err != nil {
		log.Fatalf("unable to create bucket, %v", err)
	}

	return client, nil
}

func LoadS3(s3Address string) (*s3.Client, error) {
	// Load the Shared AWS Configuration (~/.aws/config)
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-west-2"))
	if err != nil {
		log.Fatal(err)
	}

	cfg.Credentials = aws.CredentialsProviderFunc(func(ctx context.Context) (aws.Credentials, error) {
		return aws.Credentials{
			AccessKeyID:     config2.AppConfig.Storage.AccessKey,
			SecretAccessKey: config2.AppConfig.Storage.SecretKey,
		}, nil
	})
	cfg.BaseEndpoint = aws.String(s3Address)
	fmt.Printf("AWS S3 options -> endpoint: %s, access key: %s, secret key: %s\n", s3Address, config2.AppConfig.Storage.AccessKey, config2.AppConfig.Storage.SecretKey)

	// Create an Amazon S3 service client
	client := s3.NewFromConfig(cfg, func(o *s3.Options) {
		o.BaseEndpoint = aws.String(s3Address) // Ensure we use localstack URL
		o.UsePathStyle = true                  // Force path-style addressing
	})

	// Create the bucket in localstack
	_, err = client.CreateBucket(context.Background(), &s3.CreateBucketInput{
		Bucket: aws.String(config2.AppConfig.Storage.Bucket),
	})
	if err != nil {
		log.Fatalf("unable to create bucket, %v", err)
	}

	return client, nil
}
