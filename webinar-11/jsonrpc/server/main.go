package main

import (
	"errors"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"

	"github.com/rs/zerolog/log"
)

func main() {
	var c = &Calculator{}

	server := rpc.NewServer()

	if err := server.Register(c); err != nil {
		log.Fatal().Err(err).Msg("Failed to register calculator")
	}

	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to create listener")
	}

	log.Info().Msg("Started server!")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Error().Err(err).Msg("Failed to accept connection")
		}

		log.Info().Msg("Connection accepted")

		go func() {
			server.ServeCodec(jsonrpc.NewServerCodec(conn))
			log.Info().Msg("Closed connection")
		}()
	}
}

type Calculator struct{}

type AddArgs struct {
	A, B int
}

func (c *Calculator) Add(arg AddArgs, res *int) error {
	log.Debug().
		Int("a", arg.A).
		Int("b", arg.B).
		Msg("Adding two numbers")

	*res = arg.A + arg.B
	return nil
}

type DivideArgs = AddArgs

func (c *Calculator) Divide(arg DivideArgs, res *int) error {
	log.Debug().
		Int("a", arg.A).
		Int("b", arg.B).
		Msg("Dividing two numbers")

	if arg.B == 0 {
		return errors.New("division by zero")
	}

	*res = arg.A / arg.B
	return nil
}
