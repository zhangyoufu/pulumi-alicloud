// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.slb;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.slb.AclArgs;
import com.pulumi.alicloud.slb.inputs.AclState;
import com.pulumi.alicloud.slb.outputs.AclEntryList;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Object;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * An access control list contains multiple IP addresses or CIDR blocks.
 * The access control list can help you to define multiple instance listening dimension,
 * and to meet the multiple usage for single access control list.
 * 
 * Server Load Balancer allows you to configure access control for listeners.
 * You can configure different whitelists or blacklists for different listeners.
 * 
 * You can configure access control
 * when you create a listener or change access control configuration after a listener is created.
 * 
 * &gt; **NOTE:** One access control list can be attached to many Listeners in different load balancer as whitelists or blacklists.
 * 
 * &gt; **NOTE:** The maximum number of access control lists per region  is 50.
 * 
 * &gt; **NOTE:** The maximum number of IP addresses added each time is 50.
 * 
 * &gt; **NOTE:** The maximum number of entries per access control list is 300.
 * 
 * &gt; **NOTE:** The maximum number of listeners that an access control list can be added to is 50.
 * 
 * For information about slb and how to use it, see [What is Server Load Balancer](https://www.alibabacloud.com/help/doc-detail/27539.htm).
 * 
 * For information about acl and how to use it, see [Configure an access control list](https://www.alibabacloud.com/help/doc-detail/70015.htm).
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.alicloud.slb.Acl;
 * import com.pulumi.alicloud.slb.AclArgs;
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
 *         var acl = new Acl(&#34;acl&#34;, AclArgs.builder()        
 *             .ipVersion(&#34;ipv4&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * ## Entry Block
 * 
 * The entry mapping supports the following:
 * 
 * * `entry` - (Optional, Computed) The CIDR blocks.
 * * `comment` - (Optional, Computed) The comment of the entry.
 * 
 * ## Import
 * 
 * Server Load balancer access control list can be imported using the id, e.g.
 * 
 * ```sh
 * $ pulumi import alicloud:slb/acl:Acl example acl-abc123456
 * ```
 * 
 */
@ResourceType(type="alicloud:slb/acl:Acl")
public class Acl extends com.pulumi.resources.CustomResource {
    /**
     * A list of entry (CIDR blocks) to be added. It contains two sub-fields as `Entry Block` follows. **NOTE:** &#34;Field &#39;entry_list&#39; has been deprecated from provider version 1.162.0 and it will be removed in the future version. Please use the new resource &#39;alicloud_slb_acl_entry_attachment&#39;.&#34;,
     * 
     * @deprecated
     * Field &#39;entry_list&#39; has been deprecated from provider version 1.162.0 and it will be removed in the future version. Please use the new resource &#39;alicloud_slb_acl_entry_attachment&#39;.
     * 
     */
    @Deprecated /* Field 'entry_list' has been deprecated from provider version 1.162.0 and it will be removed in the future version. Please use the new resource 'alicloud_slb_acl_entry_attachment'. */
    @Export(name="entryLists", refs={List.class,AclEntryList.class}, tree="[0,1]")
    private Output<List<AclEntryList>> entryLists;

    /**
     * @return A list of entry (CIDR blocks) to be added. It contains two sub-fields as `Entry Block` follows. **NOTE:** &#34;Field &#39;entry_list&#39; has been deprecated from provider version 1.162.0 and it will be removed in the future version. Please use the new resource &#39;alicloud_slb_acl_entry_attachment&#39;.&#34;,
     * 
     */
    public Output<List<AclEntryList>> entryLists() {
        return this.entryLists;
    }
    /**
     * The IP Version of access control list is the type of its entry (IP addresses or CIDR blocks). It values ipv4/ipv6. Our plugin provides a default ip_version: &#34;ipv4&#34;.
     * 
     */
    @Export(name="ipVersion", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> ipVersion;

    /**
     * @return The IP Version of access control list is the type of its entry (IP addresses or CIDR blocks). It values ipv4/ipv6. Our plugin provides a default ip_version: &#34;ipv4&#34;.
     * 
     */
    public Output<Optional<String>> ipVersion() {
        return Codegen.optional(this.ipVersion);
    }
    /**
     * Name of the access control list.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return Name of the access control list.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * Resource group ID.
     * 
     */
    @Export(name="resourceGroupId", refs={String.class}, tree="[0]")
    private Output<String> resourceGroupId;

    /**
     * @return Resource group ID.
     * 
     */
    public Output<String> resourceGroupId() {
        return this.resourceGroupId;
    }
    /**
     * A mapping of tags to assign to the resource.
     * 
     */
    @Export(name="tags", refs={Map.class,String.class,Object.class}, tree="[0,1,2]")
    private Output</* @Nullable */ Map<String,Object>> tags;

    /**
     * @return A mapping of tags to assign to the resource.
     * 
     */
    public Output<Optional<Map<String,Object>>> tags() {
        return Codegen.optional(this.tags);
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
    public Acl(String name, @Nullable AclArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Acl(String name, @Nullable AclArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:slb/acl:Acl", name, args == null ? AclArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Acl(String name, Output<String> id, @Nullable AclState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:slb/acl:Acl", name, state, makeResourceOptions(options, id));
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
