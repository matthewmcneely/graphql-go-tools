package astnormalization

import (
	"github.com/matthewmcneely/graphql-go-tools/v2/pkg/ast"
	"github.com/matthewmcneely/graphql-go-tools/v2/pkg/astvisitor"
)

type FragmentDefinitionRemoval struct {
}

func removeFragmentDefinitions(walker *astvisitor.Walker) {
	visitor := removeFragmentDefinitionsVisitor{}
	walker.RegisterLeaveDocumentVisitor(visitor)
}

type removeFragmentDefinitionsVisitor struct {
}

func (r removeFragmentDefinitionsVisitor) LeaveDocument(operation, definition *ast.Document) {
	for i := range operation.RootNodes {
		if operation.RootNodes[i].Kind == ast.NodeKindFragmentDefinition {
			operation.RootNodes[i].Kind = ast.NodeKindUnknown
		}
	}
}
