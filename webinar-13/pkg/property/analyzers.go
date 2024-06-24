package property

import "webinar13/internal/analyzer"

type FilenameAnalyzer struct{}

func (a FilenameAnalyzer) GetPropertyName() string {
	return "filename"
}

func (a FilenameAnalyzer) GetPropertyValue(f analyzer.File) (any, error) {
	return f.Name, nil
}

type ContentLengthAnalyzer struct{}

func (a ContentLengthAnalyzer) GetPropertyName() string {
	return "content_length"
}

func (a ContentLengthAnalyzer) GetPropertyValue(f analyzer.File) (any, error) {
	return len(f.Content), nil
}
