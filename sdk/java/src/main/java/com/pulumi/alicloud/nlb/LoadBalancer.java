// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.nlb;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.nlb.LoadBalancerArgs;
import com.pulumi.alicloud.nlb.inputs.LoadBalancerState;
import com.pulumi.alicloud.nlb.outputs.LoadBalancerDeletionProtectionConfig;
import com.pulumi.alicloud.nlb.outputs.LoadBalancerModificationProtectionConfig;
import com.pulumi.alicloud.nlb.outputs.LoadBalancerZoneMapping;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a NLB Load Balancer resource.
 * 
 * For information about NLB Load Balancer and how to use it, see [What is Load Balancer](https://www.alibabacloud.com/help/en/server-load-balancer/latest/createloadbalancer).
 * 
 * &gt; **NOTE:** Available since v1.191.0.
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
 * import com.pulumi.alicloud.resourcemanager.ResourcemanagerFunctions;
 * import com.pulumi.alicloud.resourcemanager.inputs.GetResourceGroupsArgs;
 * import com.pulumi.alicloud.nlb.NlbFunctions;
 * import com.pulumi.alicloud.nlb.inputs.GetZonesArgs;
 * import com.pulumi.alicloud.vpc.Network;
 * import com.pulumi.alicloud.vpc.NetworkArgs;
 * import com.pulumi.alicloud.vpc.Switch;
 * import com.pulumi.alicloud.vpc.SwitchArgs;
 * import com.pulumi.alicloud.nlb.LoadBalancer;
 * import com.pulumi.alicloud.nlb.LoadBalancerArgs;
 * import com.pulumi.alicloud.nlb.inputs.LoadBalancerZoneMappingArgs;
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
 *         final var name = config.get("name").orElse("tf-example");
 *         final var default = ResourcemanagerFunctions.getResourceGroups();
 * 
 *         final var defaultGetZones = NlbFunctions.getZones();
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
 *             .zoneId(defaultGetZones.applyValue(getZonesResult -> getZonesResult.zones()[0].id()))
 *             .build());
 * 
 *         var default1 = new Switch("default1", SwitchArgs.builder()
 *             .vswitchName(name)
 *             .cidrBlock("10.4.1.0/24")
 *             .vpcId(defaultNetwork.id())
 *             .zoneId(defaultGetZones.applyValue(getZonesResult -> getZonesResult.zones()[1].id()))
 *             .build());
 * 
 *         var defaultLoadBalancer = new LoadBalancer("defaultLoadBalancer", LoadBalancerArgs.builder()
 *             .loadBalancerName(name)
 *             .resourceGroupId(default_.ids()[0])
 *             .loadBalancerType("Network")
 *             .addressType("Internet")
 *             .addressIpVersion("Ipv4")
 *             .vpcId(defaultNetwork.id())
 *             .tags(Map.ofEntries(
 *                 Map.entry("Created", "TF"),
 *                 Map.entry("For", "example")
 *             ))
 *             .zoneMappings(            
 *                 LoadBalancerZoneMappingArgs.builder()
 *                     .vswitchId(defaultSwitch.id())
 *                     .zoneId(defaultGetZones.applyValue(getZonesResult -> getZonesResult.zones()[0].id()))
 *                     .build(),
 *                 LoadBalancerZoneMappingArgs.builder()
 *                     .vswitchId(default1.id())
 *                     .zoneId(defaultGetZones.applyValue(getZonesResult -> getZonesResult.zones()[1].id()))
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
 * NLB Load Balancer can be imported using the id, e.g.
 * 
 * ```sh
 * $ pulumi import alicloud:nlb/loadBalancer:LoadBalancer example &lt;id&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:nlb/loadBalancer:LoadBalancer")
public class LoadBalancer extends com.pulumi.resources.CustomResource {
    /**
     * Protocol version. Value:
     * - **Ipv4**:IPv4 type.
     * - **DualStack**: Double Stack type.
     * 
     */
    @Export(name="addressIpVersion", refs={String.class}, tree="[0]")
    private Output<String> addressIpVersion;

    /**
     * @return Protocol version. Value:
     * - **Ipv4**:IPv4 type.
     * - **DualStack**: Double Stack type.
     * 
     */
    public Output<String> addressIpVersion() {
        return this.addressIpVersion;
    }
    /**
     * The network address type of IPv4 for network load balancing. Value:
     * - **Internet**: public network. Load balancer has a public network IP address, and the DNS domain name is resolved to a public network IP address, so it can be accessed in a public network environment.
     * - **Intranet**: private network. The server load balancer only has a private IP address, and the DNS domain name is resolved to the private IP address, so it can only be accessed by the intranet environment of the VPC where the server load balancer is located.
     * 
     */
    @Export(name="addressType", refs={String.class}, tree="[0]")
    private Output<String> addressType;

    /**
     * @return The network address type of IPv4 for network load balancing. Value:
     * - **Internet**: public network. Load balancer has a public network IP address, and the DNS domain name is resolved to a public network IP address, so it can be accessed in a public network environment.
     * - **Intranet**: private network. The server load balancer only has a private IP address, and the DNS domain name is resolved to the private IP address, so it can only be accessed by the intranet environment of the VPC where the server load balancer is located.
     * 
     */
    public Output<String> addressType() {
        return this.addressType;
    }
    /**
     * The ID of the shared bandwidth package associated with the public network instance.
     * 
     */
    @Export(name="bandwidthPackageId", refs={String.class}, tree="[0]")
    private Output<String> bandwidthPackageId;

    /**
     * @return The ID of the shared bandwidth package associated with the public network instance.
     * 
     */
    public Output<String> bandwidthPackageId() {
        return this.bandwidthPackageId;
    }
    /**
     * Resource creation time, using Greenwich Mean Time, formating&#39; yyyy-MM-ddTHH:mm:ssZ &#39;.
     * 
     */
    @Export(name="createTime", refs={String.class}, tree="[0]")
    private Output<String> createTime;

    /**
     * @return Resource creation time, using Greenwich Mean Time, formating&#39; yyyy-MM-ddTHH:mm:ssZ &#39;.
     * 
     */
    public Output<String> createTime() {
        return this.createTime;
    }
    /**
     * Whether cross-zone is enabled for a network-based load balancing instance. Value:
     * - **true**: on.
     * - **false**: closed.
     * 
     */
    @Export(name="crossZoneEnabled", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> crossZoneEnabled;

    /**
     * @return Whether cross-zone is enabled for a network-based load balancing instance. Value:
     * - **true**: on.
     * - **false**: closed.
     * 
     */
    public Output<Boolean> crossZoneEnabled() {
        return this.crossZoneEnabled;
    }
    /**
     * Delete protection. See `deletion_protection_config` below.
     * 
     */
    @Export(name="deletionProtectionConfig", refs={LoadBalancerDeletionProtectionConfig.class}, tree="[0]")
    private Output<LoadBalancerDeletionProtectionConfig> deletionProtectionConfig;

    /**
     * @return Delete protection. See `deletion_protection_config` below.
     * 
     */
    public Output<LoadBalancerDeletionProtectionConfig> deletionProtectionConfig() {
        return this.deletionProtectionConfig;
    }
    /**
     * Specifies whether to enable deletion protection. Default value: `false`. Valid values:
     * 
     */
    @Export(name="deletionProtectionEnabled", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> deletionProtectionEnabled;

    /**
     * @return Specifies whether to enable deletion protection. Default value: `false`. Valid values:
     * 
     */
    public Output<Boolean> deletionProtectionEnabled() {
        return this.deletionProtectionEnabled;
    }
    /**
     * The reason why the deletion protection feature is enabled or disabled. The `deletion_protection_reason` takes effect only when `deletion_protection_enabled` is set to `true`.
     * 
     */
    @Export(name="deletionProtectionReason", refs={String.class}, tree="[0]")
    private Output<String> deletionProtectionReason;

    /**
     * @return The reason why the deletion protection feature is enabled or disabled. The `deletion_protection_reason` takes effect only when `deletion_protection_enabled` is set to `true`.
     * 
     */
    public Output<String> deletionProtectionReason() {
        return this.deletionProtectionReason;
    }
    /**
     * The domain name of the NLB instance.
     * 
     */
    @Export(name="dnsName", refs={String.class}, tree="[0]")
    private Output<String> dnsName;

    /**
     * @return The domain name of the NLB instance.
     * 
     */
    public Output<String> dnsName() {
        return this.dnsName;
    }
    /**
     * The IPv6 address type of network load balancing. Value:
     * - **Internet**: Server Load Balancer has a public IP address, and the DNS domain name is resolved to a public IP address, so it can be accessed in a public network environment.
     * - **Intranet**: SLB only has the private IP address, and the DNS domain name is resolved to the private IP address, so it can only be accessed by the Intranet environment of the VPC where SLB is located.
     * 
     */
    @Export(name="ipv6AddressType", refs={String.class}, tree="[0]")
    private Output<String> ipv6AddressType;

    /**
     * @return The IPv6 address type of network load balancing. Value:
     * - **Internet**: Server Load Balancer has a public IP address, and the DNS domain name is resolved to a public IP address, so it can be accessed in a public network environment.
     * - **Intranet**: SLB only has the private IP address, and the DNS domain name is resolved to the private IP address, so it can only be accessed by the Intranet environment of the VPC where SLB is located.
     * 
     */
    public Output<String> ipv6AddressType() {
        return this.ipv6AddressType;
    }
    /**
     * The business status of the NLB instance.
     * 
     */
    @Export(name="loadBalancerBusinessStatus", refs={String.class}, tree="[0]")
    private Output<String> loadBalancerBusinessStatus;

    /**
     * @return The business status of the NLB instance.
     * 
     */
    public Output<String> loadBalancerBusinessStatus() {
        return this.loadBalancerBusinessStatus;
    }
    /**
     * The name of the network-based load balancing instance.  2 to 128 English or Chinese characters in length, which must start with a letter or Chinese, and can contain numbers, half-width periods (.), underscores (_), and dashes (-).
     * 
     */
    @Export(name="loadBalancerName", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> loadBalancerName;

    /**
     * @return The name of the network-based load balancing instance.  2 to 128 English or Chinese characters in length, which must start with a letter or Chinese, and can contain numbers, half-width periods (.), underscores (_), and dashes (-).
     * 
     */
    public Output<Optional<String>> loadBalancerName() {
        return Codegen.optional(this.loadBalancerName);
    }
    /**
     * Load balancing type. Only value: **network**, which indicates network-based load balancing.
     * 
     */
    @Export(name="loadBalancerType", refs={String.class}, tree="[0]")
    private Output<String> loadBalancerType;

    /**
     * @return Load balancing type. Only value: **network**, which indicates network-based load balancing.
     * 
     */
    public Output<String> loadBalancerType() {
        return this.loadBalancerType;
    }
    /**
     * Modify protection. See `modification_protection_config` below.
     * 
     */
    @Export(name="modificationProtectionConfig", refs={LoadBalancerModificationProtectionConfig.class}, tree="[0]")
    private Output<LoadBalancerModificationProtectionConfig> modificationProtectionConfig;

    /**
     * @return Modify protection. See `modification_protection_config` below.
     * 
     */
    public Output<LoadBalancerModificationProtectionConfig> modificationProtectionConfig() {
        return this.modificationProtectionConfig;
    }
    /**
     * The reason why the configuration read-only mode is enabled. The `modification_protection_reason` takes effect only when `modification_protection_status` is set to `ConsoleProtection`.
     * 
     */
    @Export(name="modificationProtectionReason", refs={String.class}, tree="[0]")
    private Output<String> modificationProtectionReason;

    /**
     * @return The reason why the configuration read-only mode is enabled. The `modification_protection_reason` takes effect only when `modification_protection_status` is set to `ConsoleProtection`.
     * 
     */
    public Output<String> modificationProtectionReason() {
        return this.modificationProtectionReason;
    }
    /**
     * Specifies whether to enable the configuration read-only mode. Default value: `NonProtection`. Valid values:
     * - `NonProtection`: Does not enable the configuration read-only mode. You cannot set the `modification_protection_reason`. If the `modification_protection_reason` is set, the value is cleared.
     * - `ConsoleProtection`: Enables the configuration read-only mode. You can set the `modification_protection_reason`.
     * 
     */
    @Export(name="modificationProtectionStatus", refs={String.class}, tree="[0]")
    private Output<String> modificationProtectionStatus;

    /**
     * @return Specifies whether to enable the configuration read-only mode. Default value: `NonProtection`. Valid values:
     * - `NonProtection`: Does not enable the configuration read-only mode. You cannot set the `modification_protection_reason`. If the `modification_protection_reason` is set, the value is cleared.
     * - `ConsoleProtection`: Enables the configuration read-only mode. You can set the `modification_protection_reason`.
     * 
     */
    public Output<String> modificationProtectionStatus() {
        return this.modificationProtectionStatus;
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
     * The security group to which the network-based SLB instance belongs.
     * 
     */
    @Export(name="securityGroupIds", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> securityGroupIds;

    /**
     * @return The security group to which the network-based SLB instance belongs.
     * 
     */
    public Output<List<String>> securityGroupIds() {
        return this.securityGroupIds;
    }
    /**
     * The status of the resource.
     * 
     */
    @Export(name="status", refs={String.class}, tree="[0]")
    private Output<String> status;

    /**
     * @return The status of the resource.
     * 
     */
    public Output<String> status() {
        return this.status;
    }
    /**
     * List of labels.
     * 
     */
    @Export(name="tags", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output</* @Nullable */ Map<String,String>> tags;

    /**
     * @return List of labels.
     * 
     */
    public Output<Optional<Map<String,String>>> tags() {
        return Codegen.optional(this.tags);
    }
    /**
     * The ID of the network-based SLB instance.
     * 
     */
    @Export(name="vpcId", refs={String.class}, tree="[0]")
    private Output<String> vpcId;

    /**
     * @return The ID of the network-based SLB instance.
     * 
     */
    public Output<String> vpcId() {
        return this.vpcId;
    }
    /**
     * The list of zones and vSwitch mappings. You must add at least two zones and a maximum of 10 zones. See `zone_mappings` below.
     * 
     */
    @Export(name="zoneMappings", refs={List.class,LoadBalancerZoneMapping.class}, tree="[0,1]")
    private Output<List<LoadBalancerZoneMapping>> zoneMappings;

    /**
     * @return The list of zones and vSwitch mappings. You must add at least two zones and a maximum of 10 zones. See `zone_mappings` below.
     * 
     */
    public Output<List<LoadBalancerZoneMapping>> zoneMappings() {
        return this.zoneMappings;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public LoadBalancer(java.lang.String name) {
        this(name, LoadBalancerArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public LoadBalancer(java.lang.String name, LoadBalancerArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public LoadBalancer(java.lang.String name, LoadBalancerArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:nlb/loadBalancer:LoadBalancer", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private LoadBalancer(java.lang.String name, Output<java.lang.String> id, @Nullable LoadBalancerState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:nlb/loadBalancer:LoadBalancer", name, state, makeResourceOptions(options, id), false);
    }

    private static LoadBalancerArgs makeArgs(LoadBalancerArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? LoadBalancerArgs.Empty : args;
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
    public static LoadBalancer get(java.lang.String name, Output<java.lang.String> id, @Nullable LoadBalancerState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new LoadBalancer(name, id, state, options);
    }
}
