package gvalidator

const (
	localPhonePattern  = "^0[^0][0-9]+$"
	globalPhonePattern = "^\\+[^0][0-9]+$"
	phonePattern       = "((" + localPhonePattern + ")|(" + globalPhonePattern + "))"
)

type (
	PhoneNumberValidator struct {
		TagValidator
	}
)

func NewPhoneNumberValidator() *PhoneNumberValidator {
	validator := &PhoneNumberValidator{}
	validator.Initialize()
	return validator
}

func (validator *PhoneNumberValidator) Initialize() {
	initializeTagValidator("phone", validator.Validate)
}

func (validator *PhoneNumberValidator) Validate(str string) bool {
	return validateRegex(str, phonePattern)
}
