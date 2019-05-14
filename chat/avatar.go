package main

import (
	"crypto/md5"
	"errors"
	"io"
	"io/ioutil"
	"path"
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

// UseGravatar ...
var UseGravatar GravatarAvatar

// GetAvatarURL ...
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

// FileSystemAvatar ...
type FileSystemAvatar struct{}

// UseFileSystemAvatar ...
var UseFileSystemAvatar FileSystemAvatar

// GetAvatarURL ...
func (FileSystemAvatar) GetAvatarURL(c *client) (string, error) {
	// if userid, ok := c.userData["userid"]; ok {
	// 	if useridStr, ok := userid.(string); ok {
	// 		return "/avatars/" + useridStr + ".jpg", nil
	// 	}
	// }
	// return "", ErrNoAvatar
	if userid, ok := c.userData["userid"]; ok {
		if useridStr, ok := userid.(string); ok {
			files, err := ioutil.ReadDir("avatars")
			if err != nil {
				return "", ErrNoAvatar
			}
			for _, file := range files {
				if file.IsDir() {
					continue
				}
				if match, _ := path.Match(useridStr+"*", file.Name()); match {
					return "/avatars/" + file.Name(), nil
				}
			}
		}
	}
	return "", ErrNoAvatar
}
