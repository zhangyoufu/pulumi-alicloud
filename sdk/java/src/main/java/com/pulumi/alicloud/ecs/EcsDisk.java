// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.ecs;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.ecs.EcsDiskArgs;
import com.pulumi.alicloud.ecs.inputs.EcsDiskState;
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
 * ## Import
 * 
 * ECS Disk can be imported using the id, e.g.
 * 
 * ```sh
 * $ pulumi import alicloud:ecs/ecsDisk:EcsDisk example d-abcd12345
 * ```
 * 
 */
@ResourceType(type="alicloud:ecs/ecsDisk:EcsDisk")
public class EcsDisk extends com.pulumi.resources.CustomResource {
    @Export(name="advancedFeatures", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> advancedFeatures;

    public Output<Optional<String>> advancedFeatures() {
        return Codegen.optional(this.advancedFeatures);
    }
    /**
     * Field `availability_zone` has been deprecated from provider version 1.122.0. New field `zone_id` instead.
     * 
     * @deprecated
     * Field &#39;availability_zone&#39; has been deprecated from provider version 1.122.0. New field &#39;zone_id&#39; instead
     * 
     */
    @Deprecated /* Field 'availability_zone' has been deprecated from provider version 1.122.0. New field 'zone_id' instead */
    @Export(name="availabilityZone", refs={String.class}, tree="[0]")
    private Output<String> availabilityZone;

    /**
     * @return Field `availability_zone` has been deprecated from provider version 1.122.0. New field `zone_id` instead.
     * 
     */
    public Output<String> availabilityZone() {
        return this.availabilityZone;
    }
    /**
     * Category of the disk. Valid values are `cloud`, `cloud_efficiency`, `cloud_ssd`, `cloud_essd`, `cloud_auto`, `cloud_essd_entry`, `elastic_ephemeral_disk_standard`, `elastic_ephemeral_disk_premium`. Default is `cloud_efficiency`.
     * 
     */
    @Export(name="category", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> category;

    /**
     * @return Category of the disk. Valid values are `cloud`, `cloud_efficiency`, `cloud_ssd`, `cloud_essd`, `cloud_auto`, `cloud_essd_entry`, `elastic_ephemeral_disk_standard`, `elastic_ephemeral_disk_premium`. Default is `cloud_efficiency`.
     * 
     */
    public Output<Optional<String>> category() {
        return Codegen.optional(this.category);
    }
    /**
     * Indicates whether the automatic snapshot is deleted when the disk is released. Default value: `false`.
     * 
     */
    @Export(name="deleteAutoSnapshot", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> deleteAutoSnapshot;

    /**
     * @return Indicates whether the automatic snapshot is deleted when the disk is released. Default value: `false`.
     * 
     */
    public Output<Optional<Boolean>> deleteAutoSnapshot() {
        return Codegen.optional(this.deleteAutoSnapshot);
    }
    /**
     * Indicates whether the disk is released together with the instance. Default value: `false`.
     * 
     */
    @Export(name="deleteWithInstance", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> deleteWithInstance;

    /**
     * @return Indicates whether the disk is released together with the instance. Default value: `false`.
     * 
     */
    public Output<Boolean> deleteWithInstance() {
        return this.deleteWithInstance;
    }
    /**
     * Description of the disk. This description can have a string of 2 to 256 characters, It cannot begin with http:// or https://. Default value is null.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return Description of the disk. This description can have a string of 2 to 256 characters, It cannot begin with http:// or https://. Default value is null.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * Name of the ECS disk. This name can have a string of 2 to 128 characters, must contain only alphanumeric characters or hyphens, such as &#34;-&#34;,&#34;.&#34;,&#34;_&#34;, and must not begin or end with a hyphen, and must not begin with `http://` or `https://`. Default value is `null`.
     * 
     */
    @Export(name="diskName", refs={String.class}, tree="[0]")
    private Output<String> diskName;

    /**
     * @return Name of the ECS disk. This name can have a string of 2 to 128 characters, must contain only alphanumeric characters or hyphens, such as &#34;-&#34;,&#34;.&#34;,&#34;_&#34;, and must not begin or end with a hyphen, and must not begin with `http://` or `https://`. Default value is `null`.
     * 
     */
    public Output<String> diskName() {
        return this.diskName;
    }
    /**
     * Specifies whether to check the validity of the request without actually making the request.request Default value: false. Valid values:
     * 
     */
    @Export(name="dryRun", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> dryRun;

    /**
     * @return Specifies whether to check the validity of the request without actually making the request.request Default value: false. Valid values:
     * 
     */
    public Output<Optional<Boolean>> dryRun() {
        return Codegen.optional(this.dryRun);
    }
    /**
     * Indicates whether to enable creating snapshot automatically.
     * 
     */
    @Export(name="enableAutoSnapshot", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> enableAutoSnapshot;

    /**
     * @return Indicates whether to enable creating snapshot automatically.
     * 
     */
    public Output<Boolean> enableAutoSnapshot() {
        return this.enableAutoSnapshot;
    }
    @Export(name="encryptAlgorithm", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> encryptAlgorithm;

    public Output<Optional<String>> encryptAlgorithm() {
        return Codegen.optional(this.encryptAlgorithm);
    }
    /**
     * If true, the disk will be encrypted, conflict with `snapshot_id`.
     * 
     */
    @Export(name="encrypted", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> encrypted;

    /**
     * @return If true, the disk will be encrypted, conflict with `snapshot_id`.
     * 
     */
    public Output<Optional<Boolean>> encrypted() {
        return Codegen.optional(this.encrypted);
    }
    /**
     * The ID of the instance to which the created subscription disk is automatically attached.
     * * After you specify the instance ID, the specified `resource_group_id`, `tags`, and `kms_key_id` parameters are ignored.
     * * One of the `zone_id` and `instance_id` must be set but can not be set at the same time.
     * 
     */
    @Export(name="instanceId", refs={String.class}, tree="[0]")
    private Output<String> instanceId;

    /**
     * @return The ID of the instance to which the created subscription disk is automatically attached.
     * * After you specify the instance ID, the specified `resource_group_id`, `tags`, and `kms_key_id` parameters are ignored.
     * * One of the `zone_id` and `instance_id` must be set but can not be set at the same time.
     * 
     */
    public Output<String> instanceId() {
        return this.instanceId;
    }
    /**
     * The ID of the KMS key corresponding to the data disk, The specified parameter `Encrypted` must be `true` when KmsKeyId is not empty.
     * 
     */
    @Export(name="kmsKeyId", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> kmsKeyId;

    /**
     * @return The ID of the KMS key corresponding to the data disk, The specified parameter `Encrypted` must be `true` when KmsKeyId is not empty.
     * 
     */
    public Output<Optional<String>> kmsKeyId() {
        return Codegen.optional(this.kmsKeyId);
    }
    /**
     * Field `name` has been deprecated from provider version 1.122.0. New field `disk_name` instead.
     * 
     * @deprecated
     * Field &#39;name&#39; has been deprecated from provider version 1.122.0. New field &#39;disk_name&#39; instead.
     * 
     */
    @Deprecated /* Field 'name' has been deprecated from provider version 1.122.0. New field 'disk_name' instead. */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return Field `name` has been deprecated from provider version 1.122.0. New field `disk_name` instead.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * Payment method for disk. Valid values: `PayAsYouGo`, `Subscription`. Default to `PayAsYouGo`. If you want to change the disk payment type, the `instance_id` is required.
     * 
     */
    @Export(name="paymentType", refs={String.class}, tree="[0]")
    private Output<String> paymentType;

    /**
     * @return Payment method for disk. Valid values: `PayAsYouGo`, `Subscription`. Default to `PayAsYouGo`. If you want to change the disk payment type, the `instance_id` is required.
     * 
     */
    public Output<String> paymentType() {
        return this.paymentType;
    }
    /**
     * Specifies the performance level of an ESSD when you create the ESSD. Valid values:
     * * `PL0`: A single ESSD delivers up to 10,000 random read/write IOPS.
     * * `PL1`: A single ESSD delivers up to 50,000 random read/write IOPS.
     * * `PL2`: A single ESSD delivers up to 100,000 random read/write IOPS.
     * * `PL3`: A single ESSD delivers up to 1,000,000 random read/write IOPS.
     * 
     */
    @Export(name="performanceLevel", refs={String.class}, tree="[0]")
    private Output<String> performanceLevel;

    /**
     * @return Specifies the performance level of an ESSD when you create the ESSD. Valid values:
     * * `PL0`: A single ESSD delivers up to 10,000 random read/write IOPS.
     * * `PL1`: A single ESSD delivers up to 50,000 random read/write IOPS.
     * * `PL2`: A single ESSD delivers up to 100,000 random read/write IOPS.
     * * `PL3`: A single ESSD delivers up to 1,000,000 random read/write IOPS.
     * 
     */
    public Output<String> performanceLevel() {
        return this.performanceLevel;
    }
    /**
     * The Id of resource group which the disk belongs. This attribute only supports adding or updating, not destroying.
     * 
     */
    @Export(name="resourceGroupId", refs={String.class}, tree="[0]")
    private Output<String> resourceGroupId;

    /**
     * @return The Id of resource group which the disk belongs. This attribute only supports adding or updating, not destroying.
     * 
     */
    public Output<String> resourceGroupId() {
        return this.resourceGroupId;
    }
    /**
     * The size of the disk in GiBs. When resize the disk, the new size must be greater than the former value, or you would get an error `InvalidDiskSize.TooSmall`.
     * 
     */
    @Export(name="size", refs={Integer.class}, tree="[0]")
    private Output<Integer> size;

    /**
     * @return The size of the disk in GiBs. When resize the disk, the new size must be greater than the former value, or you would get an error `InvalidDiskSize.TooSmall`.
     * 
     */
    public Output<Integer> size() {
        return this.size;
    }
    /**
     * A snapshot to base the disk off of. If the disk size required by snapshot is greater than `size`, the `size` will be ignored, conflict with `encrypted`.
     * 
     */
    @Export(name="snapshotId", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> snapshotId;

    /**
     * @return A snapshot to base the disk off of. If the disk size required by snapshot is greater than `size`, the `size` will be ignored, conflict with `encrypted`.
     * 
     */
    public Output<Optional<String>> snapshotId() {
        return Codegen.optional(this.snapshotId);
    }
    /**
     * The disk status.
     * 
     */
    @Export(name="status", refs={String.class}, tree="[0]")
    private Output<String> status;

    /**
     * @return The disk status.
     * 
     */
    public Output<String> status() {
        return this.status;
    }
    /**
     * The ID of the storage set.
     * 
     */
    @Export(name="storageSetId", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> storageSetId;

    /**
     * @return The ID of the storage set.
     * 
     */
    public Output<Optional<String>> storageSetId() {
        return Codegen.optional(this.storageSetId);
    }
    /**
     * The number of partitions in the storage set.
     * 
     */
    @Export(name="storageSetPartitionNumber", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> storageSetPartitionNumber;

    /**
     * @return The number of partitions in the storage set.
     * 
     */
    public Output<Optional<Integer>> storageSetPartitionNumber() {
        return Codegen.optional(this.storageSetPartitionNumber);
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
     * The type to expand cloud disks. Valid Values: `online`, `offline`. Default to `offline`.
     * 
     */
    @Export(name="type", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> type;

    /**
     * @return The type to expand cloud disks. Valid Values: `online`, `offline`. Default to `offline`.
     * 
     */
    public Output<Optional<String>> type() {
        return Codegen.optional(this.type);
    }
    /**
     * ID of the free zone to which the disk belongs. One of the `zone_id` and `instance_id` must be set but can not be set at the same time.
     * 
     */
    @Export(name="zoneId", refs={String.class}, tree="[0]")
    private Output<String> zoneId;

    /**
     * @return ID of the free zone to which the disk belongs. One of the `zone_id` and `instance_id` must be set but can not be set at the same time.
     * 
     */
    public Output<String> zoneId() {
        return this.zoneId;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public EcsDisk(String name) {
        this(name, EcsDiskArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public EcsDisk(String name, @Nullable EcsDiskArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public EcsDisk(String name, @Nullable EcsDiskArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:ecs/ecsDisk:EcsDisk", name, args == null ? EcsDiskArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private EcsDisk(String name, Output<String> id, @Nullable EcsDiskState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:ecs/ecsDisk:EcsDisk", name, state, makeResourceOptions(options, id));
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
    public static EcsDisk get(String name, Output<String> id, @Nullable EcsDiskState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new EcsDisk(name, id, state, options);
    }
}
