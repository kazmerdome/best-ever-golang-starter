// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package graph

import (
	"bytes"
	"context"
	"errors"
	"sync/atomic"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/introspection"
	"github.com/google/uuid"
	gqlparser "github.com/vektah/gqlparser/v2"
	"github.com/vektah/gqlparser/v2/ast"
	"gitlab.com/kazmerdome/best-ever-golang-starter/internal/module/category"
	"gitlab.com/kazmerdome/best-ever-golang-starter/internal/module/post"
)

// NewExecutableSchema creates an ExecutableSchema from the ResolverRoot interface.
func NewExecutableSchema(cfg Config) graphql.ExecutableSchema {
	return &executableSchema{
		schema:     cfg.Schema,
		resolvers:  cfg.Resolvers,
		directives: cfg.Directives,
		complexity: cfg.Complexity,
	}
}

type Config struct {
	Schema     *ast.Schema
	Resolvers  ResolverRoot
	Directives DirectiveRoot
	Complexity ComplexityRoot
}

type ResolverRoot interface {
	Mutation() MutationResolver
	Post() PostResolver
	Query() QueryResolver
}

type DirectiveRoot struct {
}

type ComplexityRoot struct {
	Category struct {
		CreatedAt func(childComplexity int) int
		Id        func(childComplexity int) int
		Name      func(childComplexity int) int
		Slug      func(childComplexity int) int
		Status    func(childComplexity int) int
		UpdatedAt func(childComplexity int) int
	}

	Mutation struct {
		CreateCategory func(childComplexity int, data category.CreateDto) int
		CreatePost     func(childComplexity int, data post.CreateDto) int
		DeleteCategory func(childComplexity int, id uuid.UUID) int
		DeletePost     func(childComplexity int, id uuid.UUID) int
		UpdateCategory func(childComplexity int, id uuid.UUID, data category.UpdateDto) int
		UpdatePost     func(childComplexity int, id uuid.UUID, data post.UpdateDto) int
	}

	Post struct {
		Category  func(childComplexity int) int
		Content   func(childComplexity int) int
		CreatedAt func(childComplexity int) int
		Id        func(childComplexity int) int
		Slug      func(childComplexity int) int
		Status    func(childComplexity int) int
		Title     func(childComplexity int) int
		UpdatedAt func(childComplexity int) int
	}

	Query struct {
		GetCategory        func(childComplexity int, id uuid.UUID) int
		GetPost            func(childComplexity int, id uuid.UUID) int
		ListCategories     func(childComplexity int, where *category.WhereDto) int
		ListPosts          func(childComplexity int, where *post.WhereDto) int
		__resolve__service func(childComplexity int) int
	}

	_Service struct {
		SDL func(childComplexity int) int
	}
}

type executableSchema struct {
	schema     *ast.Schema
	resolvers  ResolverRoot
	directives DirectiveRoot
	complexity ComplexityRoot
}

func (e *executableSchema) Schema() *ast.Schema {
	if e.schema != nil {
		return e.schema
	}
	return parsedSchema
}

func (e *executableSchema) Complexity(typeName, field string, childComplexity int, rawArgs map[string]interface{}) (int, bool) {
	ec := executionContext{nil, e, 0, 0, nil}
	_ = ec
	switch typeName + "." + field {

	case "Category.created_at":
		if e.complexity.Category.CreatedAt == nil {
			break
		}

		return e.complexity.Category.CreatedAt(childComplexity), true

	case "Category.id":
		if e.complexity.Category.Id == nil {
			break
		}

		return e.complexity.Category.Id(childComplexity), true

	case "Category.name":
		if e.complexity.Category.Name == nil {
			break
		}

		return e.complexity.Category.Name(childComplexity), true

	case "Category.slug":
		if e.complexity.Category.Slug == nil {
			break
		}

		return e.complexity.Category.Slug(childComplexity), true

	case "Category.status":
		if e.complexity.Category.Status == nil {
			break
		}

		return e.complexity.Category.Status(childComplexity), true

	case "Category.updated_at":
		if e.complexity.Category.UpdatedAt == nil {
			break
		}

		return e.complexity.Category.UpdatedAt(childComplexity), true

	case "Mutation.createCategory":
		if e.complexity.Mutation.CreateCategory == nil {
			break
		}

		args, err := ec.field_Mutation_createCategory_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Mutation.CreateCategory(childComplexity, args["data"].(category.CreateDto)), true

	case "Mutation.createPost":
		if e.complexity.Mutation.CreatePost == nil {
			break
		}

		args, err := ec.field_Mutation_createPost_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Mutation.CreatePost(childComplexity, args["data"].(post.CreateDto)), true

	case "Mutation.deleteCategory":
		if e.complexity.Mutation.DeleteCategory == nil {
			break
		}

		args, err := ec.field_Mutation_deleteCategory_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Mutation.DeleteCategory(childComplexity, args["id"].(uuid.UUID)), true

	case "Mutation.deletePost":
		if e.complexity.Mutation.DeletePost == nil {
			break
		}

		args, err := ec.field_Mutation_deletePost_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Mutation.DeletePost(childComplexity, args["id"].(uuid.UUID)), true

	case "Mutation.updateCategory":
		if e.complexity.Mutation.UpdateCategory == nil {
			break
		}

		args, err := ec.field_Mutation_updateCategory_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Mutation.UpdateCategory(childComplexity, args["id"].(uuid.UUID), args["data"].(category.UpdateDto)), true

	case "Mutation.updatePost":
		if e.complexity.Mutation.UpdatePost == nil {
			break
		}

		args, err := ec.field_Mutation_updatePost_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Mutation.UpdatePost(childComplexity, args["id"].(uuid.UUID), args["data"].(post.UpdateDto)), true

	case "Post.category":
		if e.complexity.Post.Category == nil {
			break
		}

		return e.complexity.Post.Category(childComplexity), true

	case "Post.content":
		if e.complexity.Post.Content == nil {
			break
		}

		return e.complexity.Post.Content(childComplexity), true

	case "Post.created_at":
		if e.complexity.Post.CreatedAt == nil {
			break
		}

		return e.complexity.Post.CreatedAt(childComplexity), true

	case "Post.id":
		if e.complexity.Post.Id == nil {
			break
		}

		return e.complexity.Post.Id(childComplexity), true

	case "Post.slug":
		if e.complexity.Post.Slug == nil {
			break
		}

		return e.complexity.Post.Slug(childComplexity), true

	case "Post.status":
		if e.complexity.Post.Status == nil {
			break
		}

		return e.complexity.Post.Status(childComplexity), true

	case "Post.title":
		if e.complexity.Post.Title == nil {
			break
		}

		return e.complexity.Post.Title(childComplexity), true

	case "Post.updated_at":
		if e.complexity.Post.UpdatedAt == nil {
			break
		}

		return e.complexity.Post.UpdatedAt(childComplexity), true

	case "Query.getCategory":
		if e.complexity.Query.GetCategory == nil {
			break
		}

		args, err := ec.field_Query_getCategory_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Query.GetCategory(childComplexity, args["id"].(uuid.UUID)), true

	case "Query.getPost":
		if e.complexity.Query.GetPost == nil {
			break
		}

		args, err := ec.field_Query_getPost_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Query.GetPost(childComplexity, args["id"].(uuid.UUID)), true

	case "Query.listCategories":
		if e.complexity.Query.ListCategories == nil {
			break
		}

		args, err := ec.field_Query_listCategories_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Query.ListCategories(childComplexity, args["where"].(*category.WhereDto)), true

	case "Query.listPosts":
		if e.complexity.Query.ListPosts == nil {
			break
		}

		args, err := ec.field_Query_listPosts_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Query.ListPosts(childComplexity, args["where"].(*post.WhereDto)), true

	case "Query._service":
		if e.complexity.Query.__resolve__service == nil {
			break
		}

		return e.complexity.Query.__resolve__service(childComplexity), true

	case "_Service.sdl":
		if e.complexity._Service.SDL == nil {
			break
		}

		return e.complexity._Service.SDL(childComplexity), true

	}
	return 0, false
}

func (e *executableSchema) Exec(ctx context.Context) graphql.ResponseHandler {
	rc := graphql.GetOperationContext(ctx)
	ec := executionContext{rc, e, 0, 0, make(chan graphql.DeferredResult)}
	inputUnmarshalMap := graphql.BuildUnmarshalerMap(
		ec.unmarshalInputCreateCategoryDto,
		ec.unmarshalInputCreatePostDto,
		ec.unmarshalInputDateFilter,
		ec.unmarshalInputFloat64Filter,
		ec.unmarshalInputIntFilter,
		ec.unmarshalInputPaginationFilter,
		ec.unmarshalInputSortFilter,
		ec.unmarshalInputStringFilter,
		ec.unmarshalInputUpdateCategoryDto,
		ec.unmarshalInputUpdatePostDto,
		ec.unmarshalInputUuidFilter,
		ec.unmarshalInputWhereCategoriesDto,
		ec.unmarshalInputWherePostsDto,
	)
	first := true

	switch rc.Operation.Operation {
	case ast.Query:
		return func(ctx context.Context) *graphql.Response {
			var response graphql.Response
			var data graphql.Marshaler
			if first {
				first = false
				ctx = graphql.WithUnmarshalerMap(ctx, inputUnmarshalMap)
				data = ec._Query(ctx, rc.Operation.SelectionSet)
			} else {
				if atomic.LoadInt32(&ec.pendingDeferred) > 0 {
					result := <-ec.deferredResults
					atomic.AddInt32(&ec.pendingDeferred, -1)
					data = result.Result
					response.Path = result.Path
					response.Label = result.Label
					response.Errors = result.Errors
				} else {
					return nil
				}
			}
			var buf bytes.Buffer
			data.MarshalGQL(&buf)
			response.Data = buf.Bytes()
			if atomic.LoadInt32(&ec.deferred) > 0 {
				hasNext := atomic.LoadInt32(&ec.pendingDeferred) > 0
				response.HasNext = &hasNext
			}

			return &response
		}
	case ast.Mutation:
		return func(ctx context.Context) *graphql.Response {
			if !first {
				return nil
			}
			first = false
			ctx = graphql.WithUnmarshalerMap(ctx, inputUnmarshalMap)
			data := ec._Mutation(ctx, rc.Operation.SelectionSet)
			var buf bytes.Buffer
			data.MarshalGQL(&buf)

			return &graphql.Response{
				Data: buf.Bytes(),
			}
		}

	default:
		return graphql.OneShot(graphql.ErrorResponse(ctx, "unsupported GraphQL operation"))
	}
}

type executionContext struct {
	*graphql.OperationContext
	*executableSchema
	deferred        int32
	pendingDeferred int32
	deferredResults chan graphql.DeferredResult
}

func (ec *executionContext) processDeferredGroup(dg graphql.DeferredGroup) {
	atomic.AddInt32(&ec.pendingDeferred, 1)
	go func() {
		ctx := graphql.WithFreshResponseContext(dg.Context)
		dg.FieldSet.Dispatch(ctx)
		ds := graphql.DeferredResult{
			Path:   dg.Path,
			Label:  dg.Label,
			Result: dg.FieldSet,
			Errors: graphql.GetErrors(ctx),
		}
		// null fields should bubble up
		if dg.FieldSet.Invalids > 0 {
			ds.Result = graphql.Null
		}
		ec.deferredResults <- ds
	}()
}

func (ec *executionContext) introspectSchema() (*introspection.Schema, error) {
	if ec.DisableIntrospection {
		return nil, errors.New("introspection disabled")
	}
	return introspection.WrapSchema(ec.Schema()), nil
}

func (ec *executionContext) introspectType(name string) (*introspection.Type, error) {
	if ec.DisableIntrospection {
		return nil, errors.New("introspection disabled")
	}
	return introspection.WrapTypeFromDef(ec.Schema(), ec.Schema().Types[name]), nil
}

var sources = []*ast.Source{
	{Name: "../../../internal/util/graph/schema.graphql", Input: `# Directives
# 

directive @goModel(
	model: String
	models: [String!]
) on OBJECT | INPUT_OBJECT | SCALAR | ENUM | INTERFACE | UNION

directive @goField(
	forceResolver: Boolean
	name: String
  omittable: Boolean
) on INPUT_FIELD_DEFINITION | FIELD_DEFINITION

directive @goTag(
	key: String!
	value: String
) on INPUT_FIELD_DEFINITION | FIELD_DEFINITION

# Scalars
#
scalar Time
scalar Uuid @goModel(model: "gitlab.com/kazmerdome/best-ever-golang-starter/internal/util/graph/scalar.Uuid")
scalar Json @goModel(model: "gitlab.com/kazmerdome/best-ever-golang-starter/internal/util/graph/scalar.Json")

# Inputs
# 
input IntFilter @goModel(model: "gitlab.com/kazmerdome/best-ever-golang-starter/internal/util/filter.IntFilter") {
	"Specifies equality condition. The $eq operator matches documents where the value of a field equals the specified value."
  eq: Int
	"$gt selects those documents where the value of the field is greater than (i.e. >) the specified value."
	gt: Int
	"$gte selects the documents where the value of the field is greater than or equal to (i.e. >=) a specified value (e.g. value.)"
	gte: Int
	"$lt selects the documents where the value of the field is less than (i.e. <) the specified value."
	lt: Int
	"$lte selects the documents where the value of the field is less than or equal to (i.e. <=) the specified value."
	lte: Int
}

input Float64Filter @goModel(model: "gitlab.com/kazmerdome/best-ever-golang-starter/internal/util/filter.Float64Filter") {
	"Specifies equality condition. The $eq operator matches documents where the value of a field equals the specified value."
  eq: Float
	"$gt selects those documents where the value of the field is greater than (i.e. >) the specified value."
	gt: Float
	"$gte selects the documents where the value of the field is greater than or equal to (i.e. >=) a specified value (e.g. value.)"
	gte: Float
	"$lt selects the documents where the value of the field is less than (i.e. <) the specified value."
	lt: Float
	"$lte selects the documents where the value of the field is less than or equal to (i.e. <=) the specified value."
	lte: Float
}

input DateFilter @goModel(model: "gitlab.com/kazmerdome/best-ever-golang-starter/internal/util/filter.DateFilter") {
	"Specifies equality condition. The $eq operator matches documents where the value of a field equals the specified value."
  eq: Time
	"$gt selects those documents where the value of the field is greater than (i.e. >) the specified value."
	gt: Time
	"$gte selects the documents where the value of the field is greater than or equal to (i.e. >=) a specified value (e.g. value.)"
	gte: Time
	"$lt selects the documents where the value of the field is less than (i.e. <) the specified value."
	lt: Time
	"$lte selects the documents where the value of the field is less than or equal to (i.e. <=) the specified value."
	lte: Time
}

input StringFilter @goModel(model: "gitlab.com/kazmerdome/best-ever-golang-starter/internal/util/filter.StringFilter") {
  "Specifies equality condition. The $eq operator matches documents where the value of a field equals the specified value."
  eq: String
  "Specifies a regular expression pattern for matching strings."
  regex: String
}

input UuidFilter @goModel(model: "gitlab.com/kazmerdome/best-ever-golang-starter/internal/util/filter.UuidFilter") {
  "Specifies equality condition. The $eq operator matches documents where the value of a field equals the specified value."
  eq: Uuid
  "Specifies a list of UUID values. The $in operator matches documents where the value of a field equals any value in the specified list."
  in: [Uuid!]
}

input SortFilter @goModel(model: "gitlab.com/kazmerdome/best-ever-golang-starter/internal/util/filter.SortFilter") {
  sortBy: String
	sortOrder: SortOrder
}

input PaginationFilter @goModel(model: "gitlab.com/kazmerdome/best-ever-golang-starter/internal/util/filter.PaginationFilter") {
  limit: Int
	skip: Int
}

# Enums
# 
enum SortOrder @goModel(model: "gitlab.com/kazmerdome/best-ever-golang-starter/internal/util/filter.SortOrder") {
  asc
	desc
}
`, BuiltIn: false},
	{Name: "../../../internal/module/category/schema.graphql", Input: `# Models (types) & Enums
# 
type Category @goModel(model: "gitlab.com/kazmerdome/best-ever-golang-starter/internal/module/category.Category") {
  id: Uuid!
  name: String!
  slug: String
  status: CategoryStatusEnum!
  created_at: Time!
  updated_at: Time!
}

enum CategoryStatusEnum @goModel(model: "gitlab.com/kazmerdome/best-ever-golang-starter/internal/module/category.CategoryStatus") {
  ACTIVE
  PENDING
  ARCHIVED
}

# Dto (inputs)
# 
input CreateCategoryDto @goModel(model: "gitlab.com/kazmerdome/best-ever-golang-starter/internal/module/category.CreateDto") {
  name: String!
  slug: String
  status: CategoryStatusEnum
}

input WhereCategoriesDto @goModel(model: "gitlab.com/kazmerdome/best-ever-golang-starter/internal/module/category.WhereDto") {
  name: StringFilter
  slug: StringFilter
  status: CategoryStatusEnum
  sort: SortFilter
  pagination: PaginationFilter
}

input UpdateCategoryDto @goModel(model: "gitlab.com/kazmerdome/best-ever-golang-starter/internal/module/category.UpdateDto") {
  name: String
  slug: String
  status: CategoryStatusEnum
}

# Queries and Mutations
# 
extend type Query {
  getCategory(id: Uuid!): Category
  listCategories(where: WhereCategoriesDto): [Category!]!
}

extend type Mutation {
  createCategory(data: CreateCategoryDto!): Category             
  updateCategory(id: Uuid!, data: UpdateCategoryDto!): Category  
  deleteCategory(id: Uuid!): Boolean                          
}
`, BuiltIn: false},
	{Name: "../../../internal/module/post/schema.graphql", Input: `# Models (types) & Enums
# 
type Post @goModel(model: "gitlab.com/kazmerdome/best-ever-golang-starter/internal/module/post.Post") {
  id: Uuid!
  title: String!
  slug: String!
  category: Category!
  status: PostStatusEnum!
  content: String
  created_at: Time!
  updated_at: Time!
}

enum PostStatusEnum @goModel(model: "gitlab.com/kazmerdome/best-ever-golang-starter/internal/module/post.PostStatus") {
  ACTIVE
  PENDING
  ARCHIVED
}

# Dto (inputs)
# 
input CreatePostDto @goModel(model: "gitlab.com/kazmerdome/best-ever-golang-starter/internal/module/post.CreateDto") {
  title: String!
  slug: String
  status: PostStatusEnum!
  category: Uuid!
  content: String
}

input WherePostsDto @goModel(model: "gitlab.com/kazmerdome/best-ever-golang-starter/internal/module/post.WhereDto") {
  title: StringFilter
  slug: StringFilter
  category: UuidFilter
  status: PostStatusEnum
  sort: SortFilter
  pagination: PaginationFilter
}

input UpdatePostDto @goModel(model: "gitlab.com/kazmerdome/best-ever-golang-starter/internal/module/post.UpdateDto") {
  title: String
  slug: String
  category: Uuid
  status: PostStatusEnum
  content: String
}


# Queries and Mutations
# 
extend type Query {
  getPost(id: Uuid!): Post
  listPosts(where: WherePostsDto): [Post!]!
}

extend type Mutation {
  createPost(data: CreatePostDto!): Post             
  updatePost(id: Uuid!, data: UpdatePostDto!): Post  
  deletePost(id: Uuid!): Boolean                          
}
`, BuiltIn: false},
	{Name: "../../../federation/directives.graphql", Input: `
	directive @key(fields: _FieldSet!) repeatable on OBJECT | INTERFACE
	directive @requires(fields: _FieldSet!) on FIELD_DEFINITION
	directive @provides(fields: _FieldSet!) on FIELD_DEFINITION
	directive @extends on OBJECT | INTERFACE
	directive @external on FIELD_DEFINITION
	scalar _Any
	scalar _FieldSet
`, BuiltIn: true},
	{Name: "../../../federation/entity.graphql", Input: `
type _Service {
  sdl: String
}

extend type Query {
  _service: _Service!
}
`, BuiltIn: true},
}
var parsedSchema = gqlparser.MustLoadSchema(sources...)
