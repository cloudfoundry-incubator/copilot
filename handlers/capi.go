package handlers

import (
	"context"
	"errors"
	"io"

	"code.cloudfoundry.org/copilot/api"
	"code.cloudfoundry.org/copilot/models"
	"code.cloudfoundry.org/lager"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CAPI struct {
	Logger                           lager.Logger
	RoutesRepo                       routesRepoInterface
	RouteMappingsRepo                routeMappingsRepoInterface
	CAPIDiegoProcessAssociationsRepo capiDiegoProcessAssociationsRepoInterface
	VIPProvider                      vipProvider
}

//go:generate counterfeiter -o fakes/vip_provider.go --fake-name VIPProvider . vipProvider
type vipProvider interface {
	Get(hostname string) string
}

//go:generate counterfeiter -o fakes/routes_repo.go --fake-name RoutesRepo . routesRepoInterface
type routesRepoInterface interface {
	Upsert(route *models.Route)
	Delete(guid models.RouteGUID)
	Sync(routes []*models.Route)
	Get(guid models.RouteGUID) (*models.Route, bool)
	List() map[string]string
}

//go:generate counterfeiter -o fakes/route_mappings_repo.go --fake-name RouteMappingsRepo . routeMappingsRepoInterface
type routeMappingsRepoInterface interface {
	GetCalculatedWeight(rm *models.RouteMapping) int32
	Map(routeMapping *models.RouteMapping)
	Unmap(routeMapping *models.RouteMapping)
	Sync(routeMappings []*models.RouteMapping)
	List() map[string]*models.RouteMapping
}

//go:generate counterfeiter -o fakes/capi_diego_process_associations_repo.go --fake-name CAPIDiegoProcessAssociationsRepo . capiDiegoProcessAssociationsRepoInterface
type capiDiegoProcessAssociationsRepoInterface interface {
	Upsert(capiDiegoProcessAssociation *models.CAPIDiegoProcessAssociation)
	Delete(capiProcessGUID *models.CAPIProcessGUID)
	Sync(capiDiegoProcessAssociations []*models.CAPIDiegoProcessAssociation)
	List() map[models.CAPIProcessGUID]*models.DiegoProcessGUIDs
	Get(capiProcessGUID *models.CAPIProcessGUID) *models.CAPIDiegoProcessAssociation
}

func (c *CAPI) Health(context.Context, *api.HealthRequest) (*api.HealthResponse, error) {
	c.Logger.Info("capi health check...")
	return &api.HealthResponse{Healthy: true}, nil
}

// TODO: probably remove or test these eventually, currently using for debugging
func (c *CAPI) ListCfRoutes(context.Context, *api.ListCfRoutesRequest) (*api.ListCfRoutesResponse, error) {
	c.Logger.Info("listing cf routes...")
	return &api.ListCfRoutesResponse{Routes: c.RoutesRepo.List()}, nil
}

// TODO: probably remove or test these eventually, currently using for debugging
func (c *CAPI) ListCfRouteMappings(context.Context, *api.ListCfRouteMappingsRequest) (*api.ListCfRouteMappingsResponse, error) {
	c.Logger.Info("listing cf route mappings...")
	routeMappings := c.RouteMappingsRepo.List()
	apiRoutMappings := make(map[string]*api.RouteMapping)
	for k, v := range routeMappings {
		apiRoutMappings[k] = &api.RouteMapping{
			CapiProcessGuid: string(v.CAPIProcessGUID),
			RouteGuid:       string(v.RouteGUID),
			RouteWeight:     v.RouteWeight,
		}
	}
	return &api.ListCfRouteMappingsResponse{RouteMappings: apiRoutMappings}, nil
}

// TODO: probably remove or test these eventually, currently using for debugging
func (c *CAPI) ListCapiDiegoProcessAssociations(context.Context, *api.ListCapiDiegoProcessAssociationsRequest) (*api.ListCapiDiegoProcessAssociationsResponse, error) {
	c.Logger.Info("listing capi/diego process associations...")

	response := &api.ListCapiDiegoProcessAssociationsResponse{
		CapiDiegoProcessAssociations: make(map[string]*api.DiegoProcessGuids),
	}
	for capiProcessGUID, diegoProcessGUIDs := range c.CAPIDiegoProcessAssociationsRepo.List() {
		response.CapiDiegoProcessAssociations[string(capiProcessGUID)] = &api.DiegoProcessGuids{DiegoProcessGuids: diegoProcessGUIDs.ToStringSlice()}
	}
	return response, nil
}

func (c *CAPI) UpsertRoute(context context.Context, request *api.UpsertRouteRequest) (*api.UpsertRouteResponse, error) {
	c.Logger.Info("upserting route...")
	err := validateUpsertRouteRequest(request)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Route %#v is invalid:\n %v", request, err)
	}
	route := &models.Route{
		GUID:     models.RouteGUID(request.Route.Guid),
		Host:     request.Route.Host,
		Path:     request.Route.Path,
		Internal: request.Route.Internal,
		VIP:      request.Route.Vip,
	}
	c.RoutesRepo.Upsert(route)
	return &api.UpsertRouteResponse{}, nil
}

func (c *CAPI) DeleteRoute(context context.Context, request *api.DeleteRouteRequest) (*api.DeleteRouteResponse, error) {
	c.Logger.Info("deleting route...")
	err := validateDeleteRouteRequest(request)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "%s", err)
	}
	c.RoutesRepo.Delete(models.RouteGUID(request.Guid))
	return &api.DeleteRouteResponse{}, nil
}

func (c *CAPI) MapRoute(context context.Context, request *api.MapRouteRequest) (*api.MapRouteResponse, error) {
	c.Logger.Info("mapping route...")
	err := validateMapRouteRequest(request)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Route Mapping %#v is invalid:\n %v", request, err)
	}
	r := &models.RouteMapping{
		RouteGUID:       models.RouteGUID(request.RouteMapping.RouteGuid),
		CAPIProcessGUID: models.CAPIProcessGUID(request.RouteMapping.CapiProcessGuid),
		RouteWeight:     request.RouteMapping.RouteWeight,
	}
	c.RouteMappingsRepo.Map(r)
	return &api.MapRouteResponse{}, nil
}

func (c *CAPI) UnmapRoute(context context.Context, request *api.UnmapRouteRequest) (*api.UnmapRouteResponse, error) {
	c.Logger.Info("unmapping route...")
	err := validateUnmapRouteRequest(request)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Route Mapping %#v is invalid:\n %v", request, err)
	}
	r := &models.RouteMapping{
		RouteGUID:       models.RouteGUID(request.RouteMapping.RouteGuid),
		CAPIProcessGUID: models.CAPIProcessGUID(request.RouteMapping.CapiProcessGuid),
		RouteWeight:     request.RouteMapping.RouteWeight,
	}
	c.RouteMappingsRepo.Unmap(r)
	return &api.UnmapRouteResponse{}, nil
}

func (c *CAPI) UpsertCapiDiegoProcessAssociation(context context.Context, request *api.UpsertCapiDiegoProcessAssociationRequest) (*api.UpsertCapiDiegoProcessAssociationResponse, error) {
	c.Logger.Info("upserting capi/diego process association...")
	err := validateUpsertCAPIDiegoProcessAssociationRequest(request)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Capi/Diego Process Association %#v is invalid:\n %v", request, err)
	}
	association := &models.CAPIDiegoProcessAssociation{
		CAPIProcessGUID:   models.CAPIProcessGUID(request.CapiDiegoProcessAssociation.CapiProcessGuid),
		DiegoProcessGUIDs: models.DiegoProcessGUIDsFromStringSlice(request.CapiDiegoProcessAssociation.DiegoProcessGuids),
	}
	c.CAPIDiegoProcessAssociationsRepo.Upsert(association)
	return &api.UpsertCapiDiegoProcessAssociationResponse{}, nil
}

func (c *CAPI) DeleteCapiDiegoProcessAssociation(context context.Context, request *api.DeleteCapiDiegoProcessAssociationRequest) (*api.DeleteCapiDiegoProcessAssociationResponse, error) {
	c.Logger.Info("deleting capi/diego process association...")
	err := validateDeleteCAPIDiegoProcessAssociationRequest(request)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "%s", err)
	}

	cpg := models.CAPIProcessGUID(request.CapiProcessGuid)
	c.CAPIDiegoProcessAssociationsRepo.Delete(&cpg)

	return &api.DeleteCapiDiegoProcessAssociationResponse{}, nil
}

func (c *CAPI) BulkSync(stream api.CloudControllerCopilot_BulkSyncServer) error {
	c.Logger.Info("bulk sync...")
	var (
		requestChunk *api.BulkSyncRequestChunk
		err          error
	)

	allChunks := []byte{}
	request := &api.BulkSyncRequest{}

	for {
		requestChunk, err = stream.Recv()
		if err == io.EOF {
			if len(allChunks) > 0 {
				err = request.XXX_Unmarshal(allChunks)
				if err != nil {
					return err
				}
			}
			c.syncRequest(request)

			return stream.SendAndClose(&api.BulkSyncResponse{
				TotalBytesReceived: int32(len(allChunks)),
			})
		}

		if err != nil {
			return err
		}
		allChunks = append(allChunks, requestChunk.Chunk...)
	}
}

func (c *CAPI) syncRequest(request *api.BulkSyncRequest) {
	routeMappings := make([]*models.RouteMapping, len(request.RouteMappings))
	for i, routeMapping := range request.RouteMappings {
		routeMappings[i] = &models.RouteMapping{
			RouteGUID:       models.RouteGUID(routeMapping.RouteGuid),
			CAPIProcessGUID: models.CAPIProcessGUID(routeMapping.CapiProcessGuid),
			RouteWeight:     routeMapping.RouteWeight,
		}
	}

	routes := make([]*models.Route, len(request.Routes))

	for i, route := range request.Routes {
		routes[i] = &models.Route{
			GUID:     models.RouteGUID(route.GetGuid()),
			Host:     route.GetHost(),
			Path:     route.GetPath(),
			Internal: route.GetInternal(),
			VIP:      route.GetVip(),
		}
	}

	cdpas := make([]*models.CAPIDiegoProcessAssociation, len(request.CapiDiegoProcessAssociations))

	for i, cdpa := range request.CapiDiegoProcessAssociations {
		diegoProcessGuids := make([]models.DiegoProcessGUID, len(cdpa.DiegoProcessGuids))
		for j, diegoProcessGuid := range cdpa.DiegoProcessGuids {
			diegoProcessGuids[j] = models.DiegoProcessGUID(diegoProcessGuid)
		}
		cdpas[i] = &models.CAPIDiegoProcessAssociation{
			CAPIProcessGUID:   models.CAPIProcessGUID(cdpa.CapiProcessGuid),
			DiegoProcessGUIDs: diegoProcessGuids,
		}
	}

	c.RouteMappingsRepo.Sync(routeMappings)
	c.RoutesRepo.Sync(routes)
	c.CAPIDiegoProcessAssociationsRepo.Sync(cdpas)
}

func validateUpsertRouteRequest(r *api.UpsertRouteRequest) error {
	route := r.Route
	if route == nil {
		return errors.New("route is required")
	}
	if route.Guid == "" || route.Host == "" {
		return errors.New("route Guid and Host are required")
	}
	return nil
}

func validateDeleteRouteRequest(r *api.DeleteRouteRequest) error {
	if r.Guid == "" {
		return errors.New("route Guid is required")
	}
	return nil
}

func validateMapRouteRequest(r *api.MapRouteRequest) error {
	rm := r.RouteMapping
	if rm == nil {
		return errors.New("RouteMapping is required")
	}
	if rm.RouteGuid == "" || rm.CapiProcessGuid == "" {
		return errors.New("RouteGUID and CapiProcessGUID are required")
	}
	if rm.RouteWeight == 0 || rm.RouteWeight > 128 {
		return errors.New("RouteWeight must be between 1 and 128")
	}
	return nil
}

func validateUnmapRouteRequest(r *api.UnmapRouteRequest) error {
	rm := r.RouteMapping
	if rm == nil {
		return errors.New("RouteMapping is required")
	}
	if rm.RouteGuid == "" || rm.CapiProcessGuid == "" {
		return errors.New("RouteGuid and CapiProcessGuid are required")
	}
	if rm.RouteWeight == 0 || rm.RouteWeight > 128 {
		return errors.New("RouteWeight must be between 1 and 128")
	}
	return nil
}

func validateUpsertCAPIDiegoProcessAssociationRequest(r *api.UpsertCapiDiegoProcessAssociationRequest) error {
	association := r.CapiDiegoProcessAssociation
	if association == nil {
		return errors.New("CapiDiegoProcessAssociation is required")
	}
	if association.CapiProcessGuid == "" || len(association.DiegoProcessGuids) == 0 {
		return errors.New("CapiProcessGuid and DiegoProcessGuids are required")
	}
	return nil
}

func validateDeleteCAPIDiegoProcessAssociationRequest(r *api.DeleteCapiDiegoProcessAssociationRequest) error {
	if r.CapiProcessGuid == "" {
		return errors.New("CapiProcessGuid is required")
	}
	return nil
}
