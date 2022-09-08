/**
 * @Author: Sun
 * @Description:
 * @File:  main
 * @Version: 1.0.0
 * @Date: 2022/9/8 16:12
 */

package main

import (
	"context"
	"cron/service/monitor/server"
	"fmt"
	"github.com/jasonlvhit/gocron"
	"google.golang.org/grpc"
	"log"
)

func main() {

	MonitorTaskStatus()

	gocron.Every(10).Second().Do(MonitorTaskStatus)
	<-gocron.Start()
}

func MonitorTaskStatus() {
	conn, err := grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	rpcClient := task.NewTaskClient(conn)
	fmt.Println("监听任务执行状态...")
	fmt.Println("发现数据库任务1状态执行失败，调用RPC服务")

	ExecTaskRpc(rpcClient, 1)
}

func ExecTaskRpc(t task.TaskClient, id int64) {
	resp, err := t.TaskToMonitor(context.Background(), &task.Req{
		Id: id,
	})

	if err != nil {
		log.Fatal("调用gRPC方法错误: ", err)
	}
	fmt.Println("重新执行任务成功")
	fmt.Println(resp)
}
