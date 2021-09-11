package dao

import "errors"

var ErrNoRows = errors.New("sql: no rows in result set")

// MockQuerySql 模拟返回 ErrNoRows 错误
func MockQuerySql() error {
	return ErrNoRows
}
