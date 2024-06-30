package cronjob

import (
	"github.com/robfig/cron/v3"
	"testing"
)

func TestCron(t *testing.T) {
	_ = cron.New(cron.WithSeconds())
}
