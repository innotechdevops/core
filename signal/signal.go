package signal

import (
	"os"
	"os/signal"
	"syscall"
)

// Wait signal function
//
// 	app.Hooks().OnShutdown(func() error {
//		log.Info("On server shutting down")
//		return nil
//	})
//
//	go func() {
//		err := app.Listen(":9001")
//		if err != nil {
//			log.Fatal(err)
//		}
//	}()
//
//	log.Info("Running...")

//	signal.Wait(func(sig os.Signal) {
//			log.Info("Gracefully shutting down...")
//			log.Info("Waiting for all request to finish")
//			err := app.Shutdown()
//			if err != nil {
//				log.Fatal(err)
//			}
//			log.Info("Running cleanup tasks...")
//			log.Info("Server was successful shutdown.")
//		})
func Wait(fn func(os.Signal)) {
	termChan := make(chan os.Signal, 1)
	signal.Notify(termChan, os.Interrupt, syscall.SIGTERM)
	sig := <-termChan
	fn(sig)
}
