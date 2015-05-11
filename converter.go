package ssconverter

type Converter interface {
	CheckCommand() bool
	BuildSite() (sitePath string, err error)
}