// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.dcdn;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.dcdn.WafDomainArgs;
import com.pulumi.alicloud.dcdn.inputs.WafDomainState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a DCDN Waf Domain resource.
 * 
 * For information about DCDN Waf Domain and how to use it, see [What is Waf Domain](https://www.alibabacloud.com/help/en/dcdn/developer-reference/api-dcdn-2018-01-15-batchsetdcdnwafdomainconfigs).
 * 
 * &gt; **NOTE:** Available since v1.185.0.
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
 * import com.pulumi.alicloud.dcdn.Domain;
 * import com.pulumi.alicloud.dcdn.DomainArgs;
 * import com.pulumi.alicloud.dcdn.inputs.DomainSourceArgs;
 * import com.pulumi.alicloud.dcdn.WafDomain;
 * import com.pulumi.alicloud.dcdn.WafDomainArgs;
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
 *         final var domainName = config.get("domainName").orElse("tf-example.com");
 *         var default_ = new Integer("default", IntegerArgs.builder()        
 *             .min(10000)
 *             .max(99999)
 *             .build());
 * 
 *         var example = new Domain("example", DomainArgs.builder()        
 *             .domainName(String.format("%s-%s", domainName,default_.result()))
 *             .scope("overseas")
 *             .sources(DomainSourceArgs.builder()
 *                 .content("1.1.1.1")
 *                 .port("80")
 *                 .priority("20")
 *                 .type("ipaddr")
 *                 .weight("10")
 *                 .build())
 *             .build());
 * 
 *         var exampleWafDomain = new WafDomain("exampleWafDomain", WafDomainArgs.builder()        
 *             .domainName(example.domainName())
 *             .clientIpTag("X-Forwarded-For")
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
 * DCDN Waf Domain can be imported using the id, e.g.
 * 
 * ```sh
 * $ pulumi import alicloud:dcdn/wafDomain:WafDomain example &lt;domain_name&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:dcdn/wafDomain:WafDomain")
public class WafDomain extends com.pulumi.resources.CustomResource {
    /**
     * The client ip tag.
     * 
     */
    @Export(name="clientIpTag", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> clientIpTag;

    /**
     * @return The client ip tag.
     * 
     */
    public Output<Optional<String>> clientIpTag() {
        return Codegen.optional(this.clientIpTag);
    }
    /**
     * The accelerated domain name.
     * 
     */
    @Export(name="domainName", refs={String.class}, tree="[0]")
    private Output<String> domainName;

    /**
     * @return The accelerated domain name.
     * 
     */
    public Output<String> domainName() {
        return this.domainName;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public WafDomain(String name) {
        this(name, WafDomainArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public WafDomain(String name, WafDomainArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public WafDomain(String name, WafDomainArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:dcdn/wafDomain:WafDomain", name, args == null ? WafDomainArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private WafDomain(String name, Output<String> id, @Nullable WafDomainState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:dcdn/wafDomain:WafDomain", name, state, makeResourceOptions(options, id));
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
    public static WafDomain get(String name, Output<String> id, @Nullable WafDomainState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new WafDomain(name, id, state, options);
    }
}
