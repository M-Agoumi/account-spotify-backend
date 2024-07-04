package user

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"gorm.io/gorm"
	"regexp"
	"slices"
	"time"
)

type User struct {
	gorm.Model
	ID           uint    `gorm:"uniqueIndex"` // Standard field for the primary key
	Username     *string `gorm:"uniqueIndex"` // please save
	Name         string
	Email        *string    `gorm:"uniqueIndex"` // A pointer to a string, allowing for null values
	Password     string     `json:"-"`           // Exclude from JSON response
	Birthday     *time.Time // A pointer to time.Time, can be null
	Gender       string     // M/F/N
	Terms        bool
	MemberNumber sql.NullString // Uses sql.NullString to handle nullable strings
	ActivatedAt  sql.NullTime   // Uses sql.NullTime for nullable time fields
	CreatedAt    time.Time      // Automatically managed by GORM for creation time
	UpdatedAt    time.Time      // Automatically managed by GORM for update time
}

// UnmarshalJSON Custom unmarshalling logic for the User struct
func (u *User) UnmarshalJSON(data []byte) error {
	type Alias User // Create an alias to avoid recursion
	aux := &struct {
		Password string `json:"password"`
		Birthday string `json:"birthday"`
		*Alias
	}{
		Alias: (*Alias)(u),
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		return fmt.Errorf("invalid request payload")
	}

	// Validate email
	if aux.Email != nil && !IsValidEmail(*aux.Email) {
		return fmt.Errorf("invalid email format")
	}

	if aux.Gender != "" && !isValidGender(aux.Gender) {
		return fmt.Errorf("invalid gender format")
	}

	// parse the password
	if aux.Password != "" {
		u.Password = aux.Password
	}
	// Parse the custom birthday format
	if aux.Birthday != "" {
		birthday, err := time.Parse("01/02/2006", aux.Birthday)
		if err != nil {
			return fmt.Errorf("invalid birthday format")
		}
		u.Birthday = &birthday
	}

	return nil
}

func isValidGender(gender string) bool {
	gendersList := []string{"m", "f", "n"}

	return slices.Contains(gendersList, gender)
}

// IsValidEmail Email validation function
func IsValidEmail(email string) bool {
	regex := `^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`
	re := regexp.MustCompile(regex)
	return re.MatchString(email)
}
