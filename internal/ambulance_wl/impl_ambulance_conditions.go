package ambulance_wl

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type implAmbulanceConditionsAPI struct {
}

func NewAmbulanceConditionsApi() AmbulanceConditionsAPI {
	return &implAmbulanceConditionsAPI{}
}

func (o implAmbulanceConditionsAPI) GetConditions(c *gin.Context) {
	c.AbortWithStatus(http.StatusNotImplemented)
}
