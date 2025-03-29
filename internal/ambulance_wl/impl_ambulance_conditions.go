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
	// update ambulance document
	updateAmbulanceFunc(c, func(
		c *gin.Context,
		ambulance *Ambulance,
	) (updatedAmbulance *Ambulance, responseContent interface{}, status int) {
		result := ambulance.PredefinedConditions
		if result == nil {
			result = []Condition{}
		}
		return nil, result, http.StatusOK
	})
}
