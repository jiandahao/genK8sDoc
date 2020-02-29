# Service Object
**PATH:**  Service

KIND:     Service
VERSION:  v1

DESCRIPTION:
     Service is a named abstraction of software service (for example, mysql)
     consisting of local port (for example 3306) that the proxy listens on, and
     the selector that determines which pods will answer requests sent through
     the proxy.

## apiVersion	\<string\>
**PATH:**  Service.apiVersion

     APIVersion defines the versioned schema of this representation of an
     object. Servers should convert recognized schemas to the latest internal
     value, and may reject unrecognized values. More info:
     https://git.k8s.io/community/contributors/devel/api-conventions.md#resources

## kind	\<string\>
**PATH:**  Service.kind

     Kind is a string value representing the REST resource this object
     represents. Servers may infer this from the endpoint the client submits
     requests to. Cannot be updated. In CamelCase. More info:
     https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds

## metadata \<Object\>
**PATH:**  Service.metadata

KIND:     Service
VERSION:  v1

RESOURCE: metadata <Object>

DESCRIPTION:
     Standard object's metadata. More info:
     https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata

     ObjectMeta is metadata that all persisted resources must have, which
     includes all objects users must create.

### annotations	\<map[string]string\>
**PATH:**  Service.metadata.annotations

     Annotations is an unstructured key value map stored with a resource that
     may be set by external tools to store and retrieve arbitrary metadata. They
     are not queryable and should be preserved when modifying objects. More
     info: http://kubernetes.io/docs/user-guide/annotations

### clusterName	\<string\>
**PATH:**  Service.metadata.clusterName

     The name of the cluster which the object belongs to. This is used to
     distinguish resources with same name and namespace in different clusters.
     This field is not set anywhere right now and apiserver is going to ignore
     it if set in create or update request.

### creationTimestamp	\<string\>
**PATH:**  Service.metadata.creationTimestamp

     CreationTimestamp is a timestamp representing the server time when this
     object was created. It is not guaranteed to be set in happens-before order
     across separate operations. Clients may not set this value. It is
     represented in RFC3339 form and is in UTC. Populated by the system.
     Read-only. Null for lists. More info:
     https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata

### deletionGracePeriodSeconds	\<integer\>
**PATH:**  Service.metadata.deletionGracePeriodSeconds

     Number of seconds allowed for this object to gracefully terminate before it
     will be removed from the system. Only set when deletionTimestamp is also
     set. May only be shortened. Read-only.

### deletionTimestamp	\<string\>
**PATH:**  Service.metadata.deletionTimestamp

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
**PATH:**  Service.metadata.finalizers

     Must be empty before the object is deleted from the registry. Each entry is
     an identifier for the responsible component that will remove the entry from
     the list. If the deletionTimestamp of the object is non-nil, entries in
     this list can only be removed.

### generateName	\<string\>
**PATH:**  Service.metadata.generateName

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
**PATH:**  Service.metadata.generation

     A sequence number representing a specific generation of the desired state.
     Populated by the system. Read-only.

### initializers	\<Object\>
**PATH:**  Service.metadata.initializers

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
**PATH:**  Service.metadata.labels

     Map of string keys and values that can be used to organize and categorize
     (scope and select) objects. May match selectors of replication controllers
     and services. More info: http://kubernetes.io/docs/user-guide/labels

### managedFields	\<[]Object\>
**PATH:**  Service.metadata.managedFields

     ManagedFields maps workflow-id and version to the set of fields that are
     managed by that workflow. This is mostly for internal housekeeping, and
     users typically shouldn't need to set or understand this field. A workflow
     can be the user's name, a controller's name, or the name of a specific
     apply path like "ci-cd". The set of fields is always in the version that
     the workflow used when modifying the object. This field is alpha and can be
     changed or removed without notice.

### name	\<string\>
**PATH:**  Service.metadata.name

     Name must be unique within a namespace. Is required when creating
     resources, although some resources may allow a client to request the
     generation of an appropriate name automatically. Name is primarily intended
     for creation idempotence and configuration definition. Cannot be updated.
     More info: http://kubernetes.io/docs/user-guide/identifiers#names

### namespace	\<string\>
**PATH:**  Service.metadata.namespace

     Namespace defines the space within each name must be unique. An empty
     namespace is equivalent to the "default" namespace, but "default" is the
     canonical representation. Not all objects are required to be scoped to a
     namespace - the value of this field for those objects will be empty. Must
     be a DNS_LABEL. Cannot be updated. More info:
     http://kubernetes.io/docs/user-guide/namespaces

### ownerReferences	\<[]Object\>
**PATH:**  Service.metadata.ownerReferences

     List of objects depended by this object. If ALL objects in the list have
     been deleted, this object will be garbage collected. If this object is
     managed by a controller, then an entry in this list will point to this
     controller, with the controller field set to true. There cannot be more
     than one managing controller.

### resourceVersion	\<string\>
**PATH:**  Service.metadata.resourceVersion

     An opaque value that represents the internal version of this object that
     can be used by clients to determine when objects have changed. May be used
     for optimistic concurrency, change detection, and the watch operation on a
     resource or set of resources. Clients must treat these values as opaque and
     passed unmodified back to the server. They may only be valid for a
     particular resource or set of resources. Populated by the system.
     Read-only. Value must be treated as opaque by clients and . More info:
     https://git.k8s.io/community/contributors/devel/api-conventions.md#concurrency-control-and-consistency

### selfLink	\<string\>
**PATH:**  Service.metadata.selfLink

     SelfLink is a URL representing this object. Populated by the system.
     Read-only.

## metadata	\<Object\>
**PATH:**  Service.metadata

     Standard object's metadata. More info:
     https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata

## spec \<Object\>
**PATH:**  Service.spec

KIND:     Service
VERSION:  v1

RESOURCE: spec <Object>

DESCRIPTION:
     Spec defines the behavior of a service.
     https://git.k8s.io/community/contributors/devel/api-conventions.md#spec-and-status

     ServiceSpec describes the attributes that a user creates on a service.

### clusterIP	\<string\>
**PATH:**  Service.spec.clusterIP

     clusterIP is the IP address of the service and is usually assigned randomly
     by the master. If an address is specified manually and is not in use by
     others, it will be allocated to the service; otherwise, creation of the
     service will fail. This field can not be changed through updates. Valid
     values are "None", empty string (""), or a valid IP address. "None" can be
     specified for headless services when proxying is not required. Only applies
     to types ClusterIP, NodePort, and LoadBalancer. Ignored if type is
     ExternalName. More info:
     https://kubernetes.io/docs/concepts/services-networking/service/#virtual-ips-and-service-proxies

### externalIPs	\<[]string\>
**PATH:**  Service.spec.externalIPs

     externalIPs is a list of IP addresses for which nodes in the cluster will
     also accept traffic for this service. These IPs are not managed by
     Kubernetes. The user is responsible for ensuring that traffic arrives at a
     node with this IP. A common example is external load-balancers that are not
     part of the Kubernetes system.

### externalName	\<string\>
**PATH:**  Service.spec.externalName

     externalName is the external reference that kubedns or equivalent will
     return as a CNAME record for this service. No proxying will be involved.
     Must be a valid RFC-1123 hostname (https://tools.ietf.org/html/rfc1123) and
     requires Type to be ExternalName.

### externalTrafficPolicy	\<string\>
**PATH:**  Service.spec.externalTrafficPolicy

     externalTrafficPolicy denotes if this Service desires to route external
     traffic to node-local or cluster-wide endpoints. "Local" preserves the
     client source IP and avoids a second hop for LoadBalancer and Nodeport type
     services, but risks potentially imbalanced traffic spreading. "Cluster"
     obscures the client source IP and may cause a second hop to another node,
     but should have good overall load-spreading.

### healthCheckNodePort	\<integer\>
**PATH:**  Service.spec.healthCheckNodePort

     healthCheckNodePort specifies the healthcheck nodePort for the service. If
     not specified, HealthCheckNodePort is created by the service api backend
     with the allocated nodePort. Will use user-specified nodePort value if
     specified by the client. Only effects when Type is set to LoadBalancer and
     ExternalTrafficPolicy is set to Local.

### loadBalancerIP	\<string\>
**PATH:**  Service.spec.loadBalancerIP

     Only applies to Service Type: LoadBalancer LoadBalancer will get created
     with the IP specified in this field. This feature depends on whether the
     underlying cloud-provider supports specifying the loadBalancerIP when a
     load balancer is created. This field will be ignored if the cloud-provider
     does not support the feature.

### loadBalancerSourceRanges	\<[]string\>
**PATH:**  Service.spec.loadBalancerSourceRanges

     If specified and supported by the platform, this will restrict traffic
     through the cloud-provider load-balancer will be restricted to the
     specified client IPs. This field will be ignored if the cloud-provider does
     not support the feature." More info:
     https://kubernetes.io/docs/tasks/access-application-cluster/configure-cloud-provider-firewall/

### ports	\<[]Object\>
**PATH:**  Service.spec.ports

     The list of ports that are exposed by this service. More info:
     https://kubernetes.io/docs/concepts/services-networking/service/#virtual-ips-and-service-proxies

### publishNotReadyAddresses	\<boolean\>
**PATH:**  Service.spec.publishNotReadyAddresses

     publishNotReadyAddresses, when set to true, indicates that DNS
     implementations must publish the notReadyAddresses of subsets for the
     Endpoints associated with the Service. The default value is false. The
     primary use case for setting this field is to use a StatefulSet's Headless
     Service to propagate SRV records for its Pods without respect to their
     readiness for purpose of peer discovery.

### selector	\<map[string]string\>
**PATH:**  Service.spec.selector

     Route service traffic to pods with label keys and values matching this
     selector. If empty or not present, the service is assumed to have an
     external process managing its endpoints, which Kubernetes will not modify.
     Only applies to types ClusterIP, NodePort, and LoadBalancer. Ignored if
     type is ExternalName. More info:
     https://kubernetes.io/docs/concepts/services-networking/service/

### sessionAffinity	\<string\>
**PATH:**  Service.spec.sessionAffinity

     Supports "ClientIP" and "None". Used to maintain session affinity. Enable
     client IP based session affinity. Must be ClientIP or None. Defaults to
     None. More info:
     https://kubernetes.io/docs/concepts/services-networking/service/#virtual-ips-and-service-proxies

### sessionAffinityConfig	\<Object\>
**PATH:**  Service.spec.sessionAffinityConfig

     sessionAffinityConfig contains the configurations of session affinity.

## spec	\<Object\>
**PATH:**  Service.spec

     Spec defines the behavior of a service.
     https://git.k8s.io/community/contributors/devel/api-conventions.md#spec-and-status

## status \<Object\>
**PATH:**  Service.status

KIND:     Service
VERSION:  v1

RESOURCE: status <Object>

DESCRIPTION:
     Most recently observed status of the service. Populated by the system.
     Read-only. More info:
     https://git.k8s.io/community/contributors/devel/api-conventions.md#spec-and-status

     ServiceStatus represents the current status of a service.

### loadBalancer \<Object\>
**PATH:**  Service.status.loadBalancer

KIND:     Service
VERSION:  v1

RESOURCE: loadBalancer <Object>

DESCRIPTION:
     LoadBalancer contains the current status of the load-balancer, if one is
     present.

     LoadBalancerStatus represents the status of a load-balancer.

#### ingress \<[]Object\>
**PATH:**  Service.status.loadBalancer.ingress

KIND:     Service
VERSION:  v1

RESOURCE: ingress <[]Object>

DESCRIPTION:
     Ingress is a list containing ingress points for the load-balancer. Traffic
     intended for the service should be sent to these ingress points.

     LoadBalancerIngress represents the status of a load-balancer ingress point:
     traffic intended for the service should be sent to an ingress point.

##### hostname	\<string\>
**PATH:**  Service.status.loadBalancer.ingress.hostname

     Hostname is set for load-balancer ingress points that are DNS based
     (typically AWS load-balancers)

