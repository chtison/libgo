package shellcolors

type CodeSGR uint8

// Start and end of an SGR code: "\033[...;...;...;m".
const (
	codeSGR_start     = "\033["
	codeSGR_separator = ";"
	codeSGR_end       = "m"
)

const (
	Reset CodeSGR = iota // Cancel all previous parameters
	Bold
	Faint
	Italic
	Underline
	BlinkSlow
	BlinkRapid
	Negative
	Conceal
	CrossedOut
)

const (
	Fraktur CodeSGR = 20 + iota
	NoBold
	NoBoldAndFaint
	NoItalicAndFraktur
	NoUnderline
	NoBlink
	Reserved26
	NoNegative
	NoConceal
	NoCrossedOut
)

// Text color
const (
	Black CodeSGR = 30 + iota
	Red
	Green
	Yellow
	Blue
	Magenta
	Cyan
	White
	CustomColor // See notes
	DefaultColor
)

// Background color
const (
	BgBlack CodeSGR = 40 + iota
	BgRed
	BgGreen
	BgYellow
	BgBlue
	BgMagenta
	BgCyan
	BgWhite
	BgCustomColor
	BgDefaultColor
)

// Text color high intensity
const (
	BlackHI CodeSGR = 90 + iota
	RedHI
	GreenHI
	YellowHI
	BlueHI
	MagentaHI
	CyanHI
	WhiteHI
)

// Background color high intensity
const (
	BgBlackHI CodeSGR = 100 + iota
	BgRedHI
	BgGreenHI
	BgYellowHI
	BgBlueHI
	BgMagentaHI
	BgCyanHI
	BgWhiteHI
)
