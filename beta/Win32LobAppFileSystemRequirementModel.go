// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// Win32LobAppFileSystemRequirement undocumented
type Win32LobAppFileSystemRequirement struct {
	Win32LobAppRequirement
	// Path The file or folder path to detect Win32 Line of Business (LoB) app
	Path *string `json:"path,omitempty"`
	// FileOrFolderName The file or folder name to detect Win32 Line of Business (LoB) app
	FileOrFolderName *string `json:"fileOrFolderName,omitempty"`
	// Check32BitOn64System A value indicating whether this file or folder is for checking 32-bit app on 64-bit system
	Check32BitOn64System *bool `json:"check32BitOn64System,omitempty"`
	// DetectionType The file system detection type
	DetectionType *Win32LobAppFileSystemDetectionType `json:"detectionType,omitempty"`
}