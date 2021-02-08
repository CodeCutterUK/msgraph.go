// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// TranslationLanguageOverride undocumented
type TranslationLanguageOverride struct {
	// Object is the base model of TranslationLanguageOverride
	Object
	// LanguageTag undocumented
	LanguageTag *string `json:"languageTag,omitempty"`
	// TranslationBehavior undocumented
	TranslationBehavior *TranslationBehavior `json:"translationBehavior,omitempty"`
}

// TranslationPreferences undocumented
type TranslationPreferences struct {
	// Object is the base model of TranslationPreferences
	Object
	// LanguageOverrides undocumented
	LanguageOverrides []TranslationLanguageOverride `json:"languageOverrides,omitempty"`
	// TranslationBehavior undocumented
	TranslationBehavior *TranslationBehavior `json:"translationBehavior,omitempty"`
	// UntranslatedLanguages undocumented
	UntranslatedLanguages []string `json:"untranslatedLanguages,omitempty"`
}