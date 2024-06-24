package analyzer

import (
	"fmt"
	"os"
)

type Analyzer struct {
	propertyAnalyzers []PropertyAnalyzer
}

type PropertyAnalyzer interface {
	GetPropertyName() string
	GetPropertyValue(File) (any, error)
}

func New(a ...PropertyAnalyzer) *Analyzer {
	return &Analyzer{
		propertyAnalyzers: a,
	}
}

type Report map[string]any

type File struct {
	Name    string
	Content []byte
}

func (a *Analyzer) Analyze(filename string) (Report, error) {
	fileContent, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("reading file: %w", err)
	}

	file := File{
		Name:    filename,
		Content: fileContent,
	}

	report := make(map[string]any)

	for _, analyzer := range a.propertyAnalyzers {
		propName := analyzer.GetPropertyName()

		propV, err := analyzer.GetPropertyValue(file)
		if err != nil {
			return nil, fmt.Errorf("analyzing property %v: %w", propName, err)
		}

		report[propName] = propV
	}

	return report, nil
}
