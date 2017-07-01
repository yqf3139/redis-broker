# Redis Broker

This is an implementation of a Service Broker that uses Helm to provision
instances of [Redis](https://kubeapps.com/charts/stable/redis). This is a
**proof-of-concept** for the [Kubernetes Service
Catalog](https://github.com/kubernetes-incubator/service-catalog), and should not
be used in production. Thanks to the [mariadb broker repo](https://github.com/prydonius/mariadb-broker).

## Prerequisites

1. Kubernetes cluster
2. [Helm 2.x](https://github.com/kubernetes/helm)
3. [Service Catalog API](https://github.com/kubernetes-incubator/service-catalog) - follow the [walkthrough](https://github.com/kubernetes-incubator/service-catalog/blob/master/docs/walkthrough.md)

## Installing the Broker

The Redis Service Broker can be installed using the Helm chart in this
repository.

```
$ git clone https://github.com/yqf3139/redis-broker.git
$ cd redis-broker
$ helm install --name redis-broker --namespace redis-broker charts/redis-broker
```

To register the Broker with the Service Catalog, create the Broker object:

```
$ kubectl --context service-catalog create -f examples/redis-broker.yaml
```

If the Broker was successfully registered, the `redis` ServiceClass will now
be available in the catalog:

```
$ kubectl --context service-catalog get serviceclasses
NAME      KIND
redis   ServiceClass.v1alpha1.servicecatalog.k8s.io
```

## Usage

### Create the Instance object

```
$ kubectl --context service-catalog create -f examples/redis-instance.yaml
```

This will result in the installation of a new Redis chart:

```
$ helm list
NAME                                  	REVISION	UPDATED                 	STATUS  	CHART               	NAMESPACE
i-3e0e9973-a072-49ba-8308-19568e7f4669	1       	Sat May 13 17:28:35 2017	DEPLOYED	redis-0.6.1       	3e0e9973-a072-49ba-8308-19568e7f4669
```

### Create a Binding to fetch credentials

```
$ kubectl --context service-catalog create -f examples/redis-binding.yaml
```

A secret called `redis-instance-credentials` will be created containing the
connection details for this Redis instance.

```
$ kubectl get secret redis-instance-credentials -o yaml
```
