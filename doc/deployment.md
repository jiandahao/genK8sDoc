# deployment Object
**PATH:**  deployment

KIND:     Deployment
VERSION:  extensions/v1beta1

DESCRIPTION:
     DEPRECATED - This group version of Deployment is deprecated by
     apps/v1beta2/Deployment. See the release notes for more information.
     Deployment enables declarative updates for Pods and ReplicaSets.

## apiVersion	\<string\>
**PATH:**  deployment.apiVersion

     APIVersion defines the versioned schema of this representation of an
     object. Servers should convert recognized schemas to the latest internal
     value, and may reject unrecognized values. More info:
     https://git.k8s.io/community/contributors/devel/api-conventions.md#resources

## kind	\<string\>
**PATH:**  deployment.kind

     Kind is a string value representing the REST resource this object
     represents. Servers may infer this from the endpoint the client submits
     requests to. Cannot be updated. In CamelCase. More info:
     https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds

## metadata \<Object\>
**PATH:**  deployment.metadata

KIND:     Deployment
VERSION:  extensions/v1beta1

RESOURCE: metadata <Object>

DESCRIPTION:
     Standard object metadata.

     ObjectMeta is metadata that all persisted resources must have, which
     includes all objects users must create.

### annotations	\<map[string]string\>
**PATH:**  deployment.metadata.annotations

     Annotations is an unstructured key value map stored with a resource that
     may be set by external tools to store and retrieve arbitrary metadata. They
     are not queryable and should be preserved when modifying objects. More
     info: http://kubernetes.io/docs/user-guide/annotations

### clusterName	\<string\>
**PATH:**  deployment.metadata.clusterName

     The name of the cluster which the object belongs to. This is used to
     distinguish resources with same name and namespace in different clusters.
     This field is not set anywhere right now and apiserver is going to ignore
     it if set in create or update request.

### creationTimestamp	\<string\>
**PATH:**  deployment.metadata.creationTimestamp

     CreationTimestamp is a timestamp representing the server time when this
     object was created. It is not guaranteed to be set in happens-before order
     across separate operations. Clients may not set this value. It is
     represented in RFC3339 form and is in UTC. Populated by the system.
     Read-only. Null for lists. More info:
     https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata

### deletionGracePeriodSeconds	\<integer\>
**PATH:**  deployment.metadata.deletionGracePeriodSeconds

     Number of seconds allowed for this object to gracefully terminate before it
     will be removed from the system. Only set when deletionTimestamp is also
     set. May only be shortened. Read-only.

### deletionTimestamp	\<string\>
**PATH:**  deployment.metadata.deletionTimestamp

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

### finalizers	\<[]string\>
**PATH:**  deployment.metadata.finalizers

     Must be empty before the object is deleted from the registry. Each entry is
     an identifier for the responsible component that will remove the entry from
     the list. If the deletionTimestamp of the object is non-nil, entries in
     this list can only be removed.

### generateName	\<string\>
**PATH:**  deployment.metadata.generateName

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

### generation	\<integer\>
**PATH:**  deployment.metadata.generation

     A sequence number representing a specific generation of the desired state.
     Populated by the system. Read-only.

### initializers	\<Object\>
**PATH:**  deployment.metadata.initializers

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

### labels	\<map[string]string\>
**PATH:**  deployment.metadata.labels

     Map of string keys and values that can be used to organize and categorize
     (scope and select) objects. May match selectors of replication controllers
     and services. More info: http://kubernetes.io/docs/user-guide/labels

### managedFields	\<[]Object\>
**PATH:**  deployment.metadata.managedFields

     ManagedFields maps workflow-id and version to the set of fields that are
     managed by that workflow. This is mostly for internal housekeeping, and
     users typically shouldn't need to set or understand this field. A workflow
     can be the user's name, a controller's name, or the name of a specific
     apply path like "ci-cd". The set of fields is always in the version that
     the workflow used when modifying the object. This field is alpha and can be
     changed or removed without notice.

### name	\<string\>
**PATH:**  deployment.metadata.name

     Name must be unique within a namespace. Is required when creating
     resources, although some resources may allow a client to request the
     generation of an appropriate name automatically. Name is primarily intended
     for creation idempotence and configuration definition. Cannot be updated.
     More info: http://kubernetes.io/docs/user-guide/identifiers#names

### namespace	\<string\>
**PATH:**  deployment.metadata.namespace

     Namespace defines the space within each name must be unique. An empty
     namespace is equivalent to the "default" namespace, but "default" is the
     canonical representation. Not all objects are required to be scoped to a
     namespace - the value of this field for those objects will be empty. Must
     be a DNS_LABEL. Cannot be updated. More info:
     http://kubernetes.io/docs/user-guide/namespaces

### ownerReferences	\<[]Object\>
**PATH:**  deployment.metadata.ownerReferences

     List of objects depended by this object. If ALL objects in the list have
     been deleted, this object will be garbage collected. If this object is
     managed by a controller, then an entry in this list will point to this
     controller, with the controller field set to true. There cannot be more
     than one managing controller.

### resourceVersion	\<string\>
**PATH:**  deployment.metadata.resourceVersion

     An opaque value that represents the internal version of this object that
     can be used by clients to determine when objects have changed. May be used
     for optimistic concurrency, change detection, and the watch operation on a
     resource or set of resources. Clients must treat these values as opaque and
     passed unmodified back to the server. They may only be valid for a
     particular resource or set of resources. Populated by the system.
     Read-only. Value must be treated as opaque by clients and . More info:
     https://git.k8s.io/community/contributors/devel/api-conventions.md#concurrency-control-and-consistency

### selfLink	\<string\>
**PATH:**  deployment.metadata.selfLink

     SelfLink is a URL representing this object. Populated by the system.
     Read-only.

## metadata	\<Object\>
**PATH:**  deployment.metadata

     Standard object metadata.

## spec \<Object\>
**PATH:**  deployment.spec

KIND:     Deployment
VERSION:  extensions/v1beta1

RESOURCE: spec <Object>

DESCRIPTION:
     Specification of the desired behavior of the Deployment.

     DeploymentSpec is the specification of the desired behavior of the
     Deployment.

### minReadySeconds	\<integer\>
**PATH:**  deployment.spec.minReadySeconds

     Minimum number of seconds for which a newly created pod should be ready
     without any of its container crashing, for it to be considered available.
     Defaults to 0 (pod will be considered available as soon as it is ready)

### paused	\<boolean\>
**PATH:**  deployment.spec.paused

     Indicates that the deployment is paused and will not be processed by the
     deployment controller.

### progressDeadlineSeconds	\<integer\>
**PATH:**  deployment.spec.progressDeadlineSeconds

     The maximum time in seconds for a deployment to make progress before it is
     considered to be failed. The deployment controller will continue to process
     failed deployments and a condition with a ProgressDeadlineExceeded reason
     will be surfaced in the deployment status. Note that progress will not be
     estimated during the time a deployment is paused. This is set to the max
     value of int32 (i.e. 2147483647) by default, which means "no deadline".

### replicas	\<integer\>
**PATH:**  deployment.spec.replicas

     Number of desired pods. This is a pointer to distinguish between explicit
     zero and not specified. Defaults to 1.

### revisionHistoryLimit	\<integer\>
**PATH:**  deployment.spec.revisionHistoryLimit

     The number of old ReplicaSets to retain to allow rollback. This is a
     pointer to distinguish between explicit zero and not specified. This is set
     to the max value of int32 (i.e. 2147483647) by default, which means
     "retaining all old RelicaSets".

### rollbackTo \<Object\>
**PATH:**  deployment.spec.rollbackTo

KIND:     Deployment
VERSION:  extensions/v1beta1

RESOURCE: rollbackTo <Object>

DESCRIPTION:
     DEPRECATED. The config this deployment is rolling back to. Will be cleared
     after rollback is done.

     DEPRECATED.

### rollbackTo	\<Object\>
**PATH:**  deployment.spec.rollbackTo

     DEPRECATED. The config this deployment is rolling back to. Will be cleared
     after rollback is done.

### selector \<Object\>
**PATH:**  deployment.spec.selector

KIND:     Deployment
VERSION:  extensions/v1beta1

RESOURCE: selector <Object>

DESCRIPTION:
     Label selector for pods. Existing ReplicaSets whose pods are selected by
     this will be the ones affected by this deployment.

     A label selector is a label query over a set of resources. The result of
     matchLabels and matchExpressions are ANDed. An empty label selector matches
     all objects. A null label selector matches no objects.

#### matchExpressions	\<[]Object\>
**PATH:**  deployment.spec.selector.matchExpressions

     matchExpressions is a list of label selector requirements. The requirements
     are ANDed.

### selector	\<Object\>
**PATH:**  deployment.spec.selector

     Label selector for pods. Existing ReplicaSets whose pods are selected by
     this will be the ones affected by this deployment.

### strategy \<Object\>
**PATH:**  deployment.spec.strategy

KIND:     Deployment
VERSION:  extensions/v1beta1

RESOURCE: strategy <Object>

DESCRIPTION:
     The deployment strategy to use to replace existing pods with new ones.

     DeploymentStrategy describes how to replace existing pods with new ones.

#### rollingUpdate	\<Object\>
**PATH:**  deployment.spec.strategy.rollingUpdate

     Rolling update config params. Present only if DeploymentStrategyType =
     RollingUpdate.

## spec	\<Object\>
**PATH:**  deployment.spec

     Specification of the desired behavior of the Deployment.

## status \<Object\>
**PATH:**  deployment.status

KIND:     Deployment
VERSION:  extensions/v1beta1

RESOURCE: status <Object>

DESCRIPTION:
     Most recently observed status of the Deployment.

     DeploymentStatus is the most recently observed status of the Deployment.

### availableReplicas	\<integer\>
**PATH:**  deployment.status.availableReplicas

     Total number of available pods (ready for at least minReadySeconds)
     targeted by this deployment.

### collisionCount	\<integer\>
**PATH:**  deployment.status.collisionCount

     Count of hash collisions for the Deployment. The Deployment controller uses
     this field as a collision avoidance mechanism when it needs to create the
     name for the newest ReplicaSet.

### conditions	\<[]Object\>
**PATH:**  deployment.status.conditions

     Represents the latest available observations of a deployment's current
     state.

### observedGeneration	\<integer\>
**PATH:**  deployment.status.observedGeneration

     The generation observed by the deployment controller.

### readyReplicas	\<integer\>
**PATH:**  deployment.status.readyReplicas

     Total number of ready pods targeted by this deployment.

### replicas	\<integer\>
**PATH:**  deployment.status.replicas

     Total number of non-terminated pods targeted by this deployment (their
     labels match the selector).

### unavailableReplicas	\<integer\>
**PATH:**  deployment.status.unavailableReplicas

     Total number of unavailable pods targeted by this deployment. This is the
     total number of pods that are still required for the deployment to have
     100% available capacity. They may either be pods that are running but not
     yet available or pods that still have not been created.

