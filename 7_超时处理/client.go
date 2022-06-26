package overTime

import (
	"context"
	"time"
)

// package overTime

// client
func client() {
	ctx, err := context.WithTimeout(context.Background(), time.Second*3)
	if err != nil {
	}

	// call function
	// ctx 代替 context.Background()
}
