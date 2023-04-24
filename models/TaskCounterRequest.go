package models

type TaskCounterRequest struct {
	Start  int `json:"start" validate:"min=1"`
	End    int `json:"end" validate:"max=60"`
	TaskId string
}
