package errors

import (
	"fmt"
	"reflect"
	"runtime"
)

const (
	CodeRequestParamsFormatError = 1000
	CodeInitConfigError = 1002
	CodeJsonUnmarshalError = 1003

	CodeInvalidRequestParam = 2000

	CodeDaoEmptyResult = 3000
	CodeDaoError = 3001
)

type ErrorInfo struct {
	Code int
	Msg  string
}


func (e *ErrorInfo) Error() string {
	return fmt.Sprintf("code: %d, msg: %s", e.Code, e.Msg)
}

func NewErrInfo(code int, msg string) *ErrorInfo {
	return &ErrorInfo{
		Code: code,
		Msg:  msg,
	}
}

func RequestParamsFormatError(err error) *ErrorInfo {
	return NewErrInfo(CodeRequestParamsFormatError, fmt.Sprintf("Request params format error: %v", err))
}

func InitConfigError(err error) *ErrorInfo {
	return NewErrInfo(CodeInitConfigError, fmt.Sprintf("Init config from file error: %v", err))
}

func InvalidRequestParam(param string, val interface{}) *ErrorInfo {
	return NewErrInfo(CodeInvalidRequestParam, fmt.Sprintf("Invalid request param: %s, value: '%s'", param, val))
}

func DaoEmptyResultError() *ErrorInfo {
	return NewErrInfo(CodeDaoEmptyResult, "empty result")
}

func ErrDaoError(fun interface{}, err error, format string, params ...interface{}) *ErrorInfo {
	return NewErrInfo(CodeDaoError, fmt.Sprintf("Dao error, fun: %s, param: %s, err: %v", getFuncName(fun), fmt.Sprintf(format, params...), err))
}

func ErrJsonUnmarshal(err error, data []byte) *ErrorInfo {
	return NewErrInfo(CodeJsonUnmarshalError, fmt.Sprintf("JsonUnMarshal Error, data: %s, err: %v", string(data), err))
}


func getFuncName(fun interface{}) string {
	switch reflect.TypeOf(fun).Kind() {
	case reflect.Func:
		return runtime.FuncForPC(reflect.ValueOf(fun).Pointer()).Name()
	default:
		return fmt.Sprintf("%v", fun)
	}
}