package gvalidator

const (
	uuidPattern = "^[0-9a-fA-F]+-[0-9a-fA-F]+-[0-9a-fA-F]+-[0-9a-fA-F]+-[0-9a-fA-F]+$"
)

type (
	SimpleUUIDValidator struct {
		TagValidator
	}
)

func NewSimpleUUIDValidator() *SimpleUUIDValidator {
	validator := &SimpleUUIDValidator{}
	validator.Initialize()
	return validator
}

func (validator *SimpleUUIDValidator) Initialize() {
	initializeTagValidator("simple_uuid", validator.Validate)
}

func (validator *SimpleUUIDValidator) Validate(str string) bool {
	return validateRegex(str, uuidPattern)
}
