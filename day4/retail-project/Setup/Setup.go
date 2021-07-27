package Setup

import (
	"github.com/kshitij-nawandar9/freshers-bootcamp/day4/retail-project/Config"
	"github.com/kshitij-nawandar9/freshers-bootcamp/day4/retail-project/Models/Order"
)

func AlertWorkers(){
	for i:=1;i< Config.NumberOfWorkers;i++{
		go Order.Worker()
	}
}
