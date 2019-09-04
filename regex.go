package gvalidator

import (
	"regexp"
)

type (
	RegexValidator struct {
		ParamValidator
	}
)

func NewRegexValidator() *RegexValidator {
	validator := &RegexValidator{}
	validator.Initialize()
	return validator
}

func (validator *RegexValidator) Initialize() {
	initializeParamValidator("regex", "^regex\\((.+)\\)$", validator.Validate)
}

func (validator *RegexValidator) Validate(str string, params ...string) bool {
	return validateRegex(str, params...)
}

func validateRegex(str string, params ...string) bool {
	if len(params) == 0 {
		logger.Tracef("Validate if value '%s' matches pattern ''", str)
		return true
	}
	pattern := params[0]
	logger.Tracef("Validate if value '%s' matches pattern '%s'", str, pattern)
	regex := regexp.MustCompile(pattern)
	return regex.MatchString(str)
}
