package edgecostumer

import (
	"negocio_backend/models"

	"github.com/pjmd89/gogql/lib/gql/definitionError"
	"github.com/pjmd89/gogql/lib/resolvers"
	"github.com/pjmd89/goutils/dbutils"
)

type EdgeCostumer struct {
	model models.Costumer
}

func NewEdgeCostumer(db dbutils.DBInterface) (o resolvers.ObjectTypeInterface) {
	o = &EdgeCostumer{
		model: models.Costumer{},
	}
	o.(*EdgeCostumer).model.Init(models.Costumer{}, db)
	return o
}
func (o *EdgeCostumer) Resolver(info resolvers.ResolverInfo) (r resolvers.DataReturn, err definitionError.GQLError) {
	switch info.Operation {
	case "query":
		switch info.Resolver {
		case "getCostumers":
			r, err = o.getCostumers(info)

		}

	case "mutation":
		break
	}
	return
}
func (o *EdgeCostumer) Subscribe(info resolvers.ResolverInfo) (r bool) {
	return
}
