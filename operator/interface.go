package operator

import "handle-system/model"

type Operator interface {
	Do(header model.MsgHeader, body model.MsgBody) (model.MsgHeader, model.MsgBody)
}
