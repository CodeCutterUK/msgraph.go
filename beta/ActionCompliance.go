// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// Ediscovery is navigation property
func (b *ComplianceRequestBuilder) Ediscovery() *EdiscoveryrootRequestBuilder {
	bb := &EdiscoveryrootRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/ediscovery"
	return bb
}
