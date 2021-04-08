package engine

import (
	"context"
	"fmt"
	"hermes-service/model"

)

type QueueWorker interface {
	Start(engine Spec) error
}


type queueWorkerImpl struct {
	worker *queues.SQSWorker
}

func NewQueueWorker(conf *model.AWSCredentials, queueConf *model.QueueListenerConfig) (QueueWorker, error) {
	worker, err := queues.NewWorker(&queues.Credentials{
		AccessKey: conf.AccessKey,
		SecretKey: conf.SecretKey,
		Region:    conf.Region,
	}, &queues.Config{
		Enabled:          queueConf.Enabled,
		LogActivity:      queueConf.LogActivity,
		SleepSeconds:     queueConf.SleepInSeconds,
		SqsName:          queueConf.SQSName,
		Workers:          queueConf.Workers,
		NumberOfMessages: queueConf.NumberOfMessages,
	})

	if err != nil {
		err = fmt.Errorf("creating Sqs worker %w", err)
		logs.Error(err)
		return nil, err
	}

	return &queueWorkerImpl{
		worker: worker,
	}, nil
}

type Spec interface {
	ProcessMessage(ctx context.Context, msg *queues.SQSMessage) error
}

func (q queueWorkerImpl) Start(engine Spec) error {
	logs.Info("Starting Worker")
	q.worker.Start(func(ctx context.Context, msg *queues.SQSMessage) error {
		err := engine.ProcessMessage(ctx, msg)
		if err != nil {
			return err
		}

		return q.worker.DeleteMessage(*msg)
	})

	return nil
}

type Engine struct {}

func (Engine) ProcessMessage(ctx context.Context, msg *queues.SQSMessage) error {
	var email model.EmailData
	err := msg.ParseJSONBody(&email)
	if err != nil {
		logs.WithContext(ctx).Error(err)
		return err
	}

	logs.WithContext(ctx).Infof("Sent email to %+v", email.To)

	return nil
}



// New engine instantiation
func New(cfg model.Config) (Spec, error) {
	return &Engine{
	}, nil
}