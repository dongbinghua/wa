// Code generated by "stringer -linecomment -type NameTableKind"; DO NOT EDIT.

package llenum

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[NameTableKindDefault-0]
	_ = x[NameTableKindGNU-1]
	_ = x[NameTableKindNone-2]
}

const _NameTableKind_name = "DefaultGNUNone"

var _NameTableKind_index = [...]uint8{0, 7, 10, 14}

func (i NameTableKind) String() string {
	if i >= NameTableKind(len(_NameTableKind_index)-1) {
		return "NameTableKind(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _NameTableKind_name[_NameTableKind_index[i]:_NameTableKind_index[i+1]]
}
