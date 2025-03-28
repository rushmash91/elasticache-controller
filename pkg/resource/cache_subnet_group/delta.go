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

package cache_subnet_group

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

	if ackcompare.HasNilDifference(a.ko.Spec.CacheSubnetGroupDescription, b.ko.Spec.CacheSubnetGroupDescription) {
		delta.Add("Spec.CacheSubnetGroupDescription", a.ko.Spec.CacheSubnetGroupDescription, b.ko.Spec.CacheSubnetGroupDescription)
	} else if a.ko.Spec.CacheSubnetGroupDescription != nil && b.ko.Spec.CacheSubnetGroupDescription != nil {
		if *a.ko.Spec.CacheSubnetGroupDescription != *b.ko.Spec.CacheSubnetGroupDescription {
			delta.Add("Spec.CacheSubnetGroupDescription", a.ko.Spec.CacheSubnetGroupDescription, b.ko.Spec.CacheSubnetGroupDescription)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.CacheSubnetGroupName, b.ko.Spec.CacheSubnetGroupName) {
		delta.Add("Spec.CacheSubnetGroupName", a.ko.Spec.CacheSubnetGroupName, b.ko.Spec.CacheSubnetGroupName)
	} else if a.ko.Spec.CacheSubnetGroupName != nil && b.ko.Spec.CacheSubnetGroupName != nil {
		if *a.ko.Spec.CacheSubnetGroupName != *b.ko.Spec.CacheSubnetGroupName {
			delta.Add("Spec.CacheSubnetGroupName", a.ko.Spec.CacheSubnetGroupName, b.ko.Spec.CacheSubnetGroupName)
		}
	}
	if len(a.ko.Spec.SubnetIDs) != len(b.ko.Spec.SubnetIDs) {
		delta.Add("Spec.SubnetIDs", a.ko.Spec.SubnetIDs, b.ko.Spec.SubnetIDs)
	} else if len(a.ko.Spec.SubnetIDs) > 0 {
		if !ackcompare.SliceStringPEqual(a.ko.Spec.SubnetIDs, b.ko.Spec.SubnetIDs) {
			delta.Add("Spec.SubnetIDs", a.ko.Spec.SubnetIDs, b.ko.Spec.SubnetIDs)
		}
	}
	if !reflect.DeepEqual(a.ko.Spec.SubnetRefs, b.ko.Spec.SubnetRefs) {
		delta.Add("Spec.SubnetRefs", a.ko.Spec.SubnetRefs, b.ko.Spec.SubnetRefs)
	}
	desiredACKTags, _ := convertToOrderedACKTags(a.ko.Spec.Tags)
	latestACKTags, _ := convertToOrderedACKTags(b.ko.Spec.Tags)
	if !ackcompare.MapStringStringEqual(desiredACKTags, latestACKTags) {
		delta.Add("Spec.Tags", a.ko.Spec.Tags, b.ko.Spec.Tags)
	}

	return delta
}
