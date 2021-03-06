package models_test

import (
	"code.cloudfoundry.org/copilot/models"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("RoutesRepo", func() {
	var (
		routesRepo *models.RoutesRepo
	)

	BeforeEach(func() {
		routesRepo = models.NewRoutesRepo()
	})

	Describe("Delete", func() {
		It("deletes upsert route", func() {
			route := &models.Route{
				Host: "host.example.com",
				GUID: "some-route-guid",
			}

			go routesRepo.Upsert(route)

			Eventually(func() *models.Route {
				r, _ := routesRepo.Get("some-route-guid")
				return r
			}).Should(Equal(route))

			routesRepo.Delete(route.GUID)

			r, ok := routesRepo.Get("some-route-guid")
			Expect(ok).To(BeFalse())
			Expect(r).To(BeNil())
		})

		Context("when deleting a route that does not exist", func() {
			It("does not return an error", func() {
				route := models.Route{
					Host: "host.example.com",
					GUID: "delete-me",
				}

				routesRepo.Delete(route.GUID)
				routesRepo.Delete(route.GUID)

				_, ok := routesRepo.Get("delete-me")
				Expect(ok).To(BeFalse())
			})
		})
	})

	Describe("Upsert", func() {
		It("updates the same route", func() {
			route := &models.Route{
				Host: "host.example.com",
				GUID: "some-route-guid",
			}

			updatedRoute := &models.Route{
				Host: "something.different.com",
				GUID: route.GUID,
			}

			routesRepo.Upsert(updatedRoute)

			r, _ := routesRepo.Get("some-route-guid")
			Expect(r).To(Equal(updatedRoute))
		})

		It("downcases hosts", func() {
			route := &models.Route{
				Host: "HOST.example.com",
				GUID: "some-route-guid",
			}

			routesRepo.Upsert(route)
			r, _ := routesRepo.Get("some-route-guid")
			Expect(r.Hostname()).To(Equal("host.example.com"))
		})
	})

	Describe("Sync", func() {
		It("saves routes", func() {
			route := &models.Route{
				Host: "host.example.com",
				GUID: "some-route-guid",
			}

			go routesRepo.Upsert(route)

			Eventually(func() *models.Route {
				r, _ := routesRepo.Get("some-route-guid")
				return r
			}).Should(Equal(route))

			newRoute := &models.Route{
				Host: "host.example.com",
				GUID: "some-other-route-guid",
			}

			routesRepo.Sync([]*models.Route{newRoute})
			Expect(routesRepo.List()).To(Equal(map[string]string{
				string(newRoute.GUID): newRoute.Host,
			}))
		})
	})

	Describe("GetVIPByName", func() {
		BeforeEach(func() {
			route := &models.Route{
				Host: "host.example.com",
				GUID: "some-route-guid",
				VIP:  "4.5.7.8",
			}

			routesRepo.Sync([]*models.Route{route})

			otherRoute := &models.Route{
				Host: "other.example.com",
				GUID: "some-other-route-guid",
				VIP:  "4.5.7.9",
			}
			routesRepo.Upsert(otherRoute)
		})

		It("returns the VIP for a route", func() {
			vip, ok := routesRepo.GetVIPByName("host.example.com")
			Expect(vip).To(Equal("4.5.7.8"))
			Expect(ok).To(BeTrue())

			vip, ok = routesRepo.GetVIPByName("other.example.com")
			Expect(vip).To(Equal("4.5.7.9"))
			Expect(ok).To(BeTrue())
		})

		It("returns not ok if the route doesn't exist", func() {
			routesRepo.Delete("some-route-guid")

			_, ok := routesRepo.GetVIPByName("host.example.com")
			Expect(ok).To(BeFalse())
		})
	})
})
