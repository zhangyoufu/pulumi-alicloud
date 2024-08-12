// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.dns;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.dns.AlidnsDomainArgs;
import com.pulumi.alicloud.dns.inputs.AlidnsDomainState;
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
 * Provides a Alidns domain resource.
 * 
 * &gt; **NOTE:** The domain name which you want to add must be already registered and had not added by another account. Every domain name can only exist in a unique group.
 * 
 * &gt; **NOTE:** Available since v1.95.0.
 * 
 * ## Example Usage
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.alicloud.dns.DomainGroup;
 * import com.pulumi.alicloud.dns.DomainGroupArgs;
 * import com.pulumi.alicloud.dns.AlidnsDomain;
 * import com.pulumi.alicloud.dns.AlidnsDomainArgs;
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
 *         var default_ = new DomainGroup("default", DomainGroupArgs.builder()
 *             .domainGroupName("tf-example")
 *             .build());
 * 
 *         var defaultAlidnsDomain = new AlidnsDomain("defaultAlidnsDomain", AlidnsDomainArgs.builder()
 *             .domainName("starmove.com")
 *             .groupId(default_.id())
 *             .tags(Map.ofEntries(
 *                 Map.entry("Created", "TF"),
 *                 Map.entry("For", "example")
 *             ))
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
 * Alidns domain can be imported using the id or domain name, e.g.
 * 
 * ```sh
 * $ pulumi import alicloud:dns/alidnsDomain:AlidnsDomain example aliyun.com
 * ```
 * 
 */
@ResourceType(type="alicloud:dns/alidnsDomain:AlidnsDomain")
public class AlidnsDomain extends com.pulumi.resources.CustomResource {
    /**
     * A list of the dns server name.
     * 
     */
    @Export(name="dnsServers", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> dnsServers;

    /**
     * @return A list of the dns server name.
     * 
     */
    public Output<List<String>> dnsServers() {
        return this.dnsServers;
    }
    /**
     * The domain ID.
     * 
     */
    @Export(name="domainId", refs={String.class}, tree="[0]")
    private Output<String> domainId;

    /**
     * @return The domain ID.
     * 
     */
    public Output<String> domainId() {
        return this.domainId;
    }
    /**
     * Name of the domain. This name without suffix can have a string of 1 to 63 characters(domain name subject, excluding suffix), must contain only alphanumeric characters or &#34;-&#34;, and must not begin or end with &#34;-&#34;, and &#34;-&#34; must not in the 3th and 4th character positions at the same time. Suffix `.sh` and `.tel` are not supported.
     * 
     */
    @Export(name="domainName", refs={String.class}, tree="[0]")
    private Output<String> domainName;

    /**
     * @return Name of the domain. This name without suffix can have a string of 1 to 63 characters(domain name subject, excluding suffix), must contain only alphanumeric characters or &#34;-&#34;, and must not begin or end with &#34;-&#34;, and &#34;-&#34; must not in the 3th and 4th character positions at the same time. Suffix `.sh` and `.tel` are not supported.
     * 
     */
    public Output<String> domainName() {
        return this.domainName;
    }
    /**
     * Id of the group in which the domain will add. If not supplied, then use default group.
     * 
     */
    @Export(name="groupId", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> groupId;

    /**
     * @return Id of the group in which the domain will add. If not supplied, then use default group.
     * 
     */
    public Output<Optional<String>> groupId() {
        return Codegen.optional(this.groupId);
    }
    /**
     * Domain name group name.
     * 
     */
    @Export(name="groupName", refs={String.class}, tree="[0]")
    private Output<String> groupName;

    /**
     * @return Domain name group name.
     * 
     */
    public Output<String> groupName() {
        return this.groupName;
    }
    /**
     * User language.
     * 
     */
    @Export(name="lang", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> lang;

    /**
     * @return User language.
     * 
     */
    public Output<Optional<String>> lang() {
        return Codegen.optional(this.lang);
    }
    /**
     * Only return punycode codes for Chinese domain names.
     * 
     */
    @Export(name="punyCode", refs={String.class}, tree="[0]")
    private Output<String> punyCode;

    /**
     * @return Only return punycode codes for Chinese domain names.
     * 
     */
    public Output<String> punyCode() {
        return this.punyCode;
    }
    /**
     * Remarks information for your domain name.
     * 
     */
    @Export(name="remark", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> remark;

    /**
     * @return Remarks information for your domain name.
     * 
     */
    public Output<Optional<String>> remark() {
        return Codegen.optional(this.remark);
    }
    /**
     * The Id of resource group which the dns domain belongs.
     * 
     */
    @Export(name="resourceGroupId", refs={String.class}, tree="[0]")
    private Output<String> resourceGroupId;

    /**
     * @return The Id of resource group which the dns domain belongs.
     * 
     */
    public Output<String> resourceGroupId() {
        return this.resourceGroupId;
    }
    /**
     * A mapping of tags to assign to the resource.
     * - Key: It can be [1, 20] characters in length. It can contain A-Z, a-z, numbers, underscores (_), and hyphens (-). It cannot begin with &#34;aliyun&#34;, &#34;acs:&#34;, &#34;http://&#34;, or &#34;https://&#34;. It cannot be a null string.
     * - Value: It can be [1, 20] characters in length. It can contain A-Z, a-z, numbers, underscores (_), and hyphens (-). It cannot begin with &#34;aliyun&#34;, &#34;acs:&#34;, &#34;http://&#34;, or &#34;https://&#34;. It can be a null string.
     * 
     */
    @Export(name="tags", refs={Map.class,String.class,Object.class}, tree="[0,1,2]")
    private Output</* @Nullable */ Map<String,Object>> tags;

    /**
     * @return A mapping of tags to assign to the resource.
     * - Key: It can be [1, 20] characters in length. It can contain A-Z, a-z, numbers, underscores (_), and hyphens (-). It cannot begin with &#34;aliyun&#34;, &#34;acs:&#34;, &#34;http://&#34;, or &#34;https://&#34;. It cannot be a null string.
     * - Value: It can be [1, 20] characters in length. It can contain A-Z, a-z, numbers, underscores (_), and hyphens (-). It cannot begin with &#34;aliyun&#34;, &#34;acs:&#34;, &#34;http://&#34;, or &#34;https://&#34;. It can be a null string.
     * 
     */
    public Output<Optional<Map<String,Object>>> tags() {
        return Codegen.optional(this.tags);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public AlidnsDomain(java.lang.String name) {
        this(name, AlidnsDomainArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public AlidnsDomain(java.lang.String name, AlidnsDomainArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public AlidnsDomain(java.lang.String name, AlidnsDomainArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:dns/alidnsDomain:AlidnsDomain", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private AlidnsDomain(java.lang.String name, Output<java.lang.String> id, @Nullable AlidnsDomainState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:dns/alidnsDomain:AlidnsDomain", name, state, makeResourceOptions(options, id), false);
    }

    private static AlidnsDomainArgs makeArgs(AlidnsDomainArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? AlidnsDomainArgs.Empty : args;
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
    public static AlidnsDomain get(java.lang.String name, Output<java.lang.String> id, @Nullable AlidnsDomainState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new AlidnsDomain(name, id, state, options);
    }
}
