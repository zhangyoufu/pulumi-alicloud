// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.vpc;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.vpc.SnatEntryArgs;
import com.pulumi.alicloud.vpc.inputs.SnatEntryState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a snat resource.
 * 
 * &gt; **NOTE:** Available since v1.119.0.
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
 * import com.pulumi.alicloud.AlicloudFunctions;
 * import com.pulumi.alicloud.inputs.GetZonesArgs;
 * import com.pulumi.alicloud.vpc.Network;
 * import com.pulumi.alicloud.vpc.NetworkArgs;
 * import com.pulumi.alicloud.vpc.Switch;
 * import com.pulumi.alicloud.vpc.SwitchArgs;
 * import com.pulumi.alicloud.vpc.NatGateway;
 * import com.pulumi.alicloud.vpc.NatGatewayArgs;
 * import com.pulumi.alicloud.ecs.EipAddress;
 * import com.pulumi.alicloud.ecs.EipAddressArgs;
 * import com.pulumi.alicloud.ecs.EipAssociation;
 * import com.pulumi.alicloud.ecs.EipAssociationArgs;
 * import com.pulumi.alicloud.vpc.SnatEntry;
 * import com.pulumi.alicloud.vpc.SnatEntryArgs;
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
 *         final var name = config.get(&#34;name&#34;).orElse(&#34;tf_example&#34;);
 *         final var default = AlicloudFunctions.getZones(GetZonesArgs.builder()
 *             .availableResourceCreation(&#34;VSwitch&#34;)
 *             .build());
 * 
 *         var defaultNetwork = new Network(&#34;defaultNetwork&#34;, NetworkArgs.builder()        
 *             .vpcName(name)
 *             .cidrBlock(&#34;172.16.0.0/12&#34;)
 *             .build());
 * 
 *         var defaultSwitch = new Switch(&#34;defaultSwitch&#34;, SwitchArgs.builder()        
 *             .vpcId(defaultNetwork.id())
 *             .cidrBlock(&#34;172.16.0.0/21&#34;)
 *             .zoneId(default_.zones()[0].id())
 *             .vswitchName(name)
 *             .build());
 * 
 *         var defaultNatGateway = new NatGateway(&#34;defaultNatGateway&#34;, NatGatewayArgs.builder()        
 *             .vpcId(defaultNetwork.id())
 *             .natGatewayName(name)
 *             .paymentType(&#34;PayAsYouGo&#34;)
 *             .vswitchId(defaultSwitch.id())
 *             .natType(&#34;Enhanced&#34;)
 *             .build());
 * 
 *         var defaultEipAddress = new EipAddress(&#34;defaultEipAddress&#34;, EipAddressArgs.builder()        
 *             .addressName(name)
 *             .build());
 * 
 *         var defaultEipAssociation = new EipAssociation(&#34;defaultEipAssociation&#34;, EipAssociationArgs.builder()        
 *             .allocationId(defaultEipAddress.id())
 *             .instanceId(defaultNatGateway.id())
 *             .build());
 * 
 *         var defaultSnatEntry = new SnatEntry(&#34;defaultSnatEntry&#34;, SnatEntryArgs.builder()        
 *             .snatTableId(defaultNatGateway.snatTableIds())
 *             .sourceVswitchId(defaultSwitch.id())
 *             .snatIp(defaultEipAddress.ipAddress())
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * Snat Entry can be imported using the id, e.g.
 * 
 * ```sh
 * $ pulumi import alicloud:vpc/snatEntry:SnatEntry foo stb-1aece3:snat-232ce2
 * ```
 * 
 */
@ResourceType(type="alicloud:vpc/snatEntry:SnatEntry")
public class SnatEntry extends com.pulumi.resources.CustomResource {
    /**
     * The id of the snat entry on the server.
     * 
     */
    @Export(name="snatEntryId", refs={String.class}, tree="[0]")
    private Output<String> snatEntryId;

    /**
     * @return The id of the snat entry on the server.
     * 
     */
    public Output<String> snatEntryId() {
        return this.snatEntryId;
    }
    /**
     * The name of snat entry.
     * 
     */
    @Export(name="snatEntryName", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> snatEntryName;

    /**
     * @return The name of snat entry.
     * 
     */
    public Output<Optional<String>> snatEntryName() {
        return Codegen.optional(this.snatEntryName);
    }
    /**
     * The SNAT ip address, the ip must along bandwidth package public ip which `alicloud.vpc.NatGateway` argument `bandwidth_packages`.
     * 
     */
    @Export(name="snatIp", refs={String.class}, tree="[0]")
    private Output<String> snatIp;

    /**
     * @return The SNAT ip address, the ip must along bandwidth package public ip which `alicloud.vpc.NatGateway` argument `bandwidth_packages`.
     * 
     */
    public Output<String> snatIp() {
        return this.snatIp;
    }
    /**
     * The value can get from `alicloud.vpc.NatGateway` Attributes &#34;snat_table_ids&#34;.
     * 
     */
    @Export(name="snatTableId", refs={String.class}, tree="[0]")
    private Output<String> snatTableId;

    /**
     * @return The value can get from `alicloud.vpc.NatGateway` Attributes &#34;snat_table_ids&#34;.
     * 
     */
    public Output<String> snatTableId() {
        return this.snatTableId;
    }
    /**
     * The private network segment of Ecs. This parameter and the `source_vswitch_id` parameter are mutually exclusive and cannot appear at the same time.
     * 
     */
    @Export(name="sourceCidr", refs={String.class}, tree="[0]")
    private Output<String> sourceCidr;

    /**
     * @return The private network segment of Ecs. This parameter and the `source_vswitch_id` parameter are mutually exclusive and cannot appear at the same time.
     * 
     */
    public Output<String> sourceCidr() {
        return this.sourceCidr;
    }
    /**
     * The vswitch ID.
     * 
     */
    @Export(name="sourceVswitchId", refs={String.class}, tree="[0]")
    private Output<String> sourceVswitchId;

    /**
     * @return The vswitch ID.
     * 
     */
    public Output<String> sourceVswitchId() {
        return this.sourceVswitchId;
    }
    /**
     * (Available since v1.119.1) The status of snat entry.
     * 
     */
    @Export(name="status", refs={String.class}, tree="[0]")
    private Output<String> status;

    /**
     * @return (Available since v1.119.1) The status of snat entry.
     * 
     */
    public Output<String> status() {
        return this.status;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public SnatEntry(String name) {
        this(name, SnatEntryArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public SnatEntry(String name, SnatEntryArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public SnatEntry(String name, SnatEntryArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:vpc/snatEntry:SnatEntry", name, args == null ? SnatEntryArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private SnatEntry(String name, Output<String> id, @Nullable SnatEntryState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:vpc/snatEntry:SnatEntry", name, state, makeResourceOptions(options, id));
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
    public static SnatEntry get(String name, Output<String> id, @Nullable SnatEntryState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new SnatEntry(name, id, state, options);
    }
}
