package endpass

import "errors"

var (
	ErrNoAccessToken = errors.New("Before doing requests you must pass authorization and get access token. See \"AuthCodeURL\" and \"Exchange\" methods")
)
