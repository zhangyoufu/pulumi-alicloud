// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.alb;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.alb.LoadBalancerArgs;
import com.pulumi.alicloud.alb.inputs.LoadBalancerState;
import com.pulumi.alicloud.alb.outputs.LoadBalancerAccessLogConfig;
import com.pulumi.alicloud.alb.outputs.LoadBalancerLoadBalancerBillingConfig;
import com.pulumi.alicloud.alb.outputs.LoadBalancerModificationProtectionConfig;
import com.pulumi.alicloud.alb.outputs.LoadBalancerZoneMapping;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Boolean;
import java.lang.Object;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a ALB Load Balancer resource.
 * 
 * For information about ALB Load Balancer and how to use it, see [What is Load Balancer](https://www.alibabacloud.com/help/doc-detail/197341.htm).
 * 
 * &gt; **NOTE:** Available in v1.132.0+.
 * 
 * ## Import
 * 
 * ALB Load Balancer can be imported using the id, e.g.
 * 
 * ```sh
 *  $ pulumi import alicloud:alb/loadBalancer:LoadBalancer example &lt;id&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:alb/loadBalancer:LoadBalancer")
public class LoadBalancer extends com.pulumi.resources.CustomResource {
    /**
     * The Access Logging Configuration Structure. See the following `Block access_log_config`.
     * 
     */
    @Export(name="accessLogConfig", type=LoadBalancerAccessLogConfig.class, parameters={})
    private Output</* @Nullable */ LoadBalancerAccessLogConfig> accessLogConfig;

    /**
     * @return The Access Logging Configuration Structure. See the following `Block access_log_config`.
     * 
     */
    public Output<Optional<LoadBalancerAccessLogConfig>> accessLogConfig() {
        return Codegen.optional(this.accessLogConfig);
    }
    /**
     * The method in which IP addresses are assigned. Valid values: `Fixed` and `Dynamic`. Default value: `Dynamic`.
     * 
     */
    @Export(name="addressAllocatedMode", type=String.class, parameters={})
    private Output</* @Nullable */ String> addressAllocatedMode;

    /**
     * @return The method in which IP addresses are assigned. Valid values: `Fixed` and `Dynamic`. Default value: `Dynamic`.
     * 
     */
    public Output<Optional<String>> addressAllocatedMode() {
        return Codegen.optional(this.addressAllocatedMode);
    }
    /**
     * The IP version. Valid values: `Ipv4`, `DualStack`.
     * 
     */
    @Export(name="addressIpVersion", type=String.class, parameters={})
    private Output<String> addressIpVersion;

    /**
     * @return The IP version. Valid values: `Ipv4`, `DualStack`.
     * 
     */
    public Output<String> addressIpVersion() {
        return this.addressIpVersion;
    }
    /**
     * The type of IP address that the ALB instance uses to provide services. Valid values: `Intranet`, `Internet`. **NOTE:** From version 1.193.1, `address_type` can be modified.
     * 
     */
    @Export(name="addressType", type=String.class, parameters={})
    private Output<String> addressType;

    /**
     * @return The type of IP address that the ALB instance uses to provide services. Valid values: `Intranet`, `Internet`. **NOTE:** From version 1.193.1, `address_type` can be modified.
     * 
     */
    public Output<String> addressType() {
        return this.addressType;
    }
    /**
     * The deletion protection enabled. Valid values: `true` and `false`. Default value: `false`.
     * 
     */
    @Export(name="deletionProtectionEnabled", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> deletionProtectionEnabled;

    /**
     * @return The deletion protection enabled. Valid values: `true` and `false`. Default value: `false`.
     * 
     */
    public Output<Optional<Boolean>> deletionProtectionEnabled() {
        return Codegen.optional(this.deletionProtectionEnabled);
    }
    /**
     * The domain name of the ALB instance. **NOTE:** Available in v1.158.0+.
     * 
     */
    @Export(name="dnsName", type=String.class, parameters={})
    private Output<String> dnsName;

    /**
     * @return The domain name of the ALB instance. **NOTE:** Available in v1.158.0+.
     * 
     */
    public Output<String> dnsName() {
        return this.dnsName;
    }
    /**
     * Specifies whether to precheck the API request. Valid values: `true` and `false`.
     * 
     */
    @Export(name="dryRun", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> dryRun;

    /**
     * @return Specifies whether to precheck the API request. Valid values: `true` and `false`.
     * 
     */
    public Output<Optional<Boolean>> dryRun() {
        return Codegen.optional(this.dryRun);
    }
    /**
     * The configuration of the billing method. See the following `Block load_balancer_billing_config`.
     * 
     */
    @Export(name="loadBalancerBillingConfig", type=LoadBalancerLoadBalancerBillingConfig.class, parameters={})
    private Output<LoadBalancerLoadBalancerBillingConfig> loadBalancerBillingConfig;

    /**
     * @return The configuration of the billing method. See the following `Block load_balancer_billing_config`.
     * 
     */
    public Output<LoadBalancerLoadBalancerBillingConfig> loadBalancerBillingConfig() {
        return this.loadBalancerBillingConfig;
    }
    /**
     * The edition of the ALB instance. Different editions have different limits and billing methods. Valid values: `Basic`, `Standard` and `StandardWithWaf`(Available in v1.193.1+).
     * 
     */
    @Export(name="loadBalancerEdition", type=String.class, parameters={})
    private Output<String> loadBalancerEdition;

    /**
     * @return The edition of the ALB instance. Different editions have different limits and billing methods. Valid values: `Basic`, `Standard` and `StandardWithWaf`(Available in v1.193.1+).
     * 
     */
    public Output<String> loadBalancerEdition() {
        return this.loadBalancerEdition;
    }
    /**
     * The name of the resource.
     * 
     */
    @Export(name="loadBalancerName", type=String.class, parameters={})
    private Output<String> loadBalancerName;

    /**
     * @return The name of the resource.
     * 
     */
    public Output<String> loadBalancerName() {
        return this.loadBalancerName;
    }
    /**
     * Modify the Protection Configuration. See the following `Block modification_protection_config`.
     * 
     */
    @Export(name="modificationProtectionConfig", type=LoadBalancerModificationProtectionConfig.class, parameters={})
    private Output<LoadBalancerModificationProtectionConfig> modificationProtectionConfig;

    /**
     * @return Modify the Protection Configuration. See the following `Block modification_protection_config`.
     * 
     */
    public Output<LoadBalancerModificationProtectionConfig> modificationProtectionConfig() {
        return this.modificationProtectionConfig;
    }
    /**
     * The ID of the resource group.
     * 
     */
    @Export(name="resourceGroupId", type=String.class, parameters={})
    private Output<String> resourceGroupId;

    /**
     * @return The ID of the resource group.
     * 
     */
    public Output<String> resourceGroupId() {
        return this.resourceGroupId;
    }
    /**
     * Specifies whether to enable the configuration read-only mode for the ALB instance. Valid values: `NonProtection` and `ConsoleProtection`.
     * 
     */
    @Export(name="status", type=String.class, parameters={})
    private Output<String> status;

    /**
     * @return Specifies whether to enable the configuration read-only mode for the ALB instance. Valid values: `NonProtection` and `ConsoleProtection`.
     * 
     */
    public Output<String> status() {
        return this.status;
    }
    /**
     * A mapping of tags to assign to the resource. **NOTE:** The Key of `tags` cannot begin with &#34;aliyun&#34;, &#34;acs:&#34;, &#34;http://&#34;, &#34;https://&#34;, &#34;ack&#34; or &#34;ingress&#34;.
     * 
     */
    @Export(name="tags", type=Map.class, parameters={String.class, Object.class})
    private Output</* @Nullable */ Map<String,Object>> tags;

    /**
     * @return A mapping of tags to assign to the resource. **NOTE:** The Key of `tags` cannot begin with &#34;aliyun&#34;, &#34;acs:&#34;, &#34;http://&#34;, &#34;https://&#34;, &#34;ack&#34; or &#34;ingress&#34;.
     * 
     */
    public Output<Optional<Map<String,Object>>> tags() {
        return Codegen.optional(this.tags);
    }
    /**
     * The ID of the virtual private cloud (VPC) where the ALB instance is deployed.
     * 
     */
    @Export(name="vpcId", type=String.class, parameters={})
    private Output<String> vpcId;

    /**
     * @return The ID of the virtual private cloud (VPC) where the ALB instance is deployed.
     * 
     */
    public Output<String> vpcId() {
        return this.vpcId;
    }
    /**
     * The zones and vSwitches. You must specify at least two zones. See the following `Block zone_mappings`.
     * 
     */
    @Export(name="zoneMappings", type=List.class, parameters={LoadBalancerZoneMapping.class})
    private Output<List<LoadBalancerZoneMapping>> zoneMappings;

    /**
     * @return The zones and vSwitches. You must specify at least two zones. See the following `Block zone_mappings`.
     * 
     */
    public Output<List<LoadBalancerZoneMapping>> zoneMappings() {
        return this.zoneMappings;
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
    public LoadBalancer(String name, LoadBalancerArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public LoadBalancer(String name, LoadBalancerArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:alb/loadBalancer:LoadBalancer", name, args == null ? LoadBalancerArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private LoadBalancer(String name, Output<String> id, @Nullable LoadBalancerState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:alb/loadBalancer:LoadBalancer", name, state, makeResourceOptions(options, id));
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
