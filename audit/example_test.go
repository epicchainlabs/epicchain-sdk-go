package audit_test

import (
	"github.com/epicchainlabs/epicchain-sdk-go/audit"
	cid "github.com/epicchainlabs/epicchain-sdk-go/container/id"
)

func ExampleResult() {
	var res audit.Result
	var cnr cid.ID

	res.ForEpoch(32)
	res.ForContainer(cnr)
	// ...
	res.Complete()

	// Result instances can be stored in a binary format on client side.
	data := res.Marshal()

	// ...

	// Result instances can be restored from a binary format on server side.
	var auditResult audit.Result
	_ = auditResult.Unmarshal(data)
}
