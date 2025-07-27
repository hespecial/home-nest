package svc

import (
	"fmt"
	"github.com/hibiken/asynq"
	"home-nest/app/mqueue/cmd/scheduler/internal/config"
	"time"
)

// create scheduler
func newScheduler(c config.Config) *asynq.Scheduler {
	location, _ := time.LoadLocation("Asia/Shanghai")
	return asynq.NewScheduler(
		asynq.RedisClientOpt{
			Addr: c.Redis.Host,
		},
		&asynq.SchedulerOpts{
			Location: location,
			PostEnqueueFunc: func(info *asynq.TaskInfo, err error) {
				if err != nil {
					fmt.Printf("Scheduler EnqueueErrorHandler <<<<<<<===>>>>> err : %+v , task : %+v", err, info)
				}
			},
		},
	)
}
