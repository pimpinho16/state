package state

import (
	"fmt"
	"os"
	"strings"
)

type PendingApprove struct{}

func (pa PendingApprove) Run(po *PurchaseOrder) bool {
	approved := ""

	fmt.Println("Aceptas la orden? S/N")
	_, err := fmt.Fscanf(os.Stdin, "%s\n", &approved)
	if err != nil {
		fmt.Println("No se pudo procesasr la opcion")
		return false
	}
	if !strings.EqualFold(approved, "s") {
		po.State = Rejected{}
		return true
	}
	po.State = Accepted{}

	return true
}
