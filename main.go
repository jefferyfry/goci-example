package main

import (
	"github.com/jefferyfry/funclog"
	"goci-example/api"
)

var (
	LogE = funclog.NewErrorLogger("ERROR: ")
)

func main() {
	LogE.Fatal(api.StartApiService())
}