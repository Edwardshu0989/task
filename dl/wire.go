// +build wireinject
// The build tag makes sure the stub is not built in the final build.
package dl

import (
	"awesomeProject/dao"
	"awesomeProject/server"
	"awesomeProject/service"
	"github.com/google/wire"
)

func initApp() (*App, func(), error) {
	panic(wire.Build(dao.Provider, dao.NewDB, server.New, service.New, NewApp))
	return &App{}, nil, nil
}
