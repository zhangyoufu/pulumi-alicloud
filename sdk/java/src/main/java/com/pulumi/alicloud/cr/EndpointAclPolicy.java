// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.cr;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.cr.EndpointAclPolicyArgs;
import com.pulumi.alicloud.cr.inputs.EndpointAclPolicyState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a CR Endpoint Acl Policy resource.
 * 
 * For information about CR Endpoint Acl Policy and how to use it, see [What is Endpoint Acl Policy](https://www.alibabacloud.com/help/doc-detail/145275.htm).
 * 
 * &gt; **NOTE:** Available since v1.139.0.
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
 * import com.pulumi.alicloud.cr.RegistryEnterpriseInstance;
 * import com.pulumi.alicloud.cr.RegistryEnterpriseInstanceArgs;
 * import com.pulumi.alicloud.cr.CrFunctions;
 * import com.pulumi.alicloud.cr.inputs.GetEndpointAclServiceArgs;
 * import com.pulumi.alicloud.cr.EndpointAclPolicy;
 * import com.pulumi.alicloud.cr.EndpointAclPolicyArgs;
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
 *         final var name = config.get(&#34;name&#34;).orElse(&#34;tf-example&#34;);
 *         var defaultRegistryEnterpriseInstance = new RegistryEnterpriseInstance(&#34;defaultRegistryEnterpriseInstance&#34;, RegistryEnterpriseInstanceArgs.builder()        
 *             .paymentType(&#34;Subscription&#34;)
 *             .period(1)
 *             .renewalStatus(&#34;ManualRenewal&#34;)
 *             .instanceType(&#34;Advanced&#34;)
 *             .instanceName(name)
 *             .build());
 * 
 *         final var default = CrFunctions.getEndpointAclService(GetEndpointAclServiceArgs.builder()
 *             .endpointType(&#34;internet&#34;)
 *             .enable(true)
 *             .instanceId(defaultRegistryEnterpriseInstance.id())
 *             .moduleName(&#34;Registry&#34;)
 *             .build());
 * 
 *         var defaultEndpointAclPolicy = new EndpointAclPolicy(&#34;defaultEndpointAclPolicy&#34;, EndpointAclPolicyArgs.builder()        
 *             .instanceId(default_.applyValue(default_ -&gt; default_.instanceId()))
 *             .entry(&#34;192.168.1.0/24&#34;)
 *             .description(name)
 *             .moduleName(&#34;Registry&#34;)
 *             .endpointType(&#34;internet&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * CR Endpoint Acl Policy can be imported using the id, e.g.
 * 
 * ```sh
 * $ pulumi import alicloud:cr/endpointAclPolicy:EndpointAclPolicy example &lt;instance_id&gt;:&lt;endpoint_type&gt;:&lt;entry&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:cr/endpointAclPolicy:EndpointAclPolicy")
public class EndpointAclPolicy extends com.pulumi.resources.CustomResource {
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
     * The type of endpoint. Valid values: `internet`.
     * 
     */
    @Export(name="endpointType", refs={String.class}, tree="[0]")
    private Output<String> endpointType;

    /**
     * @return The type of endpoint. Valid values: `internet`.
     * 
     */
    public Output<String> endpointType() {
        return this.endpointType;
    }
    /**
     * The IP segment that allowed to access.
     * 
     */
    @Export(name="entry", refs={String.class}, tree="[0]")
    private Output<String> entry;

    /**
     * @return The IP segment that allowed to access.
     * 
     */
    public Output<String> entry() {
        return this.entry;
    }
    /**
     * The ID of the CR Instance.
     * 
     */
    @Export(name="instanceId", refs={String.class}, tree="[0]")
    private Output<String> instanceId;

    /**
     * @return The ID of the CR Instance.
     * 
     */
    public Output<String> instanceId() {
        return this.instanceId;
    }
    /**
     * The module that needs to set the access policy. Valid values: `Registry`.
     * 
     */
    @Export(name="moduleName", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> moduleName;

    /**
     * @return The module that needs to set the access policy. Valid values: `Registry`.
     * 
     */
    public Output<Optional<String>> moduleName() {
        return Codegen.optional(this.moduleName);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public EndpointAclPolicy(String name) {
        this(name, EndpointAclPolicyArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public EndpointAclPolicy(String name, EndpointAclPolicyArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public EndpointAclPolicy(String name, EndpointAclPolicyArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:cr/endpointAclPolicy:EndpointAclPolicy", name, args == null ? EndpointAclPolicyArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private EndpointAclPolicy(String name, Output<String> id, @Nullable EndpointAclPolicyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:cr/endpointAclPolicy:EndpointAclPolicy", name, state, makeResourceOptions(options, id));
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
    public static EndpointAclPolicy get(String name, Output<String> id, @Nullable EndpointAclPolicyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new EndpointAclPolicy(name, id, state, options);
    }
}
