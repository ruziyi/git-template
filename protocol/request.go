package protocol

import (
	"github.com/astaxie/beego/validation"
)

type (
	EmptyRequest struct {
	}
	QueryIdRequest struct {
		Id int64 `query:"id"`
	}
	PostIdRequest struct {
		Id int64 `form:"id"`
	}
	BuyerConfirmPay struct {
		Id  int64  `form:"id"`
		Pic string `form:"pic"`
	}
	MatchOrderRequest struct {
		Id       int64  `form:"id"`
		Password string `form:"password"`
	}
	HireMachineRequest struct {
		Password string `form:"password"`
		Id       int64  `form:"id"`
	}
	SendSmsRequest struct {
		Phone string `json:"phone" form:"phone"`
	}
	RegisterRequest struct {
		Phone       string `json:"phone" form:"phone"`
		Code        int    `json:"code" form:"code"`
		Password    string `json:"password" form:"password"`
		PayPassword string `json:"pay_password" form:"pay_password"`
		ShareCode   string `json:"share_code" form:"share_code"`
	}
	ForgetPasswordRequest struct {
		Phone           string `json:"phone" form:"phone"`
		Code            int    `json:"code" form:"code"`
		Password        string `json:"password" form:"password"`
		PasswordConfirm string `json:"password_confirm" form:"password_confirm"`
		Type            int    `form:"type"`
	}
	LoginRequest struct {
		Phone    string `json:"phone" form:"phone"`
		Password string `json:"password" form:"password"`
	}
	UpdateBankRequest struct {
		BankName    string `json:"bank_name" form:"bank_name"`
		BankAccount string `json:"bank_account" form:"bank_account"`
		BankUser    string `json:"bank_user" form:"bank_user"`
	}
	IdsRequest struct {
		List string `form:"list" json:"list"`
	}
	OrderRequest struct {
		Num      int    `form:"num"`
		Password string `form:"password"`
	}
	TransferRequest struct {
		Num      float64 `form:"num"`
		Password string  `form:"password"`
	}
	TransferLogRequest struct {
		From int `query:"from"`
	}
	UidRequest struct {
		Uid int64 `query:"uid"`
	}
	ShareConfigRequest struct {
		Rate1 string `json:"rate1"`
		Rate2 string `json:"rate2"`
		Rate3 string `json:"rate3"`
		Rate4 string `json:"rate4"`
		Rate5 string `json:"rate5"`
		Rate6 string `json:"rate6"`
	}
)

func (req *SendSmsRequest) Valid(valid *validation.Validation) {
	valid.Length(req.Phone, 11, "phone").Message("手机格式错误")
}

func (req *RegisterRequest) Valid(valid *validation.Validation) {
	valid.Length(req.Phone, 11, "phone").Message("手机格式错误")
	valid.Required(req.Code, "code").Message("请输入验证码")
	valid.Required(req.Password, "password").Message("请输入密码")
	valid.Required(req.PayPassword, "pay_password").Message("请输入交易密码")
	valid.Required(req.ShareCode, "share_code").Message("请输入邀请码")
}
func (req *LoginRequest) Valid(valid *validation.Validation) {
	valid.Length(req.Phone, 11, "phone").Message("手机格式错误")
	valid.Required(req.Password, "password").Message("请输入密码")
}
func (req *UpdateBankRequest) Valid(valid *validation.Validation) {
	valid.Required(req.BankName, "bank_name").Message("请输入银行名称")
	valid.Required(req.BankAccount, "bank_account").Message("请输入银行账号")
	valid.Required(req.BankUser, "bank_user").Message("请输入持卡人姓名")
}
func (req *OrderRequest) Valid(valid *validation.Validation) {
	if req.Num < 0 || req.Num%10 > 0 {
		valid.SetError("num", "购买数量必须为10的倍数")
	}
}
func (req *TransferRequest) Valid(valid *validation.Validation) {
	if req.Num < 0 {
		valid.SetError("num", "金额不能小于0")
	}
}
func (req *TransferLogRequest) Valid(valid *validation.Validation) {
	if req.From != 1 && req.From != 2 && req.From != 3 {
		valid.SetError("from", "参数错误")
	}
}
func (req *HireMachineRequest) Valid(valid *validation.Validation) {
	valid.Required("password", "请输入交易密码")
}
