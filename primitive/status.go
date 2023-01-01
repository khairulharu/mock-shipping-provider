package primitive

type Status uint8

const (
	StatusUnspecified Status = iota
	// StatusOrderPlaced states order has been placed/manifested on courier partner
	StatusOrderPlaced
	// StatusPickupPending states pending pickup
	StatusPickupPending
	// StatusPickupFailed states failed pickup
	StatusPickupFailed
	// StatusPickedUp states pickup has been done
	StatusPickedUp
	// StatusInTransit states in transit
	StatusInTransit
	// StatusOutForDelivery states shipment is out for delivery
	StatusOutForDelivery
	// StatusDelivered states shipment delivered
	StatusDelivered
)

func (s Status) String() string {
	switch s {
	case StatusOrderPlaced:
		return "ORDER_PLACED"
	case StatusPickupPending:
		return "PENDING"
	case StatusPickupFailed:
		return "PICKUP_FAILED"
	case StatusPickedUp:
		return "PICKED_UP"
	case StatusInTransit:
		return "IN_TRANSIT"
	case StatusOutForDelivery:
		return "OUT_FOR_DELIVERY"
	case StatusDelivered:
		return "DELIVERED"
	case StatusUnspecified:
		fallthrough
	default:
		return "UNSPECIFIED"
	}
}
