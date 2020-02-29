KIND:     Pod
VERSION:  v1

DESCRIPTION:
     Pod is a collection of containers that can run on a host. This resource is
     created by clients and scheduled onto hosts.

FIELDS:
     APIVersion defines the versioned schema of this representation of an
     object. Servers should convert recognized schemas to the latest internal
     value, and may reject unrecognized values. More info:
     https://git.k8s.io/community/contributors/devel/api-conventions.md#resources

#      Kind is a string value representing the REST resource this object
     represents. Servers may infer this from the endpoint the client submits
     requests to. Cannot be updated. In CamelCase. More info:
     https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds

#      Standard object's metadata. More info:
     https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata

KIND:     Pod
VERSION:  v1

RESOURCE: metadata <Object>

DESCRIPTION:
     Standard object's metadata. More info:
     https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata

     ObjectMeta is metadata that all persisted resources must have, which
     includes all objects users must create.

FIELDS:
     Annotations is an unstructured key value map stored with a resource that
     may be set by external tools to store and retrieve arbitrary metadata. They
     are not queryable and should be preserved when modifying objects. More
     info: http://kubernetes.io/docs/user-guide/annotations

##       The name of the cluster which the object belongs to. This is used to
     distinguish resources with same name and namespace in different clusters.
     This field is not set anywhere right now and apiserver is going to ignore
     it if set in create or update request.

##       CreationTimestamp is a timestamp representing the server time when this
     object was created. It is not guaranteed to be set in happens-before order
     across separate operations. Clients may not set this value. It is
     represented in RFC3339 form and is in UTC. Populated by the system.
     Read-only. Null for lists. More info:
     https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata

##       Number of seconds allowed for this object to gracefully terminate before it
     will be removed from the system. Only set when deletionTimestamp is also
     set. May only be shortened. Read-only.

##       DeletionTimestamp is RFC 3339 date and time at which this resource will be
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

##       Must be empty before the object is deleted from the registry. Each entry is
     an identifier for the responsible component that will remove the entry from
     the list. If the deletionTimestamp of the object is non-nil, entries in
     this list can only be removed.

##       GenerateName is an optional prefix, used by the server, to generate a
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

##       A sequence number representing a specific generation of the desired state.
     Populated by the system. Read-only.

##       An initializer is a controller which enforces some system invariant at
     object creation time. This field is a list of initializers that have not
     yet acted on this object. If nil or empty, this object has been completely
     initialized. Otherwise, the object is considered uninitialized and is
     hidden (in list/watch and get calls) from clients that haven't explicitly
     asked to observe uninitialized objects. When an object is created, the
     system will populate this list with the current set of initializers. Only
     privileged users may set or modify this list. Once it is empty, it may not
     be modified further by any user. DEPRECATED - initializers are an alpha
     field and will be removed in v1.15.

##       Map of string keys and values that can be used to organize and categorize
     (scope and select) objects. May match selectors of replication controllers
     and services. More info: http://kubernetes.io/docs/user-guide/labels

##       ManagedFields maps workflow-id and version to the set of fields that are
     managed by that workflow. This is mostly for internal housekeeping, and
     users typically shouldn't need to set or understand this field. A workflow
     can be the user's name, a controller's name, or the name of a specific
     apply path like "ci-cd". The set of fields is always in the version that
     the workflow used when modifying the object. This field is alpha and can be
     changed or removed without notice.

##       Name must be unique within a namespace. Is required when creating
     resources, although some resources may allow a client to request the
     generation of an appropriate name automatically. Name is primarily intended
     for creation idempotence and configuration definition. Cannot be updated.
     More info: http://kubernetes.io/docs/user-guide/identifiers#names

##       Namespace defines the space within each name must be unique. An empty
     namespace is equivalent to the "default" namespace, but "default" is the
     canonical representation. Not all objects are required to be scoped to a
     namespace - the value of this field for those objects will be empty. Must
     be a DNS_LABEL. Cannot be updated. More info:
     http://kubernetes.io/docs/user-guide/namespaces

##       List of objects depended by this object. If ALL objects in the list have
     been deleted, this object will be garbage collected. If this object is
     managed by a controller, then an entry in this list will point to this
     controller, with the controller field set to true. There cannot be more
     than one managing controller.

##       An opaque value that represents the internal version of this object that
     can be used by clients to determine when objects have changed. May be used
     for optimistic concurrency, change detection, and the watch operation on a
     resource or set of resources. Clients must treat these values as opaque and
     passed unmodified back to the server. They may only be valid for a
     particular resource or set of resources. Populated by the system.
     Read-only. Value must be treated as opaque by clients and . More info:
     https://git.k8s.io/community/contributors/devel/api-conventions.md#concurrency-control-and-consistency

##       SelfLink is a URL representing this object. Populated by the system.
     Read-only.

#      Specification of the desired behavior of the pod. More info:
     https://git.k8s.io/community/contributors/devel/api-conventions.md#spec-and-status

KIND:     Pod
VERSION:  v1

RESOURCE: spec <Object>

DESCRIPTION:
     Specification of the desired behavior of the pod. More info:
     https://git.k8s.io/community/contributors/devel/api-conventions.md#spec-and-status

     PodSpec is a description of a pod.

FIELDS:
     Optional duration in seconds the pod may be active on the node relative to
     StartTime before the system will actively try to mark it failed and kill
     associated containers. Value must be a positive integer.

##       If specified, the pod's scheduling constraints

##       AutomountServiceAccountToken indicates whether a service account token
     should be automatically mounted.

   containers	<[]Object> -required-
     List of containers belonging to the pod. Containers cannot currently be
     added or removed. There must be at least one container in a Pod. Cannot be
     updated.

##       Specifies the DNS parameters of a pod. Parameters specified here will be
     merged to the generated DNS configuration based on DNSPolicy.

##       Set DNS policy for the pod. Defaults to "ClusterFirst". Valid values are
     'ClusterFirstWithHostNet', 'ClusterFirst', 'Default' or 'None'. DNS
     parameters given in DNSConfig will be merged with the policy selected with
     DNSPolicy. To have DNS options set along with hostNetwork, you have to
     specify DNS policy explicitly to 'ClusterFirstWithHostNet'.

##       EnableServiceLinks indicates whether information about services should be
     injected into pod's environment variables, matching the syntax of Docker
     links. Optional: Defaults to true.

##       HostAliases is an optional list of hosts and IPs that will be injected into
     the pod's hosts file if specified. This is only valid for non-hostNetwork
     pods.

##       Use the host's ipc namespace. Optional: Default to false.

##       Host networking requested for this pod. Use the host's network namespace.
     If this option is set, the ports that will be used must be specified.
     Default to false.

##       Use the host's pid namespace. Optional: Default to false.

##       Specifies the hostname of the Pod If not specified, the pod's hostname will
     be set to a system-defined value.

##       ImagePullSecrets is an optional list of references to secrets in the same
     namespace to use for pulling any of the images used by this PodSpec. If
     specified, these secrets will be passed to individual puller
     implementations for them to use. For example, in the case of docker, only
     DockerConfig type secrets are honored. More info:
     https://kubernetes.io/docs/concepts/containers/images#specifying-imagepullsecrets-on-a-pod

##       List of initialization containers belonging to the pod. Init containers are
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

##       NodeName is a request to schedule this pod onto a specific node. If it is
     non-empty, the scheduler simply schedules this pod onto that node, assuming
     that it fits resource requirements.

##       NodeSelector is a selector which must be true for the pod to fit on a node.
     Selector which must match a node's labels for the pod to be scheduled on
     that node. More info:
     https://kubernetes.io/docs/concepts/configuration/assign-pod-node/

##       The priority value. Various system components use this field to find the
     priority of the pod. When Priority Admission Controller is enabled, it
     prevents users from setting this field. The admission controller populates
     this field from PriorityClassName. The higher the value, the higher the
     priority.

##       If specified, indicates the pod's priority. "system-node-critical" and
     "system-cluster-critical" are two special keywords which indicate the
     highest priorities with the former being the highest priority. Any other
     name must be defined by creating a PriorityClass object with that name. If
     not specified, the pod priority will be default or zero if there is no
     default.

##       If specified, all readiness gates will be evaluated for pod readiness. A
     pod is ready when all its containers are ready AND all conditions specified
     in the readiness gates have status equal to "True" More info:
     https://git.k8s.io/enhancements/keps/sig-network/0007-pod-ready%2B%2B.md

##       Restart policy for all containers within the pod. One of Always, OnFailure,
     Never. Default to Always. More info:
     https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle/#restart-policy

##       RuntimeClassName refers to a RuntimeClass object in the node.k8s.io group,
     which should be used to run this pod. If no RuntimeClass resource matches
     the named class, the pod will not be run. If unset or empty, the "legacy"
     RuntimeClass will be used, which is an implicit class with an empty
     definition that uses the default runtime handler. More info:
     https://git.k8s.io/enhancements/keps/sig-node/runtime-class.md This is an
     alpha feature and may change in the future.

##       If specified, the pod will be dispatched by specified scheduler. If not
     specified, the pod will be dispatched by default scheduler.

##       SecurityContext holds pod-level security attributes and common container
     settings. Optional: Defaults to empty. See type description for default
     values of each field.

##       DeprecatedServiceAccount is a depreciated alias for ServiceAccountName.
     Deprecated: Use serviceAccountName instead.

##       ServiceAccountName is the name of the ServiceAccount to use to run this
     pod. More info:
     https://kubernetes.io/docs/tasks/configure-pod-container/configure-service-account/

##       Share a single process namespace between all of the containers in a pod.
     When this is set containers will be able to view and signal processes from
     other containers in the same pod, and the first process in each container
     will not be assigned PID 1. HostPID and ShareProcessNamespace cannot both
     be set. Optional: Default to false. This field is beta-level and may be
     disabled with the PodShareProcessNamespace feature.

##       If specified, the fully qualified Pod hostname will be
     "<hostname>.<subdomain>.<pod namespace>.svc.<cluster domain>". If not
     specified, the pod will not have a domainname at all.

##       Optional duration in seconds the pod needs to terminate gracefully. May be
     decreased in delete request. Value must be non-negative integer. The value
     zero indicates delete immediately. If this value is nil, the default grace
     period will be used instead. The grace period is the duration in seconds
     after the processes running in the pod are sent a termination signal and
     the time when the processes are forcibly halted with a kill signal. Set
     this value longer than the expected cleanup time for your process. Defaults
     to 30 seconds.

##       If specified, the pod's tolerations.

KIND:     Pod
VERSION:  v1

RESOURCE: tolerations <[]Object>

DESCRIPTION:
     If specified, the pod's tolerations.

     The pod this Toleration is attached to tolerates any taint that matches the
     triple <key,value,effect> using the matching operator <operator>.

FIELDS:
     Effect indicates the taint effect to match. Empty means match all taint
     effects. When specified, allowed values are NoSchedule, PreferNoSchedule
     and NoExecute.

## #       Key is the taint key that the toleration applies to. Empty means match all
     taint keys. If the key is empty, operator must be Exists; this combination
     means to match all values and all keys.

## #       Operator represents a key's relationship to the value. Valid operators are
     Exists and Equal. Defaults to Equal. Exists is equivalent to wildcard for
     value, so that a pod can tolerate all taints of a particular category.

## #       TolerationSeconds represents the period of time the toleration (which must
     be of effect NoExecute, otherwise this field is ignored) tolerates the
     taint. By default, it is not set, which means tolerate the taint forever
     (do not evict). Zero and negative values will be treated as 0 (evict
     immediately) by the system.

KIND:     Pod
VERSION:  v1

RESOURCE: volumes <[]Object>

DESCRIPTION:
     List of volumes that can be mounted by containers belonging to the pod.
     More info: https://kubernetes.io/docs/concepts/storage/volumes

     Volume represents a named volume in a pod that may be accessed by any
     container in the pod.

FIELDS:
     AWSElasticBlockStore represents an AWS Disk resource that is attached to a
     kubelet's host machine and then exposed to the pod. More info:
     https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore

KIND:     Pod
VERSION:  v1

RESOURCE: awsElasticBlockStore <Object>

DESCRIPTION:
     AWSElasticBlockStore represents an AWS Disk resource that is attached to a
     kubelet's host machine and then exposed to the pod. More info:
     https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore

     Represents a Persistent Disk resource in AWS. An AWS EBS disk must exist
     before mounting to a container. The disk must also be in the same AWS zone
     as the kubelet. An AWS EBS disk can only be mounted as read/write once. AWS
     EBS volumes support ownership management and SELinux relabeling.

FIELDS:
     Filesystem type of the volume that you want to mount. Tip: Ensure that the
     filesystem type is supported by the host operating system. Examples:
     "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified.
     More info:
     https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore

## # #       The partition in the volume that you want to mount. If omitted, the default
     is to mount by volume name. Examples: For volume /dev/sda1, you specify the
     partition as "1". Similarly, the volume partition for /dev/sda is "0" (or
     you can leave the property empty).

## #       AzureDisk represents an Azure Data Disk mount on the host and bind mount to
     the pod.

KIND:     Pod
VERSION:  v1

RESOURCE: azureDisk <Object>

DESCRIPTION:
     AzureDisk represents an Azure Data Disk mount on the host and bind mount to
     the pod.

     AzureDisk represents an Azure Data Disk mount on the host and bind mount to
     the pod.

FIELDS:
     Host Caching mode: None, Read Only, Read Write.

   diskName	<string> -required-
     The Name of the data disk in the blob storage

   diskURI	<string> -required-
     The URI the data disk in the blob storage

## # #       Filesystem type to mount. Must be a filesystem type supported by the host
     operating system. Ex. "ext4", "xfs", "ntfs". Implicitly inferred to be
     "ext4" if unspecified.

## # #       Expected values Shared: multiple blob disks per storage account Dedicated:
     single blob disk per storage account Managed: azure managed data disk (only
     in managed availability set). defaults to shared

## #       AzureFile represents an Azure File Service mount on the host and bind mount
     to the pod.

## #       CephFS represents a Ceph FS mount on the host that shares a pod's lifetime

KIND:     Pod
VERSION:  v1

RESOURCE: cephfs <Object>

DESCRIPTION:
     CephFS represents a Ceph FS mount on the host that shares a pod's lifetime

     Represents a Ceph Filesystem mount that lasts the lifetime of a pod Cephfs
     volumes do not support ownership management or SELinux relabeling.

FIELDS:
   monitors	<[]string> -required-
     Required: Monitors is a collection of Ceph monitors More info:
     https://releases.k8s.io/HEAD/examples/volumes/cephfs/README.md#how-to-use-it

     Optional: Used as the mounted root, rather than the full Ceph tree, default
     is /

## # #       Optional: Defaults to false (read/write). ReadOnly here will force the
     ReadOnly setting in VolumeMounts. More info:
     https://releases.k8s.io/HEAD/examples/volumes/cephfs/README.md#how-to-use-it

## # #       Optional: SecretFile is the path to key ring for User, default is
     /etc/ceph/user.secret More info:
     https://releases.k8s.io/HEAD/examples/volumes/cephfs/README.md#how-to-use-it

## # #       Optional: SecretRef is reference to the authentication secret for User,
     default is empty. More info:
     https://releases.k8s.io/HEAD/examples/volumes/cephfs/README.md#how-to-use-it

## #       Cinder represents a cinder volume attached and mounted on kubelets host
     machine More info:
     https://releases.k8s.io/HEAD/examples/mysql-cinder-pd/README.md

KIND:     Pod
VERSION:  v1

RESOURCE: cinder <Object>

DESCRIPTION:
     Cinder represents a cinder volume attached and mounted on kubelets host
     machine More info:
     https://releases.k8s.io/HEAD/examples/mysql-cinder-pd/README.md

     Represents a cinder volume resource in Openstack. A Cinder volume must
     exist before mounting to a container. The volume must also be in the same
     region as the kubelet. Cinder volumes support ownership management and
     SELinux relabeling.

FIELDS:
     Filesystem type to mount. Must be a filesystem type supported by the host
     operating system. Examples: "ext4", "xfs", "ntfs". Implicitly inferred to
     be "ext4" if unspecified. More info:
     https://releases.k8s.io/HEAD/examples/mysql-cinder-pd/README.md

## # #       Optional: Defaults to false (read/write). ReadOnly here will force the
     ReadOnly setting in VolumeMounts. More info:
     https://releases.k8s.io/HEAD/examples/mysql-cinder-pd/README.md

## #       ConfigMap represents a configMap that should populate this volume

KIND:     Pod
VERSION:  v1

RESOURCE: configMap <Object>

DESCRIPTION:
     ConfigMap represents a configMap that should populate this volume

     Adapts a ConfigMap into a volume. The contents of the target ConfigMap's
     Data field will be presented in a volume as files using the keys in the
     Data field as the file names, unless the items element is populated with
     specific mappings of keys to paths. ConfigMap volumes support ownership
     management and SELinux relabeling.

FIELDS:
     Optional: mode bits to use on created files by default. Must be a value
     between 0 and 0777. Defaults to 0644. Directories within the path are not
     affected by this setting. This might be in conflict with other options that
     affect the file mode, like fsGroup, and the result can be other mode bits
     set.

## # #       If unspecified, each key-value pair in the Data field of the referenced
     ConfigMap will be projected into the volume as a file whose name is the key
     and content is the value. If specified, the listed keys will be projected
     into the specified paths, and unlisted keys will not be present. If a key
     is specified which is not present in the ConfigMap, the volume setup will
     error unless it is marked optional. Paths must be relative and may not
     contain the '..' path or start with '..'.

## # #       Name of the referent. More info:
     https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names

## #       CSI (Container Storage Interface) represents storage that is handled by an
     external CSI driver (Alpha feature).

KIND:     Pod
VERSION:  v1

RESOURCE: csi <Object>

DESCRIPTION:
     CSI (Container Storage Interface) represents storage that is handled by an
     external CSI driver (Alpha feature).

     Represents a source location of a volume to mount, managed by an external
     CSI driver

FIELDS:
   driver	<string> -required-
     Driver is the name of the CSI driver that handles this volume. Consult with
     your admin for the correct name as registered in the cluster.

     Filesystem type to mount. Ex. "ext4", "xfs", "ntfs". If not provided, the
     empty value is passed to the associated CSI driver which will determine the
     default filesystem to apply.

## # #       NodePublishSecretRef is a reference to the secret object containing
     sensitive information to pass to the CSI driver to complete the CSI
     NodePublishVolume and NodeUnpublishVolume calls. This field is optional,
     and may be empty if no secret is required. If the secret object contains
     more than one secret, all secret references are passed.

## # #       Specifies a read-only configuration for the volume. Defaults to false
     (read/write).

## #       DownwardAPI represents downward API about the pod that should populate this
     volume

KIND:     Pod
VERSION:  v1

RESOURCE: downwardAPI <Object>

DESCRIPTION:
     DownwardAPI represents downward API about the pod that should populate this
     volume

     DownwardAPIVolumeSource represents a volume containing downward API info.
     Downward API volumes support ownership management and SELinux relabeling.

FIELDS:
     Optional: mode bits to use on created files by default. Must be a value
     between 0 and 0777. Defaults to 0644. Directories within the path are not
     affected by this setting. This might be in conflict with other options that
     affect the file mode, like fsGroup, and the result can be other mode bits
     set.

KIND:     Pod
VERSION:  v1

RESOURCE: items <[]Object>

DESCRIPTION:
     Items is a list of downward API volume file

     DownwardAPIVolumeFile represents information to create the file containing
     the pod field

FIELDS:
     Required: Selects a field of the pod: only annotations, labels, name and
     namespace are supported.

## # # #       Optional: mode bits to use on this file, must be a value between 0 and
     0777. If not specified, the volume defaultMode will be used. This might be
     in conflict with other options that affect the file mode, like fsGroup, and
     the result can be other mode bits set.

   path	<string> -required-
     Required: Path is the relative path name of the file to be created. Must
     not be absolute or contain the '..' path. Must be utf-8 encoded. The first
     item of the relative path must not start with '..'

KIND:     Pod
VERSION:  v1

RESOURCE: resourceFieldRef <Object>

DESCRIPTION:
     Selects a resource of the container: only resources limits and requests
     (limits.cpu, limits.memory, requests.cpu and requests.memory) are currently
     supported.

     ResourceFieldSelector represents container resources (cpu, memory) and
     their output format

FIELDS:
     Container name: required for volumes, optional for env vars

## #       EmptyDir represents a temporary directory that shares a pod's lifetime.
     More info: https://kubernetes.io/docs/concepts/storage/volumes#emptydir

KIND:     Pod
VERSION:  v1

RESOURCE: emptyDir <Object>

DESCRIPTION:
     EmptyDir represents a temporary directory that shares a pod's lifetime.
     More info: https://kubernetes.io/docs/concepts/storage/volumes#emptydir

     Represents an empty directory for a pod. Empty directory volumes support
     ownership management and SELinux relabeling.

FIELDS:
     What type of storage medium should back this directory. The default is ""
     which means to use the node's default medium. Must be an empty string
     (default) or Memory. More info:
     https://kubernetes.io/docs/concepts/storage/volumes#emptydir

## #       FC represents a Fibre Channel resource that is attached to a kubelet's host
     machine and then exposed to the pod.

KIND:     Pod
VERSION:  v1

RESOURCE: fc <Object>

DESCRIPTION:
     FC represents a Fibre Channel resource that is attached to a kubelet's host
     machine and then exposed to the pod.

     Represents a Fibre Channel volume. Fibre Channel volumes can only be
     mounted as read/write once. Fibre Channel volumes support ownership
     management and SELinux relabeling.

FIELDS:
     Filesystem type to mount. Must be a filesystem type supported by the host
     operating system. Ex. "ext4", "xfs", "ntfs". Implicitly inferred to be
     "ext4" if unspecified.

## # #       Optional: FC target lun number

## # #       Optional: Defaults to false (read/write). ReadOnly here will force the
     ReadOnly setting in VolumeMounts.

## # #       Optional: FC target worldwide names (WWNs)

## #       FlexVolume represents a generic volume resource that is
     provisioned/attached using an exec based plugin.

KIND:     Pod
VERSION:  v1

RESOURCE: flexVolume <Object>

DESCRIPTION:
     FlexVolume represents a generic volume resource that is
     provisioned/attached using an exec based plugin.

     FlexVolume represents a generic volume resource that is
     provisioned/attached using an exec based plugin.

FIELDS:
   driver	<string> -required-
     Driver is the name of the driver to use for this volume.

     Filesystem type to mount. Must be a filesystem type supported by the host
     operating system. Ex. "ext4", "xfs", "ntfs". The default filesystem depends
     on FlexVolume script.

## # #       Optional: Extra command options if any.

## # #       Optional: Defaults to false (read/write). ReadOnly here will force the
     ReadOnly setting in VolumeMounts.

## #       Flocker represents a Flocker volume attached to a kubelet's host machine.
     This depends on the Flocker control service being running

KIND:     Pod
VERSION:  v1

RESOURCE: flocker <Object>

DESCRIPTION:
     Flocker represents a Flocker volume attached to a kubelet's host machine.
     This depends on the Flocker control service being running

     Represents a Flocker volume mounted by the Flocker agent. One and only one
     of datasetName and datasetUUID should be set. Flocker volumes do not
     support ownership management or SELinux relabeling.

FIELDS:
     Name of the dataset stored as metadata -> name on the dataset for Flocker
     should be considered as deprecated

## #       GCEPersistentDisk represents a GCE Disk resource that is attached to a
     kubelet's host machine and then exposed to the pod. More info:
     https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk

KIND:     Pod
VERSION:  v1

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

FIELDS:
     Filesystem type of the volume that you want to mount. Tip: Ensure that the
     filesystem type is supported by the host operating system. Examples:
     "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified.
     More info:
     https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk

## # #       The partition in the volume that you want to mount. If omitted, the default
     is to mount by volume name. Examples: For volume /dev/sda1, you specify the
     partition as "1". Similarly, the volume partition for /dev/sda is "0" (or
     you can leave the property empty). More info:
     https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk

   pdName	<string> -required-
     Unique name of the PD resource in GCE. Used to identify the disk in GCE.
     More info:
     https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk

## #       GitRepo represents a git repository at a particular revision. DEPRECATED:
     GitRepo is deprecated. To provision a container with a git repo, mount an
     EmptyDir into an InitContainer that clones the repo using git, then mount
     the EmptyDir into the Pod's container.

KIND:     Pod
VERSION:  v1

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

FIELDS:
     Target directory name. Must not contain or start with '..'. If '.' is
     supplied, the volume directory will be the git repository. Otherwise, if
     specified, the volume will contain the git repository in the subdirectory
     with the given name.

   repository	<string> -required-
     Repository URL

## #       Glusterfs represents a Glusterfs mount on the host that shares a pod's
     lifetime. More info:
     https://releases.k8s.io/HEAD/examples/volumes/glusterfs/README.md

## #       HostPath represents a pre-existing file or directory on the host machine
     that is directly exposed to the container. This is generally used for
     system agents or other privileged things that are allowed to see the host
     machine. Most containers will NOT need this. More info:
     https://kubernetes.io/docs/concepts/storage/volumes#hostpath

## #       ISCSI represents an ISCSI Disk resource that is attached to a kubelet's
     host machine and then exposed to the pod. More info:
     https://releases.k8s.io/HEAD/examples/volumes/iscsi/README.md

   name	<string> -required-
     Volume's name. Must be a DNS_LABEL and unique within the pod. More info:
     https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names

KIND:     Pod
VERSION:  v1

RESOURCE: iscsi <Object>

DESCRIPTION:
     ISCSI represents an ISCSI Disk resource that is attached to a kubelet's
     host machine and then exposed to the pod. More info:
     https://releases.k8s.io/HEAD/examples/volumes/iscsi/README.md

     Represents an ISCSI disk. ISCSI volumes can only be mounted as read/write
     once. ISCSI volumes support ownership management and SELinux relabeling.

FIELDS:
     whether support iSCSI Discovery CHAP authentication

## # #       whether support iSCSI Session CHAP authentication

## # #       Filesystem type of the volume that you want to mount. Tip: Ensure that the
     filesystem type is supported by the host operating system. Examples:
     "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified.
     More info: https://kubernetes.io/docs/concepts/storage/volumes#iscsi

## # #       Custom iSCSI Initiator Name. If initiatorName is specified with
     iscsiInterface simultaneously, new iSCSI interface <target portal>:<volume
     name> will be created for the connection.

   iqn	<string> -required-
     Target iSCSI Qualified Name.

## # #       iSCSI Interface Name that uses an iSCSI transport. Defaults to 'default'
     (tcp).

   lun	<integer> -required-
     iSCSI Target Lun number.

## # #       iSCSI Target Portal List. The portal is either an IP or ip_addr:port if the
     port is other than default (typically TCP ports 860 and 3260).

## # #       ReadOnly here will force the ReadOnly setting in VolumeMounts. Defaults to
     false.

## #       NFS represents an NFS mount on the host that shares a pod's lifetime More
     info: https://kubernetes.io/docs/concepts/storage/volumes#nfs

## #       PersistentVolumeClaimVolumeSource represents a reference to a
     PersistentVolumeClaim in the same namespace. More info:
     https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims

## #       PhotonPersistentDisk represents a PhotonController persistent disk attached
     and mounted on kubelets host machine

## #       PortworxVolume represents a portworx volume attached and mounted on
     kubelets host machine

KIND:     Pod
VERSION:  v1

RESOURCE: portworxVolume <Object>

DESCRIPTION:
     PortworxVolume represents a portworx volume attached and mounted on
     kubelets host machine

     PortworxVolumeSource represents a Portworx volume resource.

FIELDS:
     FSType represents the filesystem type to mount Must be a filesystem type
     supported by the host operating system. Ex. "ext4", "xfs". Implicitly
     inferred to be "ext4" if unspecified.

## #       Items for all in one resources secrets, configmaps, and downward API

## #       Quobyte represents a Quobyte mount on the host that shares a pod's lifetime

KIND:     Pod
VERSION:  v1

RESOURCE: quobyte <Object>

DESCRIPTION:
     Quobyte represents a Quobyte mount on the host that shares a pod's lifetime

     Represents a Quobyte mount that lasts the lifetime of a pod. Quobyte
     volumes do not support ownership management or SELinux relabeling.

FIELDS:
     Group to map volume access to Default is no group

## # #       ReadOnly here will force the Quobyte volume to be mounted with read-only
     permissions. Defaults to false.

   registry	<string> -required-
     Registry represents a single or multiple Quobyte Registry services
     specified as a string as host:port pair (multiple entries are separated
     with commas) which acts as the central registry for volumes

## # #       Tenant owning the given Quobyte volume in the Backend Used with dynamically
     provisioned Quobyte volumes, value is set by the plugin

## #       RBD represents a Rados Block Device mount on the host that shares a pod's
     lifetime. More info:
     https://releases.k8s.io/HEAD/examples/volumes/rbd/README.md

KIND:     Pod
VERSION:  v1

RESOURCE: rbd <Object>

DESCRIPTION:
     RBD represents a Rados Block Device mount on the host that shares a pod's
     lifetime. More info:
     https://releases.k8s.io/HEAD/examples/volumes/rbd/README.md

     Represents a Rados Block Device mount that lasts the lifetime of a pod. RBD
     volumes support ownership management and SELinux relabeling.

FIELDS:
     Filesystem type of the volume that you want to mount. Tip: Ensure that the
     filesystem type is supported by the host operating system. Examples:
     "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified.
     More info: https://kubernetes.io/docs/concepts/storage/volumes#rbd

   image	<string> -required-
     The rados image name. More info:
     https://releases.k8s.io/HEAD/examples/volumes/rbd/README.md#how-to-use-it

## # #       Keyring is the path to key ring for RBDUser. Default is /etc/ceph/keyring.
     More info:
     https://releases.k8s.io/HEAD/examples/volumes/rbd/README.md#how-to-use-it

   monitors	<[]string> -required-
     A collection of Ceph monitors. More info:
     https://releases.k8s.io/HEAD/examples/volumes/rbd/README.md#how-to-use-it

## # #       The rados pool name. Default is rbd. More info:
     https://releases.k8s.io/HEAD/examples/volumes/rbd/README.md#how-to-use-it

## # #       ReadOnly here will force the ReadOnly setting in VolumeMounts. Defaults to
     false. More info:
     https://releases.k8s.io/HEAD/examples/volumes/rbd/README.md#how-to-use-it

## # #       SecretRef is name of the authentication secret for RBDUser. If provided
     overrides keyring. Default is nil. More info:
     https://releases.k8s.io/HEAD/examples/volumes/rbd/README.md#how-to-use-it

## #       ScaleIO represents a ScaleIO persistent volume attached and mounted on
     Kubernetes nodes.

KIND:     Pod
VERSION:  v1

RESOURCE: scaleIO <Object>

DESCRIPTION:
     ScaleIO represents a ScaleIO persistent volume attached and mounted on
     Kubernetes nodes.

     ScaleIOVolumeSource represents a persistent ScaleIO volume

FIELDS:
     Filesystem type to mount. Must be a filesystem type supported by the host
     operating system. Ex. "ext4", "xfs", "ntfs". Default is "xfs".

   gateway	<string> -required-
     The host address of the ScaleIO API Gateway.

## # #       The name of the ScaleIO Protection Domain for the configured storage.

## # #       Defaults to false (read/write). ReadOnly here will force the ReadOnly
     setting in VolumeMounts.

   secretRef	<Object> -required-
     SecretRef references to the secret for ScaleIO user and other sensitive
     information. If this is not provided, Login operation will fail.

## # #       Flag to enable/disable SSL communication with Gateway, default false

## # #       Indicates whether the storage for a volume should be ThickProvisioned or
     ThinProvisioned. Default is ThinProvisioned.

## # #       The ScaleIO Storage Pool associated with the protection domain.

   system	<string> -required-
     The name of the storage system as configured in ScaleIO.

## #       Secret represents a secret that should populate this volume. More info:
     https://kubernetes.io/docs/concepts/storage/volumes#secret

KIND:     Pod
VERSION:  v1

RESOURCE: secret <Object>

DESCRIPTION:
     Secret represents a secret that should populate this volume. More info:
     https://kubernetes.io/docs/concepts/storage/volumes#secret

     Adapts a Secret into a volume. The contents of the target Secret's Data
     field will be presented in a volume as files using the keys in the Data
     field as the file names. Secret volumes support ownership management and
     SELinux relabeling.

FIELDS:
     Optional: mode bits to use on created files by default. Must be a value
     between 0 and 0777. Defaults to 0644. Directories within the path are not
     affected by this setting. This might be in conflict with other options that
     affect the file mode, like fsGroup, and the result can be other mode bits
     set.

## # #       If unspecified, each key-value pair in the Data field of the referenced
     Secret will be projected into the volume as a file whose name is the key
     and content is the value. If specified, the listed keys will be projected
     into the specified paths, and unlisted keys will not be present. If a key
     is specified which is not present in the Secret, the volume setup will
     error unless it is marked optional. Paths must be relative and may not
     contain the '..' path or start with '..'.

## # #       Specify whether the Secret or it's keys must be defined

## #       StorageOS represents a StorageOS volume attached and mounted on Kubernetes
     nodes.

KIND:     Pod
VERSION:  v1

RESOURCE: storageos <Object>

DESCRIPTION:
     StorageOS represents a StorageOS volume attached and mounted on Kubernetes
     nodes.

     Represents a StorageOS persistent volume resource.

FIELDS:
     Filesystem type to mount. Must be a filesystem type supported by the host
     operating system. Ex. "ext4", "xfs", "ntfs". Implicitly inferred to be
     "ext4" if unspecified.

## # #       Defaults to false (read/write). ReadOnly here will force the ReadOnly
     setting in VolumeMounts.

## # #       SecretRef specifies the secret to use for obtaining the StorageOS API
     credentials. If not specified, default values will be attempted.

## # #       VolumeName is the human-readable name of the StorageOS volume. Volume names
     are only unique within a namespace.

KIND:     Pod
VERSION:  v1

RESOURCE: vsphereVolume <Object>

DESCRIPTION:
     VsphereVolume represents a vSphere volume attached and mounted on kubelets
     host machine

     Represents a vSphere volume resource.

FIELDS:
     Filesystem type to mount. Must be a filesystem type supported by the host
     operating system. Ex. "ext4", "xfs", "ntfs". Implicitly inferred to be
     "ext4" if unspecified.

## # #       Storage Policy Based Management (SPBM) profile ID associated with the
     StoragePolicyName.

KIND:     Pod
VERSION:  v1

RESOURCE: status <Object>

DESCRIPTION:
     Most recently observed status of the pod. This data may not be up to date.
     Populated by the system. Read-only. More info:
     https://git.k8s.io/community/contributors/devel/api-conventions.md#spec-and-status

     PodStatus represents information about the status of a pod. Status may
     trail the actual state of a system, especially if the node that hosts the
     pod cannot contact the control plane.

FIELDS:
     Current service state of pod. More info:
     https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#pod-conditions

KIND:     Pod
VERSION:  v1

RESOURCE: conditions <[]Object>

DESCRIPTION:
     Current service state of pod. More info:
     https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#pod-conditions

     PodCondition contains details for the current condition of this pod.

FIELDS:
     Last time we probed the condition.

## #       Last time the condition transitioned from one status to another.

## #       Human-readable message indicating details about last transition.

##       The list has one entry per container in the manifest. Each entry is
     currently the output of `docker inspect`. More info:
     https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#pod-and-container-status

##       IP address of the host to which the pod is assigned. Empty if not yet
     scheduled.

##       The list has one entry per init container in the manifest. The most recent
     successful init container will have ready = true, the most recently started
     container will have startTime set. More info:
     https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#pod-and-container-status

##       A human readable message indicating details about why the pod is in this
     condition.

##       nominatedNodeName is set only when this pod preempts other pods on the
     node, but it cannot be scheduled right away as preemption victims receive
     their graceful termination periods. This field does not guarantee that the
     pod will be scheduled on this node. Scheduler may decide to place the pod
     elsewhere if other nodes become available sooner. Scheduler may also decide
     to give the resources on this node to a higher priority pod that is created
     after preemption. As a result, this field may be different than
     PodSpec.nodeName when the pod is scheduled.

##       The phase of a Pod is a simple, high-level summary of where the Pod is in
     its lifecycle. The conditions array, the reason and message fields, and the
     individual container status arrays contain more detail about the pod's
     status. There are five possible phase values: Pending: The pod has been
     accepted by the Kubernetes system, but one or more of the container images
     has not been created. This includes time before being scheduled as well as
     time spent downloading images over the network, which could take a while.
     Running: The pod has been bound to a node, and all of the containers have
     been created. At least one container is still running, or is in the process
     of starting or restarting. Succeeded: All containers in the pod have
     terminated in success, and will not be restarted. Failed: All containers in
     the pod have terminated, and at least one container has terminated in
     failure. The container either exited with non-zero status or was terminated
     by the system. Unknown: For some reason the state of the pod could not be
     obtained, typically due to an error in communicating with the host of the
     pod. More info:
     https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#pod-phase

##       IP address allocated to the pod. Routable at least within the cluster.
     Empty if not yet allocated.

##       The Quality of Service (QOS) classification assigned to the pod based on
     resource requirements See PodQOSClass type for available QOS classes More
     info:
     https://git.k8s.io/community/contributors/design-proposals/node/resource-qos.md

##       A brief CamelCase message indicating details about why the pod is in this
     state. e.g. 'Evicted'

