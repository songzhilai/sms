package server

import (
	"github.com/Unknwon/goconfig"
)

type Options struct {
	ConfigFile          string
	LogConfigFile       string
	Mysql               MysqlOpts
	NatsStreaming       NatsStreamingOpts
	MysqlDataSourceName string
	ExcelName           string
	ExcelSheetName      string
	SmsContent          string
}

type MysqlOpts struct {
	Host     string
	User     string
	Password string
	Port     int
}

type NatsOpts struct {
	Host string
}

type NatsStreamingOpts struct {
	Host string
}

func processOptions(opts *Options) {
	// Setup non-standard Go defaults
	if opts.ConfigFile == "" {
		opts.ConfigFile = DEFAULT_CONFIG_FILE
	}
	if opts.LogConfigFile == "" {
		opts.LogConfigFile = DEFAULT_LOG_CONFIG_FILE
	}
}

func (o *Options) ProcessConfigFile(configFile string) error {
	o.ConfigFile = configFile
	if configFile == "" {
		return nil
	}
	config, err := goconfig.LoadConfigFile(configFile)
	if err != nil {
		return err
	}
	o.ExcelName = config.MustValue("base", "excelname", "test.xlsx")
	o.ExcelSheetName = config.MustValue("base", "excelSheetname", "Sheet1")
	o.SmsContent = config.MustValue("base", "smswxnum", "y781502032")
	// // mysql
	// o.Mysql.Host = config.MustValue("mysql", "host", "")
	// o.Mysql.User = config.MustValue("mysql", "user", "")
	// o.Mysql.Password = config.MustValue("mysql", "password", "")
	// o.Mysql.Port = config.MustInt("mysql", "port", 3306)
	// o.MysqlDataSourceName = fmt.Sprintf(
	// 	"%s:%s@tcp(%s:%d)/%s?charset=utf8",
	// 	o.Mysql.User,
	// 	o.Mysql.Password,
	// 	o.Mysql.Host,
	// 	o.Mysql.Port,
	// 	"imhdb")

	// //nats streaming
	// o.NatsStreaming.Host = config.MustValue("nats_streaming", "host", "")

	return nil
}

func ProcessConfigFile(configFile string) (*Options, error) {
	opts := &Options{}
	if err := opts.ProcessConfigFile(configFile); err != nil {
		return nil, err
	}
	return opts, nil
}
