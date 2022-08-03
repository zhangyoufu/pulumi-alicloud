// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.vpc;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.vpc.HAVipArgs;
import com.pulumi.alicloud.vpc.inputs.HAVipState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * ## Example Usage
 * 
 * Basic Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.alicloud.vpc.HAVip;
 * import com.pulumi.alicloud.vpc.HAVipArgs;
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
 *         var foo = new HAVip(&#34;foo&#34;, HAVipArgs.builder()        
 *             .description(&#34;test_havip&#34;)
 *             .vswitchId(&#34;vsw-fakeid&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * The havip can be imported using the id, e.g.
 * 
 * ```sh
 *  $ pulumi import alicloud:vpc/hAVip:HAVip foo havip-abc123456
 * ```
 * 
 */
@ResourceType(type="alicloud:vpc/hAVip:HAVip")
public class HAVip extends com.pulumi.resources.CustomResource {
    /**
     * The description of the HaVip instance.
     * 
     */
    @Export(name="description", type=String.class, parameters={})
    private Output</* @Nullable */ String> description;

    /**
     * @return The description of the HaVip instance.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * The name of the HaVip instance.
     * 
     */
    @Export(name="havipName", type=String.class, parameters={})
    private Output</* @Nullable */ String> havipName;

    /**
     * @return The name of the HaVip instance.
     * 
     */
    public Output<Optional<String>> havipName() {
        return Codegen.optional(this.havipName);
    }
    /**
     * The ip address of the HaVip. If not filled, the default will be assigned one from the vswitch.
     * 
     */
    @Export(name="ipAddress", type=String.class, parameters={})
    private Output<String> ipAddress;

    /**
     * @return The ip address of the HaVip. If not filled, the default will be assigned one from the vswitch.
     * 
     */
    public Output<String> ipAddress() {
        return this.ipAddress;
    }
    /**
     * (Available in v1.120.0+) The status of the HaVip instance.
     * 
     */
    @Export(name="status", type=String.class, parameters={})
    private Output<String> status;

    /**
     * @return (Available in v1.120.0+) The status of the HaVip instance.
     * 
     */
    public Output<String> status() {
        return this.status;
    }
    /**
     * The vswitch_id of the HaVip, the field can&#39;t be changed.
     * 
     */
    @Export(name="vswitchId", type=String.class, parameters={})
    private Output<String> vswitchId;

    /**
     * @return The vswitch_id of the HaVip, the field can&#39;t be changed.
     * 
     */
    public Output<String> vswitchId() {
        return this.vswitchId;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public HAVip(String name) {
        this(name, HAVipArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public HAVip(String name, HAVipArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public HAVip(String name, HAVipArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:vpc/hAVip:HAVip", name, args == null ? HAVipArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private HAVip(String name, Output<String> id, @Nullable HAVipState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:vpc/hAVip:HAVip", name, state, makeResourceOptions(options, id));
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
    public static HAVip get(String name, Output<String> id, @Nullable HAVipState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new HAVip(name, id, state, options);
    }
}
