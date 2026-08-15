package main

import (
	"context"
	"encoding/base64"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/kms"
)

type Environment string

const (
	Local      Environment = "local"
	Production Environment = "production"
)

func NewAWSConfig(ctx context.Context, env Environment, region, endpoint string) (aws.Config, error) {
	switch env {
	case Local:
		cfg, err := config.LoadDefaultConfig(ctx,
			config.WithRegion(region),
			config.WithBaseEndpoint(endpoint),
		)
		if err != nil {
			return aws.Config{}, fmt.Errorf("failed to load local AWS config: %w", err)
		}
		return cfg, nil

	case Production:
		cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion(region))
		if err != nil {
			return aws.Config{}, fmt.Errorf("failed to load production AWS config: %w", err)
		}
		return cfg, nil

	default:
		return aws.Config{}, fmt.Errorf("invalid environment: %s", env)
	}
}

func main() {
	ctx := context.Background()

	awsConfig, err := NewAWSConfig(
		ctx,
		Local,
		"ap-southeast-1",
		"http://localhost:4566",
	)
	if err != nil {
		log.Fatalf("failed to create AWS config: %v", err)
	}

	kmsClient := kms.NewFromConfig(awsConfig)

	keyID := "ae08f2eb-498f-4820-b427-e584c39bb8ac"

	enc, err := kmsClient.Encrypt(ctx, &kms.EncryptInput{
		KeyId:     aws.String(keyID),
		Plaintext: []byte("Hello, World!"),
	})
	if err != nil {
		log.Fatalf("failed to encrypt data: %v", err)
	}

	log.Printf("Encrypted data: %s", base64.StdEncoding.EncodeToString(enc.CiphertextBlob))
}
