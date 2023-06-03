package db

import (
	"context"
	"fmt"
	"github.com/LeoUraltsev/todoapp/internal/task"
	"github.com/LeoUraltsev/todoapp/pkg/client/postgesql"
	"github.com/LeoUraltsev/todoapp/pkg/logger"
	"github.com/jackc/pgx/v5/pgconn"
	"strings"
)

type repository struct {
	client postgesql.Client
	logger *logger.Logger
}

func (r *repository) Create(ctx context.Context, task task.Task) (string, error) {
	q := `
		insert into task 
    		(title, description)
		values 
		    ($1, $2)
		returning id
	`
	if err := r.client.QueryRow(ctx, q, task.Title, task.Description).Scan(&task.ID); err != nil {
		if pgErr, ok := err.(*pgconn.PgError); ok {
			r.logger.Sugar().Errorf("SQL Error: %s, Detail: %s, Where: %s", pgErr.Message, pgErr.Detail, pgErr.Where)
			return "", nil
		}
		return "", err
	}

	return task.ID, nil

}

func (r *repository) FindAll(ctx context.Context) ([]task.Task, error) {
	q := `SELECT * FROM task`

	rows, err := r.client.Query(ctx, q)
	if err != nil {
		return nil, err
	}

	tasks := make([]task.Task, 0)

	for rows.Next() {
		var task task.Task

		err := rows.Scan(&task.ID, &task.Title, &task.Description)
		if err != nil {
			return nil, err
		}

		tasks = append(tasks, task)

	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func (r *repository) FindOne(ctx context.Context, id string) (task.Task, error) {
	q := `
		select
		    * 
		from 
		    task 
		where 
		    id=$1
	`

	var task task.Task

	err := r.client.QueryRow(ctx, q, id).Scan(&task.ID, &task.Title, &task.Description)
	if err != nil {
		return task, err
	}

	return task, nil
}

func (r *repository) Update(ctx context.Context, task task.Task) error {

	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argID := 1

	if task.Title != "" {
		setValues = append(setValues, fmt.Sprintf("title = $%d", argID))
		args = append(args, task.Title)
		argID++
	}

	if task.Description != "" {
		setValues = append(setValues, fmt.Sprintf("description = $%d", argID))
		args = append(args, task.Description)
		argID++
	}

	setQuery := strings.Join(setValues, ", ")

	q := fmt.Sprintf(`UPDATE task
		SET %s
		WHERE id = $%d
	`, setQuery, argID)

	r.logger.Sugar().Infof("Query string: %s. Args: %s", q, args)
	args = append(args, task.ID)

	res, err := r.client.Exec(ctx, q, args...)
	if err != nil {
		r.logger.Sugar().Errorf("Update error: %v", err)
		return err
	}

	r.logger.Sugar().Infof("Update data: %s, %d", res.String(), res.RowsAffected())

	return nil
}

func (r *repository) Delete(ctx context.Context, id string) error {
	q := `DELETE from task WHERE id=$1`
	exec, err := r.client.Exec(ctx, q, id)
	if err != nil {
		return err
	}
	r.logger.Sugar().Infof("Task delete: %s", exec.String())
	return nil
}

func NewRepository(client postgesql.Client, logger *logger.Logger) task.Repository {
	return &repository{
		client: client,
		logger: logger,
	}
}
