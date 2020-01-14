package model

type MsgWrapper struct {
	Envelope MsgEnvelope
	Packet   IRPMessage
}

type IRPMessage struct {
	Header     MsgHeader
	Body       MsgBody
	Credential MsgCredential
}

type MsgEnvelope [20]byte
type MsgHeader [24]byte
type MsgBody []byte
type MsgCredential []byte

type OpCode int
type RespCode int
