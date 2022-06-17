package desired

import "github.com/nukleros/desired/internal/compare"

func Equal(desired, actual interface{}) (bool, error) {
	return compare.Compare(desired, actual)
}
