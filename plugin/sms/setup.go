package sms

import (
	"github.com/spf13/viper"
	"project/basic"
	"project/pkg/sms"
)

type SmsClient interface {
	Send(phones string, content string) error
}

var client SmsClient

func init() {
	basic.AppendInitFn(setUp)
}

func setUp() {
	client = sms.NewLexinClient(
		viper.GetString("lexin.user"),
		viper.GetString("lexin.password"),
		viper.GetString("lexin.sign"),
	)
}

func GetClient() SmsClient {
	return client
}
