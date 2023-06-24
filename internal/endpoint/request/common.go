package request

type Pages struct {
	Size uint `json:"size" binding:"max:100"`
	Num  uint `json:"num"`
}
