package main

import (
	"context"
	"fmt"
	"time"

	"github.com/anandki1209/gator/internal/database"
	"github.com/google/uuid"
)

func handlerFollow(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("No url provided")
	}

	url := cmd.Args[0]

	feed, err := s.db.GetFeedByUrl(context.Background(), url)
	if err != nil {
		return fmt.Errorf("error in getting feed: %w", err)
	}

	currentUserName := s.cfg.CurrentUserName
	currentUser, err := s.db.GetUser(context.Background(), currentUserName)
	if err != nil {
		return fmt.Errorf("error in trying to get current user: %w", err)
	}

	feedFollowParams := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    currentUser.ID,
		FeedID:    feed.ID,
	}

	followFeed, err := s.db.CreateFeedFollow(context.Background(), feedFollowParams)
	if err != nil {
		return fmt.Errorf("error in creating followFeed record: %w", err)
	}

	fmt.Println(followFeed.FeedName)
	fmt.Println(followFeed.UserName)

	return nil

}

func handlerFollowing(s *state, cmd command) error {
	currentUserName := s.cfg.CurrentUserName
	currentUser, err := s.db.GetUser(context.Background(), currentUserName)
	if err != nil {
		return fmt.Errorf("error in trying to get current user: %w", err)
	}

	feedfollows, err := s.db.GetFeedFollowsForUser(context.Background(), currentUser.ID)
	if err != nil {
		return fmt.Errorf("error in getting feedfollows for user: %w", err)
	}

	for _, feedfollow := range feedfollows {
		fmt.Println(feedfollow.FeedName)
	}
	return nil
}
