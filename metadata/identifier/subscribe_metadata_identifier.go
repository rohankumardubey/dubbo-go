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

package identifier

import (
	"dubbo.apache.org/dubbo-go/v3/common/constant"
)

// SubscriberMetadataIdentifier is inherit baseMetaIdentifier with service params: Revision
type SubscriberMetadataIdentifier struct {
	Revision string
	BaseApplicationMetadataIdentifier
}

func NewSubscriberMetadataIdentifier(application string, revision string) *SubscriberMetadataIdentifier {
	return &SubscriberMetadataIdentifier{
		Revision: revision,
		BaseApplicationMetadataIdentifier: BaseApplicationMetadataIdentifier{
			Application: application,
			Group:       constant.Dubbo,
		},
	}
}

// GetIdentifierKey returns string that format is service:Version:Group:Side:Revision
func (mdi *SubscriberMetadataIdentifier) GetIdentifierKey() string {
	return mdi.BaseApplicationMetadataIdentifier.getIdentifierKey(mdi.Revision)
}

// GetFilePathKey returns string that format is metadata/path/Version/Group/Side/Revision
func (mdi *SubscriberMetadataIdentifier) GetFilePathKey() string {
	return mdi.BaseApplicationMetadataIdentifier.getFilePathKey(mdi.Revision)
}
