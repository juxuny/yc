package validator

type IValidator interface {
	Run(value interface{}) error
}
