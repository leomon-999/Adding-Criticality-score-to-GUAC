package criticalityscore

import (
	"encoding/json"
	"fmt"

	"github.com/guacsec/guac/pkg/handler/processor"
)

type CriticalityscoreProcessor struct {
}

// 对criticalityscore的json数据结果进行验证
func (p *CriticalityscoreProcessor) ValidateSchema(d *processor.Document) error {
	if d.Type != processor.DocumentCriticalityscore {
		return fmt.Errorf("expected document type: %v, actual document type: %v", processor.DocumentCriticalityscore, d.Type)
	}

	switch d.Format {
	case processor.FormatJSON:
		var criticalityscore JSONCriticalityScoreResult
		if err := json.Unmarshal(d.Blob, &criticalityscore); err != nil {
			return err
		}
		if criticalityscore.Repo.License == "" ||
			criticalityscore.Repo.URL == "" {
			return fmt.Errorf("missing required criticalityscore fields")
		}

		return nil
	}

	return fmt.Errorf("unable to support parsing of criticalityscore document format: %v", d.Format)
}

// Unpack takes in the document and tries to unpack it
// if there is a valid decomposition of sub-documents.
//
// Returns empty list and nil error if nothing to unpack
// Returns unpacked list and nil error if successfully unpacked
func (p *CriticalityscoreProcessor) Unpack(d *processor.Document) ([]*processor.Document, error) {
	if d.Type != processor.DocumentCriticalityscore {
		return nil, fmt.Errorf("expected document type: %v, actual document type: %v", processor.DocumentCriticalityscore, d.Type)
	}

	// Criticalityscore doesn't unpack into additional documents at the moment.
	return []*processor.Document{}, nil
}
