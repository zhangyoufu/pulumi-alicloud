// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.ebs;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.ebs.DedicatedBlockStorageClusterArgs;
import com.pulumi.alicloud.ebs.inputs.DedicatedBlockStorageClusterState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * Provides a Ebs Dedicated Block Storage Cluster resource.
 * 
 * For information about Ebs Dedicated Block Storage Cluster and how to use it, see [What is Dedicated Block Storage Cluster](https://www.alibabacloud.com/help/en/ecs/developer-reference/api-ebs-2021-07-30-creatededicatedblockstoragecluster).
 * 
 * &gt; **NOTE:** Available in v1.195.0+.
 * 
 * ## Example Usage
 * 
 * Basic Usage
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.alicloud.ebs.DedicatedBlockStorageCluster;
 * import com.pulumi.alicloud.ebs.DedicatedBlockStorageClusterArgs;
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
 *         var default_ = new DedicatedBlockStorageCluster(&#34;default&#34;, DedicatedBlockStorageClusterArgs.builder()        
 *             .dedicatedBlockStorageClusterName(&#34;dedicated_block_storage_cluster_name&#34;)
 *             .regionId(&#34;cn-heyuan&#34;)
 *             .totalCapacity(61440)
 *             .type(&#34;Premium&#34;)
 *             .zoneId(&#34;cn-heyuan-b&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * Ebs Dedicated Block Storage Cluster can be imported using the id, e.g.
 * 
 * ```sh
 * $ pulumi import alicloud:ebs/dedicatedBlockStorageCluster:DedicatedBlockStorageCluster example &lt;id&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:ebs/dedicatedBlockStorageCluster:DedicatedBlockStorageCluster")
public class DedicatedBlockStorageCluster extends com.pulumi.resources.CustomResource {
    /**
     * The available capacity of the dedicated block storage cluster. Unit: GiB.
     * 
     */
    @Export(name="availableCapacity", refs={String.class}, tree="[0]")
    private Output<String> availableCapacity;

    /**
     * @return The available capacity of the dedicated block storage cluster. Unit: GiB.
     * 
     */
    public Output<String> availableCapacity() {
        return this.availableCapacity;
    }
    /**
     * The type of cloud disk that can be created by a dedicated block storage cluster.
     * 
     */
    @Export(name="category", refs={String.class}, tree="[0]")
    private Output<String> category;

    /**
     * @return The type of cloud disk that can be created by a dedicated block storage cluster.
     * 
     */
    public Output<String> category() {
        return this.category;
    }
    /**
     * The creation time of the resource
     * 
     */
    @Export(name="createTime", refs={String.class}, tree="[0]")
    private Output<String> createTime;

    /**
     * @return The creation time of the resource
     * 
     */
    public Output<String> createTime() {
        return this.createTime;
    }
    /**
     * The first ID of the resource
     * 
     */
    @Export(name="dedicatedBlockStorageClusterId", refs={String.class}, tree="[0]")
    private Output<String> dedicatedBlockStorageClusterId;

    /**
     * @return The first ID of the resource
     * 
     */
    public Output<String> dedicatedBlockStorageClusterId() {
        return this.dedicatedBlockStorageClusterId;
    }
    /**
     * The name of the resource
     * 
     */
    @Export(name="dedicatedBlockStorageClusterName", refs={String.class}, tree="[0]")
    private Output<String> dedicatedBlockStorageClusterName;

    /**
     * @return The name of the resource
     * 
     */
    public Output<String> dedicatedBlockStorageClusterName() {
        return this.dedicatedBlockStorageClusterName;
    }
    /**
     * Capacity to be delivered in GB.
     * 
     */
    @Export(name="deliveryCapacity", refs={String.class}, tree="[0]")
    private Output<String> deliveryCapacity;

    /**
     * @return Capacity to be delivered in GB.
     * 
     */
    public Output<String> deliveryCapacity() {
        return this.deliveryCapacity;
    }
    /**
     * The description of the dedicated block storage cluster.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output<String> description;

    /**
     * @return The description of the dedicated block storage cluster.
     * 
     */
    public Output<String> description() {
        return this.description;
    }
    /**
     * The expiration time of the dedicated block storage cluster, in the Unix timestamp format, in seconds.
     * 
     */
    @Export(name="expiredTime", refs={String.class}, tree="[0]")
    private Output<String> expiredTime;

    /**
     * @return The expiration time of the dedicated block storage cluster, in the Unix timestamp format, in seconds.
     * 
     */
    public Output<String> expiredTime() {
        return this.expiredTime;
    }
    /**
     * Cloud disk performance level, possible values:-PL0.-PL1.-PL2.-PL3.&gt; Only valid in SupportedCategory = cloud_essd.
     * 
     */
    @Export(name="performanceLevel", refs={String.class}, tree="[0]")
    private Output<String> performanceLevel;

    /**
     * @return Cloud disk performance level, possible values:-PL0.-PL1.-PL2.-PL3.&gt; Only valid in SupportedCategory = cloud_essd.
     * 
     */
    public Output<String> performanceLevel() {
        return this.performanceLevel;
    }
    /**
     * The ID of the resource group
     * 
     */
    @Export(name="resourceGroupId", refs={String.class}, tree="[0]")
    private Output<String> resourceGroupId;

    /**
     * @return The ID of the resource group
     * 
     */
    public Output<String> resourceGroupId() {
        return this.resourceGroupId;
    }
    /**
     * The status of the resource
     * 
     */
    @Export(name="status", refs={String.class}, tree="[0]")
    private Output<String> status;

    /**
     * @return The status of the resource
     * 
     */
    public Output<String> status() {
        return this.status;
    }
    /**
     * This parameter is not supported.
     * 
     */
    @Export(name="supportedCategory", refs={String.class}, tree="[0]")
    private Output<String> supportedCategory;

    /**
     * @return This parameter is not supported.
     * 
     */
    public Output<String> supportedCategory() {
        return this.supportedCategory;
    }
    /**
     * The total capacity of the dedicated block storage cluster. Unit: GiB.
     * 
     */
    @Export(name="totalCapacity", refs={String.class}, tree="[0]")
    private Output<String> totalCapacity;

    /**
     * @return The total capacity of the dedicated block storage cluster. Unit: GiB.
     * 
     */
    public Output<String> totalCapacity() {
        return this.totalCapacity;
    }
    /**
     * The dedicated block storage cluster performance type. Possible values:-Standard: Basic type. This type of dedicated block storage cluster can create an ESSD PL0 cloud disk.-Premium: performance type. This type of dedicated block storage cluster can create an ESSD PL1 cloud disk.
     * 
     */
    @Export(name="type", refs={String.class}, tree="[0]")
    private Output<String> type;

    /**
     * @return The dedicated block storage cluster performance type. Possible values:-Standard: Basic type. This type of dedicated block storage cluster can create an ESSD PL0 cloud disk.-Premium: performance type. This type of dedicated block storage cluster can create an ESSD PL1 cloud disk.
     * 
     */
    public Output<String> type() {
        return this.type;
    }
    /**
     * The used (created disk) capacity of the current cluster, in GB
     * 
     */
    @Export(name="usedCapacity", refs={String.class}, tree="[0]")
    private Output<String> usedCapacity;

    /**
     * @return The used (created disk) capacity of the current cluster, in GB
     * 
     */
    public Output<String> usedCapacity() {
        return this.usedCapacity;
    }
    /**
     * The zone ID  of the resource
     * 
     */
    @Export(name="zoneId", refs={String.class}, tree="[0]")
    private Output<String> zoneId;

    /**
     * @return The zone ID  of the resource
     * 
     */
    public Output<String> zoneId() {
        return this.zoneId;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public DedicatedBlockStorageCluster(String name) {
        this(name, DedicatedBlockStorageClusterArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public DedicatedBlockStorageCluster(String name, DedicatedBlockStorageClusterArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public DedicatedBlockStorageCluster(String name, DedicatedBlockStorageClusterArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:ebs/dedicatedBlockStorageCluster:DedicatedBlockStorageCluster", name, args == null ? DedicatedBlockStorageClusterArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private DedicatedBlockStorageCluster(String name, Output<String> id, @Nullable DedicatedBlockStorageClusterState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:ebs/dedicatedBlockStorageCluster:DedicatedBlockStorageCluster", name, state, makeResourceOptions(options, id));
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
    public static DedicatedBlockStorageCluster get(String name, Output<String> id, @Nullable DedicatedBlockStorageClusterState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new DedicatedBlockStorageCluster(name, id, state, options);
    }
}
