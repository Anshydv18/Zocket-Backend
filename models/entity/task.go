package entity

import (
	"backend/base"
	"backend/constants"
	"backend/models/dto"
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
)

type Task struct {
	Title         string `json:"title"`
	Description   string `json:"description"`
	Deadline      string `json:"deadline"`
	Priority      string `json:"priority"`
	Progess       string `json:"progress"`
	AssigneeEmail string `json:"assigneeEmail"`
	CreatedBy     string `json:"created_by"`
}

func (task *Task) CreateTask(ctx *context.Context) error {
	dbClient := base.DBInstance
	if dbClient == nil {
		return errors.New("error in db connection")
	}

	collection := dbClient.Database(constants.ZOCKETDB).Collection(constants.TASK_COLLECTION)
	document := bson.D{
		{Key: "title", Value: task.Title},
		{Key: "description", Value: task.Description},
		{Key: "deadline", Value: task.Deadline},
		{Key: "priority", Value: task.Priority},
		{Key: "progress", Value: task.Progess},
		{Key: "assigneeEmail", Value: task.AssigneeEmail},
		{Key: "created_by", Value: task.CreatedBy},
	}
	_, err := collection.InsertOne(*ctx, document)
	return err
}

func GetTaskByAssigneEmail(ctx *context.Context, email string) ([]*dto.Task, error) {
	dbClient := base.DBInstance
	if dbClient == nil {
		return nil, errors.New("error in db connection")
	}

	collection := dbClient.Database(constants.ZOCKETDB).Collection(constants.TASK_COLLECTION)
	filter := bson.M{
		"assigneeEmail": email,
	}

	data, err := collection.Find(*ctx, filter)
	if err != nil {
		return nil, err
	}

	var Tasks []*dto.Task
	for data.Next(*ctx) {
		var task dto.Task
		if err := data.Decode(&task); err != nil {
			return nil, err
		}
		Tasks = append(Tasks, &task)
	}

	return Tasks, nil
}
