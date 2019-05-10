package main

import (
	"errors"
)

// ErrNoAvatar is the error that is returned when the
// Avatar instance is unable to provide an avatar URL.
var ErrNoAvatar = errors.New("chat: unable to get an Avatar URL.")

// Avatar represents types capable of representing user profile pictures.
type Avatar interface {
	// GetAvatarURL gets the avatar URL for the specified client,
	// or returns an error if something goes wrong.
	// ErrNoAvatarURL is returned if the object is unable to get
	// a URL for the specified client.
	GetAvatarURL(c *client) (string, error)
}

// AuthAvatar ...
type AuthAvatar struct{}

// UseAuthAvatar ...
var UseAuthAvatar AuthAvatar

// GetAvatarURL ...
func (AuthAvatar) GetAvatarURL(c *client) (string, error) {
	url, ok := c.userData["avatar_url"]
	if !ok {
		return "", ErrNoAvatar
	}
	urlStr, ok := url.(string)
	if !ok {
		return "", ErrNoAvatar
	}
	return urlStr, nil
}
