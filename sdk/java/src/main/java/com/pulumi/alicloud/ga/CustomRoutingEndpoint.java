// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.ga;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.ga.CustomRoutingEndpointArgs;
import com.pulumi.alicloud.ga.inputs.CustomRoutingEndpointState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * Provides a Global Accelerator (GA) Custom Routing Endpoint resource.
 * 
 * For information about Global Accelerator (GA) Custom Routing Endpoint and how to use it, see [What is Custom Routing Endpoint](https://www.alibabacloud.com/help/en/global-accelerator/latest/api-ga-2019-11-20-createcustomroutingendpoints).
 * 
 * &gt; **NOTE:** Available since v1.197.0.
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
 * import com.pulumi.alicloud.ga.Accelerator;
 * import com.pulumi.alicloud.ga.AcceleratorArgs;
 * import com.pulumi.alicloud.ga.BandwidthPackage;
 * import com.pulumi.alicloud.ga.BandwidthPackageArgs;
 * import com.pulumi.alicloud.ga.BandwidthPackageAttachment;
 * import com.pulumi.alicloud.ga.BandwidthPackageAttachmentArgs;
 * import com.pulumi.alicloud.ga.Listener;
 * import com.pulumi.alicloud.ga.ListenerArgs;
 * import com.pulumi.alicloud.ga.inputs.ListenerPortRangeArgs;
 * import com.pulumi.alicloud.ga.CustomRoutingEndpointGroup;
 * import com.pulumi.alicloud.ga.CustomRoutingEndpointGroupArgs;
 * import com.pulumi.alicloud.ga.CustomRoutingEndpoint;
 * import com.pulumi.alicloud.ga.CustomRoutingEndpointArgs;
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
 *         final var region = config.get(&#34;region&#34;).orElse(&#34;cn-hangzhou&#34;);
 *         final var defaultZones = AlicloudFunctions.getZones(GetZonesArgs.builder()
 *             .availableResourceCreation(&#34;VSwitch&#34;)
 *             .build());
 * 
 *         var defaultNetwork = new Network(&#34;defaultNetwork&#34;, NetworkArgs.builder()        
 *             .vpcName(&#34;terraform-example&#34;)
 *             .cidrBlock(&#34;172.17.3.0/24&#34;)
 *             .build());
 * 
 *         var defaultSwitch = new Switch(&#34;defaultSwitch&#34;, SwitchArgs.builder()        
 *             .vswitchName(&#34;terraform-example&#34;)
 *             .cidrBlock(&#34;172.17.3.0/24&#34;)
 *             .vpcId(defaultNetwork.id())
 *             .zoneId(defaultZones.applyValue(getZonesResult -&gt; getZonesResult.zones()[0].id()))
 *             .build());
 * 
 *         var defaultAccelerator = new Accelerator(&#34;defaultAccelerator&#34;, AcceleratorArgs.builder()        
 *             .duration(1)
 *             .autoUseCoupon(true)
 *             .spec(&#34;1&#34;)
 *             .build());
 * 
 *         var defaultBandwidthPackage = new BandwidthPackage(&#34;defaultBandwidthPackage&#34;, BandwidthPackageArgs.builder()        
 *             .bandwidth(100)
 *             .type(&#34;Basic&#34;)
 *             .bandwidthType(&#34;Basic&#34;)
 *             .paymentType(&#34;PayAsYouGo&#34;)
 *             .billingType(&#34;PayBy95&#34;)
 *             .ratio(30)
 *             .build());
 * 
 *         var defaultBandwidthPackageAttachment = new BandwidthPackageAttachment(&#34;defaultBandwidthPackageAttachment&#34;, BandwidthPackageAttachmentArgs.builder()        
 *             .acceleratorId(defaultAccelerator.id())
 *             .bandwidthPackageId(defaultBandwidthPackage.id())
 *             .build());
 * 
 *         var defaultListener = new Listener(&#34;defaultListener&#34;, ListenerArgs.builder()        
 *             .acceleratorId(defaultBandwidthPackageAttachment.acceleratorId())
 *             .listenerType(&#34;CustomRouting&#34;)
 *             .portRanges(ListenerPortRangeArgs.builder()
 *                 .fromPort(10000)
 *                 .toPort(16000)
 *                 .build())
 *             .build());
 * 
 *         var defaultCustomRoutingEndpointGroup = new CustomRoutingEndpointGroup(&#34;defaultCustomRoutingEndpointGroup&#34;, CustomRoutingEndpointGroupArgs.builder()        
 *             .acceleratorId(defaultListener.acceleratorId())
 *             .listenerId(defaultListener.id())
 *             .endpointGroupRegion(region)
 *             .customRoutingEndpointGroupName(&#34;terraform-example&#34;)
 *             .description(&#34;terraform-example&#34;)
 *             .build());
 * 
 *         var defaultCustomRoutingEndpoint = new CustomRoutingEndpoint(&#34;defaultCustomRoutingEndpoint&#34;, CustomRoutingEndpointArgs.builder()        
 *             .endpointGroupId(defaultCustomRoutingEndpointGroup.id())
 *             .endpoint(defaultSwitch.id())
 *             .type(&#34;PrivateSubNet&#34;)
 *             .trafficToEndpointPolicy(&#34;DenyAll&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * Global Accelerator (GA) Custom Routing Endpoint can be imported using the id, e.g.
 * 
 * ```sh
 * $ pulumi import alicloud:ga/customRoutingEndpoint:CustomRoutingEndpoint example &lt;endpoint_group_id&gt;:&lt;custom_routing_endpoint_id&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:ga/customRoutingEndpoint:CustomRoutingEndpoint")
public class CustomRoutingEndpoint extends com.pulumi.resources.CustomResource {
    /**
     * The ID of the GA instance with which the endpoint is associated.
     * 
     */
    @Export(name="acceleratorId", refs={String.class}, tree="[0]")
    private Output<String> acceleratorId;

    /**
     * @return The ID of the GA instance with which the endpoint is associated.
     * 
     */
    public Output<String> acceleratorId() {
        return this.acceleratorId;
    }
    /**
     * The ID of the Custom Routing Endpoint.
     * 
     */
    @Export(name="customRoutingEndpointId", refs={String.class}, tree="[0]")
    private Output<String> customRoutingEndpointId;

    /**
     * @return The ID of the Custom Routing Endpoint.
     * 
     */
    public Output<String> customRoutingEndpointId() {
        return this.customRoutingEndpointId;
    }
    /**
     * The ID of the endpoint (vSwitch).
     * 
     */
    @Export(name="endpoint", refs={String.class}, tree="[0]")
    private Output<String> endpoint;

    /**
     * @return The ID of the endpoint (vSwitch).
     * 
     */
    public Output<String> endpoint() {
        return this.endpoint;
    }
    /**
     * The ID of the endpoint group in which to create endpoints.
     * 
     */
    @Export(name="endpointGroupId", refs={String.class}, tree="[0]")
    private Output<String> endpointGroupId;

    /**
     * @return The ID of the endpoint group in which to create endpoints.
     * 
     */
    public Output<String> endpointGroupId() {
        return this.endpointGroupId;
    }
    /**
     * The ID of the listener with which the endpoint is associated.
     * 
     */
    @Export(name="listenerId", refs={String.class}, tree="[0]")
    private Output<String> listenerId;

    /**
     * @return The ID of the listener with which the endpoint is associated.
     * 
     */
    public Output<String> listenerId() {
        return this.listenerId;
    }
    /**
     * The status of the Custom Routing Endpoint.
     * 
     */
    @Export(name="status", refs={String.class}, tree="[0]")
    private Output<String> status;

    /**
     * @return The status of the Custom Routing Endpoint.
     * 
     */
    public Output<String> status() {
        return this.status;
    }
    /**
     * The access policy of traffic to the endpoint. Default value: `DenyAll`. Valid values:
     * 
     */
    @Export(name="trafficToEndpointPolicy", refs={String.class}, tree="[0]")
    private Output<String> trafficToEndpointPolicy;

    /**
     * @return The access policy of traffic to the endpoint. Default value: `DenyAll`. Valid values:
     * 
     */
    public Output<String> trafficToEndpointPolicy() {
        return this.trafficToEndpointPolicy;
    }
    /**
     * The backend service type of the endpoint. Valid values: `PrivateSubNet`.
     * 
     */
    @Export(name="type", refs={String.class}, tree="[0]")
    private Output<String> type;

    /**
     * @return The backend service type of the endpoint. Valid values: `PrivateSubNet`.
     * 
     */
    public Output<String> type() {
        return this.type;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public CustomRoutingEndpoint(String name) {
        this(name, CustomRoutingEndpointArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public CustomRoutingEndpoint(String name, CustomRoutingEndpointArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public CustomRoutingEndpoint(String name, CustomRoutingEndpointArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:ga/customRoutingEndpoint:CustomRoutingEndpoint", name, args == null ? CustomRoutingEndpointArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private CustomRoutingEndpoint(String name, Output<String> id, @Nullable CustomRoutingEndpointState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:ga/customRoutingEndpoint:CustomRoutingEndpoint", name, state, makeResourceOptions(options, id));
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
    public static CustomRoutingEndpoint get(String name, Output<String> id, @Nullable CustomRoutingEndpointState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new CustomRoutingEndpoint(name, id, state, options);
    }
}
