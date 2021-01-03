package main


import (
	log "github.com/sirupsen/logrus"
	"github.com/MustaphaAlioglou/DEIMOS/shell"
)
func main(){
	log.Info("Checking logging: OK")
	log.Error("Checking logging: OK")
	shell.GetShell();

}