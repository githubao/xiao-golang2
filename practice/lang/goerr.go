// err 处理
// author: baoqiang
// time: 2018/12/21 下午9:45
package lang

import (
	"google.golang.org/grpc/status"
	"fmt"
	"strings"
	"google.golang.org/grpc/codes"
	"runtime"
	"path/filepath"
)

type ErrorCoder interface {
	Error() string
	Code() uint32
	Msg() string
	Where() string
}

type GRPCStatuser interface {
	GRPCStatus() *status.Status
	Error() string
}

// struct ErrorCode
type ErrorCode struct {
	code  uint32
	msg   string
	where string
}

// ErrorCode impl ErrorCoder
func (e *ErrorCode) Error() string {
	return fmt.Sprintf("code = %d, msg = %s", e.code, e.msg)
}

func (e *ErrorCode) Msg() string {
	return e.msg
}

func (e *ErrorCode) Code() uint32 {
	return e.code
}

func (e *ErrorCode) Where() string {
	return e.where
}

// new ErrorCoder
func New(msg string) *ErrorCode {
	where := caller(1, false)
	return &ErrorCode{code: 0, msg: msg, where: where}
}

func NewCoder(code uint32, msg string, extMsg ...string) *ErrorCode {
	if len(extMsg) != 0 {
		msg = strings.Join(extMsg, " : ") + " : " + msg
	}
	where := caller(1, false)
	return &ErrorCode{code: code, msg: msg, where: where}
}

func NewCoderWhere(code uint32, callDepth int, msg string, extMsg ...string) *ErrorCode {
	if len(extMsg) != 0 {
		msg = strings.Join(extMsg, " : ") + " : " + msg
	}
	where := caller(callDepth, false)
	return &ErrorCode{code: code, msg: msg, where: where}
}

func NewCoderErr(code uint32, err error, extMsg ...string) *ErrorCode {
	var msg string
	if err != nil {
		msg = err.Error()
	}

	if len(extMsg) != 0 {
		msg = strings.Join(extMsg, " : ") + " : " + msg
	}

	where := caller(1, false)
	return &ErrorCode{code: code, msg: msg, where: where}
}

//wrap err code, 根据不同的err code 返回新的err code
func Wrap(err error, extMsg ...string) *ErrorCode {
	var msg, where string
	var code uint32

	switch v := err.(type) {
	case ErrorCoder:
		msg = v.Msg()
		code = v.Code()
		where = v.Where()
	case GRPCStatuser:
		s := v.GRPCStatus()
		if s.Code() == codes.Unknown {
			code = 0
		} else if s.Code() < 20 {
			// grpc自带的错误是系统错误
			code = 500
		} else {
			code = uint32(s.Code())
		}
		msg = s.Message()
		where = caller(1, false)
	default:
		msg = v.Error()
		code = 0
		where = caller(1, false)
	}

	if len(extMsg) != 0 {
		msg = strings.Join(extMsg, " : ") + " : " + msg
	}

	return &ErrorCode{code: code, msg: msg, where: where}
}

func caller(callDepth int, short bool) string {
	_, file, line, ok := runtime.Caller(callDepth + 1)
	if !ok {
		file = "???"
		line = 0
	} else if short {
		file = filepath.Base(file)
	}
	return fmt.Sprintf("%s:%d", file, line)
}
