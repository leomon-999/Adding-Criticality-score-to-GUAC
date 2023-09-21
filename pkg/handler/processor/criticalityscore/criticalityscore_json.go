package criticalityscore

//根据criticalityscore的json数据结果定义数据类型
type JSONCriticalityScoreResult struct {
	DefaultScore string     `json:"default_score"`
	Legacy       jsonLegacy `json:"legacy"`
	Repo         jsonRepo   `json:"repo"`
}

//二级类型定义
type jsonLegacy struct {
	ClosedIssuesCount     int     `json:"closed_issues_count"`
	CommitFrequency       float64 `json:"commit_frequency"`
	ContributorCount      int     `json:"contributor_count"`
	CreatedSince          int     `json:"created_since"`
	GithubMentionCount    int     `json:"github_mention_count"`
	IssueCommentFrequency float64 `json:"issue_comment_frequency"`
	OrgCount              int     `json:"org_count"`
	RecentReleaseCount    int     `json:"recent_release_count"`
	UpdatedIssuesCount    int     `json:"updated_issues_count"`
	UpdatedSince          int     `json:"updated_since"`
}

//二级类型定义
type jsonRepo struct {
	CreatedAt string `json:"created_at"`
	Language  string `json:"language"`
	License   string `json:"license"`
	StarCount int    `json:"star_count"`
	UpdatedAt string `json:"updated_at"`
	URL       string `json:"url"`
}
