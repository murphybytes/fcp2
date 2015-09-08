package main

// process return codes
const (
  ERROR_GENERIC int         = 1
  ERROR_IMPROPER_USAGE int  = 2

  CONTROL_BUFFER_SZ                  = 512

  // Control messages are comprised of upper and lowercase ascii alphabet
  // letters, numbers, and the following characters _,.,-
  // Messages are delimited by spaces on the wire

  // Handshake messages
  INITIAL_SERVER_ID                   =  "fcp_server"
  INITIAL_CLIENT_ID                   =  "fcp_client"
  SERVER_CONTROL_PROTOCOL_VERSION     =  "1.0"
  CLIENT_CONTROL_PROTOCOL_VERSION     =  "1.0"


)
