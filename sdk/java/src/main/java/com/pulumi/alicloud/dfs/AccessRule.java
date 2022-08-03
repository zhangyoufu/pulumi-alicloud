// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.dfs;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.dfs.AccessRuleArgs;
import com.pulumi.alicloud.dfs.inputs.AccessRuleState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Integer;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a DFS Access Rule resource.
 * 
 * For information about DFS Access Rule and how to use it, see [What is Access Rule](https://www.alibabacloud.com/help/doc-detail/207144.htm).
 * 
 * &gt; **NOTE:** Available in v1.140.0+.
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
 * import com.pulumi.alicloud.dfs.AccessGroup;
 * import com.pulumi.alicloud.dfs.AccessGroupArgs;
 * import com.pulumi.alicloud.dfs.AccessRule;
 * import com.pulumi.alicloud.dfs.AccessRuleArgs;
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
 *         final var name = config.get(&#34;name&#34;).orElse(&#34;example_name&#34;);
 *         var defaultAccessGroup = new AccessGroup(&#34;defaultAccessGroup&#34;, AccessGroupArgs.builder()        
 *             .networkType(&#34;VPC&#34;)
 *             .accessGroupName(name)
 *             .description(name)
 *             .build());
 * 
 *         var defaultAccessRule = new AccessRule(&#34;defaultAccessRule&#34;, AccessRuleArgs.builder()        
 *             .networkSegment(&#34;192.0.2.0/24&#34;)
 *             .accessGroupId(defaultAccessGroup.id())
 *             .description(name)
 *             .rwAccessType(&#34;RDWR&#34;)
 *             .priority(&#34;10&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * DFS Access Rule can be imported using the id, e.g.
 * 
 * ```sh
 *  $ pulumi import alicloud:dfs/accessRule:AccessRule example &lt;access_group_id&gt;:&lt;access_rule_id&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:dfs/accessRule:AccessRule")
public class AccessRule extends com.pulumi.resources.CustomResource {
    /**
     * The resource ID of Access Group.
     * 
     */
    @Export(name="accessGroupId", type=String.class, parameters={})
    private Output<String> accessGroupId;

    /**
     * @return The resource ID of Access Group.
     * 
     */
    public Output<String> accessGroupId() {
        return this.accessGroupId;
    }
    /**
     * The ID of the Access Rule.
     * 
     */
    @Export(name="accessRuleId", type=String.class, parameters={})
    private Output<String> accessRuleId;

    /**
     * @return The ID of the Access Rule.
     * 
     */
    public Output<String> accessRuleId() {
        return this.accessRuleId;
    }
    /**
     * The Description of the Access Rule.
     * 
     */
    @Export(name="description", type=String.class, parameters={})
    private Output</* @Nullable */ String> description;

    /**
     * @return The Description of the Access Rule.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * The NetworkSegment of the Access Rule.
     * 
     */
    @Export(name="networkSegment", type=String.class, parameters={})
    private Output<String> networkSegment;

    /**
     * @return The NetworkSegment of the Access Rule.
     * 
     */
    public Output<String> networkSegment() {
        return this.networkSegment;
    }
    /**
     * The Priority of the Access Rule. Valid values: `1` to `100`. **NOTE:** When multiple rules are matched by the same authorized object, the high-priority rule takes effect. `1` is the highest priority.
     * 
     */
    @Export(name="priority", type=Integer.class, parameters={})
    private Output<Integer> priority;

    /**
     * @return The Priority of the Access Rule. Valid values: `1` to `100`. **NOTE:** When multiple rules are matched by the same authorized object, the high-priority rule takes effect. `1` is the highest priority.
     * 
     */
    public Output<Integer> priority() {
        return this.priority;
    }
    /**
     * The RWAccessType of the Access Rule. Valid values: `RDONLY`, `RDWR`.
     * 
     */
    @Export(name="rwAccessType", type=String.class, parameters={})
    private Output<String> rwAccessType;

    /**
     * @return The RWAccessType of the Access Rule. Valid values: `RDONLY`, `RDWR`.
     * 
     */
    public Output<String> rwAccessType() {
        return this.rwAccessType;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public AccessRule(String name) {
        this(name, AccessRuleArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public AccessRule(String name, AccessRuleArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public AccessRule(String name, AccessRuleArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:dfs/accessRule:AccessRule", name, args == null ? AccessRuleArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private AccessRule(String name, Output<String> id, @Nullable AccessRuleState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:dfs/accessRule:AccessRule", name, state, makeResourceOptions(options, id));
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
    public static AccessRule get(String name, Output<String> id, @Nullable AccessRuleState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new AccessRule(name, id, state, options);
    }
}
