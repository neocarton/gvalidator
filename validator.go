package gvalidator

import (
	"regexp"

	"github.com/asaskevich/govalidator"
)

type (
	SimpleValidatorFunc func(str string) bool

	ValidatorFunc func(str string, params ...string) bool

	TagValidator interface {
		Initialize()
		Validate(str string) bool
	}

	ParamValidator interface {
		Initialize()
		Validate(str string, params ...string) bool
	}
)

func initializeTagValidator(name string, validator SimpleValidatorFunc) {
	govalidator.TagMap[name] = govalidator.Validator(validator)
}

func initializeParamValidator(name string, pattern string, validator ValidatorFunc) {
	govalidator.ParamTagMap[name] = govalidator.ParamValidator(validator)
	govalidator.ParamTagRegexMap[name] = regexp.MustCompile(pattern)
}
