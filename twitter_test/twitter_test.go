package main

import (
	"go-program-sd-concepts/twitter_mock"
	"testing"
)

func TestTweet(t *testing.T) {
	msg := "Hello there! #welcome"
	tweet := twitter_mock.MakeTweet(msg)
	if tweet.Msg != msg {
		t.Fatalf("%s is not %s", tweet.Msg, msg)
	}

	farewell := "Goodbye! #farewell"
	tweet.UpdateMsg(farewell).UpdateMsg(msg).UpdateMsg(farewell)

	if tweet.Msg != farewell {
		t.Fatalf("%s is not %s", tweet.Msg, farewell)
	}

	tweet2 := twitter_mock.MakeTweet(msg)

	if tweet.Id >= tweet2.Id {
		t.Fatalf("%d should be less than %d", tweet.Id, tweet2.Id)
	}

	twitter := twitter_mock.MakeTwitter()
	twitter.AddTweet(&tweet).AddTweet(&tweet2)

	if len(twitter.Tweets) != 2 {
		t.Fatalf("Got %d tweets, expected 2", len(twitter.Tweets))
	}

	lastTweet, ok := twitter.PeekTweet()

	if lastTweet.Msg != tweet2.Msg && ok != true {
		t.Fatalf("Failed to get the last message. got %s %d", lastTweet.Msg, lastTweet.Id)
	}
}
