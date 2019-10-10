package db

import "errors"

var ErrModelNotFound = errors.New("model not found")
var ErrZeroLineAffected = errors.New("没有数据改动")
