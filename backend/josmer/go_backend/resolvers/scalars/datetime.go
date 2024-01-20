package scalars

import (
	"fmt"
	"regexp"
	"strconv"
	"time"

	"github.com/pjmd89/gogql/lib/gql/definitionError"
	"github.com/pjmd89/gogql/lib/resolvers"
)

type DateTime int64

func NewDateTimeScalar() (r resolvers.Scalar) {
	var scalar *DateTime
	r = scalar
	return
}

func (o *DateTime) Set(value interface{}) (r interface{}, err definitionError.GQLError) {
	r = int64(0)
	switch valueString := value.(type) {
	case string:

		re := regexp.MustCompile(`^(\d{2})/(\d{2})/(\d{4})\s(0[0-9]|1\d|2[0-3]):(0[0-9]|1\d|2\d|3\d|4\d|5\d)$`)
		if !re.MatchString(valueString) {
			err = definitionError.NewFatal("Date string ("+valueString+") is wrong", nil)
			fmt.Println("Date string (" + valueString + ") is wrong")
			return
		}
		result := re.FindStringSubmatch(valueString)
		year, _ := strconv.ParseInt(result[3], 10, 32)
		month, _ := strconv.ParseInt(result[2], 10, 32)
		day, _ := strconv.ParseInt(result[1], 10, 32)
		hour, _ := strconv.ParseInt(result[4], 10, 32)
		minutes, _ := strconv.ParseInt(result[5], 10, 32)

		location, _ := time.LoadLocation("America/Caracas")
		timeObject := time.Date(int(year), time.Month(month), int(day), int(hour), int(minutes), 0, 0, location)
		//r = Birthday(timeObject.Unix())
		r = int64(timeObject.Unix())
	}
	return
}

func (o *DateTime) Assess(resolved resolvers.ScalarResolved) (r interface{}, err definitionError.GQLError) {
	r = resolved.Value
	return
}
