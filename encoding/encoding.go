package encoding

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/Yandex-Practicum/final-project-encoding-go/models"
	"gopkg.in/yaml.v2"
)

// JSONData тип для перекодирования из JSON в YAML
type JSONData struct {
	DockerCompose *models.DockerCompose
	FileInput     string
	FileOutput    string
}

// YAMLData тип для перекодирования из YAML в JSON
type YAMLData struct {
	DockerCompose *models.DockerCompose
	FileInput     string
	FileOutput    string
}

// MyEncoder интерфейс для структур YAMLData и JSONData
type MyEncoder interface {
	Encoding() error
}

// Encoding перекодирует файл из JSON в YAML
func (j *JSONData) Encoding() error {
	// ниже реализуйте метод
	yamlData, err := yaml.Marshal(j)
	if err != nil {
		return fmt.Errorf("ошибка преобразования в YAML: %w", err)
	}

	// Запись YAML в файл
	err = os.WriteFile("output.yaml", yamlData, 0)
	if err != nil {
		return fmt.Errorf("ошибка записи файла: %w", err)
	}

	return nil
}

// Encoding перекодирует файл из YAML в JSON
func (y *YAMLData) Encoding() error {
	// Ниже реализуйте метод
	jsonData, err := json.Marshal(y)
	if err != nil {
		return fmt.Errorf("ошибка преобразования в JSON: %w", err)
	}

	// Запись YAML в файл
	err = os.WriteFile("output.yaml", jsonData, 0)
	if err != nil {
		return fmt.Errorf("ошибка записи файла: %w", err)
	}

	return nil
}
