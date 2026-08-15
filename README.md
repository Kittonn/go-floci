# go-floci

## Getting Started

1. Start Floci
```bash
docker compose up -d
```

2. Create a KMS key

After creating the KMS key, you will receive a KeyId or ARN. Use this KeyId in the `main.go` file to encrypt data.

```bash
awslocal kms create-key
```

3. Install dependencies
```bash
go mod tidy
```

4. Run the application
```bash
go run main.go
```