// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.cen;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.cen.TransitRouterRouteEntryArgs;
import com.pulumi.alicloud.cen.inputs.TransitRouterRouteEntryState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Boolean;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a CEN transit router route entry resource.[What is Cen Transit Router Route Entry](https://www.alibabacloud.com/help/en/cloud-enterprise-network/latest/api-cbn-2017-09-12-createtransitrouterrouteentry)
 * 
 * &gt; **NOTE:** Available since v1.126.0.
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
 * import com.pulumi.alicloud.cen.Instance;
 * import com.pulumi.alicloud.cen.InstanceArgs;
 * import com.pulumi.alicloud.cen.TransitRouter;
 * import com.pulumi.alicloud.cen.TransitRouterArgs;
 * import com.pulumi.alicloud.cen.TransitRouterRouteTable;
 * import com.pulumi.alicloud.cen.TransitRouterRouteTableArgs;
 * import com.pulumi.alicloud.expressconnect.ExpressconnectFunctions;
 * import com.pulumi.alicloud.expressconnect.inputs.GetPhysicalConnectionsArgs;
 * import com.pulumi.random.integer;
 * import com.pulumi.random.IntegerArgs;
 * import com.pulumi.alicloud.expressconnect.VirtualBorderRouter;
 * import com.pulumi.alicloud.expressconnect.VirtualBorderRouterArgs;
 * import com.pulumi.alicloud.cen.TransitRouterVbrAttachment;
 * import com.pulumi.alicloud.cen.TransitRouterVbrAttachmentArgs;
 * import com.pulumi.alicloud.cen.TransitRouterRouteEntry;
 * import com.pulumi.alicloud.cen.TransitRouterRouteEntryArgs;
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
 *         var exampleInstance = new Instance(&#34;exampleInstance&#34;, InstanceArgs.builder()        
 *             .cenInstanceName(name)
 *             .description(&#34;an example for cen&#34;)
 *             .build());
 * 
 *         var exampleTransitRouter = new TransitRouter(&#34;exampleTransitRouter&#34;, TransitRouterArgs.builder()        
 *             .transitRouterName(name)
 *             .cenId(exampleInstance.id())
 *             .build());
 * 
 *         var exampleTransitRouterRouteTable = new TransitRouterRouteTable(&#34;exampleTransitRouterRouteTable&#34;, TransitRouterRouteTableArgs.builder()        
 *             .transitRouterId(exampleTransitRouter.transitRouterId())
 *             .build());
 * 
 *         final var example = ExpressconnectFunctions.getPhysicalConnections(GetPhysicalConnectionsArgs.builder()
 *             .nameRegex(&#34;^preserved-NODELETING&#34;)
 *             .build());
 * 
 *         var vlanId = new Integer(&#34;vlanId&#34;, IntegerArgs.builder()        
 *             .max(2999)
 *             .min(1)
 *             .build());
 * 
 *         var exampleVirtualBorderRouter = new VirtualBorderRouter(&#34;exampleVirtualBorderRouter&#34;, VirtualBorderRouterArgs.builder()        
 *             .localGatewayIp(&#34;10.0.0.1&#34;)
 *             .peerGatewayIp(&#34;10.0.0.2&#34;)
 *             .peeringSubnetMask(&#34;255.255.255.252&#34;)
 *             .physicalConnectionId(example.applyValue(getPhysicalConnectionsResult -&gt; getPhysicalConnectionsResult.connections()[0].id()))
 *             .virtualBorderRouterName(name)
 *             .vlanId(vlanId.id())
 *             .minRxInterval(1000)
 *             .minTxInterval(1000)
 *             .detectMultiplier(10)
 *             .build());
 * 
 *         var exampleTransitRouterVbrAttachment = new TransitRouterVbrAttachment(&#34;exampleTransitRouterVbrAttachment&#34;, TransitRouterVbrAttachmentArgs.builder()        
 *             .vbrId(exampleVirtualBorderRouter.id())
 *             .cenId(exampleInstance.id())
 *             .transitRouterId(exampleTransitRouter.transitRouterId())
 *             .autoPublishRouteEnabled(true)
 *             .transitRouterAttachmentName(name)
 *             .transitRouterAttachmentDescription(name)
 *             .build());
 * 
 *         var exampleTransitRouterRouteEntry = new TransitRouterRouteEntry(&#34;exampleTransitRouterRouteEntry&#34;, TransitRouterRouteEntryArgs.builder()        
 *             .transitRouterRouteTableId(exampleTransitRouterRouteTable.transitRouterRouteTableId())
 *             .transitRouterRouteEntryDestinationCidrBlock(&#34;192.168.0.0/24&#34;)
 *             .transitRouterRouteEntryNextHopType(&#34;Attachment&#34;)
 *             .transitRouterRouteEntryName(name)
 *             .transitRouterRouteEntryDescription(name)
 *             .transitRouterRouteEntryNextHopId(exampleTransitRouterVbrAttachment.transitRouterAttachmentId())
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
 * $ pulumi import alicloud:cen/transitRouterRouteEntry:TransitRouterRouteEntry default vtb-*********:rte-*******
 * ```
 * 
 */
@ResourceType(type="alicloud:cen/transitRouterRouteEntry:TransitRouterRouteEntry")
public class TransitRouterRouteEntry extends com.pulumi.resources.CustomResource {
    /**
     * The dry run.
     * 
     * &gt; **NOTE:** If transit_router_route_entry_next_hop_type is `Attachment`, transit_router_route_entry_next_hop_id is required.
     * If transit_router_route_entry_next_hop_type is `BlackHole`, transit_router_route_entry_next_hop_id cannot be filled.
     * 
     */
    @Export(name="dryRun", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> dryRun;

    /**
     * @return The dry run.
     * 
     * &gt; **NOTE:** If transit_router_route_entry_next_hop_type is `Attachment`, transit_router_route_entry_next_hop_id is required.
     * If transit_router_route_entry_next_hop_type is `BlackHole`, transit_router_route_entry_next_hop_id cannot be filled.
     * 
     */
    public Output<Optional<Boolean>> dryRun() {
        return Codegen.optional(this.dryRun);
    }
    /**
     * The associating status of the Transit Router.
     * 
     */
    @Export(name="status", refs={String.class}, tree="[0]")
    private Output<String> status;

    /**
     * @return The associating status of the Transit Router.
     * 
     */
    public Output<String> status() {
        return this.status;
    }
    /**
     * The description of the transit router route entry.
     * 
     */
    @Export(name="transitRouterRouteEntryDescription", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> transitRouterRouteEntryDescription;

    /**
     * @return The description of the transit router route entry.
     * 
     */
    public Output<Optional<String>> transitRouterRouteEntryDescription() {
        return Codegen.optional(this.transitRouterRouteEntryDescription);
    }
    /**
     * The CIDR of the transit router route entry.
     * 
     */
    @Export(name="transitRouterRouteEntryDestinationCidrBlock", refs={String.class}, tree="[0]")
    private Output<String> transitRouterRouteEntryDestinationCidrBlock;

    /**
     * @return The CIDR of the transit router route entry.
     * 
     */
    public Output<String> transitRouterRouteEntryDestinationCidrBlock() {
        return this.transitRouterRouteEntryDestinationCidrBlock;
    }
    /**
     * The ID of the route entry.
     * 
     */
    @Export(name="transitRouterRouteEntryId", refs={String.class}, tree="[0]")
    private Output<String> transitRouterRouteEntryId;

    /**
     * @return The ID of the route entry.
     * 
     */
    public Output<String> transitRouterRouteEntryId() {
        return this.transitRouterRouteEntryId;
    }
    /**
     * The name of the transit router route entry.
     * 
     */
    @Export(name="transitRouterRouteEntryName", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> transitRouterRouteEntryName;

    /**
     * @return The name of the transit router route entry.
     * 
     */
    public Output<Optional<String>> transitRouterRouteEntryName() {
        return Codegen.optional(this.transitRouterRouteEntryName);
    }
    /**
     * The ID of the transit router route entry next hop.
     * 
     */
    @Export(name="transitRouterRouteEntryNextHopId", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> transitRouterRouteEntryNextHopId;

    /**
     * @return The ID of the transit router route entry next hop.
     * 
     */
    public Output<Optional<String>> transitRouterRouteEntryNextHopId() {
        return Codegen.optional(this.transitRouterRouteEntryNextHopId);
    }
    /**
     * The Type of the transit router route entry next hop,Valid values `Attachment` and `BlackHole`.
     * 
     */
    @Export(name="transitRouterRouteEntryNextHopType", refs={String.class}, tree="[0]")
    private Output<String> transitRouterRouteEntryNextHopType;

    /**
     * @return The Type of the transit router route entry next hop,Valid values `Attachment` and `BlackHole`.
     * 
     */
    public Output<String> transitRouterRouteEntryNextHopType() {
        return this.transitRouterRouteEntryNextHopType;
    }
    /**
     * The ID of the transit router route table.
     * 
     */
    @Export(name="transitRouterRouteTableId", refs={String.class}, tree="[0]")
    private Output<String> transitRouterRouteTableId;

    /**
     * @return The ID of the transit router route table.
     * 
     */
    public Output<String> transitRouterRouteTableId() {
        return this.transitRouterRouteTableId;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public TransitRouterRouteEntry(String name) {
        this(name, TransitRouterRouteEntryArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public TransitRouterRouteEntry(String name, TransitRouterRouteEntryArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public TransitRouterRouteEntry(String name, TransitRouterRouteEntryArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:cen/transitRouterRouteEntry:TransitRouterRouteEntry", name, args == null ? TransitRouterRouteEntryArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private TransitRouterRouteEntry(String name, Output<String> id, @Nullable TransitRouterRouteEntryState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:cen/transitRouterRouteEntry:TransitRouterRouteEntry", name, state, makeResourceOptions(options, id));
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
    public static TransitRouterRouteEntry get(String name, Output<String> id, @Nullable TransitRouterRouteEntryState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new TransitRouterRouteEntry(name, id, state, options);
    }
}
