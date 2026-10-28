package main

import (
	"context"
	"fmt"
)

func handlerReset(s *state, cmd command) error {
	if err := s.db.DeleteUsers(context.Background()); err != nil {
		return fmt.Errorf("could not delete all records from users table: %w", err)
	}
	fmt.Println("deleted all records from users table")
	return nil
}
