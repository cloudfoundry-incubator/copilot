package models_test

import (
	"code.cloudfoundry.org/copilot/models"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("RouteMappingsRepo", func() {
	var routeMappingsRepo *models.RouteMappingsRepo
	BeforeEach(func() {
		routeMappingsRepo = models.NewRouteMappingsRepo()
	})

	Describe("GetCalculatedWeight", func() {
		It("calculates the weight of a route mapping", func() {
			rmOne := &models.RouteMapping{
				RouteGUID:       "some-route-guid",
				CAPIProcessGUID: "some-capi-guid",
				RouteWeight:     1,
			}
			routeMappingsRepo.Map(rmOne)

			rmTwo := &models.RouteMapping{
				RouteGUID:       "some-route-guid",
				CAPIProcessGUID: "some-other-capi-guid",
				RouteWeight:     2,
			}
			routeMappingsRepo.Map(rmTwo)

			Expect(routeMappingsRepo.GetCalculatedWeight(rmOne)).To(Equal(int32(33)))
			Expect(routeMappingsRepo.GetCalculatedWeight(rmTwo)).To(Equal(int32(66)))
		})

		It("calculates the weight of a route mapping with 3 apps", func() {
			rmOne := &models.RouteMapping{
				RouteGUID:       "some-route-guid",
				CAPIProcessGUID: "some-capi-guid",
				RouteWeight:     1,
			}
			routeMappingsRepo.Map(rmOne)

			rmTwo := &models.RouteMapping{
				RouteGUID:       "some-route-guid",
				CAPIProcessGUID: "some-other-capi-guid",
				RouteWeight:     1,
			}
			routeMappingsRepo.Map(rmTwo)

			rmThree := &models.RouteMapping{
				RouteGUID:       "some-route-guid",
				CAPIProcessGUID: "some-other-other-capi-guid",
				RouteWeight:     1,
			}
			routeMappingsRepo.Map(rmThree)

			Expect(routeMappingsRepo.GetCalculatedWeight(rmOne)).To(Equal(int32(33)))
			Expect(routeMappingsRepo.GetCalculatedWeight(rmTwo)).To(Equal(int32(33)))
			Expect(routeMappingsRepo.GetCalculatedWeight(rmThree)).To(Equal(int32(33)))
		})
	})

	Describe("Key", func() {
		It("is unique for process guid and route guid", func() {
			rmA := models.RouteMapping{
				RouteGUID:       "route-guid-1",
				CAPIProcessGUID: "some-capi-guid-1",
			}

			rmB := models.RouteMapping{
				RouteGUID:       "route-guid-1",
				CAPIProcessGUID: "some-capi-guid-2",
			}

			rmC := models.RouteMapping{
				RouteGUID:       "route-guid-2",
				CAPIProcessGUID: "some-capi-guid-1",
			}

			Expect(rmA.Key()).ToNot(Equal(rmB.Key()))
			Expect(rmA.Key()).ToNot(Equal(rmC.Key()))
			Expect(rmB.Key()).ToNot(Equal(rmC.Key()))
		})
	})

	Describe("Map", func() {
		It("adds a route to the repo", func() {
			routeMapping := models.RouteMapping{
				RouteGUID:       "some-route-guid",
				CAPIProcessGUID: "some-capi-guid",
				RouteWeight:     1,
			}

			go routeMappingsRepo.Map(&routeMapping)

			Eventually(routeMappingsRepo.List).Should(Equal(map[string]*models.RouteMapping{
				"some-route-guid-some-capi-guid": &routeMapping,
			}))
		})

		Context("when submitting the same mapping", func() {
			It("does not add to the repo", func() {
				routeMapping := &models.RouteMapping{
					RouteGUID:       "some-route-guid",
					CAPIProcessGUID: "some-capi-guid",
				}

				routeMappingsRepo.Map(routeMapping)
				routeMappingsRepo.Map(routeMapping)
				routeMappingsRepo.Map(routeMapping)

				Expect(routeMappingsRepo.List()).To(HaveLen(1))
			})
		})
	})

	Describe("Unmap", func() {
		It("removes a route from the repo", func() {
			routeMapping := &models.RouteMapping{
				RouteGUID:       "some-route-guid",
				CAPIProcessGUID: "some-capi-guid",
			}

			routeMappingsRepo.Map(routeMapping)
			routeMappingsRepo.Unmap(routeMapping)

			Expect(routeMappingsRepo.List()).To(HaveLen(0))
		})

		It("zeroes out the denominator", func() {
			routeMapping := &models.RouteMapping{
				RouteGUID:       "some-route-guid",
				CAPIProcessGUID: "some-capi-guid",
				RouteWeight:     2,
			}

			routeMappingsRepo.Map(routeMapping)
			routeMappingsRepo.Unmap(routeMapping)
			routeMappingsRepo.Map(routeMapping)

			Expect(routeMappingsRepo.GetCalculatedWeight(routeMapping)).To(Equal(int32(100)))
		})
	})

	Describe("Sync", func() {
		It("adds a list of routes to the repo", func() {
			routeMapping := &models.RouteMapping{
				RouteGUID:       "some-route-guid",
				CAPIProcessGUID: "some-capi-guid",
				RouteWeight:     2,
			}

			routeMappingsRepo.Map(routeMapping)

			newRouteMapping := &models.RouteMapping{
				RouteGUID:       "some-other-route-guid",
				CAPIProcessGUID: "some-other-capi-guid",
				RouteWeight:     1,
			}

			otherRouteMapping := &models.RouteMapping{
				RouteGUID:       "some-other-route-guid",
				CAPIProcessGUID: "a-different-capi-guid",
				RouteWeight:     2,
			}
			updatedRouteMappings := []*models.RouteMapping{newRouteMapping, otherRouteMapping}

			routeMappingsRepo.Sync(updatedRouteMappings)

			Expect(routeMappingsRepo.List()).Should(Equal(map[string]*models.RouteMapping{
				"some-other-route-guid-some-other-capi-guid":  newRouteMapping,
				"some-other-route-guid-a-different-capi-guid": otherRouteMapping,
			}))

			Expect(routeMappingsRepo.GetCalculatedWeight(otherRouteMapping)).To(Equal(int32(66)))
		})

		It("syncs correctly after multiple times", func() {
			routeMapping := &models.RouteMapping{
				RouteGUID:       "some-route-guid",
				CAPIProcessGUID: "some-capi-guid",
				RouteWeight:     2,
			}

			updatedRouteMappings := []*models.RouteMapping{routeMapping}

			routeMappingsRepo.Sync(updatedRouteMappings)
			routeMappingsRepo.Sync(updatedRouteMappings)

			Expect(routeMappingsRepo.List()).Should(Equal(map[string]*models.RouteMapping{
				"some-route-guid-some-capi-guid": routeMapping,
			}))

			Expect(routeMappingsRepo.GetCalculatedWeight(routeMapping)).To(Equal(int32(100)))
		})
	})
})
