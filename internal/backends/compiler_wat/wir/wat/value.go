// 版权 @2022 凹语言 作者。保留所有权利。

package wat

func NewVar(name string, typ ValueType) Value { return &aVar{name: name, typ: typ} }
func NewVarI32(name string) Value             { return &aVar{typ: I32{}, name: name} }
func NewVarU32(name string) Value             { return &aVar{typ: U32{}, name: name} }
func NewVarI64(name string) Value             { return &aVar{typ: I64{}, name: name} }
func NewVarU64(name string) Value             { return &aVar{typ: U64{}, name: name} }
func NewVarF32(name string) Value             { return &aVar{typ: F32{}, name: name} }
func NewVarF64(name string) Value             { return &aVar{typ: F64{}, name: name} }

/**************************************
I32:
**************************************/
type I32 struct {
}

func (t I32) Name() string           { return "i32" }
func (t I32) Equal(u ValueType) bool { _, ok := u.(I32); return ok }

/**************************************
U32:
**************************************/
type U32 struct {
}

func (t U32) Name() string           { return "i32" }
func (t U32) Equal(u ValueType) bool { _, ok := u.(U32); return ok }

/**************************************
I64:
**************************************/
type I64 struct {
}

func (t I64) Name() string           { return "i64" }
func (t I64) Equal(u ValueType) bool { _, ok := u.(I64); return ok }

/**************************************
U64:
**************************************/
type U64 struct {
}

func (t U64) Name() string           { return "i64" }
func (t U64) Equal(u ValueType) bool { _, ok := u.(U64); return ok }

/**************************************
F32:
**************************************/
type F32 struct {
}

func (t F32) Name() string           { return "f32" }
func (t F32) Equal(u ValueType) bool { _, ok := u.(F32); return ok }

/**************************************
F64:
**************************************/
type F64 struct {
}

func (t F64) Name() string           { return "f64" }
func (t F64) Equal(u ValueType) bool { _, ok := u.(F64); return ok }

/**************************************
aVar:
**************************************/
type aVar struct {
	name string
	typ  ValueType
}

func (v *aVar) Name() string    { return v.name }
func (v *aVar) Type() ValueType { return v.typ }
