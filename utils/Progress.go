package utils

import (
	"goPractice/models"
	"time"

	"github.com/google/uuid"
)

const (
	IN_PROGRESS string = "In Progress"
	CANCELLED   string = "Cancelled"
	COMPLETED   string = "Complete"
)

type progress struct {
	status   string
	progress int
}

type Progress interface {
	GetStatus() string
	GetProgress() int
}

var MyMap map[uuid.UUID]int
var status string

func calculate(request models.TaskCounterRequest) int {

	for i := request.Start; i <= request.End; i++ {
		time.Sleep(1000)

	}

	return 100
}

func (p *progress) GetProgress() {
	return
}
