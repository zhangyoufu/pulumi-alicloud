// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.slb;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.slb.LoadBalancerArgs;
import com.pulumi.alicloud.slb.inputs.LoadBalancerState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Integer;
import java.lang.Object;
import java.lang.String;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * &gt; **NOTE:** Deprecated since v1.123.1.
 * 
 * &gt; **DEPRECATED:** This resource has been renamed to alicloud.slb.ApplicationLoadBalancer from version 1.123.1.
 * 
 * Provides an Application Load Balancer resource.
 * 
 * &gt; **NOTE:** At present, to avoid some unnecessary regulation confusion, SLB can not support alicloud international account to create &#34;paybybandwidth&#34; instance.
 * 
 * &gt; **NOTE:** The supported specifications vary by region. Currently not all regions support guaranteed-performance instances.
 * For more details about guaranteed-performance instance, see [Guaranteed-performance instances](https://www.alibabacloud.com/help/en/slb/classic-load-balancer/developer-reference/api-createloadbalancer-2#t4182.html).
 * 
 * ## Example Usage
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
 * import com.pulumi.alicloud.slb.LoadBalancer;
 * import com.pulumi.alicloud.slb.LoadBalancerArgs;
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
 *         final var name = config.get(&#34;name&#34;).orElse(&#34;terraformslbconfig&#34;);
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
 *         var defaultLoadBalancer = new LoadBalancer(&#34;defaultLoadBalancer&#34;, LoadBalancerArgs.builder()        
 *             .loadBalancerName(name)
 *             .loadBalancerSpec(&#34;slb.s2.small&#34;)
 *             .vswitchId(defaultSwitch.id())
 *             .tags(Map.ofEntries(
 *                 Map.entry(&#34;tag_a&#34;, 1),
 *                 Map.entry(&#34;tag_b&#34;, 2),
 *                 Map.entry(&#34;tag_c&#34;, 3),
 *                 Map.entry(&#34;tag_d&#34;, 4),
 *                 Map.entry(&#34;tag_e&#34;, 5),
 *                 Map.entry(&#34;tag_f&#34;, 6),
 *                 Map.entry(&#34;tag_g&#34;, 7),
 *                 Map.entry(&#34;tag_h&#34;, 8),
 *                 Map.entry(&#34;tag_i&#34;, 9),
 *                 Map.entry(&#34;tag_j&#34;, 10)
 *             ))
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * Load balancer can be imported using the id, e.g.
 * 
 * ```sh
 * $ pulumi import alicloud:slb/loadBalancer:LoadBalancer example lb-abc123456
 * ```
 * 
 * @deprecated
 * This resource has been deprecated in favour of the ApplicationLoadBalancer resource
 * 
 */
@Deprecated /* This resource has been deprecated in favour of the ApplicationLoadBalancer resource */
@ResourceType(type="alicloud:slb/loadBalancer:LoadBalancer")
public class LoadBalancer extends com.pulumi.resources.CustomResource {
    /**
     * Specify the IP address of the private network for the SLB instance, which must be in the destination CIDR block of the correspond ing switch.
     * 
     */
    @Export(name="address", refs={String.class}, tree="[0]")
    private Output<String> address;

    /**
     * @return Specify the IP address of the private network for the SLB instance, which must be in the destination CIDR block of the correspond ing switch.
     * 
     */
    public Output<String> address() {
        return this.address;
    }
    /**
     * The IP version of the SLB instance to be created, which can be set to ipv4 or ipv6 . Default to &#34;ipv4&#34;. Now, only internet instance support ipv6 address.
     * 
     */
    @Export(name="addressIpVersion", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> addressIpVersion;

    /**
     * @return The IP version of the SLB instance to be created, which can be set to ipv4 or ipv6 . Default to &#34;ipv4&#34;. Now, only internet instance support ipv6 address.
     * 
     */
    public Output<Optional<String>> addressIpVersion() {
        return Codegen.optional(this.addressIpVersion);
    }
    /**
     * The network type of the SLB instance. Valid values: [&#34;internet&#34;, &#34;intranet&#34;]. If load balancer launched in VPC, this value must be &#34;intranet&#34;.
     * - internet: After an Internet SLB instance is created, the system allocates a public IP address so that the instance can forward requests from the Internet.
     * - intranet: After an intranet SLB instance is created, the system allocates an intranet IP address so that the instance can only forward intranet requests.
     * 
     */
    @Export(name="addressType", refs={String.class}, tree="[0]")
    private Output<String> addressType;

    /**
     * @return The network type of the SLB instance. Valid values: [&#34;internet&#34;, &#34;intranet&#34;]. If load balancer launched in VPC, this value must be &#34;intranet&#34;.
     * - internet: After an Internet SLB instance is created, the system allocates a public IP address so that the instance can forward requests from the Internet.
     * - intranet: After an intranet SLB instance is created, the system allocates an intranet IP address so that the instance can only forward intranet requests.
     * 
     */
    public Output<String> addressType() {
        return this.addressType;
    }
    /**
     * Valid
     * value is between 1 and 1000, If argument &#34;internet_charge_type&#34; is &#34;paybytraffic&#34;, then this value will be ignore.
     * 
     */
    @Export(name="bandwidth", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> bandwidth;

    /**
     * @return Valid
     * value is between 1 and 1000, If argument &#34;internet_charge_type&#34; is &#34;paybytraffic&#34;, then this value will be ignore.
     * 
     */
    public Output<Optional<Integer>> bandwidth() {
        return Codegen.optional(this.bandwidth);
    }
    /**
     * Whether enable the deletion protection or not. on: Enable deletion protection. off: Disable deletion protection. Default to off. Only postpaid instance support this function.
     * 
     */
    @Export(name="deleteProtection", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> deleteProtection;

    /**
     * @return Whether enable the deletion protection or not. on: Enable deletion protection. off: Disable deletion protection. Default to off. Only postpaid instance support this function.
     * 
     */
    public Output<Optional<String>> deleteProtection() {
        return Codegen.optional(this.deleteProtection);
    }
    /**
     * The billing method of the load balancer. Valid values are &#34;PrePaid&#34; and &#34;PostPaid&#34;. Default to &#34;PostPaid&#34;.
     * 
     */
    @Export(name="instanceChargeType", refs={String.class}, tree="[0]")
    private Output<String> instanceChargeType;

    /**
     * @return The billing method of the load balancer. Valid values are &#34;PrePaid&#34; and &#34;PostPaid&#34;. Default to &#34;PostPaid&#34;.
     * 
     */
    public Output<String> instanceChargeType() {
        return this.instanceChargeType;
    }
    /**
     * Valid
     * values are `PayByBandwidth`, `PayByTraffic`. If this value is &#34;PayByBandwidth&#34;, then argument &#34;internet&#34; must be &#34;true&#34;. Default is &#34;PayByTraffic&#34;. If load balancer launched in VPC, this value must be &#34;PayByTraffic&#34;.
     * Before version 1.10.1, the valid values are &#34;paybybandwidth&#34; and &#34;paybytraffic&#34;.
     * 
     */
    @Export(name="internetChargeType", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> internetChargeType;

    /**
     * @return Valid
     * values are `PayByBandwidth`, `PayByTraffic`. If this value is &#34;PayByBandwidth&#34;, then argument &#34;internet&#34; must be &#34;true&#34;. Default is &#34;PayByTraffic&#34;. If load balancer launched in VPC, this value must be &#34;PayByTraffic&#34;.
     * Before version 1.10.1, the valid values are &#34;paybybandwidth&#34; and &#34;paybytraffic&#34;.
     * 
     */
    public Output<Optional<String>> internetChargeType() {
        return Codegen.optional(this.internetChargeType);
    }
    @Export(name="loadBalancerName", refs={String.class}, tree="[0]")
    private Output<String> loadBalancerName;

    public Output<String> loadBalancerName() {
        return this.loadBalancerName;
    }
    /**
     * The specification of the Server Load Balancer instance. Default to empty string indicating it is &#34;Shared-Performance&#34; instance. Launching &#34;Performance-guaranteed&#34; instance, it must be specified. Valid values: `slb.s1.small`, `slb.s2.small`, `slb.s2.medium`.
     * 
     */
    @Export(name="loadBalancerSpec", refs={String.class}, tree="[0]")
    private Output<String> loadBalancerSpec;

    /**
     * @return The specification of the Server Load Balancer instance. Default to empty string indicating it is &#34;Shared-Performance&#34; instance. Launching &#34;Performance-guaranteed&#34; instance, it must be specified. Valid values: `slb.s1.small`, `slb.s2.small`, `slb.s2.medium`.
     * 
     */
    public Output<String> loadBalancerSpec() {
        return this.loadBalancerSpec;
    }
    /**
     * The primary zone ID of the SLB instance. If not specified, the system will be randomly assigned. You can query the primary and standby zones in a region by calling the DescribeZone API.
     * 
     */
    @Export(name="masterZoneId", refs={String.class}, tree="[0]")
    private Output<String> masterZoneId;

    /**
     * @return The primary zone ID of the SLB instance. If not specified, the system will be randomly assigned. You can query the primary and standby zones in a region by calling the DescribeZone API.
     * 
     */
    public Output<String> masterZoneId() {
        return this.masterZoneId;
    }
    /**
     * The reason of modification protection. It&#39;s effective when `modification_protection_status` is `ConsoleProtection`.
     * 
     */
    @Export(name="modificationProtectionReason", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> modificationProtectionReason;

    /**
     * @return The reason of modification protection. It&#39;s effective when `modification_protection_status` is `ConsoleProtection`.
     * 
     */
    public Output<Optional<String>> modificationProtectionReason() {
        return Codegen.optional(this.modificationProtectionReason);
    }
    /**
     * The status of modification protection. Valid values: `ConsoleProtection` and `NonProtection`. Default value: `NonProtection`.
     * 
     */
    @Export(name="modificationProtectionStatus", refs={String.class}, tree="[0]")
    private Output<String> modificationProtectionStatus;

    /**
     * @return The status of modification protection. Valid values: `ConsoleProtection` and `NonProtection`. Default value: `NonProtection`.
     * 
     */
    public Output<String> modificationProtectionStatus() {
        return this.modificationProtectionStatus;
    }
    /**
     * Field `name` has been deprecated from provider version 1.123.1 New field `load_balancer_name` instead.
     * 
     * @deprecated
     * Field &#39;name&#39; has been deprecated from provider version 1.123.1. New field &#39;load_balancer_name&#39; instead
     * 
     */
    @Deprecated /* Field 'name' has been deprecated from provider version 1.123.1. New field 'load_balancer_name' instead */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return Field `name` has been deprecated from provider version 1.123.1 New field `load_balancer_name` instead.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The billing method of the load balancer. Valid values are `PayAsYouGo` and `Subscription`. Default to `PayAsYouGo`.
     * 
     */
    @Export(name="paymentType", refs={String.class}, tree="[0]")
    private Output<String> paymentType;

    /**
     * @return The billing method of the load balancer. Valid values are `PayAsYouGo` and `Subscription`. Default to `PayAsYouGo`.
     * 
     */
    public Output<String> paymentType() {
        return this.paymentType;
    }
    /**
     * The duration that you will buy the resource, in month. It is valid when `instance_charge_type` is `PrePaid`. Valid values: [1-9, 12, 24, 36].
     * &gt; **NOTE:** The attribute `period` is only used to create Subscription instance or modify the PayAsYouGo instance to Subscription. Once effect, it will not be modified that means running `pulumi up` will not effect the resource.
     * 
     */
    @Export(name="period", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> period;

    /**
     * @return The duration that you will buy the resource, in month. It is valid when `instance_charge_type` is `PrePaid`. Valid values: [1-9, 12, 24, 36].
     * &gt; **NOTE:** The attribute `period` is only used to create Subscription instance or modify the PayAsYouGo instance to Subscription. Once effect, it will not be modified that means running `pulumi up` will not effect the resource.
     * 
     */
    public Output<Optional<Integer>> period() {
        return Codegen.optional(this.period);
    }
    /**
     * The Id of resource group which the SLB belongs.
     * 
     */
    @Export(name="resourceGroupId", refs={String.class}, tree="[0]")
    private Output<String> resourceGroupId;

    /**
     * @return The Id of resource group which the SLB belongs.
     * 
     */
    public Output<String> resourceGroupId() {
        return this.resourceGroupId;
    }
    /**
     * The standby zone ID of the SLB instance. If not specified, the system will be randomly assigned. You can query the primary and standby zones in a region by calling the DescribeZone API.
     * 
     */
    @Export(name="slaveZoneId", refs={String.class}, tree="[0]")
    private Output<String> slaveZoneId;

    /**
     * @return The standby zone ID of the SLB instance. If not specified, the system will be randomly assigned. You can query the primary and standby zones in a region by calling the DescribeZone API.
     * 
     */
    public Output<String> slaveZoneId() {
        return this.slaveZoneId;
    }
    /**
     * The specification of the Server Load Balancer instance. Default to empty string indicating it is &#34;Shared-Performance&#34; instance.
     * Launching &#34;[Performance-guaranteed](https://www.alibabacloud.com/help/en/slb/product-overview/announcements-and-updates)&#34; instance, it is must be specified and it valid values are: &#34;slb.s1.small&#34;, &#34;slb.s2.small&#34;, &#34;slb.s2.medium&#34;,
     * &#34;slb.s3.small&#34;, &#34;slb.s3.medium&#34;, &#34;slb.s3.large&#34; and &#34;slb.s4.large&#34;.
     * 
     * @deprecated
     * Field &#39;specification&#39; has been deprecated from provider version 1.123.1. New field &#39;load_balancer_spec&#39; instead
     * 
     */
    @Deprecated /* Field 'specification' has been deprecated from provider version 1.123.1. New field 'load_balancer_spec' instead */
    @Export(name="specification", refs={String.class}, tree="[0]")
    private Output<String> specification;

    /**
     * @return The specification of the Server Load Balancer instance. Default to empty string indicating it is &#34;Shared-Performance&#34; instance.
     * Launching &#34;[Performance-guaranteed](https://www.alibabacloud.com/help/en/slb/product-overview/announcements-and-updates)&#34; instance, it is must be specified and it valid values are: &#34;slb.s1.small&#34;, &#34;slb.s2.small&#34;, &#34;slb.s2.medium&#34;,
     * &#34;slb.s3.small&#34;, &#34;slb.s3.medium&#34;, &#34;slb.s3.large&#34; and &#34;slb.s4.large&#34;.
     * 
     */
    public Output<String> specification() {
        return this.specification;
    }
    /**
     * The status of slb load balancer. Valid values: `active` and `inactice`. The system default value is `active`.
     * 
     * &gt; **NOTE:** A &#34;Shared-Performance&#34; instance can be changed to &#34;Performance-guaranteed&#34;, but the change is irreversible.
     * 
     * &gt; **NOTE:** To change a &#34;Shared-Performance&#34; instance to a &#34;Performance-guaranteed&#34; instance, the SLB will have a short probability of business interruption (10 seconds-30 seconds). Advise to change it during the business downturn, or migrate business to other SLB Instances by using GSLB before changing.
     * 
     * &gt; **NOTE:** Currently, the alibaba cloud international account does not support creating a PrePaid SLB instance.
     * 
     */
    @Export(name="status", refs={String.class}, tree="[0]")
    private Output<String> status;

    /**
     * @return The status of slb load balancer. Valid values: `active` and `inactice`. The system default value is `active`.
     * 
     * &gt; **NOTE:** A &#34;Shared-Performance&#34; instance can be changed to &#34;Performance-guaranteed&#34;, but the change is irreversible.
     * 
     * &gt; **NOTE:** To change a &#34;Shared-Performance&#34; instance to a &#34;Performance-guaranteed&#34; instance, the SLB will have a short probability of business interruption (10 seconds-30 seconds). Advise to change it during the business downturn, or migrate business to other SLB Instances by using GSLB before changing.
     * 
     * &gt; **NOTE:** Currently, the alibaba cloud international account does not support creating a PrePaid SLB instance.
     * 
     */
    public Output<String> status() {
        return this.status;
    }
    /**
     * A mapping of tags to assign to the resource. The `tags` can have a maximum of 10 tag for every load balancer instance.
     * 
     */
    @Export(name="tags", refs={Map.class,String.class,Object.class}, tree="[0,1,2]")
    private Output<Map<String,Object>> tags;

    /**
     * @return A mapping of tags to assign to the resource. The `tags` can have a maximum of 10 tag for every load balancer instance.
     * 
     */
    public Output<Map<String,Object>> tags() {
        return this.tags;
    }
    /**
     * The VSwitch ID to launch in. If `address_type` is internet, it will be ignore.
     * 
     */
    @Export(name="vswitchId", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> vswitchId;

    /**
     * @return The VSwitch ID to launch in. If `address_type` is internet, it will be ignore.
     * 
     */
    public Output<Optional<String>> vswitchId() {
        return Codegen.optional(this.vswitchId);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public LoadBalancer(String name) {
        this(name, LoadBalancerArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public LoadBalancer(String name, @Nullable LoadBalancerArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public LoadBalancer(String name, @Nullable LoadBalancerArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:slb/loadBalancer:LoadBalancer", name, args == null ? LoadBalancerArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private LoadBalancer(String name, Output<String> id, @Nullable LoadBalancerState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:slb/loadBalancer:LoadBalancer", name, state, makeResourceOptions(options, id));
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
    public static LoadBalancer get(String name, Output<String> id, @Nullable LoadBalancerState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new LoadBalancer(name, id, state, options);
    }
}
