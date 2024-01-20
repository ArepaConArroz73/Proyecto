package edgecostumer

import (
	"negocio_backend/lib"

	"github.com/pjmd89/gogql/lib/gql/definitionError"
	"github.com/pjmd89/gogql/lib/resolvers"
)

func (o *EdgeCostumer) getCostumers(info resolvers.ResolverInfo) (r resolvers.DataReturn, err definitionError.GQLError) {
	r = lib.NewEdge(info, "costumer")
	return
}
