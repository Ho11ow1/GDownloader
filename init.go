package main

import (
	"flag"
	"os"
	"strings"

	"GDownloader/Common"
	"GDownloader/Utils"
)

func init(){
	//
	url, filePath,  limit,  prefix, extension := InitFlags()
	flag.Parse()
	AssignFlags(url, filePath, limit, prefix, extension)
}

func InitFlags() (*string, *string, *uint, *string, *string){
	//
	return flag.String("url", "", "Single url"),
		flag.String("file", "", "Path to multi-url file"),
		flag.Uint("limit", 0, "Only download this many files"),
		flag.String("prefix", "", "Only download files starting with this string"),
		flag.String("extension", "", "Only download files with this extension")
}

func AssignFlags(url *string, filePath *string, limit *uint, prefix *string, extension *string){
	//
	set := map[string]bool{}
	flag.Visit(func (f *flag.Flag){
		set[f.Name] = true
	});

	if !set["url"] && !set["file"]{
        Utils.Logger.LogError("Must provide -url or -file")
        os.Exit(1)
    }
    if set["url"] && set["file"]{
        Utils.Logger.LogError("Wse either -url or -file, not both")
        os.Exit(1)
    }

	if set["url"]{
		Common.AppConfig.Urls = make([]string, 1)
		Common.AppConfig.Urls[0] = *url
	}

	if set["file"]{
		_, err := os.Open(*filePath)
		if err != nil{
			Utils.Logger.LogError("File does not exist")
        	os.Exit(1)
		}
		Common.AppConfig.Urls = GetUrlsFromFile(*filePath)
	}

	if set["limit"]{
		Common.AppConfig.Limit = limit
	}

	if set["prefix"]{
		Common.AppConfig.Prefix = prefix
	}

	if set["extension"]{
		Common.AppConfig.Extension = extension
	}
}

func GetUrlsFromFile(filePath string) []string{
	//
	content, err := os.ReadFile(filePath)
	if err != nil{
		Utils.Logger.LogError("Could not read file contents at " + filePath)
		os.Exit(1)
	}

	lines := strings.Split(strings.ReplaceAll(strings.TrimSpace(string(content)), "\r\n", "\n"), "\n")

	var urls []string

	for _, line := range lines{
		if strings.TrimSpace(line) != ""{
			urls = append(urls, line)
		}
	}

	return urls
}
