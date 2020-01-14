package handleRecorder

import "handle-system/model"

type HandleRecorder interface {
	Do(oc model.OpCode, body model.MsgBody) (model.RespCode, model.MsgBody)
}
