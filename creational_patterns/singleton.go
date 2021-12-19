package creational

import (
	"sync/atomic"
)

var IDGenerator *idGenerator

func GetIdGenerator() *idGenerator {
	return IDGenerator
}

type idGenerator struct {
	id int32
}

func (g *idGenerator) GenerateID() int32 {
	return atomic.AddInt32(&g.id, 1)
}

func init() {
	IDGenerator = &idGenerator{0}
}
