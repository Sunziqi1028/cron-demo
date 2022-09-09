/**
 * @Author: Sun
 * @Description:
 * @File:  worker
 * @Version: 1.0.0
 * @Date: 2022/9/9 16:06
 */

package workqueue

import "sync"

type Worker struct {
	JobQueue JonChan
	Quit     chan bool // 退出标志
}

func NewWorker() Worker {
	return Worker{
		make(JonChan),
		make(chan bool),
	}
}

type WorkerPool struct {
	Size      int
	JobQueue  JonChan
	WorkQueue chan *Worker
}

func NewWorkerPool(poolSize, jobQueueLen int) *WorkerPool {
	return &WorkerPool{
		poolSize,
		make(JonChan, jobQueueLen),
		make(chan *Worker, poolSize),
	}
}

func (wp *WorkerPool) Start() {
	for i := 0; i < wp.Size; i++ {
		worker := NewWorker()
		worker.Start(wp, nil)
	}

	go func() {

		for {
			select {
			case Job := <-wp.JobQueue:
				worker := <-wp.WorkQueue
				worker.JobQueue <- Job
			}
		}
	}()
}

func (w Worker) Start(workerPool *WorkerPool, ser interface{}) {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		for {
			workerPool.WorkQueue <- &w
			select {
			case Job := <-w.JobQueue:
				Job.RunTask(nil, ser)
			case <-w.Quit:
				return
			}
		}
	}()
}
