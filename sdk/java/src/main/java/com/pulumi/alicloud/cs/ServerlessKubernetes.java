// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.cs;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.cs.ServerlessKubernetesArgs;
import com.pulumi.alicloud.cs.inputs.ServerlessKubernetesState;
import com.pulumi.alicloud.cs.outputs.ServerlessKubernetesAddon;
import com.pulumi.alicloud.cs.outputs.ServerlessKubernetesRrsaMetadata;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Boolean;
import java.lang.Object;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * This resource will help you to manager a Serverless Kubernetes Cluster, see [What is serverless kubernetes](https://www.alibabacloud.com/help/en/ack/ack-managed-and-ack-dedicated/developer-reference/create-a-dedicated-kubernetes-cluster-that-supports-sandboxed-containers). The cluster is same as container service created by web console.
 * 
 * &gt; **NOTE:** Available since v1.58.0.
 * 
 * &gt; **NOTE:** Serverless Kubernetes cluster only supports VPC network and it can access internet while creating kubernetes cluster.
 * A Nat Gateway and configuring a SNAT for it can ensure one VPC network access internet. If there is no nat gateway in the
 * VPC, you can set `new_nat_gateway` to &#34;true&#34; to create one automatically.
 * 
 * &gt; **NOTE:** Creating serverless kubernetes cluster need to install several packages and it will cost about 5 minutes. Please be patient.
 * 
 * &gt; **NOTE:** The provider supports to download kube config, client certificate, client key and cluster ca certificate
 * after creating cluster successfully, and you can put them into the specified location, like &#39;~/.kube/config&#39;.
 * 
 * &gt; **NOTE:** If you want to manage serverless Kubernetes, you can use Kubernetes Provider.
 * 
 * &gt; **NOTE:** You need to activate several other products and confirm Authorization Policy used by Container Service before using this resource.
 * Please refer to the `Authorization management` and `Cluster management` sections in the [Document Center](https://www.alibabacloud.com/help/doc-detail/86488.htm).
 * 
 * &gt; **NOTE:** From version 1.162.0, support for creating professional serverless cluster.
 * 
 * ## Example Usage
 * 
 * Basic Usage
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.alicloud.AlicloudFunctions;
 * import com.pulumi.alicloud.inputs.GetZonesArgs;
 * import com.pulumi.alicloud.vpc.Network;
 * import com.pulumi.alicloud.vpc.NetworkArgs;
 * import com.pulumi.alicloud.vpc.Switch;
 * import com.pulumi.alicloud.vpc.SwitchArgs;
 * import com.pulumi.alicloud.cs.ServerlessKubernetes;
 * import com.pulumi.alicloud.cs.ServerlessKubernetesArgs;
 * import com.pulumi.alicloud.cs.inputs.ServerlessKubernetesAddonArgs;
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
 *         final var name = config.get("name").orElse("ask-example");
 *         final var default = AlicloudFunctions.getZones(GetZonesArgs.builder()
 *             .availableResourceCreation("VSwitch")
 *             .build());
 * 
 *         var defaultNetwork = new Network("defaultNetwork", NetworkArgs.builder()
 *             .vpcName(name)
 *             .cidrBlock("10.1.0.0/21")
 *             .build());
 * 
 *         var defaultSwitch = new Switch("defaultSwitch", SwitchArgs.builder()
 *             .vswitchName(name)
 *             .vpcId(defaultNetwork.id())
 *             .cidrBlock("10.1.1.0/24")
 *             .zoneId(default_.zones()[0].id())
 *             .build());
 * 
 *         var serverless = new ServerlessKubernetes("serverless", ServerlessKubernetesArgs.builder()
 *             .namePrefix(name)
 *             .clusterSpec("ack.pro.small")
 *             .vpcId(defaultNetwork.id())
 *             .vswitchIds(defaultSwitch.id())
 *             .newNatGateway(true)
 *             .endpointPublicAccessEnabled(true)
 *             .deletionProtection(false)
 *             .loadBalancerSpec("slb.s2.small")
 *             .timeZone("Asia/Shanghai")
 *             .serviceCidr("172.21.0.0/20")
 *             .serviceDiscoveryTypes("PrivateZone")
 *             .loggingType("SLS")
 *             .tags(Map.ofEntries(
 *                 Map.entry("k-aa", "v-aa"),
 *                 Map.entry("k-bb", "v-aa")
 *             ))
 *             .addons(            
 *                 ServerlessKubernetesAddonArgs.builder()
 *                     .name("alb-ingress-controller")
 *                     .build(),
 *                 ServerlessKubernetesAddonArgs.builder()
 *                     .name("metrics-server")
 *                     .build(),
 *                 ServerlessKubernetesAddonArgs.builder()
 *                     .name("knative")
 *                     .build())
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * Serverless Kubernetes cluster can be imported using the id, e.g. Then complete the main.tf accords to the result of `pulumi preview`.
 * 
 * ```sh
 * $ pulumi import alicloud:cs/serverlessKubernetes:ServerlessKubernetes main ce4273f9156874b46bb
 * ```
 * 
 */
@ResourceType(type="alicloud:cs/serverlessKubernetes:ServerlessKubernetes")
public class ServerlessKubernetes extends com.pulumi.resources.CustomResource {
    /**
     * You can specific network plugin,log component,ingress component and so on. See `addons` below.
     * 
     */
    @Export(name="addons", refs={List.class,ServerlessKubernetesAddon.class}, tree="[0,1]")
    private Output<List<ServerlessKubernetesAddon>> addons;

    /**
     * @return You can specific network plugin,log component,ingress component and so on. See `addons` below.
     * 
     */
    public Output<List<ServerlessKubernetesAddon>> addons() {
        return this.addons;
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
     */
    @Export(name="clusterCaCert", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> clusterCaCert;

    /**
     * @return The path of cluster ca certificate, like `~/.kube/cluster-ca-cert.pem`
     * 
     */
    public Output<Optional<String>> clusterCaCert() {
        return Codegen.optional(this.clusterCaCert);
    }
    /**
     * The cluster specifications of serverless kubernetes cluster, which can be empty. Valid values:
     * - ack.standard: Standard serverless clusters.
     * - ack.pro.small: Professional serverless clusters.
     * 
     */
    @Export(name="clusterSpec", refs={String.class}, tree="[0]")
    private Output<String> clusterSpec;

    /**
     * @return The cluster specifications of serverless kubernetes cluster, which can be empty. Valid values:
     * - ack.standard: Standard serverless clusters.
     * - ack.pro.small: Professional serverless clusters.
     * 
     */
    public Output<String> clusterSpec() {
        return this.clusterSpec;
    }
    /**
     * whether to create a v2 version cluster.
     * 
     * *Removed params*
     * 
     */
    @Export(name="createV2Cluster", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> createV2Cluster;

    /**
     * @return whether to create a v2 version cluster.
     * 
     * *Removed params*
     * 
     */
    public Output<Boolean> createV2Cluster() {
        return this.createV2Cluster;
    }
    /**
     * Whether enable the deletion protection or not.
     * - true: Enable deletion protection.
     * - false: Disable deletion protection.
     * 
     */
    @Export(name="deletionProtection", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> deletionProtection;

    /**
     * @return Whether enable the deletion protection or not.
     * - true: Enable deletion protection.
     * - false: Disable deletion protection.
     * 
     */
    public Output<Optional<Boolean>> deletionProtection() {
        return Codegen.optional(this.deletionProtection);
    }
    /**
     * Whether to enable cluster to support RRSA for version 1.22.3+. Default to `false`. Once the RRSA function is turned on, it is not allowed to turn off. If your cluster has enabled this function, please manually modify your tf file and add the rrsa configuration to the file, learn more [RAM Roles for Service Accounts](https://www.alibabacloud.com/help/zh/container-service-for-kubernetes/latest/use-rrsa-to-enforce-access-control).
     * 
     */
    @Export(name="enableRrsa", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> enableRrsa;

    /**
     * @return Whether to enable cluster to support RRSA for version 1.22.3+. Default to `false`. Once the RRSA function is turned on, it is not allowed to turn off. If your cluster has enabled this function, please manually modify your tf file and add the rrsa configuration to the file, learn more [RAM Roles for Service Accounts](https://www.alibabacloud.com/help/zh/container-service-for-kubernetes/latest/use-rrsa-to-enforce-access-control).
     * 
     */
    public Output<Optional<Boolean>> enableRrsa() {
        return Codegen.optional(this.enableRrsa);
    }
    /**
     * Whether to create internet  eip for API Server. Default to false.
     * 
     */
    @Export(name="endpointPublicAccessEnabled", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> endpointPublicAccessEnabled;

    /**
     * @return Whether to create internet  eip for API Server. Default to false.
     * 
     */
    public Output<Optional<Boolean>> endpointPublicAccessEnabled() {
        return Codegen.optional(this.endpointPublicAccessEnabled);
    }
    /**
     * Default false, when you want to change `vpc_id` and `vswitch_id`, you have to set this field to true, then the cluster will be recreated.
     * 
     */
    @Export(name="forceUpdate", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> forceUpdate;

    /**
     * @return Default false, when you want to change `vpc_id` and `vswitch_id`, you have to set this field to true, then the cluster will be recreated.
     * 
     */
    public Output<Optional<Boolean>> forceUpdate() {
        return Codegen.optional(this.forceUpdate);
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
     * The cluster api server load balance instance specification, default `slb.s2.small`. For more information on how to select a LB instance specification, see [SLB instance overview](https://help.aliyun.com/document_detail/85931.html).
     * 
     */
    @Export(name="loadBalancerSpec", refs={String.class}, tree="[0]")
    private Output<String> loadBalancerSpec;

    /**
     * @return The cluster api server load balance instance specification, default `slb.s2.small`. For more information on how to select a LB instance specification, see [SLB instance overview](https://help.aliyun.com/document_detail/85931.html).
     * 
     */
    public Output<String> loadBalancerSpec() {
        return this.loadBalancerSpec;
    }
    /**
     * Enable log service, Valid value `SLS`.
     * 
     */
    @Export(name="loggingType", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> loggingType;

    /**
     * @return Enable log service, Valid value `SLS`.
     * 
     */
    public Output<Optional<String>> loggingType() {
        return Codegen.optional(this.loggingType);
    }
    /**
     * The kubernetes cluster&#39;s name. It is the only in one Alicloud account.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The kubernetes cluster&#39;s name. It is the only in one Alicloud account.
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
     * Whether to create a new nat gateway while creating kubernetes cluster. SNAT must be configured when a new VPC is automatically created. Default is `true`.
     * 
     */
    @Export(name="newNatGateway", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> newNatGateway;

    /**
     * @return Whether to create a new nat gateway while creating kubernetes cluster. SNAT must be configured when a new VPC is automatically created. Default is `true`.
     * 
     */
    public Output<Optional<Boolean>> newNatGateway() {
        return Codegen.optional(this.newNatGateway);
    }
    /**
     * Has been deprecated from provider version 1.123.1. `PrivateZone` is used as the enumeration value of `service_discovery_types`.
     * 
     * @deprecated
     * Field &#39;private_zone&#39; has been deprecated from provider version 1.123.1. New field &#39;service_discovery_types&#39; replace it.
     * 
     */
    @Deprecated /* Field 'private_zone' has been deprecated from provider version 1.123.1. New field 'service_discovery_types' replace it. */
    @Export(name="privateZone", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> privateZone;

    /**
     * @return Has been deprecated from provider version 1.123.1. `PrivateZone` is used as the enumeration value of `service_discovery_types`.
     * 
     */
    public Output<Optional<Boolean>> privateZone() {
        return Codegen.optional(this.privateZone);
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
     * Nested attribute containing RRSA related data for your cluster. See `rrsa_metadata` below.
     * 
     */
    @Export(name="rrsaMetadata", refs={ServerlessKubernetesRrsaMetadata.class}, tree="[0]")
    private Output<ServerlessKubernetesRrsaMetadata> rrsaMetadata;

    /**
     * @return Nested attribute containing RRSA related data for your cluster. See `rrsa_metadata` below.
     * 
     */
    public Output<ServerlessKubernetesRrsaMetadata> rrsaMetadata() {
        return this.rrsaMetadata;
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
     * CIDR block of the service network. The specified CIDR block cannot overlap with that of the VPC or those of the ACK clusters that are deployed in the VPC. The CIDR block cannot be modified after the cluster is created.
     * 
     */
    @Export(name="serviceCidr", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> serviceCidr;

    /**
     * @return CIDR block of the service network. The specified CIDR block cannot overlap with that of the VPC or those of the ACK clusters that are deployed in the VPC. The CIDR block cannot be modified after the cluster is created.
     * 
     */
    public Output<Optional<String>> serviceCidr() {
        return Codegen.optional(this.serviceCidr);
    }
    /**
     * Service discovery type. If the value is empty, it means that service discovery is not enabled. Valid values are `CoreDNS` and `PrivateZone`.
     * 
     */
    @Export(name="serviceDiscoveryTypes", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> serviceDiscoveryTypes;

    /**
     * @return Service discovery type. If the value is empty, it means that service discovery is not enabled. Valid values are `CoreDNS` and `PrivateZone`.
     * 
     */
    public Output<Optional<List<String>>> serviceDiscoveryTypes() {
        return Codegen.optional(this.serviceDiscoveryTypes);
    }
    /**
     * If you use an existing SLS project, you must specify `sls_project_name`.
     * 
     */
    @Export(name="slsProjectName", refs={String.class}, tree="[0]")
    private Output<String> slsProjectName;

    /**
     * @return If you use an existing SLS project, you must specify `sls_project_name`.
     * 
     */
    public Output<String> slsProjectName() {
        return this.slsProjectName;
    }
    /**
     * Default nil, A map of tags assigned to the kubernetes cluster and work nodes.
     * 
     */
    @Export(name="tags", refs={Map.class,String.class,Object.class}, tree="[0,1,2]")
    private Output</* @Nullable */ Map<String,Object>> tags;

    /**
     * @return Default nil, A map of tags assigned to the kubernetes cluster and work nodes.
     * 
     */
    public Output<Optional<Map<String,Object>>> tags() {
        return Codegen.optional(this.tags);
    }
    /**
     * The time zone of the cluster.
     * 
     */
    @Export(name="timeZone", refs={String.class}, tree="[0]")
    private Output<String> timeZone;

    /**
     * @return The time zone of the cluster.
     * 
     */
    public Output<String> timeZone() {
        return this.timeZone;
    }
    /**
     * Desired Kubernetes version. If you do not specify a value, the latest available version at resource creation is used.
     * 
     */
    @Export(name="version", refs={String.class}, tree="[0]")
    private Output<String> version;

    /**
     * @return Desired Kubernetes version. If you do not specify a value, the latest available version at resource creation is used.
     * 
     */
    public Output<String> version() {
        return this.version;
    }
    /**
     * The vpc where new kubernetes cluster will be located. Specify one vpc&#39;s id, if it is not specified, a new VPC  will be built.
     * 
     */
    @Export(name="vpcId", refs={String.class}, tree="[0]")
    private Output<String> vpcId;

    /**
     * @return The vpc where new kubernetes cluster will be located. Specify one vpc&#39;s id, if it is not specified, a new VPC  will be built.
     * 
     */
    public Output<String> vpcId() {
        return this.vpcId;
    }
    /**
     * The vswitch where new kubernetes cluster will be located. Specify one vswitch&#39;s id, if it is not specified, a new VPC and VSwicth will be built. It must be in the zone which `availability_zone` specified.
     * 
     * @deprecated
     * Field &#39;vswitch_id&#39; has been deprecated from provider version 1.91.0. New field &#39;vswitch_ids&#39; replace it.
     * 
     */
    @Deprecated /* Field 'vswitch_id' has been deprecated from provider version 1.91.0. New field 'vswitch_ids' replace it. */
    @Export(name="vswitchId", refs={String.class}, tree="[0]")
    private Output<String> vswitchId;

    /**
     * @return The vswitch where new kubernetes cluster will be located. Specify one vswitch&#39;s id, if it is not specified, a new VPC and VSwicth will be built. It must be in the zone which `availability_zone` specified.
     * 
     */
    public Output<String> vswitchId() {
        return this.vswitchId;
    }
    /**
     * The vswitches where new kubernetes cluster will be located.
     * 
     */
    @Export(name="vswitchIds", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> vswitchIds;

    /**
     * @return The vswitches where new kubernetes cluster will be located.
     * 
     */
    public Output<List<String>> vswitchIds() {
        return this.vswitchIds;
    }
    /**
     * When creating a cluster using automatic VPC creation, you need to specify the zone where the VPC is located.
     * 
     */
    @Export(name="zoneId", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> zoneId;

    /**
     * @return When creating a cluster using automatic VPC creation, you need to specify the zone where the VPC is located.
     * 
     */
    public Output<Optional<String>> zoneId() {
        return Codegen.optional(this.zoneId);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public ServerlessKubernetes(java.lang.String name) {
        this(name, ServerlessKubernetesArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ServerlessKubernetes(java.lang.String name, ServerlessKubernetesArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ServerlessKubernetes(java.lang.String name, ServerlessKubernetesArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:cs/serverlessKubernetes:ServerlessKubernetes", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private ServerlessKubernetes(java.lang.String name, Output<java.lang.String> id, @Nullable ServerlessKubernetesState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:cs/serverlessKubernetes:ServerlessKubernetes", name, state, makeResourceOptions(options, id), false);
    }

    private static ServerlessKubernetesArgs makeArgs(ServerlessKubernetesArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? ServerlessKubernetesArgs.Empty : args;
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<java.lang.String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
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
    public static ServerlessKubernetes get(java.lang.String name, Output<java.lang.String> id, @Nullable ServerlessKubernetesState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ServerlessKubernetes(name, id, state, options);
    }
}
