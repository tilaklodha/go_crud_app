package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreatesNewUser(t *testing.T) {
	expectedUser := User{
		FirstName: "abc",
		LastName:  "xyz",
		City:      "city",
	}
	observedUser := NewUser("abc", "xyz", "city")
	assert.Equal(t, expectedUser.FirstName, observedUser.FirstName)
	assert.Equal(t, expectedUser.LastName, observedUser.LastName)
	assert.Equal(t, expectedUser.City, observedUser.City)
}
