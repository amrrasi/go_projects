package main

import (
	"github.com/rs/zerolog/log"
)

func main() {

	log.Print("Test")
	log.Info().Str("Category", "Search").Msg("Searching ...")

}
