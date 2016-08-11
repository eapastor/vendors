package oauth2

import (
	"errors"
	"github.com/lastbackend/oauth2/internal"
)

func GetClient(vendor, clientID, clientSecretID, redirectURI string) (internal.IOAuth2, error) {

	if _, ok := internal.OAuthVendors[vendor]; !ok {
		return nil, errors.New("vendor not supported")
	}

	return internal.OAuthVendors[vendor](clientID, clientSecretID, redirectURI)
}
