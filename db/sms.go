package db

import (
	"errors"
	"project/db/models"
)

var ErrSmsNodeFound = errors.New("短信验证失败")

func InsertSms(u *models.SmsSend) error {
	if u == nil {
		return nil
	}
	_, err := dbEngine.Insert(u)
	return err
}

func QuerySms(phone string, code int) (*models.SmsSend, error) {
	var ss = &models.SmsSend{
		Phone: phone,
		Code:  code,
	}
	has, err := dbEngine.Get(ss)
	if !has {
		err = ErrSmsNodeFound
	}

	if err != nil {
		return nil, err
	}

	return ss, nil
}

func UpdateSmsUsed(s *models.SmsSend) error {
	if s == nil {
		return nil
	}

	s.Used = 1
	_, err := dbEngine.Where("id=?", s.Id).Cols("used").Update(s)
	return err
}
