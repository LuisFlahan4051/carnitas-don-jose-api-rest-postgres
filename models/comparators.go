package models

/*import "time"

//Return old values if new values are empty or are equals and return new values otherwise
func CompareString(newString string, oldString string) string {
	if newString != "" && newString != oldString {
		return newString
	}
	return oldString
}

//Return old values if new values are empty or are equals and return new values otherwise
func CompareStringPointer(newString *string, oldString *string) *string {
	if newString != nil && *newString != *oldString {
		return newString
	}
	return oldString
}

//Return old values if new values are empty or are equals and return new values otherwise
func CompareTime(newTime time.Time, oldTime time.Time) time.Time {
	if !newTime.IsZero() && newTime != oldTime {
		return newTime
	}
	return oldTime
}

//Return old values if new values are empty or are equals and return new values otherwise
func CompareBool(newBool bool, oldBool bool) bool {
	if newBool != oldBool {
		return newBool
	}
	return oldBool
}

//Return old values if new values are empty or are equals and return new values otherwise
func CompareBoolPointer(newBool *bool, oldBool *bool) *bool {
	if newBool != nil && *newBool != *oldBool {
		return newBool
	}
	return oldBool
}

//Return old values if new values are empty or are equals and return new values otherwise
func CompareUint(newUint uint, oldUint uint) uint {
	if newUint != 0 && newUint != oldUint {
		return newUint
	}
	return oldUint
}

//Return old values if new values are empty or are equals and return new values otherwise
func CompareUintPointer(newUint *uint, oldUint *uint) *uint {
	if newUint != nil && *newUint != *oldUint {
		return newUint
	}
	return oldUint
}

//Return old values if new values are empty or NIL and return new values otherwise
func CompareTimePointer(newTime *time.Time, oldTime *time.Time) *time.Time {
	if newTime != nil && *newTime != *oldTime {
		return newTime
	}
	return oldTime
}

//Return old values if new values are empty or are equals and return new values otherwise
func CompareFloat(newFloat float64, oldFloat float64) float64 {
	if newFloat != 0 && newFloat != oldFloat {
		return newFloat
	}
	return oldFloat
}

//Return old values if new values are empty or are equals and return new values otherwise
func CompareFloatPointer(newFloat *float64, oldFloat *float64) *float64 {
	if newFloat != nil && *newFloat != *oldFloat {
		return newFloat
	}
	return oldFloat
}

//Update struct branch with new values
func (oldBranch *Branch) CompareBranch(newBranch Branch) {

	oldBranch.Id = CompareUint(newBranch.Id, oldBranch.Id)
	oldBranch.CreatedAt = CompareTimePointer(newBranch.CreatedAt, oldBranch.CreatedAt)
	oldBranch.UpdatedAt = CompareTimePointer(newBranch.UpdatedAt, oldBranch.UpdatedAt)
	oldBranch.DeletedAt = CompareTimePointer(newBranch.DeletedAt, oldBranch.DeletedAt)
	oldBranch.Name = CompareString(newBranch.Name, oldBranch.Name)
	oldBranch.Address = CompareString(newBranch.Address, oldBranch.Address)
}
*/
