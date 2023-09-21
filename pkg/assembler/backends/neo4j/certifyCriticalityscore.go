package neo4jBackend

import (
	"context"

	"strings"

	"github.com/guacsec/guac/pkg/assembler/graphql/model"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j/dbtype"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

const (
	//扫描时间
	csTimeScanned string = "csTimeScanned"
	//得分
	defaultScore string = "defaultScore"
	//提交次数
	legacyCommitFrequency string = "legacyCommitFrequency"
	//贡献者数量
	legacyContributorCount string = "legacyContributorCount"
	//最近发布的数量
	legacyRecentReleaseCount string = "legacyRecentReleaseCount"
	//更新问题数
	legacyUpdatedIssuesCount string = "legacyUpdatedIssuesCount"
	//获得的星数
	repoStarCount string = "repoStarCount"
)

// 将Criticalityscore和neo4j客户端建立联系
func (c *neo4jClient) Criticalityscore(ctx context.Context, certifyCriticalityscoreSpec *model.CertifyCriticalityscoreSpec) ([]*model.CertifyCriticalityscore, error) {
	session := c.driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeRead})
	defer session.Close()

	var sb strings.Builder
	var firstMatch bool = true
	queryValues := map[string]any{}

	query := "MATCH (root:Src)-[:SrcHasType]->(type:SrcType)-[:SrcHasNamespace]->(namespace:SrcNamespace)" +
		"-[:SrcHasName]->(name:SrcName)-[:subject]-(certifyCriticalityscore:CertifyCriticalityscore)"
	sb.WriteString(query)

	setSrcMatchValues(&sb, certifyCriticalityscoreSpec.Source, false, &firstMatch, queryValues)
	setCertifyCriticalityscoreValues(&sb, certifyCriticalityscoreSpec, &firstMatch, queryValues)
	sb.WriteString(" RETURN type.type, namespace.namespace, name.name, name.tag, name.commit, certifyCriticalityscore")

	result, err := session.ReadTransaction(
		func(tx neo4j.Transaction) (interface{}, error) {
			result, err := tx.Run(sb.String(), queryValues)
			if err != nil {
				return nil, err
			}

			collectedCertifyCriticalityscore := []*model.CertifyCriticalityscore{}

			for result.Next() {
				commitString := ""
				if result.Record().Values[4] != nil {
					commitString = result.Record().Values[4].(string)
				}

				tagString := ""
				if result.Record().Values[3] != nil {
					tagString = result.Record().Values[3].(string)
				}

				nameString := result.Record().Values[2].(string)
				namespaceString := result.Record().Values[1].(string)
				typeString := result.Record().Values[0].(string)

				srcName := &model.SourceName{
					Name:   nameString,
					Tag:    &tagString,
					Commit: &commitString,
				}

				srcNamespace := &model.SourceNamespace{
					Namespace: namespaceString,
					Names:     []*model.SourceName{srcName},
				}

				src := model.Source{
					Type:       typeString,
					Namespaces: []*model.SourceNamespace{srcNamespace},
				}

				certifyCriticalityscoreNode := dbtype.Node{}
				if result.Record().Values[5] != nil {
					certifyCriticalityscoreNode = result.Record().Values[5].(dbtype.Node)
				} else {
					return nil, gqlerror.Errorf("certifyCriticalityscore Node not found in neo4j")
				}

				//设置数据类型
				criticalityscore := model.Criticalityscore{
					TimeScanned:              certifyCriticalityscoreNode.Props[timeScanned].(string),
					DefaultScore:             certifyCriticalityscoreNode.Props[defaultScore].(string),
					LegacyCommitFrequency:    certifyCriticalityscoreNode.Props[legacyCommitFrequency].(float64),
					LegacyContributorCount:   certifyCriticalityscoreNode.Props[legacyContributorCount].(int),
					LegacyRecentReleaseCount: certifyCriticalityscoreNode.Props[legacyRecentReleaseCount].(int),
					LegacyUpdatedIssuesCount: certifyCriticalityscoreNode.Props[legacyUpdatedIssuesCount].(int),
					RepoStarCount:            certifyCriticalityscoreNode.Props[repoStarCount].(int),
					Origin:                   certifyCriticalityscoreNode.Props[origin].(string),
					Collector:                certifyCriticalityscoreNode.Props[collector].(string),
				}

				certifyCriticalityscore := &model.CertifyCriticalityscore{
					Source:           &src,
					Criticalityscore: &criticalityscore,
				}

				collectedCertifyCriticalityscore = append(collectedCertifyCriticalityscore, certifyCriticalityscore)
			}
			if err = result.Err(); err != nil {
				return nil, err
			}

			return collectedCertifyCriticalityscore, nil
		})
	if err != nil {
		return nil, err
	}

	return result.([]*model.CertifyCriticalityscore), nil
}

// 设置验证criticcalityscore的值
func setCertifyCriticalityscoreValues(sb *strings.Builder, certifyCriticalityscoreSpec *model.CertifyCriticalityscoreSpec, firstMatch *bool, queryValues map[string]any) {
	if certifyCriticalityscoreSpec.TimeScanned != nil {
		matchProperties(sb, *firstMatch, "certifyCriticalityscore", timeScanned, "$"+timeScanned)
		*firstMatch = false
		queryValues[timeScanned] = certifyCriticalityscoreSpec.TimeScanned
	}
	if certifyCriticalityscoreSpec.DefaultScore != nil {
		matchProperties(sb, *firstMatch, "certifyCriticalityscore", defaultScore, "$"+defaultScore)
		*firstMatch = false
		queryValues[defaultScore] = certifyCriticalityscoreSpec.DefaultScore
	}
	if certifyCriticalityscoreSpec.LegacyCommitFrequency != nil {
		matchProperties(sb, *firstMatch, "certifyCriticalityscore", legacyCommitFrequency, "$"+legacyCommitFrequency)
		*firstMatch = false
		queryValues[legacyCommitFrequency] = certifyCriticalityscoreSpec.LegacyCommitFrequency
	}
	if certifyCriticalityscoreSpec.LegacyContributorCount != nil {
		matchProperties(sb, *firstMatch, "certifyCriticalityscore", legacyContributorCount, "$"+legacyContributorCount)
		*firstMatch = false
		queryValues[legacyContributorCount] = certifyCriticalityscoreSpec.LegacyContributorCount
	}
	if certifyCriticalityscoreSpec.LegacyRecentReleaseCount != nil {
		matchProperties(sb, *firstMatch, "certifyCriticalityscore", legacyRecentReleaseCount, "$"+legacyRecentReleaseCount)
		*firstMatch = false
		queryValues[legacyRecentReleaseCount] = certifyCriticalityscoreSpec.LegacyRecentReleaseCount
	}
	if certifyCriticalityscoreSpec.LegacyUpdatedIssuesCount != nil {
		matchProperties(sb, *firstMatch, "certifyCriticalityscore", legacyUpdatedIssuesCount, "$"+legacyUpdatedIssuesCount)
		*firstMatch = false
		queryValues[legacyUpdatedIssuesCount] = certifyCriticalityscoreSpec.LegacyUpdatedIssuesCount
	}
	if certifyCriticalityscoreSpec.RepoStarCount != nil {
		matchProperties(sb, *firstMatch, "certifyCriticalityscore", repoStarCount, "$"+repoStarCount)
		*firstMatch = false
		queryValues[repoStarCount] = certifyCriticalityscoreSpec.RepoStarCount
	}
	if certifyCriticalityscoreSpec.Origin != nil {
		matchProperties(sb, *firstMatch, "certifyCriticalityscore", "origin", "$origin")
		*firstMatch = false
		queryValues["origin"] = certifyCriticalityscoreSpec.Origin
	}
	if certifyCriticalityscoreSpec.Collector != nil {
		matchProperties(sb, *firstMatch, "certifyCriticalityscore", "collector", "$collector")
		*firstMatch = false
		queryValues["collector"] = certifyCriticalityscoreSpec.Collector
	}
}
