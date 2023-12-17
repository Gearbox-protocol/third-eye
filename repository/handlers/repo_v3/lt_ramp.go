package repo_v3

import (
	"github.com/Gearbox-protocol/third-eye/repository/handlers"
)

type LTRampHandler struct {
	blocks *handlers.BlocksRepo
}

func NewLTRampHandler(blocks *handlers.BlocksRepo) LTRampHandler {
	return LTRampHandler{
		blocks: blocks,
	}
}
