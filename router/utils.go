package router

import "fmt"

const (
	UsernameParam = "username"
)

const (
	RootRoute               = "/"
	EchoRoute               = "/echo"
	LoginRoute              = "/login"
	RegisterRoute           = "/register"
	ChallengeRoute          = "/challenge"
	AccountPasswordRoute    = "/account/password"
	LogsRoute               = "/logs"
	UserRoute               = "/user"
	LogsUserRoute           = LogsRoute + UserRoute
	DateRoute               = "/date"
	LogsDateRoute           = LogsRoute + DateRoute
	UserUUIDRoute           = "/user/uuid"
	UserUUIDWithParamsRoute = UserUUIDRoute + "/:" + UsernameParam
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
