// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.vpc;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.vpc.RouteTableArgs;
import com.pulumi.alicloud.vpc.inputs.RouteTableState;
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
 * Provides a VPC Route Table resource. Currently, customized route tables are available in most regions apart from China (Beijing), China (Hangzhou), and China (Shenzhen) regions.
 * 
 * For information about VPC Route Table and how to use it, see [What is Route Table](https://www.alibabacloud.com/help/doc-detail/87057.htm).
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
 *         var defaultVpc = new Network(&#34;defaultVpc&#34;, NetworkArgs.builder()        
 *             .vpcName(name)
 *             .build());
 * 
 *         var default_ = new RouteTable(&#34;default&#34;, RouteTableArgs.builder()        
 *             .description(&#34;test-description&#34;)
 *             .vpcId(defaultVpc.id())
 *             .routeTableName(name)
 *             .associateType(&#34;VSwitch&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * VPC Route Table can be imported using the id, e.g.
 * 
 * ```sh
 *  $ pulumi import alicloud:vpc/routeTable:RouteTable example &lt;id&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:vpc/routeTable:RouteTable")
public class RouteTable extends com.pulumi.resources.CustomResource {
    /**
     * The type of cloud resource that is bound to the routing table. Value:
     * - **VSwitch**: switch.
     * - **Gateway**:IPv4 Gateway.
     * 
     */
    @Export(name="associateType", type=String.class, parameters={})
    private Output<String> associateType;

    /**
     * @return The type of cloud resource that is bound to the routing table. Value:
     * - **VSwitch**: switch.
     * - **Gateway**:IPv4 Gateway.
     * 
     */
    public Output<String> associateType() {
        return this.associateType;
    }
    /**
     * The creation time of the routing table.
     * 
     */
    @Export(name="createTime", type=String.class, parameters={})
    private Output<String> createTime;

    /**
     * @return The creation time of the routing table.
     * 
     */
    public Output<String> createTime() {
        return this.createTime;
    }
    /**
     * Description of the routing table.
     * 
     */
    @Export(name="description", type=String.class, parameters={})
    private Output</* @Nullable */ String> description;

    /**
     * @return Description of the routing table.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * Field &#39;name&#39; has been deprecated from provider version 1.119.1. New field &#39;route_table_name&#39; instead.
     * 
     * @deprecated
     * Field &#39;name&#39; has been deprecated from provider version 1.119.1. New field &#39;route_table_name&#39; instead.
     * 
     */
    @Deprecated /* Field 'name' has been deprecated from provider version 1.119.1. New field 'route_table_name' instead. */
    @Export(name="name", type=String.class, parameters={})
    private Output<String> name;

    /**
     * @return Field &#39;name&#39; has been deprecated from provider version 1.119.1. New field &#39;route_table_name&#39; instead.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * Resource group ID.
     * 
     */
    @Export(name="resourceGroupId", type=String.class, parameters={})
    private Output<String> resourceGroupId;

    /**
     * @return Resource group ID.
     * 
     */
    public Output<String> resourceGroupId() {
        return this.resourceGroupId;
    }
    /**
     * The name of the routing table.
     * 
     */
    @Export(name="routeTableName", type=String.class, parameters={})
    private Output<String> routeTableName;

    /**
     * @return The name of the routing table.
     * 
     */
    public Output<String> routeTableName() {
        return this.routeTableName;
    }
    /**
     * Routing table state.
     * 
     */
    @Export(name="status", type=String.class, parameters={})
    private Output<String> status;

    /**
     * @return Routing table state.
     * 
     */
    public Output<String> status() {
        return this.status;
    }
    /**
     * The tag.
     * 
     */
    @Export(name="tags", type=Map.class, parameters={String.class, Object.class})
    private Output</* @Nullable */ Map<String,Object>> tags;

    /**
     * @return The tag.
     * 
     */
    public Output<Optional<Map<String,Object>>> tags() {
        return Codegen.optional(this.tags);
    }
    /**
     * The ID of VPC.
     * 
     * The following arguments will be discarded. Please use new fields as soon as possible:
     * 
     */
    @Export(name="vpcId", type=String.class, parameters={})
    private Output<String> vpcId;

    /**
     * @return The ID of VPC.
     * 
     * The following arguments will be discarded. Please use new fields as soon as possible:
     * 
     */
    public Output<String> vpcId() {
        return this.vpcId;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public RouteTable(String name) {
        this(name, RouteTableArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public RouteTable(String name, RouteTableArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public RouteTable(String name, RouteTableArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:vpc/routeTable:RouteTable", name, args == null ? RouteTableArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private RouteTable(String name, Output<String> id, @Nullable RouteTableState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:vpc/routeTable:RouteTable", name, state, makeResourceOptions(options, id));
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
    public static RouteTable get(String name, Output<String> id, @Nullable RouteTableState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new RouteTable(name, id, state, options);
    }
}
