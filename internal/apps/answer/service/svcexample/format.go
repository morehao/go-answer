package svcexample

import (
	"go-answer/internal/apps/answer/dto/dtoexample"

	"github.com/gin-gonic/gin"
)

type FormatSvc interface {
	FormatRes(ctx *gin.Context) *dtoexample.FormatResResp
}

var _ FormatSvc = (*formatSvc)(nil)

type formatSvc struct {
}

func NewFormatSvc() FormatSvc {
	return &formatSvc{}
}

func (svc *formatSvc) FormatRes(ctx *gin.Context) *dtoexample.FormatResResp {
	return &dtoexample.FormatResResp{
		Items: []dtoexample.FormatDataItem{
			{
				Price:     1.22245,
				PriceList: []float64{1.22245, 1.22255},
			},
		},
		FormatDataItem: dtoexample.FormatDataItem{
			Price:     1.22245,
			PriceList: []float64{1.22245, 1.22255},
		},
		ItemMap: map[string]dtoexample.FormatDataItem{
			"1": {
				Price:     1.22245,
				PriceList: []float64{1.22245, 1.22255},
			},
			"2": {
				Price: 1.22245,
			},
		},
		NameMap: map[string][]string{
			"a": []string{},
		},
		PriceList: []float64{1.22245, 1.22255},
		PriceMap: map[string]float64{
			"1": 1.22245,
		},
	}
}
