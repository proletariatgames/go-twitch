package twitch

import (
	"net/http"
	"testing"
)

func TestChannelsVideos(t *testing.T) {
	tc := NewClient(&http.Client{})

	opt := &ListOptions{
		Limit:  1,
		Offset: 1,
	}

	_, err := tc.Channels.Videos("7236692", opt) // DansGaming

	if err != nil {
		t.Errorf("error not nil: %+v", err)
	}
}

func TestChannelsFollows(t *testing.T) {
	tc := NewClient(&http.Client{})

	opt := &ListOptions{
		Limit:  1,
		Offset: 1,
	}

	_, err := tc.Channels.Follows("7236692", opt)

	if err != nil {
		t.Errorf("error not nil: %+v", err)
	}
}

func TestChannelsChannel(t *testing.T) {
	tc := NewClient(&http.Client{})

	_, err := tc.Channels.Channel("7236692")

	if err != nil {
		t.Errorf("error not nil: %+v", err)
	}
}
