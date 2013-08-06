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

func (r *Relay) Listen() {
	req, _ := http.NewRequest("GET", r.Config.SourceURL, nil)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Printf("%s", err)
		return
	}

	reader := bufio.NewReader(res.Body)

	for {
		msg, err := reader.ReadString('\n')
		if err != nil {
			return
		}

		go r.Send(msg)
	}
}

func (r *Relay) Send(msg string) {
	body := strings.NewReader(msg)
	req, _ := http.NewRequest("POST", r.Config.DestURL, body)

	client := &http.Client{}
	client.Do(req)
}

func (r *Relay) Run() {
	for {
		r.Listen()
		time.Sleep(time.Second)
	}
}
