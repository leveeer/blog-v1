package logging

import (
	"context"
	nested "github.com/antonfisher/nested-logrus-formatter"
	"github.com/sirupsen/logrus"
	"os"
)

var Entry *logrus.Entry
var Logger *logrus.Logger

func init() {
	Logger = logrus.New()
	Logger.Out = os.Stdout
	//filePath := getLogFileFullPath()
	//file := openLogFile(filePath)
	//Logger.Out = file
	Logger.SetLevel(logrus.DebugLevel)
	Logger.SetReportCaller(true)
	f := &nested.Formatter{
		HideKeys:        true,
		FieldsOrder:     []string{},
		TimestampFormat: "2006/01/02 15:04:05",
	}
	Logger.SetFormatter(f)
	Entry = Logger.WithContext(context.Background())
}
