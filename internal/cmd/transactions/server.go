package transactions

import (
	"fmt"

	"github.com/fair-n-square-co/transactions/internal/config"
	"github.com/fair-n-square-co/transactions/internal/controller"
)

type ServerGroup struct {
	UserServer  *UserServer
	GroupServer *GroupServer
}

func NewServerGroup() (*ServerGroup, error) {
	config := config.NewConfig()

	ctrl, err := controller.NewController(config.Database)
	if err != nil {
		return nil, fmt.Errorf("failed to create controller: %v", err)
	}

	userServer, err := NewUserServer(ctrl.UserController())
	if err != nil {
		return nil, fmt.Errorf("failed to create user server: %v", err)
	}

	groupServer, err := NewGroupServer(ctrl.GroupController())
	if err != nil {
		return nil, fmt.Errorf("failed to create group server: %v", err)
	}

	// Create server group
	serverGroup := &ServerGroup{
		UserServer:  userServer,
		GroupServer: groupServer,
	}

	// Use the server group
	fmt.Println(serverGroup)

	return serverGroup, nil
}
