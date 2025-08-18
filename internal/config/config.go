package config

import (
	"time"

	"github.com/rs/zerolog"
)

type HttpInfo struct {
	HTTPAddr string
}

type GrpcInfo struct {
	GRPCAddr string
}

type LoggerInfo struct {
	LogLevel zerolog.Level
}

type Config struct {
	HttpInfo   *HttpInfo
	GrpcInfo   *GrpcInfo
	LoggerInfo *LoggerInfo

	ReadTimeout     time.Duration
	WriteTimeout    time.Duration
	IdleTimeout     time.Duration
	ShutdownTimeout time.Duration
}

func GetDefault() *Config {
	defaultLogLevel := zerolog.InfoLevel

	return &Config{
		HttpInfo: &HttpInfo{
			HTTPAddr: ":8080",
		},
		GrpcInfo: &GrpcInfo{
			GRPCAddr: ":9090",
		},
		LoggerInfo: &LoggerInfo{LogLevel: defaultLogLevel},

		ReadTimeout:     time.Second * 20,
		WriteTimeout:    time.Second * 20,
		IdleTimeout:     time.Second * 20,
		ShutdownTimeout: time.Second * 20,
	}
}
