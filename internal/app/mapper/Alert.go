package mapper

import (
	"urbathon-backend-2023/internal/app/model/entity"
	"urbathon-backend-2023/internal/app/model/response"
)

func AlertToAlertResponse(alert *entity.Alert) *response.Alert {
	r := &response.Alert{
		Id:       alert.Id,
		Incident: IncidentToIncidentResponse(alert.Incident),
		IsSent:   alert.IsSent,
	}

	return r
}

func AlertToAlertResponses(alerts *[]entity.Alert) *[]response.Alert {
	rs := make([]response.Alert, 0)

	for _, alert := range *alerts {
		rs = append(rs, *AlertToAlertResponse(&alert))
	}

	return &rs
}
