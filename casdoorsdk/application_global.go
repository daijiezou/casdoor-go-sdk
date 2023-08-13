// Copyright 2023 The Casdoor Authors. All Rights Reserved.
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

package casdoorsdk

func GetApplications() ([]*Application, error) {
	return globalClient.GetApplications()
}

func GetApplication(name string) ([]*Application, error) {
	return globalClient.GetApplications()
}

func AddApplication(application *Application) (bool, error) {
	return globalClient.AddApplication(application)
}

func DeleteApplication(name string) (bool, error) {
	return globalClient.DeleteApplication(name)
}

func UpdateApplication(application *Application) (bool, error) {
	return globalClient.UpdateApplication(application)
}
