//logrus的使用
package main

import (
	"github.com/sirupsen/logrus"
	"os"
)

func init() {
	//日志格式化为JSON而不是默认的ASCII
	logrus.SetFormatter(&logrus.JSONFormatter{})

	//输出stdout而不是默认的stderr，也可以是一个文件
	logrus.SetOutput(os.Stdout)

	//只记录严重或以上警告
	logrus.SetLevel(logrus.InfoLevel)
}

func main() {
	//简单示例：
	logrus.WithFields(logrus.Fields{
		"animal": "walrus",
	}).Info("Info信息******")

	logrus.WithFields(logrus.Fields{
		"params1": "walrus",
		"params2": 10,
	}).Info("info信息...")

	logrus.WithFields(logrus.Fields{
		"params3": true,
		"params4": 122,
	}).Warn("warn信息...")

	logrus.WithFields(logrus.Fields{
		"params5": true,
		"params6": 100,
	}).Fatal("fatal信息...")
}
