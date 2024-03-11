// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.cen;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.cen.RouteEntryArgs;
import com.pulumi.alicloud.cen.inputs.RouteEntryState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * Provides a CEN route entry resource. Cloud Enterprise Network (CEN) supports publishing and withdrawing route entries of attached networks. You can publish a route entry of an attached VPC or VBR to a CEN instance, then other attached networks can learn the route if there is no route conflict. You can withdraw a published route entry when CEN does not need it any more.
 * 
 * For information about CEN route entries publishment and how to use it, see [Manage network routes](https://www.alibabacloud.com/help/doc-detail/86980.htm).
 * 
 * &gt; **NOTE:** Available since v1.20.0.
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
 * import com.pulumi.alicloud.inputs.GetRegionsArgs;
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
 * import com.pulumi.alicloud.ecs.Instance;
 * import com.pulumi.alicloud.ecs.InstanceArgs;
 * import com.pulumi.alicloud.cen.Instance;
 * import com.pulumi.alicloud.cen.InstanceArgs;
 * import com.pulumi.alicloud.cen.InstanceAttachment;
 * import com.pulumi.alicloud.cen.InstanceAttachmentArgs;
 * import com.pulumi.alicloud.vpc.RouteEntry;
 * import com.pulumi.alicloud.vpc.RouteEntryArgs;
 * import com.pulumi.alicloud.cen.RouteEntry;
 * import com.pulumi.alicloud.cen.RouteEntryArgs;
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
 *         final var default = AlicloudFunctions.getRegions(GetRegionsArgs.builder()
 *             .current(true)
 *             .build());
 * 
 *         final var exampleZones = AlicloudFunctions.getZones(GetZonesArgs.builder()
 *             .availableResourceCreation(&#34;Instance&#34;)
 *             .build());
 * 
 *         final var exampleInstanceTypes = EcsFunctions.getInstanceTypes(GetInstanceTypesArgs.builder()
 *             .availabilityZone(exampleZones.applyValue(getZonesResult -&gt; getZonesResult.zones()[0].id()))
 *             .cpuCoreCount(1)
 *             .memorySize(2)
 *             .build());
 * 
 *         final var exampleImages = EcsFunctions.getImages(GetImagesArgs.builder()
 *             .nameRegex(&#34;^ubuntu_[0-9]+_[0-9]+_x64*&#34;)
 *             .owners(&#34;system&#34;)
 *             .build());
 * 
 *         var exampleNetwork = new Network(&#34;exampleNetwork&#34;, NetworkArgs.builder()        
 *             .vpcName(&#34;terraform-example&#34;)
 *             .cidrBlock(&#34;172.17.3.0/24&#34;)
 *             .build());
 * 
 *         var exampleSwitch = new Switch(&#34;exampleSwitch&#34;, SwitchArgs.builder()        
 *             .vswitchName(&#34;terraform-example&#34;)
 *             .cidrBlock(&#34;172.17.3.0/24&#34;)
 *             .vpcId(exampleNetwork.id())
 *             .zoneId(exampleZones.applyValue(getZonesResult -&gt; getZonesResult.zones()[0].id()))
 *             .build());
 * 
 *         var exampleSecurityGroup = new SecurityGroup(&#34;exampleSecurityGroup&#34;, SecurityGroupArgs.builder()        
 *             .vpcId(exampleNetwork.id())
 *             .build());
 * 
 *         var exampleInstance = new Instance(&#34;exampleInstance&#34;, InstanceArgs.builder()        
 *             .availabilityZone(exampleZones.applyValue(getZonesResult -&gt; getZonesResult.zones()[0].id()))
 *             .instanceName(&#34;terraform-example&#34;)
 *             .imageId(exampleImages.applyValue(getImagesResult -&gt; getImagesResult.images()[0].id()))
 *             .instanceType(exampleInstanceTypes.applyValue(getInstanceTypesResult -&gt; getInstanceTypesResult.instanceTypes()[0].id()))
 *             .securityGroups(exampleSecurityGroup.id())
 *             .vswitchId(exampleSwitch.id())
 *             .internetMaxBandwidthOut(5)
 *             .build());
 * 
 *         var exampleCen_instanceInstance = new Instance(&#34;exampleCen/instanceInstance&#34;, InstanceArgs.builder()        
 *             .cenInstanceName(&#34;tf_example&#34;)
 *             .description(&#34;an example for cen&#34;)
 *             .build());
 * 
 *         var exampleInstanceAttachment = new InstanceAttachment(&#34;exampleInstanceAttachment&#34;, InstanceAttachmentArgs.builder()        
 *             .instanceId(exampleCen / instanceInstance.id())
 *             .childInstanceId(exampleNetwork.id())
 *             .childInstanceType(&#34;VPC&#34;)
 *             .childInstanceRegionId(default_.regions()[0].id())
 *             .build());
 * 
 *         var exampleRouteEntry = new RouteEntry(&#34;exampleRouteEntry&#34;, RouteEntryArgs.builder()        
 *             .routeTableId(exampleNetwork.routeTableId())
 *             .destinationCidrblock(&#34;11.0.0.0/16&#34;)
 *             .nexthopType(&#34;Instance&#34;)
 *             .nexthopId(exampleInstance.id())
 *             .build());
 * 
 *         var exampleCen_routeEntryRouteEntry = new RouteEntry(&#34;exampleCen/routeEntryRouteEntry&#34;, RouteEntryArgs.builder()        
 *             .instanceId(exampleInstanceAttachment.instanceId())
 *             .routeTableId(exampleNetwork.routeTableId())
 *             .cidrBlock(exampleRouteEntry.destinationCidrblock())
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * CEN instance can be imported using the id, e.g.
 * 
 * ```sh
 * $ pulumi import alicloud:cen/routeEntry:RouteEntry example cen-abc123456:vtb-abc123:192.168.0.0/24
 * ```
 * 
 */
@ResourceType(type="alicloud:cen/routeEntry:RouteEntry")
public class RouteEntry extends com.pulumi.resources.CustomResource {
    /**
     * The destination CIDR block of the route entry to publish.
     * 
     * -&gt;**NOTE:** The &#34;alicloud_cen_instance_route_entries&#34; resource depends on the related &#34;alicloud.cen.InstanceAttachment&#34; resource.
     * 
     * -&gt;**NOTE:** The &#34;alicloud.cen.InstanceAttachment&#34; resource should depend on the related &#34;alicloud.vpc.Switch&#34; resource.
     * 
     */
    @Export(name="cidrBlock", refs={String.class}, tree="[0]")
    private Output<String> cidrBlock;

    /**
     * @return The destination CIDR block of the route entry to publish.
     * 
     * -&gt;**NOTE:** The &#34;alicloud_cen_instance_route_entries&#34; resource depends on the related &#34;alicloud.cen.InstanceAttachment&#34; resource.
     * 
     * -&gt;**NOTE:** The &#34;alicloud.cen.InstanceAttachment&#34; resource should depend on the related &#34;alicloud.vpc.Switch&#34; resource.
     * 
     */
    public Output<String> cidrBlock() {
        return this.cidrBlock;
    }
    /**
     * The ID of the CEN.
     * 
     */
    @Export(name="instanceId", refs={String.class}, tree="[0]")
    private Output<String> instanceId;

    /**
     * @return The ID of the CEN.
     * 
     */
    public Output<String> instanceId() {
        return this.instanceId;
    }
    /**
     * The route table of the attached VBR or VPC.
     * 
     */
    @Export(name="routeTableId", refs={String.class}, tree="[0]")
    private Output<String> routeTableId;

    /**
     * @return The route table of the attached VBR or VPC.
     * 
     */
    public Output<String> routeTableId() {
        return this.routeTableId;
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
        super("alicloud:cen/routeEntry:RouteEntry", name, args == null ? RouteEntryArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private RouteEntry(String name, Output<String> id, @Nullable RouteEntryState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:cen/routeEntry:RouteEntry", name, state, makeResourceOptions(options, id));
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
