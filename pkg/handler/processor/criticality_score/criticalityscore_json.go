package criticalityscore

type JSONCriticalityScoreResult struct {
	DefaultScore string     `json:"default_score"`
	Legacy       jsonLegacy `json:"legacy"`
	Repo         jsonRepo   `json:"repo"`
}
type jsonLegacy struct {
	ClosedIssuesCount     int `json:"closed_issues_count"`
	CommitFrequency       int `json:"commit_frequency"`
	ContributorCount      int `json:"contributor_count"`
	CreatedSince          int `json:"created_since"`
	GithubMentionCount    int `json:"github_mention_count"`
	IssueCommentFrequency int `json:"issue_comment_frequency"`
	OrgCount              int `json:"org_count"`
	RecentReleaseCount    int `json:"recent_release_count"`
	UpdatedIssuesCount    int `json:"updated_issues_count"`
	UpdatedSince          int `json:"updated_since"`
}
type jsonRepo struct {
	CreatedAt string `json:"created_at"`
	Language  string `json:"language"`
	License   string `json:"license"`
	StarCount int    `json:"star_count"`
	UpdatedAt string `json:"updated_at"`
	URL       string `json:"url"`
}
