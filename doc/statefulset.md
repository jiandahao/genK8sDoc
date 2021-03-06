# statefulset
 statefulset Object
**PATH:**  statefulset

KIND:     StatefulSet
VERSION:  apps/v1

DESCRIPTION:
     StatefulSet represents a set of pods with consistent identities. Identities
     are defined as: - Network: A single stable DNS and hostname. - Storage: As
     many VolumeClaims as requested. The StatefulSet guarantees that a given
     network identity will always map to the same storage identity.

# apiVersion	\<string\>
**PATH:**  statefulset.apiVersion

     APIVersion defines the versioned schema of this representation of an
     object. Servers should convert recognized schemas to the latest internal
     value, and may reject unrecognized values. More info:
     https://git.k8s.io/community/contributors/devel/api-conventions.md#resources

# kind	\<string\>
**PATH:**  statefulset.kind

     Kind is a string value representing the REST resource this object
     represents. Servers may infer this from the endpoint the client submits
     requests to. Cannot be updated. In CamelCase. More info:
     https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds

# metadata \<Object\>
**PATH:**  statefulset.metadata

KIND:     StatefulSet
VERSION:  apps/v1

RESOURCE: metadata <Object>

DESCRIPTION:

     ObjectMeta is metadata that all persisted resources must have, which
     includes all objects users must create.

## annotations	\<map[string]string\>
**PATH:**  statefulset.metadata.annotations

     Annotations is an unstructured key value map stored with a resource that
     may be set by external tools to store and retrieve arbitrary metadata. They
     are not queryable and should be preserved when modifying objects. More
     info: http://kubernetes.io/docs/user-guide/annotations

## clusterName	\<string\>
**PATH:**  statefulset.metadata.clusterName

     The name of the cluster which the object belongs to. This is used to
     distinguish resources with same name and namespace in different clusters.
     This field is not set anywhere right now and apiserver is going to ignore
     it if set in create or update request.

## creationTimestamp	\<string\>
**PATH:**  statefulset.metadata.creationTimestamp

     CreationTimestamp is a timestamp representing the server time when this
     object was created. It is not guaranteed to be set in happens-before order
     across separate operations. Clients may not set this value. It is
     represented in RFC3339 form and is in UTC. Populated by the system.
     Read-only. Null for lists. More info:
     https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata

## deletionGracePeriodSeconds	\<integer\>
**PATH:**  statefulset.metadata.deletionGracePeriodSeconds

     Number of seconds allowed for this object to gracefully terminate before it
     will be removed from the system. Only set when deletionTimestamp is also
     set. May only be shortened. Read-only.

## deletionTimestamp	\<string\>
**PATH:**  statefulset.metadata.deletionTimestamp

     DeletionTimestamp is RFC 3339 date and time at which this resource will be
     deleted. This field is set by the server when a graceful deletion is
     requested by the user, and is not directly settable by a client. The
     resource is expected to be deleted (no longer visible from resource lists,
     and not reachable by name) after the time in this field, once the
     finalizers list is empty. As long as the finalizers list contains items,
     deletion is blocked. Once the deletionTimestamp is set, this value may not
     be unset or be set further into the future, although it may be shortened or
     the resource may be deleted prior to this time. For example, a user may
     request that a pod is deleted in 30 seconds. The Kubelet will react by
     sending a graceful termination signal to the containers in the pod. After
     that 30 seconds, the Kubelet will send a hard termination signal (SIGKILL)
     to the container and after cleanup, remove the pod from the API. In the
     presence of network partitions, this object may still exist after this
     timestamp, until an administrator or automated process can determine the
     resource is fully terminated. If not set, graceful deletion of the object
     has not been requested. Populated by the system when a graceful deletion is
     requested. Read-only. More info:
     https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata

## finalizers	\<[]string\>
**PATH:**  statefulset.metadata.finalizers

     Must be empty before the object is deleted from the registry. Each entry is
     an identifier for the responsible component that will remove the entry from
     the list. If the deletionTimestamp of the object is non-nil, entries in
     this list can only be removed.

## generateName	\<string\>
**PATH:**  statefulset.metadata.generateName

     GenerateName is an optional prefix, used by the server, to generate a
     unique name ONLY IF the Name field has not been provided. If this field is
     used, the name returned to the client will be different than the name
     passed. This value will also be combined with a unique suffix. The provided
     value has the same validation rules as the Name field, and may be truncated
     by the length of the suffix required to make the value unique on the
     server. If this field is specified and the generated name exists, the
     server will NOT return a 409 - instead, it will either return 201 Created
     or 500 with Reason ServerTimeout indicating a unique name could not be
     found in the time allotted, and the client should retry (optionally after
     the time indicated in the Retry-After header). Applied only if Name is not
     specified. More info:
     https://git.k8s.io/community/contributors/devel/api-conventions.md#idempotency

## generation	\<integer\>
**PATH:**  statefulset.metadata.generation

     A sequence number representing a specific generation of the desired state.
     Populated by the system. Read-only.

## initializers \<Object\>
**PATH:**  statefulset.metadata.initializers

KIND:     StatefulSet
VERSION:  apps/v1

RESOURCE: initializers <Object>

DESCRIPTION:
     An initializer is a controller which enforces some system invariant at
     object creation time. This field is a list of initializers that have not
     yet acted on this object. If nil or empty, this object has been completely
     initialized. Otherwise, the object is considered uninitialized and is
     hidden (in list/watch and get calls) from clients that haven't explicitly
     asked to observe uninitialized objects. When an object is created, the
     system will populate this list with the current set of initializers. Only
     privileged users may set or modify this list. Once it is empty, it may not
     be modified further by any user. DEPRECATED - initializers are an alpha
     field and will be removed in v1.15.

     Initializers tracks the progress of initialization.

### pending \<[]Object\>
**PATH:**  statefulset.metadata.initializers.pending

KIND:     StatefulSet
VERSION:  apps/v1

RESOURCE: pending <[]Object>

DESCRIPTION:
     Pending is a list of initializers that must execute in order before this
     object is visible. When the last pending initializer is removed, and no
     failing result is set, the initializers struct will be set to nil and the
     object is considered as initialized and visible to all clients.

     Initializer is information about an initializer that has not yet completed.

### pending	\<[]Object\> -required-
**PATH:**  statefulset.metadata.initializers.pending

     Pending is a list of initializers that must execute in order before this
     object is visible. When the last pending initializer is removed, and no
     failing result is set, the initializers struct will be set to nil and the
     object is considered as initialized and visible to all clients.

### result \<Object\>
**PATH:**  statefulset.metadata.initializers.result

KIND:     StatefulSet
VERSION:  apps/v1

RESOURCE: result <Object>

DESCRIPTION:
     If result is set with the Failure field, the object will be persisted to
     storage and then deleted, ensuring that other clients can observe the
     deletion.

     Status is a return value for calls that don't return other objects.

#### apiVersion	\<string\>
**PATH:**  statefulset.metadata.initializers.result.apiVersion

     APIVersion defines the versioned schema of this representation of an
     object. Servers should convert recognized schemas to the latest internal
     value, and may reject unrecognized values. More info:
     https://git.k8s.io/community/contributors/devel/api-conventions.md#resources

#### code	\<integer\>
**PATH:**  statefulset.metadata.initializers.result.code

     Suggested HTTP return code for this status, 0 if not set.

#### details \<Object\>
**PATH:**  statefulset.metadata.initializers.result.details

KIND:     StatefulSet
VERSION:  apps/v1

RESOURCE: details <Object>

DESCRIPTION:
     Extended data associated with the reason. Each reason may define its own
     extended details. This field is optional and the data returned is not
     guaranteed to conform to any schema except that defined by the reason type.

     StatusDetails is a set of additional properties that MAY be set by the
     server to provide additional information about a response. The Reason field
     of a Status object defines what attributes will be set. Clients must ignore
     fields that do not match the defined type of each attribute, and should
     assume that any attribute may be empty, invalid, or under defined.

##### causes \<[]Object\>
**PATH:**  statefulset.metadata.initializers.result.details.causes

KIND:     StatefulSet
VERSION:  apps/v1

RESOURCE: causes <[]Object>

DESCRIPTION:
     The Causes array includes more details associated with the StatusReason
     failure. Not all StatusReasons may provide detailed causes.

     StatusCause provides more information about an api.Status failure,
     including cases when multiple errors are encountered.

###### field	\<string\>
**PATH:**  statefulset.metadata.initializers.result.details.causes.field

     The field of the resource that has caused this error, as named by its JSON
     serialization. May include dot and postfix notation for nested attributes.
     Arrays are zero-indexed. Fields may appear more than once in an array of
     causes due to fields having multiple errors. Optional. Examples: "name" -
     the field "name" on the current resource "items[0].name" - the field "name"
     on the first array entry in "items"

###### message	\<string\>
**PATH:**  statefulset.metadata.initializers.result.details.causes.message

     A human-readable description of the cause of the error. This field may be
     presented as-is to a reader.

##### causes	\<[]Object\>
**PATH:**  statefulset.metadata.initializers.result.details.causes

     The Causes array includes more details associated with the StatusReason
     failure. Not all StatusReasons may provide detailed causes.

##### group	\<string\>
**PATH:**  statefulset.metadata.initializers.result.details.group

     The group attribute of the resource associated with the status
     StatusReason.

##### kind	\<string\>
**PATH:**  statefulset.metadata.initializers.result.details.kind

     The kind attribute of the resource associated with the status StatusReason.
     On some operations may differ from the requested resource Kind. More info:
     https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds

##### name	\<string\>
**PATH:**  statefulset.metadata.initializers.result.details.name

     The name attribute of the resource associated with the status StatusReason
     (when there is a single name which can be described).

##### retryAfterSeconds	\<integer\>
**PATH:**  statefulset.metadata.initializers.result.details.retryAfterSeconds

     If specified, the time in seconds before the operation should be retried.
     Some errors may indicate the client must take an alternate action - for
     those errors this field may indicate how long to wait before taking the
     alternate action.

#### details	\<Object\>
**PATH:**  statefulset.metadata.initializers.result.details

     Extended data associated with the reason. Each reason may define its own
     extended details. This field is optional and the data returned is not
     guaranteed to conform to any schema except that defined by the reason type.

#### kind	\<string\>
**PATH:**  statefulset.metadata.initializers.result.kind

     Kind is a string value representing the REST resource this object
     represents. Servers may infer this from the endpoint the client submits
     requests to. Cannot be updated. In CamelCase. More info:
     https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds

#### message	\<string\>
**PATH:**  statefulset.metadata.initializers.result.message

     A human-readable description of the status of this operation.

#### metadata \<Object\>
**PATH:**  statefulset.metadata.initializers.result.metadata

KIND:     StatefulSet
VERSION:  apps/v1

RESOURCE: metadata <Object>

DESCRIPTION:
     Standard list metadata. More info:
     https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds

     ListMeta describes metadata that synthetic resources must have, including
     lists and various status objects. A resource may have only one of
     {ObjectMeta, ListMeta}.

##### continue	\<string\>
**PATH:**  statefulset.metadata.initializers.result.metadata.continue

     continue may be set if the user set a limit on the number of items
     returned, and indicates that the server has more data available. The value
     is opaque and may be used to issue another request to the endpoint that
     served this list to retrieve the next set of available objects. Continuing
     a consistent list may not be possible if the server configuration has
     changed or more than a few minutes have passed. The resourceVersion field
     returned when using this continue value will be identical to the value in
     the first response, unless you have received this token from an error
     message.

##### resourceVersion	\<string\>
**PATH:**  statefulset.metadata.initializers.result.metadata.resourceVersion

     String that identifies the server's internal version of this object that
     can be used by clients to determine when objects have changed. Value must
     be treated as opaque by clients and passed unmodified back to the server.
     Populated by the system. Read-only. More info:
     https://git.k8s.io/community/contributors/devel/api-conventions.md#concurrency-control-and-consistency

#### metadata	\<Object\>
**PATH:**  statefulset.metadata.initializers.result.metadata

     Standard list metadata. More info:
     https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds

#### reason	\<string\>
**PATH:**  statefulset.metadata.initializers.result.reason

     A machine-readable description of why this operation is in the "Failure"
     status. If this value is empty there is no information available. A Reason
     clarifies an HTTP status code but does not override it.

## initializers	\<Object\>
**PATH:**  statefulset.metadata.initializers

     An initializer is a controller which enforces some system invariant at
     object creation time. This field is a list of initializers that have not
     yet acted on this object. If nil or empty, this object has been completely
     initialized. Otherwise, the object is considered uninitialized and is
     hidden (in list/watch and get calls) from clients that haven't explicitly
     asked to observe uninitialized objects. When an object is created, the
     system will populate this list with the current set of initializers. Only
     privileged users may set or modify this list. Once it is empty, it may not
     be modified further by any user. DEPRECATED - initializers are an alpha
     field and will be removed in v1.15.

## labels	\<map[string]string\>
**PATH:**  statefulset.metadata.labels

     Map of string keys and values that can be used to organize and categorize
     (scope and select) objects. May match selectors of replication controllers
     and services. More info: http://kubernetes.io/docs/user-guide/labels

## managedFields \<[]Object\>
**PATH:**  statefulset.metadata.managedFields

KIND:     StatefulSet
VERSION:  apps/v1

RESOURCE: managedFields <[]Object>

DESCRIPTION:
     ManagedFields maps workflow-id and version to the set of fields that are
     managed by that workflow. This is mostly for internal housekeeping, and
     users typically shouldn't need to set or understand this field. A workflow
     can be the user's name, a controller's name, or the name of a specific
     apply path like "ci-cd". The set of fields is always in the version that
     the workflow used when modifying the object. This field is alpha and can be
     changed or removed without notice.

     ManagedFieldsEntry is a workflow-id, a FieldSet and the group version of
     the resource that the fieldset applies to.

### apiVersion	\<string\>
**PATH:**  statefulset.metadata.managedFields.apiVersion

     APIVersion defines the version of this resource that this field set applies
     to. The format is "group/version" just like the top-level APIVersion field.
     It is necessary to track the version of a field set because it cannot be
     automatically converted.

   fields	<map[string]>
     Fields identifies a set of fields.

### manager	\<string\>
**PATH:**  statefulset.metadata.managedFields.manager

     Manager is an identifier of the workflow managing these fields.

### operation	\<string\>
**PATH:**  statefulset.metadata.managedFields.operation

     Operation is the type of operation which lead to this ManagedFieldsEntry
     being created. The only valid values for this field are 'Apply' and
     'Update'.

## managedFields	\<[]Object\>
**PATH:**  statefulset.metadata.managedFields

     ManagedFields maps workflow-id and version to the set of fields that are
     managed by that workflow. This is mostly for internal housekeeping, and
     users typically shouldn't need to set or understand this field. A workflow
     can be the user's name, a controller's name, or the name of a specific
     apply path like "ci-cd". The set of fields is always in the version that
     the workflow used when modifying the object. This field is alpha and can be
     changed or removed without notice.

## name	\<string\>
**PATH:**  statefulset.metadata.name

     Name must be unique within a namespace. Is required when creating
     resources, although some resources may allow a client to request the
     generation of an appropriate name automatically. Name is primarily intended
     for creation idempotence and configuration definition. Cannot be updated.
     More info: http://kubernetes.io/docs/user-guide/identifiers#names

## namespace	\<string\>
**PATH:**  statefulset.metadata.namespace

     Namespace defines the space within each name must be unique. An empty
     namespace is equivalent to the "default" namespace, but "default" is the
     canonical representation. Not all objects are required to be scoped to a
     namespace - the value of this field for those objects will be empty. Must
     be a DNS_LABEL. Cannot be updated. More info:
     http://kubernetes.io/docs/user-guide/namespaces

## ownerReferences \<[]Object\>
**PATH:**  statefulset.metadata.ownerReferences

KIND:     StatefulSet
VERSION:  apps/v1

RESOURCE: ownerReferences <[]Object>

DESCRIPTION:
     List of objects depended by this object. If ALL objects in the list have
     been deleted, this object will be garbage collected. If this object is
     managed by a controller, then an entry in this list will point to this
     controller, with the controller field set to true. There cannot be more
     than one managing controller.

     OwnerReference contains enough information to let you identify an owning
     object. An owning object must be in the same namespace as the dependent, or
     be cluster-scoped, so there is no namespace field.

### apiVersion	\<string\> -required-
**PATH:**  statefulset.metadata.ownerReferences.apiVersion

     API version of the referent.

### blockOwnerDeletion	\<boolean\>
**PATH:**  statefulset.metadata.ownerReferences.blockOwnerDeletion

     If true, AND if the owner has the "foregroundDeletion" finalizer, then the
     owner cannot be deleted from the key-value store until this reference is
     removed. Defaults to false. To set this field, a user needs "delete"
     permission of the owner, otherwise 422 (Unprocessable Entity) will be
     returned.

### controller	\<boolean\>
**PATH:**  statefulset.metadata.ownerReferences.controller

     If true, this reference points to the managing controller.

### kind	\<string\> -required-
**PATH:**  statefulset.metadata.ownerReferences.kind

     Kind of the referent. More info:
     https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds

### name	\<string\> -required-
**PATH:**  statefulset.metadata.ownerReferences.name

     Name of the referent. More info:
     http://kubernetes.io/docs/user-guide/identifiers#names

## ownerReferences	\<[]Object\>
**PATH:**  statefulset.metadata.ownerReferences

     List of objects depended by this object. If ALL objects in the list have
     been deleted, this object will be garbage collected. If this object is
     managed by a controller, then an entry in this list will point to this
     controller, with the controller field set to true. There cannot be more
     than one managing controller.

## resourceVersion	\<string\>
**PATH:**  statefulset.metadata.resourceVersion

     An opaque value that represents the internal version of this object that
     can be used by clients to determine when objects have changed. May be used
     for optimistic concurrency, change detection, and the watch operation on a
     resource or set of resources. Clients must treat these values as opaque and
     passed unmodified back to the server. They may only be valid for a
     particular resource or set of resources. Populated by the system.
     Read-only. Value must be treated as opaque by clients and . More info:
     https://git.k8s.io/community/contributors/devel/api-conventions.md#concurrency-control-and-consistency

## selfLink	\<string\>
**PATH:**  statefulset.metadata.selfLink

     SelfLink is a URL representing this object. Populated by the system.
     Read-only.

# metadata	\<Object\>
**PATH:**  statefulset.metadata


# spec \<Object\>
**PATH:**  statefulset.spec

KIND:     StatefulSet
VERSION:  apps/v1

RESOURCE: spec <Object>

DESCRIPTION:
     Spec defines the desired identities of pods in this set.

     A StatefulSetSpec is the specification of a StatefulSet.

## podManagementPolicy	\<string\>
**PATH:**  statefulset.spec.podManagementPolicy

     podManagementPolicy controls how pods are created during initial scale up,
     when replacing pods on nodes, or when scaling down. The default policy is
     `OrderedReady`, where pods are created in increasing order (pod-0, then
     pod-1, etc) and the controller will wait until each pod is ready before
     continuing. When scaling down, the pods are removed in the opposite order.
     The alternative policy is `Parallel` which will create pods in parallel to
     match the desired scale without waiting, and on scale down will delete all
     pods at once.

## replicas	\<integer\>
**PATH:**  statefulset.spec.replicas

     replicas is the desired number of replicas of the given Template. These are
     replicas in the sense that they are instantiations of the same Template,
     but individual replicas also have a consistent identity. If unspecified,
     defaults to 1.

## revisionHistoryLimit	\<integer\>
**PATH:**  statefulset.spec.revisionHistoryLimit

     revisionHistoryLimit is the maximum number of revisions that will be
     maintained in the StatefulSet's revision history. The revision history
     consists of all revisions not represented by a currently applied
     StatefulSetSpec version. The default value is 10.

## selector \<Object\>
**PATH:**  statefulset.spec.selector

KIND:     StatefulSet
VERSION:  apps/v1

RESOURCE: selector <Object>

DESCRIPTION:
     selector is a label query over pods that should match the replica count. It
     must match the pod template's labels. More info:
     https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/#label-selectors

     A label selector is a label query over a set of resources. The result of
     matchLabels and matchExpressions are ANDed. An empty label selector matches
     all objects. A null label selector matches no objects.

### matchExpressions \<[]Object\>
**PATH:**  statefulset.spec.selector.matchExpressions

KIND:     StatefulSet
VERSION:  apps/v1

RESOURCE: matchExpressions <[]Object>

DESCRIPTION:
     matchExpressions is a list of label selector requirements. The requirements
     are ANDed.

     A label selector requirement is a selector that contains values, a key, and
     an operator that relates the key and values.

#### key	\<string\> -required-
**PATH:**  statefulset.spec.selector.matchExpressions.key

     key is the label key that the selector applies to.

#### operator	\<string\> -required-
**PATH:**  statefulset.spec.selector.matchExpressions.operator

     operator represents a key's relationship to a set of values. Valid
     operators are In, NotIn, Exists and DoesNotExist.

### matchExpressions	\<[]Object\>
**PATH:**  statefulset.spec.selector.matchExpressions

     matchExpressions is a list of label selector requirements. The requirements
     are ANDed.

## selector	\<Object\> -required-
**PATH:**  statefulset.spec.selector

     selector is a label query over pods that should match the replica count. It
     must match the pod template's labels. More info:
     https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/#label-selectors

## serviceName	\<string\> -required-
**PATH:**  statefulset.spec.serviceName

     serviceName is the name of the service that governs this StatefulSet. This
     service must exist before the StatefulSet, and is responsible for the
     network identity of the set. Pods get DNS/hostnames that follow the
     pattern: pod-specific-string.serviceName.default.svc.cluster.local where
     "pod-specific-string" is managed by the StatefulSet controller.

## template \<Object\>
**PATH:**  statefulset.spec.template

KIND:     StatefulSet
VERSION:  apps/v1

RESOURCE: template <Object>

DESCRIPTION:
     template is the object that describes the pod that will be created if
     insufficient replicas are detected. Each pod stamped out by the StatefulSet
     will fulfill this Template, but have a unique identity from the rest of the
     StatefulSet.

     PodTemplateSpec describes the data a pod should have when created from a
     template

### metadata \<Object\>
**PATH:**  statefulset.spec.template.metadata

KIND:     StatefulSet
VERSION:  apps/v1

RESOURCE: metadata <Object>

DESCRIPTION:
     Standard object's metadata. More info:
     https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata

     ObjectMeta is metadata that all persisted resources must have, which
     includes all objects users must create.

#### annotations	\<map[string]string\>
**PATH:**  statefulset.spec.template.metadata.annotations

     Annotations is an unstructured key value map stored with a resource that
     may be set by external tools to store and retrieve arbitrary metadata. They
     are not queryable and should be preserved when modifying objects. More
     info: http://kubernetes.io/docs/user-guide/annotations

#### clusterName	\<string\>
**PATH:**  statefulset.spec.template.metadata.clusterName

     The name of the cluster which the object belongs to. This is used to
     distinguish resources with same name and namespace in different clusters.
     This field is not set anywhere right now and apiserver is going to ignore
     it if set in create or update request.

#### creationTimestamp	\<string\>
**PATH:**  statefulset.spec.template.metadata.creationTimestamp

     CreationTimestamp is a timestamp representing the server time when this
     object was created. It is not guaranteed to be set in happens-before order
     across separate operations. Clients may not set this value. It is
     represented in RFC3339 form and is in UTC. Populated by the system.
     Read-only. Null for lists. More info:
     https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata

#### deletionGracePeriodSeconds	\<integer\>
**PATH:**  statefulset.spec.template.metadata.deletionGracePeriodSeconds

     Number of seconds allowed for this object to gracefully terminate before it
     will be removed from the system. Only set when deletionTimestamp is also
     set. May only be shortened. Read-only.

#### deletionTimestamp	\<string\>
**PATH:**  statefulset.spec.template.metadata.deletionTimestamp

     DeletionTimestamp is RFC 3339 date and time at which this resource will be
     deleted. This field is set by the server when a graceful deletion is
     requested by the user, and is not directly settable by a client. The
     resource is expected to be deleted (no longer visible from resource lists,
     and not reachable by name) after the time in this field, once the
     finalizers list is empty. As long as the finalizers list contains items,
     deletion is blocked. Once the deletionTimestamp is set, this value may not
     be unset or be set further into the future, although it may be shortened or
     the resource may be deleted prior to this time. For example, a user may
     request that a pod is deleted in 30 seconds. The Kubelet will react by
     sending a graceful termination signal to the containers in the pod. After
     that 30 seconds, the Kubelet will send a hard termination signal (SIGKILL)
     to the container and after cleanup, remove the pod from the API. In the
     presence of network partitions, this object may still exist after this
     timestamp, until an administrator or automated process can determine the
     resource is fully terminated. If not set, graceful deletion of the object
     has not been requested. Populated by the system when a graceful deletion is
     requested. Read-only. More info:
     https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata

#### finalizers	\<[]string\>
**PATH:**  statefulset.spec.template.metadata.finalizers

     Must be empty before the object is deleted from the registry. Each entry is
     an identifier for the responsible component that will remove the entry from
     the list. If the deletionTimestamp of the object is non-nil, entries in
     this list can only be removed.

#### generateName	\<string\>
**PATH:**  statefulset.spec.template.metadata.generateName

     GenerateName is an optional prefix, used by the server, to generate a
     unique name ONLY IF the Name field has not been provided. If this field is
     used, the name returned to the client will be different than the name
     passed. This value will also be combined with a unique suffix. The provided
     value has the same validation rules as the Name field, and may be truncated
     by the length of the suffix required to make the value unique on the
     server. If this field is specified and the generated name exists, the
     server will NOT return a 409 - instead, it will either return 201 Created
     or 500 with Reason ServerTimeout indicating a unique name could not be
     found in the time allotted, and the client should retry (optionally after
     the time indicated in the Retry-After header). Applied only if Name is not
     specified. More info:
     https://git.k8s.io/community/contributors/devel/api-conventions.md#idempotency

#### generation	\<integer\>
**PATH:**  statefulset.spec.template.metadata.generation

     A sequence number representing a specific generation of the desired state.
     Populated by the system. Read-only.

#### initializers \<Object\>
**PATH:**  statefulset.spec.template.metadata.initializers

KIND:     StatefulSet
VERSION:  apps/v1

RESOURCE: initializers <Object>

DESCRIPTION:
     An initializer is a controller which enforces some system invariant at
     object creation time. This field is a list of initializers that have not
     yet acted on this object. If nil or empty, this object has been completely
     initialized. Otherwise, the object is considered uninitialized and is
     hidden (in list/watch and get calls) from clients that haven't explicitly
     asked to observe uninitialized objects. When an object is created, the
     system will populate this list with the current set of initializers. Only
     privileged users may set or modify this list. Once it is empty, it may not
     be modified further by any user. DEPRECATED - initializers are an alpha
     field and will be removed in v1.15.

     Initializers tracks the progress of initialization.

##### pending \<[]Object\>
**PATH:**  statefulset.spec.template.metadata.initializers.pending

KIND:     StatefulSet
VERSION:  apps/v1

RESOURCE: pending <[]Object>

DESCRIPTION:
     Pending is a list of initializers that must execute in order before this
     object is visible. When the last pending initializer is removed, and no
     failing result is set, the initializers struct will be set to nil and the
     object is considered as initialized and visible to all clients.

     Initializer is information about an initializer that has not yet completed.

##### pending	\<[]Object\> -required-
**PATH:**  statefulset.spec.template.metadata.initializers.pending

     Pending is a list of initializers that must execute in order before this
     object is visible. When the last pending initializer is removed, and no
     failing result is set, the initializers struct will be set to nil and the
     object is considered as initialized and visible to all clients.

##### result \<Object\>
**PATH:**  statefulset.spec.template.metadata.initializers.result

KIND:     StatefulSet
VERSION:  apps/v1

RESOURCE: result <Object>

DESCRIPTION:
     If result is set with the Failure field, the object will be persisted to
     storage and then deleted, ensuring that other clients can observe the
     deletion.

     Status is a return value for calls that don't return other objects.

###### apiVersion	\<string\>
**PATH:**  statefulset.spec.template.metadata.initializers.result.apiVersion

     APIVersion defines the versioned schema of this representation of an
     object. Servers should convert recognized schemas to the latest internal
     value, and may reject unrecognized values. More info:
     https://git.k8s.io/community/contributors/devel/api-conventions.md#resources

###### code	\<integer\>
**PATH:**  statefulset.spec.template.metadata.initializers.result.code

     Suggested HTTP return code for this status, 0 if not set.

###### details \<Object\>
**PATH:**  statefulset.spec.template.metadata.initializers.result.details

KIND:     StatefulSet
VERSION:  apps/v1

RESOURCE: details <Object>

DESCRIPTION:
     Extended data associated with the reason. Each reason may define its own
     extended details. This field is optional and the data returned is not
     guaranteed to conform to any schema except that defined by the reason type.

     StatusDetails is a set of additional properties that MAY be set by the
     server to provide additional information about a response. The Reason field
     of a Status object defines what attributes will be set. Clients must ignore
     fields that do not match the defined type of each attribute, and should
     assume that any attribute may be empty, invalid, or under defined.

- causes \<[]Object\>
**PATH:**  statefulset.spec.template.metadata.initializers.result.details.causes

KIND:     StatefulSet
VERSION:  apps/v1

RESOURCE: causes <[]Object>

DESCRIPTION:
     The Causes array includes more details associated with the StatusReason
     failure. Not all StatusReasons may provide detailed causes.

     StatusCause provides more information about an api.Status failure,
     including cases when multiple errors are encountered.

- field	\<string\>
**PATH:**  statefulset.spec.template.metadata.initializers.result.details.causes.field

     The field of the resource that has caused this error, as named by its JSON
     serialization. May include dot and postfix notation for nested attributes.
     Arrays are zero-indexed. Fields may appear more than once in an array of
     causes due to fields having multiple errors. Optional. Examples: "name" -
     the field "name" on the current resource "items[0].name" - the field "name"
     on the first array entry in "items"

- message	\<string\>
**PATH:**  statefulset.spec.template.metadata.initializers.result.details.causes.message

     A human-readable description of the cause of the error. This field may be
     presented as-is to a reader.

- causes	\<[]Object\>
**PATH:**  statefulset.spec.template.metadata.initializers.result.details.causes

     The Causes array includes more details associated with the StatusReason
     failure. Not all StatusReasons may provide detailed causes.

- group	\<string\>
**PATH:**  statefulset.spec.template.metadata.initializers.result.details.group

     The group attribute of the resource associated with the status
     StatusReason.

- kind	\<string\>
**PATH:**  statefulset.spec.template.metadata.initializers.result.details.kind

     The kind attribute of the resource associated with the status StatusReason.
     On some operations may differ from the requested resource Kind. More info:
     https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds

- name	\<string\>
**PATH:**  statefulset.spec.template.metadata.initializers.result.details.name

     The name attribute of the resource associated with the status StatusReason
     (when there is a single name which can be described).

- retryAfterSeconds	\<integer\>
**PATH:**  statefulset.spec.template.metadata.initializers.result.details.retryAfterSeconds

     If specified, the time in seconds before the operation should be retried.
     Some errors may indicate the client must take an alternate action - for
     those errors this field may indicate how long to wait before taking the
     alternate action.

###### details	\<Object\>
**PATH:**  statefulset.spec.template.metadata.initializers.result.details

     Extended data associated with the reason. Each reason may define its own
     extended details. This field is optional and the data returned is not
     guaranteed to conform to any schema except that defined by the reason type.

###### kind	\<string\>
**PATH:**  statefulset.spec.template.metadata.initializers.result.kind

     Kind is a string value representing the REST resource this object
     represents. Servers may infer this from the endpoint the client submits
     requests to. Cannot be updated. In CamelCase. More info:
     https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds

###### message	\<string\>
**PATH:**  statefulset.spec.template.metadata.initializers.result.message

     A human-readable description of the status of this operation.

###### metadata \<Object\>
**PATH:**  statefulset.spec.template.metadata.initializers.result.metadata

KIND:     StatefulSet
VERSION:  apps/v1

RESOURCE: metadata <Object>

DESCRIPTION:
     Standard list metadata. More info:
     https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds

     ListMeta describes metadata that synthetic resources must have, including
     lists and various status objects. A resource may have only one of
     {ObjectMeta, ListMeta}.

- continue	\<string\>
**PATH:**  statefulset.spec.template.metadata.initializers.result.metadata.continue

     continue may be set if the user set a limit on the number of items
     returned, and indicates that the server has more data available. The value
     is opaque and may be used to issue another request to the endpoint that
     served this list to retrieve the next set of available objects. Continuing
     a consistent list may not be possible if the server configuration has
     changed or more than a few minutes have passed. The resourceVersion field
     returned when using this continue value will be identical to the value in
     the first response, unless you have received this token from an error
     message.

- resourceVersion	\<string\>
**PATH:**  statefulset.spec.template.metadata.initializers.result.metadata.resourceVersion

     String that identifies the server's internal version of this object that
     can be used by clients to determine when objects have changed. Value must
     be treated as opaque by clients and passed unmodified back to the server.
     Populated by the system. Read-only. More info:
     https://git.k8s.io/community/contributors/devel/api-conventions.md#concurrency-control-and-consistency

###### metadata	\<Object\>
**PATH:**  statefulset.spec.template.metadata.initializers.result.metadata

     Standard list metadata. More info:
     https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds

###### reason	\<string\>
**PATH:**  statefulset.spec.template.metadata.initializers.result.reason

     A machine-readable description of why this operation is in the "Failure"
     status. If this value is empty there is no information available. A Reason
     clarifies an HTTP status code but does not override it.

#### initializers	\<Object\>
**PATH:**  statefulset.spec.template.metadata.initializers

     An initializer is a controller which enforces some system invariant at
     object creation time. This field is a list of initializers that have not
     yet acted on this object. If nil or empty, this object has been completely
     initialized. Otherwise, the object is considered uninitialized and is
     hidden (in list/watch and get calls) from clients that haven't explicitly
     asked to observe uninitialized objects. When an object is created, the
     system will populate this list with the current set of initializers. Only
     privileged users may set or modify this list. Once it is empty, it may not
     be modified further by any user. DEPRECATED - initializers are an alpha
     field and will be removed in v1.15.

#### labels	\<map[string]string\>
**PATH:**  statefulset.spec.template.metadata.labels

     Map of string keys and values that can be used to organize and categorize
     (scope and select) objects. May match selectors of replication controllers
     and services. More info: http://kubernetes.io/docs/user-guide/labels

#### managedFields \<[]Object\>
**PATH:**  statefulset.spec.template.metadata.managedFields

KIND:     StatefulSet
VERSION:  apps/v1

RESOURCE: managedFields <[]Object>

DESCRIPTION:
     ManagedFields maps workflow-id and version to the set of fields that are
     managed by that workflow. This is mostly for internal housekeeping, and
     users typically shouldn't need to set or understand this field. A workflow
     can be the user's name, a controller's name, or the name of a specific
     apply path like "ci-cd". The set of fields is always in the version that
     the workflow used when modifying the object. This field is alpha and can be
     changed or removed without notice.

     ManagedFieldsEntry is a workflow-id, a FieldSet and the group version of
     the resource that the fieldset applies to.

##### apiVersion	\<string\>
**PATH:**  statefulset.spec.template.metadata.managedFields.apiVersion

     APIVersion defines the version of this resource that this field set applies
     to. The format is "group/version" just like the top-level APIVersion field.
     It is necessary to track the version of a field set because it cannot be
     automatically converted.

   fields	<map[string]>
     Fields identifies a set of fields.

##### manager	\<string\>
**PATH:**  statefulset.spec.template.metadata.managedFields.manager

     Manager is an identifier of the workflow managing these fields.

##### operation	\<string\>
**PATH:**  statefulset.spec.template.metadata.managedFields.operation

     Operation is the type of operation which lead to this ManagedFieldsEntry
     being created. The only valid values for this field are 'Apply' and
     'Update'.

#### managedFields	\<[]Object\>
**PATH:**  statefulset.spec.template.metadata.managedFields

     ManagedFields maps workflow-id and version to the set of fields that are
     managed by that workflow. This is mostly for internal housekeeping, and
     users typically shouldn't need to set or understand this field. A workflow
     can be the user's name, a controller's name, or the name of a specific
     apply path like "ci-cd". The set of fields is always in the version that
     the workflow used when modifying the object. This field is alpha and can be
     changed or removed without notice.

#### name	\<string\>
**PATH:**  statefulset.spec.template.metadata.name

     Name must be unique within a namespace. Is required when creating
     resources, although some resources may allow a client to request the
     generation of an appropriate name automatically. Name is primarily intended
     for creation idempotence and configuration definition. Cannot be updated.
     More info: http://kubernetes.io/docs/user-guide/identifiers#names

#### namespace	\<string\>
**PATH:**  statefulset.spec.template.metadata.namespace

     Namespace defines the space within each name must be unique. An empty
     namespace is equivalent to the "default" namespace, but "default" is the
     canonical representation. Not all objects are required to be scoped to a
     namespace - the value of this field for those objects will be empty. Must
     be a DNS_LABEL. Cannot be updated. More info:
     http://kubernetes.io/docs/user-guide/namespaces

#### ownerReferences \<[]Object\>
**PATH:**  statefulset.spec.template.metadata.ownerReferences

KIND:     StatefulSet
VERSION:  apps/v1

RESOURCE: ownerReferences <[]Object>

DESCRIPTION:
     List of objects depended by this object. If ALL objects in the list have
     been deleted, this object will be garbage collected. If this object is
     managed by a controller, then an entry in this list will point to this
     controller, with the controller field set to true. There cannot be more
     than one managing controller.

     OwnerReference contains enough information to let you identify an owning
     object. An owning object must be in the same namespace as the dependent, or
     be cluster-scoped, so there is no namespace field.

##### apiVersion	\<string\> -required-
**PATH:**  statefulset.spec.template.metadata.ownerReferences.apiVersion

     API version of the referent.

##### blockOwnerDeletion	\<boolean\>
**PATH:**  statefulset.spec.template.metadata.ownerReferences.blockOwnerDeletion

     If true, AND if the owner has the "foregroundDeletion" finalizer, then the
     owner cannot be deleted from the key-value store until this reference is
     removed. Defaults to false. To set this field, a user needs "delete"
     permission of the owner, otherwise 422 (Unprocessable Entity) will be
     returned.

##### controller	\<boolean\>
**PATH:**  statefulset.spec.template.metadata.ownerReferences.controller

     If true, this reference points to the managing controller.

##### kind	\<string\> -required-
**PATH:**  statefulset.spec.template.metadata.ownerReferences.kind

     Kind of the referent. More info:
     https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds

##### name	\<string\> -required-
**PATH:**  statefulset.spec.template.metadata.ownerReferences.name

     Name of the referent. More info:
     http://kubernetes.io/docs/user-guide/identifiers#names

#### ownerReferences	\<[]Object\>
**PATH:**  statefulset.spec.template.metadata.ownerReferences

     List of objects depended by this object. If ALL objects in the list have
     been deleted, this object will be garbage collected. If this object is
     managed by a controller, then an entry in this list will point to this
     controller, with the controller field set to true. There cannot be more
     than one managing controller.

#### resourceVersion	\<string\>
**PATH:**  statefulset.spec.template.metadata.resourceVersion

     An opaque value that represents the internal version of this object that
     can be used by clients to determine when objects have changed. May be used
     for optimistic concurrency, change detection, and the watch operation on a
     resource or set of resources. Clients must treat these values as opaque and
     passed unmodified back to the server. They may only be valid for a
     particular resource or set of resources. Populated by the system.
     Read-only. Value must be treated as opaque by clients and . More info:
     https://git.k8s.io/community/contributors/devel/api-conventions.md#concurrency-control-and-consistency

#### selfLink	\<string\>
**PATH:**  statefulset.spec.template.metadata.selfLink

     SelfLink is a URL representing this object. Populated by the system.
     Read-only.

### metadata	\<Object\>
**PATH:**  statefulset.spec.template.metadata

     Standard object's metadata. More info:
     https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata

### spec \<Object\>
**PATH:**  statefulset.spec.template.spec

KIND:     StatefulSet
VERSION:  apps/v1

RESOURCE: spec <Object>

DESCRIPTION:
     Specification of the desired behavior of the pod. More info:
     https://git.k8s.io/community/contributors/devel/api-conventions.md#spec-and-status

     PodSpec is a description of a pod.

#### activeDeadlineSeconds	\<integer\>
**PATH:**  statefulset.spec.template.spec.activeDeadlineSeconds

     Optional duration in seconds the pod may be active on the node relative to
     StartTime before the system will actively try to mark it failed and kill
     associated containers. Value must be a positive integer.

#### affinity \<Object\>
**PATH:**  statefulset.spec.template.spec.affinity

KIND:     StatefulSet
VERSION:  apps/v1

RESOURCE: affinity <Object>

DESCRIPTION:
     If specified, the pod's scheduling constraints

     Affinity is a group of affinity scheduling rules.

##### nodeAffinity \<Object\>
**PATH:**  statefulset.spec.template.spec.affinity.nodeAffinity

KIND:     StatefulSet
VERSION:  apps/v1

RESOURCE: nodeAffinity <Object>

DESCRIPTION:
     Describes node affinity scheduling rules for the pod.

     Node affinity is a group of node affinity scheduling rules.

###### preferredDuringSchedulingIgnoredDuringExecution \<[]Object\>
**PATH:**  statefulset.spec.template.spec.affinity.nodeAffinity.preferredDuringSchedulingIgnoredDuringExecution

KIND:     StatefulSet
VERSION:  apps/v1

RESOURCE: preferredDuringSchedulingIgnoredDuringExecution <[]Object>

DESCRIPTION:
     The scheduler will prefer to schedule pods to nodes that satisfy the
     affinity expressions specified by this field, but it may choose a node that
     violates one or more of the expressions. The node that is most preferred is
     the one with the greatest sum of weights, i.e. for each node that meets all
     of the scheduling requirements (resource request, requiredDuringScheduling
     affinity expressions, etc.), compute a sum by iterating through the
     elements of this field and adding "weight" to the sum if the node matches
     the corresponding matchExpressions; the node(s) with the highest sum are
     the most preferred.

     An empty preferred scheduling term matches all objects with implicit weight
     0 (i.e. it's a no-op). A null preferred scheduling term matches no objects
     (i.e. is also a no-op).

- preference \<Object\>
**PATH:**  statefulset.spec.template.spec.affinity.nodeAffinity.preferredDuringSchedulingIgnoredDuringExecution.preference

KIND:     StatefulSet
VERSION:  apps/v1

RESOURCE: preference <Object>

DESCRIPTION:
     A node selector term, associated with the corresponding weight.

     A null or empty node selector term matches no objects. The requirements of
     them are ANDed. The TopologySelectorTerm type implements a subset of the
     NodeSelectorTerm.

- matchExpressions \<[]Object\>
**PATH:**  statefulset.spec.template.spec.affinity.nodeAffinity.preferredDuringSchedulingIgnoredDuringExecution.preference.matchExpressions

KIND:     StatefulSet
VERSION:  apps/v1

RESOURCE: matchExpressions <[]Object>

DESCRIPTION:
     A list of node selector requirements by node's labels.

     A node selector requirement is a selector that contains values, a key, and
     an operator that relates the key and values.

- key	\<string\> -required-
**PATH:**  statefulset.spec.template.spec.affinity.nodeAffinity.preferredDuringSchedulingIgnoredDuringExecution.preference.matchExpressions.key

     The label key that the selector applies to.

- operator	\<string\> -required-
**PATH:**  statefulset.spec.template.spec.affinity.nodeAffinity.preferredDuringSchedulingIgnoredDuringExecution.preference.matchExpressions.operator

     Represents a key's relationship to a set of values. Valid operators are In,
     NotIn, Exists, DoesNotExist. Gt, and Lt.

- matchExpressions	\<[]Object\>
**PATH:**  statefulset.spec.template.spec.affinity.nodeAffinity.preferredDuringSchedulingIgnoredDuringExecution.preference.matchExpressions

     A list of node selector requirements by node's labels.

- matchFields \<[]Object\>
**PATH:**  statefulset.spec.template.spec.affinity.nodeAffinity.preferredDuringSchedulingIgnoredDuringExecution.preference.matchFields

KIND:     StatefulSet
VERSION:  apps/v1

RESOURCE: matchFields <[]Object>

DESCRIPTION:
     A list of node selector requirements by node's fields.

     A node selector requirement is a selector that contains values, a key, and
     an operator that relates the key and values.

- key	\<string\> -required-
**PATH:**  statefulset.spec.template.spec.affinity.nodeAffinity.preferredDuringSchedulingIgnoredDuringExecution.preference.matchFields.key

     The label key that the selector applies to.

- operator	\<string\> -required-
**PATH:**  statefulset.spec.template.spec.affinity.nodeAffinity.preferredDuringSchedulingIgnoredDuringExecution.preference.matchFields.operator

     Represents a key's relationship to a set of values. Valid operators are In,
     NotIn, Exists, DoesNotExist. Gt, and Lt.

- preference	\<Object\> -required-
**PATH:**  statefulset.spec.template.spec.affinity.nodeAffinity.preferredDuringSchedulingIgnoredDuringExecution.preference

     A node selector term, associated with the corresponding weight.

###### preferredDuringSchedulingIgnoredDuringExecution	\<[]Object\>
**PATH:**  statefulset.spec.template.spec.affinity.nodeAffinity.preferredDuringSchedulingIgnoredDuringExecution

     The scheduler will prefer to schedule pods to nodes that satisfy the
     affinity expressions specified by this field, but it may choose a node that
     violates one or more of the expressions. The node that is most preferred is
     the one with the greatest sum of weights, i.e. for each node that meets all
     of the scheduling requirements (resource request, requiredDuringScheduling
     affinity expressions, etc.), compute a sum by iterating through the
     elements of this field and adding "weight" to the sum if the node matches
     the corresponding matchExpressions; the node(s) with the highest sum are
     the most preferred.

###### requiredDuringSchedulingIgnoredDuringExecution \<Object\>
**PATH:**  statefulset.spec.template.spec.affinity.nodeAffinity.requiredDuringSchedulingIgnoredDuringExecution

KIND:     StatefulSet
VERSION:  apps/v1

RESOURCE: requiredDuringSchedulingIgnoredDuringExecution <Object>

DESCRIPTION:
     If the affinity requirements specified by this field are not met at
     scheduling time, the pod will not be scheduled onto the node. If the
     affinity requirements specified by this field cease to be met at some point
     during pod execution (e.g. due to an update), the system may or may not try
     to eventually evict the pod from its node.

     A node selector represents the union of the results of one or more label
     queries over a set of nodes; that is, it represents the OR of the selectors
     represented by the node selector terms.

- nodeSelectorTerms \<[]Object\>
**PATH:**  statefulset.spec.template.spec.affinity.nodeAffinity.requiredDuringSchedulingIgnoredDuringExecution.nodeSelectorTerms

KIND:     StatefulSet
VERSION:  apps/v1

RESOURCE: nodeSelectorTerms <[]Object>

DESCRIPTION:
     Required. A list of node selector terms. The terms are ORed.

     A null or empty node selector term matches no objects. The requirements of
     them are ANDed. The TopologySelectorTerm type implements a subset of the
     NodeSelectorTerm.

- matchExpressions \<[]Object\>
**PATH:**  statefulset.spec.template.spec.affinity.nodeAffinity.requiredDuringSchedulingIgnoredDuringExecution.nodeSelectorTerms.matchExpressions

KIND:     StatefulSet
VERSION:  apps/v1

RESOURCE: matchExpressions <[]Object>

DESCRIPTION:
     A list of node selector requirements by node's labels.

     A node selector requirement is a selector that contains values, a key, and
     an operator that relates the key and values.

- key	\<string\> -required-
**PATH:**  statefulset.spec.template.spec.affinity.nodeAffinity.requiredDuringSchedulingIgnoredDuringExecution.nodeSelectorTerms.matchExpressions.key

     The label key that the selector applies to.

- operator	\<string\> -required-
**PATH:**  statefulset.spec.template.spec.affinity.nodeAffinity.requiredDuringSchedulingIgnoredDuringExecution.nodeSelectorTerms.matchExpressions.operator

     Represents a key's relationship to a set of values. Valid operators are In,
     NotIn, Exists, DoesNotExist. Gt, and Lt.

- matchExpressions	\<[]Object\>
**PATH:**  statefulset.spec.template.spec.affinity.nodeAffinity.requiredDuringSchedulingIgnoredDuringExecution.nodeSelectorTerms.matchExpressions

     A list of node selector requirements by node's labels.

- matchFields \<[]Object\>
**PATH:**  statefulset.spec.template.spec.affinity.nodeAffinity.requiredDuringSchedulingIgnoredDuringExecution.nodeSelectorTerms.matchFields

KIND:     StatefulSet
VERSION:  apps/v1

RESOURCE: matchFields <[]Object>

DESCRIPTION:
     A list of node selector requirements by node's fields.

     A node selector requirement is a selector that contains values, a key, and
     an operator that relates the key and values.

- key	\<string\> -required-
**PATH:**  statefulset.spec.template.spec.affinity.nodeAffinity.requiredDuringSchedulingIgnoredDuringExecution.nodeSelectorTerms.matchFields.key

     The label key that the selector applies to.

- operator	\<string\> -required-
**PATH:**  statefulset.spec.template.spec.affinity.nodeAffinity.requiredDuringSchedulingIgnoredDuringExecution.nodeSelectorTerms.matchFields.operator

     Represents a key's relationship to a set of values. Valid operators are In,
     NotIn, Exists, DoesNotExist. Gt, and Lt.

##### nodeAffinity	\<Object\>
**PATH:**  statefulset.spec.template.spec.affinity.nodeAffinity

     Describes node affinity scheduling rules for the pod.

##### podAffinity \<Object\>
**PATH:**  statefulset.spec.template.spec.affinity.podAffinity

KIND:     StatefulSet
VERSION:  apps/v1

RESOURCE: podAffinity <Object>

DESCRIPTION:
     Describes pod affinity scheduling rules (e.g. co-locate this pod in the
     same node, zone, etc. as some other pod(s)).

     Pod affinity is a group of inter pod affinity scheduling rules.

###### preferredDuringSchedulingIgnoredDuringExecution \<[]Object\>
**PATH:**  statefulset.spec.template.spec.affinity.podAffinity.preferredDuringSchedulingIgnoredDuringExecution

KIND:     StatefulSet
VERSION:  apps/v1

RESOURCE: preferredDuringSchedulingIgnoredDuringExecution <[]Object>

DESCRIPTION:
     The scheduler will prefer to schedule pods to nodes that satisfy the
     affinity expressions specified by this field, but it may choose a node that
     violates one or more of the expressions. The node that is most preferred is
     the one with the greatest sum of weights, i.e. for each node that meets all
     of the scheduling requirements (resource request, requiredDuringScheduling
     affinity expressions, etc.), compute a sum by iterating through the
     elements of this field and adding "weight" to the sum if the node has pods
     which matches the corresponding podAffinityTerm; the node(s) with the
     highest sum are the most preferred.

     The weights of all of the matched WeightedPodAffinityTerm fields are added
     per-node to find the most preferred node(s)

- podAffinityTerm \<Object\>
**PATH:**  statefulset.spec.template.spec.affinity.podAffinity.preferredDuringSchedulingIgnoredDuringExecution.podAffinityTerm

KIND:     StatefulSet
VERSION:  apps/v1

RESOURCE: podAffinityTerm <Object>

DESCRIPTION:
     Required. A pod affinity term, associated with the corresponding weight.

     Defines a set of pods (namely those matching the labelSelector relative to
     the given namespace(s)) that this pod should be co-located (affinity) or
     not co-located (anti-affinity) with, where co-located is defined as running
     on a node whose value of the label with key <topologyKey> matches that of
     any node on which a pod of the set of pods is running

- labelSelector \<Object\>
**PATH:**  statefulset.spec.template.spec.affinity.podAffinity.preferredDuringSchedulingIgnoredDuringExecution.podAffinityTerm.labelSelector

KIND:     StatefulSet
VERSION:  apps/v1

RESOURCE: labelSelector <Object>

DESCRIPTION:
     A label query over a set of resources, in this case pods.

     A label selector is a label query over a set of resources. The result of
     matchLabels and matchExpressions are ANDed. An empty label selector matches
     all objects. A null label selector matches no objects.

- matchExpressions \<[]Object\>
**PATH:**  statefulset.spec.template.spec.affinity.podAffinity.preferredDuringSchedulingIgnoredDuringExecution.podAffinityTerm.labelSelector.matchExpressions

KIND:     StatefulSet
VERSION:  apps/v1

RESOURCE: matchExpressions <[]Object>

DESCRIPTION:
     matchExpressions is a list of label selector requirements. The requirements
     are ANDed.

     A label selector requirement is a selector that contains values, a key, and
     an operator that relates the key and values.

- key	\<string\> -required-
**PATH:**  statefulset.spec.template.spec.affinity.podAffinity.preferredDuringSchedulingIgnoredDuringExecution.podAffinityTerm.labelSelector.matchExpressions.key

     key is the label key that the selector applies to.

- operator	\<string\> -required-
**PATH:**  statefulset.spec.template.spec.affinity.podAffinity.preferredDuringSchedulingIgnoredDuringExecution.podAffinityTerm.labelSelector.matchExpressions.operator

     operator represents a key's relationship to a set of values. Valid
     operators are In, NotIn, Exists and DoesNotExist.

- matchExpressions	\<[]Object\>
**PATH:**  statefulset.spec.template.spec.affinity.podAffinity.preferredDuringSchedulingIgnoredDuringExecution.podAffinityTerm.labelSelector.matchExpressions

     matchExpressions is a list of label selector requirements. The requirements
     are ANDed.

- labelSelector	\<Object\>
**PATH:**  statefulset.spec.template.spec.affinity.podAffinity.preferredDuringSchedulingIgnoredDuringExecution.podAffinityTerm.labelSelector

     A label query over a set of resources, in this case pods.

- namespaces	\<[]string\>
**PATH:**  statefulset.spec.template.spec.affinity.podAffinity.preferredDuringSchedulingIgnoredDuringExecution.podAffinityTerm.namespaces

     namespaces specifies which namespaces the labelSelector applies to (matches
     against); null or empty list means "this pod's namespace"

- podAffinityTerm	\<Object\> -required-
**PATH:**  statefulset.spec.template.spec.affinity.podAffinity.preferredDuringSchedulingIgnoredDuringExecution.podAffinityTerm

     Required. A pod affinity term, associated with the corresponding weight.

###### preferredDuringSchedulingIgnoredDuringExecution	\<[]Object\>
**PATH:**  statefulset.spec.template.spec.affinity.podAffinity.preferredDuringSchedulingIgnoredDuringExecution

     The scheduler will prefer to schedule pods to nodes that satisfy the
     affinity expressions specified by this field, but it may choose a node that
     violates one or more of the expressions. The node that is most preferred is
     the one with the greatest sum of weights, i.e. for each node that meets all
     of the scheduling requirements (resource request, requiredDuringScheduling
     affinity expressions, etc.), compute a sum by iterating through the
     elements of this field and adding "weight" to the sum if the node has pods
     which matches the corresponding podAffinityTerm; the node(s) with the
     highest sum are the most preferred.

###### requiredDuringSchedulingIgnoredDuringExecution \<[]Object\>
**PATH:**  statefulset.spec.template.spec.affinity.podAffinity.requiredDuringSchedulingIgnoredDuringExecution

KIND:     StatefulSet
VERSION:  apps/v1

RESOURCE: requiredDuringSchedulingIgnoredDuringExecution <[]Object>

DESCRIPTION:
     If the affinity requirements specified by this field are not met at
     scheduling time, the pod will not be scheduled onto the node. If the
     affinity requirements specified by this field cease to be met at some point
     during pod execution (e.g. due to a pod label update), the system may or
     may not try to eventually evict the pod from its node. When there are
     multiple elements, the lists of nodes corresponding to each podAffinityTerm
     are intersected, i.e. all terms must be satisfied.

     Defines a set of pods (namely those matching the labelSelector relative to
     the given namespace(s)) that this pod should be co-located (affinity) or
     not co-located (anti-affinity) with, where co-located is defined as running
     on a node whose value of the label with key <topologyKey> matches that of
     any node on which a pod of the set of pods is running

- labelSelector \<Object\>
**PATH:**  statefulset.spec.template.spec.affinity.podAffinity.requiredDuringSchedulingIgnoredDuringExecution.labelSelector

KIND:     StatefulSet
VERSION:  apps/v1

RESOURCE: labelSelector <Object>

DESCRIPTION:
     A label query over a set of resources, in this case pods.

     A label selector is a label query over a set of resources. The result of
     matchLabels and matchExpressions are ANDed. An empty label selector matches
     all objects. A null label selector matches no objects.

- matchExpressions \<[]Object\>
**PATH:**  statefulset.spec.template.spec.affinity.podAffinity.requiredDuringSchedulingIgnoredDuringExecution.labelSelector.matchExpressions

KIND:     StatefulSet
VERSION:  apps/v1

RESOURCE: matchExpressions <[]Object>

DESCRIPTION:
     matchExpressions is a list of label selector requirements. The requirements
     are ANDed.

     A label selector requirement is a selector that contains values, a key, and
     an operator that relates the key and values.

- key	\<string\> -required-
**PATH:**  statefulset.spec.template.spec.affinity.podAffinity.requiredDuringSchedulingIgnoredDuringExecution.labelSelector.matchExpressions.key

     key is the label key that the selector applies to.

- operator	\<string\> -required-
**PATH:**  statefulset.spec.template.spec.affinity.podAffinity.requiredDuringSchedulingIgnoredDuringExecution.labelSelector.matchExpressions.operator

     operator represents a key's relationship to a set of values. Valid
     operators are In, NotIn, Exists and DoesNotExist.

- matchExpressions	\<[]Object\>
**PATH:**  statefulset.spec.template.spec.affinity.podAffinity.requiredDuringSchedulingIgnoredDuringExecution.labelSelector.matchExpressions

     matchExpressions is a list of label selector requirements. The requirements
     are ANDed.

- labelSelector	\<Object\>
**PATH:**  statefulset.spec.template.spec.affinity.podAffinity.requiredDuringSchedulingIgnoredDuringExecution.labelSelector

     A label query over a set of resources, in this case pods.

- namespaces	\<[]string\>
**PATH:**  statefulset.spec.template.spec.affinity.podAffinity.requiredDuringSchedulingIgnoredDuringExecution.namespaces

     namespaces specifies which namespaces the labelSelector applies to (matches
     against); null or empty list means "this pod's namespace"

##### podAffinity	\<Object\>
**PATH:**  statefulset.spec.template.spec.affinity.podAffinity

     Describes pod affinity scheduling rules (e.g. co-locate this pod in the
     same node, zone, etc. as some other pod(s)).

##### podAntiAffinity \<Object\>
**PATH:**  statefulset.spec.template.spec.affinity.podAntiAffinity

KIND:     StatefulSet
VERSION:  apps/v1

RESOURCE: podAntiAffinity <Object>

DESCRIPTION:
     Describes pod anti-affinity scheduling rules (e.g. avoid putting this pod
     in the same node, zone, etc. as some other pod(s)).

     Pod anti affinity is a group of inter pod anti affinity scheduling rules.

###### preferredDuringSchedulingIgnoredDuringExecution \<[]Object\>
**PATH:**  statefulset.spec.template.spec.affinity.podAntiAffinity.preferredDuringSchedulingIgnoredDuringExecution

KIND:     StatefulSet
VERSION:  apps/v1

RESOURCE: preferredDuringSchedulingIgnoredDuringExecution <[]Object>

DESCRIPTION:
     The scheduler will prefer to schedule pods to nodes that satisfy the
     anti-affinity expressions specified by this field, but it may choose a node
     that violates one or more of the expressions. The node that is most
     preferred is the one with the greatest sum of weights, i.e. for each node
     that meets all of the scheduling requirements (resource request,
     requiredDuringScheduling anti-affinity expressions, etc.), compute a sum by
     iterating through the elements of this field and adding "weight" to the sum
     if the node has pods which matches the corresponding podAffinityTerm; the
     node(s) with the highest sum are the most preferred.

     The weights of all of the matched WeightedPodAffinityTerm fields are added
     per-node to find the most preferred node(s)

- podAffinityTerm \<Object\>
**PATH:**  statefulset.spec.template.spec.affinity.podAntiAffinity.preferredDuringSchedulingIgnoredDuringExecution.podAffinityTerm

KIND:     StatefulSet
VERSION:  apps/v1

RESOURCE: podAffinityTerm <Object>

DESCRIPTION:
     Required. A pod affinity term, associated with the corresponding weight.

     Defines a set of pods (namely those matching the labelSelector relative to
     the given namespace(s)) that this pod should be co-located (affinity) or
     not co-located (anti-affinity) with, where co-located is defined as running
     on a node whose value of the label with key <topologyKey> matches that of
     any node on which a pod of the set of pods is running

- labelSelector \<Object\>
**PATH:**  statefulset.spec.template.spec.affinity.podAntiAffinity.preferredDuringSchedulingIgnoredDuringExecution.podAffinityTerm.labelSelector

KIND:     StatefulSet
VERSION:  apps/v1

RESOURCE: labelSelector <Object>

DESCRIPTION:
     A label query over a set of resources, in this case pods.

     A label selector is a label query over a set of resources. The result of
     matchLabels and matchExpressions are ANDed. An empty label selector matches
     all objects. A null label selector matches no objects.

- matchExpressions \<[]Object\>
**PATH:**  statefulset.spec.template.spec.affinity.podAntiAffinity.preferredDuringSchedulingIgnoredDuringExecution.podAffinityTerm.labelSelector.matchExpressions

KIND:     StatefulSet
VERSION:  apps/v1

RESOURCE: matchExpressions <[]Object>

DESCRIPTION:
     matchExpressions is a list of label selector requirements. The requirements
     are ANDed.

     A label selector requirement is a selector that contains values, a key, and
     an operator that relates the key and values.

- key	\<string\> -required-
**PATH:**  statefulset.spec.template.spec.affinity.podAntiAffinity.preferredDuringSchedulingIgnoredDuringExecution.podAffinityTerm.labelSelector.matchExpressions.key

     key is the label key that the selector applies to.

- operator	\<string\> -required-
**PATH:**  statefulset.spec.template.spec.affinity.podAntiAffinity.preferredDuringSchedulingIgnoredDuringExecution.podAffinityTerm.labelSelector.matchExpressions.operator

     operator represents a key's relationship to a set of values. Valid
     operators are In, NotIn, Exists and DoesNotExist.

- matchExpressions	\<[]Object\>
**PATH:**  statefulset.spec.template.spec.affinity.podAntiAffinity.preferredDuringSchedulingIgnoredDuringExecution.podAffinityTerm.labelSelector.matchExpressions

     matchExpressions is a list of label selector requirements. The requirements
     are ANDed.

- labelSelector	\<Object\>
**PATH:**  statefulset.spec.template.spec.affinity.podAntiAffinity.preferredDuringSchedulingIgnoredDuringExecution.podAffinityTerm.labelSelector

     A label query over a set of resources, in this case pods.

- namespaces	\<[]string\>
**PATH:**  statefulset.spec.template.spec.affinity.podAntiAffinity.preferredDuringSchedulingIgnoredDuringExecution.podAffinityTerm.namespaces

     namespaces specifies which namespaces the labelSelector applies to (matches
     against); null or empty list means "this pod's namespace"

- podAffinityTerm	\<Object\> -required-
**PATH:**  statefulset.spec.template.spec.affinity.podAntiAffinity.preferredDuringSchedulingIgnoredDuringExecution.podAffinityTerm

     Required. A pod affinity term, associated with the corresponding weight.

###### preferredDuringSchedulingIgnoredDuringExecution	\<[]Object\>
**PATH:**  statefulset.spec.template.spec.affinity.podAntiAffinity.preferredDuringSchedulingIgnoredDuringExecution

     The scheduler will prefer to schedule pods to nodes that satisfy the
     anti-affinity expressions specified by this field, but it may choose a node
     that violates one or more of the expressions. The node that is most
     preferred is the one with the greatest sum of weights, i.e. for each node
     that meets all of the scheduling requirements (resource request,
     requiredDuringScheduling anti-affinity expressions, etc.), compute a sum by
     iterating through the elements of this field and adding "weight" to the sum
     if the node has pods which matches the corresponding podAffinityTerm; the
     node(s) with the highest sum are the most preferred.

###### requiredDuringSchedulingIgnoredDuringExecution \<[]Object\>
**PATH:**  statefulset.spec.template.spec.affinity.podAntiAffinity.requiredDuringSchedulingIgnoredDuringExecution

KIND:     StatefulSet
VERSION:  apps/v1

RESOURCE: requiredDuringSchedulingIgnoredDuringExecution <[]Object>

DESCRIPTION:
     If the anti-affinity requirements specified by this field are not met at
     scheduling time, the pod will not be scheduled onto the node. If the
     anti-affinity requirements specified by this field cease to be met at some
     point during pod execution (e.g. due to a pod label update), the system may
     or may not try to eventually evict the pod from its node. When there are
     multiple elements, the lists of nodes corresponding to each podAffinityTerm
     are intersected, i.e. all terms must be satisfied.

     Defines a set of pods (namely those matching the labelSelector relative to
     the given namespace(s)) that this pod should be co-located (affinity) or
     not co-located (anti-affinity) with, where co-located is defined as running
     on a node whose value of the label with key <topologyKey> matches that of
     any node on which a pod of the set of pods is running

- labelSelector \<Object\>
**PATH:**  statefulset.spec.template.spec.affinity.podAntiAffinity.requiredDuringSchedulingIgnoredDuringExecution.labelSelector

KIND:     StatefulSet
VERSION:  apps/v1

RESOURCE: labelSelector <Object>

DESCRIPTION:
     A label query over a set of resources, in this case pods.

     A label selector is a label query over a set of resources. The result of
     matchLabels and matchExpressions are ANDed. An empty label selector matches
     all objects. A null label selector matches no objects.

- matchExpressions \<[]Object\>
**PATH:**  statefulset.spec.template.spec.affinity.podAntiAffinity.requiredDuringSchedulingIgnoredDuringExecution.labelSelector.matchExpressions

KIND:     StatefulSet
VERSION:  apps/v1

RESOURCE: matchExpressions <[]Object>

DESCRIPTION:
     matchExpressions is a list of label selector requirements. The requirements
     are ANDed.

     A label selector requirement is a selector that contains values, a key, and
     an operator that relates the key and values.

- key	\<string\> -required-
**PATH:**  statefulset.spec.template.spec.affinity.podAntiAffinity.requiredDuringSchedulingIgnoredDuringExecution.labelSelector.matchExpressions.key

     key is the label key that the selector applies to.

- operator	\<string\> -required-
**PATH:**  statefulset.spec.template.spec.affinity.podAntiAffinity.requiredDuringSchedulingIgnoredDuringExecution.labelSelector.matchExpressions.operator

     operator represents a key's relationship to a set of values. Valid
     operators are In, NotIn, Exists and DoesNotExist.

- matchExpressions	\<[]Object\>
**PATH:**  statefulset.spec.template.spec.affinity.podAntiAffinity.requiredDuringSchedulingIgnoredDuringExecution.labelSelector.matchExpressions

     matchExpressions is a list of label selector requirements. The requirements
     are ANDed.

- labelSelector	\<Object\>
**PATH:**  statefulset.spec.template.spec.affinity.podAntiAffinity.requiredDuringSchedulingIgnoredDuringExecution.labelSelector

     A label query over a set of resources, in this case pods.

- namespaces	\<[]string\>
**PATH:**  statefulset.spec.template.spec.affinity.podAntiAffinity.requiredDuringSchedulingIgnoredDuringExecution.namespaces

     namespaces specifies which namespaces the labelSelector applies to (matches
     against); null or empty list means "this pod's namespace"

#### affinity	\<Object\>
**PATH:**  statefulset.spec.template.spec.affinity

     If specified, the pod's scheduling constraints

#### automountServiceAccountToken	\<boolean\>
**PATH:**  statefulset.spec.template.spec.automountServiceAccountToken

     AutomountServiceAccountToken indicates whether a service account token
     should be automatically mounted.

#### containers \<[]Object\>
**PATH:**  statefulset.spec.template.spec.containers

KIND:     StatefulSet
VERSION:  apps/v1

RESOURCE: containers <[]Object>

DESCRIPTION:
     List of containers belonging to the pod. Containers cannot currently be
     added or removed. There must be at least one container in a Pod. Cannot be
     updated.

     A single application container that you want to run within a pod.

##### args	\<[]string\>
**PATH:**  statefulset.spec.template.spec.containers.args

     Arguments to the entrypoint. The docker image's CMD is used if this is not
     provided. Variable references $(VAR_NAME) are expanded using the
     container's environment. If a variable cannot be resolved, the reference in
     the input string will be unchanged. The $(VAR_NAME) syntax can be escaped
     with a double $$, ie: $$(VAR_NAME). Escaped references will never be
     expanded, regardless of whether the variable exists or not. Cannot be
     updated. More info:
     https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/#running-a-command-in-a-shell

##### command	\<[]string\>
**PATH:**  statefulset.spec.template.spec.containers.command

     Entrypoint array. Not executed within a shell. The docker image's
     ENTRYPOINT is used if this is not provided. Variable references $(VAR_NAME)
     are expanded using the container's environment. If a variable cannot be
     resolved, the reference in the input string will be unchanged. The
     $(VAR_NAME) syntax can be escaped with a double $$, ie: $$(VAR_NAME).
     Escaped references will never be expanded, regardless of whether the
     variable exists or not. Cannot be updated. More info:
     https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/#running-a-command-in-a-shell

##### env \<[]Object\>
**PATH:**  statefulset.spec.template.spec.containers.env

KIND:     StatefulSet
VERSION:  apps/v1

RESOURCE: env <[]Object>

DESCRIPTION:
     List of environment variables to set in the container. Cannot be updated.

     EnvVar represents an environment variable present in a Container.

###### name	\<string\> -required-
**PATH:**  statefulset.spec.template.spec.containers.env.name

     Name of the environment variable. Must be a C_IDENTIFIER.

###### value	\<string\>
**PATH:**  statefulset.spec.template.spec.containers.env.value

     Variable references $(VAR_NAME) are expanded using the previous defined
     environment variables in the container and any service environment
     variables. If a variable cannot be resolved, the reference in the input
     string will be unchanged. The $(VAR_NAME) syntax can be escaped with a
     double $$, ie: $$(VAR_NAME). Escaped references will never be expanded,
     regardless of whether the variable exists or not. Defaults to "".

###### valueFrom \<Object\>
**PATH:**  statefulset.spec.template.spec.containers.env.valueFrom

KIND:     StatefulSet
VERSION:  apps/v1

RESOURCE: valueFrom <Object>

DESCRIPTION:
     Source for the environment variable's value. Cannot be used if value is not
     empty.

     EnvVarSource represents a source for the value of an EnvVar.

- configMapKeyRef \<Object\>
**PATH:**  statefulset.spec.template.spec.containers.env.valueFrom.configMapKeyRef

KIND:     StatefulSet
VERSION:  apps/v1

RESOURCE: configMapKeyRef <Object>

DESCRIPTION:
     Selects a key of a ConfigMap.

     Selects a key from a ConfigMap.

- key	\<string\> -required-
**PATH:**  statefulset.spec.template.spec.containers.env.valueFrom.configMapKeyRef.key

     The key to select.

- name	\<string\>
**PATH:**  statefulset.spec.template.spec.containers.env.valueFrom.configMapKeyRef.name

     Name of the referent. More info:
     https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names

- configMapKeyRef	\<Object\>
**PATH:**  statefulset.spec.template.spec.containers.env.valueFrom.configMapKeyRef

     Selects a key of a ConfigMap.

- fieldRef \<Object\>
**PATH:**  statefulset.spec.template.spec.containers.env.valueFrom.fieldRef

KIND:     StatefulSet
VERSION:  apps/v1

RESOURCE: fieldRef <Object>

DESCRIPTION:
     Selects a field of the pod: supports metadata.name, metadata.namespace,
     metadata.labels, metadata.annotations, spec.nodeName,
     spec.serviceAccountName, status.hostIP, status.podIP.

     ObjectFieldSelector selects an APIVersioned field of an object.

- apiVersion	\<string\>
**PATH:**  statefulset.spec.template.spec.containers.env.valueFrom.fieldRef.apiVersion

     Version of the schema the FieldPath is written in terms of, defaults to
     "v1".

- fieldRef	\<Object\>
**PATH:**  statefulset.spec.template.spec.containers.env.valueFrom.fieldRef

     Selects a field of the pod: supports metadata.name, metadata.namespace,
     metadata.labels, metadata.annotations, spec.nodeName,
     spec.serviceAccountName, status.hostIP, status.podIP.

- resourceFieldRef \<Object\>
**PATH:**  statefulset.spec.template.spec.containers.env.valueFrom.resourceFieldRef

KIND:     StatefulSet
VERSION:  apps/v1

RESOURCE: resourceFieldRef <Object>

DESCRIPTION:
     Selects a resource of the container: only resources limits and requests
     (limits.cpu, limits.memory, limits.ephemeral-storage, requests.cpu,
     requests.memory and requests.ephemeral-storage) are currently supported.

     ResourceFieldSelector represents container resources (cpu, memory) and
     their output format

- containerName	\<string\>
**PATH:**  statefulset.spec.template.spec.containers.env.valueFrom.resourceFieldRef.containerName

     Container name: required for volumes, optional for env vars

- divisor	\<string\>
**PATH:**  statefulset.spec.template.spec.containers.env.valueFrom.resourceFieldRef.divisor

     Specifies the output format of the exposed resources, defaults to "1"

- resourceFieldRef	\<Object\>
**PATH:**  statefulset.spec.template.spec.containers.env.valueFrom.resourceFieldRef

     Selects a resource of the container: only resources limits and requests
     (limits.cpu, limits.memory, limits.ephemeral-storage, requests.cpu,
     requests.memory and requests.ephemeral-storage) are currently supported.

- secretKeyRef \<Object\>
**PATH:**  statefulset.spec.template.spec.containers.env.valueFrom.secretKeyRef

KIND:     StatefulSet
VERSION:  apps/v1

RESOURCE: secretKeyRef <Object>

DESCRIPTION:
     Selects a key of a secret in the pod's namespace

     SecretKeySelector selects a key of a Secret.

- key	\<string\> -required-
**PATH:**  statefulset.spec.template.spec.containers.env.valueFrom.secretKeyRef.key

     The key of the secret to select from. Must be a valid secret key.

- name	\<string\>
**PATH:**  statefulset.spec.template.spec.containers.env.valueFrom.secretKeyRef.name

     Name of the referent. More info:
     https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names

##### env	\<[]Object\>
**PATH:**  statefulset.spec.template.spec.containers.env

     List of environment variables to set in the container. Cannot be updated.

##### envFrom \<[]Object\>
**PATH:**  statefulset.spec.template.spec.containers.envFrom

KIND:     StatefulSet
VERSION:  apps/v1

RESOURCE: envFrom <[]Object>

DESCRIPTION:
     List of sources to populate environment variables in the container. The
     keys defined within a source must be a C_IDENTIFIER. All invalid keys will
     be reported as an event when the container is starting. When a key exists
     in multiple sources, the value associated with the last source will take
     precedence. Values defined by an Env with a duplicate key will take
     precedence. Cannot be updated.

     EnvFromSource represents the source of a set of ConfigMaps

###### configMapRef \<Object\>
**PATH:**  statefulset.spec.template.spec.containers.envFrom.configMapRef

KIND:     StatefulSet
VERSION:  apps/v1

RESOURCE: configMapRef <Object>

DESCRIPTION:
     The ConfigMap to select from

     ConfigMapEnvSource selects a ConfigMap to populate the environment
     variables with. The contents of the target ConfigMap's Data field will
     represent the key-value pairs as environment variables.

- name	\<string\>
**PATH:**  statefulset.spec.template.spec.containers.envFrom.configMapRef.name

     Name of the referent. More info:
     https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names

###### configMapRef	\<Object\>
**PATH:**  statefulset.spec.template.spec.containers.envFrom.configMapRef

     The ConfigMap to select from

###### prefix	\<string\>
**PATH:**  statefulset.spec.template.spec.containers.envFrom.prefix

     An optional identifier to prepend to each key in the ConfigMap. Must be a
     C_IDENTIFIER.

###### secretRef \<Object\>
**PATH:**  statefulset.spec.template.spec.containers.envFrom.secretRef

KIND:     StatefulSet
VERSION:  apps/v1

RESOURCE: secretRef <Object>

DESCRIPTION:
     The Secret to select from

     SecretEnvSource selects a Secret to populate the environment variables
     with. The contents of the target Secret's Data field will represent the
     key-value pairs as environment variables.

- name	\<string\>
**PATH:**  statefulset.spec.template.spec.containers.envFrom.secretRef.name

     Name of the referent. More info:
     https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names

##### envFrom	\<[]Object\>
**PATH:**  statefulset.spec.template.spec.containers.envFrom

     List of sources to populate environment variables in the container. The
     keys defined within a source must be a C_IDENTIFIER. All invalid keys will
     be reported as an event when the container is starting. When a key exists
     in multiple sources, the value associated with the last source will take
     precedence. Values defined by an Env with a duplicate key will take
     precedence. Cannot be updated.

##### image	\<string\>
**PATH:**  statefulset.spec.template.spec.containers.image

     Docker image name. More info:
     https://kubernetes.io/docs/concepts/containers/images This field is
     optional to allow higher level config management to default or override
     container images in workload controllers like Deployments and StatefulSets.

##### imagePullPolicy	\<string\>
**PATH:**  statefulset.spec.template.spec.containers.imagePullPolicy

     Image pull policy. One of Always, Never, IfNotPresent. Defaults to Always
     if :latest tag is specified, or IfNotPresent otherwise. Cannot be updated.
     More info:
     https://kubernetes.io/docs/concepts/containers/images#updating-images

##### lifecycle \<Object\>
**PATH:**  statefulset.spec.template.spec.containers.lifecycle

KIND:     StatefulSet
VERSION:  apps/v1

RESOURCE: lifecycle <Object>

DESCRIPTION:
     Actions that the management system should take in response to container
     lifecycle events. Cannot be updated.

     Lifecycle describes actions that the management system should take in
     response to container lifecycle events. For the PostStart and PreStop
     lifecycle handlers, management of the container blocks until the action is
     complete, unless the container process fails, in which case the handler is
     aborted.

###### postStart \<Object\>
**PATH:**  statefulset.spec.template.spec.containers.lifecycle.postStart

KIND:     StatefulSet
VERSION:  apps/v1

RESOURCE: postStart <Object>

DESCRIPTION:
     PostStart is called immediately after a container is created. If the
     handler fails, the container is terminated and restarted according to its
     restart policy. Other management of the container blocks until the hook
     completes. More info:
     https://kubernetes.io/docs/concepts/containers/container-lifecycle-hooks/#container-hooks

     Handler defines a specific action that should be taken

- exec \<Object\>
**PATH:**  statefulset.spec.template.spec.containers.lifecycle.postStart.exec

KIND:     StatefulSet
VERSION:  apps/v1

RESOURCE: exec <Object>

DESCRIPTION:
     One and only one of the following should be specified. Exec specifies the
     action to take.

     ExecAction describes a "run in container" action.

- exec	\<Object\>
**PATH:**  statefulset.spec.template.spec.containers.lifecycle.postStart.exec

     One and only one of the following should be specified. Exec specifies the
     action to take.

- httpGet \<Object\>
**PATH:**  statefulset.spec.template.spec.containers.lifecycle.postStart.httpGet

KIND:     StatefulSet
VERSION:  apps/v1

RESOURCE: httpGet <Object>

DESCRIPTION:
     HTTPGet specifies the http request to perform.

     HTTPGetAction describes an action based on HTTP Get requests.

- host	\<string\>
**PATH:**  statefulset.spec.template.spec.containers.lifecycle.postStart.httpGet.host

     Host name to connect to, defaults to the pod IP. You probably want to set
     "Host" in httpHeaders instead.

- httpHeaders \<[]Object\>
**PATH:**  statefulset.spec.template.spec.containers.lifecycle.postStart.httpGet.httpHeaders

KIND:     StatefulSet
VERSION:  apps/v1

RESOURCE: httpHeaders <[]Object>

DESCRIPTION:
     Custom headers to set in the request. HTTP allows repeated headers.

     HTTPHeader describes a custom header to be used in HTTP probes

- name	\<string\> -required-
**PATH:**  statefulset.spec.template.spec.containers.lifecycle.postStart.httpGet.httpHeaders.name

     The header field name

- httpHeaders	\<[]Object\>
**PATH:**  statefulset.spec.template.spec.containers.lifecycle.postStart.httpGet.httpHeaders

     Custom headers to set in the request. HTTP allows repeated headers.

- path	\<string\>
**PATH:**  statefulset.spec.template.spec.containers.lifecycle.postStart.httpGet.path

     Path to access on the HTTP server.

- port	\<string\> -required-
**PATH:**  statefulset.spec.template.spec.containers.lifecycle.postStart.httpGet.port

     Name or number of the port to access on the container. Number must be in
     the range 1 to 65535. Name must be an IANA_SVC_NAME.

- httpGet	\<Object\>
**PATH:**  statefulset.spec.template.spec.containers.lifecycle.postStart.httpGet

     HTTPGet specifies the http request to perform.

- tcpSocket \<Object\>
**PATH:**  statefulset.spec.template.spec.containers.lifecycle.postStart.tcpSocket

KIND:     StatefulSet
VERSION:  apps/v1

RESOURCE: tcpSocket <Object>

DESCRIPTION:
     TCPSocket specifies an action involving a TCP port. TCP hooks not yet
     supported

     TCPSocketAction describes an action based on opening a socket

- host	\<string\>
**PATH:**  statefulset.spec.template.spec.containers.lifecycle.postStart.tcpSocket.host

     Optional: Host name to connect to, defaults to the pod IP.

###### postStart	\<Object\>
**PATH:**  statefulset.spec.template.spec.containers.lifecycle.postStart

     PostStart is called immediately after a container is created. If the
     handler fails, the container is terminated and restarted according to its
     restart policy. Other management of the container blocks until the hook
     completes. More info:
     https://kubernetes.io/docs/concepts/containers/container-lifecycle-hooks/#container-hooks

###### preStop \<Object\>
**PATH:**  statefulset.spec.template.spec.containers.lifecycle.preStop

KIND:     StatefulSet
VERSION:  apps/v1

RESOURCE: preStop <Object>

DESCRIPTION:
     PreStop is called immediately before a container is terminated due to an
     API request or management event such as liveness probe failure, preemption,
     resource contention, etc. The handler is not called if the container
     crashes or exits. The reason for termination is passed to the handler. The
     Pod's termination grace period countdown begins before the PreStop hooked
     is executed. Regardless of the outcome of the handler, the container will
     eventually terminate within the Pod's termination grace period. Other
     management of the container blocks until the hook completes or until the
     termination grace period is reached. More info:
     https://kubernetes.io/docs/concepts/containers/container-lifecycle-hooks/#container-hooks

     Handler defines a specific action that should be taken

- exec \<Object\>
**PATH:**  statefulset.spec.template.spec.containers.lifecycle.preStop.exec

KIND:     StatefulSet
VERSION:  apps/v1

RESOURCE: exec <Object>

DESCRIPTION:
     One and only one of the following should be specified. Exec specifies the
     action to take.

     ExecAction describes a "run in container" action.

- exec	\<Object\>
**PATH:**  statefulset.spec.template.spec.containers.lifecycle.preStop.exec

     One and only one of the following should be specified. Exec specifies the
     action to take.

- httpGet \<Object\>
**PATH:**  statefulset.spec.template.spec.containers.lifecycle.preStop.httpGet

KIND:     StatefulSet
VERSION:  apps/v1

RESOURCE: httpGet <Object>

DESCRIPTION:
     HTTPGet specifies the http request to perform.

     HTTPGetAction describes an action based on HTTP Get requests.

- host	\<string\>
**PATH:**  statefulset.spec.template.spec.containers.lifecycle.preStop.httpGet.host

     Host name to connect to, defaults to the pod IP. You probably want to set
     "Host" in httpHeaders instead.

- httpHeaders \<[]Object\>
**PATH:**  statefulset.spec.template.spec.containers.lifecycle.preStop.httpGet.httpHeaders

KIND:     StatefulSet
VERSION:  apps/v1

RESOURCE: httpHeaders <[]Object>

DESCRIPTION:
     Custom headers to set in the request. HTTP allows repeated headers.

     HTTPHeader describes a custom header to be used in HTTP probes

- name	\<string\> -required-
**PATH:**  statefulset.spec.template.spec.containers.lifecycle.preStop.httpGet.httpHeaders.name

     The header field name

- httpHeaders	\<[]Object\>
**PATH:**  statefulset.spec.template.spec.containers.lifecycle.preStop.httpGet.httpHeaders

     Custom headers to set in the request. HTTP allows repeated headers.

- path	\<string\>
**PATH:**  statefulset.spec.template.spec.containers.lifecycle.preStop.httpGet.path

     Path to access on the HTTP server.

- port	\<string\> -required-
**PATH:**  statefulset.spec.template.spec.containers.lifecycle.preStop.httpGet.port

     Name or number of the port to access on the container. Number must be in
     the range 1 to 65535. Name must be an IANA_SVC_NAME.

- httpGet	\<Object\>
**PATH:**  statefulset.spec.template.spec.containers.lifecycle.preStop.httpGet

     HTTPGet specifies the http request to perform.

- tcpSocket \<Object\>
**PATH:**  statefulset.spec.template.spec.containers.lifecycle.preStop.tcpSocket

KIND:     StatefulSet
VERSION:  apps/v1

RESOURCE: tcpSocket <Object>

DESCRIPTION:
     TCPSocket specifies an action involving a TCP port. TCP hooks not yet
     supported

     TCPSocketAction describes an action based on opening a socket

- host	\<string\>
**PATH:**  statefulset.spec.template.spec.containers.lifecycle.preStop.tcpSocket.host

     Optional: Host name to connect to, defaults to the pod IP.

##### lifecycle	\<Object\>
**PATH:**  statefulset.spec.template.spec.containers.lifecycle

     Actions that the management system should take in response to container
     lifecycle events. Cannot be updated.

##### livenessProbe \<Object\>
**PATH:**  statefulset.spec.template.spec.containers.livenessProbe

KIND:     StatefulSet
VERSION:  apps/v1

RESOURCE: livenessProbe <Object>

DESCRIPTION:
     Periodic probe of container liveness. Container will be restarted if the
     probe fails. Cannot be updated. More info:
     https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes

     Probe describes a health check to be performed against a container to
     determine whether it is alive or ready to receive traffic.

###### exec \<Object\>
**PATH:**  statefulset.spec.template.spec.containers.livenessProbe.exec

KIND:     StatefulSet
VERSION:  apps/v1

RESOURCE: exec <Object>

DESCRIPTION:
     One and only one of the following should be specified. Exec specifies the
     action to take.

     ExecAction describes a "run in container" action.

###### exec	\<Object\>
**PATH:**  statefulset.spec.template.spec.containers.livenessProbe.exec

     One and only one of the following should be specified. Exec specifies the
     action to take.

###### failureThreshold	\<integer\>
**PATH:**  statefulset.spec.template.spec.containers.livenessProbe.failureThreshold

     Minimum consecutive failures for the probe to be considered failed after
     having succeeded. Defaults to 3. Minimum value is 1.

###### httpGet \<Object\>
**PATH:**  statefulset.spec.template.spec.containers.livenessProbe.httpGet

KIND:     StatefulSet
VERSION:  apps/v1

RESOURCE: httpGet <Object>

DESCRIPTION:
     HTTPGet specifies the http request to perform.

     HTTPGetAction describes an action based on HTTP Get requests.

- host	\<string\>
**PATH:**  statefulset.spec.template.spec.containers.livenessProbe.httpGet.host

     Host name to connect to, defaults to the pod IP. You probably want to set
     "Host" in httpHeaders instead.

- httpHeaders \<[]Object\>
**PATH:**  statefulset.spec.template.spec.containers.livenessProbe.httpGet.httpHeaders

KIND:     StatefulSet
VERSION:  apps/v1

RESOURCE: httpHeaders <[]Object>

DESCRIPTION:
     Custom headers to set in the request. HTTP allows repeated headers.

     HTTPHeader describes a custom header to be used in HTTP probes

- name	\<string\> -required-
**PATH:**  statefulset.spec.template.spec.containers.livenessProbe.httpGet.httpHeaders.name

     The header field name

- httpHeaders	\<[]Object\>
**PATH:**  statefulset.spec.template.spec.containers.livenessProbe.httpGet.httpHeaders

     Custom headers to set in the request. HTTP allows repeated headers.

- path	\<string\>
**PATH:**  statefulset.spec.template.spec.containers.livenessProbe.httpGet.path

     Path to access on the HTTP server.

- port	\<string\> -required-
**PATH:**  statefulset.spec.template.spec.containers.livenessProbe.httpGet.port

     Name or number of the port to access on the container. Number must be in
     the range 1 to 65535. Name must be an IANA_SVC_NAME.

###### httpGet	\<Object\>
**PATH:**  statefulset.spec.template.spec.containers.livenessProbe.httpGet

     HTTPGet specifies the http request to perform.

###### initialDelaySeconds	\<integer\>
**PATH:**  statefulset.spec.template.spec.containers.livenessProbe.initialDelaySeconds

     Number of seconds after the container has started before liveness probes
     are initiated. More info:
     https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes

###### periodSeconds	\<integer\>
**PATH:**  statefulset.spec.template.spec.containers.livenessProbe.periodSeconds

     How often (in seconds) to perform the probe. Default to 10 seconds. Minimum
     value is 1.

###### successThreshold	\<integer\>
**PATH:**  statefulset.spec.template.spec.containers.livenessProbe.successThreshold

     Minimum consecutive successes for the probe to be considered successful
     after having failed. Defaults to 1. Must be 1 for liveness. Minimum value
     is 1.

###### tcpSocket \<Object\>
**PATH:**  statefulset.spec.template.spec.containers.livenessProbe.tcpSocket

KIND:     StatefulSet
VERSION:  apps/v1

RESOURCE: tcpSocket <Object>

DESCRIPTION:
     TCPSocket specifies an action involving a TCP port. TCP hooks not yet
     supported

     TCPSocketAction describes an action based on opening a socket

- host	\<string\>
**PATH:**  statefulset.spec.template.spec.containers.livenessProbe.tcpSocket.host

     Optional: Host name to connect to, defaults to the pod IP.

###### tcpSocket	\<Object\>
**PATH:**  statefulset.spec.template.spec.containers.livenessProbe.tcpSocket

     TCPSocket specifies an action involving a TCP port. TCP hooks not yet
     supported

##### livenessProbe	\<Object\>
**PATH:**  statefulset.spec.template.spec.containers.livenessProbe

     Periodic probe of container liveness. Container will be restarted if the
     probe fails. Cannot be updated. More info:
     https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes

##### name	\<string\> -required-
**PATH:**  statefulset.spec.template.spec.containers.name

     Name of the container specified as a DNS_LABEL. Each container in a pod
     must have a unique name (DNS_LABEL). Cannot be updated.

##### ports \<[]Object\>
**PATH:**  statefulset.spec.template.spec.containers.ports

KIND:     StatefulSet
VERSION:  apps/v1

RESOURCE: ports <[]Object>

DESCRIPTION:
     List of ports to expose from the container. Exposing a port here gives the
     system additional information about the network connections a container
     uses, but is primarily informational. Not specifying a port here DOES NOT
     prevent that port from being exposed. Any port which is listening on the
     default "0.0.0.0" address inside a container will be accessible from the
     network. Cannot be updated.

     ContainerPort represents a network port in a single container.

###### containerPort	\<integer\> -required-
**PATH:**  statefulset.spec.template.spec.containers.ports.containerPort

     Number of port to expose on the pod's IP address. This must be a valid port
     number, 0 < x < 65536.

###### hostIP	\<string\>
**PATH:**  statefulset.spec.template.spec.containers.ports.hostIP

     What host IP to bind the external port to.

###### hostPort	\<integer\>
**PATH:**  statefulset.spec.template.spec.containers.ports.hostPort

     Number of port to expose on the host. If specified, this must be a valid
     port number, 0 < x < 65536. If HostNetwork is specified, this must match
     ContainerPort. Most containers do not need this.

###### name	\<string\>
**PATH:**  statefulset.spec.template.spec.containers.ports.name

     If specified, this must be an IANA_SVC_NAME and unique within the pod. Each
     named port in a pod must have a unique name. Name for the port that can be
     referred to by services.

##### ports	\<[]Object\>
**PATH:**  statefulset.spec.template.spec.containers.ports

     List of ports to expose from the container. Exposing a port here gives the
     system additional information about the network connections a container
     uses, but is primarily informational. Not specifying a port here DOES NOT
     prevent that port from being exposed. Any port which is listening on the
     default "0.0.0.0" address inside a container will be accessible from the
     network. Cannot be updated.

##### readinessProbe \<Object\>
**PATH:**  statefulset.spec.template.spec.containers.readinessProbe

KIND:     StatefulSet
VERSION:  apps/v1

RESOURCE: readinessProbe <Object>

DESCRIPTION:
     Periodic probe of container service readiness. Container will be removed
     from service endpoints if the probe fails. Cannot be updated. More info:
     https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes

     Probe describes a health check to be performed against a container to
     determine whether it is alive or ready to receive traffic.

###### exec \<Object\>
**PATH:**  statefulset.spec.template.spec.containers.readinessProbe.exec

KIND:     StatefulSet
VERSION:  apps/v1

RESOURCE: exec <Object>

DESCRIPTION:
     One and only one of the following should be specified. Exec specifies the
     action to take.

     ExecAction describes a "run in container" action.

###### exec	\<Object\>
**PATH:**  statefulset.spec.template.spec.containers.readinessProbe.exec

     One and only one of the following should be specified. Exec specifies the
     action to take.

###### failureThreshold	\<integer\>
**PATH:**  statefulset.spec.template.spec.containers.readinessProbe.failureThreshold

     Minimum consecutive failures for the probe to be considered failed after
     having succeeded. Defaults to 3. Minimum value is 1.

###### httpGet \<Object\>
**PATH:**  statefulset.spec.template.spec.containers.readinessProbe.httpGet

KIND:     StatefulSet
VERSION:  apps/v1

RESOURCE: httpGet <Object>

DESCRIPTION:
     HTTPGet specifies the http request to perform.

     HTTPGetAction describes an action based on HTTP Get requests.

- host	\<string\>
**PATH:**  statefulset.spec.template.spec.containers.readinessProbe.httpGet.host

     Host name to connect to, defaults to the pod IP. You probably want to set
     "Host" in httpHeaders instead.

- httpHeaders \<[]Object\>
**PATH:**  statefulset.spec.template.spec.containers.readinessProbe.httpGet.httpHeaders

KIND:     StatefulSet
VERSION:  apps/v1

RESOURCE: httpHeaders <[]Object>

DESCRIPTION:
     Custom headers to set in the request. HTTP allows repeated headers.

     HTTPHeader describes a custom header to be used in HTTP probes

- name	\<string\> -required-
**PATH:**  statefulset.spec.template.spec.containers.readinessProbe.httpGet.httpHeaders.name

     The header field name

- httpHeaders	\<[]Object\>
**PATH:**  statefulset.spec.template.spec.containers.readinessProbe.httpGet.httpHeaders

     Custom headers to set in the request. HTTP allows repeated headers.

- path	\<string\>
**PATH:**  statefulset.spec.template.spec.containers.readinessProbe.httpGet.path

     Path to access on the HTTP server.

- port	\<string\> -required-
**PATH:**  statefulset.spec.template.spec.containers.readinessProbe.httpGet.port

     Name or number of the port to access on the container. Number must be in
     the range 1 to 65535. Name must be an IANA_SVC_NAME.

###### httpGet	\<Object\>
**PATH:**  statefulset.spec.template.spec.containers.readinessProbe.httpGet

     HTTPGet specifies the http request to perform.

###### initialDelaySeconds	\<integer\>
**PATH:**  statefulset.spec.template.spec.containers.readinessProbe.initialDelaySeconds

     Number of seconds after the container has started before liveness probes
     are initiated. More info:
     https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes

###### periodSeconds	\<integer\>
**PATH:**  statefulset.spec.template.spec.containers.readinessProbe.periodSeconds

     How often (in seconds) to perform the probe. Default to 10 seconds. Minimum
     value is 1.

###### successThreshold	\<integer\>
**PATH:**  statefulset.spec.template.spec.containers.readinessProbe.successThreshold

     Minimum consecutive successes for the probe to be considered successful
     after having failed. Defaults to 1. Must be 1 for liveness. Minimum value
     is 1.

###### tcpSocket \<Object\>
**PATH:**  statefulset.spec.template.spec.containers.readinessProbe.tcpSocket

KIND:     StatefulSet
VERSION:  apps/v1

RESOURCE: tcpSocket <Object>

DESCRIPTION:
     TCPSocket specifies an action involving a TCP port. TCP hooks not yet
     supported

     TCPSocketAction describes an action based on opening a socket

- host	\<string\>
**PATH:**  statefulset.spec.template.spec.containers.readinessProbe.tcpSocket.host

     Optional: Host name to connect to, defaults to the pod IP.

###### tcpSocket	\<Object\>
**PATH:**  statefulset.spec.template.spec.containers.readinessProbe.tcpSocket

     TCPSocket specifies an action involving a TCP port. TCP hooks not yet
     supported

##### readinessProbe	\<Object\>
**PATH:**  statefulset.spec.template.spec.containers.readinessProbe

     Periodic probe of container service readiness. Container will be removed
     from service endpoints if the probe fails. Cannot be updated. More info:
     https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes

##### resources \<Object\>
**PATH:**  statefulset.spec.template.spec.containers.resources

KIND:     StatefulSet
VERSION:  apps/v1

RESOURCE: resources <Object>

DESCRIPTION:
     Compute Resources required by this container. Cannot be updated. More info:
     https://kubernetes.io/docs/concepts/configuration/manage-compute-resources-container/

     ResourceRequirements describes the compute resource requirements.

###### limits	\<map[string]string\>
**PATH:**  statefulset.spec.template.spec.containers.resources.limits

     Limits describes the maximum amount of compute resources allowed. More
     info:
     https://kubernetes.io/docs/concepts/configuration/manage-compute-resources-container/

##### resources	\<Object\>
**PATH:**  statefulset.spec.template.spec.containers.resources

     Compute Resources required by this container. Cannot be updated. More info:
     https://kubernetes.io/docs/concepts/configuration/manage-compute-resources-container/

##### securityContext \<Object\>
**PATH:**  statefulset.spec.template.spec.containers.securityContext

KIND:     StatefulSet
VERSION:  apps/v1

RESOURCE: securityContext <Object>

DESCRIPTION:
     Security options the pod should run with. More info:
     https://kubernetes.io/docs/concepts/policy/security-context/ More info:
     https://kubernetes.io/docs/tasks/configure-pod-container/security-context/

     SecurityContext holds security configuration that will be applied to a
     container. Some fields are present in both SecurityContext and
     PodSecurityContext. When both are set, the values in SecurityContext take
     precedence.

###### allowPrivilegeEscalation	\<boolean\>
**PATH:**  statefulset.spec.template.spec.containers.securityContext.allowPrivilegeEscalation

     AllowPrivilegeEscalation controls whether a process can gain more
     privileges than its parent process. This bool directly controls if the
     no_new_privs flag will be set on the container process.
     AllowPrivilegeEscalation is true always when the container is: 1) run as
     Privileged 2) has CAP_SYS_ADMIN

###### capabilities \<Object\>
**PATH:**  statefulset.spec.template.spec.containers.securityContext.capabilities

KIND:     StatefulSet
VERSION:  apps/v1

RESOURCE: capabilities <Object>

DESCRIPTION:
     The capabilities to add/drop when running containers. Defaults to the
     default set of capabilities granted by the container runtime.

     Adds and removes POSIX capabilities from running containers.

- add	\<[]string\>
**PATH:**  statefulset.spec.template.spec.containers.securityContext.capabilities.add

     Added capabilities

###### capabilities	\<Object\>
**PATH:**  statefulset.spec.template.spec.containers.securityContext.capabilities

     The capabilities to add/drop when running containers. Defaults to the
     default set of capabilities granted by the container runtime.

###### privileged	\<boolean\>
**PATH:**  statefulset.spec.template.spec.containers.securityContext.privileged

     Run container in privileged mode. Processes in privileged containers are
     essentially equivalent to root on the host. Defaults to false.

###### procMount	\<string\>
**PATH:**  statefulset.spec.template.spec.containers.securityContext.procMount

     procMount denotes the type of proc mount to use for the containers. The
     default is DefaultProcMount which uses the container runtime defaults for
     readonly paths and masked paths. This requires the ProcMountType feature
     flag to be enabled.

###### readOnlyRootFilesystem	\<boolean\>
**PATH:**  statefulset.spec.template.spec.containers.securityContext.readOnlyRootFilesystem

     Whether this container has a read-only root filesystem. Default is false.

###### runAsGroup	\<integer\>
**PATH:**  statefulset.spec.template.spec.containers.securityContext.runAsGroup

     The GID to run the entrypoint of the container process. Uses runtime
     default if unset. May also be set in PodSecurityContext. If set in both
     SecurityContext and PodSecurityContext, the value specified in
     SecurityContext takes precedence.

###### runAsNonRoot	\<boolean\>
**PATH:**  statefulset.spec.template.spec.containers.securityContext.runAsNonRoot

     Indicates that the container must run as a non-root user. If true, the
     Kubelet will validate the image at runtime to ensure that it does not run
     as UID 0 (root) and fail to start the container if it does. If unset or
     false, no such validation will be performed. May also be set in
     PodSecurityContext. If set in both SecurityContext and PodSecurityContext,
     the value specified in SecurityContext takes precedence.

###### runAsUser	\<integer\>
**PATH:**  statefulset.spec.template.spec.containers.securityContext.runAsUser

     The UID to run the entrypoint of the container process. Defaults to user
     specified in image metadata if unspecified. May also be set in
     PodSecurityContext. If set in both SecurityContext and PodSecurityContext,
     the value specified in SecurityContext takes precedence.

###### seLinuxOptions \<Object\>
**PATH:**  statefulset.spec.template.spec.containers.securityContext.seLinuxOptions

KIND:     StatefulSet
VERSION:  apps/v1

RESOURCE: seLinuxOptions <Object>

DESCRIPTION:
     The SELinux context to be applied to the container. If unspecified, the
     container runtime will allocate a random SELinux context for each
     container. May also be set in PodSecurityContext. If set in both
     SecurityContext and PodSecurityContext, the value specified in
     SecurityContext takes precedence.

     SELinuxOptions are the labels to be applied to the container

- level	\<string\>
**PATH:**  statefulset.spec.template.spec.containers.securityContext.seLinuxOptions.level

     Level is SELinux level label that applies to the container.

- role	\<string\>
**PATH:**  statefulset.spec.template.spec.containers.securityContext.seLinuxOptions.role

     Role is a SELinux role label that applies to the container.

- type	\<string\>
**PATH:**  statefulset.spec.template.spec.containers.securityContext.seLinuxOptions.type

     Type is a SELinux type label that applies to the container.

##### securityContext	\<Object\>
**PATH:**  statefulset.spec.template.spec.containers.securityContext

     Security options the pod should run with. More info:
     https://kubernetes.io/docs/concepts/policy/security-context/ More info:
     https://kubernetes.io/docs/tasks/configure-pod-container/security-context/

##### stdin	\<boolean\>
**PATH:**  statefulset.spec.template.spec.containers.stdin

     Whether this container should allocate a buffer for stdin in the container
     runtime. If this is not set, reads from stdin in the container will always
     result in EOF. Default is false.

##### stdinOnce	\<boolean\>
**PATH:**  statefulset.spec.template.spec.containers.stdinOnce

     Whether the container runtime should close the stdin channel after it has
     been opened by a single attach. When stdin is true the stdin stream will
     remain open across multiple attach sessions. If stdinOnce is set to true,
     stdin is opened on container start, is empty until the first client
     attaches to stdin, and then remains open and accepts data until the client
     disconnects, at which time stdin is closed and remains closed until the
     container is restarted. If this flag is false, a container processes that
     reads from stdin will never receive an EOF. Default is false

##### terminationMessagePath	\<string\>
**PATH:**  statefulset.spec.template.spec.containers.terminationMessagePath

     Optional: Path at which the file to which the container's termination
     message will be written is mounted into the container's filesystem. Message
     written is intended to be brief final status, such as an assertion failure
     message. Will be truncated by the node if greater than 4096 bytes. The
     total message length across all containers will be limited to 12kb.
     Defaults to /dev/termination-log. Cannot be updated.

##### terminationMessagePolicy	\<string\>
**PATH:**  statefulset.spec.template.spec.containers.terminationMessagePolicy

     Indicate how the termination message should be populated. File will use the
     contents of terminationMessagePath to populate the container status message
     on both success and failure. FallbackToLogsOnError will use the last chunk
     of container log output if the termination message file is empty and the
     container exited with an error. The log output is limited to 2048 bytes or
     80 lines, whichever is smaller. Defaults to File. Cannot be updated.

##### tty	\<boolean\>
**PATH:**  statefulset.spec.template.spec.containers.tty

     Whether this container should allocate a TTY for itself, also requires
     'stdin' to be true. Default is false.

##### volumeDevices \<[]Object\>
**PATH:**  statefulset.spec.template.spec.containers.volumeDevices

KIND:     StatefulSet
VERSION:  apps/v1

RESOURCE: volumeDevices <[]Object>

DESCRIPTION:
     volumeDevices is the list of block devices to be used by the container.
     This is a beta feature.

     volumeDevice describes a mapping of a raw block device within a container.

###### devicePath	\<string\> -required-
**PATH:**  statefulset.spec.template.spec.containers.volumeDevices.devicePath

     devicePath is the path inside of the container that the device will be
     mapped to.

##### volumeDevices	\<[]Object\>
**PATH:**  statefulset.spec.template.spec.containers.volumeDevices

     volumeDevices is the list of block devices to be used by the container.
     This is a beta feature.

##### volumeMounts \<[]Object\>
**PATH:**  statefulset.spec.template.spec.containers.volumeMounts

KIND:     StatefulSet
VERSION:  apps/v1

RESOURCE: volumeMounts <[]Object>

DESCRIPTION:
     Pod volumes to mount into the container's filesystem. Cannot be updated.

     VolumeMount describes a mounting of a Volume within a container.

###### mountPath	\<string\> -required-
**PATH:**  statefulset.spec.template.spec.containers.volumeMounts.mountPath

     Path within the container at which the volume should be mounted. Must not
     contain ':'.

###### mountPropagation	\<string\>
**PATH:**  statefulset.spec.template.spec.containers.volumeMounts.mountPropagation

     mountPropagation determines how mounts are propagated from the host to
     container and the other way around. When not set, MountPropagationNone is
     used. This field is beta in 1.10.

###### name	\<string\> -required-
**PATH:**  statefulset.spec.template.spec.containers.volumeMounts.name

     This must match the Name of a Volume.

###### readOnly	\<boolean\>
**PATH:**  statefulset.spec.template.spec.containers.volumeMounts.readOnly

     Mounted read-only if true, read-write otherwise (false or unspecified).
     Defaults to false.

###### subPath	\<string\>
**PATH:**  statefulset.spec.template.spec.containers.volumeMounts.subPath

     Path within the volume from which the container's volume should be mounted.
     Defaults to "" (volume's root).

##### volumeMounts	\<[]Object\>
**PATH:**  statefulset.spec.template.spec.containers.volumeMounts

     Pod volumes to mount into the container's filesystem. Cannot be updated.

#### containers	\<[]Object\> -required-
**PATH:**  statefulset.spec.template.spec.containers

     List of containers belonging to the pod. Containers cannot currently be
     added or removed. There must be at least one container in a Pod. Cannot be
     updated.

#### dnsConfig \<Object\>
**PATH:**  statefulset.spec.template.spec.dnsConfig

KIND:     StatefulSet
VERSION:  apps/v1

RESOURCE: dnsConfig <Object>

DESCRIPTION:
     Specifies the DNS parameters of a pod. Parameters specified here will be
     merged to the generated DNS configuration based on DNSPolicy.

     PodDNSConfig defines the DNS parameters of a pod in addition to those
     generated from DNSPolicy.

##### nameservers	\<[]string\>
**PATH:**  statefulset.spec.template.spec.dnsConfig.nameservers

     A list of DNS name server IP addresses. This will be appended to the base
     nameservers generated from DNSPolicy. Duplicated nameservers will be
     removed.

##### options \<[]Object\>
**PATH:**  statefulset.spec.template.spec.dnsConfig.options

KIND:     StatefulSet
VERSION:  apps/v1

RESOURCE: options <[]Object>

DESCRIPTION:
     A list of DNS resolver options. This will be merged with the base options
     generated from DNSPolicy. Duplicated entries will be removed. Resolution
     options given in Options will override those that appear in the base
     DNSPolicy.

     PodDNSConfigOption defines DNS resolver options of a pod.

###### name	\<string\>
**PATH:**  statefulset.spec.template.spec.dnsConfig.options.name

     Required.

##### options	\<[]Object\>
**PATH:**  statefulset.spec.template.spec.dnsConfig.options

     A list of DNS resolver options. This will be merged with the base options
     generated from DNSPolicy. Duplicated entries will be removed. Resolution
     options given in Options will override those that appear in the base
     DNSPolicy.

#### dnsConfig	\<Object\>
**PATH:**  statefulset.spec.template.spec.dnsConfig

     Specifies the DNS parameters of a pod. Parameters specified here will be
     merged to the generated DNS configuration based on DNSPolicy.

#### dnsPolicy	\<string\>
**PATH:**  statefulset.spec.template.spec.dnsPolicy

     Set DNS policy for the pod. Defaults to "ClusterFirst". Valid values are
     'ClusterFirstWithHostNet', 'ClusterFirst', 'Default' or 'None'. DNS
     parameters given in DNSConfig will be merged with the policy selected with
     DNSPolicy. To have DNS options set along with hostNetwork, you have to
     specify DNS policy explicitly to 'ClusterFirstWithHostNet'.

#### enableServiceLinks	\<boolean\>
**PATH:**  statefulset.spec.template.spec.enableServiceLinks

     EnableServiceLinks indicates whether information about services should be
     injected into pod's environment variables, matching the syntax of Docker
     links. Optional: Defaults to true.

#### hostAliases \<[]Object\>
**PATH:**  statefulset.spec.template.spec.hostAliases

KIND:     StatefulSet
VERSION:  apps/v1

RESOURCE: hostAliases <[]Object>

DESCRIPTION:
     HostAliases is an optional list of hosts and IPs that will be injected into
     the pod's hosts file if specified. This is only valid for non-hostNetwork
     pods.

     HostAlias holds the mapping between IP and hostnames that will be injected
     as an entry in the pod's hosts file.

##### hostnames	\<[]string\>
**PATH:**  statefulset.spec.template.spec.hostAliases.hostnames

     Hostnames for the above IP address.

#### hostAliases	\<[]Object\>
**PATH:**  statefulset.spec.template.spec.hostAliases

     HostAliases is an optional list of hosts and IPs that will be injected into
     the pod's hosts file if specified. This is only valid for non-hostNetwork
     pods.

#### hostIPC	\<boolean\>
**PATH:**  statefulset.spec.template.spec.hostIPC

     Use the host's ipc namespace. Optional: Default to false.

#### hostNetwork	\<boolean\>
**PATH:**  statefulset.spec.template.spec.hostNetwork

     Host networking requested for this pod. Use the host's network namespace.
     If this option is set, the ports that will be used must be specified.
     Default to false.

#### hostPID	\<boolean\>
**PATH:**  statefulset.spec.template.spec.hostPID

     Use the host's pid namespace. Optional: Default to false.

#### hostname	\<string\>
**PATH:**  statefulset.spec.template.spec.hostname

     Specifies the hostname of the Pod If not specified, the pod's hostname will
     be set to a system-defined value.

#### imagePullSecrets \<[]Object\>
**PATH:**  statefulset.spec.template.spec.imagePullSecrets

KIND:     StatefulSet
VERSION:  apps/v1

RESOURCE: imagePullSecrets <[]Object>

DESCRIPTION:
     ImagePullSecrets is an optional list of references to secrets in the same
     namespace to use for pulling any of the images used by this PodSpec. If
     specified, these secrets will be passed to individual puller
     implementations for them to use. For example, in the case of docker, only
     DockerConfig type secrets are honored. More info:
     https://kubernetes.io/docs/concepts/containers/images#specifying-imagepullsecrets-on-a-pod

     LocalObjectReference contains enough information to let you locate the
     referenced object inside the same namespace.

#### imagePullSecrets	\<[]Object\>
**PATH:**  statefulset.spec.template.spec.imagePullSecrets

     ImagePullSecrets is an optional list of references to secrets in the same
     namespace to use for pulling any of the images used by this PodSpec. If
     specified, these secrets will be passed to individual puller
     implementations for them to use. For example, in the case of docker, only
     DockerConfig type secrets are honored. More info:
     https://kubernetes.io/docs/concepts/containers/images#specifying-imagepullsecrets-on-a-pod

#### initContainers \<[]Object\>
**PATH:**  statefulset.spec.template.spec.initContainers

KIND:     StatefulSet
VERSION:  apps/v1

RESOURCE: initContainers <[]Object>

DESCRIPTION:
     List of initialization containers belonging to the pod. Init containers are
     executed in order prior to containers being started. If any init container
     fails, the pod is considered to have failed and is handled according to its
     restartPolicy. The name for an init container or normal container must be
     unique among all containers. Init containers may not have Lifecycle
     actions, Readiness probes, or Liveness probes. The resourceRequirements of
     an init container are taken into account during scheduling by finding the
     highest request/limit for each resource type, and then using the max of of
     that value or the sum of the normal containers. Limits are applied to init
     containers in a similar fashion. Init containers cannot currently be added
     or removed. Cannot be updated. More info:
     https://kubernetes.io/docs/concepts/workloads/pods/init-containers/

     A single application container that you want to run within a pod.

##### args	\<[]string\>
**PATH:**  statefulset.spec.template.spec.initContainers.args

     Arguments to the entrypoint. The docker image's CMD is used if this is not
     provided. Variable references $(VAR_NAME) are expanded using the
     container's environment. If a variable cannot be resolved, the reference in
     the input string will be unchanged. The $(VAR_NAME) syntax can be escaped
     with a double $$, ie: $$(VAR_NAME). Escaped references will never be
     expanded, regardless of whether the variable exists or not. Cannot be
     updated. More info:
     https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/#running-a-command-in-a-shell

##### command	\<[]string\>
**PATH:**  statefulset.spec.template.spec.initContainers.command

     Entrypoint array. Not executed within a shell. The docker image's
     ENTRYPOINT is used if this is not provided. Variable references $(VAR_NAME)
     are expanded using the container's environment. If a variable cannot be
     resolved, the reference in the input string will be unchanged. The
     $(VAR_NAME) syntax can be escaped with a double $$, ie: $$(VAR_NAME).
     Escaped references will never be expanded, regardless of whether the
     variable exists or not. Cannot be updated. More info:
     https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/#running-a-command-in-a-shell

##### env \<[]Object\>
**PATH:**  statefulset.spec.template.spec.initContainers.env

KIND:     StatefulSet
VERSION:  apps/v1

RESOURCE: env <[]Object>

DESCRIPTION:
     List of environment variables to set in the container. Cannot be updated.

     EnvVar represents an environment variable present in a Container.

###### name	\<string\> -required-
**PATH:**  statefulset.spec.template.spec.initContainers.env.name

     Name of the environment variable. Must be a C_IDENTIFIER.

###### value	\<string\>
**PATH:**  statefulset.spec.template.spec.initContainers.env.value

     Variable references $(VAR_NAME) are expanded using the previous defined
     environment variables in the container and any service environment
     variables. If a variable cannot be resolved, the reference in the input
     string will be unchanged. The $(VAR_NAME) syntax can be escaped with a
     double $$, ie: $$(VAR_NAME). Escaped references will never be expanded,
     regardless of whether the variable exists or not. Defaults to "".

###### valueFrom \<Object\>
**PATH:**  statefulset.spec.template.spec.initContainers.env.valueFrom

KIND:     StatefulSet
VERSION:  apps/v1

RESOURCE: valueFrom <Object>

DESCRIPTION:
     Source for the environment variable's value. Cannot be used if value is not
     empty.

     EnvVarSource represents a source for the value of an EnvVar.

- configMapKeyRef \<Object\>
**PATH:**  statefulset.spec.template.spec.initContainers.env.valueFrom.configMapKeyRef

KIND:     StatefulSet
VERSION:  apps/v1

RESOURCE: configMapKeyRef <Object>

DESCRIPTION:
     Selects a key of a ConfigMap.

     Selects a key from a ConfigMap.

- key	\<string\> -required-
**PATH:**  statefulset.spec.template.spec.initContainers.env.valueFrom.configMapKeyRef.key

     The key to select.

- name	\<string\>
**PATH:**  statefulset.spec.template.spec.initContainers.env.valueFrom.configMapKeyRef.name

     Name of the referent. More info:
     https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names

- configMapKeyRef	\<Object\>
**PATH:**  statefulset.spec.template.spec.initContainers.env.valueFrom.configMapKeyRef

     Selects a key of a ConfigMap.

- fieldRef \<Object\>
**PATH:**  statefulset.spec.template.spec.initContainers.env.valueFrom.fieldRef

KIND:     StatefulSet
VERSION:  apps/v1

RESOURCE: fieldRef <Object>

DESCRIPTION:
     Selects a field of the pod: supports metadata.name, metadata.namespace,
     metadata.labels, metadata.annotations, spec.nodeName,
     spec.serviceAccountName, status.hostIP, status.podIP.

     ObjectFieldSelector selects an APIVersioned field of an object.

- apiVersion	\<string\>
**PATH:**  statefulset.spec.template.spec.initContainers.env.valueFrom.fieldRef.apiVersion

     Version of the schema the FieldPath is written in terms of, defaults to
     "v1".

- fieldRef	\<Object\>
**PATH:**  statefulset.spec.template.spec.initContainers.env.valueFrom.fieldRef

     Selects a field of the pod: supports metadata.name, metadata.namespace,
     metadata.labels, metadata.annotations, spec.nodeName,
     spec.serviceAccountName, status.hostIP, status.podIP.

- resourceFieldRef \<Object\>
**PATH:**  statefulset.spec.template.spec.initContainers.env.valueFrom.resourceFieldRef

KIND:     StatefulSet
VERSION:  apps/v1

RESOURCE: resourceFieldRef <Object>

DESCRIPTION:
     Selects a resource of the container: only resources limits and requests
     (limits.cpu, limits.memory, limits.ephemeral-storage, requests.cpu,
     requests.memory and requests.ephemeral-storage) are currently supported.

     ResourceFieldSelector represents container resources (cpu, memory) and
     their output format

- containerName	\<string\>
**PATH:**  statefulset.spec.template.spec.initContainers.env.valueFrom.resourceFieldRef.containerName

     Container name: required for volumes, optional for env vars

- divisor	\<string\>
**PATH:**  statefulset.spec.template.spec.initContainers.env.valueFrom.resourceFieldRef.divisor

     Specifies the output format of the exposed resources, defaults to "1"

- resourceFieldRef	\<Object\>
**PATH:**  statefulset.spec.template.spec.initContainers.env.valueFrom.resourceFieldRef

     Selects a resource of the container: only resources limits and requests
     (limits.cpu, limits.memory, limits.ephemeral-storage, requests.cpu,
     requests.memory and requests.ephemeral-storage) are currently supported.

- secretKeyRef \<Object\>
**PATH:**  statefulset.spec.template.spec.initContainers.env.valueFrom.secretKeyRef

KIND:     StatefulSet
VERSION:  apps/v1

RESOURCE: secretKeyRef <Object>

DESCRIPTION:
     Selects a key of a secret in the pod's namespace

     SecretKeySelector selects a key of a Secret.

- key	\<string\> -required-
**PATH:**  statefulset.spec.template.spec.initContainers.env.valueFrom.secretKeyRef.key

     The key of the secret to select from. Must be a valid secret key.

- name	\<string\>
**PATH:**  statefulset.spec.template.spec.initContainers.env.valueFrom.secretKeyRef.name

     Name of the referent. More info:
     https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names

##### env	\<[]Object\>
**PATH:**  statefulset.spec.template.spec.initContainers.env

     List of environment variables to set in the container. Cannot be updated.

##### envFrom \<[]Object\>
**PATH:**  statefulset.spec.template.spec.initContainers.envFrom

KIND:     StatefulSet
VERSION:  apps/v1

RESOURCE: envFrom <[]Object>

DESCRIPTION:
     List of sources to populate environment variables in the container. The
     keys defined within a source must be a C_IDENTIFIER. All invalid keys will
     be reported as an event when the container is starting. When a key exists
     in multiple sources, the value associated with the last source will take
     precedence. Values defined by an Env with a duplicate key will take
     precedence. Cannot be updated.

     EnvFromSource represents the source of a set of ConfigMaps

###### configMapRef \<Object\>
**PATH:**  statefulset.spec.template.spec.initContainers.envFrom.configMapRef

KIND:     StatefulSet
VERSION:  apps/v1

RESOURCE: configMapRef <Object>

DESCRIPTION:
     The ConfigMap to select from

     ConfigMapEnvSource selects a ConfigMap to populate the environment
     variables with. The contents of the target ConfigMap's Data field will
     represent the key-value pairs as environment variables.

- name	\<string\>
**PATH:**  statefulset.spec.template.spec.initContainers.envFrom.configMapRef.name

     Name of the referent. More info:
     https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names

###### configMapRef	\<Object\>
**PATH:**  statefulset.spec.template.spec.initContainers.envFrom.configMapRef

     The ConfigMap to select from

###### prefix	\<string\>
**PATH:**  statefulset.spec.template.spec.initContainers.envFrom.prefix

     An optional identifier to prepend to each key in the ConfigMap. Must be a
     C_IDENTIFIER.

###### secretRef \<Object\>
**PATH:**  statefulset.spec.template.spec.initContainers.envFrom.secretRef

KIND:     StatefulSet
VERSION:  apps/v1

RESOURCE: secretRef <Object>

DESCRIPTION:
     The Secret to select from

     SecretEnvSource selects a Secret to populate the environment variables
     with. The contents of the target Secret's Data field will represent the
     key-value pairs as environment variables.

- name	\<string\>
**PATH:**  statefulset.spec.template.spec.initContainers.envFrom.secretRef.name

     Name of the referent. More info:
     https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names

##### envFrom	\<[]Object\>
**PATH:**  statefulset.spec.template.spec.initContainers.envFrom

     List of sources to populate environment variables in the container. The
     keys defined within a source must be a C_IDENTIFIER. All invalid keys will
     be reported as an event when the container is starting. When a key exists
     in multiple sources, the value associated with the last source will take
     precedence. Values defined by an Env with a duplicate key will take
     precedence. Cannot be updated.

##### image	\<string\>
**PATH:**  statefulset.spec.template.spec.initContainers.image

     Docker image name. More info:
     https://kubernetes.io/docs/concepts/containers/images This field is
     optional to allow higher level config management to default or override
     container images in workload controllers like Deployments and StatefulSets.

##### imagePullPolicy	\<string\>
**PATH:**  statefulset.spec.template.spec.initContainers.imagePullPolicy

     Image pull policy. One of Always, Never, IfNotPresent. Defaults to Always
     if :latest tag is specified, or IfNotPresent otherwise. Cannot be updated.
     More info:
     https://kubernetes.io/docs/concepts/containers/images#updating-images

##### lifecycle \<Object\>
**PATH:**  statefulset.spec.template.spec.initContainers.lifecycle

KIND:     StatefulSet
VERSION:  apps/v1

RESOURCE: lifecycle <Object>

DESCRIPTION:
     Actions that the management system should take in response to container
     lifecycle events. Cannot be updated.

     Lifecycle describes actions that the management system should take in
     response to container lifecycle events. For the PostStart and PreStop
     lifecycle handlers, management of the container blocks until the action is
     complete, unless the container process fails, in which case the handler is
     aborted.

###### postStart \<Object\>
**PATH:**  statefulset.spec.template.spec.initContainers.lifecycle.postStart

KIND:     StatefulSet
VERSION:  apps/v1

RESOURCE: postStart <Object>

DESCRIPTION:
     PostStart is called immediately after a container is created. If the
     handler fails, the container is terminated and restarted according to its
     restart policy. Other management of the container blocks until the hook
     completes. More info:
     https://kubernetes.io/docs/concepts/containers/container-lifecycle-hooks/#container-hooks

     Handler defines a specific action that should be taken

- exec \<Object\>
**PATH:**  statefulset.spec.template.spec.initContainers.lifecycle.postStart.exec

KIND:     StatefulSet
VERSION:  apps/v1

RESOURCE: exec <Object>

DESCRIPTION:
     One and only one of the following should be specified. Exec specifies the
     action to take.

     ExecAction describes a "run in container" action.

- exec	\<Object\>
**PATH:**  statefulset.spec.template.spec.initContainers.lifecycle.postStart.exec

     One and only one of the following should be specified. Exec specifies the
     action to take.

- httpGet \<Object\>
**PATH:**  statefulset.spec.template.spec.initContainers.lifecycle.postStart.httpGet

KIND:     StatefulSet
VERSION:  apps/v1

RESOURCE: httpGet <Object>

DESCRIPTION:
     HTTPGet specifies the http request to perform.

     HTTPGetAction describes an action based on HTTP Get requests.

- host	\<string\>
**PATH:**  statefulset.spec.template.spec.initContainers.lifecycle.postStart.httpGet.host

     Host name to connect to, defaults to the pod IP. You probably want to set
     "Host" in httpHeaders instead.

- httpHeaders \<[]Object\>
**PATH:**  statefulset.spec.template.spec.initContainers.lifecycle.postStart.httpGet.httpHeaders

KIND:     StatefulSet
VERSION:  apps/v1

RESOURCE: httpHeaders <[]Object>

DESCRIPTION:
     Custom headers to set in the request. HTTP allows repeated headers.

     HTTPHeader describes a custom header to be used in HTTP probes

- name	\<string\> -required-
**PATH:**  statefulset.spec.template.spec.initContainers.lifecycle.postStart.httpGet.httpHeaders.name

     The header field name

- httpHeaders	\<[]Object\>
**PATH:**  statefulset.spec.template.spec.initContainers.lifecycle.postStart.httpGet.httpHeaders

     Custom headers to set in the request. HTTP allows repeated headers.

- path	\<string\>
**PATH:**  statefulset.spec.template.spec.initContainers.lifecycle.postStart.httpGet.path

     Path to access on the HTTP server.

- port	\<string\> -required-
**PATH:**  statefulset.spec.template.spec.initContainers.lifecycle.postStart.httpGet.port

     Name or number of the port to access on the container. Number must be in
     the range 1 to 65535. Name must be an IANA_SVC_NAME.

- httpGet	\<Object\>
**PATH:**  statefulset.spec.template.spec.initContainers.lifecycle.postStart.httpGet

     HTTPGet specifies the http request to perform.

- tcpSocket \<Object\>
**PATH:**  statefulset.spec.template.spec.initContainers.lifecycle.postStart.tcpSocket

KIND:     StatefulSet
VERSION:  apps/v1

RESOURCE: tcpSocket <Object>

DESCRIPTION:
     TCPSocket specifies an action involving a TCP port. TCP hooks not yet
     supported

     TCPSocketAction describes an action based on opening a socket

- host	\<string\>
**PATH:**  statefulset.spec.template.spec.initContainers.lifecycle.postStart.tcpSocket.host

     Optional: Host name to connect to, defaults to the pod IP.

###### postStart	\<Object\>
**PATH:**  statefulset.spec.template.spec.initContainers.lifecycle.postStart

     PostStart is called immediately after a container is created. If the
     handler fails, the container is terminated and restarted according to its
     restart policy. Other management of the container blocks until the hook
     completes. More info:
     https://kubernetes.io/docs/concepts/containers/container-lifecycle-hooks/#container-hooks

###### preStop \<Object\>
**PATH:**  statefulset.spec.template.spec.initContainers.lifecycle.preStop

KIND:     StatefulSet
VERSION:  apps/v1

RESOURCE: preStop <Object>

DESCRIPTION:
     PreStop is called immediately before a container is terminated due to an
     API request or management event such as liveness probe failure, preemption,
     resource contention, etc. The handler is not called if the container
     crashes or exits. The reason for termination is passed to the handler. The
     Pod's termination grace period countdown begins before the PreStop hooked
     is executed. Regardless of the outcome of the handler, the container will
     eventually terminate within the Pod's termination grace period. Other
     management of the container blocks until the hook completes or until the
     termination grace period is reached. More info:
     https://kubernetes.io/docs/concepts/containers/container-lifecycle-hooks/#container-hooks

     Handler defines a specific action that should be taken

- exec \<Object\>
**PATH:**  statefulset.spec.template.spec.initContainers.lifecycle.preStop.exec

KIND:     StatefulSet
VERSION:  apps/v1

RESOURCE: exec <Object>

DESCRIPTION:
     One and only one of the following should be specified. Exec specifies the
     action to take.

     ExecAction describes a "run in container" action.

- exec	\<Object\>
**PATH:**  statefulset.spec.template.spec.initContainers.lifecycle.preStop.exec

     One and only one of the following should be specified. Exec specifies the
     action to take.

- httpGet \<Object\>
**PATH:**  statefulset.spec.template.spec.initContainers.lifecycle.preStop.httpGet

KIND:     StatefulSet
VERSION:  apps/v1

RESOURCE: httpGet <Object>

DESCRIPTION:
     HTTPGet specifies the http request to perform.

     HTTPGetAction describes an action based on HTTP Get requests.

- host	\<string\>
**PATH:**  statefulset.spec.template.spec.initContainers.lifecycle.preStop.httpGet.host

     Host name to connect to, defaults to the pod IP. You probably want to set
     "Host" in httpHeaders instead.

- httpHeaders \<[]Object\>
**PATH:**  statefulset.spec.template.spec.initContainers.lifecycle.preStop.httpGet.httpHeaders

KIND:     StatefulSet
VERSION:  apps/v1

RESOURCE: httpHeaders <[]Object>

DESCRIPTION:
     Custom headers to set in the request. HTTP allows repeated headers.

     HTTPHeader describes a custom header to be used in HTTP probes

- name	\<string\> -required-
**PATH:**  statefulset.spec.template.spec.initContainers.lifecycle.preStop.httpGet.httpHeaders.name

     The header field name

- httpHeaders	\<[]Object\>
**PATH:**  statefulset.spec.template.spec.initContainers.lifecycle.preStop.httpGet.httpHeaders

     Custom headers to set in the request. HTTP allows repeated headers.

- path	\<string\>
**PATH:**  statefulset.spec.template.spec.initContainers.lifecycle.preStop.httpGet.path

     Path to access on the HTTP server.

- port	\<string\> -required-
**PATH:**  statefulset.spec.template.spec.initContainers.lifecycle.preStop.httpGet.port

     Name or number of the port to access on the container. Number must be in
     the range 1 to 65535. Name must be an IANA_SVC_NAME.

- httpGet	\<Object\>
**PATH:**  statefulset.spec.template.spec.initContainers.lifecycle.preStop.httpGet

     HTTPGet specifies the http request to perform.

- tcpSocket \<Object\>
**PATH:**  statefulset.spec.template.spec.initContainers.lifecycle.preStop.tcpSocket

KIND:     StatefulSet
VERSION:  apps/v1

RESOURCE: tcpSocket <Object>

DESCRIPTION:
     TCPSocket specifies an action involving a TCP port. TCP hooks not yet
     supported

     TCPSocketAction describes an action based on opening a socket

- host	\<string\>
**PATH:**  statefulset.spec.template.spec.initContainers.lifecycle.preStop.tcpSocket.host

     Optional: Host name to connect to, defaults to the pod IP.

##### lifecycle	\<Object\>
**PATH:**  statefulset.spec.template.spec.initContainers.lifecycle

     Actions that the management system should take in response to container
     lifecycle events. Cannot be updated.

##### livenessProbe \<Object\>
**PATH:**  statefulset.spec.template.spec.initContainers.livenessProbe

KIND:     StatefulSet
VERSION:  apps/v1

RESOURCE: livenessProbe <Object>

DESCRIPTION:
     Periodic probe of container liveness. Container will be restarted if the
     probe fails. Cannot be updated. More info:
     https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes

     Probe describes a health check to be performed against a container to
     determine whether it is alive or ready to receive traffic.

###### exec \<Object\>
**PATH:**  statefulset.spec.template.spec.initContainers.livenessProbe.exec

KIND:     StatefulSet
VERSION:  apps/v1

RESOURCE: exec <Object>

DESCRIPTION:
     One and only one of the following should be specified. Exec specifies the
     action to take.

     ExecAction describes a "run in container" action.

###### exec	\<Object\>
**PATH:**  statefulset.spec.template.spec.initContainers.livenessProbe.exec

     One and only one of the following should be specified. Exec specifies the
     action to take.

###### failureThreshold	\<integer\>
**PATH:**  statefulset.spec.template.spec.initContainers.livenessProbe.failureThreshold

     Minimum consecutive failures for the probe to be considered failed after
     having succeeded. Defaults to 3. Minimum value is 1.

###### httpGet \<Object\>
**PATH:**  statefulset.spec.template.spec.initContainers.livenessProbe.httpGet

KIND:     StatefulSet
VERSION:  apps/v1

RESOURCE: httpGet <Object>

DESCRIPTION:
     HTTPGet specifies the http request to perform.

     HTTPGetAction describes an action based on HTTP Get requests.

- host	\<string\>
**PATH:**  statefulset.spec.template.spec.initContainers.livenessProbe.httpGet.host

     Host name to connect to, defaults to the pod IP. You probably want to set
     "Host" in httpHeaders instead.

- httpHeaders \<[]Object\>
**PATH:**  statefulset.spec.template.spec.initContainers.livenessProbe.httpGet.httpHeaders

KIND:     StatefulSet
VERSION:  apps/v1

RESOURCE: httpHeaders <[]Object>

DESCRIPTION:
     Custom headers to set in the request. HTTP allows repeated headers.

     HTTPHeader describes a custom header to be used in HTTP probes

- name	\<string\> -required-
**PATH:**  statefulset.spec.template.spec.initContainers.livenessProbe.httpGet.httpHeaders.name

     The header field name

- httpHeaders	\<[]Object\>
**PATH:**  statefulset.spec.template.spec.initContainers.livenessProbe.httpGet.httpHeaders

     Custom headers to set in the request. HTTP allows repeated headers.

- path	\<string\>
**PATH:**  statefulset.spec.template.spec.initContainers.livenessProbe.httpGet.path

     Path to access on the HTTP server.

- port	\<string\> -required-
**PATH:**  statefulset.spec.template.spec.initContainers.livenessProbe.httpGet.port

     Name or number of the port to access on the container. Number must be in
     the range 1 to 65535. Name must be an IANA_SVC_NAME.

###### httpGet	\<Object\>
**PATH:**  statefulset.spec.template.spec.initContainers.livenessProbe.httpGet

     HTTPGet specifies the http request to perform.

###### initialDelaySeconds	\<integer\>
**PATH:**  statefulset.spec.template.spec.initContainers.livenessProbe.initialDelaySeconds

     Number of seconds after the container has started before liveness probes
     are initiated. More info:
     https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes

###### periodSeconds	\<integer\>
**PATH:**  statefulset.spec.template.spec.initContainers.livenessProbe.periodSeconds

     How often (in seconds) to perform the probe. Default to 10 seconds. Minimum
     value is 1.

###### successThreshold	\<integer\>
**PATH:**  statefulset.spec.template.spec.initContainers.livenessProbe.successThreshold

     Minimum consecutive successes for the probe to be considered successful
     after having failed. Defaults to 1. Must be 1 for liveness. Minimum value
     is 1.

###### tcpSocket \<Object\>
**PATH:**  statefulset.spec.template.spec.initContainers.livenessProbe.tcpSocket

KIND:     StatefulSet
VERSION:  apps/v1

RESOURCE: tcpSocket <Object>

DESCRIPTION:
     TCPSocket specifies an action involving a TCP port. TCP hooks not yet
     supported

     TCPSocketAction describes an action based on opening a socket

- host	\<string\>
**PATH:**  statefulset.spec.template.spec.initContainers.livenessProbe.tcpSocket.host

     Optional: Host name to connect to, defaults to the pod IP.

###### tcpSocket	\<Object\>
**PATH:**  statefulset.spec.template.spec.initContainers.livenessProbe.tcpSocket

     TCPSocket specifies an action involving a TCP port. TCP hooks not yet
     supported

##### livenessProbe	\<Object\>
**PATH:**  statefulset.spec.template.spec.initContainers.livenessProbe

     Periodic probe of container liveness. Container will be restarted if the
     probe fails. Cannot be updated. More info:
     https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes

##### name	\<string\> -required-
**PATH:**  statefulset.spec.template.spec.initContainers.name

     Name of the container specified as a DNS_LABEL. Each container in a pod
     must have a unique name (DNS_LABEL). Cannot be updated.

##### ports \<[]Object\>
**PATH:**  statefulset.spec.template.spec.initContainers.ports

KIND:     StatefulSet
VERSION:  apps/v1

RESOURCE: ports <[]Object>

DESCRIPTION:
     List of ports to expose from the container. Exposing a port here gives the
     system additional information about the network connections a container
     uses, but is primarily informational. Not specifying a port here DOES NOT
     prevent that port from being exposed. Any port which is listening on the
     default "0.0.0.0" address inside a container will be accessible from the
     network. Cannot be updated.

     ContainerPort represents a network port in a single container.

###### containerPort	\<integer\> -required-
**PATH:**  statefulset.spec.template.spec.initContainers.ports.containerPort

     Number of port to expose on the pod's IP address. This must be a valid port
     number, 0 < x < 65536.

###### hostIP	\<string\>
**PATH:**  statefulset.spec.template.spec.initContainers.ports.hostIP

     What host IP to bind the external port to.

###### hostPort	\<integer\>
**PATH:**  statefulset.spec.template.spec.initContainers.ports.hostPort

     Number of port to expose on the host. If specified, this must be a valid
     port number, 0 < x < 65536. If HostNetwork is specified, this must match
     ContainerPort. Most containers do not need this.

###### name	\<string\>
**PATH:**  statefulset.spec.template.spec.initContainers.ports.name

     If specified, this must be an IANA_SVC_NAME and unique within the pod. Each
     named port in a pod must have a unique name. Name for the port that can be
     referred to by services.

##### ports	\<[]Object\>
**PATH:**  statefulset.spec.template.spec.initContainers.ports

     List of ports to expose from the container. Exposing a port here gives the
     system additional information about the network connections a container
     uses, but is primarily informational. Not specifying a port here DOES NOT
     prevent that port from being exposed. Any port which is listening on the
     default "0.0.0.0" address inside a container will be accessible from the
     network. Cannot be updated.

##### readinessProbe \<Object\>
**PATH:**  statefulset.spec.template.spec.initContainers.readinessProbe

KIND:     StatefulSet
VERSION:  apps/v1

RESOURCE: readinessProbe <Object>

DESCRIPTION:
     Periodic probe of container service readiness. Container will be removed
     from service endpoints if the probe fails. Cannot be updated. More info:
     https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes

     Probe describes a health check to be performed against a container to
     determine whether it is alive or ready to receive traffic.

###### exec \<Object\>
**PATH:**  statefulset.spec.template.spec.initContainers.readinessProbe.exec

KIND:     StatefulSet
VERSION:  apps/v1

RESOURCE: exec <Object>

DESCRIPTION:
     One and only one of the following should be specified. Exec specifies the
     action to take.

     ExecAction describes a "run in container" action.

###### exec	\<Object\>
**PATH:**  statefulset.spec.template.spec.initContainers.readinessProbe.exec

     One and only one of the following should be specified. Exec specifies the
     action to take.

###### failureThreshold	\<integer\>
**PATH:**  statefulset.spec.template.spec.initContainers.readinessProbe.failureThreshold

     Minimum consecutive failures for the probe to be considered failed after
     having succeeded. Defaults to 3. Minimum value is 1.

###### httpGet \<Object\>
**PATH:**  statefulset.spec.template.spec.initContainers.readinessProbe.httpGet

KIND:     StatefulSet
VERSION:  apps/v1

RESOURCE: httpGet <Object>

DESCRIPTION:
     HTTPGet specifies the http request to perform.

     HTTPGetAction describes an action based on HTTP Get requests.

- host	\<string\>
**PATH:**  statefulset.spec.template.spec.initContainers.readinessProbe.httpGet.host

     Host name to connect to, defaults to the pod IP. You probably want to set
     "Host" in httpHeaders instead.

- httpHeaders \<[]Object\>
**PATH:**  statefulset.spec.template.spec.initContainers.readinessProbe.httpGet.httpHeaders

KIND:     StatefulSet
VERSION:  apps/v1

RESOURCE: httpHeaders <[]Object>

DESCRIPTION:
     Custom headers to set in the request. HTTP allows repeated headers.

     HTTPHeader describes a custom header to be used in HTTP probes

- name	\<string\> -required-
**PATH:**  statefulset.spec.template.spec.initContainers.readinessProbe.httpGet.httpHeaders.name

     The header field name

- httpHeaders	\<[]Object\>
**PATH:**  statefulset.spec.template.spec.initContainers.readinessProbe.httpGet.httpHeaders

     Custom headers to set in the request. HTTP allows repeated headers.

- path	\<string\>
**PATH:**  statefulset.spec.template.spec.initContainers.readinessProbe.httpGet.path

     Path to access on the HTTP server.

- port	\<string\> -required-
**PATH:**  statefulset.spec.template.spec.initContainers.readinessProbe.httpGet.port

     Name or number of the port to access on the container. Number must be in
     the range 1 to 65535. Name must be an IANA_SVC_NAME.

###### httpGet	\<Object\>
**PATH:**  statefulset.spec.template.spec.initContainers.readinessProbe.httpGet

     HTTPGet specifies the http request to perform.

###### initialDelaySeconds	\<integer\>
**PATH:**  statefulset.spec.template.spec.initContainers.readinessProbe.initialDelaySeconds

     Number of seconds after the container has started before liveness probes
     are initiated. More info:
     https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes

###### periodSeconds	\<integer\>
**PATH:**  statefulset.spec.template.spec.initContainers.readinessProbe.periodSeconds

     How often (in seconds) to perform the probe. Default to 10 seconds. Minimum
     value is 1.

###### successThreshold	\<integer\>
**PATH:**  statefulset.spec.template.spec.initContainers.readinessProbe.successThreshold

     Minimum consecutive successes for the probe to be considered successful
     after having failed. Defaults to 1. Must be 1 for liveness. Minimum value
     is 1.

###### tcpSocket \<Object\>
**PATH:**  statefulset.spec.template.spec.initContainers.readinessProbe.tcpSocket

KIND:     StatefulSet
VERSION:  apps/v1

RESOURCE: tcpSocket <Object>

DESCRIPTION:
     TCPSocket specifies an action involving a TCP port. TCP hooks not yet
     supported

     TCPSocketAction describes an action based on opening a socket

- host	\<string\>
**PATH:**  statefulset.spec.template.spec.initContainers.readinessProbe.tcpSocket.host

     Optional: Host name to connect to, defaults to the pod IP.

###### tcpSocket	\<Object\>
**PATH:**  statefulset.spec.template.spec.initContainers.readinessProbe.tcpSocket

     TCPSocket specifies an action involving a TCP port. TCP hooks not yet
     supported

##### readinessProbe	\<Object\>
**PATH:**  statefulset.spec.template.spec.initContainers.readinessProbe

     Periodic probe of container service readiness. Container will be removed
     from service endpoints if the probe fails. Cannot be updated. More info:
     https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes

##### resources \<Object\>
**PATH:**  statefulset.spec.template.spec.initContainers.resources

KIND:     StatefulSet
VERSION:  apps/v1

RESOURCE: resources <Object>

DESCRIPTION:
     Compute Resources required by this container. Cannot be updated. More info:
     https://kubernetes.io/docs/concepts/configuration/manage-compute-resources-container/

     ResourceRequirements describes the compute resource requirements.

###### limits	\<map[string]string\>
**PATH:**  statefulset.spec.template.spec.initContainers.resources.limits

     Limits describes the maximum amount of compute resources allowed. More
     info:
     https://kubernetes.io/docs/concepts/configuration/manage-compute-resources-container/

##### resources	\<Object\>
**PATH:**  statefulset.spec.template.spec.initContainers.resources

     Compute Resources required by this container. Cannot be updated. More info:
     https://kubernetes.io/docs/concepts/configuration/manage-compute-resources-container/

##### securityContext \<Object\>
**PATH:**  statefulset.spec.template.spec.initContainers.securityContext

KIND:     StatefulSet
VERSION:  apps/v1

RESOURCE: securityContext <Object>

DESCRIPTION:
     Security options the pod should run with. More info:
     https://kubernetes.io/docs/concepts/policy/security-context/ More info:
     https://kubernetes.io/docs/tasks/configure-pod-container/security-context/

     SecurityContext holds security configuration that will be applied to a
     container. Some fields are present in both SecurityContext and
     PodSecurityContext. When both are set, the values in SecurityContext take
     precedence.

###### allowPrivilegeEscalation	\<boolean\>
**PATH:**  statefulset.spec.template.spec.initContainers.securityContext.allowPrivilegeEscalation

     AllowPrivilegeEscalation controls whether a process can gain more
     privileges than its parent process. This bool directly controls if the
     no_new_privs flag will be set on the container process.
     AllowPrivilegeEscalation is true always when the container is: 1) run as
     Privileged 2) has CAP_SYS_ADMIN

###### capabilities \<Object\>
**PATH:**  statefulset.spec.template.spec.initContainers.securityContext.capabilities

KIND:     StatefulSet
VERSION:  apps/v1

RESOURCE: capabilities <Object>

DESCRIPTION:
     The capabilities to add/drop when running containers. Defaults to the
     default set of capabilities granted by the container runtime.

     Adds and removes POSIX capabilities from running containers.

- add	\<[]string\>
**PATH:**  statefulset.spec.template.spec.initContainers.securityContext.capabilities.add

     Added capabilities

###### capabilities	\<Object\>
**PATH:**  statefulset.spec.template.spec.initContainers.securityContext.capabilities

     The capabilities to add/drop when running containers. Defaults to the
     default set of capabilities granted by the container runtime.

###### privileged	\<boolean\>
**PATH:**  statefulset.spec.template.spec.initContainers.securityContext.privileged

     Run container in privileged mode. Processes in privileged containers are
     essentially equivalent to root on the host. Defaults to false.

###### procMount	\<string\>
**PATH:**  statefulset.spec.template.spec.initContainers.securityContext.procMount

     procMount denotes the type of proc mount to use for the containers. The
     default is DefaultProcMount which uses the container runtime defaults for
     readonly paths and masked paths. This requires the ProcMountType feature
     flag to be enabled.

###### readOnlyRootFilesystem	\<boolean\>
**PATH:**  statefulset.spec.template.spec.initContainers.securityContext.readOnlyRootFilesystem

     Whether this container has a read-only root filesystem. Default is false.

###### runAsGroup	\<integer\>
**PATH:**  statefulset.spec.template.spec.initContainers.securityContext.runAsGroup

     The GID to run the entrypoint of the container process. Uses runtime
     default if unset. May also be set in PodSecurityContext. If set in both
     SecurityContext and PodSecurityContext, the value specified in
     SecurityContext takes precedence.

###### runAsNonRoot	\<boolean\>
**PATH:**  statefulset.spec.template.spec.initContainers.securityContext.runAsNonRoot

     Indicates that the container must run as a non-root user. If true, the
     Kubelet will validate the image at runtime to ensure that it does not run
     as UID 0 (root) and fail to start the container if it does. If unset or
     false, no such validation will be performed. May also be set in
     PodSecurityContext. If set in both SecurityContext and PodSecurityContext,
     the value specified in SecurityContext takes precedence.

###### runAsUser	\<integer\>
**PATH:**  statefulset.spec.template.spec.initContainers.securityContext.runAsUser

     The UID to run the entrypoint of the container process. Defaults to user
     specified in image metadata if unspecified. May also be set in
     PodSecurityContext. If set in both SecurityContext and PodSecurityContext,
     the value specified in SecurityContext takes precedence.

###### seLinuxOptions \<Object\>
**PATH:**  statefulset.spec.template.spec.initContainers.securityContext.seLinuxOptions

KIND:     StatefulSet
VERSION:  apps/v1

RESOURCE: seLinuxOptions <Object>

DESCRIPTION:
     The SELinux context to be applied to the container. If unspecified, the
     container runtime will allocate a random SELinux context for each
     container. May also be set in PodSecurityContext. If set in both
     SecurityContext and PodSecurityContext, the value specified in
     SecurityContext takes precedence.

     SELinuxOptions are the labels to be applied to the container

- level	\<string\>
**PATH:**  statefulset.spec.template.spec.initContainers.securityContext.seLinuxOptions.level

     Level is SELinux level label that applies to the container.

- role	\<string\>
**PATH:**  statefulset.spec.template.spec.initContainers.securityContext.seLinuxOptions.role

     Role is a SELinux role label that applies to the container.

- type	\<string\>
**PATH:**  statefulset.spec.template.spec.initContainers.securityContext.seLinuxOptions.type

     Type is a SELinux type label that applies to the container.

##### securityContext	\<Object\>
**PATH:**  statefulset.spec.template.spec.initContainers.securityContext

     Security options the pod should run with. More info:
     https://kubernetes.io/docs/concepts/policy/security-context/ More info:
     https://kubernetes.io/docs/tasks/configure-pod-container/security-context/

##### stdin	\<boolean\>
**PATH:**  statefulset.spec.template.spec.initContainers.stdin

     Whether this container should allocate a buffer for stdin in the container
     runtime. If this is not set, reads from stdin in the container will always
     result in EOF. Default is false.

##### stdinOnce	\<boolean\>
**PATH:**  statefulset.spec.template.spec.initContainers.stdinOnce

     Whether the container runtime should close the stdin channel after it has
     been opened by a single attach. When stdin is true the stdin stream will
     remain open across multiple attach sessions. If stdinOnce is set to true,
     stdin is opened on container start, is empty until the first client
     attaches to stdin, and then remains open and accepts data until the client
     disconnects, at which time stdin is closed and remains closed until the
     container is restarted. If this flag is false, a container processes that
     reads from stdin will never receive an EOF. Default is false

##### terminationMessagePath	\<string\>
**PATH:**  statefulset.spec.template.spec.initContainers.terminationMessagePath

     Optional: Path at which the file to which the container's termination
     message will be written is mounted into the container's filesystem. Message
     written is intended to be brief final status, such as an assertion failure
     message. Will be truncated by the node if greater than 4096 bytes. The
     total message length across all containers will be limited to 12kb.
     Defaults to /dev/termination-log. Cannot be updated.

##### terminationMessagePolicy	\<string\>
**PATH:**  statefulset.spec.template.spec.initContainers.terminationMessagePolicy

     Indicate how the termination message should be populated. File will use the
     contents of terminationMessagePath to populate the container status message
     on both success and failure. FallbackToLogsOnError will use the last chunk
     of container log output if the termination message file is empty and the
     container exited with an error. The log output is limited to 2048 bytes or
     80 lines, whichever is smaller. Defaults to File. Cannot be updated.

##### tty	\<boolean\>
**PATH:**  statefulset.spec.template.spec.initContainers.tty

     Whether this container should allocate a TTY for itself, also requires
     'stdin' to be true. Default is false.

##### volumeDevices \<[]Object\>
**PATH:**  statefulset.spec.template.spec.initContainers.volumeDevices

KIND:     StatefulSet
VERSION:  apps/v1

RESOURCE: volumeDevices <[]Object>

DESCRIPTION:
     volumeDevices is the list of block devices to be used by the container.
     This is a beta feature.

     volumeDevice describes a mapping of a raw block device within a container.

###### devicePath	\<string\> -required-
**PATH:**  statefulset.spec.template.spec.initContainers.volumeDevices.devicePath

     devicePath is the path inside of the container that the device will be
     mapped to.

##### volumeDevices	\<[]Object\>
**PATH:**  statefulset.spec.template.spec.initContainers.volumeDevices

     volumeDevices is the list of block devices to be used by the container.
     This is a beta feature.

##### volumeMounts \<[]Object\>
**PATH:**  statefulset.spec.template.spec.initContainers.volumeMounts

KIND:     StatefulSet
VERSION:  apps/v1

RESOURCE: volumeMounts <[]Object>

DESCRIPTION:
     Pod volumes to mount into the container's filesystem. Cannot be updated.

     VolumeMount describes a mounting of a Volume within a container.

###### mountPath	\<string\> -required-
**PATH:**  statefulset.spec.template.spec.initContainers.volumeMounts.mountPath

     Path within the container at which the volume should be mounted. Must not
     contain ':'.

###### mountPropagation	\<string\>
**PATH:**  statefulset.spec.template.spec.initContainers.volumeMounts.mountPropagation

     mountPropagation determines how mounts are propagated from the host to
     container and the other way around. When not set, MountPropagationNone is
     used. This field is beta in 1.10.

###### name	\<string\> -required-
**PATH:**  statefulset.spec.template.spec.initContainers.volumeMounts.name

     This must match the Name of a Volume.

###### readOnly	\<boolean\>
**PATH:**  statefulset.spec.template.spec.initContainers.volumeMounts.readOnly

     Mounted read-only if true, read-write otherwise (false or unspecified).
     Defaults to false.

###### subPath	\<string\>
**PATH:**  statefulset.spec.template.spec.initContainers.volumeMounts.subPath

     Path within the volume from which the container's volume should be mounted.
     Defaults to "" (volume's root).

##### volumeMounts	\<[]Object\>
**PATH:**  statefulset.spec.template.spec.initContainers.volumeMounts

     Pod volumes to mount into the container's filesystem. Cannot be updated.

#### initContainers	\<[]Object\>
**PATH:**  statefulset.spec.template.spec.initContainers

     List of initialization containers belonging to the pod. Init containers are
     executed in order prior to containers being started. If any init container
     fails, the pod is considered to have failed and is handled according to its
     restartPolicy. The name for an init container or normal container must be
     unique among all containers. Init containers may not have Lifecycle
     actions, Readiness probes, or Liveness probes. The resourceRequirements of
     an init container are taken into account during scheduling by finding the
     highest request/limit for each resource type, and then using the max of of
     that value or the sum of the normal containers. Limits are applied to init
     containers in a similar fashion. Init containers cannot currently be added
     or removed. Cannot be updated. More info:
     https://kubernetes.io/docs/concepts/workloads/pods/init-containers/

#### nodeName	\<string\>
**PATH:**  statefulset.spec.template.spec.nodeName

     NodeName is a request to schedule this pod onto a specific node. If it is
     non-empty, the scheduler simply schedules this pod onto that node, assuming
     that it fits resource requirements.

#### nodeSelector	\<map[string]string\>
**PATH:**  statefulset.spec.template.spec.nodeSelector

     NodeSelector is a selector which must be true for the pod to fit on a node.
     Selector which must match a node's labels for the pod to be scheduled on
     that node. More info:
     https://kubernetes.io/docs/concepts/configuration/assign-pod-node/

#### priority	\<integer\>
**PATH:**  statefulset.spec.template.spec.priority

     The priority value. Various system components use this field to find the
     priority of the pod. When Priority Admission Controller is enabled, it
     prevents users from setting this field. The admission controller populates
     this field from PriorityClassName. The higher the value, the higher the
     priority.

#### priorityClassName	\<string\>
**PATH:**  statefulset.spec.template.spec.priorityClassName

     If specified, indicates the pod's priority. "system-node-critical" and
     "system-cluster-critical" are two special keywords which indicate the
     highest priorities with the former being the highest priority. Any other
     name must be defined by creating a PriorityClass object with that name. If
     not specified, the pod priority will be default or zero if there is no
     default.

#### readinessGates \<[]Object\>
**PATH:**  statefulset.spec.template.spec.readinessGates

KIND:     StatefulSet
VERSION:  apps/v1

RESOURCE: readinessGates <[]Object>

DESCRIPTION:
     If specified, all readiness gates will be evaluated for pod readiness. A
     pod is ready when all its containers are ready AND all conditions specified
     in the readiness gates have status equal to "True" More info:
     https://git.k8s.io/enhancements/keps/sig-network/0007-pod-ready%!B(MISSING)%!B(MISSING).md

     PodReadinessGate contains the reference to a pod condition

#### readinessGates	\<[]Object\>
**PATH:**  statefulset.spec.template.spec.readinessGates

     If specified, all readiness gates will be evaluated for pod readiness. A
     pod is ready when all its containers are ready AND all conditions specified
     in the readiness gates have status equal to "True" More info:
     https://git.k8s.io/enhancements/keps/sig-network/0007-pod-ready%2B%2B.md

#### restartPolicy	\<string\>
**PATH:**  statefulset.spec.template.spec.restartPolicy

     Restart policy for all containers within the pod. One of Always, OnFailure,
     Never. Default to Always. More info:
     https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle/#restart-policy

#### runtimeClassName	\<string\>
**PATH:**  statefulset.spec.template.spec.runtimeClassName

     RuntimeClassName refers to a RuntimeClass object in the node.k8s.io group,
     which should be used to run this pod. If no RuntimeClass resource matches
     the named class, the pod will not be run. If unset or empty, the "legacy"
     RuntimeClass will be used, which is an implicit class with an empty
     definition that uses the default runtime handler. More info:
     https://git.k8s.io/enhancements/keps/sig-node/runtime-class.md This is an
     alpha feature and may change in the future.

#### schedulerName	\<string\>
**PATH:**  statefulset.spec.template.spec.schedulerName

     If specified, the pod will be dispatched by specified scheduler. If not
     specified, the pod will be dispatched by default scheduler.

#### securityContext \<Object\>
**PATH:**  statefulset.spec.template.spec.securityContext

KIND:     StatefulSet
VERSION:  apps/v1

RESOURCE: securityContext <Object>

DESCRIPTION:
     SecurityContext holds pod-level security attributes and common container
     settings. Optional: Defaults to empty. See type description for default
     values of each field.

     PodSecurityContext holds pod-level security attributes and common container
     settings. Some fields are also present in container.securityContext. Field
     values of container.securityContext take precedence over field values of
     PodSecurityContext.

##### fsGroup	\<integer\>
**PATH:**  statefulset.spec.template.spec.securityContext.fsGroup

     A special supplemental group that applies to all containers in a pod. Some
     volume types allow the Kubelet to change the ownership of that volume to be
     owned by the pod: 1. The owning GID will be the FSGroup 2. The setgid bit
     is set (new files created in the volume will be owned by FSGroup) 3. The
     permission bits are OR'd with rw-rw---- If unset, the Kubelet will not
     modify the ownership and permissions of any volume.

##### runAsGroup	\<integer\>
**PATH:**  statefulset.spec.template.spec.securityContext.runAsGroup

     The GID to run the entrypoint of the container process. Uses runtime
     default if unset. May also be set in SecurityContext. If set in both
     SecurityContext and PodSecurityContext, the value specified in
     SecurityContext takes precedence for that container.

##### runAsNonRoot	\<boolean\>
**PATH:**  statefulset.spec.template.spec.securityContext.runAsNonRoot

     Indicates that the container must run as a non-root user. If true, the
     Kubelet will validate the image at runtime to ensure that it does not run
     as UID 0 (root) and fail to start the container if it does. If unset or
     false, no such validation will be performed. May also be set in
     SecurityContext. If set in both SecurityContext and PodSecurityContext, the
     value specified in SecurityContext takes precedence.

##### runAsUser	\<integer\>
**PATH:**  statefulset.spec.template.spec.securityContext.runAsUser

     The UID to run the entrypoint of the container process. Defaults to user
     specified in image metadata if unspecified. May also be set in
     SecurityContext. If set in both SecurityContext and PodSecurityContext, the
     value specified in SecurityContext takes precedence for that container.

##### seLinuxOptions \<Object\>
**PATH:**  statefulset.spec.template.spec.securityContext.seLinuxOptions

KIND:     StatefulSet
VERSION:  apps/v1

RESOURCE: seLinuxOptions <Object>

DESCRIPTION:
     The SELinux context to be applied to all containers. If unspecified, the
     container runtime will allocate a random SELinux context for each
     container. May also be set in SecurityContext. If set in both
     SecurityContext and PodSecurityContext, the value specified in
     SecurityContext takes precedence for that container.

     SELinuxOptions are the labels to be applied to the container

###### level	\<string\>
**PATH:**  statefulset.spec.template.spec.securityContext.seLinuxOptions.level

     Level is SELinux level label that applies to the container.

###### role	\<string\>
**PATH:**  statefulset.spec.template.spec.securityContext.seLinuxOptions.role

     Role is a SELinux role label that applies to the container.

###### type	\<string\>
**PATH:**  statefulset.spec.template.spec.securityContext.seLinuxOptions.type

     Type is a SELinux type label that applies to the container.

##### seLinuxOptions	\<Object\>
**PATH:**  statefulset.spec.template.spec.securityContext.seLinuxOptions

     The SELinux context to be applied to all containers. If unspecified, the
     container runtime will allocate a random SELinux context for each
     container. May also be set in SecurityContext. If set in both
     SecurityContext and PodSecurityContext, the value specified in
     SecurityContext takes precedence for that container.

   supplementalGroups	<[]integer>
     A list of groups applied to the first process run in each container, in
     addition to the container's primary GID. If unspecified, no groups will be
     added to any container.

##### sysctls \<[]Object\>
**PATH:**  statefulset.spec.template.spec.securityContext.sysctls

KIND:     StatefulSet
VERSION:  apps/v1

RESOURCE: sysctls <[]Object>

DESCRIPTION:
     Sysctls hold a list of namespaced sysctls used for the pod. Pods with
     unsupported sysctls (by the container runtime) might fail to launch.

     Sysctl defines a kernel parameter to be set

###### name	\<string\> -required-
**PATH:**  statefulset.spec.template.spec.securityContext.sysctls.name

     Name of a property to set

#### securityContext	\<Object\>
**PATH:**  statefulset.spec.template.spec.securityContext

     SecurityContext holds pod-level security attributes and common container
     settings. Optional: Defaults to empty. See type description for default
     values of each field.

#### serviceAccount	\<string\>
**PATH:**  statefulset.spec.template.spec.serviceAccount

     DeprecatedServiceAccount is a depreciated alias for ServiceAccountName.
     Deprecated: Use serviceAccountName instead.

#### serviceAccountName	\<string\>
**PATH:**  statefulset.spec.template.spec.serviceAccountName

     ServiceAccountName is the name of the ServiceAccount to use to run this
     pod. More info:
     https://kubernetes.io/docs/tasks/configure-pod-container/configure-service-account/

#### shareProcessNamespace	\<boolean\>
**PATH:**  statefulset.spec.template.spec.shareProcessNamespace

     Share a single process namespace between all of the containers in a pod.
     When this is set containers will be able to view and signal processes from
     other containers in the same pod, and the first process in each container
     will not be assigned PID 1. HostPID and ShareProcessNamespace cannot both
     be set. Optional: Default to false. This field is beta-level and may be
     disabled with the PodShareProcessNamespace feature.

#### subdomain	\<string\>
**PATH:**  statefulset.spec.template.spec.subdomain

     If specified, the fully qualified Pod hostname will be
     "<hostname>.<subdomain>.<pod namespace>.svc.<cluster domain>". If not
     specified, the pod will not have a domainname at all.

#### terminationGracePeriodSeconds	\<integer\>
**PATH:**  statefulset.spec.template.spec.terminationGracePeriodSeconds

     Optional duration in seconds the pod needs to terminate gracefully. May be
     decreased in delete request. Value must be non-negative integer. The value
     zero indicates delete immediately. If this value is nil, the default grace
     period will be used instead. The grace period is the duration in seconds
     after the processes running in the pod are sent a termination signal and
     the time when the processes are forcibly halted with a kill signal. Set
     this value longer than the expected cleanup time for your process. Defaults
     to 30 seconds.

#### tolerations \<[]Object\>
**PATH:**  statefulset.spec.template.spec.tolerations

KIND:     StatefulSet
VERSION:  apps/v1

RESOURCE: tolerations <[]Object>

DESCRIPTION:
     If specified, the pod's tolerations.

     The pod this Toleration is attached to tolerates any taint that matches the
     triple <key,value,effect> using the matching operator <operator>.

##### effect	\<string\>
**PATH:**  statefulset.spec.template.spec.tolerations.effect

     Effect indicates the taint effect to match. Empty means match all taint
     effects. When specified, allowed values are NoSchedule, PreferNoSchedule
     and NoExecute.

##### key	\<string\>
**PATH:**  statefulset.spec.template.spec.tolerations.key

     Key is the taint key that the toleration applies to. Empty means match all
     taint keys. If the key is empty, operator must be Exists; this combination
     means to match all values and all keys.

##### operator	\<string\>
**PATH:**  statefulset.spec.template.spec.tolerations.operator

     Operator represents a key's relationship to the value. Valid operators are
     Exists and Equal. Defaults to Equal. Exists is equivalent to wildcard for
     value, so that a pod can tolerate all taints of a particular category.

##### tolerationSeconds	\<integer\>
**PATH:**  statefulset.spec.template.spec.tolerations.tolerationSeconds

     TolerationSeconds represents the period of time the toleration (which must
     be of effect NoExecute, otherwise this field is ignored) tolerates the
     taint. By default, it is not set, which means tolerate the taint forever
     (do not evict). Zero and negative values will be treated as 0 (evict
     immediately) by the system.

#### tolerations	\<[]Object\>
**PATH:**  statefulset.spec.template.spec.tolerations

     If specified, the pod's tolerations.

#### volumes \<[]Object\>
**PATH:**  statefulset.spec.template.spec.volumes

KIND:     StatefulSet
VERSION:  apps/v1

RESOURCE: volumes <[]Object>

DESCRIPTION:
     List of volumes that can be mounted by containers belonging to the pod.
     More info: https://kubernetes.io/docs/concepts/storage/volumes

     Volume represents a named volume in a pod that may be accessed by any
     container in the pod.

##### awsElasticBlockStore \<Object\>
**PATH:**  statefulset.spec.template.spec.volumes.awsElasticBlockStore

KIND:     StatefulSet
VERSION:  apps/v1

RESOURCE: awsElasticBlockStore <Object>

DESCRIPTION:
     AWSElasticBlockStore represents an AWS Disk resource that is attached to a
     kubelet's host machine and then exposed to the pod. More info:
     https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore

     Represents a Persistent Disk resource in AWS. An AWS EBS disk must exist
     before mounting to a container. The disk must also be in the same AWS zone
     as the kubelet. An AWS EBS disk can only be mounted as read/write once. AWS
     EBS volumes support ownership management and SELinux relabeling.

###### fsType	\<string\>
**PATH:**  statefulset.spec.template.spec.volumes.awsElasticBlockStore.fsType

     Filesystem type of the volume that you want to mount. Tip: Ensure that the
     filesystem type is supported by the host operating system. Examples:
     "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified.
     More info:
     https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore

###### partition	\<integer\>
**PATH:**  statefulset.spec.template.spec.volumes.awsElasticBlockStore.partition

     The partition in the volume that you want to mount. If omitted, the default
     is to mount by volume name. Examples: For volume /dev/sda1, you specify the
     partition as "1". Similarly, the volume partition for /dev/sda is "0" (or
     you can leave the property empty).

###### readOnly	\<boolean\>
**PATH:**  statefulset.spec.template.spec.volumes.awsElasticBlockStore.readOnly

     Specify "true" to force and set the ReadOnly property in VolumeMounts to
     "true". If omitted, the default is "false". More info:
     https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore

##### awsElasticBlockStore	\<Object\>
**PATH:**  statefulset.spec.template.spec.volumes.awsElasticBlockStore

     AWSElasticBlockStore represents an AWS Disk resource that is attached to a
     kubelet's host machine and then exposed to the pod. More info:
     https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore

##### azureDisk \<Object\>
**PATH:**  statefulset.spec.template.spec.volumes.azureDisk

KIND:     StatefulSet
VERSION:  apps/v1

RESOURCE: azureDisk <Object>

DESCRIPTION:
     AzureDisk represents an Azure Data Disk mount on the host and bind mount to
     the pod.

     AzureDisk represents an Azure Data Disk mount on the host and bind mount to
     the pod.

###### cachingMode	\<string\>
**PATH:**  statefulset.spec.template.spec.volumes.azureDisk.cachingMode

     Host Caching mode: None, Read Only, Read Write.

###### diskName	\<string\> -required-
**PATH:**  statefulset.spec.template.spec.volumes.azureDisk.diskName

     The Name of the data disk in the blob storage

###### diskURI	\<string\> -required-
**PATH:**  statefulset.spec.template.spec.volumes.azureDisk.diskURI

     The URI the data disk in the blob storage

###### fsType	\<string\>
**PATH:**  statefulset.spec.template.spec.volumes.azureDisk.fsType

     Filesystem type to mount. Must be a filesystem type supported by the host
     operating system. Ex. "ext4", "xfs", "ntfs". Implicitly inferred to be
     "ext4" if unspecified.

###### kind	\<string\>
**PATH:**  statefulset.spec.template.spec.volumes.azureDisk.kind

     Expected values Shared: multiple blob disks per storage account Dedicated:
     single blob disk per storage account Managed: azure managed data disk (only
     in managed availability set). defaults to shared

##### azureDisk	\<Object\>
**PATH:**  statefulset.spec.template.spec.volumes.azureDisk

     AzureDisk represents an Azure Data Disk mount on the host and bind mount to
     the pod.

##### azureFile \<Object\>
**PATH:**  statefulset.spec.template.spec.volumes.azureFile

KIND:     StatefulSet
VERSION:  apps/v1

RESOURCE: azureFile <Object>

DESCRIPTION:
     AzureFile represents an Azure File Service mount on the host and bind mount
     to the pod.

     AzureFile represents an Azure File Service mount on the host and bind mount
     to the pod.

###### readOnly	\<boolean\>
**PATH:**  statefulset.spec.template.spec.volumes.azureFile.readOnly

     Defaults to false (read/write). ReadOnly here will force the ReadOnly
     setting in VolumeMounts.

###### secretName	\<string\> -required-
**PATH:**  statefulset.spec.template.spec.volumes.azureFile.secretName

     the name of secret that contains Azure Storage Account Name and Key

##### azureFile	\<Object\>
**PATH:**  statefulset.spec.template.spec.volumes.azureFile

     AzureFile represents an Azure File Service mount on the host and bind mount
     to the pod.

##### cephfs \<Object\>
**PATH:**  statefulset.spec.template.spec.volumes.cephfs

KIND:     StatefulSet
VERSION:  apps/v1

RESOURCE: cephfs <Object>

DESCRIPTION:
     CephFS represents a Ceph FS mount on the host that shares a pod's lifetime

     Represents a Ceph Filesystem mount that lasts the lifetime of a pod Cephfs
     volumes do not support ownership management or SELinux relabeling.

###### monitors	\<[]string\> -required-
**PATH:**  statefulset.spec.template.spec.volumes.cephfs.monitors

     Required: Monitors is a collection of Ceph monitors More info:
     https://releases.k8s.io/HEAD/examples/volumes/cephfs/README.md#how-to-use-it

###### path	\<string\>
**PATH:**  statefulset.spec.template.spec.volumes.cephfs.path

     Optional: Used as the mounted root, rather than the full Ceph tree, default
     is /

###### readOnly	\<boolean\>
**PATH:**  statefulset.spec.template.spec.volumes.cephfs.readOnly

     Optional: Defaults to false (read/write). ReadOnly here will force the
     ReadOnly setting in VolumeMounts. More info:
     https://releases.k8s.io/HEAD/examples/volumes/cephfs/README.md#how-to-use-it

###### secretFile	\<string\>
**PATH:**  statefulset.spec.template.spec.volumes.cephfs.secretFile

     Optional: SecretFile is the path to key ring for User, default is
     /etc/ceph/user.secret More info:
     https://releases.k8s.io/HEAD/examples/volumes/cephfs/README.md#how-to-use-it

###### secretRef \<Object\>
**PATH:**  statefulset.spec.template.spec.volumes.cephfs.secretRef

KIND:     StatefulSet
VERSION:  apps/v1

RESOURCE: secretRef <Object>

DESCRIPTION:
     Optional: SecretRef is reference to the authentication secret for User,
     default is empty. More info:
     https://releases.k8s.io/HEAD/examples/volumes/cephfs/README.md#how-to-use-it

     LocalObjectReference contains enough information to let you locate the
     referenced object inside the same namespace.

###### secretRef	\<Object\>
**PATH:**  statefulset.spec.template.spec.volumes.cephfs.secretRef

     Optional: SecretRef is reference to the authentication secret for User,
     default is empty. More info:
     https://releases.k8s.io/HEAD/examples/volumes/cephfs/README.md#how-to-use-it

##### cephfs	\<Object\>
**PATH:**  statefulset.spec.template.spec.volumes.cephfs

     CephFS represents a Ceph FS mount on the host that shares a pod's lifetime

##### cinder \<Object\>
**PATH:**  statefulset.spec.template.spec.volumes.cinder

KIND:     StatefulSet
VERSION:  apps/v1

RESOURCE: cinder <Object>

DESCRIPTION:
     Cinder represents a cinder volume attached and mounted on kubelets host
     machine More info:
     https://releases.k8s.io/HEAD/examples/mysql-cinder-pd/README.md

     Represents a cinder volume resource in Openstack. A Cinder volume must
     exist before mounting to a container. The volume must also be in the same
     region as the kubelet. Cinder volumes support ownership management and
     SELinux relabeling.

###### fsType	\<string\>
**PATH:**  statefulset.spec.template.spec.volumes.cinder.fsType

     Filesystem type to mount. Must be a filesystem type supported by the host
     operating system. Examples: "ext4", "xfs", "ntfs". Implicitly inferred to
     be "ext4" if unspecified. More info:
     https://releases.k8s.io/HEAD/examples/mysql-cinder-pd/README.md

###### readOnly	\<boolean\>
**PATH:**  statefulset.spec.template.spec.volumes.cinder.readOnly

     Optional: Defaults to false (read/write). ReadOnly here will force the
     ReadOnly setting in VolumeMounts. More info:
     https://releases.k8s.io/HEAD/examples/mysql-cinder-pd/README.md

###### secretRef \<Object\>
**PATH:**  statefulset.spec.template.spec.volumes.cinder.secretRef

KIND:     StatefulSet
VERSION:  apps/v1

RESOURCE: secretRef <Object>

DESCRIPTION:
     Optional: points to a secret object containing parameters used to connect
     to OpenStack.

     LocalObjectReference contains enough information to let you locate the
     referenced object inside the same namespace.

###### secretRef	\<Object\>
**PATH:**  statefulset.spec.template.spec.volumes.cinder.secretRef

     Optional: points to a secret object containing parameters used to connect
     to OpenStack.

##### cinder	\<Object\>
**PATH:**  statefulset.spec.template.spec.volumes.cinder

     Cinder represents a cinder volume attached and mounted on kubelets host
     machine More info:
     https://releases.k8s.io/HEAD/examples/mysql-cinder-pd/README.md

##### configMap \<Object\>
**PATH:**  statefulset.spec.template.spec.volumes.configMap

KIND:     StatefulSet
VERSION:  apps/v1

RESOURCE: configMap <Object>

DESCRIPTION:
     ConfigMap represents a configMap that should populate this volume

     Adapts a ConfigMap into a volume. The contents of the target ConfigMap's
     Data field will be presented in a volume as files using the keys in the
     Data field as the file names, unless the items element is populated with
     specific mappings of keys to paths. ConfigMap volumes support ownership
     management and SELinux relabeling.

###### defaultMode	\<integer\>
**PATH:**  statefulset.spec.template.spec.volumes.configMap.defaultMode

     Optional: mode bits to use on created files by default. Must be a value
     between 0 and 0777. Defaults to 0644. Directories within the path are not
     affected by this setting. This might be in conflict with other options that
     affect the file mode, like fsGroup, and the result can be other mode bits
     set.

###### items \<[]Object\>
**PATH:**  statefulset.spec.template.spec.volumes.configMap.items

KIND:     StatefulSet
VERSION:  apps/v1

RESOURCE: items <[]Object>

DESCRIPTION:
     If unspecified, each key-value pair in the Data field of the referenced
     ConfigMap will be projected into the volume as a file whose name is the key
     and content is the value. If specified, the listed keys will be projected
     into the specified paths, and unlisted keys will not be present. If a key
     is specified which is not present in the ConfigMap, the volume setup will
     error unless it is marked optional. Paths must be relative and may not
     contain the '..' path or start with '..'.

     Maps a string key to a path within a volume.

- key	\<string\> -required-
**PATH:**  statefulset.spec.template.spec.volumes.configMap.items.key

     The key to project.

- mode	\<integer\>
**PATH:**  statefulset.spec.template.spec.volumes.configMap.items.mode

     Optional: mode bits to use on this file, must be a value between 0 and
     0777. If not specified, the volume defaultMode will be used. This might be
     in conflict with other options that affect the file mode, like fsGroup, and
     the result can be other mode bits set.

###### items	\<[]Object\>
**PATH:**  statefulset.spec.template.spec.volumes.configMap.items

     If unspecified, each key-value pair in the Data field of the referenced
     ConfigMap will be projected into the volume as a file whose name is the key
     and content is the value. If specified, the listed keys will be projected
     into the specified paths, and unlisted keys will not be present. If a key
     is specified which is not present in the ConfigMap, the volume setup will
     error unless it is marked optional. Paths must be relative and may not
     contain the '..' path or start with '..'.

###### name	\<string\>
**PATH:**  statefulset.spec.template.spec.volumes.configMap.name

     Name of the referent. More info:
     https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names

##### configMap	\<Object\>
**PATH:**  statefulset.spec.template.spec.volumes.configMap

     ConfigMap represents a configMap that should populate this volume

##### csi \<Object\>
**PATH:**  statefulset.spec.template.spec.volumes.csi

KIND:     StatefulSet
VERSION:  apps/v1

RESOURCE: csi <Object>

DESCRIPTION:
     CSI (Container Storage Interface) represents storage that is handled by an
     external CSI driver (Alpha feature).

     Represents a source location of a volume to mount, managed by an external
     CSI driver

###### driver	\<string\> -required-
**PATH:**  statefulset.spec.template.spec.volumes.csi.driver

     Driver is the name of the CSI driver that handles this volume. Consult with
     your admin for the correct name as registered in the cluster.

###### fsType	\<string\>
**PATH:**  statefulset.spec.template.spec.volumes.csi.fsType

     Filesystem type to mount. Ex. "ext4", "xfs", "ntfs". If not provided, the
     empty value is passed to the associated CSI driver which will determine the
     default filesystem to apply.

###### nodePublishSecretRef \<Object\>
**PATH:**  statefulset.spec.template.spec.volumes.csi.nodePublishSecretRef

KIND:     StatefulSet
VERSION:  apps/v1

RESOURCE: nodePublishSecretRef <Object>

DESCRIPTION:
     NodePublishSecretRef is a reference to the secret object containing
     sensitive information to pass to the CSI driver to complete the CSI
     NodePublishVolume and NodeUnpublishVolume calls. This field is optional,
     and may be empty if no secret is required. If the secret object contains
     more than one secret, all secret references are passed.

     LocalObjectReference contains enough information to let you locate the
     referenced object inside the same namespace.

###### nodePublishSecretRef	\<Object\>
**PATH:**  statefulset.spec.template.spec.volumes.csi.nodePublishSecretRef

     NodePublishSecretRef is a reference to the secret object containing
     sensitive information to pass to the CSI driver to complete the CSI
     NodePublishVolume and NodeUnpublishVolume calls. This field is optional,
     and may be empty if no secret is required. If the secret object contains
     more than one secret, all secret references are passed.

###### readOnly	\<boolean\>
**PATH:**  statefulset.spec.template.spec.volumes.csi.readOnly

     Specifies a read-only configuration for the volume. Defaults to false
     (read/write).

##### csi	\<Object\>
**PATH:**  statefulset.spec.template.spec.volumes.csi

     CSI (Container Storage Interface) represents storage that is handled by an
     external CSI driver (Alpha feature).

##### downwardAPI \<Object\>
**PATH:**  statefulset.spec.template.spec.volumes.downwardAPI

KIND:     StatefulSet
VERSION:  apps/v1

RESOURCE: downwardAPI <Object>

DESCRIPTION:
     DownwardAPI represents downward API about the pod that should populate this
     volume

     DownwardAPIVolumeSource represents a volume containing downward API info.
     Downward API volumes support ownership management and SELinux relabeling.

###### defaultMode	\<integer\>
**PATH:**  statefulset.spec.template.spec.volumes.downwardAPI.defaultMode

     Optional: mode bits to use on created files by default. Must be a value
     between 0 and 0777. Defaults to 0644. Directories within the path are not
     affected by this setting. This might be in conflict with other options that
     affect the file mode, like fsGroup, and the result can be other mode bits
     set.

###### items \<[]Object\>
**PATH:**  statefulset.spec.template.spec.volumes.downwardAPI.items

KIND:     StatefulSet
VERSION:  apps/v1

RESOURCE: items <[]Object>

DESCRIPTION:
     Items is a list of downward API volume file

     DownwardAPIVolumeFile represents information to create the file containing
     the pod field

- fieldRef \<Object\>
**PATH:**  statefulset.spec.template.spec.volumes.downwardAPI.items.fieldRef

KIND:     StatefulSet
VERSION:  apps/v1

RESOURCE: fieldRef <Object>

DESCRIPTION:
     Required: Selects a field of the pod: only annotations, labels, name and
     namespace are supported.

     ObjectFieldSelector selects an APIVersioned field of an object.

- apiVersion	\<string\>
**PATH:**  statefulset.spec.template.spec.volumes.downwardAPI.items.fieldRef.apiVersion

     Version of the schema the FieldPath is written in terms of, defaults to
     "v1".

- fieldRef	\<Object\>
**PATH:**  statefulset.spec.template.spec.volumes.downwardAPI.items.fieldRef

     Required: Selects a field of the pod: only annotations, labels, name and
     namespace are supported.

- mode	\<integer\>
**PATH:**  statefulset.spec.template.spec.volumes.downwardAPI.items.mode

     Optional: mode bits to use on this file, must be a value between 0 and
     0777. If not specified, the volume defaultMode will be used. This might be
     in conflict with other options that affect the file mode, like fsGroup, and
     the result can be other mode bits set.

- path	\<string\> -required-
**PATH:**  statefulset.spec.template.spec.volumes.downwardAPI.items.path

     Required: Path is the relative path name of the file to be created. Must
     not be absolute or contain the '..' path. Must be utf-8 encoded. The first
     item of the relative path must not start with '..'

- resourceFieldRef \<Object\>
**PATH:**  statefulset.spec.template.spec.volumes.downwardAPI.items.resourceFieldRef

KIND:     StatefulSet
VERSION:  apps/v1

RESOURCE: resourceFieldRef <Object>

DESCRIPTION:
     Selects a resource of the container: only resources limits and requests
     (limits.cpu, limits.memory, requests.cpu and requests.memory) are currently
     supported.

     ResourceFieldSelector represents container resources (cpu, memory) and
     their output format

- containerName	\<string\>
**PATH:**  statefulset.spec.template.spec.volumes.downwardAPI.items.resourceFieldRef.containerName

     Container name: required for volumes, optional for env vars

- divisor	\<string\>
**PATH:**  statefulset.spec.template.spec.volumes.downwardAPI.items.resourceFieldRef.divisor

     Specifies the output format of the exposed resources, defaults to "1"

##### downwardAPI	\<Object\>
**PATH:**  statefulset.spec.template.spec.volumes.downwardAPI

     DownwardAPI represents downward API about the pod that should populate this
     volume

##### emptyDir \<Object\>
**PATH:**  statefulset.spec.template.spec.volumes.emptyDir

KIND:     StatefulSet
VERSION:  apps/v1

RESOURCE: emptyDir <Object>

DESCRIPTION:
     EmptyDir represents a temporary directory that shares a pod's lifetime.
     More info: https://kubernetes.io/docs/concepts/storage/volumes#emptydir

     Represents an empty directory for a pod. Empty directory volumes support
     ownership management and SELinux relabeling.

###### medium	\<string\>
**PATH:**  statefulset.spec.template.spec.volumes.emptyDir.medium

     What type of storage medium should back this directory. The default is ""
     which means to use the node's default medium. Must be an empty string
     (default) or Memory. More info:
     https://kubernetes.io/docs/concepts/storage/volumes#emptydir

##### emptyDir	\<Object\>
**PATH:**  statefulset.spec.template.spec.volumes.emptyDir

     EmptyDir represents a temporary directory that shares a pod's lifetime.
     More info: https://kubernetes.io/docs/concepts/storage/volumes#emptydir

##### fc \<Object\>
**PATH:**  statefulset.spec.template.spec.volumes.fc

KIND:     StatefulSet
VERSION:  apps/v1

RESOURCE: fc <Object>

DESCRIPTION:
     FC represents a Fibre Channel resource that is attached to a kubelet's host
     machine and then exposed to the pod.

     Represents a Fibre Channel volume. Fibre Channel volumes can only be
     mounted as read/write once. Fibre Channel volumes support ownership
     management and SELinux relabeling.

###### fsType	\<string\>
**PATH:**  statefulset.spec.template.spec.volumes.fc.fsType

     Filesystem type to mount. Must be a filesystem type supported by the host
     operating system. Ex. "ext4", "xfs", "ntfs". Implicitly inferred to be
     "ext4" if unspecified.

###### lun	\<integer\>
**PATH:**  statefulset.spec.template.spec.volumes.fc.lun

     Optional: FC target lun number

###### readOnly	\<boolean\>
**PATH:**  statefulset.spec.template.spec.volumes.fc.readOnly

     Optional: Defaults to false (read/write). ReadOnly here will force the
     ReadOnly setting in VolumeMounts.

###### targetWWNs	\<[]string\>
**PATH:**  statefulset.spec.template.spec.volumes.fc.targetWWNs

     Optional: FC target worldwide names (WWNs)

##### fc	\<Object\>
**PATH:**  statefulset.spec.template.spec.volumes.fc

     FC represents a Fibre Channel resource that is attached to a kubelet's host
     machine and then exposed to the pod.

##### flexVolume \<Object\>
**PATH:**  statefulset.spec.template.spec.volumes.flexVolume

KIND:     StatefulSet
VERSION:  apps/v1

RESOURCE: flexVolume <Object>

DESCRIPTION:
     FlexVolume represents a generic volume resource that is
     provisioned/attached using an exec based plugin.

     FlexVolume represents a generic volume resource that is
     provisioned/attached using an exec based plugin.

###### driver	\<string\> -required-
**PATH:**  statefulset.spec.template.spec.volumes.flexVolume.driver

     Driver is the name of the driver to use for this volume.

###### fsType	\<string\>
**PATH:**  statefulset.spec.template.spec.volumes.flexVolume.fsType

     Filesystem type to mount. Must be a filesystem type supported by the host
     operating system. Ex. "ext4", "xfs", "ntfs". The default filesystem depends
     on FlexVolume script.

###### options	\<map[string]string\>
**PATH:**  statefulset.spec.template.spec.volumes.flexVolume.options

     Optional: Extra command options if any.

###### readOnly	\<boolean\>
**PATH:**  statefulset.spec.template.spec.volumes.flexVolume.readOnly

     Optional: Defaults to false (read/write). ReadOnly here will force the
     ReadOnly setting in VolumeMounts.

###### secretRef \<Object\>
**PATH:**  statefulset.spec.template.spec.volumes.flexVolume.secretRef

KIND:     StatefulSet
VERSION:  apps/v1

RESOURCE: secretRef <Object>

DESCRIPTION:
     Optional: SecretRef is reference to the secret object containing sensitive
     information to pass to the plugin scripts. This may be empty if no secret
     object is specified. If the secret object contains more than one secret,
     all secrets are passed to the plugin scripts.

     LocalObjectReference contains enough information to let you locate the
     referenced object inside the same namespace.

##### flexVolume	\<Object\>
**PATH:**  statefulset.spec.template.spec.volumes.flexVolume

     FlexVolume represents a generic volume resource that is
     provisioned/attached using an exec based plugin.

##### flocker \<Object\>
**PATH:**  statefulset.spec.template.spec.volumes.flocker

KIND:     StatefulSet
VERSION:  apps/v1

RESOURCE: flocker <Object>

DESCRIPTION:
     Flocker represents a Flocker volume attached to a kubelet's host machine.
     This depends on the Flocker control service being running

     Represents a Flocker volume mounted by the Flocker agent. One and only one
     of datasetName and datasetUUID should be set. Flocker volumes do not
     support ownership management or SELinux relabeling.

###### datasetName	\<string\>
**PATH:**  statefulset.spec.template.spec.volumes.flocker.datasetName

     Name of the dataset stored as metadata -> name on the dataset for Flocker
     should be considered as deprecated

##### flocker	\<Object\>
**PATH:**  statefulset.spec.template.spec.volumes.flocker

     Flocker represents a Flocker volume attached to a kubelet's host machine.
     This depends on the Flocker control service being running

##### gcePersistentDisk \<Object\>
**PATH:**  statefulset.spec.template.spec.volumes.gcePersistentDisk

KIND:     StatefulSet
VERSION:  apps/v1

RESOURCE: gcePersistentDisk <Object>

DESCRIPTION:
     GCEPersistentDisk represents a GCE Disk resource that is attached to a
     kubelet's host machine and then exposed to the pod. More info:
     https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk

     Represents a Persistent Disk resource in Google Compute Engine. A GCE PD
     must exist before mounting to a container. The disk must also be in the
     same GCE project and zone as the kubelet. A GCE PD can only be mounted as
     read/write once or read-only many times. GCE PDs support ownership
     management and SELinux relabeling.

###### fsType	\<string\>
**PATH:**  statefulset.spec.template.spec.volumes.gcePersistentDisk.fsType

     Filesystem type of the volume that you want to mount. Tip: Ensure that the
     filesystem type is supported by the host operating system. Examples:
     "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified.
     More info:
     https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk

###### partition	\<integer\>
**PATH:**  statefulset.spec.template.spec.volumes.gcePersistentDisk.partition

     The partition in the volume that you want to mount. If omitted, the default
     is to mount by volume name. Examples: For volume /dev/sda1, you specify the
     partition as "1". Similarly, the volume partition for /dev/sda is "0" (or
     you can leave the property empty). More info:
     https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk

###### pdName	\<string\> -required-
**PATH:**  statefulset.spec.template.spec.volumes.gcePersistentDisk.pdName

     Unique name of the PD resource in GCE. Used to identify the disk in GCE.
     More info:
     https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk

##### gcePersistentDisk	\<Object\>
**PATH:**  statefulset.spec.template.spec.volumes.gcePersistentDisk

     GCEPersistentDisk represents a GCE Disk resource that is attached to a
     kubelet's host machine and then exposed to the pod. More info:
     https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk

##### gitRepo \<Object\>
**PATH:**  statefulset.spec.template.spec.volumes.gitRepo

KIND:     StatefulSet
VERSION:  apps/v1

RESOURCE: gitRepo <Object>

DESCRIPTION:
     GitRepo represents a git repository at a particular revision. DEPRECATED:
     GitRepo is deprecated. To provision a container with a git repo, mount an
     EmptyDir into an InitContainer that clones the repo using git, then mount
     the EmptyDir into the Pod's container.

     Represents a volume that is populated with the contents of a git
     repository. Git repo volumes do not support ownership management. Git repo
     volumes support SELinux relabeling. DEPRECATED: GitRepo is deprecated. To
     provision a container with a git repo, mount an EmptyDir into an
     InitContainer that clones the repo using git, then mount the EmptyDir into
     the Pod's container.

###### directory	\<string\>
**PATH:**  statefulset.spec.template.spec.volumes.gitRepo.directory

     Target directory name. Must not contain or start with '..'. If '.' is
     supplied, the volume directory will be the git repository. Otherwise, if
     specified, the volume will contain the git repository in the subdirectory
     with the given name.

###### repository	\<string\> -required-
**PATH:**  statefulset.spec.template.spec.volumes.gitRepo.repository

     Repository URL

##### gitRepo	\<Object\>
**PATH:**  statefulset.spec.template.spec.volumes.gitRepo

     GitRepo represents a git repository at a particular revision. DEPRECATED:
     GitRepo is deprecated. To provision a container with a git repo, mount an
     EmptyDir into an InitContainer that clones the repo using git, then mount
     the EmptyDir into the Pod's container.

##### glusterfs \<Object\>
**PATH:**  statefulset.spec.template.spec.volumes.glusterfs

KIND:     StatefulSet
VERSION:  apps/v1

RESOURCE: glusterfs <Object>

DESCRIPTION:
     Glusterfs represents a Glusterfs mount on the host that shares a pod's
     lifetime. More info:
     https://releases.k8s.io/HEAD/examples/volumes/glusterfs/README.md

     Represents a Glusterfs mount that lasts the lifetime of a pod. Glusterfs
     volumes do not support ownership management or SELinux relabeling.

###### endpoints	\<string\> -required-
**PATH:**  statefulset.spec.template.spec.volumes.glusterfs.endpoints

     EndpointsName is the endpoint name that details Glusterfs topology. More
     info:
     https://releases.k8s.io/HEAD/examples/volumes/glusterfs/README.md#create-a-pod

###### path	\<string\> -required-
**PATH:**  statefulset.spec.template.spec.volumes.glusterfs.path

     Path is the Glusterfs volume path. More info:
     https://releases.k8s.io/HEAD/examples/volumes/glusterfs/README.md#create-a-pod

##### glusterfs	\<Object\>
**PATH:**  statefulset.spec.template.spec.volumes.glusterfs

     Glusterfs represents a Glusterfs mount on the host that shares a pod's
     lifetime. More info:
     https://releases.k8s.io/HEAD/examples/volumes/glusterfs/README.md

##### hostPath \<Object\>
**PATH:**  statefulset.spec.template.spec.volumes.hostPath

KIND:     StatefulSet
VERSION:  apps/v1

RESOURCE: hostPath <Object>

DESCRIPTION:
     HostPath represents a pre-existing file or directory on the host machine
     that is directly exposed to the container. This is generally used for
     system agents or other privileged things that are allowed to see the host
     machine. Most containers will NOT need this. More info:
     https://kubernetes.io/docs/concepts/storage/volumes#hostpath

     Represents a host path mapped into a pod. Host path volumes do not support
     ownership management or SELinux relabeling.

###### path	\<string\> -required-
**PATH:**  statefulset.spec.template.spec.volumes.hostPath.path

     Path of the directory on the host. If the path is a symlink, it will follow
     the link to the real path. More info:
     https://kubernetes.io/docs/concepts/storage/volumes#hostpath

##### hostPath	\<Object\>
**PATH:**  statefulset.spec.template.spec.volumes.hostPath

     HostPath represents a pre-existing file or directory on the host machine
     that is directly exposed to the container. This is generally used for
     system agents or other privileged things that are allowed to see the host
     machine. Most containers will NOT need this. More info:
     https://kubernetes.io/docs/concepts/storage/volumes#hostpath

##### iscsi \<Object\>
**PATH:**  statefulset.spec.template.spec.volumes.iscsi

KIND:     StatefulSet
VERSION:  apps/v1

RESOURCE: iscsi <Object>

DESCRIPTION:
     ISCSI represents an ISCSI Disk resource that is attached to a kubelet's
     host machine and then exposed to the pod. More info:
     https://releases.k8s.io/HEAD/examples/volumes/iscsi/README.md

     Represents an ISCSI disk. ISCSI volumes can only be mounted as read/write
     once. ISCSI volumes support ownership management and SELinux relabeling.

###### chapAuthDiscovery	\<boolean\>
**PATH:**  statefulset.spec.template.spec.volumes.iscsi.chapAuthDiscovery

     whether support iSCSI Discovery CHAP authentication

###### chapAuthSession	\<boolean\>
**PATH:**  statefulset.spec.template.spec.volumes.iscsi.chapAuthSession

     whether support iSCSI Session CHAP authentication

###### fsType	\<string\>
**PATH:**  statefulset.spec.template.spec.volumes.iscsi.fsType

     Filesystem type of the volume that you want to mount. Tip: Ensure that the
     filesystem type is supported by the host operating system. Examples:
     "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified.
     More info: https://kubernetes.io/docs/concepts/storage/volumes#iscsi

###### initiatorName	\<string\>
**PATH:**  statefulset.spec.template.spec.volumes.iscsi.initiatorName

     Custom iSCSI Initiator Name. If initiatorName is specified with
     iscsiInterface simultaneously, new iSCSI interface <target portal>:<volume
     name> will be created for the connection.

###### iqn	\<string\> -required-
**PATH:**  statefulset.spec.template.spec.volumes.iscsi.iqn

     Target iSCSI Qualified Name.

###### iscsiInterface	\<string\>
**PATH:**  statefulset.spec.template.spec.volumes.iscsi.iscsiInterface

     iSCSI Interface Name that uses an iSCSI transport. Defaults to 'default'
     (tcp).

###### lun	\<integer\> -required-
**PATH:**  statefulset.spec.template.spec.volumes.iscsi.lun

     iSCSI Target Lun number.

###### portals	\<[]string\>
**PATH:**  statefulset.spec.template.spec.volumes.iscsi.portals

     iSCSI Target Portal List. The portal is either an IP or ip_addr:port if the
     port is other than default (typically TCP ports 860 and 3260).

###### readOnly	\<boolean\>
**PATH:**  statefulset.spec.template.spec.volumes.iscsi.readOnly

     ReadOnly here will force the ReadOnly setting in VolumeMounts. Defaults to
     false.

###### secretRef \<Object\>
**PATH:**  statefulset.spec.template.spec.volumes.iscsi.secretRef

KIND:     StatefulSet
VERSION:  apps/v1

RESOURCE: secretRef <Object>

DESCRIPTION:
     CHAP Secret for iSCSI target and initiator authentication

     LocalObjectReference contains enough information to let you locate the
     referenced object inside the same namespace.

###### secretRef	\<Object\>
**PATH:**  statefulset.spec.template.spec.volumes.iscsi.secretRef

     CHAP Secret for iSCSI target and initiator authentication

##### iscsi	\<Object\>
**PATH:**  statefulset.spec.template.spec.volumes.iscsi

     ISCSI represents an ISCSI Disk resource that is attached to a kubelet's
     host machine and then exposed to the pod. More info:
     https://releases.k8s.io/HEAD/examples/volumes/iscsi/README.md

##### name	\<string\> -required-
**PATH:**  statefulset.spec.template.spec.volumes.name

     Volume's name. Must be a DNS_LABEL and unique within the pod. More info:
     https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names

##### nfs \<Object\>
**PATH:**  statefulset.spec.template.spec.volumes.nfs

KIND:     StatefulSet
VERSION:  apps/v1

RESOURCE: nfs <Object>

DESCRIPTION:
     NFS represents an NFS mount on the host that shares a pod's lifetime More
     info: https://kubernetes.io/docs/concepts/storage/volumes#nfs

     Represents an NFS mount that lasts the lifetime of a pod. NFS volumes do
     not support ownership management or SELinux relabeling.

###### path	\<string\> -required-
**PATH:**  statefulset.spec.template.spec.volumes.nfs.path

     Path that is exported by the NFS server. More info:
     https://kubernetes.io/docs/concepts/storage/volumes#nfs

###### readOnly	\<boolean\>
**PATH:**  statefulset.spec.template.spec.volumes.nfs.readOnly

     ReadOnly here will force the NFS export to be mounted with read-only
     permissions. Defaults to false. More info:
     https://kubernetes.io/docs/concepts/storage/volumes#nfs

##### nfs	\<Object\>
**PATH:**  statefulset.spec.template.spec.volumes.nfs

     NFS represents an NFS mount on the host that shares a pod's lifetime More
     info: https://kubernetes.io/docs/concepts/storage/volumes#nfs

##### persistentVolumeClaim \<Object\>
**PATH:**  statefulset.spec.template.spec.volumes.persistentVolumeClaim

KIND:     StatefulSet
VERSION:  apps/v1

RESOURCE: persistentVolumeClaim <Object>

DESCRIPTION:
     PersistentVolumeClaimVolumeSource represents a reference to a
     PersistentVolumeClaim in the same namespace. More info:
     https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims

     PersistentVolumeClaimVolumeSource references the user's PVC in the same
     namespace. This volume finds the bound PV and mounts that volume for the
     pod. A PersistentVolumeClaimVolumeSource is, essentially, a wrapper around
     another type of volume that is owned by someone else (the system).

###### claimName	\<string\> -required-
**PATH:**  statefulset.spec.template.spec.volumes.persistentVolumeClaim.claimName

     ClaimName is the name of a PersistentVolumeClaim in the same namespace as
     the pod using this volume. More info:
     https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims

##### persistentVolumeClaim	\<Object\>
**PATH:**  statefulset.spec.template.spec.volumes.persistentVolumeClaim

     PersistentVolumeClaimVolumeSource represents a reference to a
     PersistentVolumeClaim in the same namespace. More info:
     https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims

##### photonPersistentDisk \<Object\>
**PATH:**  statefulset.spec.template.spec.volumes.photonPersistentDisk

KIND:     StatefulSet
VERSION:  apps/v1

RESOURCE: photonPersistentDisk <Object>

DESCRIPTION:
     PhotonPersistentDisk represents a PhotonController persistent disk attached
     and mounted on kubelets host machine

     Represents a Photon Controller persistent disk resource.

###### fsType	\<string\>
**PATH:**  statefulset.spec.template.spec.volumes.photonPersistentDisk.fsType

     Filesystem type to mount. Must be a filesystem type supported by the host
     operating system. Ex. "ext4", "xfs", "ntfs". Implicitly inferred to be
     "ext4" if unspecified.

##### photonPersistentDisk	\<Object\>
**PATH:**  statefulset.spec.template.spec.volumes.photonPersistentDisk

     PhotonPersistentDisk represents a PhotonController persistent disk attached
     and mounted on kubelets host machine

##### portworxVolume \<Object\>
**PATH:**  statefulset.spec.template.spec.volumes.portworxVolume

KIND:     StatefulSet
VERSION:  apps/v1

RESOURCE: portworxVolume <Object>

DESCRIPTION:
     PortworxVolume represents a portworx volume attached and mounted on
     kubelets host machine

     PortworxVolumeSource represents a Portworx volume resource.

###### fsType	\<string\>
**PATH:**  statefulset.spec.template.spec.volumes.portworxVolume.fsType

     FSType represents the filesystem type to mount Must be a filesystem type
     supported by the host operating system. Ex. "ext4", "xfs". Implicitly
     inferred to be "ext4" if unspecified.

###### readOnly	\<boolean\>
**PATH:**  statefulset.spec.template.spec.volumes.portworxVolume.readOnly

     Defaults to false (read/write). ReadOnly here will force the ReadOnly
     setting in VolumeMounts.

##### portworxVolume	\<Object\>
**PATH:**  statefulset.spec.template.spec.volumes.portworxVolume

     PortworxVolume represents a portworx volume attached and mounted on
     kubelets host machine

##### projected \<Object\>
**PATH:**  statefulset.spec.template.spec.volumes.projected

KIND:     StatefulSet
VERSION:  apps/v1

RESOURCE: projected <Object>

DESCRIPTION:
     Items for all in one resources secrets, configmaps, and downward API

     Represents a projected volume source

###### defaultMode	\<integer\>
**PATH:**  statefulset.spec.template.spec.volumes.projected.defaultMode

     Mode bits to use on created files by default. Must be a value between 0 and
     0777. Directories within the path are not affected by this setting. This
     might be in conflict with other options that affect the file mode, like
     fsGroup, and the result can be other mode bits set.

###### sources \<[]Object\>
**PATH:**  statefulset.spec.template.spec.volumes.projected.sources

KIND:     StatefulSet
VERSION:  apps/v1

RESOURCE: sources <[]Object>

DESCRIPTION:
     list of volume projections

     Projection that may be projected along with other supported volume types

- configMap \<Object\>
**PATH:**  statefulset.spec.template.spec.volumes.projected.sources.configMap

KIND:     StatefulSet
VERSION:  apps/v1

RESOURCE: configMap <Object>

DESCRIPTION:
     information about the configMap data to project

     Adapts a ConfigMap into a projected volume. The contents of the target
     ConfigMap's Data field will be presented in a projected volume as files
     using the keys in the Data field as the file names, unless the items
     element is populated with specific mappings of keys to paths. Note that
     this is identical to a configmap volume source without the default mode.

- items \<[]Object\>
**PATH:**  statefulset.spec.template.spec.volumes.projected.sources.configMap.items

KIND:     StatefulSet
VERSION:  apps/v1

RESOURCE: items <[]Object>

DESCRIPTION:
     If unspecified, each key-value pair in the Data field of the referenced
     ConfigMap will be projected into the volume as a file whose name is the key
     and content is the value. If specified, the listed keys will be projected
     into the specified paths, and unlisted keys will not be present. If a key
     is specified which is not present in the ConfigMap, the volume setup will
     error unless it is marked optional. Paths must be relative and may not
     contain the '..' path or start with '..'.

     Maps a string key to a path within a volume.

- key	\<string\> -required-
**PATH:**  statefulset.spec.template.spec.volumes.projected.sources.configMap.items.key

     The key to project.

- mode	\<integer\>
**PATH:**  statefulset.spec.template.spec.volumes.projected.sources.configMap.items.mode

     Optional: mode bits to use on this file, must be a value between 0 and
     0777. If not specified, the volume defaultMode will be used. This might be
     in conflict with other options that affect the file mode, like fsGroup, and
     the result can be other mode bits set.

- items	\<[]Object\>
**PATH:**  statefulset.spec.template.spec.volumes.projected.sources.configMap.items

     If unspecified, each key-value pair in the Data field of the referenced
     ConfigMap will be projected into the volume as a file whose name is the key
     and content is the value. If specified, the listed keys will be projected
     into the specified paths, and unlisted keys will not be present. If a key
     is specified which is not present in the ConfigMap, the volume setup will
     error unless it is marked optional. Paths must be relative and may not
     contain the '..' path or start with '..'.

- name	\<string\>
**PATH:**  statefulset.spec.template.spec.volumes.projected.sources.configMap.name

     Name of the referent. More info:
     https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names

- configMap	\<Object\>
**PATH:**  statefulset.spec.template.spec.volumes.projected.sources.configMap

     information about the configMap data to project

- downwardAPI \<Object\>
**PATH:**  statefulset.spec.template.spec.volumes.projected.sources.downwardAPI

KIND:     StatefulSet
VERSION:  apps/v1

RESOURCE: downwardAPI <Object>

DESCRIPTION:
     information about the downwardAPI data to project

     Represents downward API info for projecting into a projected volume. Note
     that this is identical to a downwardAPI volume source without the default
     mode.

- items \<[]Object\>
**PATH:**  statefulset.spec.template.spec.volumes.projected.sources.downwardAPI.items

KIND:     StatefulSet
VERSION:  apps/v1

RESOURCE: items <[]Object>

DESCRIPTION:
     Items is a list of DownwardAPIVolume file

     DownwardAPIVolumeFile represents information to create the file containing
     the pod field

- fieldRef \<Object\>
**PATH:**  statefulset.spec.template.spec.volumes.projected.sources.downwardAPI.items.fieldRef

KIND:     StatefulSet
VERSION:  apps/v1

RESOURCE: fieldRef <Object>

DESCRIPTION:
     Required: Selects a field of the pod: only annotations, labels, name and
     namespace are supported.

     ObjectFieldSelector selects an APIVersioned field of an object.

- apiVersion	\<string\>
**PATH:**  statefulset.spec.template.spec.volumes.projected.sources.downwardAPI.items.fieldRef.apiVersion

     Version of the schema the FieldPath is written in terms of, defaults to
     "v1".

- fieldRef	\<Object\>
**PATH:**  statefulset.spec.template.spec.volumes.projected.sources.downwardAPI.items.fieldRef

     Required: Selects a field of the pod: only annotations, labels, name and
     namespace are supported.

- mode	\<integer\>
**PATH:**  statefulset.spec.template.spec.volumes.projected.sources.downwardAPI.items.mode

     Optional: mode bits to use on this file, must be a value between 0 and
     0777. If not specified, the volume defaultMode will be used. This might be
     in conflict with other options that affect the file mode, like fsGroup, and
     the result can be other mode bits set.

- path	\<string\> -required-
**PATH:**  statefulset.spec.template.spec.volumes.projected.sources.downwardAPI.items.path

     Required: Path is the relative path name of the file to be created. Must
     not be absolute or contain the '..' path. Must be utf-8 encoded. The first
     item of the relative path must not start with '..'

- resourceFieldRef \<Object\>
**PATH:**  statefulset.spec.template.spec.volumes.projected.sources.downwardAPI.items.resourceFieldRef

KIND:     StatefulSet
VERSION:  apps/v1

RESOURCE: resourceFieldRef <Object>

DESCRIPTION:
     Selects a resource of the container: only resources limits and requests
     (limits.cpu, limits.memory, requests.cpu and requests.memory) are currently
     supported.

     ResourceFieldSelector represents container resources (cpu, memory) and
     their output format

- containerName	\<string\>
**PATH:**  statefulset.spec.template.spec.volumes.projected.sources.downwardAPI.items.resourceFieldRef.containerName

     Container name: required for volumes, optional for env vars

- divisor	\<string\>
**PATH:**  statefulset.spec.template.spec.volumes.projected.sources.downwardAPI.items.resourceFieldRef.divisor

     Specifies the output format of the exposed resources, defaults to "1"

- downwardAPI	\<Object\>
**PATH:**  statefulset.spec.template.spec.volumes.projected.sources.downwardAPI

     information about the downwardAPI data to project

- secret \<Object\>
**PATH:**  statefulset.spec.template.spec.volumes.projected.sources.secret

KIND:     StatefulSet
VERSION:  apps/v1

RESOURCE: secret <Object>

DESCRIPTION:
     information about the secret data to project

     Adapts a secret into a projected volume. The contents of the target
     Secret's Data field will be presented in a projected volume as files using
     the keys in the Data field as the file names. Note that this is identical
     to a secret volume source without the default mode.

- items \<[]Object\>
**PATH:**  statefulset.spec.template.spec.volumes.projected.sources.secret.items

KIND:     StatefulSet
VERSION:  apps/v1

RESOURCE: items <[]Object>

DESCRIPTION:
     If unspecified, each key-value pair in the Data field of the referenced
     Secret will be projected into the volume as a file whose name is the key
     and content is the value. If specified, the listed keys will be projected
     into the specified paths, and unlisted keys will not be present. If a key
     is specified which is not present in the Secret, the volume setup will
     error unless it is marked optional. Paths must be relative and may not
     contain the '..' path or start with '..'.

     Maps a string key to a path within a volume.

- key	\<string\> -required-
**PATH:**  statefulset.spec.template.spec.volumes.projected.sources.secret.items.key

     The key to project.

- mode	\<integer\>
**PATH:**  statefulset.spec.template.spec.volumes.projected.sources.secret.items.mode

     Optional: mode bits to use on this file, must be a value between 0 and
     0777. If not specified, the volume defaultMode will be used. This might be
     in conflict with other options that affect the file mode, like fsGroup, and
     the result can be other mode bits set.

- items	\<[]Object\>
**PATH:**  statefulset.spec.template.spec.volumes.projected.sources.secret.items

     If unspecified, each key-value pair in the Data field of the referenced
     Secret will be projected into the volume as a file whose name is the key
     and content is the value. If specified, the listed keys will be projected
     into the specified paths, and unlisted keys will not be present. If a key
     is specified which is not present in the Secret, the volume setup will
     error unless it is marked optional. Paths must be relative and may not
     contain the '..' path or start with '..'.

- name	\<string\>
**PATH:**  statefulset.spec.template.spec.volumes.projected.sources.secret.name

     Name of the referent. More info:
     https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names

- secret	\<Object\>
**PATH:**  statefulset.spec.template.spec.volumes.projected.sources.secret

     information about the secret data to project

- serviceAccountToken \<Object\>
**PATH:**  statefulset.spec.template.spec.volumes.projected.sources.serviceAccountToken

KIND:     StatefulSet
VERSION:  apps/v1

RESOURCE: serviceAccountToken <Object>

DESCRIPTION:
     information about the serviceAccountToken data to project

     ServiceAccountTokenProjection represents a projected service account token
     volume. This projection can be used to insert a service account token into
     the pods runtime filesystem for use against APIs (Kubernetes API Server or
     otherwise).

- audience	\<string\>
**PATH:**  statefulset.spec.template.spec.volumes.projected.sources.serviceAccountToken.audience

     Audience is the intended audience of the token. A recipient of a token must
     identify itself with an identifier specified in the audience of the token,
     and otherwise should reject the token. The audience defaults to the
     identifier of the apiserver.

- expirationSeconds	\<integer\>
**PATH:**  statefulset.spec.template.spec.volumes.projected.sources.serviceAccountToken.expirationSeconds

     ExpirationSeconds is the requested duration of validity of the service
     account token. As the token approaches expiration, the kubelet volume
     plugin will proactively rotate the service account token. The kubelet will
     start trying to rotate the token if the token is older than 80 percent of
     its time to live or if the token is older than 24 hours.Defaults to 1 hour
     and must be at least 10 minutes.

##### projected	\<Object\>
**PATH:**  statefulset.spec.template.spec.volumes.projected

     Items for all in one resources secrets, configmaps, and downward API

##### quobyte \<Object\>
**PATH:**  statefulset.spec.template.spec.volumes.quobyte

KIND:     StatefulSet
VERSION:  apps/v1

RESOURCE: quobyte <Object>

DESCRIPTION:
     Quobyte represents a Quobyte mount on the host that shares a pod's lifetime

     Represents a Quobyte mount that lasts the lifetime of a pod. Quobyte
     volumes do not support ownership management or SELinux relabeling.

###### group	\<string\>
**PATH:**  statefulset.spec.template.spec.volumes.quobyte.group

     Group to map volume access to Default is no group

###### readOnly	\<boolean\>
**PATH:**  statefulset.spec.template.spec.volumes.quobyte.readOnly

     ReadOnly here will force the Quobyte volume to be mounted with read-only
     permissions. Defaults to false.

###### registry	\<string\> -required-
**PATH:**  statefulset.spec.template.spec.volumes.quobyte.registry

     Registry represents a single or multiple Quobyte Registry services
     specified as a string as host:port pair (multiple entries are separated
     with commas) which acts as the central registry for volumes

###### tenant	\<string\>
**PATH:**  statefulset.spec.template.spec.volumes.quobyte.tenant

     Tenant owning the given Quobyte volume in the Backend Used with dynamically
     provisioned Quobyte volumes, value is set by the plugin

###### user	\<string\>
**PATH:**  statefulset.spec.template.spec.volumes.quobyte.user

     User to map volume access to Defaults to serivceaccount user

##### quobyte	\<Object\>
**PATH:**  statefulset.spec.template.spec.volumes.quobyte

     Quobyte represents a Quobyte mount on the host that shares a pod's lifetime

##### rbd \<Object\>
**PATH:**  statefulset.spec.template.spec.volumes.rbd

KIND:     StatefulSet
VERSION:  apps/v1

RESOURCE: rbd <Object>

DESCRIPTION:
     RBD represents a Rados Block Device mount on the host that shares a pod's
     lifetime. More info:
     https://releases.k8s.io/HEAD/examples/volumes/rbd/README.md

     Represents a Rados Block Device mount that lasts the lifetime of a pod. RBD
     volumes support ownership management and SELinux relabeling.

###### fsType	\<string\>
**PATH:**  statefulset.spec.template.spec.volumes.rbd.fsType

     Filesystem type of the volume that you want to mount. Tip: Ensure that the
     filesystem type is supported by the host operating system. Examples:
     "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified.
     More info: https://kubernetes.io/docs/concepts/storage/volumes#rbd

###### image	\<string\> -required-
**PATH:**  statefulset.spec.template.spec.volumes.rbd.image

     The rados image name. More info:
     https://releases.k8s.io/HEAD/examples/volumes/rbd/README.md#how-to-use-it

###### keyring	\<string\>
**PATH:**  statefulset.spec.template.spec.volumes.rbd.keyring

     Keyring is the path to key ring for RBDUser. Default is /etc/ceph/keyring.
     More info:
     https://releases.k8s.io/HEAD/examples/volumes/rbd/README.md#how-to-use-it

###### monitors	\<[]string\> -required-
**PATH:**  statefulset.spec.template.spec.volumes.rbd.monitors

     A collection of Ceph monitors. More info:
     https://releases.k8s.io/HEAD/examples/volumes/rbd/README.md#how-to-use-it

###### pool	\<string\>
**PATH:**  statefulset.spec.template.spec.volumes.rbd.pool

     The rados pool name. Default is rbd. More info:
     https://releases.k8s.io/HEAD/examples/volumes/rbd/README.md#how-to-use-it

###### readOnly	\<boolean\>
**PATH:**  statefulset.spec.template.spec.volumes.rbd.readOnly

     ReadOnly here will force the ReadOnly setting in VolumeMounts. Defaults to
     false. More info:
     https://releases.k8s.io/HEAD/examples/volumes/rbd/README.md#how-to-use-it

###### secretRef \<Object\>
**PATH:**  statefulset.spec.template.spec.volumes.rbd.secretRef

KIND:     StatefulSet
VERSION:  apps/v1

RESOURCE: secretRef <Object>

DESCRIPTION:
     SecretRef is name of the authentication secret for RBDUser. If provided
     overrides keyring. Default is nil. More info:
     https://releases.k8s.io/HEAD/examples/volumes/rbd/README.md#how-to-use-it

     LocalObjectReference contains enough information to let you locate the
     referenced object inside the same namespace.

###### secretRef	\<Object\>
**PATH:**  statefulset.spec.template.spec.volumes.rbd.secretRef

     SecretRef is name of the authentication secret for RBDUser. If provided
     overrides keyring. Default is nil. More info:
     https://releases.k8s.io/HEAD/examples/volumes/rbd/README.md#how-to-use-it

##### rbd	\<Object\>
**PATH:**  statefulset.spec.template.spec.volumes.rbd

     RBD represents a Rados Block Device mount on the host that shares a pod's
     lifetime. More info:
     https://releases.k8s.io/HEAD/examples/volumes/rbd/README.md

##### scaleIO \<Object\>
**PATH:**  statefulset.spec.template.spec.volumes.scaleIO

KIND:     StatefulSet
VERSION:  apps/v1

RESOURCE: scaleIO <Object>

DESCRIPTION:
     ScaleIO represents a ScaleIO persistent volume attached and mounted on
     Kubernetes nodes.

     ScaleIOVolumeSource represents a persistent ScaleIO volume

###### fsType	\<string\>
**PATH:**  statefulset.spec.template.spec.volumes.scaleIO.fsType

     Filesystem type to mount. Must be a filesystem type supported by the host
     operating system. Ex. "ext4", "xfs", "ntfs". Default is "xfs".

###### gateway	\<string\> -required-
**PATH:**  statefulset.spec.template.spec.volumes.scaleIO.gateway

     The host address of the ScaleIO API Gateway.

###### protectionDomain	\<string\>
**PATH:**  statefulset.spec.template.spec.volumes.scaleIO.protectionDomain

     The name of the ScaleIO Protection Domain for the configured storage.

###### readOnly	\<boolean\>
**PATH:**  statefulset.spec.template.spec.volumes.scaleIO.readOnly

     Defaults to false (read/write). ReadOnly here will force the ReadOnly
     setting in VolumeMounts.

###### secretRef \<Object\>
**PATH:**  statefulset.spec.template.spec.volumes.scaleIO.secretRef

KIND:     StatefulSet
VERSION:  apps/v1

RESOURCE: secretRef <Object>

DESCRIPTION:
     SecretRef references to the secret for ScaleIO user and other sensitive
     information. If this is not provided, Login operation will fail.

     LocalObjectReference contains enough information to let you locate the
     referenced object inside the same namespace.

###### secretRef	\<Object\> -required-
**PATH:**  statefulset.spec.template.spec.volumes.scaleIO.secretRef

     SecretRef references to the secret for ScaleIO user and other sensitive
     information. If this is not provided, Login operation will fail.

###### sslEnabled	\<boolean\>
**PATH:**  statefulset.spec.template.spec.volumes.scaleIO.sslEnabled

     Flag to enable/disable SSL communication with Gateway, default false

###### storageMode	\<string\>
**PATH:**  statefulset.spec.template.spec.volumes.scaleIO.storageMode

     Indicates whether the storage for a volume should be ThickProvisioned or
     ThinProvisioned. Default is ThinProvisioned.

###### storagePool	\<string\>
**PATH:**  statefulset.spec.template.spec.volumes.scaleIO.storagePool

     The ScaleIO Storage Pool associated with the protection domain.

###### system	\<string\> -required-
**PATH:**  statefulset.spec.template.spec.volumes.scaleIO.system

     The name of the storage system as configured in ScaleIO.

##### scaleIO	\<Object\>
**PATH:**  statefulset.spec.template.spec.volumes.scaleIO

     ScaleIO represents a ScaleIO persistent volume attached and mounted on
     Kubernetes nodes.

##### secret \<Object\>
**PATH:**  statefulset.spec.template.spec.volumes.secret

KIND:     StatefulSet
VERSION:  apps/v1

RESOURCE: secret <Object>

DESCRIPTION:
     Secret represents a secret that should populate this volume. More info:
     https://kubernetes.io/docs/concepts/storage/volumes#secret

     Adapts a Secret into a volume. The contents of the target Secret's Data
     field will be presented in a volume as files using the keys in the Data
     field as the file names. Secret volumes support ownership management and
     SELinux relabeling.

###### defaultMode	\<integer\>
**PATH:**  statefulset.spec.template.spec.volumes.secret.defaultMode

     Optional: mode bits to use on created files by default. Must be a value
     between 0 and 0777. Defaults to 0644. Directories within the path are not
     affected by this setting. This might be in conflict with other options that
     affect the file mode, like fsGroup, and the result can be other mode bits
     set.

###### items \<[]Object\>
**PATH:**  statefulset.spec.template.spec.volumes.secret.items

KIND:     StatefulSet
VERSION:  apps/v1

RESOURCE: items <[]Object>

DESCRIPTION:
     If unspecified, each key-value pair in the Data field of the referenced
     Secret will be projected into the volume as a file whose name is the key
     and content is the value. If specified, the listed keys will be projected
     into the specified paths, and unlisted keys will not be present. If a key
     is specified which is not present in the Secret, the volume setup will
     error unless it is marked optional. Paths must be relative and may not
     contain the '..' path or start with '..'.

     Maps a string key to a path within a volume.

- key	\<string\> -required-
**PATH:**  statefulset.spec.template.spec.volumes.secret.items.key

     The key to project.

- mode	\<integer\>
**PATH:**  statefulset.spec.template.spec.volumes.secret.items.mode

     Optional: mode bits to use on this file, must be a value between 0 and
     0777. If not specified, the volume defaultMode will be used. This might be
     in conflict with other options that affect the file mode, like fsGroup, and
     the result can be other mode bits set.

###### items	\<[]Object\>
**PATH:**  statefulset.spec.template.spec.volumes.secret.items

     If unspecified, each key-value pair in the Data field of the referenced
     Secret will be projected into the volume as a file whose name is the key
     and content is the value. If specified, the listed keys will be projected
     into the specified paths, and unlisted keys will not be present. If a key
     is specified which is not present in the Secret, the volume setup will
     error unless it is marked optional. Paths must be relative and may not
     contain the '..' path or start with '..'.

###### optional	\<boolean\>
**PATH:**  statefulset.spec.template.spec.volumes.secret.optional

     Specify whether the Secret or it's keys must be defined

##### secret	\<Object\>
**PATH:**  statefulset.spec.template.spec.volumes.secret

     Secret represents a secret that should populate this volume. More info:
     https://kubernetes.io/docs/concepts/storage/volumes#secret

##### storageos \<Object\>
**PATH:**  statefulset.spec.template.spec.volumes.storageos

KIND:     StatefulSet
VERSION:  apps/v1

RESOURCE: storageos <Object>

DESCRIPTION:
     StorageOS represents a StorageOS volume attached and mounted on Kubernetes
     nodes.

     Represents a StorageOS persistent volume resource.

###### fsType	\<string\>
**PATH:**  statefulset.spec.template.spec.volumes.storageos.fsType

     Filesystem type to mount. Must be a filesystem type supported by the host
     operating system. Ex. "ext4", "xfs", "ntfs". Implicitly inferred to be
     "ext4" if unspecified.

###### readOnly	\<boolean\>
**PATH:**  statefulset.spec.template.spec.volumes.storageos.readOnly

     Defaults to false (read/write). ReadOnly here will force the ReadOnly
     setting in VolumeMounts.

###### secretRef \<Object\>
**PATH:**  statefulset.spec.template.spec.volumes.storageos.secretRef

KIND:     StatefulSet
VERSION:  apps/v1

RESOURCE: secretRef <Object>

DESCRIPTION:
     SecretRef specifies the secret to use for obtaining the StorageOS API
     credentials. If not specified, default values will be attempted.

     LocalObjectReference contains enough information to let you locate the
     referenced object inside the same namespace.

###### secretRef	\<Object\>
**PATH:**  statefulset.spec.template.spec.volumes.storageos.secretRef

     SecretRef specifies the secret to use for obtaining the StorageOS API
     credentials. If not specified, default values will be attempted.

###### volumeName	\<string\>
**PATH:**  statefulset.spec.template.spec.volumes.storageos.volumeName

     VolumeName is the human-readable name of the StorageOS volume. Volume names
     are only unique within a namespace.

##### storageos	\<Object\>
**PATH:**  statefulset.spec.template.spec.volumes.storageos

     StorageOS represents a StorageOS volume attached and mounted on Kubernetes
     nodes.

##### vsphereVolume \<Object\>
**PATH:**  statefulset.spec.template.spec.volumes.vsphereVolume

KIND:     StatefulSet
VERSION:  apps/v1

RESOURCE: vsphereVolume <Object>

DESCRIPTION:
     VsphereVolume represents a vSphere volume attached and mounted on kubelets
     host machine

     Represents a vSphere volume resource.

###### fsType	\<string\>
**PATH:**  statefulset.spec.template.spec.volumes.vsphereVolume.fsType

     Filesystem type to mount. Must be a filesystem type supported by the host
     operating system. Ex. "ext4", "xfs", "ntfs". Implicitly inferred to be
     "ext4" if unspecified.

###### storagePolicyID	\<string\>
**PATH:**  statefulset.spec.template.spec.volumes.vsphereVolume.storagePolicyID

     Storage Policy Based Management (SPBM) profile ID associated with the
     StoragePolicyName.

###### storagePolicyName	\<string\>
**PATH:**  statefulset.spec.template.spec.volumes.vsphereVolume.storagePolicyName

     Storage Policy Based Management (SPBM) profile name.

## template	\<Object\> -required-
**PATH:**  statefulset.spec.template

     template is the object that describes the pod that will be created if
     insufficient replicas are detected. Each pod stamped out by the StatefulSet
     will fulfill this Template, but have a unique identity from the rest of the
     StatefulSet.

## updateStrategy \<Object\>
**PATH:**  statefulset.spec.updateStrategy

KIND:     StatefulSet
VERSION:  apps/v1

RESOURCE: updateStrategy <Object>

DESCRIPTION:
     updateStrategy indicates the StatefulSetUpdateStrategy that will be
     employed to update Pods in the StatefulSet when a revision is made to
     Template.

     StatefulSetUpdateStrategy indicates the strategy that the StatefulSet
     controller will use to perform updates. It includes any additional
     parameters necessary to perform the update for the indicated strategy.

### rollingUpdate \<Object\>
**PATH:**  statefulset.spec.updateStrategy.rollingUpdate

KIND:     StatefulSet
VERSION:  apps/v1

RESOURCE: rollingUpdate <Object>

DESCRIPTION:
     RollingUpdate is used to communicate parameters when Type is
     RollingUpdateStatefulSetStrategyType.

     RollingUpdateStatefulSetStrategy is used to communicate parameter for
     RollingUpdateStatefulSetStrategyType.

### rollingUpdate	\<Object\>
**PATH:**  statefulset.spec.updateStrategy.rollingUpdate

     RollingUpdate is used to communicate parameters when Type is
     RollingUpdateStatefulSetStrategyType.

## updateStrategy	\<Object\>
**PATH:**  statefulset.spec.updateStrategy

     updateStrategy indicates the StatefulSetUpdateStrategy that will be
     employed to update Pods in the StatefulSet when a revision is made to
     Template.

## volumeClaimTemplates \<[]Object\>
**PATH:**  statefulset.spec.volumeClaimTemplates

KIND:     StatefulSet
VERSION:  apps/v1

RESOURCE: volumeClaimTemplates <[]Object>

DESCRIPTION:
     volumeClaimTemplates is a list of claims that pods are allowed to
     reference. The StatefulSet controller is responsible for mapping network
     identities to claims in a way that maintains the identity of a pod. Every
     claim in this list must have at least one matching (by name) volumeMount in
     one container in the template. A claim in this list takes precedence over
     any volumes in the template, with the same name.

     PersistentVolumeClaim is a user's request for and claim to a persistent
     volume

### apiVersion	\<string\>
**PATH:**  statefulset.spec.volumeClaimTemplates.apiVersion

     APIVersion defines the versioned schema of this representation of an
     object. Servers should convert recognized schemas to the latest internal
     value, and may reject unrecognized values. More info:
     https://git.k8s.io/community/contributors/devel/api-conventions.md#resources

### kind	\<string\>
**PATH:**  statefulset.spec.volumeClaimTemplates.kind

     Kind is a string value representing the REST resource this object
     represents. Servers may infer this from the endpoint the client submits
     requests to. Cannot be updated. In CamelCase. More info:
     https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds

### metadata \<Object\>
**PATH:**  statefulset.spec.volumeClaimTemplates.metadata

KIND:     StatefulSet
VERSION:  apps/v1

RESOURCE: metadata <Object>

DESCRIPTION:
     Standard object's metadata. More info:
     https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata

     ObjectMeta is metadata that all persisted resources must have, which
     includes all objects users must create.

#### annotations	\<map[string]string\>
**PATH:**  statefulset.spec.volumeClaimTemplates.metadata.annotations

     Annotations is an unstructured key value map stored with a resource that
     may be set by external tools to store and retrieve arbitrary metadata. They
     are not queryable and should be preserved when modifying objects. More
     info: http://kubernetes.io/docs/user-guide/annotations

#### clusterName	\<string\>
**PATH:**  statefulset.spec.volumeClaimTemplates.metadata.clusterName

     The name of the cluster which the object belongs to. This is used to
     distinguish resources with same name and namespace in different clusters.
     This field is not set anywhere right now and apiserver is going to ignore
     it if set in create or update request.

#### creationTimestamp	\<string\>
**PATH:**  statefulset.spec.volumeClaimTemplates.metadata.creationTimestamp

     CreationTimestamp is a timestamp representing the server time when this
     object was created. It is not guaranteed to be set in happens-before order
     across separate operations. Clients may not set this value. It is
     represented in RFC3339 form and is in UTC. Populated by the system.
     Read-only. Null for lists. More info:
     https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata

#### deletionGracePeriodSeconds	\<integer\>
**PATH:**  statefulset.spec.volumeClaimTemplates.metadata.deletionGracePeriodSeconds

     Number of seconds allowed for this object to gracefully terminate before it
     will be removed from the system. Only set when deletionTimestamp is also
     set. May only be shortened. Read-only.

#### deletionTimestamp	\<string\>
**PATH:**  statefulset.spec.volumeClaimTemplates.metadata.deletionTimestamp

     DeletionTimestamp is RFC 3339 date and time at which this resource will be
     deleted. This field is set by the server when a graceful deletion is
     requested by the user, and is not directly settable by a client. The
     resource is expected to be deleted (no longer visible from resource lists,
     and not reachable by name) after the time in this field, once the
     finalizers list is empty. As long as the finalizers list contains items,
     deletion is blocked. Once the deletionTimestamp is set, this value may not
     be unset or be set further into the future, although it may be shortened or
     the resource may be deleted prior to this time. For example, a user may
     request that a pod is deleted in 30 seconds. The Kubelet will react by
     sending a graceful termination signal to the containers in the pod. After
     that 30 seconds, the Kubelet will send a hard termination signal (SIGKILL)
     to the container and after cleanup, remove the pod from the API. In the
     presence of network partitions, this object may still exist after this
     timestamp, until an administrator or automated process can determine the
     resource is fully terminated. If not set, graceful deletion of the object
     has not been requested. Populated by the system when a graceful deletion is
     requested. Read-only. More info:
     https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata

#### finalizers	\<[]string\>
**PATH:**  statefulset.spec.volumeClaimTemplates.metadata.finalizers

     Must be empty before the object is deleted from the registry. Each entry is
     an identifier for the responsible component that will remove the entry from
     the list. If the deletionTimestamp of the object is non-nil, entries in
     this list can only be removed.

#### generateName	\<string\>
**PATH:**  statefulset.spec.volumeClaimTemplates.metadata.generateName

     GenerateName is an optional prefix, used by the server, to generate a
     unique name ONLY IF the Name field has not been provided. If this field is
     used, the name returned to the client will be different than the name
     passed. This value will also be combined with a unique suffix. The provided
     value has the same validation rules as the Name field, and may be truncated
     by the length of the suffix required to make the value unique on the
     server. If this field is specified and the generated name exists, the
     server will NOT return a 409 - instead, it will either return 201 Created
     or 500 with Reason ServerTimeout indicating a unique name could not be
     found in the time allotted, and the client should retry (optionally after
     the time indicated in the Retry-After header). Applied only if Name is not
     specified. More info:
     https://git.k8s.io/community/contributors/devel/api-conventions.md#idempotency

#### generation	\<integer\>
**PATH:**  statefulset.spec.volumeClaimTemplates.metadata.generation

     A sequence number representing a specific generation of the desired state.
     Populated by the system. Read-only.

#### initializers \<Object\>
**PATH:**  statefulset.spec.volumeClaimTemplates.metadata.initializers

KIND:     StatefulSet
VERSION:  apps/v1

RESOURCE: initializers <Object>

DESCRIPTION:
     An initializer is a controller which enforces some system invariant at
     object creation time. This field is a list of initializers that have not
     yet acted on this object. If nil or empty, this object has been completely
     initialized. Otherwise, the object is considered uninitialized and is
     hidden (in list/watch and get calls) from clients that haven't explicitly
     asked to observe uninitialized objects. When an object is created, the
     system will populate this list with the current set of initializers. Only
     privileged users may set or modify this list. Once it is empty, it may not
     be modified further by any user. DEPRECATED - initializers are an alpha
     field and will be removed in v1.15.

     Initializers tracks the progress of initialization.

##### pending \<[]Object\>
**PATH:**  statefulset.spec.volumeClaimTemplates.metadata.initializers.pending

KIND:     StatefulSet
VERSION:  apps/v1

RESOURCE: pending <[]Object>

DESCRIPTION:
     Pending is a list of initializers that must execute in order before this
     object is visible. When the last pending initializer is removed, and no
     failing result is set, the initializers struct will be set to nil and the
     object is considered as initialized and visible to all clients.

     Initializer is information about an initializer that has not yet completed.

##### pending	\<[]Object\> -required-
**PATH:**  statefulset.spec.volumeClaimTemplates.metadata.initializers.pending

     Pending is a list of initializers that must execute in order before this
     object is visible. When the last pending initializer is removed, and no
     failing result is set, the initializers struct will be set to nil and the
     object is considered as initialized and visible to all clients.

##### result \<Object\>
**PATH:**  statefulset.spec.volumeClaimTemplates.metadata.initializers.result

KIND:     StatefulSet
VERSION:  apps/v1

RESOURCE: result <Object>

DESCRIPTION:
     If result is set with the Failure field, the object will be persisted to
     storage and then deleted, ensuring that other clients can observe the
     deletion.

     Status is a return value for calls that don't return other objects.

###### apiVersion	\<string\>
**PATH:**  statefulset.spec.volumeClaimTemplates.metadata.initializers.result.apiVersion

     APIVersion defines the versioned schema of this representation of an
     object. Servers should convert recognized schemas to the latest internal
     value, and may reject unrecognized values. More info:
     https://git.k8s.io/community/contributors/devel/api-conventions.md#resources

###### code	\<integer\>
**PATH:**  statefulset.spec.volumeClaimTemplates.metadata.initializers.result.code

     Suggested HTTP return code for this status, 0 if not set.

###### details \<Object\>
**PATH:**  statefulset.spec.volumeClaimTemplates.metadata.initializers.result.details

KIND:     StatefulSet
VERSION:  apps/v1

RESOURCE: details <Object>

DESCRIPTION:
     Extended data associated with the reason. Each reason may define its own
     extended details. This field is optional and the data returned is not
     guaranteed to conform to any schema except that defined by the reason type.

     StatusDetails is a set of additional properties that MAY be set by the
     server to provide additional information about a response. The Reason field
     of a Status object defines what attributes will be set. Clients must ignore
     fields that do not match the defined type of each attribute, and should
     assume that any attribute may be empty, invalid, or under defined.

- causes \<[]Object\>
**PATH:**  statefulset.spec.volumeClaimTemplates.metadata.initializers.result.details.causes

KIND:     StatefulSet
VERSION:  apps/v1

RESOURCE: causes <[]Object>

DESCRIPTION:
     The Causes array includes more details associated with the StatusReason
     failure. Not all StatusReasons may provide detailed causes.

     StatusCause provides more information about an api.Status failure,
     including cases when multiple errors are encountered.

- field	\<string\>
**PATH:**  statefulset.spec.volumeClaimTemplates.metadata.initializers.result.details.causes.field

     The field of the resource that has caused this error, as named by its JSON
     serialization. May include dot and postfix notation for nested attributes.
     Arrays are zero-indexed. Fields may appear more than once in an array of
     causes due to fields having multiple errors. Optional. Examples: "name" -
     the field "name" on the current resource "items[0].name" - the field "name"
     on the first array entry in "items"

- message	\<string\>
**PATH:**  statefulset.spec.volumeClaimTemplates.metadata.initializers.result.details.causes.message

     A human-readable description of the cause of the error. This field may be
     presented as-is to a reader.

- causes	\<[]Object\>
**PATH:**  statefulset.spec.volumeClaimTemplates.metadata.initializers.result.details.causes

     The Causes array includes more details associated with the StatusReason
     failure. Not all StatusReasons may provide detailed causes.

- group	\<string\>
**PATH:**  statefulset.spec.volumeClaimTemplates.metadata.initializers.result.details.group

     The group attribute of the resource associated with the status
     StatusReason.

- kind	\<string\>
**PATH:**  statefulset.spec.volumeClaimTemplates.metadata.initializers.result.details.kind

     The kind attribute of the resource associated with the status StatusReason.
     On some operations may differ from the requested resource Kind. More info:
     https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds

- name	\<string\>
**PATH:**  statefulset.spec.volumeClaimTemplates.metadata.initializers.result.details.name

     The name attribute of the resource associated with the status StatusReason
     (when there is a single name which can be described).

- retryAfterSeconds	\<integer\>
**PATH:**  statefulset.spec.volumeClaimTemplates.metadata.initializers.result.details.retryAfterSeconds

     If specified, the time in seconds before the operation should be retried.
     Some errors may indicate the client must take an alternate action - for
     those errors this field may indicate how long to wait before taking the
     alternate action.

###### details	\<Object\>
**PATH:**  statefulset.spec.volumeClaimTemplates.metadata.initializers.result.details

     Extended data associated with the reason. Each reason may define its own
     extended details. This field is optional and the data returned is not
     guaranteed to conform to any schema except that defined by the reason type.

###### kind	\<string\>
**PATH:**  statefulset.spec.volumeClaimTemplates.metadata.initializers.result.kind

     Kind is a string value representing the REST resource this object
     represents. Servers may infer this from the endpoint the client submits
     requests to. Cannot be updated. In CamelCase. More info:
     https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds

###### message	\<string\>
**PATH:**  statefulset.spec.volumeClaimTemplates.metadata.initializers.result.message

     A human-readable description of the status of this operation.

###### metadata \<Object\>
**PATH:**  statefulset.spec.volumeClaimTemplates.metadata.initializers.result.metadata

KIND:     StatefulSet
VERSION:  apps/v1

RESOURCE: metadata <Object>

DESCRIPTION:
     Standard list metadata. More info:
     https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds

     ListMeta describes metadata that synthetic resources must have, including
     lists and various status objects. A resource may have only one of
     {ObjectMeta, ListMeta}.

- continue	\<string\>
**PATH:**  statefulset.spec.volumeClaimTemplates.metadata.initializers.result.metadata.continue

     continue may be set if the user set a limit on the number of items
     returned, and indicates that the server has more data available. The value
     is opaque and may be used to issue another request to the endpoint that
     served this list to retrieve the next set of available objects. Continuing
     a consistent list may not be possible if the server configuration has
     changed or more than a few minutes have passed. The resourceVersion field
     returned when using this continue value will be identical to the value in
     the first response, unless you have received this token from an error
     message.

- resourceVersion	\<string\>
**PATH:**  statefulset.spec.volumeClaimTemplates.metadata.initializers.result.metadata.resourceVersion

     String that identifies the server's internal version of this object that
     can be used by clients to determine when objects have changed. Value must
     be treated as opaque by clients and passed unmodified back to the server.
     Populated by the system. Read-only. More info:
     https://git.k8s.io/community/contributors/devel/api-conventions.md#concurrency-control-and-consistency

###### metadata	\<Object\>
**PATH:**  statefulset.spec.volumeClaimTemplates.metadata.initializers.result.metadata

     Standard list metadata. More info:
     https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds

###### reason	\<string\>
**PATH:**  statefulset.spec.volumeClaimTemplates.metadata.initializers.result.reason

     A machine-readable description of why this operation is in the "Failure"
     status. If this value is empty there is no information available. A Reason
     clarifies an HTTP status code but does not override it.

#### initializers	\<Object\>
**PATH:**  statefulset.spec.volumeClaimTemplates.metadata.initializers

     An initializer is a controller which enforces some system invariant at
     object creation time. This field is a list of initializers that have not
     yet acted on this object. If nil or empty, this object has been completely
     initialized. Otherwise, the object is considered uninitialized and is
     hidden (in list/watch and get calls) from clients that haven't explicitly
     asked to observe uninitialized objects. When an object is created, the
     system will populate this list with the current set of initializers. Only
     privileged users may set or modify this list. Once it is empty, it may not
     be modified further by any user. DEPRECATED - initializers are an alpha
     field and will be removed in v1.15.

#### labels	\<map[string]string\>
**PATH:**  statefulset.spec.volumeClaimTemplates.metadata.labels

     Map of string keys and values that can be used to organize and categorize
     (scope and select) objects. May match selectors of replication controllers
     and services. More info: http://kubernetes.io/docs/user-guide/labels

#### managedFields \<[]Object\>
**PATH:**  statefulset.spec.volumeClaimTemplates.metadata.managedFields

KIND:     StatefulSet
VERSION:  apps/v1

RESOURCE: managedFields <[]Object>

DESCRIPTION:
     ManagedFields maps workflow-id and version to the set of fields that are
     managed by that workflow. This is mostly for internal housekeeping, and
     users typically shouldn't need to set or understand this field. A workflow
     can be the user's name, a controller's name, or the name of a specific
     apply path like "ci-cd". The set of fields is always in the version that
     the workflow used when modifying the object. This field is alpha and can be
     changed or removed without notice.

     ManagedFieldsEntry is a workflow-id, a FieldSet and the group version of
     the resource that the fieldset applies to.

##### apiVersion	\<string\>
**PATH:**  statefulset.spec.volumeClaimTemplates.metadata.managedFields.apiVersion

     APIVersion defines the version of this resource that this field set applies
     to. The format is "group/version" just like the top-level APIVersion field.
     It is necessary to track the version of a field set because it cannot be
     automatically converted.

   fields	<map[string]>
     Fields identifies a set of fields.

##### manager	\<string\>
**PATH:**  statefulset.spec.volumeClaimTemplates.metadata.managedFields.manager

     Manager is an identifier of the workflow managing these fields.

##### operation	\<string\>
**PATH:**  statefulset.spec.volumeClaimTemplates.metadata.managedFields.operation

     Operation is the type of operation which lead to this ManagedFieldsEntry
     being created. The only valid values for this field are 'Apply' and
     'Update'.

#### managedFields	\<[]Object\>
**PATH:**  statefulset.spec.volumeClaimTemplates.metadata.managedFields

     ManagedFields maps workflow-id and version to the set of fields that are
     managed by that workflow. This is mostly for internal housekeeping, and
     users typically shouldn't need to set or understand this field. A workflow
     can be the user's name, a controller's name, or the name of a specific
     apply path like "ci-cd". The set of fields is always in the version that
     the workflow used when modifying the object. This field is alpha and can be
     changed or removed without notice.

#### name	\<string\>
**PATH:**  statefulset.spec.volumeClaimTemplates.metadata.name

     Name must be unique within a namespace. Is required when creating
     resources, although some resources may allow a client to request the
     generation of an appropriate name automatically. Name is primarily intended
     for creation idempotence and configuration definition. Cannot be updated.
     More info: http://kubernetes.io/docs/user-guide/identifiers#names

#### namespace	\<string\>
**PATH:**  statefulset.spec.volumeClaimTemplates.metadata.namespace

     Namespace defines the space within each name must be unique. An empty
     namespace is equivalent to the "default" namespace, but "default" is the
     canonical representation. Not all objects are required to be scoped to a
     namespace - the value of this field for those objects will be empty. Must
     be a DNS_LABEL. Cannot be updated. More info:
     http://kubernetes.io/docs/user-guide/namespaces

#### ownerReferences \<[]Object\>
**PATH:**  statefulset.spec.volumeClaimTemplates.metadata.ownerReferences

KIND:     StatefulSet
VERSION:  apps/v1

RESOURCE: ownerReferences <[]Object>

DESCRIPTION:
     List of objects depended by this object. If ALL objects in the list have
     been deleted, this object will be garbage collected. If this object is
     managed by a controller, then an entry in this list will point to this
     controller, with the controller field set to true. There cannot be more
     than one managing controller.

     OwnerReference contains enough information to let you identify an owning
     object. An owning object must be in the same namespace as the dependent, or
     be cluster-scoped, so there is no namespace field.

##### apiVersion	\<string\> -required-
**PATH:**  statefulset.spec.volumeClaimTemplates.metadata.ownerReferences.apiVersion

     API version of the referent.

##### blockOwnerDeletion	\<boolean\>
**PATH:**  statefulset.spec.volumeClaimTemplates.metadata.ownerReferences.blockOwnerDeletion

     If true, AND if the owner has the "foregroundDeletion" finalizer, then the
     owner cannot be deleted from the key-value store until this reference is
     removed. Defaults to false. To set this field, a user needs "delete"
     permission of the owner, otherwise 422 (Unprocessable Entity) will be
     returned.

##### controller	\<boolean\>
**PATH:**  statefulset.spec.volumeClaimTemplates.metadata.ownerReferences.controller

     If true, this reference points to the managing controller.

##### kind	\<string\> -required-
**PATH:**  statefulset.spec.volumeClaimTemplates.metadata.ownerReferences.kind

     Kind of the referent. More info:
     https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds

##### name	\<string\> -required-
**PATH:**  statefulset.spec.volumeClaimTemplates.metadata.ownerReferences.name

     Name of the referent. More info:
     http://kubernetes.io/docs/user-guide/identifiers#names

#### ownerReferences	\<[]Object\>
**PATH:**  statefulset.spec.volumeClaimTemplates.metadata.ownerReferences

     List of objects depended by this object. If ALL objects in the list have
     been deleted, this object will be garbage collected. If this object is
     managed by a controller, then an entry in this list will point to this
     controller, with the controller field set to true. There cannot be more
     than one managing controller.

#### resourceVersion	\<string\>
**PATH:**  statefulset.spec.volumeClaimTemplates.metadata.resourceVersion

     An opaque value that represents the internal version of this object that
     can be used by clients to determine when objects have changed. May be used
     for optimistic concurrency, change detection, and the watch operation on a
     resource or set of resources. Clients must treat these values as opaque and
     passed unmodified back to the server. They may only be valid for a
     particular resource or set of resources. Populated by the system.
     Read-only. Value must be treated as opaque by clients and . More info:
     https://git.k8s.io/community/contributors/devel/api-conventions.md#concurrency-control-and-consistency

#### selfLink	\<string\>
**PATH:**  statefulset.spec.volumeClaimTemplates.metadata.selfLink

     SelfLink is a URL representing this object. Populated by the system.
     Read-only.

### metadata	\<Object\>
**PATH:**  statefulset.spec.volumeClaimTemplates.metadata

     Standard object's metadata. More info:
     https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata

### spec \<Object\>
**PATH:**  statefulset.spec.volumeClaimTemplates.spec

KIND:     StatefulSet
VERSION:  apps/v1

RESOURCE: spec <Object>

DESCRIPTION:
     Spec defines the desired characteristics of a volume requested by a pod
     author. More info:
     https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims

     PersistentVolumeClaimSpec describes the common attributes of storage
     devices and allows a Source for provider-specific attributes

#### accessModes	\<[]string\>
**PATH:**  statefulset.spec.volumeClaimTemplates.spec.accessModes

     AccessModes contains the desired access modes the volume should have. More
     info:
     https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes-1

#### dataSource \<Object\>
**PATH:**  statefulset.spec.volumeClaimTemplates.spec.dataSource

KIND:     StatefulSet
VERSION:  apps/v1

RESOURCE: dataSource <Object>

DESCRIPTION:
     This field requires the VolumeSnapshotDataSource alpha feature gate to be
     enabled and currently VolumeSnapshot is the only supported data source. If
     the provisioner can support VolumeSnapshot data source, it will create a
     new volume and data will be restored to the volume at the same time. If the
     provisioner does not support VolumeSnapshot data source, volume will not be
     created and the failure will be reported as an event. In the future, we
     plan to support more data source types and the behavior of the provisioner
     may change.

     TypedLocalObjectReference contains enough information to let you locate the
     typed referenced object inside the same namespace.

##### apiGroup	\<string\>
**PATH:**  statefulset.spec.volumeClaimTemplates.spec.dataSource.apiGroup

     APIGroup is the group for the resource being referenced. If APIGroup is not
     specified, the specified Kind must be in the core API group. For any other
     third-party types, APIGroup is required.

##### kind	\<string\> -required-
**PATH:**  statefulset.spec.volumeClaimTemplates.spec.dataSource.kind

     Kind is the type of resource being referenced

#### dataSource	\<Object\>
**PATH:**  statefulset.spec.volumeClaimTemplates.spec.dataSource

     This field requires the VolumeSnapshotDataSource alpha feature gate to be
     enabled and currently VolumeSnapshot is the only supported data source. If
     the provisioner can support VolumeSnapshot data source, it will create a
     new volume and data will be restored to the volume at the same time. If the
     provisioner does not support VolumeSnapshot data source, volume will not be
     created and the failure will be reported as an event. In the future, we
     plan to support more data source types and the behavior of the provisioner
     may change.

#### resources \<Object\>
**PATH:**  statefulset.spec.volumeClaimTemplates.spec.resources

KIND:     StatefulSet
VERSION:  apps/v1

RESOURCE: resources <Object>

DESCRIPTION:
     Resources represents the minimum resources the volume should have. More
     info:
     https://kubernetes.io/docs/concepts/storage/persistent-volumes#resources

     ResourceRequirements describes the compute resource requirements.

##### limits	\<map[string]string\>
**PATH:**  statefulset.spec.volumeClaimTemplates.spec.resources.limits

     Limits describes the maximum amount of compute resources allowed. More
     info:
     https://kubernetes.io/docs/concepts/configuration/manage-compute-resources-container/

#### resources	\<Object\>
**PATH:**  statefulset.spec.volumeClaimTemplates.spec.resources

     Resources represents the minimum resources the volume should have. More
     info:
     https://kubernetes.io/docs/concepts/storage/persistent-volumes#resources

#### selector \<Object\>
**PATH:**  statefulset.spec.volumeClaimTemplates.spec.selector

KIND:     StatefulSet
VERSION:  apps/v1

RESOURCE: selector <Object>

DESCRIPTION:
     A label query over volumes to consider for binding.

     A label selector is a label query over a set of resources. The result of
     matchLabels and matchExpressions are ANDed. An empty label selector matches
     all objects. A null label selector matches no objects.

##### matchExpressions \<[]Object\>
**PATH:**  statefulset.spec.volumeClaimTemplates.spec.selector.matchExpressions

KIND:     StatefulSet
VERSION:  apps/v1

RESOURCE: matchExpressions <[]Object>

DESCRIPTION:
     matchExpressions is a list of label selector requirements. The requirements
     are ANDed.

     A label selector requirement is a selector that contains values, a key, and
     an operator that relates the key and values.

###### key	\<string\> -required-
**PATH:**  statefulset.spec.volumeClaimTemplates.spec.selector.matchExpressions.key

     key is the label key that the selector applies to.

###### operator	\<string\> -required-
**PATH:**  statefulset.spec.volumeClaimTemplates.spec.selector.matchExpressions.operator

     operator represents a key's relationship to a set of values. Valid
     operators are In, NotIn, Exists and DoesNotExist.

##### matchExpressions	\<[]Object\>
**PATH:**  statefulset.spec.volumeClaimTemplates.spec.selector.matchExpressions

     matchExpressions is a list of label selector requirements. The requirements
     are ANDed.

#### selector	\<Object\>
**PATH:**  statefulset.spec.volumeClaimTemplates.spec.selector

     A label query over volumes to consider for binding.

#### storageClassName	\<string\>
**PATH:**  statefulset.spec.volumeClaimTemplates.spec.storageClassName

     Name of the StorageClass required by the claim. More info:
     https://kubernetes.io/docs/concepts/storage/persistent-volumes#class-1

#### volumeMode	\<string\>
**PATH:**  statefulset.spec.volumeClaimTemplates.spec.volumeMode

     volumeMode defines what type of volume is required by the claim. Value of
     Filesystem is implied when not included in claim spec. This is a beta
     feature.

### spec	\<Object\>
**PATH:**  statefulset.spec.volumeClaimTemplates.spec

     Spec defines the desired characteristics of a volume requested by a pod
     author. More info:
     https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims

### status \<Object\>
**PATH:**  statefulset.spec.volumeClaimTemplates.status

KIND:     StatefulSet
VERSION:  apps/v1

RESOURCE: status <Object>

DESCRIPTION:
     Status represents the current information/status of a persistent volume
     claim. Read-only. More info:
     https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims

     PersistentVolumeClaimStatus is the current status of a persistent volume
     claim.

#### accessModes	\<[]string\>
**PATH:**  statefulset.spec.volumeClaimTemplates.status.accessModes

     AccessModes contains the actual access modes the volume backing the PVC
     has. More info:
     https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes-1

#### capacity	\<map[string]string\>
**PATH:**  statefulset.spec.volumeClaimTemplates.status.capacity

     Represents the actual resources of the underlying volume.

#### conditions \<[]Object\>
**PATH:**  statefulset.spec.volumeClaimTemplates.status.conditions

KIND:     StatefulSet
VERSION:  apps/v1

RESOURCE: conditions <[]Object>

DESCRIPTION:
     Current Condition of persistent volume claim. If underlying persistent
     volume is being resized then the Condition will be set to 'ResizeStarted'.

     PersistentVolumeClaimCondition contails details about state of pvc

##### lastProbeTime	\<string\>
**PATH:**  statefulset.spec.volumeClaimTemplates.status.conditions.lastProbeTime

     Last time we probed the condition.

##### lastTransitionTime	\<string\>
**PATH:**  statefulset.spec.volumeClaimTemplates.status.conditions.lastTransitionTime

     Last time the condition transitioned from one status to another.

##### message	\<string\>
**PATH:**  statefulset.spec.volumeClaimTemplates.status.conditions.message

     Human-readable message indicating details about last transition.

##### reason	\<string\>
**PATH:**  statefulset.spec.volumeClaimTemplates.status.conditions.reason

     Unique, this should be a short, machine understandable string that gives
     the reason for condition's last transition. If it reports "ResizeStarted"
     that means the underlying persistent volume is being resized.

##### status	\<string\> -required-
**PATH:**  statefulset.spec.volumeClaimTemplates.status.conditions.status


#### conditions	\<[]Object\>
**PATH:**  statefulset.spec.volumeClaimTemplates.status.conditions

     Current Condition of persistent volume claim. If underlying persistent
     volume is being resized then the Condition will be set to 'ResizeStarted'.

# spec	\<Object\>
**PATH:**  statefulset.spec

     Spec defines the desired identities of pods in this set.

# status \<Object\>
**PATH:**  statefulset.status

KIND:     StatefulSet
VERSION:  apps/v1

RESOURCE: status <Object>

DESCRIPTION:
     Status is the current status of Pods in this StatefulSet. This data may be
     out of date by some window of time.

     StatefulSetStatus represents the current state of a StatefulSet.

## collisionCount	\<integer\>
**PATH:**  statefulset.status.collisionCount

     collisionCount is the count of hash collisions for the StatefulSet. The
     StatefulSet controller uses this field as a collision avoidance mechanism
     when it needs to create the name for the newest ControllerRevision.

## conditions \<[]Object\>
**PATH:**  statefulset.status.conditions

KIND:     StatefulSet
VERSION:  apps/v1

RESOURCE: conditions <[]Object>

DESCRIPTION:
     Represents the latest available observations of a statefulset's current
     state.

     StatefulSetCondition describes the state of a statefulset at a certain
     point.

### lastTransitionTime	\<string\>
**PATH:**  statefulset.status.conditions.lastTransitionTime

     Last time the condition transitioned from one status to another.

### message	\<string\>
**PATH:**  statefulset.status.conditions.message

     A human readable message indicating details about the transition.

### reason	\<string\>
**PATH:**  statefulset.status.conditions.reason

     The reason for the condition's last transition.

### status	\<string\> -required-
**PATH:**  statefulset.status.conditions.status

     Status of the condition, one of True, False, Unknown.

## conditions	\<[]Object\>
**PATH:**  statefulset.status.conditions

     Represents the latest available observations of a statefulset's current
     state.

## currentReplicas	\<integer\>
**PATH:**  statefulset.status.currentReplicas

     currentReplicas is the number of Pods created by the StatefulSet controller
     from the StatefulSet version indicated by currentRevision.

## currentRevision	\<string\>
**PATH:**  statefulset.status.currentRevision

     currentRevision, if not empty, indicates the version of the StatefulSet
     used to generate Pods in the sequence [0,currentReplicas).

## observedGeneration	\<integer\>
**PATH:**  statefulset.status.observedGeneration

     observedGeneration is the most recent generation observed for this
     StatefulSet. It corresponds to the StatefulSet's generation, which is
     updated on mutation by the API Server.

## readyReplicas	\<integer\>
**PATH:**  statefulset.status.readyReplicas

     readyReplicas is the number of Pods created by the StatefulSet controller
     that have a Ready Condition.

## replicas	\<integer\> -required-
**PATH:**  statefulset.status.replicas

     replicas is the number of Pods created by the StatefulSet controller.

## updateRevision	\<string\>
**PATH:**  statefulset.status.updateRevision

     updateRevision, if not empty, indicates the version of the StatefulSet used
     to generate Pods in the sequence [replicas-updatedReplicas,replicas)

