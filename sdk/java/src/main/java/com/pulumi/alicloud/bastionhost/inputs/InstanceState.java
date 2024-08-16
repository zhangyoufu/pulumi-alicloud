// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.bastionhost.inputs;

import com.pulumi.alicloud.bastionhost.inputs.InstanceAdAuthServerArgs;
import com.pulumi.alicloud.bastionhost.inputs.InstanceLdapAuthServerArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class InstanceState extends com.pulumi.resources.ResourceArgs {

    public static final InstanceState Empty = new InstanceState();

    /**
     * The AD auth server of the Instance. See `ad_auth_server` below.
     * 
     */
    @Import(name="adAuthServers")
    private @Nullable Output<List<InstanceAdAuthServerArgs>> adAuthServers;

    /**
     * @return The AD auth server of the Instance. See `ad_auth_server` below.
     * 
     */
    public Optional<Output<List<InstanceAdAuthServerArgs>>> adAuthServers() {
        return Optional.ofNullable(this.adAuthServers);
    }

    /**
     * The bandwidth of Cloud Bastionhost instance.
     * If China-Site Account, its valid values: 0 to 150. Unit: Mbit/s. The value must be a multiple of 5.
     * If International-Site Account, its valid values: 0 to 200. Unit: Mbit/s. The value must be a multiple of 10.
     * 
     */
    @Import(name="bandwidth")
    private @Nullable Output<String> bandwidth;

    /**
     * @return The bandwidth of Cloud Bastionhost instance.
     * If China-Site Account, its valid values: 0 to 150. Unit: Mbit/s. The value must be a multiple of 5.
     * If International-Site Account, its valid values: 0 to 200. Unit: Mbit/s. The value must be a multiple of 10.
     * 
     */
    public Optional<Output<String>> bandwidth() {
        return Optional.ofNullable(this.bandwidth);
    }

    /**
     * Description of the instance. This name can have a string of 1 to 63 characters.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return Description of the instance. This name can have a string of 1 to 63 characters.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
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
     * The LDAP auth server of the Instance. See `ldap_auth_server` below.
     * 
     */
    @Import(name="ldapAuthServers")
    private @Nullable Output<List<InstanceLdapAuthServerArgs>> ldapAuthServers;

    /**
     * @return The LDAP auth server of the Instance. See `ldap_auth_server` below.
     * 
     */
    public Optional<Output<List<InstanceLdapAuthServerArgs>>> ldapAuthServers() {
        return Optional.ofNullable(this.ldapAuthServers);
    }

    /**
     * The package type of Cloud Bastionhost instance. You can query more supported types through the [DescribePricingModule](https://help.aliyun.com/document_detail/96469.html).
     * 
     */
    @Import(name="licenseCode")
    private @Nullable Output<String> licenseCode;

    /**
     * @return The package type of Cloud Bastionhost instance. You can query more supported types through the [DescribePricingModule](https://help.aliyun.com/document_detail/96469.html).
     * 
     */
    public Optional<Output<String>> licenseCode() {
        return Optional.ofNullable(this.licenseCode);
    }

    /**
     * Duration for initially producing the instance. Valid values: [1~9], 12, 24, 36. At present, the provider does not support modify &#34;period&#34;.
     * &gt; **NOTE:** The attribute `period` is only used to create Subscription instance or modify the PayAsYouGo instance to Subscription. Once effect, it will not be modified that means running `pulumi up` will not effect the resource.
     * 
     */
    @Import(name="period")
    private @Nullable Output<Integer> period;

    /**
     * @return Duration for initially producing the instance. Valid values: [1~9], 12, 24, 36. At present, the provider does not support modify &#34;period&#34;.
     * &gt; **NOTE:** The attribute `period` is only used to create Subscription instance or modify the PayAsYouGo instance to Subscription. Once effect, it will not be modified that means running `pulumi up` will not effect the resource.
     * 
     */
    public Optional<Output<Integer>> period() {
        return Optional.ofNullable(this.period);
    }

    /**
     * The plan code of Cloud Bastionhost instance. Valid values:
     * 
     */
    @Import(name="planCode")
    private @Nullable Output<String> planCode;

    /**
     * @return The plan code of Cloud Bastionhost instance. Valid values:
     * 
     */
    public Optional<Output<String>> planCode() {
        return Optional.ofNullable(this.planCode);
    }

    @Import(name="publicWhiteLists")
    private @Nullable Output<List<String>> publicWhiteLists;

    public Optional<Output<List<String>>> publicWhiteLists() {
        return Optional.ofNullable(this.publicWhiteLists);
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
     * - `M`: months.
     * - `Y`: years.
     * 
     */
    @Import(name="renewalPeriodUnit")
    private @Nullable Output<String> renewalPeriodUnit;

    /**
     * @return The unit of the auto-renewal period. Valid values:  **NOTE:** The `renewal_period_unit` is required under the condition that `renewal_status` is `AutoRenewal`.
     * - `M`: months.
     * - `Y`: years.
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

    /**
     * security group IDs configured to Bastionhost.
     * **NOTE:** There is a potential diff error because of the order of `security_group_ids` values indefinite.
     * So, from version 1.160.0, `security_group_ids` type has been updated as `set` from `list`,
     * and you can use tolist to convert it to a list.
     * 
     */
    @Import(name="securityGroupIds")
    private @Nullable Output<List<String>> securityGroupIds;

    /**
     * @return security group IDs configured to Bastionhost.
     * **NOTE:** There is a potential diff error because of the order of `security_group_ids` values indefinite.
     * So, from version 1.160.0, `security_group_ids` type has been updated as `set` from `list`,
     * and you can use tolist to convert it to a list.
     * 
     */
    public Optional<Output<List<String>>> securityGroupIds() {
        return Optional.ofNullable(this.securityGroupIds);
    }

    /**
     * The storage of Cloud Bastionhost instance. Valid values: 0 to 500. Unit: TB.
     * 
     */
    @Import(name="storage")
    private @Nullable Output<String> storage;

    /**
     * @return The storage of Cloud Bastionhost instance. Valid values: 0 to 500. Unit: TB.
     * 
     */
    public Optional<Output<String>> storage() {
        return Optional.ofNullable(this.storage);
    }

    /**
     * A mapping of tags to assign to the resource.
     * 
     */
    @Import(name="tags")
    private @Nullable Output<Map<String,String>> tags;

    /**
     * @return A mapping of tags to assign to the resource.
     * 
     */
    public Optional<Output<Map<String,String>>> tags() {
        return Optional.ofNullable(this.tags);
    }

    /**
     * VSwitch ID configured to Bastionhost.
     * 
     */
    @Import(name="vswitchId")
    private @Nullable Output<String> vswitchId;

    /**
     * @return VSwitch ID configured to Bastionhost.
     * 
     */
    public Optional<Output<String>> vswitchId() {
        return Optional.ofNullable(this.vswitchId);
    }

    private InstanceState() {}

    private InstanceState(InstanceState $) {
        this.adAuthServers = $.adAuthServers;
        this.bandwidth = $.bandwidth;
        this.description = $.description;
        this.enablePublicAccess = $.enablePublicAccess;
        this.ldapAuthServers = $.ldapAuthServers;
        this.licenseCode = $.licenseCode;
        this.period = $.period;
        this.planCode = $.planCode;
        this.publicWhiteLists = $.publicWhiteLists;
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
    public static Builder builder(InstanceState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private InstanceState $;

        public Builder() {
            $ = new InstanceState();
        }

        public Builder(InstanceState defaults) {
            $ = new InstanceState(Objects.requireNonNull(defaults));
        }

        /**
         * @param adAuthServers The AD auth server of the Instance. See `ad_auth_server` below.
         * 
         * @return builder
         * 
         */
        public Builder adAuthServers(@Nullable Output<List<InstanceAdAuthServerArgs>> adAuthServers) {
            $.adAuthServers = adAuthServers;
            return this;
        }

        /**
         * @param adAuthServers The AD auth server of the Instance. See `ad_auth_server` below.
         * 
         * @return builder
         * 
         */
        public Builder adAuthServers(List<InstanceAdAuthServerArgs> adAuthServers) {
            return adAuthServers(Output.of(adAuthServers));
        }

        /**
         * @param adAuthServers The AD auth server of the Instance. See `ad_auth_server` below.
         * 
         * @return builder
         * 
         */
        public Builder adAuthServers(InstanceAdAuthServerArgs... adAuthServers) {
            return adAuthServers(List.of(adAuthServers));
        }

        /**
         * @param bandwidth The bandwidth of Cloud Bastionhost instance.
         * If China-Site Account, its valid values: 0 to 150. Unit: Mbit/s. The value must be a multiple of 5.
         * If International-Site Account, its valid values: 0 to 200. Unit: Mbit/s. The value must be a multiple of 10.
         * 
         * @return builder
         * 
         */
        public Builder bandwidth(@Nullable Output<String> bandwidth) {
            $.bandwidth = bandwidth;
            return this;
        }

        /**
         * @param bandwidth The bandwidth of Cloud Bastionhost instance.
         * If China-Site Account, its valid values: 0 to 150. Unit: Mbit/s. The value must be a multiple of 5.
         * If International-Site Account, its valid values: 0 to 200. Unit: Mbit/s. The value must be a multiple of 10.
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
        public Builder description(@Nullable Output<String> description) {
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
         * @param ldapAuthServers The LDAP auth server of the Instance. See `ldap_auth_server` below.
         * 
         * @return builder
         * 
         */
        public Builder ldapAuthServers(@Nullable Output<List<InstanceLdapAuthServerArgs>> ldapAuthServers) {
            $.ldapAuthServers = ldapAuthServers;
            return this;
        }

        /**
         * @param ldapAuthServers The LDAP auth server of the Instance. See `ldap_auth_server` below.
         * 
         * @return builder
         * 
         */
        public Builder ldapAuthServers(List<InstanceLdapAuthServerArgs> ldapAuthServers) {
            return ldapAuthServers(Output.of(ldapAuthServers));
        }

        /**
         * @param ldapAuthServers The LDAP auth server of the Instance. See `ldap_auth_server` below.
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
        public Builder licenseCode(@Nullable Output<String> licenseCode) {
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

        /**
         * @param period Duration for initially producing the instance. Valid values: [1~9], 12, 24, 36. At present, the provider does not support modify &#34;period&#34;.
         * &gt; **NOTE:** The attribute `period` is only used to create Subscription instance or modify the PayAsYouGo instance to Subscription. Once effect, it will not be modified that means running `pulumi up` will not effect the resource.
         * 
         * @return builder
         * 
         */
        public Builder period(@Nullable Output<Integer> period) {
            $.period = period;
            return this;
        }

        /**
         * @param period Duration for initially producing the instance. Valid values: [1~9], 12, 24, 36. At present, the provider does not support modify &#34;period&#34;.
         * &gt; **NOTE:** The attribute `period` is only used to create Subscription instance or modify the PayAsYouGo instance to Subscription. Once effect, it will not be modified that means running `pulumi up` will not effect the resource.
         * 
         * @return builder
         * 
         */
        public Builder period(Integer period) {
            return period(Output.of(period));
        }

        /**
         * @param planCode The plan code of Cloud Bastionhost instance. Valid values:
         * 
         * @return builder
         * 
         */
        public Builder planCode(@Nullable Output<String> planCode) {
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

        public Builder publicWhiteLists(@Nullable Output<List<String>> publicWhiteLists) {
            $.publicWhiteLists = publicWhiteLists;
            return this;
        }

        public Builder publicWhiteLists(List<String> publicWhiteLists) {
            return publicWhiteLists(Output.of(publicWhiteLists));
        }

        public Builder publicWhiteLists(String... publicWhiteLists) {
            return publicWhiteLists(List.of(publicWhiteLists));
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
         * - `M`: months.
         * - `Y`: years.
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
         * - `M`: months.
         * - `Y`: years.
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

        /**
         * @param securityGroupIds security group IDs configured to Bastionhost.
         * **NOTE:** There is a potential diff error because of the order of `security_group_ids` values indefinite.
         * So, from version 1.160.0, `security_group_ids` type has been updated as `set` from `list`,
         * and you can use tolist to convert it to a list.
         * 
         * @return builder
         * 
         */
        public Builder securityGroupIds(@Nullable Output<List<String>> securityGroupIds) {
            $.securityGroupIds = securityGroupIds;
            return this;
        }

        /**
         * @param securityGroupIds security group IDs configured to Bastionhost.
         * **NOTE:** There is a potential diff error because of the order of `security_group_ids` values indefinite.
         * So, from version 1.160.0, `security_group_ids` type has been updated as `set` from `list`,
         * and you can use tolist to convert it to a list.
         * 
         * @return builder
         * 
         */
        public Builder securityGroupIds(List<String> securityGroupIds) {
            return securityGroupIds(Output.of(securityGroupIds));
        }

        /**
         * @param securityGroupIds security group IDs configured to Bastionhost.
         * **NOTE:** There is a potential diff error because of the order of `security_group_ids` values indefinite.
         * So, from version 1.160.0, `security_group_ids` type has been updated as `set` from `list`,
         * and you can use tolist to convert it to a list.
         * 
         * @return builder
         * 
         */
        public Builder securityGroupIds(String... securityGroupIds) {
            return securityGroupIds(List.of(securityGroupIds));
        }

        /**
         * @param storage The storage of Cloud Bastionhost instance. Valid values: 0 to 500. Unit: TB.
         * 
         * @return builder
         * 
         */
        public Builder storage(@Nullable Output<String> storage) {
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
        public Builder tags(@Nullable Output<Map<String,String>> tags) {
            $.tags = tags;
            return this;
        }

        /**
         * @param tags A mapping of tags to assign to the resource.
         * 
         * @return builder
         * 
         */
        public Builder tags(Map<String,String> tags) {
            return tags(Output.of(tags));
        }

        /**
         * @param vswitchId VSwitch ID configured to Bastionhost.
         * 
         * @return builder
         * 
         */
        public Builder vswitchId(@Nullable Output<String> vswitchId) {
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

        public InstanceState build() {
            return $;
        }
    }

}
