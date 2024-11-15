package worker

import (
	"context"
	"google.golang.org/grpc"
	//"google.golang.org/protobuf/proto"
	"fleetgo/logger"
	"fleetgo/worker/pb"
	"fmt"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/anypb"
	"log"
	"net"
	"reflect"
)

// Worker 实现了 Worker 服务
type Worker struct {
	pb.UnimplementedWorkerServer
	actor      interface{}
	actorType  reflect.Type
	actorValue reflect.Value
	server     *grpc.Server
	listener   net.Listener
}

// NewWorker 创建一个新的 Worker 实例，并绑定到一个具体的 Actor
func NewWorker(actor interface{}, port int) (*Worker, error) {
	actorType := reflect.TypeOf(actor)
	actorValue := reflect.ValueOf(actor)

	// 创建 gRPC 服务器
	server := grpc.NewServer()

	// 监听端口
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return nil, err
	}

	workerInstance := &Worker{
		actor:      actor,
		actorType:  actorType,
		actorValue: actorValue,
		server:     server,
		listener:   l,
	}
	// 注册 Worker 服务
	pb.RegisterWorkerServer(server, workerInstance)
	return workerInstance, nil
}

// Execute 是 Worker 服务的方法
func (w *Worker) Execute(ctx context.Context, req *pb.ExecuteRequest) (*pb.ExecuteResponse, error) {
	log := logger.GetLogger()
	log.Info("Received execute request for method: ", req.MethodName)

	method, found := w.actorType.MethodByName(req.MethodName)
	if !found {
		return &pb.ExecuteResponse{Status: 1, Message: fmt.Sprintf("Method %s not found", req.MethodName)}, nil
	}

	// 解析参数
	var args []reflect.Value
	for _, arg := range req.MethodArgs {
		argValue, err := arg.UnmarshalNew()
		if err != nil {
			return &pb.ExecuteResponse{Status: 2, Message: "Failed to parse arguments"}, nil
		}
		// 使用类型断言检查类型
		if simpleInt, ok := argValue.(*pb.SimpleInt); ok {
			// 安全地访问 value 字段
			intValue := simpleInt.GetValue()
			args = append(args, reflect.ValueOf(intValue))
			fmt.Println(intValue) // 输出: 12345
		} else {
			args = append(args, reflect.ValueOf(argValue))
			fmt.Println("Not a SimpleInt type")
		}
	}

	// 调用方法
	results := method.Func.Call(append([]reflect.Value{w.actorValue}, args...))
	var result *anypb.Any
	// 处理返回值
	if len(results) > 0 {
		resValue := results[0].Interface()
		log.Info("Received result: ", resValue)
		if resValue != nil {
			// 直接将返回值转换为 proto.Message 类型
			var err error
			switch v := resValue.(type) {
			case string:
				result, err = anypb.New(&pb.SimpleString{Value: v})
			case int:
				result, err = anypb.New(&pb.SimpleInt{Value: int64(v)})
			case int64:
				result, err = anypb.New(&pb.SimpleInt{Value: v})
			case float32, float64:
				result, err = anypb.New(&pb.SimpleFloat{Value: float64(v.(float64))}) // 直接使用 float64 类型
			default:
				// 如果不是简单类型，尝试直接转换
				// 注意：这里假设 arg 实现了 ProtoMessage 接口
				result, err = anypb.New(resValue.(protoreflect.ProtoMessage))
			}
			if err != nil {
				return &pb.ExecuteResponse{Status: 3, Message: fmt.Sprintf("Failed to create Any message: %v", err)}, nil
			}
		} else {
			log.Info("Invalid result type", result)
			return &pb.ExecuteResponse{Status: 3, Message: "Result is not a valid type"}, nil
		}
	} else {
		log.Info("No results found.")
	}

	return &pb.ExecuteResponse{Status: 0, Message: "", Result: result}, nil
}

// Start 启动 gRPC 服务器
func (w *Worker) Start() error {
	log.Println("Starting gRPC server...")
	return w.server.Serve(w.listener)
}

// Stop 停止 gRPC 服务器
func (w *Worker) Stop() {
	log.Println("Stopping gRPC server...")
	w.server.Stop()
}
