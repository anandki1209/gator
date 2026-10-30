package main

import (
	"context"
	"fmt"
	"time"

	"github.com/anandki1209/gator/internal/database"
	"github.com/google/uuid"
)

func handlerFollow(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("No url provided")
	}

	url := cmd.Args[0]

	feed, err := s.db.GetFeedByUrl(context.Background(), url)
	if err != nil {
		return fmt.Errorf("error in getting feed: %w", err)
	}

	feedFollowParams := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
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

func handlerFollowing(s *state, cmd command, user database.User) error {
	feedfollows, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("error in getting feedfollows for user: %w", err)
	}

	for _, feedfollow := range feedfollows {
		fmt.Println(feedfollow.FeedName)
	}
	return nil
}
