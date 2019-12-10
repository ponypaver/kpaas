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

package check

import (
	"github.com/kpaas-io/kpaas/pkg/deploy/assets"
	"github.com/kpaas-io/kpaas/pkg/deploy/command"
	"github.com/kpaas-io/kpaas/pkg/deploy/machine"
	"github.com/kpaas-io/kpaas/pkg/deploy/operation"
	pb "github.com/kpaas-io/kpaas/pkg/deploy/protos"
)

const (
	memoryScript    = "/scripts/check_memory_capacity.sh"
	memoryRemoteDir = "/tmp"
)

type CheckMemoryOperation struct {
	operation.BaseOperation
	CheckOperations
}

func (ckops *CheckMemoryOperation) getScript() string {
	ckops.Script = memoryScript
	return ckops.Script
}

func (ckops *CheckMemoryOperation) getScriptPath() string {
	ckops.ScriptPath = memoryRemoteDir
	return ckops.ScriptPath
}

func (ckops *CheckMemoryOperation) GetOperations(config *pb.NodeCheckConfig) (operation.Operation, error) {
	ops := &CheckMemoryOperation{}
	m, err := machine.NewMachine(config.Node)
	if err != nil {
		return nil, err
	}

	scriptFile, err := assets.Assets.Open(ckops.getScript())
	if err != nil {
		return nil, err
	}

	if err := m.PutFile(scriptFile, ckops.getScriptPath()+ckops.getScript()); err != nil {
		return nil, err
	}

	ops.AddCommands(command.NewShellCommand(m, "bash", ckops.getScriptPath()+ckops.getScript(), nil))
	return ops, nil
}

// check if memory capacity satisfied with minimal requirement
func CheckMemoryCapacity(comparedMemory string, desiredMemory float64) error {
	err := operation.CheckEntity(comparedMemory, desiredMemory)
	if err != nil {
		return err
	}

	return nil
}