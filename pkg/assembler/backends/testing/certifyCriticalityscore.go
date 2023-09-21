package testing

import (
	"context"
	"time"

	"github.com/guacsec/guac/pkg/assembler/graphql/model"
)

// 注册所有认证关键度分数
func registerAllCertifyCriticalityscore(client *demoClient) error {
	// "git", "github", "https://github.com/laiyoufafa/aafwk_aafwk_lite", "tag=3.1.1"
	selectedSourceType := "git"
	selectedSourceNameSpace := "github"
	selectedSourceName := "https://github.com/laiyoufafa/aafwk_aafwk_lite"
	selectedTag := "3.1.1"
	selectedSourceSpec := &model.SourceSpec{Type: &selectedSourceType, Namespace: &selectedSourceNameSpace, Name: &selectedSourceName, Tag: &selectedTag}
	selectedSource, err := client.Sources(context.TODO(), selectedSourceSpec)
	if err != nil {
		return err
	}

	err = client.registerCertifyCriticalityscore(selectedSource[0], time.Now(), "0.26270", 0, 24, 6, 0, 0)
	if err != nil {
		return err
	}
	return nil
}

// Ingest CertifyCriticalityscore
func (c *demoClient) registerCertifyCriticalityscore(selectedSource *model.Source, timeScanned time.Time, defaultScore string, legacyCommitFrequency float64, legacyContributorCount, legacyRecentReleaseCount, legacyUpdatedIssuesCount, repoStarCount int) error {
	for _, h := range c.certifyCriticalityscore {
		if h.Source == selectedSource &&
			h.Criticalityscore.DefaultScore == defaultScore &&
			h.Criticalityscore.LegacyCommitFrequency == legacyCommitFrequency &&
			h.Criticalityscore.LegacyContributorCount == legacyContributorCount &&
			h.Criticalityscore.LegacyRecentReleaseCount == legacyRecentReleaseCount &&
			h.Criticalityscore.LegacyUpdatedIssuesCount == legacyUpdatedIssuesCount &&
			h.Criticalityscore.RepoStarCount == repoStarCount {
			return nil
		}
	}

	newCertifyCriticalityscore := &model.CertifyCriticalityscore{
		Source: selectedSource,
		Criticalityscore: &model.Criticalityscore{
			TimeScanned:              timeScanned.String(),
			DefaultScore:             defaultScore,
			LegacyCommitFrequency:    legacyCommitFrequency,
			LegacyContributorCount:   legacyContributorCount,
			LegacyRecentReleaseCount: legacyRecentReleaseCount,
			LegacyUpdatedIssuesCount: legacyUpdatedIssuesCount,
			RepoStarCount:            repoStarCount,
			Origin:                   "testing backend",
			Collector:                "testing backend",
		},
	}
	c.certifyCriticalityscore = append(c.certifyCriticalityscore, newCertifyCriticalityscore)

	return nil
}

// Query CertifyCriticalityscore
func (c *demoClient) Criticalityscore(ctx context.Context, certifyCriticalityscoreSpec *model.CertifyCriticalityscoreSpec) ([]*model.CertifyCriticalityscore, error) {
	var collectedHasSourceAt []*model.CertifyCriticalityscore

	for _, h := range c.certifyCriticalityscore {
		matchOrSkip := true

		if certifyCriticalityscoreSpec.LegacyCommitFrequency != nil &&
			h.Criticalityscore.LegacyCommitFrequency != *certifyCriticalityscoreSpec.LegacyCommitFrequency {
			matchOrSkip = false
		}
		if certifyCriticalityscoreSpec.LegacyContributorCount != nil &&
			h.Criticalityscore.LegacyContributorCount != *certifyCriticalityscoreSpec.LegacyContributorCount {
			matchOrSkip = false
		}
		if certifyCriticalityscoreSpec.LegacyRecentReleaseCount != nil &&
			h.Criticalityscore.LegacyRecentReleaseCount != *certifyCriticalityscoreSpec.LegacyRecentReleaseCount {
			matchOrSkip = false
		}
		if certifyCriticalityscoreSpec.LegacyUpdatedIssuesCount != nil &&
			h.Criticalityscore.LegacyUpdatedIssuesCount != *certifyCriticalityscoreSpec.LegacyUpdatedIssuesCount {
			matchOrSkip = false
		}
		if certifyCriticalityscoreSpec.RepoStarCount != nil &&
			h.Criticalityscore.RepoStarCount != *certifyCriticalityscoreSpec.RepoStarCount {
			matchOrSkip = false
		}
		if certifyCriticalityscoreSpec.Collector != nil &&
			h.Criticalityscore.Collector != *certifyCriticalityscoreSpec.Collector {
			matchOrSkip = false
		}
		if certifyCriticalityscoreSpec.Origin != nil &&
			h.Criticalityscore.Origin != *certifyCriticalityscoreSpec.Origin {
			matchOrSkip = false
		}

		if certifyCriticalityscoreSpec.Source != nil &&
			h.Source != nil {
			if certifyCriticalityscoreSpec.Source.Type == nil || h.Source.Type == *certifyCriticalityscoreSpec.Source.Type {
				newSource, err := filterSourceNamespace(h.Source, certifyCriticalityscoreSpec.Source)
				if err != nil {
					return nil, err
				}
				if newSource == nil {
					matchOrSkip = false
				}
			}
		}

		if matchOrSkip {
			collectedHasSourceAt = append(collectedHasSourceAt, h)
		}
	}
	return collectedHasSourceAt, nil
}
