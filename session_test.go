package gomniauth

import (
	"github.com/stretchr/gomniauth/common"
	"github.com/stretchr/stew/objects"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestNewSession(t *testing.T) {

	manager := new(Manager)
	id := "abc123"
	s := NewSession(manager, id)

	if assert.NotNil(t, s) {
		assert.Equal(t, manager, s.Manager())
		assert.Equal(t, id, s.ID())
	}

}

func TestSessionGetAuthURL(t *testing.T) {

	manager := new(Manager)
	id := "abc123"
	targetUrl := "http://www.google.com/"

	s := NewSession(manager, id)

	state := objects.NewMap("id", id, "targetUrl", targetUrl)

	provider := new(TestProvider)
	provider.On("Config").Return(objects.NewMap("clientId", "CLIENTID",
		"redirectURL", "http://www.test.com/",
		"accessType", "online",
		"approvalPrompt", "force"))

	provider.On("AuthType").Return(common.AuthTypeOAuth2)

	url, err := s.GetAuthURL(provider, state)

	if assert.NoError(t, err) {
		assert.Equal(t, "?access_type=online&approval_prompt=force&client_id=CLIENTID&redirect_uri=http%3A%2F%2Fwww.test.com%2F&response_type=code&scope=&state=state", url)
	}

	mock.AssertExpectationsForObjects(t, provider.Mock)

}
