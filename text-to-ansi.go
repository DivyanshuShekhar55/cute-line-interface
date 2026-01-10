package main

import (
	"fmt"
	"strconv"
	"strings"
)

func getColorAnsi(color string) string {
	colorMap := map[string]string{
		"black":         "30",
		"red":           "31",
		"green":         "32",
		"yellow":        "33",
		"blue":          "34",
		"magenta":       "35",
		"cyan":          "36",
		"white":         "37",
		"blackBright":   "90",
		"redBright":     "91",
		"greenBright":   "92",
		"yellowBright":  "93",
		"blueBright":    "94",
		"magentaBright": "95",
		"cyanBright":    "96",
		"whiteBright":   "97",
		// Add hex like "#af4cab" for true color
	}
	if code, ok := colorMap[color]; ok {
		return code
	}
	// Handle hex colors (true color)
	if strings.HasPrefix(color, "#") && len(color) == 7 {
		r, _ := strconv.ParseUint(color[1:3], 16, 8)
		g, _ := strconv.ParseUint(color[3:5], 16, 8)
		b, _ := strconv.ParseUint(color[5:7], 16, 8)
		return fmt.Sprintf("38;2;%d;%d;%d", r, g, b) // FG true color
	}
	return "37" // Default white
}

func turnText(text, color string, isBold, isUnderlined bool) string {
	var sgr []string
	sgr = append(sgr, "0") // Start with reset, but override

	colorCode := getColorAnsi(color)
	if strings.Contains(colorCode, ";") || strings.Contains(colorCode, "2;") {
		// True color or complex: use directly
		sgr[0] = colorCode
	} else {
		sgr = append(sgr, colorCode)
	}

	if isBold {
		sgr = append(sgr, "1")
	}
	if isUnderlined {
		sgr = append(sgr, "4")
	}

	return fmt.Sprintf("\033[%sm%s\033[0m", strings.Join(sgr, ";"), text)
}

func addDivider(color string, len int) string {
	
	var sgr []string
	sgr = append(sgr, "0") // Start with reset, but override

	colorCode := getColorAnsi(color)
	if strings.Contains(colorCode, ";") || strings.Contains(colorCode, "2;") {
		// True color or complex: use directly
		sgr[0] = colorCode
	} else {
		sgr = append(sgr, colorCode)
	}

	var divider string
	for i:=0; i<len; i++ {
		divider = divider + "-"
	}

	return fmt.Sprintf("\033[%sm%s\033[0m", strings.Join(sgr, ";"), divider)

}

func logError(err error) {
	err_str := "error occured"+ err.Error()
	err_log := turnText(err_str, "redBright", false, false)
	fmt.Println(err_log)
}