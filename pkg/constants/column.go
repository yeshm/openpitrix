// Copyright 2018 The OpenPitrix Authors. All rights reserved.
// Use of this source code is governed by a Apache license
// that can be found in the LICENSE file.

package constants

const (
	ColumnAppId            = "app_id"
	ColumnCategoryId       = "category_id"
	ColumnChartName        = "chart_name"
	ColumnClusterId        = "cluster_id"
	ColumnClusterType      = "cluster_type"
	ColumnCreateTime       = "create_time"
	ColumnCredential       = "credential"
	ColumnDescription      = "description"
	ColumnExecutor         = "executor"
	ColumnFrontgateId      = "frontgate_id"
	ColumnHome             = "home"
	ColumnIcon             = "icon"
	ColumnInstanceId       = "instance_id"
	ColumnJobAction        = "job_action"
	ColumnJobId            = "job_id"
	ColumnKeywords         = "keywords"
	ColumnLabelKey         = "label_key"
	ColumnLabelValue       = "label_value"
	ColumnLocale           = "locale"
	ColumnMaintainers      = "maintainers"
	ColumnName             = "name"
	ColumnNodeId           = "node_id"
	ColumnKeyPairId        = "key_pair_id"
	ColumnOwner            = "owner"
	ColumnPackageName      = "package_name"
	ColumnPrivateIp        = "private_ip"
	ColumnProvider         = "provider"
	ColumnReadme           = "readme"
	ColumnRepoEventId      = "repo_event_id"
	ColumnRepoId           = "repo_id"
	ColumnRepoLabelId      = "repo_label_id"
	ColumnRepoSelectorId   = "repo_selector_id"
	ColumnResouceId        = "resource_id"
	ColumnRole             = "role"
	ColumnRuntimeId        = "runtime_id"
	ColumnRuntimeLabelId   = "runtime_label_id"
	ColumnScreenshots      = "screenshots"
	ColumnSelectorKey      = "selector_key"
	ColumnSelectorValue    = "selector_value"
	ColumnSequence         = "sequence"
	ColumnSources          = "sources"
	ColumnStatus           = "status"
	ColumnStatusTime       = "status_time"
	ColumnTarget           = "target"
	ColumnTaskAction       = "task_action"
	ColumnTaskId           = "task_id"
	ColumnTransitionStatus = "transition_status"
	ColumnType             = "type"
	ColumnUpdateTime       = "update_time"
	ColumnUrl              = "url"
	ColumnVersionId        = "version_id"
	ColumnVisibility       = "visibility"
	ColumnVolumeId         = "volume_id"
	ColumnZone             = "zone"

	ColumnUserId       = "user_id"
	ColumnGroupId      = "group_id"
	ColumnResetId      = "reset_id"
	ColumnPassword     = "password"
	ColumnEmail        = "email"
	ColumnClientId     = "client_id"
	ColumnClientSecret = "client_secret"
	ColumnRefreshToken = "refresh_token"
	ColumnAccessToken  = "access_token"
	ColumnTokenId      = "token_id"
	ColumnScope        = "scope"
	ColumnUsername     = "username"

	ColumnMessage = "message"

	ColumnAppDefaultStatus = "app_default_status"
)

var PushEventTables = map[string][]string{
	TableRepoEvent: {
		ColumnRepoEventId, ColumnRepoId, ColumnStatus,
	},
	TableCluster: {
		ColumnClusterId, ColumnStatus, ColumnTransitionStatus,
	},
	TableClusterNode: {
		ColumnNodeId, ColumnStatus, ColumnTransitionStatus,
	},
	TableJob: {
		ColumnJobId, ColumnStatus, ColumnClusterId, ColumnAppId, ColumnAppId,
	},
}

// columns that can be search through sql '=' operator
var IndexedColumns = map[string][]string{
	TableApp: {
		ColumnAppId, ColumnName, ColumnRepoId, ColumnDescription, ColumnStatus,
		ColumnHome, ColumnIcon, ColumnScreenshots, ColumnMaintainers, ColumnSources,
		ColumnReadme, ColumnOwner, ColumnChartName,
	},
	TableAppVersion: {
		ColumnVersionId, ColumnAppId, ColumnName, ColumnOwner, ColumnDescription,
		ColumnPackageName, ColumnStatus,
	},
	TableJob: {
		ColumnJobId, ColumnClusterId, ColumnAppId, ColumnVersionId,
		ColumnExecutor, ColumnProvider, ColumnStatus, ColumnOwner,
	},
	TableTask: {
		ColumnJobId, ColumnTaskId, ColumnExecutor, ColumnStatus, ColumnOwner,
	},
	TableRepo: {
		ColumnRepoId, ColumnName, ColumnType, ColumnVisibility, ColumnStatus, ColumnAppDefaultStatus, ColumnOwner,
	},
	TableRuntime: {
		ColumnRuntimeId, ColumnProvider, ColumnZone, ColumnStatus, ColumnOwner,
	},
	TableRepoLabel: {
		ColumnRepoId, ColumnRepoLabelId, ColumnStatus,
	},
	TableRepoSelector: {
		ColumnRepoId, ColumnRepoSelectorId, ColumnStatus,
	},
	TableRepoEvent: {
		ColumnRepoEventId, ColumnRepoId, ColumnStatus, ColumnOwner,
	},
	TableCluster: {
		ColumnClusterId, ColumnAppId, ColumnVersionId, ColumnStatus,
		ColumnRuntimeId, ColumnFrontgateId, ColumnOwner, ColumnClusterType,
	},
	TableClusterNode: {
		ColumnClusterId, ColumnNodeId, ColumnStatus, ColumnOwner,
	},
	TableCategory: {
		ColumnCategoryId, ColumnStatus, ColumnLocale, ColumnOwner, ColumnName,
	},
	TableUser: {
		ColumnGroupId, ColumnUserId, ColumnStatus, ColumnRole,
	},
}

var SearchWordColumnTable = []string{
	TableRuntime,
	TableApp,
	TableAppVersion,
	TableRepo,
	TableJob,
	TableTask,
	TableCluster,
	TableClusterNode,
	TableUser,
}

// columns that can be search through sql 'like' operator
var SearchColumns = map[string][]string{
	TableApp: {
		ColumnAppId, ColumnName, ColumnRepoId, ColumnOwner, ColumnChartName, ColumnKeywords,
	},
	TableAppVersion: {
		ColumnVersionId, ColumnAppId, ColumnName, ColumnDescription, ColumnOwner, ColumnPackageName,
	},
	TableJob: {
		ColumnJobId, ColumnClusterId, ColumnOwner, ColumnJobAction, ColumnExecutor, ColumnProvider, ColumnExecutor, ColumnProvider,
	},
	TableTask: {
		ColumnJobId, ColumnTaskId, ColumnTaskAction, ColumnOwner, ColumnNodeId, ColumnTarget,
	},
	TableRuntime: {
		ColumnRuntimeId, ColumnName, ColumnOwner, ColumnProvider, ColumnZone,
	},
	TableCluster: {
		ColumnClusterId, ColumnName, ColumnOwner, ColumnAppId, ColumnVersionId, ColumnRuntimeId,
	},
	TableClusterNode: {
		ColumnNodeId, ColumnClusterId, ColumnName, ColumnInstanceId, ColumnVolumeId, ColumnPrivateIp, ColumnRole, ColumnOwner,
	},
	TableRepo: {
		ColumnName, ColumnDescription,
	},
	TableUser: {
		ColumnUserId, ColumnDescription, ColumnEmail, ColumnUsername,
	},
}
