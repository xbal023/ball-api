package plugin 

import (
	"os"
	"os/exec"
	"fmt"

	x "github.com/bolaxd/dumn/bot.wa/utils/parse"
	y "github.com/bolaxd/dumn/bot.wa/utils/simple"
	"github.com/bolaxd/dumn/helper"
	a "github.com/bolaxd/dumn/config"
	)
	
func Pull(ball *y.S, m *x.Parse)  {
	if !m.IsOwn { ball.Reply(a.FOwner, true); return }
	res, err := exec.Command("bash", "-c", "git pull").Output()
	if err != nil {
		ball.Reply(a.CustomError(string(err)), true)
		return
	}
	ball.Reply(a.CustomSuccess(string(res)), true)
}