package controller

import (
	"QuickAuth/internal/endpoint/request"
	"QuickAuth/internal/endpoint/resp"
	"github.com/gin-gonic/gin"
)

// @Summary	get redirect uri list
// @Schemes
// @Description	get redirect uri list
// @Tags		client
// @Param		clientId	path		string	true	"client id"
// @Success		200			{object}	[]string
// @Router		/api/quick/clients/{clientId}/redirect-uri [get]
func (o Controller) listRedirectUri(c *gin.Context) {
	clients, err := o.svc.ListRedirectUri(c.Param("clientId"))
	if err != nil {
		resp.ErrorSelect(c, err, "list redirect uri err")
		return
	}
	resp.SuccessArray(c, len(clients), clients)
}

// @Summary	create client
// @Schemes
// @Description	create client
// @Tags		client
// @Param		clientId	path	string				true	"client id"
// @Param		bd			body	request.RedirectUriReq	true	"body"
// @Success		200
// @Router		/api/quick/clients/{clientId}/redirect-uri [post]
func (o Controller) createRedirectUri(c *gin.Context) {
	var in request.RedirectUriReq
	if err := o.SetCtx(c).BindUriAndJson(&in).Error; err != nil {
		resp.ErrorRequest(c, err, "init redirect uri req err")
		return
	}

	if err := o.svc.CreateRedirectUri(in.ClientId, in.Uri); err != nil {
		resp.ErrorUnknown(c, err, "create redirect uri err")
		return
	}
	resp.Success(c)
}

// @Summary	modify client
// @Schemes
// @Description	modify client
// @Tags		client
// @Param		clientId	path	string				true	"client id"
// @Param		uriId		path	string				true	"uri id"
// @Param		bd			body	request.RedirectUriReq	true	"body"
// @Success		200
// @Router		/api/quick/clients/{clientId}/redirect-uri/{uriId} [put]
func (o Controller) modifyRedirectUri(c *gin.Context) {
	var in request.RedirectUriReq
	if err := o.SetCtx(c).BindUriAndJson(&in).Error; err != nil {
		resp.ErrorRequest(c, err, "init redirect uri req err")
		return
	}
	if err := o.svc.ModifyRedirectUri(in.ClientId, in.UriId, in.Uri); err != nil {
		resp.ErrorUnknown(c, err, "modify redirect uri err")
		return
	}
	resp.Success(c)
}

// @Summary	delete client
// @Schemes
// @Description	delete client
// @Tags		client
// @Param		clientId	path	string	true	"client id"
// @Param		uri			path	string	true	"uri name"
// @Success		200
// @Router		/api/quick/clients/{clientId}/redirect-uri/{uri} [delete]
func (o Controller) deleteRedirectUri(c *gin.Context) {
	clientId := c.Param("clientId")
	uri := c.Param("uri")
	if err := o.svc.DeleteRedirectUri(clientId, uri); err != nil {
		resp.ErrorUnknown(c, err, "delete redirect uri err")
		return
	}
	resp.Success(c)
}
