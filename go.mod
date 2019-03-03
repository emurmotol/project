module github.com/emurmotol/project

require (
	github.com/99designs/gqlgen v0.7.2
	github.com/emurmotol/project/api v0.0.0
	github.com/emurmotol/project/auth_api v0.0.0
	github.com/emurmotol/project/user_api v0.0.0
	github.com/go-kit/kit v0.8.0
	github.com/kujtimiihoxha/kit v0.0.0-20190129201007-e40b4aff8f4e // indirect
	github.com/pkg/errors v0.8.1 // indirect
	github.com/urfave/cli v1.20.0 // indirect
	github.com/vektah/gqlparser v1.1.1
)

replace (
	github.com/emurmotol/project/api => ./api
	github.com/emurmotol/project/auth_api => ./auth_api
	github.com/emurmotol/project/user_api => ./user_api
)
