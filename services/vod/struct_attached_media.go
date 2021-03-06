package vod

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

// AttachedMedia is a nested struct in vod response
type AttachedMedia struct {
	BusinessType     string     `json:"BusinessType" xml:"BusinessType"`
	Tags             string     `json:"Tags" xml:"Tags"`
	Title            string     `json:"Title" xml:"Title"`
	ModificationTime string     `json:"ModificationTime" xml:"ModificationTime"`
	StorageLocation  string     `json:"StorageLocation" xml:"StorageLocation"`
	CreationTime     string     `json:"CreationTime" xml:"CreationTime"`
	Icon             string     `json:"Icon" xml:"Icon"`
	URL              string     `json:"URL" xml:"URL"`
	AppId            string     `json:"AppId" xml:"AppId"`
	RegionId         string     `json:"RegionId" xml:"RegionId"`
	FileSize         int64      `json:"FileSize" xml:"FileSize"`
	Status           string     `json:"Status" xml:"Status"`
	MediaId          string     `json:"MediaId" xml:"MediaId"`
	Description      string     `json:"Description" xml:"Description"`
	Ext              string     `json:"Ext" xml:"Ext"`
	OnlineStatus     string     `json:"OnlineStatus" xml:"OnlineStatus"`
	Type             string     `json:"Type" xml:"Type"`
	Categories       []Category `json:"Categories" xml:"Categories"`
}
