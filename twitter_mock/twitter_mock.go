package twitter_mock

import (
	"regexp"
	"strings"
)

var tweetID int

type Tweet struct {
	Id  int
	Msg string
}

func (t *Tweet) UpdateMsg(msg string) *Tweet {
	t.Msg = msg
	return t
}

type Twitter struct {
	Tweets     []*Tweet
	searchTags map[string][]int
}

func getTags(msg string) []string {
	tags := make([]string, 0)
	r, _ := regexp.Compile("^#")
	words := strings.Split(msg, " ")

	for _, word := range words {
		word = strings.TrimSpace(word)
		if r.MatchString(word) {
			tags = append(tags, word[1:])
		}
	}

	return tags
}

func (t *Twitter) UpdateTagIndex(tw Tweet) {
	tags := getTags(tw.Msg)
	if len(tags) > 0 {
		for _, tag := range tags {
			if tweetIds, ok := t.searchTags[tag]; ok {
				tweetIds = append(tweetIds, tw.Id)
			} else {
				t.searchTags[tag] = []int{tw.Id}
			}
		}
	}
}

func (t Twitter) SearchTweetsByTag(tag string) []int {
	if tweetIds, ok := t.searchTags[tag]; ok {
		return tweetIds
	}

	return []int{}
}

func (t *Twitter) AddTweet(tw *Tweet) *Twitter {
	t.UpdateTagIndex(*tw)
	t.Tweets = append([]*Tweet{tw}, t.Tweets...)
	return t
}

func (t *Twitter) PeekTweet() (*Tweet, bool) {
	if len(t.Tweets) == 0 {
		return nil, false
	}

	return t.Tweets[0], true
}

func MakeTwitter() Twitter {
	tweets := make([]*Tweet, 0)
	s := make(map[string][]int, 0)
	return Twitter{
		Tweets:     tweets,
		searchTags: s,
	}
}

func MakeTweet(msg string) Tweet {
	tweetID++
	return Tweet{
		Id:  tweetID,
		Msg: msg,
	}
}
