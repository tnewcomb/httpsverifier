package servicetypes

//Used for response to a fingerprint request
type FingerprintResponse struct {
	Results []DomainResult
}

//Used for fingerprint requests
type FingerprintRequest struct {
	Domains []string
}

type DomainResult struct {
	Domain      string
	Fingerprint string
	Found       bool
}

type Page struct {
	Title   string
	Body    []byte
	Domains []string
	Results FingerprintResponse
}
