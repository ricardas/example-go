package datetime

import "time"

func CurrentTime() string {
	return time.Now().Format(time.RFC3339)
}
