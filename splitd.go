package main

import (
	"./splitd"
	"flag"
)

var config *splitd.Config

func init() {
	config = splitd.NewConfig()

	flag.StringVar(&config.SourceURL, "source-url", "", "URL from which to read messages (required)")
	flag.StringVar(&config.DestURL, "dest-url", "", "URL to which messages are sent (required)")
}

func main() {
	flag.Parse()

	r := splitd.NewRelay(config)

	r.Run()
}
