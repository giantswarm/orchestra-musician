package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"syscall"

	"github.com/hoisie/web"
)

const (
	cliName        = "musician"
	cliDescription = "A small tool that fetches an url and starts mplayer"
)

var (
	globalFlagset = flag.NewFlagSet(cliName, flag.ExitOnError)
	globalFlags   = struct {
		debug    bool
		version  bool
		dep      string
		filename string
	}{}

	projectVersion = "dev"
	projectBuild   string
)

func init() {
	globalFlagset.BoolVar(&globalFlags.debug, "debug", false, "Print out more debug information to stderr")
	globalFlagset.BoolVar(&globalFlags.version, "version", false, "Print the version and exit")
	globalFlagset.StringVar(&globalFlags.dep, "dep", "", "Dependency to request data from")
	globalFlagset.StringVar(&globalFlags.filename, "filename", "", "Filename to play when requested to start")
}

func musician(val string) string {
	fmt.Printf("Received start signal from conductor. Starting to play %s\n", globalFlags.filename)

	cmd := exec.Command("mplayer", "-ao", "jack", globalFlags.filename)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Printf("Error executing mplayer command: %s\n", err)
	}
	return "OK"
}

func main() {
	globalFlagset.Parse(os.Args[1:])

	// deal specially with --version
	if globalFlags.version {
		fmt.Println("musician version", projectVersion, projectBuild)
		os.Exit(0)
	}

	web.Get("/(.*)", musician)
	go web.Run("0.0.0.0:80")

	// Handle SIGINT and SIGTERM.
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	log.Println(<-ch)

	// Stop the service gracefully.
	web.Close()
}
