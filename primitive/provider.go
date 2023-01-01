package primitive

// Provider specifies a shipping provider
type Provider uint8

const (
	ProviderUnspecified Provider = iota
	ProviderJNE
	ProviderSiCepat
	ProviderJNT
	ProviderAnterAja
)

func (p Provider) String() string {
	switch p {
	case ProviderJNE:
		return "JNE"
	case ProviderSiCepat:
		return "SICEPAT"
	case ProviderJNT:
		return "JNT"
	case ProviderAnterAja:
		return "ANTERAJA"
	case ProviderUnspecified:
		fallthrough
	default:
		return "UNSPECIFIED"
	}
}
