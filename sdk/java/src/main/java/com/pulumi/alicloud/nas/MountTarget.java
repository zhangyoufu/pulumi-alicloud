// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.nas;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.nas.MountTargetArgs;
import com.pulumi.alicloud.nas.inputs.MountTargetState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a NAS Mount Target resource.
 * For information about NAS Mount Target and how to use it, see [Manage NAS Mount Targets](https://www.alibabacloud.com/help/en/doc-detail/27531.htm).
 * 
 * &gt; **NOTE:** Available since v1.34.0.
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
 * import com.pulumi.alicloud.nas.NasFunctions;
 * import com.pulumi.alicloud.nas.inputs.GetZonesArgs;
 * import com.pulumi.alicloud.vpc.Network;
 * import com.pulumi.alicloud.vpc.NetworkArgs;
 * import com.pulumi.alicloud.vpc.Switch;
 * import com.pulumi.alicloud.vpc.SwitchArgs;
 * import com.pulumi.alicloud.nas.FileSystem;
 * import com.pulumi.alicloud.nas.FileSystemArgs;
 * import com.pulumi.alicloud.nas.AccessGroup;
 * import com.pulumi.alicloud.nas.AccessGroupArgs;
 * import com.pulumi.alicloud.nas.MountTarget;
 * import com.pulumi.alicloud.nas.MountTargetArgs;
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
 *         final var default = NasFunctions.getZones(GetZonesArgs.builder()
 *             .fileSystemType(&#34;extreme&#34;)
 *             .build());
 * 
 *         final var countSize = default_.zones().length();
 * 
 *         final var zoneId = default_.zones()[countSize - 1].zoneId();
 * 
 *         var example = new Network(&#34;example&#34;, NetworkArgs.builder()        
 *             .vpcName(&#34;terraform-example&#34;)
 *             .cidrBlock(&#34;172.17.3.0/24&#34;)
 *             .build());
 * 
 *         var exampleSwitch = new Switch(&#34;exampleSwitch&#34;, SwitchArgs.builder()        
 *             .vswitchName(example.vpcName())
 *             .cidrBlock(example.cidrBlock())
 *             .vpcId(example.id())
 *             .zoneId(zoneId)
 *             .build());
 * 
 *         var exampleFileSystem = new FileSystem(&#34;exampleFileSystem&#34;, FileSystemArgs.builder()        
 *             .protocolType(&#34;NFS&#34;)
 *             .storageType(&#34;advance&#34;)
 *             .fileSystemType(&#34;extreme&#34;)
 *             .capacity(&#34;100&#34;)
 *             .zoneId(zoneId)
 *             .build());
 * 
 *         var exampleAccessGroup = new AccessGroup(&#34;exampleAccessGroup&#34;, AccessGroupArgs.builder()        
 *             .accessGroupName(&#34;access_group_xxx&#34;)
 *             .accessGroupType(&#34;Vpc&#34;)
 *             .description(&#34;test_access_group&#34;)
 *             .fileSystemType(&#34;extreme&#34;)
 *             .build());
 * 
 *         var exampleMountTarget = new MountTarget(&#34;exampleMountTarget&#34;, MountTargetArgs.builder()        
 *             .fileSystemId(exampleFileSystem.id())
 *             .accessGroupName(exampleAccessGroup.accessGroupName())
 *             .vswitchId(exampleSwitch.id())
 *             .vpcId(example.id())
 *             .networkType(exampleAccessGroup.accessGroupType())
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * NAS MountTarget can be imported using the id, e.g.
 * 
 * ```sh
 * $ pulumi import alicloud:nas/mountTarget:MountTarget foo 192094b415:192094b415-luw38.cn-beijing.nas.aliyuncs.com
 * ```
 * 
 */
@ResourceType(type="alicloud:nas/mountTarget:MountTarget")
public class MountTarget extends com.pulumi.resources.CustomResource {
    /**
     * The name of the permission group that applies to the mount target.
     * 
     */
    @Export(name="accessGroupName", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> accessGroupName;

    /**
     * @return The name of the permission group that applies to the mount target.
     * 
     */
    public Output<Optional<String>> accessGroupName() {
        return Codegen.optional(this.accessGroupName);
    }
    /**
     * The ID of the file system.
     * 
     */
    @Export(name="fileSystemId", refs={String.class}, tree="[0]")
    private Output<String> fileSystemId;

    /**
     * @return The ID of the file system.
     * 
     */
    public Output<String> fileSystemId() {
        return this.fileSystemId;
    }
    /**
     * The IPv4 domain name of the mount target. **NOTE:** Available since v1.161.0.
     * 
     */
    @Export(name="mountTargetDomain", refs={String.class}, tree="[0]")
    private Output<String> mountTargetDomain;

    /**
     * @return The IPv4 domain name of the mount target. **NOTE:** Available since v1.161.0.
     * 
     */
    public Output<String> mountTargetDomain() {
        return this.mountTargetDomain;
    }
    /**
     * mount target network type. Valid values: `VPC`. The classic network&#39;s mount targets are not supported.
     * 
     */
    @Export(name="networkType", refs={String.class}, tree="[0]")
    private Output<String> networkType;

    /**
     * @return mount target network type. Valid values: `VPC`. The classic network&#39;s mount targets are not supported.
     * 
     */
    public Output<String> networkType() {
        return this.networkType;
    }
    /**
     * The ID of security group.
     * 
     */
    @Export(name="securityGroupId", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> securityGroupId;

    /**
     * @return The ID of security group.
     * 
     */
    public Output<Optional<String>> securityGroupId() {
        return Codegen.optional(this.securityGroupId);
    }
    /**
     * Whether the MountTarget is active. The status of the mount target. Valid values: `Active` and `Inactive`, Default value is `Active`. Before you mount a file system, make sure that the mount target is in the Active state.
     * 
     */
    @Export(name="status", refs={String.class}, tree="[0]")
    private Output<String> status;

    /**
     * @return Whether the MountTarget is active. The status of the mount target. Valid values: `Active` and `Inactive`, Default value is `Active`. Before you mount a file system, make sure that the mount target is in the Active state.
     * 
     */
    public Output<String> status() {
        return this.status;
    }
    /**
     * The ID of VPC.
     * 
     */
    @Export(name="vpcId", refs={String.class}, tree="[0]")
    private Output<String> vpcId;

    /**
     * @return The ID of VPC.
     * 
     */
    public Output<String> vpcId() {
        return this.vpcId;
    }
    /**
     * The ID of the VSwitch in the VPC where the mount target resides.
     * 
     */
    @Export(name="vswitchId", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> vswitchId;

    /**
     * @return The ID of the VSwitch in the VPC where the mount target resides.
     * 
     */
    public Output<Optional<String>> vswitchId() {
        return Codegen.optional(this.vswitchId);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public MountTarget(String name) {
        this(name, MountTargetArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public MountTarget(String name, MountTargetArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public MountTarget(String name, MountTargetArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:nas/mountTarget:MountTarget", name, args == null ? MountTargetArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private MountTarget(String name, Output<String> id, @Nullable MountTargetState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:nas/mountTarget:MountTarget", name, state, makeResourceOptions(options, id));
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
    public static MountTarget get(String name, Output<String> id, @Nullable MountTargetState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new MountTarget(name, id, state, options);
    }
}
