package main

import (
	"path/filepath"

	"github.com/johnayoung/go-wow-bot/pkg/inject"
	"github.com/johnayoung/go-wow-bot/utils"
)

func main() {
	srcPath := filepath.Join(utils.RootDir(), "bloog", "bloogq.exe")
	dllPath := filepath.Join(utils.RootDir(), "cmd", "dllmain", "mainSRC.dll")

	inject.InjectDLL32(srcPath, dllPath)

	for {
	}
}
