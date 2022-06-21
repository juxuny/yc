package router

type ValidatorHandler interface {
	Validate() error
}
