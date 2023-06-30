package controller

import (
	"QuickAuth/internal/endpoint/request"
	"QuickAuth/internal/endpoint/resp"
	"github.com/gin-gonic/gin"
)

// @Summary	client secret info
// @Schemes
// @Description	list client secret
// @Tags		client
// @Param		clientId	path	string	true	"client id"
// @Success		200
// @Router		/api/quick/clients/{clientId}/secrets [get]
func (o Controller) listClientSecret(c *gin.Context) {
	clients, err := o.svc.ListClientSecrets(c.Param("clientId"))
	if err != nil {
		resp.ErrorSelect(c, err, "list client secret err")
		return
	}
	resp.SuccessArray(c, len(clients), clients)
}

// @Summary	create client secret
// @Schemes
// @Description	create client secret
// @Tags		client
// @Param		clientId	path	string					true	"client id"
// @Param		bd			body	request.ClientSecretReq	true	"body"
// @Success		200
// @Router		/api/quick/clients/{clientId}/secrets [post]
func (o Controller) createClientSecret(c *gin.Context) {
	var in request.ClientSecretReq
	if err := o.SetCtx(c).BindUriAndJson(&in).Error; err != nil {
		resp.ErrorRequest(c, err, "init client secret req err")
		return
	}

	client, err := o.svc.CreateClientSecret(in.ToModel())
	if err != nil {
		resp.ErrorUnknown(c, err, "create client secret err")
		return
	}
	resp.SuccessWithData(c, client)
}

// @Summary	delete client secret
// @Schemes
// @Description	delete client secret
// @Tags		client
// @Param		clientId	path	string	true	"client id"
// @Success		200
// @Router		/api/quick/clients/{clientId}/secrets/{secretId} [delete]
func (o Controller) deleteClientSecret(c *gin.Context) {
	clientId := c.Param("clientId")
	secretId := c.Param("secretId")
	if err := o.svc.DeleteClientSecret(clientId, secretId); err != nil {
		resp.ErrorUnknown(c, err, "delete client secret err")
		return
	}
	resp.Success(c)
}
