package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/anandki1209/gator/internal/config"
	"github.com/anandki1209/gator/internal/database"

	_ "github.com/lib/pq"
)

type state struct {
	db  *database.Queries
	cfg *config.Config
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	db, err := sql.Open("postgres", cfg.DbURL)
	if err != nil {
		log.Fatalf("error opening database connection: %v", err)
	}
	defer db.Close()
	dbQueries := database.New(db)

	programState := state{
		db:  dbQueries,
		cfg: &cfg,
	}

	cmds := commands{
		registeredCommands: make(map[string]func(*state, command) error),
	}

	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)
	cmds.register("reset", handlerReset)
	cmds.register("users", handlerListUsers)
	cmds.register("agg", handlerAgg)
	cmds.register("addfeed", handlerAddFeed)
	cmds.register("feeds", handlerGetFeeds)
	cmds.register("follow", handlerFollow)
	cmds.register("following", handlerFollowing)

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
