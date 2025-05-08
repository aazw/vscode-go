package api

import (
	"github.com/gin-gonic/gin"
	openapi_types "github.com/oapi-codegen/runtime/types"

	"github.com/aazw/vscode-go/playgrounds/openapi_codegen/pkg/api/openapi"
)

type ServerInterfaceImpl struct {
}

func NewGinServerHandler() openapi.ServerInterface {
	return &ServerInterfaceImpl{}
}

// ServerInterfaceの実装
func (p *ServerInterfaceImpl) ListUsers(c *gin.Context) {

}

func (p *ServerInterfaceImpl) CreateUser(c *gin.Context) {

}

func (p *ServerInterfaceImpl) DeleteUserById(c *gin.Context, userId openapi_types.UUID) {

}

func (p *ServerInterfaceImpl) GetUserById(c *gin.Context, userId openapi_types.UUID) {

}

func (p *ServerInterfaceImpl) UpdateUserById(c *gin.Context, userId openapi_types.UUID) {

}
