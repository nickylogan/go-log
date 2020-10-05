package log

import "github.com/logrusorgru/aurora/v3"

type colorScheme struct {
	debugLevel aurora.Color
	infoLevel  aurora.Color
	warnLevel  aurora.Color
	errorLevel aurora.Color
	fatalLevel aurora.Color
	panicLevel aurora.Color

	funcName   aurora.Color
	sourceLine aurora.Color
	timestamp  aurora.Color
}

var defaultScheme = colorScheme{
	debugLevel: aurora.BlueFg,
	infoLevel:  aurora.GreenFg,
	warnLevel:  aurora.YellowFg,
	errorLevel: aurora.RedFg,
	fatalLevel: aurora.RedBg | aurora.WhiteFg | aurora.BoldFm,
	panicLevel: aurora.RedBg | aurora.WhiteFg | aurora.BoldFm,

	funcName:   aurora.CyanFg,
	sourceLine: aurora.BlackFg | aurora.BrightFg,
	timestamp:  aurora.BlackFg | aurora.BrightFg,
}
