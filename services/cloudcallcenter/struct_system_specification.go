package cloudcallcenter

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// SystemSpecification is a nested struct in cloudcallcenter response
type SystemSpecification struct {
	Account         string `json:"Account" xml:"Account"`
	FreeNumberCount int    `json:"FreeNumberCount" xml:"FreeNumberCount"`
	StorageMaxDays  int    `json:"StorageMaxDays" xml:"StorageMaxDays"`
	MaxAgents       int    `json:"MaxAgents" xml:"MaxAgents"`
	MaxRoles        int    `json:"MaxRoles" xml:"MaxRoles"`
	MaxInstances    int    `json:"MaxInstances" xml:"MaxInstances"`
	SpecificationId string `json:"SpecificationId" xml:"SpecificationId"`
	MaxOnlineAgents int    `json:"MaxOnlineAgents" xml:"MaxOnlineAgents"`
	MaxPhoneNumbers int    `json:"MaxPhoneNumbers" xml:"MaxPhoneNumbers"`
	StorageMaxSize  int    `json:"StorageMaxSize" xml:"StorageMaxSize"`
	MaxSkillGroups  int    `json:"MaxSkillGroups" xml:"MaxSkillGroups"`
	MaxContactFlows int    `json:"MaxContactFlows" xml:"MaxContactFlows"`
}
