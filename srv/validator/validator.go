package validator

type Validator struct {
	Errors map[string]string
}

func NewValidator() *Validator {
	return &Validator{
		Errors: make(map[string]string),
	}
}

func (v *Validator) SetError(key string, value string) {
	v.Errors[key] = value
}
