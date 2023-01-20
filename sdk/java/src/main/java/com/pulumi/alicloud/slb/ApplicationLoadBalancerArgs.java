// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.slb;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.Object;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ApplicationLoadBalancerArgs extends com.pulumi.resources.ResourceArgs {

    public static final ApplicationLoadBalancerArgs Empty = new ApplicationLoadBalancerArgs();

    /**
     * Specify the IP address of the private network for the SLB instance, which must be in the destination CIDR block of the corresponding switch.
     * 
     */
    @Import(name="address")
    private @Nullable Output<String> address;

    /**
     * @return Specify the IP address of the private network for the SLB instance, which must be in the destination CIDR block of the corresponding switch.
     * 
     */
    public Optional<Output<String>> address() {
        return Optional.ofNullable(this.address);
    }

    /**
     * The IP version of the SLB instance to be created, which can be set to `ipv4` or `ipv6` . Default to `ipv4`. Now, only internet instance support `ipv6` address.
     * 
     */
    @Import(name="addressIpVersion")
    private @Nullable Output<String> addressIpVersion;

    /**
     * @return The IP version of the SLB instance to be created, which can be set to `ipv4` or `ipv6` . Default to `ipv4`. Now, only internet instance support `ipv6` address.
     * 
     */
    public Optional<Output<String>> addressIpVersion() {
        return Optional.ofNullable(this.addressIpVersion);
    }

    /**
     * The network type of the SLB instance. Valid values: [&#34;internet&#34;, &#34;intranet&#34;]. If load balancer launched in VPC, this value must be `intranet`.
     * - internet: After an Internet SLB instance is created, the system allocates a public IP address so that the instance can forward requests from the Internet.
     * - intranet: After an intranet SLB instance is created, the system allocates an intranet IP address so that the instance can only forward intranet requests.
     * 
     */
    @Import(name="addressType")
    private @Nullable Output<String> addressType;

    /**
     * @return The network type of the SLB instance. Valid values: [&#34;internet&#34;, &#34;intranet&#34;]. If load balancer launched in VPC, this value must be `intranet`.
     * - internet: After an Internet SLB instance is created, the system allocates a public IP address so that the instance can forward requests from the Internet.
     * - intranet: After an intranet SLB instance is created, the system allocates an intranet IP address so that the instance can only forward intranet requests.
     * 
     */
    public Optional<Output<String>> addressType() {
        return Optional.ofNullable(this.addressType);
    }

    /**
     * Valid value is between 1 and 5120, If argument `internet_charge_type` is `PayByTraffic`, then this value will be ignored.
     * 
     */
    @Import(name="bandwidth")
    private @Nullable Output<Integer> bandwidth;

    /**
     * @return Valid value is between 1 and 5120, If argument `internet_charge_type` is `PayByTraffic`, then this value will be ignored.
     * 
     */
    public Optional<Output<Integer>> bandwidth() {
        return Optional.ofNullable(this.bandwidth);
    }

    /**
     * Whether enable the deletion protection or not. on: Enable deletion protection. off: Disable deletion protection. Default to off. Only postpaid instance support this function.
     * 
     */
    @Import(name="deleteProtection")
    private @Nullable Output<String> deleteProtection;

    /**
     * @return Whether enable the deletion protection or not. on: Enable deletion protection. off: Disable deletion protection. Default to off. Only postpaid instance support this function.
     * 
     */
    public Optional<Output<String>> deleteProtection() {
        return Optional.ofNullable(this.deleteProtection);
    }

    /**
     * Support `PayBySpec` (default) and `PayByCLCU`, This parameter takes effect when the value of **payment_type** (instance payment mode) is **PayAsYouGo** (pay-as-you-go).
     * 
     */
    @Import(name="instanceChargeType")
    private @Nullable Output<String> instanceChargeType;

    /**
     * @return Support `PayBySpec` (default) and `PayByCLCU`, This parameter takes effect when the value of **payment_type** (instance payment mode) is **PayAsYouGo** (pay-as-you-go).
     * 
     */
    public Optional<Output<String>> instanceChargeType() {
        return Optional.ofNullable(this.instanceChargeType);
    }

    /**
     * Valid values are `PayByBandwidth`, `PayByTraffic`. If this value is `PayByBandwidth`, then argument `address_type` must be `internet`. Default is `PayByTraffic`. If load balancer launched in VPC, this value must be `PayByTraffic`. Before version 1.10.1, the valid values are `paybybandwidth` and `paybytraffic`.
     * 
     */
    @Import(name="internetChargeType")
    private @Nullable Output<String> internetChargeType;

    /**
     * @return Valid values are `PayByBandwidth`, `PayByTraffic`. If this value is `PayByBandwidth`, then argument `address_type` must be `internet`. Default is `PayByTraffic`. If load balancer launched in VPC, this value must be `PayByTraffic`. Before version 1.10.1, the valid values are `paybybandwidth` and `paybytraffic`.
     * 
     */
    public Optional<Output<String>> internetChargeType() {
        return Optional.ofNullable(this.internetChargeType);
    }

    @Import(name="loadBalancerName")
    private @Nullable Output<String> loadBalancerName;

    public Optional<Output<String>> loadBalancerName() {
        return Optional.ofNullable(this.loadBalancerName);
    }

    /**
     * The specification of the Server Load Balancer instance. Default to empty string indicating it is &#34;Shared-Performance&#34; instance.
     * Launching &#34;[Performance-guaranteed](https://www.alibabacloud.com/help/doc-detail/27657.htm)&#34; instance, it must be specified. Valid values: `slb.s1.small`, `slb.s2.small`, `slb.s2.medium`,
     * `slb.s3.small`, `slb.s3.medium`, `slb.s3.large` and `slb.s4.large`. It will be ignored when `instance_charge_type = &#34;PayByCLCU&#34;`.
     * 
     */
    @Import(name="loadBalancerSpec")
    private @Nullable Output<String> loadBalancerSpec;

    /**
     * @return The specification of the Server Load Balancer instance. Default to empty string indicating it is &#34;Shared-Performance&#34; instance.
     * Launching &#34;[Performance-guaranteed](https://www.alibabacloud.com/help/doc-detail/27657.htm)&#34; instance, it must be specified. Valid values: `slb.s1.small`, `slb.s2.small`, `slb.s2.medium`,
     * `slb.s3.small`, `slb.s3.medium`, `slb.s3.large` and `slb.s4.large`. It will be ignored when `instance_charge_type = &#34;PayByCLCU&#34;`.
     * 
     */
    public Optional<Output<String>> loadBalancerSpec() {
        return Optional.ofNullable(this.loadBalancerSpec);
    }

    /**
     * The primary zone ID of the SLB instance. If not specified, the system will be randomly assigned. You can query the primary and standby zones in a region by calling the [DescribeZone](https://help.aliyun.com/document_detail/27585.htm) API.
     * 
     */
    @Import(name="masterZoneId")
    private @Nullable Output<String> masterZoneId;

    /**
     * @return The primary zone ID of the SLB instance. If not specified, the system will be randomly assigned. You can query the primary and standby zones in a region by calling the [DescribeZone](https://help.aliyun.com/document_detail/27585.htm) API.
     * 
     */
    public Optional<Output<String>> masterZoneId() {
        return Optional.ofNullable(this.masterZoneId);
    }

    /**
     * The reason of modification protection. It&#39;s effective when `modification_protection_status` is `ConsoleProtection`.
     * 
     */
    @Import(name="modificationProtectionReason")
    private @Nullable Output<String> modificationProtectionReason;

    /**
     * @return The reason of modification protection. It&#39;s effective when `modification_protection_status` is `ConsoleProtection`.
     * 
     */
    public Optional<Output<String>> modificationProtectionReason() {
        return Optional.ofNullable(this.modificationProtectionReason);
    }

    /**
     * The status of modification protection. Valid values: `ConsoleProtection` and `NonProtection`. Default value is `NonProtection`.
     * 
     */
    @Import(name="modificationProtectionStatus")
    private @Nullable Output<String> modificationProtectionStatus;

    /**
     * @return The status of modification protection. Valid values: `ConsoleProtection` and `NonProtection`. Default value is `NonProtection`.
     * 
     */
    public Optional<Output<String>> modificationProtectionStatus() {
        return Optional.ofNullable(this.modificationProtectionStatus);
    }

    /**
     * Field `name` has been deprecated from provider version 1.123.1 New field `load_balancer_name` instead.
     * 
     * @deprecated
     * Field &#39;name&#39; has been deprecated from provider version 1.123.1. New field &#39;load_balancer_name&#39; instead
     * 
     */
    @Deprecated /* Field 'name' has been deprecated from provider version 1.123.1. New field 'load_balancer_name' instead */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return Field `name` has been deprecated from provider version 1.123.1 New field `load_balancer_name` instead.
     * 
     * @deprecated
     * Field &#39;name&#39; has been deprecated from provider version 1.123.1. New field &#39;load_balancer_name&#39; instead
     * 
     */
    @Deprecated /* Field 'name' has been deprecated from provider version 1.123.1. New field 'load_balancer_name' instead */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * The billing method of the load balancer. Valid values are `PayAsYouGo` and `Subscription`. Default to `PayAsYouGo`.
     * 
     */
    @Import(name="paymentType")
    private @Nullable Output<String> paymentType;

    /**
     * @return The billing method of the load balancer. Valid values are `PayAsYouGo` and `Subscription`. Default to `PayAsYouGo`.
     * 
     */
    public Optional<Output<String>> paymentType() {
        return Optional.ofNullable(this.paymentType);
    }

    @Import(name="period")
    private @Nullable Output<Integer> period;

    public Optional<Output<Integer>> period() {
        return Optional.ofNullable(this.period);
    }

    /**
     * The id of resource group which the SLB belongs.
     * 
     */
    @Import(name="resourceGroupId")
    private @Nullable Output<String> resourceGroupId;

    /**
     * @return The id of resource group which the SLB belongs.
     * 
     */
    public Optional<Output<String>> resourceGroupId() {
        return Optional.ofNullable(this.resourceGroupId);
    }

    /**
     * The standby zone ID of the SLB instance. If not specified, the system will be randomly assigned. You can query the primary and standby zones in a region by calling the DescribeZone API.
     * 
     */
    @Import(name="slaveZoneId")
    private @Nullable Output<String> slaveZoneId;

    /**
     * @return The standby zone ID of the SLB instance. If not specified, the system will be randomly assigned. You can query the primary and standby zones in a region by calling the DescribeZone API.
     * 
     */
    public Optional<Output<String>> slaveZoneId() {
        return Optional.ofNullable(this.slaveZoneId);
    }

    /**
     * Field `specification` has been deprecated from provider version 1.123.1 New field `load_balancer_spec` instead.
     * 
     * @deprecated
     * Field &#39;specification&#39; has been deprecated from provider version 1.123.1. New field &#39;load_balancer_spec&#39; instead
     * 
     */
    @Deprecated /* Field 'specification' has been deprecated from provider version 1.123.1. New field 'load_balancer_spec' instead */
    @Import(name="specification")
    private @Nullable Output<String> specification;

    /**
     * @return Field `specification` has been deprecated from provider version 1.123.1 New field `load_balancer_spec` instead.
     * 
     * @deprecated
     * Field &#39;specification&#39; has been deprecated from provider version 1.123.1. New field &#39;load_balancer_spec&#39; instead
     * 
     */
    @Deprecated /* Field 'specification' has been deprecated from provider version 1.123.1. New field 'load_balancer_spec' instead */
    public Optional<Output<String>> specification() {
        return Optional.ofNullable(this.specification);
    }

    /**
     * The status of slb load balancer. Valid values: `active` and `inactice`. The system default value is `active`.
     * 
     */
    @Import(name="status")
    private @Nullable Output<String> status;

    /**
     * @return The status of slb load balancer. Valid values: `active` and `inactice`. The system default value is `active`.
     * 
     */
    public Optional<Output<String>> status() {
        return Optional.ofNullable(this.status);
    }

    /**
     * A mapping of tags to assign to the resource. The `tags` can have a maximum of 10 tag for every load balancer instance.
     * 
     */
    @Import(name="tags")
    private @Nullable Output<Map<String,Object>> tags;

    /**
     * @return A mapping of tags to assign to the resource. The `tags` can have a maximum of 10 tag for every load balancer instance.
     * 
     */
    public Optional<Output<Map<String,Object>>> tags() {
        return Optional.ofNullable(this.tags);
    }

    /**
     * The VSwitch ID to launch in. **Note:** Required for a VPC SLB. If `address_type` is internet, it will be ignored.
     * 
     */
    @Import(name="vswitchId")
    private @Nullable Output<String> vswitchId;

    /**
     * @return The VSwitch ID to launch in. **Note:** Required for a VPC SLB. If `address_type` is internet, it will be ignored.
     * 
     */
    public Optional<Output<String>> vswitchId() {
        return Optional.ofNullable(this.vswitchId);
    }

    private ApplicationLoadBalancerArgs() {}

    private ApplicationLoadBalancerArgs(ApplicationLoadBalancerArgs $) {
        this.address = $.address;
        this.addressIpVersion = $.addressIpVersion;
        this.addressType = $.addressType;
        this.bandwidth = $.bandwidth;
        this.deleteProtection = $.deleteProtection;
        this.instanceChargeType = $.instanceChargeType;
        this.internetChargeType = $.internetChargeType;
        this.loadBalancerName = $.loadBalancerName;
        this.loadBalancerSpec = $.loadBalancerSpec;
        this.masterZoneId = $.masterZoneId;
        this.modificationProtectionReason = $.modificationProtectionReason;
        this.modificationProtectionStatus = $.modificationProtectionStatus;
        this.name = $.name;
        this.paymentType = $.paymentType;
        this.period = $.period;
        this.resourceGroupId = $.resourceGroupId;
        this.slaveZoneId = $.slaveZoneId;
        this.specification = $.specification;
        this.status = $.status;
        this.tags = $.tags;
        this.vswitchId = $.vswitchId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ApplicationLoadBalancerArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ApplicationLoadBalancerArgs $;

        public Builder() {
            $ = new ApplicationLoadBalancerArgs();
        }

        public Builder(ApplicationLoadBalancerArgs defaults) {
            $ = new ApplicationLoadBalancerArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param address Specify the IP address of the private network for the SLB instance, which must be in the destination CIDR block of the corresponding switch.
         * 
         * @return builder
         * 
         */
        public Builder address(@Nullable Output<String> address) {
            $.address = address;
            return this;
        }

        /**
         * @param address Specify the IP address of the private network for the SLB instance, which must be in the destination CIDR block of the corresponding switch.
         * 
         * @return builder
         * 
         */
        public Builder address(String address) {
            return address(Output.of(address));
        }

        /**
         * @param addressIpVersion The IP version of the SLB instance to be created, which can be set to `ipv4` or `ipv6` . Default to `ipv4`. Now, only internet instance support `ipv6` address.
         * 
         * @return builder
         * 
         */
        public Builder addressIpVersion(@Nullable Output<String> addressIpVersion) {
            $.addressIpVersion = addressIpVersion;
            return this;
        }

        /**
         * @param addressIpVersion The IP version of the SLB instance to be created, which can be set to `ipv4` or `ipv6` . Default to `ipv4`. Now, only internet instance support `ipv6` address.
         * 
         * @return builder
         * 
         */
        public Builder addressIpVersion(String addressIpVersion) {
            return addressIpVersion(Output.of(addressIpVersion));
        }

        /**
         * @param addressType The network type of the SLB instance. Valid values: [&#34;internet&#34;, &#34;intranet&#34;]. If load balancer launched in VPC, this value must be `intranet`.
         * - internet: After an Internet SLB instance is created, the system allocates a public IP address so that the instance can forward requests from the Internet.
         * - intranet: After an intranet SLB instance is created, the system allocates an intranet IP address so that the instance can only forward intranet requests.
         * 
         * @return builder
         * 
         */
        public Builder addressType(@Nullable Output<String> addressType) {
            $.addressType = addressType;
            return this;
        }

        /**
         * @param addressType The network type of the SLB instance. Valid values: [&#34;internet&#34;, &#34;intranet&#34;]. If load balancer launched in VPC, this value must be `intranet`.
         * - internet: After an Internet SLB instance is created, the system allocates a public IP address so that the instance can forward requests from the Internet.
         * - intranet: After an intranet SLB instance is created, the system allocates an intranet IP address so that the instance can only forward intranet requests.
         * 
         * @return builder
         * 
         */
        public Builder addressType(String addressType) {
            return addressType(Output.of(addressType));
        }

        /**
         * @param bandwidth Valid value is between 1 and 5120, If argument `internet_charge_type` is `PayByTraffic`, then this value will be ignored.
         * 
         * @return builder
         * 
         */
        public Builder bandwidth(@Nullable Output<Integer> bandwidth) {
            $.bandwidth = bandwidth;
            return this;
        }

        /**
         * @param bandwidth Valid value is between 1 and 5120, If argument `internet_charge_type` is `PayByTraffic`, then this value will be ignored.
         * 
         * @return builder
         * 
         */
        public Builder bandwidth(Integer bandwidth) {
            return bandwidth(Output.of(bandwidth));
        }

        /**
         * @param deleteProtection Whether enable the deletion protection or not. on: Enable deletion protection. off: Disable deletion protection. Default to off. Only postpaid instance support this function.
         * 
         * @return builder
         * 
         */
        public Builder deleteProtection(@Nullable Output<String> deleteProtection) {
            $.deleteProtection = deleteProtection;
            return this;
        }

        /**
         * @param deleteProtection Whether enable the deletion protection or not. on: Enable deletion protection. off: Disable deletion protection. Default to off. Only postpaid instance support this function.
         * 
         * @return builder
         * 
         */
        public Builder deleteProtection(String deleteProtection) {
            return deleteProtection(Output.of(deleteProtection));
        }

        /**
         * @param instanceChargeType Support `PayBySpec` (default) and `PayByCLCU`, This parameter takes effect when the value of **payment_type** (instance payment mode) is **PayAsYouGo** (pay-as-you-go).
         * 
         * @return builder
         * 
         */
        public Builder instanceChargeType(@Nullable Output<String> instanceChargeType) {
            $.instanceChargeType = instanceChargeType;
            return this;
        }

        /**
         * @param instanceChargeType Support `PayBySpec` (default) and `PayByCLCU`, This parameter takes effect when the value of **payment_type** (instance payment mode) is **PayAsYouGo** (pay-as-you-go).
         * 
         * @return builder
         * 
         */
        public Builder instanceChargeType(String instanceChargeType) {
            return instanceChargeType(Output.of(instanceChargeType));
        }

        /**
         * @param internetChargeType Valid values are `PayByBandwidth`, `PayByTraffic`. If this value is `PayByBandwidth`, then argument `address_type` must be `internet`. Default is `PayByTraffic`. If load balancer launched in VPC, this value must be `PayByTraffic`. Before version 1.10.1, the valid values are `paybybandwidth` and `paybytraffic`.
         * 
         * @return builder
         * 
         */
        public Builder internetChargeType(@Nullable Output<String> internetChargeType) {
            $.internetChargeType = internetChargeType;
            return this;
        }

        /**
         * @param internetChargeType Valid values are `PayByBandwidth`, `PayByTraffic`. If this value is `PayByBandwidth`, then argument `address_type` must be `internet`. Default is `PayByTraffic`. If load balancer launched in VPC, this value must be `PayByTraffic`. Before version 1.10.1, the valid values are `paybybandwidth` and `paybytraffic`.
         * 
         * @return builder
         * 
         */
        public Builder internetChargeType(String internetChargeType) {
            return internetChargeType(Output.of(internetChargeType));
        }

        public Builder loadBalancerName(@Nullable Output<String> loadBalancerName) {
            $.loadBalancerName = loadBalancerName;
            return this;
        }

        public Builder loadBalancerName(String loadBalancerName) {
            return loadBalancerName(Output.of(loadBalancerName));
        }

        /**
         * @param loadBalancerSpec The specification of the Server Load Balancer instance. Default to empty string indicating it is &#34;Shared-Performance&#34; instance.
         * Launching &#34;[Performance-guaranteed](https://www.alibabacloud.com/help/doc-detail/27657.htm)&#34; instance, it must be specified. Valid values: `slb.s1.small`, `slb.s2.small`, `slb.s2.medium`,
         * `slb.s3.small`, `slb.s3.medium`, `slb.s3.large` and `slb.s4.large`. It will be ignored when `instance_charge_type = &#34;PayByCLCU&#34;`.
         * 
         * @return builder
         * 
         */
        public Builder loadBalancerSpec(@Nullable Output<String> loadBalancerSpec) {
            $.loadBalancerSpec = loadBalancerSpec;
            return this;
        }

        /**
         * @param loadBalancerSpec The specification of the Server Load Balancer instance. Default to empty string indicating it is &#34;Shared-Performance&#34; instance.
         * Launching &#34;[Performance-guaranteed](https://www.alibabacloud.com/help/doc-detail/27657.htm)&#34; instance, it must be specified. Valid values: `slb.s1.small`, `slb.s2.small`, `slb.s2.medium`,
         * `slb.s3.small`, `slb.s3.medium`, `slb.s3.large` and `slb.s4.large`. It will be ignored when `instance_charge_type = &#34;PayByCLCU&#34;`.
         * 
         * @return builder
         * 
         */
        public Builder loadBalancerSpec(String loadBalancerSpec) {
            return loadBalancerSpec(Output.of(loadBalancerSpec));
        }

        /**
         * @param masterZoneId The primary zone ID of the SLB instance. If not specified, the system will be randomly assigned. You can query the primary and standby zones in a region by calling the [DescribeZone](https://help.aliyun.com/document_detail/27585.htm) API.
         * 
         * @return builder
         * 
         */
        public Builder masterZoneId(@Nullable Output<String> masterZoneId) {
            $.masterZoneId = masterZoneId;
            return this;
        }

        /**
         * @param masterZoneId The primary zone ID of the SLB instance. If not specified, the system will be randomly assigned. You can query the primary and standby zones in a region by calling the [DescribeZone](https://help.aliyun.com/document_detail/27585.htm) API.
         * 
         * @return builder
         * 
         */
        public Builder masterZoneId(String masterZoneId) {
            return masterZoneId(Output.of(masterZoneId));
        }

        /**
         * @param modificationProtectionReason The reason of modification protection. It&#39;s effective when `modification_protection_status` is `ConsoleProtection`.
         * 
         * @return builder
         * 
         */
        public Builder modificationProtectionReason(@Nullable Output<String> modificationProtectionReason) {
            $.modificationProtectionReason = modificationProtectionReason;
            return this;
        }

        /**
         * @param modificationProtectionReason The reason of modification protection. It&#39;s effective when `modification_protection_status` is `ConsoleProtection`.
         * 
         * @return builder
         * 
         */
        public Builder modificationProtectionReason(String modificationProtectionReason) {
            return modificationProtectionReason(Output.of(modificationProtectionReason));
        }

        /**
         * @param modificationProtectionStatus The status of modification protection. Valid values: `ConsoleProtection` and `NonProtection`. Default value is `NonProtection`.
         * 
         * @return builder
         * 
         */
        public Builder modificationProtectionStatus(@Nullable Output<String> modificationProtectionStatus) {
            $.modificationProtectionStatus = modificationProtectionStatus;
            return this;
        }

        /**
         * @param modificationProtectionStatus The status of modification protection. Valid values: `ConsoleProtection` and `NonProtection`. Default value is `NonProtection`.
         * 
         * @return builder
         * 
         */
        public Builder modificationProtectionStatus(String modificationProtectionStatus) {
            return modificationProtectionStatus(Output.of(modificationProtectionStatus));
        }

        /**
         * @param name Field `name` has been deprecated from provider version 1.123.1 New field `load_balancer_name` instead.
         * 
         * @return builder
         * 
         * @deprecated
         * Field &#39;name&#39; has been deprecated from provider version 1.123.1. New field &#39;load_balancer_name&#39; instead
         * 
         */
        @Deprecated /* Field 'name' has been deprecated from provider version 1.123.1. New field 'load_balancer_name' instead */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name Field `name` has been deprecated from provider version 1.123.1 New field `load_balancer_name` instead.
         * 
         * @return builder
         * 
         * @deprecated
         * Field &#39;name&#39; has been deprecated from provider version 1.123.1. New field &#39;load_balancer_name&#39; instead
         * 
         */
        @Deprecated /* Field 'name' has been deprecated from provider version 1.123.1. New field 'load_balancer_name' instead */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param paymentType The billing method of the load balancer. Valid values are `PayAsYouGo` and `Subscription`. Default to `PayAsYouGo`.
         * 
         * @return builder
         * 
         */
        public Builder paymentType(@Nullable Output<String> paymentType) {
            $.paymentType = paymentType;
            return this;
        }

        /**
         * @param paymentType The billing method of the load balancer. Valid values are `PayAsYouGo` and `Subscription`. Default to `PayAsYouGo`.
         * 
         * @return builder
         * 
         */
        public Builder paymentType(String paymentType) {
            return paymentType(Output.of(paymentType));
        }

        public Builder period(@Nullable Output<Integer> period) {
            $.period = period;
            return this;
        }

        public Builder period(Integer period) {
            return period(Output.of(period));
        }

        /**
         * @param resourceGroupId The id of resource group which the SLB belongs.
         * 
         * @return builder
         * 
         */
        public Builder resourceGroupId(@Nullable Output<String> resourceGroupId) {
            $.resourceGroupId = resourceGroupId;
            return this;
        }

        /**
         * @param resourceGroupId The id of resource group which the SLB belongs.
         * 
         * @return builder
         * 
         */
        public Builder resourceGroupId(String resourceGroupId) {
            return resourceGroupId(Output.of(resourceGroupId));
        }

        /**
         * @param slaveZoneId The standby zone ID of the SLB instance. If not specified, the system will be randomly assigned. You can query the primary and standby zones in a region by calling the DescribeZone API.
         * 
         * @return builder
         * 
         */
        public Builder slaveZoneId(@Nullable Output<String> slaveZoneId) {
            $.slaveZoneId = slaveZoneId;
            return this;
        }

        /**
         * @param slaveZoneId The standby zone ID of the SLB instance. If not specified, the system will be randomly assigned. You can query the primary and standby zones in a region by calling the DescribeZone API.
         * 
         * @return builder
         * 
         */
        public Builder slaveZoneId(String slaveZoneId) {
            return slaveZoneId(Output.of(slaveZoneId));
        }

        /**
         * @param specification Field `specification` has been deprecated from provider version 1.123.1 New field `load_balancer_spec` instead.
         * 
         * @return builder
         * 
         * @deprecated
         * Field &#39;specification&#39; has been deprecated from provider version 1.123.1. New field &#39;load_balancer_spec&#39; instead
         * 
         */
        @Deprecated /* Field 'specification' has been deprecated from provider version 1.123.1. New field 'load_balancer_spec' instead */
        public Builder specification(@Nullable Output<String> specification) {
            $.specification = specification;
            return this;
        }

        /**
         * @param specification Field `specification` has been deprecated from provider version 1.123.1 New field `load_balancer_spec` instead.
         * 
         * @return builder
         * 
         * @deprecated
         * Field &#39;specification&#39; has been deprecated from provider version 1.123.1. New field &#39;load_balancer_spec&#39; instead
         * 
         */
        @Deprecated /* Field 'specification' has been deprecated from provider version 1.123.1. New field 'load_balancer_spec' instead */
        public Builder specification(String specification) {
            return specification(Output.of(specification));
        }

        /**
         * @param status The status of slb load balancer. Valid values: `active` and `inactice`. The system default value is `active`.
         * 
         * @return builder
         * 
         */
        public Builder status(@Nullable Output<String> status) {
            $.status = status;
            return this;
        }

        /**
         * @param status The status of slb load balancer. Valid values: `active` and `inactice`. The system default value is `active`.
         * 
         * @return builder
         * 
         */
        public Builder status(String status) {
            return status(Output.of(status));
        }

        /**
         * @param tags A mapping of tags to assign to the resource. The `tags` can have a maximum of 10 tag for every load balancer instance.
         * 
         * @return builder
         * 
         */
        public Builder tags(@Nullable Output<Map<String,Object>> tags) {
            $.tags = tags;
            return this;
        }

        /**
         * @param tags A mapping of tags to assign to the resource. The `tags` can have a maximum of 10 tag for every load balancer instance.
         * 
         * @return builder
         * 
         */
        public Builder tags(Map<String,Object> tags) {
            return tags(Output.of(tags));
        }

        /**
         * @param vswitchId The VSwitch ID to launch in. **Note:** Required for a VPC SLB. If `address_type` is internet, it will be ignored.
         * 
         * @return builder
         * 
         */
        public Builder vswitchId(@Nullable Output<String> vswitchId) {
            $.vswitchId = vswitchId;
            return this;
        }

        /**
         * @param vswitchId The VSwitch ID to launch in. **Note:** Required for a VPC SLB. If `address_type` is internet, it will be ignored.
         * 
         * @return builder
         * 
         */
        public Builder vswitchId(String vswitchId) {
            return vswitchId(Output.of(vswitchId));
        }

        public ApplicationLoadBalancerArgs build() {
            return $;
        }
    }

}
