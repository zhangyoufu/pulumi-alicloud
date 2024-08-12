// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.vpc;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.vpc.TrafficMirrorFilterEgressRuleArgs;
import com.pulumi.alicloud.vpc.inputs.TrafficMirrorFilterEgressRuleState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a VPC Traffic Mirror Filter Egress Rule resource.
 * 
 * For information about VPC Traffic Mirror Filter Egress Rule and how to use it, see [What is Traffic Mirror Filter Egress Rule](https://www.alibabacloud.com/help/doc-detail/261357.htm).
 * 
 * &gt; **NOTE:** Available since v1.140.0.
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
 * import com.pulumi.alicloud.vpc.TrafficMirrorFilter;
 * import com.pulumi.alicloud.vpc.TrafficMirrorFilterArgs;
 * import com.pulumi.alicloud.vpc.TrafficMirrorFilterEgressRule;
 * import com.pulumi.alicloud.vpc.TrafficMirrorFilterEgressRuleArgs;
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
 *         var example = new TrafficMirrorFilter("example", TrafficMirrorFilterArgs.builder()
 *             .trafficMirrorFilterName("example_value")
 *             .build());
 * 
 *         var default_ = new TrafficMirrorFilterEgressRule("default", TrafficMirrorFilterEgressRuleArgs.builder()
 *             .action("drop")
 *             .priority("2")
 *             .sourceCidrBlock("10.0.0.0/11")
 *             .destinationCidrBlock("10.0.0.0/12")
 *             .trafficMirrorFilterId(example.id())
 *             .protocol("ALL")
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
 * VPC Traffic Mirror Filter Egress Rule can be imported using the id, e.g.
 * 
 * ```sh
 * $ pulumi import alicloud:vpc/trafficMirrorFilterEgressRule:TrafficMirrorFilterEgressRule example &lt;traffic_mirror_filter_id&gt;:&lt;traffic_mirror_filter_egress_rule_id&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:vpc/trafficMirrorFilterEgressRule:TrafficMirrorFilterEgressRule")
public class TrafficMirrorFilterEgressRule extends com.pulumi.resources.CustomResource {
    /**
     * The collection policy of the inbound rule. Valid values: `accept` or `drop`. `accept`: collects network traffic. `drop`: does not collect network traffic.
     * 
     */
    @Export(name="action", refs={String.class}, tree="[0]")
    private Output<String> action;

    /**
     * @return The collection policy of the inbound rule. Valid values: `accept` or `drop`. `accept`: collects network traffic. `drop`: does not collect network traffic.
     * 
     */
    public Output<String> action() {
        return this.action;
    }
    /**
     * The destination CIDR block of the outbound traffic.
     * 
     */
    @Export(name="destinationCidrBlock", refs={String.class}, tree="[0]")
    private Output<String> destinationCidrBlock;

    /**
     * @return The destination CIDR block of the outbound traffic.
     * 
     */
    public Output<String> destinationCidrBlock() {
        return this.destinationCidrBlock;
    }
    /**
     * The destination CIDR block of the outbound traffic. Valid values: `1` to `65535`. Separate the first port and last port with a forward slash (/), for example, `1/200` or `80/80`. A value of `-1/-1` indicates that all ports are available. Therefore, do not set the value to `-1/-1`. **NOTE:** When `protocol` is `ICMP`, this parameter is invalid.
     * 
     */
    @Export(name="destinationPortRange", refs={String.class}, tree="[0]")
    private Output<String> destinationPortRange;

    /**
     * @return The destination CIDR block of the outbound traffic. Valid values: `1` to `65535`. Separate the first port and last port with a forward slash (/), for example, `1/200` or `80/80`. A value of `-1/-1` indicates that all ports are available. Therefore, do not set the value to `-1/-1`. **NOTE:** When `protocol` is `ICMP`, this parameter is invalid.
     * 
     */
    public Output<String> destinationPortRange() {
        return this.destinationPortRange;
    }
    /**
     * Whether to PreCheck this request only. Value:
     * - **true**: sends a check request and does not create inbound or outbound rules. Check items include whether required parameters are filled in, request format, and restrictions. If the check fails, the corresponding error is returned. If the check passes, the error code &#39;DryRunOperation&#39; is returned &#39;.
     * - **false** (default): Sends a normal request and directly creates an inbound or outbound direction rule after checking.
     * 
     */
    @Export(name="dryRun", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> dryRun;

    /**
     * @return Whether to PreCheck this request only. Value:
     * - **true**: sends a check request and does not create inbound or outbound rules. Check items include whether required parameters are filled in, request format, and restrictions. If the check fails, the corresponding error is returned. If the check passes, the error code &#39;DryRunOperation&#39; is returned &#39;.
     * - **false** (default): Sends a normal request and directly creates an inbound or outbound direction rule after checking.
     * 
     */
    public Output<Optional<Boolean>> dryRun() {
        return Codegen.optional(this.dryRun);
    }
    /**
     * The priority of the inbound rule. A smaller value indicates a higher priority. The maximum value is `10`, which indicates that you can configure at most 10 inbound rules for a filter.
     * 
     */
    @Export(name="priority", refs={Integer.class}, tree="[0]")
    private Output<Integer> priority;

    /**
     * @return The priority of the inbound rule. A smaller value indicates a higher priority. The maximum value is `10`, which indicates that you can configure at most 10 inbound rules for a filter.
     * 
     */
    public Output<Integer> priority() {
        return this.priority;
    }
    /**
     * The transport protocol used by outbound traffic that needs to be mirrored. Valid values: `ALL`, `ICMP`, `TCP`, `UDP`.
     * 
     */
    @Export(name="protocol", refs={String.class}, tree="[0]")
    private Output<String> protocol;

    /**
     * @return The transport protocol used by outbound traffic that needs to be mirrored. Valid values: `ALL`, `ICMP`, `TCP`, `UDP`.
     * 
     */
    public Output<String> protocol() {
        return this.protocol;
    }
    /**
     * . Field &#39;rule_action&#39; has been deprecated from provider version 1.211.0. New field &#39;action&#39; instead.
     * 
     * @deprecated
     * Field &#39;rule_action&#39; has been deprecated since provider version 1.211.0. New field &#39;action&#39; instead.
     * 
     */
    @Deprecated /* Field 'rule_action' has been deprecated since provider version 1.211.0. New field 'action' instead. */
    @Export(name="ruleAction", refs={String.class}, tree="[0]")
    private Output<String> ruleAction;

    /**
     * @return . Field &#39;rule_action&#39; has been deprecated from provider version 1.211.0. New field &#39;action&#39; instead.
     * 
     */
    public Output<String> ruleAction() {
        return this.ruleAction;
    }
    /**
     * The source CIDR block of the outbound traffic.
     * 
     */
    @Export(name="sourceCidrBlock", refs={String.class}, tree="[0]")
    private Output<String> sourceCidrBlock;

    /**
     * @return The source CIDR block of the outbound traffic.
     * 
     */
    public Output<String> sourceCidrBlock() {
        return this.sourceCidrBlock;
    }
    /**
     * The source port range of the outbound traffic. Valid values: `1` to `65535`. Separate the first port and last port with a forward slash (/), for example, `1/200` or `80/80`. A value of `-1/-1` indicates that all ports are available. Therefore, do not set the value to `-1/-1`. **NOTE:** When `protocol` is `ICMP`, this parameter is invalid.
     * 
     */
    @Export(name="sourcePortRange", refs={String.class}, tree="[0]")
    private Output<String> sourcePortRange;

    /**
     * @return The source port range of the outbound traffic. Valid values: `1` to `65535`. Separate the first port and last port with a forward slash (/), for example, `1/200` or `80/80`. A value of `-1/-1` indicates that all ports are available. Therefore, do not set the value to `-1/-1`. **NOTE:** When `protocol` is `ICMP`, this parameter is invalid.
     * 
     */
    public Output<String> sourcePortRange() {
        return this.sourcePortRange;
    }
    /**
     * The state of the inbound rule. `Creating`, `Created`, `Modifying` and `Deleting`.
     * 
     */
    @Export(name="status", refs={String.class}, tree="[0]")
    private Output<String> status;

    /**
     * @return The state of the inbound rule. `Creating`, `Created`, `Modifying` and `Deleting`.
     * 
     */
    public Output<String> status() {
        return this.status;
    }
    /**
     * The ID of the outbound rule.
     * 
     */
    @Export(name="trafficMirrorFilterEgressRuleId", refs={String.class}, tree="[0]")
    private Output<String> trafficMirrorFilterEgressRuleId;

    /**
     * @return The ID of the outbound rule.
     * 
     */
    public Output<String> trafficMirrorFilterEgressRuleId() {
        return this.trafficMirrorFilterEgressRuleId;
    }
    /**
     * The ID of the filter.
     * 
     * The following arguments will be discarded. Please use new fields as soon as possible:
     * 
     */
    @Export(name="trafficMirrorFilterId", refs={String.class}, tree="[0]")
    private Output<String> trafficMirrorFilterId;

    /**
     * @return The ID of the filter.
     * 
     * The following arguments will be discarded. Please use new fields as soon as possible:
     * 
     */
    public Output<String> trafficMirrorFilterId() {
        return this.trafficMirrorFilterId;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public TrafficMirrorFilterEgressRule(java.lang.String name) {
        this(name, TrafficMirrorFilterEgressRuleArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public TrafficMirrorFilterEgressRule(java.lang.String name, TrafficMirrorFilterEgressRuleArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public TrafficMirrorFilterEgressRule(java.lang.String name, TrafficMirrorFilterEgressRuleArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:vpc/trafficMirrorFilterEgressRule:TrafficMirrorFilterEgressRule", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private TrafficMirrorFilterEgressRule(java.lang.String name, Output<java.lang.String> id, @Nullable TrafficMirrorFilterEgressRuleState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:vpc/trafficMirrorFilterEgressRule:TrafficMirrorFilterEgressRule", name, state, makeResourceOptions(options, id), false);
    }

    private static TrafficMirrorFilterEgressRuleArgs makeArgs(TrafficMirrorFilterEgressRuleArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? TrafficMirrorFilterEgressRuleArgs.Empty : args;
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
    public static TrafficMirrorFilterEgressRule get(java.lang.String name, Output<java.lang.String> id, @Nullable TrafficMirrorFilterEgressRuleState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new TrafficMirrorFilterEgressRule(name, id, state, options);
    }
}
