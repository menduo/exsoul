// Created by @menduo @ 2020/7/14
package exsoul

import (
	"errors"
	"fmt"
)

var (
	ErrorOutOfRange   = errors.New("[ExSoul] Out of range")
	ErrorInvalidInput = errors.New("[ExSoul] Invalid Input")
	ErrorParam        = errors.New("[ExSoul] Invalid Param")
)

func newInvalidInputError(format string, values ...interface{}) error {
	format = fmt.Sprintf("%s: %s", ErrorInvalidInput, format)
	return errors.New(fmt.Sprintf(format, values...))
}

func newParamError(format string, values ...interface{}) error {
	format = fmt.Sprintf("%s: %s", ErrorParam, format)
	return errors.New(fmt.Sprintf(format, values...))
}

func newParamErrorWithError(prefix string, err error) error {
	msg := fmt.Sprintf("%s: %s %s", ErrorParam, prefix, err.Error())
	return errors.New(msg)
}

func newError(format string, values ...interface{}) error {
	format = fmt.Sprintf("%s: %s", "[ExSoul]", format)
	return errors.New(fmt.Sprintf(format, values...))
}
