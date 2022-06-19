package gui

import (
	"github.com/jroimartin/gocui"
)

var gocuiColorMap = map[string]gocui.Attribute{
	"default":   gocui.ColorDefault,
	"black":     gocui.ColorBlack,
	"red":       gocui.ColorRed,
	"green":     gocui.ColorGreen,
	"yellow":    gocui.ColorYellow,
	"blue":      gocui.ColorBlue,
	"magenta":   gocui.ColorMagenta,
	"cyan":      gocui.ColorCyan,
	"white":     gocui.ColorWhite,
	"bold":      gocui.AttrBold,
	"reverse":   gocui.AttrReverse,
	"underline": gocui.AttrUnderline,
}

func GetGocuiAttribute(key string) gocui.Attribute {
	value, present := gocuiColorMap[key]
	if present {
		return value
	}
	return gocui.ColorWhite
}

func GetGocuiStyle(keys []string) gocui.Attribute {
	var attribute gocui.Attribute
	for _, key := range keys {
		attribute |= GetGocuiAttribute(key)
	}

	return attribute
}
