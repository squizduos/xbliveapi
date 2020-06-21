package xbliveapi

import (
	"fmt"
	"time"
)

type credentials struct {
	token        string
	uhs          string
	xid          string
	gamertag     string
	accessToken  string
	refreshToken string
	userID       string
	expiresAt    time.Time
}

func (c *credentials) authHeader() string {
	return fmt.Sprintf("XBL3.0 x=%s;%s", c.uhs, c.token)
}
