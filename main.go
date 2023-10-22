package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/jm199seo/screenshot_copy/cmd"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill, syscall.SIGTERM, syscall.SIGINT)
	defer cancel()

	ret := cmd.Execute(ctx)
	os.Exit(ret)
}
