/*
 * @FilePath: /workflow-server/pkg/errors/types.go
 * @Author: maggot-code
 * @Date: 2023-08-16 15:25:06
 * @LastEditors: maggot-code
 * @LastEditTime: 2023-08-18 16:59:46
 * @Description:
 */
package errors

import "net/http"

const (
	UnknownStatusCode   = http.StatusInternalServerError
	UnknownStatusReason = ""
)

func StatusOK(reason, message string) *Error {
	return New(http.StatusOK, reason, message)
}

func IsStatusOK(err error) bool {
	return Code(err) == http.StatusOK
}

func StatusBadRequest(reason, message string) *Error {
	return New(http.StatusBadRequest, reason, message)
}

func IsStatusBadRequest(err error) bool {
	return Code(err) == http.StatusBadRequest
}

func StatusUnauthorized(reason, message string) *Error {
	return New(http.StatusUnauthorized, reason, message)
}

func IsStatusUnauthorized(err error) bool {
	return Code(err) == http.StatusUnauthorized
}

func StatusForbidden(reason, message string) *Error {
	return New(http.StatusForbidden, reason, message)
}

func IsStatusForbidden(err error) bool {
	return Code(err) == http.StatusForbidden
}

func StatusNotFound(reason, message string) *Error {
	return New(http.StatusNotFound, reason, message)
}

func IsStatusNotFound(err error) bool {
	return Code(err) == http.StatusNotFound
}

func StatusInternalServerError(reason, message string) *Error {
	return New(http.StatusInternalServerError, reason, message)
}

func IsStatusInternalServerError(err error) bool {
	return Code(err) == http.StatusInternalServerError
}

func StatusBadGateway(reason, message string) *Error {
	return New(http.StatusBadGateway, reason, message)
}

func IsStatusBadGateway(err error) bool {
	return Code(err) == http.StatusBadGateway
}

func StatusServiceUnavailable(reason, message string) *Error {
	return New(http.StatusServiceUnavailable, reason, message)
}

func IsStatusServiceUnavailable(err error) bool {
	return Code(err) == http.StatusServiceUnavailable
}

func StatusGatewayTimeout(reason, message string) *Error {
	return New(http.StatusGatewayTimeout, reason, message)
}

func IsStatusGatewayTimeout(err error) bool {
	return Code(err) == http.StatusGatewayTimeout
}
