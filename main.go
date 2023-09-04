package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/user"

	"gopkg.in/yaml.v2"
)

type Credentials struct {
	Hostname string `yaml:"hostname"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

// https://support.google.com/domains/answer/6147083?hl=en&ref_topic=9018335&sjid=8521529886545416990-NA
func main() {

	// Load config
	currentUser, _ := user.Current()
	data, err := os.ReadFile(currentUser.HomeDir + `/.ddns_gdomains`)
	if err != nil {
		log.Fatalf("Cant find config file in home dir, please place yaml at `~/.ddns_gdomains: %s", err)
	}
	var cfg Credentials
	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		log.Fatalf("Cant parse config file: %s", err)
	}

	// Create request object
	request_url := fmt.Sprintf("https://%s:%s@domains.google.com/nic/update?hostname=%s", cfg.Username, cfg.Password, cfg.Hostname)
	req, err := http.NewRequest("GET", request_url, nil)

	if err != nil {
		log.Fatalf("Cannot create http request object, %s", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("Request failed: %s", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read response")
	}
	log.Println(string(body))
}
