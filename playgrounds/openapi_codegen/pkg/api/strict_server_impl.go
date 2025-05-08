package api

import (
	"context"

	"github.com/aazw/vscode-go/playgrounds/openapi_codegen/pkg/api/openapi"
)

type StrictServerInterfaceImpl struct {
}

func NewStrictServerInterfaceImpl() openapi.StrictServerInterface {
	return &StrictServerInterfaceImpl{}
}

func (p *StrictServerInterfaceImpl) ListUsers(ctx context.Context, request openapi.ListUsersRequestObject) (openapi.ListUsersResponseObject, error) {

	return openapi.ListUsers200JSONResponse{}, nil
}

func (p *StrictServerInterfaceImpl) CreateUser(ctx context.Context, request openapi.CreateUserRequestObject) (openapi.CreateUserResponseObject, error) {

	return openapi.CreateUser201JSONResponse{}, nil
}

func (p *StrictServerInterfaceImpl) GetUserById(ctx context.Context, request openapi.GetUserByIdRequestObject) (openapi.GetUserByIdResponseObject, error) {

	return openapi.GetUserById200JSONResponse{}, nil
}

func (p *StrictServerInterfaceImpl) UpdateUserById(ctx context.Context, request openapi.UpdateUserByIdRequestObject) (openapi.UpdateUserByIdResponseObject, error) {

	return openapi.UpdateUserById200JSONResponse{}, nil
}

func (p *StrictServerInterfaceImpl) DeleteUserById(ctx context.Context, request openapi.DeleteUserByIdRequestObject) (openapi.DeleteUserByIdResponseObject, error) {

	return openapi.DeleteUserById204Response{}, nil
}
