// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.ga;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.ga.AclArgs;
import com.pulumi.alicloud.ga.inputs.AclState;
import com.pulumi.alicloud.ga.outputs.AclAclEntry;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a Global Accelerator (GA) Acl resource.
 * 
 * For information about Global Accelerator (GA) Acl and how to use it, see [What is Acl](https://www.alibabacloud.com/help/en/global-accelerator/latest/api-doc-ga-2019-11-20-api-doc-createacl).
 * 
 * &gt; **NOTE:** Available in v1.150.0+.
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
 * import com.pulumi.alicloud.ga.inputs.AclAclEntryArgs;
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
 *         var default_ = new Acl(&#34;default&#34;, AclArgs.builder()        
 *             .aclEntries(AclAclEntryArgs.builder()
 *                 .entry(&#34;192.168.1.0/24&#34;)
 *                 .entryDescription(&#34;tf-test1&#34;)
 *                 .build())
 *             .aclName(&#34;tf-testAccAcl&#34;)
 *             .addressIpVersion(&#34;IPv4&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * Global Accelerator (GA) Acl can be imported using the id, e.g.
 * 
 * ```sh
 *  $ pulumi import alicloud:ga/acl:Acl example &lt;id&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:ga/acl:Acl")
public class Acl extends com.pulumi.resources.CustomResource {
    /**
     * The entries of the Acl. See the following `Block acl_entries`. **NOTE:** &#34;Field &#39;acl_entries&#39; has been deprecated from provider version 1.190.0 and it will be removed in the future version. Please use the new resource `alicloud.ga.AclEntryAttachment`.&#34;
     * 
     * @deprecated
     * Field &#39;acl_entries&#39; has been deprecated from provider version 1.190.0 and it will be removed in the future version. Please use the new resource &#39;alicloud_ga_acl_entry_attachment&#39;.
     * 
     */
    @Deprecated /* Field 'acl_entries' has been deprecated from provider version 1.190.0 and it will be removed in the future version. Please use the new resource 'alicloud_ga_acl_entry_attachment'. */
    @Export(name="aclEntries", type=List.class, parameters={AclAclEntry.class})
    private Output<List<AclAclEntry>> aclEntries;

    /**
     * @return The entries of the Acl. See the following `Block acl_entries`. **NOTE:** &#34;Field &#39;acl_entries&#39; has been deprecated from provider version 1.190.0 and it will be removed in the future version. Please use the new resource `alicloud.ga.AclEntryAttachment`.&#34;
     * 
     */
    public Output<List<AclAclEntry>> aclEntries() {
        return this.aclEntries;
    }
    /**
     * The name of the ACL. The name must be `2` to `128` characters in length, and can contain letters, digits, periods (.), hyphens (-) and underscores (_). It must start with a letter.
     * 
     */
    @Export(name="aclName", type=String.class, parameters={})
    private Output</* @Nullable */ String> aclName;

    /**
     * @return The name of the ACL. The name must be `2` to `128` characters in length, and can contain letters, digits, periods (.), hyphens (-) and underscores (_). It must start with a letter.
     * 
     */
    public Output<Optional<String>> aclName() {
        return Codegen.optional(this.aclName);
    }
    /**
     * The IP version. Valid values: `IPv4` and `IPv6`.
     * 
     */
    @Export(name="addressIpVersion", type=String.class, parameters={})
    private Output<String> addressIpVersion;

    /**
     * @return The IP version. Valid values: `IPv4` and `IPv6`.
     * 
     */
    public Output<String> addressIpVersion() {
        return this.addressIpVersion;
    }
    /**
     * The dry run.
     * 
     */
    @Export(name="dryRun", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> dryRun;

    /**
     * @return The dry run.
     * 
     */
    public Output<Optional<Boolean>> dryRun() {
        return Codegen.optional(this.dryRun);
    }
    /**
     * The status of the resource.
     * 
     */
    @Export(name="status", type=String.class, parameters={})
    private Output<String> status;

    /**
     * @return The status of the resource.
     * 
     */
    public Output<String> status() {
        return this.status;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Acl(String name) {
        this(name, AclArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Acl(String name, AclArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Acl(String name, AclArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:ga/acl:Acl", name, args == null ? AclArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Acl(String name, Output<String> id, @Nullable AclState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:ga/acl:Acl", name, state, makeResourceOptions(options, id));
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
    public static Acl get(String name, Output<String> id, @Nullable AclState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Acl(name, id, state, options);
    }
}
