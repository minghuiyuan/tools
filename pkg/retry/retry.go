package retry

import (
	"time"

	"github.com/pkg/errors"
)

// BackOffDelayDurationByMins : retry until function return nil
// you should put you retry contents into inputF, if no need to retry, return nil
// for every retry, will sleep for 1,2,4,8,16, maxDuration,maxDuration, ... minutes
func BackOffRetryMaxTimesMaxDurationUntilNoError(maxTimes int64, maxDuration time.Duration, inputFunc func() error) error {
	var times int64 = 1
	var dur int64 = 1
	var lastError error
	if maxTimes <= 0 {
		return errors.Errorf("first input parameter error: retry times should be at least 1")
	}
	if maxDuration < 1*time.Minute {
		return errors.Errorf("second input parameter error, maxDuration should be at least 1 minite.\n")
	}
	for true {
		lastError = inputFunc()
		if lastError == nil {
			break
		} else {
			// retry too max times
			if times > maxTimes {
				return lastError
			}
			// sleep and wait
			if time.Duration(dur*int64(time.Minute)) > maxDuration {
				time.Sleep(maxDuration)
			} else {
				time.Sleep(time.Duration(dur * int64(time.Minute)))
				dur *= 2
			}
			times += 1
		}
	}
	return lastError
}
