package splitd

import (
	"bufio"
	"fmt"
	"net/http"
	"strings"
)

type Relay struct {
	Config *Config
}

func NewRelay(config *Config) *Relay {
	r := &Relay{config}
	return r
}

func (r *Relay) Run() {
	client := &http.Client{}

	req, err := http.NewRequest("GET", r.Config.SourceURL, nil)
	if err != nil {
		panic(fmt.Sprintf("%s", err))
	}

	res, err := client.Do(req)
	if err != nil {
		// TODO: New request
		return
	}

	reader := bufio.NewReader(res.Body)

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			// TODO: New request
			break
		}

		body := strings.NewReader(line)
		req, err = http.NewRequest("POST", r.Config.DestURL, body)
		if err != nil {
			panic(fmt.Sprintf("%s", err))
		}

		c := &http.Client{}
		c.Do(req)
	}
}
