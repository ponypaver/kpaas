// Copyright 2019 Shanghai JingDuo Information Technology co., Ltd.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package action

import (
	"fmt"
	"io"
	"path/filepath"
	"time"

	pb "github.com/kpaas-io/kpaas/pkg/deploy/protos"
	"github.com/kpaas-io/kpaas/pkg/utils/idcreator"
)

// Type represents the type of an action
type Type string

// Status represents the status of an action
type Status string

const (
	ActionPending Status = "pending"
	ActionDoing   Status = "doing"
	ActionDone    Status = "done" // means success
	ActionFailed  Status = "failed"
)

// ItemStatus represents the status of an action item
type ItemStatus string

const (
	ItemPending ItemStatus = "pending"
	ItemDoing   ItemStatus = "doing"
	ItemDone    ItemStatus = "done" // means success
	ItemFailed  ItemStatus = "failed"
)

// Action repsents the definition of executable command(s) in a node,
// multiple actions can be executed concurrently.
type Action interface {
	GetName() string
	GetStatus() Status
	SetStatus(Status)
	GetType() Type
	GetErr() *pb.Error
	SetErr(*pb.Error)
	GetLogFilePath() string
	SetLogFilePath(string)
	GetCreationTimestamp() time.Time
	GetNode() *pb.Node
	GetExecuteLogBuffer() io.ReadWriter
	SetExecuteLogBuffer(io.ReadWriter)
}

// Base is the basic metadata of an action
type Base struct {
	Name              string
	ActionType        Type
	Status            Status
	Err               *pb.Error
	LogFilePath       string
	CreationTimestamp time.Time
	Node              *pb.Node
	ExecuteLogBuffer  io.ReadWriter
}

func (b *Base) GetName() string {
	return b.Name
}

func (b *Base) GetStatus() Status {
	return b.Status
}

func (b *Base) SetStatus(status Status) {
	b.Status = status
}

func (b *Base) GetType() Type {
	return b.ActionType
}

func (b *Base) GetErr() *pb.Error {
	return b.Err
}

func (b *Base) SetErr(err *pb.Error) {
	b.Err = err
}

func (b *Base) GetLogFilePath() string {
	return b.LogFilePath
}

func (b *Base) SetLogFilePath(path string) {
	b.LogFilePath = path
}

func (b *Base) GetCreationTimestamp() time.Time {
	return b.CreationTimestamp
}

func (b *Base) GetNode() *pb.Node {
	return b.Node
}

func (b *Base) GetExecuteLogBuffer() io.ReadWriter {
	return b.ExecuteLogBuffer
}

func (b *Base) SetExecuteLogBuffer(buf io.ReadWriter) {
	b.ExecuteLogBuffer = buf
}

// GenActionLogFilePath is a helper to return a file path based on the base path and aciton name
func GenActionLogFilePath(basePath, actionName string, nodeName string) string {
	if basePath == "" || actionName == "" || nodeName == "" {
		return ""
	}
	fileName := fmt.Sprintf("%s-%s.log", nodeName, actionName)
	return filepath.Join(basePath, fileName)
}

// GenActionName generates a unique action name with the action type as prefix.
func GenActionName(actionType Type) string {
	return fmt.Sprintf("%s-%s", actionType, idcreator.NextString())
}
