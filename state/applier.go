package state

import (
	"github.com/filecoin-project/chain-validation/chain/types"
)

// Applier applies abstract messages to states.
type Applier interface {
	ApplyMessage(context *types.ExecutionContext, state Wrapper, msg *types.Message) (types.MessageReceipt, error)
}
