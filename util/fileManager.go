package util

type FileManager struct {
	Input  string
	Output string
}

func NewFileManager(inputPath, outputFileName string) *FileManager {
	return &FileManager{
		Input:  inputPath,
		Output: outputFileName,
	}
}
