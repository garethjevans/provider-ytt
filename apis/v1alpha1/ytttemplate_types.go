/*
Copyright 2022 The Crossplane Authors.

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

package v1alpha1

import (
	"reflect"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

// YttTemplateParameters are the configurable fields of a YttTemplate.
type YttTemplateParameters struct {
	ConfigurableField string `json:"configurableField"`
}

// YttTemplateObservation are the observable fields of a YttTemplate.
type YttTemplateObservation struct {
	ObservableField string `json:"observableField,omitempty"`
}

// A YttTemplateSpec defines the desired state of a YttTemplate.
type YttTemplateSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       YttTemplateParameters `json:"forProvider"`
	Ytt               string                `json:"ytt"`
}

// A YttTemplateStatus represents the observed state of a YttTemplate.
type YttTemplateStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          YttTemplateObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// A YttTemplate is an example API type.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,ytt}
type YttTemplate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   YttTemplateSpec   `json:"spec"`
	Status YttTemplateStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// YttTemplateList contains a list of YttTemplate
type YttTemplateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []YttTemplate `json:"items"`
}

// YttTemplate type metadata.
var (
	YttTemplateKind             = reflect.TypeOf(YttTemplate{}).Name()
	YttTemplateGroupKind        = schema.GroupKind{Group: Group, Kind: YttTemplateKind}.String()
	YttTemplateKindAPIVersion   = YttTemplateKind + "." + SchemeGroupVersion.String()
	YttTemplateGroupVersionKind = SchemeGroupVersion.WithKind(YttTemplateKind)
)

func init() {
	SchemeBuilder.Register(&YttTemplate{}, &YttTemplateList{})
}
