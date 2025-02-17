package actions

var OutputDirPath = "./output/"

var config = OutputDirPath + "config.json"

const (
	FILES = "files"
)

type JsonCategoryStructure struct {
	Category string   `json:"category"`
	Keys     []string `json:"keys"`
}

type JsonStructure struct {
	Config JsonCategoryStructure `json:"config"`
}
