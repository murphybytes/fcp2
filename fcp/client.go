package main

import (
	"github.com/murphybytes/fcp2/tufer"
	"fmt"
)

type Client struct {
	ctx *Context
}

func NewClient( ctx *Context ) ( *Client, error ) {
	return &Client{ ctx }, nil
}

//
// Entry point for client code
//
func ( c *Client ) Execute( ) {
	c.ctx.LogDebug( "Starting client" )
	commandParser, _ := NewClientCommandParser( c.ctx )
	_, dest, err := commandParser.Parse()

	if err != nil {
		c.ctx.LogFatal( err )
	}

	destService := c.GetService( dest )

	client, err := tufer.NewClient(destService, c.ctx )

	if err != nil {
		c.ctx.LogFatal( err )
	}

	//var connection *tufer.Connection
	connection, err := client.Connect()

	if err != nil {
		c.ctx.LogFatal( err )
	}

	defer connection.Close()

	c.ctx.LogDebug( "Connection to ", destService, " succeeds" )

  err = c.Handshake( connection )

	if err != nil {
		c.ctx.LogFatal( err )
	}

	c.ctx.LogDebug( "Handshake succeeds")

}




func (c *Client) Handshake(conn *tufer.Connection)( err error) {
	msgs, err := conn.ReadControlMessages()

	if err != nil {
		return
	}

	v := NewFatalValidator( c.ctx )

	v.ValidateInt( len(msgs), 2 )
	v.ValidateString( msgs[0], INITIAL_SERVER_ID )


	c.ctx.LogInfo( "Handshake:", msgs )

	err = conn.WriteControlMessages( INITIAL_CLIENT_ID, CLIENT_CONTROL_PROTOCOL_VERSION )

	return
}

func (c *Client) GetService( fileInfo *FileInformation )( string ) {
	return fmt.Sprintf( "%s:%d", fileInfo.hostName, *c.ctx.arguments.port )
}
