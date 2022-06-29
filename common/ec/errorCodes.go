package ec

import "google.golang.org/grpc/status"

var (
	// 1000~1999 系统统一错误
	ParamInvalid = status.Error(10000, "入参错误")

	// 10000~19999 业务错误

)
