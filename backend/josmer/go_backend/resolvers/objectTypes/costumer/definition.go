package costumer

import (
	"negocio_backend/models"

	"github.com/pjmd89/gogql/lib/gql/definitionError"
	"github.com/pjmd89/gogql/lib/resolvers"
	"github.com/pjmd89/goutils/dbutils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// models alias
type costumerModel = models.Costumer

// map alias
type myMap = map[string]interface{}

// mongo ID alias
type objectID = primitive.ObjectID

// function alias
var isValidObjectID = primitive.IsValidObjectID

type Costumer struct {
	model models.Costumer
}

func NewCostumer(db dbutils.DBInterface) (o resolvers.ObjectTypeInterface) {
	o = &Costumer{
		model: models.Costumer{},
	}
	o.(*Costumer).model.Init(models.Costumer{}, db)
	return o
}
func (o *Costumer) Resolver(info resolvers.ResolverInfo) (r resolvers.DataReturn, err definitionError.GQLError) {
	switch info.Operation {
	case "query":
		switch info.Resolver {
		case "edges":
			r, err = o.edges(info)
		}
	case "mutation":
		switch info.Resolver {
		case "createCostumer":
			r, err = o.createCostumer(info)
		case "updateCostumer":
			r, err = o.updateCostumer(info)
		}
	}
	return
}
func (o *Costumer) Subscribe(info resolvers.ResolverInfo) (r bool) {
	return
}
