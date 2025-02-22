package dto

import "go.mongodb.org/mongo-driver/bson/primitive"

type Task struct {
	Id            primitive.ObjectID `bson:"_id" json:"id"`
	Title         string             `bson:"title" json:"title"`
	Description   string             `bson:"description" json:"description"`
	Deadline      string             `bson:"deadline" json:"deadline"`
	Priority      string             `bson:"priority" json:"priority"`
	Progress      string             `bson:"progress" json:"progress"`
	AssigneeEmail string             `bson:"assigneeEmail" json:"assigneeEmail"`
	CreatedBy     string             `bson:"created_by" json:"created_by"`
}
