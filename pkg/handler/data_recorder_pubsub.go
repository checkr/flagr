package handler

import (
	"context"
	"encoding/json"

	"cloud.google.com/go/pubsub"
	"github.com/rexmont/flagr/pkg/config"
	"github.com/rexmont/flagr/swagger_gen/models"
	"github.com/sirupsen/logrus"
	"google.golang.org/api/option"
)

type pubsubRecorder struct {
	producer *pubsub.Client
	topic    *pubsub.Topic
}

var (
	pubsubClient = func() (*pubsub.Client, error) {
		return pubsub.NewClient(
			context.Background(),
			config.Config.RecorderPubsubProjectID,
			option.WithCredentialsFile(config.Config.RecorderPubsubKeyFile),
		)
	}
)

// NewPubsubRecorder creates a new Pubsub recorder
var NewPubsubRecorder = func() DataRecorder {
	client, err := pubsubClient()
	if err != nil {
		logrus.WithField("pubsub_error", err).Fatal("error getting pubsub client")
	}

	return &pubsubRecorder{
		producer: client,
		topic:    client.Topic(config.Config.RecorderPubsubTopicName),
	}
}

func (p *pubsubRecorder) AsyncRecord(r *models.EvalResult) {
	pr := &pubsubEvalResult{
		EvalResult: r,
	}

	payload, err := pr.Payload()
	if err != nil {
		logrus.WithField("pubsub_error", err).Error("error marshaling payload")
		return
	}

	messageFrame := pubsubMessageFrame{
		Payload:   string(payload),
		Encrypted: false,
	}

	message, err := messageFrame.encode()
	if err != nil {
		logrus.WithField("pubsub_error", err).Error("error marshaling message frame")
		return
	}

	ctx := context.Background()
	res := p.topic.Publish(ctx, &pubsub.Message{Data: message})
	if config.Config.RecorderPubsubVerbose {
		go func() {
			ctx, cancel := context.WithTimeout(ctx, config.Config.RecorderPubsubVerboseCancelTimeout)
			defer cancel()
			id, err := res.Get(ctx)
			if err != nil {
				logrus.WithFields(logrus.Fields{"pubsub_error": err, "id": id}).Error("error pushing to pubsub")
			}
		}()
	}
}

type pubsubEvalResult struct {
	*models.EvalResult
}

type pubsubMessageFrame struct {
	Payload   string `json:"payload"`
	Encrypted bool   `json:"encrypted"`
}

func (pmf *pubsubMessageFrame) encode() ([]byte, error) {
	return json.MarshalIndent(pmf, "", "  ")
}

// Payload marshals the EvalResult
func (r *pubsubEvalResult) Payload() ([]byte, error) {
	return r.EvalResult.MarshalBinary()
}
