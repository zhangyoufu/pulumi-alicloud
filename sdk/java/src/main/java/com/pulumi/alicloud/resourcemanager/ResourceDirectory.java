// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.resourcemanager;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.resourcemanager.ResourceDirectoryArgs;
import com.pulumi.alicloud.resourcemanager.inputs.ResourceDirectoryState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * Provides a Resource Manager Resource Directory resource. Resource Directory enables you to establish an organizational structure for the resources used by applications of your enterprise. You can plan, build, and manage the resources in a centralized manner by using only one resource directory.
 * 
 * For information about Resource Manager Resource Directory and how to use it, see [What is Resource Manager Resource Directory](https://www.alibabacloud.com/help/en/doc-detail/94475.htm).
 * 
 * &gt; **NOTE:** Available since v1.84.0.
 * 
 * &gt; **NOTE:** An account can only be used to enable a resource directory after it passes enterprise real-name verification. An account that only passed individual real-name verification cannot be used to enable a resource directory.
 * 
 * &gt; **NOTE:** Before you destroy the resource, make sure that the following requirements are met:
 *   - All member accounts must be removed from the resource directory.
 *   - All folders except the root folder must be deleted from the resource directory.
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
 * import com.pulumi.alicloud.resourcemanager.ResourcemanagerFunctions;
 * import com.pulumi.alicloud.resourcemanager.inputs.GetResourceDirectoriesArgs;
 * import com.pulumi.alicloud.resourcemanager.ResourceDirectory;
 * import com.pulumi.alicloud.resourcemanager.ResourceDirectoryArgs;
 * import com.pulumi.codegen.internal.KeyedValue;
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
 *         final var default = ResourcemanagerFunctions.getResourceDirectories();
 * 
 *         for (var i = 0; i < default_.directories().length() > 0 ? 0 : 1; i++) {
 *             new ResourceDirectory("defaultResourceDirectory-" + i, ResourceDirectoryArgs.builder()            
 *                 .status("Enabled")
 *                 .build());
 * 
 *         
 * }
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * Resource Manager Resource Directory can be imported using the id, e.g.
 * 
 * ```sh
 * $ pulumi import alicloud:resourcemanager/resourceDirectory:ResourceDirectory example rd-s3****
 * ```
 * 
 */
@ResourceType(type="alicloud:resourcemanager/resourceDirectory:ResourceDirectory")
public class ResourceDirectory extends com.pulumi.resources.CustomResource {
    /**
     * The ID of the master account.
     * 
     */
    @Export(name="masterAccountId", refs={String.class}, tree="[0]")
    private Output<String> masterAccountId;

    /**
     * @return The ID of the master account.
     * 
     */
    public Output<String> masterAccountId() {
        return this.masterAccountId;
    }
    /**
     * The name of the master account.
     * 
     */
    @Export(name="masterAccountName", refs={String.class}, tree="[0]")
    private Output<String> masterAccountName;

    /**
     * @return The name of the master account.
     * 
     */
    public Output<String> masterAccountName() {
        return this.masterAccountName;
    }
    /**
     * Specifies whether to enable the member deletion feature. Valid values:`Enabled` and `Disabled`.
     * 
     */
    @Export(name="memberDeletionStatus", refs={String.class}, tree="[0]")
    private Output<String> memberDeletionStatus;

    /**
     * @return Specifies whether to enable the member deletion feature. Valid values:`Enabled` and `Disabled`.
     * 
     */
    public Output<String> memberDeletionStatus() {
        return this.memberDeletionStatus;
    }
    /**
     * The ID of the root folder.
     * 
     */
    @Export(name="rootFolderId", refs={String.class}, tree="[0]")
    private Output<String> rootFolderId;

    /**
     * @return The ID of the root folder.
     * 
     */
    public Output<String> rootFolderId() {
        return this.rootFolderId;
    }
    /**
     * The status of control policy. Valid values:`Enabled` and `Disabled`.
     * 
     */
    @Export(name="status", refs={String.class}, tree="[0]")
    private Output<String> status;

    /**
     * @return The status of control policy. Valid values:`Enabled` and `Disabled`.
     * 
     */
    public Output<String> status() {
        return this.status;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public ResourceDirectory(String name) {
        this(name, ResourceDirectoryArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ResourceDirectory(String name, @Nullable ResourceDirectoryArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ResourceDirectory(String name, @Nullable ResourceDirectoryArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:resourcemanager/resourceDirectory:ResourceDirectory", name, args == null ? ResourceDirectoryArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private ResourceDirectory(String name, Output<String> id, @Nullable ResourceDirectoryState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:resourcemanager/resourceDirectory:ResourceDirectory", name, state, makeResourceOptions(options, id));
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
    public static ResourceDirectory get(String name, Output<String> id, @Nullable ResourceDirectoryState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ResourceDirectory(name, id, state, options);
    }
}
