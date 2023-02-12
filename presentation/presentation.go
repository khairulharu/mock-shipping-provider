package presentation

import "mock-shipping-provider/primitive"

// providerMap converts the given string
// into the correct enum value
var providerMap = map[string]primitive.Provider{
	"JNE":      primitive.ProviderJNE,
	"SICEPAT":  primitive.ProviderSiCepat,
	"JNT":      primitive.ProviderJNT,
	"ANTERAJA": primitive.ProviderAnterAja,
}
