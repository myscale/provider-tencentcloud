/*
Copyright 2021 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by terrajet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type SubnetObservation struct {
	AvailableIPCount *float64 `json:"availableIpCount,omitempty" tf:"available_ip_count,omitempty"`

	CreateTime *string `json:"createTime,omitempty" tf:"create_time,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	IsDefault *bool `json:"isDefault,omitempty" tf:"is_default,omitempty"`
}

type SubnetParameters struct {

	// The availability zone within which the subnet should be created.
	// +kubebuilder:validation:Required
	AvailabilityZone *string `json:"availabilityZone" tf:"availability_zone,omitempty"`

	// A network address block of the subnet.
	// +kubebuilder:validation:Required
	CidrBlock *string `json:"cidrBlock" tf:"cidr_block,omitempty"`

	// Indicates whether multicast is enabled. The default value is 'true'.
	// +kubebuilder:validation:Optional
	IsMulticast *bool `json:"isMulticast,omitempty" tf:"is_multicast,omitempty"`

	// The name of subnet to be created.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// ID of a routing table to which the subnet should be associated.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-jet-tencentcloud/apis/routetable/v1alpha1.RouteTable
	// +crossplane:generate:reference:refFieldName=RouteTableIdRefs
	// +crossplane:generate:reference:selectorFieldName=RouteTableIdSelector
	// +kubebuilder:validation:Optional
	RouteTableID *string `json:"routeTableId,omitempty" tf:"route_table_id,omitempty"`

	// +kubebuilder:validation:Optional
	RouteTableIdRefs *v1.Reference `json:"routeTableIdRefs,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	RouteTableIdSelector *v1.Selector `json:"routeTableIdSelector,omitempty" tf:"-"`

	// Tags of the subnet.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// ID of the VPC to be associated.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-jet-tencentcloud/apis/vpc/v1alpha1.VPC
	// +crossplane:generate:reference:refFieldName=VpcIdRefs
	// +crossplane:generate:reference:selectorFieldName=VpcIdSelector
	// +kubebuilder:validation:Optional
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`

	// +kubebuilder:validation:Optional
	VpcIdRefs *v1.Reference `json:"vpcIdRefs,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	VpcIdSelector *v1.Selector `json:"vpcIdSelector,omitempty" tf:"-"`
}

// SubnetSpec defines the desired state of Subnet
type SubnetSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SubnetParameters `json:"forProvider"`
}

// SubnetStatus defines the observed state of Subnet.
type SubnetStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SubnetObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Subnet is the Schema for the Subnets API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tencentcloudjet}
type Subnet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SubnetSpec   `json:"spec"`
	Status            SubnetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SubnetList contains a list of Subnets
type SubnetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Subnet `json:"items"`
}

// Repository type metadata.
var (
	Subnet_Kind             = "Subnet"
	Subnet_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Subnet_Kind}.String()
	Subnet_KindAPIVersion   = Subnet_Kind + "." + CRDGroupVersion.String()
	Subnet_GroupVersionKind = CRDGroupVersion.WithKind(Subnet_Kind)
)

func init() {
	SchemeBuilder.Register(&Subnet{}, &SubnetList{})
}
