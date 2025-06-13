package upload

type File struct {
	FileName string `json:"file_name"`
	FileSize string `json:"file_size"`
	FileType string `json:"file_type"`
	FilePath string `json:"file_path"`
}