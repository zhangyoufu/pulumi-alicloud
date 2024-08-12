// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.vpc;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.vpc.BgpGroupArgs;
import com.pulumi.alicloud.vpc.inputs.BgpGroupState;
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
 * Provides a VPC Bgp Group resource.
 * 
 * For information about VPC Bgp Group and how to use it, see [What is Bgp Group](https://www.alibabacloud.com/help/en/doc-detail/91267.html).
 * 
 * &gt; **NOTE:** Available since v1.152.0.
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
 * import com.pulumi.alicloud.expressconnect.ExpressconnectFunctions;
 * import com.pulumi.alicloud.expressconnect.inputs.GetPhysicalConnectionsArgs;
 * import com.pulumi.random.integer;
 * import com.pulumi.random.IntegerArgs;
 * import com.pulumi.alicloud.expressconnect.VirtualBorderRouter;
 * import com.pulumi.alicloud.expressconnect.VirtualBorderRouterArgs;
 * import com.pulumi.alicloud.vpc.BgpGroup;
 * import com.pulumi.alicloud.vpc.BgpGroupArgs;
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
 *         final var name = config.get("name").orElse("tf-example");
 *         final var example = ExpressconnectFunctions.getPhysicalConnections(GetPhysicalConnectionsArgs.builder()
 *             .nameRegex("^preserved-NODELETING")
 *             .build());
 * 
 *         var vlanId = new Integer("vlanId", IntegerArgs.builder()
 *             .max(2999)
 *             .min(1)
 *             .build());
 * 
 *         var exampleVirtualBorderRouter = new VirtualBorderRouter("exampleVirtualBorderRouter", VirtualBorderRouterArgs.builder()
 *             .localGatewayIp("10.0.0.1")
 *             .peerGatewayIp("10.0.0.2")
 *             .peeringSubnetMask("255.255.255.252")
 *             .physicalConnectionId(example.applyValue(getPhysicalConnectionsResult -> getPhysicalConnectionsResult.connections()[0].id()))
 *             .virtualBorderRouterName(name)
 *             .vlanId(vlanId.id())
 *             .minRxInterval(1000)
 *             .minTxInterval(1000)
 *             .detectMultiplier(10)
 *             .build());
 * 
 *         var exampleBgpGroup = new BgpGroup("exampleBgpGroup", BgpGroupArgs.builder()
 *             .authKey("YourPassword+12345678")
 *             .bgpGroupName(name)
 *             .description(name)
 *             .peerAsn(1111)
 *             .routerId(exampleVirtualBorderRouter.id())
 *             .isFakeAsn(true)
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
 * VPC Bgp Group can be imported using the id, e.g.
 * 
 * ```sh
 * $ pulumi import alicloud:vpc/bgpGroup:BgpGroup example &lt;id&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:vpc/bgpGroup:BgpGroup")
public class BgpGroup extends com.pulumi.resources.CustomResource {
    /**
     * The authentication key of the BGP group.
     * 
     */
    @Export(name="authKey", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> authKey;

    /**
     * @return The authentication key of the BGP group.
     * 
     */
    public Output<Optional<String>> authKey() {
        return Codegen.optional(this.authKey);
    }
    /**
     * The name of the BGP group. The name must be `2` to `128` characters in length and can contain digits, periods (.), underscores (_), and hyphens (-). The name must start with a letter but cannot start with `http://` or `https://`.
     * 
     */
    @Export(name="bgpGroupName", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> bgpGroupName;

    /**
     * @return The name of the BGP group. The name must be `2` to `128` characters in length and can contain digits, periods (.), underscores (_), and hyphens (-). The name must start with a letter but cannot start with `http://` or `https://`.
     * 
     */
    public Output<Optional<String>> bgpGroupName() {
        return Codegen.optional(this.bgpGroupName);
    }
    /**
     * The description of the BGP group. The description must be `2` to `256` characters in length. It must start with a letter but cannot start with `http://` or `https://`.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return The description of the BGP group. The description must be `2` to `256` characters in length. It must start with a letter but cannot start with `http://` or `https://`.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * The is fake asn. A router that runs BGP typically belongs to only one AS. In some cases, for example, the AS needs to be migrated or is merged with another AS, a new AS number replaces the original one.
     * 
     */
    @Export(name="isFakeAsn", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> isFakeAsn;

    /**
     * @return The is fake asn. A router that runs BGP typically belongs to only one AS. In some cases, for example, the AS needs to be migrated or is merged with another AS, a new AS number replaces the original one.
     * 
     */
    public Output<Boolean> isFakeAsn() {
        return this.isFakeAsn;
    }
    /**
     * The AS number on the Alibaba Cloud side.
     * 
     */
    @Export(name="localAsn", refs={Integer.class}, tree="[0]")
    private Output<Integer> localAsn;

    /**
     * @return The AS number on the Alibaba Cloud side.
     * 
     */
    public Output<Integer> localAsn() {
        return this.localAsn;
    }
    /**
     * The AS number of the BGP peer.
     * 
     */
    @Export(name="peerAsn", refs={Integer.class}, tree="[0]")
    private Output<Integer> peerAsn;

    /**
     * @return The AS number of the BGP peer.
     * 
     */
    public Output<Integer> peerAsn() {
        return this.peerAsn;
    }
    /**
     * The ID of the VBR.
     * 
     */
    @Export(name="routerId", refs={String.class}, tree="[0]")
    private Output<String> routerId;

    /**
     * @return The ID of the VBR.
     * 
     */
    public Output<String> routerId() {
        return this.routerId;
    }
    /**
     * The status of the resource.
     * 
     */
    @Export(name="status", refs={String.class}, tree="[0]")
    private Output<String> status;

    /**
     * @return The status of the resource.
     * 
     */
    public Output<String> status() {
        return this.status;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public BgpGroup(java.lang.String name) {
        this(name, BgpGroupArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public BgpGroup(java.lang.String name, BgpGroupArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public BgpGroup(java.lang.String name, BgpGroupArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:vpc/bgpGroup:BgpGroup", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private BgpGroup(java.lang.String name, Output<java.lang.String> id, @Nullable BgpGroupState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:vpc/bgpGroup:BgpGroup", name, state, makeResourceOptions(options, id), false);
    }

    private static BgpGroupArgs makeArgs(BgpGroupArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? BgpGroupArgs.Empty : args;
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
    public static BgpGroup get(java.lang.String name, Output<java.lang.String> id, @Nullable BgpGroupState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new BgpGroup(name, id, state, options);
    }
}
