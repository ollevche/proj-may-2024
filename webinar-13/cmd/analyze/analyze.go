package main

import (
	"fmt"
	"os"
	"webinar13/internal/analyzer"
	"webinar13/pkg/property"

	"github.com/rs/zerolog/log"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal().Msg("Expected file as an argument")
	}

	filename := os.Args[1]

	analyzer := analyzer.New(
		analyzer.NewPrefixWrapper(property.FilenameAnalyzer{}, "dev_"),
		property.ContentLengthAnalyzer{},
	)

	report, err := analyzer.Analyze(filename)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to analyze file")
	}

	for k, v := range report {
		fmt.Printf("%s\t%v\n", k, v)
	}
}
