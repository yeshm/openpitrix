// Copyright 2018 The OpenPitrix Authors. All rights reserved.
// Use of this source code is governed by a Apache license
// that can be found in the LICENSE file.

package models

import (
	"time"

	"openpitrix.io/openpitrix/pkg/constants"
	"openpitrix.io/openpitrix/pkg/db"
	"openpitrix.io/openpitrix/pkg/logger"
	"openpitrix.io/openpitrix/pkg/pb"
	"openpitrix.io/openpitrix/pkg/util/idutil"
	"openpitrix.io/openpitrix/pkg/util/jsonutil"
	"openpitrix.io/openpitrix/pkg/util/pbutil"
)

func NewTaskId() string {
	return idutil.GetUuid("t-")
}

type Task struct {
	TaskId         string
	JobId          string
	TaskAction     string
	Directive      string
	Owner          string
	Status         string
	ErrorCode      uint32
	Executor       string
	Target         string
	NodeId         string
	FailureAllowed bool
	CreateTime     time.Time
	StatusTime     time.Time
}

var TaskColumns = db.GetColumnsFromStruct(&Task{})

func NewTask(taskId, jobId, nodeId, target, taskAction, directive, userId string, failureAllowed bool) *Task {
	if taskId == "" {
		taskId = NewTaskId()
	} else if taskId == constants.PlaceHolder {
		taskId = ""
	}
	return &Task{
		TaskId:         taskId,
		JobId:          jobId,
		NodeId:         nodeId,
		Target:         target,
		TaskAction:     taskAction,
		Directive:      directive,
		Owner:          userId,
		Status:         constants.StatusPending,
		CreateTime:     time.Now(),
		StatusTime:     time.Now(),
		FailureAllowed: failureAllowed,
	}
}

func TaskToPb(task *Task) *pb.Task {
	pbTask := pb.Task{}
	pbTask.TaskId = pbutil.ToProtoString(task.TaskId)
	pbTask.JobId = pbutil.ToProtoString(task.JobId)
	pbTask.TaskAction = pbutil.ToProtoString(task.TaskAction)
	pbTask.Directive = pbutil.ToProtoString(task.Directive)
	pbTask.Owner = pbutil.ToProtoString(task.Owner)
	pbTask.Status = pbutil.ToProtoString(task.Status)
	pbTask.ErrorCode = pbutil.ToProtoUInt32(task.ErrorCode)
	pbTask.Executor = pbutil.ToProtoString(task.Executor)
	pbTask.Target = pbutil.ToProtoString(task.Target)
	pbTask.NodeId = pbutil.ToProtoString(task.NodeId)
	pbTask.CreateTime = pbutil.ToProtoTimestamp(task.CreateTime)
	pbTask.StatusTime = pbutil.ToProtoTimestamp(task.StatusTime)
	pbTask.FailureAllowed = pbutil.ToProtoBool(task.FailureAllowed)
	return &pbTask
}

func TasksToPbs(tasks []*Task) (pbTasks []*pb.Task) {
	for _, task := range tasks {
		pbTasks = append(pbTasks, TaskToPb(task))
	}
	return
}

func (t *Task) GetTimeout(defaultTimeout time.Duration) time.Duration {
	if t.Directive == "" {
		return defaultTimeout
	}

	directive := make(map[string]interface{})
	err := jsonutil.Decode([]byte(t.Directive), &directive)
	if err != nil {
		logger.Error(nil, "Decode task [%s] directive [%s] failed: %+v.", t.TaskId, t.Directive, err)
		return defaultTimeout
	}

	timeout, exist := directive[constants.TimeoutName]
	if !exist {
		return defaultTimeout
	}
	tm := timeout.(float64)
	if tm <= 0 {
		return defaultTimeout
	}
	return time.Duration(tm) * time.Second
}
