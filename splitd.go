package main

import (
	"flag"
	"fmt"
	"github.com/scttnlsn/splitd/splitd"
	"net/url"
	"os"
)

var config *splitd.Config

func init() {
	config = splitd.NewConfig()

	flag.StringVar(&config.SourceURL, "source-url", "", "URL from which to read messages (required)")
	flag.StringVar(&config.SourceAuth, "source-auth", "", "HTTP basic auth required for source requests")
	flag.StringVar(&config.SourceMethod, "source-method", "GET", "HTTP request method for source requests")
	flag.StringVar(&config.DestURL, "dest-url", "", "URL to which messages are sent (required)")
	flag.StringVar(&config.DestAuth, "dest-auth", "", "HTTP basic auth required for destination requests")
	flag.StringVar(&config.DestMethod, "dest-method", "POST", "HTTP request method for destination requests")
}

func main() {
	version := flag.Bool("v", false, "output the version number")

	flag.Parse()

	if *version {
		fmt.Println(splitd.Version)
		os.Exit(0)
	}

	_, err := url.Parse(config.SourceURL)
	if config.SourceURL == "" || err != nil {
		panic(fmt.Sprintf("Source URL malformed: %s", err))
	}

	_, err = url.Parse(config.DestURL)
	if config.DestURL == "" || err != nil {
		panic(fmt.Sprintf("Destination URL malformed: %s", err))
	}

	r := splitd.NewRelay(config)
	r.Run()
}
