package model

// AWSCredentials representation
type AWSCredentials struct {
	Region    string `json:"region"`
	AccessKey string `json:"access_key"`
	SecretKey string `json:"secret_key"`
}

type QueueListenerConfig struct {
	Enabled          bool   `json:"enabled"`
	SQSName          string `json:"sqs_name"`
	Workers          int    `json:"workers"`
	SleepInSeconds   int    `json:"sleep_in_seconds"`
	LogActivity      bool   `json:"log_activity"`
	NumberOfMessages int64  `json:"number_of_messages"`
}

// Config for service
// Extend this struct to parse config values from env_* files
type Config struct {
	Env         string            `json:"-"`
	Version     string            `json:"string"`
	AWS         *AWSCredentials   `json:"aws"`
	Sqs         SQSConfig         `json:"sqs"`
}

type SQSConfig struct {
	Listener QueueListenerConfig `json:"listener"`
}

