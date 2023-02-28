package filter

import (
	"fmt"
	"git.supremind.info/gobase/2023/03-base/context"
	"time"
)

/*
 * @Author: ych
 * @Description: ...
 * @File: filter
 * @Version: ...
 * @Date: 2023-02-22 11:49:15
 */

type HandlerFunc func(c *context.Context)

type FilterBuilder func(next Filter) Filter

type Filter func(c *context.Context)

var _ FilterBuilder = MetricsFilterBuilder

func MetricsFilterBuilder(next Filter) Filter {
	return func(c *context.Context) {
		start := time.Now().Nanosecond()
		next(c)
		end := time.Now().Nanosecond()
		fmt.Printf("用了: %d 纳秒", end-start)
	}
}
