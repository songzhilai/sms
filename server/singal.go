package server

// Signal Handling
// func (s *Server) handleSignals() {

// 	c := make(chan os.Signal, 1)

// 	signal.Notify(c, syscall.SIGUSR2)

// 	go func() {
// 		for {
// 			<-c
// 			s.Reload()
// 		}
// 	}()
// }
