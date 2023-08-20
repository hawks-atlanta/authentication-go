package config

import (
	_ "github.com/caarlos0/env/v9"
)

type Environment struct {
	DatabaseEngine string `env:"DATABASE_ENGINE" envDefault:"sqlite"`
	DatabaseDSN    string `env:"DATABASE_DSN" envDefault:":memory:"`
}
