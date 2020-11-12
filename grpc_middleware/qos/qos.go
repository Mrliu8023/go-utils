package qos

import (
	"context"
	"time"

	"google.golang.org/grpc"
)

type Qos struct {
	Times     int
	TimeStemp time.Time
}

func Qos() *Qos {
	return &Qos{0, time.Now()}
}

func UnaryServerInterceptor(qos *Qos) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		if qos != nil {
			if time.Until(qos.TimeStemp) > time.Second {
				qos = Qos()
			}
			qos.Times++
		}
		return handler(ctx, req)
	}
}
