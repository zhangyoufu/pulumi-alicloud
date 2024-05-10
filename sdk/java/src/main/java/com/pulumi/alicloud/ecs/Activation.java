// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.ecs;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.ecs.ActivationArgs;
import com.pulumi.alicloud.ecs.inputs.ActivationState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Integer;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a ECS Activation resource.
 * 
 * For information about ECS Activation and how to use it, see [What is Activation](https://www.alibabacloud.com/help/en/elastic-compute-service/latest/createactivation#doc-api-Ecs-CreateActivation).
 * 
 * &gt; **NOTE:** Available in v1.177.0+.
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
 * import com.pulumi.alicloud.ecs.Activation;
 * import com.pulumi.alicloud.ecs.ActivationArgs;
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
 *         var example = new Activation("example", ActivationArgs.builder()        
 *             .description("terraform-example")
 *             .instanceCount(10)
 *             .instanceName("terraform-example")
 *             .ipAddressRange("0.0.0.0/0")
 *             .timeToLiveInHours(4)
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
 * ECS Activation can be imported using the id, e.g.
 * 
 * ```sh
 * $ pulumi import alicloud:ecs/activation:Activation example &lt;id&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:ecs/activation:Activation")
public class Activation extends com.pulumi.resources.CustomResource {
    /**
     * The description of the activation code. The description can be 1 to 100 characters in length and cannot start with `http://` or `https://`.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return The description of the activation code. The description can be 1 to 100 characters in length and cannot start with `http://` or `https://`.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * The maximum number of times that the activation code can be used to register managed instances. Valid values: `1` to `1000`. Default value: `10`.
     * 
     */
    @Export(name="instanceCount", refs={Integer.class}, tree="[0]")
    private Output<Integer> instanceCount;

    /**
     * @return The maximum number of times that the activation code can be used to register managed instances. Valid values: `1` to `1000`. Default value: `10`.
     * 
     */
    public Output<Integer> instanceCount() {
        return this.instanceCount;
    }
    /**
     * The default instance name prefix. The instance name prefix must be 1 to 50 characters in length. It must start with a letter and cannot start with `http://` or `https://`. The instance name prefix can contain only letters, digits, periods (.), underscores (_), hyphens (-), and colons (:).
     * - If you use the activation code created by the CreateActivation operation to register managed instances, the instances are assigned sequential names that are prefixed by the value of this parameter. You can also specify a new instance name to override the assigned sequential name when you register a managed instance.
     * - If you specify InstanceName when you register a managed instance, an instance name in the format of `&lt;InstanceName&gt;-&lt;Number&gt;` is generated. The number of digits in the &lt;Number&gt; value is determined by that in the InstanceCount value. Example: 001. If you do not specify InstanceName, the hostname (Hostname) is used as the instance name.
     * 
     */
    @Export(name="instanceName", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> instanceName;

    /**
     * @return The default instance name prefix. The instance name prefix must be 1 to 50 characters in length. It must start with a letter and cannot start with `http://` or `https://`. The instance name prefix can contain only letters, digits, periods (.), underscores (_), hyphens (-), and colons (:).
     * - If you use the activation code created by the CreateActivation operation to register managed instances, the instances are assigned sequential names that are prefixed by the value of this parameter. You can also specify a new instance name to override the assigned sequential name when you register a managed instance.
     * - If you specify InstanceName when you register a managed instance, an instance name in the format of `&lt;InstanceName&gt;-&lt;Number&gt;` is generated. The number of digits in the &lt;Number&gt; value is determined by that in the InstanceCount value. Example: 001. If you do not specify InstanceName, the hostname (Hostname) is used as the instance name.
     * 
     */
    public Output<Optional<String>> instanceName() {
        return Codegen.optional(this.instanceName);
    }
    /**
     * The IP addresses of hosts that are allowed to use the activation code. The value can be IPv4 addresses, IPv6 addresses, or CIDR blocks.
     * 
     */
    @Export(name="ipAddressRange", refs={String.class}, tree="[0]")
    private Output<String> ipAddressRange;

    /**
     * @return The IP addresses of hosts that are allowed to use the activation code. The value can be IPv4 addresses, IPv6 addresses, or CIDR blocks.
     * 
     */
    public Output<String> ipAddressRange() {
        return this.ipAddressRange;
    }
    /**
     * The validity period of the activation code. The activation code cannot be used to register new instances after the validity period expires. Unit: hours. Valid values: `1` to `24`. Default value: `4`.
     * 
     */
    @Export(name="timeToLiveInHours", refs={Integer.class}, tree="[0]")
    private Output<Integer> timeToLiveInHours;

    /**
     * @return The validity period of the activation code. The activation code cannot be used to register new instances after the validity period expires. Unit: hours. Valid values: `1` to `24`. Default value: `4`.
     * 
     */
    public Output<Integer> timeToLiveInHours() {
        return this.timeToLiveInHours;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Activation(String name) {
        this(name, ActivationArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Activation(String name, @Nullable ActivationArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Activation(String name, @Nullable ActivationArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:ecs/activation:Activation", name, args == null ? ActivationArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Activation(String name, Output<String> id, @Nullable ActivationState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:ecs/activation:Activation", name, state, makeResourceOptions(options, id));
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
    public static Activation get(String name, Output<String> id, @Nullable ActivationState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Activation(name, id, state, options);
    }
}
