package config

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	AwsConfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/rs/zerolog/log"
)

func (cfg Config) LoadAwsConfig() aws.Config {
	conf, err := AwsConfig.LoadDefaultConfig(context.TODO(),
		AwsConfig.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(
			cfg.R2.ApiKey, cfg.R2.ApiSecret, "",
		)), AwsConfig.WithRegion("auto"))
	if err != nil {
		log.Fatal().Msgf("unable to load AWS Config, %v", err)
	}

	log.Info().Msg("Load AWS Config success")
	return conf
}
