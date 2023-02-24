package lib

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	msgraphsdkgo "github.com/microsoftgraph/msgraph-sdk-go"
	"github.com/microsoftgraph/msgraph-sdk-go/me"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"log"
	"os"
	"strings"
)

func CheckError(msg string, err error) {
	if err != nil {
		log.Fatal(msg, err)
	}
}

type GraphHelper struct {
	interactiveBrowserCredential *azidentity.InteractiveBrowserCredential
	graphUserScopes              []string
	appClient                    *msgraphsdkgo.GraphServiceClient
}

func NewGraphHelper() *GraphHelper {
	g := &GraphHelper{}
	return g
}

func (g *GraphHelper) InitializeGraphForAppAuth() (models.Userable, error) {
	clientId := os.Getenv("CLIENT_ID")
	tenantId := os.Getenv("TENANT_ID")
	scope := os.Getenv("GRAPH_USER_SCOPES")
	g.graphUserScopes = strings.Split(scope, ",")
	cred, err := azidentity.NewInteractiveBrowserCredential(&azidentity.InteractiveBrowserCredentialOptions{
		TenantID:    tenantId,
		ClientID:    clientId,
		RedirectURL: "http://localhost:8080/",
	})
	CheckError("New Browser Cred Error: ", err)
	g.interactiveBrowserCredential = cred

	query := me.MeRequestBuilderGetQueryParameters{
		// Only request specific properties
		Select: []string{"displayName", "mail", "userPrincipalName"},
	}

	client, err := msgraphsdkgo.NewGraphServiceClientWithCredentials(cred, g.graphUserScopes)
	g.appClient = client

	CheckError("Client Cred Fail ", err)
	user, err := client.Me().Get(context.Background(), &me.MeRequestBuilderGetRequestConfiguration{
		QueryParameters: &query,
	})
	return user, err
}
