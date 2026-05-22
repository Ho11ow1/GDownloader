package Utils

import "github.com/pterm/pterm"

type LoggerUtils struct{}

func (this LoggerUtils) Log(message string){
	//
	pterm.Info.Println(message)
}

func (this LoggerUtils) LogMessage(message string){
	//
	pterm.Println(message)
}

func (this LoggerUtils) LogError(message string){
	//
	pterm.Error.Println(message)
}

func (this LoggerUtils) LogWarning(message string){
	//
	pterm.Warning.Println(message)
}

func (this LoggerUtils) LogSuccess(message string){
	//
	pterm.Success.Println(message)
}

var Logger = &LoggerUtils{}
