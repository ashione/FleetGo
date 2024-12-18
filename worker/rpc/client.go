package rpc

import (
	"context"
	"errors"
	flog "fleetgo/logger"
	"fmt"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/anypb"

	worker_pb "fleetgo/worker/pb" // 替换为您实际的包路径
)

// WorkerClient 封装了与 Worker 服务交互的功能
type WorkerClient struct {
	client     worker_pb.WorkerClient
	conn       *grpc.ClientConn
	serverAddr string
	logger     *log.Logger
}

// NewWorkerClient 创建一个新的 WorkerClient 实例
func NewWorkerClient(serverAddr string) (*WorkerClient, error) {
	conn, err := grpc.Dial(serverAddr, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return nil, err
	}
	client := worker_pb.NewWorkerClient(conn)
	return &WorkerClient{
		client:     client,
		conn:       conn,
		serverAddr: serverAddr,
		logger:     flog.GetLogger(),
	}, nil
}

// Close 关闭 gRPC 连接
func (c *WorkerClient) Close() error {
	return c.conn.Close()
}

func (c *WorkerClient) ConvertArgs(args ...interface{}) ([]*anypb.Any, error) {
	var methodArgs []*anypb.Any
	for _, arg := range args {
		var argAny *anypb.Any
		var err error

		// 检查是否为简单类型（如 string, int, float 等）
		switch v := arg.(type) {
		case string:
			argAny, err = anypb.New(&worker_pb.SimpleString{Value: v})
		case int, int32, int64:
			argAny, err = anypb.New(&worker_pb.SimpleInt{Value: int64(v.(int))}) // 直接使用 int64 类型
		case float32, float64:
			argAny, err = anypb.New(&worker_pb.SimpleFloat{Value: float64(v.(float64))}) // 直接使用 float64 类型
		default:
			// 如果不是简单类型，尝试直接转换
			// 注意：这里假设 arg 实现了 ProtoMessage 接口
			argAny, err = anypb.New(arg.(protoreflect.ProtoMessage))
		}
		if err != nil {
			c.logger.Printf("Error creating Any from argument: %v\n", err)
			return nil, err
		}
		methodArgs = append(methodArgs, argAny)
	}
	return methodArgs, nil
}

// CallMethod 调用 Worker 服务上的指定方法
// 示例主函数，展示如何使用 WorkerClient
func (c *WorkerClient) CallMethod(ctx context.Context, methodName string, args ...interface{}) (interface{}, error) {
	methodArgs, convertErr := c.ConvertArgs(args...)
	if convertErr != nil {
		return nil, convertErr
	}
	resp, err := c.client.Execute(ctx, &worker_pb.ExecuteRequest{
		MethodName: methodName,
		MethodArgs: methodArgs,
	})
	if err != nil {
		c.logger.Printf("Error executing method %s: %v\n", methodName, err)
		return nil, err
	}

	if resp.Status != 0 {
		c.logger.Printf("Method %s returned with status %d: %s\n", methodName, resp.Status, resp.Message)
		return nil, errors.New(resp.Message)
	}

	if resp.Result == nil {
		return nil, fmt.Errorf("response result is nil")
	}

	// 提取日志和错误处理到一个公共函数中
	handleUnmarshal := func(result interface{}, err error) (interface{}, error) {
		if err != nil {
			c.logger.Printf("Error unmarshalling response: %v\n", err)
			return nil, err
		}
		return result, nil
	}

	switch resp.Result.TypeUrl {
	case "type.googleapis.com/worker.pb.SimpleInt":
		var result worker_pb.SimpleInt
		c.logger.Printf("result type %v\n", resp.Result)
		return handleUnmarshal(result.Value, resp.Result.UnmarshalTo(&result))

	case "type.googleapis.com/worker.pb.SimpleFloat":
		var result worker_pb.SimpleFloat
		c.logger.Printf("result type %v\n", resp.Result)
		return handleUnmarshal(result.Value, resp.Result.UnmarshalTo(&result))

	case "type.googleapis.com/worker.pb.SimpleString":
		var result worker_pb.SimpleString
		c.logger.Printf("result type %v\n", resp.Result)
		return handleUnmarshal(result.Value, resp.Result.UnmarshalTo(&result))

	default:
		c.logger.Printf("Unknown result type: %v\n", resp.Result.TypeUrl)
		return nil, fmt.Errorf("unknown result type: %v", resp.Result.TypeUrl)
	}
}
