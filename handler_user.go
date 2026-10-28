package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/anandki1209/gator/internal/database"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}

	userInput := cmd.Args[0]

	if _, err := s.db.GetUser(context.Background(), userInput); err != nil {
		return fmt.Errorf("error getting user details: %w", err)
	}
	if err := s.cfg.SetUser(userInput); err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}

	fmt.Printf("user name has been set to %s\n", userInput)

	return nil
}

func handlerRegister(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}

	userInput := cmd.Args[0]
	user, err := s.db.CreateUser(context.Background(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      userInput,
	})
	if err != nil {
		return fmt.Errorf("error in registering user: %w", err)

	}

	if err := s.cfg.SetUser(userInput); err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}

	fmt.Println("user was created succesfully")
	printUser(user)
	return nil

}

func printUser(user database.User) {
	fmt.Printf(" * ID: %v/n", user.ID)
	fmt.Printf(" * Name: %v/n", user.Name)
}
