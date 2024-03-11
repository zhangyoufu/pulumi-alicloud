// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.vpc;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.vpc.RouteTableAttachmentArgs;
import com.pulumi.alicloud.vpc.inputs.RouteTableAttachmentState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * Provides a VPC Route Table Attachment resource. Routing table associated resource type.
 * 
 * For information about VPC Route Table Attachment and how to use it, see [What is Route Table Attachment](https://www.alibabacloud.com/help/doc-detail/174112.htm).
 * 
 * &gt; **NOTE:** Available since v1.194.0.
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
 * import com.pulumi.alicloud.vpc.Network;
 * import com.pulumi.alicloud.vpc.NetworkArgs;
 * import com.pulumi.alicloud.AlicloudFunctions;
 * import com.pulumi.alicloud.inputs.GetZonesArgs;
 * import com.pulumi.alicloud.vpc.Switch;
 * import com.pulumi.alicloud.vpc.SwitchArgs;
 * import com.pulumi.alicloud.vpc.RouteTable;
 * import com.pulumi.alicloud.vpc.RouteTableArgs;
 * import com.pulumi.alicloud.vpc.RouteTableAttachment;
 * import com.pulumi.alicloud.vpc.RouteTableAttachmentArgs;
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
 *         var fooNetwork = new Network(&#34;fooNetwork&#34;, NetworkArgs.builder()        
 *             .cidrBlock(&#34;172.16.0.0/12&#34;)
 *             .build());
 * 
 *         final var default = AlicloudFunctions.getZones(GetZonesArgs.builder()
 *             .availableResourceCreation(&#34;VSwitch&#34;)
 *             .build());
 * 
 *         var fooSwitch = new Switch(&#34;fooSwitch&#34;, SwitchArgs.builder()        
 *             .vpcId(fooNetwork.id())
 *             .cidrBlock(&#34;172.16.0.0/21&#34;)
 *             .zoneId(default_.zones()[0].id())
 *             .build());
 * 
 *         var fooRouteTable = new RouteTable(&#34;fooRouteTable&#34;, RouteTableArgs.builder()        
 *             .vpcId(fooNetwork.id())
 *             .routeTableName(name)
 *             .description(&#34;route_table_attachment&#34;)
 *             .build());
 * 
 *         var fooRouteTableAttachment = new RouteTableAttachment(&#34;fooRouteTableAttachment&#34;, RouteTableAttachmentArgs.builder()        
 *             .vswitchId(fooSwitch.id())
 *             .routeTableId(fooRouteTable.id())
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * VPC Route Table Attachment can be imported using the id, e.g.
 * 
 * ```sh
 * $ pulumi import alicloud:vpc/routeTableAttachment:RouteTableAttachment example &lt;route_table_id&gt;:&lt;vswitch_id&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:vpc/routeTableAttachment:RouteTableAttachment")
public class RouteTableAttachment extends com.pulumi.resources.CustomResource {
    /**
     * The ID of the route table to be bound to the switch.
     * 
     */
    @Export(name="routeTableId", refs={String.class}, tree="[0]")
    private Output<String> routeTableId;

    /**
     * @return The ID of the route table to be bound to the switch.
     * 
     */
    public Output<String> routeTableId() {
        return this.routeTableId;
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
     * The ID of the switch to bind the route table.
     * 
     */
    @Export(name="vswitchId", refs={String.class}, tree="[0]")
    private Output<String> vswitchId;

    /**
     * @return The ID of the switch to bind the route table.
     * 
     */
    public Output<String> vswitchId() {
        return this.vswitchId;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public RouteTableAttachment(String name) {
        this(name, RouteTableAttachmentArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public RouteTableAttachment(String name, RouteTableAttachmentArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public RouteTableAttachment(String name, RouteTableAttachmentArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:vpc/routeTableAttachment:RouteTableAttachment", name, args == null ? RouteTableAttachmentArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private RouteTableAttachment(String name, Output<String> id, @Nullable RouteTableAttachmentState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:vpc/routeTableAttachment:RouteTableAttachment", name, state, makeResourceOptions(options, id));
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
    public static RouteTableAttachment get(String name, Output<String> id, @Nullable RouteTableAttachmentState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new RouteTableAttachment(name, id, state, options);
    }
}
