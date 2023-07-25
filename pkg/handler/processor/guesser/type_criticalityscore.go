package guesser

import (
	"encoding/json"

	"github.com/guacsec/guac/pkg/handler/processor"
	cs "github.com/guacsec/guac/pkg/handler/processor/criticality_score"
)

type criticalityscoreTypeGuesser struct{}

func (_ *criticalityscoreTypeGuesser) GuessDocumentType(blob []byte, format processor.FormatType) processor.DocumentType {
	var criticalityscore cs.JSONCriticalityScoreResult
	if json.Unmarshal(blob, &criticalityscore) == nil && format == processor.FormatJSON {
		if criticalityscore.Repo.License != "" || criticalityscore.Repo.URL != "" {
			return processor.DocumentCriticalityscore
		}
	}
	return processor.DocumentUnknown
}
