package ginUtil

import (
	"github.com/sirupsen/logrus"
)

type ErrRespopnse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type ClientError interface {
	Error() string
	// ResponseBody returns response body.
	ResponseBody() interface{}
	// ResponseHeaders returns http status code and headers.
	ResponseHeaders() (int, map[string]string)
}

type HTTPError struct {
	Cause  error  `json:"-"`
	Detail string `json:"detail"`
	Status int    `json:"-"`
}

func (e *HTTPError) Error() string {
	if e.Cause == nil {
		return e.Detail
	}
	return e.Detail + " : " + e.Cause.Error()
}

// ResponseBody returns JSON response body.
func (e *HTTPError) ResponseBody() interface{} {
	return ErrRespopnse{
		Code:    -1,
		Message: e.Detail,
	}
}

// ResponseHeaders returns http status code and headers.
func (e *HTTPError) ResponseHeaders() (int, map[string]string) {
	return e.Status, map[string]string{
		"Content-Type": "application/json; charset=utf-8",
	}
}

func NewHTTPError(err error, status int, detail string) error {
	return &HTTPError{
		Cause:  err,
		Detail: detail,
		Status: status,
	}
}

func NewDefaultHTTPError(err error, detail string) error {
	if err != nil {
		logrus.Warnf("api request error:%+v", err)
	}
	return NewHTTPError(err, 400, detail)
}
