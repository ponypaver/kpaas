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

package wizard

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/kpaas-io/kpaas/pkg/constant"
	"github.com/kpaas-io/kpaas/pkg/service/model/common"
)

func TestNewNode(t *testing.T) {

	node := NewNode()
	assert.IsType(t, &Node{}, node)
	assert.NotNil(t, node.MachineRoles)
	assert.NotNil(t, node.Labels)
	assert.NotNil(t, node.Taints)
	assert.NotNil(t, node.CheckReport)
	assert.NotNil(t, node.DeploymentReports)
	assert.Equal(t, AuthenticationTypePassword, node.AuthenticationType)
}

func TestNewDeploymentReport(t *testing.T) {

	report := NewDeploymentReport()
	assert.NotNil(t, report)
	assert.Equal(t, DeployStatusPending, report.Status)
	assert.Nil(t, report.Error)
}

func TestNewCheckItem(t *testing.T) {

	item := NewCheckItem()
	assert.Equal(t, constant.CheckResultPending, item.CheckResult)
}

func TestNode_SetCheckResult(t *testing.T) {

	tests := []struct {
		Input struct {
			Node          *Node
			CheckResult   constant.CheckResult
			FailureDetail *common.FailureDetail
		}
		Want *Node
	}{
		{
			Input: struct {
				Node          *Node
				CheckResult   constant.CheckResult
				FailureDetail *common.FailureDetail
			}{
				Node: &Node{
					CheckReport: &CheckReport{
						CheckResult: constant.CheckResultRunning,
					},
				},
				CheckResult:   constant.CheckResultSuccessful,
				FailureDetail: nil,
			},
			Want: &Node{
				CheckReport: &CheckReport{
					CheckResult: constant.CheckResultSuccessful,
				},
			},
		},
		{
			Input: struct {
				Node          *Node
				CheckResult   constant.CheckResult
				FailureDetail *common.FailureDetail
			}{
				Node: &Node{
					CheckReport: &CheckReport{
						CheckResult: constant.CheckResultRunning,
					},
				},
				CheckResult: constant.CheckResultFailed,
				FailureDetail: &common.FailureDetail{
					Reason:     "reason",
					Detail:     "detail",
					FixMethods: "fix",
					LogId:      1,
				},
			},
			Want: &Node{
				CheckReport: &CheckReport{
					CheckResult: constant.CheckResultFailed,
					CheckedError: &common.FailureDetail{
						Reason:     "reason",
						Detail:     "detail",
						FixMethods: "fix",
						LogId:      1,
					},
				},
			},
		},
	}

	for _, item := range tests {

		item.Input.Node.SetCheckResult(item.Input.CheckResult, item.Input.FailureDetail)
		assert.Equal(t, item.Want, item.Input.Node)
	}
}

func TestNode_SetCheckItem(t *testing.T) {

	tests := []struct {
		Input struct {
			Node          *Node
			ItemName      string
			CheckResult   constant.CheckResult
			FailureDetail *common.FailureDetail
		}
		Want *Node
	}{
		{
			Input: struct {
				Node          *Node
				ItemName      string
				CheckResult   constant.CheckResult
				FailureDetail *common.FailureDetail
			}{
				Node: &Node{
					CheckReport: &CheckReport{
						CheckItems: []*CheckItem{},
					},
				},
				ItemName:      "item 1",
				CheckResult:   constant.CheckResultRunning,
				FailureDetail: nil,
			},
			Want: &Node{
				CheckReport: &CheckReport{
					CheckItems: []*CheckItem{
						{
							ItemName:    "item 1",
							CheckResult: constant.CheckResultRunning,
							Error:       nil,
						},
					},
				},
			},
		},
		{
			Input: struct {
				Node          *Node
				ItemName      string
				CheckResult   constant.CheckResult
				FailureDetail *common.FailureDetail
			}{
				Node: &Node{
					CheckReport: &CheckReport{
						CheckItems: []*CheckItem{
							{
								ItemName:    "item 1",
								CheckResult: constant.CheckResultRunning,
								Error:       nil,
							},
						},
					},
				},
				ItemName:      "item 1",
				CheckResult:   constant.CheckResultSuccessful,
				FailureDetail: nil,
			},
			Want: &Node{
				CheckReport: &CheckReport{
					CheckItems: []*CheckItem{
						{
							ItemName:    "item 1",
							CheckResult: constant.CheckResultSuccessful,
							Error:       nil,
						},
					},
				},
			},
		},
		{
			Input: struct {
				Node          *Node
				ItemName      string
				CheckResult   constant.CheckResult
				FailureDetail *common.FailureDetail
			}{
				Node: &Node{
					CheckReport: &CheckReport{
						CheckItems: []*CheckItem{
							{
								ItemName:    "item 2",
								CheckResult: constant.CheckResultRunning,
								Error:       nil,
							},
						},
					},
				},
				ItemName:    "item 2",
				CheckResult: constant.CheckResultFailed,
				FailureDetail: &common.FailureDetail{
					Reason:     "reason",
					Detail:     "detail",
					FixMethods: "fix",
					LogId:      1,
				},
			},
			Want: &Node{
				CheckReport: &CheckReport{
					CheckItems: []*CheckItem{
						{
							ItemName:    "item 2",
							CheckResult: constant.CheckResultFailed,
							Error: &common.FailureDetail{
								Reason:     "reason",
								Detail:     "detail",
								FixMethods: "fix",
								LogId:      1,
							},
						},
					},
				},
			},
		},
		{
			Input: struct {
				Node          *Node
				ItemName      string
				CheckResult   constant.CheckResult
				FailureDetail *common.FailureDetail
			}{
				Node: &Node{
					CheckReport: &CheckReport{
						CheckItems: []*CheckItem{
							{
								ItemName:    "item 1",
								CheckResult: constant.CheckResultRunning,
								Error:       nil,
							},
						},
					},
				},
				ItemName:    "item 2",
				CheckResult: constant.CheckResultFailed,
				FailureDetail: &common.FailureDetail{
					Reason:     "reason",
					Detail:     "detail",
					FixMethods: "fix",
					LogId:      1,
				},
			},
			Want: &Node{
				CheckReport: &CheckReport{
					CheckItems: []*CheckItem{
						{
							ItemName:    "item 1",
							CheckResult: constant.CheckResultRunning,
							Error:       nil,
						},
						{
							ItemName:    "item 2",
							CheckResult: constant.CheckResultFailed,
							Error: &common.FailureDetail{
								Reason:     "reason",
								Detail:     "detail",
								FixMethods: "fix",
								LogId:      1,
							},
						},
					},
				},
			},
		},
	}

	for _, item := range tests {

		item.Input.Node.SetCheckItem(item.Input.ItemName, item.Input.CheckResult, item.Input.FailureDetail)
		assert.Equal(t, item.Want, item.Input.Node)
	}
}

func TestNode_SetDeployResult(t *testing.T) {

	tests := []struct {
		Input struct {
			Node          *Node
			Item          constant.DeployItem
			Status        DeployStatus
			FailureDetail *common.FailureDetail
		}
		Want *Node
	}{
		{
			Input: struct {
				Node          *Node
				Item          constant.DeployItem
				Status        DeployStatus
				FailureDetail *common.FailureDetail
			}{
				Node: &Node{
					DeploymentReports: map[constant.DeployItem]*DeploymentReport{},
				},
				Item:          constant.DeployItemMaster,
				Status:        DeployStatusPending,
				FailureDetail: nil,
			},
			Want: &Node{
				DeploymentReports: map[constant.DeployItem]*DeploymentReport{
					constant.DeployItemMaster: {
						DeployItem: constant.DeployItemMaster,
						Status:     DeployStatusPending,
						Error:      nil,
					},
				},
			},
		},
		{
			Input: struct {
				Node          *Node
				Item          constant.DeployItem
				Status        DeployStatus
				FailureDetail *common.FailureDetail
			}{
				Node: &Node{
					DeploymentReports: map[constant.DeployItem]*DeploymentReport{
						constant.DeployItemMaster: {
							DeployItem: constant.DeployItemMaster,
							Status:     DeployStatusPending,
							Error:      nil,
						},
					},
				},
				Item:          constant.DeployItemMaster,
				Status:        DeployStatusRunning,
				FailureDetail: nil,
			},
			Want: &Node{
				DeploymentReports: map[constant.DeployItem]*DeploymentReport{
					constant.DeployItemMaster: {
						DeployItem: constant.DeployItemMaster,
						Status:     DeployStatusRunning,
						Error:      nil,
					},
				},
			},
		},
		{
			Input: struct {
				Node          *Node
				Item          constant.DeployItem
				Status        DeployStatus
				FailureDetail *common.FailureDetail
			}{
				Node: &Node{
					DeploymentReports: map[constant.DeployItem]*DeploymentReport{
						constant.DeployItemMaster: {
							DeployItem: constant.DeployItemMaster,
							Status:     DeployStatusRunning,
							Error:      nil,
						},
					},
				},
				Item:   constant.DeployItemMaster,
				Status: DeployStatusFailed,
				FailureDetail: &common.FailureDetail{
					Reason:     "reason",
					Detail:     "detail",
					FixMethods: "fix",
					LogId:      1,
				},
			},
			Want: &Node{
				DeploymentReports: map[constant.DeployItem]*DeploymentReport{
					constant.DeployItemMaster: {
						DeployItem: constant.DeployItemMaster,
						Status:     DeployStatusFailed,
						Error: &common.FailureDetail{
							Reason:     "reason",
							Detail:     "detail",
							FixMethods: "fix",
							LogId:      1,
						},
					},
				},
			},
		},
		{
			Input: struct {
				Node          *Node
				Item          constant.DeployItem
				Status        DeployStatus
				FailureDetail *common.FailureDetail
			}{
				Node: &Node{
					DeploymentReports: map[constant.DeployItem]*DeploymentReport{
						constant.DeployItemMaster: {
							DeployItem: constant.DeployItemMaster,
							Status:     DeployStatusRunning,
							Error:      nil,
						},
					},
				},
				Item:          constant.DeployItemEtcd,
				Status:        DeployStatusSuccessful,
				FailureDetail: nil,
			},
			Want: &Node{
				DeploymentReports: map[constant.DeployItem]*DeploymentReport{
					constant.DeployItemMaster: {
						DeployItem: constant.DeployItemMaster,
						Status:     DeployStatusRunning,
						Error:      nil,
					},
					constant.DeployItemEtcd: {
						DeployItem: constant.DeployItemEtcd,
						Status:     DeployStatusSuccessful,
						Error:      nil,
					},
				},
			},
		},
	}

	for _, item := range tests {

		item.Input.Node.SetDeployResult(item.Input.Item, item.Input.Status, item.Input.FailureDetail)
		assert.Equal(t, item.Want, item.Input.Node)
	}
}
