package main

import(
	"github.com/yqh231/CoinQuantitave/monitor/decimal"
	"github.com/yqh231/CoinQuantitave/monitor/mongo"
	"time"
)


func taskBreakPoint(){
	markets := mongo.MarketAll()
	var flag bool
	for _, market:= range markets{
		top := topPrice(market.Lookup("market").StringValue())
		cur := curPrice(market.Lookup("market").StringValue())
		if decimal.CompareStr(cur, top) == 1{
			flag = true	
		}
		mongo.BpInsertOne(market.Lookup("market").StringValue(), cur, top, flag)
		time.Sleep(5000 * time.Millisecond)
	}
}