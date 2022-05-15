package event

import (
	"github.com/geiqin/oceanengine/marketing-api/core"
	"github.com/geiqin/oceanengine/marketing-api/model/tools/event"
)

// EventConfigsGet 获取已创建事件列表
// 此接口用于查询资产下已经创建的事件及其相关信息
func EventConfigsGet(clt *core.SDKClient, accessToken string, req *event.EventConfigsGetRequest) ([]event.EventConfig, error) {
	var resp event.EventConfigsGetResponse
	err := clt.Get("2/tools/event_manager/event_configs/get/", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	return resp.Data.List, nil
}
