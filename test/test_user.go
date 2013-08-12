package test

import (
	"github.com/stretchr/gomniauth/common"
	"github.com/stretchr/stew/objects"
	"github.com/stretchr/testify/mock"
)

type TestUser struct {
	mock.Mock
}

// Email gets the users email address.
func (u *TestUser) Email() string {
	return u.Called().String(0)
}

// Name gets the users full name.
func (u *TestUser) Name() string {
	return u.Called().String(0)
}

// Nickname gets the users nickname or username.
func (u *TestUser) Nickname() string {
	return u.Called().String(0)
}

// AvatarURL gets the URL of an image representing the user.
func (u *TestUser) AvatarURL() string {
	return u.Called().String(0)
}

// ProviderCredentials gets a map of Credentials (by provider name).
func (u *TestUser) ProviderCredentials() map[string]*common.Credentials {
	return u.Called().Get(0).(map[string]*common.Credentials)
}

// IDForProvider gets the ID value for the specified provider name for
// this user from the ProviderCredentials data.
func (u *TestUser) IDForProvider(provider string) string {
	return u.Called().String(0)
}

// ID gets this user's globally unique ID.
func (u *TestUser) ID() string {
	return u.Called().String(0)
}

// Data gets the underlying data that makes up this User.
func (u *TestUser) Data() objects.Map {
	return u.Called().Get(0).(objects.Map)
}
