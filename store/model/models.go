// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package model

import (
	git "github.com/bluebrown/kobold/git"
	store "github.com/bluebrown/kobold/store"
	null "github.com/volatiletech/null/v8"
)

type Channel struct {
	Name        string      `json:"name"`
	DecoderName null.String `json:"decoder_name"`
}

type Decoder struct {
	Name   string `json:"name"`
	Script []byte `json:"script"`
}

type Pipeline struct {
	Name         string         `json:"name"`
	RepoUri      git.PackageURI `json:"repo_uri"`
	DestBranch   null.String    `json:"dest_branch"`
	PostHookName null.String    `json:"post_hook_name"`
}

type PipelineListItem struct {
	Name         string         `json:"name"`
	RepoUri      git.PackageURI `json:"repo_uri"`
	DestBranch   null.String    `json:"dest_branch"`
	PostHookName null.String    `json:"post_hook_name"`
	Channels     store.FlatList `json:"channels"`
}

type PostHook struct {
	Name   string `json:"name"`
	Script []byte `json:"script"`
}

type Run struct {
	Fingerprint string         `json:"fingerprint"`
	RepoUri     git.PackageURI `json:"repo_uri"`
	DestBranch  null.String    `json:"dest_branch"`
	PostHook    null.String    `json:"post_hook"`
	Status      string         `json:"status"`
	Timestamp   interface{}    `json:"timestamp"`
	Warnings    store.FlatList `json:"warnings"`
	Error       interface{}    `json:"error"`
	Msgs        store.FlatList `json:"msgs"`
}

type Subscription struct {
	PipelineName string `json:"pipeline_name"`
	ChannelName  string `json:"channel_name"`
}

type Task struct {
	ID                   string         `json:"id"`
	Msgs                 store.FlatList `json:"msgs"`
	RepoUri              git.PackageURI `json:"repo_uri"`
	DestBranch           null.String    `json:"dest_branch"`
	PostHookName         null.String    `json:"post_hook_name"`
	Status               string         `json:"status"`
	Timestamp            string         `json:"timestamp"`
	Warnings             store.FlatList `json:"warnings"`
	FailureReason        null.String    `json:"failure_reason"`
	TaskGroupFingerprint null.String    `json:"task_group_fingerprint"`
}

type TaskGroup struct {
	Fingerprint string         `json:"fingerprint"`
	RepoUri     git.PackageURI `json:"repo_uri"`
	DestBranch  null.String    `json:"dest_branch"`
	PostHook    []byte         `json:"post_hook"`
	TaskIds     store.FlatList `json:"task_ids"`
	Msgs        store.FlatList `json:"msgs"`
}
