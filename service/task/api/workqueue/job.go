/**
 * @Author: Sun
 * @Description:
 * @File:  job
 * @Version: 1.0.0
 * @Date: 2022/9/9 16:05
 */

package workqueue

type Job interface {
	RunTask(request interface{}, svcCtx interface{})
}

type JonChan chan Job
