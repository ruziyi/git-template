package protocol

var SuccessResponse = &StringResponse{Message: "操作成功"}

type (
	StringResponse struct {
		Message string `json:"message"`
	}
	NoticeResponse struct {
		Id        int64  `json:"id"`
		Title     string `json:"title"`
		Content   string `json:"content"`
		CreatedAt int64  `json:"created_at"`
	}
	LoginResponse struct {
		Uid        int64  `json:"uid"`
		Token      string `json:"token"`
		Phone      string `json:"phone"`
		MinerLevel int    `json:"miner_level"`
		TeamLevel  int    `json:"team_level"`
	}
	VerifyInfoResponse struct {
		Name   string `json:"name"`
		CardId string `json:"card_id"`
		Pic1   string `json:"pic1"`
		Pic2   string `json:"pic2"`
		State  int    `json:"state"`
	}
	TeamInfo struct {
		Uid        int64  `json:"uid"`
		Phone      string `json:"phone"`
		MinerLevel int    `json:"miner_level"`
		TeamLevel  int    `json:"team_level"`
		ShareCount int    `json:"share_count"`
		TeamCount  int    `json:"team_count"`
	}
	MachineInfo struct {
		Id     int64   `json:"id"`
		Pic    string  `json:"pic"`
		Name   string  `json:"name"`
		Price  int     `json:"price"`
		Output float64 `json:"output"`
	}
	MyMachineInfo struct {
		Id          int64   `json:"id"`
		Pic         string  `json:"pic"`
		Name        string  `json:"name"`
		Output      float64 `json:"output"`
		Created     float64 `json:"created"`
		Withdraw    float64 `json:"withdraw"`
		RunningTime int64   `json:"running_time"`
		Expired     bool    `json:"expired"`
		CountDown   int     `json:"count_down"`
		IsGift      int     `json:"is_gift"`
	}
	NormalMachinesResponse struct {
		Total  float64         `json:"total"`
		Remain float64         `json:"remain"`
		List   []MyMachineInfo `json:"list"`
	}
	TradeOrderItem struct {
		OrderId int64   `json:"order_id"`
		Num     int     `json:"num"`
		Price   float64 `json:"price"`
		Total   float64 `json:"total"`
		State   int     `json:"state"`
	}
	NotMatchOrderItem struct {
		OrderId int64   `json:"order_id"`
		Uid     int64   `json:"uid"`
		Num     int     `json:"num"`
		Price   float64 `json:"price"`
	}
	TraderOrderInfo struct {
		Id     int64   `json:"id"`
		Num    int     `json:"num"`
		Price  float64 `json:"price"`
		Total  float64 `json:"total"`
		State  int     `json:"state"`
		PayPic string  `json:"pay_pic"`

		BuyerUid   int64  `json:"buyer_uid"`
		BuyerPhone string `json:"buyer_phone"`

		SellerUid   int64  `json:"seller_uid"`
		SellerPhone string `json:"seller_phone"`

		SellerAliAccount  string `json:"seller_ali_account"`
		SellerAliPic      string `json:"seller_ali_pic"`
		SellerWxAccount   string `json:"seller_wx_account"`
		SellerWxPic       string `json:"seller_wx_pic"`
		SellerBankAccount string `json:"seller_bank_account"`
		SellerBankName    string `json:"seller_bank_name"`
	}

	MyFundResponse struct {
		Income      float64 `json:"income"`       //总收益
		SaleMoney   float64 `json:"sale_money"`   //可售余额
		SaleAmount  float64 `json:"sale_amount"`  //可售额度
		MinerIncome float64 `json:"miner_income"` //矿主分红
		ShareIncome float64 `json:"share_income"` //分享收益
		MineWallet  float64 `json:"mine_wallet"`  //矿池钱包
		MinePool    float64 `json:"mine_pool"`    //矿池资产
	}
	TransferLog struct {
		Time    int64   `json:"time"`
		Num     float64 `json:"num"`
		Comment string  `json:"comment"`
	}
	UserInfo struct {
		Uid        int64 `json:"uid"`
		MinerLevel int   `json:"miner_level"`
		TeamLevel  int   `json:"team_level"`

		ShareCount int     `json:"share_count"`
		TeamCount  int     `json:"team_count"`
		TeamDegree float64 `json:"team_degree"`

		Verified string `json:"verified"`
	}

	PriceInfoResponse struct {
		Price    float64    `json:"price"`
		Increase float64    `json:"increase"`
		List     []DayPrice `json:"list"`
	}
	DayPrice struct {
		Time  int64   `json:"time"`
		Price float64 `json:"price"`
	}
	NowPrice struct {
		Price float64 `json:"price"`
	}

	CountDownInfo struct {
		Time int `json:"time"`
	}
	RemainResponse struct {
		Num float64 `json:"num"`
	}
)
