package secrets

import (
	"context"
	"os"

	"github.com/swethabhageerath/aws/lib/models"
	"github.com/swethabhageerath/logging/lib/constants"
	mod "github.com/swethabhageerath/logging/lib/models"
	"github.com/swethabhageerath/logging/lib/writers"

	"github.com/aws/aws-sdk-go-v2/aws"
	cfg "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"github.com/pkg/errors"
	er "github.com/swethabhageerath/aws/lib/errors"
)

type Secrets struct{}

func (s Secrets) GetValue(ctx context.Context, secretName string, out chan<- models.SecretsManagerResponse) {
	client, err := s.getClient(ctx)

	if err != nil {
		er := errors.Wrap(err, er.NewErrRetrievingAwsSecretsManagerClient().Message)
		s.log(er)
		out <- models.SecretsManagerResponse{
			Data:  nil,
			Error: er,
		}
	}

	input := s.getRequestInput(secretName)

	output, err := client.GetSecretValue(ctx, input)

	if err != nil {
		er := errors.Wrap(err, er.NewErrRetrievingSecretFromAwsSecretsManager().Message)
		s.log(er)
		out <- models.SecretsManagerResponse{
			Data:  nil,
			Error: er,
		}
	}

	r := models.SecretsManagerResponse{
		Data:  output,
		Error: nil,
	}

	out <- r
}

func (s Secrets) getConfig(ctx context.Context) (aws.Config, error) {
	region := s.getRegion(KEY_AWS_SECRETSMANAGER_REGION)

	if region == "" {
		er := errors.New(er.NewErrRegionNotSpecifiedForSecretsManager().Message)
		s.log(er)
		return aws.Config{}, er
	}

	c, err := cfg.LoadDefaultConfig(ctx, cfg.WithRegion(region))
	if err != nil {
		er := errors.Wrap(err, er.NewErrLoadingConfiguringForAwsSecretsManager().Message)
		s.log(er)
		return aws.Config{}, er
	}

	return c, nil
}

func (s Secrets) getRequestInput(secretName string) *secretsmanager.GetSecretValueInput {
	return &secretsmanager.GetSecretValueInput{
		SecretId:     aws.String(secretName),
		VersionStage: aws.String(KEY_AWS_SECRETSMANAGER_VERSION_STAGE),
	}
}

func (s Secrets) getClient(ctx context.Context) (*secretsmanager.Client, error) {
	config, err := s.getConfig(ctx)
	if err != nil {
		return nil, err
	}
	return secretsmanager.NewFromConfig(config), nil
}

func (s Secrets) getRegion(key string) string {
	return os.Getenv(key)
}

func (s Secrets) log(err error) {
	l := mod.New(mod.WithMandatoryFields("AWSSecretsManager", "bmoola", constants.ERROR),
		mod.WithRequestId("abc123"), mod.WithStackTrace(err))
	_, e := l.Attach(writers.FileWriter{})
	if e != nil {
		panic(e)
	}
	l.Notify()
}
