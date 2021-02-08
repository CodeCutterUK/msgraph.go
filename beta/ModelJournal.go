// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "time"

// Journal undocumented
type Journal struct {
	// Entity is the base model of Journal
	Entity
	// BalancingAccountID undocumented
	BalancingAccountID *UUID `json:"balancingAccountId,omitempty"`
	// BalancingAccountNumber undocumented
	BalancingAccountNumber *string `json:"balancingAccountNumber,omitempty"`
	// Code undocumented
	Code *string `json:"code,omitempty"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// LastModifiedDateTime undocumented
	LastModifiedDateTime *time.Time `json:"lastModifiedDateTime,omitempty"`
	// Account undocumented
	Account *Account `json:"account,omitempty"`
	// JournalLines undocumented
	JournalLines []JournalLine `json:"journalLines,omitempty"`
}

// JournalLine undocumented
type JournalLine struct {
	// Entity is the base model of JournalLine
	Entity
	// AccountID undocumented
	AccountID *UUID `json:"accountId,omitempty"`
	// AccountNumber undocumented
	AccountNumber *string `json:"accountNumber,omitempty"`
	// Amount undocumented
	Amount *int `json:"amount,omitempty"`
	// Comment undocumented
	Comment *string `json:"comment,omitempty"`
	// Description undocumented
	Description *string `json:"description,omitempty"`
	// DocumentNumber undocumented
	DocumentNumber *string `json:"documentNumber,omitempty"`
	// ExternalDocumentNumber undocumented
	ExternalDocumentNumber *string `json:"externalDocumentNumber,omitempty"`
	// JournalDisplayName undocumented
	JournalDisplayName *string `json:"journalDisplayName,omitempty"`
	// LastModifiedDateTime undocumented
	LastModifiedDateTime *time.Time `json:"lastModifiedDateTime,omitempty"`
	// LineNumber undocumented
	LineNumber *int `json:"lineNumber,omitempty"`
	// PostingDate undocumented
	PostingDate *Date `json:"postingDate,omitempty"`
	// Account undocumented
	Account *Account `json:"account,omitempty"`
}
