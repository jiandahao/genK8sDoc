# ClusterRoleBinding
 ClusterRoleBinding Object
**PATH:**  ClusterRoleBinding

KIND:     ClusterRoleBinding
VERSION:  rbac.authorization.k8s.io/v1

DESCRIPTION:
     ClusterRoleBinding references a ClusterRole, but not contain it. It can
     reference a ClusterRole in the global namespace, and adds who information
     via Subject.

# apiVersion	\<string\>
**PATH:**  ClusterRoleBinding.apiVersion

     APIVersion defines the versioned schema of this representation of an
     object. Servers should convert recognized schemas to the latest internal
     value, and may reject unrecognized values. More info:
     https://git.k8s.io/community/contributors/devel/api-conventions.md#resources

# kind	\<string\>
**PATH:**  ClusterRoleBinding.kind

     Kind is a string value representing the REST resource this object
     represents. Servers may infer this from the endpoint the client submits
     requests to. Cannot be updated. In CamelCase. More info:
     https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds

# metadata \<Object\>
**PATH:**  ClusterRoleBinding.metadata

KIND:     ClusterRoleBinding
VERSION:  rbac.authorization.k8s.io/v1

RESOURCE: metadata <Object>

DESCRIPTION:
     Standard object's metadata.

     ObjectMeta is metadata that all persisted resources must have, which
     includes all objects users must create.

## annotations	\<map[string]string\>
**PATH:**  ClusterRoleBinding.metadata.annotations

     Annotations is an unstructured key value map stored with a resource that
     may be set by external tools to store and retrieve arbitrary metadata. They
     are not queryable and should be preserved when modifying objects. More
     info: http://kubernetes.io/docs/user-guide/annotations

## clusterName	\<string\>
**PATH:**  ClusterRoleBinding.metadata.clusterName

     The name of the cluster which the object belongs to. This is used to
     distinguish resources with same name and namespace in different clusters.
     This field is not set anywhere right now and apiserver is going to ignore
     it if set in create or update request.

## creationTimestamp	\<string\>
**PATH:**  ClusterRoleBinding.metadata.creationTimestamp

     CreationTimestamp is a timestamp representing the server time when this
     object was created. It is not guaranteed to be set in happens-before order
     across separate operations. Clients may not set this value. It is
     represented in RFC3339 form and is in UTC. Populated by the system.
     Read-only. Null for lists. More info:
     https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata

## deletionGracePeriodSeconds	\<integer\>
**PATH:**  ClusterRoleBinding.metadata.deletionGracePeriodSeconds

     Number of seconds allowed for this object to gracefully terminate before it
     will be removed from the system. Only set when deletionTimestamp is also
     set. May only be shortened. Read-only.

## deletionTimestamp	\<string\>
**PATH:**  ClusterRoleBinding.metadata.deletionTimestamp

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
**PATH:**  ClusterRoleBinding.metadata.finalizers

     Must be empty before the object is deleted from the registry. Each entry is
     an identifier for the responsible component that will remove the entry from
     the list. If the deletionTimestamp of the object is non-nil, entries in
     this list can only be removed.

## generateName	\<string\>
**PATH:**  ClusterRoleBinding.metadata.generateName

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
**PATH:**  ClusterRoleBinding.metadata.generation

     A sequence number representing a specific generation of the desired state.
     Populated by the system. Read-only.

## initializers \<Object\>
**PATH:**  ClusterRoleBinding.metadata.initializers

KIND:     ClusterRoleBinding
VERSION:  rbac.authorization.k8s.io/v1

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
**PATH:**  ClusterRoleBinding.metadata.initializers.pending

KIND:     ClusterRoleBinding
VERSION:  rbac.authorization.k8s.io/v1

RESOURCE: pending <[]Object>

DESCRIPTION:
     Pending is a list of initializers that must execute in order before this
     object is visible. When the last pending initializer is removed, and no
     failing result is set, the initializers struct will be set to nil and the
     object is considered as initialized and visible to all clients.

     Initializer is information about an initializer that has not yet completed.

### pending	\<[]Object\> -required-
**PATH:**  ClusterRoleBinding.metadata.initializers.pending

     Pending is a list of initializers that must execute in order before this
     object is visible. When the last pending initializer is removed, and no
     failing result is set, the initializers struct will be set to nil and the
     object is considered as initialized and visible to all clients.

### result \<Object\>
**PATH:**  ClusterRoleBinding.metadata.initializers.result

KIND:     ClusterRoleBinding
VERSION:  rbac.authorization.k8s.io/v1

RESOURCE: result <Object>

DESCRIPTION:
     If result is set with the Failure field, the object will be persisted to
     storage and then deleted, ensuring that other clients can observe the
     deletion.

     Status is a return value for calls that don't return other objects.

#### apiVersion	\<string\>
**PATH:**  ClusterRoleBinding.metadata.initializers.result.apiVersion

     APIVersion defines the versioned schema of this representation of an
     object. Servers should convert recognized schemas to the latest internal
     value, and may reject unrecognized values. More info:
     https://git.k8s.io/community/contributors/devel/api-conventions.md#resources

#### code	\<integer\>
**PATH:**  ClusterRoleBinding.metadata.initializers.result.code

     Suggested HTTP return code for this status, 0 if not set.

#### details \<Object\>
**PATH:**  ClusterRoleBinding.metadata.initializers.result.details

KIND:     ClusterRoleBinding
VERSION:  rbac.authorization.k8s.io/v1

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
**PATH:**  ClusterRoleBinding.metadata.initializers.result.details.causes

KIND:     ClusterRoleBinding
VERSION:  rbac.authorization.k8s.io/v1

RESOURCE: causes <[]Object>

DESCRIPTION:
     The Causes array includes more details associated with the StatusReason
     failure. Not all StatusReasons may provide detailed causes.

     StatusCause provides more information about an api.Status failure,
     including cases when multiple errors are encountered.

###### field	\<string\>
**PATH:**  ClusterRoleBinding.metadata.initializers.result.details.causes.field

     The field of the resource that has caused this error, as named by its JSON
     serialization. May include dot and postfix notation for nested attributes.
     Arrays are zero-indexed. Fields may appear more than once in an array of
     causes due to fields having multiple errors. Optional. Examples: "name" -
     the field "name" on the current resource "items[0].name" - the field "name"
     on the first array entry in "items"

###### message	\<string\>
**PATH:**  ClusterRoleBinding.metadata.initializers.result.details.causes.message

     A human-readable description of the cause of the error. This field may be
     presented as-is to a reader.

##### causes	\<[]Object\>
**PATH:**  ClusterRoleBinding.metadata.initializers.result.details.causes

     The Causes array includes more details associated with the StatusReason
     failure. Not all StatusReasons may provide detailed causes.

##### group	\<string\>
**PATH:**  ClusterRoleBinding.metadata.initializers.result.details.group

     The group attribute of the resource associated with the status
     StatusReason.

##### kind	\<string\>
**PATH:**  ClusterRoleBinding.metadata.initializers.result.details.kind

     The kind attribute of the resource associated with the status StatusReason.
     On some operations may differ from the requested resource Kind. More info:
     https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds

##### name	\<string\>
**PATH:**  ClusterRoleBinding.metadata.initializers.result.details.name

     The name attribute of the resource associated with the status StatusReason
     (when there is a single name which can be described).

##### retryAfterSeconds	\<integer\>
**PATH:**  ClusterRoleBinding.metadata.initializers.result.details.retryAfterSeconds

     If specified, the time in seconds before the operation should be retried.
     Some errors may indicate the client must take an alternate action - for
     those errors this field may indicate how long to wait before taking the
     alternate action.

#### details	\<Object\>
**PATH:**  ClusterRoleBinding.metadata.initializers.result.details

     Extended data associated with the reason. Each reason may define its own
     extended details. This field is optional and the data returned is not
     guaranteed to conform to any schema except that defined by the reason type.

#### kind	\<string\>
**PATH:**  ClusterRoleBinding.metadata.initializers.result.kind

     Kind is a string value representing the REST resource this object
     represents. Servers may infer this from the endpoint the client submits
     requests to. Cannot be updated. In CamelCase. More info:
     https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds

#### message	\<string\>
**PATH:**  ClusterRoleBinding.metadata.initializers.result.message

     A human-readable description of the status of this operation.

#### metadata \<Object\>
**PATH:**  ClusterRoleBinding.metadata.initializers.result.metadata

KIND:     ClusterRoleBinding
VERSION:  rbac.authorization.k8s.io/v1

RESOURCE: metadata <Object>

DESCRIPTION:
     Standard list metadata. More info:
     https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds

     ListMeta describes metadata that synthetic resources must have, including
     lists and various status objects. A resource may have only one of
     {ObjectMeta, ListMeta}.

##### continue	\<string\>
**PATH:**  ClusterRoleBinding.metadata.initializers.result.metadata.continue

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
**PATH:**  ClusterRoleBinding.metadata.initializers.result.metadata.resourceVersion

     String that identifies the server's internal version of this object that
     can be used by clients to determine when objects have changed. Value must
     be treated as opaque by clients and passed unmodified back to the server.
     Populated by the system. Read-only. More info:
     https://git.k8s.io/community/contributors/devel/api-conventions.md#concurrency-control-and-consistency

#### metadata	\<Object\>
**PATH:**  ClusterRoleBinding.metadata.initializers.result.metadata

     Standard list metadata. More info:
     https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds

#### reason	\<string\>
**PATH:**  ClusterRoleBinding.metadata.initializers.result.reason

     A machine-readable description of why this operation is in the "Failure"
     status. If this value is empty there is no information available. A Reason
     clarifies an HTTP status code but does not override it.

## initializers	\<Object\>
**PATH:**  ClusterRoleBinding.metadata.initializers

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
**PATH:**  ClusterRoleBinding.metadata.labels

     Map of string keys and values that can be used to organize and categorize
     (scope and select) objects. May match selectors of replication controllers
     and services. More info: http://kubernetes.io/docs/user-guide/labels

## managedFields \<[]Object\>
**PATH:**  ClusterRoleBinding.metadata.managedFields

KIND:     ClusterRoleBinding
VERSION:  rbac.authorization.k8s.io/v1

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
**PATH:**  ClusterRoleBinding.metadata.managedFields.apiVersion

     APIVersion defines the version of this resource that this field set applies
     to. The format is "group/version" just like the top-level APIVersion field.
     It is necessary to track the version of a field set because it cannot be
     automatically converted.

   fields	<map[string]>
     Fields identifies a set of fields.

### manager	\<string\>
**PATH:**  ClusterRoleBinding.metadata.managedFields.manager

     Manager is an identifier of the workflow managing these fields.

### operation	\<string\>
**PATH:**  ClusterRoleBinding.metadata.managedFields.operation

     Operation is the type of operation which lead to this ManagedFieldsEntry
     being created. The only valid values for this field are 'Apply' and
     'Update'.

## managedFields	\<[]Object\>
**PATH:**  ClusterRoleBinding.metadata.managedFields

     ManagedFields maps workflow-id and version to the set of fields that are
     managed by that workflow. This is mostly for internal housekeeping, and
     users typically shouldn't need to set or understand this field. A workflow
     can be the user's name, a controller's name, or the name of a specific
     apply path like "ci-cd". The set of fields is always in the version that
     the workflow used when modifying the object. This field is alpha and can be
     changed or removed without notice.

## name	\<string\>
**PATH:**  ClusterRoleBinding.metadata.name

     Name must be unique within a namespace. Is required when creating
     resources, although some resources may allow a client to request the
     generation of an appropriate name automatically. Name is primarily intended
     for creation idempotence and configuration definition. Cannot be updated.
     More info: http://kubernetes.io/docs/user-guide/identifiers#names

## namespace	\<string\>
**PATH:**  ClusterRoleBinding.metadata.namespace

     Namespace defines the space within each name must be unique. An empty
     namespace is equivalent to the "default" namespace, but "default" is the
     canonical representation. Not all objects are required to be scoped to a
     namespace - the value of this field for those objects will be empty. Must
     be a DNS_LABEL. Cannot be updated. More info:
     http://kubernetes.io/docs/user-guide/namespaces

## ownerReferences \<[]Object\>
**PATH:**  ClusterRoleBinding.metadata.ownerReferences

KIND:     ClusterRoleBinding
VERSION:  rbac.authorization.k8s.io/v1

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
**PATH:**  ClusterRoleBinding.metadata.ownerReferences.apiVersion

     API version of the referent.

### blockOwnerDeletion	\<boolean\>
**PATH:**  ClusterRoleBinding.metadata.ownerReferences.blockOwnerDeletion

     If true, AND if the owner has the "foregroundDeletion" finalizer, then the
     owner cannot be deleted from the key-value store until this reference is
     removed. Defaults to false. To set this field, a user needs "delete"
     permission of the owner, otherwise 422 (Unprocessable Entity) will be
     returned.

### controller	\<boolean\>
**PATH:**  ClusterRoleBinding.metadata.ownerReferences.controller

     If true, this reference points to the managing controller.

### kind	\<string\> -required-
**PATH:**  ClusterRoleBinding.metadata.ownerReferences.kind

     Kind of the referent. More info:
     https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds

### name	\<string\> -required-
**PATH:**  ClusterRoleBinding.metadata.ownerReferences.name

     Name of the referent. More info:
     http://kubernetes.io/docs/user-guide/identifiers#names

## ownerReferences	\<[]Object\>
**PATH:**  ClusterRoleBinding.metadata.ownerReferences

     List of objects depended by this object. If ALL objects in the list have
     been deleted, this object will be garbage collected. If this object is
     managed by a controller, then an entry in this list will point to this
     controller, with the controller field set to true. There cannot be more
     than one managing controller.

## resourceVersion	\<string\>
**PATH:**  ClusterRoleBinding.metadata.resourceVersion

     An opaque value that represents the internal version of this object that
     can be used by clients to determine when objects have changed. May be used
     for optimistic concurrency, change detection, and the watch operation on a
     resource or set of resources. Clients must treat these values as opaque and
     passed unmodified back to the server. They may only be valid for a
     particular resource or set of resources. Populated by the system.
     Read-only. Value must be treated as opaque by clients and . More info:
     https://git.k8s.io/community/contributors/devel/api-conventions.md#concurrency-control-and-consistency

## selfLink	\<string\>
**PATH:**  ClusterRoleBinding.metadata.selfLink

     SelfLink is a URL representing this object. Populated by the system.
     Read-only.

# metadata	\<Object\>
**PATH:**  ClusterRoleBinding.metadata

     Standard object's metadata.

# roleRef \<Object\>
**PATH:**  ClusterRoleBinding.roleRef

KIND:     ClusterRoleBinding
VERSION:  rbac.authorization.k8s.io/v1

RESOURCE: roleRef <Object>

DESCRIPTION:
     RoleRef can only reference a ClusterRole in the global namespace. If the
     RoleRef cannot be resolved, the Authorizer must return an error.

     RoleRef contains information that points to the role being used

## apiGroup	\<string\> -required-
**PATH:**  ClusterRoleBinding.roleRef.apiGroup

     APIGroup is the group for the resource being referenced

## kind	\<string\> -required-
**PATH:**  ClusterRoleBinding.roleRef.kind

     Kind is the type of resource being referenced

# roleRef	\<Object\> -required-
**PATH:**  ClusterRoleBinding.roleRef

     RoleRef can only reference a ClusterRole in the global namespace. If the
     RoleRef cannot be resolved, the Authorizer must return an error.

# subjects \<[]Object\>
**PATH:**  ClusterRoleBinding.subjects

KIND:     ClusterRoleBinding
VERSION:  rbac.authorization.k8s.io/v1

RESOURCE: subjects <[]Object>

DESCRIPTION:
     Subjects holds references to the objects the role applies to.

     Subject contains a reference to the object or user identities a role
     binding applies to. This can either hold a direct API object reference, or
     a value for non-objects such as user and group names.

## apiGroup	\<string\>
**PATH:**  ClusterRoleBinding.subjects.apiGroup

     APIGroup holds the API group of the referenced subject. Defaults to "" for
     ServiceAccount subjects. Defaults to "rbac.authorization.k8s.io" for User
     and Group subjects.

## kind	\<string\> -required-
**PATH:**  ClusterRoleBinding.subjects.kind

     Kind of object being referenced. Values defined by this API group are
     "User", "Group", and "ServiceAccount". If the Authorizer does not
     recognized the kind value, the Authorizer should report an error.

## name	\<string\> -required-
**PATH:**  ClusterRoleBinding.subjects.name

     Name of the object being referenced.

