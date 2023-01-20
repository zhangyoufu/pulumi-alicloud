// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.bastionhost;

import com.pulumi.alicloud.bastionhost.inputs.InstanceAdAuthServerArgs;
import com.pulumi.alicloud.bastionhost.inputs.InstanceLdapAuthServerArgs;
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


public final class InstanceArgs extends com.pulumi.resources.ResourceArgs {

    public static final InstanceArgs Empty = new InstanceArgs();

    /**
     * The AD auth server of the Instance. See the following `Block ad_auth_server`.
     * 
     */
    @Import(name="adAuthServers")
    private @Nullable Output<List<InstanceAdAuthServerArgs>> adAuthServers;

    /**
     * @return The AD auth server of the Instance. See the following `Block ad_auth_server`.
     * 
     */
    public Optional<Output<List<InstanceAdAuthServerArgs>>> adAuthServers() {
        return Optional.ofNullable(this.adAuthServers);
    }

    /**
     * The bandwidth of Cloud Bastionhost instance. Valid values: 0 to 500. Unit: Mbit/s.
     * 
     */
    @Import(name="bandwidth", required=true)
    private Output<String> bandwidth;

    /**
     * @return The bandwidth of Cloud Bastionhost instance. Valid values: 0 to 500. Unit: Mbit/s.
     * 
     */
    public Output<String> bandwidth() {
        return this.bandwidth;
    }

    /**
     * Description of the instance. This name can have a string of 1 to 63 characters.
     * 
     */
    @Import(name="description", required=true)
    private Output<String> description;

    /**
     * @return Description of the instance. This name can have a string of 1 to 63 characters.
     * 
     */
    public Output<String> description() {
        return this.description;
    }

    /**
     * Whether to Enable the public internet access to a specified Bastionhost instance. The valid values: `true`, `false`.
     * 
     */
    @Import(name="enablePublicAccess")
    private @Nullable Output<Boolean> enablePublicAccess;

    /**
     * @return Whether to Enable the public internet access to a specified Bastionhost instance. The valid values: `true`, `false`.
     * 
     */
    public Optional<Output<Boolean>> enablePublicAccess() {
        return Optional.ofNullable(this.enablePublicAccess);
    }

    /**
     * The LDAP auth server of the Instance. See the following `Block ldap_auth_server`.
     * 
     */
    @Import(name="ldapAuthServers")
    private @Nullable Output<List<InstanceLdapAuthServerArgs>> ldapAuthServers;

    /**
     * @return The LDAP auth server of the Instance. See the following `Block ldap_auth_server`.
     * 
     */
    public Optional<Output<List<InstanceLdapAuthServerArgs>>> ldapAuthServers() {
        return Optional.ofNullable(this.ldapAuthServers);
    }

    /**
     * The package type of Cloud Bastionhost instance. You can query more supported types through the [DescribePricingModule](https://help.aliyun.com/document_detail/96469.html).
     * 
     */
    @Import(name="licenseCode", required=true)
    private Output<String> licenseCode;

    /**
     * @return The package type of Cloud Bastionhost instance. You can query more supported types through the [DescribePricingModule](https://help.aliyun.com/document_detail/96469.html).
     * 
     */
    public Output<String> licenseCode() {
        return this.licenseCode;
    }

    @Import(name="period")
    private @Nullable Output<Integer> period;

    public Optional<Output<Integer>> period() {
        return Optional.ofNullable(this.period);
    }

    /**
     * The plan code of Cloud Bastionhost instance. Valid values:
     * 
     */
    @Import(name="planCode", required=true)
    private Output<String> planCode;

    /**
     * @return The plan code of Cloud Bastionhost instance. Valid values:
     * 
     */
    public Output<String> planCode() {
        return this.planCode;
    }

    /**
     * Automatic renewal period. Valid values: `1` to `9`, `12`, `24`, `36`. **NOTE:** The `renew_period` is required under the condition that `renewal_status` is `AutoRenewal`. From version 1.193.0, `renew_period` can be modified.
     * 
     */
    @Import(name="renewPeriod")
    private @Nullable Output<Integer> renewPeriod;

    /**
     * @return Automatic renewal period. Valid values: `1` to `9`, `12`, `24`, `36`. **NOTE:** The `renew_period` is required under the condition that `renewal_status` is `AutoRenewal`. From version 1.193.0, `renew_period` can be modified.
     * 
     */
    public Optional<Output<Integer>> renewPeriod() {
        return Optional.ofNullable(this.renewPeriod);
    }

    /**
     * The unit of the auto-renewal period. Valid values:  **NOTE:** The `renewal_period_unit` is required under the condition that `renewal_status` is `AutoRenewal`.
     * 
     */
    @Import(name="renewalPeriodUnit")
    private @Nullable Output<String> renewalPeriodUnit;

    /**
     * @return The unit of the auto-renewal period. Valid values:  **NOTE:** The `renewal_period_unit` is required under the condition that `renewal_status` is `AutoRenewal`.
     * 
     */
    public Optional<Output<String>> renewalPeriodUnit() {
        return Optional.ofNullable(this.renewalPeriodUnit);
    }

    /**
     * Automatic renewal status. Valid values: `AutoRenewal`, `ManualRenewal`, `NotRenewal`. From version 1.193.0, `renewal_status` can be modified.
     * 
     */
    @Import(name="renewalStatus")
    private @Nullable Output<String> renewalStatus;

    /**
     * @return Automatic renewal status. Valid values: `AutoRenewal`, `ManualRenewal`, `NotRenewal`. From version 1.193.0, `renewal_status` can be modified.
     * 
     */
    public Optional<Output<String>> renewalStatus() {
        return Optional.ofNullable(this.renewalStatus);
    }

    /**
     * The Id of resource group which the Bastionhost Instance belongs. If not set, the resource is created in the default resource group.
     * 
     */
    @Import(name="resourceGroupId")
    private @Nullable Output<String> resourceGroupId;

    /**
     * @return The Id of resource group which the Bastionhost Instance belongs. If not set, the resource is created in the default resource group.
     * 
     */
    public Optional<Output<String>> resourceGroupId() {
        return Optional.ofNullable(this.resourceGroupId);
    }

    @Import(name="securityGroupIds", required=true)
    private Output<List<String>> securityGroupIds;

    public Output<List<String>> securityGroupIds() {
        return this.securityGroupIds;
    }

    /**
     * The storage of Cloud Bastionhost instance. Valid values: 0 to 500. Unit: TB.
     * 
     */
    @Import(name="storage", required=true)
    private Output<String> storage;

    /**
     * @return The storage of Cloud Bastionhost instance. Valid values: 0 to 500. Unit: TB.
     * 
     */
    public Output<String> storage() {
        return this.storage;
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

    /**
     * VSwitch ID configured to Bastionhost.
     * 
     */
    @Import(name="vswitchId", required=true)
    private Output<String> vswitchId;

    /**
     * @return VSwitch ID configured to Bastionhost.
     * 
     */
    public Output<String> vswitchId() {
        return this.vswitchId;
    }

    private InstanceArgs() {}

    private InstanceArgs(InstanceArgs $) {
        this.adAuthServers = $.adAuthServers;
        this.bandwidth = $.bandwidth;
        this.description = $.description;
        this.enablePublicAccess = $.enablePublicAccess;
        this.ldapAuthServers = $.ldapAuthServers;
        this.licenseCode = $.licenseCode;
        this.period = $.period;
        this.planCode = $.planCode;
        this.renewPeriod = $.renewPeriod;
        this.renewalPeriodUnit = $.renewalPeriodUnit;
        this.renewalStatus = $.renewalStatus;
        this.resourceGroupId = $.resourceGroupId;
        this.securityGroupIds = $.securityGroupIds;
        this.storage = $.storage;
        this.tags = $.tags;
        this.vswitchId = $.vswitchId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(InstanceArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private InstanceArgs $;

        public Builder() {
            $ = new InstanceArgs();
        }

        public Builder(InstanceArgs defaults) {
            $ = new InstanceArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param adAuthServers The AD auth server of the Instance. See the following `Block ad_auth_server`.
         * 
         * @return builder
         * 
         */
        public Builder adAuthServers(@Nullable Output<List<InstanceAdAuthServerArgs>> adAuthServers) {
            $.adAuthServers = adAuthServers;
            return this;
        }

        /**
         * @param adAuthServers The AD auth server of the Instance. See the following `Block ad_auth_server`.
         * 
         * @return builder
         * 
         */
        public Builder adAuthServers(List<InstanceAdAuthServerArgs> adAuthServers) {
            return adAuthServers(Output.of(adAuthServers));
        }

        /**
         * @param adAuthServers The AD auth server of the Instance. See the following `Block ad_auth_server`.
         * 
         * @return builder
         * 
         */
        public Builder adAuthServers(InstanceAdAuthServerArgs... adAuthServers) {
            return adAuthServers(List.of(adAuthServers));
        }

        /**
         * @param bandwidth The bandwidth of Cloud Bastionhost instance. Valid values: 0 to 500. Unit: Mbit/s.
         * 
         * @return builder
         * 
         */
        public Builder bandwidth(Output<String> bandwidth) {
            $.bandwidth = bandwidth;
            return this;
        }

        /**
         * @param bandwidth The bandwidth of Cloud Bastionhost instance. Valid values: 0 to 500. Unit: Mbit/s.
         * 
         * @return builder
         * 
         */
        public Builder bandwidth(String bandwidth) {
            return bandwidth(Output.of(bandwidth));
        }

        /**
         * @param description Description of the instance. This name can have a string of 1 to 63 characters.
         * 
         * @return builder
         * 
         */
        public Builder description(Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description Description of the instance. This name can have a string of 1 to 63 characters.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param enablePublicAccess Whether to Enable the public internet access to a specified Bastionhost instance. The valid values: `true`, `false`.
         * 
         * @return builder
         * 
         */
        public Builder enablePublicAccess(@Nullable Output<Boolean> enablePublicAccess) {
            $.enablePublicAccess = enablePublicAccess;
            return this;
        }

        /**
         * @param enablePublicAccess Whether to Enable the public internet access to a specified Bastionhost instance. The valid values: `true`, `false`.
         * 
         * @return builder
         * 
         */
        public Builder enablePublicAccess(Boolean enablePublicAccess) {
            return enablePublicAccess(Output.of(enablePublicAccess));
        }

        /**
         * @param ldapAuthServers The LDAP auth server of the Instance. See the following `Block ldap_auth_server`.
         * 
         * @return builder
         * 
         */
        public Builder ldapAuthServers(@Nullable Output<List<InstanceLdapAuthServerArgs>> ldapAuthServers) {
            $.ldapAuthServers = ldapAuthServers;
            return this;
        }

        /**
         * @param ldapAuthServers The LDAP auth server of the Instance. See the following `Block ldap_auth_server`.
         * 
         * @return builder
         * 
         */
        public Builder ldapAuthServers(List<InstanceLdapAuthServerArgs> ldapAuthServers) {
            return ldapAuthServers(Output.of(ldapAuthServers));
        }

        /**
         * @param ldapAuthServers The LDAP auth server of the Instance. See the following `Block ldap_auth_server`.
         * 
         * @return builder
         * 
         */
        public Builder ldapAuthServers(InstanceLdapAuthServerArgs... ldapAuthServers) {
            return ldapAuthServers(List.of(ldapAuthServers));
        }

        /**
         * @param licenseCode The package type of Cloud Bastionhost instance. You can query more supported types through the [DescribePricingModule](https://help.aliyun.com/document_detail/96469.html).
         * 
         * @return builder
         * 
         */
        public Builder licenseCode(Output<String> licenseCode) {
            $.licenseCode = licenseCode;
            return this;
        }

        /**
         * @param licenseCode The package type of Cloud Bastionhost instance. You can query more supported types through the [DescribePricingModule](https://help.aliyun.com/document_detail/96469.html).
         * 
         * @return builder
         * 
         */
        public Builder licenseCode(String licenseCode) {
            return licenseCode(Output.of(licenseCode));
        }

        public Builder period(@Nullable Output<Integer> period) {
            $.period = period;
            return this;
        }

        public Builder period(Integer period) {
            return period(Output.of(period));
        }

        /**
         * @param planCode The plan code of Cloud Bastionhost instance. Valid values:
         * 
         * @return builder
         * 
         */
        public Builder planCode(Output<String> planCode) {
            $.planCode = planCode;
            return this;
        }

        /**
         * @param planCode The plan code of Cloud Bastionhost instance. Valid values:
         * 
         * @return builder
         * 
         */
        public Builder planCode(String planCode) {
            return planCode(Output.of(planCode));
        }

        /**
         * @param renewPeriod Automatic renewal period. Valid values: `1` to `9`, `12`, `24`, `36`. **NOTE:** The `renew_period` is required under the condition that `renewal_status` is `AutoRenewal`. From version 1.193.0, `renew_period` can be modified.
         * 
         * @return builder
         * 
         */
        public Builder renewPeriod(@Nullable Output<Integer> renewPeriod) {
            $.renewPeriod = renewPeriod;
            return this;
        }

        /**
         * @param renewPeriod Automatic renewal period. Valid values: `1` to `9`, `12`, `24`, `36`. **NOTE:** The `renew_period` is required under the condition that `renewal_status` is `AutoRenewal`. From version 1.193.0, `renew_period` can be modified.
         * 
         * @return builder
         * 
         */
        public Builder renewPeriod(Integer renewPeriod) {
            return renewPeriod(Output.of(renewPeriod));
        }

        /**
         * @param renewalPeriodUnit The unit of the auto-renewal period. Valid values:  **NOTE:** The `renewal_period_unit` is required under the condition that `renewal_status` is `AutoRenewal`.
         * 
         * @return builder
         * 
         */
        public Builder renewalPeriodUnit(@Nullable Output<String> renewalPeriodUnit) {
            $.renewalPeriodUnit = renewalPeriodUnit;
            return this;
        }

        /**
         * @param renewalPeriodUnit The unit of the auto-renewal period. Valid values:  **NOTE:** The `renewal_period_unit` is required under the condition that `renewal_status` is `AutoRenewal`.
         * 
         * @return builder
         * 
         */
        public Builder renewalPeriodUnit(String renewalPeriodUnit) {
            return renewalPeriodUnit(Output.of(renewalPeriodUnit));
        }

        /**
         * @param renewalStatus Automatic renewal status. Valid values: `AutoRenewal`, `ManualRenewal`, `NotRenewal`. From version 1.193.0, `renewal_status` can be modified.
         * 
         * @return builder
         * 
         */
        public Builder renewalStatus(@Nullable Output<String> renewalStatus) {
            $.renewalStatus = renewalStatus;
            return this;
        }

        /**
         * @param renewalStatus Automatic renewal status. Valid values: `AutoRenewal`, `ManualRenewal`, `NotRenewal`. From version 1.193.0, `renewal_status` can be modified.
         * 
         * @return builder
         * 
         */
        public Builder renewalStatus(String renewalStatus) {
            return renewalStatus(Output.of(renewalStatus));
        }

        /**
         * @param resourceGroupId The Id of resource group which the Bastionhost Instance belongs. If not set, the resource is created in the default resource group.
         * 
         * @return builder
         * 
         */
        public Builder resourceGroupId(@Nullable Output<String> resourceGroupId) {
            $.resourceGroupId = resourceGroupId;
            return this;
        }

        /**
         * @param resourceGroupId The Id of resource group which the Bastionhost Instance belongs. If not set, the resource is created in the default resource group.
         * 
         * @return builder
         * 
         */
        public Builder resourceGroupId(String resourceGroupId) {
            return resourceGroupId(Output.of(resourceGroupId));
        }

        public Builder securityGroupIds(Output<List<String>> securityGroupIds) {
            $.securityGroupIds = securityGroupIds;
            return this;
        }

        public Builder securityGroupIds(List<String> securityGroupIds) {
            return securityGroupIds(Output.of(securityGroupIds));
        }

        public Builder securityGroupIds(String... securityGroupIds) {
            return securityGroupIds(List.of(securityGroupIds));
        }

        /**
         * @param storage The storage of Cloud Bastionhost instance. Valid values: 0 to 500. Unit: TB.
         * 
         * @return builder
         * 
         */
        public Builder storage(Output<String> storage) {
            $.storage = storage;
            return this;
        }

        /**
         * @param storage The storage of Cloud Bastionhost instance. Valid values: 0 to 500. Unit: TB.
         * 
         * @return builder
         * 
         */
        public Builder storage(String storage) {
            return storage(Output.of(storage));
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

        /**
         * @param vswitchId VSwitch ID configured to Bastionhost.
         * 
         * @return builder
         * 
         */
        public Builder vswitchId(Output<String> vswitchId) {
            $.vswitchId = vswitchId;
            return this;
        }

        /**
         * @param vswitchId VSwitch ID configured to Bastionhost.
         * 
         * @return builder
         * 
         */
        public Builder vswitchId(String vswitchId) {
            return vswitchId(Output.of(vswitchId));
        }

        public InstanceArgs build() {
            $.bandwidth = Objects.requireNonNull($.bandwidth, "expected parameter 'bandwidth' to be non-null");
            $.description = Objects.requireNonNull($.description, "expected parameter 'description' to be non-null");
            $.licenseCode = Objects.requireNonNull($.licenseCode, "expected parameter 'licenseCode' to be non-null");
            $.planCode = Objects.requireNonNull($.planCode, "expected parameter 'planCode' to be non-null");
            $.securityGroupIds = Objects.requireNonNull($.securityGroupIds, "expected parameter 'securityGroupIds' to be non-null");
            $.storage = Objects.requireNonNull($.storage, "expected parameter 'storage' to be non-null");
            $.vswitchId = Objects.requireNonNull($.vswitchId, "expected parameter 'vswitchId' to be non-null");
            return $;
        }
    }

}
