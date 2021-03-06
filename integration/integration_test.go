package integration_test

import (
	"code.cloudfoundry.org/policy_client"
	"context"
	"crypto/tls"
	"fmt"
	"os"
	"os/exec"
	"time"

	bbsmodels "code.cloudfoundry.org/bbs/models"
	"code.cloudfoundry.org/copilot"
	"code.cloudfoundry.org/copilot/api"
	"code.cloudfoundry.org/copilot/certs"
	"code.cloudfoundry.org/copilot/config"
	copilotsnapshot "code.cloudfoundry.org/copilot/snapshot"
	"code.cloudfoundry.org/copilot/testhelpers"
	"code.cloudfoundry.org/durationjson"
	"github.com/gogo/protobuf/proto"
	"github.com/gogo/protobuf/types"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
	"github.com/onsi/gomega/gexec"
	"google.golang.org/grpc"
	"istio.io/api/networking/v1alpha3"
)

var _ = Describe("Copilot", func() {
	var (
		session                        *gexec.Session
		serverConfig                   *config.Config
		pilotClientTLSConfig           *tls.Config
		cloudControllerClientTLSConfig *tls.Config
		configFilePath                 string

		mcpClient         *testhelpers.MockPilotMCPClient
		ccClient          copilot.CloudControllerClient
		vipResolverClient copilot.VIPResolverCopilotClient
		mockBBS           *testhelpers.MockBBSServer
		mockPolicyServer  *testhelpers.MockPolicyServer

		cleanupFuncs      []func()
		routeHost         string
		internalRouteHost string
	)

	BeforeEach(func() {
		mockPolicyServer = testhelpers.NewMockPolicyServer()
		policies := []*policy_client.Policy{
			{
				Source: policy_client.Source{
					ID: "capi-process-guid-a",
				},
				Destination: policy_client.Destination{
					ID:       "capi-process-guid-a",
					Protocol: "http",
					Ports: policy_client.Ports{
						Start: 1,
						End:   2,
					},
				},
			},
		}
		mockPolicyServer.SetGetPoliciesResponse(policies)

		mockBBS = testhelpers.NewMockBBSServer()
		mockBBS.SetPostV1EventsResponse(&bbsmodels.ActualLRP{
			ActualLRPKey: bbsmodels.ActualLRPKey{
				ProcessGuid: "diego-process-guid-a",
			},
			State: bbsmodels.ActualLRPStateRunning,
			ActualLRPNetInfo: bbsmodels.ActualLRPNetInfo{
				Address:         "10.10.1.3",
				InstanceAddress: "10.255.1.13",
				Ports: []*bbsmodels.PortMapping{
					{ContainerPort: 8080, HostPort: 61003},
				},
			},
		})

		mockBBS.SetPostV1ActualLRPsList(
			[]*bbsmodels.ActualLRP{
				{
					ActualLRPKey: bbsmodels.ActualLRPKey{
						ProcessGuid: "diego-process-guid-a",
					},
					State:    bbsmodels.ActualLRPStateRunning,
					Presence: bbsmodels.ActualLRP_Ordinary,
					ActualLRPNetInfo: bbsmodels.ActualLRPNetInfo{
						Address:         "10.10.1.3",
						InstanceAddress: "10.255.1.13",
						Ports: []*bbsmodels.PortMapping{
							{ContainerPort: 8080, HostPort: 61003},
						},
					},
				},
				{ // this instance only has SSH port, not app port.  it shouldn't show up in route results
					ActualLRPKey: bbsmodels.NewActualLRPKey("diego-process-guid-a", 1, "domain1"),
					State:        bbsmodels.ActualLRPStateRunning,
					Presence:     bbsmodels.ActualLRP_Ordinary,
					ActualLRPNetInfo: bbsmodels.ActualLRPNetInfo{
						Address:         "10.10.1.4",
						InstanceAddress: "10.255.1.15",
						Ports: []*bbsmodels.PortMapping{
							{ContainerPort: 2222, HostPort: 61004},
						},
					},
				},
				{
					ActualLRPKey: bbsmodels.NewActualLRPKey("diego-process-guid-a", 1, "domain1"),
					State:        bbsmodels.ActualLRPStateRunning,
					Presence:     bbsmodels.ActualLRP_Ordinary,
					ActualLRPNetInfo: bbsmodels.ActualLRPNetInfo{
						Address:         "10.10.1.5",
						InstanceAddress: "10.255.1.16",
						Ports: []*bbsmodels.PortMapping{
							{ContainerPort: 8080, HostPort: 61005},
						},
					},
				},
				{
					ActualLRPKey: bbsmodels.NewActualLRPKey("diego-process-guid-b", 1, "domain1"),
					State:        bbsmodels.ActualLRPStateRunning,
					Presence:     bbsmodels.ActualLRP_Ordinary,
					ActualLRPNetInfo: bbsmodels.ActualLRPNetInfo{
						Address:         "10.10.1.6",
						InstanceAddress: "10.255.0.34",
						Ports: []*bbsmodels.PortMapping{
							{ContainerPort: 2222, HostPort: 61008},
							{ContainerPort: 8080, HostPort: 61006},
						},
					},
				},
				{
					ActualLRPKey: bbsmodels.NewActualLRPKey("diego-process-guid-other", 1, "domain1"),
					State:        bbsmodels.ActualLRPStateRunning,
					Presence:     bbsmodels.ActualLRP_Ordinary,
					ActualLRPNetInfo: bbsmodels.ActualLRPNetInfo{
						Address:         "10.10.1.7",
						InstanceAddress: "10.255.0.35",
						Ports: []*bbsmodels.PortMapping{
							{ContainerPort: 8080, HostPort: 61111},
						},
					},
				},
			})

		copilotCreds := testhelpers.GenerateMTLS()
		cleanupFuncs = append(cleanupFuncs, copilotCreds.CleanupTempFiles)
		listenAddrForCloudController := fmt.Sprintf("127.0.0.1:%d", testhelpers.PickAPort())
		listenAddrForVIPResolver := fmt.Sprintf("127.0.0.1:%d", testhelpers.PickAPort())
		listenAddrForMCP := fmt.Sprintf("127.0.0.1:%d", testhelpers.PickAPort())
		copilotTLSFiles := copilotCreds.CreateServerTLSFiles()
		bbsCreds := testhelpers.GenerateMTLS()
		bbsTLSFiles := bbsCreds.CreateClientTLSFiles()
		mockBBS.Server.HTTPTestServer.TLS = bbsCreds.ServerTLSConfig()

		mockBBS.Server.HTTPTestServer.StartTLS()
		cleanupFuncs = append(cleanupFuncs, mockBBS.Server.Close)

		policyServerCreds := testhelpers.GenerateMTLS()
		policyServerTLSFiles := policyServerCreds.CreateClientTLSFiles()
		mockPolicyServer.Server.HTTPTestServer.TLS = policyServerCreds.ServerTLSConfig()
		mockPolicyServer.Server.HTTPTestServer.StartTLS()

		serverConfig = &config.Config{
			ListenAddressForCloudController: listenAddrForCloudController,
			ListenAddressForVIPResolver:     listenAddrForVIPResolver,
			ListenAddressForMCP:             listenAddrForMCP,
			PilotClientCAPath:               copilotTLSFiles.ClientCA,
			CloudControllerClientCAPath:     copilotTLSFiles.OtherClientCA,
			ServerCertPath:                  copilotTLSFiles.ServerCert,
			ServerKeyPath:                   copilotTLSFiles.ServerKey,
			VIPCIDR:                         "127.128.0.0/9",
			MCPConvergeInterval:             durationjson.Duration(10 * time.Millisecond),
			LogLevel:                        "info",
			BBS: &config.BBSConfig{
				ServerCACertPath: bbsTLSFiles.ServerCA,
				ClientCertPath:   bbsTLSFiles.ClientCert,
				ClientKeyPath:    bbsTLSFiles.ClientKey,
				Address:          mockBBS.Server.URL(),
				SyncInterval:     durationjson.Duration(10 * time.Millisecond),
			},
			TLSPems: []certs.CertChainKeyPair{
				{
					CertChain:  "-----BEGIN CERTIFICATE-----\nMIIC/DCCAeSgAwIBAgIRAPf9lECQDqNwfP1KpPxMqmIwDQYJKoZIhvcNAQELBQAw\nEjEQMA4GA1UEChMHQWNtZSBDbzAeFw0xODExMTQxODU0MThaFw0xOTExMTQxODU0\nMThaMBIxEDAOBgNVBAoTB0FjbWUgQ28wggEiMA0GCSqGSIb3DQEBAQUAA4IBDwAw\nggEKAoIBAQCgDRgg0xFS7Hw4yN/EMTVYp9+My7R6mZ/s7qYVEcrKVnLYZAAUsXsA\nLG1BVeTWfQSPvshi1EP4SAsRpZ8sO/o3GybVfm5ejBVOC0seA1zm2LHMwPyjeIXU\neM/7S3VdBkve+37vj78uZe149Jj+IkLL3zkfRtI+coG9mw4FpP0TqaRQ41cKqnQS\nD2iRbSfBW/nMRcFQr7aK+z+LQg6LPez7CxCsdXgcMf8kNVdceQSatEFnufnK/Gyy\nDs+P2ovlqLpVC05SsO/dTQp+QtVYMNeCA/eLixNzwfiCXhDZ993JFUWj3TkCr7f6\nBY5U/2naXAGS8ZZVzXlweX2SO0BYicPNAgMBAAGjTTBLMA4GA1UdDwEB/wQEAwIF\noDATBgNVHSUEDDAKBggrBgEFBQcDATAMBgNVHRMBAf8EAjAAMBYGA1UdEQQPMA2C\nC2V4YW1wbGUuY29tMA0GCSqGSIb3DQEBCwUAA4IBAQAbgykDDrDA00rKNx/B4G2j\nAeDAHkAnMK5IjdrgH2KUNeI07eRkLhYobrquwcKRYa9RJcM/eImX8BkviwjlOkDz\noJdU0LMVrsrjBuwj9qYg+D7IywvPrrrdrjgF05BxUfwoH1lTKm8Q9SVnpEWEdJj9\n+sP10reX+O7L4xiqgyuKHjWPEK4NJD+Wsw5n8UvIq5LVvTt3bLWsgN2Mole0lJb4\nvgR4N1absHZN6/yju5s7cY0lLBcEitJNUQeW3lHSOWXJ8xiw9aayFnJ4tmNWpALU\nvqettFbN38gfsH8JElHwyeKLthGL/Kj1Cvb//SbK30RnG8vY8kuqKKs3iwuHB+UU\n-----END CERTIFICATE-----\n",
					PrivateKey: "-----BEGIN RSA PRIVATE KEY-----\nMIIEogIBAAKCAQEAoA0YINMRUux8OMjfxDE1WKffjMu0epmf7O6mFRHKylZy2GQA\nFLF7ACxtQVXk1n0Ej77IYtRD+EgLEaWfLDv6Nxsm1X5uXowVTgtLHgNc5tixzMD8\no3iF1HjP+0t1XQZL3vt+74+/LmXtePSY/iJCy985H0bSPnKBvZsOBaT9E6mkUONX\nCqp0Eg9okW0nwVv5zEXBUK+2ivs/i0IOiz3s+wsQrHV4HDH/JDVXXHkEmrRBZ7n5\nyvxssg7Pj9qL5ai6VQtOUrDv3U0KfkLVWDDXggP3i4sTc8H4gl4Q2ffdyRVFo905\nAq+3+gWOVP9p2lwBkvGWVc15cHl9kjtAWInDzQIDAQABAoIBABUV5InehLfCBBOP\nEzvLp9WIOEFaTOqh9pnGTwcTkv3ZKcQsWH5ha2z4bWRgJofDbKhrYAb1JAc/poWq\npi+zryE3aIRT5cJ6/guMHVdU5hZbkgEBo8b9h9QYHn5i0JFy1OgJhg2ViIBaWVDI\nGKfSZ65oOCRQtj4X49PQ66X+uICwcWhJ3tZnFVODPQU6uDaUZsJzESTaEYaTEkpH\nKCbYdKL4dqt76SIxzKwy1tQlV7R/5Vl5iGhIq143iqNVEAHnCDzJZyonoFpvzT3A\nKfxYjwbatzDdDDujlzyUEwdzy+ZSkMtb/b2Asd0QseY4LgsjnkyQKTtuemjxLw7F\nrMbD3ZkCgYEAx5hWLceS1li3h4UVQFPEdqLAyBW5xTGAVKP4dEirlDWWlkVNpPSw\nD/ZAMieL7WT40JExsGYovrtly9BgyOkTbhbs3dlTDsd2++2/gycjEiNIw08Q5F0v\nz0TgV5psUb1E2Mvubf+Ns04C/NwXHX+A8ClcHuVy/qw/y92s+r+H2wcCgYEAzUf2\nFu4d+CO2JqcvPY2YikDFNT6pIzO/Ux0W41FwJVDHRa+42vqW5qPrr6ThWoVEEs5h\nzeBgh0X6K+2AbELDm3kxW43ceHo6KmPCPyQcMxff+A8LyxZWxn/8wb4CKso2zc1L\ncm6w5E0NsCmt/4WP5EeIIUmUXIpcNP9uNCZ6sYsCgYAn1q802gXkBLc1NIoGWfH3\n4ApspXF7+6JqwoO/6hVdMskI23Jg/3n45aTwndYfHy1Oq/xoAiwVzd/Gq6P11hfL\nvIWwzkT2yTdll5HHQtOMNkC6wxhTDIqTa2L/+VGviwCn6SSBDiYhaOvNvrxaZe29\ngfPiMtgeHxFoxqlVL0+VlwKBgApa+PULKgPceVHV2TI3tFw1DD21XX7jG2Gr8/2f\nnBKl0oeXZ7HUNkyINFl17dBNLLPuKUzjZrssMoSIxJOxgoCTSoQd0eNZ9xkwUxow\nTiPdrnSq/aNPCy2UQ0Have0+qikTlBy/rLi3klsynw5mxG11lk5nkc5hRGmAASUs\nU8AlAoGAVMzVMvNOC5Q3uiji0HRnZRa9XrOFdLGZjIUqtLAEdCyGG2Q1WBy2aVgX\nHb/NjnkfmroOSCKUOyqFt0N3sHAv65E5rUdY46uyfczyaQ4wjEhxHPCID6aQc/4f\npBr58YgMa/6k4d3H6arh4cXXPZ16r2gxOcwrVeHecxGpfSAtzBg=\n-----END RSA PRIVATE KEY-----\n",
				},
			},
			PolicyServerAddress:        mockPolicyServer.Server.URL(),
			PolicyServerClientCertPath: policyServerTLSFiles.ClientCert,
			PolicyServerClientKeyPath:  policyServerTLSFiles.ClientKey,
			PolicyServerCAPath:         policyServerTLSFiles.ServerCA,
		}
		fmt.Printf("%+v\n", serverConfig)
		configFilePath = testhelpers.TempFileName()
		cleanupFuncs = append(cleanupFuncs, func() { os.Remove(configFilePath) })

		Expect(serverConfig.Save(configFilePath)).To(Succeed())

		cmd := exec.Command(binaryPath, "-config", configFilePath)
		var err error
		session, err = gexec.Start(cmd, GinkgoWriter, GinkgoWriter)
		Expect(err).NotTo(HaveOccurred())
		Eventually(session.Out).Should(gbytes.Say(`started`))

		pilotClientTLSConfig = copilotCreds.ClientTLSConfig()
		cloudControllerClientTLSConfig = copilotCreds.OtherClientTLSConfig()

		ccClient, err = copilot.NewCloudControllerClient(serverConfig.ListenAddressForCloudController, cloudControllerClientTLSConfig)
		Expect(err).NotTo(HaveOccurred())
		vipResolverClient, err = copilot.NewVIPResolverCopilotClient(serverConfig.ListenAddressForVIPResolver, grpc.WithInsecure())
		Expect(err).NotTo(HaveOccurred())
		mcpClient, err = testhelpers.NewMockPilotMCPClient(pilotClientTLSConfig, serverConfig.ListenAddressForMCP)
		Expect(err).NotTo(HaveOccurred())
	})

	AfterEach(func() {
		err := mcpClient.Close()
		Expect(err).To(BeNil())
		session.Interrupt()
		Eventually(session, "10s").Should(gexec.Exit())

		for i := len(cleanupFuncs) - 1; i >= 0; i-- {
			cleanupFuncs[i]()
		}
	})

	Context("Bulk Sync", func() {
		It("Receives chunks and returns the number of bytes synced", func() {
			client, err := ccClient.BulkSync(context.Background())
			Expect(err).NotTo(HaveOccurred())

			By("Sending the first chunk")
			firstBulkReq := bulkSyncRequest(routeHost)
			data, err := proto.Marshal(firstBulkReq)
			Expect(err).NotTo(HaveOccurred())

			err = client.Send(&api.BulkSyncRequestChunk{Chunk: data})
			Expect(err).NotTo(HaveOccurred())

			By("Sending the second chunk")
			secondBulkReq := bulkSyncRequest(routeHost)
			data2, err := proto.Marshal(secondBulkReq)
			Expect(err).NotTo(HaveOccurred())

			err = client.Send(&api.BulkSyncRequestChunk{Chunk: data2})
			Expect(err).NotTo(HaveOccurred())

			resp, err := client.CloseAndRecv()
			Expect(err).NotTo(HaveOccurred())
			totalSent := len(data) + len(data2)
			Expect(resp.TotalBytesReceived).To(Equal(int32(totalSent)))
		})
	})

	Context("vip resolver client queries a route's vip", func() {
		BeforeEach(func() {
			WaitForHealthy(ccClient)
			WaitForHealthy(vipResolverClient)

			routeHost = "amelia.apps.internal"
			_, err := ccClient.UpsertRoute(context.Background(), &api.UpsertRouteRequest{
				Route: &api.Route{
					Guid: "route-guid-a",
					Host: routeHost,
					Vip:  "127.0.0.76",
				}})
			Expect(err).NotTo(HaveOccurred())
		})

		It("returns the route's vip", func() {
			resp, err := vipResolverClient.GetVIPByName(
				context.Background(),
				&api.GetVIPByNameRequest{
					Fqdn: "amelia.apps.internal",
				},
			)
			Expect(err).NotTo(HaveOccurred())
			Expect(resp.Ip).To(Equal("127.0.0.76"))
		})
	})

	Specify("a journey", func() {
		WaitForHealthy(ccClient)

		By("CC creates and maps a route")
		routeHost = "some-url"
		_, err := ccClient.UpsertRoute(context.Background(), &api.UpsertRouteRequest{
			Route: &api.Route{
				Guid: "route-guid-a",
				Host: routeHost,
			}})
		Expect(err).NotTo(HaveOccurred())

		_, err = ccClient.MapRoute(context.Background(), &api.MapRouteRequest{
			RouteMapping: &api.RouteMapping{
				RouteGuid:       "route-guid-a",
				CapiProcessGuid: "capi-process-guid-a",
				RouteWeight:     1,
			},
		})
		Expect(err).NotTo(HaveOccurred())

		_, err = ccClient.UpsertCapiDiegoProcessAssociation(context.Background(), &api.UpsertCapiDiegoProcessAssociationRequest{
			CapiDiegoProcessAssociation: &api.CapiDiegoProcessAssociation{
				CapiProcessGuid:   "capi-process-guid-a",
				DiegoProcessGuids: []string{"diego-process-guid-a"},
			},
		})
		Expect(err).NotTo(HaveOccurred())

		By("CC creates and maps an internal route")
		internalRouteHost = "some-internal-url"
		_, err = ccClient.UpsertRoute(context.Background(), &api.UpsertRouteRequest{
			Route: &api.Route{
				Guid:     "internal-route-guid-a",
				Host:     internalRouteHost,
				Internal: true,
				Vip:      "127.0.0.44",
			}})
		Expect(err).NotTo(HaveOccurred())

		_, err = ccClient.MapRoute(context.Background(), &api.MapRouteRequest{
			RouteMapping: &api.RouteMapping{
				RouteGuid:       "internal-route-guid-a",
				CapiProcessGuid: "capi-process-guid-a",
				RouteWeight:     1,
			},
		})
		Expect(err).NotTo(HaveOccurred())

		_, err = ccClient.UpsertCapiDiegoProcessAssociation(context.Background(), &api.UpsertCapiDiegoProcessAssociationRequest{
			CapiDiegoProcessAssociation: &api.CapiDiegoProcessAssociation{
				CapiProcessGuid:   "capi-process-guid-a",
				DiegoProcessGuids: []string{"diego-process-guid-a"},
			},
		})
		Expect(err).NotTo(HaveOccurred())

		By("istio pilot MCP client sees the correct messages and objects")
		Eventually(mcpClient.GetAllMessageNames, "1s").Should(ConsistOf(
			"istio/networking/v1alpha3/destinationrules",
			"istio/networking/v1alpha3/virtualservices",
			"istio/networking/v1alpha3/serviceentries",
			"istio/networking/v1alpha3/gateways",
			"istio/networking/v1alpha3/sidecars",
			"istio/networking/v1alpha3/envoyfilters",
			"istio/rbac/v1alpha1/rbacconfigs",
			"istio/rbac/v1alpha1/authorizationpolicies",
			"istio/authentication/v1alpha1/policies",
			"istio/authentication/v1alpha1/meshpolicies",
			"istio/mixer/v1/config/client/quotaspecbindings",
			"istio/mixer/v1/config/client/quotaspecs",
			"istio/config/v1alpha2/httpapispecs",
			"istio/config/v1alpha2/httpapispecbindings",
			"istio/rbac/v1alpha1/servicerolebindings",
			"istio/rbac/v1alpha1/serviceroles",
		))

		Eventually(mcpClient.GetAllObjectNames, "1s").Should(Equal(map[string][]string{
			"istio/networking/v1alpha3/destinationrules": []string{fmt.Sprintf("copilot-rule-for-%s", routeHost), fmt.Sprintf("internal/copilot-rule-for-%s", internalRouteHost)},
			"istio/networking/v1alpha3/virtualservices":  []string{fmt.Sprintf("copilot-service-for-%s", routeHost), fmt.Sprintf("internal/copilot-service-for-%s", internalRouteHost)},
			"istio/networking/v1alpha3/serviceentries":   []string{fmt.Sprintf("copilot-service-entry-for-%s", routeHost), fmt.Sprintf("internal/copilot-service-entry-for-%s", internalRouteHost)},
			"istio/networking/v1alpha3/gateways":         []string{copilotsnapshot.DefaultGatewayName},
			"istio/networking/v1alpha3/sidecars":         []string{copilotsnapshot.DefaultSidecarName, "capi-process-guid-a"},
		}))

		expectedRoutes := []Route{
			{
				dest: generateDestination([]RouteDestination{
					{
						port:   8080,
						weight: 100,
						subset: "capi-process-guid-a",
						host:   "some-url",
					},
				}),
			},
		}
		expectedRoutesInternal := []Route{
			{
				dest: generateDestination([]RouteDestination{
					{
						port:   8080,
						weight: 100,
						subset: "capi-process-guid-a",
						host:   "some-internal-url",
					},
				}),
			},
		}
		expectedVS := expectedVirtualService("some-url", "cloudfoundry-ingress", expectedRoutes)
		expectedInternalVS := expectedVirtualServiceWithRetries("some-internal-url", "", expectedRoutesInternal)
		Eventually(mcpClient.GetAllVirtualServices, "1s").Should(ConsistOf(expectedVS, expectedInternalVS))

		expectedDR := expectedDestinationRule("some-url", []string{"capi-process-guid-a"})
		expectedInternalDR := expectedDestinationRule("some-internal-url", []string{"capi-process-guid-a"})
		Eventually(mcpClient.GetAllDestinationRules, "1s").Should(ConsistOf([]*v1alpha3.DestinationRule{expectedDR, expectedInternalDR}))

		expectedGW := expectedGateway(80)
		Eventually(mcpClient.GetAllGateways, "1s").Should(Equal([]*v1alpha3.Gateway{expectedGW}))

		expectedSidecar := expectedSidecarResource("capi-process-guid-a", []string{"internal/some-internal-url"})
		Eventually(mcpClient.GetAllSidecars, "1s").Should(ContainElement(expectedSidecar))

		expectedSE := expectedServiceEntry(
			"some-url",
			"",
			"http",
			[]Endpoint{
				{
					port:   61003,
					addr:   "10.10.1.3",
					subset: "capi-process-guid-a",
				},
				{
					port:   61005,
					addr:   "10.10.1.5",
					subset: "capi-process-guid-a",
				},
			},
		)
		expectedInternalSE := expectedServiceEntry(
			"some-internal-url",
			"127.0.0.44",
			"http",
			[]Endpoint{
				{
					port:   8080,
					addr:   "10.255.1.13",
					subset: "capi-process-guid-a",
				},
				{
					port:   8080,
					addr:   "10.255.1.16",
					subset: "capi-process-guid-a",
				},
			},
		)
		Eventually(mcpClient.GetAllServiceEntries, "1s").Should(ConsistOf(expectedSE, expectedInternalSE))

		By("cc maps another backend to the same route")
		_, err = ccClient.MapRoute(context.Background(), &api.MapRouteRequest{
			RouteMapping: &api.RouteMapping{
				RouteGuid:       "route-guid-a",
				CapiProcessGuid: "capi-process-guid-b",
				RouteWeight:     1,
			},
		})
		Expect(err).NotTo(HaveOccurred())

		_, err = ccClient.UpsertCapiDiegoProcessAssociation(context.Background(), &api.UpsertCapiDiegoProcessAssociationRequest{
			CapiDiegoProcessAssociation: &api.CapiDiegoProcessAssociation{
				CapiProcessGuid: "capi-process-guid-b",
				DiegoProcessGuids: []string{
					"diego-process-guid-b",
				},
			},
		})
		Expect(err).NotTo(HaveOccurred())

		By("cc adds a second route and maps it to the second backend")
		_, err = ccClient.UpsertRoute(context.Background(), &api.UpsertRouteRequest{
			Route: &api.Route{
				Guid: "route-guid-b",
				Host: "some-other-url",
				Path: "/some/path",
			}})
		Expect(err).NotTo(HaveOccurred())

		_, err = ccClient.MapRoute(context.Background(), &api.MapRouteRequest{
			RouteMapping: &api.RouteMapping{
				RouteGuid:       "route-guid-b",
				CapiProcessGuid: "capi-process-guid-other",
				RouteWeight:     1,
			},
		})
		Expect(err).NotTo(HaveOccurred())

		_, err = ccClient.UpsertCapiDiegoProcessAssociation(context.Background(), &api.UpsertCapiDiegoProcessAssociationRequest{
			CapiDiegoProcessAssociation: &api.CapiDiegoProcessAssociation{
				CapiProcessGuid: "capi-process-guid-other",
				DiegoProcessGuids: []string{
					"diego-process-guid-other",
				},
			},
		})
		Expect(err).NotTo(HaveOccurred())

		By("istio mcp client sees both routes and their respective backends")
		expectedSE = expectedServiceEntry(
			"some-url",
			"",
			"http",
			[]Endpoint{
				{
					port:   61003,
					addr:   "10.10.1.3",
					subset: "capi-process-guid-a",
				},
				{
					port:   61005,
					addr:   "10.10.1.5",
					subset: "capi-process-guid-a",
				},
				{
					port:   61006,
					addr:   "10.10.1.6",
					subset: "capi-process-guid-b",
				},
			},
		)
		expectedOtherSE := expectedServiceEntry(
			"some-other-url",
			"",
			"http",
			[]Endpoint{
				{
					port:   61111,
					addr:   "10.10.1.7",
					subset: "capi-process-guid-other",
				},
			},
		)
		Eventually(mcpClient.GetAllServiceEntries, "1s").Should(ConsistOf(expectedSE, expectedOtherSE, expectedInternalSE))

		expectedRoutes = []Route{
			{
				dest: generateDestination([]RouteDestination{
					{
						port:   8080,
						weight: 50,
						subset: "capi-process-guid-a",
						host:   "some-url",
					},
					{
						port:   8080,
						weight: 50,
						subset: "capi-process-guid-b",
						host:   "some-url",
					},
				}),
			},
		}
		expectedOtherRoutes := []Route{
			{
				dest: generateDestination([]RouteDestination{
					{
						port:   8080,
						weight: 100,
						subset: "capi-process-guid-other",
						host:   "some-other-url",
					},
				}),
				match: generateMatch([]string{"/some/path"}),
			},
		}
		expectedVS = expectedVirtualService("some-url", "cloudfoundry-ingress", expectedRoutes)
		expectedOtherVS := expectedVirtualService("some-other-url", "cloudfoundry-ingress", expectedOtherRoutes)
		Eventually(mcpClient.GetAllVirtualServices, "1s").Should(ConsistOf(
			expectedVS,
			expectedOtherVS,
			expectedInternalVS,
		))

		expectedDR = expectedDestinationRule("some-url",
			[]string{"capi-process-guid-a", "capi-process-guid-b"})
		expectedOtherDR := expectedDestinationRule("some-other-url",
			[]string{"capi-process-guid-other"})
		Eventually(mcpClient.GetAllDestinationRules, "1s").Should(ConsistOf([]*v1alpha3.DestinationRule{expectedDR, expectedOtherDR, expectedInternalDR}))

		expectedGW = expectedGateway(80)
		Eventually(mcpClient.GetAllGateways, "1s").Should(Equal([]*v1alpha3.Gateway{expectedGW}))

		By("cc unmaps the first backend from the first route")
		_, err = ccClient.UnmapRoute(context.Background(), &api.UnmapRouteRequest{RouteMapping: &api.RouteMapping{
			RouteGuid:       "route-guid-a",
			CapiProcessGuid: "capi-process-guid-a",
			RouteWeight:     1,
		}})
		Expect(err).NotTo(HaveOccurred())

		By("cc deletes the second route")
		_, err = ccClient.DeleteRoute(context.Background(), &api.DeleteRouteRequest{
			Guid: "route-guid-b",
		})
		Expect(err).NotTo(HaveOccurred())

		By("istio mcp client sees the updated stuff")
		expectedSE = expectedServiceEntry(
			"some-url",
			"",
			"http",
			[]Endpoint{
				{
					port:   61006,
					addr:   "10.10.1.6",
					subset: "capi-process-guid-b",
				},
			},
		)
		Eventually(mcpClient.GetAllServiceEntries, "1s").Should(ConsistOf(expectedSE, expectedInternalSE))

		expectedRoutes = []Route{
			{
				dest: generateDestination([]RouteDestination{
					{
						port:   8080,
						weight: 100,
						subset: "capi-process-guid-b",
						host:   "some-url",
					},
				}),
			},
		}
		expectedVS = expectedVirtualService("some-url", "cloudfoundry-ingress", expectedRoutes)
		Eventually(mcpClient.GetAllVirtualServices, "3s").Should(ConsistOf(
			expectedVS,
			expectedInternalVS,
		))

		expectedDR = expectedDestinationRule("some-url",
			[]string{"capi-process-guid-b"})
		Eventually(mcpClient.GetAllDestinationRules, "1s").Should(ConsistOf([]*v1alpha3.DestinationRule{expectedDR, expectedInternalDR}))

		expectedGW = expectedGateway(80)
		Eventually(mcpClient.GetAllGateways, "1s").Should(Equal([]*v1alpha3.Gateway{expectedGW}))
	})

	Context("when the BBS is not available", func() {
		BeforeEach(func() {
			mockBBS.Server.Close()

			// stop copilot
			gexec.KillAndWait(time.Second * 10)
			Eventually(session, "2s").Should(gexec.Exit())
		})

		It("crashes and prints a useful error log", func() {
			// re-start copilot
			cmd := exec.Command(binaryPath, "-config", configFilePath)
			var err error
			session, err = gexec.Start(cmd, GinkgoWriter, GinkgoWriter)
			Expect(err).NotTo(HaveOccurred())

			Eventually(session, "2s").Should(gexec.Exit(1))
			Eventually(session.Out).Should(gbytes.Say(`unable to reach BBS`))
		})

		Context("but if the user sets config BBS.Disable", func() {
			BeforeEach(func() {
				serverConfig.BBS.Disable = true
				Expect(serverConfig.Save(configFilePath)).To(Succeed())
			})

			It("boots successfully and serves requests on the Cloud Controller-facing server", func() {
				cmd := exec.Command(binaryPath, "-config", configFilePath)
				var err error
				session, err = gexec.Start(cmd, GinkgoWriter, GinkgoWriter)
				Expect(err).NotTo(HaveOccurred())
				Eventually(session.Out).Should(gbytes.Say(`BBS support is disabled`))

				WaitForHealthy(ccClient)
				_, err = ccClient.UpsertRoute(context.Background(), &api.UpsertRouteRequest{
					Route: &api.Route{
						Guid: "route-guid-xyz",
						Host: "some-url",
					}})
				Expect(err).NotTo(HaveOccurred())
			})
		})
	})
})

type healthable interface {
	Health(ctx context.Context, in *api.HealthRequest, opts ...grpc.CallOption) (*api.HealthResponse, error)
}

func WaitForHealthy(client healthable) {
	By("waiting for the server to become healthy")
	serverForCloudControllerIsHealthy := func() error {
		ctx, cancelFunc := context.WithTimeout(context.Background(), 100*time.Millisecond)
		defer cancelFunc()
		_, err := client.Health(ctx, new(api.HealthRequest))
		return err
	}
	Eventually(serverForCloudControllerIsHealthy, 2*time.Second).Should(Succeed())
}

type Route struct {
	match []*v1alpha3.HTTPMatchRequest
	dest  []*v1alpha3.HTTPRouteDestination
}

type RouteDestination struct {
	host   string
	port   uint32
	weight int32
	subset string
}

func generateMatch(paths []string) []*v1alpha3.HTTPMatchRequest {
	matches := []*v1alpha3.HTTPMatchRequest{}

	for _, p := range paths {
		matches = append(matches, &v1alpha3.HTTPMatchRequest{
			Uri: &v1alpha3.StringMatch{
				MatchType: &v1alpha3.StringMatch_Prefix{Prefix: p},
			},
			Scheme:       nil,
			Method:       nil,
			Authority:    nil,
			Headers:      nil,
			Port:         0,
			SourceLabels: nil,
			Gateways:     nil,
		})
	}

	return matches
}

func generateDestination(dests []RouteDestination) []*v1alpha3.HTTPRouteDestination {
	newDests := []*v1alpha3.HTTPRouteDestination{}

	for _, d := range dests {
		newDests = append(newDests, &v1alpha3.HTTPRouteDestination{
			Destination: &v1alpha3.Destination{
				Host:   d.host,
				Subset: d.subset,
				Port: &v1alpha3.PortSelector{
					Port: &v1alpha3.PortSelector_Number{Number: d.port},
				},
			},
			Weight:                d.weight,
			RemoveResponseHeaders: nil,
			AppendResponseHeaders: nil,
			RemoveRequestHeaders:  nil,
			AppendRequestHeaders:  nil,
		})
	}

	return newDests
}

func expectedVirtualService(host, gateway string, routes []Route) *v1alpha3.VirtualService {
	newRoutes := []*v1alpha3.HTTPRoute{}
	for _, r := range routes {
		newRoutes = append(newRoutes, &v1alpha3.HTTPRoute{
			Match:                 r.match,
			Route:                 r.dest,
			Redirect:              nil,
			Rewrite:               nil,
			WebsocketUpgrade:      false,
			Timeout:               nil,
			Retries:               nil,
			Fault:                 nil,
			Mirror:                nil,
			CorsPolicy:            nil,
			AppendHeaders:         nil,
			RemoveResponseHeaders: nil,
			AppendResponseHeaders: nil,
			RemoveRequestHeaders:  nil,
			AppendRequestHeaders:  nil,
		})
	}

	var gateways []string
	if gateway != "" {
		gateways = []string{gateway}
	}
	return &v1alpha3.VirtualService{
		Hosts:    []string{host},
		Gateways: gateways,
		Tls:      nil,
		Tcp:      nil,
		Http:     newRoutes,
	}
}

func expectedVirtualServiceWithRetries(host, gateway string, routes []Route) *v1alpha3.VirtualService {
	newRoutes := []*v1alpha3.HTTPRoute{}
	for _, r := range routes {
		newRoutes = append(newRoutes, &v1alpha3.HTTPRoute{
			Match:            r.match,
			Route:            r.dest,
			Redirect:         nil,
			Rewrite:          nil,
			WebsocketUpgrade: false,
			Timeout:          types.DurationProto(15 * time.Second),
			Retries: &v1alpha3.HTTPRetry{
				Attempts: 3,
				RetryOn:  "5xx",
			},
			Fault:                 nil,
			Mirror:                nil,
			CorsPolicy:            nil,
			AppendHeaders:         nil,
			RemoveResponseHeaders: nil,
			AppendResponseHeaders: nil,
			RemoveRequestHeaders:  nil,
			AppendRequestHeaders:  nil,
		})
	}

	var gateways []string
	if gateway != "" {
		gateways = []string{gateway}
	}
	return &v1alpha3.VirtualService{
		Hosts:    []string{host},
		Gateways: gateways,
		Tls:      nil,
		Tcp:      nil,
		Http:     newRoutes,
	}
}

func expectedDestinationRule(host string, subsets []string) *v1alpha3.DestinationRule {
	sets := []*v1alpha3.Subset{}
	for _, s := range subsets {
		sets = append(sets, &v1alpha3.Subset{
			Name: s,
			Labels: map[string]string{
				"cfapp": s,
			},
			TrafficPolicy: nil,
		})
	}

	return &v1alpha3.DestinationRule{
		Host:          host,
		TrafficPolicy: nil,
		Subsets:       sets,
	}
}

func expectedSidecarResource(sourceAppGUID string, destinations []string) *v1alpha3.Sidecar {
	return &v1alpha3.Sidecar{
		WorkloadSelector: &v1alpha3.WorkloadSelector{
			Labels: map[string]string{
				"cfapp": sourceAppGUID,
			},
		},
		Egress: []*v1alpha3.IstioEgressListener{
			{
				Hosts: destinations,
			},
		},
	}
}

func expectedGateway(port uint32) *v1alpha3.Gateway {
	return &v1alpha3.Gateway{
		Servers: []*v1alpha3.Server{
			{
				Port:  &v1alpha3.Port{Number: port, Protocol: "http", Name: "http"},
				Hosts: []string{"*"},
				Tls:   nil,
			},
			{
				Port: &v1alpha3.Port{
					Number:   443,
					Protocol: "https",
					Name:     "example.com",
				},
				Tls: &v1alpha3.Server_TLSOptions{
					Mode:              v1alpha3.Server_TLSOptions_SIMPLE,
					ServerCertificate: "/etc/istio/example.com/tls.crt",
					PrivateKey:        "/etc/istio/example.com/tls.key",
				},
				Hosts: []string{"example.com"},
			},
		},
		Selector: nil,
	}
}

type Endpoint struct {
	addr   string
	port   uint32
	subset string
}

func expectedServiceEntry(host, address, protocol string, newEndpoints []Endpoint) *v1alpha3.ServiceEntry {
	endpoints := []*v1alpha3.ServiceEntry_Endpoint{}
	for i, _ := range newEndpoints {
		endpoints = append(endpoints, &v1alpha3.ServiceEntry_Endpoint{
			Address: newEndpoints[i].addr,
			Ports:   map[string]uint32{protocol: newEndpoints[i].port},
			Labels: map[string]string{
				"cfapp": newEndpoints[i].subset,
			},
		})
	}

	var addresses []string
	if address != "" {
		addresses = []string{address}
	}

	return &v1alpha3.ServiceEntry{
		Hosts:     []string{host},
		Addresses: addresses,
		Ports: []*v1alpha3.Port{
			{Number: 8080, Protocol: protocol, Name: protocol},
		},
		Location:   1,
		Resolution: 1,
		Endpoints:  endpoints,
	}
}

func bulkSyncRequest(routeHost string) *api.BulkSyncRequest {
	randTag := time.Now().UnixNano()
	request := &api.BulkSyncRequest{
		RouteMappings: []*api.RouteMapping{
			{
				RouteGuid:       fmt.Sprintf("route-guid-%d", randTag),
				CapiProcessGuid: fmt.Sprintf("capi-process-guid-%d", randTag),
				RouteWeight:     1,
			},
		},
		Routes: []*api.Route{
			{
				Guid: fmt.Sprintf("route-guid-%d", randTag),
				Host: routeHost,
			},
		},
		CapiDiegoProcessAssociations: []*api.CapiDiegoProcessAssociation{
			{
				CapiProcessGuid:   fmt.Sprintf("capi-process-guid-%d", randTag),
				DiegoProcessGuids: []string{fmt.Sprintf("diego-process-guid-%d", randTag)},
			},
		},
	}
	return request
}
