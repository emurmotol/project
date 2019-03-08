package endpoint

import (
	"context"
	"fmt"
	"time"

	gormadapter "github.com/casbin/gorm-adapter"
	"github.com/emurmotol/project/user/pkg/utils"
	"github.com/go-kit/kit/auth/casbin"
	endpoint "github.com/go-kit/kit/endpoint"
	log "github.com/go-kit/kit/log"
	metrics "github.com/go-kit/kit/metrics"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/spf13/viper"
	stdgrpc "google.golang.org/grpc"
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

func PostgresMiddleware() endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			db, err := gorm.Open("postgres", viper.GetString("POSTGRES_DATABASE"))
			if err != nil {
				return nil, err
			}
			defer db.Close()
			ctx = context.WithValue(ctx, utils.DBContextKey, db)
			return next(ctx, request)
		}
	}
}

func AuthorizerMiddleware() endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			// setup casbin context from db
			db := utils.GetDB(ctx)
			ctx = context.WithValue(ctx, casbin.CasbinModelContextKey, viper.GetString("CASBIN_MODEL"))
			ctx = context.WithValue(ctx, casbin.CasbinPolicyContextKey, gormadapter.NewAdapterByDB(db))

			subject := utils.GetClaims(ctx).Subject
			object := stdgrpc.ServerTransportStreamFromContext(ctx).Method()
			action := "read"
			e := func(ctx1 context.Context, i interface{}) (interface{}, error) { return ctx1, nil }
			mw := casbin.NewEnforcer(subject, object, action)(e)
			ctx2, err := mw(ctx, nil)
			if err != nil {
				return nil, err
			}
			return next(ctx2.(context.Context), request)
		}
	}
}
