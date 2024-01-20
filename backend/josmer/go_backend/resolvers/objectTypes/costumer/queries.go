package costumer

import (
	"negocio_backend/lib"
	"negocio_backend/lib/utils"

	"github.com/pjmd89/gogql/lib/gql/definitionError"
	"github.com/pjmd89/gogql/lib/resolvers"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// pagination related function
func (o *Costumer) edges(info resolvers.ResolverInfo) (r resolvers.DataReturn, err definitionError.GQLError) {

	var opts []*options.FindOptions
	if info.Parent.(lib.Edge).FindOpts != nil {
		opts = append(opts, info.Parent.(lib.Edge).FindOpts...)

	}
	if info.Directives["sort"] != nil {
		opts = append(opts, info.Directives["sort"].(*options.FindOptions))
	}
	rr, rerr := o.model.Read(info.Parent.(lib.Edge).Where, opts)
	if rerr != nil {
		utils.Logs.System.Fatal().Println(rerr.Error())
		err = definitionError.NewFatal(rerr.Error(), nil)
		return
	}
	r = rr.([]costumerModel)
	return
}
