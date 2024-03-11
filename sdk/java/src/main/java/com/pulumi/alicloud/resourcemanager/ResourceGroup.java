// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.resourcemanager;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.resourcemanager.ResourceGroupArgs;
import com.pulumi.alicloud.resourcemanager.inputs.ResourceGroupState;
import com.pulumi.alicloud.resourcemanager.outputs.ResourceGroupRegionStatus;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.List;
import javax.annotation.Nullable;

/**
 * Provides a Resource Manager Resource Group resource. If you need to group cloud resources according to business departments, projects, and other dimensions, you can create resource groups.
 * For information about Resource Manager Resoource Group and how to use it, see [What is Resource Manager Resource Group](https://www.alibabacloud.com/help/en/doc-detail/94485.htm)
 * 
 * &gt; **NOTE:** Available since v1.82.0.
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
 * import com.pulumi.alicloud.resourcemanager.ResourceGroup;
 * import com.pulumi.alicloud.resourcemanager.ResourceGroupArgs;
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
 *         final var name = config.get(&#34;name&#34;).orElse(&#34;tfexample&#34;);
 *         var example = new ResourceGroup(&#34;example&#34;, ResourceGroupArgs.builder()        
 *             .resourceGroupName(name)
 *             .displayName(name)
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * Resource Manager Resource Group can be imported using the id, e.g.
 * 
 * ```sh
 * $ pulumi import alicloud:resourcemanager/resourceGroup:ResourceGroup example abc123456
 * ```
 * 
 */
@ResourceType(type="alicloud:resourcemanager/resourceGroup:ResourceGroup")
public class ResourceGroup extends com.pulumi.resources.CustomResource {
    /**
     * The ID of the Alibaba Cloud account to which the resource group belongs.
     * 
     */
    @Export(name="accountId", refs={String.class}, tree="[0]")
    private Output<String> accountId;

    /**
     * @return The ID of the Alibaba Cloud account to which the resource group belongs.
     * 
     */
    public Output<String> accountId() {
        return this.accountId;
    }
    /**
     * The display name of the resource group. The name must be 1 to 30 characters in length and can contain letters, digits, periods (.), at signs (@), and hyphens (-).
     * 
     */
    @Export(name="displayName", refs={String.class}, tree="[0]")
    private Output<String> displayName;

    /**
     * @return The display name of the resource group. The name must be 1 to 30 characters in length and can contain letters, digits, periods (.), at signs (@), and hyphens (-).
     * 
     */
    public Output<String> displayName() {
        return this.displayName;
    }
    /**
     * Field `name` has been deprecated from version 1.114.0. Use `resource_group_name` instead.
     * 
     * @deprecated
     * Field &#39;name&#39; has been deprecated from version 1.114.0. Use &#39;resource_group_name&#39; instead.
     * 
     */
    @Deprecated /* Field 'name' has been deprecated from version 1.114.0. Use 'resource_group_name' instead. */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return Field `name` has been deprecated from version 1.114.0. Use `resource_group_name` instead.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The status of the resource group in all regions. See `region_statuses` below.
     * 
     */
    @Export(name="regionStatuses", refs={List.class,ResourceGroupRegionStatus.class}, tree="[0,1]")
    private Output<List<ResourceGroupRegionStatus>> regionStatuses;

    /**
     * @return The status of the resource group in all regions. See `region_statuses` below.
     * 
     */
    public Output<List<ResourceGroupRegionStatus>> regionStatuses() {
        return this.regionStatuses;
    }
    /**
     * The unique identifier of the resource group.The identifier must be 3 to 12 characters in length and can contain letters, digits, periods (.), hyphens (-), and underscores (_). The identifier must start with a letter.
     * 
     */
    @Export(name="resourceGroupName", refs={String.class}, tree="[0]")
    private Output<String> resourceGroupName;

    /**
     * @return The unique identifier of the resource group.The identifier must be 3 to 12 characters in length and can contain letters, digits, periods (.), hyphens (-), and underscores (_). The identifier must start with a letter.
     * 
     */
    public Output<String> resourceGroupName() {
        return this.resourceGroupName;
    }
    /**
     * The status of the regional resource group.
     * 
     */
    @Export(name="status", refs={String.class}, tree="[0]")
    private Output<String> status;

    /**
     * @return The status of the regional resource group.
     * 
     */
    public Output<String> status() {
        return this.status;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public ResourceGroup(String name) {
        this(name, ResourceGroupArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ResourceGroup(String name, ResourceGroupArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ResourceGroup(String name, ResourceGroupArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:resourcemanager/resourceGroup:ResourceGroup", name, args == null ? ResourceGroupArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private ResourceGroup(String name, Output<String> id, @Nullable ResourceGroupState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:resourcemanager/resourceGroup:ResourceGroup", name, state, makeResourceOptions(options, id));
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
    public static ResourceGroup get(String name, Output<String> id, @Nullable ResourceGroupState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ResourceGroup(name, id, state, options);
    }
}
