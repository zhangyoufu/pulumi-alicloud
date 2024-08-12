// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.ebs;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.ebs.ReplicaGroupDrillArgs;
import com.pulumi.alicloud.ebs.inputs.ReplicaGroupDrillState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * Provides a EBS Replica Group Drill resource.
 * 
 * For information about EBS Replica Group Drill and how to use it, see [What is Replica Group Drill](https://www.alibabacloud.com/help/en/).
 * 
 * &gt; **NOTE:** Available since v1.215.0.
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
 * import com.pulumi.alicloud.ebs.ReplicaGroupDrill;
 * import com.pulumi.alicloud.ebs.ReplicaGroupDrillArgs;
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
 *         var default_ = new ReplicaGroupDrill("default", ReplicaGroupDrillArgs.builder()
 *             .groupId("pg-m1H9aaOUIGsDUwgZ")
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
 * EBS Replica Group Drill can be imported using the id, e.g.
 * 
 * ```sh
 * $ pulumi import alicloud:ebs/replicaGroupDrill:ReplicaGroupDrill example &lt;group_id&gt;:&lt;replica_group_drill_id&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:ebs/replicaGroupDrill:ReplicaGroupDrill")
public class ReplicaGroupDrill extends com.pulumi.resources.CustomResource {
    /**
     * The ID of the replication group. You can use the describediskreplicaggroups interface to query the asynchronous replication group list to obtain the value of the replication group ID input parameter.
     * 
     */
    @Export(name="groupId", refs={String.class}, tree="[0]")
    private Output<String> groupId;

    /**
     * @return The ID of the replication group. You can use the describediskreplicaggroups interface to query the asynchronous replication group list to obtain the value of the replication group ID input parameter.
     * 
     */
    public Output<String> groupId() {
        return this.groupId;
    }
    /**
     * The first ID of the resource.
     * 
     */
    @Export(name="replicaGroupDrillId", refs={String.class}, tree="[0]")
    private Output<String> replicaGroupDrillId;

    /**
     * @return The first ID of the resource.
     * 
     */
    public Output<String> replicaGroupDrillId() {
        return this.replicaGroupDrillId;
    }
    /**
     * Walkthrough status. _failed: Execution failed._failed: Cleanup failed.
     * 
     */
    @Export(name="status", refs={String.class}, tree="[0]")
    private Output<String> status;

    /**
     * @return Walkthrough status. _failed: Execution failed._failed: Cleanup failed.
     * 
     */
    public Output<String> status() {
        return this.status;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public ReplicaGroupDrill(java.lang.String name) {
        this(name, ReplicaGroupDrillArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ReplicaGroupDrill(java.lang.String name, ReplicaGroupDrillArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ReplicaGroupDrill(java.lang.String name, ReplicaGroupDrillArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:ebs/replicaGroupDrill:ReplicaGroupDrill", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private ReplicaGroupDrill(java.lang.String name, Output<java.lang.String> id, @Nullable ReplicaGroupDrillState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:ebs/replicaGroupDrill:ReplicaGroupDrill", name, state, makeResourceOptions(options, id), false);
    }

    private static ReplicaGroupDrillArgs makeArgs(ReplicaGroupDrillArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? ReplicaGroupDrillArgs.Empty : args;
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
    public static ReplicaGroupDrill get(java.lang.String name, Output<java.lang.String> id, @Nullable ReplicaGroupDrillState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ReplicaGroupDrill(name, id, state, options);
    }
}
