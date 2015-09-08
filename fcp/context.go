package main

import (
	"log"
	"os"
	"io/ioutil"
	"strconv"
)


type Context struct {
	arguments *Arguments
	debugLog       *log.Logger
	infoLog        *log.Logger
	warnLog        *log.Logger
	errorLog       *log.Logger
	fatalLog       *log.Logger
}

func NewContext()( *Context ) {
	return &Context{
		NewArguments(),
		log.New(os.Stderr, "DEBUG ", log.Lshortfile | log.LstdFlags ),
		log.New(os.Stderr, "INFO ", log.Lshortfile | log.LstdFlags ),
		log.New(os.Stderr, "WARN ", log.Lshortfile | log.LstdFlags ),
		log.New(os.Stderr, "ERROR ", log.Lshortfile | log.LstdFlags),
		log.New(os.Stderr, "FATAL ", log.Lshortfile | log.LstdFlags),
	}
}


func (ctx *Context) LogDebug(  v ...interface{} ) {
	if *ctx.arguments.verbose {
		ctx.debugLog.Println( v... )
	}
}

func (ctx *Context ) LogInfo( v ...interface{} ) {
	ctx.infoLog.Println( v... )
}

func (ctx *Context ) LogWarn( v ...interface{} ) {
	ctx.warnLog.Println( v... )
}

func (ctx *Context ) LogError( v ...interface{} ) {
	ctx.errorLog.Println( v... )
}

// logs message and terminates application
func (ctx *Context ) LogFatal( v ...interface{} ) {
	ctx.RemovePidFile()
	ctx.fatalLog.Fatalln( v... )
}

func (ctx *Context) CreatePidFile() {
	if *ctx.arguments.pidDir != "" {
		pidPath := *ctx.arguments.pidDir + "/fcp.pid"
		ctx.LogInfo( "Pid file will written to:", pidPath )

		if _, err := os.Stat( pidPath ); os.IsNotExist(err) {
			buff := []byte( strconv.Itoa( os.Getpid() ) )

			ioutil.WriteFile( pidPath, buff, 0644 )
		} else {
			ctx.LogFatal( "PID file exists, fcp server is either running or crashed unexpectantly.")
		}

	} else {
		ctx.LogDebug( "Pid file will not be used.")
	}
}

func (ctx *Context) RemovePidFile() {
	if *ctx.arguments.pidDir != "" {
		pidPath := *ctx.arguments.pidDir + "/fcp.pid"
		ctx.LogInfo( "Removing Pidfile:", pidPath )
		err := os.Remove( pidPath )

		if err != nil {
			ctx.LogWarn( "Error occurred while removing pid file: ", err )
		}

	}

}
