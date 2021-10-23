package twitter_mock

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
	Tweets []*Tweet
}

func (t *Twitter) AddTweet(tw *Tweet) *Twitter {
	t.Tweets = append([]*Tweet{tw}, t.Tweets...)
	return t
}

// will return Tweet and ok when tweets are available
// ok will be false on an empty list
func (t *Twitter) PeekTweet() (*Tweet, bool) {
	if len(t.Tweets) == 0 {
		return nil, false
	}

	return t.Tweets[0], true
}

func MakeTwitter() Twitter {
	tweets := make([]*Tweet, 0)
	t := Twitter{
		Tweets: tweets,
	}

	return t
}

func MakeTweet(msg string) Tweet {
	tweetID++
	return Tweet{
		Id:  tweetID,
		Msg: msg,
	}
}
