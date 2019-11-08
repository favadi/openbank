// Code generated by "scopegen"; DO NOT EDIT.
package consent

type AuthScope string

const (
	Scope_Read AuthScope = "https://auth.bnk.to/consent.read"
	Scope_Write AuthScope = "https://auth.bnk.to/consent.write"
)

var AuthScopes = map[string][]AuthScope{
	"/consent.ConsentService/AnswerConsentChallenge": []AuthScope{Scope_Write},
	"/consent.ConsentService/CreateConsentEmail": []AuthScope{Scope_Write},
	"/consent.ConsentService/CreateConsentSMS": []AuthScope{Scope_Write},
	"/consent.ConsentService/GetConsents": []AuthScope{Scope_Read},
	"/consent.ConsentService/RevokeConsent": []AuthScope{Scope_Write},
}