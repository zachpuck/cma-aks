package azureutil

import (
	"context"
	"fmt"
	"time"

	"github.com/Azure/azure-sdk-for-go/services/graphrbac/1.6/graphrbac"
	"github.com/Azure/go-autorest/autorest/azure/auth"

	"github.com/Azure/go-autorest/autorest/date"
	"github.com/Azure/go-autorest/autorest/to"
)

type ClusterAccount struct {
	ClientID     string
	ClientSecret string
}

// GetServicePrincipalsClient get an authorized service principals client
func GetServicePrincipalsClient(tenantID string, clientID string, clientSecret string) (graphrbac.ServicePrincipalsClient, error) {
	// create new service principals client
	spClient := graphrbac.NewServicePrincipalsClient(tenantID)

	// authorize request for graph api
	activeDirectoryEndpoint := "https://login.microsoftonline.com"
	graphResourceID := "https://graph.windows.net"
	graphClientConfig := auth.ClientCredentialsConfig{
		ClientID: clientID,
		ClientSecret: clientSecret,
		TenantID: tenantID,
		Resource:  graphResourceID,
		AADEndpoint: activeDirectoryEndpoint,
	}
	authorizer, err := graphClientConfig.Authorizer()
	if err != nil {
		return spClient, fmt.Errorf("failed to initialize authorizer: %v", err)
	}

	spClient.Authorizer = authorizer
	return spClient, nil
}

// GetApplicationsClient gets an authorized active directory application client
func GetApplicationsClient(tenantID string, clientID string, clientSecret string) (graphrbac.ApplicationsClient, error) {
	// create new applications client
	appClient := graphrbac.NewApplicationsClient(tenantID)

	// authorize request for graph api
	activeDirectoryEndpoint := "https://login.microsoftonline.com"
	graphResourceID := "https://graph.windows.net"
	graphClientConfig := auth.ClientCredentialsConfig{
		ClientID: clientID,
		ClientSecret: clientSecret,
		TenantID: tenantID,
		Resource:  graphResourceID,
		AADEndpoint: activeDirectoryEndpoint,
	}
	authorizer, err := graphClientConfig.Authorizer()
	if err != nil {
		return appClient, fmt.Errorf("failed to initialize authorizer: %v", err)
	}

	appClient.Authorizer = authorizer
	return appClient, nil
}

// CreateADApplication creates an Active Directory application
func CreateADApplication(ctx context.Context, appClient graphrbac.ApplicationsClient, resourceName string) (graphrbac.Application, error) {
	appName := resourceName + "-sp"
	appURL := "https://" + appName

	appParameters := graphrbac.ApplicationCreateParameters{
		AvailableToOtherTenants: to.BoolPtr(false),
		DisplayName:             to.StringPtr(appName),
		Homepage:                to.StringPtr(appURL),
		IdentifierUris:          &[]string{appURL},
	}

	// FIXME: Check if it already exist first.
	result, err := appClient.Create(ctx, appParameters)
	if err != nil {
		return result, fmt.Errorf("cannot create application: %v", err)
	}

	return result, nil
}

// CreateServicePrincipal creates a service principals associated with an AD application
func CreateServicePrincipal(ctx context.Context, spClient graphrbac.ServicePrincipalsClient, appID string) (graphrbac.ServicePrincipal, error) {

	spParameters := graphrbac.ServicePrincipalCreateParameters{
		AppID:          to.StringPtr(appID),
		AccountEnabled: to.BoolPtr(true),
	}
	result, err := spClient.Create(ctx, spParameters)
	if err != nil {
		return result, fmt.Errorf("cannot create service principal: %v", err)
	}

	return result, nil
}

// AddClientSecret adds a secret to a specific AD app
func AddClientSecret(ctx context.Context, appClient graphrbac.ApplicationsClient, appID string) (*ClusterAccount, error) {
	var clusterAccount ClusterAccount
	
	// TODO: create Key and Value
	secret := ""
	keyID := ""
	
	passwordParameters := graphrbac.PasswordCredentialsUpdateParameters{
		Value: &[]graphrbac.PasswordCredential{
			{
				StartDate: &date.Time{time.Now()},
				// FIXME: fix EndDate
				EndDate:   &date.Time{time.Date(2018, time.December, 20, 22, 0, 0, 0, time.UTC)},
				Value:     to.StringPtr(secret),
				KeyID:     to.StringPtr(keyID),
			},
		},
	}
	result, err := appClient.UpdatePasswordCredentials(ctx, appID, passwordParameters)
	if err != nil {
		return nil, fmt.Errorf("cannot create password: %v", err)
	}

	clusterAccount.ClientID = appID
	clusterAccount.ClientSecret = secret

	fmt.Println("result ", result)

	return &clusterAccount, nil
}
