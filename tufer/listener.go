package tufer


import (
	"net"
)


type Listener struct {
	listener *net.TCPListener
	tcpAddr  *net.TCPAddr
	logger   Logger
}


func NewListener(service string, logger Logger )( *Listener, error ) {

	tcpAddr, err := net.ResolveTCPAddr( "tcp4", service )

	if err != nil {
		return nil, err
	}

	listener, err := net.ListenTCP( "tcp", tcpAddr )

	if err != nil {
		return nil, err
	}

	return &Listener{listener, tcpAddr, logger}, nil

}

func ( l *Listener ) Accept()( *Connection, error ) {
	controlChannel, err := l.listener.Accept()

	if err != nil {
		return nil, err
	}

	return NewConnection(controlChannel.(*net.TCPConn), l.logger ), nil

}

func (l *Listener ) Close() {
	l.listener.Close()
}
