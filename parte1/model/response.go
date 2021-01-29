package model

// {
// 	"total": 12001.00,
// 	"comprasPorTDC":{
// 	  "oro":1000,
// 	  "amex":9401
// 	},
// 	"nocompraron":100,
// 	"compraMasAlta":500
//   }

type (
	SummaryPurchase struct {
		Total         float64
		ComprasPorTDC map[string]float64
		NoCompraron   uint
		CompraMasAlta float64
	}
)

//NewSummaryPurchase return a new SummaryPurchase
func NewSummaryPurchase() *SummaryPurchase {
	return &SummaryPurchase{
		0.0,
		make(map[string]float64),
		0,
		0.0,
	}
}
