// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.ga;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.ga.EndpointGroupArgs;
import com.pulumi.alicloud.ga.inputs.EndpointGroupState;
import com.pulumi.alicloud.ga.outputs.EndpointGroupEndpointConfiguration;
import com.pulumi.alicloud.ga.outputs.EndpointGroupPortOverrides;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a Global Accelerator (GA) Endpoint Group resource.
 * 
 * For information about Global Accelerator (GA) Endpoint Group and how to use it, see [What is Endpoint Group](https://www.alibabacloud.com/help/en/doc-detail/153259.htm).
 * 
 * &gt; **NOTE:** Available in v1.113.0+.
 * 
 * &gt; **NOTE:** Listeners that use different protocols support different types of endpoint groups:
 * * For a TCP or UDP listener, you can create only one default endpoint group.
 * * For an HTTP or HTTPS listener, you can create one default endpoint group and one virtual endpoint group. By default, you can create only one virtual endpoint group.
 *   * A default endpoint group refers to the endpoint group that you configure when you create an HTTP or HTTPS listener.
 *   * A virtual endpoint group refers to the endpoint group that you can create on the Endpoint Group page after you create a listener.
 * * After you create a virtual endpoint group for an HTTP or HTTPS listener, you can create a forwarding rule and associate the forwarding rule with the virtual endpoint group. Then, the HTTP or HTTPS listener forwards requests with different destination domain names or paths to the default or virtual endpoint group based on the forwarding rule. This way, you can use one Global Accelerator (GA) instance to accelerate access to multiple domain names or paths. For more information about how to create a forwarding rule, see [Manage forwarding rules](https://www.alibabacloud.com/help/en/doc-detail/204224.htm).
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
 * import com.pulumi.alicloud.ga.Accelerator;
 * import com.pulumi.alicloud.ga.AcceleratorArgs;
 * import com.pulumi.alicloud.ga.BandwidthPackage;
 * import com.pulumi.alicloud.ga.BandwidthPackageArgs;
 * import com.pulumi.alicloud.ga.BandwidthPackageAttachment;
 * import com.pulumi.alicloud.ga.BandwidthPackageAttachmentArgs;
 * import com.pulumi.alicloud.ga.Listener;
 * import com.pulumi.alicloud.ga.ListenerArgs;
 * import com.pulumi.alicloud.ga.inputs.ListenerPortRangeArgs;
 * import com.pulumi.alicloud.ecs.EipAddress;
 * import com.pulumi.alicloud.ecs.EipAddressArgs;
 * import com.pulumi.alicloud.ga.EndpointGroup;
 * import com.pulumi.alicloud.ga.EndpointGroupArgs;
 * import com.pulumi.alicloud.ga.inputs.EndpointGroupEndpointConfigurationArgs;
 * import com.pulumi.resources.CustomResourceOptions;
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
 *         var exampleAccelerator = new Accelerator(&#34;exampleAccelerator&#34;, AcceleratorArgs.builder()        
 *             .duration(1)
 *             .autoUseCoupon(true)
 *             .spec(&#34;1&#34;)
 *             .build());
 * 
 *         var deBandwidthPackage = new BandwidthPackage(&#34;deBandwidthPackage&#34;, BandwidthPackageArgs.builder()        
 *             .bandwidth(&#34;100&#34;)
 *             .type(&#34;Basic&#34;)
 *             .bandwidthType(&#34;Basic&#34;)
 *             .paymentType(&#34;PayAsYouGo&#34;)
 *             .billingType(&#34;PayBy95&#34;)
 *             .ratio(30)
 *             .build());
 * 
 *         var deBandwidthPackageAttachment = new BandwidthPackageAttachment(&#34;deBandwidthPackageAttachment&#34;, BandwidthPackageAttachmentArgs.builder()        
 *             .acceleratorId(exampleAccelerator.id())
 *             .bandwidthPackageId(deBandwidthPackage.id())
 *             .build());
 * 
 *         var exampleListener = new Listener(&#34;exampleListener&#34;, ListenerArgs.builder()        
 *             .acceleratorId(exampleAccelerator.id())
 *             .portRanges(ListenerPortRangeArgs.builder()
 *                 .fromPort(60)
 *                 .toPort(70)
 *                 .build())
 *             .build(), CustomResourceOptions.builder()
 *                 .dependsOn(deBandwidthPackageAttachment)
 *                 .build());
 * 
 *         var exampleEipAddress = new EipAddress(&#34;exampleEipAddress&#34;, EipAddressArgs.builder()        
 *             .bandwidth(&#34;10&#34;)
 *             .internetChargeType(&#34;PayByBandwidth&#34;)
 *             .build());
 * 
 *         var exampleEndpointGroup = new EndpointGroup(&#34;exampleEndpointGroup&#34;, EndpointGroupArgs.builder()        
 *             .acceleratorId(exampleAccelerator.id())
 *             .endpointConfigurations(EndpointGroupEndpointConfigurationArgs.builder()
 *                 .endpoint(exampleEipAddress.ipAddress())
 *                 .type(&#34;PublicIp&#34;)
 *                 .weight(&#34;20&#34;)
 *                 .build())
 *             .endpointGroupRegion(&#34;cn-hangzhou&#34;)
 *             .listenerId(exampleListener.id())
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * Ga Endpoint Group can be imported using the id, e.g.
 * 
 * ```sh
 *  $ pulumi import alicloud:ga/endpointGroup:EndpointGroup example &lt;id&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:ga/endpointGroup:EndpointGroup")
public class EndpointGroup extends com.pulumi.resources.CustomResource {
    /**
     * The ID of the Global Accelerator instance to which the endpoint group will be added.
     * 
     */
    @Export(name="acceleratorId", type=String.class, parameters={})
    private Output<String> acceleratorId;

    /**
     * @return The ID of the Global Accelerator instance to which the endpoint group will be added.
     * 
     */
    public Output<String> acceleratorId() {
        return this.acceleratorId;
    }
    /**
     * The description of the endpoint group.
     * 
     */
    @Export(name="description", type=String.class, parameters={})
    private Output</* @Nullable */ String> description;

    /**
     * @return The description of the endpoint group.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * The endpointConfigurations of the endpoint group.
     * 
     */
    @Export(name="endpointConfigurations", type=List.class, parameters={EndpointGroupEndpointConfiguration.class})
    private Output<List<EndpointGroupEndpointConfiguration>> endpointConfigurations;

    /**
     * @return The endpointConfigurations of the endpoint group.
     * 
     */
    public Output<List<EndpointGroupEndpointConfiguration>> endpointConfigurations() {
        return this.endpointConfigurations;
    }
    /**
     * The ID of the region where the endpoint group is deployed.
     * 
     */
    @Export(name="endpointGroupRegion", type=String.class, parameters={})
    private Output<String> endpointGroupRegion;

    /**
     * @return The ID of the region where the endpoint group is deployed.
     * 
     */
    public Output<String> endpointGroupRegion() {
        return this.endpointGroupRegion;
    }
    /**
     * The endpoint group type. Valid values: `default`, `virtual`. Default value is `default`.
     * 
     */
    @Export(name="endpointGroupType", type=String.class, parameters={})
    private Output</* @Nullable */ String> endpointGroupType;

    /**
     * @return The endpoint group type. Valid values: `default`, `virtual`. Default value is `default`.
     * 
     */
    public Output<Optional<String>> endpointGroupType() {
        return Codegen.optional(this.endpointGroupType);
    }
    /**
     * The endpoint request protocol. Valid value: `HTTP`, `HTTPS`.
     * 
     */
    @Export(name="endpointRequestProtocol", type=String.class, parameters={})
    private Output</* @Nullable */ String> endpointRequestProtocol;

    /**
     * @return The endpoint request protocol. Valid value: `HTTP`, `HTTPS`.
     * 
     */
    public Output<Optional<String>> endpointRequestProtocol() {
        return Codegen.optional(this.endpointRequestProtocol);
    }
    /**
     * The interval between two consecutive health checks. Unit: seconds.
     * 
     */
    @Export(name="healthCheckIntervalSeconds", type=Integer.class, parameters={})
    private Output</* @Nullable */ Integer> healthCheckIntervalSeconds;

    /**
     * @return The interval between two consecutive health checks. Unit: seconds.
     * 
     */
    public Output<Optional<Integer>> healthCheckIntervalSeconds() {
        return Codegen.optional(this.healthCheckIntervalSeconds);
    }
    /**
     * The path specified as the destination of the targets for health checks.
     * 
     */
    @Export(name="healthCheckPath", type=String.class, parameters={})
    private Output</* @Nullable */ String> healthCheckPath;

    /**
     * @return The path specified as the destination of the targets for health checks.
     * 
     */
    public Output<Optional<String>> healthCheckPath() {
        return Codegen.optional(this.healthCheckPath);
    }
    /**
     * The port that is used for health checks.
     * 
     */
    @Export(name="healthCheckPort", type=Integer.class, parameters={})
    private Output</* @Nullable */ Integer> healthCheckPort;

    /**
     * @return The port that is used for health checks.
     * 
     */
    public Output<Optional<Integer>> healthCheckPort() {
        return Codegen.optional(this.healthCheckPort);
    }
    /**
     * The protocol that is used to connect to the targets for health checks. Valid values: `http`, `https`, `tcp`.
     * 
     */
    @Export(name="healthCheckProtocol", type=String.class, parameters={})
    private Output</* @Nullable */ String> healthCheckProtocol;

    /**
     * @return The protocol that is used to connect to the targets for health checks. Valid values: `http`, `https`, `tcp`.
     * 
     */
    public Output<Optional<String>> healthCheckProtocol() {
        return Codegen.optional(this.healthCheckProtocol);
    }
    /**
     * The ID of the listener that is associated with the endpoint group.
     * 
     */
    @Export(name="listenerId", type=String.class, parameters={})
    private Output<String> listenerId;

    /**
     * @return The ID of the listener that is associated with the endpoint group.
     * 
     */
    public Output<String> listenerId() {
        return this.listenerId;
    }
    /**
     * The name of the endpoint group.
     * 
     */
    @Export(name="name", type=String.class, parameters={})
    private Output<String> name;

    /**
     * @return The name of the endpoint group.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * Mapping between listening port and forwarding port of boarding point.
     * 
     */
    @Export(name="portOverrides", type=EndpointGroupPortOverrides.class, parameters={})
    private Output</* @Nullable */ EndpointGroupPortOverrides> portOverrides;

    /**
     * @return Mapping between listening port and forwarding port of boarding point.
     * 
     */
    public Output<Optional<EndpointGroupPortOverrides>> portOverrides() {
        return Codegen.optional(this.portOverrides);
    }
    /**
     * The status of the endpoint group.
     * 
     */
    @Export(name="status", type=String.class, parameters={})
    private Output<String> status;

    /**
     * @return The status of the endpoint group.
     * 
     */
    public Output<String> status() {
        return this.status;
    }
    /**
     * The number of consecutive failed heath checks that must occur before the endpoint is deemed unhealthy. Default value is `3`.
     * 
     */
    @Export(name="thresholdCount", type=Integer.class, parameters={})
    private Output</* @Nullable */ Integer> thresholdCount;

    /**
     * @return The number of consecutive failed heath checks that must occur before the endpoint is deemed unhealthy. Default value is `3`.
     * 
     */
    public Output<Optional<Integer>> thresholdCount() {
        return Codegen.optional(this.thresholdCount);
    }
    /**
     * The weight of the endpoint group when the corresponding listener is associated with multiple endpoint groups.
     * 
     */
    @Export(name="trafficPercentage", type=Integer.class, parameters={})
    private Output</* @Nullable */ Integer> trafficPercentage;

    /**
     * @return The weight of the endpoint group when the corresponding listener is associated with multiple endpoint groups.
     * 
     */
    public Output<Optional<Integer>> trafficPercentage() {
        return Codegen.optional(this.trafficPercentage);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public EndpointGroup(String name) {
        this(name, EndpointGroupArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public EndpointGroup(String name, EndpointGroupArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public EndpointGroup(String name, EndpointGroupArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:ga/endpointGroup:EndpointGroup", name, args == null ? EndpointGroupArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private EndpointGroup(String name, Output<String> id, @Nullable EndpointGroupState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:ga/endpointGroup:EndpointGroup", name, state, makeResourceOptions(options, id));
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
    public static EndpointGroup get(String name, Output<String> id, @Nullable EndpointGroupState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new EndpointGroup(name, id, state, options);
    }
}
