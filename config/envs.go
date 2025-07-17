package config

import (
	"github.com/caarlos0/env/v11"
)

type Envs struct {
	LogLevel             string `env:"LOG_LEVEL" envDefault:"info"`
	ServiceVersion       string `env:"SERVICE_VERSION" envDefault:"1.0.0"`
	HttpPort             int    `env:"HTTP_PORT" envDefault:"3000"`
	DelayShutdownSeconds int    `env:"DELAY_SHUTDOWN_SECONDS" envDefault:"30"`
	Foo                  string `env:"FOO" envDefault:"bar"`
}

func MustGetEnvs() Envs {
	envs, err := env.ParseAs[Envs]()
	if err != nil {
		panic(err.Error())
	}

	return envs
}
