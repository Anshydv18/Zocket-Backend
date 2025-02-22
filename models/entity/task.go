package entity

import (
	"backend/base"
	"backend/constants"
	"backend/models/dto"
	"backend/models/requests"
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
)

type Task struct {
	dto.Task
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
		{Key: "progress", Value: task.Progress},
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

func GetTaskByCreatedByEmail(ctx *context.Context, email string) ([]*dto.Task, error) {
	dbClient := base.DBInstance
	if dbClient == nil {
		return nil, errors.New("error in db connection")
	}

	collection := dbClient.Database(constants.ZOCKETDB).Collection(constants.TASK_COLLECTION)
	filter := bson.M{
		"created_by": email,
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

func UpdateTaskValue(ctx *context.Context, request *requests.CreateTaskRequest) error {
	dbClient := base.DBInstance
	if dbClient == nil {
		return errors.New("error in db connection")
	}

	collection := dbClient.Database(constants.ZOCKETDB).Collection(constants.TASK_COLLECTION)
	filter := bson.M{
		"_id": request.Id,
	}

	_, err := collection.UpdateOne(*ctx, filter, bson.M{
		"$set": bson.M{
			"title":         request.Title,
			"description":   request.Description,
			"deadline":      request.Deadline,
			"priority":      request.Priority,
			"progress":      request.Progress,
			"assigneeEmail": request.AssigneeEmail,
			"created_by":    request.CreatedBy,
		},
	})

	return err
}
