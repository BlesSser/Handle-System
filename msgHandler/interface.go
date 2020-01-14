package msgHandler

import "handle-system/model"

type MsgHandler interface {
	Do(msg model.MsgWrapper) model.MsgWrapper
}
