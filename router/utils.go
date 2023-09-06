package router

const (
	EchoRoute     = "/echo"
	LoginRoute    = "/login"
	RegisterRoute = "/register"
)

type Result struct {
	Succeed bool   `json:"succeed"`
	Message string `json:"msg"`
}

var UnauthorizedResult = Result{Succeed: false, Message: "unauthorized"}

func InternalServerError(err error) Result {
	return Result{Succeed: false, Message: err.Error()}
}
