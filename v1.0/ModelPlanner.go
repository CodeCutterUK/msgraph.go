// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "time"

// Planner undocumented
type Planner struct {
	// Entity is the base model of Planner
	Entity
	// Buckets undocumented
	Buckets []PlannerBucket `json:"buckets,omitempty"`
	// Plans undocumented
	Plans []PlannerPlan `json:"plans,omitempty"`
	// Tasks undocumented
	Tasks []PlannerTask `json:"tasks,omitempty"`
}

// PlannerAppliedCategories undocumented
type PlannerAppliedCategories struct {
	// Object is the base model of PlannerAppliedCategories
	Object
}

// PlannerAssignedToTaskBoardTaskFormat undocumented
type PlannerAssignedToTaskBoardTaskFormat struct {
	// Entity is the base model of PlannerAssignedToTaskBoardTaskFormat
	Entity
	// OrderHintsByAssignee undocumented
	OrderHintsByAssignee *PlannerOrderHintsByAssignee `json:"orderHintsByAssignee,omitempty"`
	// UnassignedOrderHint undocumented
	UnassignedOrderHint *string `json:"unassignedOrderHint,omitempty"`
}

// PlannerAssignment undocumented
type PlannerAssignment struct {
	// Object is the base model of PlannerAssignment
	Object
	// AssignedBy undocumented
	AssignedBy *IdentitySet `json:"assignedBy,omitempty"`
	// AssignedDateTime undocumented
	AssignedDateTime *time.Time `json:"assignedDateTime,omitempty"`
	// OrderHint undocumented
	OrderHint *string `json:"orderHint,omitempty"`
}

// PlannerAssignments undocumented
type PlannerAssignments struct {
	// Object is the base model of PlannerAssignments
	Object
}

// PlannerBucket undocumented
type PlannerBucket struct {
	// Entity is the base model of PlannerBucket
	Entity
	// Name undocumented
	Name *string `json:"name,omitempty"`
	// OrderHint undocumented
	OrderHint *string `json:"orderHint,omitempty"`
	// PlanID undocumented
	PlanID *string `json:"planId,omitempty"`
	// Tasks undocumented
	Tasks []PlannerTask `json:"tasks,omitempty"`
}

// PlannerBucketTaskBoardTaskFormat undocumented
type PlannerBucketTaskBoardTaskFormat struct {
	// Entity is the base model of PlannerBucketTaskBoardTaskFormat
	Entity
	// OrderHint undocumented
	OrderHint *string `json:"orderHint,omitempty"`
}

// PlannerCategoryDescriptions undocumented
type PlannerCategoryDescriptions struct {
	// Object is the base model of PlannerCategoryDescriptions
	Object
	// Category1 undocumented
	Category1 *string `json:"category1,omitempty"`
	// Category2 undocumented
	Category2 *string `json:"category2,omitempty"`
	// Category3 undocumented
	Category3 *string `json:"category3,omitempty"`
	// Category4 undocumented
	Category4 *string `json:"category4,omitempty"`
	// Category5 undocumented
	Category5 *string `json:"category5,omitempty"`
	// Category6 undocumented
	Category6 *string `json:"category6,omitempty"`
}

// PlannerChecklistItem undocumented
type PlannerChecklistItem struct {
	// Object is the base model of PlannerChecklistItem
	Object
	// IsChecked undocumented
	IsChecked *bool `json:"isChecked,omitempty"`
	// LastModifiedBy undocumented
	LastModifiedBy *IdentitySet `json:"lastModifiedBy,omitempty"`
	// LastModifiedDateTime undocumented
	LastModifiedDateTime *time.Time `json:"lastModifiedDateTime,omitempty"`
	// OrderHint undocumented
	OrderHint *string `json:"orderHint,omitempty"`
	// Title undocumented
	Title *string `json:"title,omitempty"`
}

// PlannerChecklistItems undocumented
type PlannerChecklistItems struct {
	// Object is the base model of PlannerChecklistItems
	Object
}

// PlannerExternalReference undocumented
type PlannerExternalReference struct {
	// Object is the base model of PlannerExternalReference
	Object
	// Alias undocumented
	Alias *string `json:"alias,omitempty"`
	// LastModifiedBy undocumented
	LastModifiedBy *IdentitySet `json:"lastModifiedBy,omitempty"`
	// LastModifiedDateTime undocumented
	LastModifiedDateTime *time.Time `json:"lastModifiedDateTime,omitempty"`
	// PreviewPriority undocumented
	PreviewPriority *string `json:"previewPriority,omitempty"`
	// Type undocumented
	Type *string `json:"type,omitempty"`
}

// PlannerExternalReferences undocumented
type PlannerExternalReferences struct {
	// Object is the base model of PlannerExternalReferences
	Object
}

// PlannerGroup undocumented
type PlannerGroup struct {
	// Entity is the base model of PlannerGroup
	Entity
	// Plans undocumented
	Plans []PlannerPlan `json:"plans,omitempty"`
}

// PlannerOrderHintsByAssignee undocumented
type PlannerOrderHintsByAssignee struct {
	// Object is the base model of PlannerOrderHintsByAssignee
	Object
}

// PlannerPlan undocumented
type PlannerPlan struct {
	// Entity is the base model of PlannerPlan
	Entity
	// CreatedBy undocumented
	CreatedBy *IdentitySet `json:"createdBy,omitempty"`
	// CreatedDateTime undocumented
	CreatedDateTime *time.Time `json:"createdDateTime,omitempty"`
	// Owner undocumented
	Owner *string `json:"owner,omitempty"`
	// Title undocumented
	Title *string `json:"title,omitempty"`
	// Buckets undocumented
	Buckets []PlannerBucket `json:"buckets,omitempty"`
	// Details undocumented
	Details *PlannerPlanDetails `json:"details,omitempty"`
	// Tasks undocumented
	Tasks []PlannerTask `json:"tasks,omitempty"`
}

// PlannerPlanDetails undocumented
type PlannerPlanDetails struct {
	// Entity is the base model of PlannerPlanDetails
	Entity
	// CategoryDescriptions undocumented
	CategoryDescriptions *PlannerCategoryDescriptions `json:"categoryDescriptions,omitempty"`
	// SharedWith undocumented
	SharedWith *PlannerUserIDs `json:"sharedWith,omitempty"`
}

// PlannerProgressTaskBoardTaskFormat undocumented
type PlannerProgressTaskBoardTaskFormat struct {
	// Entity is the base model of PlannerProgressTaskBoardTaskFormat
	Entity
	// OrderHint undocumented
	OrderHint *string `json:"orderHint,omitempty"`
}

// PlannerTask undocumented
type PlannerTask struct {
	// Entity is the base model of PlannerTask
	Entity
	// ActiveChecklistItemCount undocumented
	ActiveChecklistItemCount *int `json:"activeChecklistItemCount,omitempty"`
	// AppliedCategories undocumented
	AppliedCategories *PlannerAppliedCategories `json:"appliedCategories,omitempty"`
	// AssigneePriority undocumented
	AssigneePriority *string `json:"assigneePriority,omitempty"`
	// Assignments undocumented
	Assignments *PlannerAssignments `json:"assignments,omitempty"`
	// BucketID undocumented
	BucketID *string `json:"bucketId,omitempty"`
	// ChecklistItemCount undocumented
	ChecklistItemCount *int `json:"checklistItemCount,omitempty"`
	// CompletedBy undocumented
	CompletedBy *IdentitySet `json:"completedBy,omitempty"`
	// CompletedDateTime undocumented
	CompletedDateTime *time.Time `json:"completedDateTime,omitempty"`
	// ConversationThreadID undocumented
	ConversationThreadID *string `json:"conversationThreadId,omitempty"`
	// CreatedBy undocumented
	CreatedBy *IdentitySet `json:"createdBy,omitempty"`
	// CreatedDateTime undocumented
	CreatedDateTime *time.Time `json:"createdDateTime,omitempty"`
	// DueDateTime undocumented
	DueDateTime *time.Time `json:"dueDateTime,omitempty"`
	// HasDescription undocumented
	HasDescription *bool `json:"hasDescription,omitempty"`
	// OrderHint undocumented
	OrderHint *string `json:"orderHint,omitempty"`
	// PercentComplete undocumented
	PercentComplete *int `json:"percentComplete,omitempty"`
	// PlanID undocumented
	PlanID *string `json:"planId,omitempty"`
	// PreviewType undocumented
	PreviewType *PlannerPreviewType `json:"previewType,omitempty"`
	// ReferenceCount undocumented
	ReferenceCount *int `json:"referenceCount,omitempty"`
	// StartDateTime undocumented
	StartDateTime *time.Time `json:"startDateTime,omitempty"`
	// Title undocumented
	Title *string `json:"title,omitempty"`
	// AssignedToTaskBoardFormat undocumented
	AssignedToTaskBoardFormat *PlannerAssignedToTaskBoardTaskFormat `json:"assignedToTaskBoardFormat,omitempty"`
	// BucketTaskBoardFormat undocumented
	BucketTaskBoardFormat *PlannerBucketTaskBoardTaskFormat `json:"bucketTaskBoardFormat,omitempty"`
	// Details undocumented
	Details *PlannerTaskDetails `json:"details,omitempty"`
	// ProgressTaskBoardFormat undocumented
	ProgressTaskBoardFormat *PlannerProgressTaskBoardTaskFormat `json:"progressTaskBoardFormat,omitempty"`
}

// PlannerTaskDetails undocumented
type PlannerTaskDetails struct {
	// Entity is the base model of PlannerTaskDetails
	Entity
	// Checklist undocumented
	Checklist *PlannerChecklistItems `json:"checklist,omitempty"`
	// Description undocumented
	Description *string `json:"description,omitempty"`
	// PreviewType undocumented
	PreviewType *PlannerPreviewType `json:"previewType,omitempty"`
	// References undocumented
	References *PlannerExternalReferences `json:"references,omitempty"`
}

// PlannerUser undocumented
type PlannerUser struct {
	// Entity is the base model of PlannerUser
	Entity
	// Plans undocumented
	Plans []PlannerPlan `json:"plans,omitempty"`
	// Tasks undocumented
	Tasks []PlannerTask `json:"tasks,omitempty"`
}

// PlannerUserIDs undocumented
type PlannerUserIDs struct {
	// Object is the base model of PlannerUserIDs
	Object
}
