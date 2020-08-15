//+build wireinject

package main

import (
	"github.com/google/wire"
	"payment/pkg"
)

func SetupApplication() (pkg.Application, error) {
	wire.Build(pkg.Container)
	return pkg.Application{}, nil
}
