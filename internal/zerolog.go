package internal

import (
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func InitZerolog(verbose bool) {
	zerolog.TimestampFunc = func() time.Time {
		return time.Now().In(time.Local)
	}
	if verbose {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC822}).With().Caller().Logger()
	} else {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC822})
	}
}
