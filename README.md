# lambda-go

Lambda em Go com LocalStack.

## Pré-requisitos

- Docker
- Go 1.22+
- AWS CLI

## Setup inicial (apenas uma vez)

Adicionar usuário ao grupo docker:
```bash
sudo usermod -aG docker $USER
newgrp docker
```
## aws configure

```bash
AWS Access Key ID [None]: fakeAccessKeyId
AWS Secret Access Key [None]: fakeSecretAccessKey
Default region name [us-east-1]: us-east-1
Default output format [None]: json
```
## aws configure --profile localstack

```bash
AWS Access Key ID [None]: nome_perfil_novo
AWS Secret Access Key [None]: senha_perfil_novo
Default region name [None]: us-east-1
Default output format [None]: json
```

## Subir o LocalStack
```bash
docker compose up -d
```

## Compilar o binário

```bash
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o bootstrap ./internal/lambda/
zip function.zip bootstrap
```

## Criar função no localStack
```bash
aws --endpoint-url=http://localhost:4566 lambda create-function \
  --function-name hello-lambda \
  --runtime provided.al2023 \
  --handler bootstrap \
  --role arn:aws:iam::000000000000:role/lambda-role \
  --zip-file fileb://function.zip \
  --region us-east-1
```

## Deletar função no localStack
```bash
aws --endpoint-url=http://localhost:4566 lambda delete-function \
  --function-name hello-lambda \
  --region us-east-1
```