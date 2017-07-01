package controller

import (
	"errors"
	"fmt"

	"github.com/golang/glog"
	"github.com/kubernetes-incubator/service-catalog/contrib/pkg/broker/controller"
	"github.com/kubernetes-incubator/service-catalog/pkg/brokerapi"
	"github.com/yqf3139/redis-broker/client"
)

type errNoSuchInstance struct {
	instanceID string
}

func (e errNoSuchInstance) Error() string {
	return fmt.Sprintf("no such instance with ID %s", e.instanceID)
}

type redisController struct {
}

// CreateController creates an instance of a User Provided service broker controller.
func CreateController() controller.Controller {
	return &redisController{}
}

func (c *redisController) Catalog() (*brokerapi.Catalog, error) {
	return &brokerapi.Catalog{
		Services: []*brokerapi.Service{
			{
				Name:        "redis",
				ID:          "b624ec28-5e3c-11e7-88d3-a79dd6001a67",
				Description: "Redis cache",
				Plans: []brokerapi.ServicePlan{
					{
						Name:        "default",
						ID:          "c5e56cf0-5e3c-11e7-b090-23414af6e79c",
						Description: "Redis cache",
						Free:        true,
					},
				},
				Bindable: true,
			},
		},
	}, nil
}

func (c *redisController) CreateServiceInstance(id string, req *brokerapi.CreateServiceInstanceRequest) (*brokerapi.CreateServiceInstanceResponse, error) {
	if err := client.Install(releaseName(id), id); err != nil {
		return nil, err
	}
	glog.Infof("Created Redis Service Instance:\n%v\n", id)
	return &brokerapi.CreateServiceInstanceResponse{}, nil
}

func (c *redisController) GetServiceInstance(id string) (string, error) {
	return "", errors.New("Unimplemented")
}

func (c *redisController) RemoveServiceInstance(id string) (*brokerapi.DeleteServiceInstanceResponse, error) {
	if err := client.Delete(releaseName(id)); err != nil {
		return nil, err
	}
	return &brokerapi.DeleteServiceInstanceResponse{}, nil
}

func (c *redisController) Bind(instanceID, bindingID string, req *brokerapi.BindingRequest) (*brokerapi.CreateServiceBindingResponse, error) {
	host := releaseName(instanceID) + "-redis." + instanceID + ".svc.cluster.local"
	port := "6379"
	password, err := client.GetPassword(releaseName(instanceID), instanceID)
	if err != nil {
		return nil, err
	}
	return &brokerapi.CreateServiceBindingResponse{
		Credentials: brokerapi.Credential{
			"password": password,
			"host":     host,
			"port":     port,
		},
	}, nil
}

func (c *redisController) UnBind(instanceID string, bindingID string) error {
	// Since we don't persist the binding, there's nothing to do here.
	return nil
}

func releaseName(id string) string {
	return "i-" + id
}
