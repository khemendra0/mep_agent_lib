/*
 *  Copyright 2020 Huawei Technologies Co., Ltd.
 *
 *  Licensed under the Apache License, Version 2.0 (the "License");
 *  you may not use this file except in compliance with the License.
 *  You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 *  Unless required by applicable law or agreed to in writing, software
 *  distributed under the License is distributed on an "AS IS" BASIS,
 *  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *  See the License for the specific language governing permissions and
 *  limitations under the License.
 */

package test

import (
	"github.com/agiledragon/gomonkey"
	"github.com/smartystreets/goconvey/convey"
	"mep-agent/src/service"
	"testing"
)

func TestStartSuccess(t *testing.T) {

	convey.Convey("RegisterToMepTest", t, func() {
		patch := gomonkey.ApplyFunc(service.RegisterToMep, func(param string, url string) (string, error) {
			return "", nil
		})

		_, err := service.BeginService().Start("../../conf/app_instance_info.yaml")
		if err != nil {
			t.Error("error")
		}
		defer patch.Reset()

	})

}
