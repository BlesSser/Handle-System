package certificate

import "handle-system/model"

type Certificate interface {
	Verify(msg model.IRPMessage) bool
	Sign(header model.MsgHeader, body model.MsgBody) model.IRPMessage
}
