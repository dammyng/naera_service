package v1

import (
	"naerarauth/internals/persistence"
	"naerarshared/interfaces"
)

type NaerarAuthRouteHandler struct {
	Db persistence.NaerarAuthDBHandler
	MemStore interfaces.MemStorage
}

func NewNaerarAuthRouteHandler(db persistence.NaerarAuthDBHandler, redis interfaces.MemStorage) *NaerarAuthRouteHandler {
	return &NaerarAuthRouteHandler{
		Db: db,
		MemStore: redis,
	}
}
