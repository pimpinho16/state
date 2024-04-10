package main

import "state/state"

func main() {
	initialized := state.Initialized{}
	//se inicia con el primer estado
	po := state.PurchaseOrder{
		Client: "Alexander Andrade",
		State:  initialized,
	}
	for po.State.Run(&po) { //se ejecutara hasta que el metodo run de falso
	}

}
