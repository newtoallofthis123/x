package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/newtoallofthis123/x/db"
	"github.com/newtoallofthis123/x/utils"
)

var (
	output  = flag.String("o", "", "Output file")
	help    = flag.Bool("h", false, "Show help")
	version = flag.Bool("v", false, "Show version")
	list    = flag.Bool("l", false, "List out all parsed tasks and exit")
	add     = flag.Bool("a", false, "Add a task to the database")
	rename  = flag.Bool("r", false, "Rename a task in the database")
	del     = flag.Bool("d", false, "Delete a task from the database")
	sup     = flag.Bool("s", false, "Suppress output")
	sync    = flag.Bool("sync", false, "Sync the database with the current config")
	edit    = flag.Bool("e", false, "Edit a task")
	config  = flag.String("c", "exec.conf", "Specify a config file, defaults to exec.conf")
)

func printHelp() {
	fmt.Println("=============")
	fmt.Println("X v.0.1.2")
	fmt.Println("=============")
	flag.PrintDefaults()
}

func main() {
	flag.Parse()

	if *help {
		printHelp()
		return
	}

	if *version {
		fmt.Println("X v.0.1.3")
		return
	}

	cmd := strings.Join(flag.Args(), " ")
	cmd = strings.TrimSpace(cmd)
	utils.InitPaths()

	if !*add && !*edit && !*del && !*list && !*sync && cmd == "" && !*rename {
		printHelp()
		return
	}

	db, err := db.MakeDb(utils.GetDbPath())
	if err != nil {
		panic(err)
	}

	err = db.Init()
	if err != nil {
		panic(err)
	}

	if *add {
		if cmd == "" {
			panic("No task provided")
		}
		name := strings.Split(cmd, " ")[0]
		c := strings.Join(strings.Split(cmd, " ")[1:], " ")
		sp := strings.Split(c, ",")
		cm := strings.Join(sp, "&&")

		err = db.AddTask(name, cm)
		if err != nil {
			panic(err)
		}

		fmt.Println("Task added")
		return
	}

	if *rename {
		if cmd == "" {
			panic("No task provided")
		}

		name := strings.Split(cmd, " ")[0]
		ren := strings.Split(cmd, " ")[1]

		task, ok := db.GetTask(name)
		if !ok {
			panic("Task not found")
		}

		err = db.DeleteTaskByName(name)
		if err != nil {
			panic(err)
		}

		err = db.AddTask(ren, task.Cmd)
		if err != nil {
			panic(err)
		}

		fmt.Println("Task renamed")
		return
	}

	if *del {
		if cmd == "" {
			panic("No task provided")
		}

		err = db.DeleteTask(cmd)
		if err != nil {
			panic(err)
		}

		fmt.Println("Task deleted")
		return
	}

	if *edit {
		if cmd == "" {
			panic("No task provided")
		}

		name := strings.Split(cmd, " ")[0]
		c := strings.Join(strings.Split(cmd, " ")[1:], " ")
		sp := strings.Split(c, ",")
		cm := strings.Join(sp, "&&")

		err = db.UpdateTask(name, cm)
		if err != nil {
			panic(err)
		}

		fmt.Println("Task updated")
		return
	}

	cmds, err := utils.CompileTasks(utils.GetConfigPaths(*config), &db)
	if err != nil {
		panic(err)
	}

	if *list {
		for k, v := range cmds {
			fmt.Print(k + ": [" + strings.Join(v, ",") + "]\n")
		}
		return
	}

	if *sync {
		confirm := ""
		fmt.Print("Are you sure you want to sync the database with the current config? (y/n): ")
		fmt.Scanln(&confirm)
		if confirm != "y" {
			return
		}

		err = db.Truncate()
		if err != nil {
			panic(err)
		}

		for k, v := range cmds {
			err = db.AddTask(k, strings.Join(v, "&&"))
			if err != nil {
				panic(err)
			}
		}

		fmt.Println("Database synced")
		return
	}

	if *sup {
		*output = "!"
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
