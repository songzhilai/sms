package server

import (
	"fmt"

	"github.com/cihub/seelog"
)

// ConfigureLogger 配置日志模块
func (s *Server) ConfigureLogger() {

	opts := s.getOpts()
	if opts.LogConfigFile != "" {
		logger, err := seelog.LoggerFromConfigAsFile(opts.LogConfigFile)
		if err != nil {
			fmt.Printf("日志配置文件加载失败：%s", err)
		} else {
			s.logger = logger
			defer s.logger.Flush()
		}
	}
}
