package twitch

import (
	"net/http"
	"testing"
)

func TestUsersUser(t *testing.T) {

	tc := NewClient(&http.Client{})

	_, err := tc.Users.User("22747064") // test_user1

	if err != nil {
		t.Errorf("error not nil: %+v", err)
	}

}

func TestUsersFollows(t *testing.T) {

	tc := NewClient(&http.Client{})

	opt := &ListOptions{
		Limit:  1,
		Offset: 0,
	}

	_, err := tc.Users.Follows("22747064", opt)

	if err != nil {
		t.Errorf("error not nil: %+v", err)
	}

}

func TestUsersFollow(t *testing.T) {

	tc := NewClient(&http.Client{})

	_, err := tc.Users.Follow("19407332", "7236692") // Roybot1911, DansGaming

	if err != nil {
		t.Errorf("error not nil: %+v", err)
	}

}
