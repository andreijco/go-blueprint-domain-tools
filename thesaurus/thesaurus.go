package thesaurus
// Thesaurus interface
type Thesaurus interface {
	Synonyms(term string) ([]string, error)
}