package common

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"luckgo/common/util"
	"os"
	"strings"
)
var(
	AppConf = &AppConfig{}
	log = logrus.New()
)



func InitConfig()  {
	v := viper.New()
	v.AddConfigPath("./conf")
	//v.AddConfigPath("C:\\Users\\home\\go\\src\\luckgo\\conf\\")
	v.SetConfigType("ini")
	v.SetConfigFile("server.ini")

	err := v.ReadInConfig()
	if err != nil{
		fmt.Printf("加载配置文件错误，错误码[%s]\n",err.Error() )
		return
	}

   // var appConf AppConfig
	err = v.Unmarshal(&AppConf)
	if err != nil{
		fmt.Printf("配置文件序列号错误，错误码[%s]\n",err.Error() )
		return
	}
	fmt.Printf("logLevel[%s]\n",AppConf.LogConf.LogLevel)
	fmt.Printf("logPath[%s]\n",AppConf.LogConf.LogPath)

}
func transformLogLevel(logLevel string)  logrus.Level {
	switch logLevel {
	case "panic":
		return logrus.PanicLevel
	case "fata":
		return logrus.FatalLevel
	case "error":
		return logrus.ErrorLevel
	case "warn":
		return logrus.WarnLevel
	case "info":
		return logrus.InfoLevel
	case "debug":
		fmt.Println("1111")
		return logrus.DebugLevel
	case "trace":
		return logrus.TraceLevel
	default:
		return logrus.DebugLevel
	}
}

func InitLog()  {
	if util.IsPathExist(AppConf.LogConf.LogPath){
	//if(!util.isPathExist(AppConf.LogConf.LogPath)){
		err := os.MkdirAll(AppConf.LogConf.LogPath, os.ModePerm)
		if err != nil{
			fmt.Println(err)
			return
		}
	}
	logFilePath := AppConf.LogConf.LogPath
	if (!strings.HasSuffix(AppConf.LogConf.LogPath,"/")){
		logFilePath += "/"
	}
	logFilePath += AppConf.LogConf.LogFileName

	logFile, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil{
		log.Info("Failed to log to file, using default stderr")
	}
	log.Out = logFile
	log.Level = transformLogLevel(AppConf.LogConf.LogLevel)

	//log.WithFields(logrus.Fields{"filename":"123.txt"}).Info("info123")
	//log.WithFields(logrus.Fields{"filename":"123.txt"}).Debug("debug123")
	//log.WithFields(logrus.Fields{"filename":"123.txt"}).Trace("debug123")
}

func Init()  {
	InitConfig()
	InitLog()
}
