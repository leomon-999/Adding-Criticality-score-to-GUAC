package guesser

import (
	"testing"

	"github.com/guacsec/guac/internal/testing/testdata"
	"github.com/guacsec/guac/pkg/handler/processor"
)

// 测试文档类型猜测方法
func Test_criticalityscoreTypeGuesser_GuessDocumentType(t *testing.T) {
	testCases := []struct {
		name     string
		blob     []byte
		expected processor.DocumentType
	}{{
		name: "invalid criticalityscore Document",
		blob: []byte(`{
			"abc": "def"
		}`),
		expected: processor.DocumentUnknown,
	}, {
		name:     "valid criticalityscore Document",
		blob:     testdata.CriticalityscoreExample,
		expected: processor.DocumentCriticalityscore,
	}}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			guesser := &criticalityscoreTypeGuesser{}
			f := guesser.GuessDocumentType(tt.blob, processor.FormatJSON)
			if f != tt.expected {
				t.Errorf("got the wrong format, got %v, expected %v", f, tt.expected)
			}
		})
	}
}
