package listener

import (
	"github.com/sirupsen/logrus"
	"net"
)

func HandleRequest(c net.Conn) {

	requestByte := make([]byte, 10)
	_, err := c.Read(requestByte)
	if err != nil {
		logrus.Error("read buffer error")
	}

}
