// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.bastionhost;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.bastionhost.HostGroupArgs;
import com.pulumi.alicloud.bastionhost.inputs.HostGroupState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a Bastion Host Host Group resource.
 * 
 * For information about Bastion Host Host Group and how to use it, see [What is Host Group](https://www.alibabacloud.com/help/en/doc-detail/204307.htm).
 * 
 * &gt; **NOTE:** Available in v1.134.0+.
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
 * import com.pulumi.alicloud.bastionhost.HostGroup;
 * import com.pulumi.alicloud.bastionhost.HostGroupArgs;
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
 *         var example = new HostGroup(&#34;example&#34;, HostGroupArgs.builder()        
 *             .hostGroupName(&#34;example_value&#34;)
 *             .instanceId(&#34;bastionhost-cn-tl3xxxxxxx&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * Bastion Host Host Group can be imported using the id, e.g.
 * 
 * ```sh
 *  $ pulumi import alicloud:bastionhost/hostGroup:HostGroup example &lt;instance_id&gt;:&lt;host_group_id&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:bastionhost/hostGroup:HostGroup")
public class HostGroup extends com.pulumi.resources.CustomResource {
    /**
     * Specify the New Host Group of Notes, Supports up to 500 Characters.
     * 
     */
    @Export(name="comment", type=String.class, parameters={})
    private Output</* @Nullable */ String> comment;

    /**
     * @return Specify the New Host Group of Notes, Supports up to 500 Characters.
     * 
     */
    public Output<Optional<String>> comment() {
        return Codegen.optional(this.comment);
    }
    /**
     * Host Group ID.
     * 
     */
    @Export(name="hostGroupId", type=String.class, parameters={})
    private Output<String> hostGroupId;

    /**
     * @return Host Group ID.
     * 
     */
    public Output<String> hostGroupId() {
        return this.hostGroupId;
    }
    /**
     * Specify the New Host Group Name, Supports up to 128 Characters.
     * 
     */
    @Export(name="hostGroupName", type=String.class, parameters={})
    private Output<String> hostGroupName;

    /**
     * @return Specify the New Host Group Name, Supports up to 128 Characters.
     * 
     */
    public Output<String> hostGroupName() {
        return this.hostGroupName;
    }
    /**
     * Specify the New Host Group Where the Bastion Host ID of.
     * 
     */
    @Export(name="instanceId", type=String.class, parameters={})
    private Output<String> instanceId;

    /**
     * @return Specify the New Host Group Where the Bastion Host ID of.
     * 
     */
    public Output<String> instanceId() {
        return this.instanceId;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public HostGroup(String name) {
        this(name, HostGroupArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public HostGroup(String name, HostGroupArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public HostGroup(String name, HostGroupArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:bastionhost/hostGroup:HostGroup", name, args == null ? HostGroupArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private HostGroup(String name, Output<String> id, @Nullable HostGroupState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:bastionhost/hostGroup:HostGroup", name, state, makeResourceOptions(options, id));
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
    public static HostGroup get(String name, Output<String> id, @Nullable HostGroupState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new HostGroup(name, id, state, options);
    }
}
