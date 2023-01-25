package handler

import (
	"strings"
	
	"go.mau.fi/whatsmeow"
	ev "go.mau.fi/whatsmeow/types/events"
	"github.com/bolaxd/dumn/bot.wa/utils/cmd"
	)
	
func MessageUpsert(conn *whatsmeow.Client) func(evt interface{}) {
	return func(evt interface{}) {
		switch v := evt.(type) {
		case *ev.Message:
			conn.MarkRead([]string{ v.Info.ID }, v.Info.Timestamp, v.Info.Chat, v.Info.Sender)
			if strings.HasPrefix(v.Info.ID, "BAE5") || strings.HasPrefix(v.Info.ID, "3EB0") || len(v.Info.ID) < 32 {
				return
			}
			go cmd.Cmd(conn, v)
		break
		}
	}
}
