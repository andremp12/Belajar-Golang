package response

import database "booking-konstruksi/database/migration"

type Satuan struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func SatuanResponse(data *database.Satuan) Satuan {
	return Satuan{
		ID:   data.ID,
		Name: data.Name,
	}
}
