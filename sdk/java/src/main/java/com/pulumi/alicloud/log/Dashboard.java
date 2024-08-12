// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.log;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.log.DashboardArgs;
import com.pulumi.alicloud.log.inputs.DashboardState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * The dashboard is a real-time data analysis platform provided by the log service. You can display frequently used query and analysis statements in the form of charts and save statistical charts to the dashboard.
 * [Refer to details](https://www.alibabacloud.com/help/doc-detail/102530.htm).
 * 
 * &gt; **NOTE:** Available since v1.86.0.
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
 * import com.pulumi.random.integer;
 * import com.pulumi.random.IntegerArgs;
 * import com.pulumi.alicloud.log.Project;
 * import com.pulumi.alicloud.log.ProjectArgs;
 * import com.pulumi.alicloud.log.Store;
 * import com.pulumi.alicloud.log.StoreArgs;
 * import com.pulumi.alicloud.log.Dashboard;
 * import com.pulumi.alicloud.log.DashboardArgs;
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
 *         var default_ = new Integer("default", IntegerArgs.builder()
 *             .max(99999)
 *             .min(10000)
 *             .build());
 * 
 *         var example = new Project("example", ProjectArgs.builder()
 *             .name(String.format("terraform-example-%s", default_.result()))
 *             .description("terraform-example")
 *             .build());
 * 
 *         var exampleStore = new Store("exampleStore", StoreArgs.builder()
 *             .project(example.name())
 *             .name("example-store")
 *             .shardCount(3)
 *             .autoSplit(true)
 *             .maxSplitShardCount(60)
 *             .appendMeta(true)
 *             .build());
 * 
 *         var exampleDashboard = new Dashboard("exampleDashboard", DashboardArgs.builder()
 *             .projectName(example.name())
 *             .dashboardName("terraform-example")
 *             .displayName("terraform-example")
 *             .attribute("""
 *   {
 *     "type":"grid"
 *   }
 *             """)
 *             .charList("""
 *   [
 *     {
 *       "action": {},
 *       "title":"new_title",
 *       "type":"map",
 *       "search":{
 *         "logstore":"example-store",
 *         "topic":"new_topic",
 *         "query":"* | SELECT COUNT(name) as ct_name, COUNT(product) as ct_product, name,product GROUP BY name,product",
 *         "start":"-86400s",
 *         "end":"now"
 *       },
 *       "display":{
 *         "xAxis":[
 *           "ct_name"
 *         ],
 *         "yAxis":[
 *           "ct_product"
 *         ],
 *         "xPos":0,
 *         "yPos":0,
 *         "width":10,
 *         "height":12,
 *         "displayName":"terraform-example"
 *       }
 *     }
 *   ]
 *             """)
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
 * Log Dashboard can be imported using the id, e.g.
 * 
 * ```sh
 * $ pulumi import alicloud:log/dashboard:Dashboard example &lt;project_name&gt;:&lt;dashboard_name&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:log/dashboard:Dashboard")
public class Dashboard extends com.pulumi.resources.CustomResource {
    /**
     * Dashboard attribute.
     * 
     */
    @Export(name="attribute", refs={String.class}, tree="[0]")
    private Output<String> attribute;

    /**
     * @return Dashboard attribute.
     * 
     */
    public Output<String> attribute() {
        return this.attribute;
    }
    /**
     * Configuration of charts in the dashboard.
     * **Note:** From version 1.164.0, `char_list` can set parameter &#34;action&#34;.
     * 
     */
    @Export(name="charList", refs={String.class}, tree="[0]")
    private Output<String> charList;

    /**
     * @return Configuration of charts in the dashboard.
     * **Note:** From version 1.164.0, `char_list` can set parameter &#34;action&#34;.
     * 
     */
    public Output<String> charList() {
        return this.charList;
    }
    /**
     * The name of the Log Dashboard.
     * 
     */
    @Export(name="dashboardName", refs={String.class}, tree="[0]")
    private Output<String> dashboardName;

    /**
     * @return The name of the Log Dashboard.
     * 
     */
    public Output<String> dashboardName() {
        return this.dashboardName;
    }
    /**
     * Dashboard alias.
     * 
     */
    @Export(name="displayName", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> displayName;

    /**
     * @return Dashboard alias.
     * 
     */
    public Output<Optional<String>> displayName() {
        return Codegen.optional(this.displayName);
    }
    /**
     * The name of the log project. It is the only in one Alicloud account.
     * 
     */
    @Export(name="projectName", refs={String.class}, tree="[0]")
    private Output<String> projectName;

    /**
     * @return The name of the log project. It is the only in one Alicloud account.
     * 
     */
    public Output<String> projectName() {
        return this.projectName;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Dashboard(java.lang.String name) {
        this(name, DashboardArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Dashboard(java.lang.String name, DashboardArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Dashboard(java.lang.String name, DashboardArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:log/dashboard:Dashboard", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private Dashboard(java.lang.String name, Output<java.lang.String> id, @Nullable DashboardState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:log/dashboard:Dashboard", name, state, makeResourceOptions(options, id), false);
    }

    private static DashboardArgs makeArgs(DashboardArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? DashboardArgs.Empty : args;
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
    public static Dashboard get(java.lang.String name, Output<java.lang.String> id, @Nullable DashboardState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Dashboard(name, id, state, options);
    }
}
