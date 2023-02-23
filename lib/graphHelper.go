package lib

import (
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	msgraphsdkgo "github.com/microsoftgraph/msgraph-sdk-go"
	"log"
	"os"
)

func CheckError(msg string, err error) {
	if err != nil {
		log.Fatal(msg, err)
	}
}

type GraphHelper struct {
	interactiveBrowserCredential *azidentity.InteractiveBrowserCredential
	appClient                    *msgraphsdkgo.GraphServiceClient
}

func NewGraphHelper() *GraphHelper {
	g := &GraphHelper{}
	return g
}

func (g *GraphHelper) InitializeGraphForAppAuth() error {
	clientId := os.Getenv("CLIENT_ID")
	tenantId := os.Getenv("TENANT_ID")
	cred, err := azidentity.NewInteractiveBrowserCredential(&azidentity.InteractiveBrowserCredentialOptions{
		TenantID:    tenantId,
		ClientID:    clientId,
		RedirectURL: "https://www.google.com/",
	})
	g.interactiveBrowserCredential = cred
	//CheckError("Error Creating Credentials: ", err)
	//auth, err := authentication.NewAzureIdentityAuthenticationProviderWithScopes(cred, []string{"Files.Read"})
	//CheckError("Error authentication provider", err)
	//adapter, err := msgraphsdkgo.NewGraphRequestAdapter(auth)
	//CheckError("Error Creating Adapter: ", err)

	// the following block create a new Graph Service Client using the previously created request adapter.

	client, err := msgraphsdkgo.NewGraphServiceClientWithCredentials(cred, []string{"User.Read"})
	CheckError("Client Cred Fail ", err)
	result, err := client.Me().Get(nil, nil)
	fmt.Println(result.GetAboutMe())
	g.appClient = client
	return err
}
