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

package v1alpha1

import (
	ackv1alpha1 "github.com/aws-controllers-k8s/runtime/apis/core/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// SnapshotSpec defines the desired state of Snapshot.
//
// Represents a copy of an entire Valkey or Redis OSS cluster as of the time
// when the snapshot was taken.
type SnapshotSpec struct {

	// The identifier of an existing cluster. The snapshot is created from this
	// cluster.
	CacheClusterID *string `json:"cacheClusterID,omitempty"`
	// The ID of the KMS key used to encrypt the snapshot.
	KMSKeyID *string `json:"kmsKeyID,omitempty"`
	// The identifier of an existing replication group. The snapshot is created
	// from this replication group.
	ReplicationGroupID *string `json:"replicationGroupID,omitempty"`
	// A name for the snapshot being created.
	// +kubebuilder:validation:Required
	SnapshotName *string `json:"snapshotName"`
	// The name of an existing snapshot from which to make a copy.
	SourceSnapshotName *string `json:"sourceSnapshotName,omitempty"`
	// A list of tags to be added to this resource. A tag is a key-value pair. A
	// tag key must be accompanied by a tag value, although null is accepted.
	Tags []*Tag `json:"tags,omitempty"`
}

// SnapshotStatus defines the observed state of Snapshot
type SnapshotStatus struct {
	// All CRs managed by ACK have a common `Status.ACKResourceMetadata` member
	// that is used to contain resource sync state, account ownership,
	// constructed ARN for the resource
	// +kubebuilder:validation:Optional
	ACKResourceMetadata *ackv1alpha1.ResourceMetadata `json:"ackResourceMetadata"`
	// All CRs managed by ACK have a common `Status.Conditions` member that
	// contains a collection of `ackv1alpha1.Condition` objects that describe
	// the various terminal states of the CR and its backend AWS service API
	// resource
	// +kubebuilder:validation:Optional
	Conditions []*ackv1alpha1.Condition `json:"conditions"`
	// If you are running Valkey 7.2 and above or Redis OSS engine version 6.0 and
	// above, set this parameter to yes if you want to opt-in to the next auto minor
	// version upgrade campaign. This parameter is disabled for previous versions.
	// +kubebuilder:validation:Optional
	AutoMinorVersionUpgrade *bool `json:"autoMinorVersionUpgrade,omitempty"`
	// Indicates the status of automatic failover for the source Valkey or Redis
	// OSS replication group.
	// +kubebuilder:validation:Optional
	AutomaticFailover *string `json:"automaticFailover,omitempty"`
	// The date and time when the source cluster was created.
	// +kubebuilder:validation:Optional
	CacheClusterCreateTime *metav1.Time `json:"cacheClusterCreateTime,omitempty"`
	// The name of the compute and memory capacity node type for the source cluster.
	//
	// The following node types are supported by ElastiCache. Generally speaking,
	// the current generation types provide more memory and computational power
	// at lower cost when compared to their equivalent previous generation counterparts.
	//
	//    * General purpose: Current generation: M7g node types: cache.m7g.large,
	//    cache.m7g.xlarge, cache.m7g.2xlarge, cache.m7g.4xlarge, cache.m7g.8xlarge,
	//    cache.m7g.12xlarge, cache.m7g.16xlarge For region availability, see Supported
	//    Node Types (https://docs.aws.amazon.com/AmazonElastiCache/latest/dg/CacheNodes.SupportedTypes.html#CacheNodes.SupportedTypesByRegion)
	//    M6g node types (available only for Redis OSS engine version 5.0.6 onward
	//    and for Memcached engine version 1.5.16 onward): cache.m6g.large, cache.m6g.xlarge,
	//    cache.m6g.2xlarge, cache.m6g.4xlarge, cache.m6g.8xlarge, cache.m6g.12xlarge,
	//    cache.m6g.16xlarge M5 node types: cache.m5.large, cache.m5.xlarge, cache.m5.2xlarge,
	//    cache.m5.4xlarge, cache.m5.12xlarge, cache.m5.24xlarge M4 node types:
	//    cache.m4.large, cache.m4.xlarge, cache.m4.2xlarge, cache.m4.4xlarge, cache.m4.10xlarge
	//    T4g node types (available only for Redis OSS engine version 5.0.6 onward
	//    and Memcached engine version 1.5.16 onward): cache.t4g.micro, cache.t4g.small,
	//    cache.t4g.medium T3 node types: cache.t3.micro, cache.t3.small, cache.t3.medium
	//    T2 node types: cache.t2.micro, cache.t2.small, cache.t2.medium Previous
	//    generation: (not recommended. Existing clusters are still supported but
	//    creation of new clusters is not supported for these types.) T1 node types:
	//    cache.t1.micro M1 node types: cache.m1.small, cache.m1.medium, cache.m1.large,
	//    cache.m1.xlarge M3 node types: cache.m3.medium, cache.m3.large, cache.m3.xlarge,
	//    cache.m3.2xlarge
	//
	//    * Compute optimized: Previous generation: (not recommended. Existing clusters
	//    are still supported but creation of new clusters is not supported for
	//    these types.) C1 node types: cache.c1.xlarge
	//
	//    * Memory optimized: Current generation: R7g node types: cache.r7g.large,
	//    cache.r7g.xlarge, cache.r7g.2xlarge, cache.r7g.4xlarge, cache.r7g.8xlarge,
	//    cache.r7g.12xlarge, cache.r7g.16xlarge For region availability, see Supported
	//    Node Types (https://docs.aws.amazon.com/AmazonElastiCache/latest/dg/CacheNodes.SupportedTypes.html#CacheNodes.SupportedTypesByRegion)
	//    R6g node types (available only for Redis OSS engine version 5.0.6 onward
	//    and for Memcached engine version 1.5.16 onward): cache.r6g.large, cache.r6g.xlarge,
	//    cache.r6g.2xlarge, cache.r6g.4xlarge, cache.r6g.8xlarge, cache.r6g.12xlarge,
	//    cache.r6g.16xlarge R5 node types: cache.r5.large, cache.r5.xlarge, cache.r5.2xlarge,
	//    cache.r5.4xlarge, cache.r5.12xlarge, cache.r5.24xlarge R4 node types:
	//    cache.r4.large, cache.r4.xlarge, cache.r4.2xlarge, cache.r4.4xlarge, cache.r4.8xlarge,
	//    cache.r4.16xlarge Previous generation: (not recommended. Existing clusters
	//    are still supported but creation of new clusters is not supported for
	//    these types.) M2 node types: cache.m2.xlarge, cache.m2.2xlarge, cache.m2.4xlarge
	//    R3 node types: cache.r3.large, cache.r3.xlarge, cache.r3.2xlarge, cache.r3.4xlarge,
	//    cache.r3.8xlarge
	//
	// Additional node type info
	//
	//    * All current generation instance types are created in Amazon VPC by default.
	//
	//    * Valkey or Redis OSS append-only files (AOF) are not supported for T1
	//    or T2 instances.
	//
	//    * Valkey or Redis OSS Multi-AZ with automatic failover is not supported
	//    on T1 instances.
	//
	//    * The configuration variables appendonly and appendfsync are not supported
	//    on Valkey, or on Redis OSS version 2.8.22 and later.
	// +kubebuilder:validation:Optional
	CacheNodeType *string `json:"cacheNodeType,omitempty"`
	// The cache parameter group that is associated with the source cluster.
	// +kubebuilder:validation:Optional
	CacheParameterGroupName *string `json:"cacheParameterGroupName,omitempty"`
	// The name of the cache subnet group associated with the source cluster.
	// +kubebuilder:validation:Optional
	CacheSubnetGroupName *string `json:"cacheSubnetGroupName,omitempty"`
	// Enables data tiering. Data tiering is only supported for replication groups
	// using the r6gd node type. This parameter must be set to true when using r6gd
	// nodes. For more information, see Data tiering (https://docs.aws.amazon.com/AmazonElastiCache/latest/dg/data-tiering.html).
	// +kubebuilder:validation:Optional
	DataTiering *string `json:"dataTiering,omitempty"`
	// The name of the cache engine (memcached or redis) used by the source cluster.
	// +kubebuilder:validation:Optional
	Engine *string `json:"engine,omitempty"`
	// The version of the cache engine version that is used by the source cluster.
	// +kubebuilder:validation:Optional
	EngineVersion *string `json:"engineVersion,omitempty"`
	// A list of the cache nodes in the source cluster.
	// +kubebuilder:validation:Optional
	NodeSnapshots []*NodeSnapshot `json:"nodeSnapshots,omitempty"`
	// The number of cache nodes in the source cluster.
	//
	// For clusters running Valkey or Redis OSS, this value must be 1. For clusters
	// running Memcached, this value must be between 1 and 40.
	// +kubebuilder:validation:Optional
	NumCacheNodes *int64 `json:"numCacheNodes,omitempty"`
	// The number of node groups (shards) in this snapshot. When restoring from
	// a snapshot, the number of node groups (shards) in the snapshot and in the
	// restored replication group must be the same.
	// +kubebuilder:validation:Optional
	NumNodeGroups *int64 `json:"numNodeGroups,omitempty"`
	// The port number used by each cache nodes in the source cluster.
	// +kubebuilder:validation:Optional
	Port *int64 `json:"port,omitempty"`
	// The name of the Availability Zone in which the source cluster is located.
	// +kubebuilder:validation:Optional
	PreferredAvailabilityZone *string `json:"preferredAvailabilityZone,omitempty"`
	// Specifies the weekly time range during which maintenance on the cluster is
	// performed. It is specified as a range in the format ddd:hh24:mi-ddd:hh24:mi
	// (24H Clock UTC). The minimum maintenance window is a 60 minute period.
	//
	// Valid values for ddd are:
	//
	//    * sun
	//
	//    * mon
	//
	//    * tue
	//
	//    * wed
	//
	//    * thu
	//
	//    * fri
	//
	//    * sat
	//
	// Example: sun:23:00-mon:01:30
	// +kubebuilder:validation:Optional
	PreferredMaintenanceWindow *string `json:"preferredMaintenanceWindow,omitempty"`
	// The ARN (Amazon Resource Name) of the preferred outpost.
	// +kubebuilder:validation:Optional
	PreferredOutpostARN *string `json:"preferredOutpostARN,omitempty"`
	// A description of the source replication group.
	// +kubebuilder:validation:Optional
	ReplicationGroupDescription *string `json:"replicationGroupDescription,omitempty"`
	// For an automatic snapshot, the number of days for which ElastiCache retains
	// the snapshot before deleting it.
	//
	// For manual snapshots, this field reflects the SnapshotRetentionLimit for
	// the source cluster when the snapshot was created. This field is otherwise
	// ignored: Manual snapshots do not expire, and can only be deleted using the
	// DeleteSnapshot operation.
	//
	// Important If the value of SnapshotRetentionLimit is set to zero (0), backups
	// are turned off.
	// +kubebuilder:validation:Optional
	SnapshotRetentionLimit *int64 `json:"snapshotRetentionLimit,omitempty"`
	// Indicates whether the snapshot is from an automatic backup (automated) or
	// was created manually (manual).
	// +kubebuilder:validation:Optional
	SnapshotSource *string `json:"snapshotSource,omitempty"`
	// The status of the snapshot. Valid values: creating | available | restoring
	// | copying | deleting.
	// +kubebuilder:validation:Optional
	SnapshotStatus *string `json:"snapshotStatus,omitempty"`
	// The daily time range during which ElastiCache takes daily snapshots of the
	// source cluster.
	// +kubebuilder:validation:Optional
	SnapshotWindow *string `json:"snapshotWindow,omitempty"`
	// The Amazon Resource Name (ARN) for the topic used by the source cluster for
	// publishing notifications.
	// +kubebuilder:validation:Optional
	TopicARN *string `json:"topicARN,omitempty"`
	// The Amazon Virtual Private Cloud identifier (VPC ID) of the cache subnet
	// group for the source cluster.
	// +kubebuilder:validation:Optional
	VPCID *string `json:"vpcID,omitempty"`
}

// Snapshot is the Schema for the Snapshots API
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
type Snapshot struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SnapshotSpec   `json:"spec,omitempty"`
	Status            SnapshotStatus `json:"status,omitempty"`
}

// SnapshotList contains a list of Snapshot
// +kubebuilder:object:root=true
type SnapshotList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Snapshot `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Snapshot{}, &SnapshotList{})
}
