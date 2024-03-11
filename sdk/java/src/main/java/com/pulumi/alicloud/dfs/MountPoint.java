// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.dfs;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.dfs.MountPointArgs;
import com.pulumi.alicloud.dfs.inputs.MountPointState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a DFS Mount Point resource.
 * 
 * For information about DFS Mount Point and how to use it, see [What is Mount Point](https://www.alibabacloud.com/help/en/aibaba-cloud-storage-services/latest/apsara-file-storage-for-hdfs).
 * 
 * &gt; **NOTE:** Available since v1.140.0.
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
 * import com.pulumi.alicloud.dfs.DfsFunctions;
 * import com.pulumi.alicloud.dfs.inputs.GetZonesArgs;
 * import com.pulumi.random.RandomInteger;
 * import com.pulumi.random.RandomIntegerArgs;
 * import com.pulumi.alicloud.vpc.Network;
 * import com.pulumi.alicloud.vpc.NetworkArgs;
 * import com.pulumi.alicloud.vpc.Switch;
 * import com.pulumi.alicloud.vpc.SwitchArgs;
 * import com.pulumi.alicloud.dfs.AccessGroup;
 * import com.pulumi.alicloud.dfs.AccessGroupArgs;
 * import com.pulumi.alicloud.dfs.FileSystem;
 * import com.pulumi.alicloud.dfs.FileSystemArgs;
 * import com.pulumi.alicloud.dfs.MountPoint;
 * import com.pulumi.alicloud.dfs.MountPointArgs;
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
 *         final var defaultZones = DfsFunctions.getZones();
 * 
 *         var defaultRandomInteger = new RandomInteger(&#34;defaultRandomInteger&#34;, RandomIntegerArgs.builder()        
 *             .min(10000)
 *             .max(99999)
 *             .build());
 * 
 *         var defaultVPC = new Network(&#34;defaultVPC&#34;, NetworkArgs.builder()        
 *             .cidrBlock(&#34;172.16.0.0/12&#34;)
 *             .vpcName(name)
 *             .build());
 * 
 *         var defaultVSwitch = new Switch(&#34;defaultVSwitch&#34;, SwitchArgs.builder()        
 *             .description(&#34;example&#34;)
 *             .vpcId(defaultVPC.id())
 *             .cidrBlock(&#34;172.16.0.0/24&#34;)
 *             .vswitchName(name)
 *             .zoneId(defaultZones.applyValue(getZonesResult -&gt; getZonesResult.zones()[0].zoneId()))
 *             .build());
 * 
 *         var defaultAccessGroup = new AccessGroup(&#34;defaultAccessGroup&#34;, AccessGroupArgs.builder()        
 *             .description(&#34;AccessGroup resource manager center example&#34;)
 *             .networkType(&#34;VPC&#34;)
 *             .accessGroupName(defaultRandomInteger.result().applyValue(result -&gt; String.format(&#34;%s-%s&#34;, name,result)))
 *             .build());
 * 
 *         var updateAccessGroup = new AccessGroup(&#34;updateAccessGroup&#34;, AccessGroupArgs.builder()        
 *             .description(&#34;Second AccessGroup resource manager center example&#34;)
 *             .networkType(&#34;VPC&#34;)
 *             .accessGroupName(defaultRandomInteger.result().applyValue(result -&gt; String.format(&#34;%s-update-%s&#34;, name,result)))
 *             .build());
 * 
 *         var defaultFs = new FileSystem(&#34;defaultFs&#34;, FileSystemArgs.builder()        
 *             .spaceCapacity(&#34;1024&#34;)
 *             .description(&#34;for mountpoint  example&#34;)
 *             .storageType(&#34;STANDARD&#34;)
 *             .zoneId(defaultZones.applyValue(getZonesResult -&gt; getZonesResult.zones()[0].zoneId()))
 *             .protocolType(&#34;HDFS&#34;)
 *             .dataRedundancyType(&#34;LRS&#34;)
 *             .fileSystemName(defaultRandomInteger.result().applyValue(result -&gt; String.format(&#34;%s-%s&#34;, name,result)))
 *             .build());
 * 
 *         var defaultMountPoint = new MountPoint(&#34;defaultMountPoint&#34;, MountPointArgs.builder()        
 *             .vpcId(defaultVPC.id())
 *             .description(&#34;mountpoint example&#34;)
 *             .networkType(&#34;VPC&#34;)
 *             .vswitchId(defaultVSwitch.id())
 *             .fileSystemId(defaultFs.id())
 *             .accessGroupId(defaultAccessGroup.id())
 *             .status(&#34;Active&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * DFS Mount Point can be imported using the id, e.g.
 * 
 * ```sh
 * $ pulumi import alicloud:dfs/mountPoint:MountPoint example &lt;file_system_id&gt;:&lt;mount_point_id&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:dfs/mountPoint:MountPoint")
public class MountPoint extends com.pulumi.resources.CustomResource {
    /**
     * The id of the permission group associated with the Mount point, which is used to set the access permissions of the Mount point.
     * 
     */
    @Export(name="accessGroupId", refs={String.class}, tree="[0]")
    private Output<String> accessGroupId;

    /**
     * @return The id of the permission group associated with the Mount point, which is used to set the access permissions of the Mount point.
     * 
     */
    public Output<String> accessGroupId() {
        return this.accessGroupId;
    }
    /**
     * The mount point alias prefix, which specifies the mount point alias prefix.
     * 
     */
    @Export(name="aliasPrefix", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> aliasPrefix;

    /**
     * @return The mount point alias prefix, which specifies the mount point alias prefix.
     * 
     */
    public Output<Optional<String>> aliasPrefix() {
        return Codegen.optional(this.aliasPrefix);
    }
    /**
     * The creation time of the Mount point resource.
     * 
     */
    @Export(name="createTime", refs={String.class}, tree="[0]")
    private Output<String> createTime;

    /**
     * @return The creation time of the Mount point resource.
     * 
     */
    public Output<String> createTime() {
        return this.createTime;
    }
    /**
     * The description of the Mount point.  No more than 32 characters in length.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return The description of the Mount point.  No more than 32 characters in length.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * Unique file system identifier, used to retrieve specified file system resources.
     * 
     */
    @Export(name="fileSystemId", refs={String.class}, tree="[0]")
    private Output<String> fileSystemId;

    /**
     * @return Unique file system identifier, used to retrieve specified file system resources.
     * 
     */
    public Output<String> fileSystemId() {
        return this.fileSystemId;
    }
    /**
     * The unique identifier of the Mount point, which is used to retrieve the specified mount point resources.
     * 
     */
    @Export(name="mountPointId", refs={String.class}, tree="[0]")
    private Output<String> mountPointId;

    /**
     * @return The unique identifier of the Mount point, which is used to retrieve the specified mount point resources.
     * 
     */
    public Output<String> mountPointId() {
        return this.mountPointId;
    }
    /**
     * The network type of the Mount point.  Only VPC (VPC) is supported.
     * 
     */
    @Export(name="networkType", refs={String.class}, tree="[0]")
    private Output<String> networkType;

    /**
     * @return The network type of the Mount point.  Only VPC (VPC) is supported.
     * 
     */
    public Output<String> networkType() {
        return this.networkType;
    }
    /**
     * Mount point status. Value: Inactive: Disable mount points Active: Activate the mount point.
     * 
     */
    @Export(name="status", refs={String.class}, tree="[0]")
    private Output<String> status;

    /**
     * @return Mount point status. Value: Inactive: Disable mount points Active: Activate the mount point.
     * 
     */
    public Output<String> status() {
        return this.status;
    }
    /**
     * The ID of the VPC. Specifies the VPC environment to which the mount point belongs.
     * 
     */
    @Export(name="vpcId", refs={String.class}, tree="[0]")
    private Output<String> vpcId;

    /**
     * @return The ID of the VPC. Specifies the VPC environment to which the mount point belongs.
     * 
     */
    public Output<String> vpcId() {
        return this.vpcId;
    }
    /**
     * VSwitch ID, which specifies the VSwitch resource used to create the mount point.
     * 
     */
    @Export(name="vswitchId", refs={String.class}, tree="[0]")
    private Output<String> vswitchId;

    /**
     * @return VSwitch ID, which specifies the VSwitch resource used to create the mount point.
     * 
     */
    public Output<String> vswitchId() {
        return this.vswitchId;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public MountPoint(String name) {
        this(name, MountPointArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public MountPoint(String name, MountPointArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public MountPoint(String name, MountPointArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:dfs/mountPoint:MountPoint", name, args == null ? MountPointArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private MountPoint(String name, Output<String> id, @Nullable MountPointState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:dfs/mountPoint:MountPoint", name, state, makeResourceOptions(options, id));
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
    public static MountPoint get(String name, Output<String> id, @Nullable MountPointState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new MountPoint(name, id, state, options);
    }
}
