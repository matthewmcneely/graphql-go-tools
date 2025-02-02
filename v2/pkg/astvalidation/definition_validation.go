package astvalidation

import (
	"github.com/matthewmcneely/graphql-go-tools/v2/pkg/ast"
	"github.com/matthewmcneely/graphql-go-tools/v2/pkg/astvisitor"
	"github.com/matthewmcneely/graphql-go-tools/v2/pkg/operationreport"
)

func DefaultDefinitionValidator() *DefinitionValidator {
	return NewDefinitionValidator(
		PopulatedTypeBodies(),
		UniqueOperationTypes(),
		UniqueTypeNames(),
		UniqueFieldDefinitionNames(),
		UniqueEnumValueNames(),
		UniqueUnionMemberTypes(),
		KnownTypeNames(),
		RequireDefinedTypesForExtensions(),
		ImplementTransitiveInterfaces(),
		ImplementingTypesAreSupersets(),
		DirectivesAreUniquePerLocation(),
	)
}

func NewDefinitionValidator(rules ...Rule) *DefinitionValidator {
	validator := &DefinitionValidator{
		walker: astvisitor.NewWalker(48),
	}

	for _, rule := range rules {
		validator.RegisterRule(rule)
	}

	return validator
}

type DefinitionValidator struct {
	walker astvisitor.Walker
}

func (d *DefinitionValidator) RegisterRule(rule Rule) {
	rule(&d.walker)
}

func (d *DefinitionValidator) Validate(definition *ast.Document, report *operationreport.Report) ValidationState {
	if report == nil {
		report = &operationreport.Report{}
	}

	d.walker.Walk(definition, definition, report)

	if report.HasErrors() {
		return Invalid
	}
	return Valid
}
