package console

import (
	"fmt"
	"strings"
	x "github.com/bolaxd/dumn/bot.wa/utils/parse"
	v "github.com/TwiN/go-color"
	)
	
func ResponseAll(m *x.Parse)  {
	isCmd := strings.HasPrefix(m.Body, m.Pref)
	send := strings.Split(m.Sender.String(), "@")[0]
	if strings.HasPrefix(m.Chat.String(), "status@") {
		return
	}
	if m.IsGc {
		if isCmd {
			fmt.Println(v.InPurple("Command :") + " " + v.InBlue(m.CmdP))
			fmt.Println(v.InCyan("Query :") + " " + v.InBlue(m.Query))
		} else if m.Body != "" {
			fmt.Println(v.InRed("Teks :") + " " + v.InBlue(m.Body))
		} else {
			fmt.Println(v.InRed("Send Media :") + " " + v.InBlue(m.TypeM))
		}
		fmt.Println(v.InYellow("Chat :") + " " + v.InBlue(m.Chat.String()))
		fmt.Println(v.InGreen("Sender :") + " " + v.InBlue(send))
		fmt.Println(v.InGray("Nama :") + " " + v.InBlue(m.Pushname))
		fmt.Println(v.InPurple("------------------------------------------------"))
	} else if !m.IsGc {
		if isCmd {
			fmt.Println(v.InPurple("Command :") + " " + v.InBlue(m.CmdP))
		} else if m.Body != "" {
			fmt.Println(v.InRed("Teks :") + " " + v.InBlue(m.Body))
		} else {
			fmt.Println(v.InRed("Send Media :") + " " + v.InBlue(m.TypeM))
		}
		fmt.Println(v.InYellow("Chat :") + " " + v.InBlue(m.Chat.String()))
		fmt.Println(v.InGreen("Sender :") + " " + v.InBlue(send))
		fmt.Println(v.InGray("Nama :") + " " + v.InBlue(m.Pushname))
		fmt.Println(v.InPurple("------------------------------------------------"))
	}
}