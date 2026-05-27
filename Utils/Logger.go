package Utils

import "github.com/pterm/pterm"

type LoggerUtils struct {}

func (this *LoggerUtils) Log(message string) {
	//
	pterm.Info.Printfln("%s\n", message)
}

func (this *LoggerUtils) LogError(message string) {
	//
	pterm.Error.Printfln("%s\n", message)
}

func (this *LoggerUtils) LogSuccess(message string) {
	//
	pterm.Success.Printfln("%s\n", message)
}

var Logger = &LoggerUtils {}
