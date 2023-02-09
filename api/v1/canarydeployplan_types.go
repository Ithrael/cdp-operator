/*
Copyright 2023.

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

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// CanaryDeployPlanSpec defines the desired state of CanaryDeployPlan
type CanaryDeployPlanSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of CanaryDeployPlan. Edit canarydeployplan_types.go to remove/update
	Service     string       `json:"service"`
	Policies    PoliciesSpec `json:"policies"`
	Destination string       `json:"destination,omitempty"`
}

type PoliciesSpec struct {
	Combine CombineSpec `json:"combine"`
}

type CombineSpec struct {
	Must []PolicySpec `json:"must"`
}

type PolicySpec struct {
	Policy string `json:"policy"`
	Op     string `json:"op"` // contains/Is/range/mod/weight
	Value  string `json:"value"`
}

// CanaryDeployPlanStatus defines the observed state of CanaryDeployPlan
type CanaryDeployPlanStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// CanaryDeployPlan is the Schema for the canarydeployplans API
type CanaryDeployPlan struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CanaryDeployPlanSpec   `json:"spec,omitempty"`
	Status CanaryDeployPlanStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// CanaryDeployPlanList contains a list of CanaryDeployPlan
type CanaryDeployPlanList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CanaryDeployPlan `json:"items"`
}

func init() {
	SchemeBuilder.Register(&CanaryDeployPlan{}, &CanaryDeployPlanList{})
}
