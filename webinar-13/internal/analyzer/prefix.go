package analyzer

type PrefixWrapper struct {
	PropertyAnalyzer
	prefix string
}

func NewPrefixWrapper(original PropertyAnalyzer, prefix string) PrefixWrapper {
	return PrefixWrapper{
		PropertyAnalyzer: original,
		prefix:           prefix,
	}
}

func (w PrefixWrapper) GetPropertyName() string {
	return w.prefix + w.PropertyAnalyzer.GetPropertyName()
}
