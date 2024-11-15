package main

import (
	"fleetgo/logger"
	"fleetgo/worker"
	"fmt"
	"os"
	"os/signal"
	"reflect"
)

// ExampleActor 是一个示例的 Actor 结构体
type ExampleActor struct{}

// Echo 是 ExampleActor 的一个方法，接受任意类型的参数
func (a *ExampleActor) Echo(arg interface{}) (interface{}, error) {
	log := logger.GetLogger()
	log.Info("Received message:", arg)
	return fmt.Sprintf("Echo: %v", arg), nil
}

func (a *ExampleActor) Add(arg interface{}) (interface{}, error) {
	log := logger.GetLogger()
	log.Info("Received message:", arg, " type :", reflect.TypeOf(arg))
	switch v := arg.(type) {
	case int64:
		result := v + 1
		log.Info("Add:", result)
		return result, nil
	default:
		fmt.Println("Unknown type")
	}
	fmt.Println("Invalid argument type")
	return fmt.Sprintf("Error"), nil
}

func main() {
	var logPath *string = nil
	logger.InitLogger("info", logPath, "", 100, 30, 7)
	log := logger.GetLogger()
	// 创建 ExampleActor 实例
	actor := &ExampleActor{}

	// 创建 Worker 实例
	workerInstance, err := worker.NewWorker(actor, 50051)
	if err != nil {
		log.Error("failed to create workerInstance: %v", err)
	}

	// 启动 gRPC 服务器
	if err := workerInstance.Start(); err != nil {
		log.Fatal("failed to start gRPC server: %v", err)
	}

	// 等待中断信号来优雅地关闭服务器
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	<-ch

	// 停止 gRPC 服务器
	workerInstance.Stop()
}
