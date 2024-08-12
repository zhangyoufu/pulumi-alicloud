// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.cms;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.cms.MonitorGroupInstancesArgs;
import com.pulumi.alicloud.cms.inputs.MonitorGroupInstancesState;
import com.pulumi.alicloud.cms.outputs.MonitorGroupInstancesInstance;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.List;
import javax.annotation.Nullable;

/**
 * Provides a Cloud Monitor Service Monitor Group Instances resource.
 * 
 * For information about Cloud Monitor Service Monitor Group Instances and how to use it, see [What is Monitor Group Instances](https://www.alibabacloud.com/help/en/cloudmonitor/latest/createmonitorgroupinstances).
 * 
 * &gt; **NOTE:** Available since v1.115.0.
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
 * import com.pulumi.alicloud.cms.MonitorGroup;
 * import com.pulumi.alicloud.cms.MonitorGroupArgs;
 * import com.pulumi.alicloud.AlicloudFunctions;
 * import com.pulumi.alicloud.inputs.GetRegionsArgs;
 * import com.pulumi.alicloud.cms.MonitorGroupInstances;
 * import com.pulumi.alicloud.cms.MonitorGroupInstancesArgs;
 * import com.pulumi.alicloud.cms.inputs.MonitorGroupInstancesInstanceArgs;
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
 *         final var name = config.get("name").orElse("tf_example");
 *         var defaultNetwork = new Network("defaultNetwork", NetworkArgs.builder()
 *             .vpcName(name)
 *             .cidrBlock("192.168.0.0/16")
 *             .build());
 * 
 *         var defaultMonitorGroup = new MonitorGroup("defaultMonitorGroup", MonitorGroupArgs.builder()
 *             .monitorGroupName(name)
 *             .build());
 * 
 *         final var default = AlicloudFunctions.getRegions(GetRegionsArgs.builder()
 *             .current(true)
 *             .build());
 * 
 *         var example = new MonitorGroupInstances("example", MonitorGroupInstancesArgs.builder()
 *             .groupId(defaultMonitorGroup.id())
 *             .instances(MonitorGroupInstancesInstanceArgs.builder()
 *                 .instanceId(defaultNetwork.id())
 *                 .instanceName(name)
 *                 .regionId(default_.regions()[0].id())
 *                 .category("vpc")
 *                 .build())
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
 * Cloud Monitor Service Monitor Group Instances can be imported using the id, e.g.
 * 
 * ```sh
 * $ pulumi import alicloud:cms/monitorGroupInstances:MonitorGroupInstances example &lt;group_id&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:cms/monitorGroupInstances:MonitorGroupInstances")
public class MonitorGroupInstances extends com.pulumi.resources.CustomResource {
    /**
     * The id of Cms Group.
     * 
     */
    @Export(name="groupId", refs={String.class}, tree="[0]")
    private Output<String> groupId;

    /**
     * @return The id of Cms Group.
     * 
     */
    public Output<String> groupId() {
        return this.groupId;
    }
    /**
     * Instance information added to the Cms Group. See `instances` below.
     * 
     */
    @Export(name="instances", refs={List.class,MonitorGroupInstancesInstance.class}, tree="[0,1]")
    private Output<List<MonitorGroupInstancesInstance>> instances;

    /**
     * @return Instance information added to the Cms Group. See `instances` below.
     * 
     */
    public Output<List<MonitorGroupInstancesInstance>> instances() {
        return this.instances;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public MonitorGroupInstances(java.lang.String name) {
        this(name, MonitorGroupInstancesArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public MonitorGroupInstances(java.lang.String name, MonitorGroupInstancesArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public MonitorGroupInstances(java.lang.String name, MonitorGroupInstancesArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:cms/monitorGroupInstances:MonitorGroupInstances", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private MonitorGroupInstances(java.lang.String name, Output<java.lang.String> id, @Nullable MonitorGroupInstancesState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:cms/monitorGroupInstances:MonitorGroupInstances", name, state, makeResourceOptions(options, id), false);
    }

    private static MonitorGroupInstancesArgs makeArgs(MonitorGroupInstancesArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? MonitorGroupInstancesArgs.Empty : args;
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
    public static MonitorGroupInstances get(java.lang.String name, Output<java.lang.String> id, @Nullable MonitorGroupInstancesState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new MonitorGroupInstances(name, id, state, options);
    }
}
