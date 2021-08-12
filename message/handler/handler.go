package handler

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	log "github.com/sirupsen/logrus"
)

// Handler accepts SNS requests
func Handler(ctx context.Context, snsEvent events.SNSEvent) error {
	for _, message := range snsEvent.Records {
		log.Info(message)
	}
	return nil
}
