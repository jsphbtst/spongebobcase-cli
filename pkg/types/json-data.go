package types

type JsonFile struct {
	Data []string `json:"data"`
}

func (j *JsonFile) AddData(text string) {
	for _, record := range j.Data {
		if record == text {
			return
		}
	}

	j.Data = append(j.Data, text)

	if len(j.Data) > 50 {
		j.Data = j.Data[1:]
	}
}
