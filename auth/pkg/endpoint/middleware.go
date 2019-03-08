package endpoint

import (
	"context"
	"fmt"
	"time"

	"github.com/emurmotol/project/auth/pkg/grpc/pb"
	"github.com/emurmotol/project/auth/pkg/utils"
	endpoint "github.com/go-kit/kit/endpoint"
	log "github.com/go-kit/kit/log"
	metrics "github.com/go-kit/kit/metrics"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

// InstrumentingMiddleware returns an endpoint middleware that records
// the duration of each invocation to the passed histogram. The middleware adds
// a single field: "success", which is "true" if no error is returned, and
// "false" otherwise.
func InstrumentingMiddleware(duration metrics.Histogram) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			defer func(begin time.Time) {
				duration.With("success", fmt.Sprint(err == nil)).Observe(time.Since(begin).Seconds())
			}(time.Now())
			return next(ctx, request)
		}
	}
}

// LoggingMiddleware returns an endpoint middleware that logs the
// duration of each invocation, and the resulting error, if any.
func LoggingMiddleware(logger log.Logger) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			defer func(begin time.Time) {
				logger.Log("transport_error", err, "took", time.Since(begin))
			}(time.Now())
			return next(ctx, request)
		}
	}
}

func UserClientMiddleware() endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			conn, err := grpc.Dial(viper.GetString("USER_API_GRPC_ADDRESS"), grpc.WithInsecure())
			if err != nil {
				return nil, err
			}
			svc := pb.NewUserClient(conn)
			defer conn.Close()
			ctx = context.WithValue(ctx, utils.UserClientContextKey, svc)
			return next(ctx, request)
		}
	}
}
