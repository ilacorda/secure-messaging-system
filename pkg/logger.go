package pkg

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
)

func Setup() {
	writer := zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: "15:04:05"}

	log.Logger = log.Output(writer)
	zerolog.SetGlobalLevel(zerolog.InfoLevel) // Default Log Level
}
