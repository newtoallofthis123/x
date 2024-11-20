package utils

import (
	"strings"

	"github.com/newtoallofthis123/x/db"
	"github.com/newtoallofthis123/x/parser"
)

// CompileTasks compiles tasks from the provided paths and stores them in the database.
// It returns a map of task names to their respective paths and an error if any issues are encountered.
func CompileTasks(paths []string, db *db.Db) (map[string][]string, error) {
	tasks := make(map[string][]string)
	fromDb, err := db.GetAllTasks()
	if err != nil {
		return nil, err
	}

	for _, t := range fromDb {
		cmds := strings.Split(t.Cmd, "&&")
		cleaned := make([]string, 0)
		for _, c := range cmds {
			cleaned = append(cleaned, strings.TrimSpace(c))
		}

		tasks[t.Name] = append(tasks[t.Name], cleaned...)
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
