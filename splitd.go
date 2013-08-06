package main

import (
	"./splitd"
	"flag"
	"fmt"
	"net/url"
)

var config *splitd.Config

func init() {
	config = splitd.NewConfig()

	flag.StringVar(&config.SourceURL, "source-url", "", "URL from which to read messages (required)")
	flag.StringVar(&config.DestURL, "dest-url", "", "URL to which messages are sent (required)")
}

func main() {
	flag.Parse()

	_, err := url.Parse(config.SourceURL)
	if err != nil {
		panic(fmt.Sprintf("Source URL malformed: %s", err))
	}

	_, err = url.Parse(config.DestURL)
	if err != nil {
		panic(fmt.Sprintf("Destination URL malformed: %s", err))
	}

	r := splitd.NewRelay(config)

	r.Run()
}
