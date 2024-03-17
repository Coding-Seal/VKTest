package main

import (
	"VKTest/web/handlers"
	"log/slog"
	"os"
)

func main() {
	opts := &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}
	logger := slog.New(slog.NewJSONHandler(os.Stdout, opts))
	slog.SetDefault(logger)

	serv := NewServer()
	serv.Server.Handler = handlers.Handlers()
	serv.Server.ListenAndServe()
}
