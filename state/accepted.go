package state

import "fmt"

type Accepted struct{}

func (a Accepted) Run(po *PurchaseOrder) bool {
	fmt.Printf("%s,Gracias por tu compra, te enviaremos la factura", po.Client)
	return false
}
