package cmd

import (
	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/types/events"
	console "github.com/bolaxd/dumn/bot.wa/utils"
	simple "github.com/bolaxd/dumn/bot.wa/utils/simple"
	mek "github.com/bolaxd/dumn/bot.wa/utils/parse"
	i "github.com/bolaxd/dumn/bot.wa/plugin"
)

func Cmd(conn *whatsmeow.Client, up *events.Message) {
	ball := simple.SimpleGo(conn, up)
	m := mek.Parser(ball, up)
	console.ResponseAll(m)
	switch m.Cmd {
		// BOT
	// case "menu":
	// 	go i.Menu(ball, m);
	// break;
	// case "ping":
	// 	go i.Ping(ball, m);
	// break;
	// // OWNER
	// case "bongkar":
	// 	go i.Bongkar(ball, m);
	// break;
	// case "upload":
	// 	go i.Upload(ball, m);
	// break;
	// case "metadata":
	// 	go i.Metadata(ball, m);
	// break;
	// case "getinfo":
	// 	go i.GetInfoGroup(ball, m);
	// break;
	// case "out":
	// 	go i.Out(ball, m);
	// break;
	// case "download":
	// 	go i.Down(ball, m);
	// break;
	case "exec":
		go i.Execute(ball, m);
	break;
	case "push":
		go i.Push(ball, m);
	break;
	case "pull":
		go i.Pull(ball, m);
	break;
	// case "stiker":
	// 	go i.Stiker(ball, m);
	// break;
	// // GROUP
	// case "polling":
	// 	go i.Polling(ball, m);
	// break;
	// case "setppgc":
	// 	go i.SetPPGC(ball, m);
	// break;
	// case "setname":
	// 	go i.SetName(ball, m);
	// break;
	// case "setdesk", "setdesc":
	// 	go i.SetDesc(ball, m);
	// break;
	// case "gclink", "link":
	// 	go i.Link(ball, m);
	// break;
	// case "gcrevoke", "revoke":
	// 	go i.Revoke(ball, m);
	// break;
	// case "gcopen", "open":
	// 	go i.GcOpen(ball, m);
	// break;
	// case "gcclose", "close":
	// 	go i.GcClose(ball, m);
	// break;
	// case "gclock", "lock":
	// 	go i.GcLock(ball, m);
	// break;
	// case "gcunlock", "unlock":
	// 	go i.GcUnlock(ball, m);
	// break;
	}
}