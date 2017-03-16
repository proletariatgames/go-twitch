// Channels methods of the twitch api.
// https://github.com/justintv/Twitch-API/blob/master/v3_resources/channels.md

package twitch

import (
	"fmt"
	"math/rand"

	"github.com/google/go-querystring/query"
)

// used with GET /channels/:channel/videos
type VideosS struct {
	Videos []VideoS `json:"videos,omitempty"`
	Links  LinksS   `json:"_links,omitempty"`
}

// used with GET /channels/:channel/editors
type EditorsS struct {
	Users []UserS `json:"users,omitempty"`
	Links LinksS  `json:"_links,omitempty"`
}

// used with GET /channels/:channel/follows
type FollowsS struct {
	Follows []FollowS `json:"follows,omitempty"`
	Total   int       `json:"_total,omitempty"`
	Links   LinksS    `json:"_links,omitempty"`
}

type FollowS struct {
	User UserS `json:"user,omitempty"`
}

type PanelDataS struct {
	Link        string `json:"link,omitempty"`
	Image       string `json:"image,omitempty"`
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
}

type PanelS struct {
	Id              int        `json:"_id,omitempty"`
	DisplayOrder    int        `json:"display_order,omitempty"`
	Kind            string     `json:"kind,omitempty"`
	HtmlDescription string     `json:"html_description,omitempty"`
	UserId          int        `json:"user_id,omitempty"`
	Data            PanelDataS `json:"data,omitempty"`
	Channel         string     `json:"channel,omitempty"`
}

type SubsS struct {
	Total         int    `json:"_total,omitempty"`
	Links         LinksS `json:"_links,omitempty"`
	Subscriptions []SubS `json:"subscriptions,omitempty"`
}

type SubS struct {
	Id   string `json:"_id,omitempty"`
	User UserS  `json:"user,omitempty"`
}

type ChannelsMethod struct {
	client *Client
}

// type ChanSubS struct {
// 	ViewUntil          int      `json:"view_until,omitempty"`
// 	RestrictedBitrates []string `json:"restricted_bitrates"` // needs double checking
// }

// type PrivateS struct {
// 	AllowedToView bool `json:"allowed_to_view,omitempty"`
// }

// type TokenS struct {
// 	Adblock          bool     `json:"adblock"`
// 	PlayerType       *string  `json:"player_type"` // needs double checking
// 	Platform         *string  `json:"platform"`    // needs double checking
// 	UserID           *string  `json:"user_id"`     // needs double checking
// 	Channel          string   `json:"channel"`
// 	Expires          int      `json:"expires,omitempty"`
// 	Chansub          ChanSubS `json:"chansub,omitempty"`
// 	Private          PrivateS `json:"private,omitempty"`
// 	Privileged       bool     `json:"privileged,omitempty"`
// 	SourceRestricted bool     `json:"source_restricted,omitempty"`
// 	HTTPSRequired    bool     `json:"https_required,omitempty"`
// 	ShowAds          bool     `json:"show_ads,omitempty"`
// 	DeviceID         string   `json:"device_id"`
// }

type AccessTokenS struct {
	Sig              string `json:"sig,omitempty"`
	MobileRestricted bool   `json:"mobile_restricted"`
	Token            string `json:"token,omitempty"`
}

// Returns a channel object. If `name` is an empty string, returns the channel
// object of authenticated user.
func (c *ChannelsMethod) Channel(name string) (*ChannelS, error) {
	rel := "channel" // get authenticated channel
	if name != "" {
		rel = "channels/" + name
	}

	channel := new(ChannelS)
	_, err := c.client.Get(rel, channel)
	return channel, err
}

// Returns a list of users who are editors of channel `name`.
func (c *ChannelsMethod) editors(name string) (*EditorsS, error) {
	rel := "channels/" + name + "/editors"

	editors := new(EditorsS)
	_, err := c.client.Get(rel, editors)
	return editors, err
}

// Returns a list of videos ordered by time of creation, starting with the most
// recent from channel `name`.
func (c *ChannelsMethod) Videos(name string, opt *ListOptions) (*VideosS, error) {
	rel := "channels/" + name + "/videos"
	if opt != nil {
		v, err := query.Values(opt)
		if err != nil {
			return nil, err
		}
		rel += "?" + v.Encode()
	}

	videos := new(VideosS)
	_, err := c.client.Get(rel, videos)
	return videos, err
}

// Returns a list of users the channel `name` is following.
func (c *ChannelsMethod) Follows(name string, opt *ListOptions) (*FollowsS, error) {
	rel := "channels/" + name + "/follows"
	if opt != nil {
		v, err := query.Values(opt)
		if err != nil {
			return nil, err
		}
		rel += "?" + v.Encode()
	}

	follow := new(FollowsS)
	_, err := c.client.Get(rel, follow)
	return follow, err
}

// Returns the list of panels the channel `name` has.
func (c *ChannelsMethod) Panels(name string, opt *ListOptions) (*[]PanelS, error) {
	rel := "channels/" + name + "/panels"
	if opt != nil {
		v, err := query.Values(opt)
		if err != nil {
			return nil, err
		}
		rel += "?" + v.Encode()
	}

	panels := new([]PanelS)
	_, err := c.client.GetAPI(rel, panels)
	return panels, err
}

func (c *ChannelsMethod) AccessToken(name string) (*AccessTokenS, error) {
	rel := "channels/" + name + "/access_token"
	accessToken := new(AccessTokenS)
	_, err := c.client.GetAPI(rel, accessToken)
	return accessToken, err
}

func (c *ChannelsMethod) M3U8(name string, opt *M3U8Options) (string, error) {
	rel := fmt.Sprintf("channel/hls/%s.m3u8", name)

	if opt != nil {
		opt.AllowAudioOnly = true
		opt.AllowSource = true
		opt.Type = "any"
		opt.Random = rand.Int()
		opt.Player = "twitchweb"
	}

	v, err := query.Values(opt)
	if err != nil {
		return "", err
	}
	rel += "?" + v.Encode()

	return c.client.GetUsher(rel)
}

func (c *ChannelsMethod) subscriptions(name string, opt *ListOptions) (*SubsS, error) {
	rel := "channels/" + name + "/subscriptions"
	if opt != nil {
		v, err := query.Values(opt)
		if err != nil {
			return nil, err
		}
		rel += "?" + v.Encode()
	}

	subs := new(SubsS)
	_, err := c.client.Get(rel, subs)
	return subs, err
}

func (c *ChannelsMethod) subscription(name string, user string) (*SubS, error) {
	rel := fmt.Sprintf("channels/%s/subscriptions/%s", name, user)

	sub := new(SubS)
	_, err := c.client.Get(rel, sub)
	return sub, err
}

// TODO PUT /channels/:channel/

// TODO POST /channels/:channel/commercial
