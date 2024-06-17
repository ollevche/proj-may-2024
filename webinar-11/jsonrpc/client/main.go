package main

import (
	"fmt"
	"net/rpc"
	"net/rpc/jsonrpc"

	"github.com/rs/zerolog/log"
)

func main() {
	client, err := NewCalculatorClient(":8080")
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to create client")
	}
	defer client.Close()

	log.Info().Msg("Trying to send JSON-RPC requests")

	result, err := client.Add(1, 3)
	if err != nil {
		log.Error().Err(err).Msg("Failed to add values")
	} else {
		log.Info().Int("result", result).Msg("Got OKAY response")
	}

	result, err = client.Divide(1, 0)
	if err != nil {
		log.Error().Err(err).Msg("Failed to divide values")
	} else {
		log.Info().Int("result", result).Msg("Got OKAY response")
	}

	result, err = client.Divide(25, 5)
	if err != nil {
		log.Error().Err(err).Msg("Failed to divide values")
	} else {
		log.Info().Int("result", result).Msg("Got OKAY response")
	}
}

type CalculatorClient struct {
	c *rpc.Client
}

func NewCalculatorClient(addr string) (*CalculatorClient, error) {
	c, err := jsonrpc.Dial("tcp", addr)
	if err != nil {
		return nil, fmt.Errorf("dialing jsonrpc: %w", err)
	}

	return &CalculatorClient{c: c}, nil
}

type Args struct {
	A, B int
}

func (c *CalculatorClient) Add(a, b int) (int, error) {
	var result int

	err := c.c.Call("Calculator.Add", Args{a, b}, &result)
	if err != nil {
		return 0, fmt.Errorf("calling Add: %w", err)
	}

	return result, nil
}

func (c *CalculatorClient) Divide(a, b int) (int, error) {
	var result int

	err := c.c.Call("Calculator.Divide", Args{a, b}, &result)
	if err != nil {
		return 0, fmt.Errorf("calling Divide with %d and %d: %w", a, b, err)
	}

	return result, nil
}

func (c *CalculatorClient) Close() {
	c.c.Close()
}
