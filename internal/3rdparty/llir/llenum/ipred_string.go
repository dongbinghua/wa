// Code generated by "stringer -linecomment -type IPred"; DO NOT EDIT.

package llenum

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[IPredEQ-0]
	_ = x[IPredNE-1]
	_ = x[IPredSGE-2]
	_ = x[IPredSGT-3]
	_ = x[IPredSLE-4]
	_ = x[IPredSLT-5]
	_ = x[IPredUGE-6]
	_ = x[IPredUGT-7]
	_ = x[IPredULE-8]
	_ = x[IPredULT-9]
}

const _IPred_name = "eqnesgesgtslesltugeugtuleult"

var _IPred_index = [...]uint8{0, 2, 4, 7, 10, 13, 16, 19, 22, 25, 28}

func (i IPred) String() string {
	if i >= IPred(len(_IPred_index)-1) {
		return "IPred(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _IPred_name[_IPred_index[i]:_IPred_index[i+1]]
}
