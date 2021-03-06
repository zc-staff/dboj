package util

import "time"

// global config
const (
	ListenAddr    = ":8087"
	Database      = "dboj:dboj@tcp(127.0.0.1:3306)/dboj?charset=utf8&parseTime=true"
	SessionExpire = 3600 * time.Second
	Salt1         = "807dbeb2318b7ba9343158b6a6e1b50d"
	RedirectCode  = 302
	CacheMem      = 32 << 20
	SidName       = "dboj-session"
	PageSize      = 15
	PageDelta     = 3

	Pending     = 0
	Accepted    = 1
	WrongAnswer = 2
	TimeLimit   = 3
	SystemError = 4
)
