// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.vpc;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.vpc.BgpNetworkArgs;
import com.pulumi.alicloud.vpc.inputs.BgpNetworkState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * Provides a VPC Bgp Network resource.
 * 
 * For information about VPC Bgp Network and how to use it, see [What is Bgp Network](https://www.alibabacloud.com/help/en/doc-detail/91267.html).
 * 
 * &gt; **NOTE:** Available since v1.153.0.
 * 
 * ## Example Usage
 * 
 * Basic Usage
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.alicloud.expressconnect.ExpressconnectFunctions;
 * import com.pulumi.alicloud.expressconnect.inputs.GetPhysicalConnectionsArgs;
 * import com.pulumi.random.RandomInteger;
 * import com.pulumi.random.RandomIntegerArgs;
 * import com.pulumi.alicloud.expressconnect.VirtualBorderRouter;
 * import com.pulumi.alicloud.expressconnect.VirtualBorderRouterArgs;
 * import com.pulumi.alicloud.vpc.BgpNetwork;
 * import com.pulumi.alicloud.vpc.BgpNetworkArgs;
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
 *         final var name = config.get(&#34;name&#34;).orElse(&#34;tf-example&#34;);
 *         final var examplePhysicalConnections = ExpressconnectFunctions.getPhysicalConnections(GetPhysicalConnectionsArgs.builder()
 *             .nameRegex(&#34;^preserved-NODELETING&#34;)
 *             .build());
 * 
 *         var vlanId = new RandomInteger(&#34;vlanId&#34;, RandomIntegerArgs.builder()        
 *             .max(2999)
 *             .min(1)
 *             .build());
 * 
 *         var exampleVirtualBorderRouter = new VirtualBorderRouter(&#34;exampleVirtualBorderRouter&#34;, VirtualBorderRouterArgs.builder()        
 *             .localGatewayIp(&#34;10.0.0.1&#34;)
 *             .peerGatewayIp(&#34;10.0.0.2&#34;)
 *             .peeringSubnetMask(&#34;255.255.255.252&#34;)
 *             .physicalConnectionId(examplePhysicalConnections.applyValue(getPhysicalConnectionsResult -&gt; getPhysicalConnectionsResult.connections()[0].id()))
 *             .virtualBorderRouterName(name)
 *             .vlanId(vlanId.id())
 *             .minRxInterval(1000)
 *             .minTxInterval(1000)
 *             .detectMultiplier(10)
 *             .build());
 * 
 *         var exampleBgpNetwork = new BgpNetwork(&#34;exampleBgpNetwork&#34;, BgpNetworkArgs.builder()        
 *             .dstCidrBlock(&#34;192.168.0.0/24&#34;)
 *             .routerId(exampleVirtualBorderRouter.id())
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * VPC Bgp Network can be imported using the id, e.g.
 * 
 * ```sh
 * $ pulumi import alicloud:vpc/bgpNetwork:BgpNetwork example &lt;router_id&gt;:&lt;dst_cidr_block&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:vpc/bgpNetwork:BgpNetwork")
public class BgpNetwork extends com.pulumi.resources.CustomResource {
    /**
     * The CIDR block of the virtual private cloud (VPC) or vSwitch that you want to connect to a data center.
     * 
     */
    @Export(name="dstCidrBlock", refs={String.class}, tree="[0]")
    private Output<String> dstCidrBlock;

    /**
     * @return The CIDR block of the virtual private cloud (VPC) or vSwitch that you want to connect to a data center.
     * 
     */
    public Output<String> dstCidrBlock() {
        return this.dstCidrBlock;
    }
    /**
     * The ID of the vRouter associated with the router interface.
     * 
     */
    @Export(name="routerId", refs={String.class}, tree="[0]")
    private Output<String> routerId;

    /**
     * @return The ID of the vRouter associated with the router interface.
     * 
     */
    public Output<String> routerId() {
        return this.routerId;
    }
    /**
     * The state of the advertised BGP network.
     * 
     */
    @Export(name="status", refs={String.class}, tree="[0]")
    private Output<String> status;

    /**
     * @return The state of the advertised BGP network.
     * 
     */
    public Output<String> status() {
        return this.status;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public BgpNetwork(String name) {
        this(name, BgpNetworkArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public BgpNetwork(String name, BgpNetworkArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public BgpNetwork(String name, BgpNetworkArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:vpc/bgpNetwork:BgpNetwork", name, args == null ? BgpNetworkArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private BgpNetwork(String name, Output<String> id, @Nullable BgpNetworkState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:vpc/bgpNetwork:BgpNetwork", name, state, makeResourceOptions(options, id));
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
    public static BgpNetwork get(String name, Output<String> id, @Nullable BgpNetworkState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new BgpNetwork(name, id, state, options);
    }
}
