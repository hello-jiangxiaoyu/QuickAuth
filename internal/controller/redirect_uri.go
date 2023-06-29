package controller

import (
	"QuickAuth/internal/endpoint/resp"
	"QuickAuth/pkg/model"
	"github.com/gin-gonic/gin"
)

// @Summary	get redirect uri list
// @Schemes
// @Description	get redirect uri list
// @Tags		client
// @Param		clientId	path	string	true	"client id"
// @Success		200
// @Router		/api/quick/clients/{clientId}/redirect-uri [get]
func (o Controller) listRedirectUri(c *gin.Context) {
	clients, err := o.svc.ListRedirectUri(c.Param("clientId"))
	if err != nil {
		resp.ErrorSelect(c, err, "list clients err")
		return
	}
	resp.SuccessArray(c, len(clients), clients)
}

// @Summary	create client
// @Schemes
// @Description	create client
// @Tags		client
// @Param		clientId	path	string	true	"client id"
// @Param		bd			body	model.RedirectURI	true	"body"
// @Success		200
// @Router		/api/quick/clients/{clientId}/redirect-uri [post]
func (o Controller) createRedirectUri(c *gin.Context) {
	var in model.RedirectURI
	if err := o.SetCtx(c).BindUri(&in).Error; err != nil {
		resp.ErrorRequest(c, err, "init client req err")
		return
	}
	in.ClientID = c.Param("clientId")
	client, err := o.svc.CreateRedirectUri(in)
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
// @Param		bd			body	model.RedirectURI	true	"body"
// @Success		200
// @Router		/api/quick/clients/{clientId}/redirect-uri/{uriId} [put]
func (o Controller) modifyRedirectUri(c *gin.Context) {
	var in model.RedirectURI
	if err := o.SetCtx(c).BindUriAndJson(&in).Error; err != nil {
		resp.ErrorRequest(c, err, "init client req err")
		return
	}
	in.ClientID = c.Param("clientId")
	if err := o.svc.ModifyRedirectUri(in); err != nil {
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
// @Router		/api/quick/clients/{clientId}/redirect-uri/{uriId} [delete]
func (o Controller) deleteRedirectUri(c *gin.Context) {
	clientId := c.Param("clientId")
	uriId := c.Param("uriId")
	if err := o.svc.DeleteRedirectUri(clientId, uriId); err != nil {
		resp.ErrorUnknown(c, err, "delete client err")
		return
	}
	resp.Success(c)
}
