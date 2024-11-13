package main

import (
	"github.com/newtoallofthis123/exec/db"
	"github.com/newtoallofthis123/exec/utils"
)

func main() {
	utils.InitPaths()

	db, err := db.MakeDb(utils.GetDbPath())
	if err != nil {
		panic(err)
	}

	err = db.Init()
	if err != nil {
		panic(err)
	}
}
