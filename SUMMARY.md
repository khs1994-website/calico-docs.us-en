# [Calico 英文文档](https://github.com/projectcalico/calico)

- [Api](api/README.md)
- [Apiserver](apiserver/README.md)
- [App Policy](app-policy/README.md)
- [Calico](calico/README.md)
- [Calicoctl](calicoctl/README.md)
- Charts
  - [Calico](charts/calico/README.md)
  - [crd.projectcalico.org.v1](charts/crd.projectcalico.org.v1/README.md)
  - [projectcalico.org.v3](charts/projectcalico.org.v3/README.md)
  - [Tigera Operator](charts/tigera-operator/README.md)
- [Cni Plugin](cni-plugin/README.md)
  - Pkg
    - Ipamplugin
      * [DESIGN](cni-plugin/pkg/ipamplugin/DESIGN.md)
  * [Configuration](cni-plugin/configuration.md)
- [Confd](confd/README.md)
  - Docs
    * [Bgpfilter Bird Config Examples](confd/docs/bgpfilter-bird-config-examples.md)
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
- Design
  - Ipam
    * [DESIGN](design/ipam/DESIGN.md)
    * [Ipam Cni](design/ipam/ipam-cni.md)
    * [Ipam Core Library](design/ipam/ipam-core-library.md)
    * [Ipam Datastore](design/ipam/ipam-datastore.md)
    * [Ipam Gc](design/ipam/ipam-gc.md)
    * [Ipam Other Callers](design/ipam/ipam-other-callers.md)
- [E 2 E](e2e/README.md)
  - Images
    - [Rapidclient](e2e/images/rapidclient/README.md)
      * [DESIGN](e2e/images/rapidclient/DESIGN.md)
- [Felix](felix/README.md)
  - Design
    * [Bpf Conntrack Flowstate](felix/design/bpf-conntrack-flowstate.md)
    * [Bpf Encap Fragments Icmp](felix/design/bpf-encap-fragments-icmp.md)
    * [Bpf Host Networking](felix/design/bpf-host-networking.md)
    * [Bpf Observability](felix/design/bpf-observability.md)
    * [Bpf Overview](felix/design/bpf-overview.md)
    * [Bpf Services](felix/design/bpf-services.md)
    * [Bpf Tc Programs](felix/design/bpf-tc-programs.md)
    * [Bpf Tests](felix/design/bpf-tests.md)
    * [Bpf Xdp](felix/design/bpf-xdp.md)
    * [Calc Graph](felix/design/calc-graph.md)
    * [Dataplane](felix/design/dataplane.md)
  - Docs
    * [Calc Graph Diagram](felix/docs/calc-graph-diagram.md)
    * [Config Params](felix/docs/config-params.md)
  - [K 8 Sfv](felix/k8sfv/README.md)
    * [Monitoring](felix/k8sfv/monitoring.md)
  - [Nfnetlink](felix/nfnetlink/README.md)
  * [CHANGES](felix/CHANGES.md)
  * [CLAUDE](felix/CLAUDE.md)
  * [DESIGN](felix/DESIGN.md)
- [Goldmane](goldmane/README.md)
  * [CLAUDE](goldmane/CLAUDE.md)
  * [DESIGN](goldmane/DESIGN.md)
- Hack
  - Cmd
    - [Ipam Hammer](hack/cmd/ipam-hammer/README.md)
  - Docs
    * [Adding An Api](hack/docs/adding-an-api.md)
  - [Perf](hack/perf/README.md)
  - Rpms
    - [Nftables](hack/rpms/nftables/README.md)
  - Test
    - [Certs](hack/test/certs/README.md)
- [Key Cert Provisioner](key-cert-provisioner/README.md)
  - [Test Signer](key-cert-provisioner/test-signer/README.md)
- [Kube Controllers](kube-controllers/README.md)
  - Pkg
    - Controllers
      - Node
        * [DESIGN](kube-controllers/pkg/controllers/node/DESIGN.md)
- [Lib](lib/README.md)
- [Libcalico Go](libcalico-go/README.md)
  - Docs
    * [Data Model](libcalico-go/docs/data-model.md)
  - Lib
    - Apis
      - crd.projectcalico.org
        - [V 1](libcalico-go/lib/apis/crd.projectcalico.org/v1/README.md)
      - [Internalapi](libcalico-go/lib/apis/internalapi/README.md)
      - [V 1](libcalico-go/lib/apis/v1/README.md)
    - Ipam
      * [DESIGN](libcalico-go/lib/ipam/DESIGN.md)
  - [Test](libcalico-go/test/README.md)
- [Manifests](manifests/README.md)
- [Networking Calico](networking-calico/README.md)
- [Node](node/README.md)
  - Filesystem
    - Etc
      - Calico
        - Confd
          - [Templates](node/filesystem/etc/calico/confd/templates/README.md)
  * [DESIGN](node/DESIGN.md)
- [Pod 2 Daemon](pod2daemon/README.md)
- Process
  - Testing
    - [Aso](process/testing/aso/README.md)
    - [Win Connections](process/testing/win-connections/README.md)
- [Release](release/README.md)
  - [Packaging](release/packaging/README.md)
  * [RELEASING](release/RELEASING.md)
- Third Party
  - [Cni Plugins](third_party/cni-plugins/README.md)
- [Typha](typha/README.md)
  * [DESIGN](typha/DESIGN.md)
- [Whisker](whisker/README.md)
  * [Pull Request Template](whisker/pull_request_template.md)
* [AUTHORS](AUTHORS.md)
* [CONTRIBUTING DOCS](CONTRIBUTING_DOCS.md)
* [CONTRIBUTING](CONTRIBUTING.md)
* [DESIGN](DESIGN.md)
* [DEVELOPER GUIDE](DEVELOPER_GUIDE.md)
* [LICENSE](LICENSE.md)
* [SECURITY](SECURITY.md)
