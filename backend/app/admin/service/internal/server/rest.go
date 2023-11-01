package server

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/go-kratos/kratos/v2/transport/http"

	authnEngine "github.com/tx7do/kratos-authn/engine"
	authn "github.com/tx7do/kratos-authn/middleware"

	authzEngine "github.com/tx7do/kratos-authz/engine"
	authz "github.com/tx7do/kratos-authz/middleware"

	swaggerUI "github.com/tx7do/kratos-swagger-ui"

	bootstrap "github.com/tx7do/kratos-bootstrap"
	"github.com/tx7do/kratos-bootstrap/gen/api/go/conf/v1"

	"kratos-monolithic-demo/app/admin/service/cmd/server/assets"
	"kratos-monolithic-demo/app/admin/service/internal/service"

	v1 "kratos-monolithic-demo/gen/api/go/admin/service/v1"
	"kratos-monolithic-demo/pkg/middleware/auth"
)

// NewWhiteListMatcher 创建jwt白名单
func newRestWhiteListMatcher() selector.MatchFunc {
	whiteList := make(map[string]bool)
	whiteList[v1.OperationAuthenticationServiceLogin] = true
	return func(ctx context.Context, operation string) bool {
		if _, ok := whiteList[operation]; ok {
			return false
		}
		return true
	}
}

// NewMiddleware 创建中间件
func newRestMiddleware(authenticator authnEngine.Authenticator, authorizer authzEngine.Engine, logger log.Logger) []middleware.Middleware {
	var ms []middleware.Middleware
	ms = append(ms, logging.Server(logger))
	ms = append(ms, selector.Server(
		authn.Server(authenticator),
		auth.Server(),
		authz.Server(authorizer),
	).Match(newRestWhiteListMatcher()).Build())
	return ms
}

// NewRESTServer new an HTTP server.
func NewRESTServer(
	cfg *conf.Bootstrap, logger log.Logger,
	authenticator authnEngine.Authenticator, authorizer authzEngine.Engine,
	authnSvc *service.AuthenticationService,
	userSvc *service.UserService,
	dictSvc *service.DictService,
	dictDetailSvc *service.DictDetailService,
	menuSvc *service.MenuService,
	routerSvc *service.RouterService,
	orgSvc *service.OrganizationService,
	roleSvc *service.RoleService,
	positionSvc *service.PositionService,
) *http.Server {
	srv := bootstrap.CreateRestServer(cfg, newRestMiddleware(authenticator, authorizer, logger)...)

	v1.RegisterAuthenticationServiceHTTPServer(srv, authnSvc)
	v1.RegisterUserServiceHTTPServer(srv, userSvc)
	v1.RegisterDictServiceHTTPServer(srv, dictSvc)
	v1.RegisterDictDetailServiceHTTPServer(srv, dictDetailSvc)
	v1.RegisterMenuServiceHTTPServer(srv, menuSvc)
	v1.RegisterRouterServiceHTTPServer(srv, routerSvc)
	v1.RegisterOrganizationServiceHTTPServer(srv, orgSvc)
	v1.RegisterRoleServiceHTTPServer(srv, roleSvc)
	v1.RegisterPositionServiceHTTPServer(srv, positionSvc)

	if cfg.GetServer().GetRest().GetEnableSwagger() {
		swaggerUI.RegisterSwaggerUIServerWithOption(
			srv,
			swaggerUI.WithTitle("Kratos巨石应用实践"),
			swaggerUI.WithMemoryData(assets.OpenApiData, "yaml"),
		)
	}

	return srv
}
