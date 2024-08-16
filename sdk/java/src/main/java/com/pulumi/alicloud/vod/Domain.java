// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.vod;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.vod.DomainArgs;
import com.pulumi.alicloud.vod.inputs.DomainState;
import com.pulumi.alicloud.vod.outputs.DomainSource;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a VOD Domain resource.
 * 
 * For information about VOD Domain and how to use it, see [What is Domain](https://www.alibabacloud.com/help/product/29932.html).
 * 
 * &gt; **NOTE:** Available since v1.136.0+.
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
 * import com.pulumi.alicloud.vod.Domain;
 * import com.pulumi.alicloud.vod.DomainArgs;
 * import com.pulumi.alicloud.vod.inputs.DomainSourceArgs;
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
 *         var default_ = new Integer("default", IntegerArgs.builder()
 *             .min(10000)
 *             .max(99999)
 *             .build());
 * 
 *         var defaultDomain = new Domain("defaultDomain", DomainArgs.builder()
 *             .domainName(String.format("example-%s.com", default_.result()))
 *             .scope("domestic")
 *             .sources(DomainSourceArgs.builder()
 *                 .sourceType("domain")
 *                 .sourceContent("outin-c7405446108111ec9a7100163e0eb78b.oss-cn-beijing.aliyuncs.com")
 *                 .sourcePort("443")
 *                 .build())
 *             .tags(Map.ofEntries(
 *                 Map.entry("Created", "terraform"),
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
 * VOD Domain can be imported using the id, e.g.
 * 
 * ```sh
 * $ pulumi import alicloud:vod/domain:Domain example &lt;domain_name&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:vod/domain:Domain")
public class Domain extends com.pulumi.resources.CustomResource {
    /**
     * The name of the certificate. The value of this parameter is returned if HTTPS is enabled.
     * 
     */
    @Export(name="certName", refs={String.class}, tree="[0]")
    private Output<String> certName;

    /**
     * @return The name of the certificate. The value of this parameter is returned if HTTPS is enabled.
     * 
     */
    public Output<String> certName() {
        return this.certName;
    }
    /**
     * The URL that is used for health checks.
     * 
     */
    @Export(name="checkUrl", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> checkUrl;

    /**
     * @return The URL that is used for health checks.
     * 
     */
    public Output<Optional<String>> checkUrl() {
        return Codegen.optional(this.checkUrl);
    }
    /**
     * The CNAME that is assigned to the domain name for CDN. You must add a CNAME record in the system of your Domain Name System (DNS) service provider to map the domain name for CDN to the CNAME.
     * 
     */
    @Export(name="cname", refs={String.class}, tree="[0]")
    private Output<String> cname;

    /**
     * @return The CNAME that is assigned to the domain name for CDN. You must add a CNAME record in the system of your Domain Name System (DNS) service provider to map the domain name for CDN to the CNAME.
     * 
     */
    public Output<String> cname() {
        return this.cname;
    }
    /**
     * The description of the domain name for CDN.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output<String> description;

    /**
     * @return The description of the domain name for CDN.
     * 
     */
    public Output<String> description() {
        return this.description;
    }
    /**
     * The domain name for CDN that you want to add to ApsaraVideo VOD. Wildcard domain names are supported. Start the domain name with a period (.). Example: `.example.com.`.
     * 
     */
    @Export(name="domainName", refs={String.class}, tree="[0]")
    private Output<String> domainName;

    /**
     * @return The domain name for CDN that you want to add to ApsaraVideo VOD. Wildcard domain names are supported. Start the domain name with a period (.). Example: `.example.com.`.
     * 
     */
    public Output<String> domainName() {
        return this.domainName;
    }
    /**
     * The time when the domain name for CDN was added. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
     * 
     */
    @Export(name="gmtCreated", refs={String.class}, tree="[0]")
    private Output<String> gmtCreated;

    /**
     * @return The time when the domain name for CDN was added. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
     * 
     */
    public Output<String> gmtCreated() {
        return this.gmtCreated;
    }
    /**
     * The last time when the domain name for CDN was modified. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
     * 
     */
    @Export(name="gmtModified", refs={String.class}, tree="[0]")
    private Output<String> gmtModified;

    /**
     * @return The last time when the domain name for CDN was modified. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
     * 
     */
    public Output<String> gmtModified() {
        return this.gmtModified;
    }
    /**
     * This parameter is applicable to users of level 3 or higher in mainland China and users outside mainland China. Valid values:
     * 
     */
    @Export(name="scope", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> scope;

    /**
     * @return This parameter is applicable to users of level 3 or higher in mainland China and users outside mainland China. Valid values:
     * 
     */
    public Output<Optional<String>> scope() {
        return Codegen.optional(this.scope);
    }
    /**
     * The information about the address of the origin server. For more information about the Sources parameter, See the following `Block sources`.
     * 
     */
    @Export(name="sources", refs={List.class,DomainSource.class}, tree="[0,1]")
    private Output<List<DomainSource>> sources;

    /**
     * @return The information about the address of the origin server. For more information about the Sources parameter, See the following `Block sources`.
     * 
     */
    public Output<List<DomainSource>> sources() {
        return this.sources;
    }
    /**
     * Indicates whether the Secure Sockets Layer (SSL) certificate is enabled. Valid values: `on`,`off`.
     * 
     */
    @Export(name="sslProtocol", refs={String.class}, tree="[0]")
    private Output<String> sslProtocol;

    /**
     * @return Indicates whether the Secure Sockets Layer (SSL) certificate is enabled. Valid values: `on`,`off`.
     * 
     */
    public Output<String> sslProtocol() {
        return this.sslProtocol;
    }
    /**
     * The public key of the certificate. The value of this parameter is returned if HTTPS is enabled.
     * 
     */
    @Export(name="sslPub", refs={String.class}, tree="[0]")
    private Output<String> sslPub;

    /**
     * @return The public key of the certificate. The value of this parameter is returned if HTTPS is enabled.
     * 
     */
    public Output<String> sslPub() {
        return this.sslPub;
    }
    /**
     * The status of the domain name for CDN. Valid values:
     * 
     */
    @Export(name="status", refs={String.class}, tree="[0]")
    private Output<String> status;

    /**
     * @return The status of the domain name for CDN. Valid values:
     * 
     */
    public Output<String> status() {
        return this.status;
    }
    /**
     * A mapping of tags to assign to the resource.
     * * `Key`: It can be up to 64 characters in length. It cannot be a null string.
     * * `Value`: It can be up to 128 characters in length. It can be a null string.
     * 
     */
    @Export(name="tags", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output</* @Nullable */ Map<String,String>> tags;

    /**
     * @return A mapping of tags to assign to the resource.
     * * `Key`: It can be up to 64 characters in length. It cannot be a null string.
     * * `Value`: It can be up to 128 characters in length. It can be a null string.
     * 
     */
    public Output<Optional<Map<String,String>>> tags() {
        return Codegen.optional(this.tags);
    }
    /**
     * The top-level domain name.
     * 
     */
    @Export(name="topLevelDomain", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> topLevelDomain;

    /**
     * @return The top-level domain name.
     * 
     */
    public Output<Optional<String>> topLevelDomain() {
        return Codegen.optional(this.topLevelDomain);
    }
    /**
     * The weight of the origin server.
     * 
     */
    @Export(name="weight", refs={String.class}, tree="[0]")
    private Output<String> weight;

    /**
     * @return The weight of the origin server.
     * 
     */
    public Output<String> weight() {
        return this.weight;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Domain(java.lang.String name) {
        this(name, DomainArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Domain(java.lang.String name, DomainArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Domain(java.lang.String name, DomainArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:vod/domain:Domain", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private Domain(java.lang.String name, Output<java.lang.String> id, @Nullable DomainState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:vod/domain:Domain", name, state, makeResourceOptions(options, id), false);
    }

    private static DomainArgs makeArgs(DomainArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? DomainArgs.Empty : args;
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
    public static Domain get(java.lang.String name, Output<java.lang.String> id, @Nullable DomainState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Domain(name, id, state, options);
    }
}
