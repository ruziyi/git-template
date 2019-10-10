package models

import "time"

func (sc *SystemConfig) IsOpen() bool {
	now := time.Now()
	open := time.Date(now.Year(), now.Month(), now.Day(), sc.Open.Hour(), sc.Open.Minute(), sc.Open.Second(), 0, now.Location())
	end := time.Date(now.Year(), now.Month(), now.Day(), sc.End.Hour(), sc.End.Minute(), sc.End.Second(), 0, now.Location())
	return now.After(open) && now.Before(end)
}
