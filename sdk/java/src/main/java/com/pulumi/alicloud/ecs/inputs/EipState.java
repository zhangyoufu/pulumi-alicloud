// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.ecs.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.Object;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class EipState extends com.pulumi.resources.ResourceArgs {

    public static final EipState Empty = new EipState();

    @Import(name="activityId")
    private @Nullable Output<String> activityId;

    public Optional<Output<String>> activityId() {
        return Optional.ofNullable(this.activityId);
    }

    /**
     * The name of the EIP instance. This name can have a string of 2 to 128 characters, must contain only alphanumeric characters or hyphens, such as &#34;-&#34;,&#34;.&#34;,&#34;_&#34;, and must not begin or end with a hyphen, and must not begin with http:// or https://.
     * 
     */
    @Import(name="addressName")
    private @Nullable Output<String> addressName;

    /**
     * @return The name of the EIP instance. This name can have a string of 2 to 128 characters, must contain only alphanumeric characters or hyphens, such as &#34;-&#34;,&#34;.&#34;,&#34;_&#34;, and must not begin or end with a hyphen, and must not begin with http:// or https://.
     * 
     */
    public Optional<Output<String>> addressName() {
        return Optional.ofNullable(this.addressName);
    }

    @Import(name="autoPay")
    private @Nullable Output<Boolean> autoPay;

    public Optional<Output<Boolean>> autoPay() {
        return Optional.ofNullable(this.autoPay);
    }

    /**
     * Maximum bandwidth to the elastic public network, measured in Mbps (Mega bit per second). If this value is not specified, then automatically sets it to 5 Mbps.
     * 
     */
    @Import(name="bandwidth")
    private @Nullable Output<String> bandwidth;

    /**
     * @return Maximum bandwidth to the elastic public network, measured in Mbps (Mega bit per second). If this value is not specified, then automatically sets it to 5 Mbps.
     * 
     */
    public Optional<Output<String>> bandwidth() {
        return Optional.ofNullable(this.bandwidth);
    }

    @Import(name="createTime")
    private @Nullable Output<String> createTime;

    public Optional<Output<String>> createTime() {
        return Optional.ofNullable(this.createTime);
    }

    /**
     * Whether enable the deletion protection or not. Default value: `false`.
     * - true: Enable deletion protection.
     * - false: Disable deletion protection.
     * 
     */
    @Import(name="deletionProtection")
    private @Nullable Output<Boolean> deletionProtection;

    /**
     * @return Whether enable the deletion protection or not. Default value: `false`.
     * - true: Enable deletion protection.
     * - false: Disable deletion protection.
     * 
     */
    public Optional<Output<Boolean>> deletionProtection() {
        return Optional.ofNullable(this.deletionProtection);
    }

    /**
     * Description of the EIP instance, This description can have a string of 2 to 256 characters, It cannot begin with http:// or https://. Default value is null.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return Description of the EIP instance, This description can have a string of 2 to 256 characters, It cannot begin with http:// or https://. Default value is null.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    @Import(name="highDefinitionMonitorLogStatus")
    private @Nullable Output<String> highDefinitionMonitorLogStatus;

    public Optional<Output<String>> highDefinitionMonitorLogStatus() {
        return Optional.ofNullable(this.highDefinitionMonitorLogStatus);
    }

    /**
     * (It has been deprecated from version 1.126.0 and using new attribute `payment_type` instead) Elastic IP instance charge type. Valid values are &#34;PrePaid&#34; and &#34;PostPaid&#34;. Default to &#34;PostPaid&#34;.
     * 
     * @deprecated
     * Field &#39;instance_charge_type&#39; has been deprecated since provider version 1.126.0. New field &#39;payment_type&#39; instead.
     * 
     */
    @Deprecated /* Field 'instance_charge_type' has been deprecated since provider version 1.126.0. New field 'payment_type' instead. */
    @Import(name="instanceChargeType")
    private @Nullable Output<String> instanceChargeType;

    /**
     * @return (It has been deprecated from version 1.126.0 and using new attribute `payment_type` instead) Elastic IP instance charge type. Valid values are &#34;PrePaid&#34; and &#34;PostPaid&#34;. Default to &#34;PostPaid&#34;.
     * 
     * @deprecated
     * Field &#39;instance_charge_type&#39; has been deprecated since provider version 1.126.0. New field &#39;payment_type&#39; instead.
     * 
     */
    @Deprecated /* Field 'instance_charge_type' has been deprecated since provider version 1.126.0. New field 'payment_type' instead. */
    public Optional<Output<String>> instanceChargeType() {
        return Optional.ofNullable(this.instanceChargeType);
    }

    /**
     * Internet charge type of the EIP, Valid values are `PayByBandwidth`, `PayByTraffic`. Default to `PayByBandwidth`. **NOTE:** From version `1.7.1` to `1.125.0`, it defaults to `PayByTraffic`. It is only &#34;PayByBandwidth&#34; when `instance_charge_type` is PrePaid.
     * 
     */
    @Import(name="internetChargeType")
    private @Nullable Output<String> internetChargeType;

    /**
     * @return Internet charge type of the EIP, Valid values are `PayByBandwidth`, `PayByTraffic`. Default to `PayByBandwidth`. **NOTE:** From version `1.7.1` to `1.125.0`, it defaults to `PayByTraffic`. It is only &#34;PayByBandwidth&#34; when `instance_charge_type` is PrePaid.
     * 
     */
    public Optional<Output<String>> internetChargeType() {
        return Optional.ofNullable(this.internetChargeType);
    }

    /**
     * The elastic ip address
     * 
     */
    @Import(name="ipAddress")
    private @Nullable Output<String> ipAddress;

    /**
     * @return The elastic ip address
     * 
     */
    public Optional<Output<String>> ipAddress() {
        return Optional.ofNullable(this.ipAddress);
    }

    /**
     * The line type of the Elastic IP instance. Default to `BGP`. Other type of the isp need to open a whitelist.
     * 
     */
    @Import(name="isp")
    private @Nullable Output<String> isp;

    /**
     * @return The line type of the Elastic IP instance. Default to `BGP`. Other type of the isp need to open a whitelist.
     * 
     */
    public Optional<Output<String>> isp() {
        return Optional.ofNullable(this.isp);
    }

    @Import(name="logProject")
    private @Nullable Output<String> logProject;

    public Optional<Output<String>> logProject() {
        return Optional.ofNullable(this.logProject);
    }

    @Import(name="logStore")
    private @Nullable Output<String> logStore;

    public Optional<Output<String>> logStore() {
        return Optional.ofNullable(this.logStore);
    }

    /**
     * It has been deprecated from version 1.126.0 and using new attribute `address_name` instead.
     * 
     * @deprecated
     * Field &#39;name&#39; has been deprecated since provider version 1.126.0. New field &#39;address_name&#39; instead.
     * 
     */
    @Deprecated /* Field 'name' has been deprecated since provider version 1.126.0. New field 'address_name' instead. */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return It has been deprecated from version 1.126.0 and using new attribute `address_name` instead.
     * 
     * @deprecated
     * Field &#39;name&#39; has been deprecated since provider version 1.126.0. New field &#39;address_name&#39; instead.
     * 
     */
    @Deprecated /* Field 'name' has been deprecated since provider version 1.126.0. New field 'address_name' instead. */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    @Import(name="netmode")
    private @Nullable Output<String> netmode;

    public Optional<Output<String>> netmode() {
        return Optional.ofNullable(this.netmode);
    }

    /**
     * The billing method of the EIP. Valid values: `Subscription` and `PayAsYouGo`. Default value is `PayAsYouGo`.
     * 
     */
    @Import(name="paymentType")
    private @Nullable Output<String> paymentType;

    /**
     * @return The billing method of the EIP. Valid values: `Subscription` and `PayAsYouGo`. Default value is `PayAsYouGo`.
     * 
     */
    public Optional<Output<String>> paymentType() {
        return Optional.ofNullable(this.paymentType);
    }

    /**
     * The duration that you will buy the resource, in month. It is valid when `instance_charge_type` is `PrePaid`. Valid values: [1-9, 12, 24, 36]. At present, the provider does not support modify &#34;period&#34; and you can do that via web console.
     * **NOTE:** The attribute `period` is only used to create Subscription instance or modify the PayAsYouGo instance to Subscription. Once effect, it will not be modified that means running `pulumi up` will not effect the resource.
     * 
     */
    @Import(name="period")
    private @Nullable Output<Integer> period;

    /**
     * @return The duration that you will buy the resource, in month. It is valid when `instance_charge_type` is `PrePaid`. Valid values: [1-9, 12, 24, 36]. At present, the provider does not support modify &#34;period&#34; and you can do that via web console.
     * **NOTE:** The attribute `period` is only used to create Subscription instance or modify the PayAsYouGo instance to Subscription. Once effect, it will not be modified that means running `pulumi up` will not effect the resource.
     * 
     */
    public Optional<Output<Integer>> period() {
        return Optional.ofNullable(this.period);
    }

    @Import(name="pricingCycle")
    private @Nullable Output<String> pricingCycle;

    public Optional<Output<String>> pricingCycle() {
        return Optional.ofNullable(this.pricingCycle);
    }

    @Import(name="publicIpAddressPoolId")
    private @Nullable Output<String> publicIpAddressPoolId;

    public Optional<Output<String>> publicIpAddressPoolId() {
        return Optional.ofNullable(this.publicIpAddressPoolId);
    }

    /**
     * The Id of resource group which the eip belongs.
     * 
     */
    @Import(name="resourceGroupId")
    private @Nullable Output<String> resourceGroupId;

    /**
     * @return The Id of resource group which the eip belongs.
     * 
     */
    public Optional<Output<String>> resourceGroupId() {
        return Optional.ofNullable(this.resourceGroupId);
    }

    @Import(name="securityProtectionTypes")
    private @Nullable Output<List<String>> securityProtectionTypes;

    public Optional<Output<List<String>>> securityProtectionTypes() {
        return Optional.ofNullable(this.securityProtectionTypes);
    }

    /**
     * The EIP current status.
     * 
     */
    @Import(name="status")
    private @Nullable Output<String> status;

    /**
     * @return The EIP current status.
     * 
     */
    public Optional<Output<String>> status() {
        return Optional.ofNullable(this.status);
    }

    /**
     * A mapping of tags to assign to the resource.
     * 
     */
    @Import(name="tags")
    private @Nullable Output<Map<String,Object>> tags;

    /**
     * @return A mapping of tags to assign to the resource.
     * 
     */
    public Optional<Output<Map<String,Object>>> tags() {
        return Optional.ofNullable(this.tags);
    }

    @Import(name="zone")
    private @Nullable Output<String> zone;

    public Optional<Output<String>> zone() {
        return Optional.ofNullable(this.zone);
    }

    private EipState() {}

    private EipState(EipState $) {
        this.activityId = $.activityId;
        this.addressName = $.addressName;
        this.autoPay = $.autoPay;
        this.bandwidth = $.bandwidth;
        this.createTime = $.createTime;
        this.deletionProtection = $.deletionProtection;
        this.description = $.description;
        this.highDefinitionMonitorLogStatus = $.highDefinitionMonitorLogStatus;
        this.instanceChargeType = $.instanceChargeType;
        this.internetChargeType = $.internetChargeType;
        this.ipAddress = $.ipAddress;
        this.isp = $.isp;
        this.logProject = $.logProject;
        this.logStore = $.logStore;
        this.name = $.name;
        this.netmode = $.netmode;
        this.paymentType = $.paymentType;
        this.period = $.period;
        this.pricingCycle = $.pricingCycle;
        this.publicIpAddressPoolId = $.publicIpAddressPoolId;
        this.resourceGroupId = $.resourceGroupId;
        this.securityProtectionTypes = $.securityProtectionTypes;
        this.status = $.status;
        this.tags = $.tags;
        this.zone = $.zone;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(EipState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private EipState $;

        public Builder() {
            $ = new EipState();
        }

        public Builder(EipState defaults) {
            $ = new EipState(Objects.requireNonNull(defaults));
        }

        public Builder activityId(@Nullable Output<String> activityId) {
            $.activityId = activityId;
            return this;
        }

        public Builder activityId(String activityId) {
            return activityId(Output.of(activityId));
        }

        /**
         * @param addressName The name of the EIP instance. This name can have a string of 2 to 128 characters, must contain only alphanumeric characters or hyphens, such as &#34;-&#34;,&#34;.&#34;,&#34;_&#34;, and must not begin or end with a hyphen, and must not begin with http:// or https://.
         * 
         * @return builder
         * 
         */
        public Builder addressName(@Nullable Output<String> addressName) {
            $.addressName = addressName;
            return this;
        }

        /**
         * @param addressName The name of the EIP instance. This name can have a string of 2 to 128 characters, must contain only alphanumeric characters or hyphens, such as &#34;-&#34;,&#34;.&#34;,&#34;_&#34;, and must not begin or end with a hyphen, and must not begin with http:// or https://.
         * 
         * @return builder
         * 
         */
        public Builder addressName(String addressName) {
            return addressName(Output.of(addressName));
        }

        public Builder autoPay(@Nullable Output<Boolean> autoPay) {
            $.autoPay = autoPay;
            return this;
        }

        public Builder autoPay(Boolean autoPay) {
            return autoPay(Output.of(autoPay));
        }

        /**
         * @param bandwidth Maximum bandwidth to the elastic public network, measured in Mbps (Mega bit per second). If this value is not specified, then automatically sets it to 5 Mbps.
         * 
         * @return builder
         * 
         */
        public Builder bandwidth(@Nullable Output<String> bandwidth) {
            $.bandwidth = bandwidth;
            return this;
        }

        /**
         * @param bandwidth Maximum bandwidth to the elastic public network, measured in Mbps (Mega bit per second). If this value is not specified, then automatically sets it to 5 Mbps.
         * 
         * @return builder
         * 
         */
        public Builder bandwidth(String bandwidth) {
            return bandwidth(Output.of(bandwidth));
        }

        public Builder createTime(@Nullable Output<String> createTime) {
            $.createTime = createTime;
            return this;
        }

        public Builder createTime(String createTime) {
            return createTime(Output.of(createTime));
        }

        /**
         * @param deletionProtection Whether enable the deletion protection or not. Default value: `false`.
         * - true: Enable deletion protection.
         * - false: Disable deletion protection.
         * 
         * @return builder
         * 
         */
        public Builder deletionProtection(@Nullable Output<Boolean> deletionProtection) {
            $.deletionProtection = deletionProtection;
            return this;
        }

        /**
         * @param deletionProtection Whether enable the deletion protection or not. Default value: `false`.
         * - true: Enable deletion protection.
         * - false: Disable deletion protection.
         * 
         * @return builder
         * 
         */
        public Builder deletionProtection(Boolean deletionProtection) {
            return deletionProtection(Output.of(deletionProtection));
        }

        /**
         * @param description Description of the EIP instance, This description can have a string of 2 to 256 characters, It cannot begin with http:// or https://. Default value is null.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description Description of the EIP instance, This description can have a string of 2 to 256 characters, It cannot begin with http:// or https://. Default value is null.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        public Builder highDefinitionMonitorLogStatus(@Nullable Output<String> highDefinitionMonitorLogStatus) {
            $.highDefinitionMonitorLogStatus = highDefinitionMonitorLogStatus;
            return this;
        }

        public Builder highDefinitionMonitorLogStatus(String highDefinitionMonitorLogStatus) {
            return highDefinitionMonitorLogStatus(Output.of(highDefinitionMonitorLogStatus));
        }

        /**
         * @param instanceChargeType (It has been deprecated from version 1.126.0 and using new attribute `payment_type` instead) Elastic IP instance charge type. Valid values are &#34;PrePaid&#34; and &#34;PostPaid&#34;. Default to &#34;PostPaid&#34;.
         * 
         * @return builder
         * 
         * @deprecated
         * Field &#39;instance_charge_type&#39; has been deprecated since provider version 1.126.0. New field &#39;payment_type&#39; instead.
         * 
         */
        @Deprecated /* Field 'instance_charge_type' has been deprecated since provider version 1.126.0. New field 'payment_type' instead. */
        public Builder instanceChargeType(@Nullable Output<String> instanceChargeType) {
            $.instanceChargeType = instanceChargeType;
            return this;
        }

        /**
         * @param instanceChargeType (It has been deprecated from version 1.126.0 and using new attribute `payment_type` instead) Elastic IP instance charge type. Valid values are &#34;PrePaid&#34; and &#34;PostPaid&#34;. Default to &#34;PostPaid&#34;.
         * 
         * @return builder
         * 
         * @deprecated
         * Field &#39;instance_charge_type&#39; has been deprecated since provider version 1.126.0. New field &#39;payment_type&#39; instead.
         * 
         */
        @Deprecated /* Field 'instance_charge_type' has been deprecated since provider version 1.126.0. New field 'payment_type' instead. */
        public Builder instanceChargeType(String instanceChargeType) {
            return instanceChargeType(Output.of(instanceChargeType));
        }

        /**
         * @param internetChargeType Internet charge type of the EIP, Valid values are `PayByBandwidth`, `PayByTraffic`. Default to `PayByBandwidth`. **NOTE:** From version `1.7.1` to `1.125.0`, it defaults to `PayByTraffic`. It is only &#34;PayByBandwidth&#34; when `instance_charge_type` is PrePaid.
         * 
         * @return builder
         * 
         */
        public Builder internetChargeType(@Nullable Output<String> internetChargeType) {
            $.internetChargeType = internetChargeType;
            return this;
        }

        /**
         * @param internetChargeType Internet charge type of the EIP, Valid values are `PayByBandwidth`, `PayByTraffic`. Default to `PayByBandwidth`. **NOTE:** From version `1.7.1` to `1.125.0`, it defaults to `PayByTraffic`. It is only &#34;PayByBandwidth&#34; when `instance_charge_type` is PrePaid.
         * 
         * @return builder
         * 
         */
        public Builder internetChargeType(String internetChargeType) {
            return internetChargeType(Output.of(internetChargeType));
        }

        /**
         * @param ipAddress The elastic ip address
         * 
         * @return builder
         * 
         */
        public Builder ipAddress(@Nullable Output<String> ipAddress) {
            $.ipAddress = ipAddress;
            return this;
        }

        /**
         * @param ipAddress The elastic ip address
         * 
         * @return builder
         * 
         */
        public Builder ipAddress(String ipAddress) {
            return ipAddress(Output.of(ipAddress));
        }

        /**
         * @param isp The line type of the Elastic IP instance. Default to `BGP`. Other type of the isp need to open a whitelist.
         * 
         * @return builder
         * 
         */
        public Builder isp(@Nullable Output<String> isp) {
            $.isp = isp;
            return this;
        }

        /**
         * @param isp The line type of the Elastic IP instance. Default to `BGP`. Other type of the isp need to open a whitelist.
         * 
         * @return builder
         * 
         */
        public Builder isp(String isp) {
            return isp(Output.of(isp));
        }

        public Builder logProject(@Nullable Output<String> logProject) {
            $.logProject = logProject;
            return this;
        }

        public Builder logProject(String logProject) {
            return logProject(Output.of(logProject));
        }

        public Builder logStore(@Nullable Output<String> logStore) {
            $.logStore = logStore;
            return this;
        }

        public Builder logStore(String logStore) {
            return logStore(Output.of(logStore));
        }

        /**
         * @param name It has been deprecated from version 1.126.0 and using new attribute `address_name` instead.
         * 
         * @return builder
         * 
         * @deprecated
         * Field &#39;name&#39; has been deprecated since provider version 1.126.0. New field &#39;address_name&#39; instead.
         * 
         */
        @Deprecated /* Field 'name' has been deprecated since provider version 1.126.0. New field 'address_name' instead. */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name It has been deprecated from version 1.126.0 and using new attribute `address_name` instead.
         * 
         * @return builder
         * 
         * @deprecated
         * Field &#39;name&#39; has been deprecated since provider version 1.126.0. New field &#39;address_name&#39; instead.
         * 
         */
        @Deprecated /* Field 'name' has been deprecated since provider version 1.126.0. New field 'address_name' instead. */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        public Builder netmode(@Nullable Output<String> netmode) {
            $.netmode = netmode;
            return this;
        }

        public Builder netmode(String netmode) {
            return netmode(Output.of(netmode));
        }

        /**
         * @param paymentType The billing method of the EIP. Valid values: `Subscription` and `PayAsYouGo`. Default value is `PayAsYouGo`.
         * 
         * @return builder
         * 
         */
        public Builder paymentType(@Nullable Output<String> paymentType) {
            $.paymentType = paymentType;
            return this;
        }

        /**
         * @param paymentType The billing method of the EIP. Valid values: `Subscription` and `PayAsYouGo`. Default value is `PayAsYouGo`.
         * 
         * @return builder
         * 
         */
        public Builder paymentType(String paymentType) {
            return paymentType(Output.of(paymentType));
        }

        /**
         * @param period The duration that you will buy the resource, in month. It is valid when `instance_charge_type` is `PrePaid`. Valid values: [1-9, 12, 24, 36]. At present, the provider does not support modify &#34;period&#34; and you can do that via web console.
         * **NOTE:** The attribute `period` is only used to create Subscription instance or modify the PayAsYouGo instance to Subscription. Once effect, it will not be modified that means running `pulumi up` will not effect the resource.
         * 
         * @return builder
         * 
         */
        public Builder period(@Nullable Output<Integer> period) {
            $.period = period;
            return this;
        }

        /**
         * @param period The duration that you will buy the resource, in month. It is valid when `instance_charge_type` is `PrePaid`. Valid values: [1-9, 12, 24, 36]. At present, the provider does not support modify &#34;period&#34; and you can do that via web console.
         * **NOTE:** The attribute `period` is only used to create Subscription instance or modify the PayAsYouGo instance to Subscription. Once effect, it will not be modified that means running `pulumi up` will not effect the resource.
         * 
         * @return builder
         * 
         */
        public Builder period(Integer period) {
            return period(Output.of(period));
        }

        public Builder pricingCycle(@Nullable Output<String> pricingCycle) {
            $.pricingCycle = pricingCycle;
            return this;
        }

        public Builder pricingCycle(String pricingCycle) {
            return pricingCycle(Output.of(pricingCycle));
        }

        public Builder publicIpAddressPoolId(@Nullable Output<String> publicIpAddressPoolId) {
            $.publicIpAddressPoolId = publicIpAddressPoolId;
            return this;
        }

        public Builder publicIpAddressPoolId(String publicIpAddressPoolId) {
            return publicIpAddressPoolId(Output.of(publicIpAddressPoolId));
        }

        /**
         * @param resourceGroupId The Id of resource group which the eip belongs.
         * 
         * @return builder
         * 
         */
        public Builder resourceGroupId(@Nullable Output<String> resourceGroupId) {
            $.resourceGroupId = resourceGroupId;
            return this;
        }

        /**
         * @param resourceGroupId The Id of resource group which the eip belongs.
         * 
         * @return builder
         * 
         */
        public Builder resourceGroupId(String resourceGroupId) {
            return resourceGroupId(Output.of(resourceGroupId));
        }

        public Builder securityProtectionTypes(@Nullable Output<List<String>> securityProtectionTypes) {
            $.securityProtectionTypes = securityProtectionTypes;
            return this;
        }

        public Builder securityProtectionTypes(List<String> securityProtectionTypes) {
            return securityProtectionTypes(Output.of(securityProtectionTypes));
        }

        public Builder securityProtectionTypes(String... securityProtectionTypes) {
            return securityProtectionTypes(List.of(securityProtectionTypes));
        }

        /**
         * @param status The EIP current status.
         * 
         * @return builder
         * 
         */
        public Builder status(@Nullable Output<String> status) {
            $.status = status;
            return this;
        }

        /**
         * @param status The EIP current status.
         * 
         * @return builder
         * 
         */
        public Builder status(String status) {
            return status(Output.of(status));
        }

        /**
         * @param tags A mapping of tags to assign to the resource.
         * 
         * @return builder
         * 
         */
        public Builder tags(@Nullable Output<Map<String,Object>> tags) {
            $.tags = tags;
            return this;
        }

        /**
         * @param tags A mapping of tags to assign to the resource.
         * 
         * @return builder
         * 
         */
        public Builder tags(Map<String,Object> tags) {
            return tags(Output.of(tags));
        }

        public Builder zone(@Nullable Output<String> zone) {
            $.zone = zone;
            return this;
        }

        public Builder zone(String zone) {
            return zone(Output.of(zone));
        }

        public EipState build() {
            return $;
        }
    }

}
