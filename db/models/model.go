package models

import (
	"time"
)

type SmsSend struct {
	Id        int64
	UserId    int64  `xorm:"not null default 0 index BIGINT(20)"`
	Phone     string `xorm:"not null default '' index CHAR(20)"`
	Code      int    `xorm:"not null default 0 INT(10)"`
	Content   string `xorm:"not null default '' VARCHAR(512)"`
	Channel   string `xorm:"not null default '' CHAR(20)"`
	Success   int    `xorm:"not null default 0 index TINYINT(1)"`
	Extra     string `xorm:"not null default '' VARCHAR(512)"`
	Used      int    `xorm:"not null default 0 index TINYINT(1)"`
	CreatedAt int64  `xorm:"BIGINT(20) created"`
	UpdatedAt int64  `xorm:"not null BIGINT(20) updated"`
}

type User struct {
	Id          int64
	Phone       string `xorm:"not null default '' unique CHAR(20)"`
	Name        string `xorm:"not null default '' VARCHAR(100)"`
	Password    string `xorm:"not null VARCHAR(128)"`
	PayPassword string `xorm:"not null VARCHAR(128)"`
	Pid         int64  `xorm:"not null BIGINT(20)"`
	Pids        string `xorm:"not null TEXT"`
	Level       int    `xorm:"not null INT(10)"`

	MinerLevel int `xorm:"not null INT(10) comment('矿工等级')"`
	TeamLevel  int `xorm:"not null INT(10) comment('矿主等级')"`

	Token     string `xorm:"not null unique default '' VARCHAR(255)"`
	ShareCode string `xorm:"not null CHAR(10) unique"`
	State     int    `xorm:"not null TINYINT(1)"`

	Wallet      float64 `xorm:"not null DECIMAL(20,4) comment('矿池钱包')"`
	Pool        float64 `xorm:"not null DECIMAL(20,4)  comment('矿池资产')"`
	Money       float64 `xorm:"not null DECIMAL(20,4)  comment('可售余额')"`
	MaxSale     float64 `xorm:"not null DECIMAL(20,4)  comment('可售额度')"`
	MinerIncome float64 `xorm:"not null DECIMAL(20,4)  comment('矿主分红')"`
	ShareIncome float64 `xorm:"not null DECIMAL(20,4)  comment('分享收益')"`

	TotalIncome float64 `xorm:"not null DECIMAL(20,4)  comment('总收益')"`

	IsGift int `xorm:"not null TINYINT(1) comment('计算后台赠送使用')"`

	CreatedAt time.Time `xorm:"BIGINT(20) created"`
	UpdatedAt time.Time `xorm:"not null BIGINT(20) updated"`
}

type MineMachine struct {
	Id        int64
	Name      string    `xorm:"not null CHAR(50)"`
	Pic       string    `xorm:"not null VARCHAR(512)"`
	Price     int       `xorm:"not null int(10)"`
	Output    float64   `xorm:"not null DECIMAL(15,4) default(0.0000)"`
	Degree    float64   `xorm:"not null DECIMAL(10,2) COMMENT('活跃度')"`
	CreatedAt time.Time `xorm:"BIGINT(20) created"`
	UpdatedAt time.Time `xorm:"not null BIGINT(20) updated"`
}

type Notice struct {
	Id        int64
	Title     string `xorm:"not null VARCHAR(128)"`
	Content   string `xorm:"not null TEXT"`
	CreatedAt int64  `xorm:"not null BIGINT(20) created"`
	UpdatedAt int64  `xorm:"not null BIGINT(20) updated"`
}

type UserMachine struct {
	Id             int64
	Mid            int64     `xorm:"BIGINT(20)"`
	Uid            int64     `xorm:"BIGINT(20) index"`
	CreatedAt      time.Time `xorm:"BIGINT(20) created"`
	UpdatedAt      time.Time `xorm:"not null BIGINT(20) updated"`
	ExpiredAt      int64     `xorm:"not null BIGINT(20) default(0)"`
	LatestWithdraw int64     `xorm:"not null BIGINT(20)  default(0)"`
	Withdraw       float64   `xorm:"not null DECIMAL(20,4)  default(0.0000)"`
	IsGift         int       `xorm:"not null TINYINT(1) default(0) comment('是否系统赠送')"`
}
type UserGift struct {
	Id         int64
	Mid        int64     `xorm:"BIGINT(20)"`
	Uid        int64     `xorm:"BIGINT(20) index"`
	CreatedAt  time.Time `xorm:"not null BIGINT(20) created"`
	UpdatedAt  time.Time `xorm:"not null BIGINT(20) updated"`
	ReceivedAt int64     `xorm:"not null BIGINT(20)"`
}

type GiftConfig struct {
	Id        int64
	ShareNum  int       `xorm:"INT(10)"`
	Mid       int64     `xorm:"BIGINT(20)"`
	GiftMid   int64     `xorm:"BIGINT(20)"`
	GiftNum   int       `xorm:"BIGINT(20)"`
	CreatedAt time.Time `xorm:"not null BIGINT(20) created"`
	UpdatedAt time.Time `xorm:"not null BIGINT(20) updated"`
}

func (um *UserMachine) IsExpired() bool {
	expireTime := time.Unix(um.ExpiredAt, 0)
	return expireTime.Before(time.Now())
}

type UserVerify struct {
	Id        int64  `json:"-"`
	Uid       int64  `xorm:"not null BIGINT(20) index"`
	Name      string `xorm:"not null VARCHAR(32)"`
	CardId    string `xorm:"not null CHAR(18)"`
	Pic1      string `xorm:"not null VARCHAR(512)"`
	Pic2      string `xorm:"not null VARCHAR(512)"`
	State     int    `xorm:"not null TINYINT(1)"`
	CreatedAt int64  `xorm:"not null BIGINT(20) created"`
	UpdatedAt int64  `xorm:"not null BIGINT(20) updated"`
}

const (
	UserVerifyStateNormal = 0
	UserVerifyRejected    = -1
	UserVerifyConfirmed   = 1
)

type UserPayConfig struct {
	Id  int64 `json:"-"`
	Uid int64 `json:"-"xorm:"not null BIGINT(20) index"`

	AlipayName    string `json:"alipay_name" xorm:"not null VARCHAR(32)" json:"alipay_name"`
	AlipayAccount string `json:"alipay_account" xorm:"not null VARCHAR(32)" json:"alipay_account"`
	AlipayQr      string `json:"alipay_qr" xorm:"not null VARCHAR(512)" json:"alipay_qr"`

	WechatName    string `json:"wechat_name" xorm:"not null VARCHAR(32)" json:"wechat_name"`
	WechatAccount string `json:"wechat_account" xorm:"not null VARCHAR(32)" json:"wechat_account"`
	WechatQr      string `json:"wechat_qr" xorm:"not null VARCHAR(512)" json:"wechat_qr"`

	BankName    string `json:"bank_name" xorm:"not null VARCHAR(32)" json:"bank_name"`
	BankAccount string `json:"bank_account" xorm:"not null VARCHAR(100)" json:"bank_account"`
	BankUser    string `json:"bank_user" xorm:"not null VARCHAR(32)" json:"bank_user"`

	CreatedAt int64 `json:"-" xorm:"not null BIGINT(20) created"`
	UpdatedAt int64 `json:"-" xorm:"not null BIGINT(20) updated"`
}

type Comment struct {
	Id        int64  `json:"-"`
	Uid       int64  `xorm:"not null BIGINT(20) index" json:"uid"`
	Message   string `xorm:"not null VARCHAR(512)" json:"message"`
	Pic       string `xorm:"not null VARCHAR(512)" json:"pic"`
	Reply     string `xorm:"not null VARCHAR(512)" json:"reply"`
	CreatedAt int64  `xorm:"not null BIGINT(20) created" json:"created_at"`
	UpdatedAt int64  `xorm:"not null BIGINT(20) updated" json:"updated_at"`
}
type SignIn struct {
	Id        int64     `json:"-"`
	Uid       int64     `xorm:"not null BIGINT(20) index" json:"uid"`
	CreatedAt time.Time `xorm:"not null BIGINT(20) created" json:"created_at"`
}

type Order struct {
	Id         int64     `json:"-"`
	Uid        int64     `xorm:"not null BIGINT(20) index"`
	SellerUid  int64     `xorm:"not null BIGINT(20) index"`
	Num        int       `xorm:"not null INT(10)"`
	Pic        string    `xorm:"not null VARCHAR(512)"`
	Price      float64   `xorm:"not null DECIMAL(10,2)"`
	TotalPrice float64   `xorm:"not null DECIMAL(10,2)"`
	Fee        float64   `xorm:"not null DECIMAL(10,2) comment('手续费')"`
	State      int       `xorm:"not null TINYINT(1)"`
	CreatedAt  time.Time `xorm:"not null BIGINT(20) created" json:"created_at"`
	UpdatedAt  time.Time `xorm:"not null BIGINT(20) updated" json:"updated_at"`
}

const (
	OrderStateNormal       = 0 //初始待匹配
	OrderStateMatched      = 1 //已匹配
	OrderStatePayed        = 2 //玩家确认付款
	OrderStateConfirmPayed = 3 //卖家确认付款
)

type SystemConfig struct {
	Id int64

	MaxSale         float64 `xorm:"not null DECIMAL(10,2) comment('可售最大额度')"`
	GtsPrice        float64 `xorm:"not null DECIMAL(10,2) comment('gts价格')"`
	MaxOrder        int     `xorm:"not null INT(10) comment('每日交易订单最大数目')"`
	AccountMaxOrder int     `xorm:"not null INT(10) comment('每日账号交易订单最大数目')"`
	TradeFee        float64 `xorm:"not null DECIMAL(2,2) comment('交易手续费')"`
	ReRate          int     `xorm:"not null TINYINT(1) comment('复投比例')"`
	RemainGts       float64 `xorm:"not null DECIMAL(20,4) comment('gts剩余量')"`

	Open time.Time `xorm:"not null TIME"`
	End  time.Time `xorm:"not null TIME"`

	CreatedAt int64 `xorm:"not null BIGINT(20) created"`
	UpdatedAt int64 `xorm:"not null BIGINT(20) updated"`
}

type MoneyLog struct {
	Id  int64
	Uid int64   `xorm:"not null BIGINT(20)"`
	Num float64 `xorm:"not null DECIMAL(20,4) comment('金额')"`

	From    string `xorm:"VARCHAR(20)"`
	To      string `xorm:"VARCHAR(20)"`
	Comment string `xorm:"VARCHAR(200)"`

	CreatedAt time.Time `xorm:"not null BIGINT(20) created"`
	UpdatedAt time.Time `xorm:"not null BIGINT(20) updated"`
}

type IncomeLog struct {
	Id      int64
	Uid     int64   `xorm:"not null BIGINT(20)"`
	Num     float64 `xorm:"not null DECIMAL(20,4) comment('金额')"`
	Comment string  `xorm:"VARCHAR(200)"`

	CreatedAt time.Time `xorm:"not null BIGINT(20) created"`
	UpdatedAt time.Time `xorm:"not null BIGINT(20) updated"`
}

type PriceLog struct {
	Id    int64
	Price float64 `xorm:"not null DECIMAL(20,4) comment('价格')"`

	CreatedAt time.Time `xorm:"not null BIGINT(20) created"`
	UpdatedAt time.Time `xorm:"not null BIGINT(20) updated"`
}
type ShareConfig struct {
	Id        int64
	Rate1     float64   `xorm:"not null DECIMAL(5,5)"`
	Rate2     float64   `xorm:"not null DECIMAL(5,5)"`
	Rate3     float64   `xorm:"not null DECIMAL(5,5)"`
	Rate4     float64   `xorm:"not null DECIMAL(5,5)"`
	Rate5     float64   `xorm:"not null DECIMAL(5,5)"`
	Rate6     float64   `xorm:"not null DECIMAL(5,5)"`
	Rate7     float64   `xorm:"not null DECIMAL(5,5)"`
	CreatedAt time.Time `xorm:"not null BIGINT(20) created" json:"created_at"`
	UpdatedAt time.Time `xorm:"not null BIGINT(20) updated" json:"updated_at"`
}

type TeamConfig struct {
	Id        int64
	Level1    string    `xorm:"not null VARCHAR(200) default('')"`
	Level2    string    `xorm:"not null VARCHAR(200) default('')"`
	Level3    string    `xorm:"not null VARCHAR(200) default('')"`
	Level4    string    `xorm:"not null VARCHAR(200) default('')"`
	Level5    string    `xorm:"not null VARCHAR(200) default('')"`
	CreatedAt time.Time `xorm:"not null BIGINT(20) created" json:"created_at"`
	UpdatedAt time.Time `xorm:"not null BIGINT(20) updated" json:"updated_at"`
}

type WalletLog struct {
	Id           int64
	Uid          int64     `xorm:"not null BIGINT(20) index"`
	WalletName   string    `xorm:"not null CHAR(30) index"`
	Delta        float64   `xorm:"not null DECIMAL(20,5) comment('变动金额')"`
	ChangedValue float64   `xorm:"not null DECIMAL(20,5) comment('变动后余额')"`
	Comment      string    `xorm:"VARCHAR(200)"`
	CreatedAt    time.Time `xorm:"not null BIGINT(20) created"`
	UpdatedAt    time.Time `xorm:"not null BIGINT(20) updated"`
}
