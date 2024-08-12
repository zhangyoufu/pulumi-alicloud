// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.ecs;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.ecs.EcsSnapshotArgs;
import com.pulumi.alicloud.ecs.inputs.EcsSnapshotState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.Object;
import java.lang.String;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a ECS Snapshot resource.
 * 
 * For information about ECS Snapshot and how to use it, see [What is Snapshot](https://www.alibabacloud.com/help/en/doc-detail/25524.htm).
 * 
 * &gt; **NOTE:** Available in v1.120.0+.
 * 
 * ## Import
 * 
 * ECS Snapshot can be imported using the id, e.g.
 * 
 * ```sh
 * $ pulumi import alicloud:ecs/ecsSnapshot:EcsSnapshot example &lt;id&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:ecs/ecsSnapshot:EcsSnapshot")
public class EcsSnapshot extends com.pulumi.resources.CustomResource {
    /**
     * The category of the snapshot. Valid Values: `standard` and `flash`.
     * 
     */
    @Export(name="category", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> category;

    /**
     * @return The category of the snapshot. Valid Values: `standard` and `flash`.
     * 
     */
    public Output<Optional<String>> category() {
        return Codegen.optional(this.category);
    }
    /**
     * The description of the snapshot.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return The description of the snapshot.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * The ID of the disk.
     * 
     */
    @Export(name="diskId", refs={String.class}, tree="[0]")
    private Output<String> diskId;

    /**
     * @return The ID of the disk.
     * 
     */
    public Output<String> diskId() {
        return this.diskId;
    }
    /**
     * Specifies whether to forcibly delete the snapshot that has been used to create disks.
     * 
     */
    @Export(name="force", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> force;

    /**
     * @return Specifies whether to forcibly delete the snapshot that has been used to create disks.
     * 
     */
    public Output<Optional<Boolean>> force() {
        return Codegen.optional(this.force);
    }
    /**
     * Specifies whether to enable the instant access feature.
     * 
     */
    @Export(name="instantAccess", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> instantAccess;

    /**
     * @return Specifies whether to enable the instant access feature.
     * 
     */
    public Output<Optional<Boolean>> instantAccess() {
        return Codegen.optional(this.instantAccess);
    }
    /**
     * Specifies the retention period of the instant access feature. After the retention period ends, the snapshot is automatically released.
     * 
     */
    @Export(name="instantAccessRetentionDays", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> instantAccessRetentionDays;

    /**
     * @return Specifies the retention period of the instant access feature. After the retention period ends, the snapshot is automatically released.
     * 
     */
    public Output<Optional<Integer>> instantAccessRetentionDays() {
        return Codegen.optional(this.instantAccessRetentionDays);
    }
    /**
     * Field `name` has been deprecated from provider version 1.120.0. New field `snapshot_name` instead.
     * 
     * @deprecated
     * Field &#39;name&#39; has been deprecated from provider version 1.120.0. New field &#39;snapshot_name&#39; instead.
     * 
     */
    @Deprecated /* Field 'name' has been deprecated from provider version 1.120.0. New field 'snapshot_name' instead. */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return Field `name` has been deprecated from provider version 1.120.0. New field `snapshot_name` instead.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The resource group id.
     * 
     */
    @Export(name="resourceGroupId", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> resourceGroupId;

    /**
     * @return The resource group id.
     * 
     */
    public Output<Optional<String>> resourceGroupId() {
        return Codegen.optional(this.resourceGroupId);
    }
    /**
     * The retention period of the snapshot.
     * 
     */
    @Export(name="retentionDays", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> retentionDays;

    /**
     * @return The retention period of the snapshot.
     * 
     */
    public Output<Optional<Integer>> retentionDays() {
        return Codegen.optional(this.retentionDays);
    }
    /**
     * The name of the snapshot.
     * 
     */
    @Export(name="snapshotName", refs={String.class}, tree="[0]")
    private Output<String> snapshotName;

    /**
     * @return The name of the snapshot.
     * 
     */
    public Output<String> snapshotName() {
        return this.snapshotName;
    }
    /**
     * The status of snapshot.
     * 
     */
    @Export(name="status", refs={String.class}, tree="[0]")
    private Output<String> status;

    /**
     * @return The status of snapshot.
     * 
     */
    public Output<String> status() {
        return this.status;
    }
    /**
     * A mapping of tags to assign to the snapshot.
     * 
     * &gt; **NOTE:** If `force` is true, After an snapshot is deleted, the disks created from this snapshot cannot be re-initialized.
     * 
     */
    @Export(name="tags", refs={Map.class,String.class,Object.class}, tree="[0,1,2]")
    private Output</* @Nullable */ Map<String,Object>> tags;

    /**
     * @return A mapping of tags to assign to the snapshot.
     * 
     * &gt; **NOTE:** If `force` is true, After an snapshot is deleted, the disks created from this snapshot cannot be re-initialized.
     * 
     */
    public Output<Optional<Map<String,Object>>> tags() {
        return Codegen.optional(this.tags);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public EcsSnapshot(java.lang.String name) {
        this(name, EcsSnapshotArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public EcsSnapshot(java.lang.String name, EcsSnapshotArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public EcsSnapshot(java.lang.String name, EcsSnapshotArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:ecs/ecsSnapshot:EcsSnapshot", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private EcsSnapshot(java.lang.String name, Output<java.lang.String> id, @Nullable EcsSnapshotState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:ecs/ecsSnapshot:EcsSnapshot", name, state, makeResourceOptions(options, id), false);
    }

    private static EcsSnapshotArgs makeArgs(EcsSnapshotArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? EcsSnapshotArgs.Empty : args;
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
    public static EcsSnapshot get(java.lang.String name, Output<java.lang.String> id, @Nullable EcsSnapshotState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new EcsSnapshot(name, id, state, options);
    }
}
