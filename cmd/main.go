package main

import (
	"flag"
	"strings"

	"github.com/newtoallofthis123/exec/db"
	"github.com/newtoallofthis123/exec/utils"
)

var (
	output = flag.String("o", "", "Output file")
	help   = flag.Bool("h", false, "Show help")
	// TODO: Implement this
	// config = flag.String("c", "./exec.conf", "Specify a config file, defaults to ./exec.conf")
)

func main() {
	flag.Parse()

	if *help {
		flag.PrintDefaults()
		return
	}

	cmd := strings.Join(flag.Args(), " ")
	cmd = strings.TrimSpace(cmd)
	if cmd == "" {
		return
	}

	utils.InitPaths()

	db, err := db.MakeDb(utils.GetDbPath())
	if err != nil {
		panic(err)
	}

	err = db.Init()
	if err != nil {
		panic(err)
	}

	cmds, err := utils.CompileTasks(utils.GetConfigPaths(), &db)
	if err != nil {
		panic(err)
	}

	if cmds[cmd] != nil {
		err = utils.Run(*output, cmds[cmd])
		if err != nil {
			panic(err)
		}
	} else {
		panic("Task not found")
	}
}
