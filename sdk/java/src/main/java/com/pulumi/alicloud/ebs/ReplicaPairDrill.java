// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.ebs;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.ebs.ReplicaPairDrillArgs;
import com.pulumi.alicloud.ebs.inputs.ReplicaPairDrillState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * Provides a EBS Replica Pair Drill resource.
 * 
 * For information about EBS Replica Pair Drill and how to use it, see [What is Replica Pair Drill](https://www.alibabacloud.com/help/en/).
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
 * import com.pulumi.alicloud.ebs.ReplicaPairDrill;
 * import com.pulumi.alicloud.ebs.ReplicaPairDrillArgs;
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
 *         var default_ = new ReplicaPairDrill("default", ReplicaPairDrillArgs.builder()
 *             .pairId("pair-cn-wwo3kjfq5001")
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
 * EBS Replica Pair Drill can be imported using the id, e.g.
 * 
 * ```sh
 * $ pulumi import alicloud:ebs/replicaPairDrill:ReplicaPairDrill example &lt;pair_id&gt;:&lt;replica_pair_drill_id&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:ebs/replicaPairDrill:ReplicaPairDrill")
public class ReplicaPairDrill extends com.pulumi.resources.CustomResource {
    /**
     * Copy the ID of the pair. You can call DescribeDiskReplicaPairs to query the list of asynchronous replication pairs to obtain the replication pair ID.
     * 
     */
    @Export(name="pairId", refs={String.class}, tree="[0]")
    private Output<String> pairId;

    /**
     * @return Copy the ID of the pair. You can call DescribeDiskReplicaPairs to query the list of asynchronous replication pairs to obtain the replication pair ID.
     * 
     */
    public Output<String> pairId() {
        return this.pairId;
    }
    /**
     * The first ID of the resource.
     * 
     */
    @Export(name="replicaPairDrillId", refs={String.class}, tree="[0]")
    private Output<String> replicaPairDrillId;

    /**
     * @return The first ID of the resource.
     * 
     */
    public Output<String> replicaPairDrillId() {
        return this.replicaPairDrillId;
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
    public ReplicaPairDrill(java.lang.String name) {
        this(name, ReplicaPairDrillArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ReplicaPairDrill(java.lang.String name, ReplicaPairDrillArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ReplicaPairDrill(java.lang.String name, ReplicaPairDrillArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:ebs/replicaPairDrill:ReplicaPairDrill", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private ReplicaPairDrill(java.lang.String name, Output<java.lang.String> id, @Nullable ReplicaPairDrillState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:ebs/replicaPairDrill:ReplicaPairDrill", name, state, makeResourceOptions(options, id), false);
    }

    private static ReplicaPairDrillArgs makeArgs(ReplicaPairDrillArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? ReplicaPairDrillArgs.Empty : args;
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
    public static ReplicaPairDrill get(java.lang.String name, Output<java.lang.String> id, @Nullable ReplicaPairDrillState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ReplicaPairDrill(name, id, state, options);
    }
}
