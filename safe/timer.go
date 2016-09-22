package safe
import (
	"fmt"
	"time"
)

func WaitCondition(timeout, retry time.Duration, cond func() bool) error {
	ok := make(chan struct{})
	exit := make(chan struct{})
	go func() {
		for {
			select {
			case <-exit:
			default:
				if cond() {
					close(ok)
					return
				}
				time.Sleep(retry)
			}
		}
	}()
	select {
	case <-ok:
		return nil
	case <-time.After(timeout):
		close(exit)
		return fmt.Errorf("timed out waiting for condition after %s", timeout)
	}
}
