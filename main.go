package main

import (
	"fleetgo/logger"
	//"fmt"
)

func main() {
	var logPath *string = nil
	logger.InitLogger("info", logPath, "", 100, 30, 7)

	log := logger.GetLogger()

	log.Info("Hello FleetGo")

}
