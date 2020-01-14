package handleRecorder

import "handle-system/model"

type SimpleRecorder struct {
}

//CreateSimpleHandleRecorder
func CreateSimpleHandleRecorder() HandleRecorder {

	return &SimpleRecorder{}
}

func (sr *SimpleRecorder) Do(oc model.OpCode, body model.MsgBody) (model.RespCode, model.MsgBody) {

	return 0, nil
}
