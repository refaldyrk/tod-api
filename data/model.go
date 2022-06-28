package data

type DataApiModel struct {
	Type string `bson:"type"`
	Data string `bson:"data"`
}
