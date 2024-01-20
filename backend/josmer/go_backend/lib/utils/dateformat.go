package utils

type DateFormat int

const (
	FORMATDATE_UNKNOWN DateFormat = iota
	FORMATDATE_Y_M_D
	FORMATDATE_Y__M__D
	FORMATDATE_M_D_Y
	FORMATDATE_M__D__Y
	FORMATDATE_D_M_Y
	FORMATDATE_D__M__Y
)
const (
	FORMATDATE_UNKNOWN_REVERSE string = ""
	FORMATDATE_Y_M_D_REVERSE          = "Y-m-d"
	FORMATDATE_Y__M__D_REVERSE        = "Y/m/d"
	FORMATDATE_M_D_Y_REVERSE          = "m-d-Y"
	FORMATDATE_M__D__Y_REVERSE        = "m/d/Y"
	FORMATDATE_D_M_Y_REVERSE          = "d-m-Y"
	FORMATDATE_D__M__Y_REVERSE        = "d/m/Y"
)
const (
	FORMATDATE_UNKNOWN_ENUM string = ""
	FORMATDATE_Y_M_D_ENUM          = "Y_m_d"
	FORMATDATE_Y__M__D_ENUM        = "Y__m__d"
	FORMATDATE_M_D_Y_ENUM          = "m_d_Y"
	FORMATDATE_M__D__Y_ENUM        = "m__d__Y"
	FORMATDATE_D_M_Y_ENUM          = "d_m_Y"
	FORMATDATE_D__M__Y_ENUM        = "d__m__Y"
)

func (o DateFormat) Set(format string) DateFormat {
	switch format {
	case FORMATDATE_Y_M_D_REVERSE:
		return FORMATDATE_Y_M_D
	case FORMATDATE_Y__M__D_REVERSE:
		return FORMATDATE_Y__M__D
	case FORMATDATE_M_D_Y_REVERSE:
		return FORMATDATE_M_D_Y
	case FORMATDATE_M__D__Y_REVERSE:
		return FORMATDATE_M__D__Y
	case FORMATDATE_D_M_Y_REVERSE:
		return FORMATDATE_D_M_Y
	case FORMATDATE_D__M__Y_REVERSE:
		return FORMATDATE_D__M__Y
	default:
		return FORMATDATE_Y_M_D
	}
}
func (o DateFormat) Reverse(enum string) string {
	switch enum {
	case FORMATDATE_Y_M_D_ENUM:
		return FORMATDATE_Y_M_D_REVERSE
	case FORMATDATE_Y__M__D_ENUM:
		return FORMATDATE_Y__M__D_REVERSE
	case FORMATDATE_M_D_Y_ENUM:
		return FORMATDATE_M_D_Y_REVERSE
	case FORMATDATE_M__D__Y_ENUM:
		return FORMATDATE_M__D__Y_REVERSE
	case FORMATDATE_D_M_Y_ENUM:
		return FORMATDATE_D_M_Y_REVERSE
	case FORMATDATE_D__M__Y_ENUM:
		return FORMATDATE_D__M__Y_REVERSE
	default:
		return FORMATDATE_Y_M_D_REVERSE
	}
}
func (o DateFormat) String() string {
	switch o {
	case FORMATDATE_Y_M_D:
		return "2006-01-02"
	case FORMATDATE_Y__M__D:
		return "2006/01/02"
	case FORMATDATE_M_D_Y:
		return "01-02-2006"
	case FORMATDATE_M__D__Y:
		return "01/02/2006"
	case FORMATDATE_D_M_Y:
		return "02-01-2006"
	case FORMATDATE_D__M__Y:
		return "02/01/2006"
	default:
		return "2006-01-02"
	}
}
