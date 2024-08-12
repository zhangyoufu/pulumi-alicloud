// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.eds;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.eds.RamDirectoryArgs;
import com.pulumi.alicloud.eds.inputs.RamDirectoryState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import javax.annotation.Nullable;

/**
 * Provides a ECD Ram Directory resource.
 * 
 * For information about ECD Ram Directory and how to use it, see [What is Ram Directory](https://www.alibabacloud.com/help/en/wuying-workspace/developer-reference/api-ecd-2020-09-30-createramdirectory).
 * 
 * &gt; **NOTE:** Available since v1.174.0.
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
 * import com.pulumi.alicloud.eds.EdsFunctions;
 * import com.pulumi.alicloud.eds.inputs.GetZonesArgs;
 * import com.pulumi.alicloud.vpc.Network;
 * import com.pulumi.alicloud.vpc.NetworkArgs;
 * import com.pulumi.alicloud.vpc.Switch;
 * import com.pulumi.alicloud.vpc.SwitchArgs;
 * import com.pulumi.alicloud.eds.RamDirectory;
 * import com.pulumi.alicloud.eds.RamDirectoryArgs;
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
 *         final var default = EdsFunctions.getZones();
 * 
 *         var defaultNetwork = new Network("defaultNetwork", NetworkArgs.builder()
 *             .vpcName(name)
 *             .cidrBlock("172.16.0.0/16")
 *             .build());
 * 
 *         var defaultSwitch = new Switch("defaultSwitch", SwitchArgs.builder()
 *             .vpcId(defaultNetwork.id())
 *             .cidrBlock("172.16.0.0/24")
 *             .zoneId(default_.ids()[0])
 *             .vswitchName(name)
 *             .build());
 * 
 *         var defaultRamDirectory = new RamDirectory("defaultRamDirectory", RamDirectoryArgs.builder()
 *             .desktopAccessType("INTERNET")
 *             .enableAdminAccess(true)
 *             .ramDirectoryName(name)
 *             .vswitchIds(defaultSwitch.id())
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
 * ECD Ram Directory can be imported using the id, e.g.
 * 
 * ```sh
 * $ pulumi import alicloud:eds/ramDirectory:RamDirectory example &lt;id&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:eds/ramDirectory:RamDirectory")
public class RamDirectory extends com.pulumi.resources.CustomResource {
    /**
     * The desktop access type. Valid values: `VPC`, `INTERNET`, `ANY`.
     * 
     */
    @Export(name="desktopAccessType", refs={String.class}, tree="[0]")
    private Output<String> desktopAccessType;

    /**
     * @return The desktop access type. Valid values: `VPC`, `INTERNET`, `ANY`.
     * 
     */
    public Output<String> desktopAccessType() {
        return this.desktopAccessType;
    }
    /**
     * Whether to enable public network access.
     * 
     */
    @Export(name="enableAdminAccess", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> enableAdminAccess;

    /**
     * @return Whether to enable public network access.
     * 
     */
    public Output<Boolean> enableAdminAccess() {
        return this.enableAdminAccess;
    }
    /**
     * Whether to grant local administrator rights to users who use cloud desktops.
     * 
     */
    @Export(name="enableInternetAccess", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> enableInternetAccess;

    /**
     * @return Whether to grant local administrator rights to users who use cloud desktops.
     * 
     */
    public Output<Boolean> enableInternetAccess() {
        return this.enableInternetAccess;
    }
    /**
     * The name of the directory. The name must be 2 to 255 characters in length. It must start with a letter but cannot start with `http://` or `https://`. It can contain letters, digits, colons (:), underscores (_), and hyphens (-).
     * 
     */
    @Export(name="ramDirectoryName", refs={String.class}, tree="[0]")
    private Output<String> ramDirectoryName;

    /**
     * @return The name of the directory. The name must be 2 to 255 characters in length. It must start with a letter but cannot start with `http://` or `https://`. It can contain letters, digits, colons (:), underscores (_), and hyphens (-).
     * 
     */
    public Output<String> ramDirectoryName() {
        return this.ramDirectoryName;
    }
    /**
     * The status of directory.
     * 
     */
    @Export(name="status", refs={String.class}, tree="[0]")
    private Output<String> status;

    /**
     * @return The status of directory.
     * 
     */
    public Output<String> status() {
        return this.status;
    }
    /**
     * List of VSwitch IDs in the directory.
     * 
     */
    @Export(name="vswitchIds", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> vswitchIds;

    /**
     * @return List of VSwitch IDs in the directory.
     * 
     */
    public Output<List<String>> vswitchIds() {
        return this.vswitchIds;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public RamDirectory(java.lang.String name) {
        this(name, RamDirectoryArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public RamDirectory(java.lang.String name, RamDirectoryArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public RamDirectory(java.lang.String name, RamDirectoryArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:eds/ramDirectory:RamDirectory", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private RamDirectory(java.lang.String name, Output<java.lang.String> id, @Nullable RamDirectoryState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:eds/ramDirectory:RamDirectory", name, state, makeResourceOptions(options, id), false);
    }

    private static RamDirectoryArgs makeArgs(RamDirectoryArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? RamDirectoryArgs.Empty : args;
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
    public static RamDirectory get(java.lang.String name, Output<java.lang.String> id, @Nullable RamDirectoryState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new RamDirectory(name, id, state, options);
    }
}
