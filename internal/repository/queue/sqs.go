package queue

import (
	"codechallenge/internal/repository"
	"codechallenge/internal/service/service_models"
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
)

type queue struct {
	queueUrl string
	client   *sqs.Client
}

func NewQueue(queueUrl string, client *sqs.Client) repository.QueueRepository {
	return &queue{
		queueUrl: queueUrl,
		client:   client,
	}
}

func (s *queue) PushTodoItem(ctx context.Context, message service_models.TodoItem) error {
	body, err := json.Marshal(message)
	if err != nil {
		return err
	}

	// Create the SendMessageInput object
	input := &sqs.SendMessageInput{
		MessageBody: aws.String(string(body)),
		QueueUrl:    aws.String(s.queueUrl),
	}

	// Send the message to the SQS queue
	resp, err := s.client.SendMessage(ctx, input)
	if err != nil {
		return err
	}
	fmt.Printf("Message sent successfully with ID: %s\n", *resp.MessageId)
	return nil
}
