package twitch

import (
	"net/http"
	"testing"
)

func TestChatEmoticons(t *testing.T) {

	tc := NewClient(&http.Client{})

	_, err := tc.Chat.Emoticons()

	if err != nil {
		t.Errorf("error not nil: %+v", err)
	}

}
