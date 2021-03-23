# BGPLB

BGPLB is a simple load-balancer implementation for bare metal Kubernetes clusters, 
using calico bgp to advertise service ip.

## Core Feature

* LoadBalancerIP assignment in Kubernetes services
* Support specify IP for services
* Auto detector Calico cidr config

## How to Build

This Project build by [kubebuilder](https://github.com/kubernetes-sigs/kubebuilder).

### Prerequisites

1、Go 1.11+

2、Kubernetes 1.18+

3、Calico 3.10+ with BGP

### Deploy

1、make docker-build ${IMG}

2、make  docker-push ${IMG}

3、make deploy
