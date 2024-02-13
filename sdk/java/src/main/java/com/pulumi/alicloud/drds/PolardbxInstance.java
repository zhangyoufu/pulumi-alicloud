// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.drds;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.drds.PolardbxInstanceArgs;
import com.pulumi.alicloud.drds.inputs.PolardbxInstanceState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Integer;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a DRDS Polardb X Instance resource.
 * 
 * For information about DRDS Polardb X Instance and how to use it, see [What is Polardb X Instance](https://www.alibabacloud.com/help/en/polardb/polardb-for-xscale/api-createdbinstance-1).
 * 
 * &gt; **NOTE:** Available since v1.211.0.
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
 * import com.pulumi.alicloud.AlicloudFunctions;
 * import com.pulumi.alicloud.inputs.GetZonesArgs;
 * import com.pulumi.alicloud.vpc.Network;
 * import com.pulumi.alicloud.vpc.NetworkArgs;
 * import com.pulumi.alicloud.vpc.Switch;
 * import com.pulumi.alicloud.vpc.SwitchArgs;
 * import com.pulumi.alicloud.drds.PolardbxInstance;
 * import com.pulumi.alicloud.drds.PolardbxInstanceArgs;
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
 *         final var name = config.get(&#34;name&#34;).orElse(&#34;terraform-example&#34;);
 *         final var defaultZones = AlicloudFunctions.getZones(GetZonesArgs.builder()
 *             .availableResourceCreation(&#34;VSwitch&#34;)
 *             .build());
 * 
 *         var exampleNetwork = new Network(&#34;exampleNetwork&#34;, NetworkArgs.builder()        
 *             .vpcName(name)
 *             .build());
 * 
 *         var exampleSwitch = new Switch(&#34;exampleSwitch&#34;, SwitchArgs.builder()        
 *             .vpcId(exampleNetwork.id())
 *             .zoneId(defaultZones.applyValue(getZonesResult -&gt; getZonesResult.zones()[0].id()))
 *             .cidrBlock(&#34;172.16.0.0/24&#34;)
 *             .vswitchName(name)
 *             .build());
 * 
 *         var defaultPolardbxInstance = new PolardbxInstance(&#34;defaultPolardbxInstance&#34;, PolardbxInstanceArgs.builder()        
 *             .topologyType(&#34;3azones&#34;)
 *             .vswitchId(exampleSwitch.id())
 *             .primaryZone(&#34;ap-southeast-1a&#34;)
 *             .cnNodeCount(&#34;2&#34;)
 *             .dnClass(&#34;mysql.n4.medium.25&#34;)
 *             .cnClass(&#34;polarx.x4.medium.2e&#34;)
 *             .dnNodeCount(&#34;2&#34;)
 *             .secondaryZone(&#34;ap-southeast-1b&#34;)
 *             .tertiaryZone(&#34;ap-southeast-1c&#34;)
 *             .vpcId(exampleNetwork.id())
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * DRDS Polardb X Instance can be imported using the id, e.g.
 * 
 * ```sh
 * $ pulumi import alicloud:drds/polardbxInstance:PolardbxInstance example &lt;id&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:drds/polardbxInstance:PolardbxInstance")
public class PolardbxInstance extends com.pulumi.resources.CustomResource {
    /**
     * Compute node specifications.
     * 
     */
    @Export(name="cnClass", refs={String.class}, tree="[0]")
    private Output<String> cnClass;

    /**
     * @return Compute node specifications.
     * 
     */
    public Output<String> cnClass() {
        return this.cnClass;
    }
    /**
     * Number of computing nodes.
     * 
     */
    @Export(name="cnNodeCount", refs={Integer.class}, tree="[0]")
    private Output<Integer> cnNodeCount;

    /**
     * @return Number of computing nodes.
     * 
     */
    public Output<Integer> cnNodeCount() {
        return this.cnNodeCount;
    }
    /**
     * The creation time of the resource.
     * 
     */
    @Export(name="createTime", refs={String.class}, tree="[0]")
    private Output<String> createTime;

    /**
     * @return The creation time of the resource.
     * 
     */
    public Output<String> createTime() {
        return this.createTime;
    }
    /**
     * Storage node specifications.
     * 
     */
    @Export(name="dnClass", refs={String.class}, tree="[0]")
    private Output<String> dnClass;

    /**
     * @return Storage node specifications.
     * 
     */
    public Output<String> dnClass() {
        return this.dnClass;
    }
    /**
     * The number of storage nodes.
     * 
     */
    @Export(name="dnNodeCount", refs={Integer.class}, tree="[0]")
    private Output<Integer> dnNodeCount;

    /**
     * @return The number of storage nodes.
     * 
     */
    public Output<Integer> dnNodeCount() {
        return this.dnNodeCount;
    }
    /**
     * Primary Availability Zone.
     * 
     */
    @Export(name="primaryZone", refs={String.class}, tree="[0]")
    private Output<String> primaryZone;

    /**
     * @return Primary Availability Zone.
     * 
     */
    public Output<String> primaryZone() {
        return this.primaryZone;
    }
    /**
     * The resource group ID can be empty. This parameter is not supported for the time being.
     * 
     */
    @Export(name="resourceGroupId", refs={String.class}, tree="[0]")
    private Output<String> resourceGroupId;

    /**
     * @return The resource group ID can be empty. This parameter is not supported for the time being.
     * 
     */
    public Output<String> resourceGroupId() {
        return this.resourceGroupId;
    }
    /**
     * Secondary availability zone.
     * 
     */
    @Export(name="secondaryZone", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> secondaryZone;

    /**
     * @return Secondary availability zone.
     * 
     */
    public Output<Optional<String>> secondaryZone() {
        return Codegen.optional(this.secondaryZone);
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
     * Third Availability Zone.
     * 
     */
    @Export(name="tertiaryZone", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> tertiaryZone;

    /**
     * @return Third Availability Zone.
     * 
     */
    public Output<Optional<String>> tertiaryZone() {
        return Codegen.optional(this.tertiaryZone);
    }
    /**
     * Topology type:
     * - **3azones**: three available areas;
     * - **1azone**: Single zone.
     * 
     */
    @Export(name="topologyType", refs={String.class}, tree="[0]")
    private Output<String> topologyType;

    /**
     * @return Topology type:
     * - **3azones**: three available areas;
     * - **1azone**: Single zone.
     * 
     */
    public Output<String> topologyType() {
        return this.topologyType;
    }
    /**
     * The VPC ID.
     * 
     */
    @Export(name="vpcId", refs={String.class}, tree="[0]")
    private Output<String> vpcId;

    /**
     * @return The VPC ID.
     * 
     */
    public Output<String> vpcId() {
        return this.vpcId;
    }
    /**
     * The ID of the virtual switch.
     * 
     */
    @Export(name="vswitchId", refs={String.class}, tree="[0]")
    private Output<String> vswitchId;

    /**
     * @return The ID of the virtual switch.
     * 
     */
    public Output<String> vswitchId() {
        return this.vswitchId;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public PolardbxInstance(String name) {
        this(name, PolardbxInstanceArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public PolardbxInstance(String name, PolardbxInstanceArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public PolardbxInstance(String name, PolardbxInstanceArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:drds/polardbxInstance:PolardbxInstance", name, args == null ? PolardbxInstanceArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private PolardbxInstance(String name, Output<String> id, @Nullable PolardbxInstanceState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:drds/polardbxInstance:PolardbxInstance", name, state, makeResourceOptions(options, id));
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
    public static PolardbxInstance get(String name, Output<String> id, @Nullable PolardbxInstanceState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new PolardbxInstance(name, id, state, options);
    }
}
