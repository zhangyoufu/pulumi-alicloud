// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.vpc.inputs;

import com.pulumi.alicloud.vpc.inputs.DhcpOptionsSetAssociateVpcArgs;
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


public final class DhcpOptionsSetState extends com.pulumi.resources.ResourceArgs {

    public static final DhcpOptionsSetState Empty = new DhcpOptionsSetState();

    /**
     * Field &#39;associate_vpcs&#39; has been deprecated from provider version 1.153.0 and it will be removed in the future version. Please use the new resource &#39;alicloud_vpc_dhcp_options_set_attachment&#39; to attach DhcpOptionsSet and Vpc. See `associate_vpcs` below.
     * 
     * @deprecated
     * Field &#39;associate_vpcs&#39; has been deprecated from provider version 1.211.0. Field &#39;associate_vpcs&#39; has been deprecated from provider version 1.153.0 and it will be removed in the future version. Please use the new resource &#39;alicloud_vpc_dhcp_options_set_attachment&#39; to attach DhcpOptionsSet and Vpc.
     * 
     */
    @Deprecated /* Field 'associate_vpcs' has been deprecated from provider version 1.211.0. Field 'associate_vpcs' has been deprecated from provider version 1.153.0 and it will be removed in the future version. Please use the new resource 'alicloud_vpc_dhcp_options_set_attachment' to attach DhcpOptionsSet and Vpc. */
    @Import(name="associateVpcs")
    private @Nullable Output<List<DhcpOptionsSetAssociateVpcArgs>> associateVpcs;

    /**
     * @return Field &#39;associate_vpcs&#39; has been deprecated from provider version 1.153.0 and it will be removed in the future version. Please use the new resource &#39;alicloud_vpc_dhcp_options_set_attachment&#39; to attach DhcpOptionsSet and Vpc. See `associate_vpcs` below.
     * 
     * @deprecated
     * Field &#39;associate_vpcs&#39; has been deprecated from provider version 1.211.0. Field &#39;associate_vpcs&#39; has been deprecated from provider version 1.153.0 and it will be removed in the future version. Please use the new resource &#39;alicloud_vpc_dhcp_options_set_attachment&#39; to attach DhcpOptionsSet and Vpc.
     * 
     */
    @Deprecated /* Field 'associate_vpcs' has been deprecated from provider version 1.211.0. Field 'associate_vpcs' has been deprecated from provider version 1.153.0 and it will be removed in the future version. Please use the new resource 'alicloud_vpc_dhcp_options_set_attachment' to attach DhcpOptionsSet and Vpc. */
    public Optional<Output<List<DhcpOptionsSetAssociateVpcArgs>>> associateVpcs() {
        return Optional.ofNullable(this.associateVpcs);
    }

    /**
     * The description can be blank or contain 1 to 256 characters. It must start with a letter or Chinese character but cannot start with http:// or https://.
     * 
     */
    @Import(name="dhcpOptionsSetDescription")
    private @Nullable Output<String> dhcpOptionsSetDescription;

    /**
     * @return The description can be blank or contain 1 to 256 characters. It must start with a letter or Chinese character but cannot start with http:// or https://.
     * 
     */
    public Optional<Output<String>> dhcpOptionsSetDescription() {
        return Optional.ofNullable(this.dhcpOptionsSetDescription);
    }

    /**
     * The name must be 2 to 128 characters in length and can contain letters, Chinese characters, digits, underscores (_), and hyphens (-). It must start with a letter or a Chinese character.
     * 
     */
    @Import(name="dhcpOptionsSetName")
    private @Nullable Output<String> dhcpOptionsSetName;

    /**
     * @return The name must be 2 to 128 characters in length and can contain letters, Chinese characters, digits, underscores (_), and hyphens (-). It must start with a letter or a Chinese character.
     * 
     */
    public Optional<Output<String>> dhcpOptionsSetName() {
        return Optional.ofNullable(this.dhcpOptionsSetName);
    }

    /**
     * The root domain, for example, example.com. After a DHCP options set is associated with a Virtual Private Cloud (VPC) network, the root domain in the DHCP options set is automatically synchronized to the ECS instances in the VPC network.
     * 
     */
    @Import(name="domainName")
    private @Nullable Output<String> domainName;

    /**
     * @return The root domain, for example, example.com. After a DHCP options set is associated with a Virtual Private Cloud (VPC) network, the root domain in the DHCP options set is automatically synchronized to the ECS instances in the VPC network.
     * 
     */
    public Optional<Output<String>> domainName() {
        return Optional.ofNullable(this.domainName);
    }

    /**
     * The DNS server IP addresses. Up to four DNS server IP addresses can be specified. IP addresses must be separated with commas (,).Before you specify any DNS server IP address, all ECS instances in the associated VPC network use the IP addresses of the Alibaba Cloud DNS servers, which are 100.100.2.136 and 100.100.2.138.
     * 
     */
    @Import(name="domainNameServers")
    private @Nullable Output<String> domainNameServers;

    /**
     * @return The DNS server IP addresses. Up to four DNS server IP addresses can be specified. IP addresses must be separated with commas (,).Before you specify any DNS server IP address, all ECS instances in the associated VPC network use the IP addresses of the Alibaba Cloud DNS servers, which are 100.100.2.136 and 100.100.2.138.
     * 
     */
    public Optional<Output<String>> domainNameServers() {
        return Optional.ofNullable(this.domainNameServers);
    }

    /**
     * Whether to PreCheck only this request, value:
     * - **true**: sends a check request and does not delete the DHCP option set. Check items include whether required parameters are filled in, request format, and restrictions. If the check fails, the corresponding error is returned. If the check passes, the error code &#39;DryRunOperation&#39; is returned &#39;.
     * - **false** (default): Sends a normal request and directly deletes the DHCP option set after checking.
     * 
     */
    @Import(name="dryRun")
    private @Nullable Output<Boolean> dryRun;

    /**
     * @return Whether to PreCheck only this request, value:
     * - **true**: sends a check request and does not delete the DHCP option set. Check items include whether required parameters are filled in, request format, and restrictions. If the check fails, the corresponding error is returned. If the check passes, the error code &#39;DryRunOperation&#39; is returned &#39;.
     * - **false** (default): Sends a normal request and directly deletes the DHCP option set after checking.
     * 
     */
    public Optional<Output<Boolean>> dryRun() {
        return Optional.ofNullable(this.dryRun);
    }

    /**
     * The lease time of the IPv6 DHCP option set.When the lease time is set to hours: Unit: h. Value range: 24h ~ 1176h,87600h ~ 175200h. Default value: 87600h.When the lease time is set to day: Unit: d. Value range: 1d ~ 49d,3650d ~ 7300d. Default value: 3650d.
     * 
     */
    @Import(name="ipv6LeaseTime")
    private @Nullable Output<String> ipv6LeaseTime;

    /**
     * @return The lease time of the IPv6 DHCP option set.When the lease time is set to hours: Unit: h. Value range: 24h ~ 1176h,87600h ~ 175200h. Default value: 87600h.When the lease time is set to day: Unit: d. Value range: 1d ~ 49d,3650d ~ 7300d. Default value: 3650d.
     * 
     */
    public Optional<Output<String>> ipv6LeaseTime() {
        return Optional.ofNullable(this.ipv6LeaseTime);
    }

    /**
     * The lease time of the IPv4 DHCP option set.When the lease time is set to hours: Unit: h. Value range: 24h ~ 1176h,87600h ~ 175200h. Default value: 87600h.When the lease time is set to day: Unit: d. Value range: 1d ~ 49d,3650d ~ 7300d. Default value: 3650d.
     * 
     */
    @Import(name="leaseTime")
    private @Nullable Output<String> leaseTime;

    /**
     * @return The lease time of the IPv4 DHCP option set.When the lease time is set to hours: Unit: h. Value range: 24h ~ 1176h,87600h ~ 175200h. Default value: 87600h.When the lease time is set to day: Unit: d. Value range: 1d ~ 49d,3650d ~ 7300d. Default value: 3650d.
     * 
     */
    public Optional<Output<String>> leaseTime() {
        return Optional.ofNullable(this.leaseTime);
    }

    /**
     * The ID of the account to which the DHCP options set belongs.
     * 
     */
    @Import(name="ownerId")
    private @Nullable Output<Integer> ownerId;

    /**
     * @return The ID of the account to which the DHCP options set belongs.
     * 
     */
    public Optional<Output<Integer>> ownerId() {
        return Optional.ofNullable(this.ownerId);
    }

    /**
     * The ID of the resource group.
     * 
     */
    @Import(name="resourceGroupId")
    private @Nullable Output<String> resourceGroupId;

    /**
     * @return The ID of the resource group.
     * 
     */
    public Optional<Output<String>> resourceGroupId() {
        return Optional.ofNullable(this.resourceGroupId);
    }

    /**
     * The status of the resource.
     * 
     */
    @Import(name="status")
    private @Nullable Output<String> status;

    /**
     * @return The status of the resource.
     * 
     */
    public Optional<Output<String>> status() {
        return Optional.ofNullable(this.status);
    }

    /**
     * Tags of the current resource.
     * 
     */
    @Import(name="tags")
    private @Nullable Output<Map<String,Object>> tags;

    /**
     * @return Tags of the current resource.
     * 
     */
    public Optional<Output<Map<String,Object>>> tags() {
        return Optional.ofNullable(this.tags);
    }

    private DhcpOptionsSetState() {}

    private DhcpOptionsSetState(DhcpOptionsSetState $) {
        this.associateVpcs = $.associateVpcs;
        this.dhcpOptionsSetDescription = $.dhcpOptionsSetDescription;
        this.dhcpOptionsSetName = $.dhcpOptionsSetName;
        this.domainName = $.domainName;
        this.domainNameServers = $.domainNameServers;
        this.dryRun = $.dryRun;
        this.ipv6LeaseTime = $.ipv6LeaseTime;
        this.leaseTime = $.leaseTime;
        this.ownerId = $.ownerId;
        this.resourceGroupId = $.resourceGroupId;
        this.status = $.status;
        this.tags = $.tags;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(DhcpOptionsSetState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private DhcpOptionsSetState $;

        public Builder() {
            $ = new DhcpOptionsSetState();
        }

        public Builder(DhcpOptionsSetState defaults) {
            $ = new DhcpOptionsSetState(Objects.requireNonNull(defaults));
        }

        /**
         * @param associateVpcs Field &#39;associate_vpcs&#39; has been deprecated from provider version 1.153.0 and it will be removed in the future version. Please use the new resource &#39;alicloud_vpc_dhcp_options_set_attachment&#39; to attach DhcpOptionsSet and Vpc. See `associate_vpcs` below.
         * 
         * @return builder
         * 
         * @deprecated
         * Field &#39;associate_vpcs&#39; has been deprecated from provider version 1.211.0. Field &#39;associate_vpcs&#39; has been deprecated from provider version 1.153.0 and it will be removed in the future version. Please use the new resource &#39;alicloud_vpc_dhcp_options_set_attachment&#39; to attach DhcpOptionsSet and Vpc.
         * 
         */
        @Deprecated /* Field 'associate_vpcs' has been deprecated from provider version 1.211.0. Field 'associate_vpcs' has been deprecated from provider version 1.153.0 and it will be removed in the future version. Please use the new resource 'alicloud_vpc_dhcp_options_set_attachment' to attach DhcpOptionsSet and Vpc. */
        public Builder associateVpcs(@Nullable Output<List<DhcpOptionsSetAssociateVpcArgs>> associateVpcs) {
            $.associateVpcs = associateVpcs;
            return this;
        }

        /**
         * @param associateVpcs Field &#39;associate_vpcs&#39; has been deprecated from provider version 1.153.0 and it will be removed in the future version. Please use the new resource &#39;alicloud_vpc_dhcp_options_set_attachment&#39; to attach DhcpOptionsSet and Vpc. See `associate_vpcs` below.
         * 
         * @return builder
         * 
         * @deprecated
         * Field &#39;associate_vpcs&#39; has been deprecated from provider version 1.211.0. Field &#39;associate_vpcs&#39; has been deprecated from provider version 1.153.0 and it will be removed in the future version. Please use the new resource &#39;alicloud_vpc_dhcp_options_set_attachment&#39; to attach DhcpOptionsSet and Vpc.
         * 
         */
        @Deprecated /* Field 'associate_vpcs' has been deprecated from provider version 1.211.0. Field 'associate_vpcs' has been deprecated from provider version 1.153.0 and it will be removed in the future version. Please use the new resource 'alicloud_vpc_dhcp_options_set_attachment' to attach DhcpOptionsSet and Vpc. */
        public Builder associateVpcs(List<DhcpOptionsSetAssociateVpcArgs> associateVpcs) {
            return associateVpcs(Output.of(associateVpcs));
        }

        /**
         * @param associateVpcs Field &#39;associate_vpcs&#39; has been deprecated from provider version 1.153.0 and it will be removed in the future version. Please use the new resource &#39;alicloud_vpc_dhcp_options_set_attachment&#39; to attach DhcpOptionsSet and Vpc. See `associate_vpcs` below.
         * 
         * @return builder
         * 
         * @deprecated
         * Field &#39;associate_vpcs&#39; has been deprecated from provider version 1.211.0. Field &#39;associate_vpcs&#39; has been deprecated from provider version 1.153.0 and it will be removed in the future version. Please use the new resource &#39;alicloud_vpc_dhcp_options_set_attachment&#39; to attach DhcpOptionsSet and Vpc.
         * 
         */
        @Deprecated /* Field 'associate_vpcs' has been deprecated from provider version 1.211.0. Field 'associate_vpcs' has been deprecated from provider version 1.153.0 and it will be removed in the future version. Please use the new resource 'alicloud_vpc_dhcp_options_set_attachment' to attach DhcpOptionsSet and Vpc. */
        public Builder associateVpcs(DhcpOptionsSetAssociateVpcArgs... associateVpcs) {
            return associateVpcs(List.of(associateVpcs));
        }

        /**
         * @param dhcpOptionsSetDescription The description can be blank or contain 1 to 256 characters. It must start with a letter or Chinese character but cannot start with http:// or https://.
         * 
         * @return builder
         * 
         */
        public Builder dhcpOptionsSetDescription(@Nullable Output<String> dhcpOptionsSetDescription) {
            $.dhcpOptionsSetDescription = dhcpOptionsSetDescription;
            return this;
        }

        /**
         * @param dhcpOptionsSetDescription The description can be blank or contain 1 to 256 characters. It must start with a letter or Chinese character but cannot start with http:// or https://.
         * 
         * @return builder
         * 
         */
        public Builder dhcpOptionsSetDescription(String dhcpOptionsSetDescription) {
            return dhcpOptionsSetDescription(Output.of(dhcpOptionsSetDescription));
        }

        /**
         * @param dhcpOptionsSetName The name must be 2 to 128 characters in length and can contain letters, Chinese characters, digits, underscores (_), and hyphens (-). It must start with a letter or a Chinese character.
         * 
         * @return builder
         * 
         */
        public Builder dhcpOptionsSetName(@Nullable Output<String> dhcpOptionsSetName) {
            $.dhcpOptionsSetName = dhcpOptionsSetName;
            return this;
        }

        /**
         * @param dhcpOptionsSetName The name must be 2 to 128 characters in length and can contain letters, Chinese characters, digits, underscores (_), and hyphens (-). It must start with a letter or a Chinese character.
         * 
         * @return builder
         * 
         */
        public Builder dhcpOptionsSetName(String dhcpOptionsSetName) {
            return dhcpOptionsSetName(Output.of(dhcpOptionsSetName));
        }

        /**
         * @param domainName The root domain, for example, example.com. After a DHCP options set is associated with a Virtual Private Cloud (VPC) network, the root domain in the DHCP options set is automatically synchronized to the ECS instances in the VPC network.
         * 
         * @return builder
         * 
         */
        public Builder domainName(@Nullable Output<String> domainName) {
            $.domainName = domainName;
            return this;
        }

        /**
         * @param domainName The root domain, for example, example.com. After a DHCP options set is associated with a Virtual Private Cloud (VPC) network, the root domain in the DHCP options set is automatically synchronized to the ECS instances in the VPC network.
         * 
         * @return builder
         * 
         */
        public Builder domainName(String domainName) {
            return domainName(Output.of(domainName));
        }

        /**
         * @param domainNameServers The DNS server IP addresses. Up to four DNS server IP addresses can be specified. IP addresses must be separated with commas (,).Before you specify any DNS server IP address, all ECS instances in the associated VPC network use the IP addresses of the Alibaba Cloud DNS servers, which are 100.100.2.136 and 100.100.2.138.
         * 
         * @return builder
         * 
         */
        public Builder domainNameServers(@Nullable Output<String> domainNameServers) {
            $.domainNameServers = domainNameServers;
            return this;
        }

        /**
         * @param domainNameServers The DNS server IP addresses. Up to four DNS server IP addresses can be specified. IP addresses must be separated with commas (,).Before you specify any DNS server IP address, all ECS instances in the associated VPC network use the IP addresses of the Alibaba Cloud DNS servers, which are 100.100.2.136 and 100.100.2.138.
         * 
         * @return builder
         * 
         */
        public Builder domainNameServers(String domainNameServers) {
            return domainNameServers(Output.of(domainNameServers));
        }

        /**
         * @param dryRun Whether to PreCheck only this request, value:
         * - **true**: sends a check request and does not delete the DHCP option set. Check items include whether required parameters are filled in, request format, and restrictions. If the check fails, the corresponding error is returned. If the check passes, the error code &#39;DryRunOperation&#39; is returned &#39;.
         * - **false** (default): Sends a normal request and directly deletes the DHCP option set after checking.
         * 
         * @return builder
         * 
         */
        public Builder dryRun(@Nullable Output<Boolean> dryRun) {
            $.dryRun = dryRun;
            return this;
        }

        /**
         * @param dryRun Whether to PreCheck only this request, value:
         * - **true**: sends a check request and does not delete the DHCP option set. Check items include whether required parameters are filled in, request format, and restrictions. If the check fails, the corresponding error is returned. If the check passes, the error code &#39;DryRunOperation&#39; is returned &#39;.
         * - **false** (default): Sends a normal request and directly deletes the DHCP option set after checking.
         * 
         * @return builder
         * 
         */
        public Builder dryRun(Boolean dryRun) {
            return dryRun(Output.of(dryRun));
        }

        /**
         * @param ipv6LeaseTime The lease time of the IPv6 DHCP option set.When the lease time is set to hours: Unit: h. Value range: 24h ~ 1176h,87600h ~ 175200h. Default value: 87600h.When the lease time is set to day: Unit: d. Value range: 1d ~ 49d,3650d ~ 7300d. Default value: 3650d.
         * 
         * @return builder
         * 
         */
        public Builder ipv6LeaseTime(@Nullable Output<String> ipv6LeaseTime) {
            $.ipv6LeaseTime = ipv6LeaseTime;
            return this;
        }

        /**
         * @param ipv6LeaseTime The lease time of the IPv6 DHCP option set.When the lease time is set to hours: Unit: h. Value range: 24h ~ 1176h,87600h ~ 175200h. Default value: 87600h.When the lease time is set to day: Unit: d. Value range: 1d ~ 49d,3650d ~ 7300d. Default value: 3650d.
         * 
         * @return builder
         * 
         */
        public Builder ipv6LeaseTime(String ipv6LeaseTime) {
            return ipv6LeaseTime(Output.of(ipv6LeaseTime));
        }

        /**
         * @param leaseTime The lease time of the IPv4 DHCP option set.When the lease time is set to hours: Unit: h. Value range: 24h ~ 1176h,87600h ~ 175200h. Default value: 87600h.When the lease time is set to day: Unit: d. Value range: 1d ~ 49d,3650d ~ 7300d. Default value: 3650d.
         * 
         * @return builder
         * 
         */
        public Builder leaseTime(@Nullable Output<String> leaseTime) {
            $.leaseTime = leaseTime;
            return this;
        }

        /**
         * @param leaseTime The lease time of the IPv4 DHCP option set.When the lease time is set to hours: Unit: h. Value range: 24h ~ 1176h,87600h ~ 175200h. Default value: 87600h.When the lease time is set to day: Unit: d. Value range: 1d ~ 49d,3650d ~ 7300d. Default value: 3650d.
         * 
         * @return builder
         * 
         */
        public Builder leaseTime(String leaseTime) {
            return leaseTime(Output.of(leaseTime));
        }

        /**
         * @param ownerId The ID of the account to which the DHCP options set belongs.
         * 
         * @return builder
         * 
         */
        public Builder ownerId(@Nullable Output<Integer> ownerId) {
            $.ownerId = ownerId;
            return this;
        }

        /**
         * @param ownerId The ID of the account to which the DHCP options set belongs.
         * 
         * @return builder
         * 
         */
        public Builder ownerId(Integer ownerId) {
            return ownerId(Output.of(ownerId));
        }

        /**
         * @param resourceGroupId The ID of the resource group.
         * 
         * @return builder
         * 
         */
        public Builder resourceGroupId(@Nullable Output<String> resourceGroupId) {
            $.resourceGroupId = resourceGroupId;
            return this;
        }

        /**
         * @param resourceGroupId The ID of the resource group.
         * 
         * @return builder
         * 
         */
        public Builder resourceGroupId(String resourceGroupId) {
            return resourceGroupId(Output.of(resourceGroupId));
        }

        /**
         * @param status The status of the resource.
         * 
         * @return builder
         * 
         */
        public Builder status(@Nullable Output<String> status) {
            $.status = status;
            return this;
        }

        /**
         * @param status The status of the resource.
         * 
         * @return builder
         * 
         */
        public Builder status(String status) {
            return status(Output.of(status));
        }

        /**
         * @param tags Tags of the current resource.
         * 
         * @return builder
         * 
         */
        public Builder tags(@Nullable Output<Map<String,Object>> tags) {
            $.tags = tags;
            return this;
        }

        /**
         * @param tags Tags of the current resource.
         * 
         * @return builder
         * 
         */
        public Builder tags(Map<String,Object> tags) {
            return tags(Output.of(tags));
        }

        public DhcpOptionsSetState build() {
            return $;
        }
    }

}
