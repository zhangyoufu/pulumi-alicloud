// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.tag;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.tag.PolicyAttachmentArgs;
import com.pulumi.alicloud.tag.inputs.PolicyAttachmentState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * Provides a Tag Policy Attachment resource to attaches a policy to an object. After you attach a tag policy to an object.
 * For information about Tag Policy Attachment and how to use it,
 * see [What is Policy Attachment](https://www.alibabacloud.com/help/en/resource-management/latest/attach-policy).
 * 
 * &gt; **NOTE:** Available since v1.204.0.
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
 * import com.pulumi.alicloud.AlicloudFunctions;
 * import com.pulumi.alicloud.tag.Policy;
 * import com.pulumi.alicloud.tag.PolicyArgs;
 * import com.pulumi.alicloud.tag.PolicyAttachment;
 * import com.pulumi.alicloud.tag.PolicyAttachmentArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App }{{@code
 *     public static void main(String[] args) }{{@code
 *         Pulumi.run(App::stack);
 *     }}{@code
 * 
 *     public static void stack(Context ctx) }{{@code
 *         final var config = ctx.config();
 *         final var name = config.get("name").orElse("tf-example");
 *         final var default = AlicloudFunctions.getAccount();
 * 
 *         var example = new Policy("example", PolicyArgs.builder()
 *             .policyName(name)
 *             .policyDesc(name)
 *             .userType("USER")
 *             .policyContent("""
 * 		}{{@code "tags":}{{@code "CostCenter":}{{@code "tag_value":}{{@code "}{@literal @@}{@code assign":["Beijing","Shanghai"]}}{@code ,"tag_key":}{{@code "}{@literal @@}{@code assign":"CostCenter"}}}}}{@code
 *             """)
 *             .build());
 * 
 *         var examplePolicyAttachment = new PolicyAttachment("examplePolicyAttachment", PolicyAttachmentArgs.builder()
 *             .policyId(example.id())
 *             .targetId(default_.id())
 *             .targetType("USER")
 *             .build());
 * 
 *     }}{@code
 * }}{@code
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * Tag Policy Attachment can be imported using the id, e.g.
 * 
 * ```sh
 * $ pulumi import alicloud:tag/policyAttachment:PolicyAttachment example &lt;policy_id&gt;:&lt;target_id&gt;:&lt;target_type&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:tag/policyAttachment:PolicyAttachment")
public class PolicyAttachment extends com.pulumi.resources.CustomResource {
    /**
     * The ID of the tag policy.
     * 
     */
    @Export(name="policyId", refs={String.class}, tree="[0]")
    private Output<String> policyId;

    /**
     * @return The ID of the tag policy.
     * 
     */
    public Output<String> policyId() {
        return this.policyId;
    }
    /**
     * The ID of the object.
     * 
     */
    @Export(name="targetId", refs={String.class}, tree="[0]")
    private Output<String> targetId;

    /**
     * @return The ID of the object.
     * 
     */
    public Output<String> targetId() {
        return this.targetId;
    }
    /**
     * The type of the object. Valid values: `USER`, `ROOT`, `FOLDER`, `ACCOUNT`.
     * 
     */
    @Export(name="targetType", refs={String.class}, tree="[0]")
    private Output<String> targetType;

    /**
     * @return The type of the object. Valid values: `USER`, `ROOT`, `FOLDER`, `ACCOUNT`.
     * 
     */
    public Output<String> targetType() {
        return this.targetType;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public PolicyAttachment(java.lang.String name) {
        this(name, PolicyAttachmentArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public PolicyAttachment(java.lang.String name, PolicyAttachmentArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public PolicyAttachment(java.lang.String name, PolicyAttachmentArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:tag/policyAttachment:PolicyAttachment", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private PolicyAttachment(java.lang.String name, Output<java.lang.String> id, @Nullable PolicyAttachmentState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:tag/policyAttachment:PolicyAttachment", name, state, makeResourceOptions(options, id), false);
    }

    private static PolicyAttachmentArgs makeArgs(PolicyAttachmentArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? PolicyAttachmentArgs.Empty : args;
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
    public static PolicyAttachment get(java.lang.String name, Output<java.lang.String> id, @Nullable PolicyAttachmentState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new PolicyAttachment(name, id, state, options);
    }
}
