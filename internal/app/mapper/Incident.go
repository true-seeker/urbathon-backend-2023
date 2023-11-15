package mapper

import (
	"urbathon-backend-2023/internal/app/model/entity"
	"urbathon-backend-2023/internal/app/model/response"
)

func IncidentToIncidentResponse(incident *entity.Incident) *response.Incident {
	r := &response.Incident{
		Id:   incident.Id,
		Db:   DatabaseToDatabaseResponse(incident.Db),
		Date: incident.Date,
	}

	return r
}

func IncidentToIncidentResponses(incidents *[]entity.Incident) *[]response.Incident {
	rs := make([]response.Incident, 0)

	for _, incident := range *incidents {
		rs = append(rs, *IncidentToIncidentResponse(&incident))
	}

	return &rs
}
