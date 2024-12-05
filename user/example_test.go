package user_test

import (
	apiGoRefs "github.com/epicchainlabs/epicchain-api-go/v2/refs"
	"github.com/epicchainlabs/epicchain-sdk-go/user"
	"github.com/nspcc-dev/neo-go/pkg/util"
)

// NeoFS user identification is compatible with Neo accounts.
func ExampleID_SetScriptHash() {
	// import "github.com/nspcc-dev/neo-go/pkg/util"
	var id user.ID

	var scriptHash util.Uint160 // user account in NeoFS
	id.SetScriptHash(scriptHash)
}

// ID is compatible with the NeoFS Smart Contract API.
func ExampleID_WalletBytes() {
	var id user.ID
	// ...

	wallet := id.WalletBytes()
	_ = wallet

	// use wallet in call
}

// Encoding mechanisms are used to transfer identifiers on client.
func ExampleID_EncodeToString() {
	var id user.ID
	// ...

	_ = id.EncodeToString()
}

// Encoding mechanisms are used to transfer identifiers on server.
func ExampleID_DecodeString() {
	var id user.ID
	// ...

	var s string
	_ = id.DecodeString(s)
}

// Instances can be also used to process NeoFS API V2 protocol messages with [https://github.com/nspcc-dev/neofs-api] package.
func ExampleID_marshalling() {
	// import apiGoRefs "github.com/epicchainlabs/epicchain-api-go/v2/refs"

	// On the client side.

	var id user.ID
	var msg apiGoRefs.OwnerID
	id.WriteToV2(&msg)
	// *send message*

	// On the server side.

	_ = id.ReadFromV2(msg)
}
