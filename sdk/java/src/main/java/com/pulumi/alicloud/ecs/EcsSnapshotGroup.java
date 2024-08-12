// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.ecs;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.ecs.EcsSnapshotGroupArgs;
import com.pulumi.alicloud.ecs.inputs.EcsSnapshotGroupState;
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
 * Provides a ECS Snapshot Group resource.
 * 
 * For information about ECS Snapshot Group and how to use it, see [What is Snapshot Group](https://www.alibabacloud.com/help/en/doc-detail/210939.html).
 * 
 * &gt; **NOTE:** Available in v1.160.0+.
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
 * import com.pulumi.alicloud.AlicloudFunctions;
 * import com.pulumi.alicloud.inputs.GetZonesArgs;
 * import com.pulumi.alicloud.ecs.EcsFunctions;
 * import com.pulumi.alicloud.ecs.inputs.GetInstanceTypesArgs;
 * import com.pulumi.alicloud.ecs.inputs.GetImagesArgs;
 * import com.pulumi.alicloud.vpc.Network;
 * import com.pulumi.alicloud.vpc.NetworkArgs;
 * import com.pulumi.alicloud.vpc.Switch;
 * import com.pulumi.alicloud.vpc.SwitchArgs;
 * import com.pulumi.alicloud.ecs.SecurityGroup;
 * import com.pulumi.alicloud.ecs.SecurityGroupArgs;
 * import com.pulumi.alicloud.ecs.Instance;
 * import com.pulumi.alicloud.ecs.InstanceArgs;
 * import com.pulumi.alicloud.ecs.EcsDisk;
 * import com.pulumi.alicloud.ecs.EcsDiskArgs;
 * import com.pulumi.alicloud.ecs.DiskAttachment;
 * import com.pulumi.alicloud.ecs.DiskAttachmentArgs;
 * import com.pulumi.alicloud.ecs.EcsSnapshotGroup;
 * import com.pulumi.alicloud.ecs.EcsSnapshotGroupArgs;
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
 *         final var default = AlicloudFunctions.getZones(GetZonesArgs.builder()
 *             .availableResourceCreation("Instance")
 *             .availableDiskCategory("cloud_essd")
 *             .build());
 * 
 *         final var defaultGetInstanceTypes = EcsFunctions.getInstanceTypes(GetInstanceTypesArgs.builder()
 *             .availabilityZone(default_.zones()[0].id())
 *             .systemDiskCategory("cloud_essd")
 *             .build());
 * 
 *         final var defaultGetImages = EcsFunctions.getImages(GetImagesArgs.builder()
 *             .owners("system")
 *             .nameRegex("^ubuntu_18.*64")
 *             .mostRecent(true)
 *             .build());
 * 
 *         var defaultNetwork = new Network("defaultNetwork", NetworkArgs.builder()
 *             .vpcName("terraform-example")
 *             .cidrBlock("172.17.3.0/24")
 *             .build());
 * 
 *         var defaultSwitch = new Switch("defaultSwitch", SwitchArgs.builder()
 *             .vswitchName("terraform-example")
 *             .cidrBlock("172.17.3.0/24")
 *             .vpcId(defaultNetwork.id())
 *             .zoneId(default_.zones()[0].id())
 *             .build());
 * 
 *         var defaultSecurityGroup = new SecurityGroup("defaultSecurityGroup", SecurityGroupArgs.builder()
 *             .name("terraform-example")
 *             .vpcId(defaultNetwork.id())
 *             .build());
 * 
 *         var defaultInstance = new Instance("defaultInstance", InstanceArgs.builder()
 *             .availabilityZone(default_.zones()[0].id())
 *             .instanceName("terraform-example")
 *             .securityGroups(defaultSecurityGroup.id())
 *             .vswitchId(defaultSwitch.id())
 *             .instanceType(defaultGetInstanceTypes.applyValue(getInstanceTypesResult -> getInstanceTypesResult.instanceTypes()[0].id()))
 *             .imageId(defaultGetImages.applyValue(getImagesResult -> getImagesResult.images()[0].id()))
 *             .internetMaxBandwidthOut(10)
 *             .build());
 * 
 *         var defaultEcsDisk = new EcsDisk("defaultEcsDisk", EcsDiskArgs.builder()
 *             .zoneId(default_.zones()[0].id())
 *             .diskName("terraform-example")
 *             .description("terraform-example")
 *             .category("cloud_essd")
 *             .size("30")
 *             .build());
 * 
 *         var defaultDiskAttachment = new DiskAttachment("defaultDiskAttachment", DiskAttachmentArgs.builder()
 *             .diskId(defaultEcsDisk.id())
 *             .instanceId(defaultInstance.id())
 *             .build());
 * 
 *         var defaultEcsSnapshotGroup = new EcsSnapshotGroup("defaultEcsSnapshotGroup", EcsSnapshotGroupArgs.builder()
 *             .description("terraform-example")
 *             .diskIds(defaultDiskAttachment.diskId())
 *             .snapshotGroupName("terraform-example")
 *             .instanceId(defaultInstance.id())
 *             .instantAccess(true)
 *             .instantAccessRetentionDays(1)
 *             .tags(Map.ofEntries(
 *                 Map.entry("Created", "TF"),
 *                 Map.entry("For", "Acceptance")
 *             ))
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
 * ECS Snapshot Group can be imported using the id, e.g.
 * 
 * ```sh
 * $ pulumi import alicloud:ecs/ecsSnapshotGroup:EcsSnapshotGroup example &lt;id&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:ecs/ecsSnapshotGroup:EcsSnapshotGroup")
public class EcsSnapshotGroup extends com.pulumi.resources.CustomResource {
    /**
     * The description of the snapshot-consistent group. The description must be 2 to 256 characters in length and cannot start with http:// or https://.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return The description of the snapshot-consistent group. The description must be 2 to 256 characters in length and cannot start with http:// or https://.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * The ID of disk for which to create snapshots. You can specify multiple disk IDs across instances with the same zone.
     * 
     */
    @Export(name="diskIds", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> diskIds;

    /**
     * @return The ID of disk for which to create snapshots. You can specify multiple disk IDs across instances with the same zone.
     * 
     */
    public Output<Optional<List<String>>> diskIds() {
        return Codegen.optional(this.diskIds);
    }
    /**
     * The ID of disk N for which you do not need to create snapshots. After this parameter is specified, the created snapshot-consistent group does not contain snapshots of the disk.
     * 
     */
    @Export(name="excludeDiskIds", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> excludeDiskIds;

    /**
     * @return The ID of disk N for which you do not need to create snapshots. After this parameter is specified, the created snapshot-consistent group does not contain snapshots of the disk.
     * 
     */
    public Output<Optional<List<String>>> excludeDiskIds() {
        return Codegen.optional(this.excludeDiskIds);
    }
    /**
     * The ID of the instance.
     * 
     */
    @Export(name="instanceId", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> instanceId;

    /**
     * @return The ID of the instance.
     * 
     */
    public Output<Optional<String>> instanceId() {
        return Codegen.optional(this.instanceId);
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
     * Specify the number of days for which the instant access feature is available. Unit: days. Valid values: `1` to `65535`.
     * 
     */
    @Export(name="instantAccessRetentionDays", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> instantAccessRetentionDays;

    /**
     * @return Specify the number of days for which the instant access feature is available. Unit: days. Valid values: `1` to `65535`.
     * 
     */
    public Output<Optional<Integer>> instantAccessRetentionDays() {
        return Codegen.optional(this.instantAccessRetentionDays);
    }
    /**
     * The ID of the resource group to which the snapshot consistency group belongs.
     * 
     */
    @Export(name="resourceGroupId", refs={String.class}, tree="[0]")
    private Output<String> resourceGroupId;

    /**
     * @return The ID of the resource group to which the snapshot consistency group belongs.
     * 
     */
    public Output<String> resourceGroupId() {
        return this.resourceGroupId;
    }
    /**
     * The name of the snapshot-consistent group. The name must be `2` to `128` characters in length, and can contain letters, digits, periods (.), underscores (_), hyphens (-), and colons (:). It must start with a letter or a digit and cannot start with `http://` or `https://`.
     * 
     */
    @Export(name="snapshotGroupName", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> snapshotGroupName;

    /**
     * @return The name of the snapshot-consistent group. The name must be `2` to `128` characters in length, and can contain letters, digits, periods (.), underscores (_), hyphens (-), and colons (:). It must start with a letter or a digit and cannot start with `http://` or `https://`.
     * 
     */
    public Output<Optional<String>> snapshotGroupName() {
        return Codegen.optional(this.snapshotGroupName);
    }
    /**
     * The status of the resource.
     * 
     */
    @Export(name="status", refs={String.class}, tree="[0]")
    private Output<String> status;

    /**
     * @return The status of the resource.
     * 
     */
    public Output<String> status() {
        return this.status;
    }
    /**
     * A mapping of tags to assign to the snapshot group.
     * 
     */
    @Export(name="tags", refs={Map.class,String.class,Object.class}, tree="[0,1,2]")
    private Output</* @Nullable */ Map<String,Object>> tags;

    /**
     * @return A mapping of tags to assign to the snapshot group.
     * 
     */
    public Output<Optional<Map<String,Object>>> tags() {
        return Codegen.optional(this.tags);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public EcsSnapshotGroup(java.lang.String name) {
        this(name, EcsSnapshotGroupArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public EcsSnapshotGroup(java.lang.String name, @Nullable EcsSnapshotGroupArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public EcsSnapshotGroup(java.lang.String name, @Nullable EcsSnapshotGroupArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:ecs/ecsSnapshotGroup:EcsSnapshotGroup", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private EcsSnapshotGroup(java.lang.String name, Output<java.lang.String> id, @Nullable EcsSnapshotGroupState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:ecs/ecsSnapshotGroup:EcsSnapshotGroup", name, state, makeResourceOptions(options, id), false);
    }

    private static EcsSnapshotGroupArgs makeArgs(@Nullable EcsSnapshotGroupArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? EcsSnapshotGroupArgs.Empty : args;
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
    public static EcsSnapshotGroup get(java.lang.String name, Output<java.lang.String> id, @Nullable EcsSnapshotGroupState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new EcsSnapshotGroup(name, id, state, options);
    }
}
