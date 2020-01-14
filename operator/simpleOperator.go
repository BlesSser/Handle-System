package operator

import "handle-system/model"

type SimpleOperator struct {
}

func (so *SimpleOperator) Do(header model.MsgHeader, body model.MsgBody) (model.MsgHeader, model.MsgBody) {

	return [24]byte{}, nil
}
