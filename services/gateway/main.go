package main

import "github.com/nickonos/Spotify/packages/logging"

func main() {
	logger := logging.NewLogger("gateway")
	logger.Print("hello")
}
