// Code generated by protoc-gen-go. DO NOT EDIT.
// source: routing/v1alpha2/destination_rule.proto

/*
Package v1alpha2 is a generated protocol buffer package.

It is generated from these files:
	routing/v1alpha2/destination_rule.proto
	routing/v1alpha2/external_service.proto
	routing/v1alpha2/gateway.proto
	routing/v1alpha2/virtual_service.proto

It has these top-level messages:
	DestinationRule
	TrafficPolicy
	Subset
	LoadBalancerSettings
	ConnectionPoolSettings
	OutlierDetection
	TLSSettings
	ExternalService
	Gateway
	Server
	Port
	VirtualService
	Destination
	HTTPRoute
	TCPRoute
	HTTPMatchRequest
	DestinationWeight
	L4MatchAttributes
	HTTPRedirect
	HTTPRewrite
	StringMatch
	HTTPRetry
	CorsPolicy
	HTTPFaultInjection
	PortSelector
*/
package v1alpha2

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/duration"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Standard load balancing algorithms that require no tuning.
type LoadBalancerSettings_SimpleLB int32

const (
	// Round Robin policy. Default
	LoadBalancerSettings_ROUND_ROBIN LoadBalancerSettings_SimpleLB = 0
	// The least request load balancer uses an O(1) algorithm which selects
	// two random healthy hosts and picks the host which has fewer active
	// requests.
	LoadBalancerSettings_LEAST_CONN LoadBalancerSettings_SimpleLB = 1
	// The random load balancer selects a random healthy host. The random
	// load balancer generally performs better than round robin if no health
	// checking policy is configured.
	LoadBalancerSettings_RANDOM LoadBalancerSettings_SimpleLB = 2
	// This option will forward the connection to the original IP address
	// requested by the caller without doing any form of load
	// balancing. This option must be used with care. It is meant for
	// advanced use cases. Refer to Original Destination load balancer in
	// Envoy for further details.
	LoadBalancerSettings_PASSTHROUGH LoadBalancerSettings_SimpleLB = 3
)

var LoadBalancerSettings_SimpleLB_name = map[int32]string{
	0: "ROUND_ROBIN",
	1: "LEAST_CONN",
	2: "RANDOM",
	3: "PASSTHROUGH",
}
var LoadBalancerSettings_SimpleLB_value = map[string]int32{
	"ROUND_ROBIN": 0,
	"LEAST_CONN":  1,
	"RANDOM":      2,
	"PASSTHROUGH": 3,
}

func (x LoadBalancerSettings_SimpleLB) String() string {
	return proto.EnumName(LoadBalancerSettings_SimpleLB_name, int32(x))
}
func (LoadBalancerSettings_SimpleLB) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{3, 0}
}

// TLS connection mode
type TLSSettings_TLSmode int32

const (
	// If set to "disable", the proxy will use not setup a TLS connection to the
	// upstream server.
	TLSSettings_DISABLE TLSSettings_TLSmode = 0
	// If set to "simple", the proxy will originate a TLS connection to the
	// upstream server.
	TLSSettings_SIMPLE TLSSettings_TLSmode = 1
	// If set to "mutual", the proxy will secure connections to the
	// upstream using mutual TLS by presenting client certificates for
	// authentication.
	TLSSettings_MUTUAL TLSSettings_TLSmode = 2
)

var TLSSettings_TLSmode_name = map[int32]string{
	0: "DISABLE",
	1: "SIMPLE",
	2: "MUTUAL",
}
var TLSSettings_TLSmode_value = map[string]int32{
	"DISABLE": 0,
	"SIMPLE":  1,
	"MUTUAL":  2,
}

func (x TLSSettings_TLSmode) String() string {
	return proto.EnumName(TLSSettings_TLSmode_name, int32(x))
}
func (TLSSettings_TLSmode) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{6, 0} }

// DestinationRule defines policies that apply to traffic intended for a
// service after routing has occurred. These rules specify configuration
// for load balancing, connection pool size from the sidecar, and outlier
// detection settings to detect and evict unhealthy hosts from the load
// balancing pool. For example, a simple load balancing policy for the
// ratings service would look as follows:
//
//     apiVersion: config.istio.io/v1alpha2
//     kind: DestinationRule
//     metadata:
//       name: bookinfo-ratings
//     spec:
//       name: ratings
//       trafficPolicy:
//         loadBalancer:
//           simple: LEAST_CONN
//
// Version specific policies can be specified by defining a named
// subset and overriding the settings specified at the service level. The
// following rule uses a round robin load balancing policy for all traffic
// going to a subset named testversion that is composed of endpoints (e.g.,
// pods) with labels (version:v3).
//
//     apiVersion: config.istio.io/v1alpha2
//     kind: DestinationRule
//     metadata:
//       name: bookinfo-ratings
//     spec:
//       name: ratings
//       trafficPolicy:
//         loadBalancer:
//           simple: LEAST_CONN
//       subsets:
//       - name: testversion
//         labels:
//           version: v3
//         trafficPolicy:
//           loadBalancer:
//             simple: ROUND_ROBIN
//
// Note that policies specified for subsets will not take effect until
// a route rule explicitly sends traffic to this subset.
type DestinationRule struct {
	// REQUIRED. The destination address for traffic captured by this
	// rule.  Could be a DNS name with wildcard prefix or a CIDR
	// prefix. Depending on the platform, short-names can also be used
	// instead of a FQDN (i.e. has no dots in the name). In such a scenario,
	// the FQDN of the host would be derived based on the underlying
	// platform.
	//
	// For example on Kubernetes, when hosts contains a short name, Istio
	// will interpret the short name based on the namespace of the client
	// where rules are being applied. Thus, when a client in the "default"
	// namespace applies a rule containing a name "reviews", Istio will setup
	// routes to the "reviews.default.svc.cluster.local" service. However, if
	// a different name such as "reviews.sales" is used, it would be treated
	// as a FQDN during virtual host matching.  In Consul, a plain service
	// name would be resolved to the FQDN "reviews.service.consul".
	//
	// Note that the hosts field applies to both HTTP and TCP
	// services. Service inside the mesh, i.e. those found in the service
	// registry, must always be referred to using their alphanumeric
	// names. IP addresses or CIDR prefixes are allowed only for services
	// defined via the Gateway.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// Traffic policies to apply (load balancing policy, connection pool
	// sizes, outlier detection).
	TrafficPolicy *TrafficPolicy `protobuf:"bytes,2,opt,name=traffic_policy,json=trafficPolicy" json:"traffic_policy,omitempty"`
	// One or more named sets that represent individual versions of a
	// service. Traffic policies can be overridden at subset level.
	Subsets []*Subset `protobuf:"bytes,3,rep,name=subsets" json:"subsets,omitempty"`
}

func (m *DestinationRule) Reset()                    { *m = DestinationRule{} }
func (m *DestinationRule) String() string            { return proto.CompactTextString(m) }
func (*DestinationRule) ProtoMessage()               {}
func (*DestinationRule) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *DestinationRule) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *DestinationRule) GetTrafficPolicy() *TrafficPolicy {
	if m != nil {
		return m.TrafficPolicy
	}
	return nil
}

func (m *DestinationRule) GetSubsets() []*Subset {
	if m != nil {
		return m.Subsets
	}
	return nil
}

// Traffic policies to apply for a specific destination. See
// DestinationRule for examples.
type TrafficPolicy struct {
	// Settings controlling the load balancer algorithms.
	LoadBalancer *LoadBalancerSettings `protobuf:"bytes,1,opt,name=load_balancer,json=loadBalancer" json:"load_balancer,omitempty"`
	// Settings controlling the volume of connections to an upstream service
	ConnectionPool *ConnectionPoolSettings `protobuf:"bytes,2,opt,name=connection_pool,json=connectionPool" json:"connection_pool,omitempty"`
	// Settings controlling eviction of unhealthy hosts from the load balancing pool
	OutlierDetection *OutlierDetection `protobuf:"bytes,3,opt,name=outlier_detection,json=outlierDetection" json:"outlier_detection,omitempty"`
	// TLS related settings for connections to the upstream service.
	Tls *TLSSettings `protobuf:"bytes,4,opt,name=tls" json:"tls,omitempty"`
}

func (m *TrafficPolicy) Reset()                    { *m = TrafficPolicy{} }
func (m *TrafficPolicy) String() string            { return proto.CompactTextString(m) }
func (*TrafficPolicy) ProtoMessage()               {}
func (*TrafficPolicy) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *TrafficPolicy) GetLoadBalancer() *LoadBalancerSettings {
	if m != nil {
		return m.LoadBalancer
	}
	return nil
}

func (m *TrafficPolicy) GetConnectionPool() *ConnectionPoolSettings {
	if m != nil {
		return m.ConnectionPool
	}
	return nil
}

func (m *TrafficPolicy) GetOutlierDetection() *OutlierDetection {
	if m != nil {
		return m.OutlierDetection
	}
	return nil
}

func (m *TrafficPolicy) GetTls() *TLSSettings {
	if m != nil {
		return m.Tls
	}
	return nil
}

// A subset of endpoints of a service. Subsets can be used for scenarios
// like A/B testing, or routing to a specific version of a service. Refer
// to VirtualSerive documentation for examples of using subsets in these
// scenarios. In addition, traffic policies defined at the service-level
// can be overridden at a subset-level. The following rule uses a round
// robin load balancing policy for all traffic going to a subset named
// testversion that is composed of endpoints (e.g., pods) with labels
// (version:v3).
//
//     apiVersion: config.istio.io/v1alpha2
//     kind: DestinationRule
//     metadata:
//       name: bookinfo-ratings
//     spec:
//       name: ratings
//       trafficPolicy:
//         loadBalancer:
//           simple: LEAST_CONN
//       subsets:
//       - name: testversion
//         labels:
//           version: v3
//         trafficPolicy:
//           loadBalancer:
//             simple: ROUND_ROBIN
//
// Note that policies specified for subsets will not take effect until
// a route rule explicitly sends traffic to this subset.
type Subset struct {
	// REQUIRED. name of the subset. The service name and the subset name can
	// be used for traffic splitting in a route rule.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// REQUIRED. Labels apply a filter over the endpoints of a service in the
	// service registry. See route rules for examples of usage.
	Labels map[string]string `protobuf:"bytes,2,rep,name=labels" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// Traffic policies that apply to this subset. Subsets inherit the
	// traffic policies specified at the DestinationRule level. Settings
	// specified at the subset level will override the corresponding settings
	// specified at the DestinationRule level.
	TrafficPolicy *TrafficPolicy `protobuf:"bytes,3,opt,name=traffic_policy,json=trafficPolicy" json:"traffic_policy,omitempty"`
}

func (m *Subset) Reset()                    { *m = Subset{} }
func (m *Subset) String() string            { return proto.CompactTextString(m) }
func (*Subset) ProtoMessage()               {}
func (*Subset) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Subset) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Subset) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *Subset) GetTrafficPolicy() *TrafficPolicy {
	if m != nil {
		return m.TrafficPolicy
	}
	return nil
}

// Load balancing policies to apply for a specific destination. See Envoy's
// load balancing
// [documentation](https://www.envoyproxy.io/docs/envoy/latest/intro/arch_overview/load_balancing.html)
// for more details.
//
// For example, the following rule uses a round robin load balancing policy
// for all traffic going to the ratings service.
//
//     apiVersion: config.istio.io/v1alpha2
//     kind: DestinationRule
//     metadata:
//       name: bookinfo-ratings
//     spec:
//       name: ratings
//       trafficPolicy:
//         loadBalancer:
//           simple: ROUND_ROBIN
//
// The following example uses the consistent hashing based load balancer
// for the same ratings service using the Cookie header as the hash key.
//
//     apiVersion: config.istio.io/v1alpha2
//     kind: DestinationRule
//     metadata:
//       name: bookinfo-ratings
//     spec:
//       name: ratings
//       trafficPolicy:
//         loadBalancer:
//           consistentHash:
//             http_header: Cookie
//
type LoadBalancerSettings struct {
	// Upstream load balancing policy.
	//
	// Types that are valid to be assigned to LbPolicy:
	//	*LoadBalancerSettings_Simple
	//	*LoadBalancerSettings_ConsistentHash
	LbPolicy isLoadBalancerSettings_LbPolicy `protobuf_oneof:"lb_policy"`
}

func (m *LoadBalancerSettings) Reset()                    { *m = LoadBalancerSettings{} }
func (m *LoadBalancerSettings) String() string            { return proto.CompactTextString(m) }
func (*LoadBalancerSettings) ProtoMessage()               {}
func (*LoadBalancerSettings) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type isLoadBalancerSettings_LbPolicy interface {
	isLoadBalancerSettings_LbPolicy()
}

type LoadBalancerSettings_Simple struct {
	Simple LoadBalancerSettings_SimpleLB `protobuf:"varint,1,opt,name=simple,enum=istio.routing.v1alpha2.LoadBalancerSettings_SimpleLB,oneof"`
}
type LoadBalancerSettings_ConsistentHash struct {
	ConsistentHash *LoadBalancerSettings_ConsistentHashLB `protobuf:"bytes,2,opt,name=consistent_hash,json=consistentHash,oneof"`
}

func (*LoadBalancerSettings_Simple) isLoadBalancerSettings_LbPolicy()         {}
func (*LoadBalancerSettings_ConsistentHash) isLoadBalancerSettings_LbPolicy() {}

func (m *LoadBalancerSettings) GetLbPolicy() isLoadBalancerSettings_LbPolicy {
	if m != nil {
		return m.LbPolicy
	}
	return nil
}

func (m *LoadBalancerSettings) GetSimple() LoadBalancerSettings_SimpleLB {
	if x, ok := m.GetLbPolicy().(*LoadBalancerSettings_Simple); ok {
		return x.Simple
	}
	return LoadBalancerSettings_ROUND_ROBIN
}

func (m *LoadBalancerSettings) GetConsistentHash() *LoadBalancerSettings_ConsistentHashLB {
	if x, ok := m.GetLbPolicy().(*LoadBalancerSettings_ConsistentHash); ok {
		return x.ConsistentHash
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*LoadBalancerSettings) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _LoadBalancerSettings_OneofMarshaler, _LoadBalancerSettings_OneofUnmarshaler, _LoadBalancerSettings_OneofSizer, []interface{}{
		(*LoadBalancerSettings_Simple)(nil),
		(*LoadBalancerSettings_ConsistentHash)(nil),
	}
}

func _LoadBalancerSettings_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*LoadBalancerSettings)
	// lb_policy
	switch x := m.LbPolicy.(type) {
	case *LoadBalancerSettings_Simple:
		b.EncodeVarint(1<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.Simple))
	case *LoadBalancerSettings_ConsistentHash:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ConsistentHash); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("LoadBalancerSettings.LbPolicy has unexpected type %T", x)
	}
	return nil
}

func _LoadBalancerSettings_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*LoadBalancerSettings)
	switch tag {
	case 1: // lb_policy.simple
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.LbPolicy = &LoadBalancerSettings_Simple{LoadBalancerSettings_SimpleLB(x)}
		return true, err
	case 2: // lb_policy.consistent_hash
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(LoadBalancerSettings_ConsistentHashLB)
		err := b.DecodeMessage(msg)
		m.LbPolicy = &LoadBalancerSettings_ConsistentHash{msg}
		return true, err
	default:
		return false, nil
	}
}

func _LoadBalancerSettings_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*LoadBalancerSettings)
	// lb_policy
	switch x := m.LbPolicy.(type) {
	case *LoadBalancerSettings_Simple:
		n += proto.SizeVarint(1<<3 | proto.WireVarint)
		n += proto.SizeVarint(uint64(x.Simple))
	case *LoadBalancerSettings_ConsistentHash:
		s := proto.Size(x.ConsistentHash)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// Consistent hashing (ketama hash) based load balancer for even load
// distribution/redistribution when the connection pool changes. This
// load balancing policy is applicable only for HTTP-based
// connections. A user specified HTTP header is used as the key with
// [xxHash](http://cyan4973.github.io/xxHash) hashing.
type LoadBalancerSettings_ConsistentHashLB struct {
	// REQUIRED. The name of the HTTP request header that will be used to
	// obtain the hash key. If the request header is not present, the load
	// balancer will use a random number as the hash, effectively making
	// the load balancing policy random.
	HttpHeader string `protobuf:"bytes,1,opt,name=http_header,json=httpHeader" json:"http_header,omitempty"`
	// The minimum number of virtual nodes to use for the hash
	// ring. Defaults to 1024. Larger ring sizes result in more granular
	// load distributions. If the number of hosts in the load balancing
	// pool is larger than the ring size, each host will be assigned a
	// single virtual node.
	MinimumRingSize uint32 `protobuf:"varint,2,opt,name=minimum_ring_size,json=minimumRingSize" json:"minimum_ring_size,omitempty"`
}

func (m *LoadBalancerSettings_ConsistentHashLB) Reset()         { *m = LoadBalancerSettings_ConsistentHashLB{} }
func (m *LoadBalancerSettings_ConsistentHashLB) String() string { return proto.CompactTextString(m) }
func (*LoadBalancerSettings_ConsistentHashLB) ProtoMessage()    {}
func (*LoadBalancerSettings_ConsistentHashLB) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{3, 0}
}

func (m *LoadBalancerSettings_ConsistentHashLB) GetHttpHeader() string {
	if m != nil {
		return m.HttpHeader
	}
	return ""
}

func (m *LoadBalancerSettings_ConsistentHashLB) GetMinimumRingSize() uint32 {
	if m != nil {
		return m.MinimumRingSize
	}
	return 0
}

// Connection pool settings for an upstream host. The settings apply to
// each individual host in the upstream service.  See Envoy's [circuit
// breaker](https://www.envoyproxy.io/docs/envoy/latest/intro/arch_overview/circuit_breaking)
// for more details. Connection pool settings can be applied at the TCP
// level as well as at HTTP level.
//
// For example, the following rule sets a limit of 100 connections to redis
// service called myredissrv with a connect timeout of 30ms
//
//     apiVersion: config.istio.io/v1alpha2
//     kind: DestinationRule
//     metadata:
//       name: bookinfo-redis
//     spec:
//       destination:
//         name: myredissrv
//       trafficPolicy:
//         connectionPool:
//           tcp:
//             maxConnections: 100
//             connectTimeout: 30ms
//
type ConnectionPoolSettings struct {
	// Settings common to both HTTP and TCP upstream connections.
	Tcp *ConnectionPoolSettings_TCPSettings `protobuf:"bytes,1,opt,name=tcp" json:"tcp,omitempty"`
	// HTTP connection pool settings.
	Http *ConnectionPoolSettings_HTTPSettings `protobuf:"bytes,2,opt,name=http" json:"http,omitempty"`
}

func (m *ConnectionPoolSettings) Reset()                    { *m = ConnectionPoolSettings{} }
func (m *ConnectionPoolSettings) String() string            { return proto.CompactTextString(m) }
func (*ConnectionPoolSettings) ProtoMessage()               {}
func (*ConnectionPoolSettings) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *ConnectionPoolSettings) GetTcp() *ConnectionPoolSettings_TCPSettings {
	if m != nil {
		return m.Tcp
	}
	return nil
}

func (m *ConnectionPoolSettings) GetHttp() *ConnectionPoolSettings_HTTPSettings {
	if m != nil {
		return m.Http
	}
	return nil
}

// Settings common to both HTTP and TCP upstream connections.
type ConnectionPoolSettings_TCPSettings struct {
	// Maximum number of HTTP1 /TCP connections to a destination host.
	MaxConnections int32 `protobuf:"varint,1,opt,name=max_connections,json=maxConnections" json:"max_connections,omitempty"`
	// TCP connection timeout.
	ConnectTimeout *google_protobuf.Duration `protobuf:"bytes,2,opt,name=connect_timeout,json=connectTimeout" json:"connect_timeout,omitempty"`
}

func (m *ConnectionPoolSettings_TCPSettings) Reset()         { *m = ConnectionPoolSettings_TCPSettings{} }
func (m *ConnectionPoolSettings_TCPSettings) String() string { return proto.CompactTextString(m) }
func (*ConnectionPoolSettings_TCPSettings) ProtoMessage()    {}
func (*ConnectionPoolSettings_TCPSettings) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{4, 0}
}

func (m *ConnectionPoolSettings_TCPSettings) GetMaxConnections() int32 {
	if m != nil {
		return m.MaxConnections
	}
	return 0
}

func (m *ConnectionPoolSettings_TCPSettings) GetConnectTimeout() *google_protobuf.Duration {
	if m != nil {
		return m.ConnectTimeout
	}
	return nil
}

// Settings applicable to HTTP1.1/HTTP2/GRPC connections.
type ConnectionPoolSettings_HTTPSettings struct {
	// Maximum number of pending HTTP requests to a destination. Default 1024.
	Http1MaxPendingRequests int32 `protobuf:"varint,1,opt,name=http1_max_pending_requests,json=http1MaxPendingRequests" json:"http1_max_pending_requests,omitempty"`
	// Maximum number of requests to a backend. Default 1024.
	Http2MaxRequests int32 `protobuf:"varint,2,opt,name=http2_max_requests,json=http2MaxRequests" json:"http2_max_requests,omitempty"`
	// Maximum number of requests per connection to a backend. Setting this
	// parameter to 1 disables keep alive.
	MaxRequestsPerConnection int32 `protobuf:"varint,3,opt,name=max_requests_per_connection,json=maxRequestsPerConnection" json:"max_requests_per_connection,omitempty"`
	// Maximum number of retries that can be outstanding to all hosts in a
	// cluster at a given time. Defaults to 3.
	MaxRetries int32 `protobuf:"varint,4,opt,name=max_retries,json=maxRetries" json:"max_retries,omitempty"`
}

func (m *ConnectionPoolSettings_HTTPSettings) Reset()         { *m = ConnectionPoolSettings_HTTPSettings{} }
func (m *ConnectionPoolSettings_HTTPSettings) String() string { return proto.CompactTextString(m) }
func (*ConnectionPoolSettings_HTTPSettings) ProtoMessage()    {}
func (*ConnectionPoolSettings_HTTPSettings) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{4, 1}
}

func (m *ConnectionPoolSettings_HTTPSettings) GetHttp1MaxPendingRequests() int32 {
	if m != nil {
		return m.Http1MaxPendingRequests
	}
	return 0
}

func (m *ConnectionPoolSettings_HTTPSettings) GetHttp2MaxRequests() int32 {
	if m != nil {
		return m.Http2MaxRequests
	}
	return 0
}

func (m *ConnectionPoolSettings_HTTPSettings) GetMaxRequestsPerConnection() int32 {
	if m != nil {
		return m.MaxRequestsPerConnection
	}
	return 0
}

func (m *ConnectionPoolSettings_HTTPSettings) GetMaxRetries() int32 {
	if m != nil {
		return m.MaxRetries
	}
	return 0
}

// A Circuit breaker implementation that tracks the status of each
// individual host in the upstream service.  While currently applicable to
// only HTTP services, future versions will support opaque TCP services as
// well. For HTTP services, hosts that continually return errors for API
// calls are ejected from the pool for a pre-defined period of time. See
// Envoy's [outlier
// detection](https://www.envoyproxy.io/docs/envoy/latest/intro/arch_overview/outlier)
// for more details.
//
// The following rule sets a connection pool size of 100 connections and
// 1000 concurrent HTTP2 requests, with no more than 10 req/connection to
// "reviews" service. In addition, it configures upstream hosts to be
// scanned every 5 mins, such that any host that fails 7 consecutive times
// with 5XX error code will be ejected for 15 minutes.
//
//     apiVersion: config.istio.io/v1alpha2
//     kind: DestinationRule
//     metadata:
//       name: reviews-cb-policy
//     spec:
//       destination:
//         name: reviews
//       trafficPolicy:
//         connectionPool:
//           tcp:
//             maxConnections: 100
//           http:
//             http2MaxRequests: 1000
//             maxRequestsPerConnection: 10
//         outlierDetection:
//           http:
//             consecutiveErrors: 7
//             interval: 5m
//             baseEjectionTime: 15m
//
type OutlierDetection struct {
	// Settings for HTTP1.1/HTTP2/GRPC connections.
	Http *OutlierDetection_HTTPSettings `protobuf:"bytes,1,opt,name=http" json:"http,omitempty"`
}

func (m *OutlierDetection) Reset()                    { *m = OutlierDetection{} }
func (m *OutlierDetection) String() string            { return proto.CompactTextString(m) }
func (*OutlierDetection) ProtoMessage()               {}
func (*OutlierDetection) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *OutlierDetection) GetHttp() *OutlierDetection_HTTPSettings {
	if m != nil {
		return m.Http
	}
	return nil
}

// Outlier detection settings for HTTP1.1/HTTP2/GRPC connections.
type OutlierDetection_HTTPSettings struct {
	// Number of 5XX errors before a host is ejected from the connection
	// pool. Defaults to 5.
	ConsecutiveErrors int32 `protobuf:"varint,1,opt,name=consecutive_errors,json=consecutiveErrors" json:"consecutive_errors,omitempty"`
	// Time interval between ejection sweep analysis. format:
	// 1h/1m/1s/1ms. MUST BE >=1ms. Default is 10s.
	Interval *google_protobuf.Duration `protobuf:"bytes,2,opt,name=interval" json:"interval,omitempty"`
	// Minimum ejection duration. A host will remain ejected for a period
	// equal to the product of minimum ejection duration and the number of
	// times the host has been ejected. This technique allows the system to
	// automatically increase the ejection period for unhealthy upstream
	// servers. format: 1h/1m/1s/1ms. MUST BE >=1ms. Default is 30s.
	BaseEjectionTime *google_protobuf.Duration `protobuf:"bytes,3,opt,name=base_ejection_time,json=baseEjectionTime" json:"base_ejection_time,omitempty"`
	// Maximum % of hosts in the load balancing pool for the upstream
	// service that can be ejected. Defaults to 10%.
	MaxEjectionPercent int32 `protobuf:"varint,4,opt,name=max_ejection_percent,json=maxEjectionPercent" json:"max_ejection_percent,omitempty"`
}

func (m *OutlierDetection_HTTPSettings) Reset()         { *m = OutlierDetection_HTTPSettings{} }
func (m *OutlierDetection_HTTPSettings) String() string { return proto.CompactTextString(m) }
func (*OutlierDetection_HTTPSettings) ProtoMessage()    {}
func (*OutlierDetection_HTTPSettings) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{5, 0}
}

func (m *OutlierDetection_HTTPSettings) GetConsecutiveErrors() int32 {
	if m != nil {
		return m.ConsecutiveErrors
	}
	return 0
}

func (m *OutlierDetection_HTTPSettings) GetInterval() *google_protobuf.Duration {
	if m != nil {
		return m.Interval
	}
	return nil
}

func (m *OutlierDetection_HTTPSettings) GetBaseEjectionTime() *google_protobuf.Duration {
	if m != nil {
		return m.BaseEjectionTime
	}
	return nil
}

func (m *OutlierDetection_HTTPSettings) GetMaxEjectionPercent() int32 {
	if m != nil {
		return m.MaxEjectionPercent
	}
	return 0
}

// SSL/TLS related settings for upstream connections. See Envoy's [TLS
// context](https://www.envoyproxy.io/docs/envoy/latest/api-v1/cluster_manager/cluster_ssl.html#config-cluster-manager-cluster-ssl)
// for more details. These settings are common to both HTTP and TCP upstreams.
//
// For example, the following rule configures a client to use mutual TLS
// for connections to upstream database cluster.
//
//     apiVersion: config.istio.io/v1alpha2
//     kind: DestinationRule
//     metadata:
//       name: db-mtls
//     spec:
//       destination:
//         name: mydbserver
//       tls:
//         mode: MUTUAL
//         clientCertificate: /etc/certs/myclientcert.pem
//         privateKey: /etc/certs/client_private_key.pem
//         caCertificates: /etc/certs/rootcacerts.pem
//
// The following rule configures a client to use TLS when talking to a foreign service whose domain matches *.foo.com.
//
//     apiVersion: config.istio.io/v1alpha2
//     kind: DestinationRule
//     metadata:
//       name: tls-foo
//     spec:
//       destination:
//         name: *.foo.com
//       tls:
//         mode: SIMPLE
//
type TLSSettings struct {
	// REQUIRED: Indicates whether connections to this port should be secured
	// using TLS. The value of this field determines how TLS is enforced.
	Mode TLSSettings_TLSmode `protobuf:"varint,1,opt,name=mode,enum=istio.routing.v1alpha2.TLSSettings_TLSmode" json:"mode,omitempty"`
	// REQUIRED if mode is "mutual". The path to the file holding the
	// client-side TLS certificate to use.
	ClientCertificate string `protobuf:"bytes,2,opt,name=client_certificate,json=clientCertificate" json:"client_certificate,omitempty"`
	// REQUIRED if mode is "mutual". The path to the file holding the
	// client's private key.
	PrivateKey string `protobuf:"bytes,3,opt,name=private_key,json=privateKey" json:"private_key,omitempty"`
	// OPTIONAL: The path to the file containing certificate authority
	// certificates to use in verifying a presented server certificate. If
	// omitted, the proxy will not verify the server's certificate.
	CaCertificates string `protobuf:"bytes,4,opt,name=ca_certificates,json=caCertificates" json:"ca_certificates,omitempty"`
	// A list of alternate names to verify the subject identity in the
	// certificate. If specified, the proxy will verify that the server
	// certificate's subject alt name matches one of the specified values.
	SubjectAltNames []string `protobuf:"bytes,5,rep,name=subject_alt_names,json=subjectAltNames" json:"subject_alt_names,omitempty"`
	// SNI string to present to the server during TLS handshake.
	Sni string `protobuf:"bytes,6,opt,name=sni" json:"sni,omitempty"`
}

func (m *TLSSettings) Reset()                    { *m = TLSSettings{} }
func (m *TLSSettings) String() string            { return proto.CompactTextString(m) }
func (*TLSSettings) ProtoMessage()               {}
func (*TLSSettings) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *TLSSettings) GetMode() TLSSettings_TLSmode {
	if m != nil {
		return m.Mode
	}
	return TLSSettings_DISABLE
}

func (m *TLSSettings) GetClientCertificate() string {
	if m != nil {
		return m.ClientCertificate
	}
	return ""
}

func (m *TLSSettings) GetPrivateKey() string {
	if m != nil {
		return m.PrivateKey
	}
	return ""
}

func (m *TLSSettings) GetCaCertificates() string {
	if m != nil {
		return m.CaCertificates
	}
	return ""
}

func (m *TLSSettings) GetSubjectAltNames() []string {
	if m != nil {
		return m.SubjectAltNames
	}
	return nil
}

func (m *TLSSettings) GetSni() string {
	if m != nil {
		return m.Sni
	}
	return ""
}

func init() {
	proto.RegisterType((*DestinationRule)(nil), "istio.routing.v1alpha2.DestinationRule")
	proto.RegisterType((*TrafficPolicy)(nil), "istio.routing.v1alpha2.TrafficPolicy")
	proto.RegisterType((*Subset)(nil), "istio.routing.v1alpha2.Subset")
	proto.RegisterType((*LoadBalancerSettings)(nil), "istio.routing.v1alpha2.LoadBalancerSettings")
	proto.RegisterType((*LoadBalancerSettings_ConsistentHashLB)(nil), "istio.routing.v1alpha2.LoadBalancerSettings.ConsistentHashLB")
	proto.RegisterType((*ConnectionPoolSettings)(nil), "istio.routing.v1alpha2.ConnectionPoolSettings")
	proto.RegisterType((*ConnectionPoolSettings_TCPSettings)(nil), "istio.routing.v1alpha2.ConnectionPoolSettings.TCPSettings")
	proto.RegisterType((*ConnectionPoolSettings_HTTPSettings)(nil), "istio.routing.v1alpha2.ConnectionPoolSettings.HTTPSettings")
	proto.RegisterType((*OutlierDetection)(nil), "istio.routing.v1alpha2.OutlierDetection")
	proto.RegisterType((*OutlierDetection_HTTPSettings)(nil), "istio.routing.v1alpha2.OutlierDetection.HTTPSettings")
	proto.RegisterType((*TLSSettings)(nil), "istio.routing.v1alpha2.TLSSettings")
	proto.RegisterEnum("istio.routing.v1alpha2.LoadBalancerSettings_SimpleLB", LoadBalancerSettings_SimpleLB_name, LoadBalancerSettings_SimpleLB_value)
	proto.RegisterEnum("istio.routing.v1alpha2.TLSSettings_TLSmode", TLSSettings_TLSmode_name, TLSSettings_TLSmode_value)
}

func init() { proto.RegisterFile("routing/v1alpha2/destination_rule.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 1054 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x56, 0xdd, 0x6e, 0x1a, 0x47,
	0x14, 0x36, 0xac, 0x8d, 0xc3, 0x21, 0x86, 0xf5, 0xc8, 0x4a, 0x29, 0x55, 0x13, 0x8b, 0xaa, 0x8a,
	0x95, 0xa6, 0x4b, 0x43, 0x65, 0x29, 0x4d, 0x14, 0x55, 0x60, 0x50, 0xb0, 0x8a, 0x81, 0x0e, 0x58,
	0x95, 0x7a, 0xb3, 0x1a, 0x96, 0xb1, 0x99, 0x76, 0xff, 0xba, 0x33, 0x6b, 0xd9, 0x7e, 0x83, 0x5e,
	0xf5, 0x49, 0x7a, 0xd9, 0x87, 0xe8, 0x55, 0xaf, 0xfb, 0x08, 0xbd, 0xea, 0x2b, 0x54, 0xf3, 0x03,
	0xac, 0x2d, 0x9c, 0xc6, 0xb9, 0x5b, 0xce, 0xf9, 0xce, 0x37, 0x67, 0xce, 0xcf, 0x37, 0xc0, 0xd3,
	0x24, 0x4a, 0x05, 0x0b, 0xcf, 0x1b, 0x17, 0x2f, 0x88, 0x1f, 0xcf, 0x49, 0xb3, 0x31, 0xa3, 0x5c,
	0xb0, 0x90, 0x08, 0x16, 0x85, 0x6e, 0x92, 0xfa, 0xd4, 0x89, 0x93, 0x48, 0x44, 0xe8, 0x11, 0xe3,
	0x82, 0x45, 0x8e, 0x81, 0x3b, 0x0b, 0x78, 0xed, 0xf1, 0x79, 0x14, 0x9d, 0xfb, 0xb4, 0xa1, 0x50,
	0xd3, 0xf4, 0xac, 0x31, 0x4b, 0x13, 0x15, 0xac, 0xe3, 0xea, 0xbf, 0xe7, 0xa0, 0xd2, 0x59, 0x51,
	0xe2, 0xd4, 0xa7, 0x08, 0xc1, 0x66, 0x48, 0x02, 0x5a, 0xcd, 0xed, 0xe7, 0x0e, 0x8a, 0x58, 0x7d,
	0xa3, 0x3e, 0x94, 0x45, 0x42, 0xce, 0xce, 0x98, 0xe7, 0xc6, 0x91, 0xcf, 0xbc, 0xab, 0x6a, 0x7e,
	0x3f, 0x77, 0x50, 0x6a, 0x7e, 0xee, 0xac, 0x3f, 0xd8, 0x99, 0x68, 0xf4, 0x48, 0x81, 0xf1, 0x8e,
	0xc8, 0xfe, 0x44, 0x2f, 0x61, 0x9b, 0xa7, 0x53, 0x4e, 0x05, 0xaf, 0x5a, 0xfb, 0xd6, 0x41, 0xa9,
	0xf9, 0xf8, 0x2e, 0x9a, 0xb1, 0x82, 0xe1, 0x05, 0xbc, 0xfe, 0x67, 0x1e, 0x76, 0x6e, 0x50, 0xa3,
	0xef, 0x61, 0xc7, 0x8f, 0xc8, 0xcc, 0x9d, 0x12, 0x9f, 0x84, 0x1e, 0x4d, 0x54, 0xda, 0xa5, 0xe6,
	0xf3, 0xbb, 0x18, 0xfb, 0x11, 0x99, 0xb5, 0x0d, 0x76, 0x4c, 0x85, 0x74, 0x72, 0xfc, 0xd0, 0xcf,
	0x58, 0xd1, 0x0f, 0x50, 0xf1, 0xa2, 0x30, 0xa4, 0x9e, 0xaa, 0x72, 0x1c, 0x45, 0xbe, 0xb9, 0xad,
	0x73, 0x17, 0xe9, 0xd1, 0x12, 0x3e, 0x8a, 0x22, 0x7f, 0x49, 0x5b, 0xf6, 0x6e, 0xd8, 0xd1, 0x29,
	0xec, 0x46, 0xa9, 0xf0, 0x19, 0x4d, 0xdc, 0x19, 0x15, 0xda, 0x51, 0xb5, 0x14, 0xf5, 0xc1, 0x5d,
	0xd4, 0x43, 0x1d, 0xd0, 0x59, 0xe0, 0xb1, 0x1d, 0xdd, 0xb2, 0xa0, 0x43, 0xb0, 0x84, 0xcf, 0xab,
	0x9b, 0x8a, 0xe8, 0xb3, 0x3b, 0x3b, 0xd2, 0x1f, 0x2f, 0x13, 0x93, 0xf8, 0xfa, 0x3f, 0x39, 0x28,
	0xe8, 0xfa, 0xae, 0x6d, 0x79, 0x1b, 0x0a, 0x3e, 0x99, 0x52, 0x9f, 0x57, 0xf3, 0xaa, 0x47, 0xcf,
	0xde, 0xdd, 0x23, 0xa7, 0xaf, 0xc0, 0xdd, 0x50, 0x24, 0x57, 0xd8, 0x44, 0xae, 0x19, 0x1b, 0xeb,
	0xc3, 0xc7, 0xa6, 0xf6, 0x0d, 0x94, 0x32, 0x87, 0x20, 0x1b, 0xac, 0x9f, 0xe9, 0x95, 0xc9, 0x59,
	0x7e, 0xa2, 0x3d, 0xd8, 0xba, 0x20, 0x7e, 0x4a, 0x55, 0xbb, 0x8a, 0x58, 0xff, 0x78, 0x95, 0x7f,
	0x99, 0xab, 0xff, 0x6a, 0xc1, 0xde, 0xba, 0xce, 0xa3, 0x21, 0x14, 0x38, 0x0b, 0x62, 0x5f, 0xdf,
	0xbd, 0xdc, 0x3c, 0xbc, 0xcf, 0xdc, 0x38, 0x63, 0x15, 0xda, 0x6f, 0xf7, 0x36, 0xb0, 0xa1, 0x41,
	0x73, 0x35, 0x3c, 0x9c, 0x71, 0x41, 0x43, 0xe1, 0xce, 0x09, 0x9f, 0x9b, 0xe1, 0x79, 0x73, 0x2f,
	0xe6, 0xa3, 0x25, 0x47, 0x8f, 0xf0, 0xb9, 0x3a, 0xa1, 0xec, 0xdd, 0xb0, 0xd5, 0x5c, 0xb0, 0x6f,
	0xa3, 0xd0, 0x13, 0x28, 0xcd, 0x85, 0x88, 0xdd, 0x39, 0x25, 0x33, 0xb3, 0x0b, 0x45, 0x0c, 0xd2,
	0xd4, 0x53, 0x16, 0xf4, 0x0c, 0x76, 0x03, 0x16, 0xb2, 0x20, 0x0d, 0xdc, 0x84, 0x85, 0xe7, 0x2e,
	0x67, 0xd7, 0xba, 0x5c, 0x3b, 0xb8, 0x62, 0x1c, 0x98, 0x85, 0xe7, 0x63, 0x76, 0x4d, 0xeb, 0x3d,
	0x78, 0xb0, 0xb8, 0x20, 0xaa, 0x40, 0x09, 0x0f, 0x4f, 0x07, 0x1d, 0x17, 0x0f, 0xdb, 0xc7, 0x03,
	0x7b, 0x03, 0x95, 0x01, 0xfa, 0xdd, 0xd6, 0x78, 0xe2, 0x1e, 0x0d, 0x07, 0x03, 0x3b, 0x87, 0x00,
	0x0a, 0xb8, 0x35, 0xe8, 0x0c, 0x4f, 0xec, 0xbc, 0x04, 0x8f, 0x5a, 0xe3, 0xf1, 0xa4, 0x87, 0x87,
	0xa7, 0x6f, 0x7b, 0xb6, 0xd5, 0x2e, 0x41, 0xd1, 0x9f, 0x9a, 0x11, 0xa8, 0xff, 0xb6, 0x09, 0x8f,
	0xd6, 0x2f, 0x0c, 0xea, 0x83, 0x25, 0xbc, 0xd8, 0xac, 0xf0, 0xab, 0xfb, 0x6d, 0x9b, 0x33, 0x39,
	0x1a, 0x65, 0x06, 0xdc, 0x8b, 0xd1, 0x10, 0x36, 0xe5, 0xcd, 0x4d, 0xfd, 0x5f, 0xdf, 0x93, 0xae,
	0x37, 0x99, 0xac, 0xf8, 0x14, 0x51, 0xed, 0x1a, 0x4a, 0x99, 0x43, 0xd0, 0x53, 0xa8, 0x04, 0xe4,
	0xd2, 0x5d, 0x2d, 0x39, 0x57, 0x99, 0x6f, 0xe1, 0x72, 0x40, 0x2e, 0x57, 0xac, 0x1c, 0xb5, 0x97,
	0x82, 0xe2, 0x0a, 0x16, 0xd0, 0x28, 0x15, 0x26, 0xa7, 0x8f, 0x1d, 0xad, 0xcf, 0xce, 0x42, 0x9f,
	0x9d, 0x8e, 0xd1, 0xe7, 0xa5, 0x76, 0x4c, 0x74, 0x40, 0xed, 0xef, 0x1c, 0x3c, 0xcc, 0xa6, 0x84,
	0x5e, 0x43, 0x4d, 0x26, 0xf5, 0xc2, 0x95, 0x39, 0xc4, 0x34, 0x9c, 0xc9, 0x76, 0x26, 0xf4, 0x97,
	0x94, 0x72, 0xb1, 0x48, 0xe4, 0x23, 0x85, 0x38, 0x21, 0x97, 0x23, 0xed, 0xc7, 0xc6, 0x8d, 0x9e,
	0x03, 0x92, 0xae, 0xa6, 0x0a, 0x5e, 0x06, 0xe5, 0x55, 0x90, 0xad, 0x3c, 0x27, 0xe4, 0x72, 0x89,
	0x7e, 0x03, 0x9f, 0x64, 0x71, 0x6e, 0x4c, 0x93, 0xcc, 0xad, 0xd5, 0x4e, 0x6f, 0xe1, 0x6a, 0xb0,
	0x8a, 0x18, 0xd1, 0x64, 0x75, 0x7f, 0x39, 0x94, 0x3a, 0x5c, 0x24, 0x8c, 0x6a, 0x9d, 0xda, 0xc2,
	0xa0, 0xe0, 0xca, 0x52, 0xff, 0x2b, 0x0f, 0xf6, 0x6d, 0x9d, 0x43, 0xc7, 0xa6, 0x7b, 0x7a, 0x18,
	0x0e, 0xdf, 0x57, 0x1f, 0xd7, 0xf5, 0xed, 0xdf, 0xdb, 0xb5, 0xfb, 0x12, 0x90, 0x5c, 0x26, 0xea,
	0xa5, 0x82, 0x5d, 0x50, 0x97, 0x26, 0x49, 0x94, 0x2c, 0x6a, 0xb6, 0x9b, 0xf1, 0x74, 0x95, 0x03,
	0x1d, 0xc2, 0x03, 0x16, 0x0a, 0x9a, 0x5c, 0x10, 0xff, 0xff, 0x1b, 0xb7, 0x84, 0xa2, 0xb7, 0x80,
	0xa6, 0x84, 0x53, 0x97, 0xfe, 0x64, 0x9e, 0x12, 0xd9, 0x7c, 0xa3, 0x80, 0xef, 0x20, 0xb0, 0x65,
	0x50, 0xd7, 0xc4, 0xc8, 0xf6, 0xa3, 0xaf, 0x60, 0x4f, 0x16, 0x70, 0xc9, 0x13, 0xd3, 0xc4, 0xa3,
	0xa1, 0x30, 0x95, 0x44, 0x01, 0xb9, 0x5c, 0xc0, 0x47, 0xda, 0x53, 0xff, 0x23, 0x0f, 0xa5, 0x8c,
	0xe0, 0xa3, 0x6f, 0x61, 0x33, 0x88, 0x66, 0x0b, 0x91, 0xfb, 0xe2, 0x3d, 0xde, 0x08, 0xf9, 0x2d,
	0x43, 0xb0, 0x0a, 0x54, 0x15, 0xf3, 0x99, 0x94, 0x34, 0x8f, 0x26, 0x82, 0x9d, 0x31, 0x8f, 0x88,
	0x85, 0xce, 0xee, 0x6a, 0xcf, 0xd1, 0xca, 0x21, 0x5b, 0x1e, 0x27, 0xec, 0x82, 0x08, 0xea, 0x4a,
	0x8d, 0xb6, 0xb4, 0x0e, 0x19, 0xd3, 0x77, 0xf4, 0x4a, 0xee, 0x8e, 0x47, 0xb2, 0x5c, 0x7a, 0x2e,
	0x8a, 0xb8, 0xec, 0x91, 0x0c, 0x11, 0x97, 0x82, 0xc5, 0xd3, 0xa9, 0xbc, 0x9e, 0x4b, 0x7c, 0xe1,
	0xca, 0xa7, 0x89, 0x57, 0xb7, 0xf6, 0xad, 0x83, 0x22, 0xae, 0x18, 0x47, 0xcb, 0x17, 0x03, 0x69,
	0x96, 0x2f, 0x02, 0x0f, 0x59, 0xb5, 0xa0, 0x5f, 0x04, 0x1e, 0xb2, 0xba, 0x03, 0xdb, 0xe6, 0x1e,
	0xa8, 0x04, 0xdb, 0x9d, 0xe3, 0x71, 0xab, 0xdd, 0xef, 0xda, 0x1b, 0x52, 0xad, 0xc6, 0xc7, 0x27,
	0xa3, 0x7e, 0x57, 0x2b, 0xd7, 0xc9, 0xe9, 0xe4, 0xb4, 0xd5, 0xb7, 0xf3, 0xed, 0x27, 0x3f, 0x7e,
	0xaa, 0x4b, 0xc3, 0xa2, 0x06, 0x89, 0x59, 0xe3, 0xf6, 0xff, 0xaf, 0x69, 0x41, 0xf5, 0xeb, 0xeb,
	0xff, 0x02, 0x00, 0x00, 0xff, 0xff, 0x5f, 0x59, 0xb8, 0x8a, 0x9a, 0x09, 0x00, 0x00,
}
