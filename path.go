package ijson

import (
	"strconv"
	"strings"
)

type (
	Path uint // Type of path to get or manipulate the data
	Actn uint // Type of action to be performed on data
)

const (
	Act_Ukn Actn = iota
	Act_Get
	Act_Set
	Act_Del
)

const (
	P_Unknown Path = iota // Unknown Get path, ""

	PGet_Obj    // GET from Object - "<key>"
	PGet_ArrLen // GET from Array - "#"
	PGet_ArrIdx // GET from Array - "#<index>"
	PGet_ArrFld // GET from Array - "#~<key>"

	PSet_Obj       // Set in Object - "<key>"
	PSet_ArrIdx    // Set in Array - "#<index>"
	PSet_ArrAppend // Set in Array - "#"

	PDel_Obj      // Delete from Object - "<key>"
	PDel_ArrIdx   // Delete from Array - "#<index>"
	PDel_ArrIdxPO // Delete from Array - "#~<index>"
	PDel_ArrEnd   // Delete from Array - "#"
)

const (
	// PathArrayStart is a byte representation of "#"
	PathArrayStart byte = 35
	// PathWildCard is a byte representation of "~"
	PathWildCard byte = 126
)

const (
	PathArrayStartStr string = "#"
	PathWildCardStr   string = "#~"
)

func DetectGetPath(p string) Path {
	cnt := len(p)

	if cnt == 0 {
		return P_Unknown
	}

	if cnt >= 1 && p[0] == PathArrayStart {
		if cnt == 1 {
			return PGet_ArrLen
		}

		if p[1] == PathWildCard {
			return PGet_ArrFld
		}

		return PGet_ArrIdx
	}
	return PGet_Obj
}

func DetectSetPath(p string) Path {
	cnt := len(p)

	if cnt == 0 {
		return P_Unknown
	}

	if cnt >= 1 && p[0] == PathArrayStart {
		if cnt == 1 {
			return PSet_ArrAppend
		}

		return PSet_ArrIdx
	}

	return PSet_Obj
}

func DetectDelPath(p string) Path {
	cnt := len(p)

	if cnt == 0 {
		return P_Unknown
	}

	if cnt >= 1 && p[0] == PathArrayStart {
		if cnt == 1 {
			return PDel_ArrEnd
		}

		if p[1] == PathWildCard {
			return PDel_ArrIdxPO
		}

		return PDel_ArrIdx
	}

	return PDel_Obj
}

// DetectDelPath returns the type of path for the provided action. Returns P_Unknown for invalid path.
func DetectPath(a Actn, p string) Path {
	switch a {
	case Act_Get:
		return DetectGetPath(p)
	case Act_Set:
		return DetectSetPath(p)
	case Act_Del:
		return DetectDelPath(p)
	default:
		return P_Unknown
	}
}

func index(p string, t Path) (int, error) {
	// we have already resolved the path type and it is valid.
	var i string
	switch t {
	case PGet_ArrIdx, PSet_ArrIdx, PDel_ArrIdx:
		i = p[1:]
	case PDel_ArrIdxPO:
		i = p[2:]
		// default:
		// 	return 0, errors.New("unknown path to resolve index")
	}

	return strconv.Atoi(i)
}

func field(p string) string {
	// we have already resolved the path type and it is valid.
	return p[2:]
}

func split(p string) []string {
	return strings.Split(p, ".")
}
