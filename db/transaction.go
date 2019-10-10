package db

import "github.com/go-xorm/xorm"

func NewTransaction(fn func(session *xorm.Session) error) error {
	session := GetEngine().NewSession()
	err := session.Begin()
	defer session.Close()
	if err != nil {
		return err
	}
	err = fn(session)
	if err != nil {
		session.Rollback()
		return err
	}
	return session.Commit()
}
