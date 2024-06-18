package helpers

import (
	"context"
	"os"
	"testing"

	"rsc.io/script"
	"rsc.io/script/scripttest"
)

func TestAll(t *testing.T) {
	ctx := context.Background()

	engine := &script.Engine{
		Conds: scripttest.DefaultConds(),
		Cmds:  scripttest.DefaultCmds(),
		Quiet: true,
	}

	env := os.Environ()
	scripttest.Test(t, ctx, engine, env, "*.txtar")
}
