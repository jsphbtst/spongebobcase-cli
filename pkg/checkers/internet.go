package checkers

import (
	"net"
	"time"
)

func HasInternet() bool {
	timeout := 5 * time.Second
	conn, err := net.DialTimeout("tcp", "8.8.8.8:53", timeout)

	if err != nil {
		return false
	}

	conn.Close()
	return true
}
