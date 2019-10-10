package db

import (
	"project/db/models"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	log "github.com/sirupsen/logrus"
)

const defaultMaxConns = 10

var (
	dbEngine *xorm.Engine
	logger   *log.Entry
)

type options struct {
	showSQL      bool
	maxOpenConns int
	maxIdleConns int
}

// ModelOption specifies an option for dialing a xordefaultModel.
type ModelOption func(*options)

// MaxIdleConns specifies the max idle connect numbers.
func MaxIdleConns(i int) ModelOption {
	return func(opts *options) {
		opts.maxIdleConns = i
	}
}

// MaxOpenConns specifies the max open connect numbers.
func MaxOpenConns(i int) ModelOption {
	return func(opts *options) {
		opts.maxOpenConns = i
	}
}

// ShowSQL specifies the buffer size.
func ShowSQL(show bool) ModelOption {
	return func(opts *options) {
		opts.showSQL = show
	}
}

func envInit() {
	// 定时ping数据库, 保持连接池连接
	go func() {
		ticker := time.NewTicker(time.Minute * 5)
		for {
			select {
			case <-ticker.C:
				dbEngine.Ping()
			}
		}
	}()
}

//New create the DbEngine's connection
func MustStartup(dsn string, opts ...ModelOption) func() {
	logger = log.WithField("component", "model")
	settings := &options{
		maxIdleConns: defaultMaxConns,
		maxOpenConns: defaultMaxConns,
		showSQL:      true,
	}

	// options handle
	for _, opt := range opts {
		opt(settings)
	}

	logger.Infof("DSN=%s ShowSQL=%t MaxIdleConn=%v MaxOpenConn=%v", dsn, settings.showSQL, settings.maxIdleConns, settings.maxOpenConns)

	// create DbEngine instance
	if db, err := xorm.NewEngine("mysql", dsn); err != nil {
		panic(err)
	} else {
		dbEngine = db
	}

	// 设置日志相关
	dbEngine.SetLogger(&Logger{Entry: logger.WithField("orm", "xorm")})
	// options
	dbEngine.SetMaxIdleConns(settings.maxIdleConns)
	dbEngine.SetMaxOpenConns(settings.maxOpenConns)
	dbEngine.ShowSQL(settings.showSQL)
	dbEngine.ShowExecTime(true)

	//enableCache()
	syncSchema()
	envInit()

	closer := func() {
		dbEngine.Close()
		logger.Info("stopped")
	}

	return closer
}

func syncSchema() {
	err := dbEngine.StoreEngine("InnoDB").Sync2(
		new(models.SmsSend),
		new(models.User),
		new(models.Notice),
		new(models.MineMachine),
		new(models.UserMachine),
		new(models.UserPayConfig),
		new(models.Comment),
		new(models.SignIn),
		new(models.Order),
		new(models.SystemConfig),
		new(models.UserVerify),
		new(models.MoneyLog),
		new(models.IncomeLog),
		new(models.PriceLog),
		new(models.ShareConfig),
		new(models.TeamConfig),
		new(models.UserGift),
		new(models.GiftConfig),
		new(models.WalletLog),
	)
	if err != nil {
		logger.Println("sync db error: ", err)
	}
}

func GetEngine() *xorm.Engine {
	return dbEngine
}
