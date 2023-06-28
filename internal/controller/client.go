package controller

import (
	"QuickAuth/internal/endpoint/resp"
	"QuickAuth/pkg/model"
	"github.com/gin-gonic/gin"
)

// @Summary	clients info
// @Schemes
// @Description	list clients
// @Tags		client
// @Success		200
// @Router		/api/quick/clients [get]
func (o Controller) listClient(c *gin.Context) {
	clients, err := o.svc.ListClients()
	if err != nil {
		resp.ErrorSelect(c, err, "list clients err")
		return
	}
	resp.SuccessArray(c, len(clients), clients)
}

// @Summary	clients info
// @Schemes
// @Description	list clients
// @Tags		client
// @Param		clientId	path	string	true	"client id"
// @Success		200
// @Router		/api/quick/clients/{clientId} [get]
func (o Controller) getClient(c *gin.Context) {
	var in model.Client
	if err := o.SetCtx(c).BindUri(&in).Error; err != nil {
		resp.ErrorRequest(c, err, "init client req err")
		return
	}
	client, err := o.svc.GetClient(in.ID)
	if err != nil {
		resp.ErrorUnknown(c, err, "no such client")
		return
	}
	resp.SuccessWithData(c, client)
}

// @Summary	create client
// @Schemes
// @Description	create client
// @Tags		client
// @Success		200
// @Param		bd		body	model.Client	true	"body"
// @Router		/api/quick/clients [post]
func (o Controller) createClient(c *gin.Context) {
	var in model.Client
	if err := o.SetCtx(c).BindUri(&in).Error; err != nil {
		resp.ErrorRequest(c, err, "init client req err")
		return
	}
	client, err := o.svc.CreateClient(in)
	if err != nil {
		resp.ErrorUnknown(c, err, "create client err")
		return
	}
	resp.SuccessWithData(c, client)
}

// @Summary	modify client
// @Schemes
// @Description	modify client
// @Tags		client
// @Param		clientId	path	string	true	"client id"
// @Param		bd		body	model.Client	true	"body"
// @Success		200
// @Router		/api/quick/clients/{clientId} [put]
func (o Controller) modifyClient(c *gin.Context) {
	var in model.Client
	if err := o.SetCtx(c).BindUriAndJson(&in).Error; err != nil {
		resp.ErrorRequest(c, err, "init client req err")
		return
	}
	if err := o.svc.ModifyClient(in); err != nil {
		resp.ErrorUnknown(c, err, "modify client err")
		return
	}
	resp.Success(c)
}

// @Summary	delete client
// @Schemes
// @Description	delete client
// @Tags		client
// @Param		clientId	path	string	true	"client id"
// @Success		200
// @Router		/api/quick/clients/{clientId} [delete]
func (o Controller) deleteClient(c *gin.Context) {
	if err := o.svc.DeleteClient(c.Param("clientId")); err != nil {
		resp.ErrorUnknown(c, err, "delete client err")
		return
	}
	resp.Success(c)
}
