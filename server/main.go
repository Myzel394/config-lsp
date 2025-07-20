package main

import (
	"config-lsp/common"
	roothandler "config-lsp/root-handler"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/getsentry/sentry-go"
)

func main() {
	if len(os.Args) > 1 && (os.Args[1] == "--version" || os.Args[1] == "version") {
		fmt.Println(roothandler.Version)

		os.Exit(0)
		return
	}

	common.InitServerOptions()

	if common.ServerOptions.AllowedUsageReports != common.UsageReportNone {
		var environment string

		if common.ServerOptions.IsDebug {
			environment = "debug"
		} else {
			environment = "production"
		}

		err := sentry.Init(sentry.ClientOptions{
			// Either set your DSN here or set the SENTRY_DSN environment variable.
			Dsn: "https://76a930e8b68a92fb2993844e79e08b07@sentry.myzel394.app/2",
			// Enable printing of SDK debug messages.
			// Useful when getting started or trying to figure something out.
			Debug:       common.ServerOptions.IsDebug,
			Environment: environment,

			EnableTracing:    true,
			TracesSampleRate: 1.0,
			EnableLogs:       true,
			SampleRate:       1.0,

			SendDefaultPII: false,
		})

		if err != nil {
			log.Fatal("Sentry failed to initialize:", err.Error())
		}

		defer sentry.Flush(2 * time.Second)
	} else {
		println("config-lsp server started with anonymous usage reports disabled")
	}

	common.SetupLogging()

	roothandler.SetUpRootHandler()
}
