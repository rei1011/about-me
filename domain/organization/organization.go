package organization

import "github.com/google/uuid"

type Organization struct {
	Id   uuid.UUID
	Name string
}

func NewOrganization(name string) Organization {
	uuid, _ := uuid.NewUUID()
	return Organization{
		Id:   uuid,
		Name: name,
	}
}
