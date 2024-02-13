// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.ecs;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.ecs.AutoSnapshotPolicyArgs;
import com.pulumi.alicloud.ecs.inputs.AutoSnapshotPolicyState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.Object;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a ECS Auto Snapshot Policy resource.
 * 
 * For information about ECS Auto Snapshot Policy and how to use it, see [What is Auto Snapshot Policy](https://www.alibabacloud.com/help/en/doc-detail/25527.htm).
 * 
 * &gt; **NOTE:** Available in v1.117.0+.
 * 
 * ## Example Usage
 * 
 * ## Import
 * 
 * ECS Auto Snapshot Policy can be imported using the id, e.g.
 * 
 * ```sh
 * $ pulumi import alicloud:ecs/autoSnapshotPolicy:AutoSnapshotPolicy example &lt;id&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:ecs/autoSnapshotPolicy:AutoSnapshotPolicy")
public class AutoSnapshotPolicy extends com.pulumi.resources.CustomResource {
    /**
     * The retention period of the snapshot copied across regions.
     * - -1: The snapshot is permanently retained.
     * - [1, 65535]: The automatic snapshot is retained for the specified number of days.
     *   Default value: -1.
     * 
     */
    @Export(name="copiedSnapshotsRetentionDays", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> copiedSnapshotsRetentionDays;

    /**
     * @return The retention period of the snapshot copied across regions.
     * - -1: The snapshot is permanently retained.
     * - [1, 65535]: The automatic snapshot is retained for the specified number of days.
     *   Default value: -1.
     * 
     */
    public Output<Optional<Integer>> copiedSnapshotsRetentionDays() {
        return Codegen.optional(this.copiedSnapshotsRetentionDays);
    }
    /**
     * Specifies whether to enable the system to automatically copy snapshots across regions.
     * 
     */
    @Export(name="enableCrossRegionCopy", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> enableCrossRegionCopy;

    /**
     * @return Specifies whether to enable the system to automatically copy snapshots across regions.
     * 
     */
    public Output<Optional<Boolean>> enableCrossRegionCopy() {
        return Codegen.optional(this.enableCrossRegionCopy);
    }
    /**
     * The snapshot policy name.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The snapshot policy name.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The automatic snapshot repetition dates. The unit of measurement is day and the repeating cycle is a week. Value range: [1, 7], which represents days starting from Monday to Sunday, for example 1  indicates Monday. When you want to schedule multiple automatic snapshot tasks for a disk in a week, you can set the RepeatWeekdays to an array.
     * - A maximum of seven time points can be selected.
     * - The format is  an JSON array of [&#34;1&#34;, &#34;2&#34;, … &#34;7&#34;]  and the time points are separated by commas (,).
     * 
     */
    @Export(name="repeatWeekdays", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> repeatWeekdays;

    /**
     * @return The automatic snapshot repetition dates. The unit of measurement is day and the repeating cycle is a week. Value range: [1, 7], which represents days starting from Monday to Sunday, for example 1  indicates Monday. When you want to schedule multiple automatic snapshot tasks for a disk in a week, you can set the RepeatWeekdays to an array.
     * - A maximum of seven time points can be selected.
     * - The format is  an JSON array of [&#34;1&#34;, &#34;2&#34;, … &#34;7&#34;]  and the time points are separated by commas (,).
     * 
     */
    public Output<List<String>> repeatWeekdays() {
        return this.repeatWeekdays;
    }
    /**
     * The snapshot retention time, and the unit of measurement is day. Optional values:
     * - -1: The automatic snapshots are retained permanently.
     * - [1, 65536]: The number of days retained.
     *   Default value: -1.
     * 
     */
    @Export(name="retentionDays", refs={Integer.class}, tree="[0]")
    private Output<Integer> retentionDays;

    /**
     * @return The snapshot retention time, and the unit of measurement is day. Optional values:
     * - -1: The automatic snapshots are retained permanently.
     * - [1, 65536]: The number of days retained.
     *   Default value: -1.
     * 
     */
    public Output<Integer> retentionDays() {
        return this.retentionDays;
    }
    /**
     * The status of Auto Snapshot Policy.
     * 
     */
    @Export(name="status", refs={String.class}, tree="[0]")
    private Output<String> status;

    /**
     * @return The status of Auto Snapshot Policy.
     * 
     */
    public Output<String> status() {
        return this.status;
    }
    /**
     * A mapping of tags to assign to the resource.
     * 
     */
    @Export(name="tags", refs={Map.class,String.class,Object.class}, tree="[0,1,2]")
    private Output</* @Nullable */ Map<String,Object>> tags;

    /**
     * @return A mapping of tags to assign to the resource.
     * 
     */
    public Output<Optional<Map<String,Object>>> tags() {
        return Codegen.optional(this.tags);
    }
    /**
     * The destination region to which the snapshot is copied. You can set a destination region.
     * 
     */
    @Export(name="targetCopyRegions", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> targetCopyRegions;

    /**
     * @return The destination region to which the snapshot is copied. You can set a destination region.
     * 
     */
    public Output<Optional<List<String>>> targetCopyRegions() {
        return Codegen.optional(this.targetCopyRegions);
    }
    /**
     * The automatic snapshot creation schedule, and the unit of measurement is hour. Value range: [0, 23], which represents from 00:00 to 24:00,  for example 1 indicates 01:00. When you want to schedule multiple automatic snapshot tasks for a disk in a day, you can set the TimePoints to an array.
     * - A maximum of 24 time points can be selected.
     * - The format is  an JSON array of [&#34;0&#34;, &#34;1&#34;, … &#34;23&#34;] and the time points are separated by commas (,).
     * 
     */
    @Export(name="timePoints", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> timePoints;

    /**
     * @return The automatic snapshot creation schedule, and the unit of measurement is hour. Value range: [0, 23], which represents from 00:00 to 24:00,  for example 1 indicates 01:00. When you want to schedule multiple automatic snapshot tasks for a disk in a day, you can set the TimePoints to an array.
     * - A maximum of 24 time points can be selected.
     * - The format is  an JSON array of [&#34;0&#34;, &#34;1&#34;, … &#34;23&#34;] and the time points are separated by commas (,).
     * 
     */
    public Output<List<String>> timePoints() {
        return this.timePoints;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public AutoSnapshotPolicy(String name) {
        this(name, AutoSnapshotPolicyArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public AutoSnapshotPolicy(String name, AutoSnapshotPolicyArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public AutoSnapshotPolicy(String name, AutoSnapshotPolicyArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:ecs/autoSnapshotPolicy:AutoSnapshotPolicy", name, args == null ? AutoSnapshotPolicyArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private AutoSnapshotPolicy(String name, Output<String> id, @Nullable AutoSnapshotPolicyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:ecs/autoSnapshotPolicy:AutoSnapshotPolicy", name, state, makeResourceOptions(options, id));
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
    public static AutoSnapshotPolicy get(String name, Output<String> id, @Nullable AutoSnapshotPolicyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new AutoSnapshotPolicy(name, id, state, options);
    }
}
