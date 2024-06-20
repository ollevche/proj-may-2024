package main

import (
	"errors"
	"fmt"
	"io"
	"math/rand"
	"os"

	"github.com/rs/zerolog/log"
)

func createFile(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("creating file: %w", err)
	}
	defer file.Close()

	linesNum := rand.Intn(100)
	var fileContent string

	log.Info().Msgf("Generating %d lines", linesNum)

	for i := 0; i < linesNum; i++ {
		fileContent += fmt.Sprintf("%d:\tabcde\n", i)
	}

	if _, err = file.WriteString(fileContent); err != nil {
		return fmt.Errorf("writing string to file: %w", err)
	}

	return nil
}

func main() {
	const filename = "a.txt"

	err := createFile(filename)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to create file")
	}

	fileContent, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to read file")
	}

	fmt.Print(string(fileContent))

	fmt.Println(len(fileContent))

	buff := make([]byte, 5)

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to open file")
	}

	for {
		n, err := file.Read(buff)

		// && = ТА, || = АБО
		// ЯКЩО помилка ненульова ТА помилка є EOF (End Of File)
		// if err != nil && errors.Is(err, io.EOF) {
		// }

		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			log.Fatal().Err(err).Msg("Failed to read file")
		}

		fmt.Printf("N=%d: `%s`\n", n, string(buff))
	}
}
