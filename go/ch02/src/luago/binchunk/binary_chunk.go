package binchunk

const (
	LUA_SIGNATUR     = "\x1bLua"
	LUAC_VERSION     = 0x53
	LUAC_FORMAT      = 0
	LUAC_DATA        = "\x19\x93\r\n\x1a\n"
	CINT_SIZE        = 4
	CSIZET_SIZE      = 8
	INSTRUCTION_SIZE = 4
	LUA_INTEGER_SIZE = 8
	LUA_NUMBER_SIZE  = 8
	LUAC_INT         = 0x5678
	LUAC_NUM         = 370.5
)
const (
	TAG_NIL       = 0x00
	TAG_BOOLEAN   = 0x01
	TAG_NUMBER    = 0x02
	TAG_INTEGER   = 0x03
	TAG_SHORT_STR = 0x04
	TAG_LONG_STR  = 0x05
)

type header struct {
	signature       [4]byte
	version         byte
	format          byte
	luacData        [6]byte
	cintSize        byte
	sizetSize       byte
	instructionSize byte //Lua虚拟机指令大小
	luaIntegerSize  byte
	luaNumberSize   byte
	luacInt         int64
	luacNum         float64
}
type Upvalue struct {
	Instack byte
	Idx     byte
}
type LocVar struct {
	VarName string
	StartPC uint32 //起始位置
	EndPC   uint32 //结束位置
}
type Prototype struct {
	Source          string        // name of the source file
	LineDefined     uint32        // first line where the function is defined
	LastLineDefined uint32        // last line where the function is defined
	NumParams       byte          // number of parameters
	IsVararg        byte          // is the function vararg?是否为可变参数
	MaxStackSize    byte          //寄存器数量
	Code            []uint32      //指令表
	Constants       []interface{} //常量表
	Upvalues        []Upvalue     //
	Protos          []*Prototype  //子函数表
	LineInfo        []uint32      //行号表
	LocVars         []LocVar      //局部变量表
	UpvalueNames    []string
}

type binaryChunk struct {
	header
	sizeUpvalues byte
	mainFunc     *Prototype
}

func Undump(data []byte) *Prototype {
	reader := &reader{data}
	reader.checkHeader()
	reader.readByte()
	return reader.readProto("")
}
