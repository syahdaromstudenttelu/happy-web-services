package lib

import (
	"happy-order-service/helper"

	"github.com/jaevor/go-nanoid"
)

func GetRandomStdId(length uint) string {
	idStd, err := nanoid.Standard(int(length))
	helper.DoPanicIfError(err)

	return idStd()
}
