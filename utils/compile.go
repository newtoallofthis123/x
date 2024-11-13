package utils

import (
	"github.com/newtoallofthis123/exec/db"
	"github.com/newtoallofthis123/exec/parser"
)

func CompileTasks(paths []string, db *db.Db) (map[string][]string, error) {
	tasks := make(map[string][]string)
	fromDb, err := db.GetAllTasks()
	if err != nil {
		return nil, err
	}

	for _, t := range fromDb {
		tasks[t.Name] = append(tasks[t.Name], t.Cmd)
	}

	for _, path := range paths {
		p, err := parser.MakeParser(path)
		if err != nil {
			return nil, err
		}

		err = p.Parse()
		if err != nil {
			return nil, err
		}

		for k, v := range p.GetCmds() {
			tasks[k] = append(tasks[k], v...)
		}
	}

	return tasks, nil
}
