package utils

type TimeFormat int

const (
	FORMATTIME_UNKNOWN TimeFormat = iota
	FORMATTIME_H_I
	FORMATTIME_H_I_A
)
const (
	FORMATTIME_UNKNOWN_REVERSE string = ""
	FORMATTIME_H_I_REVERSE            = "H:i"
	FORMATTIME_H_I_A_REVERSE          = "h:ia"
)
const (
	FORMATTIME_UNKNOWN_ENUM string = ""
	FORMATTIME_H_I_ENUM            = "H_i"
	FORMATTIME_H_I_A_ENUM          = "h_ia"
)

func (o TimeFormat) Set(format string) TimeFormat {
	switch format {
	case FORMATTIME_H_I_REVERSE:
		return FORMATTIME_H_I
	case FORMATTIME_H_I_A_REVERSE:
		return FORMATTIME_H_I_A
	default:
		return FORMATTIME_H_I
	}
}
func (o TimeFormat) Reverse(format string) string {
	switch format {
	case FORMATTIME_H_I_ENUM:
		return FORMATTIME_H_I_REVERSE
	case FORMATTIME_H_I_A_ENUM:
		return FORMATTIME_H_I_A_REVERSE
	default:
		return FORMATTIME_H_I_REVERSE
	}
}
func (o TimeFormat) String() string {
	switch o {
	case FORMATTIME_H_I:
		return "15:04"
	case FORMATTIME_H_I_A:
		return "03:04 pm"
	default:
		return "15:04"
	}
}
