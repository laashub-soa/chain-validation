package chain

import (
	"github.com/filecoin-project/go-address"
	builtin_spec "github.com/filecoin-project/specs-actors/actors/builtin"
	init_spec "github.com/filecoin-project/specs-actors/actors/builtin/init"
	paych_spec "github.com/filecoin-project/specs-actors/actors/builtin/paych"

	"github.com/filecoin-project/chain-validation/chain/types"
)

var noParams []byte

// Transfer builds a simple value transfer message and returns it.
func (mp *MessageProducer) Transfer(to, from address.Address, opts ...MsgOpt) *types.Message {
	return mp.Build(to, from, builtin_spec.MethodSend, noParams, opts...)
}

func (mp *MessageProducer) CreatePaymentChannel(to, from address.Address, opts ...MsgOpt) *types.Message {
	return mp.InitExec(builtin_spec.InitActorAddr, from, init_spec.ExecParams{
		CodeCID: builtin_spec.PaymentChannelActorCodeID,
		ConstructorParams: MustSerialize(&paych_spec.ConstructorParams{
			From: from,
			To:   to,
		}),
	}, opts...)
}
