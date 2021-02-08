// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "time"

// Chat undocumented
type Chat struct {
	// Entity is the base model of Chat
	Entity
}

// ChatInfo undocumented
type ChatInfo struct {
	// Object is the base model of ChatInfo
	Object
	// MessageID undocumented
	MessageID *string `json:"messageId,omitempty"`
	// ReplyChainMessageID undocumented
	ReplyChainMessageID *string `json:"replyChainMessageId,omitempty"`
	// ThreadID undocumented
	ThreadID *string `json:"threadId,omitempty"`
}

// ChatMessage undocumented
type ChatMessage struct {
	// Entity is the base model of ChatMessage
	Entity
	// Attachments undocumented
	Attachments []ChatMessageAttachment `json:"attachments,omitempty"`
	// Body undocumented
	Body *ItemBody `json:"body,omitempty"`
	// CreatedDateTime undocumented
	CreatedDateTime *time.Time `json:"createdDateTime,omitempty"`
	// DeletedDateTime undocumented
	DeletedDateTime *time.Time `json:"deletedDateTime,omitempty"`
	// Etag undocumented
	Etag *string `json:"etag,omitempty"`
	// From undocumented
	From *IdentitySet `json:"from,omitempty"`
	// Importance undocumented
	Importance *ChatMessageImportance `json:"importance,omitempty"`
	// LastEditedDateTime undocumented
	LastEditedDateTime *time.Time `json:"lastEditedDateTime,omitempty"`
	// LastModifiedDateTime undocumented
	LastModifiedDateTime *time.Time `json:"lastModifiedDateTime,omitempty"`
	// Locale undocumented
	Locale *string `json:"locale,omitempty"`
	// Mentions undocumented
	Mentions []ChatMessageMention `json:"mentions,omitempty"`
	// MessageType undocumented
	MessageType *ChatMessageType `json:"messageType,omitempty"`
	// PolicyViolation undocumented
	PolicyViolation *ChatMessagePolicyViolation `json:"policyViolation,omitempty"`
	// Reactions undocumented
	Reactions []ChatMessageReaction `json:"reactions,omitempty"`
	// ReplyToID undocumented
	ReplyToID *string `json:"replyToId,omitempty"`
	// Subject undocumented
	Subject *string `json:"subject,omitempty"`
	// Summary undocumented
	Summary *string `json:"summary,omitempty"`
	// WebURL undocumented
	WebURL *string `json:"webUrl,omitempty"`
	// HostedContents undocumented
	HostedContents []ChatMessageHostedContent `json:"hostedContents,omitempty"`
	// Replies undocumented
	Replies []ChatMessage `json:"replies,omitempty"`
}

// ChatMessageAttachment undocumented
type ChatMessageAttachment struct {
	// Object is the base model of ChatMessageAttachment
	Object
	// Content undocumented
	Content *string `json:"content,omitempty"`
	// ContentType undocumented
	ContentType *string `json:"contentType,omitempty"`
	// ContentURL undocumented
	ContentURL *string `json:"contentUrl,omitempty"`
	// ID undocumented
	ID *string `json:"id,omitempty"`
	// Name undocumented
	Name *string `json:"name,omitempty"`
	// ThumbnailURL undocumented
	ThumbnailURL *string `json:"thumbnailUrl,omitempty"`
}

// ChatMessageHostedContent undocumented
type ChatMessageHostedContent struct {
	// Entity is the base model of ChatMessageHostedContent
	Entity
}

// ChatMessageMention undocumented
type ChatMessageMention struct {
	// Object is the base model of ChatMessageMention
	Object
	// ID undocumented
	ID *int `json:"id,omitempty"`
	// Mentioned undocumented
	Mentioned *IdentitySet `json:"mentioned,omitempty"`
	// MentionText undocumented
	MentionText *string `json:"mentionText,omitempty"`
}

// ChatMessagePolicyViolation undocumented
type ChatMessagePolicyViolation struct {
	// Object is the base model of ChatMessagePolicyViolation
	Object
	// DlpAction undocumented
	DlpAction *ChatMessagePolicyViolationDlpActionTypes `json:"dlpAction,omitempty"`
	// JustificationText undocumented
	JustificationText *string `json:"justificationText,omitempty"`
	// PolicyTip undocumented
	PolicyTip *ChatMessagePolicyViolationPolicyTip `json:"policyTip,omitempty"`
	// UserAction undocumented
	UserAction *ChatMessagePolicyViolationUserActionTypes `json:"userAction,omitempty"`
	// VerdictDetails undocumented
	VerdictDetails *ChatMessagePolicyViolationVerdictDetailsTypes `json:"verdictDetails,omitempty"`
}

// ChatMessagePolicyViolationPolicyTip undocumented
type ChatMessagePolicyViolationPolicyTip struct {
	// Object is the base model of ChatMessagePolicyViolationPolicyTip
	Object
	// ComplianceURL undocumented
	ComplianceURL *string `json:"complianceUrl,omitempty"`
	// GeneralText undocumented
	GeneralText *string `json:"generalText,omitempty"`
	// MatchedConditionDescriptions undocumented
	MatchedConditionDescriptions []string `json:"matchedConditionDescriptions,omitempty"`
}

// ChatMessageReaction undocumented
type ChatMessageReaction struct {
	// Object is the base model of ChatMessageReaction
	Object
	// CreatedDateTime undocumented
	CreatedDateTime *time.Time `json:"createdDateTime,omitempty"`
	// ReactionType undocumented
	ReactionType *string `json:"reactionType,omitempty"`
	// User undocumented
	User *IdentitySet `json:"user,omitempty"`
}
