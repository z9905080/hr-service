package enum

// LeaveType is a struct that represent leave type enum
type LeaveType string

const (
	// LeaveTypeAnnual is a constant that represent leave type annual
	LeaveTypeAnnual LeaveType = "annual"
	// LeaveTypeSick is a constant that represent leave type sick
	LeaveTypeSick LeaveType = "sick"
	// LeaveTypeUnpaid is a constant that represent leave type unpaid
	LeaveTypeUnpaid LeaveType = "unpaid"
)

func (l LeaveType) String() string {
	return string(l)
}

func (l LeaveType) IsValid() bool {
	switch l {
	case LeaveTypeAnnual, LeaveTypeSick, LeaveTypeUnpaid:
		return true
	}
	return false
}

// IsAnnual is a function to check if leave type is annual
// NIT: this function may not be needed
func (l LeaveType) IsAnnual() bool {
	return l == LeaveTypeAnnual
}

// IsSick is a function to check if leave type is sick
// NIT: this function may not be needed
func (l LeaveType) IsSick() bool {
	return l == LeaveTypeSick
}

// IsUnpaid is a function to check if leave type is unpaid
// NIT: this function may not be needed
func (l LeaveType) IsUnpaid() bool {
	return l == LeaveTypeUnpaid
}
