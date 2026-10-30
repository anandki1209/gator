package main

import (
	"context"
	"fmt"
	"time"

	"github.com/anandki1209/gator/internal/database"
	"github.com/google/uuid"
)

func handlerAddFeed(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 2 {
		return fmt.Errorf("did not provided name and url arguments")
	}

	name := cmd.Args[0]
	url := cmd.Args[1]

	feedParams := database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      name,
		Url:       url,
		UserID:    user.ID,
	}

	feed, err := s.db.CreateFeed(context.Background(), feedParams)
	if err != nil {
		return fmt.Errorf("error in creating feed: %w", err)
	}

	feedFollowParams := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	}

	followfeed, err := s.db.CreateFeedFollow(context.Background(), feedFollowParams)
	if err != nil {
		return fmt.Errorf("error in creating follow feed: %w", err)
	}

	fmt.Printf("%+v\n", feed)
	fmt.Printf("%+v\n", followfeed)
	return nil
}

func handlerGetFeeds(s *state, cmd command) error {
	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("error in retrieving all feeds from database: %w", err)
	}

	for _, feed := range feeds {
		fmt.Println(feed.FeedName)
		fmt.Println(feed.Url)
		fmt.Println(feed.UserName)
	}

	return nil

}
