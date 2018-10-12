package server

import (
	"fmt"
	"os"
	"runtime"
	"sync"

	"github.com/Unknwon/goconfig"
	"github.com/cihub/seelog"
	nats "github.com/nats-io/go-nats"
)

type Server struct {
	mu                sync.Mutex
	config            *goconfig.ConfigFile
	configFile        string
	logger            seelog.LoggerInterface
	opts              *Options
	optsMu            sync.RWMutex
	natsStreamingConn *nats.Conn
}

// New will setup a new server struct after parsing the options.
func New(opts *Options) *Server {
	processOptions(opts)

	s := &Server{
		configFile: opts.ConfigFile,
		opts:       opts,
	}

	// s.handleSignals()
	return s
}

func (s *Server) getOpts() *Options {
	s.optsMu.RLock()
	opts := s.opts
	s.optsMu.RUnlock()
	return opts
}

func (s *Server) setOpts(opts *Options) {
	s.optsMu.Lock()
	s.opts = opts
	s.optsMu.Unlock()
}

// Start up the server, this will block.
func (s *Server) Start() {

	s.logger.Info("Starting  smssrv version ", VERSION)

	s.startSms()

	runtime.Goexit()

}

// PrintAndDie is exported for access in other packages.
func PrintAndDie(msg string) {
	fmt.Fprintf(os.Stderr, "%s\n", msg)
	os.Exit(1)
}
