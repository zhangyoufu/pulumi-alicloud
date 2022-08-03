// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.slb;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.slb.TlsCipherPolicyArgs;
import com.pulumi.alicloud.slb.inputs.TlsCipherPolicyState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.List;
import javax.annotation.Nullable;

/**
 * Provides a SLB Tls Cipher Policy resource.
 * 
 * For information about SLB Tls Cipher Policy and how to use it, see [What is Tls Cipher Policy](https://www.alibabacloud.com/help/doc-detail/196714.htm).
 * 
 * &gt; **NOTE:** Available in v1.135.0+.
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
 * import com.pulumi.alicloud.slb.TlsCipherPolicy;
 * import com.pulumi.alicloud.slb.TlsCipherPolicyArgs;
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
 *         var example = new TlsCipherPolicy(&#34;example&#34;, TlsCipherPolicyArgs.builder()        
 *             .ciphers(            
 *                 &#34;AES256-SHA256&#34;,
 *                 &#34;AES128-GCM-SHA256&#34;)
 *             .tlsCipherPolicyName(&#34;Test-example_value&#34;)
 *             .tlsVersions(&#34;TLSv1.2&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * SLB Tls Cipher Policy can be imported using the id, e.g.
 * 
 * ```sh
 *  $ pulumi import alicloud:slb/tlsCipherPolicy:TlsCipherPolicy example &lt;id&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:slb/tlsCipherPolicy:TlsCipherPolicy")
public class TlsCipherPolicy extends com.pulumi.resources.CustomResource {
    /**
     * The encryption algorithms supported. It depends on the value of `tls_versions`.
     * 
     */
    @Export(name="ciphers", type=List.class, parameters={String.class})
    private Output<List<String>> ciphers;

    /**
     * @return The encryption algorithms supported. It depends on the value of `tls_versions`.
     * 
     */
    public Output<List<String>> ciphers() {
        return this.ciphers;
    }
    /**
     * TLS policy instance state.
     * 
     */
    @Export(name="status", type=String.class, parameters={})
    private Output<String> status;

    /**
     * @return TLS policy instance state.
     * 
     */
    public Output<String> status() {
        return this.status;
    }
    /**
     * TLS policy name. Length is from 2 to 128, or in both the English and Chinese characters must be with an uppercase/lowercase letter or a Chinese character and the beginning, may contain numbers, in dot `.`, underscore `_` or dash `-`.
     * 
     */
    @Export(name="tlsCipherPolicyName", type=String.class, parameters={})
    private Output<String> tlsCipherPolicyName;

    /**
     * @return TLS policy name. Length is from 2 to 128, or in both the English and Chinese characters must be with an uppercase/lowercase letter or a Chinese character and the beginning, may contain numbers, in dot `.`, underscore `_` or dash `-`.
     * 
     */
    public Output<String> tlsCipherPolicyName() {
        return this.tlsCipherPolicyName;
    }
    /**
     * The version of TLS protocol. You can find the corresponding value description in the document center [What is Tls Cipher Policy](https://www.alibabacloud.com/help/doc-detail/196714.htm).
     * 
     */
    @Export(name="tlsVersions", type=List.class, parameters={String.class})
    private Output<List<String>> tlsVersions;

    /**
     * @return The version of TLS protocol. You can find the corresponding value description in the document center [What is Tls Cipher Policy](https://www.alibabacloud.com/help/doc-detail/196714.htm).
     * 
     */
    public Output<List<String>> tlsVersions() {
        return this.tlsVersions;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public TlsCipherPolicy(String name) {
        this(name, TlsCipherPolicyArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public TlsCipherPolicy(String name, TlsCipherPolicyArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public TlsCipherPolicy(String name, TlsCipherPolicyArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:slb/tlsCipherPolicy:TlsCipherPolicy", name, args == null ? TlsCipherPolicyArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private TlsCipherPolicy(String name, Output<String> id, @Nullable TlsCipherPolicyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:slb/tlsCipherPolicy:TlsCipherPolicy", name, state, makeResourceOptions(options, id));
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
    public static TlsCipherPolicy get(String name, Output<String> id, @Nullable TlsCipherPolicyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new TlsCipherPolicy(name, id, state, options);
    }
}
