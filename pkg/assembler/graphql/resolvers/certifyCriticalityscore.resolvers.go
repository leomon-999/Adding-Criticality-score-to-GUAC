package resolvers

//该文件将根据模式自动重新生成，任何解析器实现将在生成时被复制，任何未知代码将被移至末尾。
import (
	"context"

	"github.com/guacsec/guac/pkg/assembler/graphql/model"
)

// Criticalityscore 临界分数的字段的解析器。
func (r *queryResolver) Criticalityscore(ctx context.Context, criticalityscoreSpec *model.CertifyCriticalityscoreSpec) ([]*model.CertifyCriticalityscore, error) {
	return r.Backend.Criticalityscore(ctx, criticalityscoreSpec)
}
