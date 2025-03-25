package ambulance_wl

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type implAmbulanceWaitingListAPI struct {
}

func NewAmbulanceWaitingListApi() AmbulanceWaitingListAPI {
	return &implAmbulanceWaitingListAPI{}
}

func (o implAmbulanceWaitingListAPI) CreateWaitingListEntry(c *gin.Context) {
	c.AbortWithStatus(http.StatusNotImplemented)
}

func (o implAmbulanceWaitingListAPI) DeleteWaitingListEntry(c *gin.Context) {
	c.AbortWithStatus(http.StatusNotImplemented)
}

func (o implAmbulanceWaitingListAPI) GetWaitingListEntries(c *gin.Context) {
	c.AbortWithStatus(http.StatusNotImplemented)
}

func (o implAmbulanceWaitingListAPI) GetWaitingListEntry(c *gin.Context) {
	c.AbortWithStatus(http.StatusNotImplemented)
}

func (o implAmbulanceWaitingListAPI) UpdateWaitingListEntry(c *gin.Context) {
	c.AbortWithStatus(http.StatusNotImplemented)
}
