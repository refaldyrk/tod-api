package apikey

type ApikeyModel struct {
	Owner string `bson:"owner"`
	Key   string `bson:"key"`
}
