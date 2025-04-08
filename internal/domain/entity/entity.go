package entity

// Entity arayüzü, ID ve varlık tipi ile çalışmak için genel bir yapı sağlar
type Entity[ID comparable] interface {
	GetID() ID
}

var AllEntities = []interface{}{
	&User{},
}
