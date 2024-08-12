// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.cdn;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.cdn.RealTimeLogDeliveryArgs;
import com.pulumi.alicloud.cdn.inputs.RealTimeLogDeliveryState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * Provides a CDN Real Time Log Delivery resource.
 * 
 * For information about CDN Real Time Log Delivery and how to use it, see [What is Real Time Log Delivery](https://www.alibabacloud.com/help/en/cdn/developer-reference/api-cdn-2018-05-10-createrealtimelogdelivery).
 * 
 * &gt; **NOTE:** Available since v1.134.0.
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
 * import com.pulumi.alicloud.cdn.DomainNew;
 * import com.pulumi.alicloud.cdn.DomainNewArgs;
 * import com.pulumi.alicloud.cdn.inputs.DomainNewSourceArgs;
 * import com.pulumi.alicloud.log.Project;
 * import com.pulumi.alicloud.log.ProjectArgs;
 * import com.pulumi.alicloud.log.Store;
 * import com.pulumi.alicloud.log.StoreArgs;
 * import com.pulumi.alicloud.AlicloudFunctions;
 * import com.pulumi.alicloud.inputs.GetRegionsArgs;
 * import com.pulumi.alicloud.cdn.RealTimeLogDelivery;
 * import com.pulumi.alicloud.cdn.RealTimeLogDeliveryArgs;
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
 *         var defaultInteger = new Integer("defaultInteger", IntegerArgs.builder()
 *             .max(99999)
 *             .min(10000)
 *             .build());
 * 
 *         var defaultDomainNew = new DomainNew("defaultDomainNew", DomainNewArgs.builder()
 *             .scope("overseas")
 *             .domainName(String.format("mycdndomain-%s.alicloud-provider.cn", defaultInteger.result()))
 *             .cdnType("web")
 *             .sources(DomainNewSourceArgs.builder()
 *                 .type("ipaddr")
 *                 .content("1.1.3.1")
 *                 .priority(20)
 *                 .port(80)
 *                 .weight(15)
 *                 .build())
 *             .build());
 * 
 *         var defaultProject = new Project("defaultProject", ProjectArgs.builder()
 *             .projectName(String.format("terraform-example-%s", defaultInteger.result()))
 *             .description("terraform-example")
 *             .build());
 * 
 *         var defaultStore = new Store("defaultStore", StoreArgs.builder()
 *             .projectName(defaultProject.name())
 *             .logstoreName("example-store")
 *             .shardCount(3)
 *             .autoSplit(true)
 *             .maxSplitShardCount(60)
 *             .appendMeta(true)
 *             .build());
 * 
 *         final var default = AlicloudFunctions.getRegions(GetRegionsArgs.builder()
 *             .current(true)
 *             .build());
 * 
 *         var defaultRealTimeLogDelivery = new RealTimeLogDelivery("defaultRealTimeLogDelivery", RealTimeLogDeliveryArgs.builder()
 *             .domain(defaultDomainNew.domainName())
 *             .logstore(defaultStore.logstoreName())
 *             .project(defaultProject.projectName())
 *             .slsRegion(default_.regions()[0].id())
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
 * CDN Real Time Log Delivery can be imported using the id, e.g.
 * 
 * ```sh
 * $ pulumi import alicloud:cdn/realTimeLogDelivery:RealTimeLogDelivery example &lt;domain&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:cdn/realTimeLogDelivery:RealTimeLogDelivery")
public class RealTimeLogDelivery extends com.pulumi.resources.CustomResource {
    /**
     * The accelerated domain name for which you want to configure real-time log delivery. You can specify multiple domain names and separate them with commas (,).
     * 
     */
    @Export(name="domain", refs={String.class}, tree="[0]")
    private Output<String> domain;

    /**
     * @return The accelerated domain name for which you want to configure real-time log delivery. You can specify multiple domain names and separate them with commas (,).
     * 
     */
    public Output<String> domain() {
        return this.domain;
    }
    /**
     * The name of the Logstore that collects log data from Alibaba Cloud Content Delivery Network (CDN) in real time.
     * 
     */
    @Export(name="logstore", refs={String.class}, tree="[0]")
    private Output<String> logstore;

    /**
     * @return The name of the Logstore that collects log data from Alibaba Cloud Content Delivery Network (CDN) in real time.
     * 
     */
    public Output<String> logstore() {
        return this.logstore;
    }
    /**
     * The name of the Log Service project that is used for real-time log delivery.
     * 
     */
    @Export(name="project", refs={String.class}, tree="[0]")
    private Output<String> project;

    /**
     * @return The name of the Log Service project that is used for real-time log delivery.
     * 
     */
    public Output<String> project() {
        return this.project;
    }
    /**
     * The region where the Log Service project is deployed.
     * 
     * &gt; **NOTE:** If your Project and Logstore services already exist, if you continue to create existing content, the created content will overwrite your existing indexes and custom reports. Please be careful to create your existing services to avoid affecting your online services after coverage.
     * 
     */
    @Export(name="slsRegion", refs={String.class}, tree="[0]")
    private Output<String> slsRegion;

    /**
     * @return The region where the Log Service project is deployed.
     * 
     * &gt; **NOTE:** If your Project and Logstore services already exist, if you continue to create existing content, the created content will overwrite your existing indexes and custom reports. Please be careful to create your existing services to avoid affecting your online services after coverage.
     * 
     */
    public Output<String> slsRegion() {
        return this.slsRegion;
    }
    /**
     * The status of the real-time log delivery feature. Valid Values: `online` and `offline`.
     * 
     */
    @Export(name="status", refs={String.class}, tree="[0]")
    private Output<String> status;

    /**
     * @return The status of the real-time log delivery feature. Valid Values: `online` and `offline`.
     * 
     */
    public Output<String> status() {
        return this.status;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public RealTimeLogDelivery(java.lang.String name) {
        this(name, RealTimeLogDeliveryArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public RealTimeLogDelivery(java.lang.String name, RealTimeLogDeliveryArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public RealTimeLogDelivery(java.lang.String name, RealTimeLogDeliveryArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:cdn/realTimeLogDelivery:RealTimeLogDelivery", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private RealTimeLogDelivery(java.lang.String name, Output<java.lang.String> id, @Nullable RealTimeLogDeliveryState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:cdn/realTimeLogDelivery:RealTimeLogDelivery", name, state, makeResourceOptions(options, id), false);
    }

    private static RealTimeLogDeliveryArgs makeArgs(RealTimeLogDeliveryArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? RealTimeLogDeliveryArgs.Empty : args;
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
    public static RealTimeLogDelivery get(java.lang.String name, Output<java.lang.String> id, @Nullable RealTimeLogDeliveryState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new RealTimeLogDelivery(name, id, state, options);
    }
}
