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
// Code generated by angryjet. DO NOT EDIT.

package v1alpha1

import (
	"context"
	v1alpha1 "github.com/crossplane-contrib/provider-jet-tencentcloud/apis/subnet/v1alpha1"
	v1alpha11 "github.com/crossplane-contrib/provider-jet-tencentcloud/apis/vpc/v1alpha1"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this Cluster.
func (mg *Cluster) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ClusterIntranetSubnetID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ClusterIntranetSubnetIdRefs,
		Selector:     mg.Spec.ForProvider.ClusterIntranetSubnetIdSelector,
		To: reference.To{
			List:    &v1alpha1.SubnetList{},
			Managed: &v1alpha1.Subnet{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ClusterIntranetSubnetID")
	}
	mg.Spec.ForProvider.ClusterIntranetSubnetID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ClusterIntranetSubnetIdRefs = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.VPCID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.VpcIdRefs,
		Selector:     mg.Spec.ForProvider.VpcIdSelector,
		To: reference.To{
			List:    &v1alpha11.VPCList{},
			Managed: &v1alpha11.VPC{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.VPCID")
	}
	mg.Spec.ForProvider.VPCID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.VpcIdRefs = rsp.ResolvedReference

	for i3 := 0; i3 < len(mg.Spec.ForProvider.WorkerConfig); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.WorkerConfig[i3].SubnetID),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.WorkerConfig[i3].SubnetIdRefs,
			Selector:     mg.Spec.ForProvider.WorkerConfig[i3].SubnetIdSelector,
			To: reference.To{
				List:    &v1alpha1.SubnetList{},
				Managed: &v1alpha1.Subnet{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.WorkerConfig[i3].SubnetID")
		}
		mg.Spec.ForProvider.WorkerConfig[i3].SubnetID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.WorkerConfig[i3].SubnetIdRefs = rsp.ResolvedReference

	}

	return nil
}
