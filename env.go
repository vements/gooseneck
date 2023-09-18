package gooseneck

import (
	"os"
)

type Env struct {
}

func (e Env) MustDefine(key string) string {
	value := os.Getenv(key)
	if value == "" {
		Fatal().Str("key", key).Msg("not set or empty")
	}
	return value
}

func (e Env) Optional(key string, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		value = defaultValue
	}
	return value
}

func (e Env) Port() string {
	return e.Optional(PORT, "3000")
}
