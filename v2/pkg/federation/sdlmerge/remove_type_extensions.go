package sdlmerge

import (
	"github.com/matthewmcneely/graphql-go-tools/v2/pkg/ast"
	"github.com/matthewmcneely/graphql-go-tools/v2/pkg/astvisitor"
)

func newRemoveMergedTypeExtensions() *removeMergedTypeExtensionsVisitor {
	return &removeMergedTypeExtensionsVisitor{}
}

type removeMergedTypeExtensionsVisitor struct {
}

func (r *removeMergedTypeExtensionsVisitor) Register(walker *astvisitor.Walker) {
	walker.RegisterLeaveDocumentVisitor(r)
}

func (r *removeMergedTypeExtensionsVisitor) LeaveDocument(operation, definition *ast.Document) {
	operation.RemoveMergedTypeExtensions()
}
