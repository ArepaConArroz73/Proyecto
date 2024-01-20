package models

import (
	"github.com/pjmd89/mongomodel/mongomodel"
)

type EdgeCostumer struct {
	mongomodel.Model `bson:"-" gql:"omit=true"`
	Edges            []Costumer `bson:"edges" gql:"name=edges"`
	PageInfo         PageInfo   `bson:"pageInfo" gql:"name=pageInfo"`
}
