package main

import (
	"context"
	"fleetgo/logger"
	fleetRpc "fleetgo/worker/rpc" // 替换为您实际的包路径
	"time"
)

func main() {
	// 创建日志记录器
	logger.InitLogger("info", nil, "", 100, 30, 7)
	log := logger.GetLogger()

	// 创建 WorkerClient 实例
	client, err := fleetRpc.NewWorkerClient("localhost:50051")
	if err != nil {
		log.Infof("could not create WorkerClient: %v", err)
	}
	defer func(client *fleetRpc.WorkerClient) {
		err := client.Close()
		if err != nil {
			log.Infof("Client close error %+v", err)
		}
	}(client)

	// 准备参数
	arg := "Hello, World!"

	// 创建上下文
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// 调用方法
	result, err := client.CallMethod(ctx, "Echo", arg)
	if err != nil {
		log.Fatalf("error when calling method: %v", err)
	}
	arg2 := 1
	result2, err2 := client.CallMethod(ctx, "Add", arg2)
	if err2 != nil {
		log.Fatalf("error when calling method: %v", err2)
	} else {
		log.Printf("Result %v", result2)
	}

	// 记录结果
	log.Infof("Result from method call: %v\n", result)
}
