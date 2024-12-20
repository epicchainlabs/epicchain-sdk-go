package bearer_test

import (
	"context"

	"github.com/epicchainlabs/epicchain-sdk-go/bearer"
	"github.com/epicchainlabs/epicchain-sdk-go/client"
	cid "github.com/epicchainlabs/epicchain-sdk-go/container/id"
	"github.com/epicchainlabs/epicchain-sdk-go/eacl"
	oid "github.com/epicchainlabs/epicchain-sdk-go/object/id"
	"github.com/epicchainlabs/epicchain-sdk-go/user"
)

// Define bearer token by setting correct lifetime, extended ACL and owner ID of
// the user that will attach token to its requests.
func Example() {
	// import "github.com/epicchainlabs/epicchain-sdk-go/eacl"
	// import "github.com/epicchainlabs/epicchain-sdk-go/user"

	var bearerToken bearer.Token
	var ownerID user.ID
	var eaclTable eacl.Table

	bearerToken.SetExp(500)
	bearerToken.SetIat(10)
	bearerToken.SetNbf(10)
	bearerToken.SetEACLTable(eaclTable)
	bearerToken.ForUser(ownerID)

	// Bearer token must be signed by owner of the container.
	// import neofscrypto "github.com/epicchainlabs/epicchain-sdk-go/crypto"

	var signer user.Signer
	// signer initialization, bearerToken initialization, other steps ...

	_ = bearerToken.Sign(signer)
}

// Provide signed token in JSON or binary format to the request sender. Request
// sender can attach this bearer token to the object service requests.
func ExampleToken_attachToRequest() {
	// import "github.com/epicchainlabs/epicchain-sdk-go/client"
	// import "github.com/epicchainlabs/epicchain-sdk-go/user"
	// import oid "github.com/epicchainlabs/epicchain-sdk-go/object/id"
	// import cid "github.com/epicchainlabs/epicchain-sdk-go/container/id"

	var bearerToken bearer.Token
	var sdkClient *client.Client
	var signer user.Signer

	// init bearerToken, sdkClient, signer, other steps ...

	var headParams client.PrmObjectHead
	headParams.WithBearerToken(bearerToken)
	// ...

	response, err := sdkClient.ObjectHead(context.Background(), cid.ID{}, oid.ID{}, signer, headParams)

	_ = response
	_ = err
}
