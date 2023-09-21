package criticalityscore

import (
	"reflect"
	"testing"

	"github.com/guacsec/guac/internal/testing/testdata"
	"github.com/guacsec/guac/pkg/handler/processor"
)

// 测试unpack方法
func TestCriticalityscoreProcessor_Unpack(t *testing.T) {
	testCases := []struct {
		name      string
		doc       processor.Document
		expected  []*processor.Document
		expectErr bool
	}{{
		name: "Criticalityscore document",
		doc: processor.Document{
			Blob:              testdata.CriticalityscoreExample,
			Format:            processor.FormatUnknown,
			Type:              processor.DocumentCriticalityscore,
			SourceInformation: processor.SourceInformation{},
		},
		expected:  []*processor.Document{},
		expectErr: false,
	}, {
		name: "Incorrect type",
		doc: processor.Document{
			Blob:              testdata.CriticalityscoreExample,
			Format:            processor.FormatUnknown,
			Type:              processor.DocumentUnknown,
			SourceInformation: processor.SourceInformation{},
		},
		expected:  nil,
		expectErr: true,
	}}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			d := CriticalityscoreProcessor{}
			actual, err := d.Unpack(&tt.doc)
			if (err != nil) != tt.expectErr {
				t.Errorf("CriticalityscoreProcessor.Unpack() error = %v, expectErr %v", err, tt.expectErr)
			}
			if !reflect.DeepEqual(actual, tt.expected) {
				t.Errorf("CriticalityscoreProcessor.Unpack() = %v, expected %v", actual, tt.expected)
			}
		})
	}
}

// 测试验证方法
func TestCriticalityscoreProcessor_ValidateSchema(t *testing.T) {
	testCases := []struct {
		name      string
		doc       processor.Document
		expectErr bool
	}{{
		name: "valid Criticalityscore document",
		doc: processor.Document{
			Blob:              testdata.CriticalityscoreExample,
			Format:            processor.FormatJSON,
			Type:              processor.DocumentCriticalityscore,
			SourceInformation: processor.SourceInformation{},
		},
		expectErr: false,
	}, {
		name: "invalid Criticalityscore document",
		doc: processor.Document{
			Blob:              testdata.CriticalityscoreInvalid,
			Format:            processor.FormatJSON,
			Type:              processor.DocumentCriticalityscore,
			SourceInformation: processor.SourceInformation{},
		},
		expectErr: true,
	}, {
		name: "invalid format supported",
		doc: processor.Document{
			Blob:              testdata.CriticalityscoreExample,
			Format:            processor.FormatUnknown,
			Type:              processor.DocumentCriticalityscore,
			SourceInformation: processor.SourceInformation{},
		},
		expectErr: true,
	}}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			d := CriticalityscoreProcessor{}
			err := d.ValidateSchema(&tt.doc)
			if (err != nil) != tt.expectErr {
				t.Errorf("CriticalityscoreProcessor.ValidateSchema() error = %v, expectErr %v", err, tt.expectErr)
			}
		})
	}
}
