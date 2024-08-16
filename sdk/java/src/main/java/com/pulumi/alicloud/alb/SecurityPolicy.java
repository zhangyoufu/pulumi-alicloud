// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.alb;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.alb.SecurityPolicyArgs;
import com.pulumi.alicloud.alb.inputs.SecurityPolicyState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a ALB Security Policy resource.
 * 
 * For information about ALB Security Policy and how to use it, see [What is Security Policy](https://www.alibabacloud.com/help/en/slb/application-load-balancer/developer-reference/api-alb-2020-06-16-createsecuritypolicy).
 * 
 * &gt; **NOTE:** Available since v1.130.0.
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
 * import com.pulumi.alicloud.alb.SecurityPolicy;
 * import com.pulumi.alicloud.alb.SecurityPolicyArgs;
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
 *         var default_ = new SecurityPolicy("default", SecurityPolicyArgs.builder()
 *             .securityPolicyName("tf_example")
 *             .tlsVersions("TLSv1.0")
 *             .ciphers(            
 *                 "ECDHE-ECDSA-AES128-SHA",
 *                 "AES256-SHA")
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
 * ALB Security Policy can be imported using the id, e.g.
 * 
 * ```sh
 * $ pulumi import alicloud:alb/securityPolicy:SecurityPolicy example &lt;id&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:alb/securityPolicy:SecurityPolicy")
public class SecurityPolicy extends com.pulumi.resources.CustomResource {
    /**
     * The supported cipher suites, which are determined by the TLS protocol version.The specified cipher suites must be supported by at least one TLS protocol version that you select.
     * 
     */
    @Export(name="ciphers", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> ciphers;

    /**
     * @return The supported cipher suites, which are determined by the TLS protocol version.The specified cipher suites must be supported by at least one TLS protocol version that you select.
     * 
     */
    public Output<List<String>> ciphers() {
        return this.ciphers;
    }
    /**
     * The dry run.
     * 
     */
    @Export(name="dryRun", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> dryRun;

    /**
     * @return The dry run.
     * 
     */
    public Output<Optional<Boolean>> dryRun() {
        return Codegen.optional(this.dryRun);
    }
    /**
     * The ID of the resource group.
     * 
     */
    @Export(name="resourceGroupId", refs={String.class}, tree="[0]")
    private Output<String> resourceGroupId;

    /**
     * @return The ID of the resource group.
     * 
     */
    public Output<String> resourceGroupId() {
        return this.resourceGroupId;
    }
    /**
     * The name of the resource. The name must be 2 to 128 characters in length and must start with a letter. It can contain digits, periods (.), underscores (_), and hyphens (-).
     * 
     */
    @Export(name="securityPolicyName", refs={String.class}, tree="[0]")
    private Output<String> securityPolicyName;

    /**
     * @return The name of the resource. The name must be 2 to 128 characters in length and must start with a letter. It can contain digits, periods (.), underscores (_), and hyphens (-).
     * 
     */
    public Output<String> securityPolicyName() {
        return this.securityPolicyName;
    }
    /**
     * The status of the resource.
     * 
     */
    @Export(name="status", refs={String.class}, tree="[0]")
    private Output<String> status;

    /**
     * @return The status of the resource.
     * 
     */
    public Output<String> status() {
        return this.status;
    }
    /**
     * A mapping of tags to assign to the resource.
     * 
     */
    @Export(name="tags", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output</* @Nullable */ Map<String,String>> tags;

    /**
     * @return A mapping of tags to assign to the resource.
     * 
     */
    public Output<Optional<Map<String,String>>> tags() {
        return Codegen.optional(this.tags);
    }
    /**
     * The TLS protocol versions that are supported. Valid values: TLSv1.0, TLSv1.1, TLSv1.2 and TLSv1.3.
     * 
     */
    @Export(name="tlsVersions", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> tlsVersions;

    /**
     * @return The TLS protocol versions that are supported. Valid values: TLSv1.0, TLSv1.1, TLSv1.2 and TLSv1.3.
     * 
     */
    public Output<List<String>> tlsVersions() {
        return this.tlsVersions;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public SecurityPolicy(java.lang.String name) {
        this(name, SecurityPolicyArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public SecurityPolicy(java.lang.String name, SecurityPolicyArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public SecurityPolicy(java.lang.String name, SecurityPolicyArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:alb/securityPolicy:SecurityPolicy", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private SecurityPolicy(java.lang.String name, Output<java.lang.String> id, @Nullable SecurityPolicyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:alb/securityPolicy:SecurityPolicy", name, state, makeResourceOptions(options, id), false);
    }

    private static SecurityPolicyArgs makeArgs(SecurityPolicyArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? SecurityPolicyArgs.Empty : args;
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
    public static SecurityPolicy get(java.lang.String name, Output<java.lang.String> id, @Nullable SecurityPolicyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new SecurityPolicy(name, id, state, options);
    }
}
