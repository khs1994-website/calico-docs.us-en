# [Calico 英文文档](https://github.com/projectcalico/calico)

- [Api](api/README.md)
- [Apiserver](apiserver/README.md)
- [App Policy](app-policy/README.md)
- [Calico](calico/README.md)
  - Includes
    - Charts
      - [Tigera Operator](calico/_includes/charts/tigera-operator/README.md)
    - Content
      * [Auto Hostendpoints Migrate](calico/_includes/content/auto-hostendpoints-migrate.md)
      * [Calico Upgrade](calico/_includes/content/calico-upgrade.md)
      * [Calicoctl Convert](calico/_includes/content/calicoctl-convert.md)
      * [Calicoctl Version](calico/_includes/content/calicoctl-version.md)
      * [Config Calico Upgrade Etcd](calico/_includes/content/config-calico-upgrade-etcd.md)
      * [Deleting Etcd V 2 V 3 Data](calico/_includes/content/deleting-etcd-v2-v3-data.md)
      * [Determine Ipam](calico/_includes/content/determine-ipam.md)
      * [Docker Container Service](calico/_includes/content/docker-container-service.md)
      * [Endpointport](calico/_includes/content/endpointport.md)
      * [Entityrule](calico/_includes/content/entityrule.md)
      * [Environment File](calico/_includes/content/environment-file.md)
      * [Felix Init Datastore](calico/_includes/content/felix-init-datastore.md)
      * [Hostendpoints Upgrade](calico/_includes/content/hostendpoints-upgrade.md)
      * [Httpmatch](calico/_includes/content/httpmatch.md)
      * [Icmp](calico/_includes/content/icmp.md)
      * [Install Calico Upgrade](calico/_includes/content/install-calico-upgrade.md)
      * [Install K 8 S Addons](calico/_includes/content/install-k8s-addons.md)
      * [Install Openshift Manifests](calico/_includes/content/install-openshift-manifests.md)
      * [Ipnat](calico/_includes/content/ipnat.md)
      * [Openstack Etcd Auth](calico/_includes/content/openstack-etcd-auth.md)
      * [Pod Cidr Sed](calico/_includes/content/pod-cidr-sed.md)
      * [Ports](calico/_includes/content/ports.md)
      * [Reqs Kernel](calico/_includes/content/reqs-kernel.md)
      * [Reqs Sys](calico/_includes/content/reqs-sys.md)
      * [Rule](calico/_includes/content/rule.md)
      * [Selector Scopes](calico/_includes/content/selector-scopes.md)
      * [Selectors](calico/_includes/content/selectors.md)
      * [Serviceaccountmatch](calico/_includes/content/serviceaccountmatch.md)
      * [Servicematch](calico/_includes/content/servicematch.md)
      * [Testing Data Migration](calico/_includes/content/testing-data-migration.md)
    - Geek Details
      * [Cni Aws](calico/_includes/geek-details/cni-aws.md)
      * [Cni Azure](calico/_includes/geek-details/cni-azure.md)
      * [Cni Calico](calico/_includes/geek-details/cni-calico.md)
      * [Cni Kubenet](calico/_includes/geek-details/cni-kubenet.md)
      * [Cross Subnet Ipip](calico/_includes/geek-details/cross-subnet-ipip.md)
      * [Cross Subnet Vxlan](calico/_includes/geek-details/cross-subnet-vxlan.md)
      * [Datastore Etcd](calico/_includes/geek-details/datastore-etcd.md)
      * [Datastore Kubernetes](calico/_includes/geek-details/datastore-kubernetes.md)
      * [Info](calico/_includes/geek-details/info.md)
      * [Ipam Aws](calico/_includes/geek-details/ipam-aws.md)
      * [Ipam Azure](calico/_includes/geek-details/ipam-azure.md)
      * [Ipam Calico](calico/_includes/geek-details/ipam-calico.md)
      * [Ipam Host Local](calico/_includes/geek-details/ipam-host-local.md)
      * [Overlay Ipip](calico/_includes/geek-details/overlay-ipip.md)
      * [Overlay No](calico/_includes/geek-details/overlay-no.md)
      * [Overlay Vxlan](calico/_includes/geek-details/overlay-vxlan.md)
      * [Policy Calico](calico/_includes/geek-details/policy-calico.md)
      * [Policy None](calico/_includes/geek-details/policy-none.md)
      * [Routing Bgp](calico/_includes/geek-details/routing-bgp.md)
      * [Routing Calico](calico/_includes/geek-details/routing-calico.md)
      * [Routing Static](calico/_includes/geek-details/routing-static.md)
      * [Routing Vpc Native](calico/_includes/geek-details/routing-vpc-native.md)
  - About
    * [About Calico](calico/about/about-calico.md)
    * [About E BPF](calico/about/about-ebpf.md)
    * [About Kubernetes Networking](calico/about/about-k8s-networking.md)
    * [About Kubernetes Egress](calico/about/about-kubernetes-egress.md)
    * [About Kubernetes Ingress](calico/about/about-kubernetes-ingress.md)
    * [About Kubernetes Services](calico/about/about-kubernetes-services.md)
    * [About Network Policy](calico/about/about-network-policy.md)
    * [About Networking](calico/about/about-networking.md)
    * [About](calico/about/index.md)
  - Calico Enterprise
    * [Calico Cloud](calico/calico-enterprise/index.md)
  - Getting Started
    - Bare Metal
      - Installation
        * [Binary Install With Package Manager](calico/getting-started/bare-metal/installation/binary-mgr.md)
        * [Binary Install Without Package Manager](calico/getting-started/bare-metal/installation/binary.md)
        * [Docker Container Install](calico/getting-started/bare-metal/installation/container.md)
        * [Index](calico/getting-started/bare-metal/installation/index.md)
      * [About Non Cluster Hosts](calico/getting-started/bare-metal/about.md)
      * [Non Cluster Hosts](calico/getting-started/bare-metal/index.md)
      * [System Requirements](calico/getting-started/bare-metal/requirements.md)
    - Clis
      - [Calicoctl](calico/getting-started/clis/calicoctl/index.md)
        - Configure
          * [Configure Calicoctl To Connect To An Etcd Datastore](calico/getting-started/clis/calicoctl/configure/etcd.md)
          * [Index](calico/getting-started/clis/calicoctl/configure/index.md)
          * [Configure Calicoctl To Connect To The Kubernetes API Datastore](calico/getting-started/clis/calicoctl/configure/kdd.md)
          * [Configure Calicoctl](calico/getting-started/clis/calicoctl/configure/overview.md)
    - Kubernetes
      - Flannel
        * [Install Calico For Policy And Flannel Aka Canal For Networking](calico/getting-started/kubernetes/flannel/flannel.md)
        * [Flannel](calico/getting-started/kubernetes/flannel/index.md)
        * [Migrate A Kubernetes Cluster From Flannel Canal To Calico](calico/getting-started/kubernetes/flannel/migration-from-flannel.md)
      - Hardway
        * [Configure BGP Peering](calico/getting-started/kubernetes/hardway/configure-bgp-peering.md)
        * [Configure IP Pools](calico/getting-started/kubernetes/hardway/configure-ip-pools.md)
        * [End User RBAC](calico/getting-started/kubernetes/hardway/end-user-rbac.md)
        * [Index](calico/getting-started/kubernetes/hardway/index.md)
        * [Install CNI Plugin](calico/getting-started/kubernetes/hardway/install-cni-plugin.md)
        * [Install Calico Node](calico/getting-started/kubernetes/hardway/install-node.md)
        * [Install Typha](calico/getting-started/kubernetes/hardway/install-typha.md)
        * [Istio Integration](calico/getting-started/kubernetes/hardway/istio-integration.md)
        * [Calico The Hard Way](calico/getting-started/kubernetes/hardway/overview.md)
        * [Stand Up Kubernetes](calico/getting-started/kubernetes/hardway/standing-up-kubernetes.md)
        * [Test Network Policy](calico/getting-started/kubernetes/hardway/test-network-policy.md)
        * [Test Networking](calico/getting-started/kubernetes/hardway/test-networking.md)
        * [The Calico Datastore](calico/getting-started/kubernetes/hardway/the-calico-datastore.md)
      - Installation
        * [Customize Calico Configuration](calico/getting-started/kubernetes/installation/config-options.md)
      - K 3 S
        * [K 3 S](calico/getting-started/kubernetes/k3s/index.md)
        * [K 3 S Multi Node Install](calico/getting-started/kubernetes/k3s/multi-node-install.md)
        * [Quickstart For Calico On K 3 S](calico/getting-started/kubernetes/k3s/quickstart.md)
      - Managed Public Cloud
        * [Microsoft Azure Kubernetes Service AKS](calico/getting-started/kubernetes/managed-public-cloud/aks.md)
        * [Amazon Elastic Kubernetes Service EKS](calico/getting-started/kubernetes/managed-public-cloud/eks.md)
        * [Google Kubernetes Engine GKE](calico/getting-started/kubernetes/managed-public-cloud/gke.md)
        * [IBM Cloud Kubernetes Service IKS](calico/getting-started/kubernetes/managed-public-cloud/iks.md)
        * [Managed Public Cloud](calico/getting-started/kubernetes/managed-public-cloud/index.md)
      - Self Managed Onprem
        * [Self Managed On Premises](calico/getting-started/kubernetes/self-managed-onprem/index.md)
        * [Install Calico Networking And Network Policy For On Premises Deployments](calico/getting-started/kubernetes/self-managed-onprem/onpremises.md)
      - Self Managed Public Cloud
        * [Self Managed Kubernetes In Amazon Web Services AWS](calico/getting-started/kubernetes/self-managed-public-cloud/aws.md)
        * [Self Managed Kubernetes In Microsoft Azure](calico/getting-started/kubernetes/self-managed-public-cloud/azure.md)
        * [Self Managed Kubernetes In Digital Ocean DO](calico/getting-started/kubernetes/self-managed-public-cloud/do.md)
        * [Self Managed Kubernetes In Google Compute Engine GCE](calico/getting-started/kubernetes/self-managed-public-cloud/gce.md)
        * [Self Managed Public Cloud](calico/getting-started/kubernetes/self-managed-public-cloud/index.md)
      - Vpp
        * [Get Started With VPP Networking](calico/getting-started/kubernetes/vpp/getting-started.md)
        * [VPP Dataplane Tech Preview](calico/getting-started/kubernetes/vpp/index.md)
        * [I Psec Configuration With VPP](calico/getting-started/kubernetes/vpp/ipsec.md)
      * [Install Using Helm](calico/getting-started/kubernetes/helm.md)
      * [Kubernetes](calico/getting-started/kubernetes/index.md)
      * [Quickstart For Calico On Micro K 8 S](calico/getting-started/kubernetes/microk8s.md)
      * [Quickstart For Calico On Minikube](calico/getting-started/kubernetes/minikube.md)
      * [Quickstart For Calico On Kubernetes](calico/getting-started/kubernetes/quickstart.md)
      * [Install Calico On A Rancher Kubernetes Engine Cluster](calico/getting-started/kubernetes/rancher.md)
      * [System Requirements](calico/getting-started/kubernetes/requirements.md)
    - Openshift
      * [Open Shift](calico/getting-started/openshift/index.md)
      * [Install An Open Shift 4 Cluster With Calico](calico/getting-started/openshift/installation.md)
      * [System Requirements](calico/getting-started/openshift/requirements.md)
    - Openstack
      - Installation
        * [Dev Stack](calico/getting-started/openstack/installation/devstack.md)
        * [Index](calico/getting-started/openstack/installation/index.md)
        * [Calico On Open Stack](calico/getting-started/openstack/installation/overview.md)
        * [Red Hat Enterprise Linux](calico/getting-started/openstack/installation/redhat.md)
        * [Ubuntu](calico/getting-started/openstack/installation/ubuntu.md)
      * [Index](calico/getting-started/openstack/index.md)
      * [Calico For Open Stack](calico/getting-started/openstack/overview.md)
      * [System Requirements](calico/getting-started/openstack/requirements.md)
      * [Verify Your Deployment](calico/getting-started/openstack/verification.md)
    - Windows Calico
      - Kubernetes
        * [Kubernetes](calico/getting-started/windows-calico/kubernetes/index.md)
        * [Install Calico For Windows On A Rancher Kubernetes Engine Cluster](calico/getting-started/windows-calico/kubernetes/rancher.md)
        * [Requirements](calico/getting-started/windows-calico/kubernetes/requirements.md)
        * [Install Calico For Windows](calico/getting-started/windows-calico/kubernetes/standard.md)
      - Openshift
        * [Open Shift](calico/getting-started/windows-calico/openshift/index.md)
        * [Install An Open Shift 4 Cluster On Windows Nodes](calico/getting-started/windows-calico/openshift/installation.md)
      * [Basic Policy Demo](calico/getting-started/windows-calico/demo.md)
      * [Calico For Windows](calico/getting-started/windows-calico/index.md)
      * [Create Kubeconfig For Windows Nodes](calico/getting-started/windows-calico/kubeconfig.md)
      * [Limitations And Known Issues](calico/getting-started/windows-calico/limitations.md)
      * [Start And Stop Calico For Windows Services](calico/getting-started/windows-calico/maintain.md)
      * [Quickstart](calico/getting-started/windows-calico/quickstart.md)
      * [Troubleshoot Calico For Windows](calico/getting-started/windows-calico/troubleshoot.md)
    * [Install Calico](calico/getting-started/index.md)
  - Hack
    - [Development Environment](calico/hack/development-environment/README.md)
    - [Remove Calico Policy](calico/hack/remove-calico-policy/README.md)
      * [Override Policy](calico/hack/remove-calico-policy/override-policy.md)
      * [Remove Policy](calico/hack/remove-calico-policy/remove-policy.md)
  - Maintenance
    - Clis
      - [Calicoctl](calico/maintenance/clis/calicoctl/index.md)
        - Configure
          * [Configure Calicoctl To Connect To An Etcd Datastore](calico/maintenance/clis/calicoctl/configure/etcd.md)
          * [Configure Calicoctl](calico/maintenance/clis/calicoctl/configure/overview.md)
          * [Configure Calicoctl To Connect To The Kubernetes API Datastore](calico/maintenance/clis/calicoctl/configure/kdd.md)
        * [Install Calicoctl](calico/maintenance/clis/calicoctl/install.md)
    - Ebpf
      * [Creating An EKS Cluster For E BPF Mode](calico/maintenance/ebpf/ebpf-and-eks.md)
      * [Enable The E BPF Dataplane](calico/maintenance/ebpf/enabling-bpf.md)
      * [E BPF Dataplane Mode](calico/maintenance/ebpf/index.md)
      * [E BPF Use Cases](calico/maintenance/ebpf/use-cases-ebpf.md)
    - Image Options
      * [Configure Use Of Your Image Registry](calico/maintenance/image-options/alternate-registry.md)
      * [Install Images By Registry Digest](calico/maintenance/image-options/imageset.md)
      * [Install Using An Alternate Registry](calico/maintenance/image-options/index.md)
    - Monitor
      * [Monitor](calico/maintenance/monitor/index.md)
      * [Monitor Calico Component Metrics](calico/maintenance/monitor/monitor-component-metrics.md)
      * [Visualizing Metrics Via Grafana](calico/maintenance/monitor/monitor-component-visual.md)
    - Troubleshoot
      * [Troubleshooting Commands](calico/maintenance/troubleshoot/commands.md)
      * [Component Logs](calico/maintenance/troubleshoot/component-logs.md)
      * [Troubleshoot](calico/maintenance/troubleshoot/index.md)
      * [Troubleshooting E BPF Mode](calico/maintenance/troubleshoot/troubleshoot-ebpf.md)
      * [Troubleshooting And Diagnostics](calico/maintenance/troubleshoot/troubleshooting.md)
      * [VPP Dataplane Troubleshooting](calico/maintenance/troubleshoot/vpp.md)
    * [Manage TLS Certificates Used By Calico](calico/maintenance/certificate-management.md)
    * [Migrate Calico Data From An Etcdv 3 Datastore To A Kubernetes Datastore](calico/maintenance/datastore-migration.md)
    * [Decommission A Node](calico/maintenance/decommissioning-a-node.md)
    * [Creating An EKS Cluster For E BPF Mode](calico/maintenance/ebpf-and-eks.md)
    * [Enable The E BPF Dataplane](calico/maintenance/enabling-bpf.md)
    * [Maintenance](calico/maintenance/index.md)
    * [Enable Kubectl To Manage Calico AP Is](calico/maintenance/install-apiserver.md)
    * [Upgrade Calico On Kubernetes](calico/maintenance/kubernetes-upgrade.md)
    * [Upgrade Calico On Open Shift 4](calico/maintenance/openshift-upgrade.md)
    * [Upgrade Calico On Open Stack](calico/maintenance/openstack-upgrade.md)
    * [Migrate Calico To An Operator Managed Installation](calico/maintenance/operator-migration.md)
    * [Upgrade](calico/maintenance/upgrading.md)
  - Networking
    - Calico Enterprise
      * [Federation](calico/networking/calico-enterprise/federation.md)
      * [Calico Enterprise](calico/networking/calico-enterprise/index.md)
      * [Network Visibility](calico/networking/calico-enterprise/network-visibility.md)
      * [User Console](calico/networking/calico-enterprise/user-console.md)
    - Openstack
      * [Configure Systems For Use With Calico](calico/networking/openstack/configuration.md)
      * [IP Addressing And Connectivity](calico/networking/openstack/connectivity.md)
      * [Set Up A Development Machine](calico/networking/openstack/dev-machine-setup.md)
      * [Floating I Ps](calico/networking/openstack/floating-ips.md)
      * [Host Routes](calico/networking/openstack/host-routes.md)
    * [Add A Floating IP To A Pod](calico/networking/add-floating-ip.md)
    * [Advertise Kubernetes Service IP Addresses](calico/networking/advertise-service-ips.md)
    * [Assign IP Addresses Based On Topology](calico/networking/assign-ip-addresses-topology.md)
    * [Configure BGP Peering](calico/networking/bgp.md)
    * [Change IP Pool Block Size](calico/networking/change-block-size.md)
    * [Configure Networking](calico/networking/configuring.md)
    * [Determine Best Networking Option](calico/networking/determine-best-networking.md)
    * [Enable IPVS In Kubernetes](calico/networking/enabling-ipvs.md)
    * [External Connectivity](calico/networking/external-connectivity.md)
    * [Get Started With IP Address Management](calico/networking/get-started-ip-addresses.md)
    * [Networking](calico/networking/index.md)
    * [Configure IP Autodetection](calico/networking/ip-autodetection.md)
    * [IP Address Management](calico/networking/ipam.md)
    * [Configure Kubernetes Control Plane To Operate Over I Pv 6](calico/networking/ipv6-control-plane.md)
    * [Configure Dual Stack Or I Pv 6 Only](calico/networking/ipv6.md)
    * [Restrict A Pod To Use An IP Address In A Specific Range](calico/networking/legacy-firewalls.md)
    * [Migrate From One IP Pool To Another](calico/networking/migrate-pools.md)
    * [Configure MTU To Maximize Network Performance](calico/networking/mtu.md)
    * [Accelerate Istio Network Performance](calico/networking/sidecar-acceleration.md)
    * [Use IPVS Kube Proxy](calico/networking/use-ipvs.md)
    * [Use A Specific IP Address With A Pod](calico/networking/use-specific-ip.md)
    * [Overlay Networking](calico/networking/vxlan-ipip.md)
    * [Configure Outgoing NAT](calico/networking/workloads-outside-cluster.md)
  - Reference
    - Architecture
      - Design
        * [Network Design](calico/reference/architecture/design/index.md)
        * [Calico Over Ethernet Fabrics](calico/reference/architecture/design/l2-interconnect-fabric.md)
        * [Calico Over IP Fabrics](calico/reference/architecture/design/l3-interconnect-fabric.md)
      * [The Calico Data Path IP Routing And Iptables](calico/reference/architecture/data-path.md)
      * [Architecture](calico/reference/architecture/index.md)
      * [Component Architecture](calico/reference/architecture/overview.md)
    - Calicoctl
      - Datastore
        - Migrate
          * [Calicoctl Datastore Migrate Export](calico/reference/calicoctl/datastore/migrate/export.md)
          * [Calicoctl Datastore Migrate Import](calico/reference/calicoctl/datastore/migrate/import.md)
          * [Index](calico/reference/calicoctl/datastore/migrate/index.md)
          * [Calicoctl Datastore Migrate Lock](calico/reference/calicoctl/datastore/migrate/lock.md)
          * [Calicoctl Datastore Migrate](calico/reference/calicoctl/datastore/migrate/overview.md)
          * [Calicoctl Datastore Migrate Unlock](calico/reference/calicoctl/datastore/migrate/unlock.md)
        * [Index](calico/reference/calicoctl/datastore/index.md)
        * [Calicoctl Datastore](calico/reference/calicoctl/datastore/overview.md)
      - Ipam
        * [Calicoctl Ipam Check](calico/reference/calicoctl/ipam/check.md)
        * [Calicoctl Ipam](calico/reference/calicoctl/ipam/show.md)
        * [Index](calico/reference/calicoctl/ipam/index.md)
      - Node
        * [Calicoctl Node Checksystem](calico/reference/calicoctl/node/checksystem.md)
        * [Calicoctl Node Diags](calico/reference/calicoctl/node/diags.md)
        * [Index](calico/reference/calicoctl/node/index.md)
        * [Calicoctl Node](calico/reference/calicoctl/node/overview.md)
        * [Calicoctl Node Run](calico/reference/calicoctl/node/run.md)
        * [Calicoctl Node Status](calico/reference/calicoctl/node/status.md)
      * [Calicoctl Apply](calico/reference/calicoctl/apply.md)
      * [Calicoctl Convert](calico/reference/calicoctl/convert.md)
      * [Calicoctl Create](calico/reference/calicoctl/create.md)
      * [Calicoctl Delete](calico/reference/calicoctl/delete.md)
      * [Calicoctl Get](calico/reference/calicoctl/get.md)
      * [Index](calico/reference/calicoctl/index.md)
      * [Calicoctl Label](calico/reference/calicoctl/label.md)
      * [Calicoctl User Reference](calico/reference/calicoctl/overview.md)
      * [Calicoctl Patch](calico/reference/calicoctl/patch.md)
      * [Calicoctl Replace](calico/reference/calicoctl/replace.md)
      * [Calicoctl Version](calico/reference/calicoctl/version.md)
    - Cni Plugin
      * [Configure The Calico CNI Plugins](calico/reference/cni-plugin/configuration.md)
    - Dikastes
      * [Configuring Dikastes](calico/reference/dikastes/configuration.md)
    - Etcd Rbac
      * [Calico Key And Path Prefixes](calico/reference/etcd-rbac/calico-etcdv3-paths.md)
      * [Generating Certificates](calico/reference/etcd-rbac/certificate-generation.md)
      * [Index](calico/reference/etcd-rbac/index.md)
      * [Segmenting Etcd On Kubernetes Advanced](calico/reference/etcd-rbac/kubernetes-advanced.md)
      * [Segmenting Etcd On Kubernetes Basic](calico/reference/etcd-rbac/kubernetes.md)
      * [Setting Up Etcd Certificates For RBAC](calico/reference/etcd-rbac/overview.md)
      * [Creating Users And Roles](calico/reference/etcd-rbac/users-and-roles.md)
    - Felix
      * [Configuring Felix](calico/reference/felix/configuration.md)
      * [Felix](calico/reference/felix/index.md)
      * [Prometheus Statistics](calico/reference/felix/prometheus.md)
    - Host Endpoints
      * [Creating Policy For Basic Connectivity](calico/reference/host-endpoints/connectivity.md)
      * [Connection Tracking](calico/reference/host-endpoints/conntrack.md)
      * [Failsafe Rules](calico/reference/host-endpoints/failsafe.md)
      * [Apply On Forwarded Traffic](calico/reference/host-endpoints/forwarded.md)
      * [Index](calico/reference/host-endpoints/index.md)
      * [Creating Host Endpoint Objects](calico/reference/host-endpoints/objects.md)
      * [Host Endpoints](calico/reference/host-endpoints/overview.md)
      * [Pre DNAT Policy](calico/reference/host-endpoints/pre-dnat.md)
      * [Selector Based Policies](calico/reference/host-endpoints/selector.md)
      * [Summary](calico/reference/host-endpoints/summary.md)
    - [Installation](calico/reference/installation/README.md)
    - [Kube Controllers](calico/reference/kube-controllers/index.md)
      * [Configuring The Calico Kubernetes Controllers](calico/reference/kube-controllers/configuration.md)
      * [Prometheus Statistics](calico/reference/kube-controllers/prometheus.md)
    - Legal
      * [Application Layer Policy Attributions](calico/reference/legal/alp.md)
      * [Calicoctl Attributions](calico/reference/legal/calicoctl.md)
      * [CNI Plugin Attributions](calico/reference/legal/cni.md)
      * [Confd Attributions](calico/reference/legal/confd.md)
      * [Felix Attributions](calico/reference/legal/felix.md)
      * [Attributions](calico/reference/legal/index.md)
      * [Kubernetes Controllers Attributions](calico/reference/legal/kube-controllers.md)
      * [Calico Node Attributions](calico/reference/legal/node.md)
      * [Typha Attributions](calico/reference/legal/typha.md)
    - Node
      * [Configuring Calico Node](calico/reference/node/configuration.md)
    - Public Cloud
      * [Amazon Web Services](calico/reference/public-cloud/aws.md)
      * [Azure](calico/reference/public-cloud/azure.md)
      * [Google Compute Engine](calico/reference/public-cloud/gce.md)
      * [IBM Cloud](calico/reference/public-cloud/ibm.md)
      * [Configuration On Public Clouds](calico/reference/public-cloud/index.md)
    - Resources
      * [BGP Configuration](calico/reference/resources/bgpconfig.md)
      * [BGP Peer](calico/reference/resources/bgppeer.md)
      * [Calico Node Status](calico/reference/resources/caliconodestatus.md)
      * [Felix Configuration](calico/reference/resources/felixconfig.md)
      * [Global Network Policy](calico/reference/resources/globalnetworkpolicy.md)
      * [Global Network Set](calico/reference/resources/globalnetworkset.md)
      * [Host Endpoint](calico/reference/resources/hostendpoint.md)
      * [Index](calico/reference/resources/index.md)
      * [IP Pool](calico/reference/resources/ippool.md)
      * [IP Reservation](calico/reference/resources/ipreservation.md)
      * [Kubernetes Controllers Configuration](calico/reference/resources/kubecontrollersconfig.md)
      * [Network Policy](calico/reference/resources/networkpolicy.md)
      * [Network Set](calico/reference/resources/networkset.md)
      * [Node](calico/reference/resources/node.md)
      * [Resource Definitions](calico/reference/resources/overview.md)
      * [Profile](calico/reference/resources/profile.md)
      * [Workload Endpoint](calico/reference/resources/workloadendpoint.md)
    - Typha
      * [Configuring Typha](calico/reference/typha/configuration.md)
      * [Typha](calico/reference/typha/index.md)
      * [Typha Overview](calico/reference/typha/overview.md)
      * [Prometheus Statistics](calico/reference/typha/prometheus.md)
    - Vpp
      * [Host Network Configuration](calico/reference/vpp/host-network.md)
      * [VPP Dataplane](calico/reference/vpp/index.md)
      * [VPP Dataplane Implementation Details](calico/reference/vpp/technical-details.md)
      * [Primary Interface Configuration](calico/reference/vpp/uplink-configuration.md)
    * [Calico API](calico/reference/api.md)
    * [Frequently Asked Questions](calico/reference/faq.md)
    * [Reference](calico/reference/index.md)
    * [Getting Involved](calico/reference/involved.md)
  - Release Notes
    * [Release Notes](calico/release-notes/index.md)
  - Security
    - Calico Enterprise
      * [Advanced Compliance Controls](calico/security/calico-enterprise/compliance.md)
      * [Advanced Egress Access Controls](calico/security/calico-enterprise/egress-access-controls.md)
      * [Federation](calico/security/calico-enterprise/federation.md)
      * [Calico Enterprise](calico/security/calico-enterprise/index.md)
      * [Network Visibility](calico/security/calico-enterprise/network-visibility.md)
      * [Policy Workflow](calico/security/calico-enterprise/policy-workflow.md)
      * [Threat Defense](calico/security/calico-enterprise/threat-defense.md)
      * [Calico Enterprise Manager User Interface](calico/security/calico-enterprise/user-console.md)
    - Comms
      * [Configure Encryption And Authentication To Secure Calico Components](calico/security/comms/crypto-auth.md)
      * [Secure Calico Component Communications](calico/security/comms/index.md)
      * [Schedule Typha For Scaling To Well Known Nodes](calico/security/comms/reduce-nodes.md)
      * [Secure BGP Sessions](calico/security/comms/secure-bgp.md)
      * [Secure Calico Prometheus Endpoints](calico/security/comms/secure-metrics.md)
    - Tutorials
      - App Layer Policy
        * [Enforce Calico Network Policy Using Istio Tutorial](calico/security/tutorials/app-layer-policy/enforce-policy-istio.md)
      - Kubernetes Policy Demo
        * [Kubernetes Policy Demo](calico/security/tutorials/kubernetes-policy-demo/kubernetes-demo.md)
      * [Calico Policy Tutorial](calico/security/tutorials/calico-policy.md)
      * [Kubernetes Policy Advanced Tutorial](calico/security/tutorials/kubernetes-policy-advanced.md)
      * [Kubernetes Policy Basic Tutorial](calico/security/tutorials/kubernetes-policy-basic.md)
      * [Protect Hosts Tutorial](calico/security/tutorials/protect-hosts.md)
    * [Adopt A Zero Trust Network Model For Security](calico/security/adopt-zero-trust.md)
    * [Enforce Network Policy For Istio](calico/security/app-layer-policy.md)
    * [Get Started With Calico Network Policy](calico/security/calico-network-policy.md)
    * [Calico Policy](calico/security/calico-policy.md)
    * [Defend Against Do S Attacks](calico/security/defend-dos-attack.md)
    * [Encrypt In Cluster Pod Traffic](calico/security/encrypt-cluster-pod-traffic.md)
    * [Use External I Ps Or Networks Rules In Policy](calico/security/external-ips-policy.md)
    * [Policy For Extreme Traffic](calico/security/extreme-traffic.md)
    * [Get Started With Policy](calico/security/get-started.md)
    * [Enable Extreme High Connection Workloads](calico/security/high-connection-workloads.md)
    * [Apply Policy To Forwarded Traffic](calico/security/host-forwarded-traffic.md)
    * [Policy For Hosts](calico/security/hosts.md)
    * [Use HTTP Methods And Paths In Policy Rules](calico/security/http-methods.md)
    * [Use ICMP Ping Rules In Policy](calico/security/icmp-ping.md)
    * [Security](calico/security/index.md)
    * [Policy For Istio](calico/security/istio.md)
    * [Enable Default Deny For Kubernetes Pods](calico/security/kubernetes-default-deny.md)
    * [Get Started With Kubernetes Network Policy](calico/security/kubernetes-network-policy.md)
    * [Apply Calico Policy To Kubernetes Node Ports](calico/security/kubernetes-node-ports.md)
    * [Protect Kubernetes Nodes](calico/security/kubernetes-nodes.md)
  - Upgrade
    - [v2.5](calico/upgrade/v2.5/README.md)
  * [AUTHORS](calico/AUTHORS.md)
  * [DOC STYLE GUIDE](calico/DOC_STYLE_GUIDE.md)
  * [Project Calico Documentation Archives](calico/releases.md)
- [Calicoctl](calicoctl/README.md)
  - Test Data
    - Convert
      - Input
        - V 1
          - [Migration](calicoctl/test-data/convert/input/v1/migration/README.md)
  * [CONTRIBUTING](calicoctl/CONTRIBUTING.md)
- [Cni Plugin](cni-plugin/README.md)
  * [Configuration](cni-plugin/configuration.md)
- [Confd](confd/README.md)
  - Docs
    * [Command Line Flags](confd/docs/command-line-flags.md)
    * [Configuration Guide](confd/docs/configuration-guide.md)
    * [Dns Srv Records](confd/docs/dns-srv-records.md)
    * [Installation](confd/docs/installation.md)
    * [Logging](confd/docs/logging.md)
    * [Noop Mode](confd/docs/noop-mode.md)
    * [Quick Start Guide](confd/docs/quick-start-guide.md)
    * [Release Checklist](confd/docs/release-checklist.md)
    * [Template Resources](confd/docs/template-resources.md)
    * [Templates](confd/docs/templates.md)
    * [Tomcat Sample](confd/docs/tomcat-sample.md)
  - Etc
    - Calico
      - Confd
        - [Templates](confd/etc/calico/confd/templates/README.md)
  - [Tests](confd/tests/README.md)
  * [CONTRIBUTING](confd/CONTRIBUTING.md)
- [Felix](felix/README.md)
  - Bpf Gpl
    - Include
      - [Libbpf](felix/bpf-gpl/include/libbpf/README.md)
  - [K 8 Sfv](felix/k8sfv/README.md)
    * [Monitoring](felix/k8sfv/monitoring.md)
  * [CHANGES](felix/CHANGES.md)
  * [CONTRIBUTING](felix/CONTRIBUTING.md)
- Hack
  - Release
    * [RELEASING](hack/release/RELEASING.md)
  - Test
    - [Certs](hack/test/certs/README.md)
- [Kube Controllers](kube-controllers/README.md)
- [Libcalico Go](libcalico-go/README.md)
  - Docs
    * [Data Model](libcalico-go/docs/data-model.md)
  - Lib
    - Apis
      - crd.projectcalico.org
        - [V 1](libcalico-go/lib/apis/crd.projectcalico.org/v1/README.md)
  - [Test](libcalico-go/test/README.md)
- [Networking Calico](networking-calico/README.md)
  * [CONTRIBUTING](networking-calico/CONTRIBUTING.md)
- [Node](node/README.md)
- [Pod 2 Daemon](pod2daemon/README.md)
- [Typha](typha/README.md)
  * [CONTRIBUTING](typha/CONTRIBUTING.md)
* [CONTRIBUTING CODE](CONTRIBUTING_CODE.md)
* [CONTRIBUTING DOCS](CONTRIBUTING_DOCS.md)
* [DEVELOPER GUIDE](DEVELOPER_GUIDE.md)