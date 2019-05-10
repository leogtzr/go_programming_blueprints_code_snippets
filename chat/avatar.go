package main

import (
	"crypto/md5"
	"errors"
	"io"
	"strings"
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

// GravatarAvatar
type GravatarAvatar struct{}

var UseGravatar GravatarAvatar

func (GravatarAvatar) GetAvatarURL(c *client) (string, error) {
	userID, ok := c.userData["userid"]
	if !ok {
		return "", ErrNoAvatar
	}
	userIDStr, ok := userID.(string)
	if !ok {
		return "", ErrNoAvatar
	}
	m := md5.New()
	io.WriteString(m, strings.ToLower(userIDStr))
	return "//www.gravatar.com/avatar/" + userIDStr, nil
}

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
