// Code generated by cue get go. DO NOT EDIT.

//cue:generate cue get go k8s.io/api/discovery/v1alpha1

package v1alpha1

// LabelServiceName is used to indicate the name of a Kubernetes service.
#LabelServiceName: "kubernetes.io/service-name"

// LabelManagedBy is used to indicate the controller or entity that manages
// an EndpointSlice. This label aims to enable different EndpointSlice
// objects to be managed by different controllers or entities within the
// same cluster. It is highly recommended to configure this label for all
// EndpointSlices.
#LabelManagedBy: "endpointslice.kubernetes.io/managed-by"
