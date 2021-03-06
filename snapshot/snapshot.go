package snapshot

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
	"time"

	"code.cloudfoundry.org/copilot/models"
	"code.cloudfoundry.org/lager"

	"code.cloudfoundry.org/policy_client"
	"istio.io/istio/pilot/pkg/model"
	snap "istio.io/istio/pkg/mcp/snapshot"
)

var (
	// TODO: Remove unsupported typeURLs (everything except Gateway, VirtualService, DestinationRule)
	// when mcp client is capable of only sending a subset of the types
	AuthenticationPolicyTypeURL     string
	AuthenticationMeshPolicyTypeURL string
	AuthorizationPolicyTypeURL      string
	DestinationRuleTypeURL          string
	VirtualServiceTypeURL           string
	GatewayTypeURL                  string
	ServiceEntryTypeURL             string
	EnvoyFilterTypeURL              string
	SidecarTypeURL                  string
	HTTPAPISpecTypeURL              string
	HTTPAPISpecBindingTypeURL       string
	QuotaSpecTypeURL                string
	QuotaSpecBindingTypeURL         string
	PolicyTypeURL                   string
	MeshPolicyTypeURL               string
	ServiceRoleTypeURL              string
	ServiceRoleBindingTypeURL       string
	RbacConfigTypeURL               string
	ClusterRbacConfigTypeURL        string
)

const (
	DefaultGatewayName = "cloudfoundry-ingress"
	DefaultSidecarName = "cloudfoundry-default-sidecar"

	// TODO: Do not specify the nodeID yet as it's used as a key for cache lookup
	// in snapshot, we should add this once the nodeID is configurable in pilot
	node        = "default"
	gatewayPort = 80
)

//go:generate counterfeiter -o fakes/collector.go --fake-name Collector . collector
type collector interface {
	Collect() []*models.RouteWithBackends
}

//go:generate counterfeiter -o fakes/setter.go --fake-name Setter . setter
type setter interface {
	SetSnapshot(node string, istio snap.Snapshot)
}

type Snapshot struct {
	logger             lager.Logger
	ticker             <-chan time.Time
	collector          collector
	setter             setter
	builder            *snap.InMemoryBuilder
	cachedRoutes       []*models.RouteWithBackends
	cachedPolicies     []*policy_client.Policy
	config             config
	policyServerClient policy_client.InternalPolicyClient
	ver                int
	initialized        bool
}

func New(logger lager.Logger, ticker <-chan time.Time, collector collector, setter setter, builder *snap.InMemoryBuilder, policyServerClient policy_client.InternalPolicyClient, config config) *Snapshot {
	return &Snapshot{
		logger:             logger,
		ticker:             ticker,
		collector:          collector,
		setter:             setter,
		builder:            builder,
		policyServerClient: policyServerClient,
		config:             config,
	}
}

func (s *Snapshot) Run(signals <-chan os.Signal, ready chan<- struct{}) error {
	close(ready)

	for {
		select {
		case <-signals:
			return nil
		case <-s.ticker:
			routes := s.collector.Collect()

			var policies []*policy_client.Policy
			if s.policyServerClient != nil {
				var policyError error
				policies, _, policyError = s.policyServerClient.GetPolicies()
				if policyError != nil {
					s.logger.Error("Error fetching policies", policyError)
					continue
				}
			}

			if s.initialized &&
				reflect.DeepEqual(routes, s.cachedRoutes) &&
				reflect.DeepEqual(policies, s.cachedPolicies) {
				continue
			}

			newVersion := s.increment()
			s.cachedRoutes = routes
			s.cachedPolicies = policies

			gateways := s.config.CreateGatewayResources()
			sidecars := s.config.CreateSidecarResources(routes, policies, newVersion)
			virtualServices := s.config.CreateVirtualServiceResources(routes, newVersion)
			destinationRules := s.config.CreateDestinationRuleResources(routes, newVersion)
			serviceEntries := s.config.CreateServiceEntryResources(routes, newVersion)

			s.builder.Set(GatewayTypeURL, "1", gateways)
			s.builder.Set(SidecarTypeURL, newVersion, sidecars)
			s.builder.Set(VirtualServiceTypeURL, newVersion, virtualServices)
			s.builder.Set(DestinationRuleTypeURL, newVersion, destinationRules)
			s.builder.Set(ServiceEntryTypeURL, newVersion, serviceEntries)

			//Empty responses
			s.builder.Set(AuthenticationPolicyTypeURL, "1", s.config.EmptyResponse())
			s.builder.Set(AuthorizationPolicyTypeURL, "1", s.config.EmptyResponse())
			s.builder.Set(ClusterRbacConfigTypeURL, "1", s.config.EmptyResponse())
			s.builder.Set(EnvoyFilterTypeURL, "1", s.config.EmptyResponse())
			s.builder.Set(HTTPAPISpecBindingTypeURL, "1", s.config.EmptyResponse())
			s.builder.Set(HTTPAPISpecTypeURL, "1", s.config.EmptyResponse())
			s.builder.Set(MeshPolicyTypeURL, "1", s.config.EmptyResponse())
			s.builder.Set(QuotaSpecBindingTypeURL, "1", s.config.EmptyResponse())
			s.builder.Set(QuotaSpecTypeURL, "1", s.config.EmptyResponse())
			s.builder.Set(ServiceRoleBindingTypeURL, "1", s.config.EmptyResponse())
			s.builder.Set(ServiceRoleTypeURL, "1", s.config.EmptyResponse())
			s.builder.Set(RbacConfigTypeURL, "1", s.config.EmptyResponse())

			shot := s.builder.Build()
			s.setter.SetSnapshot(node, shot)
			s.builder = shot.Builder()
			s.initialized = true
		}
	}
}

func (s *Snapshot) version() string {
	return strconv.Itoa(s.ver)
}

func (s *Snapshot) increment() string {
	s.ver++
	return s.version()
}

func getTypeURLByType(name string) string {
	protoSchema, ok := model.IstioConfigTypes.GetByType(name)
	if !ok {
		fmt.Fprintf(os.Stdout, "Istio Config Type %q does not exist.\n", name)
		os.Exit(1)
	}

	return protoSchema.Collection
}

func init() {
	AuthenticationPolicyTypeURL = getTypeURLByType("policy")
	AuthenticationMeshPolicyTypeURL = getTypeURLByType("mesh-policy")
	AuthorizationPolicyTypeURL = getTypeURLByType("authorization-policy")
	DestinationRuleTypeURL = getTypeURLByType("destination-rule")
	VirtualServiceTypeURL = getTypeURLByType("virtual-service")
	GatewayTypeURL = getTypeURLByType("gateway")
	ServiceEntryTypeURL = getTypeURLByType("service-entry")
	EnvoyFilterTypeURL = getTypeURLByType("envoy-filter")
	SidecarTypeURL = getTypeURLByType("sidecar")
	HTTPAPISpecTypeURL = getTypeURLByType("http-api-spec")
	HTTPAPISpecBindingTypeURL = getTypeURLByType("http-api-spec-binding")
	QuotaSpecTypeURL = getTypeURLByType("quota-spec")
	QuotaSpecBindingTypeURL = getTypeURLByType("quota-spec-binding")
	PolicyTypeURL = getTypeURLByType("policy")
	MeshPolicyTypeURL = getTypeURLByType("mesh-policy")
	ServiceRoleTypeURL = getTypeURLByType("service-role")
	ServiceRoleBindingTypeURL = getTypeURLByType("service-role-binding")
	RbacConfigTypeURL = getTypeURLByType("rbac-config")
	ClusterRbacConfigTypeURL = getTypeURLByType("cluster-rbac-config")
}
