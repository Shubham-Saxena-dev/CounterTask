package utils

import (
	"fmt"
	"goPractice/models"
	"time"
)

const (
	IN_PROGRESS string = "In Progress"
	CANCELLED   string = "Cancelled"
	COMPLETED   string = "Complete"
)

type progress struct {
	status   string
	progress int
	start    int
	end      int
	cancel   bool
}

type Progress interface {
	GetStatus() string
	GetProgress() int
	CancelTask()
	Calculate()
}

func NewProgressTask(request models.TaskCounterRequest) Progress {
	return &progress{
		start:    request.Start,
		end:      request.End,
		progress: 0,
		status:   IN_PROGRESS,
	}
}

func (p *progress) Calculate() {
	fmt.Println("Task started")
	for i := p.start; i <= p.end; i++ {
		time.Sleep(1 * time.Second)
		p.progress = (i - p.start) * 100 / (p.end - p.start)

		if p.cancel {
			fmt.Println("Task Cancelled")
			p.status = CANCELLED
			break
		}
	}
	fmt.Println("Task finished")
	p.status = COMPLETED
}

func (p *progress) GetProgress() int {
	return p.progress
}

func (p *progress) GetStatus() string {
	return p.status
}

func (p *progress) CancelTask() {
	p.cancel = true
}
