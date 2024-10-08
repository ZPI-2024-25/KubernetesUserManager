/*
 * KubernetesUserManager - API
 *
 * This is a backend API server documentation for KubernetesUserManager  Some useful links: - [Jira](https://samuelus.atlassian.net/jira/software/projects/ZPI/boards/4) - [Confluence](https://samuelus.atlassian.net/wiki/spaces/ZPI/overview)
 *
 * API version: 0.0.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package api

import (
	"fmt"
	"github.com/ZPI-2024-25/KubernetesUserManager/go/common"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = common.Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/api/v1/",
		Index,
	},

	Route{
		"DeleteUser",
		strings.ToUpper("Delete"),
		"/api/v1/users/{id}",
		DeleteUser,
	},

	Route{
		"GetUser",
		strings.ToUpper("Get"),
		"/api/v1/users/{id}",
		GetUser,
	},

	Route{
		"ListUsers",
		strings.ToUpper("Get"),
		"/api/v1/users",
		ListUsers,
	},

	Route{
		"UpdateUserPermissions",
		strings.ToUpper("Put"),
		"/api/v1/users/{id}/permissions",
		UpdateUserPermissions,
	},

	Route{
		"GetHelmRelease",
		strings.ToUpper("Get"),
		"/api/v1/helm/releases/{namespace}/{releaseName}",
		GetHelmRelease,
	},

	Route{
		"GetHelmReleaseHistory",
		strings.ToUpper("Get"),
		"/api/v1/helm/releases/{namespace}/{releaseName}/history",
		GetHelmReleaseHistory,
	},

	Route{
		"InstallHelmRelease",
		strings.ToUpper("Post"),
		"/api/v1/helm/releases/{namespace}",
		InstallHelmRelease,
	},

	Route{
		"ListHelmReleases",
		strings.ToUpper("Get"),
		"/api/v1/helm/releases",
		ListHelmReleases,
	},

	Route{
		"ListHelmReleasesInNamespace",
		strings.ToUpper("Get"),
		"/api/v1/helm/releases/{namespace}",
		ListHelmReleasesInNamespace,
	},

	Route{
		"RollbackHelmRelease",
		strings.ToUpper("Post"),
		"/api/v1/helm/releases/{namespace}/{releaseName}/rollback",
		RollbackHelmRelease,
	},

	Route{
		"UninstallHelmRelease",
		strings.ToUpper("Delete"),
		"/api/v1/helm/releases/{namespace}/{releaseName}",
		UninstallHelmRelease,
	},

	Route{
		"UpdateHelmRelease",
		strings.ToUpper("Put"),
		"/api/v1/helm/releases/{namespace}/{releaseName}",
		UpdateHelmRelease,
	},

	Route{
		"CreateClusterResource",
		strings.ToUpper("Post"),
		"/api/v1/k8s/non-namespaced/{resourceType}",
		CreateClusterResource,
	},

	Route{
		"CreateNamespacedResource",
		strings.ToUpper("Post"),
		"/api/v1/k8s/namespaced/{resourceType}/{namespace}",
		CreateNamespacedResource,
	},

	Route{
		"DeleteClusterResource",
		strings.ToUpper("Delete"),
		"/api/v1/k8s/non-namespaced/{resourceType}/{resourceName}",
		DeleteClusterResource,
	},

	Route{
		"DeleteNamespacedResource",
		strings.ToUpper("Delete"),
		"/api/v1/k8s/namespaced/{resourceType}/{namespace}/{resourceName}",
		DeleteNamespacedResource,
	},

	Route{
		"GetClusterResource",
		strings.ToUpper("Get"),
		"/api/v1/k8s/non-namespaced/{resourceType}/{resourceName}",
		GetClusterResource,
	},

	Route{
		"GetNamespacedResource",
		strings.ToUpper("Get"),
		"/api/v1/k8s/namespaced/{resourceType}/{namespace}/{resourceName}",
		GetNamespacedResource,
	},

	Route{
		"ListClusterResources",
		strings.ToUpper("Get"),
		"/api/v1/k8s/non-namespaced/{resourceType}",
		ListClusterResources,
	},

	Route{
		"ListNamespacedResources",
		strings.ToUpper("Get"),
		"/api/v1/k8s/namespaced/{resourceType}/{namespace}",
		ListNamespacedResources,
	},

	Route{
		"UpdateClusterResource",
		strings.ToUpper("Put"),
		"/api/v1/k8s/non-namespaced/{resourceType}/{resourceName}",
		UpdateClusterResource,
	},

	Route{
		"UpdateNamespacedResource",
		strings.ToUpper("Put"),
		"/api/v1/k8s/namespaced/{resourceType}/{namespace}/{resourceName}",
		UpdateNamespacedResource,
	},

	Route{
		"CheckLoginStatus",
		strings.ToUpper("Get"),
		"/api/v1/auth/status",
		CheckLoginStatus,
	},

	Route{
		"ExchangeAuthCode",
		strings.ToUpper("Post"),
		"/api/v1/auth/callback",
		ExchangeAuthCode,
	},

	Route{
		"LogoutUser",
		strings.ToUpper("Post"),
		"/api/v1/auth/logout",
		LogoutUser,
	},

	Route{
		"RefreshToken",
		strings.ToUpper("Post"),
		"/api/v1/auth/refresh",
		RefreshToken,
	},
}
