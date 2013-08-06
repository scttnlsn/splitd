package splitd

import (
	"bufio"
	"log"
	"net/http"
	"strings"
	"time"
)

type Relay struct {
	Config *Config
}

func NewRelay(config *Config) *Relay {
	r := &Relay{config}

	return r
}

func (r *Relay) Request() {
	req, _ := http.NewRequest("GET", r.Config.SourceURL, nil)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Printf("%s", err)
		return
	}

	reader := bufio.NewReader(res.Body)

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			return
		}

		body := strings.NewReader(line)
		req, _ = http.NewRequest("POST", r.Config.DestURL, body)

		c := &http.Client{}
		c.Do(req)
	}
}

func (r *Relay) Run() {
	for {
		r.Request()
		time.Sleep(time.Second)
	}
}
