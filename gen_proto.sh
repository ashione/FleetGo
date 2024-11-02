#!/bin/bash
protoc --go_out=. --go-grpc_out=. worker/worker.proto
