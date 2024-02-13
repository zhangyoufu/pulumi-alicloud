// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.privatelink;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.privatelink.VpcEndpointZoneArgs;
import com.pulumi.alicloud.privatelink.inputs.VpcEndpointZoneState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Boolean;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a Private Link Vpc Endpoint Zone resource.
 * 
 * For information about Private Link Vpc Endpoint Zone and how to use it, see [What is Vpc Endpoint Zone](https://www.alibabacloud.com/help/en/privatelink/latest/api-privatelink-2020-04-15-addzonetovpcendpoint).
 * 
 * &gt; **NOTE:** Available since v1.111.0.
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
 * import com.pulumi.alicloud.privatelink.VpcEndpointService;
 * import com.pulumi.alicloud.privatelink.VpcEndpointServiceArgs;
 * import com.pulumi.alicloud.vpc.Network;
 * import com.pulumi.alicloud.vpc.NetworkArgs;
 * import com.pulumi.alicloud.vpc.Switch;
 * import com.pulumi.alicloud.vpc.SwitchArgs;
 * import com.pulumi.alicloud.ecs.SecurityGroup;
 * import com.pulumi.alicloud.ecs.SecurityGroupArgs;
 * import com.pulumi.alicloud.slb.ApplicationLoadBalancer;
 * import com.pulumi.alicloud.slb.ApplicationLoadBalancerArgs;
 * import com.pulumi.alicloud.privatelink.VpcEndpointServiceResource;
 * import com.pulumi.alicloud.privatelink.VpcEndpointServiceResourceArgs;
 * import com.pulumi.alicloud.privatelink.VpcEndpoint;
 * import com.pulumi.alicloud.privatelink.VpcEndpointArgs;
 * import com.pulumi.alicloud.privatelink.VpcEndpointZone;
 * import com.pulumi.alicloud.privatelink.VpcEndpointZoneArgs;
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
 *         final var exampleZones = AlicloudFunctions.getZones(GetZonesArgs.builder()
 *             .availableResourceCreation(&#34;VSwitch&#34;)
 *             .build());
 * 
 *         var exampleVpcEndpointService = new VpcEndpointService(&#34;exampleVpcEndpointService&#34;, VpcEndpointServiceArgs.builder()        
 *             .serviceDescription(name)
 *             .connectBandwidth(103)
 *             .autoAcceptConnection(false)
 *             .build());
 * 
 *         var exampleNetwork = new Network(&#34;exampleNetwork&#34;, NetworkArgs.builder()        
 *             .vpcName(name)
 *             .cidrBlock(&#34;10.0.0.0/8&#34;)
 *             .build());
 * 
 *         var exampleSwitch = new Switch(&#34;exampleSwitch&#34;, SwitchArgs.builder()        
 *             .vswitchName(name)
 *             .cidrBlock(&#34;10.1.0.0/16&#34;)
 *             .vpcId(exampleNetwork.id())
 *             .zoneId(exampleZones.applyValue(getZonesResult -&gt; getZonesResult.zones()[0].id()))
 *             .build());
 * 
 *         var exampleSecurityGroup = new SecurityGroup(&#34;exampleSecurityGroup&#34;, SecurityGroupArgs.builder()        
 *             .vpcId(exampleNetwork.id())
 *             .build());
 * 
 *         var exampleApplicationLoadBalancer = new ApplicationLoadBalancer(&#34;exampleApplicationLoadBalancer&#34;, ApplicationLoadBalancerArgs.builder()        
 *             .loadBalancerName(name)
 *             .vswitchId(exampleSwitch.id())
 *             .loadBalancerSpec(&#34;slb.s2.small&#34;)
 *             .addressType(&#34;intranet&#34;)
 *             .build());
 * 
 *         var exampleVpcEndpointServiceResource = new VpcEndpointServiceResource(&#34;exampleVpcEndpointServiceResource&#34;, VpcEndpointServiceResourceArgs.builder()        
 *             .serviceId(exampleVpcEndpointService.id())
 *             .resourceId(exampleApplicationLoadBalancer.id())
 *             .resourceType(&#34;slb&#34;)
 *             .build());
 * 
 *         var exampleVpcEndpoint = new VpcEndpoint(&#34;exampleVpcEndpoint&#34;, VpcEndpointArgs.builder()        
 *             .serviceId(exampleVpcEndpointServiceResource.serviceId())
 *             .securityGroupIds(exampleSecurityGroup.id())
 *             .vpcId(exampleNetwork.id())
 *             .vpcEndpointName(name)
 *             .build());
 * 
 *         var exampleVpcEndpointZone = new VpcEndpointZone(&#34;exampleVpcEndpointZone&#34;, VpcEndpointZoneArgs.builder()        
 *             .endpointId(exampleVpcEndpoint.id())
 *             .vswitchId(exampleSwitch.id())
 *             .zoneId(exampleZones.applyValue(getZonesResult -&gt; getZonesResult.zones()[0].id()))
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * Private Link Vpc Endpoint Zone can be imported using the id, e.g.
 * 
 * ```sh
 * $ pulumi import alicloud:privatelink/vpcEndpointZone:VpcEndpointZone example &lt;endpoint_id&gt;:&lt;zone_id&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:privatelink/vpcEndpointZone:VpcEndpointZone")
public class VpcEndpointZone extends com.pulumi.resources.CustomResource {
    /**
     * Specifies whether to perform only a dry run, without performing the actual request. Valid values:
     * - **true**: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error message is returned. If the request passes the dry run, the DryRunOperation error code is returned.
     * - **false (default)**: performs a dry run and performs the actual request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
     * 
     */
    @Export(name="dryRun", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> dryRun;

    /**
     * @return Specifies whether to perform only a dry run, without performing the actual request. Valid values:
     * - **true**: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error message is returned. If the request passes the dry run, the DryRunOperation error code is returned.
     * - **false (default)**: performs a dry run and performs the actual request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
     * 
     */
    public Output<Optional<Boolean>> dryRun() {
        return Codegen.optional(this.dryRun);
    }
    /**
     * The endpoint ID.
     * 
     */
    @Export(name="endpointId", refs={String.class}, tree="[0]")
    private Output<String> endpointId;

    /**
     * @return The endpoint ID.
     * 
     */
    public Output<String> endpointId() {
        return this.endpointId;
    }
    /**
     * The IP address of the endpoint ENI.
     * 
     */
    @Export(name="eniIp", refs={String.class}, tree="[0]")
    private Output<String> eniIp;

    /**
     * @return The IP address of the endpoint ENI.
     * 
     */
    public Output<String> eniIp() {
        return this.eniIp;
    }
    /**
     * The state of the zone.
     * 
     */
    @Export(name="status", refs={String.class}, tree="[0]")
    private Output<String> status;

    /**
     * @return The state of the zone.
     * 
     */
    public Output<String> status() {
        return this.status;
    }
    /**
     * The ID of the vSwitch in the zone. .
     * 
     */
    @Export(name="vswitchId", refs={String.class}, tree="[0]")
    private Output<String> vswitchId;

    /**
     * @return The ID of the vSwitch in the zone. .
     * 
     */
    public Output<String> vswitchId() {
        return this.vswitchId;
    }
    /**
     * The zone ID.
     * 
     */
    @Export(name="zoneId", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> zoneId;

    /**
     * @return The zone ID.
     * 
     */
    public Output<Optional<String>> zoneId() {
        return Codegen.optional(this.zoneId);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public VpcEndpointZone(String name) {
        this(name, VpcEndpointZoneArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public VpcEndpointZone(String name, VpcEndpointZoneArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public VpcEndpointZone(String name, VpcEndpointZoneArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:privatelink/vpcEndpointZone:VpcEndpointZone", name, args == null ? VpcEndpointZoneArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private VpcEndpointZone(String name, Output<String> id, @Nullable VpcEndpointZoneState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:privatelink/vpcEndpointZone:VpcEndpointZone", name, state, makeResourceOptions(options, id));
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
    public static VpcEndpointZone get(String name, Output<String> id, @Nullable VpcEndpointZoneState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new VpcEndpointZone(name, id, state, options);
    }
}
