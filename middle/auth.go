package middle

import "context"

type AuthHandler struct{}

func (t *AuthHandler) Run(context context.Context) (isEnd bool, err error) {
	// TODO: verify validation of token
	return
}
