package user

import (
	"context"
	"fmt"
	"sync"

	"github.com/google/uuid"
)

type GrpcServerUser struct {
	UnimplementedUserManagerServer
	Users    map[string]*User
	NextUuid string
	mu       sync.Mutex
}

func (s *GrpcServerUser) CreateUser(ctx context.Context, req *UserCreateModelRequest) (*UserModelResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.NextUuid = uuid.New().String()

	user := &User{
		Uuid:       s.NextUuid,
		Identifier: req.GetIdentifier(),
		Secret:     req.GetSecret(),
		Name:       req.GetName(),
	}
	s.Users[user.Uuid] = user

	return &UserModelResponse{
		Identifier: user.GetIdentifier(),
		Name:       user.GetName(),
	}, nil
}

func (s *GrpcServerUser) ReadUser(ctx context.Context, req *ReadUserRequest) (*UserModelResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	findUuid := req.GetUuid()
	if s.Users[findUuid] == nil {
		return nil, fmt.Errorf("not found user uuid %s to read", findUuid)
	}

	foundUser := s.Users[findUuid]
	return &UserModelResponse{
		Identifier: foundUser.GetIdentifier(),
		Name:       foundUser.GetName(),
	}, nil
}

func (s *GrpcServerUser) UpdateUser(ctx context.Context, req *UserUpdateModelRequest) (*UserModelResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	findUuid := req.GetUuid()
	if s.Users[findUuid] == nil {
		return nil, fmt.Errorf("not found user uuid %s to update", findUuid)
	}

	updateUser := s.Users[findUuid]
	updateUser.Identifier = req.GetIdentifier()
	updateUser.Secret = req.GetSecret()
	updateUser.Name = req.GetName()

	return &UserModelResponse{
		Identifier: updateUser.GetIdentifier(),
		Name:       updateUser.GetName(),
	}, nil
}

func (s *GrpcServerUser) DeleteUser(ctx context.Context, req *DeleteUserRequest) (*DeleteUserRequest, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	findUuid := req.GetUuid()
	if s.Users[findUuid] == nil {
		return nil, fmt.Errorf("not found user uuid %s to delete", findUuid)
	}

	delete(s.Users, findUuid)
	return &DeleteUserRequest{
		Uuid: findUuid,
	}, nil
}
