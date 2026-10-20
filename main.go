package main

import (
	"log"
	"os"

	"github.com/anandki1209/gator/internal/config"
)

type state struct {
	cfg *config.Config
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	programState := state{
		cfg: &cfg,
	}

	cmds := commands{
		registeredCommands: make(map[string]func(*state, command) error),
	}

	cmds.register("login", handlerLogin)

	args := os.Args
	if len(args) < 2 {
		log.Fatalf("No command name provided")
	}

	cmdName := args[1]
	cmdArgs := args[2:]

	cmd := command{
		Name: cmdName,
		Args: cmdArgs,
	}

	if err := cmds.run(&programState, cmd); err != nil {
		log.Fatalf("error in trying to run this command: %v", err)
	}

}
