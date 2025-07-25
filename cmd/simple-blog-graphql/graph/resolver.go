package graph

import (
	"github.com/kazmerdome/best-ever-golang-starter/internal/module/category"
	"github.com/kazmerdome/best-ever-golang-starter/internal/module/post"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	CategoryService category.CategoryService
	PostService     post.PostService
}
