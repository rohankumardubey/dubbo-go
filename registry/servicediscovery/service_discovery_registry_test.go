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

package servicediscovery

//import (
//	"testing"
//)
//
//import (
//	gxset "github.com/dubbogo/gost/container/set"
//	"github.com/dubbogo/gost/hash/page"
//	"github.com/stretchr/testify/assert"
//)
//
//import (
//	"dubbo.apache.org/dubbo-go/v3/common"
//	"dubbo.apache.org/dubbo-go/v3/common/extension"
//	"dubbo.apache.org/dubbo-go/v3/config"
//	"dubbo.apache.org/dubbo-go/v3/metadata/mapping"
//	"dubbo.apache.org/dubbo-go/v3/metadata/service"
//	"dubbo.apache.org/dubbo-go/v3/registry"
//)
//
//var (
//	serviceInterface = "org.apache.dubbo.metadata.MetadataService"
//	group            = "dubbo-provider"
//	version          = "1.0.0"
//)
//
//func TestServiceDiscoveryRegistry_Register(t *testing.T) {
//	config.GetApplicationConfig().MetadataType = "mock"
//	extension.SetLocalMetadataService("mock", func() (service service.MetadataService, err error) {
//		service = &mockMetadataService{}
//		return
//	})
//
//	extension.SetServiceDiscovery("mock", func(name string) (discovery registry.ServiceDiscovery, err error) {
//		return &mockServiceDiscovery{}, nil
//	})
//
//	extension.SetGlobalServiceNameMapping(func() mapping.ServiceNameMapping {
//		return mapping.NewMockServiceNameMapping()
//	})
//
//	config.GetBaseConfig().ServiceDiscoveries["mock"] = &config.ServiceDiscoveryConfig{
//		Protocol: "mock",
//	}
//	registryURL, _ := common.NewURL("service-discovery://localhost:12345",
//		common.WithParamsValue("service_discovery", "mock"),
//		common.WithParamsValue("subscribed-services", "a, b , c,d,e ,"))
//	url, _ := common.NewURL("dubbo://192.168.0.102:20880/" + serviceInterface +
//		"?&application=" + group +
//		"&interface=" + serviceInterface +
//		"&group=" + group +
//		"&version=" + version +
//		"&service_discovery=mock" +
//		"&methods=getAllServiceKeys,getServiceRestMetadata,getExportedURLs,getAllExportedURLs" +
//		"&side=provider")
//	registry, err := newServiceDiscoveryRegistry(registryURL)
//	assert.Nil(t, err)
//	assert.NotNil(t, registry)
//	err = registry.Register(url)
//	assert.NoError(t, err)
//}
//
//type mockServiceDiscovery struct{}
//
//func (m *mockServiceDiscovery) String() string {
//	panic("implement me")
//}
//
//func (m *mockServiceDiscovery) Destroy() error {
//	panic("implement me")
//}
//
//func (m *mockServiceDiscovery) Register(registry.ServiceInstance) error {
//	return nil
//}
//
//func (m *mockServiceDiscovery) Update(registry.ServiceInstance) error {
//	panic("implement me")
//}
//
//func (m *mockServiceDiscovery) Unregister(registry.ServiceInstance) error {
//	panic("implement me")
//}
//
//func (m *mockServiceDiscovery) GetDefaultPageSize() int {
//	panic("implement me")
//}
//
//func (m *mockServiceDiscovery) GetServices() *gxset.HashSet {
//	panic("implement me")
//}
//
//func (m *mockServiceDiscovery) GetInstances(string) []registry.ServiceInstance {
//	panic("implement me")
//}
//
//func (m *mockServiceDiscovery) GetInstancesByPage(string, int, int) gxpage.Pager {
//	panic("implement me")
//}
//
//func (m *mockServiceDiscovery) GetHealthyInstancesByPage(string, int, int, bool) gxpage.Pager {
//	panic("implement me")
//}
//
//func (m *mockServiceDiscovery) GetRequestInstances([]string, int, int) map[string]gxpage.Pager {
//	panic("implement me")
//}
//
//func (m *mockServiceDiscovery) AddListener(registry.ServiceInstancesChangedListener) error {
//	panic("implement me")
//}
//
//func (m *mockServiceDiscovery) DispatchEventByServiceName(string) error {
//	panic("implement me")
//}
//
//func (m *mockServiceDiscovery) DispatchEventForInstances(string, []registry.ServiceInstance) error {
//	panic("implement me")
//}
//
//func (m *mockServiceDiscovery) DispatchEvent(*registry.ServiceInstancesChangedEvent) error {
//	panic("implement me")
//}
//
//type mockMetadataService struct{}
//
//func (m *mockMetadataService) GetExportedURLs(string, string, string, string) ([]*common.URL, error) {
//	panic("implement me")
//}
//
//func (m *mockMetadataService) GetMetadataInfo(revision string) (*common.MetadataInfo, error) {
//	panic("implement me")
//}
//
//func (m *mockMetadataService) GetExportedServiceURLs() []*common.URL {
//	panic("implement me")
//}
//
//func (m *mockMetadataService) GetMetadataServiceURL() *common.URL {
//	panic("implement me")
//}
//
//func (m *mockMetadataService) SetMetadataServiceURL(url *common.URL) {
//	panic("implement me")
//}
//
//func (m *mockMetadataService) Reference() string {
//	panic("implement me")
//}
//
//func (m *mockMetadataService) ServiceName() (string, error) {
//	panic("implement me")
//}
//
//func (m *mockMetadataService) ExportURL(*common.URL) (bool, error) {
//	return true, nil
//}
//
//func (m *mockMetadataService) UnexportURL(*common.URL) error {
//	panic("implement me")
//}
//
//func (m *mockMetadataService) SubscribeURL(*common.URL) (bool, error) {
//	panic("implement me")
//}
//
//func (m *mockMetadataService) UnsubscribeURL(*common.URL) error {
//	panic("implement me")
//}
//
//func (m *mockMetadataService) PublishServiceDefinition(*common.URL) error {
//	return nil
//}
//
//func (m *mockMetadataService) MethodMapper() map[string]string {
//	panic("implement me")
//}
//
//func (m *mockMetadataService) GetSubscribedURLs() ([]*common.URL, error) {
//	panic("implement me")
//}
//
//func (m *mockMetadataService) GetServiceDefinition(string, string, string) (string, error) {
//	panic("implement me")
//}
//
//func (m *mockMetadataService) GetServiceDefinitionByServiceKey(string) (string, error) {
//	panic("implement me")
//}
//
//func (m *mockMetadataService) RefreshMetadata(string, string) (bool, error) {
//	panic("implement me")
//}
//
//func (m *mockMetadataService) Version() (string, error) {
//	panic("implement me")
//}
