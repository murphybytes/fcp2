package main

import (
	"flag"
	"fmt"
)

type Arguments struct {
	server          *bool
	verbose         *bool
	noEncrypt       *bool
	pidDir          *string
	port            *int
	listenHost      *string
	fileArgs        []string

}

func NewArguments()( *Arguments ) {

	flag.Usage = func() {
		fmt.Println("Usage: fcp [options] [[[user1][@host1]]:sourcefile [[user2][@host2]]:destfile]")
		fmt.Println("Options:")
		flag.PrintDefaults()
	}


	return &Arguments{
		flag.Bool("s", false, "Run in server mode"),
		flag.Bool("v", false, "Verbose log messages"),
		flag.Bool("n", false, "Disable encryption"),
		flag.String("pid-file", "", "Path where pid file will be stored"),
		flag.Int( "p", 8069, "Port used for tcp connections." ),
		flag.String("listen-host", "localhost", "Host name for server" ),
		nil,
	}
}

func ( args *Arguments ) GetServerService( ) ( string ) {
	return fmt.Sprintf( ":%d", *args.port )
}


