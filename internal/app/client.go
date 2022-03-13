package client

import (
	"qin/internal/app/chain"
)

func Client() {
	bc := chain.NewBlockchain()

	bc.AddBlock("Send 1 BTC to han")
	bc.AddBlock("Send 2 more BTC to tuo")

}
