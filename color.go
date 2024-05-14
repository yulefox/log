package log

const (
	FMT_RESET = 0

	FMT_BOLD       = 1
	FMT_DIM        = 2
	FMT_UNDERLINED = 4
	FMT_BLINK      = 5
	FMT_MINVERTED  = 7
	FMT_HIDDEN     = 8

	FGC_DEFAULT      = 39
	FGC_BLACK        = 30
	FGC_RED          = 31
	FGC_GREEN        = 32
	FGC_YELLOW       = 33
	FGC_BLUE         = 34
	FGC_MAGENTA      = 35
	FGC_CYAN         = 36
	FGC_LIGHTGREY    = 37
	FGC_DARKGREY     = 90
	FGC_LIGHTRED     = 91
	FGC_LIGHTGREEN   = 92
	FGC_LIGHTYELLOW  = 93
	FGC_LIGHTBLUE    = 94
	FGC_LIGHTMAGENTA = 95
	FGC_LIGHTCYAN    = 96
	FGC_LIGHTWHITE   = 97

	BGC_DEFAULT      = 49
	BGC_BLACK        = 40
	BGC_RED          = 41
	BGC_GREEN        = 42
	BGC_YELLOW       = 43
	BGC_BLUE         = 44
	BGC_MAGENTA      = 45
	BGC_CYAN         = 46
	BGC_LIGHTGREY    = 47
	BGC_DARKGREY     = 100
	BGC_LIGHTRED     = 101
	BGC_LIGHTGREEN   = 102
	BGC_LIGHTYELLOW  = 103
	BGC_LIGHTBLUE    = 104
	BGC_LIGHTMAGENTA = 105
	BGC_LIGHTCYAN    = 106
	BGC_LIGHTWHITE   = 107
)

type Shader struct {
	color string
}

func (s *Shader) do(content string) string {
	cache := bufferPool.Get().(*Buffer)
	defer cache.close()

	cache.WriteString(s.color + content + "\033[0m")

	//fmt.Fprint(cache, "\033[0;")
	//fmt.Fprintf(cache, "%d", s.color[0])
	//for i := 1; i < len(s.color); i++ {
	//	fmt.Fprintf(cache, ";%d", s.color[i])
	//}
	//fmt.Fprint(cache, "m", content, "\033[0m")

	return cache.String()
}

var (
	infoShader = Shader{
		color: "\033[0;0;97;42m",
		//color: []int{
		//	FGC_LIGHTWHITE,
		//	BGC_GREEN,
		//},
	}

	debugShader = Shader{
		color: "\033[0;0;97;44m",
		//color: []int{
		//	FGC_LIGHTWHITE,
		//	BGC_LIGHTGREY,
		//},
	}

	warnShader = Shader{
		color: "\033[0;0;97;43m",
		//color: []int{
		//	FGC_LIGHTWHITE,
		//	BGC_YELLOW,
		//},
	}

	errorShader = Shader{
		color: "\033[0;0;97;41m",
		//color: []int{
		//	FGC_LIGHTWHITE,
		//	BGC_RED,
		//},
	}
)

func shaderByLv(level Level) *Shader {
	switch level {
	case DEBU:
		return &debugShader
	case INFO:
		return &infoShader
	case WARN:
		return &warnShader
	default:
		return &errorShader
	}
}
