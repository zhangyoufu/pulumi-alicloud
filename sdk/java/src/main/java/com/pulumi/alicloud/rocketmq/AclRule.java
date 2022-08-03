// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.rocketmq;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.rocketmq.AclRuleArgs;
import com.pulumi.alicloud.rocketmq.inputs.AclRuleState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Integer;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a Sag Acl Rule resource. This topic describes how to configure an access control list (ACL) rule for a target Smart Access Gateway instance to permit or deny access to or from specified IP addresses in the ACL rule.
 * 
 * For information about Sag Acl Rule and how to use it, see [What is access control list (ACL) rule](https://www.alibabacloud.com/help/doc-detail/111483.htm).
 * 
 * &gt; **NOTE:** Available in 1.60.0+
 * 
 * &gt; **NOTE:** Only the following regions support create Cloud Connect Network. [`cn-shanghai`, `cn-shanghai-finance-1`, `cn-hongkong`, `ap-southeast-1`, `ap-southeast-2`, `ap-southeast-3`, `ap-southeast-5`, `ap-northeast-1`, `eu-central-1`]
 * 
 * ## Example Usage
 * 
 * Basic Usage
 * 
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.alicloud.rocketmq.Acl;
 * import com.pulumi.alicloud.rocketmq.AclArgs;
 * import com.pulumi.alicloud.rocketmq.AclRule;
 * import com.pulumi.alicloud.rocketmq.AclRuleArgs;
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
 *         var defaultAcl = new Acl(&#34;defaultAcl&#34;, AclArgs.builder()        
 *             .sagCount(&#34;0&#34;)
 *             .build());
 * 
 *         var defaultAclRule = new AclRule(&#34;defaultAclRule&#34;, AclRuleArgs.builder()        
 *             .aclId(defaultAcl.id())
 *             .description(&#34;tf-testSagAclRule&#34;)
 *             .policy(&#34;accept&#34;)
 *             .ipProtocol(&#34;ALL&#34;)
 *             .direction(&#34;in&#34;)
 *             .sourceCidr(&#34;10.10.1.0/24&#34;)
 *             .sourcePortRange(&#34;-1/-1&#34;)
 *             .destCidr(&#34;192.168.1.0/24&#34;)
 *             .destPortRange(&#34;-1/-1&#34;)
 *             .priority(&#34;1&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * The Sag Acl Rule can be imported using the id, e.g.
 * 
 * ```sh
 *  $ pulumi import alicloud:rocketmq/aclRule:AclRule example acr-abc123456
 * ```
 * 
 */
@ResourceType(type="alicloud:rocketmq/aclRule:AclRule")
public class AclRule extends com.pulumi.resources.CustomResource {
    /**
     * The ID of the ACL.
     * 
     */
    @Export(name="aclId", type=String.class, parameters={})
    private Output<String> aclId;

    /**
     * @return The ID of the ACL.
     * 
     */
    public Output<String> aclId() {
        return this.aclId;
    }
    /**
     * The description of the ACL rule. It must be 1 to 512 characters in length.
     * 
     */
    @Export(name="description", type=String.class, parameters={})
    private Output</* @Nullable */ String> description;

    /**
     * @return The description of the ACL rule. It must be 1 to 512 characters in length.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * The destination address. It is an IPv4 address range in CIDR format. Default value: 0.0.0.0/0.
     * 
     */
    @Export(name="destCidr", type=String.class, parameters={})
    private Output<String> destCidr;

    /**
     * @return The destination address. It is an IPv4 address range in CIDR format. Default value: 0.0.0.0/0.
     * 
     */
    public Output<String> destCidr() {
        return this.destCidr;
    }
    /**
     * The range of the destination port. Valid value: 80/80.
     * 
     */
    @Export(name="destPortRange", type=String.class, parameters={})
    private Output<String> destPortRange;

    /**
     * @return The range of the destination port. Valid value: 80/80.
     * 
     */
    public Output<String> destPortRange() {
        return this.destPortRange;
    }
    /**
     * The direction of the ACL rule. Valid values: in|out.
     * 
     */
    @Export(name="direction", type=String.class, parameters={})
    private Output<String> direction;

    /**
     * @return The direction of the ACL rule. Valid values: in|out.
     * 
     */
    public Output<String> direction() {
        return this.direction;
    }
    /**
     * The protocol used by the ACL rule. The value is not case sensitive.
     * 
     */
    @Export(name="ipProtocol", type=String.class, parameters={})
    private Output<String> ipProtocol;

    /**
     * @return The protocol used by the ACL rule. The value is not case sensitive.
     * 
     */
    public Output<String> ipProtocol() {
        return this.ipProtocol;
    }
    /**
     * The policy used by the ACL rule. Valid values: accept|drop.
     * 
     */
    @Export(name="policy", type=String.class, parameters={})
    private Output<String> policy;

    /**
     * @return The policy used by the ACL rule. Valid values: accept|drop.
     * 
     */
    public Output<String> policy() {
        return this.policy;
    }
    /**
     * The priority of the ACL rule. Value range: 1 to 100.
     * 
     */
    @Export(name="priority", type=Integer.class, parameters={})
    private Output</* @Nullable */ Integer> priority;

    /**
     * @return The priority of the ACL rule. Value range: 1 to 100.
     * 
     */
    public Output<Optional<Integer>> priority() {
        return Codegen.optional(this.priority);
    }
    /**
     * The source address. It is an IPv4 address range in the CIDR format. Default value: 0.0.0.0/0.
     * 
     */
    @Export(name="sourceCidr", type=String.class, parameters={})
    private Output<String> sourceCidr;

    /**
     * @return The source address. It is an IPv4 address range in the CIDR format. Default value: 0.0.0.0/0.
     * 
     */
    public Output<String> sourceCidr() {
        return this.sourceCidr;
    }
    /**
     * The range of the source port. Valid value: 80/80.
     * 
     */
    @Export(name="sourcePortRange", type=String.class, parameters={})
    private Output<String> sourcePortRange;

    /**
     * @return The range of the source port. Valid value: 80/80.
     * 
     */
    public Output<String> sourcePortRange() {
        return this.sourcePortRange;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public AclRule(String name) {
        this(name, AclRuleArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public AclRule(String name, AclRuleArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public AclRule(String name, AclRuleArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:rocketmq/aclRule:AclRule", name, args == null ? AclRuleArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private AclRule(String name, Output<String> id, @Nullable AclRuleState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:rocketmq/aclRule:AclRule", name, state, makeResourceOptions(options, id));
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
    public static AclRule get(String name, Output<String> id, @Nullable AclRuleState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new AclRule(name, id, state, options);
    }
}
