package config

type FileConfig struct {
	JsonFilePath string
}

func GetFileConfig() *FileConfig {
	return &FileConfig{
		JsonFilePath: getEnv("JSON_FILE_PATH", ""),
	}
}
