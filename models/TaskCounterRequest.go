package models

import "github.com/google/uuid"

type TaskCounterRequest struct {
	Start  int `json:"start" validate:"min=1"`
	End    int `json:"end" validate:"max=60"`
	TaskId uuid.UUID
}
