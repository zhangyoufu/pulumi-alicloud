// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.vpc;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.vpc.GatewayEndpointRouteTableAttachmentArgs;
import com.pulumi.alicloud.vpc.inputs.GatewayEndpointRouteTableAttachmentState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * Provides a VPC Gateway Endpoint Route Table Attachment resource. VPC gateway node association route.
 * 
 * For information about VPC Gateway Endpoint Route Table Attachment and how to use it, see [What is Gateway Endpoint Route Table Attachment](https://www.alibabacloud.com/help/en/virtual-private-cloud/latest/311148).
 * 
 * &gt; **NOTE:** Available since v1.208.0.
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
 * import com.pulumi.alicloud.vpc.Network;
 * import com.pulumi.alicloud.vpc.NetworkArgs;
 * import com.pulumi.alicloud.vpc.GatewayEndpoint;
 * import com.pulumi.alicloud.vpc.GatewayEndpointArgs;
 * import com.pulumi.alicloud.vpc.RouteTable;
 * import com.pulumi.alicloud.vpc.RouteTableArgs;
 * import com.pulumi.alicloud.vpc.GatewayEndpointRouteTableAttachment;
 * import com.pulumi.alicloud.vpc.GatewayEndpointRouteTableAttachmentArgs;
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
 *         final var name = config.get("name").orElse("terraform-example");
 *         var defaulteVpc = new Network("defaulteVpc", NetworkArgs.builder()        
 *             .description("test")
 *             .build());
 * 
 *         var defaultGE = new GatewayEndpoint("defaultGE", GatewayEndpointArgs.builder()        
 *             .serviceName("com.aliyun.cn-hangzhou.oss")
 *             .policyDocument("""
 *         {
 *           "Version": "1",
 *           "Statement": [{
 *             "Effect": "Allow",
 *             "Resource": ["*"],
 *             "Action": ["*"],
 *             "Principal": ["*"]
 *           }]
 *         }
 *             """)
 *             .vpcId(defaulteVpc.id())
 *             .gatewayEndpointDescrption("test-gateway-endpoint")
 *             .gatewayEndpointName(String.format("%s1", name))
 *             .build());
 * 
 *         var defaultRT = new RouteTable("defaultRT", RouteTableArgs.builder()        
 *             .vpcId(defaulteVpc.id())
 *             .routeTableName(String.format("%s2", name))
 *             .build());
 * 
 *         var default_ = new GatewayEndpointRouteTableAttachment("default", GatewayEndpointRouteTableAttachmentArgs.builder()        
 *             .gatewayEndpointId(defaultGE.id())
 *             .routeTableId(defaultRT.id())
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
 * VPC Gateway Endpoint Route Table Attachment can be imported using the id, e.g.
 * 
 * ```sh
 * $ pulumi import alicloud:vpc/gatewayEndpointRouteTableAttachment:GatewayEndpointRouteTableAttachment example &lt;gateway_endpoint_id&gt;:&lt;route_table_id&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:vpc/gatewayEndpointRouteTableAttachment:GatewayEndpointRouteTableAttachment")
public class GatewayEndpointRouteTableAttachment extends com.pulumi.resources.CustomResource {
    /**
     * The ID of the gateway endpoint instance to which you want to associate the route table.
     * 
     */
    @Export(name="gatewayEndpointId", refs={String.class}, tree="[0]")
    private Output<String> gatewayEndpointId;

    /**
     * @return The ID of the gateway endpoint instance to which you want to associate the route table.
     * 
     */
    public Output<String> gatewayEndpointId() {
        return this.gatewayEndpointId;
    }
    /**
     * Routing table ID.
     * 
     */
    @Export(name="routeTableId", refs={String.class}, tree="[0]")
    private Output<String> routeTableId;

    /**
     * @return Routing table ID.
     * 
     */
    public Output<String> routeTableId() {
        return this.routeTableId;
    }
    /**
     * Status of the gateway endpoint.
     * 
     */
    @Export(name="status", refs={String.class}, tree="[0]")
    private Output<String> status;

    /**
     * @return Status of the gateway endpoint.
     * 
     */
    public Output<String> status() {
        return this.status;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public GatewayEndpointRouteTableAttachment(String name) {
        this(name, GatewayEndpointRouteTableAttachmentArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public GatewayEndpointRouteTableAttachment(String name, GatewayEndpointRouteTableAttachmentArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public GatewayEndpointRouteTableAttachment(String name, GatewayEndpointRouteTableAttachmentArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:vpc/gatewayEndpointRouteTableAttachment:GatewayEndpointRouteTableAttachment", name, args == null ? GatewayEndpointRouteTableAttachmentArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private GatewayEndpointRouteTableAttachment(String name, Output<String> id, @Nullable GatewayEndpointRouteTableAttachmentState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:vpc/gatewayEndpointRouteTableAttachment:GatewayEndpointRouteTableAttachment", name, state, makeResourceOptions(options, id));
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
    public static GatewayEndpointRouteTableAttachment get(String name, Output<String> id, @Nullable GatewayEndpointRouteTableAttachmentState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new GatewayEndpointRouteTableAttachment(name, id, state, options);
    }
}
