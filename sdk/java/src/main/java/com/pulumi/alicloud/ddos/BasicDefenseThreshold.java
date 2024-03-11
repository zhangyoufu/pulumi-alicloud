// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.ddos;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.ddos.BasicDefenseThresholdArgs;
import com.pulumi.alicloud.ddos.inputs.BasicDefenseThresholdState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * Provides a Ddos Basic defense threshold resource.
 * 
 * For information about Ddos Basic Antiddos and how to use it, see [What is Defense Threshold](https://www.alibabacloud.com/help/en/ddos-protection/latest/modifydefensethreshold).
 * 
 * &gt; **NOTE:** Available since v1.168.0.
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
 * import com.pulumi.alicloud.ecs.EipAddress;
 * import com.pulumi.alicloud.ecs.EipAddressArgs;
 * import com.pulumi.alicloud.ddos.BasicDefenseThreshold;
 * import com.pulumi.alicloud.ddos.BasicDefenseThresholdArgs;
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
 *         var defaultEipAddress = new EipAddress(&#34;defaultEipAddress&#34;, EipAddressArgs.builder()        
 *             .addressName(name)
 *             .isp(&#34;BGP&#34;)
 *             .internetChargeType(&#34;PayByBandwidth&#34;)
 *             .paymentType(&#34;PayAsYouGo&#34;)
 *             .build());
 * 
 *         var defaultBasicDefenseThreshold = new BasicDefenseThreshold(&#34;defaultBasicDefenseThreshold&#34;, BasicDefenseThresholdArgs.builder()        
 *             .instanceId(defaultEipAddress.id())
 *             .ddosType(&#34;defense&#34;)
 *             .instanceType(&#34;eip&#34;)
 *             .bps(390)
 *             .pps(90000)
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * Ddos Basic Antiddos can be imported using the id, e.g.
 * 
 * ```sh
 * $ pulumi import alicloud:ddos/basicDefenseThreshold:BasicDefenseThreshold example &lt;instance_id&gt;:&lt;instance_type&gt;:&lt;ddos_type&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:ddos/basicDefenseThreshold:BasicDefenseThreshold")
public class BasicDefenseThreshold extends com.pulumi.resources.CustomResource {
    /**
     * Specifies the traffic scrubbing threshold. Unit: Mbit/s. The traffic scrubbing threshold cannot exceed the peak inbound or outbound Internet traffic, whichever is larger, of the asset.
     * 
     */
    @Export(name="bps", refs={Integer.class}, tree="[0]")
    private Output<Integer> bps;

    /**
     * @return Specifies the traffic scrubbing threshold. Unit: Mbit/s. The traffic scrubbing threshold cannot exceed the peak inbound or outbound Internet traffic, whichever is larger, of the asset.
     * 
     */
    public Output<Integer> bps() {
        return this.bps;
    }
    /**
     * The type of the threshold to query. Valid values: `defense`,`blackhole`.
     * 
     */
    @Export(name="ddosType", refs={String.class}, tree="[0]")
    private Output<String> ddosType;

    /**
     * @return The type of the threshold to query. Valid values: `defense`,`blackhole`.
     * 
     */
    public Output<String> ddosType() {
        return this.ddosType;
    }
    /**
     * The ID of the instance.
     * 
     */
    @Export(name="instanceId", refs={String.class}, tree="[0]")
    private Output<String> instanceId;

    /**
     * @return The ID of the instance.
     * 
     */
    public Output<String> instanceId() {
        return this.instanceId;
    }
    /**
     * The instance type of the public IP address asset. Value: `ecs`,`slb`,`eip`.
     * 
     */
    @Export(name="instanceType", refs={String.class}, tree="[0]")
    private Output<String> instanceType;

    /**
     * @return The instance type of the public IP address asset. Value: `ecs`,`slb`,`eip`.
     * 
     */
    public Output<String> instanceType() {
        return this.instanceType;
    }
    /**
     * The Internet IP address.
     * 
     */
    @Export(name="internetIp", refs={String.class}, tree="[0]")
    private Output<String> internetIp;

    /**
     * @return The Internet IP address.
     * 
     */
    public Output<String> internetIp() {
        return this.internetIp;
    }
    /**
     * Whether it is the system default threshold. Value:
     * 
     */
    @Export(name="isAuto", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> isAuto;

    /**
     * @return Whether it is the system default threshold. Value:
     * 
     */
    public Output<Boolean> isAuto() {
        return this.isAuto;
    }
    /**
     * The maximum traffic scrubbing threshold. Unit: Mbit/s.
     * 
     */
    @Export(name="maxBps", refs={Integer.class}, tree="[0]")
    private Output<Integer> maxBps;

    /**
     * @return The maximum traffic scrubbing threshold. Unit: Mbit/s.
     * 
     */
    public Output<Integer> maxBps() {
        return this.maxBps;
    }
    /**
     * The maximum packet scrubbing threshold. Unit: pps.
     * 
     */
    @Export(name="maxPps", refs={Integer.class}, tree="[0]")
    private Output<Integer> maxPps;

    /**
     * @return The maximum packet scrubbing threshold. Unit: pps.
     * 
     */
    public Output<Integer> maxPps() {
        return this.maxPps;
    }
    /**
     * The current message number cleaning threshold. Unit: pps.
     * 
     */
    @Export(name="pps", refs={Integer.class}, tree="[0]")
    private Output<Integer> pps;

    /**
     * @return The current message number cleaning threshold. Unit: pps.
     * 
     */
    public Output<Integer> pps() {
        return this.pps;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public BasicDefenseThreshold(String name) {
        this(name, BasicDefenseThresholdArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public BasicDefenseThreshold(String name, BasicDefenseThresholdArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public BasicDefenseThreshold(String name, BasicDefenseThresholdArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:ddos/basicDefenseThreshold:BasicDefenseThreshold", name, args == null ? BasicDefenseThresholdArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private BasicDefenseThreshold(String name, Output<String> id, @Nullable BasicDefenseThresholdState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:ddos/basicDefenseThreshold:BasicDefenseThreshold", name, state, makeResourceOptions(options, id));
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
    public static BasicDefenseThreshold get(String name, Output<String> id, @Nullable BasicDefenseThresholdState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new BasicDefenseThreshold(name, id, state, options);
    }
}
