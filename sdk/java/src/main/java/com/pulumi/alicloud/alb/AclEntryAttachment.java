// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.alb;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.alb.AclEntryAttachmentArgs;
import com.pulumi.alicloud.alb.inputs.AclEntryAttachmentState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * For information about acl entry attachment and how to use it, see [Configure an acl entry](https://www.alibabacloud.com/help/en/slb/application-load-balancer/developer-reference/api-alb-2020-06-16-addentriestoacl).
 * 
 * &gt; **NOTE:** Available since v1.166.0.
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.alicloud.resourcemanager.ResourcemanagerFunctions;
 * import com.pulumi.alicloud.resourcemanager.inputs.GetResourceGroupsArgs;
 * import com.pulumi.alicloud.alb.Acl;
 * import com.pulumi.alicloud.alb.AclArgs;
 * import com.pulumi.alicloud.alb.AclEntryAttachment;
 * import com.pulumi.alicloud.alb.AclEntryAttachmentArgs;
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
 *         final var name = config.get(&#34;name&#34;).orElse(&#34;tf_example&#34;);
 *         final var defaultResourceGroups = ResourcemanagerFunctions.getResourceGroups();
 * 
 *         var defaultAcl = new Acl(&#34;defaultAcl&#34;, AclArgs.builder()        
 *             .aclName(name)
 *             .resourceGroupId(defaultResourceGroups.applyValue(getResourceGroupsResult -&gt; getResourceGroupsResult.groups()[0].id()))
 *             .build());
 * 
 *         var defaultAclEntryAttachment = new AclEntryAttachment(&#34;defaultAclEntryAttachment&#34;, AclEntryAttachmentArgs.builder()        
 *             .aclId(defaultAcl.id())
 *             .entry(&#34;168.10.10.0/24&#34;)
 *             .description(name)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * Acl entry attachment can be imported using the id, e.g.
 * 
 * ```sh
 * $ pulumi import alicloud:alb/aclEntryAttachment:AclEntryAttachment example &lt;acl_id&gt;:&lt;entry&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:alb/aclEntryAttachment:AclEntryAttachment")
public class AclEntryAttachment extends com.pulumi.resources.CustomResource {
    /**
     * The ID of the Acl.
     * 
     */
    @Export(name="aclId", refs={String.class}, tree="[0]")
    private Output<String> aclId;

    /**
     * @return The ID of the Acl.
     * 
     */
    public Output<String> aclId() {
        return this.aclId;
    }
    /**
     * The description of the entry.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return The description of the entry.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * The CIDR blocks.
     * 
     */
    @Export(name="entry", refs={String.class}, tree="[0]")
    private Output<String> entry;

    /**
     * @return The CIDR blocks.
     * 
     */
    public Output<String> entry() {
        return this.entry;
    }
    /**
     * The Status of the resource.
     * 
     */
    @Export(name="status", refs={String.class}, tree="[0]")
    private Output<String> status;

    /**
     * @return The Status of the resource.
     * 
     */
    public Output<String> status() {
        return this.status;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public AclEntryAttachment(String name) {
        this(name, AclEntryAttachmentArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public AclEntryAttachment(String name, AclEntryAttachmentArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public AclEntryAttachment(String name, AclEntryAttachmentArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:alb/aclEntryAttachment:AclEntryAttachment", name, args == null ? AclEntryAttachmentArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private AclEntryAttachment(String name, Output<String> id, @Nullable AclEntryAttachmentState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:alb/aclEntryAttachment:AclEntryAttachment", name, state, makeResourceOptions(options, id));
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
    public static AclEntryAttachment get(String name, Output<String> id, @Nullable AclEntryAttachmentState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new AclEntryAttachment(name, id, state, options);
    }
}
