package services

import (
	"time"

	"github.com/beego/beego/v2/core/logs"
)

func LoadDataWithRetry(maxRetries int, baseDelay time.Duration) error {
	var err error
	for attempt := 1; attempt <= maxRetries; attempt++ {

		logs.Info("Loading property data in memory...")
		err = LoadData()
		if err == nil {
			logs.Info("Property data successfully loaded to the memory")
			return nil
		}
		logs.Warning("Attempt %d/%d: failed to load data : %v", attempt, maxRetries, err)

		if attempt < maxRetries {
			backoff := baseDelay * time.Duration(1<<(attempt-1)) // 1x, 2x, 4x, 8x...
			logs.Informational("Retrying in %v...", backoff)
			time.Sleep(backoff)
		}
	}
	return err
}
