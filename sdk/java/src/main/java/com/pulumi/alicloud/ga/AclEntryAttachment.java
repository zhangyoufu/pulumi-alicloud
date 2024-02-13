// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.ga;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.ga.AclEntryAttachmentArgs;
import com.pulumi.alicloud.ga.inputs.AclEntryAttachmentState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a Global Accelerator (GA) Acl entry attachment resource.
 * 
 * For information about Global Accelerator (GA) Acl entry attachment and how to use it, see [What is Acl entry attachment](https://www.alibabacloud.com/help/en/global-accelerator/latest/api-ga-2019-11-20-addentriestoacl).
 * 
 * &gt; **NOTE:** Available since v1.190.0.
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
 * import com.pulumi.alicloud.ga.Acl;
 * import com.pulumi.alicloud.ga.AclArgs;
 * import com.pulumi.alicloud.ga.AclEntryAttachment;
 * import com.pulumi.alicloud.ga.AclEntryAttachmentArgs;
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
 *         var defaultAcl = new Acl(&#34;defaultAcl&#34;, AclArgs.builder()        
 *             .aclName(&#34;tf-example-value&#34;)
 *             .addressIpVersion(&#34;IPv4&#34;)
 *             .build());
 * 
 *         var defaultAclEntryAttachment = new AclEntryAttachment(&#34;defaultAclEntryAttachment&#34;, AclEntryAttachmentArgs.builder()        
 *             .aclId(defaultAcl.id())
 *             .entry(&#34;192.168.1.1/32&#34;)
 *             .entryDescription(&#34;tf-example-value&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * Global Accelerator (GA) Acl entry attachment can be imported using the id.Format to `&lt;acl_id&gt;:&lt;entry&gt;`, e.g.
 * 
 * ```sh
 * $ pulumi import alicloud:ga/aclEntryAttachment:AclEntryAttachment example your_acl_id:your_entry
 * ```
 * 
 */
@ResourceType(type="alicloud:ga/aclEntryAttachment:AclEntryAttachment")
public class AclEntryAttachment extends com.pulumi.resources.CustomResource {
    /**
     * The ID of the global acceleration instance.
     * 
     */
    @Export(name="aclId", refs={String.class}, tree="[0]")
    private Output<String> aclId;

    /**
     * @return The ID of the global acceleration instance.
     * 
     */
    public Output<String> aclId() {
        return this.aclId;
    }
    /**
     * The IP address(192.168.XX.XX) or CIDR(10.0.XX.XX/24) block that you want to add to the network ACL.
     * 
     */
    @Export(name="entry", refs={String.class}, tree="[0]")
    private Output<String> entry;

    /**
     * @return The IP address(192.168.XX.XX) or CIDR(10.0.XX.XX/24) block that you want to add to the network ACL.
     * 
     */
    public Output<String> entry() {
        return this.entry;
    }
    /**
     * The description of the entry. The description must be 1 to 256 characters in length, and can contain letters, digits, hyphens (-), forward slashes (/), periods (.), and underscores (_).
     * 
     */
    @Export(name="entryDescription", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> entryDescription;

    /**
     * @return The description of the entry. The description must be 1 to 256 characters in length, and can contain letters, digits, hyphens (-), forward slashes (/), periods (.), and underscores (_).
     * 
     */
    public Output<Optional<String>> entryDescription() {
        return Codegen.optional(this.entryDescription);
    }
    /**
     * The status of the network ACL.
     * 
     */
    @Export(name="status", refs={String.class}, tree="[0]")
    private Output<String> status;

    /**
     * @return The status of the network ACL.
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
        super("alicloud:ga/aclEntryAttachment:AclEntryAttachment", name, args == null ? AclEntryAttachmentArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private AclEntryAttachment(String name, Output<String> id, @Nullable AclEntryAttachmentState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:ga/aclEntryAttachment:AclEntryAttachment", name, state, makeResourceOptions(options, id));
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
