// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.ens;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.ens.VswitchArgs;
import com.pulumi.alicloud.ens.inputs.VswitchState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a ENS Vswitch resource.
 * 
 * For information about ENS Vswitch and how to use it, see [What is Vswitch](https://www.alibabacloud.com/help/en/ens/developer-reference/api-createvswitch).
 * 
 * &gt; **NOTE:** Available since v1.213.0.
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
 * import com.pulumi.alicloud.ens.Network;
 * import com.pulumi.alicloud.ens.NetworkArgs;
 * import com.pulumi.alicloud.ens.Vswitch;
 * import com.pulumi.alicloud.ens.VswitchArgs;
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
 *         final var name = config.get("name").orElse("terraform-example");
 *         var default_ = new Network("default", NetworkArgs.builder()
 *             .networkName(name)
 *             .description(name)
 *             .cidrBlock("192.168.2.0/24")
 *             .ensRegionId("cn-chenzhou-telecom_unicom_cmcc")
 *             .build());
 * 
 *         var defaultVswitch = new Vswitch("defaultVswitch", VswitchArgs.builder()
 *             .description(name)
 *             .cidrBlock("192.168.2.0/24")
 *             .vswitchName(name)
 *             .ensRegionId("cn-chenzhou-telecom_unicom_cmcc")
 *             .networkId(default_.id())
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
 * ENS Vswitch can be imported using the id, e.g.
 * 
 * ```sh
 * $ pulumi import alicloud:ens/vswitch:Vswitch example &lt;id&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:ens/vswitch:Vswitch")
public class Vswitch extends com.pulumi.resources.CustomResource {
    /**
     * IPv4 CIDR block of the VSwitch instance.
     * 
     */
    @Export(name="cidrBlock", refs={String.class}, tree="[0]")
    private Output<String> cidrBlock;

    /**
     * @return IPv4 CIDR block of the VSwitch instance.
     * 
     */
    public Output<String> cidrBlock() {
        return this.cidrBlock;
    }
    /**
     * The creation time of the VSwitch instance, in the UTC time format, yyyy-MM-ddTHH:mm:ssZ.
     * 
     */
    @Export(name="createTime", refs={String.class}, tree="[0]")
    private Output<String> createTime;

    /**
     * @return The creation time of the VSwitch instance, in the UTC time format, yyyy-MM-ddTHH:mm:ssZ.
     * 
     */
    public Output<String> createTime() {
        return this.createTime;
    }
    /**
     * Description of the VSwitch Instance.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return Description of the VSwitch Instance.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * ENS Region ID.
     * 
     */
    @Export(name="ensRegionId", refs={String.class}, tree="[0]")
    private Output<String> ensRegionId;

    /**
     * @return ENS Region ID.
     * 
     */
    public Output<String> ensRegionId() {
        return this.ensRegionId;
    }
    /**
     * Network ID of the VSwitch instance.
     * 
     */
    @Export(name="networkId", refs={String.class}, tree="[0]")
    private Output<String> networkId;

    /**
     * @return Network ID of the VSwitch instance.
     * 
     */
    public Output<String> networkId() {
        return this.networkId;
    }
    /**
     * Status of the switch instance.
     * 
     */
    @Export(name="status", refs={String.class}, tree="[0]")
    private Output<String> status;

    /**
     * @return Status of the switch instance.
     * 
     */
    public Output<String> status() {
        return this.status;
    }
    /**
     * Name of the switch instance.
     * 
     */
    @Export(name="vswitchName", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> vswitchName;

    /**
     * @return Name of the switch instance.
     * 
     */
    public Output<Optional<String>> vswitchName() {
        return Codegen.optional(this.vswitchName);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Vswitch(java.lang.String name) {
        this(name, VswitchArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Vswitch(java.lang.String name, VswitchArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Vswitch(java.lang.String name, VswitchArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:ens/vswitch:Vswitch", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private Vswitch(java.lang.String name, Output<java.lang.String> id, @Nullable VswitchState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:ens/vswitch:Vswitch", name, state, makeResourceOptions(options, id), false);
    }

    private static VswitchArgs makeArgs(VswitchArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? VswitchArgs.Empty : args;
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
    public static Vswitch get(java.lang.String name, Output<java.lang.String> id, @Nullable VswitchState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Vswitch(name, id, state, options);
    }
}
