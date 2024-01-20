package costumer

import (
	"fmt"

	"github.com/pjmd89/gogql/lib/gql/definitionError"
	"github.com/pjmd89/gogql/lib/resolvers"
)

func (o *Costumer) consultRecordByID(recordId objectID) (foundRecord costumerModel, consultErr definitionError.GQLError) {
	where := myMap{
		"_id": recordId,
	}

	found, readErr := o.model.Read(where, nil)
	if readErr != nil || len(found.([]costumerModel)) == 0 {
		consultErr = definitionError.NewFatal(fmt.Sprintf("costumer not found, query error: %s", readErr), nil)
		return
	}
	return found.([]costumerModel)[0], nil

}
func (o *Costumer) validateRecordID(recordId objectID, recordMap any) (idErr definitionError.GQLError) {
	if recordMap == nil || !isValidObjectID(recordId.Hex()) {
		idErr = definitionError.NewFatal("CostumerID not provided or invalid", nil)
		return
	}
	return
}

func (o *Costumer) createCostumer(info resolvers.ResolverInfo) (r resolvers.DataReturn, err definitionError.GQLError) {
	request := info.Args["input"].(myMap)
	request["totalVisits"] = int64(1)
	created, mistake := o.model.Create(request, nil)
	if mistake != nil {
		err = definitionError.NewFatal(fmt.Sprintf("Costumer not created: %s", mistake), nil)
		return
	}
	r = created
	return
}

func (o *Costumer) checkCostumerInDB(costumerId objectID, costumerMap any) (foundCostumer costumerModel, costumerErr definitionError.GQLError) {
	if idErr := o.validateRecordID(costumerId, costumerMap); idErr != nil {
		costumerErr = idErr
		return
	}
	foundRecord, consultErr := o.consultRecordByID(costumerId)
	if consultErr != nil {
		costumerErr = consultErr
		return
	}
	foundCostumer = foundRecord
	return
}
func (o *Costumer) updateCostumer(info resolvers.ResolverInfo) (r resolvers.DataReturn, err definitionError.GQLError) {
	request := info.Args["input"].(myMap)

	costumerId := request["_id"].(objectID)
	foundCostumer, costumerErr := o.checkCostumerInDB(costumerId, request)
	if costumerErr != nil {
		err = costumerErr
		return
	}

	request["totalVisits"] = (foundCostumer.TotalVisits + 1)

	where := myMap{
		"_id": costumerId,
	}

	result, mistake := o.model.Model.Update(request, where, nil)
	if mistake != nil {
		err = definitionError.NewFatal(fmt.Sprintf("costumer not updated: %s ", mistake), nil)
	}
	r = result
	return
}
