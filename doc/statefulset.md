# statefulset Object
**PATH:**  statefulset

KIND:     StatefulSet
VERSION:  apps/v1

DESCRIPTION:
     StatefulSet represents a set of pods with consistent identities. Identities
     are defined as: - Network: A single stable DNS and hostname. - Storage: As
     many VolumeClaims as requested. The StatefulSet guarantees that a given
     network identity will always map to the same storage identity.

## apiVersion	\<string\>
**PATH:**  statefulset.apiVersion

     APIVersion defines the versioned schema of this representation of an
     object. Servers should convert recognized schemas to the latest internal
     value, and may reject unrecognized values. More info:
     https://git.k8s.io/community/contributors/devel/api-conventions.md#resources

## kind	\<string\>
**PATH:**  statefulset.kind

     Kind is a string value representing the REST resource this object
     represents. Servers may infer this from the endpoint the client submits
     requests to. Cannot be updated. In CamelCase. More info:
     https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds

## metadata \<Object\>
**PATH:**  statefulset.metadata

KIND:     StatefulSet
VERSION:  apps/v1

RESOURCE: metadata <Object>

DESCRIPTION:

     ObjectMeta is metadata that all persisted resources must have, which
     includes all objects users must create.

### annotations	\<map[string]string\>
**PATH:**  statefulset.metadata.annotations

     Annotations is an unstructured key value map stored with a resource that
     may be set by external tools to store and retrieve arbitrary metadata. They
     are not queryable and should be preserved when modifying objects. More
     info: http://kubernetes.io/docs/user-guide/annotations

### clusterName	\<string\>
**PATH:**  statefulset.metadata.clusterName

     The name of the cluster which the object belongs to. This is used to
     distinguish resources with same name and namespace in different clusters.
     This field is not set anywhere right now and apiserver is going to ignore
     it if set in create or update request.

### creationTimestamp	\<string\>
**PATH:**  statefulset.metadata.creationTimestamp

     CreationTimestamp is a timestamp representing the server time when this
     object was created. It is not guaranteed to be set in happens-before order
     across separate operations. Clients may not set this value. It is
     represented in RFC3339 form and is in UTC. Populated by the system.
     Read-only. Null for lists. More info:
     https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata

### deletionGracePeriodSeconds	\<integer\>
**PATH:**  statefulset.metadata.deletionGracePeriodSeconds

     Number of seconds allowed for this object to gracefully terminate before it
     will be removed from the system. Only set when deletionTimestamp is also
     set. May only be shortened. Read-only.

### deletionTimestamp	\<string\>
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

### finalizers	\<[]string\>
**PATH:**  statefulset.metadata.finalizers

     Must be empty before the object is deleted from the registry. Each entry is
     an identifier for the responsible component that will remove the entry from
     the list. If the deletionTimestamp of the object is non-nil, entries in
     this list can only be removed.

### generateName	\<string\>
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

### generation	\<integer\>
**PATH:**  statefulset.metadata.generation

     A sequence number representing a specific generation of the desired state.
     Populated by the system. Read-only.

### initializers	\<Object\>
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

### labels	\<map[string]string\>
**PATH:**  statefulset.metadata.labels

     Map of string keys and values that can be used to organize and categorize
     (scope and select) objects. May match selectors of replication controllers
     and services. More info: http://kubernetes.io/docs/user-guide/labels

### managedFields	\<[]Object\>
**PATH:**  statefulset.metadata.managedFields

     ManagedFields maps workflow-id and version to the set of fields that are
     managed by that workflow. This is mostly for internal housekeeping, and
     users typically shouldn't need to set or understand this field. A workflow
     can be the user's name, a controller's name, or the name of a specific
     apply path like "ci-cd". The set of fields is always in the version that
     the workflow used when modifying the object. This field is alpha and can be
     changed or removed without notice.

### name	\<string\>
**PATH:**  statefulset.metadata.name

     Name must be unique within a namespace. Is required when creating
     resources, although some resources may allow a client to request the
     generation of an appropriate name automatically. Name is primarily intended
     for creation idempotence and configuration definition. Cannot be updated.
     More info: http://kubernetes.io/docs/user-guide/identifiers#names

### namespace	\<string\>
**PATH:**  statefulset.metadata.namespace

     Namespace defines the space within each name must be unique. An empty
     namespace is equivalent to the "default" namespace, but "default" is the
     canonical representation. Not all objects are required to be scoped to a
     namespace - the value of this field for those objects will be empty. Must
     be a DNS_LABEL. Cannot be updated. More info:
     http://kubernetes.io/docs/user-guide/namespaces

### ownerReferences	\<[]Object\>
**PATH:**  statefulset.metadata.ownerReferences

     List of objects depended by this object. If ALL objects in the list have
     been deleted, this object will be garbage collected. If this object is
     managed by a controller, then an entry in this list will point to this
     controller, with the controller field set to true. There cannot be more
     than one managing controller.

### resourceVersion	\<string\>
**PATH:**  statefulset.metadata.resourceVersion

     An opaque value that represents the internal version of this object that
     can be used by clients to determine when objects have changed. May be used
     for optimistic concurrency, change detection, and the watch operation on a
     resource or set of resources. Clients must treat these values as opaque and
     passed unmodified back to the server. They may only be valid for a
     particular resource or set of resources. Populated by the system.
     Read-only. Value must be treated as opaque by clients and . More info:
     https://git.k8s.io/community/contributors/devel/api-conventions.md#concurrency-control-and-consistency

### selfLink	\<string\>
**PATH:**  statefulset.metadata.selfLink

     SelfLink is a URL representing this object. Populated by the system.
     Read-only.

## metadata	\<Object\>
**PATH:**  statefulset.metadata


## spec \<Object\>
**PATH:**  statefulset.spec

KIND:     StatefulSet
VERSION:  apps/v1

RESOURCE: spec <Object>

DESCRIPTION:
     Spec defines the desired identities of pods in this set.

     A StatefulSetSpec is the specification of a StatefulSet.

### podManagementPolicy	\<string\>
**PATH:**  statefulset.spec.podManagementPolicy

     podManagementPolicy controls how pods are created during initial scale up,
     when replacing pods on nodes, or when scaling down. The default policy is
     `OrderedReady`, where pods are created in increasing order (pod-0, then
     pod-1, etc) and the controller will wait until each pod is ready before
     continuing. When scaling down, the pods are removed in the opposite order.
     The alternative policy is `Parallel` which will create pods in parallel to
     match the desired scale without waiting, and on scale down will delete all
     pods at once.

### replicas	\<integer\>
**PATH:**  statefulset.spec.replicas

     replicas is the desired number of replicas of the given Template. These are
     replicas in the sense that they are instantiations of the same Template,
     but individual replicas also have a consistent identity. If unspecified,
     defaults to 1.

### revisionHistoryLimit	\<integer\>
**PATH:**  statefulset.spec.revisionHistoryLimit

     revisionHistoryLimit is the maximum number of revisions that will be
     maintained in the StatefulSet's revision history. The revision history
     consists of all revisions not represented by a currently applied
     StatefulSetSpec version. The default value is 10.

   selector	<Object> -required-
     selector is a label query over pods that should match the replica count. It
     must match the pod template's labels. More info:
     https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/#label-selectors

   serviceName	<string> -required-
     serviceName is the name of the service that governs this StatefulSet. This
     service must exist before the StatefulSet, and is responsible for the
     network identity of the set. Pods get DNS/hostnames that follow the
     pattern: pod-specific-string.serviceName.default.svc.cluster.local where
     "pod-specific-string" is managed by the StatefulSet controller.

   template	<Object> -required-
     template is the object that describes the pod that will be created if
     insufficient replicas are detected. Each pod stamped out by the StatefulSet
     will fulfill this Template, but have a unique identity from the rest of the
     StatefulSet.

### updateStrategy \<[]Object\>
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

#### rollingUpdate	\<Object\>
**PATH:**  statefulset.spec.updateStrategy.rollingUpdate

     RollingUpdate is used to communicate parameters when Type is
     RollingUpdateStatefulSetStrategyType.

### updateStrategy	\<Object\>
**PATH:**  statefulset.spec.updateStrategy

     updateStrategy indicates the StatefulSetUpdateStrategy that will be
     employed to update Pods in the StatefulSet when a revision is made to
     Template.

### volumeClaimTemplates \<[]Object\>
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

#### apiVersion	\<string\>
**PATH:**  statefulset.spec.volumeClaimTemplates.apiVersion

     APIVersion defines the versioned schema of this representation of an
     object. Servers should convert recognized schemas to the latest internal
     value, and may reject unrecognized values. More info:
     https://git.k8s.io/community/contributors/devel/api-conventions.md#resources

#### kind	\<string\>
**PATH:**  statefulset.spec.volumeClaimTemplates.kind

     Kind is a string value representing the REST resource this object
     represents. Servers may infer this from the endpoint the client submits
     requests to. Cannot be updated. In CamelCase. More info:
     https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds

#### metadata \<Object\>
**PATH:**  statefulset.spec.volumeClaimTemplates.metadata

KIND:     StatefulSet
VERSION:  apps/v1

RESOURCE: metadata <Object>

DESCRIPTION:
     Standard object's metadata. More info:
     https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata

     ObjectMeta is metadata that all persisted resources must have, which
     includes all objects users must create.

##### annotations	\<map[string]string\>
**PATH:**  statefulset.spec.volumeClaimTemplates.metadata.annotations

     Annotations is an unstructured key value map stored with a resource that
     may be set by external tools to store and retrieve arbitrary metadata. They
     are not queryable and should be preserved when modifying objects. More
     info: http://kubernetes.io/docs/user-guide/annotations

##### clusterName	\<string\>
**PATH:**  statefulset.spec.volumeClaimTemplates.metadata.clusterName

     The name of the cluster which the object belongs to. This is used to
     distinguish resources with same name and namespace in different clusters.
     This field is not set anywhere right now and apiserver is going to ignore
     it if set in create or update request.

##### creationTimestamp	\<string\>
**PATH:**  statefulset.spec.volumeClaimTemplates.metadata.creationTimestamp

     CreationTimestamp is a timestamp representing the server time when this
     object was created. It is not guaranteed to be set in happens-before order
     across separate operations. Clients may not set this value. It is
     represented in RFC3339 form and is in UTC. Populated by the system.
     Read-only. Null for lists. More info:
     https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata

##### deletionGracePeriodSeconds	\<integer\>
**PATH:**  statefulset.spec.volumeClaimTemplates.metadata.deletionGracePeriodSeconds

     Number of seconds allowed for this object to gracefully terminate before it
     will be removed from the system. Only set when deletionTimestamp is also
     set. May only be shortened. Read-only.

##### deletionTimestamp	\<string\>
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

##### finalizers	\<[]string\>
**PATH:**  statefulset.spec.volumeClaimTemplates.metadata.finalizers

     Must be empty before the object is deleted from the registry. Each entry is
     an identifier for the responsible component that will remove the entry from
     the list. If the deletionTimestamp of the object is non-nil, entries in
     this list can only be removed.

##### generateName	\<string\>
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

##### generation	\<integer\>
**PATH:**  statefulset.spec.volumeClaimTemplates.metadata.generation

     A sequence number representing a specific generation of the desired state.
     Populated by the system. Read-only.

##### initializers	\<Object\>
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

##### labels	\<map[string]string\>
**PATH:**  statefulset.spec.volumeClaimTemplates.metadata.labels

     Map of string keys and values that can be used to organize and categorize
     (scope and select) objects. May match selectors of replication controllers
     and services. More info: http://kubernetes.io/docs/user-guide/labels

##### managedFields	\<[]Object\>
**PATH:**  statefulset.spec.volumeClaimTemplates.metadata.managedFields

     ManagedFields maps workflow-id and version to the set of fields that are
     managed by that workflow. This is mostly for internal housekeeping, and
     users typically shouldn't need to set or understand this field. A workflow
     can be the user's name, a controller's name, or the name of a specific
     apply path like "ci-cd". The set of fields is always in the version that
     the workflow used when modifying the object. This field is alpha and can be
     changed or removed without notice.

##### name	\<string\>
**PATH:**  statefulset.spec.volumeClaimTemplates.metadata.name

     Name must be unique within a namespace. Is required when creating
     resources, although some resources may allow a client to request the
     generation of an appropriate name automatically. Name is primarily intended
     for creation idempotence and configuration definition. Cannot be updated.
     More info: http://kubernetes.io/docs/user-guide/identifiers#names

##### namespace	\<string\>
**PATH:**  statefulset.spec.volumeClaimTemplates.metadata.namespace

     Namespace defines the space within each name must be unique. An empty
     namespace is equivalent to the "default" namespace, but "default" is the
     canonical representation. Not all objects are required to be scoped to a
     namespace - the value of this field for those objects will be empty. Must
     be a DNS_LABEL. Cannot be updated. More info:
     http://kubernetes.io/docs/user-guide/namespaces

##### ownerReferences	\<[]Object\>
**PATH:**  statefulset.spec.volumeClaimTemplates.metadata.ownerReferences

     List of objects depended by this object. If ALL objects in the list have
     been deleted, this object will be garbage collected. If this object is
     managed by a controller, then an entry in this list will point to this
     controller, with the controller field set to true. There cannot be more
     than one managing controller.

##### resourceVersion	\<string\>
**PATH:**  statefulset.spec.volumeClaimTemplates.metadata.resourceVersion

     An opaque value that represents the internal version of this object that
     can be used by clients to determine when objects have changed. May be used
     for optimistic concurrency, change detection, and the watch operation on a
     resource or set of resources. Clients must treat these values as opaque and
     passed unmodified back to the server. They may only be valid for a
     particular resource or set of resources. Populated by the system.
     Read-only. Value must be treated as opaque by clients and . More info:
     https://git.k8s.io/community/contributors/devel/api-conventions.md#concurrency-control-and-consistency

##### selfLink	\<string\>
**PATH:**  statefulset.spec.volumeClaimTemplates.metadata.selfLink

     SelfLink is a URL representing this object. Populated by the system.
     Read-only.

#### metadata	\<Object\>
**PATH:**  statefulset.spec.volumeClaimTemplates.metadata

     Standard object's metadata. More info:
     https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata

#### spec \<Object\>
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

##### accessModes	\<[]string\>
**PATH:**  statefulset.spec.volumeClaimTemplates.spec.accessModes

     AccessModes contains the desired access modes the volume should have. More
     info:
     https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes-1

##### dataSource \<Object\>
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

##### dataSource	\<Object\>
**PATH:**  statefulset.spec.volumeClaimTemplates.spec.dataSource

     This field requires the VolumeSnapshotDataSource alpha feature gate to be
     enabled and currently VolumeSnapshot is the only supported data source. If
     the provisioner can support VolumeSnapshot data source, it will create a
     new volume and data will be restored to the volume at the same time. If the
     provisioner does not support VolumeSnapshot data source, volume will not be
     created and the failure will be reported as an event. In the future, we
     plan to support more data source types and the behavior of the provisioner
     may change.

##### resources \<Object\>
**PATH:**  statefulset.spec.volumeClaimTemplates.spec.resources

KIND:     StatefulSet
VERSION:  apps/v1

RESOURCE: resources <Object>

DESCRIPTION:
     Resources represents the minimum resources the volume should have. More
     info:
     https://kubernetes.io/docs/concepts/storage/persistent-volumes#resources

     ResourceRequirements describes the compute resource requirements.

###### limits	\<map[string]string\>
**PATH:**  statefulset.spec.volumeClaimTemplates.spec.resources.limits

     Limits describes the maximum amount of compute resources allowed. More
     info:
     https://kubernetes.io/docs/concepts/configuration/manage-compute-resources-container/

##### resources	\<Object\>
**PATH:**  statefulset.spec.volumeClaimTemplates.spec.resources

     Resources represents the minimum resources the volume should have. More
     info:
     https://kubernetes.io/docs/concepts/storage/persistent-volumes#resources

##### selector	\<Object\>
**PATH:**  statefulset.spec.volumeClaimTemplates.spec.selector

     A label query over volumes to consider for binding.

##### storageClassName	\<string\>
**PATH:**  statefulset.spec.volumeClaimTemplates.spec.storageClassName

     Name of the StorageClass required by the claim. More info:
     https://kubernetes.io/docs/concepts/storage/persistent-volumes#class-1

##### volumeMode	\<string\>
**PATH:**  statefulset.spec.volumeClaimTemplates.spec.volumeMode

     volumeMode defines what type of volume is required by the claim. Value of
     Filesystem is implied when not included in claim spec. This is a beta
     feature.

#### spec	\<Object\>
**PATH:**  statefulset.spec.volumeClaimTemplates.spec

     Spec defines the desired characteristics of a volume requested by a pod
     author. More info:
     https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims

#### status \<Object\>
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

##### accessModes	\<[]string\>
**PATH:**  statefulset.spec.volumeClaimTemplates.status.accessModes

     AccessModes contains the actual access modes the volume backing the PVC
     has. More info:
     https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes-1

##### capacity	\<map[string]string\>
**PATH:**  statefulset.spec.volumeClaimTemplates.status.capacity

     Represents the actual resources of the underlying volume.

##### conditions	\<[]Object\>
**PATH:**  statefulset.spec.volumeClaimTemplates.status.conditions

     Current Condition of persistent volume claim. If underlying persistent
     volume is being resized then the Condition will be set to 'ResizeStarted'.

## spec	\<Object\>
**PATH:**  statefulset.spec

     Spec defines the desired identities of pods in this set.

## status \<Object\>
**PATH:**  statefulset.status

KIND:     StatefulSet
VERSION:  apps/v1

RESOURCE: status <Object>

DESCRIPTION:
     Status is the current status of Pods in this StatefulSet. This data may be
     out of date by some window of time.

     StatefulSetStatus represents the current state of a StatefulSet.

### collisionCount	\<integer\>
**PATH:**  statefulset.status.collisionCount

     collisionCount is the count of hash collisions for the StatefulSet. The
     StatefulSet controller uses this field as a collision avoidance mechanism
     when it needs to create the name for the newest ControllerRevision.

### conditions	\<[]Object\>
**PATH:**  statefulset.status.conditions

     Represents the latest available observations of a statefulset's current
     state.

### currentReplicas	\<integer\>
**PATH:**  statefulset.status.currentReplicas

     currentReplicas is the number of Pods created by the StatefulSet controller
     from the StatefulSet version indicated by currentRevision.

### currentRevision	\<string\>
**PATH:**  statefulset.status.currentRevision

     currentRevision, if not empty, indicates the version of the StatefulSet
     used to generate Pods in the sequence [0,currentReplicas).

### observedGeneration	\<integer\>
**PATH:**  statefulset.status.observedGeneration

     observedGeneration is the most recent generation observed for this
     StatefulSet. It corresponds to the StatefulSet's generation, which is
     updated on mutation by the API Server.

### readyReplicas	\<integer\>
**PATH:**  statefulset.status.readyReplicas

     readyReplicas is the number of Pods created by the StatefulSet controller
     that have a Ready Condition.

   replicas	<integer> -required-
     replicas is the number of Pods created by the StatefulSet controller.

### updateRevision	\<string\>
**PATH:**  statefulset.status.updateRevision

     updateRevision, if not empty, indicates the version of the StatefulSet used
     to generate Pods in the sequence [replicas-updatedReplicas,replicas)

