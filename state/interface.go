package state

// State interface que recibe un puntero a una orden de compra
// return bool se devuelve si es el estado final o si seguiran cambiando los estados, si es estado final: false, sino lo es: true
type State interface {
	Run(po *PurchaseOrder) bool
}
