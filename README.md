![logo](assets/infiltrator0.jpg)

# Infiltrator 🕵️‍♂️

**Infiltrator** is a stealth-oriented input surveillance tool written in Go.  
It captures keyboard input, clipboard contents, system metadata, and exfiltrates the data through a secure Telegram bot channel.

> ⚠️ This tool is intended strictly for **educational**, **offensive security research**, and **red team** purposes.

---

## ✨ Features

- ⌨️ Keyboard input logging (raw keystrokes)
- 📋 Clipboard content capture
- 🌐 IP address and approximate geolocation
- 💻 Hostname & current user info
- 📷 Periodic screenshots
- 📤 Periodic exfiltration to Telegram bot (via HTTP API)
- 🧩 Modular structure – easily extendable (e.g. mic access, etc.)

---

## ⚙️ Requirements

- Go 1.19+
- A Telegram bot token ([@BotFather](https://t.me/BotFather))
- Your Telegram user ID or chat ID

---

## 🚀 Installation

```bash
git clone https://github.com/christbowel/Infiltrator.git
cd Infiltrator
go build -o infiltrator.exe main.go
```

## 🧠 Configuration

Edit the following variables in `main.go`:

```go
botToken = "YOUR_BOT-Token"
ChatID   = "YOUT_CHAT-ID"
```
---

## 🔒 Disclaimer

This project is provided for educational and authorized security testing only.
Using this software on devices you do not own or without explicit permission is illegal and strictly prohibited.

> ⚠️ The author declines any responsibility for misuse or damage caused by this tool.

---

🎯 **Why Infiltrator?**  
💡 Infiltrator was born from the need for a lightweight yet powerful input surveillance tool written in Go.  
Instead of relying on bulky or complex frameworks, it focuses on stealthy and efficient keylogging, clipboard capture, and system info exfiltration over Telegram.  
This makes it ideal for red teamers and security researchers who want to quickly deploy and gather input data without raising suspicion.

---

🤖 **Future Improvements**  
🛠️ Adding screenshot capture for visual context.  
⚡ Implementing startup persistence to survive reboots.  
🔐 Encrypting stored logs locally before exfiltration.  
📊 Adding alternative exfiltration channels beyond Telegram (SFTP, webhook).  
🎭 Binary obfuscation for stealthier deployment.

---

✨ **Credits**  
Developed by @Christbowel, inspired by minimalist and efficient security tools.

Special thanks to the authors of the following libraries for core functionality:  
- [github.com/atotto/clipboard](https://github.com/atotto/clipboard) clipboard management  
- [github.com/eiannone/keyboard](https://github.com/eiannone/keyboard) keyboard input capture

---

👀 Stay stealthy, stay sharp! 🚀
