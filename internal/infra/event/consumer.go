package event

import (
	"context"
	"encoding/json"
	"log/slog"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/tbtec/tremligeiro/internal/dto"
)

type IConsumerService interface {
	ConsumeMessage(ctx context.Context) (*dto.Order, error)
}

type ConsumerService struct {
	QueueUrl string
	QueueArn string
	Client   *sqs.Client
}

func NewConsumerService(QueueUrl string, config aws.Config) IConsumerService {
	return &ConsumerService{
		QueueUrl: QueueUrl,
		Client:   sqs.NewFromConfig(config),
	}
}

func (consumer *ConsumerService) ConsumeMessage(ctx context.Context) (*dto.Order, error) {
	// Receive a message from the queue
	resp, err := consumer.Client.ReceiveMessage(ctx, &sqs.ReceiveMessageInput{
		QueueUrl:            &consumer.QueueUrl,
		MaxNumberOfMessages: 1,
	})
	if err != nil {
		// return nil, err
	}

	if len(resp.Messages) == 0 {
		return nil, nil // No messages available
	}

	// Deserialize the message body to Order
	var order dto.Order
	err = json.Unmarshal([]byte(*resp.Messages[0].Body), &order)
	if err != nil {
		return nil, err
	}

	slog.InfoContext(ctx, "Received message", "MessageId", *resp.Messages[0].MessageId)
	slog.InfoContext(ctx, "Received message", "body", order)

	// Delete the message from the queue
	out, delErr := consumer.Client.DeleteMessage(ctx, &sqs.DeleteMessageInput{
		QueueUrl:      &consumer.QueueUrl,
		ReceiptHandle: resp.Messages[0].ReceiptHandle,
	})
	if delErr != nil {
		slog.ErrorContext(ctx, "Error deleting message", "error", delErr)
	}
	slog.InfoContext(ctx, "Message deleted", "recepit", *&out.ResultMetadata)

	return &order, nil
}

func (consumer *ConsumerService) DeleteMessage(ctx context.Context, receiptHandle string) error {
	_, err := consumer.Client.DeleteMessage(ctx, &sqs.DeleteMessageInput{
		QueueUrl:      &consumer.QueueUrl,
		ReceiptHandle: &receiptHandle,
	})
	if err != nil {
		return err
	}

	return nil
}
