// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.cs;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.cs.EdgeKubernetesArgs;
import com.pulumi.alicloud.cs.inputs.EdgeKubernetesState;
import com.pulumi.alicloud.cs.outputs.EdgeKubernetesAddon;
import com.pulumi.alicloud.cs.outputs.EdgeKubernetesCertificateAuthority;
import com.pulumi.alicloud.cs.outputs.EdgeKubernetesConnections;
import com.pulumi.alicloud.cs.outputs.EdgeKubernetesLogConfig;
import com.pulumi.alicloud.cs.outputs.EdgeKubernetesRuntime;
import com.pulumi.alicloud.cs.outputs.EdgeKubernetesWorkerDataDisk;
import com.pulumi.alicloud.cs.outputs.EdgeKubernetesWorkerNode;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.Object;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * This resource will help you to manage a Edge Kubernetes Cluster in Alibaba Cloud Kubernetes Service, see [What is edge kubernetes](https://www.alibabacloud.com/help/en/ack/ack-managed-and-ack-dedicated/developer-reference/create-an-ack-edge-cluster).
 * 
 * &gt; **NOTE:** Kubernetes cluster only supports VPC network and it can access internet while creating kubernetes cluster.
 * A Nat Gateway and configuring a SNAT for it can ensure one VPC network access internet. If there is no nat gateway in the
 * VPC, you can set `new_nat_gateway` to &#34;true&#34; to create one automatically.
 * 
 * &gt; **NOTE:** Creating kubernetes cluster need to install several packages and it will cost about 15 minutes. Please be patient.
 * 
 * &gt; **NOTE:** The provider supports to download kube config, client certificate, client key and cluster ca certificate
 * after creating cluster successfully, and you can put them into the specified location, like &#39;~/.kube/config&#39;.
 * 
 * &gt; **NOTE:** The provider supports disabling internet load balancer for API Server by setting `false` to `slb_internet_enabled`.
 * 
 * &gt; **NOTE:** If you want to manage Kubernetes, you can use Kubernetes Provider.
 * 
 * &gt; **NOTE:** Available since v1.103.0.
 * 
 * &gt; **NOTE:** From version 1.185.0+, support new fields `cluster_spec`, `runtime` and `load_balancer_spec`.
 * 
 * ## Example Usage
 * 
 * Basic Usage
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.alicloud.AlicloudFunctions;
 * import com.pulumi.alicloud.inputs.GetZonesArgs;
 * import com.pulumi.alicloud.ecs.EcsFunctions;
 * import com.pulumi.alicloud.ecs.inputs.GetInstanceTypesArgs;
 * import com.pulumi.alicloud.vpc.Network;
 * import com.pulumi.alicloud.vpc.NetworkArgs;
 * import com.pulumi.alicloud.vpc.Switch;
 * import com.pulumi.alicloud.vpc.SwitchArgs;
 * import com.pulumi.alicloud.cs.EdgeKubernetes;
 * import com.pulumi.alicloud.cs.EdgeKubernetesArgs;
 * import com.pulumi.alicloud.cs.inputs.EdgeKubernetesWorkerDataDiskArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         final var config = ctx.config();
 *         final var name = config.get(&#34;name&#34;).orElse(&#34;tf-example-basic-edge&#34;);
 *         final var default = AlicloudFunctions.getZones(GetZonesArgs.builder()
 *             .availableResourceCreation(&#34;VSwitch&#34;)
 *             .build());
 * 
 *         final var defaultGetInstanceTypes = EcsFunctions.getInstanceTypes(GetInstanceTypesArgs.builder()
 *             .availabilityZone(default_.zones()[0].id())
 *             .cpuCoreCount(4)
 *             .memorySize(8)
 *             .kubernetesNodeRole(&#34;Master&#34;)
 *             .build());
 * 
 *         var defaultNetwork = new Network(&#34;defaultNetwork&#34;, NetworkArgs.builder()        
 *             .vpcName(name)
 *             .cidrBlock(&#34;10.4.0.0/16&#34;)
 *             .build());
 * 
 *         var defaultSwitch = new Switch(&#34;defaultSwitch&#34;, SwitchArgs.builder()        
 *             .vswitchName(name)
 *             .cidrBlock(&#34;10.4.0.0/24&#34;)
 *             .vpcId(defaultNetwork.id())
 *             .zoneId(default_.zones()[0].id())
 *             .build());
 * 
 *         var defaultEdgeKubernetes = new EdgeKubernetes(&#34;defaultEdgeKubernetes&#34;, EdgeKubernetesArgs.builder()        
 *             .name(name)
 *             .workerVswitchIds(defaultSwitch.id())
 *             .workerInstanceTypes(defaultGetInstanceTypes.applyValue(getInstanceTypesResult -&gt; getInstanceTypesResult.instanceTypes()[0].id()))
 *             .workerNumber(&#34;1&#34;)
 *             .password(&#34;Test12345&#34;)
 *             .podCidr(&#34;10.99.0.0/16&#34;)
 *             .serviceCidr(&#34;172.16.0.0/16&#34;)
 *             .workerInstanceChargeType(&#34;PostPaid&#34;)
 *             .newNatGateway(&#34;true&#34;)
 *             .nodeCidrMask(&#34;24&#34;)
 *             .installCloudMonitor(&#34;true&#34;)
 *             .slbInternetEnabled(&#34;true&#34;)
 *             .isEnterpriseSecurityGroup(&#34;true&#34;)
 *             .workerDataDisks(EdgeKubernetesWorkerDataDiskArgs.builder()
 *                 .category(&#34;cloud_ssd&#34;)
 *                 .size(&#34;200&#34;)
 *                 .encrypted(&#34;false&#34;)
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * You could create a professional kubernetes edge cluster now.
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.alicloud.AlicloudFunctions;
 * import com.pulumi.alicloud.inputs.GetZonesArgs;
 * import com.pulumi.alicloud.ecs.EcsFunctions;
 * import com.pulumi.alicloud.ecs.inputs.GetInstanceTypesArgs;
 * import com.pulumi.alicloud.vpc.Network;
 * import com.pulumi.alicloud.vpc.NetworkArgs;
 * import com.pulumi.alicloud.vpc.Switch;
 * import com.pulumi.alicloud.vpc.SwitchArgs;
 * import com.pulumi.alicloud.cs.EdgeKubernetes;
 * import com.pulumi.alicloud.cs.EdgeKubernetesArgs;
 * import com.pulumi.alicloud.cs.inputs.EdgeKubernetesWorkerDataDiskArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         final var config = ctx.config();
 *         final var name = config.get(&#34;name&#34;).orElse(&#34;tf_example&#34;);
 *         final var default = AlicloudFunctions.getZones(GetZonesArgs.builder()
 *             .availableResourceCreation(&#34;VSwitch&#34;)
 *             .build());
 * 
 *         final var defaultGetInstanceTypes = EcsFunctions.getInstanceTypes(GetInstanceTypesArgs.builder()
 *             .availabilityZone(default_.zones()[0].id())
 *             .cpuCoreCount(4)
 *             .memorySize(8)
 *             .kubernetesNodeRole(&#34;Master&#34;)
 *             .build());
 * 
 *         var defaultNetwork = new Network(&#34;defaultNetwork&#34;, NetworkArgs.builder()        
 *             .vpcName(name)
 *             .cidrBlock(&#34;10.4.0.0/16&#34;)
 *             .build());
 * 
 *         var defaultSwitch = new Switch(&#34;defaultSwitch&#34;, SwitchArgs.builder()        
 *             .vswitchName(name)
 *             .cidrBlock(&#34;10.4.0.0/24&#34;)
 *             .vpcId(defaultNetwork.id())
 *             .zoneId(default_.zones()[0].id())
 *             .build());
 * 
 *         var defaultEdgeKubernetes = new EdgeKubernetes(&#34;defaultEdgeKubernetes&#34;, EdgeKubernetesArgs.builder()        
 *             .name(name)
 *             .workerVswitchIds(defaultSwitch.id())
 *             .workerInstanceTypes(defaultGetInstanceTypes.applyValue(getInstanceTypesResult -&gt; getInstanceTypesResult.instanceTypes()[0].id()))
 *             .clusterSpec(&#34;ack.pro.small&#34;)
 *             .workerNumber(&#34;1&#34;)
 *             .password(&#34;Test12345&#34;)
 *             .podCidr(&#34;10.99.0.0/16&#34;)
 *             .serviceCidr(&#34;172.16.0.0/16&#34;)
 *             .workerInstanceChargeType(&#34;PostPaid&#34;)
 *             .newNatGateway(&#34;true&#34;)
 *             .nodeCidrMask(&#34;24&#34;)
 *             .loadBalancerSpec(&#34;slb.s2.small&#34;)
 *             .installCloudMonitor(&#34;true&#34;)
 *             .slbInternetEnabled(&#34;true&#34;)
 *             .isEnterpriseSecurityGroup(&#34;true&#34;)
 *             .workerDataDisks(EdgeKubernetesWorkerDataDiskArgs.builder()
 *                 .category(&#34;cloud_ssd&#34;)
 *                 .size(&#34;200&#34;)
 *                 .encrypted(&#34;false&#34;)
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * Kubernetes edge cluster can be imported using the id, e.g. Then complete the main.tf accords to the result of `pulumi preview`.
 * 
 * ```sh
 * $ pulumi import alicloud:cs/edgeKubernetes:EdgeKubernetes main cluster-id
 * ```
 * 
 */
@ResourceType(type="alicloud:cs/edgeKubernetes:EdgeKubernetes")
public class EdgeKubernetes extends com.pulumi.resources.CustomResource {
    /**
     * The addon you want to install in cluster. See `addons` below.
     * 
     */
    @Export(name="addons", refs={List.class,EdgeKubernetesAddon.class}, tree="[0,1]")
    private Output</* @Nullable */ List<EdgeKubernetesAddon>> addons;

    /**
     * @return The addon you want to install in cluster. See `addons` below.
     * 
     */
    public Output<Optional<List<EdgeKubernetesAddon>>> addons() {
        return Codegen.optional(this.addons);
    }
    /**
     * The ID of availability zone.
     * 
     */
    @Export(name="availabilityZone", refs={String.class}, tree="[0]")
    private Output<String> availabilityZone;

    /**
     * @return The ID of availability zone.
     * 
     */
    public Output<String> availabilityZone() {
        return this.availabilityZone;
    }
    /**
     * Nested attribute containing certificate authority data for your cluster.
     * 
     */
    @Export(name="certificateAuthority", refs={EdgeKubernetesCertificateAuthority.class}, tree="[0]")
    private Output<EdgeKubernetesCertificateAuthority> certificateAuthority;

    /**
     * @return Nested attribute containing certificate authority data for your cluster.
     * 
     */
    public Output<EdgeKubernetesCertificateAuthority> certificateAuthority() {
        return this.certificateAuthority;
    }
    /**
     * The path of client certificate, like `~/.kube/client-cert.pem`.
     * 
     */
    @Export(name="clientCert", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> clientCert;

    /**
     * @return The path of client certificate, like `~/.kube/client-cert.pem`.
     * 
     */
    public Output<Optional<String>> clientCert() {
        return Codegen.optional(this.clientCert);
    }
    /**
     * The path of client key, like `~/.kube/client-key.pem`.
     * 
     */
    @Export(name="clientKey", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> clientKey;

    /**
     * @return The path of client key, like `~/.kube/client-key.pem`.
     * 
     */
    public Output<Optional<String>> clientKey() {
        return Codegen.optional(this.clientKey);
    }
    /**
     * The path of cluster ca certificate, like `~/.kube/cluster-ca-cert.pem`
     * 
     * *Removed params*
     * 
     */
    @Export(name="clusterCaCert", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> clusterCaCert;

    /**
     * @return The path of cluster ca certificate, like `~/.kube/cluster-ca-cert.pem`
     * 
     * *Removed params*
     * 
     */
    public Output<Optional<String>> clusterCaCert() {
        return Codegen.optional(this.clusterCaCert);
    }
    /**
     * The cluster specifications of kubernetes cluster,which can be empty. Valid values:
     * * ack.standard : Standard edge clusters.
     * * ack.pro.small : Professional edge clusters.
     * 
     */
    @Export(name="clusterSpec", refs={String.class}, tree="[0]")
    private Output<String> clusterSpec;

    /**
     * @return The cluster specifications of kubernetes cluster,which can be empty. Valid values:
     * * ack.standard : Standard edge clusters.
     * * ack.pro.small : Professional edge clusters.
     * 
     */
    public Output<String> clusterSpec() {
        return this.clusterSpec;
    }
    /**
     * Map of kubernetes cluster connection information.
     * 
     */
    @Export(name="connections", refs={EdgeKubernetesConnections.class}, tree="[0]")
    private Output<EdgeKubernetesConnections> connections;

    /**
     * @return Map of kubernetes cluster connection information.
     * 
     */
    public Output<EdgeKubernetesConnections> connections() {
        return this.connections;
    }
    /**
     * Whether to enable cluster deletion protection.
     * 
     */
    @Export(name="deletionProtection", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> deletionProtection;

    /**
     * @return Whether to enable cluster deletion protection.
     * 
     */
    public Output<Optional<Boolean>> deletionProtection() {
        return Codegen.optional(this.deletionProtection);
    }
    /**
     * Default false, when you want to change `vpc_id`, you have to set this field to true, then the cluster will be recreated.
     * 
     */
    @Export(name="forceUpdate", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> forceUpdate;

    /**
     * @return Default false, when you want to change `vpc_id`, you have to set this field to true, then the cluster will be recreated.
     * 
     */
    public Output<Optional<Boolean>> forceUpdate() {
        return Codegen.optional(this.forceUpdate);
    }
    /**
     * Install cloud monitor agent on ECS. default: `true`.
     * 
     */
    @Export(name="installCloudMonitor", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> installCloudMonitor;

    /**
     * @return Install cloud monitor agent on ECS. default: `true`.
     * 
     */
    public Output<Optional<Boolean>> installCloudMonitor() {
        return Codegen.optional(this.installCloudMonitor);
    }
    /**
     * Enable to create advanced security group. default: false. See [Advanced security group](https://www.alibabacloud.com/help/doc-detail/120621.htm).
     * 
     */
    @Export(name="isEnterpriseSecurityGroup", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> isEnterpriseSecurityGroup;

    /**
     * @return Enable to create advanced security group. default: false. See [Advanced security group](https://www.alibabacloud.com/help/doc-detail/120621.htm).
     * 
     */
    public Output<Boolean> isEnterpriseSecurityGroup() {
        return this.isEnterpriseSecurityGroup;
    }
    /**
     * The keypair of ssh login cluster node, you have to create it first. You have to specify one of `password` `key_name` `kms_encrypted_password` fields.
     * 
     */
    @Export(name="keyName", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> keyName;

    /**
     * @return The keypair of ssh login cluster node, you have to create it first. You have to specify one of `password` `key_name` `kms_encrypted_password` fields.
     * 
     */
    public Output<Optional<String>> keyName() {
        return Codegen.optional(this.keyName);
    }
    /**
     * The path of kube config, like `~/.kube/config`.
     * 
     * @deprecated
     * Field &#39;kube_config&#39; has been deprecated from provider version 1.187.0. New DataSource &#39;alicloud_cs_cluster_credential&#39; manage your cluster&#39;s kube config.
     * 
     */
    @Deprecated /* Field 'kube_config' has been deprecated from provider version 1.187.0. New DataSource 'alicloud_cs_cluster_credential' manage your cluster's kube config. */
    @Export(name="kubeConfig", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> kubeConfig;

    /**
     * @return The path of kube config, like `~/.kube/config`.
     * 
     */
    public Output<Optional<String>> kubeConfig() {
        return Codegen.optional(this.kubeConfig);
    }
    /**
     * The cluster api server load balance instance specification. For more information on how to select a LB instance specification, see [SLB instance overview](https://help.aliyun.com/document_detail/85931.html).
     * -&gt;NOTE: If you want to use `Flannel` as CNI network plugin, You need to specific the `pod_cidr` field and addons with `flannel`.
     * 
     * *Worker params*
     * 
     */
    @Export(name="loadBalancerSpec", refs={String.class}, tree="[0]")
    private Output<String> loadBalancerSpec;

    /**
     * @return The cluster api server load balance instance specification. For more information on how to select a LB instance specification, see [SLB instance overview](https://help.aliyun.com/document_detail/85931.html).
     * -&gt;NOTE: If you want to use `Flannel` as CNI network plugin, You need to specific the `pod_cidr` field and addons with `flannel`.
     * 
     * *Worker params*
     * 
     */
    public Output<String> loadBalancerSpec() {
        return this.loadBalancerSpec;
    }
    /**
     * A list of one element containing information about the associated log store. See `log_config` below.
     * 
     * @deprecated
     * Field &#39;log_config&#39; has been removed from provider version 1.103.0. New field &#39;addons&#39; replaces it.
     * 
     */
    @Deprecated /* Field 'log_config' has been removed from provider version 1.103.0. New field 'addons' replaces it. */
    @Export(name="logConfig", refs={EdgeKubernetesLogConfig.class}, tree="[0]")
    private Output</* @Nullable */ EdgeKubernetesLogConfig> logConfig;

    /**
     * @return A list of one element containing information about the associated log store. See `log_config` below.
     * 
     */
    public Output<Optional<EdgeKubernetesLogConfig>> logConfig() {
        return Codegen.optional(this.logConfig);
    }
    /**
     * The kubernetes cluster&#39;s name. It is unique in one Alicloud account.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The kubernetes cluster&#39;s name. It is unique in one Alicloud account.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    @Export(name="namePrefix", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> namePrefix;

    public Output<Optional<String>> namePrefix() {
        return Codegen.optional(this.namePrefix);
    }
    /**
     * The ID of nat gateway used to launch kubernetes cluster.
     * 
     */
    @Export(name="natGatewayId", refs={String.class}, tree="[0]")
    private Output<String> natGatewayId;

    /**
     * @return The ID of nat gateway used to launch kubernetes cluster.
     * 
     */
    public Output<String> natGatewayId() {
        return this.natGatewayId;
    }
    /**
     * Whether to create a new nat gateway while creating kubernetes cluster. Default to true. Then openapi in Alibaba Cloud are not all on intranet, So turn this option on is a good choice.
     * 
     */
    @Export(name="newNatGateway", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> newNatGateway;

    /**
     * @return Whether to create a new nat gateway while creating kubernetes cluster. Default to true. Then openapi in Alibaba Cloud are not all on intranet, So turn this option on is a good choice.
     * 
     */
    public Output<Optional<Boolean>> newNatGateway() {
        return Codegen.optional(this.newNatGateway);
    }
    /**
     * The node cidr block to specific how many pods can run on single node. 24-28 is allowed. 24 means 2^(32-24)-1=255 and the node can run at most 255 pods. default: 24
     * 
     */
    @Export(name="nodeCidrMask", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> nodeCidrMask;

    /**
     * @return The node cidr block to specific how many pods can run on single node. 24-28 is allowed. 24 means 2^(32-24)-1=255 and the node can run at most 255 pods. default: 24
     * 
     */
    public Output<Optional<Integer>> nodeCidrMask() {
        return Codegen.optional(this.nodeCidrMask);
    }
    /**
     * The password of ssh login cluster node. You have to specify one of `password`, `key_name` `kms_encrypted_password` fields.
     * 
     */
    @Export(name="password", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> password;

    /**
     * @return The password of ssh login cluster node. You have to specify one of `password`, `key_name` `kms_encrypted_password` fields.
     * 
     */
    public Output<Optional<String>> password() {
        return Codegen.optional(this.password);
    }
    /**
     * [Flannel Specific] The CIDR block for the pod network when using Flannel.
     * 
     */
    @Export(name="podCidr", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> podCidr;

    /**
     * @return [Flannel Specific] The CIDR block for the pod network when using Flannel.
     * 
     */
    public Output<Optional<String>> podCidr() {
        return Codegen.optional(this.podCidr);
    }
    /**
     * Proxy mode is option of kube-proxy. options: iptables|ipvs. default: ipvs.
     * 
     */
    @Export(name="proxyMode", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> proxyMode;

    /**
     * @return Proxy mode is option of kube-proxy. options: iptables|ipvs. default: ipvs.
     * 
     */
    public Output<Optional<String>> proxyMode() {
        return Codegen.optional(this.proxyMode);
    }
    /**
     * RDS instance list, You can choose which RDS instances whitelist to add instances to.
     * 
     */
    @Export(name="rdsInstances", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> rdsInstances;

    /**
     * @return RDS instance list, You can choose which RDS instances whitelist to add instances to.
     * 
     */
    public Output<Optional<List<String>>> rdsInstances() {
        return Codegen.optional(this.rdsInstances);
    }
    /**
     * The ID of the resource group,by default these cloud resources are automatically assigned to the default resource group.
     * 
     */
    @Export(name="resourceGroupId", refs={String.class}, tree="[0]")
    private Output<String> resourceGroupId;

    /**
     * @return The ID of the resource group,by default these cloud resources are automatically assigned to the default resource group.
     * 
     */
    public Output<String> resourceGroupId() {
        return this.resourceGroupId;
    }
    @Export(name="retainResources", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> retainResources;

    public Output<Optional<List<String>>> retainResources() {
        return Codegen.optional(this.retainResources);
    }
    /**
     * The runtime of containers. If you select another container runtime, see [Comparison of Docker, containerd, and Sandboxed-Container](https://www.alibabacloud.com/help/doc-detail/160313.htm). Detailed below.
     * 
     */
    @Export(name="runtime", refs={EdgeKubernetesRuntime.class}, tree="[0]")
    private Output</* @Nullable */ EdgeKubernetesRuntime> runtime;

    /**
     * @return The runtime of containers. If you select another container runtime, see [Comparison of Docker, containerd, and Sandboxed-Container](https://www.alibabacloud.com/help/doc-detail/160313.htm). Detailed below.
     * 
     */
    public Output<Optional<EdgeKubernetesRuntime>> runtime() {
        return Codegen.optional(this.runtime);
    }
    /**
     * The ID of the security group to which the ECS instances in the cluster belong. If it is not specified, a new Security group will be built.
     * 
     */
    @Export(name="securityGroupId", refs={String.class}, tree="[0]")
    private Output<String> securityGroupId;

    /**
     * @return The ID of the security group to which the ECS instances in the cluster belong. If it is not specified, a new Security group will be built.
     * 
     */
    public Output<String> securityGroupId() {
        return this.securityGroupId;
    }
    /**
     * The CIDR block for the service network. It cannot be duplicated with the VPC CIDR and CIDR used by Kubernetes cluster in VPC, cannot be modified after creation.
     * 
     */
    @Export(name="serviceCidr", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> serviceCidr;

    /**
     * @return The CIDR block for the service network. It cannot be duplicated with the VPC CIDR and CIDR used by Kubernetes cluster in VPC, cannot be modified after creation.
     * 
     */
    public Output<Optional<String>> serviceCidr() {
        return Codegen.optional(this.serviceCidr);
    }
    /**
     * The public ip of load balancer.
     * 
     */
    @Export(name="slbInternet", refs={String.class}, tree="[0]")
    private Output<String> slbInternet;

    /**
     * @return The public ip of load balancer.
     * 
     */
    public Output<String> slbInternet() {
        return this.slbInternet;
    }
    /**
     * Whether to create internet load balancer for API Server. Default to true.
     * 
     */
    @Export(name="slbInternetEnabled", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> slbInternetEnabled;

    /**
     * @return Whether to create internet load balancer for API Server. Default to true.
     * 
     */
    public Output<Optional<Boolean>> slbInternetEnabled() {
        return Codegen.optional(this.slbInternetEnabled);
    }
    /**
     * The ID of private load balancer where the current cluster master node is located.
     * 
     */
    @Export(name="slbIntranet", refs={String.class}, tree="[0]")
    private Output<String> slbIntranet;

    /**
     * @return The ID of private load balancer where the current cluster master node is located.
     * 
     */
    public Output<String> slbIntranet() {
        return this.slbIntranet;
    }
    /**
     * Default nil, A map of tags assigned to the kubernetes cluster and work node.
     * 
     */
    @Export(name="tags", refs={Map.class,String.class,Object.class}, tree="[0,1,2]")
    private Output</* @Nullable */ Map<String,Object>> tags;

    /**
     * @return Default nil, A map of tags assigned to the kubernetes cluster and work node.
     * 
     */
    public Output<Optional<Map<String,Object>>> tags() {
        return Codegen.optional(this.tags);
    }
    /**
     * Windows instances support batch and PowerShell scripts. If your script file is larger than 1 KB, we recommend that you upload the script to Object Storage Service (OSS) and pull it through the internal endpoint of your OSS bucket.
     * 
     */
    @Export(name="userData", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> userData;

    /**
     * @return Windows instances support batch and PowerShell scripts. If your script file is larger than 1 KB, we recommend that you upload the script to Object Storage Service (OSS) and pull it through the internal endpoint of your OSS bucket.
     * 
     */
    public Output<Optional<String>> userData() {
        return Codegen.optional(this.userData);
    }
    /**
     * Desired Kubernetes version. If you do not specify a value, the latest available version at resource creation is used and no upgrades will occur except you set a higher version number. The value must be configured and increased to upgrade the version when desired. Downgrades are not supported by ACK.
     * 
     */
    @Export(name="version", refs={String.class}, tree="[0]")
    private Output<String> version;

    /**
     * @return Desired Kubernetes version. If you do not specify a value, the latest available version at resource creation is used and no upgrades will occur except you set a higher version number. The value must be configured and increased to upgrade the version when desired. Downgrades are not supported by ACK.
     * 
     */
    public Output<String> version() {
        return this.version;
    }
    /**
     * The ID of VPC where the current cluster is located.
     * 
     */
    @Export(name="vpcId", refs={String.class}, tree="[0]")
    private Output<String> vpcId;

    /**
     * @return The ID of VPC where the current cluster is located.
     * 
     */
    public Output<String> vpcId() {
        return this.vpcId;
    }
    /**
     * The data disk configurations of worker nodes, such as the disk type and disk size. See `worker_data_disks` below.
     * 
     */
    @Export(name="workerDataDisks", refs={List.class,EdgeKubernetesWorkerDataDisk.class}, tree="[0,1]")
    private Output</* @Nullable */ List<EdgeKubernetesWorkerDataDisk>> workerDataDisks;

    /**
     * @return The data disk configurations of worker nodes, such as the disk type and disk size. See `worker_data_disks` below.
     * 
     */
    public Output<Optional<List<EdgeKubernetesWorkerDataDisk>>> workerDataDisks() {
        return Codegen.optional(this.workerDataDisks);
    }
    /**
     * The system disk category of worker node. Its valid value are `cloud_efficiency`, `cloud_ssd` and `cloud_essd` and . Default to `cloud_efficiency`.
     * 
     */
    @Export(name="workerDiskCategory", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> workerDiskCategory;

    /**
     * @return The system disk category of worker node. Its valid value are `cloud_efficiency`, `cloud_ssd` and `cloud_essd` and . Default to `cloud_efficiency`.
     * 
     */
    public Output<Optional<String>> workerDiskCategory() {
        return Codegen.optional(this.workerDiskCategory);
    }
    /**
     * Worker node system disk performance level, when `worker_disk_category` values `cloud_essd`, the optional values are `PL0`, `PL1`, `PL2` or `PL3`, but the specific performance level is related to the disk capacity. For more information, see [Enhanced SSDs](https://www.alibabacloud.com/help/doc-detail/122389.htm). Default is `PL1`.
     * 
     */
    @Export(name="workerDiskPerformanceLevel", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> workerDiskPerformanceLevel;

    /**
     * @return Worker node system disk performance level, when `worker_disk_category` values `cloud_essd`, the optional values are `PL0`, `PL1`, `PL2` or `PL3`, but the specific performance level is related to the disk capacity. For more information, see [Enhanced SSDs](https://www.alibabacloud.com/help/doc-detail/122389.htm). Default is `PL1`.
     * 
     */
    public Output<Optional<String>> workerDiskPerformanceLevel() {
        return Codegen.optional(this.workerDiskPerformanceLevel);
    }
    /**
     * The system disk size of worker node. Its valid value range [20~32768] in GB. Default to 40.
     * 
     */
    @Export(name="workerDiskSize", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> workerDiskSize;

    /**
     * @return The system disk size of worker node. Its valid value range [20~32768] in GB. Default to 40.
     * 
     */
    public Output<Optional<Integer>> workerDiskSize() {
        return Codegen.optional(this.workerDiskSize);
    }
    /**
     * Worker node system disk auto snapshot policy.
     * 
     * *Computed params*
     * 
     * You can set some file paths to save kube_config information, but this way is cumbersome. Since version 1.105.0, we&#39;ve written it to tf state file. About its use，see export attribute certificate_authority. From version 1.187.0+, new DataSource `alicloud.cs.getClusterCredential` is recommended to manage cluster&#39;s kube_config.
     * 
     */
    @Export(name="workerDiskSnapshotPolicyId", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> workerDiskSnapshotPolicyId;

    /**
     * @return Worker node system disk auto snapshot policy.
     * 
     * *Computed params*
     * 
     * You can set some file paths to save kube_config information, but this way is cumbersome. Since version 1.105.0, we&#39;ve written it to tf state file. About its use，see export attribute certificate_authority. From version 1.187.0+, new DataSource `alicloud.cs.getClusterCredential` is recommended to manage cluster&#39;s kube_config.
     * 
     */
    public Output<Optional<String>> workerDiskSnapshotPolicyId() {
        return Codegen.optional(this.workerDiskSnapshotPolicyId);
    }
    /**
     * Worker payment type, its valid value is `PostPaid`. Defaults to `PostPaid`. More charge details in [ACK@edge charge](https://help.aliyun.com/document_detail/178718.html).
     * 
     */
    @Export(name="workerInstanceChargeType", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> workerInstanceChargeType;

    /**
     * @return Worker payment type, its valid value is `PostPaid`. Defaults to `PostPaid`. More charge details in [ACK@edge charge](https://help.aliyun.com/document_detail/178718.html).
     * 
     */
    public Output<Optional<String>> workerInstanceChargeType() {
        return Codegen.optional(this.workerInstanceChargeType);
    }
    /**
     * The instance types of worker node, you can set multiple types to avoid NoStock of a certain type.
     * 
     */
    @Export(name="workerInstanceTypes", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> workerInstanceTypes;

    /**
     * @return The instance types of worker node, you can set multiple types to avoid NoStock of a certain type.
     * 
     */
    public Output<List<String>> workerInstanceTypes() {
        return this.workerInstanceTypes;
    }
    /**
     * List of cluster worker nodes.
     * 
     */
    @Export(name="workerNodes", refs={List.class,EdgeKubernetesWorkerNode.class}, tree="[0,1]")
    private Output<List<EdgeKubernetesWorkerNode>> workerNodes;

    /**
     * @return List of cluster worker nodes.
     * 
     */
    public Output<List<EdgeKubernetesWorkerNode>> workerNodes() {
        return this.workerNodes;
    }
    /**
     * The cloud worker node number of the edge kubernetes cluster. Default to 1. It is limited up to 50 and if you want to enlarge it, please apply white list or contact with us.
     * 
     */
    @Export(name="workerNumber", refs={Integer.class}, tree="[0]")
    private Output<Integer> workerNumber;

    /**
     * @return The cloud worker node number of the edge kubernetes cluster. Default to 1. It is limited up to 50 and if you want to enlarge it, please apply white list or contact with us.
     * 
     */
    public Output<Integer> workerNumber() {
        return this.workerNumber;
    }
    /**
     * The RamRole Name attached to worker node.
     * 
     */
    @Export(name="workerRamRoleName", refs={String.class}, tree="[0]")
    private Output<String> workerRamRoleName;

    /**
     * @return The RamRole Name attached to worker node.
     * 
     */
    public Output<String> workerRamRoleName() {
        return this.workerRamRoleName;
    }
    /**
     * The vswitches used by workers.
     * 
     */
    @Export(name="workerVswitchIds", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> workerVswitchIds;

    /**
     * @return The vswitches used by workers.
     * 
     */
    public Output<List<String>> workerVswitchIds() {
        return this.workerVswitchIds;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public EdgeKubernetes(String name) {
        this(name, EdgeKubernetesArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public EdgeKubernetes(String name, EdgeKubernetesArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public EdgeKubernetes(String name, EdgeKubernetesArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:cs/edgeKubernetes:EdgeKubernetes", name, args == null ? EdgeKubernetesArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private EdgeKubernetes(String name, Output<String> id, @Nullable EdgeKubernetesState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:cs/edgeKubernetes:EdgeKubernetes", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .additionalSecretOutputs(List.of(
                "password"
            ))
            .build();
        return com.pulumi.resources.CustomResourceOptions.merge(defaultOptions, options, id);
    }

    /**
     * Get an existing Host resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state
     * @param options Optional settings to control the behavior of the CustomResource.
     */
    public static EdgeKubernetes get(String name, Output<String> id, @Nullable EdgeKubernetesState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new EdgeKubernetes(name, id, state, options);
    }
}
