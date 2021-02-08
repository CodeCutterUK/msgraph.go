// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// Room undocumented
type Room struct {
	// Place is the base model of Room
	Place
	// AudioDeviceName undocumented
	AudioDeviceName *string `json:"audioDeviceName,omitempty"`
	// BookingType undocumented
	BookingType *BookingType `json:"bookingType,omitempty"`
	// Building undocumented
	Building *string `json:"building,omitempty"`
	// Capacity undocumented
	Capacity *int `json:"capacity,omitempty"`
	// DisplayDeviceName undocumented
	DisplayDeviceName *string `json:"displayDeviceName,omitempty"`
	// EmailAddress undocumented
	EmailAddress *string `json:"emailAddress,omitempty"`
	// FloorLabel undocumented
	FloorLabel *string `json:"floorLabel,omitempty"`
	// FloorNumber undocumented
	FloorNumber *int `json:"floorNumber,omitempty"`
	// IsWheelChairAccessible undocumented
	IsWheelChairAccessible *bool `json:"isWheelChairAccessible,omitempty"`
	// Label undocumented
	Label *string `json:"label,omitempty"`
	// Nickname undocumented
	Nickname *string `json:"nickname,omitempty"`
	// Tags undocumented
	Tags []string `json:"tags,omitempty"`
	// VideoDeviceName undocumented
	VideoDeviceName *string `json:"videoDeviceName,omitempty"`
}

// RoomList undocumented
type RoomList struct {
	// Place is the base model of RoomList
	Place
	// EmailAddress undocumented
	EmailAddress *string `json:"emailAddress,omitempty"`
	// Rooms undocumented
	Rooms []Room `json:"rooms,omitempty"`
}