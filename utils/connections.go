package utils

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
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

func LoadSQS(sqsAddress string) (aws.Config, error) {
	// Load the AWS configuration with a custom endpoint for LocalStack
	return config.LoadDefaultConfig(context.TODO(),
		config.WithRegion("us-east-1"), // LocalStack often assumes us-east-1 region by default
		config.WithEndpointResolver(aws.EndpointResolverFunc(
			func(service, region string) (aws.Endpoint, error) {
				if service == sqs.ServiceID {
					return aws.Endpoint{
						URL:           sqsAddress, // LocalStack's SQS URL
						SigningRegion: "us-east-1",
					}, nil
				}
				return aws.Endpoint{}, fmt.Errorf("unknown endpoint requested")
			},
		)),
	)
}
