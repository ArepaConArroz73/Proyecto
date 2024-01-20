package scalars

import (
	"fmt"

	"github.com/pjmd89/gogql/lib/gql/definitionError"
	"github.com/pjmd89/gogql/lib/resolvers"
)

type DNI int64

func NewDNIScalar() (r resolvers.Scalar) {
	var scalar *DNI
	r = scalar
	return
}

func (o *DNI) Set(value interface{}) (r interface{}, err definitionError.GQLError) {
	switch value := value.(type) {
	case float64:

		r = int64(value)
		fmt.Println(r)
	case int64:
		r = value
	}
	return
}

func (o *DNI) Assess(resolved resolvers.ScalarResolved) (r interface{}, err definitionError.GQLError) {
	r = resolved.Value
	return
}
