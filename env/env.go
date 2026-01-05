package env

import "os"

func GetModelBaseUrl() string {
	return os.Getenv("MODEL_BASE_URL")
}

func GetLLMModel() string {
	return os.Getenv("LLM_MODEL")
}

func GetModelApiKey() string {
	return os.Getenv("LLM_API_KEY")
}
