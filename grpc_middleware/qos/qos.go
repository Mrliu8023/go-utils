package qos

import (
	"context"
	"google.golang.org/grpc"
	"time"
)

type Qos struct {
	Times     int
	TimeStemp time.Time
}

func NewQos() *Qos {
	return &Qos{0, time.Now()}
}

func UnaryServerInterceptor(qos *Qos) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		if qos != nil {
			if time.Since(qos.TimeStemp) > time.Second {
				qos = NewQos()
			}
			qos.Times++
		}
		return handler(ctx, req)
	}
}
