package validation

type Validation interface {
	ValidateStruct(s interface{}) error
}
