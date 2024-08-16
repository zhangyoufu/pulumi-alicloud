// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.nas;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.nas.FileSystemArgs;
import com.pulumi.alicloud.nas.inputs.FileSystemState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Integer;
import java.lang.String;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a Nas File System resource.
 * 
 * After activating NAS, you can create a file system and purchase a storage package for it in the NAS console. The NAS console also enables you to view the file system details and remove unnecessary file systems.
 * 
 * For information about NAS file system and how to use it, see [Manage file systems](https://www.alibabacloud.com/help/doc-detail/27530.htm)
 * 
 * &gt; **NOTE:** Available in v1.33.0+.
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
 * import com.pulumi.alicloud.nas.NasFunctions;
 * import com.pulumi.alicloud.nas.inputs.GetZonesArgs;
 * import com.pulumi.alicloud.nas.FileSystem;
 * import com.pulumi.alicloud.nas.FileSystemArgs;
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
 *         final var example = NasFunctions.getZones(GetZonesArgs.builder()
 *             .fileSystemType("standard")
 *             .build());
 * 
 *         var foo = new FileSystem("foo", FileSystemArgs.builder()
 *             .protocolType("NFS")
 *             .storageType("Performance")
 *             .description("terraform-example")
 *             .encryptType("1")
 *             .zoneId(example.applyValue(getZonesResult -> getZonesResult.zones()[0].zoneId()))
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.alicloud.nas.NasFunctions;
 * import com.pulumi.alicloud.nas.inputs.GetZonesArgs;
 * import com.pulumi.alicloud.nas.FileSystem;
 * import com.pulumi.alicloud.nas.FileSystemArgs;
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
 *         final var example = NasFunctions.getZones(GetZonesArgs.builder()
 *             .fileSystemType("extreme")
 *             .build());
 * 
 *         var foo = new FileSystem("foo", FileSystemArgs.builder()
 *             .fileSystemType("extreme")
 *             .protocolType("NFS")
 *             .zoneId(example.applyValue(getZonesResult -> getZonesResult.zones()[0].zoneId()))
 *             .storageType("standard")
 *             .capacity("100")
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
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
 *         final var example = NasFunctions.getZones(GetZonesArgs.builder()
 *             .fileSystemType("cpfs")
 *             .build());
 * 
 *         var exampleNetwork = new Network("exampleNetwork", NetworkArgs.builder()
 *             .vpcName("terraform-example")
 *             .cidrBlock("172.17.3.0/24")
 *             .build());
 * 
 *         var exampleSwitch = new Switch("exampleSwitch", SwitchArgs.builder()
 *             .vswitchName("terraform-example")
 *             .cidrBlock("172.17.3.0/24")
 *             .vpcId(exampleNetwork.id())
 *             .zoneId(example.applyValue(getZonesResult -> getZonesResult.zones()[1].zoneId()))
 *             .build());
 * 
 *         var exampleFileSystem = new FileSystem("exampleFileSystem", FileSystemArgs.builder()
 *             .protocolType("cpfs")
 *             .storageType("advance_200")
 *             .fileSystemType("cpfs")
 *             .capacity(3600)
 *             .zoneId(example.applyValue(getZonesResult -> getZonesResult.zones()[1].zoneId()))
 *             .vpcId(exampleNetwork.id())
 *             .vswitchId(exampleSwitch.id())
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
 * Nas File System can be imported using the id, e.g.
 * 
 * ```sh
 * $ pulumi import alicloud:nas/fileSystem:FileSystem foo 1337849c59
 * ```
 * 
 */
@ResourceType(type="alicloud:nas/fileSystem:FileSystem")
public class FileSystem extends com.pulumi.resources.CustomResource {
    /**
     * The capacity of the file system. The `capacity` is required when the `file_system_type` is `extreme`.
     * Unit: gib; **Note**: The minimum value is 100.
     * 
     */
    @Export(name="capacity", refs={Integer.class}, tree="[0]")
    private Output<Integer> capacity;

    /**
     * @return The capacity of the file system. The `capacity` is required when the `file_system_type` is `extreme`.
     * Unit: gib; **Note**: The minimum value is 100.
     * 
     */
    public Output<Integer> capacity() {
        return this.capacity;
    }
    /**
     * The File System description.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return The File System description.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * Whether the file system is encrypted. Using kms service escrow key to encrypt and store the file system data. When reading and writing encrypted data, there is no need to decrypt.
     * * Valid values:
     * 
     */
    @Export(name="encryptType", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> encryptType;

    /**
     * @return Whether the file system is encrypted. Using kms service escrow key to encrypt and store the file system data. When reading and writing encrypted data, there is no need to decrypt.
     * * Valid values:
     * 
     */
    public Output<Optional<Integer>> encryptType() {
        return Codegen.optional(this.encryptType);
    }
    /**
     * the type of the file system.
     * Valid values:
     * `standard` (Default),
     * `extreme`,
     * `cpfs`.
     * 
     */
    @Export(name="fileSystemType", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> fileSystemType;

    /**
     * @return the type of the file system.
     * Valid values:
     * `standard` (Default),
     * `extreme`,
     * `cpfs`.
     * 
     */
    public Output<Optional<String>> fileSystemType() {
        return Codegen.optional(this.fileSystemType);
    }
    /**
     * The id of the KMS key. The `kms_key_id` is required when the `encrypt_type` is `2`.
     * 
     */
    @Export(name="kmsKeyId", refs={String.class}, tree="[0]")
    private Output<String> kmsKeyId;

    /**
     * @return The id of the KMS key. The `kms_key_id` is required when the `encrypt_type` is `2`.
     * 
     */
    public Output<String> kmsKeyId() {
        return this.kmsKeyId;
    }
    /**
     * The protocol type of the file system.
     * Valid values:
     * `NFS`,
     * `SMB` (Available when the `file_system_type` is `standard`),
     * `cpfs` (Available when the `file_system_type` is `cpfs`).
     * 
     */
    @Export(name="protocolType", refs={String.class}, tree="[0]")
    private Output<String> protocolType;

    /**
     * @return The protocol type of the file system.
     * Valid values:
     * `NFS`,
     * `SMB` (Available when the `file_system_type` is `standard`),
     * `cpfs` (Available when the `file_system_type` is `cpfs`).
     * 
     */
    public Output<String> protocolType() {
        return this.protocolType;
    }
    /**
     * The storage type of the file System.
     * * Valid values:
     * * `Performance` (Available when the `file_system_type` is `standard`)
     * * `Capacity` (Available when the `file_system_type` is `standard`)
     * 
     */
    @Export(name="storageType", refs={String.class}, tree="[0]")
    private Output<String> storageType;

    /**
     * @return The storage type of the file System.
     * * Valid values:
     * * `Performance` (Available when the `file_system_type` is `standard`)
     * * `Capacity` (Available when the `file_system_type` is `standard`)
     * 
     */
    public Output<String> storageType() {
        return this.storageType;
    }
    /**
     * A mapping of tags to assign to the resource.
     * 
     */
    @Export(name="tags", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output</* @Nullable */ Map<String,String>> tags;

    /**
     * @return A mapping of tags to assign to the resource.
     * 
     */
    public Output<Optional<Map<String,String>>> tags() {
        return Codegen.optional(this.tags);
    }
    /**
     * The id of the VPC. The `vpc_id` is required when the `file_system_type` is `cpfs`.
     * 
     */
    @Export(name="vpcId", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> vpcId;

    /**
     * @return The id of the VPC. The `vpc_id` is required when the `file_system_type` is `cpfs`.
     * 
     */
    public Output<Optional<String>> vpcId() {
        return Codegen.optional(this.vpcId);
    }
    /**
     * The id of the vSwitch. The `vswitch_id` is required when the `file_system_type` is `cpfs`.
     * 
     */
    @Export(name="vswitchId", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> vswitchId;

    /**
     * @return The id of the vSwitch. The `vswitch_id` is required when the `file_system_type` is `cpfs`.
     * 
     */
    public Output<Optional<String>> vswitchId() {
        return Codegen.optional(this.vswitchId);
    }
    /**
     * The available zones information that supports nas.When FileSystemType=standard, this parameter is not required. **Note:** By default, a qualified availability zone is randomly selected according to the `protocol_type` and `storage_type` configuration.
     * 
     */
    @Export(name="zoneId", refs={String.class}, tree="[0]")
    private Output<String> zoneId;

    /**
     * @return The available zones information that supports nas.When FileSystemType=standard, this parameter is not required. **Note:** By default, a qualified availability zone is randomly selected according to the `protocol_type` and `storage_type` configuration.
     * 
     */
    public Output<String> zoneId() {
        return this.zoneId;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public FileSystem(java.lang.String name) {
        this(name, FileSystemArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public FileSystem(java.lang.String name, FileSystemArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public FileSystem(java.lang.String name, FileSystemArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:nas/fileSystem:FileSystem", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private FileSystem(java.lang.String name, Output<java.lang.String> id, @Nullable FileSystemState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:nas/fileSystem:FileSystem", name, state, makeResourceOptions(options, id), false);
    }

    private static FileSystemArgs makeArgs(FileSystemArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? FileSystemArgs.Empty : args;
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
    public static FileSystem get(java.lang.String name, Output<java.lang.String> id, @Nullable FileSystemState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new FileSystem(name, id, state, options);
    }
}
