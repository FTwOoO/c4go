package program

// DefinitionType - conversion map from C standard library structures to
// c4go structures
var DefinitionType = map[string]string{
	// time.h
	"time_t": "github.com/FTwOoO/c4go/noarch.TimeT",
	"github.com/FTwOoO/c4go/noarch.TimeT": "int32",
	"__time_t":      "int32",
	"__suseconds_t": "int32",

	"fpos_t": "int32",

	// unistd.h
	"ssize_t": "github.com/FTwOoO/c4go/noarch.SsizeT",
	"github.com/FTwOoO/c4go/noarch.SsizeT": "long int",

	// built-in
	"bool":                   "bool",
	"char *":                 "[]byte",
	"char":                   "byte",
	"char*":                  "[]byte",
	"double":                 "float64",
	"float":                  "float32",
	"int":                    "int32",
	"long double":            "float64",
	"long int":               "int32",
	"long long":              "int64",
	"long long int":          "int64",
	"long long unsigned int": "uint64",
	"long unsigned int":      "uint32",
	"long":                   "int32",
	"short":                  "int16",
	"signed char":            "int8",
	"unsigned char":          "uint8",
	"unsigned int":           "uint32",
	"unsigned long long":     "uint64",
	"unsigned long":          "uint32",
	"unsigned short":         "uint16",
	"unsigned short int":     "uint16",
	"void":                   "",
	"_Bool":                  "int32",
	"size_t":                 "uint32",
	"ptrdiff_t":              "github.com/FTwOoO/c4go/noarch.PtrdiffT",
	"github.com/FTwOoO/c4go/noarch.PtrdiffT": "uint64",
	"wchar_t": "github.com/FTwOoO/c4go/noarch.WcharT",
	"github.com/FTwOoO/c4go/noarch.WcharT": "rune",

	// void*
	"void*":  "interface{}",
	"void *": "interface{}",

	// null is a special case (it should probably have a less ambiguous name)
	// when using the NULL macro.
	"null": "null",

	// sys/stat.h
	"mode_t":   "uint16",
	"__mode_t": "uint16",

	"FILE": "github.com/FTwOoO/c4go/noarch.File",

	// termios.h
	"tcflag_t": "uint32",
	"cc_t":     "uint8",

	// sys/resource.h
	"__rusage_who":   "int",
	"__rusage_who_t": "int",

	// time.h
	"clock_t": "github.com/FTwOoO/c4go/noarch.ClockT",
	"github.com/FTwOoO/c4go/noarch.ClockT": "int64",

	// signal.h
	"sig_atomic_t": "int64",

	// sys/types.h
	"off_t":   "int64",
	"__off_t": "int64",
}
