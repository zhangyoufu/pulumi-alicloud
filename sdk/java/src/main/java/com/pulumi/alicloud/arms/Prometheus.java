// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.arms;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.arms.PrometheusArgs;
import com.pulumi.alicloud.arms.inputs.PrometheusState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Object;
import java.lang.String;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a Application Real-Time Monitoring Service (ARMS) Prometheus resource.
 * 
 * For information about Application Real-Time Monitoring Service (ARMS) Prometheus and how to use it, see [What is Prometheus](https://www.alibabacloud.com/help/en/arms/developer-reference/api-arms-2019-08-08-createprometheusinstance).
 * 
 * &gt; **NOTE:** Available since v1.203.0.
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
 * import com.pulumi.alicloud.ecs.SecurityGroup;
 * import com.pulumi.alicloud.ecs.SecurityGroupArgs;
 * import com.pulumi.alicloud.resourcemanager.ResourcemanagerFunctions;
 * import com.pulumi.alicloud.resourcemanager.inputs.GetResourceGroupsArgs;
 * import com.pulumi.alicloud.arms.Prometheus;
 * import com.pulumi.alicloud.arms.PrometheusArgs;
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
 *         final var name = config.get("name").orElse("tf_example");
 *         final var default = AlicloudFunctions.getZones(GetZonesArgs.builder()
 *             .availableResourceCreation("VSwitch")
 *             .build());
 * 
 *         var defaultNetwork = new Network("defaultNetwork", NetworkArgs.builder()        
 *             .vpcName(name)
 *             .cidrBlock("10.4.0.0/16")
 *             .build());
 * 
 *         var defaultSwitch = new Switch("defaultSwitch", SwitchArgs.builder()        
 *             .vswitchName(name)
 *             .cidrBlock("10.4.0.0/24")
 *             .vpcId(defaultNetwork.id())
 *             .zoneId(default_.zones()[default_.zones().length() - 1].id())
 *             .build());
 * 
 *         var defaultSecurityGroup = new SecurityGroup("defaultSecurityGroup", SecurityGroupArgs.builder()        
 *             .name(name)
 *             .vpcId(defaultNetwork.id())
 *             .build());
 * 
 *         final var defaultGetResourceGroups = ResourcemanagerFunctions.getResourceGroups();
 * 
 *         var defaultPrometheus = new Prometheus("defaultPrometheus", PrometheusArgs.builder()        
 *             .clusterType("ecs")
 *             .grafanaInstanceId("free")
 *             .vpcId(defaultNetwork.id())
 *             .vswitchId(defaultSwitch.id())
 *             .securityGroupId(defaultSecurityGroup.id())
 *             .clusterName(defaultNetwork.id().applyValue(id -> String.format("%s-%s", name,id)))
 *             .resourceGroupId(defaultGetResourceGroups.applyValue(getResourceGroupsResult -> getResourceGroupsResult.groups()[0].id()))
 *             .tags(Map.ofEntries(
 *                 Map.entry("Created", "TF"),
 *                 Map.entry("For", "Prometheus")
 *             ))
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
 * Application Real-Time Monitoring Service (ARMS) Prometheus can be imported using the id, e.g.
 * 
 * ```sh
 * $ pulumi import alicloud:arms/prometheus:Prometheus example &lt;id&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:arms/prometheus:Prometheus")
public class Prometheus extends com.pulumi.resources.CustomResource {
    /**
     * The ID of the Kubernetes cluster. This parameter is required, if you set `cluster_type` to `aliyun-cs`.
     * 
     */
    @Export(name="clusterId", refs={String.class}, tree="[0]")
    private Output<String> clusterId;

    /**
     * @return The ID of the Kubernetes cluster. This parameter is required, if you set `cluster_type` to `aliyun-cs`.
     * 
     */
    public Output<String> clusterId() {
        return this.clusterId;
    }
    /**
     * The name of the created cluster. This parameter is required, if you set `cluster_type` to `remote-write`, `ecs` or `global-view`.
     * 
     */
    @Export(name="clusterName", refs={String.class}, tree="[0]")
    private Output<String> clusterName;

    /**
     * @return The name of the created cluster. This parameter is required, if you set `cluster_type` to `remote-write`, `ecs` or `global-view`.
     * 
     */
    public Output<String> clusterName() {
        return this.clusterName;
    }
    /**
     * The type of the Prometheus instance. Valid values: `remote-write`, `ecs`, `global-view`, `aliyun-cs`.
     * 
     */
    @Export(name="clusterType", refs={String.class}, tree="[0]")
    private Output<String> clusterType;

    /**
     * @return The type of the Prometheus instance. Valid values: `remote-write`, `ecs`, `global-view`, `aliyun-cs`.
     * 
     */
    public Output<String> clusterType() {
        return this.clusterType;
    }
    /**
     * The ID of the Grafana dedicated instance. When using the shared version of Grafana, you can set `grafana_instance_id` to `free`.
     * 
     */
    @Export(name="grafanaInstanceId", refs={String.class}, tree="[0]")
    private Output<String> grafanaInstanceId;

    /**
     * @return The ID of the Grafana dedicated instance. When using the shared version of Grafana, you can set `grafana_instance_id` to `free`.
     * 
     */
    public Output<String> grafanaInstanceId() {
        return this.grafanaInstanceId;
    }
    /**
     * The ID of the resource group.
     * 
     */
    @Export(name="resourceGroupId", refs={String.class}, tree="[0]")
    private Output<String> resourceGroupId;

    /**
     * @return The ID of the resource group.
     * 
     */
    public Output<String> resourceGroupId() {
        return this.resourceGroupId;
    }
    /**
     * The ID of the security group. This parameter is required, if you set `cluster_type` to `ecs` or `aliyun-cs`(ASK instance).
     * 
     */
    @Export(name="securityGroupId", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> securityGroupId;

    /**
     * @return The ID of the security group. This parameter is required, if you set `cluster_type` to `ecs` or `aliyun-cs`(ASK instance).
     * 
     */
    public Output<Optional<String>> securityGroupId() {
        return Codegen.optional(this.securityGroupId);
    }
    /**
     * The child instance json string of the globalView instance.
     * 
     */
    @Export(name="subClustersJson", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> subClustersJson;

    /**
     * @return The child instance json string of the globalView instance.
     * 
     */
    public Output<Optional<String>> subClustersJson() {
        return Codegen.optional(this.subClustersJson);
    }
    /**
     * A mapping of tags to assign to the resource.
     * 
     */
    @Export(name="tags", refs={Map.class,String.class,Object.class}, tree="[0,1,2]")
    private Output</* @Nullable */ Map<String,Object>> tags;

    /**
     * @return A mapping of tags to assign to the resource.
     * 
     */
    public Output<Optional<Map<String,Object>>> tags() {
        return Codegen.optional(this.tags);
    }
    /**
     * The ID of the VPC. This parameter is required, if you set `cluster_type` to `ecs` or `aliyun-cs`(ASK instance).
     * 
     */
    @Export(name="vpcId", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> vpcId;

    /**
     * @return The ID of the VPC. This parameter is required, if you set `cluster_type` to `ecs` or `aliyun-cs`(ASK instance).
     * 
     */
    public Output<Optional<String>> vpcId() {
        return Codegen.optional(this.vpcId);
    }
    /**
     * The ID of the VSwitch. This parameter is required, if you set `cluster_type` to `ecs` or `aliyun-cs`(ASK instance).
     * 
     */
    @Export(name="vswitchId", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> vswitchId;

    /**
     * @return The ID of the VSwitch. This parameter is required, if you set `cluster_type` to `ecs` or `aliyun-cs`(ASK instance).
     * 
     */
    public Output<Optional<String>> vswitchId() {
        return Codegen.optional(this.vswitchId);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Prometheus(String name) {
        this(name, PrometheusArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Prometheus(String name, PrometheusArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Prometheus(String name, PrometheusArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:arms/prometheus:Prometheus", name, args == null ? PrometheusArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Prometheus(String name, Output<String> id, @Nullable PrometheusState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:arms/prometheus:Prometheus", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
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
    public static Prometheus get(String name, Output<String> id, @Nullable PrometheusState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Prometheus(name, id, state, options);
    }
}
