package main

import (
	"os"

	"go-reverse-tcp/proxy"

	"golang.org/x/exp/slog"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	logger.Info("This is an Info message")

	proxy.CreateService(
		"myservice",
		"0.0.0.0",
		"0.0.0.0",
		3000,
		3000,
	)
}
