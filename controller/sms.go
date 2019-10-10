package controller

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"math/rand"
	"project/db"
	"project/db/models"
	ginUtil "project/pkg/gin_util"
	"project/plugin/sms"
	"project/protocol"
	"time"
)

const CodeExpire = 60 * 5

var (
	ErrPhone      = errors.New("手机号码格式错误")
	ErrCodeExpire = errors.New("短信验证码已过期")
)

var (
	SendSms = ginUtil.WrapRequestHandler(sendSms)
)

/**
 * showdoc
 * @catalog 账号相关
 * @title 短信验证码
 * @description 短信验证码接口
 * @method post
 * @url /api/sendSms
 * @param phone 必选 string 手机号
 * @return {"message":"success"}
 * @return_param message string 成功信息
 * @remark
 */
func sendSms(ctx *gin.Context, data *protocol.SendSmsRequest) (err error, res *protocol.StringResponse) {
	code := rand.Intn(9000) + 1000
	content := fmt.Sprintf("你的短信验证码为%d, %d分钟内有效", code, CodeExpire/60)
	err = sms.GetClient().Send(data.Phone, content)

	successCode := 0
	extra := ""

	if err != nil {
		successCode = 1
		extra = err.Error()
	}
	smsSend := &models.SmsSend{
		Phone:   data.Phone,
		Code:    code,
		Content: content,
		Channel: "lexin",
		Success: successCode,
		Extra:   extra,
	}
	dbErr := db.InsertSms(smsSend)
	if dbErr != nil {
		logrus.Warnf("保存短信失败：%s", dbErr.Error())
	}
	if err != nil {
		return ginUtil.NewDefaultHTTPError(err, "短信发送失败"), nil
	}
	return nil, protocol.SuccessResponse
}

func checkCode(phone string, code int) (*models.SmsSend, error) {
	s, err := db.QuerySms(phone, code)
	if err != nil {
		return s, err
	}
	if time.Now().Unix()-s.CreatedAt > CodeExpire || s.Used == 1 {
		return s, ErrCodeExpire
	}
	return s, nil
}
