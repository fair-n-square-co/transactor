package controller

import (
	"testing"

	dbmocks "github.com/fair-n-square-co/transactions/internal/db/mocks"
	tassert "github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestNewController(t *testing.T) {
	assert := tassert.New(t)
	ctrl := gomock.NewController(t)
	mockDBClient := dbmocks.NewMockClient(ctrl)
	newController, err := NewController(mockDBClient)

	if err != nil {
		t.Errorf("Failed to create controller: %v", err)
	}

	assert.IsType(&controller{}, newController, "Expected controller to be of type *controller")
}
