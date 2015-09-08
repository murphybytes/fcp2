package tufer

import (
	"time"
	"net"
)

const (
	CONTROL_MESSAGE_BUFFER_SIZE   = 512
	)

type Connection struct {
	controlChannel *net.TCPConn
	logger         Logger
	buffer         []byte
}

func NewConnection( controlChannel *net.TCPConn, logger Logger )( *Connection ) {
	return &Connection{ controlChannel, logger, make([]byte, CONTROL_MESSAGE_BUFFER_SIZE) }
}

func ( c *Connection ) Close() {
	c.controlChannel.Close()
}

func (c *Connection ) ReadControlMessages() ( msgs []string, err error ) {
	read, err := c.controlChannel.Read(c.buffer[0:])

	if err != nil {
		return
	}

	msgs = Unpack( string(c.buffer[0:read]))
	return
}

func( c *Connection ) WriteControlMessages( msg ...string )( err error ) {
	packed := PackArray( msg )
	_, err = c.controlChannel.Write([]byte(packed))
	return
}

func( c *Connection ) SetDeadline( seconds int ) {

	deadline := time.Now().Add( time.Duration(seconds) * time.Second )
	c.controlChannel.SetDeadline(deadline)
}

func( c *Connection ) ClearDeadline( ) {
	c.controlChannel.SetDeadline(time.Unix(0,0))
}
