package main

import "fmt"

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}

	userInput := cmd.Args[0]
	if err := s.cfg.SetUser(userInput); err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}

	fmt.Printf("user name has been set to %s\n", userInput)

	return nil
}
