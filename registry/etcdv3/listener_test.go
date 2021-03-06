/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package etcdv3

import (
	"os"
	"testing"
	"time"
)

import (
	gxtime "github.com/dubbogo/gost/time"

	"github.com/stretchr/testify/suite"

	"go.etcd.io/etcd/server/v3/embed"
)

import (
	"dubbo.apache.org/dubbo-go/v3/common"
	"dubbo.apache.org/dubbo-go/v3/config_center"
	"dubbo.apache.org/dubbo-go/v3/remoting"
)

type RegistryTestSuite struct {
	suite.Suite
	etcd *embed.Etcd
}

const defaultEtcdV3WorkDir = "/tmp/default-dubbo-go-registry.etcd"

// start etcd server
func (suite *RegistryTestSuite) SetupSuite() {
	t := suite.T()

	cfg := embed.NewConfig()
	// avoid conflict with default etcd work-dir
	cfg.Dir = defaultEtcdV3WorkDir
	e, err := embed.StartEtcd(cfg)
	if err != nil {
		t.Fatal(err)
	}

	select {
	case <-e.Server.ReadyNotify():
		t.Log("Server is ready!")
	case <-gxtime.After(60 * time.Second):
		e.Server.Stop() // trigger a shutdown
		t.Logf("Server took too long to start!")
	}

	suite.etcd = e
}

// stop etcd server
func (suite *RegistryTestSuite) TearDownSuite() {
	suite.etcd.Close()
	// clean the etcd workdir
	if err := os.RemoveAll(defaultEtcdV3WorkDir); err != nil {
		suite.FailNow(err.Error())
	}
}

func (suite *RegistryTestSuite) TestDataChange() {
	t := suite.T()

	listener := NewRegistryDataListener(&MockDataListener{})
	url, _ := common.NewURL("jsonrpc%3A%2F%2F127.0.0.1%3A20001%2Fcom.ikurento.user.UserProvider%3Fanyhost%3Dtrue%26app.version%3D0.0.1%26application%3DBDTService%26category%3Dproviders%26cluster%3Dfailover%26dubbo%3Ddubbo-provider-golang-2.6.0%26environment%3Ddev%26group%3D%26interface%3Dcom.ikurento.user.UserProvider%26ip%3D10.32.20.124%26loadbalance%3Drandom%26methods.GetUser.loadbalance%3Drandom%26methods.GetUser.retries%3D1%26methods.GetUser.weight%3D0%26module%3Ddubbogo%2Buser-info%2Bserver%26name%3DBDTService%26organization%3Dikurento.com%26owner%3DZX%26pid%3D74500%26retries%3D0%26service.filter%3Decho%26side%3Dprovider%26timestamp%3D1560155407%26version%3D%26warmup%3D100")
	listener.AddInterestedURL(url)
	if !listener.DataChange(remoting.Event{Path: "/dubbo/com.ikurento.user.UserProvider/providers/jsonrpc%3A%2F%2F127.0.0.1%3A20001%2Fcom.ikurento.user.UserProvider%3Fanyhost%3Dtrue%26app.version%3D0.0.1%26application%3DBDTService%26category%3Dproviders%26cluster%3Dfailover%26dubbo%3Ddubbo-provider-golang-2.6.0%26environment%3Ddev%26group%3D%26interface%3Dcom.ikurento.user.UserProvider%26ip%3D10.32.20.124%26loadbalance%3Drandom%26methods.GetUser.loadbalance%3Drandom%26methods.GetUser.retries%3D1%26methods.GetUser.weight%3D0%26module%3Ddubbogo%2Buser-info%2Bserver%26name%3DBDTService%26organization%3Dikurento.com%26owner%3DZX%26pid%3D74500%26retries%3D0%26service.filter%3Decho%26side%3Dprovider%26timestamp%3D1560155407%26version%3D%26warmup%3D100"}) {
		t.Fatal("data change not ok")
	}
}

func TestRegistrySuite(t *testing.T) {
	suite.Run(t, &RegistryTestSuite{})
}

type MockDataListener struct{}

func (*MockDataListener) Process(configType *config_center.ConfigChangeEvent) {}
