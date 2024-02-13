// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.vpc;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.vpc.RouteEntryArgs;
import com.pulumi.alicloud.vpc.inputs.RouteEntryState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a route entry resource. A route entry represents a route item of one VPC route table.
 * 
 * ## Example Usage
 * 
 * Basic Usage
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
 * import com.pulumi.alicloud.ecs.inputs.GetImagesArgs;
 * import com.pulumi.alicloud.vpc.Network;
 * import com.pulumi.alicloud.vpc.NetworkArgs;
 * import com.pulumi.alicloud.vpc.Switch;
 * import com.pulumi.alicloud.vpc.SwitchArgs;
 * import com.pulumi.alicloud.ecs.SecurityGroup;
 * import com.pulumi.alicloud.ecs.SecurityGroupArgs;
 * import com.pulumi.alicloud.ecs.SecurityGroupRule;
 * import com.pulumi.alicloud.ecs.SecurityGroupRuleArgs;
 * import com.pulumi.alicloud.ecs.Instance;
 * import com.pulumi.alicloud.ecs.InstanceArgs;
 * import com.pulumi.alicloud.vpc.RouteEntry;
 * import com.pulumi.alicloud.vpc.RouteEntryArgs;
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
 *         final var defaultZones = AlicloudFunctions.getZones(GetZonesArgs.builder()
 *             .availableResourceCreation(&#34;VSwitch&#34;)
 *             .build());
 * 
 *         final var defaultInstanceTypes = EcsFunctions.getInstanceTypes(GetInstanceTypesArgs.builder()
 *             .availabilityZone(defaultZones.applyValue(getZonesResult -&gt; getZonesResult.zones()[0].id()))
 *             .cpuCoreCount(1)
 *             .memorySize(2)
 *             .build());
 * 
 *         final var defaultImages = EcsFunctions.getImages(GetImagesArgs.builder()
 *             .nameRegex(&#34;^ubuntu_18.*64&#34;)
 *             .mostRecent(true)
 *             .owners(&#34;system&#34;)
 *             .build());
 * 
 *         final var name = config.get(&#34;name&#34;).orElse(&#34;RouteEntryConfig&#34;);
 *         var fooNetwork = new Network(&#34;fooNetwork&#34;, NetworkArgs.builder()        
 *             .vpcName(name)
 *             .cidrBlock(&#34;10.1.0.0/21&#34;)
 *             .build());
 * 
 *         var fooSwitch = new Switch(&#34;fooSwitch&#34;, SwitchArgs.builder()        
 *             .vpcId(fooNetwork.id())
 *             .cidrBlock(&#34;10.1.1.0/24&#34;)
 *             .zoneId(defaultZones.applyValue(getZonesResult -&gt; getZonesResult.zones()[0].id()))
 *             .vswitchName(name)
 *             .build());
 * 
 *         var tfTestFoo = new SecurityGroup(&#34;tfTestFoo&#34;, SecurityGroupArgs.builder()        
 *             .description(&#34;foo&#34;)
 *             .vpcId(fooNetwork.id())
 *             .build());
 * 
 *         var ingress = new SecurityGroupRule(&#34;ingress&#34;, SecurityGroupRuleArgs.builder()        
 *             .type(&#34;ingress&#34;)
 *             .ipProtocol(&#34;tcp&#34;)
 *             .nicType(&#34;intranet&#34;)
 *             .policy(&#34;accept&#34;)
 *             .portRange(&#34;22/22&#34;)
 *             .priority(1)
 *             .securityGroupId(tfTestFoo.id())
 *             .cidrIp(&#34;0.0.0.0/0&#34;)
 *             .build());
 * 
 *         var fooInstance = new Instance(&#34;fooInstance&#34;, InstanceArgs.builder()        
 *             .securityGroups(tfTestFoo.id())
 *             .vswitchId(fooSwitch.id())
 *             .instanceChargeType(&#34;PostPaid&#34;)
 *             .instanceType(defaultInstanceTypes.applyValue(getInstanceTypesResult -&gt; getInstanceTypesResult.instanceTypes()[0].id()))
 *             .internetChargeType(&#34;PayByTraffic&#34;)
 *             .internetMaxBandwidthOut(5)
 *             .systemDiskCategory(&#34;cloud_efficiency&#34;)
 *             .imageId(defaultImages.applyValue(getImagesResult -&gt; getImagesResult.images()[0].id()))
 *             .instanceName(name)
 *             .build());
 * 
 *         var fooRouteEntry = new RouteEntry(&#34;fooRouteEntry&#34;, RouteEntryArgs.builder()        
 *             .routeTableId(fooNetwork.routeTableId())
 *             .destinationCidrblock(&#34;172.11.1.1/32&#34;)
 *             .nexthopType(&#34;Instance&#34;)
 *             .nexthopId(fooInstance.id())
 *             .build());
 * 
 *     }
 * }
 * ```
 * ## Module Support
 * 
 * You can use to the existing vpc module
 * to create a VPC, several VSwitches and add several route entries one-click.
 * 
 * ## Import
 * 
 * Router entry can be imported using the id, e.g (formatted as&lt;route_table_id:router_id:destination_cidrblock:nexthop_type:nexthop_id&gt;).
 * 
 * ```sh
 * $ pulumi import alicloud:vpc/routeEntry:RouteEntry example vtb-123456:vrt-123456:0.0.0.0/0:NatGateway:ngw-123456
 * ```
 * 
 */
@ResourceType(type="alicloud:vpc/routeEntry:RouteEntry")
public class RouteEntry extends com.pulumi.resources.CustomResource {
    /**
     * The RouteEntry&#39;s target network segment.
     * 
     */
    @Export(name="destinationCidrblock", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> destinationCidrblock;

    /**
     * @return The RouteEntry&#39;s target network segment.
     * 
     */
    public Output<Optional<String>> destinationCidrblock() {
        return Codegen.optional(this.destinationCidrblock);
    }
    /**
     * The name of the route entry. This name can have a string of 2 to 128 characters, must contain only alphanumeric characters or hyphens, such as &#34;-&#34;,&#34;.&#34;,&#34;_&#34;, and must not begin or end with a hyphen, and must not begin with http:// or https://.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The name of the route entry. This name can have a string of 2 to 128 characters, must contain only alphanumeric characters or hyphens, such as &#34;-&#34;,&#34;.&#34;,&#34;_&#34;, and must not begin or end with a hyphen, and must not begin with http:// or https://.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The route entry&#39;s next hop. ECS instance ID or VPC router interface ID.
     * 
     */
    @Export(name="nexthopId", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> nexthopId;

    /**
     * @return The route entry&#39;s next hop. ECS instance ID or VPC router interface ID.
     * 
     */
    public Output<Optional<String>> nexthopId() {
        return Codegen.optional(this.nexthopId);
    }
    /**
     * The next hop type. Available values:
     * 
     */
    @Export(name="nexthopType", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> nexthopType;

    /**
     * @return The next hop type. Available values:
     * 
     */
    public Output<Optional<String>> nexthopType() {
        return Codegen.optional(this.nexthopType);
    }
    /**
     * The ID of the route table.
     * 
     */
    @Export(name="routeTableId", refs={String.class}, tree="[0]")
    private Output<String> routeTableId;

    /**
     * @return The ID of the route table.
     * 
     */
    public Output<String> routeTableId() {
        return this.routeTableId;
    }
    /**
     * This argument has been deprecated. Please use other arguments to launch a custom route entry.
     * 
     * @deprecated
     * Attribute router_id has been deprecated and suggest removing it from your template.
     * 
     */
    @Deprecated /* Attribute router_id has been deprecated and suggest removing it from your template. */
    @Export(name="routerId", refs={String.class}, tree="[0]")
    private Output<String> routerId;

    /**
     * @return This argument has been deprecated. Please use other arguments to launch a custom route entry.
     * 
     */
    public Output<String> routerId() {
        return this.routerId;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public RouteEntry(String name) {
        this(name, RouteEntryArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public RouteEntry(String name, RouteEntryArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public RouteEntry(String name, RouteEntryArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:vpc/routeEntry:RouteEntry", name, args == null ? RouteEntryArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private RouteEntry(String name, Output<String> id, @Nullable RouteEntryState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:vpc/routeEntry:RouteEntry", name, state, makeResourceOptions(options, id));
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
    public static RouteEntry get(String name, Output<String> id, @Nullable RouteEntryState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new RouteEntry(name, id, state, options);
    }
}
