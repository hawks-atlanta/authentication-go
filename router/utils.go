package router

import "fmt"

const (
	RootRoute            = "/"
	EchoRoute            = "/echo"
	LoginRoute           = "/login"
	RegisterRoute        = "/register"
	ChallengeRoute       = "/challenge"
	AccountPasswordRoute = "/account/password"
)

const (
	AuthorizationHeader = "Authorization"
)

const (
	SessionVariale = "SESSION"
)

type Result struct {
	Succeed bool   `json:"succeed"`
	Message string `json:"msg"`
}

var UnauthorizedResult = Result{Succeed: false, Message: "unauthorized"}

func InternalServerError(err error) Result {
	return Result{Succeed: false, Message: err.Error()}
}

func SucceedResult(msg string) Result {
	return Result{Succeed: true, Message: msg}
}

func Bearer(tok string) string {
	return fmt.Sprintf("Bearer %s", tok)
}
