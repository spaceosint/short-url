package repeatable

import "time"

func DoWithTries(fn func() error, attemtps int, duration time.Duration) (err error) {
	for attemtps > 0 {
		if err = fn(); err != nil {
			time.Sleep(duration)
			attemtps--
			continue
		}
		return nil
	}
	return
}
