package main

import (
	"context"
	"log/slog"
)

func main() {
	msg := "server starting"
	port := 8080
	// slog.Info(msg, "port", port) // port will escape to the heap
	ctx := context.Background()
	slog.LogAttrs(ctx, slog.LevelInfo, msg, slog.Int("port", port))
}
