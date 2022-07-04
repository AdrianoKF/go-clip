/*
Copyright © 2022 Adrian Rumpold <a.rumpold@gmail.com>
*/
package main

import (
	"github.com/AdrianoKF/go-clip/cmd"
	"github.com/AdrianoKF/go-clip/internal/util"
)

func main() {
	util.InitializeLogging(true)
	cmd.Execute()
}
