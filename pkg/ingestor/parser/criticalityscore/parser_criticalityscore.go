package criticalityscore

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/guacsec/guac/pkg/assembler"
	"github.com/guacsec/guac/pkg/handler/processor"
	cs "github.com/guacsec/guac/pkg/handler/processor/criticalityscore"
	"github.com/guacsec/guac/pkg/ingestor/parser/common"
)

type criticalityscoreParser struct {
	criticalityscoreNodes []assembler.MetadataNode
	// artifactNode should have a 1:1 mapping to the index
	// of criticalityscoreNodes.
	artifactNodes []assembler.ArtifactNode
}

// NewCriticalityscoreParser initializes the criticalityscoreParser
func NewCriticalityscoreParser() common.DocumentParser {
	return &criticalityscoreParser{
		criticalityscoreNodes: []assembler.MetadataNode{},
		artifactNodes:         []assembler.ArtifactNode{},
	}
}

// Parse breaks out the document into the graph components
func (p *criticalityscoreParser) Parse(ctx context.Context, doc *processor.Document) error {

	if doc.Type != processor.DocumentCriticalityscore {
		return fmt.Errorf("expected document type: %v, actual document type: %v", processor.DocumentCriticalityscore, doc.Type)
	}

	switch doc.Format {
	case processor.FormatJSON:
		var criticalityscore cs.JSONCriticalityScoreResult
		if err := json.Unmarshal(doc.Blob, &criticalityscore); err != nil {
			return err
		}
		p.criticalityscoreNodes = append(p.criticalityscoreNodes, getMetadataNode(&criticalityscore))
		p.artifactNodes = append(p.artifactNodes, getArtifactNode(&criticalityscore))
		return nil
	}
	return fmt.Errorf("unable to support parsing of Scorecard document format: %v", doc.Format)
}

// CreateNodes creates the GuacNode for the graph inputs
func (p *criticalityscoreParser) CreateNodes(ctx context.Context) []assembler.GuacNode {
	nodes := []assembler.GuacNode{}
	for _, n := range p.criticalityscoreNodes {
		nodes = append(nodes, n)
	}
	for _, n := range p.artifactNodes {
		nodes = append(nodes, n)
	}

	return nodes
}

// CreateEdges creates the GuacEdges that form the relationship for the graph inputs
func (p *criticalityscoreParser) CreateEdges(ctx context.Context, foundIdentities []assembler.IdentityNode) []assembler.GuacEdge {
	// TODO: handle identity for edges (https://github.com/guacsec/guac/issues/128)
	edges := []assembler.GuacEdge{}
	for i, s := range p.criticalityscoreNodes {
		edges = append(edges, assembler.MetadataForEdge{
			MetadataNode: s,
			ForArtifact:  p.artifactNodes[i],
		})
	}
	return edges
}

// GetIdentities gets the identity node from the document if they exist
func (p *criticalityscoreParser) GetIdentities(ctx context.Context) []assembler.IdentityNode {
	return nil
}

func metadataId(s *cs.JSONCriticalityScoreResult) string {
	return fmt.Sprintf("%v:%v", s.Repo.URL[len("https://"):len(s.Repo.URL)], s.Repo.License)
}

func getMetadataNode(s *cs.JSONCriticalityScoreResult) assembler.MetadataNode {
	mnNode := assembler.MetadataNode{
		MetadataType: "criticalityscore",
		ID:           metadataId(s),
		Details:      map[string]interface{}{},
	}

	mnNode.Details["repo"] = sourceUri(s.Repo.URL)
	mnNode.Details["language"] = s.Repo.Language
	mnNode.Details["star_count"] = s.Repo.StarCount
	mnNode.Details["commit_frequency"] = s.Legacy.CommitFrequency
	mnNode.Details["contributor_count"] = s.Legacy.ContributorCount
	mnNode.Details["github_mention_count"] = s.Legacy.GithubMentionCount
	mnNode.Details["issue_comment_frequency"] = s.Legacy.IssueCommentFrequency
	mnNode.Details["recent_release_count"] = s.Legacy.RecentReleaseCount
	mnNode.Details["org_count"] = s.Legacy.OrgCount
	mnNode.Details["updated_issues_count"] = s.Legacy.UpdatedIssuesCount
	mnNode.Details["score"] = s.DefaultScore

	return mnNode
}

func getArtifactNode(s *cs.JSONCriticalityScoreResult) assembler.ArtifactNode {
	return assembler.ArtifactNode{
		Name: sourceUri(s.Repo.URL),
	}
}

func sourceUri(s string) string {
	return "git+" + s
}

func (p *criticalityscoreParser) GetIdentifiers(ctx context.Context) (*common.IdentifierStrings, error) {
	return nil, fmt.Errorf("not yet implemented")
}
