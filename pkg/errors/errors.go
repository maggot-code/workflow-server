/*
 * @FilePath: /workflow-server/pkg/errors/errors.go
 * @Author: maggot-code
 * @Date: 2023-08-14 16:15:41
 * @LastEditors: maggot-code
 * @LastEditTime: 2023-08-16 18:02:55
 * @Description:
 */
package errors

import (
	"errors"
	"fmt"
	"net/http"
)

type Error struct {
	Status
	cause error
}

func (e *Error) Error() string {
	return fmt.Sprintf("error: code = %d reason = %s message = %s metadata = %v cause = %v", e.Code, e.Reason, e.Message, e.Metadata, e.cause)
}

func (e *Error) Unwrap() error { return e.cause }

func (e *Error) Is(err error) bool {
	if se := new(Error); errors.As(err, &se) {
		// return se.Code == e.Code && se.Reason == e.Reason
		return se.Code == e.Code
	}

	return false
}

func (e *Error) WithCause(cause error) *Error {
	e.cause = cause
	return e
}

func (e *Error) WithReason(reason string) *Error {
	e.Reason = reason
	return e
}

func (e *Error) WithMessage(message string) *Error {
	e.Message = message
	return e
}

func New(code int, reason, message string) *Error {
	return &Error{
		Status: Status{
			Code:    int32(code),
			Reason:  reason,
			Message: message,
		},
	}
}

func Newf(code int, reason, format string, a ...interface{}) *Error {
	return New(code, reason, fmt.Sprintf(format, a...))
}

func Code(err error) int {
	if err == nil {
		return http.StatusOK
	}

	return int(FromError(err).Code)
}

func Reason(err error) string {
	if err == nil {
		return UnknownStatusReason
	}

	return FromError(err).Reason
}

func FromError(err error) *Error {
	if err == nil {
		return nil
	}

	if se := new(Error); errors.As(err, &se) {
		return se
	}

	return New(http.StatusOK, "ok", "ok")
}
