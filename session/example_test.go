package session_test

import (
	apiGoSession "github.com/epicchainlabs/epicchain-api-go/v2/session"
	cid "github.com/epicchainlabs/epicchain-sdk-go/container/id"
	neofscrypto "github.com/epicchainlabs/epicchain-sdk-go/crypto"
	"github.com/epicchainlabs/epicchain-sdk-go/session"
	"github.com/epicchainlabs/epicchain-sdk-go/user"
)

// Both parties agree on a secret (private session key), the possession of which
// will be authenticated by a trusted person. The principal confirms his trust by
// signing the public part of the secret (public session key).
func ExampleContainer() {
	// import neofscrypto "github.com/epicchainlabs/epicchain-sdk-go/crypto"
	// import "github.com/epicchainlabs/epicchain-sdk-go/user"
	// import cid "github.com/epicchainlabs/epicchain-sdk-go/container/id"

	// you private key/signer, to prove you are you
	var principalSigner user.Signer
	// trusted party, who can do action on behalf of you. Usually the key maybe taken from Client.SessionCreate.
	var trustedPubKey neofscrypto.PublicKey
	var cnr cid.ID

	var tok session.Object
	tok.ForVerb(session.VerbObjectPut)
	tok.SetAuthKey(trustedPubKey)
	tok.BindContainer(cnr)
	// ...

	_ = tok.Sign(principalSigner)

	// transfer the token to a trusted party
}

// Instances can be also used to process NeoFS API V2 protocol messages with [https://github.com/nspcc-dev/neofs-api] package.
func ExampleObject_marshalling() {
	// import apiGoSession "github.com/epicchainlabs/epicchain-api-go/v2/session"

	// On the client side.

	var tok session.Object
	var msg apiGoSession.Token

	tok.WriteToV2(&msg)
	// *send message*

	// On the server side.

	_ = tok.ReadFromV2(msg)
}
