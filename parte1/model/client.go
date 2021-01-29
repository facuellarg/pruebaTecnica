// [{
// 	"clientId":1000223,
// 	"phone":"992003040",
// 	"nombre":"Juan Mata",
// 	"compro":true,
// 	"tdc":"gold",
// 	"monto":123.20,
// 	"date":"2020-02-20T14:53:00Z"
// 	}]
package model

import "time"

type Client struct {
	ClientId int
	Phone    string
	Nombre   string
	Compro   bool
	Tdc      string
	Monto    float64
	Date     time.Time
}
