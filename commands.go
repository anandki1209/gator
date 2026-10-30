package main

import (
	"context"
	"fmt"

	"github.com/anandki1209/gator/internal/database"
)

type command struct {
	Name string
	Args []string
}

type commands struct {
	registeredCommands map[string]func(*state, command) error
}

func (c *commands) run(s *state, cmd command) error {
	if handler, ok := c.registeredCommands[cmd.Name]; ok {
		return handler(s, cmd)
	}
	return fmt.Errorf("unknown command: %s", cmd.Name)

}

func (c *commands) register(name string, f func(*state, command) error) {
	c.registeredCommands[name] = f

}

// middleware for handlers that rely on logged in users
func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(*state, command) error {
	return func(s *state, cmd command) error {
		user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
		if err != nil {
			return fmt.Errorf("error in trying to get current user: %w", err)
		}
		return handler(s, cmd, user)

	}

}
