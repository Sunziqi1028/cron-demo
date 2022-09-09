/**
 * @Author: Sun
 * @Description:
 * @File:  global
 * @Version: 1.0.0
 * @Date: 2022/9/9 16:25
 */

package global

import "cron/service/task/api/workqueue"

var AccessToken string

var WorkPools *workqueue.WorkerPool

const (
	PoolNum     = 100 * 100
	JobQueueNum = 100
)
