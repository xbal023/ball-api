package plugin 

import (
	"os/exec"
	"fmt"

	x "github.com/bolaxd/dumn/bot.wa/utils/parse"
	y "github.com/bolaxd/dumn/bot.wa/utils/simple"
	a "github.com/bolaxd/dumn/config"
	)
	
func Execute(ball *y.S, m *x.Parse)  {
	if !m.IsOwn { ball.Reply(a.FOwner, true); return }
	res, err := exec.Command("bash", "-c", m.Query).Output()
	if err != nil { ball.Reply(a.CustomError(string(err)), true); return }
	ball.Reply(a.CustomSuccess(string(res)), true)
}