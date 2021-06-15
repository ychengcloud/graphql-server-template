package middleware

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	kratosclient "github.com/ory/kratos-client-go"
	"{{ .Extra.pkgpath }}/internal/models"
)

var OryKratosSession = "ory_kratos_session"

type AuthKratos struct {
	opts *models.ServiceOptions

	cfg *kratosclient.Configuration
}

func NewAuthKratos(opts *models.ServiceOptions) *AuthKratos {
	configuration := kratosclient.NewConfiguration()
	configuration.Host = opts.KratosHost
	configuration.Scheme = "http"
	configuration.Servers = kratosclient.ServerConfigurations{
		{
			// URL: opts.KratosHost,
		},
	}

	ak := &AuthKratos{
		opts: opts,
		cfg:  configuration,
	}
	return ak

}

func (ak *AuthKratos) Authenticate(c *gin.Context) (*models.CurrentUser, error) {
	apiClient := kratosclient.NewAPIClient(ak.cfg)
	sessionToken, err := c.Cookie(OryKratosSession)

	apiClient.GetConfig().AddDefaultHeader("Cookie", OryKratosSession+"="+sessionToken)
	session, _, err := apiClient.PublicApi.ToSession(c.Request.Context()).Execute()
	if err != nil {
		return nil, err
	}

	traits := session.Identity.GetTraits()
	if traits == nil {
		return nil, fmt.Errorf("traits not found")
	}

	email, ok := traits.(map[string]interface{})["email"]
	if !ok {
		return nil, fmt.Errorf("traits interface not is map[string]interface\n")
	}
	emailValue, ok := email.(string)
	if !ok {
		return nil, fmt.Errorf("email interface not is string\n")
	}
	user := &models.CurrentUser{
		Name: emailValue,
	}
	return user, nil
}
