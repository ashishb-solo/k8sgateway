---
title: "Developer Guides"
weight: 60
description: Extending Gloo Gateway's functionality with the use of plugins
---

## Intro


Gloo Gateway invites developers to extend Gloo Gateway's functionality and adapt to new use cases via the addition of plugins. 

Gloo Gateway's plugin based architecture makes it easy to extend functionality in a variety of areas:

- [Gloo Gateway's API](https://github.com/solo-io/gloo/tree/main/projects/controller/api/v1): extensible through the use of [Protocol Buffers](https://developers.google.com/protocol-buffers/) along with [Solo-Kit](https://github.com/solo-io/solo-kit)
- [Service Discovery Plugins](https://github.com/solo-io/gloo/blob/main/projects/controller/pkg/discovery/discovery.go#L21): automatically discover service endpoints from catalogs such as [Kubernetes](https://github.com/solo-io/gloo/tree/main/projects/controller/pkg/plugins/kubernetes) and [Consul](https://github.com/solo-io/gloo/tree/main/projects/controller/pkg/plugins/consul)
- [Function Discovery Plugins](https://github.com/solo-io/gloo/blob/main/projects/discovery/pkg/fds/interface.go#L31): annotate services with information discovered by polling services directly (such as OpenAPI endpoints and gRPC methods).
- [Routing Plugins](https://github.com/solo-io/gloo/blob/main/projects/controller/pkg/plugins/plugin_interface.go#L53): customize what happens to requests when they match a route or virtual host
- [Upstream Plugins](https://github.com/solo-io/gloo/tree/main/projects/controller/pkg/plugins): customize what happens to requests when they are routed to a service
- **Operators for Configuration**: Gloo Gateway exposes its intermediate language for proxy configuration via the {{< protobuf name="gloo.solo.io.Proxy" display="Proxy">}} Custom Resource, allowing operators to leverage Gloo Gateway for multiple use cases. With the optional [Gloo Gateway GraphQL module]({{< versioned_link_path fromRoot="/guides/graphql/" >}}), you can set up API gateway and GraphQL server functionality for your apps, without running in the same process (or even the same container) as Gloo Gateway.

## Gloo Gateway API Concepts


* {{< protobuf name="gloo.solo.io.Proxy" display="v1.Proxies">}} provide the routing configuration which Gloo Gateway will translate and apply to Envoy.
* {{< protobuf name="gloo.solo.io.Upstream" display="v1.Upstreams" >}} describe routable destinations for Gloo Gateway.

* **Proxies** represent a unified configuration to be applied to one or more instances of a proxy. You can think of the proxy of as tree like such:

        proxy
        ├─ bind-address
        │  ├── domain-set
        │  │  ├── /route
        │  │  ├── /route
        │  │  ├── /route
        │  │  └── tls-config
        │  └── domain-set
        │     ├── /route
        │     ├── /route
        │     ├── /route
        │     └── tls-config
        └─ bind-address
           ├── domain-set
           │  ├── /route
           │  ├── /route
           │  ├── /route
           │  └── tls-config
           └── domain-set
              ├── /route
              ├── /route
              ├── /route
              └── tls-config

  A single proxy CRD contains all the configuration necessary to be applied to an instance of Envoy. In the Gloo Gateway system, Proxies are treated as an intermediary representation of config, while user-facing config is imported from simpler, more opinionated resources such as the {{< protobuf name="gateway.solo.io.VirtualService">}} or [Kubernetes Ingress objects](https://kubernetes.io/docs/concepts/services-networking/ingress/).
  
  For this reason, a standard Gloo Gateway deployment contains one or more controllers which programmatically generate and write these CRDs to provide simpler, use-case specific APIs such as API Gateway and Ingress. You can extend these capabilities even more with Gloo Gateway modules, such as [Gloo Gateway GraphQL]({{< versioned_link_path fromRoot="/guides/graphql/" >}}). This optional module is an advanced controller which creates routing configuration for Gloo Gateway from [**GraphQL Schemas**](https://graphql.org/).

* **Upstreams** represent destinations for routing requests in Gloo Gateway. Routes in Gloo Gateway specify one or more Upstreams (by name) as their destination. Upstreams have a `type` which is provided in their `upstreamSpec` field. Each type of upstream corresponds to an **Upstream Plugin**, which tells Gloo Gateway how to translate upstreams of that type to Envoy clusters. When a route is declared for an upstream, Gloo Gateway invokes the corresponding plugin for that type 


## Guides

This Section includes the following developer guides:

{{% children description="true" %}}

> Note: the Controller tutorial does not require modifying any Gloo Gateway code.

