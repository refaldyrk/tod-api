package data

type SingleDataFormatter struct {
	Type string `json:"type"`
	Data string `json:"data"`
}

func ToSingleData(model DataApiModel) SingleDataFormatter {
	return SingleDataFormatter{
		Type: model.Type,
		Data: model.Data,
	}
}

func ToManyData(models []DataApiModel) []SingleDataFormatter {
	var data []SingleDataFormatter
	for _, model := range models {
		data = append(data, ToSingleData(model))
	}
	return data
}
