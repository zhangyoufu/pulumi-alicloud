// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.ens;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.ens.DiskArgs;
import com.pulumi.alicloud.ens.inputs.DiskState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a ENS Disk resource. The disk. When you use it for the first time, please contact the product classmates to add a resource whitelist.
 * 
 * For information about ENS Disk and how to use it, see [What is Disk](https://www.alibabacloud.com/help/en/ens/developer-reference/api-createdisk).
 * 
 * &gt; **NOTE:** Available since v1.213.0.
 * 
 * ## Example Usage
 * 
 * Basic Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.alicloud.ens.Disk;
 * import com.pulumi.alicloud.ens.DiskArgs;
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
 *         final var name = config.get(&#34;name&#34;).orElse(&#34;terraform-example&#34;);
 *         var default_ = new Disk(&#34;default&#34;, DiskArgs.builder()        
 *             .category(&#34;cloud_ssd&#34;)
 *             .ensRegionId(&#34;cn-chongqing-11&#34;)
 *             .paymentType(&#34;PayAsYouGo&#34;)
 *             .size(&#34;20&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * ENS Disk can be imported using the id, e.g.
 * 
 * ```sh
 * $ pulumi import alicloud:ens/disk:Disk example &lt;id&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:ens/disk:Disk")
public class Disk extends com.pulumi.resources.CustomResource {
    /**
     * Types of disk instancesValues: cloud_efficiency (high-efficiency cloud disk),cloud_ssd (full Flash cloud disk),local_hdd (local HDD),local_ssd (local ssd).
     * 
     */
    @Export(name="category", refs={String.class}, tree="[0]")
    private Output<String> category;

    /**
     * @return Types of disk instancesValues: cloud_efficiency (high-efficiency cloud disk),cloud_ssd (full Flash cloud disk),local_hdd (local HDD),local_ssd (local ssd).
     * 
     */
    public Output<String> category() {
        return this.category;
    }
    /**
     * Disk instance creation time.
     * 
     */
    @Export(name="createTime", refs={String.class}, tree="[0]")
    private Output<String> createTime;

    /**
     * @return Disk instance creation time.
     * 
     */
    public Output<String> createTime() {
        return this.createTime;
    }
    /**
     * Name of the disk instance.
     * 
     */
    @Export(name="diskName", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> diskName;

    /**
     * @return Name of the disk instance.
     * 
     */
    public Output<Optional<String>> diskName() {
        return Codegen.optional(this.diskName);
    }
    /**
     * Indicates whether the cloud disk is Encrypted. If Encrypted = true, the default service key is used when KMSKeyId is not entered. Value range:`true`, `false`(default).
     * 
     */
    @Export(name="encrypted", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> encrypted;

    /**
     * @return Indicates whether the cloud disk is Encrypted. If Encrypted = true, the default service key is used when KMSKeyId is not entered. Value range:`true`, `false`(default).
     * 
     */
    public Output<Optional<Boolean>> encrypted() {
        return Codegen.optional(this.encrypted);
    }
    /**
     * Ens node IDExample value: cn-chengdu-telecom.
     * 
     */
    @Export(name="ensRegionId", refs={String.class}, tree="[0]")
    private Output<String> ensRegionId;

    /**
     * @return Ens node IDExample value: cn-chengdu-telecom.
     * 
     */
    public Output<String> ensRegionId() {
        return this.ensRegionId;
    }
    /**
     * The ID of the KMS key used by the cloud disk. If Encrypted is set to true, the service default key is used when KMSKeyId is empty.
     * 
     */
    @Export(name="kmsKeyId", refs={String.class}, tree="[0]")
    private Output<String> kmsKeyId;

    /**
     * @return The ID of the KMS key used by the cloud disk. If Encrypted is set to true, the service default key is used when KMSKeyId is empty.
     * 
     */
    public Output<String> kmsKeyId() {
        return this.kmsKeyId;
    }
    /**
     * Billing type of the disk instanceValue: PayAsYouGo.
     * 
     */
    @Export(name="paymentType", refs={String.class}, tree="[0]")
    private Output<String> paymentType;

    /**
     * @return Billing type of the disk instanceValue: PayAsYouGo.
     * 
     */
    public Output<String> paymentType() {
        return this.paymentType;
    }
    /**
     * The size of the disk instance. Unit: GiB.
     * 
     */
    @Export(name="size", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> size;

    /**
     * @return The size of the disk instance. Unit: GiB.
     * 
     */
    public Output<Optional<Integer>> size() {
        return Codegen.optional(this.size);
    }
    /**
     * The ID of the snapshot used to create the cloud disk.
     * 
     * The SnapshotId and Size parameters have the following limitations:
     * - If the snapshot capacity corresponding to the **SnapshotId** parameter is greater than the specified **Size** parameter, the Size of the cloud disk created is the Size of the specified snapshot.
     * - If the snapshot capacity corresponding to the **SnapshotId** parameter is less than the set **Size** parameter value, the Size of the cloud disk created is the specified **Size** parameter value.
     * 
     */
    @Export(name="snapshotId", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> snapshotId;

    /**
     * @return The ID of the snapshot used to create the cloud disk.
     * 
     * The SnapshotId and Size parameters have the following limitations:
     * - If the snapshot capacity corresponding to the **SnapshotId** parameter is greater than the specified **Size** parameter, the Size of the cloud disk created is the Size of the specified snapshot.
     * - If the snapshot capacity corresponding to the **SnapshotId** parameter is less than the set **Size** parameter value, the Size of the cloud disk created is the specified **Size** parameter value.
     * 
     */
    public Output<Optional<String>> snapshotId() {
        return Codegen.optional(this.snapshotId);
    }
    /**
     * Status of the disk instance:Value:In-use: In useAvailable: To be mountedAttaching: AttachingDetaching: uninstallingCreating: CreatingReIniting: Resetting.
     * 
     */
    @Export(name="status", refs={String.class}, tree="[0]")
    private Output<String> status;

    /**
     * @return Status of the disk instance:Value:In-use: In useAvailable: To be mountedAttaching: AttachingDetaching: uninstallingCreating: CreatingReIniting: Resetting.
     * 
     */
    public Output<String> status() {
        return this.status;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Disk(String name) {
        this(name, DiskArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Disk(String name, DiskArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Disk(String name, DiskArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:ens/disk:Disk", name, args == null ? DiskArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Disk(String name, Output<String> id, @Nullable DiskState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:ens/disk:Disk", name, state, makeResourceOptions(options, id));
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
    public static Disk get(String name, Output<String> id, @Nullable DiskState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Disk(name, id, state, options);
    }
}
