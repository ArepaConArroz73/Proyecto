package models

import (
	"negocio_backend/resolvers/scalars"

	"github.com/pjmd89/mongomodel/mongomodel"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Costumer struct {
	mongomodel.Model `bson:"-" gql:"omit=true"`
	Id               primitive.ObjectID `bson:"_id,omitempty" gql:"name=_id,id=true,objectID=true"`
	FirstName        string             `bson:"firstName" gql:"name=firstName"`
	LastName         string             `bson:"lastName" gql:"name=lastName"`
	DNI              scalars.DNI        `bson:"dni" gql:"name=dni"`
	Address          string             `bson:"address" gql:"name=address"`
	TotalVisits      int64              `bson:"totalVisits" gql:"name=totalVisits,objectID=true"`
	Created          scalars.DateTime   `bson:"created" gql:"name=created,created=true"`
	Updated          scalars.DateTime   `bson:"updated" gql:"name=updated,updated=true"`
}
