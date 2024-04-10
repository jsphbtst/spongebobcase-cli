package types

type JsonFile struct {
	Data []string `json:"data"`
}

func (j *JsonFile) AddData(text string) {
	j.Data = append(j.Data, text)

	if len(j.Data) > 50 {
		j.Data = j.Data[1:]
	}
}
