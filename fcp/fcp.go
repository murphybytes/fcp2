package main

import (
	"flag"
	"os"
	"os/signal"
)



func main() {

	ctx := NewContext()
	flag.Parse()

	if  *ctx.arguments.server {
		// create pid if user wants it
		ctx.CreatePidFile()
		signalChannel := make(chan os.Signal, 1)
		signal.Notify( signalChannel, os.Kill, os.Interrupt )

		ctx.LogInfo("Starting in server mode.")
		// fire up in server mode
		serverChannel    := make(chan int )

		go RunServer( serverChannel,  ctx )

		sig := <-signalChannel
		ctx.LogDebug("Got signal: ", sig )
		// signal server goroutine to shut down
		serverChannel<- 1
		// block until it stops
		<-serverChannel
		// get rid of pid file if one was created
		ctx.RemovePidFile()
		ctx.LogInfo("Server stopped, exiting...")


	} else {
		// if we are here we are in interactive mode, look for file source and
		// destination as last two arguments
		ctx.arguments.fileArgs = flag.Args()

		if len(ctx.arguments.fileArgs) != 2 {
			flag.Usage()
			os.Exit( ERROR_IMPROPER_USAGE )
		}

		client, err := NewClient( ctx )

		if err != nil {
			ctx.LogFatal( "Client shut down with error:", err )
		}

		client.Execute()


	}


}
