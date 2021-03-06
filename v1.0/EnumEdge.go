// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// EdgeCookiePolicy undocumented
type EdgeCookiePolicy string

const (
	// EdgeCookiePolicyVUserDefined undocumented
	EdgeCookiePolicyVUserDefined EdgeCookiePolicy = "userDefined"
	// EdgeCookiePolicyVAllow undocumented
	EdgeCookiePolicyVAllow EdgeCookiePolicy = "allow"
	// EdgeCookiePolicyVBlockThirdParty undocumented
	EdgeCookiePolicyVBlockThirdParty EdgeCookiePolicy = "blockThirdParty"
	// EdgeCookiePolicyVBlockAll undocumented
	EdgeCookiePolicyVBlockAll EdgeCookiePolicy = "blockAll"
)

var (
	// EdgeCookiePolicyPUserDefined is a pointer to EdgeCookiePolicyVUserDefined
	EdgeCookiePolicyPUserDefined = &_EdgeCookiePolicyPUserDefined
	// EdgeCookiePolicyPAllow is a pointer to EdgeCookiePolicyVAllow
	EdgeCookiePolicyPAllow = &_EdgeCookiePolicyPAllow
	// EdgeCookiePolicyPBlockThirdParty is a pointer to EdgeCookiePolicyVBlockThirdParty
	EdgeCookiePolicyPBlockThirdParty = &_EdgeCookiePolicyPBlockThirdParty
	// EdgeCookiePolicyPBlockAll is a pointer to EdgeCookiePolicyVBlockAll
	EdgeCookiePolicyPBlockAll = &_EdgeCookiePolicyPBlockAll
)

var (
	_EdgeCookiePolicyPUserDefined     = EdgeCookiePolicyVUserDefined
	_EdgeCookiePolicyPAllow           = EdgeCookiePolicyVAllow
	_EdgeCookiePolicyPBlockThirdParty = EdgeCookiePolicyVBlockThirdParty
	_EdgeCookiePolicyPBlockAll        = EdgeCookiePolicyVBlockAll
)

// EdgeSearchEngineType undocumented
type EdgeSearchEngineType string

const (
	// EdgeSearchEngineTypeVDefault undocumented
	EdgeSearchEngineTypeVDefault EdgeSearchEngineType = "default"
	// EdgeSearchEngineTypeVBing undocumented
	EdgeSearchEngineTypeVBing EdgeSearchEngineType = "bing"
)

var (
	// EdgeSearchEngineTypePDefault is a pointer to EdgeSearchEngineTypeVDefault
	EdgeSearchEngineTypePDefault = &_EdgeSearchEngineTypePDefault
	// EdgeSearchEngineTypePBing is a pointer to EdgeSearchEngineTypeVBing
	EdgeSearchEngineTypePBing = &_EdgeSearchEngineTypePBing
)

var (
	_EdgeSearchEngineTypePDefault = EdgeSearchEngineTypeVDefault
	_EdgeSearchEngineTypePBing    = EdgeSearchEngineTypeVBing
)
