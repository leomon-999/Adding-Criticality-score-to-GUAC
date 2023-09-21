package generated

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"sync"

	"github.com/99designs/gqlgen/graphql"
	"github.com/guacsec/guac/pkg/assembler/graphql/model"
	"github.com/vektah/gqlparser/v2/ast"
)

// region    ************************** generated!.gotpl **************************

// endregion ************************** generated!.gotpl **************************

// region    ***************************** args.gotpl *****************************

// endregion ***************************** args.gotpl *****************************

// region    ************************** directives.gotpl **************************

// endregion ************************** directives.gotpl **************************

// region    **************************** field.gotpl *****************************

func (ec *executionContext) _CertifyCriticalityscore_source(ctx context.Context, field graphql.CollectedField, obj *model.CertifyCriticalityscore) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_CertifyCriticalityscore_source(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Source, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(*model.Source)
	fc.Result = res
	return ec.marshalNSource2ᚖgithubᚗcomᚋguacsecᚋguacᚋpkgᚋassemblerᚋgraphqlᚋmodelᚐSource(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_CertifyCriticalityscore_source(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "CertifyCriticalityscore",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			switch field.Name {
			case "type":
				return ec.fieldContext_Source_type(ctx, field)
			case "namespaces":
				return ec.fieldContext_Source_namespaces(ctx, field)
			}
			return nil, fmt.Errorf("no field named %q was found under type Source", field.Name)
		},
	}
	return fc, nil
}

func (ec *executionContext) _CertifyCriticalityscore_criticalityscore(ctx context.Context, field graphql.CollectedField, obj *model.CertifyCriticalityscore) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_CertifyCriticalityscore_criticalityscore(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Criticalityscore, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(*model.Criticalityscore)
	fc.Result = res
	return ec.marshalNCriticalityscore2ᚖgithubᚗcomᚋguacsecᚋguacᚋpkgᚋassemblerᚋgraphqlᚋmodelᚐCriticalityscore(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_CertifyCriticalityscore_criticalityscore(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "CertifyCriticalityscore",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			switch field.Name {
			case "defaultScore":
				return ec.fieldContext_Criticalityscore_defaultScore(ctx, field)
			case "timeScanned":
				return ec.fieldContext_Criticalityscore_timeScanned(ctx, field)
			case "legacyCommitFrequency":
				return ec.fieldContext_Criticalityscore_legacyCommitFrequency(ctx, field)
			case "legacyContributorCount":
				return ec.fieldContext_Criticalityscore_legacyContributorCount(ctx, field)
			case "legacyRecentReleaseCount":
				return ec.fieldContext_Criticalityscore_legacyRecentReleaseCount(ctx, field)
			case "legacyUpdatedIssuesCount":
				return ec.fieldContext_Criticalityscore_legacyUpdatedIssuesCount(ctx, field)
			case "repoStarCount":
				return ec.fieldContext_Criticalityscore_repoStarCount(ctx, field)
			case "origin":
				return ec.fieldContext_Criticalityscore_origin(ctx, field)
			case "collector":
				return ec.fieldContext_Criticalityscore_collector(ctx, field)
			}
			return nil, fmt.Errorf("no field named %q was found under type Criticalityscore", field.Name)
		},
	}
	return fc, nil
}

func (ec *executionContext) _Criticalityscore_defaultScore(ctx context.Context, field graphql.CollectedField, obj *model.Criticalityscore) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Criticalityscore_defaultScore(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.DefaultScore, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalNString2string(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_Criticalityscore_defaultScore(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Criticalityscore",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type Float does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _Criticalityscore_timeScanned(ctx context.Context, field graphql.CollectedField, obj *model.Criticalityscore) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Criticalityscore_timeScanned(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.TimeScanned, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalNString2string(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_Criticalityscore_timeScanned(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Criticalityscore",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type String does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _Criticalityscore_legacyCommitFrequency(ctx context.Context, field graphql.CollectedField, obj *model.Criticalityscore) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Criticalityscore_legacyCommitFrequency(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.LegacyCommitFrequency, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(int)
	fc.Result = res
	return ec.marshalNInt2int(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_Criticalityscore_legacyCommitFrequency(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Criticalityscore",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type String does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _Criticalityscore_legacyContributorCount(ctx context.Context, field graphql.CollectedField, obj *model.Criticalityscore) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Criticalityscore_legacyContributorCount(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.LegacyContributorCount, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(int)
	fc.Result = res
	return ec.marshalNInt2int(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_Criticalityscore_legacyContributorCount(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Criticalityscore",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type String does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _Criticalityscore_legacyRecentReleaseCount(ctx context.Context, field graphql.CollectedField, obj *model.Criticalityscore) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Criticalityscore_legacyRecentReleaseCount(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.LegacyRecentReleaseCount, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(int)
	fc.Result = res
	return ec.marshalNInt2int(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_Criticalityscore_legacyRecentReleaseCount(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Criticalityscore",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type String does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _Criticalityscore_legacyUpdatedIssuesCount(ctx context.Context, field graphql.CollectedField, obj *model.Criticalityscore) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Criticalityscore_legacyUpdatedIssuesCount(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.LegacyUpdatedIssuesCount, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(int)
	fc.Result = res
	return ec.marshalNInt2int(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_Criticalityscore_legacyUpdatedIssuesCount(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Criticalityscore",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type String does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _Criticalityscore_repoStarCount(ctx context.Context, field graphql.CollectedField, obj *model.Criticalityscore) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Criticalityscore_repoStarCount(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.RepoStarCount, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(int)
	fc.Result = res
	return ec.marshalNInt2int(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_Criticalityscore_repoStarCount(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Criticalityscore",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type String does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _Criticalityscore_origin(ctx context.Context, field graphql.CollectedField, obj *model.Criticalityscore) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Criticalityscore_origin(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Origin, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalNString2string(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_Criticalityscore_origin(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Criticalityscore",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type String does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _Criticalityscore_collector(ctx context.Context, field graphql.CollectedField, obj *model.Criticalityscore) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Criticalityscore_collector(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Collector, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalNString2string(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_Criticalityscore_collector(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Criticalityscore",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type String does not have child fields")
		},
	}
	return fc, nil
}

// endregion **************************** field.gotpl *****************************

// region    **************************** input.gotpl *****************************

func (ec *executionContext) unmarshalInputCertifyCriticalityscoreSpec(ctx context.Context, obj interface{}) (model.CertifyCriticalityscoreSpec, error) {
	var it model.CertifyCriticalityscoreSpec
	asMap := map[string]interface{}{}
	for k, v := range obj.(map[string]interface{}) {
		asMap[k] = v
	}

	fieldsInOrder := [...]string{"source", "timeScanned", "defaultScore", "legacyCommitFrequency", "legacyContributorCount", "legacyRecentReleaseCount", "legacyUpdatedIssuesCount", "repoStarCount", "origin", "collector"}
	for _, k := range fieldsInOrder {
		v, ok := asMap[k]
		if !ok {
			continue
		}
		switch k {
		case "source":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("source"))
			it.Source, err = ec.unmarshalOSourceSpec2ᚖgithubᚗcomᚋguacsecᚋguacᚋpkgᚋassemblerᚋgraphqlᚋmodelᚐSourceSpec(ctx, v)
			if err != nil {
				return it, err
			}
		case "timeScanned":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("timeScanned"))
			it.TimeScanned, err = ec.unmarshalOString2ᚖstring(ctx, v)
			if err != nil {
				return it, err
			}
		case "defaultScore":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("defaultScore"))
			it.DefaultScore, err = ec.unmarshalOString2ᚖstring(ctx, v)
			if err != nil {
				return it, err
			}
		case "legacyCommitFrequency":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("legacyCommitFrequency"))
			it.LegacyCommitFrequency, err = ec.unmarshalOFloat2ᚖfloat64(ctx, v)
			if err != nil {
				return it, err
			}
		case "legacyContributorCount":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("legacyContributorCount"))
			it.LegacyContributorCount, err = ec.unmarshalOInt2ᚖint(ctx, v)
			if err != nil {
				return it, err
			}
		case "legacyRecentReleaseCount":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("legacyRecentReleaseCount"))
			it.LegacyRecentReleaseCount, err = ec.unmarshalOInt2ᚖint(ctx, v)
			if err != nil {
				return it, err
			}
		case "legacyUpdatedIssuesCount":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("legacyUpdatedIssuesCount"))
			it.LegacyUpdatedIssuesCount, err = ec.unmarshalOInt2ᚖint(ctx, v)
			if err != nil {
				return it, err
			}
		case "repoStarCount":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("repoStarCount"))
			it.RepoStarCount, err = ec.unmarshalOInt2ᚖint(ctx, v)
			if err != nil {
				return it, err
			}
		case "origin":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("origin"))
			it.Origin, err = ec.unmarshalOString2ᚖstring(ctx, v)
			if err != nil {
				return it, err
			}
		case "collector":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("collector"))
			it.Collector, err = ec.unmarshalOString2ᚖstring(ctx, v)
			if err != nil {
				return it, err
			}
		}
	}

	return it, nil
}

// endregion **************************** input.gotpl *****************************

// region    ************************** interface.gotpl ***************************

// endregion ************************** interface.gotpl ***************************

// region    **************************** object.gotpl ****************************

var certifyCriticalityscoreImplementors = []string{"CertifyCriticalityscore"}

func (ec *executionContext) _CertifyCriticalityscore(ctx context.Context, sel ast.SelectionSet, obj *model.CertifyCriticalityscore) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, certifyCriticalityscoreImplementors)
	out := graphql.NewFieldSet(fields)
	var invalids uint32
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("CertifyCriticalityscore")
		case "source":

			out.Values[i] = ec._CertifyCriticalityscore_source(ctx, field, obj)

			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "criticalityscore":

			out.Values[i] = ec._CertifyCriticalityscore_criticalityscore(ctx, field, obj)

			if out.Values[i] == graphql.Null {
				invalids++
			}
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}
	out.Dispatch()
	if invalids > 0 {
		return graphql.Null
	}
	return out
}

var criticalityscoreImplementors = []string{"Criticalityscore"}

func (ec *executionContext) _Criticalityscore(ctx context.Context, sel ast.SelectionSet, obj *model.Criticalityscore) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, criticalityscoreImplementors)
	out := graphql.NewFieldSet(fields)
	var invalids uint32
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("Criticalityscore")
		case "defaultScore":

			out.Values[i] = ec._Criticalityscore_defaultScore(ctx, field, obj)

			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "timeScanned":

			out.Values[i] = ec._Criticalityscore_timeScanned(ctx, field, obj)

			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "legacyCommitFrequency":

			out.Values[i] = ec._Criticalityscore_legacyCommitFrequency(ctx, field, obj)

			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "legacyContributorCount":

			out.Values[i] = ec._Criticalityscore_legacyContributorCount(ctx, field, obj)

			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "legacyRecentReleaseCount":

			out.Values[i] = ec._Criticalityscore_legacyRecentReleaseCount(ctx, field, obj)

			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "legacyUpdatedIssuesCount":

			out.Values[i] = ec._Criticalityscore_legacyUpdatedIssuesCount(ctx, field, obj)

			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "repoStarCount":

			out.Values[i] = ec._Criticalityscore_repoStarCount(ctx, field, obj)

			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "origin":

			out.Values[i] = ec._Criticalityscore_origin(ctx, field, obj)

			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "collector":

			out.Values[i] = ec._Criticalityscore_collector(ctx, field, obj)

			if out.Values[i] == graphql.Null {
				invalids++
			}
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}
	out.Dispatch()
	if invalids > 0 {
		return graphql.Null
	}
	return out
}

// endregion **************************** object.gotpl ****************************

// region    ***************************** type.gotpl *****************************

func (ec *executionContext) marshalNCertifyCriticalityscore2ᚕᚖgithubᚗcomᚋguacsecᚋguacᚋpkgᚋassemblerᚋgraphqlᚋmodelᚐCertifyCriticalityscoreᚄ(ctx context.Context, sel ast.SelectionSet, v []*model.CertifyCriticalityscore) graphql.Marshaler {
	ret := make(graphql.Array, len(v))
	var wg sync.WaitGroup
	isLen1 := len(v) == 1
	if !isLen1 {
		wg.Add(len(v))
	}
	for i := range v {
		i := i
		fc := &graphql.FieldContext{
			Index:  &i,
			Result: &v[i],
		}
		ctx := graphql.WithFieldContext(ctx, fc)
		f := func(i int) {
			defer func() {
				if r := recover(); r != nil {
					ec.Error(ctx, ec.Recover(ctx, r))
					ret = nil
				}
			}()
			if !isLen1 {
				defer wg.Done()
			}
			ret[i] = ec.marshalNCertifyCriticalityscore2ᚖgithubᚗcomᚋguacsecᚋguacᚋpkgᚋassemblerᚋgraphqlᚋmodelᚐCertifyCriticalityscore(ctx, sel, v[i])
		}
		if isLen1 {
			f(i)
		} else {
			go f(i)
		}

	}
	wg.Wait()

	for _, e := range ret {
		if e == graphql.Null {
			return graphql.Null
		}
	}

	return ret
}

func (ec *executionContext) marshalNCertifyCriticalityscore2ᚖgithubᚗcomᚋguacsecᚋguacᚋpkgᚋassemblerᚋgraphqlᚋmodelᚐCertifyCriticalityscore(ctx context.Context, sel ast.SelectionSet, v *model.CertifyCriticalityscore) graphql.Marshaler {
	if v == nil {
		if !graphql.HasFieldError(ctx, graphql.GetFieldContext(ctx)) {
			ec.Errorf(ctx, "the requested element is null which the schema does not allow")
		}
		return graphql.Null
	}
	return ec._CertifyCriticalityscore(ctx, sel, v)
}

func (ec *executionContext) marshalNCriticalityscore2ᚖgithubᚗcomᚋguacsecᚋguacᚋpkgᚋassemblerᚋgraphqlᚋmodelᚐCriticalityscore(ctx context.Context, sel ast.SelectionSet, v *model.Criticalityscore) graphql.Marshaler {
	if v == nil {
		if !graphql.HasFieldError(ctx, graphql.GetFieldContext(ctx)) {
			ec.Errorf(ctx, "the requested element is null which the schema does not allow")
		}
		return graphql.Null
	}
	return ec._Criticalityscore(ctx, sel, v)
}

func (ec *executionContext) unmarshalOCertifyCriticalityscoreSpec2ᚖgithubᚗcomᚋguacsecᚋguacᚋpkgᚋassemblerᚋgraphqlᚋmodelᚐCertifyCriticalityscoreSpec(ctx context.Context, v interface{}) (*model.CertifyCriticalityscoreSpec, error) {
	if v == nil {
		return nil, nil
	}
	res, err := ec.unmarshalInputCertifyCriticalityscoreSpec(ctx, v)
	return &res, graphql.ErrorOnPath(ctx, err)
}

// endregion ***************************** type.gotpl *****************************
