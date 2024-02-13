// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.cen;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.cen.ChildInstanceRouteEntryToAttachmentArgs;
import com.pulumi.alicloud.cen.inputs.ChildInstanceRouteEntryToAttachmentState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Boolean;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a Cen Child Instance Route Entry To Attachment resource.
 * 
 * For information about Cen Child Instance Route Entry To Attachment and how to use it, see [What is Child Instance Route Entry To Attachment](https://www.alibabacloud.com/help/en/cen/developer-reference/api-cbn-2017-09-12-createcenchildinstancerouteentrytoattachment).
 * 
 * &gt; **NOTE:** Available since v1.195.0.
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
 * import com.pulumi.alicloud.cen.CenFunctions;
 * import com.pulumi.alicloud.cen.inputs.GetTransitRouterAvailableResourcesArgs;
 * import com.pulumi.alicloud.vpc.Network;
 * import com.pulumi.alicloud.vpc.NetworkArgs;
 * import com.pulumi.alicloud.vpc.Switch;
 * import com.pulumi.alicloud.vpc.SwitchArgs;
 * import com.pulumi.alicloud.cen.Instance;
 * import com.pulumi.alicloud.cen.InstanceArgs;
 * import com.pulumi.alicloud.cen.TransitRouter;
 * import com.pulumi.alicloud.cen.TransitRouterArgs;
 * import com.pulumi.alicloud.cen.TransitRouterVpcAttachment;
 * import com.pulumi.alicloud.cen.TransitRouterVpcAttachmentArgs;
 * import com.pulumi.alicloud.cen.inputs.TransitRouterVpcAttachmentZoneMappingArgs;
 * import com.pulumi.alicloud.vpc.RouteTable;
 * import com.pulumi.alicloud.vpc.RouteTableArgs;
 * import com.pulumi.alicloud.cen.ChildInstanceRouteEntryToAttachment;
 * import com.pulumi.alicloud.cen.ChildInstanceRouteEntryToAttachmentArgs;
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
 *         final var name = config.get(&#34;name&#34;).orElse(&#34;terraform-example&#34;);
 *         final var default = CenFunctions.getTransitRouterAvailableResources();
 * 
 *         final var masterZone = default_.resources()[0].masterZones()[0];
 * 
 *         final var slaveZone = default_.resources()[0].slaveZones()[1];
 * 
 *         var exampleNetwork = new Network(&#34;exampleNetwork&#34;, NetworkArgs.builder()        
 *             .vpcName(name)
 *             .cidrBlock(&#34;192.168.0.0/16&#34;)
 *             .build());
 * 
 *         var exampleMaster = new Switch(&#34;exampleMaster&#34;, SwitchArgs.builder()        
 *             .vswitchName(name)
 *             .cidrBlock(&#34;192.168.1.0/24&#34;)
 *             .vpcId(exampleNetwork.id())
 *             .zoneId(masterZone)
 *             .build());
 * 
 *         var exampleSlave = new Switch(&#34;exampleSlave&#34;, SwitchArgs.builder()        
 *             .vswitchName(name)
 *             .cidrBlock(&#34;192.168.2.0/24&#34;)
 *             .vpcId(exampleNetwork.id())
 *             .zoneId(slaveZone)
 *             .build());
 * 
 *         var exampleInstance = new Instance(&#34;exampleInstance&#34;, InstanceArgs.builder()        
 *             .cenInstanceName(name)
 *             .protectionLevel(&#34;REDUCED&#34;)
 *             .build());
 * 
 *         var exampleTransitRouter = new TransitRouter(&#34;exampleTransitRouter&#34;, TransitRouterArgs.builder()        
 *             .transitRouterName(name)
 *             .cenId(exampleInstance.id())
 *             .build());
 * 
 *         var exampleTransitRouterVpcAttachment = new TransitRouterVpcAttachment(&#34;exampleTransitRouterVpcAttachment&#34;, TransitRouterVpcAttachmentArgs.builder()        
 *             .cenId(exampleInstance.id())
 *             .transitRouterId(exampleTransitRouter.transitRouterId())
 *             .vpcId(exampleNetwork.id())
 *             .zoneMappings(            
 *                 TransitRouterVpcAttachmentZoneMappingArgs.builder()
 *                     .zoneId(masterZone)
 *                     .vswitchId(exampleMaster.id())
 *                     .build(),
 *                 TransitRouterVpcAttachmentZoneMappingArgs.builder()
 *                     .zoneId(slaveZone)
 *                     .vswitchId(exampleSlave.id())
 *                     .build())
 *             .transitRouterAttachmentName(name)
 *             .transitRouterAttachmentDescription(name)
 *             .build());
 * 
 *         var exampleRouteTable = new RouteTable(&#34;exampleRouteTable&#34;, RouteTableArgs.builder()        
 *             .vpcId(exampleNetwork.id())
 *             .routeTableName(name)
 *             .description(name)
 *             .build());
 * 
 *         var exampleChildInstanceRouteEntryToAttachment = new ChildInstanceRouteEntryToAttachment(&#34;exampleChildInstanceRouteEntryToAttachment&#34;, ChildInstanceRouteEntryToAttachmentArgs.builder()        
 *             .transitRouterAttachmentId(exampleTransitRouterVpcAttachment.transitRouterAttachmentId())
 *             .cenId(exampleInstance.id())
 *             .destinationCidrBlock(&#34;10.0.0.0/24&#34;)
 *             .childInstanceRouteTableId(exampleRouteTable.id())
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * Cen Child Instance Route Entry To Attachment can be imported using the id, e.g.
 * 
 * ```sh
 * $ pulumi import alicloud:cen/childInstanceRouteEntryToAttachment:ChildInstanceRouteEntryToAttachment example &lt;cen_id&gt;:&lt;child_instance_route_table_id&gt;:&lt;transit_router_attachment_id&gt;:&lt;destination_cidr_block&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:cen/childInstanceRouteEntryToAttachment:ChildInstanceRouteEntryToAttachment")
public class ChildInstanceRouteEntryToAttachment extends com.pulumi.resources.CustomResource {
    /**
     * The ID of the CEN instance.
     * 
     */
    @Export(name="cenId", refs={String.class}, tree="[0]")
    private Output<String> cenId;

    /**
     * @return The ID of the CEN instance.
     * 
     */
    public Output<String> cenId() {
        return this.cenId;
    }
    /**
     * The first ID of the resource
     * 
     */
    @Export(name="childInstanceRouteTableId", refs={String.class}, tree="[0]")
    private Output<String> childInstanceRouteTableId;

    /**
     * @return The first ID of the resource
     * 
     */
    public Output<String> childInstanceRouteTableId() {
        return this.childInstanceRouteTableId;
    }
    /**
     * DestinationCidrBlock
     * 
     */
    @Export(name="destinationCidrBlock", refs={String.class}, tree="[0]")
    private Output<String> destinationCidrBlock;

    /**
     * @return DestinationCidrBlock
     * 
     */
    public Output<String> destinationCidrBlock() {
        return this.destinationCidrBlock;
    }
    /**
     * Whether to perform pre-check on this request, including permission and instance status verification.
     * 
     */
    @Export(name="dryRun", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> dryRun;

    /**
     * @return Whether to perform pre-check on this request, including permission and instance status verification.
     * 
     */
    public Output<Optional<Boolean>> dryRun() {
        return Codegen.optional(this.dryRun);
    }
    /**
     * ServiceType
     * 
     */
    @Export(name="serviceType", refs={String.class}, tree="[0]")
    private Output<String> serviceType;

    /**
     * @return ServiceType
     * 
     */
    public Output<String> serviceType() {
        return this.serviceType;
    }
    /**
     * The status of the resource
     * 
     */
    @Export(name="status", refs={String.class}, tree="[0]")
    private Output<String> status;

    /**
     * @return The status of the resource
     * 
     */
    public Output<String> status() {
        return this.status;
    }
    /**
     * TransitRouterAttachmentId
     * 
     */
    @Export(name="transitRouterAttachmentId", refs={String.class}, tree="[0]")
    private Output<String> transitRouterAttachmentId;

    /**
     * @return TransitRouterAttachmentId
     * 
     */
    public Output<String> transitRouterAttachmentId() {
        return this.transitRouterAttachmentId;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public ChildInstanceRouteEntryToAttachment(String name) {
        this(name, ChildInstanceRouteEntryToAttachmentArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ChildInstanceRouteEntryToAttachment(String name, ChildInstanceRouteEntryToAttachmentArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ChildInstanceRouteEntryToAttachment(String name, ChildInstanceRouteEntryToAttachmentArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:cen/childInstanceRouteEntryToAttachment:ChildInstanceRouteEntryToAttachment", name, args == null ? ChildInstanceRouteEntryToAttachmentArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private ChildInstanceRouteEntryToAttachment(String name, Output<String> id, @Nullable ChildInstanceRouteEntryToAttachmentState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:cen/childInstanceRouteEntryToAttachment:ChildInstanceRouteEntryToAttachment", name, state, makeResourceOptions(options, id));
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
    public static ChildInstanceRouteEntryToAttachment get(String name, Output<String> id, @Nullable ChildInstanceRouteEntryToAttachmentState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ChildInstanceRouteEntryToAttachment(name, id, state, options);
    }
}
