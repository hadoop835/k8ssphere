package apiserver

import (
	"github.com/emicklei/go-restful"
	urlruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/klog/v2"
	"k8ssphere.io/k8ssphere/pkg/apiserver/config"
	"net/http"
)

const (
	// ApiRootPath defines the root path of all KubeSphere apis.
	ApiRootPath = "/kapis"

	// MimeMergePatchJson is the mime header used in merge request
	MimeMergePatchJson = "application/merge-patch+json"

	//
	MimeJsonPatchJson = "application/json-patch+json"
)

type APIServer struct {
	// number of kubesphere apiserver
	ServerCount int
	//
	Server *http.Server
	//
	Config *config.Config
	// webservice container, where all webservice defines
	container *restful.Container
}

func (s *APIServer) PrepareRun(stopCh <-chan struct{}) error {
	s.container = restful.NewContainer()
	s.container.Router(restful.CurlyRouter{})
	for _, ws := range s.container.RegisteredWebServices() {
		klog.V(2).Infof("%s", ws.RootPath())
	}
	s.Server.Handler = s.container

	return nil
}

func (s *APIServer) installKubeSphereAPIs() {

	urlruntime.Must(configv1alpha2.AddToContainer(s.container, s.Config))

}
