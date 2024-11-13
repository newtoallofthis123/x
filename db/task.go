package db

import "time"

func (d *Db) AddTask(name, cmd string) error {
	_, err := d.db.Exec("INSERT INTO tasks (name, cmd, last_used) VALUES (?, ?, ?)", name, cmd, time.Now().String())
	return err
}

type Task struct {
	Id       int
	Name     string
	Cmd      string
	LastUsed time.Time
}

func (d *Db) GetTask(name string) (Task, bool) {
	row := d.db.QueryRow("SELECT id, name, cmd, last_used FROM tasks WHERE name = ?", name)
	var t Task
	err := row.Scan(&t.Id, &t.Name, &t.Cmd, &t.LastUsed)
	if err != nil {
		return Task{}, false
	}
	return t, true
}

func (d *Db) GetAllTasks() ([]Task, error) {
	rows, err := d.db.Query("SELECT id, name, cmd, last_used FROM tasks")
	if err != nil {
		return []Task{}, err
	}
	var tasks []Task

	for rows.Next() {
		var t Task
		err := rows.Scan(&t.Id, &t.Name, &t.Cmd, &t.LastUsed)
		if err != nil {
			return []Task{}, err
		}
		tasks = append(tasks, t)
	}

	return tasks, nil
}

func (d *Db) UpdateTask(name, cmd string) error {
	_, err := d.db.Exec("UPDATE tasks SET cmd = ?, last_used = ? WHERE name = ?", cmd, time.Now(), name)
	return err
}

func (d *Db) DeleteTaskByName(name string) error {
	_, err := d.db.Exec("DELETE FROM tasks WHERE name = ?", name)
	return err
}

func (d *Db) DeleteTask(q string) error {
	_, err := d.db.Exec("DELETE FROM tasks WHERE id = ? OR name = ?", q, q)
	return err
}
