package state

import (
	"fmt"
	"os"
	"strings"
)

type Rejected struct{}

func (r Rejected) Run(po *PurchaseOrder) bool {
	confirm := ""
	fmt.Println("Confirmas que rechazas la orden? S/N")
	_, err := fmt.Fscanf(os.Stdin, "%s\n", &confirm)
	if err != nil {
		fmt.Println("No se pudo procesar la opcion")
		return false
	}
	if !strings.EqualFold(confirm, "s") {
		po.State = PendingApprove{}
		return true
	}
	fmt.Printf("%s Lamentamos que hayas cancelado tu compra", po.Client)
	return false
}
