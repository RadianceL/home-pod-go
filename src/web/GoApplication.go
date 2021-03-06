package application

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	models "go-gin-template/src/pkg"
	"go-gin-template/src/pkg/config"
	AppConfig "go-gin-template/src/web/config"
	"go-gin-template/src/web/middleware"
	routers "go-gin-template/src/web/router"
	"go-gin-template/support/convert"
	"go-gin-template/support/logger"
	"net/http"
	"time"
)

// Run 运行
func Run(configPath string) {
	// 获取配置路径
	if configPath == "" {
		configPath = "./config.yaml"
	}
	// 初始化日志配置
	logger.InitLog("debug", "./data/log/log.log")
	// 加载配置
	loadConfig, err := AppConfig.LoadConfig(configPath)
	if err != nil {
		panic(err)
	}
	// 初始化数据
	initDB(loadConfig)
	// 初始化web服务
	initWeb(loadConfig)
}

func initDB(config *config.Config) {
	models.InitDB(config)
	logger.Debug("数据库加载完成.......")
}

func initWeb(loadConfig *config.Config) {
	gin.SetMode(gin.DebugMode)
	app := gin.New()
	app.NoRoute(middleware.NoRouteHandler())
	app.NoMethod(middleware.NoMethodHandler())
	// 崩溃恢复
	app.Use(middleware.RecoveryMiddleware())
	// 注册路由
	routers.RegisterRouter(app)
	go initHTTPServer(loadConfig, app)
}

// InitHTTPServer 初始化http服务
func initHTTPServer(config *config.Config, handler http.Handler) {
	srv := &http.Server{
		Addr:         ":" + convert.ToString(config.Web.Port),
		Handler:      handler,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}
	_ = srv.ListenAndServe()
}
