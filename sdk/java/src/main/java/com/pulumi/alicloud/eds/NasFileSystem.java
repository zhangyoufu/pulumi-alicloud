// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.eds;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.eds.NasFileSystemArgs;
import com.pulumi.alicloud.eds.inputs.NasFileSystemState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Boolean;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a ECD Nas File System resource.
 * 
 * For information about ECD Nas File System and how to use it, see [What is Nas File System](https://www.alibabacloud.com/help/en/elastic-desktop-service/latest/api-reference-for-easy-use-1).
 * 
 * &gt; **NOTE:** Available since v1.141.0.
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
 * import com.pulumi.random.integer;
 * import com.pulumi.random.IntegerArgs;
 * import com.pulumi.alicloud.eds.SimpleOfficeSite;
 * import com.pulumi.alicloud.eds.SimpleOfficeSiteArgs;
 * import com.pulumi.alicloud.eds.NasFileSystem;
 * import com.pulumi.alicloud.eds.NasFileSystemArgs;
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
 *         var default_ = new Integer("default", IntegerArgs.builder()
 *             .min(10000)
 *             .max(99999)
 *             .build());
 * 
 *         var defaultSimpleOfficeSite = new SimpleOfficeSite("defaultSimpleOfficeSite", SimpleOfficeSiteArgs.builder()
 *             .cidrBlock("172.16.0.0/12")
 *             .enableAdminAccess(false)
 *             .desktopAccessType("Internet")
 *             .officeSiteName(String.format("%s-%s", name,default_.result()))
 *             .build());
 * 
 *         var example = new NasFileSystem("example", NasFileSystemArgs.builder()
 *             .nasFileSystemName(name)
 *             .officeSiteId(defaultSimpleOfficeSite.id())
 *             .description(name)
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
 * ECD Nas File System can be imported using the id, e.g.
 * 
 * ```sh
 * $ pulumi import alicloud:eds/nasFileSystem:NasFileSystem example &lt;id&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:eds/nasFileSystem:NasFileSystem")
public class NasFileSystem extends com.pulumi.resources.CustomResource {
    /**
     * The description of nas file system.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return The description of nas file system.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * The filesystem id of nas file system.
     * 
     */
    @Export(name="fileSystemId", refs={String.class}, tree="[0]")
    private Output<String> fileSystemId;

    /**
     * @return The filesystem id of nas file system.
     * 
     */
    public Output<String> fileSystemId() {
        return this.fileSystemId;
    }
    /**
     * The domain of mount target.
     * 
     */
    @Export(name="mountTargetDomain", refs={String.class}, tree="[0]")
    private Output<String> mountTargetDomain;

    /**
     * @return The domain of mount target.
     * 
     */
    public Output<String> mountTargetDomain() {
        return this.mountTargetDomain;
    }
    /**
     * The name of nas file system.
     * 
     */
    @Export(name="nasFileSystemName", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> nasFileSystemName;

    /**
     * @return The name of nas file system.
     * 
     */
    public Output<Optional<String>> nasFileSystemName() {
        return Codegen.optional(this.nasFileSystemName);
    }
    /**
     * The ID of office site.
     * 
     */
    @Export(name="officeSiteId", refs={String.class}, tree="[0]")
    private Output<String> officeSiteId;

    /**
     * @return The ID of office site.
     * 
     */
    public Output<String> officeSiteId() {
        return this.officeSiteId;
    }
    /**
     * The mount point is in an inactive state, reset the mount point of the NAS file system. Default to `false`.
     * 
     */
    @Export(name="reset", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> reset;

    /**
     * @return The mount point is in an inactive state, reset the mount point of the NAS file system. Default to `false`.
     * 
     */
    public Output<Optional<Boolean>> reset() {
        return Codegen.optional(this.reset);
    }
    /**
     * The status of nas file system. Valid values: `Pending`, `Running`, `Stopped`,`Deleting`, `Deleted`, `Invalid`.
     * 
     */
    @Export(name="status", refs={String.class}, tree="[0]")
    private Output<String> status;

    /**
     * @return The status of nas file system. Valid values: `Pending`, `Running`, `Stopped`,`Deleting`, `Deleted`, `Invalid`.
     * 
     */
    public Output<String> status() {
        return this.status;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public NasFileSystem(java.lang.String name) {
        this(name, NasFileSystemArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public NasFileSystem(java.lang.String name, NasFileSystemArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public NasFileSystem(java.lang.String name, NasFileSystemArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:eds/nasFileSystem:NasFileSystem", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private NasFileSystem(java.lang.String name, Output<java.lang.String> id, @Nullable NasFileSystemState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:eds/nasFileSystem:NasFileSystem", name, state, makeResourceOptions(options, id), false);
    }

    private static NasFileSystemArgs makeArgs(NasFileSystemArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? NasFileSystemArgs.Empty : args;
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
    public static NasFileSystem get(java.lang.String name, Output<java.lang.String> id, @Nullable NasFileSystemState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new NasFileSystem(name, id, state, options);
    }
}
