package desired

import "github.com/scottd018/desired/internal/compare"

func Equal(desired, actual interface{}) (bool, error) {
	return compare.Compare(desired, actual)
}
