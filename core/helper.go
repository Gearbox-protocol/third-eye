package core

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

func Topic(topic string) common.Hash {
	return crypto.Keccak256Hash([]byte(topic))
}
