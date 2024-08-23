//go:build wireinject
// +build wireinject

package service

import (
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(wire.FieldsOf(new(*Services), "Books"))
