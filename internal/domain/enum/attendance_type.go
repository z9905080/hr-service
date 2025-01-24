package enum

// AttendanceType is a struct that represent attendance type enum
type AttendanceType string

const (
	// AttendanceTypeIn is a constant that represent attendance type in
	AttendanceTypeIn AttendanceType = "in"
	// AttendanceTypeOut is a constant that represent attendance type out
	AttendanceTypeOut AttendanceType = "out"
)

func (a AttendanceType) String() string {
	return string(a)
}

func (a AttendanceType) IsValid() bool {
	switch a {
	case AttendanceTypeIn, AttendanceTypeOut:
		return true
	}
	return false
}

// IsIn is a function to check if attendance type is in
// NIT: this function may not be needed
func (a AttendanceType) IsIn() bool {
	return a == AttendanceTypeIn
}

// IsOut is a function to check if attendance type is out
// NIT: this function may not be needed
func (a AttendanceType) IsOut() bool {
	return a == AttendanceTypeOut
}
