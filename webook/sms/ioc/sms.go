package ioc

import (
	"github.com/google/wire"
	"webook/sms/_internal/service/sms"
	"webook/sms/_internal/service/sms/failover"
	"webook/sms/_internal/service/sms/memory"
	memory2 "webook/sms/_internal/service/sms/memory2"
	memory3 "webook/sms/_internal/service/sms/memory3"
	"webook/sms/_internal/service/sms/ratelimit"
	"webook/sms/_internal/service/sms/retry"
)

func SMS() sms.SMS {
	s1 := memory.NewMemory()
	s2 := memory2.NewMemory2()
	s3 := memory3.NewMemory3()
	failoverSMS := failover.NewFailoverSMS(s1, s2, s3)
	newRetry := retry.NewRetry(failoverSMS)
	return ratelimit.NewRatelimit(newRetry)
}

var InitSMS = wire.NewSet(
	SMS,
)
