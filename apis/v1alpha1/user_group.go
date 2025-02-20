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

// UserGroupSpec defines the desired state of UserGroup.
//

type UserGroupSpec struct {

	// The current supported value is Redis user.
	// +kubebuilder:validation:Required
	Engine *string `json:"engine"`
	// A list of tags to be added to this resource. A tag is a key-value pair. A
	// tag key must be accompanied by a tag value, although null is accepted. Available
	// for Valkey and Redis OSS only.
	Tags []*Tag `json:"tags,omitempty"`
	// The ID of the user group.
	// +kubebuilder:validation:Required
	UserGroupID *string `json:"userGroupID"`
	// The list of user IDs that belong to the user group.
	UserIDs []*string `json:"userIDs,omitempty"`
}

// UserGroupStatus defines the observed state of UserGroup
type UserGroupStatus struct {
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
	// The minimum engine version required, which is Redis OSS 6.0
	// +kubebuilder:validation:Optional
	MinimumEngineVersion *string `json:"minimumEngineVersion,omitempty"`
	// A list of updates being applied to the user group.
	// +kubebuilder:validation:Optional
	PendingChanges *UserGroupPendingChanges `json:"pendingChanges,omitempty"`
	// A list of replication groups that the user group can access.
	// +kubebuilder:validation:Optional
	ReplicationGroups []*string `json:"replicationGroups,omitempty"`
	// Indicates user group status. Can be "creating", "active", "modifying", "deleting".
	// +kubebuilder:validation:Optional
	Status *string `json:"status,omitempty"`
}

// UserGroup is the Schema for the UserGroups API
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
type UserGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              UserGroupSpec   `json:"spec,omitempty"`
	Status            UserGroupStatus `json:"status,omitempty"`
}

// UserGroupList contains a list of UserGroup
// +kubebuilder:object:root=true
type UserGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []UserGroup `json:"items"`
}

func init() {
	SchemeBuilder.Register(&UserGroup{}, &UserGroupList{})
}
