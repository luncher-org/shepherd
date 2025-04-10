/*
Copyright 2025 Rancher Labs, Inc.

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

// Code generated by main. DO NOT EDIT.

package v1

import (
	context "context"

	scheme "github.com/rancher/shepherd/pkg/generated/clientset/versioned/scheme"
	upgradecattleiov1 "github.com/rancher/system-upgrade-controller/pkg/apis/upgrade.cattle.io/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	gentype "k8s.io/client-go/gentype"
)

// PlansGetter has a method to return a PlanInterface.
// A group's client should implement this interface.
type PlansGetter interface {
	Plans(namespace string) PlanInterface
}

// PlanInterface has methods to work with Plan resources.
type PlanInterface interface {
	Create(ctx context.Context, plan *upgradecattleiov1.Plan, opts metav1.CreateOptions) (*upgradecattleiov1.Plan, error)
	Update(ctx context.Context, plan *upgradecattleiov1.Plan, opts metav1.UpdateOptions) (*upgradecattleiov1.Plan, error)
	// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
	UpdateStatus(ctx context.Context, plan *upgradecattleiov1.Plan, opts metav1.UpdateOptions) (*upgradecattleiov1.Plan, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*upgradecattleiov1.Plan, error)
	List(ctx context.Context, opts metav1.ListOptions) (*upgradecattleiov1.PlanList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *upgradecattleiov1.Plan, err error)
	PlanExpansion
}

// plans implements PlanInterface
type plans struct {
	*gentype.ClientWithList[*upgradecattleiov1.Plan, *upgradecattleiov1.PlanList]
}

// newPlans returns a Plans
func newPlans(c *UpgradeV1Client, namespace string) *plans {
	return &plans{
		gentype.NewClientWithList[*upgradecattleiov1.Plan, *upgradecattleiov1.PlanList](
			"plans",
			c.RESTClient(),
			scheme.ParameterCodec,
			namespace,
			func() *upgradecattleiov1.Plan { return &upgradecattleiov1.Plan{} },
			func() *upgradecattleiov1.PlanList { return &upgradecattleiov1.PlanList{} },
		),
	}
}
