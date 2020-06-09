package setting

import (
	"fmt"
	"log"
	"time"

	"github.com/go-ini/ini"
)

var (
	// Cfg 配置文件
	Cfg *ini.File

	// RunMode 运行模式
	RunMode string

	// HTTPPort 端口
	HTTPPort int
	// ReadTimeout 读超时
	ReadTimeout time.Duration
	// WriteTimeout 写超时
	WriteTimeout time.Duration

	// PageSize 页容
	PageSize int
	// JWTSecret jwt密钥
	JWTSecret string
)

func init() {
	var err error
	Cfg, err = ini.Load("conf/app.ini")
	if err != nil {
		log.Fatalf("Fail to parse 'conf/app.ini': %v", err)
	}

	LoadBase()
	LoadServer()
	LoadApp()
}

// LoadBase 加载基本配置
func LoadBase() {
	RunMode = Cfg.Section("").Key("RUN_MODE").MustString("debug")
}

// LoadServer 加载服务配置
func LoadServer() {
	sec, err := Cfg.GetSection("server")
	if err != nil {
		log.Fatalf("Fail to get section 'server': %v", err)
	}

	HTTPPort = sec.Key("HTTP_PORT").MustInt(8000)
	ReadTimeout = sec.Key("READ_TIMEOUT").MustDuration(60 * time.Second)
	WriteTimeout = sec.Key("WRITE_TIMEOUT").MustDuration(60 * time.Second)
}

// LoadApp 加载app配置
func LoadApp() {
	sec, err := Cfg.GetSection("app")
	if err != nil {
		log.Fatalf("Fail to get section 'app': %v", err)
	}

	JWTSecret = sec.Key("JWT_SECRET").MustString("!@#$%")
	fmt.Println("JWTSecret", JWTSecret)
	PageSize = sec.Key("PAGE_SIZE").MustInt(10)
}
