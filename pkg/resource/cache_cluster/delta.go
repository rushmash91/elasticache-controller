// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Code generated by ack-generate. DO NOT EDIT.

package cache_cluster

import (
	"bytes"
	"reflect"

	ackcompare "github.com/aws-controllers-k8s/runtime/pkg/compare"
	acktags "github.com/aws-controllers-k8s/runtime/pkg/tags"
)

// Hack to avoid import errors during build...
var (
	_ = &bytes.Buffer{}
	_ = &reflect.Method{}
	_ = &acktags.Tags{}
)

// newResourceDelta returns a new `ackcompare.Delta` used to compare two
// resources
func newResourceDelta(
	a *resource,
	b *resource,
) *ackcompare.Delta {
	delta := ackcompare.NewDelta()
	if (a == nil && b != nil) ||
		(a != nil && b == nil) {
		delta.Add("", a, b)
		return delta
	}

	if ackcompare.HasNilDifference(a.ko.Spec.AZMode, b.ko.Spec.AZMode) {
		delta.Add("Spec.AZMode", a.ko.Spec.AZMode, b.ko.Spec.AZMode)
	} else if a.ko.Spec.AZMode != nil && b.ko.Spec.AZMode != nil {
		if *a.ko.Spec.AZMode != *b.ko.Spec.AZMode {
			delta.Add("Spec.AZMode", a.ko.Spec.AZMode, b.ko.Spec.AZMode)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.AuthToken, b.ko.Spec.AuthToken) {
		delta.Add("Spec.AuthToken", a.ko.Spec.AuthToken, b.ko.Spec.AuthToken)
	} else if a.ko.Spec.AuthToken != nil && b.ko.Spec.AuthToken != nil {
		if *a.ko.Spec.AuthToken != *b.ko.Spec.AuthToken {
			delta.Add("Spec.AuthToken", a.ko.Spec.AuthToken, b.ko.Spec.AuthToken)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.AutoMinorVersionUpgrade, b.ko.Spec.AutoMinorVersionUpgrade) {
		delta.Add("Spec.AutoMinorVersionUpgrade", a.ko.Spec.AutoMinorVersionUpgrade, b.ko.Spec.AutoMinorVersionUpgrade)
	} else if a.ko.Spec.AutoMinorVersionUpgrade != nil && b.ko.Spec.AutoMinorVersionUpgrade != nil {
		if *a.ko.Spec.AutoMinorVersionUpgrade != *b.ko.Spec.AutoMinorVersionUpgrade {
			delta.Add("Spec.AutoMinorVersionUpgrade", a.ko.Spec.AutoMinorVersionUpgrade, b.ko.Spec.AutoMinorVersionUpgrade)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.CacheClusterID, b.ko.Spec.CacheClusterID) {
		delta.Add("Spec.CacheClusterID", a.ko.Spec.CacheClusterID, b.ko.Spec.CacheClusterID)
	} else if a.ko.Spec.CacheClusterID != nil && b.ko.Spec.CacheClusterID != nil {
		if *a.ko.Spec.CacheClusterID != *b.ko.Spec.CacheClusterID {
			delta.Add("Spec.CacheClusterID", a.ko.Spec.CacheClusterID, b.ko.Spec.CacheClusterID)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.CacheNodeType, b.ko.Spec.CacheNodeType) {
		delta.Add("Spec.CacheNodeType", a.ko.Spec.CacheNodeType, b.ko.Spec.CacheNodeType)
	} else if a.ko.Spec.CacheNodeType != nil && b.ko.Spec.CacheNodeType != nil {
		if *a.ko.Spec.CacheNodeType != *b.ko.Spec.CacheNodeType {
			delta.Add("Spec.CacheNodeType", a.ko.Spec.CacheNodeType, b.ko.Spec.CacheNodeType)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.CacheParameterGroupName, b.ko.Spec.CacheParameterGroupName) {
		delta.Add("Spec.CacheParameterGroupName", a.ko.Spec.CacheParameterGroupName, b.ko.Spec.CacheParameterGroupName)
	} else if a.ko.Spec.CacheParameterGroupName != nil && b.ko.Spec.CacheParameterGroupName != nil {
		if *a.ko.Spec.CacheParameterGroupName != *b.ko.Spec.CacheParameterGroupName {
			delta.Add("Spec.CacheParameterGroupName", a.ko.Spec.CacheParameterGroupName, b.ko.Spec.CacheParameterGroupName)
		}
	}
	if !reflect.DeepEqual(a.ko.Spec.CacheParameterGroupRef, b.ko.Spec.CacheParameterGroupRef) {
		delta.Add("Spec.CacheParameterGroupRef", a.ko.Spec.CacheParameterGroupRef, b.ko.Spec.CacheParameterGroupRef)
	}
	if len(a.ko.Spec.CacheSecurityGroupNames) != len(b.ko.Spec.CacheSecurityGroupNames) {
		delta.Add("Spec.CacheSecurityGroupNames", a.ko.Spec.CacheSecurityGroupNames, b.ko.Spec.CacheSecurityGroupNames)
	} else if len(a.ko.Spec.CacheSecurityGroupNames) > 0 {
		if !ackcompare.SliceStringPEqual(a.ko.Spec.CacheSecurityGroupNames, b.ko.Spec.CacheSecurityGroupNames) {
			delta.Add("Spec.CacheSecurityGroupNames", a.ko.Spec.CacheSecurityGroupNames, b.ko.Spec.CacheSecurityGroupNames)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.CacheSubnetGroupName, b.ko.Spec.CacheSubnetGroupName) {
		delta.Add("Spec.CacheSubnetGroupName", a.ko.Spec.CacheSubnetGroupName, b.ko.Spec.CacheSubnetGroupName)
	} else if a.ko.Spec.CacheSubnetGroupName != nil && b.ko.Spec.CacheSubnetGroupName != nil {
		if *a.ko.Spec.CacheSubnetGroupName != *b.ko.Spec.CacheSubnetGroupName {
			delta.Add("Spec.CacheSubnetGroupName", a.ko.Spec.CacheSubnetGroupName, b.ko.Spec.CacheSubnetGroupName)
		}
	}
	if !reflect.DeepEqual(a.ko.Spec.CacheSubnetGroupRef, b.ko.Spec.CacheSubnetGroupRef) {
		delta.Add("Spec.CacheSubnetGroupRef", a.ko.Spec.CacheSubnetGroupRef, b.ko.Spec.CacheSubnetGroupRef)
	}
	if ackcompare.HasNilDifference(a.ko.Spec.Engine, b.ko.Spec.Engine) {
		delta.Add("Spec.Engine", a.ko.Spec.Engine, b.ko.Spec.Engine)
	} else if a.ko.Spec.Engine != nil && b.ko.Spec.Engine != nil {
		if *a.ko.Spec.Engine != *b.ko.Spec.Engine {
			delta.Add("Spec.Engine", a.ko.Spec.Engine, b.ko.Spec.Engine)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.EngineVersion, b.ko.Spec.EngineVersion) {
		delta.Add("Spec.EngineVersion", a.ko.Spec.EngineVersion, b.ko.Spec.EngineVersion)
	} else if a.ko.Spec.EngineVersion != nil && b.ko.Spec.EngineVersion != nil {
		if *a.ko.Spec.EngineVersion != *b.ko.Spec.EngineVersion {
			delta.Add("Spec.EngineVersion", a.ko.Spec.EngineVersion, b.ko.Spec.EngineVersion)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.IPDiscovery, b.ko.Spec.IPDiscovery) {
		delta.Add("Spec.IPDiscovery", a.ko.Spec.IPDiscovery, b.ko.Spec.IPDiscovery)
	} else if a.ko.Spec.IPDiscovery != nil && b.ko.Spec.IPDiscovery != nil {
		if *a.ko.Spec.IPDiscovery != *b.ko.Spec.IPDiscovery {
			delta.Add("Spec.IPDiscovery", a.ko.Spec.IPDiscovery, b.ko.Spec.IPDiscovery)
		}
	}
	if len(a.ko.Spec.LogDeliveryConfigurations) != len(b.ko.Spec.LogDeliveryConfigurations) {
		delta.Add("Spec.LogDeliveryConfigurations", a.ko.Spec.LogDeliveryConfigurations, b.ko.Spec.LogDeliveryConfigurations)
	} else if len(a.ko.Spec.LogDeliveryConfigurations) > 0 {
		if !reflect.DeepEqual(a.ko.Spec.LogDeliveryConfigurations, b.ko.Spec.LogDeliveryConfigurations) {
			delta.Add("Spec.LogDeliveryConfigurations", a.ko.Spec.LogDeliveryConfigurations, b.ko.Spec.LogDeliveryConfigurations)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.NetworkType, b.ko.Spec.NetworkType) {
		delta.Add("Spec.NetworkType", a.ko.Spec.NetworkType, b.ko.Spec.NetworkType)
	} else if a.ko.Spec.NetworkType != nil && b.ko.Spec.NetworkType != nil {
		if *a.ko.Spec.NetworkType != *b.ko.Spec.NetworkType {
			delta.Add("Spec.NetworkType", a.ko.Spec.NetworkType, b.ko.Spec.NetworkType)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.NotificationTopicARN, b.ko.Spec.NotificationTopicARN) {
		delta.Add("Spec.NotificationTopicARN", a.ko.Spec.NotificationTopicARN, b.ko.Spec.NotificationTopicARN)
	} else if a.ko.Spec.NotificationTopicARN != nil && b.ko.Spec.NotificationTopicARN != nil {
		if *a.ko.Spec.NotificationTopicARN != *b.ko.Spec.NotificationTopicARN {
			delta.Add("Spec.NotificationTopicARN", a.ko.Spec.NotificationTopicARN, b.ko.Spec.NotificationTopicARN)
		}
	}
	if !reflect.DeepEqual(a.ko.Spec.NotificationTopicRef, b.ko.Spec.NotificationTopicRef) {
		delta.Add("Spec.NotificationTopicRef", a.ko.Spec.NotificationTopicRef, b.ko.Spec.NotificationTopicRef)
	}
	if ackcompare.HasNilDifference(a.ko.Spec.NumCacheNodes, b.ko.Spec.NumCacheNodes) {
		delta.Add("Spec.NumCacheNodes", a.ko.Spec.NumCacheNodes, b.ko.Spec.NumCacheNodes)
	} else if a.ko.Spec.NumCacheNodes != nil && b.ko.Spec.NumCacheNodes != nil {
		if *a.ko.Spec.NumCacheNodes != *b.ko.Spec.NumCacheNodes {
			delta.Add("Spec.NumCacheNodes", a.ko.Spec.NumCacheNodes, b.ko.Spec.NumCacheNodes)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.OutpostMode, b.ko.Spec.OutpostMode) {
		delta.Add("Spec.OutpostMode", a.ko.Spec.OutpostMode, b.ko.Spec.OutpostMode)
	} else if a.ko.Spec.OutpostMode != nil && b.ko.Spec.OutpostMode != nil {
		if *a.ko.Spec.OutpostMode != *b.ko.Spec.OutpostMode {
			delta.Add("Spec.OutpostMode", a.ko.Spec.OutpostMode, b.ko.Spec.OutpostMode)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.Port, b.ko.Spec.Port) {
		delta.Add("Spec.Port", a.ko.Spec.Port, b.ko.Spec.Port)
	} else if a.ko.Spec.Port != nil && b.ko.Spec.Port != nil {
		if *a.ko.Spec.Port != *b.ko.Spec.Port {
			delta.Add("Spec.Port", a.ko.Spec.Port, b.ko.Spec.Port)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.PreferredAvailabilityZone, b.ko.Spec.PreferredAvailabilityZone) {
		delta.Add("Spec.PreferredAvailabilityZone", a.ko.Spec.PreferredAvailabilityZone, b.ko.Spec.PreferredAvailabilityZone)
	} else if a.ko.Spec.PreferredAvailabilityZone != nil && b.ko.Spec.PreferredAvailabilityZone != nil {
		if *a.ko.Spec.PreferredAvailabilityZone != *b.ko.Spec.PreferredAvailabilityZone {
			delta.Add("Spec.PreferredAvailabilityZone", a.ko.Spec.PreferredAvailabilityZone, b.ko.Spec.PreferredAvailabilityZone)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.PreferredMaintenanceWindow, b.ko.Spec.PreferredMaintenanceWindow) {
		delta.Add("Spec.PreferredMaintenanceWindow", a.ko.Spec.PreferredMaintenanceWindow, b.ko.Spec.PreferredMaintenanceWindow)
	} else if a.ko.Spec.PreferredMaintenanceWindow != nil && b.ko.Spec.PreferredMaintenanceWindow != nil {
		if *a.ko.Spec.PreferredMaintenanceWindow != *b.ko.Spec.PreferredMaintenanceWindow {
			delta.Add("Spec.PreferredMaintenanceWindow", a.ko.Spec.PreferredMaintenanceWindow, b.ko.Spec.PreferredMaintenanceWindow)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.PreferredOutpostARN, b.ko.Spec.PreferredOutpostARN) {
		delta.Add("Spec.PreferredOutpostARN", a.ko.Spec.PreferredOutpostARN, b.ko.Spec.PreferredOutpostARN)
	} else if a.ko.Spec.PreferredOutpostARN != nil && b.ko.Spec.PreferredOutpostARN != nil {
		if *a.ko.Spec.PreferredOutpostARN != *b.ko.Spec.PreferredOutpostARN {
			delta.Add("Spec.PreferredOutpostARN", a.ko.Spec.PreferredOutpostARN, b.ko.Spec.PreferredOutpostARN)
		}
	}
	if len(a.ko.Spec.PreferredOutpostARNs) != len(b.ko.Spec.PreferredOutpostARNs) {
		delta.Add("Spec.PreferredOutpostARNs", a.ko.Spec.PreferredOutpostARNs, b.ko.Spec.PreferredOutpostARNs)
	} else if len(a.ko.Spec.PreferredOutpostARNs) > 0 {
		if !ackcompare.SliceStringPEqual(a.ko.Spec.PreferredOutpostARNs, b.ko.Spec.PreferredOutpostARNs) {
			delta.Add("Spec.PreferredOutpostARNs", a.ko.Spec.PreferredOutpostARNs, b.ko.Spec.PreferredOutpostARNs)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.ReplicationGroupID, b.ko.Spec.ReplicationGroupID) {
		delta.Add("Spec.ReplicationGroupID", a.ko.Spec.ReplicationGroupID, b.ko.Spec.ReplicationGroupID)
	} else if a.ko.Spec.ReplicationGroupID != nil && b.ko.Spec.ReplicationGroupID != nil {
		if *a.ko.Spec.ReplicationGroupID != *b.ko.Spec.ReplicationGroupID {
			delta.Add("Spec.ReplicationGroupID", a.ko.Spec.ReplicationGroupID, b.ko.Spec.ReplicationGroupID)
		}
	}
	if !reflect.DeepEqual(a.ko.Spec.ReplicationGroupRef, b.ko.Spec.ReplicationGroupRef) {
		delta.Add("Spec.ReplicationGroupRef", a.ko.Spec.ReplicationGroupRef, b.ko.Spec.ReplicationGroupRef)
	}
	if len(a.ko.Spec.SecurityGroupIDs) != len(b.ko.Spec.SecurityGroupIDs) {
		delta.Add("Spec.SecurityGroupIDs", a.ko.Spec.SecurityGroupIDs, b.ko.Spec.SecurityGroupIDs)
	} else if len(a.ko.Spec.SecurityGroupIDs) > 0 {
		if !ackcompare.SliceStringPEqual(a.ko.Spec.SecurityGroupIDs, b.ko.Spec.SecurityGroupIDs) {
			delta.Add("Spec.SecurityGroupIDs", a.ko.Spec.SecurityGroupIDs, b.ko.Spec.SecurityGroupIDs)
		}
	}
	if len(a.ko.Spec.SnapshotARNs) != len(b.ko.Spec.SnapshotARNs) {
		delta.Add("Spec.SnapshotARNs", a.ko.Spec.SnapshotARNs, b.ko.Spec.SnapshotARNs)
	} else if len(a.ko.Spec.SnapshotARNs) > 0 {
		if !ackcompare.SliceStringPEqual(a.ko.Spec.SnapshotARNs, b.ko.Spec.SnapshotARNs) {
			delta.Add("Spec.SnapshotARNs", a.ko.Spec.SnapshotARNs, b.ko.Spec.SnapshotARNs)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.SnapshotName, b.ko.Spec.SnapshotName) {
		delta.Add("Spec.SnapshotName", a.ko.Spec.SnapshotName, b.ko.Spec.SnapshotName)
	} else if a.ko.Spec.SnapshotName != nil && b.ko.Spec.SnapshotName != nil {
		if *a.ko.Spec.SnapshotName != *b.ko.Spec.SnapshotName {
			delta.Add("Spec.SnapshotName", a.ko.Spec.SnapshotName, b.ko.Spec.SnapshotName)
		}
	}
	if !reflect.DeepEqual(a.ko.Spec.SnapshotRef, b.ko.Spec.SnapshotRef) {
		delta.Add("Spec.SnapshotRef", a.ko.Spec.SnapshotRef, b.ko.Spec.SnapshotRef)
	}
	if ackcompare.HasNilDifference(a.ko.Spec.SnapshotRetentionLimit, b.ko.Spec.SnapshotRetentionLimit) {
		delta.Add("Spec.SnapshotRetentionLimit", a.ko.Spec.SnapshotRetentionLimit, b.ko.Spec.SnapshotRetentionLimit)
	} else if a.ko.Spec.SnapshotRetentionLimit != nil && b.ko.Spec.SnapshotRetentionLimit != nil {
		if *a.ko.Spec.SnapshotRetentionLimit != *b.ko.Spec.SnapshotRetentionLimit {
			delta.Add("Spec.SnapshotRetentionLimit", a.ko.Spec.SnapshotRetentionLimit, b.ko.Spec.SnapshotRetentionLimit)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.SnapshotWindow, b.ko.Spec.SnapshotWindow) {
		delta.Add("Spec.SnapshotWindow", a.ko.Spec.SnapshotWindow, b.ko.Spec.SnapshotWindow)
	} else if a.ko.Spec.SnapshotWindow != nil && b.ko.Spec.SnapshotWindow != nil {
		if *a.ko.Spec.SnapshotWindow != *b.ko.Spec.SnapshotWindow {
			delta.Add("Spec.SnapshotWindow", a.ko.Spec.SnapshotWindow, b.ko.Spec.SnapshotWindow)
		}
	}
	desiredACKTags, _ := convertToOrderedACKTags(a.ko.Spec.Tags)
	latestACKTags, _ := convertToOrderedACKTags(b.ko.Spec.Tags)
	if !ackcompare.MapStringStringEqual(desiredACKTags, latestACKTags) {
		delta.Add("Spec.Tags", a.ko.Spec.Tags, b.ko.Spec.Tags)
	}
	if ackcompare.HasNilDifference(a.ko.Spec.TransitEncryptionEnabled, b.ko.Spec.TransitEncryptionEnabled) {
		delta.Add("Spec.TransitEncryptionEnabled", a.ko.Spec.TransitEncryptionEnabled, b.ko.Spec.TransitEncryptionEnabled)
	} else if a.ko.Spec.TransitEncryptionEnabled != nil && b.ko.Spec.TransitEncryptionEnabled != nil {
		if *a.ko.Spec.TransitEncryptionEnabled != *b.ko.Spec.TransitEncryptionEnabled {
			delta.Add("Spec.TransitEncryptionEnabled", a.ko.Spec.TransitEncryptionEnabled, b.ko.Spec.TransitEncryptionEnabled)
		}
	}

	modifyDelta(delta, a, b)
	return delta
}
