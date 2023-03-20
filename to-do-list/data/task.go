package data

import (
	"strings"
	"time"
)

type Task struct {
	ID          int
	Title       string
	UserID      int
	CreatedAt   string
	Deadline    string
	Description string
	IsImportant bool
	IsFinished  bool
}

func UserTasksByUserID(user_id int) (tasks []Task, err error) {
	rows, err := DB.Query(
		"SELECT ID, TITLE, USER_ID, ISIMPORTANT, ISFINISHED, DESCRIPTION,  CREATED_AT FROM TASKS WHERE USER_ID = $1 ORDER BY DEADLINE",
		user_id,
	)
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		var task Task
		err = rows.Scan(
			&task.ID, &task.Title, &task.UserID, &task.IsImportant, &task.IsFinished, &task.Description,
			&task.CreatedAt,
		)
		var deadline string
		err := DB.QueryRow("SELECT DEADLINE FROM TASKS WHERE ID=$1", task.ID).Scan(&deadline)
		if err == nil {
			task.Deadline = deadline
		}
		var createdAt string
		err = DB.QueryRow("SELECT CREATED_AT FROM TASKS WHERE ID=$1", task.ID).Scan(&createdAt)
		if err == nil {
			task.CreatedAt = createdAt
		}
		task.CreatedAt = strings.Split(task.CreatedAt, "T")[0]
		task.Deadline = strings.Split(task.Deadline, "T")[0]
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	return
}

func DeleteUserTasks(user User) (err error) {
	stmt, err := DB.Prepare("DELETE FROM TASKS WHERE USER_ID = $1")
	if err != nil {
		return
	}
	defer stmt.Close()
	_, err = stmt.Exec(user.ID)
	return
}

func (task *Task) Create() (err error) {
	if task.Deadline == "" {
		st, err := DB.Prepare("INSERT INTO TASKS(USER_ID, TITLE, DESCRIPTION, ISIMPORTANT, CREATED_AT) VALUES ($1, $2, $3, $4, $5) RETURNING ID, CREATED_AT")
		if err != nil {
			return err
		}
		defer st.Close()
		err = st.QueryRow(
			task.UserID, task.Title, task.Description, task.IsImportant, time.Now(),
		).Scan(&task.ID, &task.CreatedAt)
		return err
	}
	st, err := DB.Prepare("INSERT INTO TASKS(USER_ID, TITLE, DESCRIPTION, DEADLINE, ISIMPORTANT, CREATED_AT) VALUES ($1, $2, $3, $4, $5, $6) RETURNING ID, CREATED_AT")
	if err != nil {
		return err
	}
	defer st.Close()
	err = st.QueryRow(
		task.UserID, task.Title, task.Description, task.Deadline, task.IsImportant, time.Now(),
	).Scan(&task.ID, &task.CreatedAt)
	return err
}

func (task *Task) Delete() (err error) {
	stmt, err := DB.Prepare("DELETE FROM TASKS WHERE ID = $1")
	if err != nil {
		return
	}
	defer stmt.Close()
	_, err = stmt.Exec(task.ID)
	return
}

func (task *Task) Update() (err error) {
	stmt, err := DB.Prepare("UPDATE TASKS SET TITLE = $1, DESCRIPTION=$2, ISIMPORTANT=$3 WHERE ID = $4")
	if err != nil {
		return
	}
	defer stmt.Close()
	_, err = stmt.Exec(task.Title, task.Description, task.IsImportant, task.ID)
	return
}

func DeleteTaskByID(id int) (err error) {
	stmt, err := DB.Prepare("DELETE FROM TASKS WHERE ID = $1")
	if err != nil {
		return
	}
	defer stmt.Close()
	_, err = stmt.Exec(id)
	return
}
