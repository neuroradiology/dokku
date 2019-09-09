package main

import (
	"flag"

	"github.com/dokku/dokku/plugins/resource"
)

// displays a resource report for one or more apps
func main() {
	flag.Parse()
	appName := flag.Arg(0)

	resource.ReportSingleApp(appName, "")
}
