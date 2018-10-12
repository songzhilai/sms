package server

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/cihub/seelog"
)

// option is a hot-swappable configuration setting.
type option interface {
	// Apply the server option.
	Apply(server *Server)
}

type logOption struct {
	option
}

// Apply the log change.
func (l *logOption) Apply(server *Server) {
	server.mu.Lock()
	opts := server.getOpts()
	logger, err := seelog.LoggerFromConfigAsFile(opts.LogConfigFile)
	if err != nil {
		fmt.Printf("日志配置文件加载失败：%s", err)
	} else {
		server.logger = logger
		defer server.logger.Flush()
	}

	server.mu.Unlock()
	fmt.Printf("Reloaded: log\n")
}

// type natsStreamingOption struct {
// 	option
// }

// // Apply the nats streaming change.
// func (n *natsStreamingOption) Apply(server *Server) {
// 	fmt.Printf("Apply: nats\n")
// 	if server.natsStreamingConn != nil {
// 		server.natsStreamingConn.Close()
// 		server.startNats()
// 	}
// 	fmt.Printf("Reloaded: nats\n")
// }

func (s *Server) Reload() error {

	newOpts, err := ProcessConfigFile(s.configFile)
	if err != nil {
		return err
	}
	processOptions(newOpts)
	if err := s.reloadOptions(newOpts); err != nil {
		return err
	}

	return nil
}

// reloadOptions 重新加载配置
func (s *Server) reloadOptions(newOpts *Options) error {

	changed, err := s.diffOptions(newOpts)
	if err != nil {
		return err
	}
	s.setOpts(newOpts)
	s.applyOptions(changed)
	return nil
}

func (s *Server) diffOptions(newOpts *Options) ([]option, error) {
	var (
		oldConfig = reflect.ValueOf(s.getOpts()).Elem()
		newConfig = reflect.ValueOf(newOpts).Elem()
		diffOpts  = []option{}
	)

	// log 默认重新加载
	diffOpts = append(diffOpts, &logOption{})

	for i := 0; i < oldConfig.NumField(); i++ {
		var (
			field    = oldConfig.Type().Field(i)
			oldValue = oldConfig.Field(i).Interface()
			newValue = newConfig.Field(i).Interface()
			changed  = !reflect.DeepEqual(oldValue, newValue)
		)
		if !changed {
			continue
		}
		switch strings.ToLower(field.Name) {
		// case "nats_streaming":
		// 	diffOpts = append(diffOpts, &natsStreamingOption{})
		default:
			fmt.Errorf("Config reload not supported for %s: old=%v, new=%v",
				field.Name, oldValue, newValue)
		}
	}
	return diffOpts, nil
}

func (s *Server) applyOptions(opts []option) {
	for _, opt := range opts {
		opt.Apply(s)
	}
	// fmt.Printf("Reloaded server configuration")
}
