package scheduler

import (
	"time"

	"github.com/go-co-op/gocron"
)

func startSchedulerFor(task func()) error {
	s := gocron.NewScheduler(time.UTC) //한국 시각으로 수정할것
	_, err := s.Every(1).Week().Weekday(time.Friday).Do(task)
	if err != nil {
		return err
	}
	return nil
}
