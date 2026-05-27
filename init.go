package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"sync/atomic"

	"GDownloader/Common"

	"github.com/mattn/go-colorable"
	"github.com/pterm/pterm"
)

func init() {
	//
	SetConsoleBehaviour()

	url, filePath,  limit,  prefix, extension := InitFlags()
	flag.Parse()
	AssignFlags(url, filePath, limit, prefix, extension)
}

func SetConsoleBehaviour() {
	//
	pterm.SetDefaultOutput(colorable.NewColorableStdout())
}

func InitFlags() (*string, *string, *uint64, *string, *string) {
	//
	return flag.String("url", "", "Single url"),
		flag.String("file", "", "Path to multi-url file"),
		flag.Uint64("limit", 0, "Only download this many files"),
		flag.String("prefix", "", "Only download files starting with this string"),
		flag.String("extension", "", "Only download files with this extension")
}

func AssignFlags(url *string, filePath *string, limit *uint64, prefix *string, extension *string){
	//
	set := map[string]bool{}
	flag.Visit(func (f *flag.Flag) {
		set[f.Name] = true
	});

	if !set["url"] && !set["file"] {
        fmt.Println("Must provide -url or -file")
        os.Exit(1)
    }
    if set["url"] && set["file"] {
        fmt.Println("Wse either -url or -file, not both")
        os.Exit(1)
    }

	if set["url"] {
		Common.AppConfig.Urls = make([]string, 1)
		Common.AppConfig.Urls[0] = *url
	}

	if set["file"] {
		_, err := os.Open(*filePath)
		if err != nil{
			fmt.Println("File does not exist")
        	os.Exit(1)
		}

		Common.AppConfig.Urls = GetUrlsFromFile(*filePath)
	}

	if set["limit"] {
		Common.AppConfig.Limit = &atomic.Uint64{}
        Common.AppConfig.Limit.Store(*limit)
	}

	if set["prefix"] {
		Common.AppConfig.Prefix = prefix
	}

	if set["extension"] {
		Common.AppConfig.Extension = extension
	}
}

func GetUrlsFromFile(filePath string) []string {
	//
	content, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("Could not read file contents at " + filePath)
		os.Exit(1)
	}

	lines := strings.Split(strings.ReplaceAll(strings.TrimSpace(string(content)), "\r\n", "\n"), "\n")

	var urls []string

	for _, line := range lines{ 
		if strings.TrimSpace(line) != "" {
			urls = append(urls, line)
		}
	}

	return urls
}
