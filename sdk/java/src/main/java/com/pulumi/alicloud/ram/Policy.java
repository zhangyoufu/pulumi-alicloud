// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.ram;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.ram.PolicyArgs;
import com.pulumi.alicloud.ram.inputs.PolicyState;
import com.pulumi.alicloud.ram.outputs.PolicyStatement;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a RAM Policy resource.
 * 
 * &gt; **NOTE:** When you want to destroy this resource forcefully(means remove all the relationships associated with it automatically and then destroy it) without set `force`  with `true` at beginning, you need add `force = true` to configuration file and run `pulumi preview`, then you can delete resource forcefully.
 * 
 * &gt; **NOTE:** Each policy can own at most 5 versions and the oldest version will be removed after its version achieves 5.
 * 
 * &gt; **NOTE:** If the policy has multiple versions, all non-default versions will be deleted first when deleting policy.
 * 
 * &gt; **NOTE:** Available since v1.0.0+.
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
 * import com.pulumi.random.integer;
 * import com.pulumi.random.IntegerArgs;
 * import com.pulumi.alicloud.ram.Policy;
 * import com.pulumi.alicloud.ram.PolicyArgs;
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
 *         // Create a new RAM Policy.
 *         var default_ = new Integer("default", IntegerArgs.builder()
 *             .min(10000)
 *             .max(99999)
 *             .build());
 * 
 *         var policy = new Policy("policy", PolicyArgs.builder()
 *             .policyName(String.format("tf-example-%s", default_.result()))
 *             .policyDocument("""
 *   {
 *     "Statement": [
 *       {
 *         "Action": [
 *           "oss:ListObjects",
 *           "oss:GetObject"
 *         ],
 *         "Effect": "Allow",
 *         "Resource": [
 *           "acs:oss:*:*:mybucket",
 *           "acs:oss:*:*:mybucket/*"
 *         ]
 *       }
 *     ],
 *       "Version": "1"
 *   }
 *             """)
 *             .description("this is a policy test")
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
 * RAM policy can be imported using the id or name, e.g.
 * 
 * ```sh
 * $ pulumi import alicloud:ram/policy:Policy example my-policy
 * ```
 * 
 */
@ResourceType(type="alicloud:ram/policy:Policy")
public class Policy extends com.pulumi.resources.CustomResource {
    /**
     * The policy attachment count.
     * 
     */
    @Export(name="attachmentCount", refs={Integer.class}, tree="[0]")
    private Output<Integer> attachmentCount;

    /**
     * @return The policy attachment count.
     * 
     */
    public Output<Integer> attachmentCount() {
        return this.attachmentCount;
    }
    /**
     * The default version of policy.
     * 
     */
    @Export(name="defaultVersion", refs={String.class}, tree="[0]")
    private Output<String> defaultVersion;

    /**
     * @return The default version of policy.
     * 
     */
    public Output<String> defaultVersion() {
        return this.defaultVersion;
    }
    /**
     * Description of the RAM policy. This name can have a string of 1 to 1024 characters.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return Description of the RAM policy. This name can have a string of 1 to 1024 characters.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * It has been deprecated since provider version 1.114.0 and `policy_document` instead.
     * 
     * @deprecated
     * Field &#39;document&#39; has been deprecated from provider version 1.114.0. New field &#39;policy_document&#39; instead.
     * 
     */
    @Deprecated /* Field 'document' has been deprecated from provider version 1.114.0. New field 'policy_document' instead. */
    @Export(name="document", refs={String.class}, tree="[0]")
    private Output<String> document;

    /**
     * @return It has been deprecated since provider version 1.114.0 and `policy_document` instead.
     * 
     */
    public Output<String> document() {
        return this.document;
    }
    /**
     * This parameter is used for resource destroy. Default value is `false`.
     * 
     */
    @Export(name="force", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> force;

    /**
     * @return This parameter is used for resource destroy. Default value is `false`.
     * 
     */
    public Output<Optional<Boolean>> force() {
        return Codegen.optional(this.force);
    }
    /**
     * It has been deprecated since provider version 1.114.0 and `policy_name` instead.
     * 
     * @deprecated
     * Field &#39;name&#39; has been deprecated from provider version 1.114.0. New field &#39;policy_name&#39; instead.
     * 
     */
    @Deprecated /* Field 'name' has been deprecated from provider version 1.114.0. New field 'policy_name' instead. */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return It has been deprecated since provider version 1.114.0 and `policy_name` instead.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * Document of the RAM policy. It is required when the `statement` is not specified.
     * 
     */
    @Export(name="policyDocument", refs={String.class}, tree="[0]")
    private Output<String> policyDocument;

    /**
     * @return Document of the RAM policy. It is required when the `statement` is not specified.
     * 
     */
    public Output<String> policyDocument() {
        return this.policyDocument;
    }
    /**
     * Name of the RAM policy. This name can have a string of 1 to 128 characters, must contain only alphanumeric characters or hyphen &#34;-&#34;, and must not begin with a hyphen.
     * 
     */
    @Export(name="policyName", refs={String.class}, tree="[0]")
    private Output<String> policyName;

    /**
     * @return Name of the RAM policy. This name can have a string of 1 to 128 characters, must contain only alphanumeric characters or hyphen &#34;-&#34;, and must not begin with a hyphen.
     * 
     */
    public Output<String> policyName() {
        return this.policyName;
    }
    /**
     * The rotation strategy of the policy. You can use this parameter to delete an early policy version. Valid Values: `None`, `DeleteOldestNonDefaultVersionWhenLimitExceeded`. Default to `None`.
     * 
     */
    @Export(name="rotateStrategy", refs={String.class}, tree="[0]")
    private Output<String> rotateStrategy;

    /**
     * @return The rotation strategy of the policy. You can use this parameter to delete an early policy version. Valid Values: `None`, `DeleteOldestNonDefaultVersionWhenLimitExceeded`. Default to `None`.
     * 
     */
    public Output<String> rotateStrategy() {
        return this.rotateStrategy;
    }
    /**
     * (It has been deprecated since version 1.49.0, and use field &#39;document&#39; to replace.) Statements of the RAM policy document. It is required when the `document` is not specified. See `statement` below.
     * 
     * @deprecated
     * Field &#39;statement&#39; has been deprecated from version 1.49.0, and use field &#39;document&#39; to replace.
     * 
     */
    @Deprecated /* Field 'statement' has been deprecated from version 1.49.0, and use field 'document' to replace.  */
    @Export(name="statements", refs={List.class,PolicyStatement.class}, tree="[0,1]")
    private Output<List<PolicyStatement>> statements;

    /**
     * @return (It has been deprecated since version 1.49.0, and use field &#39;document&#39; to replace.) Statements of the RAM policy document. It is required when the `document` is not specified. See `statement` below.
     * 
     */
    public Output<List<PolicyStatement>> statements() {
        return this.statements;
    }
    /**
     * The policy type.
     * 
     */
    @Export(name="type", refs={String.class}, tree="[0]")
    private Output<String> type;

    /**
     * @return The policy type.
     * 
     */
    public Output<String> type() {
        return this.type;
    }
    /**
     * (It has been deprecated since version 1.49.0, and use field &#39;document&#39; to replace.) Version of the RAM policy document. Valid value is `1`. Default value is `1`.
     * 
     * @deprecated
     * Field &#39;version&#39; has been deprecated from version 1.49.0, and use field &#39;document&#39; to replace.
     * 
     */
    @Deprecated /* Field 'version' has been deprecated from version 1.49.0, and use field 'document' to replace.  */
    @Export(name="version", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> version;

    /**
     * @return (It has been deprecated since version 1.49.0, and use field &#39;document&#39; to replace.) Version of the RAM policy document. Valid value is `1`. Default value is `1`.
     * 
     */
    public Output<Optional<String>> version() {
        return Codegen.optional(this.version);
    }
    /**
     * The ID of default version policy.
     * 
     */
    @Export(name="versionId", refs={String.class}, tree="[0]")
    private Output<String> versionId;

    /**
     * @return The ID of default version policy.
     * 
     */
    public Output<String> versionId() {
        return this.versionId;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Policy(java.lang.String name) {
        this(name, PolicyArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Policy(java.lang.String name, @Nullable PolicyArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Policy(java.lang.String name, @Nullable PolicyArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:ram/policy:Policy", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private Policy(java.lang.String name, Output<java.lang.String> id, @Nullable PolicyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:ram/policy:Policy", name, state, makeResourceOptions(options, id), false);
    }

    private static PolicyArgs makeArgs(@Nullable PolicyArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? PolicyArgs.Empty : args;
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
    public static Policy get(java.lang.String name, Output<java.lang.String> id, @Nullable PolicyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Policy(name, id, state, options);
    }
}
