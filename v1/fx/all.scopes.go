// Code generated by "scopegen"; DO NOT EDIT.
package fx

type AuthScope string

const (
	Scope_Read AuthScope = "https://auth.bnk.to/fx.read"
	Scope_Write AuthScope = "https://auth.bnk.to/fx.write"
)

var AuthScopes = map[string][]AuthScope{
	"/fx.FXService/CreateFX": []AuthScope{Scope_Write},
	"/fx.FXService/GetCurrentFXRate": []AuthScope{Scope_Read},
	"/fx.FXService/UpdateFX": []AuthScope{Scope_Write},
}