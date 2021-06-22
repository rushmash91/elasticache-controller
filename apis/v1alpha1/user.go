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

// UserSpec defines the desired state of User.
//

type UserSpec struct {
	// Access permissions string used for this user.
	// +kubebuilder:validation:Required
	AccessString *string `json:"accessString"`
	// The current supported value is Redis.
	// +kubebuilder:validation:Required
	Engine *string `json:"engine"`
	// Indicates a password is not required for this user.
	NoPasswordRequired *bool `json:"noPasswordRequired,omitempty"`
	// The ID of the user.
	// +kubebuilder:validation:Required
	UserID *string `json:"userID"`
	// The username of the user.
	// +kubebuilder:validation:Required
	UserName *string `json:"userName"`
}

// UserStatus defines the observed state of User
type UserStatus struct {
	// All CRs managed by ACK have a common `Status.ACKResourceMetadata` member
	// that is used to contain resource sync state, account ownership,
	// constructed ARN for the resource
	ACKResourceMetadata *ackv1alpha1.ResourceMetadata `json:"ackResourceMetadata"`
	// All CRS managed by ACK have a common `Status.Conditions` member that
	// contains a collection of `ackv1alpha1.Condition` objects that describe
	// the various terminal states of the CR and its backend AWS service API
	// resource
	Conditions []*ackv1alpha1.Condition `json:"conditions"`
	// Denotes whether the user requires a password to authenticate.
	Authentication *Authentication `json:"authentication,omitempty"`
	// Access permissions string used for this user.
	ExpandedAccessString *string `json:"expandedAccessString,omitempty"`
	// Access permissions string used for this user.
	LastRequestedAccessString *string `json:"lastRequestedAccessString,omitempty"`
	// Indicates the user status. Can be "active", "modifying" or "deleting".
	Status *string `json:"status,omitempty"`
	// Returns a list of the user group IDs the user belongs to.
	UserGroupIDs []*string `json:"userGroupIDs,omitempty"`
}

// User is the Schema for the Users API
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
type User struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              UserSpec   `json:"spec,omitempty"`
	Status            UserStatus `json:"status,omitempty"`
}

// UserList contains a list of User
// +kubebuilder:object:root=true
type UserList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []User `json:"items"`
}

func init() {
	SchemeBuilder.Register(&User{}, &UserList{})
}
