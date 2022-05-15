package adconvert

import (
	"github.com/geiqin/oceanengine/marketing-api/core"
	"github.com/geiqin/oceanengine/marketing-api/enum"
	"github.com/geiqin/oceanengine/marketing-api/model/tools/adconvert"
)

// DeepbidRead 查询深度优化方式
// 查询ocpc、ocpm广告计划可使用的深度优化方式及相关信息
func DeepbidRead(clt *core.SDKClient, accessToken string, req *adconvert.DeepbidReadRequest) ([]enum.DeepBidType, error) {
	var resp adconvert.DeepbidReadResponse
	if err := clt.Get("2/tools/ad_convert/deepbid/read/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data.SuccessList, nil
}
