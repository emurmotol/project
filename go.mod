module github.com/emurmotol/project

require (
	github.com/99designs/gqlgen v0.7.2
	github.com/emurmotol/project/api v0.0.0
	github.com/emurmotol/project/auth v0.0.0
	github.com/emurmotol/project/auth_api v0.0.0-20190308075526-e3ea4d08b23a // indirect
	github.com/emurmotol/project/user v0.0.0
	github.com/go-kit/kit v0.8.0
	github.com/golang/protobuf v1.3.0
	github.com/kujtimiihoxha/kit v0.0.0-20190129201007-e40b4aff8f4e // indirect
	github.com/lightstep/lightstep-tracer-go v0.15.6
	github.com/oklog/oklog v0.3.2
	github.com/opentracing/opentracing-go v1.0.2
	github.com/openzipkin/zipkin-go-opentracing v0.3.5
	github.com/pkg/errors v0.8.1 // indirect
	github.com/prometheus/client_golang v0.9.2
	github.com/urfave/cli v1.20.0 // indirect
	github.com/vektah/gqlparser v1.1.1
	golang.org/x/net v0.0.0-20190227160552-c95aed5357e7
	google.golang.org/grpc v1.19.0
	sourcegraph.com/sourcegraph/appdash v0.0.0-20190107175209-d9ea5c54f7dc
)

replace (
	github.com/emurmotol/project/api => ./api
	github.com/emurmotol/project/auth => ./auth
	github.com/emurmotol/project/user => ./user
)
