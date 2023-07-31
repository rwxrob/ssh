package ssh

import "time"

// TCPTimeout is the default number of seconds to wait to complete a TCP
// connection.
var TCPTimeout = 300 * time.Second
