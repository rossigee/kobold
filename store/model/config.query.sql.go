// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: config.query.sql

package model

import (
	"context"

	git "github.com/bluebrown/kobold/git"
	null "github.com/volatiletech/null/v8"
)

const channelPut = `-- name: ChannelPut :exec
insert into channel(name, decoder_name) values (?, ?)
on conflict(name) do update set decoder_name = excluded.decoder_name
`

type ChannelPutParams struct {
	Name        string      `json:"name"`
	DecoderName null.String `json:"decoder_name"`
}

// ChannelPut
//
//	insert into channel(name, decoder_name) values (?, ?)
//	on conflict(name) do update set decoder_name = excluded.decoder_name
func (q *Queries) ChannelPut(ctx context.Context, arg ChannelPutParams) error {
	_, err := q.db.ExecContext(ctx, channelPut, arg.Name, arg.DecoderName)
	return err
}

const decoderPut = `-- name: DecoderPut :exec
insert into decoder(name, script) values (?, ?)
on conflict(name) do update set script = excluded.script
`

type DecoderPutParams struct {
	Name   string `json:"name"`
	Script []byte `json:"script"`
}

// DecoderPut
//
//	insert into decoder(name, script) values (?, ?)
//	on conflict(name) do update set script = excluded.script
func (q *Queries) DecoderPut(ctx context.Context, arg DecoderPutParams) error {
	_, err := q.db.ExecContext(ctx, decoderPut, arg.Name, arg.Script)
	return err
}

const pipelinePut = `-- name: PipelinePut :exec
insert into pipeline(name, repo_uri, dest_branch, post_hook_name) values (?, ?, ?, ?)
on conflict(name) do update set repo_uri = excluded.repo_uri, dest_branch = excluded.dest_branch, post_hook_name = excluded.post_hook_name
`

type PipelinePutParams struct {
	Name         string         `json:"name"`
	RepoUri      git.PackageURI `json:"repo_uri"`
	DestBranch   null.String    `json:"dest_branch"`
	PostHookName null.String    `json:"post_hook_name"`
}

// PipelinePut
//
//	insert into pipeline(name, repo_uri, dest_branch, post_hook_name) values (?, ?, ?, ?)
//	on conflict(name) do update set repo_uri = excluded.repo_uri, dest_branch = excluded.dest_branch, post_hook_name = excluded.post_hook_name
func (q *Queries) PipelinePut(ctx context.Context, arg PipelinePutParams) error {
	_, err := q.db.ExecContext(ctx, pipelinePut,
		arg.Name,
		arg.RepoUri,
		arg.DestBranch,
		arg.PostHookName,
	)
	return err
}

const postHookPut = `-- name: PostHookPut :exec
insert into post_hook(name, script) values (?, ?)
on conflict(name) do update set script = excluded.script
`

type PostHookPutParams struct {
	Name   string `json:"name"`
	Script []byte `json:"script"`
}

// PostHookPut
//
//	insert into post_hook(name, script) values (?, ?)
//	on conflict(name) do update set script = excluded.script
func (q *Queries) PostHookPut(ctx context.Context, arg PostHookPutParams) error {
	_, err := q.db.ExecContext(ctx, postHookPut, arg.Name, arg.Script)
	return err
}

const subscriptionPut = `-- name: SubscriptionPut :exec
insert into subscription(pipeline_name, channel_name) values (?, ?)
on conflict(pipeline_name, channel_name) do nothing
`

type SubscriptionPutParams struct {
	PipelineName string `json:"pipeline_name"`
	ChannelName  string `json:"channel_name"`
}

// SubscriptionPut
//
//	insert into subscription(pipeline_name, channel_name) values (?, ?)
//	on conflict(pipeline_name, channel_name) do nothing
func (q *Queries) SubscriptionPut(ctx context.Context, arg SubscriptionPutParams) error {
	_, err := q.db.ExecContext(ctx, subscriptionPut, arg.PipelineName, arg.ChannelName)
	return err
}
