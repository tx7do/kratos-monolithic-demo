// Code generated by protoc-gen-go-errors. DO NOT EDIT.

package servicev1

import (
	fmt "fmt"
	errors "github.com/go-kratos/kratos/v2/errors"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
const _ = errors.SupportPackageIsVersion1

// 401
func IsNotLoggedIn(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == UserErrorReason_NOT_LOGGED_IN.String() && e.Code == 401
}

// 401
func ErrorNotLoggedIn(format string, args ...interface{}) *errors.Error {
	return errors.New(401, UserErrorReason_NOT_LOGGED_IN.String(), fmt.Sprintf(format, args...))
}

// 用户ID无效
func IsInvalidUserid(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == UserErrorReason_INVALID_USERID.String() && e.Code == 400
}

// 用户ID无效
func ErrorInvalidUserid(format string, args ...interface{}) *errors.Error {
	return errors.New(400, UserErrorReason_INVALID_USERID.String(), fmt.Sprintf(format, args...))
}

// 密码无效
func IsInvalidPassword(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == UserErrorReason_INVALID_PASSWORD.String() && e.Code == 400
}

// 密码无效
func ErrorInvalidPassword(format string, args ...interface{}) *errors.Error {
	return errors.New(400, UserErrorReason_INVALID_PASSWORD.String(), fmt.Sprintf(format, args...))
}

// token过期
func IsTokenExpired(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == UserErrorReason_TOKEN_EXPIRED.String() && e.Code == 400
}

// token过期
func ErrorTokenExpired(format string, args ...interface{}) *errors.Error {
	return errors.New(400, UserErrorReason_TOKEN_EXPIRED.String(), fmt.Sprintf(format, args...))
}

// token无效
func IsInvalidToken(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == UserErrorReason_INVALID_TOKEN.String() && e.Code == 400
}

// token无效
func ErrorInvalidToken(format string, args ...interface{}) *errors.Error {
	return errors.New(400, UserErrorReason_INVALID_TOKEN.String(), fmt.Sprintf(format, args...))
}

// token不存在
func IsTokenNotExist(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == UserErrorReason_TOKEN_NOT_EXIST.String() && e.Code == 404
}

// token不存在
func ErrorTokenNotExist(format string, args ...interface{}) *errors.Error {
	return errors.New(404, UserErrorReason_TOKEN_NOT_EXIST.String(), fmt.Sprintf(format, args...))
}

// 用户不存在
func IsUserNotExist(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == UserErrorReason_USER_NOT_EXIST.String() && e.Code == 404
}

// 用户不存在
func ErrorUserNotExist(format string, args ...interface{}) *errors.Error {
	return errors.New(404, UserErrorReason_USER_NOT_EXIST.String(), fmt.Sprintf(format, args...))
}
