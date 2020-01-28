// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/entc/integration/ent/group"
	"github.com/facebookincubator/ent/entc/integration/ent/predicate"
	"github.com/facebookincubator/ent/schema/field"
)

// GroupDelete is the builder for deleting a Group entity.
type GroupDelete struct {
	config
	predicates []predicate.Group
}

// Where adds a new predicate to the delete builder.
func (gd *GroupDelete) Where(ps ...predicate.Group) *GroupDelete {
	gd.predicates = append(gd.predicates, ps...)
	return gd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (gd *GroupDelete) Exec(ctx context.Context) (int, error) {
	return gd.sqlExec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (gd *GroupDelete) ExecX(ctx context.Context) int {
	n, err := gd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (gd *GroupDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: group.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: group.FieldID,
			},
		},
	}
	if ps := gd.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, gd.driver, _spec)
}

// GroupDeleteOne is the builder for deleting a single Group entity.
type GroupDeleteOne struct {
	gd *GroupDelete
}

// Exec executes the deletion query.
func (gdo *GroupDeleteOne) Exec(ctx context.Context) error {
	n, err := gdo.gd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{group.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (gdo *GroupDeleteOne) ExecX(ctx context.Context) {
	gdo.gd.ExecX(ctx)
}
