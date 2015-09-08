package tufer

import (
"net"

)


type Client struct {
	serverAddr *net.TCPAddr
	logger     Logger
}



func NewClient( service string, logger Logger  ) ( *Client, error ) {

	serverAddr, err := net.ResolveTCPAddr( "tcp4", service )

	if err != nil {
		return nil, err
	}

	return &Client{ serverAddr, logger }, nil

}

func ( c *Client ) Connect() ( *Connection, error ) {

	controlChannel, err := net.DialTCP("tcp", nil, c.serverAddr )

	if err != nil {
		return nil, err
	}

	return NewConnection( controlChannel, c.logger ), nil
}
