// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.ga;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.ga.BasicEndpointArgs;
import com.pulumi.alicloud.ga.inputs.BasicEndpointState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a Global Accelerator (GA) Basic Endpoint resource.
 * 
 * For information about Global Accelerator (GA) Basic Endpoint and how to use it, see [What is Basic Endpoint](https://www.alibabacloud.com/help/en/global-accelerator/latest/api-ga-2019-11-20-createbasicendpoint).
 * 
 * &gt; **NOTE:** Available since v1.194.0.
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
 * import com.pulumi.alicloud.Provider;
 * import com.pulumi.alicloud.ProviderArgs;
 * import com.pulumi.alicloud.AlicloudFunctions;
 * import com.pulumi.alicloud.inputs.GetZonesArgs;
 * import com.pulumi.alicloud.vpc.Network;
 * import com.pulumi.alicloud.vpc.NetworkArgs;
 * import com.pulumi.alicloud.vpc.Switch;
 * import com.pulumi.alicloud.vpc.SwitchArgs;
 * import com.pulumi.alicloud.ecs.SecurityGroup;
 * import com.pulumi.alicloud.ecs.SecurityGroupArgs;
 * import com.pulumi.alicloud.ecs.EcsNetworkInterface;
 * import com.pulumi.alicloud.ecs.EcsNetworkInterfaceArgs;
 * import com.pulumi.alicloud.ga.BasicAccelerator;
 * import com.pulumi.alicloud.ga.BasicAcceleratorArgs;
 * import com.pulumi.alicloud.ga.BasicEndpointGroup;
 * import com.pulumi.alicloud.ga.BasicEndpointGroupArgs;
 * import com.pulumi.alicloud.ga.BasicEndpoint;
 * import com.pulumi.alicloud.ga.BasicEndpointArgs;
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
 *         final var config = ctx.config();
 *         final var region = config.get(&#34;region&#34;).orElse(&#34;cn-shenzhen&#34;);
 *         final var endpointRegion = config.get(&#34;endpointRegion&#34;).orElse(&#34;cn-hangzhou&#34;);
 *         var sz = new Provider(&#34;sz&#34;, ProviderArgs.builder()        
 *             .region(region)
 *             .build());
 * 
 *         var hz = new Provider(&#34;hz&#34;, ProviderArgs.builder()        
 *             .region(endpointRegion)
 *             .build());
 * 
 *         final var defaultZones = AlicloudFunctions.getZones(GetZonesArgs.builder()
 *             .availableResourceCreation(&#34;VSwitch&#34;)
 *             .build());
 * 
 *         var defaultNetwork = new Network(&#34;defaultNetwork&#34;, NetworkArgs.builder()        
 *             .vpcName(&#34;terraform-example&#34;)
 *             .cidrBlock(&#34;172.17.3.0/24&#34;)
 *             .build(), CustomResourceOptions.builder()
 *                 .provider(alicloud.sz())
 *                 .build());
 * 
 *         var defaultSwitch = new Switch(&#34;defaultSwitch&#34;, SwitchArgs.builder()        
 *             .vswitchName(&#34;terraform-example&#34;)
 *             .cidrBlock(&#34;172.17.3.0/24&#34;)
 *             .vpcId(defaultNetwork.id())
 *             .zoneId(defaultZones.applyValue(getZonesResult -&gt; getZonesResult.zones()[0].id()))
 *             .build(), CustomResourceOptions.builder()
 *                 .provider(alicloud.sz())
 *                 .build());
 * 
 *         var defaultSecurityGroup = new SecurityGroup(&#34;defaultSecurityGroup&#34;, SecurityGroupArgs.builder()        
 *             .vpcId(defaultNetwork.id())
 *             .build(), CustomResourceOptions.builder()
 *                 .provider(alicloud.sz())
 *                 .build());
 * 
 *         var defaultEcsNetworkInterface = new EcsNetworkInterface(&#34;defaultEcsNetworkInterface&#34;, EcsNetworkInterfaceArgs.builder()        
 *             .vswitchId(defaultSwitch.id())
 *             .securityGroupIds(defaultSecurityGroup.id())
 *             .build(), CustomResourceOptions.builder()
 *                 .provider(alicloud.sz())
 *                 .build());
 * 
 *         var defaultBasicAccelerator = new BasicAccelerator(&#34;defaultBasicAccelerator&#34;, BasicAcceleratorArgs.builder()        
 *             .duration(1)
 *             .basicAcceleratorName(&#34;terraform-example&#34;)
 *             .description(&#34;terraform-example&#34;)
 *             .bandwidthBillingType(&#34;CDT&#34;)
 *             .autoUseCoupon(&#34;true&#34;)
 *             .autoPay(true)
 *             .build());
 * 
 *         var defaultBasicEndpointGroup = new BasicEndpointGroup(&#34;defaultBasicEndpointGroup&#34;, BasicEndpointGroupArgs.builder()        
 *             .acceleratorId(defaultBasicAccelerator.id())
 *             .endpointGroupRegion(region)
 *             .basicEndpointGroupName(&#34;terraform-example&#34;)
 *             .description(&#34;terraform-example&#34;)
 *             .build());
 * 
 *         var defaultBasicEndpoint = new BasicEndpoint(&#34;defaultBasicEndpoint&#34;, BasicEndpointArgs.builder()        
 *             .acceleratorId(defaultBasicAccelerator.id())
 *             .endpointGroupId(defaultBasicEndpointGroup.id())
 *             .endpointType(&#34;ENI&#34;)
 *             .endpointAddress(defaultEcsNetworkInterface.id())
 *             .endpointSubAddressType(&#34;secondary&#34;)
 *             .endpointSubAddress(&#34;192.168.0.1&#34;)
 *             .basicEndpointName(&#34;terraform-example&#34;)
 *             .build(), CustomResourceOptions.builder()
 *                 .provider(alicloud.hz())
 *                 .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * Global Accelerator (GA) Basic Endpoint can be imported using the id, e.g.
 * 
 * ```sh
 * $ pulumi import alicloud:ga/basicEndpoint:BasicEndpoint example &lt;endpoint_group_id&gt;:&lt;endpoint_id&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:ga/basicEndpoint:BasicEndpoint")
public class BasicEndpoint extends com.pulumi.resources.CustomResource {
    /**
     * The ID of the Basic GA instance.
     * 
     */
    @Export(name="acceleratorId", refs={String.class}, tree="[0]")
    private Output<String> acceleratorId;

    /**
     * @return The ID of the Basic GA instance.
     * 
     */
    public Output<String> acceleratorId() {
        return this.acceleratorId;
    }
    /**
     * The name of the Basic Endpoint.
     * 
     */
    @Export(name="basicEndpointName", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> basicEndpointName;

    /**
     * @return The name of the Basic Endpoint.
     * 
     */
    public Output<Optional<String>> basicEndpointName() {
        return Codegen.optional(this.basicEndpointName);
    }
    /**
     * The address of the Basic Endpoint.
     * 
     */
    @Export(name="endpointAddress", refs={String.class}, tree="[0]")
    private Output<String> endpointAddress;

    /**
     * @return The address of the Basic Endpoint.
     * 
     */
    public Output<String> endpointAddress() {
        return this.endpointAddress;
    }
    /**
     * The ID of the Basic Endpoint Group.
     * 
     */
    @Export(name="endpointGroupId", refs={String.class}, tree="[0]")
    private Output<String> endpointGroupId;

    /**
     * @return The ID of the Basic Endpoint Group.
     * 
     */
    public Output<String> endpointGroupId() {
        return this.endpointGroupId;
    }
    /**
     * The ID of the Basic Endpoint.
     * 
     */
    @Export(name="endpointId", refs={String.class}, tree="[0]")
    private Output<String> endpointId;

    /**
     * @return The ID of the Basic Endpoint.
     * 
     */
    public Output<String> endpointId() {
        return this.endpointId;
    }
    /**
     * The sub address of the Basic Endpoint.
     * 
     */
    @Export(name="endpointSubAddress", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> endpointSubAddress;

    /**
     * @return The sub address of the Basic Endpoint.
     * 
     */
    public Output<Optional<String>> endpointSubAddress() {
        return Codegen.optional(this.endpointSubAddress);
    }
    /**
     * The sub address type of the Basic Endpoint. Valid values: `primary`, `secondary`.
     * 
     */
    @Export(name="endpointSubAddressType", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> endpointSubAddressType;

    /**
     * @return The sub address type of the Basic Endpoint. Valid values: `primary`, `secondary`.
     * 
     */
    public Output<Optional<String>> endpointSubAddressType() {
        return Codegen.optional(this.endpointSubAddressType);
    }
    /**
     * The type of the Basic Endpoint. Valid values: `ENI`, `SLB`, `ECS` and `NLB`.
     * 
     */
    @Export(name="endpointType", refs={String.class}, tree="[0]")
    private Output<String> endpointType;

    /**
     * @return The type of the Basic Endpoint. Valid values: `ENI`, `SLB`, `ECS` and `NLB`.
     * 
     */
    public Output<String> endpointType() {
        return this.endpointType;
    }
    /**
     * The zone id of the Basic Endpoint.
     * 
     */
    @Export(name="endpointZoneId", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> endpointZoneId;

    /**
     * @return The zone id of the Basic Endpoint.
     * 
     */
    public Output<Optional<String>> endpointZoneId() {
        return Codegen.optional(this.endpointZoneId);
    }
    /**
     * The status of the Basic Endpoint.
     * 
     */
    @Export(name="status", refs={String.class}, tree="[0]")
    private Output<String> status;

    /**
     * @return The status of the Basic Endpoint.
     * 
     */
    public Output<String> status() {
        return this.status;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public BasicEndpoint(String name) {
        this(name, BasicEndpointArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public BasicEndpoint(String name, BasicEndpointArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public BasicEndpoint(String name, BasicEndpointArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:ga/basicEndpoint:BasicEndpoint", name, args == null ? BasicEndpointArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private BasicEndpoint(String name, Output<String> id, @Nullable BasicEndpointState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:ga/basicEndpoint:BasicEndpoint", name, state, makeResourceOptions(options, id));
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
    public static BasicEndpoint get(String name, Output<String> id, @Nullable BasicEndpointState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new BasicEndpoint(name, id, state, options);
    }
}
