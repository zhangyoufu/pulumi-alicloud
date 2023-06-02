// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.vpc;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.vpc.GatewayRouteTableAttachmentArgs;
import com.pulumi.alicloud.vpc.inputs.GatewayRouteTableAttachmentState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Boolean;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a VPC Gateway Route Table Attachment resource.
 * 
 * For information about VPC Gateway Route Table Attachment and how to use it, see [What is Gateway Route Table Attachment](https://www.alibabacloud.com/help/doc-detail/174112.htm).
 * 
 * &gt; **NOTE:** Available in v1.194.0+.
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
 * import com.pulumi.alicloud.vpc.Network;
 * import com.pulumi.alicloud.vpc.NetworkArgs;
 * import com.pulumi.alicloud.vpc.RouteTable;
 * import com.pulumi.alicloud.vpc.RouteTableArgs;
 * import com.pulumi.alicloud.vpc.Ipv4Gateway;
 * import com.pulumi.alicloud.vpc.Ipv4GatewayArgs;
 * import com.pulumi.alicloud.vpc.GatewayRouteTableAttachment;
 * import com.pulumi.alicloud.vpc.GatewayRouteTableAttachmentArgs;
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
 *         var exampleNetwork = new Network(&#34;exampleNetwork&#34;, NetworkArgs.builder()        
 *             .cidrBlock(&#34;172.16.0.0/12&#34;)
 *             .vpcName(&#34;terraform-example&#34;)
 *             .build());
 * 
 *         var exampleRouteTable = new RouteTable(&#34;exampleRouteTable&#34;, RouteTableArgs.builder()        
 *             .vpcId(exampleNetwork.id())
 *             .routeTableName(&#34;terraform-example&#34;)
 *             .description(&#34;terraform-example&#34;)
 *             .associateType(&#34;Gateway&#34;)
 *             .build());
 * 
 *         var exampleIpv4Gateway = new Ipv4Gateway(&#34;exampleIpv4Gateway&#34;, Ipv4GatewayArgs.builder()        
 *             .ipv4GatewayName(&#34;terraform-example&#34;)
 *             .vpcId(exampleNetwork.id())
 *             .enabled(&#34;true&#34;)
 *             .build());
 * 
 *         var exampleGatewayRouteTableAttachment = new GatewayRouteTableAttachment(&#34;exampleGatewayRouteTableAttachment&#34;, GatewayRouteTableAttachmentArgs.builder()        
 *             .ipv4GatewayId(exampleIpv4Gateway.id())
 *             .routeTableId(exampleRouteTable.id())
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * VPC Gateway Route Table Attachment can be imported using the id, e.g.
 * 
 * ```sh
 *  $ pulumi import alicloud:vpc/gatewayRouteTableAttachment:GatewayRouteTableAttachment example &lt;route_table_id&gt;:&lt;ipv4_gateway_id&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:vpc/gatewayRouteTableAttachment:GatewayRouteTableAttachment")
public class GatewayRouteTableAttachment extends com.pulumi.resources.CustomResource {
    /**
     * The creation time of the resource.
     * 
     */
    @Export(name="createTime", type=String.class, parameters={})
    private Output<String> createTime;

    /**
     * @return The creation time of the resource.
     * 
     */
    public Output<String> createTime() {
        return this.createTime;
    }
    /**
     * Specifies whether to only precheck this request. Default value: `false`.
     * 
     */
    @Export(name="dryRun", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> dryRun;

    /**
     * @return Specifies whether to only precheck this request. Default value: `false`.
     * 
     */
    public Output<Optional<Boolean>> dryRun() {
        return Codegen.optional(this.dryRun);
    }
    /**
     * The ID of the IPv4 Gateway instance.
     * 
     */
    @Export(name="ipv4GatewayId", type=String.class, parameters={})
    private Output<String> ipv4GatewayId;

    /**
     * @return The ID of the IPv4 Gateway instance.
     * 
     */
    public Output<String> ipv4GatewayId() {
        return this.ipv4GatewayId;
    }
    /**
     * The ID of the Gateway route table to be bound.
     * 
     */
    @Export(name="routeTableId", type=String.class, parameters={})
    private Output<String> routeTableId;

    /**
     * @return The ID of the Gateway route table to be bound.
     * 
     */
    public Output<String> routeTableId() {
        return this.routeTableId;
    }
    /**
     * The status of the IPv4 Gateway instance. Value:
     * - **Creating**: The function is being created.
     * - **Created**: Created and available.
     * - **Modifying**: is being modified.
     * - **Deleting**: Deleting.
     * - **Deleted**: Deleted.
     * - **Activating**: enabled.
     * 
     */
    @Export(name="status", type=String.class, parameters={})
    private Output<String> status;

    /**
     * @return The status of the IPv4 Gateway instance. Value:
     * - **Creating**: The function is being created.
     * - **Created**: Created and available.
     * - **Modifying**: is being modified.
     * - **Deleting**: Deleting.
     * - **Deleted**: Deleted.
     * - **Activating**: enabled.
     * 
     */
    public Output<String> status() {
        return this.status;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public GatewayRouteTableAttachment(String name) {
        this(name, GatewayRouteTableAttachmentArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public GatewayRouteTableAttachment(String name, GatewayRouteTableAttachmentArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public GatewayRouteTableAttachment(String name, GatewayRouteTableAttachmentArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:vpc/gatewayRouteTableAttachment:GatewayRouteTableAttachment", name, args == null ? GatewayRouteTableAttachmentArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private GatewayRouteTableAttachment(String name, Output<String> id, @Nullable GatewayRouteTableAttachmentState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:vpc/gatewayRouteTableAttachment:GatewayRouteTableAttachment", name, state, makeResourceOptions(options, id));
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
    public static GatewayRouteTableAttachment get(String name, Output<String> id, @Nullable GatewayRouteTableAttachmentState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new GatewayRouteTableAttachment(name, id, state, options);
    }
}
