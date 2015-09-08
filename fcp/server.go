package main

import (
	"time"
	"github.com/murphybytes/fcp2/tufer"
)

func RunServer( c chan int, ctx *Context ) {

	go Listen( ctx )

WaitLoop:
	for  {

		select {
		case <-c :
			break WaitLoop
		default:
			duration := 100 * time.Millisecond
			time.Sleep(duration)
		}

	}

	ctx.LogDebug( "Exiting server loop")
	c <- 1

}


func Listen( ctx *Context ) {

	ctx.LogDebug( "Started listener routine" )
	listener, err := tufer.NewListener( ctx.arguments.GetServerService(), ctx )

	if err != nil {
		ctx.LogFatal( "Error creating listener socket. Error:", err )
	}

	defer listener.Close()

	for {
		connection, err := listener.Accept()

		if err != nil {
			ctx.LogWarn( "Accept returned:", err )
		} else {
			go HandleConnection( connection, ctx)
			ctx.LogDebug("Accept Success")
		}

	}

	ctx.LogDebug( "listen routine exits")

}

func HandleConnection( conn *tufer.Connection, ctx *Context ) {
	defer conn.Close()
	ctx.LogDebug("Handle Connection")
	err := Handshake(conn, ctx)

	if err != nil {
		ctx.LogWarn( "Handshake failed:", err)
	}

}

func Handshake(conn *tufer.Connection, ctx *Context)( err error ) {

	conn.SetDeadline(5)
	defer conn.ClearDeadline()

	// send id to client
	err = conn.WriteControlMessages(INITIAL_SERVER_ID, SERVER_CONTROL_PROTOCOL_VERSION)


	if err != nil {
		return
	}

	ctx.LogDebug( "Wrote Control Messages")
	// client should send proper response
	msgs, err := conn.ReadControlMessages()

	if err != nil {
		return
	}

	ctx.LogDebug( "Handshake:", msgs)

	validator := NewValidator()
	err = validator.ValidateInt( 2, len( msgs ) )
	if err != nil {
		return
	}
	validator.ValidateString( INITIAL_CLIENT_ID, msgs[0] )
	if err != nil {
		return
	}


	return

}
