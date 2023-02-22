// Copyright 2019-present Facebook
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"

	"entgo.io/contrib/entgql"
	"entgo.io/contrib/entgql/internal/todo/ent/category"
	"entgo.io/contrib/entgql/internal/todo/ent/group"
	"entgo.io/contrib/entgql/internal/todo/ent/todo"
	"entgo.io/contrib/entgql/internal/todo/ent/user"
	"entgo.io/ent/dialect/sql"
	"github.com/99designs/gqlgen/graphql"
)

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (bp *BillProductQuery) CollectFields(ctx context.Context, satisfies ...string) (*BillProductQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return bp, nil
	}
	if err := bp.collectField(ctx, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return bp, nil
}

func (bp *BillProductQuery) collectField(ctx context.Context, op *graphql.OperationContext, field graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	return nil
}

type billproductPaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []BillProductPaginateOption
}

func newBillProductPaginateArgs(rv map[string]interface{}) *billproductPaginateArgs {
	args := &billproductPaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	if v, ok := rv[whereField].(*BillProductWhereInput); ok {
		args.opts = append(args.opts, WithBillProductFilter(v.Filter))
	}
	return args
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (c *CategoryQuery) CollectFields(ctx context.Context, satisfies ...string) (*CategoryQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return c, nil
	}
	if err := c.collectField(ctx, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return c, nil
}

func (c *CategoryQuery) collectField(ctx context.Context, op *graphql.OperationContext, field graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	for _, field := range graphql.CollectFields(op, field.Selections, satisfies) {
		switch field.Name {
		case "todos":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&TodoClient{config: c.config}).Query()
			)
			args := newTodoPaginateArgs(fieldArgs(ctx, new(TodoWhereInput), path...))
			if err := validateFirstLast(args.first, args.last); err != nil {
				return fmt.Errorf("validate first and last in path %q: %w", path, err)
			}
			pager, err := newTodoPager(args.opts, args.last != nil)
			if err != nil {
				return fmt.Errorf("create new pager in path %q: %w", path, err)
			}
			if query, err = pager.applyFilter(query); err != nil {
				return err
			}
			ignoredEdges := !hasCollectedField(ctx, append(path, edgesField)...)
			if hasCollectedField(ctx, append(path, totalCountField)...) || hasCollectedField(ctx, append(path, pageInfoField)...) {
				hasPagination := args.after != nil || args.first != nil || args.before != nil || args.last != nil
				if hasPagination || ignoredEdges {
					query := query.Clone()
					c.loadTotal = append(c.loadTotal, func(ctx context.Context, nodes []*Category) error {
						ids := make([]driver.Value, len(nodes))
						for i := range nodes {
							ids[i] = nodes[i].ID
						}
						var v []struct {
							NodeID int `sql:"category_id"`
							Count  int `sql:"count"`
						}
						query.Where(func(s *sql.Selector) {
							s.Where(sql.InValues(category.TodosColumn, ids...))
						})
						if err := query.GroupBy(category.TodosColumn).Aggregate(Count()).Scan(ctx, &v); err != nil {
							return err
						}
						m := make(map[int]int, len(v))
						for i := range v {
							m[v[i].NodeID] = v[i].Count
						}
						for i := range nodes {
							n := m[nodes[i].ID]
							if nodes[i].Edges.totalCount[0] == nil {
								nodes[i].Edges.totalCount[0] = make(map[string]int)
							}
							nodes[i].Edges.totalCount[0][alias] = n
						}
						return nil
					})
				} else {
					c.loadTotal = append(c.loadTotal, func(_ context.Context, nodes []*Category) error {
						for i := range nodes {
							n := len(nodes[i].Edges.Todos)
							if nodes[i].Edges.totalCount[0] == nil {
								nodes[i].Edges.totalCount[0] = make(map[string]int)
							}
							nodes[i].Edges.totalCount[0][alias] = n
						}
						return nil
					})
				}
			}
			if ignoredEdges || (args.first != nil && *args.first == 0) || (args.last != nil && *args.last == 0) {
				continue
			}

			if query, err = pager.applyCursors(query, args.after, args.before); err != nil {
				return err
			}
			if limit := paginateLimit(args.first, args.last); limit > 0 {
				modify := limitRows(category.TodosColumn, limit, pager.orderExpr())
				query.modifiers = append(query.modifiers, modify)
			} else {
				query = pager.applyOrder(query)
			}
			path = append(path, edgesField, nodeField)
			if field := collectedField(ctx, path...); field != nil {
				if err := query.collectField(ctx, op, *field, path, satisfies...); err != nil {
					return err
				}
			}
			c.WithNamedTodos(alias, func(wq *TodoQuery) {
				*wq = *query
			})
		case "subCategories":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&CategoryClient{config: c.config}).Query()
			)
			args := newCategoryPaginateArgs(fieldArgs(ctx, new(CategoryWhereInput), path...))
			if err := validateFirstLast(args.first, args.last); err != nil {
				return fmt.Errorf("validate first and last in path %q: %w", path, err)
			}
			pager, err := newCategoryPager(args.opts, args.last != nil)
			if err != nil {
				return fmt.Errorf("create new pager in path %q: %w", path, err)
			}
			if query, err = pager.applyFilter(query); err != nil {
				return err
			}
			ignoredEdges := !hasCollectedField(ctx, append(path, edgesField)...)
			if hasCollectedField(ctx, append(path, totalCountField)...) || hasCollectedField(ctx, append(path, pageInfoField)...) {
				hasPagination := args.after != nil || args.first != nil || args.before != nil || args.last != nil
				if hasPagination || ignoredEdges {
					query := query.Clone()
					c.loadTotal = append(c.loadTotal, func(ctx context.Context, nodes []*Category) error {
						ids := make([]driver.Value, len(nodes))
						for i := range nodes {
							ids[i] = nodes[i].ID
						}
						var v []struct {
							NodeID int `sql:"category_id"`
							Count  int `sql:"count"`
						}
						query.Where(func(s *sql.Selector) {
							joinT := sql.Table(category.SubCategoriesTable)
							s.Join(joinT).On(s.C(category.FieldID), joinT.C(category.SubCategoriesPrimaryKey[1]))
							s.Where(sql.InValues(joinT.C(category.SubCategoriesPrimaryKey[0]), ids...))
							s.Select(joinT.C(category.SubCategoriesPrimaryKey[0]), sql.Count("*"))
							s.GroupBy(joinT.C(category.SubCategoriesPrimaryKey[0]))
						})
						if err := query.Select().Scan(ctx, &v); err != nil {
							return err
						}
						m := make(map[int]int, len(v))
						for i := range v {
							m[v[i].NodeID] = v[i].Count
						}
						for i := range nodes {
							n := m[nodes[i].ID]
							if nodes[i].Edges.totalCount[1] == nil {
								nodes[i].Edges.totalCount[1] = make(map[string]int)
							}
							nodes[i].Edges.totalCount[1][alias] = n
						}
						return nil
					})
				} else {
					c.loadTotal = append(c.loadTotal, func(_ context.Context, nodes []*Category) error {
						for i := range nodes {
							n := len(nodes[i].Edges.SubCategories)
							if nodes[i].Edges.totalCount[1] == nil {
								nodes[i].Edges.totalCount[1] = make(map[string]int)
							}
							nodes[i].Edges.totalCount[1][alias] = n
						}
						return nil
					})
				}
			}
			if ignoredEdges || (args.first != nil && *args.first == 0) || (args.last != nil && *args.last == 0) {
				continue
			}

			if query, err = pager.applyCursors(query, args.after, args.before); err != nil {
				return err
			}
			if limit := paginateLimit(args.first, args.last); limit > 0 {
				modify := limitRows(category.SubCategoriesPrimaryKey[0], limit, pager.orderExpr())
				query.modifiers = append(query.modifiers, modify)
			} else {
				query = pager.applyOrder(query)
			}
			path = append(path, edgesField, nodeField)
			if field := collectedField(ctx, path...); field != nil {
				if err := query.collectField(ctx, op, *field, path, satisfies...); err != nil {
					return err
				}
			}
			c.WithNamedSubCategories(alias, func(wq *CategoryQuery) {
				*wq = *query
			})
		}
	}
	return nil
}

type categoryPaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []CategoryPaginateOption
}

func newCategoryPaginateArgs(rv map[string]interface{}) *categoryPaginateArgs {
	args := &categoryPaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	if v, ok := rv[orderByField]; ok {
		switch v := v.(type) {
		case []*CategoryOrder:
			args.opts = append(args.opts, WithCategoryOrder(v))
		case []interface{}:
			var orders []*CategoryOrder
			for i := range v {
				mv, ok := v[i].(map[string]interface{})
				if !ok {
					continue
				}
				var (
					err1, err2 error
					order      = &CategoryOrder{Field: &CategoryOrderField{}, Direction: entgql.OrderDirectionAsc}
				)
				if d, ok := mv[directionField]; ok {
					err1 = order.Direction.UnmarshalGQL(d)
				}
				if f, ok := mv[fieldField]; ok {
					err2 = order.Field.UnmarshalGQL(f)
				}
				if err1 == nil && err2 == nil {
					orders = append(orders, order)
				}
			}
			args.opts = append(args.opts, WithCategoryOrder(orders))
		}
	}
	if v, ok := rv[whereField].(*CategoryWhereInput); ok {
		args.opts = append(args.opts, WithCategoryFilter(v.Filter))
	}
	return args
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (f *FriendshipQuery) CollectFields(ctx context.Context, satisfies ...string) (*FriendshipQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return f, nil
	}
	if err := f.collectField(ctx, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return f, nil
}

func (f *FriendshipQuery) collectField(ctx context.Context, op *graphql.OperationContext, field graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	for _, field := range graphql.CollectFields(op, field.Selections, satisfies) {
		switch field.Name {
		case "user":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&UserClient{config: f.config}).Query()
			)
			if err := query.collectField(ctx, op, field, path, satisfies...); err != nil {
				return err
			}
			f.withUser = query
		case "friend":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&UserClient{config: f.config}).Query()
			)
			if err := query.collectField(ctx, op, field, path, satisfies...); err != nil {
				return err
			}
			f.withFriend = query
		}
	}
	return nil
}

type friendshipPaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []FriendshipPaginateOption
}

func newFriendshipPaginateArgs(rv map[string]interface{}) *friendshipPaginateArgs {
	args := &friendshipPaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	if v, ok := rv[whereField].(*FriendshipWhereInput); ok {
		args.opts = append(args.opts, WithFriendshipFilter(v.Filter))
	}
	return args
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (gr *GroupQuery) CollectFields(ctx context.Context, satisfies ...string) (*GroupQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return gr, nil
	}
	if err := gr.collectField(ctx, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return gr, nil
}

func (gr *GroupQuery) collectField(ctx context.Context, op *graphql.OperationContext, field graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	for _, field := range graphql.CollectFields(op, field.Selections, satisfies) {
		switch field.Name {
		case "users":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&UserClient{config: gr.config}).Query()
			)
			args := newUserPaginateArgs(fieldArgs(ctx, new(UserWhereInput), path...))
			if err := validateFirstLast(args.first, args.last); err != nil {
				return fmt.Errorf("validate first and last in path %q: %w", path, err)
			}
			pager, err := newUserPager(args.opts, args.last != nil)
			if err != nil {
				return fmt.Errorf("create new pager in path %q: %w", path, err)
			}
			if query, err = pager.applyFilter(query); err != nil {
				return err
			}
			ignoredEdges := !hasCollectedField(ctx, append(path, edgesField)...)
			if hasCollectedField(ctx, append(path, totalCountField)...) || hasCollectedField(ctx, append(path, pageInfoField)...) {
				hasPagination := args.after != nil || args.first != nil || args.before != nil || args.last != nil
				if hasPagination || ignoredEdges {
					query := query.Clone()
					gr.loadTotal = append(gr.loadTotal, func(ctx context.Context, nodes []*Group) error {
						ids := make([]driver.Value, len(nodes))
						for i := range nodes {
							ids[i] = nodes[i].ID
						}
						var v []struct {
							NodeID int `sql:"group_id"`
							Count  int `sql:"count"`
						}
						query.Where(func(s *sql.Selector) {
							joinT := sql.Table(group.UsersTable)
							s.Join(joinT).On(s.C(user.FieldID), joinT.C(group.UsersPrimaryKey[0]))
							s.Where(sql.InValues(joinT.C(group.UsersPrimaryKey[1]), ids...))
							s.Select(joinT.C(group.UsersPrimaryKey[1]), sql.Count("*"))
							s.GroupBy(joinT.C(group.UsersPrimaryKey[1]))
						})
						if err := query.Select().Scan(ctx, &v); err != nil {
							return err
						}
						m := make(map[int]int, len(v))
						for i := range v {
							m[v[i].NodeID] = v[i].Count
						}
						for i := range nodes {
							n := m[nodes[i].ID]
							if nodes[i].Edges.totalCount[0] == nil {
								nodes[i].Edges.totalCount[0] = make(map[string]int)
							}
							nodes[i].Edges.totalCount[0][alias] = n
						}
						return nil
					})
				} else {
					gr.loadTotal = append(gr.loadTotal, func(_ context.Context, nodes []*Group) error {
						for i := range nodes {
							n := len(nodes[i].Edges.Users)
							if nodes[i].Edges.totalCount[0] == nil {
								nodes[i].Edges.totalCount[0] = make(map[string]int)
							}
							nodes[i].Edges.totalCount[0][alias] = n
						}
						return nil
					})
				}
			}
			if ignoredEdges || (args.first != nil && *args.first == 0) || (args.last != nil && *args.last == 0) {
				continue
			}

			if query, err = pager.applyCursors(query, args.after, args.before); err != nil {
				return err
			}
			if limit := paginateLimit(args.first, args.last); limit > 0 {
				modify := limitRows(group.UsersPrimaryKey[1], limit, pager.orderExpr())
				query.modifiers = append(query.modifiers, modify)
			} else {
				query = pager.applyOrder(query)
			}
			path = append(path, edgesField, nodeField)
			if field := collectedField(ctx, path...); field != nil {
				if err := query.collectField(ctx, op, *field, path, satisfies...); err != nil {
					return err
				}
			}
			gr.WithNamedUsers(alias, func(wq *UserQuery) {
				*wq = *query
			})
		}
	}
	return nil
}

type groupPaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []GroupPaginateOption
}

func newGroupPaginateArgs(rv map[string]interface{}) *groupPaginateArgs {
	args := &groupPaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	if v, ok := rv[whereField].(*GroupWhereInput); ok {
		args.opts = append(args.opts, WithGroupFilter(v.Filter))
	}
	return args
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (t *TodoQuery) CollectFields(ctx context.Context, satisfies ...string) (*TodoQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return t, nil
	}
	if err := t.collectField(ctx, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return t, nil
}

func (t *TodoQuery) collectField(ctx context.Context, op *graphql.OperationContext, field graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	for _, field := range graphql.CollectFields(op, field.Selections, satisfies) {
		switch field.Name {
		case "parent":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&TodoClient{config: t.config}).Query()
			)
			if err := query.collectField(ctx, op, field, path, satisfies...); err != nil {
				return err
			}
			t.withParent = query
		case "children":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&TodoClient{config: t.config}).Query()
			)
			args := newTodoPaginateArgs(fieldArgs(ctx, new(TodoWhereInput), path...))
			if err := validateFirstLast(args.first, args.last); err != nil {
				return fmt.Errorf("validate first and last in path %q: %w", path, err)
			}
			pager, err := newTodoPager(args.opts, args.last != nil)
			if err != nil {
				return fmt.Errorf("create new pager in path %q: %w", path, err)
			}
			if query, err = pager.applyFilter(query); err != nil {
				return err
			}
			ignoredEdges := !hasCollectedField(ctx, append(path, edgesField)...)
			if hasCollectedField(ctx, append(path, totalCountField)...) || hasCollectedField(ctx, append(path, pageInfoField)...) {
				hasPagination := args.after != nil || args.first != nil || args.before != nil || args.last != nil
				if hasPagination || ignoredEdges {
					query := query.Clone()
					t.loadTotal = append(t.loadTotal, func(ctx context.Context, nodes []*Todo) error {
						ids := make([]driver.Value, len(nodes))
						for i := range nodes {
							ids[i] = nodes[i].ID
						}
						var v []struct {
							NodeID int `sql:"todo_children"`
							Count  int `sql:"count"`
						}
						query.Where(func(s *sql.Selector) {
							s.Where(sql.InValues(todo.ChildrenColumn, ids...))
						})
						if err := query.GroupBy(todo.ChildrenColumn).Aggregate(Count()).Scan(ctx, &v); err != nil {
							return err
						}
						m := make(map[int]int, len(v))
						for i := range v {
							m[v[i].NodeID] = v[i].Count
						}
						for i := range nodes {
							n := m[nodes[i].ID]
							if nodes[i].Edges.totalCount[1] == nil {
								nodes[i].Edges.totalCount[1] = make(map[string]int)
							}
							nodes[i].Edges.totalCount[1][alias] = n
						}
						return nil
					})
				} else {
					t.loadTotal = append(t.loadTotal, func(_ context.Context, nodes []*Todo) error {
						for i := range nodes {
							n := len(nodes[i].Edges.Children)
							if nodes[i].Edges.totalCount[1] == nil {
								nodes[i].Edges.totalCount[1] = make(map[string]int)
							}
							nodes[i].Edges.totalCount[1][alias] = n
						}
						return nil
					})
				}
			}
			if ignoredEdges || (args.first != nil && *args.first == 0) || (args.last != nil && *args.last == 0) {
				continue
			}

			if query, err = pager.applyCursors(query, args.after, args.before); err != nil {
				return err
			}
			if limit := paginateLimit(args.first, args.last); limit > 0 {
				modify := limitRows(todo.ChildrenColumn, limit, pager.orderExpr())
				query.modifiers = append(query.modifiers, modify)
			} else {
				query = pager.applyOrder(query)
			}
			path = append(path, edgesField, nodeField)
			if field := collectedField(ctx, path...); field != nil {
				if err := query.collectField(ctx, op, *field, path, satisfies...); err != nil {
					return err
				}
			}
			t.WithNamedChildren(alias, func(wq *TodoQuery) {
				*wq = *query
			})
		case "category":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&CategoryClient{config: t.config}).Query()
			)
			if err := query.collectField(ctx, op, field, path, satisfies...); err != nil {
				return err
			}
			t.withCategory = query
		}
	}
	return nil
}

type todoPaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []TodoPaginateOption
}

func newTodoPaginateArgs(rv map[string]interface{}) *todoPaginateArgs {
	args := &todoPaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	if v, ok := rv[orderByField]; ok {
		switch v := v.(type) {
		case map[string]interface{}:
			var (
				err1, err2 error
				order      = &TodoOrder{Field: &TodoOrderField{}, Direction: entgql.OrderDirectionAsc}
			)
			if d, ok := v[directionField]; ok {
				err1 = order.Direction.UnmarshalGQL(d)
			}
			if f, ok := v[fieldField]; ok {
				err2 = order.Field.UnmarshalGQL(f)
			}
			if err1 == nil && err2 == nil {
				args.opts = append(args.opts, WithTodoOrder(order))
			}
		case *TodoOrder:
			if v != nil {
				args.opts = append(args.opts, WithTodoOrder(v))
			}
		}
	}
	if v, ok := rv[whereField].(*TodoWhereInput); ok {
		args.opts = append(args.opts, WithTodoFilter(v.Filter))
	}
	return args
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (u *UserQuery) CollectFields(ctx context.Context, satisfies ...string) (*UserQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return u, nil
	}
	if err := u.collectField(ctx, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return u, nil
}

func (u *UserQuery) collectField(ctx context.Context, op *graphql.OperationContext, field graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	for _, field := range graphql.CollectFields(op, field.Selections, satisfies) {
		switch field.Name {
		case "groups":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&GroupClient{config: u.config}).Query()
			)
			args := newGroupPaginateArgs(fieldArgs(ctx, new(GroupWhereInput), path...))
			if err := validateFirstLast(args.first, args.last); err != nil {
				return fmt.Errorf("validate first and last in path %q: %w", path, err)
			}
			pager, err := newGroupPager(args.opts, args.last != nil)
			if err != nil {
				return fmt.Errorf("create new pager in path %q: %w", path, err)
			}
			if query, err = pager.applyFilter(query); err != nil {
				return err
			}
			ignoredEdges := !hasCollectedField(ctx, append(path, edgesField)...)
			if hasCollectedField(ctx, append(path, totalCountField)...) || hasCollectedField(ctx, append(path, pageInfoField)...) {
				hasPagination := args.after != nil || args.first != nil || args.before != nil || args.last != nil
				if hasPagination || ignoredEdges {
					query := query.Clone()
					u.loadTotal = append(u.loadTotal, func(ctx context.Context, nodes []*User) error {
						ids := make([]driver.Value, len(nodes))
						for i := range nodes {
							ids[i] = nodes[i].ID
						}
						var v []struct {
							NodeID int `sql:"user_id"`
							Count  int `sql:"count"`
						}
						query.Where(func(s *sql.Selector) {
							joinT := sql.Table(user.GroupsTable)
							s.Join(joinT).On(s.C(group.FieldID), joinT.C(user.GroupsPrimaryKey[1]))
							s.Where(sql.InValues(joinT.C(user.GroupsPrimaryKey[0]), ids...))
							s.Select(joinT.C(user.GroupsPrimaryKey[0]), sql.Count("*"))
							s.GroupBy(joinT.C(user.GroupsPrimaryKey[0]))
						})
						if err := query.Select().Scan(ctx, &v); err != nil {
							return err
						}
						m := make(map[int]int, len(v))
						for i := range v {
							m[v[i].NodeID] = v[i].Count
						}
						for i := range nodes {
							n := m[nodes[i].ID]
							if nodes[i].Edges.totalCount[0] == nil {
								nodes[i].Edges.totalCount[0] = make(map[string]int)
							}
							nodes[i].Edges.totalCount[0][alias] = n
						}
						return nil
					})
				} else {
					u.loadTotal = append(u.loadTotal, func(_ context.Context, nodes []*User) error {
						for i := range nodes {
							n := len(nodes[i].Edges.Groups)
							if nodes[i].Edges.totalCount[0] == nil {
								nodes[i].Edges.totalCount[0] = make(map[string]int)
							}
							nodes[i].Edges.totalCount[0][alias] = n
						}
						return nil
					})
				}
			}
			if ignoredEdges || (args.first != nil && *args.first == 0) || (args.last != nil && *args.last == 0) {
				continue
			}

			if query, err = pager.applyCursors(query, args.after, args.before); err != nil {
				return err
			}
			if limit := paginateLimit(args.first, args.last); limit > 0 {
				modify := limitRows(user.GroupsPrimaryKey[0], limit, pager.orderExpr())
				query.modifiers = append(query.modifiers, modify)
			} else {
				query = pager.applyOrder(query)
			}
			path = append(path, edgesField, nodeField)
			if field := collectedField(ctx, path...); field != nil {
				if err := query.collectField(ctx, op, *field, path, satisfies...); err != nil {
					return err
				}
			}
			u.WithNamedGroups(alias, func(wq *GroupQuery) {
				*wq = *query
			})
		case "friends":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&UserClient{config: u.config}).Query()
			)
			args := newUserPaginateArgs(fieldArgs(ctx, new(UserWhereInput), path...))
			if err := validateFirstLast(args.first, args.last); err != nil {
				return fmt.Errorf("validate first and last in path %q: %w", path, err)
			}
			pager, err := newUserPager(args.opts, args.last != nil)
			if err != nil {
				return fmt.Errorf("create new pager in path %q: %w", path, err)
			}
			if query, err = pager.applyFilter(query); err != nil {
				return err
			}
			ignoredEdges := !hasCollectedField(ctx, append(path, edgesField)...)
			if hasCollectedField(ctx, append(path, totalCountField)...) || hasCollectedField(ctx, append(path, pageInfoField)...) {
				hasPagination := args.after != nil || args.first != nil || args.before != nil || args.last != nil
				if hasPagination || ignoredEdges {
					query := query.Clone()
					u.loadTotal = append(u.loadTotal, func(ctx context.Context, nodes []*User) error {
						ids := make([]driver.Value, len(nodes))
						for i := range nodes {
							ids[i] = nodes[i].ID
						}
						var v []struct {
							NodeID int `sql:"user_id"`
							Count  int `sql:"count"`
						}
						query.Where(func(s *sql.Selector) {
							joinT := sql.Table(user.FriendsTable)
							s.Join(joinT).On(s.C(user.FieldID), joinT.C(user.FriendsPrimaryKey[1]))
							s.Where(sql.InValues(joinT.C(user.FriendsPrimaryKey[0]), ids...))
							s.Select(joinT.C(user.FriendsPrimaryKey[0]), sql.Count("*"))
							s.GroupBy(joinT.C(user.FriendsPrimaryKey[0]))
						})
						if err := query.Select().Scan(ctx, &v); err != nil {
							return err
						}
						m := make(map[int]int, len(v))
						for i := range v {
							m[v[i].NodeID] = v[i].Count
						}
						for i := range nodes {
							n := m[nodes[i].ID]
							if nodes[i].Edges.totalCount[1] == nil {
								nodes[i].Edges.totalCount[1] = make(map[string]int)
							}
							nodes[i].Edges.totalCount[1][alias] = n
						}
						return nil
					})
				} else {
					u.loadTotal = append(u.loadTotal, func(_ context.Context, nodes []*User) error {
						for i := range nodes {
							n := len(nodes[i].Edges.Friends)
							if nodes[i].Edges.totalCount[1] == nil {
								nodes[i].Edges.totalCount[1] = make(map[string]int)
							}
							nodes[i].Edges.totalCount[1][alias] = n
						}
						return nil
					})
				}
			}
			if ignoredEdges || (args.first != nil && *args.first == 0) || (args.last != nil && *args.last == 0) {
				continue
			}

			if query, err = pager.applyCursors(query, args.after, args.before); err != nil {
				return err
			}
			if limit := paginateLimit(args.first, args.last); limit > 0 {
				modify := limitRows(user.FriendsPrimaryKey[0], limit, pager.orderExpr())
				query.modifiers = append(query.modifiers, modify)
			} else {
				query = pager.applyOrder(query)
			}
			path = append(path, edgesField, nodeField)
			if field := collectedField(ctx, path...); field != nil {
				if err := query.collectField(ctx, op, *field, path, satisfies...); err != nil {
					return err
				}
			}
			u.WithNamedFriends(alias, func(wq *UserQuery) {
				*wq = *query
			})
		case "friendships":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&FriendshipClient{config: u.config}).Query()
			)
			args := newFriendshipPaginateArgs(fieldArgs(ctx, new(FriendshipWhereInput), path...))
			if err := validateFirstLast(args.first, args.last); err != nil {
				return fmt.Errorf("validate first and last in path %q: %w", path, err)
			}
			pager, err := newFriendshipPager(args.opts, args.last != nil)
			if err != nil {
				return fmt.Errorf("create new pager in path %q: %w", path, err)
			}
			if query, err = pager.applyFilter(query); err != nil {
				return err
			}
			ignoredEdges := !hasCollectedField(ctx, append(path, edgesField)...)
			if hasCollectedField(ctx, append(path, totalCountField)...) || hasCollectedField(ctx, append(path, pageInfoField)...) {
				hasPagination := args.after != nil || args.first != nil || args.before != nil || args.last != nil
				if hasPagination || ignoredEdges {
					query := query.Clone()
					u.loadTotal = append(u.loadTotal, func(ctx context.Context, nodes []*User) error {
						ids := make([]driver.Value, len(nodes))
						for i := range nodes {
							ids[i] = nodes[i].ID
						}
						var v []struct {
							NodeID int `sql:"user_id"`
							Count  int `sql:"count"`
						}
						query.Where(func(s *sql.Selector) {
							s.Where(sql.InValues(user.FriendshipsColumn, ids...))
						})
						if err := query.GroupBy(user.FriendshipsColumn).Aggregate(Count()).Scan(ctx, &v); err != nil {
							return err
						}
						m := make(map[int]int, len(v))
						for i := range v {
							m[v[i].NodeID] = v[i].Count
						}
						for i := range nodes {
							n := m[nodes[i].ID]
							if nodes[i].Edges.totalCount[2] == nil {
								nodes[i].Edges.totalCount[2] = make(map[string]int)
							}
							nodes[i].Edges.totalCount[2][alias] = n
						}
						return nil
					})
				} else {
					u.loadTotal = append(u.loadTotal, func(_ context.Context, nodes []*User) error {
						for i := range nodes {
							n := len(nodes[i].Edges.Friendships)
							if nodes[i].Edges.totalCount[2] == nil {
								nodes[i].Edges.totalCount[2] = make(map[string]int)
							}
							nodes[i].Edges.totalCount[2][alias] = n
						}
						return nil
					})
				}
			}
			if ignoredEdges || (args.first != nil && *args.first == 0) || (args.last != nil && *args.last == 0) {
				continue
			}

			if query, err = pager.applyCursors(query, args.after, args.before); err != nil {
				return err
			}
			if limit := paginateLimit(args.first, args.last); limit > 0 {
				modify := limitRows(user.FriendshipsColumn, limit, pager.orderExpr())
				query.modifiers = append(query.modifiers, modify)
			} else {
				query = pager.applyOrder(query)
			}
			path = append(path, edgesField, nodeField)
			if field := collectedField(ctx, path...); field != nil {
				if err := query.collectField(ctx, op, *field, path, satisfies...); err != nil {
					return err
				}
			}
			u.WithNamedFriendships(alias, func(wq *FriendshipQuery) {
				*wq = *query
			})
		}
	}
	return nil
}

type userPaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []UserPaginateOption
}

func newUserPaginateArgs(rv map[string]interface{}) *userPaginateArgs {
	args := &userPaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	if v, ok := rv[whereField].(*UserWhereInput); ok {
		args.opts = append(args.opts, WithUserFilter(v.Filter))
	}
	return args
}

const (
	afterField     = "after"
	firstField     = "first"
	beforeField    = "before"
	lastField      = "last"
	orderByField   = "orderBy"
	directionField = "direction"
	fieldField     = "field"
	whereField     = "where"
)

func fieldArgs(ctx context.Context, whereInput interface{}, path ...string) map[string]interface{} {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return nil
	}
	oc := graphql.GetOperationContext(ctx)
	for _, name := range path {
		var field *graphql.CollectedField
		for _, f := range graphql.CollectFields(oc, fc.Field.Selections, nil) {
			if f.Alias == name {
				field = &f
				break
			}
		}
		if field == nil {
			return nil
		}
		cf, err := fc.Child(ctx, *field)
		if err != nil {
			args := field.ArgumentMap(oc.Variables)
			return unmarshalArgs(ctx, whereInput, args)
		}
		fc = cf
	}
	return fc.Args
}

// unmarshalArgs allows extracting the field arguments from their raw representation.
func unmarshalArgs(ctx context.Context, whereInput interface{}, args map[string]interface{}) map[string]interface{} {
	for _, k := range []string{firstField, lastField} {
		v, ok := args[k]
		if !ok {
			continue
		}
		i, err := graphql.UnmarshalInt(v)
		if err == nil {
			args[k] = &i
		}
	}
	for _, k := range []string{beforeField, afterField} {
		v, ok := args[k]
		if !ok {
			continue
		}
		c := &Cursor{}
		if c.UnmarshalGQL(v) == nil {
			args[k] = c
		}
	}
	if v, ok := args[whereField]; ok && whereInput != nil {
		if err := graphql.UnmarshalInputFromContext(ctx, v, whereInput); err == nil {
			args[whereField] = whereInput
		}
	}

	return args
}

func limitRows(partitionBy string, limit int, orderBy ...sql.Querier) func(s *sql.Selector) {
	return func(s *sql.Selector) {
		d := sql.Dialect(s.Dialect())
		s.SetDistinct(false)
		with := d.With("src_query").
			As(s.Clone()).
			With("limited_query").
			As(
				d.Select("*").
					AppendSelectExprAs(
						sql.RowNumber().PartitionBy(partitionBy).OrderExpr(orderBy...),
						"row_number",
					).
					From(d.Table("src_query")),
			)
		t := d.Table("limited_query").As(s.TableName())
		*s = *d.Select(s.UnqualifiedColumns()...).
			From(t).
			Where(sql.LTE(t.C("row_number"), limit)).
			Prefix(with)
	}
}
