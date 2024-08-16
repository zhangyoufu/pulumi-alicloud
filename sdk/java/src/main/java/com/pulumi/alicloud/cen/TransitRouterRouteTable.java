// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.cen;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.cen.TransitRouterRouteTableArgs;
import com.pulumi.alicloud.cen.inputs.TransitRouterRouteTableState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Boolean;
import java.lang.String;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a CEN transit router route table resource.[What is Cen Transit Router Route Table](https://www.alibabacloud.com/help/en/cen/developer-reference/api-cbn-2017-09-12-createtransitrouterroutetable)
 * 
 * &gt; **NOTE:** Available since v1.126.0.
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
 * import com.pulumi.alicloud.cen.Instance;
 * import com.pulumi.alicloud.cen.InstanceArgs;
 * import com.pulumi.alicloud.cen.TransitRouter;
 * import com.pulumi.alicloud.cen.TransitRouterArgs;
 * import com.pulumi.alicloud.cen.TransitRouterRouteTable;
 * import com.pulumi.alicloud.cen.TransitRouterRouteTableArgs;
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
 *         var example = new Instance("example", InstanceArgs.builder()
 *             .cenInstanceName("tf_example")
 *             .description("an example for cen")
 *             .build());
 * 
 *         var exampleTransitRouter = new TransitRouter("exampleTransitRouter", TransitRouterArgs.builder()
 *             .transitRouterName("tf_example")
 *             .cenId(example.id())
 *             .build());
 * 
 *         var exampleTransitRouterRouteTable = new TransitRouterRouteTable("exampleTransitRouterRouteTable", TransitRouterRouteTableArgs.builder()
 *             .transitRouterId(exampleTransitRouter.transitRouterId())
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
 * CEN transit router route table  can be imported using the id, e.g.
 * 
 * ```sh
 * $ pulumi import alicloud:cen/transitRouterRouteTable:TransitRouterRouteTable default tr-*********:vtb-********
 * ```
 * 
 */
@ResourceType(type="alicloud:cen/transitRouterRouteTable:TransitRouterRouteTable")
public class TransitRouterRouteTable extends com.pulumi.resources.CustomResource {
    /**
     * The dry run.
     * 
     */
    @Export(name="dryRun", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> dryRun;

    /**
     * @return The dry run.
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
     * A mapping of tags to assign to the resource.
     * 
     */
    @Export(name="tags", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output</* @Nullable */ Map<String,String>> tags;

    /**
     * @return A mapping of tags to assign to the resource.
     * 
     */
    public Output<Optional<Map<String,String>>> tags() {
        return Codegen.optional(this.tags);
    }
    /**
     * The ID of the transit router.
     * 
     */
    @Export(name="transitRouterId", refs={String.class}, tree="[0]")
    private Output<String> transitRouterId;

    /**
     * @return The ID of the transit router.
     * 
     */
    public Output<String> transitRouterId() {
        return this.transitRouterId;
    }
    /**
     * The description of the transit router route table.
     * 
     */
    @Export(name="transitRouterRouteTableDescription", refs={String.class}, tree="[0]")
    private Output<String> transitRouterRouteTableDescription;

    /**
     * @return The description of the transit router route table.
     * 
     */
    public Output<String> transitRouterRouteTableDescription() {
        return this.transitRouterRouteTableDescription;
    }
    /**
     * The id of the transit router route table.
     * 
     */
    @Export(name="transitRouterRouteTableId", refs={String.class}, tree="[0]")
    private Output<String> transitRouterRouteTableId;

    /**
     * @return The id of the transit router route table.
     * 
     */
    public Output<String> transitRouterRouteTableId() {
        return this.transitRouterRouteTableId;
    }
    /**
     * The name of the transit router route table.
     * 
     */
    @Export(name="transitRouterRouteTableName", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> transitRouterRouteTableName;

    /**
     * @return The name of the transit router route table.
     * 
     */
    public Output<Optional<String>> transitRouterRouteTableName() {
        return Codegen.optional(this.transitRouterRouteTableName);
    }
    /**
     * The type of the transit router route table. Valid values: `Custom`, `System`.
     * 
     */
    @Export(name="transitRouterRouteTableType", refs={String.class}, tree="[0]")
    private Output<String> transitRouterRouteTableType;

    /**
     * @return The type of the transit router route table. Valid values: `Custom`, `System`.
     * 
     */
    public Output<String> transitRouterRouteTableType() {
        return this.transitRouterRouteTableType;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public TransitRouterRouteTable(java.lang.String name) {
        this(name, TransitRouterRouteTableArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public TransitRouterRouteTable(java.lang.String name, TransitRouterRouteTableArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public TransitRouterRouteTable(java.lang.String name, TransitRouterRouteTableArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:cen/transitRouterRouteTable:TransitRouterRouteTable", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private TransitRouterRouteTable(java.lang.String name, Output<java.lang.String> id, @Nullable TransitRouterRouteTableState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:cen/transitRouterRouteTable:TransitRouterRouteTable", name, state, makeResourceOptions(options, id), false);
    }

    private static TransitRouterRouteTableArgs makeArgs(TransitRouterRouteTableArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? TransitRouterRouteTableArgs.Empty : args;
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
    public static TransitRouterRouteTable get(java.lang.String name, Output<java.lang.String> id, @Nullable TransitRouterRouteTableState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new TransitRouterRouteTable(name, id, state, options);
    }
}
