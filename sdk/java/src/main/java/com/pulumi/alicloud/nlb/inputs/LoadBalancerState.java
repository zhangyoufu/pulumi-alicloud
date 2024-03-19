// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.nlb.inputs;

import com.pulumi.alicloud.nlb.inputs.LoadBalancerDeletionProtectionConfigArgs;
import com.pulumi.alicloud.nlb.inputs.LoadBalancerModificationProtectionConfigArgs;
import com.pulumi.alicloud.nlb.inputs.LoadBalancerZoneMappingArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.Object;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class LoadBalancerState extends com.pulumi.resources.ResourceArgs {

    public static final LoadBalancerState Empty = new LoadBalancerState();

    /**
     * Protocol version. Value:
     * - **Ipv4**:IPv4 type.
     * - **DualStack**: Double Stack type.
     * 
     */
    @Import(name="addressIpVersion")
    private @Nullable Output<String> addressIpVersion;

    /**
     * @return Protocol version. Value:
     * - **Ipv4**:IPv4 type.
     * - **DualStack**: Double Stack type.
     * 
     */
    public Optional<Output<String>> addressIpVersion() {
        return Optional.ofNullable(this.addressIpVersion);
    }

    /**
     * The network address type of IPv4 for network load balancing. Value:
     * - **Internet**: public network. Load balancer has a public network IP address, and the DNS domain name is resolved to a public network IP address, so it can be accessed in a public network environment.
     * - **Intranet**: private network. The server load balancer only has a private IP address, and the DNS domain name is resolved to the private IP address, so it can only be accessed by the intranet environment of the VPC where the server load balancer is located.
     * 
     */
    @Import(name="addressType")
    private @Nullable Output<String> addressType;

    /**
     * @return The network address type of IPv4 for network load balancing. Value:
     * - **Internet**: public network. Load balancer has a public network IP address, and the DNS domain name is resolved to a public network IP address, so it can be accessed in a public network environment.
     * - **Intranet**: private network. The server load balancer only has a private IP address, and the DNS domain name is resolved to the private IP address, so it can only be accessed by the intranet environment of the VPC where the server load balancer is located.
     * 
     */
    public Optional<Output<String>> addressType() {
        return Optional.ofNullable(this.addressType);
    }

    /**
     * The ID of the shared bandwidth package associated with the public network instance.
     * 
     */
    @Import(name="bandwidthPackageId")
    private @Nullable Output<String> bandwidthPackageId;

    /**
     * @return The ID of the shared bandwidth package associated with the public network instance.
     * 
     */
    public Optional<Output<String>> bandwidthPackageId() {
        return Optional.ofNullable(this.bandwidthPackageId);
    }

    /**
     * Resource creation time, using Greenwich Mean Time, formating&#39; yyyy-MM-ddTHH:mm:ssZ &#39;.
     * 
     */
    @Import(name="createTime")
    private @Nullable Output<String> createTime;

    /**
     * @return Resource creation time, using Greenwich Mean Time, formating&#39; yyyy-MM-ddTHH:mm:ssZ &#39;.
     * 
     */
    public Optional<Output<String>> createTime() {
        return Optional.ofNullable(this.createTime);
    }

    /**
     * Whether cross-zone is enabled for a network-based load balancing instance. Value:
     * - **true**: on.
     * - **false**: closed.
     * 
     */
    @Import(name="crossZoneEnabled")
    private @Nullable Output<Boolean> crossZoneEnabled;

    /**
     * @return Whether cross-zone is enabled for a network-based load balancing instance. Value:
     * - **true**: on.
     * - **false**: closed.
     * 
     */
    public Optional<Output<Boolean>> crossZoneEnabled() {
        return Optional.ofNullable(this.crossZoneEnabled);
    }

    /**
     * Delete protection. See `deletion_protection_config` below.
     * 
     */
    @Import(name="deletionProtectionConfig")
    private @Nullable Output<LoadBalancerDeletionProtectionConfigArgs> deletionProtectionConfig;

    /**
     * @return Delete protection. See `deletion_protection_config` below.
     * 
     */
    public Optional<Output<LoadBalancerDeletionProtectionConfigArgs>> deletionProtectionConfig() {
        return Optional.ofNullable(this.deletionProtectionConfig);
    }

    /**
     * Specifies whether to enable deletion protection. Default value: `false`. Valid values:
     * 
     */
    @Import(name="deletionProtectionEnabled")
    private @Nullable Output<Boolean> deletionProtectionEnabled;

    /**
     * @return Specifies whether to enable deletion protection. Default value: `false`. Valid values:
     * 
     */
    public Optional<Output<Boolean>> deletionProtectionEnabled() {
        return Optional.ofNullable(this.deletionProtectionEnabled);
    }

    /**
     * The reason why the deletion protection feature is enabled or disabled. The `deletion_protection_reason` takes effect only when `deletion_protection_enabled` is set to `true`.
     * 
     */
    @Import(name="deletionProtectionReason")
    private @Nullable Output<String> deletionProtectionReason;

    /**
     * @return The reason why the deletion protection feature is enabled or disabled. The `deletion_protection_reason` takes effect only when `deletion_protection_enabled` is set to `true`.
     * 
     */
    public Optional<Output<String>> deletionProtectionReason() {
        return Optional.ofNullable(this.deletionProtectionReason);
    }

    /**
     * The domain name of the NLB instance.
     * 
     */
    @Import(name="dnsName")
    private @Nullable Output<String> dnsName;

    /**
     * @return The domain name of the NLB instance.
     * 
     */
    public Optional<Output<String>> dnsName() {
        return Optional.ofNullable(this.dnsName);
    }

    /**
     * The IPv6 address type of network load balancing. Value:
     * - **Internet**: Server Load Balancer has a public IP address, and the DNS domain name is resolved to a public IP address, so it can be accessed in a public network environment.
     * - **Intranet**: SLB only has the private IP address, and the DNS domain name is resolved to the private IP address, so it can only be accessed by the Intranet environment of the VPC where SLB is located.
     * 
     */
    @Import(name="ipv6AddressType")
    private @Nullable Output<String> ipv6AddressType;

    /**
     * @return The IPv6 address type of network load balancing. Value:
     * - **Internet**: Server Load Balancer has a public IP address, and the DNS domain name is resolved to a public IP address, so it can be accessed in a public network environment.
     * - **Intranet**: SLB only has the private IP address, and the DNS domain name is resolved to the private IP address, so it can only be accessed by the Intranet environment of the VPC where SLB is located.
     * 
     */
    public Optional<Output<String>> ipv6AddressType() {
        return Optional.ofNullable(this.ipv6AddressType);
    }

    /**
     * The business status of the NLB instance.
     * 
     */
    @Import(name="loadBalancerBusinessStatus")
    private @Nullable Output<String> loadBalancerBusinessStatus;

    /**
     * @return The business status of the NLB instance.
     * 
     */
    public Optional<Output<String>> loadBalancerBusinessStatus() {
        return Optional.ofNullable(this.loadBalancerBusinessStatus);
    }

    /**
     * The name of the network-based load balancing instance.  2 to 128 English or Chinese characters in length, which must start with a letter or Chinese, and can contain numbers, half-width periods (.), underscores (_), and dashes (-).
     * 
     */
    @Import(name="loadBalancerName")
    private @Nullable Output<String> loadBalancerName;

    /**
     * @return The name of the network-based load balancing instance.  2 to 128 English or Chinese characters in length, which must start with a letter or Chinese, and can contain numbers, half-width periods (.), underscores (_), and dashes (-).
     * 
     */
    public Optional<Output<String>> loadBalancerName() {
        return Optional.ofNullable(this.loadBalancerName);
    }

    /**
     * Load balancing type. Only value: **network**, which indicates network-based load balancing.
     * 
     */
    @Import(name="loadBalancerType")
    private @Nullable Output<String> loadBalancerType;

    /**
     * @return Load balancing type. Only value: **network**, which indicates network-based load balancing.
     * 
     */
    public Optional<Output<String>> loadBalancerType() {
        return Optional.ofNullable(this.loadBalancerType);
    }

    /**
     * Modify protection. See `modification_protection_config` below.
     * 
     */
    @Import(name="modificationProtectionConfig")
    private @Nullable Output<LoadBalancerModificationProtectionConfigArgs> modificationProtectionConfig;

    /**
     * @return Modify protection. See `modification_protection_config` below.
     * 
     */
    public Optional<Output<LoadBalancerModificationProtectionConfigArgs>> modificationProtectionConfig() {
        return Optional.ofNullable(this.modificationProtectionConfig);
    }

    /**
     * The reason why the configuration read-only mode is enabled. The `modification_protection_reason` takes effect only when `modification_protection_status` is set to `ConsoleProtection`.
     * 
     */
    @Import(name="modificationProtectionReason")
    private @Nullable Output<String> modificationProtectionReason;

    /**
     * @return The reason why the configuration read-only mode is enabled. The `modification_protection_reason` takes effect only when `modification_protection_status` is set to `ConsoleProtection`.
     * 
     */
    public Optional<Output<String>> modificationProtectionReason() {
        return Optional.ofNullable(this.modificationProtectionReason);
    }

    /**
     * Specifies whether to enable the configuration read-only mode. Default value: `NonProtection`. Valid values:
     * 
     */
    @Import(name="modificationProtectionStatus")
    private @Nullable Output<String> modificationProtectionStatus;

    /**
     * @return Specifies whether to enable the configuration read-only mode. Default value: `NonProtection`. Valid values:
     * 
     */
    public Optional<Output<String>> modificationProtectionStatus() {
        return Optional.ofNullable(this.modificationProtectionStatus);
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
     * The security group to which the network-based SLB instance belongs.
     * 
     */
    @Import(name="securityGroupIds")
    private @Nullable Output<List<String>> securityGroupIds;

    /**
     * @return The security group to which the network-based SLB instance belongs.
     * 
     */
    public Optional<Output<List<String>>> securityGroupIds() {
        return Optional.ofNullable(this.securityGroupIds);
    }

    /**
     * ON.
     * 
     */
    @Import(name="status")
    private @Nullable Output<String> status;

    /**
     * @return ON.
     * 
     */
    public Optional<Output<String>> status() {
        return Optional.ofNullable(this.status);
    }

    /**
     * List of labels.
     * 
     */
    @Import(name="tags")
    private @Nullable Output<Map<String,Object>> tags;

    /**
     * @return List of labels.
     * 
     */
    public Optional<Output<Map<String,Object>>> tags() {
        return Optional.ofNullable(this.tags);
    }

    /**
     * The ID of the network-based SLB instance.
     * 
     */
    @Import(name="vpcId")
    private @Nullable Output<String> vpcId;

    /**
     * @return The ID of the network-based SLB instance.
     * 
     */
    public Optional<Output<String>> vpcId() {
        return Optional.ofNullable(this.vpcId);
    }

    /**
     * The list of zones and vSwitch mappings. You must add at least two zones and a maximum of 10 zones. See `zone_mappings` below.
     * 
     */
    @Import(name="zoneMappings")
    private @Nullable Output<List<LoadBalancerZoneMappingArgs>> zoneMappings;

    /**
     * @return The list of zones and vSwitch mappings. You must add at least two zones and a maximum of 10 zones. See `zone_mappings` below.
     * 
     */
    public Optional<Output<List<LoadBalancerZoneMappingArgs>>> zoneMappings() {
        return Optional.ofNullable(this.zoneMappings);
    }

    private LoadBalancerState() {}

    private LoadBalancerState(LoadBalancerState $) {
        this.addressIpVersion = $.addressIpVersion;
        this.addressType = $.addressType;
        this.bandwidthPackageId = $.bandwidthPackageId;
        this.createTime = $.createTime;
        this.crossZoneEnabled = $.crossZoneEnabled;
        this.deletionProtectionConfig = $.deletionProtectionConfig;
        this.deletionProtectionEnabled = $.deletionProtectionEnabled;
        this.deletionProtectionReason = $.deletionProtectionReason;
        this.dnsName = $.dnsName;
        this.ipv6AddressType = $.ipv6AddressType;
        this.loadBalancerBusinessStatus = $.loadBalancerBusinessStatus;
        this.loadBalancerName = $.loadBalancerName;
        this.loadBalancerType = $.loadBalancerType;
        this.modificationProtectionConfig = $.modificationProtectionConfig;
        this.modificationProtectionReason = $.modificationProtectionReason;
        this.modificationProtectionStatus = $.modificationProtectionStatus;
        this.resourceGroupId = $.resourceGroupId;
        this.securityGroupIds = $.securityGroupIds;
        this.status = $.status;
        this.tags = $.tags;
        this.vpcId = $.vpcId;
        this.zoneMappings = $.zoneMappings;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(LoadBalancerState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private LoadBalancerState $;

        public Builder() {
            $ = new LoadBalancerState();
        }

        public Builder(LoadBalancerState defaults) {
            $ = new LoadBalancerState(Objects.requireNonNull(defaults));
        }

        /**
         * @param addressIpVersion Protocol version. Value:
         * - **Ipv4**:IPv4 type.
         * - **DualStack**: Double Stack type.
         * 
         * @return builder
         * 
         */
        public Builder addressIpVersion(@Nullable Output<String> addressIpVersion) {
            $.addressIpVersion = addressIpVersion;
            return this;
        }

        /**
         * @param addressIpVersion Protocol version. Value:
         * - **Ipv4**:IPv4 type.
         * - **DualStack**: Double Stack type.
         * 
         * @return builder
         * 
         */
        public Builder addressIpVersion(String addressIpVersion) {
            return addressIpVersion(Output.of(addressIpVersion));
        }

        /**
         * @param addressType The network address type of IPv4 for network load balancing. Value:
         * - **Internet**: public network. Load balancer has a public network IP address, and the DNS domain name is resolved to a public network IP address, so it can be accessed in a public network environment.
         * - **Intranet**: private network. The server load balancer only has a private IP address, and the DNS domain name is resolved to the private IP address, so it can only be accessed by the intranet environment of the VPC where the server load balancer is located.
         * 
         * @return builder
         * 
         */
        public Builder addressType(@Nullable Output<String> addressType) {
            $.addressType = addressType;
            return this;
        }

        /**
         * @param addressType The network address type of IPv4 for network load balancing. Value:
         * - **Internet**: public network. Load balancer has a public network IP address, and the DNS domain name is resolved to a public network IP address, so it can be accessed in a public network environment.
         * - **Intranet**: private network. The server load balancer only has a private IP address, and the DNS domain name is resolved to the private IP address, so it can only be accessed by the intranet environment of the VPC where the server load balancer is located.
         * 
         * @return builder
         * 
         */
        public Builder addressType(String addressType) {
            return addressType(Output.of(addressType));
        }

        /**
         * @param bandwidthPackageId The ID of the shared bandwidth package associated with the public network instance.
         * 
         * @return builder
         * 
         */
        public Builder bandwidthPackageId(@Nullable Output<String> bandwidthPackageId) {
            $.bandwidthPackageId = bandwidthPackageId;
            return this;
        }

        /**
         * @param bandwidthPackageId The ID of the shared bandwidth package associated with the public network instance.
         * 
         * @return builder
         * 
         */
        public Builder bandwidthPackageId(String bandwidthPackageId) {
            return bandwidthPackageId(Output.of(bandwidthPackageId));
        }

        /**
         * @param createTime Resource creation time, using Greenwich Mean Time, formating&#39; yyyy-MM-ddTHH:mm:ssZ &#39;.
         * 
         * @return builder
         * 
         */
        public Builder createTime(@Nullable Output<String> createTime) {
            $.createTime = createTime;
            return this;
        }

        /**
         * @param createTime Resource creation time, using Greenwich Mean Time, formating&#39; yyyy-MM-ddTHH:mm:ssZ &#39;.
         * 
         * @return builder
         * 
         */
        public Builder createTime(String createTime) {
            return createTime(Output.of(createTime));
        }

        /**
         * @param crossZoneEnabled Whether cross-zone is enabled for a network-based load balancing instance. Value:
         * - **true**: on.
         * - **false**: closed.
         * 
         * @return builder
         * 
         */
        public Builder crossZoneEnabled(@Nullable Output<Boolean> crossZoneEnabled) {
            $.crossZoneEnabled = crossZoneEnabled;
            return this;
        }

        /**
         * @param crossZoneEnabled Whether cross-zone is enabled for a network-based load balancing instance. Value:
         * - **true**: on.
         * - **false**: closed.
         * 
         * @return builder
         * 
         */
        public Builder crossZoneEnabled(Boolean crossZoneEnabled) {
            return crossZoneEnabled(Output.of(crossZoneEnabled));
        }

        /**
         * @param deletionProtectionConfig Delete protection. See `deletion_protection_config` below.
         * 
         * @return builder
         * 
         */
        public Builder deletionProtectionConfig(@Nullable Output<LoadBalancerDeletionProtectionConfigArgs> deletionProtectionConfig) {
            $.deletionProtectionConfig = deletionProtectionConfig;
            return this;
        }

        /**
         * @param deletionProtectionConfig Delete protection. See `deletion_protection_config` below.
         * 
         * @return builder
         * 
         */
        public Builder deletionProtectionConfig(LoadBalancerDeletionProtectionConfigArgs deletionProtectionConfig) {
            return deletionProtectionConfig(Output.of(deletionProtectionConfig));
        }

        /**
         * @param deletionProtectionEnabled Specifies whether to enable deletion protection. Default value: `false`. Valid values:
         * 
         * @return builder
         * 
         */
        public Builder deletionProtectionEnabled(@Nullable Output<Boolean> deletionProtectionEnabled) {
            $.deletionProtectionEnabled = deletionProtectionEnabled;
            return this;
        }

        /**
         * @param deletionProtectionEnabled Specifies whether to enable deletion protection. Default value: `false`. Valid values:
         * 
         * @return builder
         * 
         */
        public Builder deletionProtectionEnabled(Boolean deletionProtectionEnabled) {
            return deletionProtectionEnabled(Output.of(deletionProtectionEnabled));
        }

        /**
         * @param deletionProtectionReason The reason why the deletion protection feature is enabled or disabled. The `deletion_protection_reason` takes effect only when `deletion_protection_enabled` is set to `true`.
         * 
         * @return builder
         * 
         */
        public Builder deletionProtectionReason(@Nullable Output<String> deletionProtectionReason) {
            $.deletionProtectionReason = deletionProtectionReason;
            return this;
        }

        /**
         * @param deletionProtectionReason The reason why the deletion protection feature is enabled or disabled. The `deletion_protection_reason` takes effect only when `deletion_protection_enabled` is set to `true`.
         * 
         * @return builder
         * 
         */
        public Builder deletionProtectionReason(String deletionProtectionReason) {
            return deletionProtectionReason(Output.of(deletionProtectionReason));
        }

        /**
         * @param dnsName The domain name of the NLB instance.
         * 
         * @return builder
         * 
         */
        public Builder dnsName(@Nullable Output<String> dnsName) {
            $.dnsName = dnsName;
            return this;
        }

        /**
         * @param dnsName The domain name of the NLB instance.
         * 
         * @return builder
         * 
         */
        public Builder dnsName(String dnsName) {
            return dnsName(Output.of(dnsName));
        }

        /**
         * @param ipv6AddressType The IPv6 address type of network load balancing. Value:
         * - **Internet**: Server Load Balancer has a public IP address, and the DNS domain name is resolved to a public IP address, so it can be accessed in a public network environment.
         * - **Intranet**: SLB only has the private IP address, and the DNS domain name is resolved to the private IP address, so it can only be accessed by the Intranet environment of the VPC where SLB is located.
         * 
         * @return builder
         * 
         */
        public Builder ipv6AddressType(@Nullable Output<String> ipv6AddressType) {
            $.ipv6AddressType = ipv6AddressType;
            return this;
        }

        /**
         * @param ipv6AddressType The IPv6 address type of network load balancing. Value:
         * - **Internet**: Server Load Balancer has a public IP address, and the DNS domain name is resolved to a public IP address, so it can be accessed in a public network environment.
         * - **Intranet**: SLB only has the private IP address, and the DNS domain name is resolved to the private IP address, so it can only be accessed by the Intranet environment of the VPC where SLB is located.
         * 
         * @return builder
         * 
         */
        public Builder ipv6AddressType(String ipv6AddressType) {
            return ipv6AddressType(Output.of(ipv6AddressType));
        }

        /**
         * @param loadBalancerBusinessStatus The business status of the NLB instance.
         * 
         * @return builder
         * 
         */
        public Builder loadBalancerBusinessStatus(@Nullable Output<String> loadBalancerBusinessStatus) {
            $.loadBalancerBusinessStatus = loadBalancerBusinessStatus;
            return this;
        }

        /**
         * @param loadBalancerBusinessStatus The business status of the NLB instance.
         * 
         * @return builder
         * 
         */
        public Builder loadBalancerBusinessStatus(String loadBalancerBusinessStatus) {
            return loadBalancerBusinessStatus(Output.of(loadBalancerBusinessStatus));
        }

        /**
         * @param loadBalancerName The name of the network-based load balancing instance.  2 to 128 English or Chinese characters in length, which must start with a letter or Chinese, and can contain numbers, half-width periods (.), underscores (_), and dashes (-).
         * 
         * @return builder
         * 
         */
        public Builder loadBalancerName(@Nullable Output<String> loadBalancerName) {
            $.loadBalancerName = loadBalancerName;
            return this;
        }

        /**
         * @param loadBalancerName The name of the network-based load balancing instance.  2 to 128 English or Chinese characters in length, which must start with a letter or Chinese, and can contain numbers, half-width periods (.), underscores (_), and dashes (-).
         * 
         * @return builder
         * 
         */
        public Builder loadBalancerName(String loadBalancerName) {
            return loadBalancerName(Output.of(loadBalancerName));
        }

        /**
         * @param loadBalancerType Load balancing type. Only value: **network**, which indicates network-based load balancing.
         * 
         * @return builder
         * 
         */
        public Builder loadBalancerType(@Nullable Output<String> loadBalancerType) {
            $.loadBalancerType = loadBalancerType;
            return this;
        }

        /**
         * @param loadBalancerType Load balancing type. Only value: **network**, which indicates network-based load balancing.
         * 
         * @return builder
         * 
         */
        public Builder loadBalancerType(String loadBalancerType) {
            return loadBalancerType(Output.of(loadBalancerType));
        }

        /**
         * @param modificationProtectionConfig Modify protection. See `modification_protection_config` below.
         * 
         * @return builder
         * 
         */
        public Builder modificationProtectionConfig(@Nullable Output<LoadBalancerModificationProtectionConfigArgs> modificationProtectionConfig) {
            $.modificationProtectionConfig = modificationProtectionConfig;
            return this;
        }

        /**
         * @param modificationProtectionConfig Modify protection. See `modification_protection_config` below.
         * 
         * @return builder
         * 
         */
        public Builder modificationProtectionConfig(LoadBalancerModificationProtectionConfigArgs modificationProtectionConfig) {
            return modificationProtectionConfig(Output.of(modificationProtectionConfig));
        }

        /**
         * @param modificationProtectionReason The reason why the configuration read-only mode is enabled. The `modification_protection_reason` takes effect only when `modification_protection_status` is set to `ConsoleProtection`.
         * 
         * @return builder
         * 
         */
        public Builder modificationProtectionReason(@Nullable Output<String> modificationProtectionReason) {
            $.modificationProtectionReason = modificationProtectionReason;
            return this;
        }

        /**
         * @param modificationProtectionReason The reason why the configuration read-only mode is enabled. The `modification_protection_reason` takes effect only when `modification_protection_status` is set to `ConsoleProtection`.
         * 
         * @return builder
         * 
         */
        public Builder modificationProtectionReason(String modificationProtectionReason) {
            return modificationProtectionReason(Output.of(modificationProtectionReason));
        }

        /**
         * @param modificationProtectionStatus Specifies whether to enable the configuration read-only mode. Default value: `NonProtection`. Valid values:
         * 
         * @return builder
         * 
         */
        public Builder modificationProtectionStatus(@Nullable Output<String> modificationProtectionStatus) {
            $.modificationProtectionStatus = modificationProtectionStatus;
            return this;
        }

        /**
         * @param modificationProtectionStatus Specifies whether to enable the configuration read-only mode. Default value: `NonProtection`. Valid values:
         * 
         * @return builder
         * 
         */
        public Builder modificationProtectionStatus(String modificationProtectionStatus) {
            return modificationProtectionStatus(Output.of(modificationProtectionStatus));
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
         * @param securityGroupIds The security group to which the network-based SLB instance belongs.
         * 
         * @return builder
         * 
         */
        public Builder securityGroupIds(@Nullable Output<List<String>> securityGroupIds) {
            $.securityGroupIds = securityGroupIds;
            return this;
        }

        /**
         * @param securityGroupIds The security group to which the network-based SLB instance belongs.
         * 
         * @return builder
         * 
         */
        public Builder securityGroupIds(List<String> securityGroupIds) {
            return securityGroupIds(Output.of(securityGroupIds));
        }

        /**
         * @param securityGroupIds The security group to which the network-based SLB instance belongs.
         * 
         * @return builder
         * 
         */
        public Builder securityGroupIds(String... securityGroupIds) {
            return securityGroupIds(List.of(securityGroupIds));
        }

        /**
         * @param status ON.
         * 
         * @return builder
         * 
         */
        public Builder status(@Nullable Output<String> status) {
            $.status = status;
            return this;
        }

        /**
         * @param status ON.
         * 
         * @return builder
         * 
         */
        public Builder status(String status) {
            return status(Output.of(status));
        }

        /**
         * @param tags List of labels.
         * 
         * @return builder
         * 
         */
        public Builder tags(@Nullable Output<Map<String,Object>> tags) {
            $.tags = tags;
            return this;
        }

        /**
         * @param tags List of labels.
         * 
         * @return builder
         * 
         */
        public Builder tags(Map<String,Object> tags) {
            return tags(Output.of(tags));
        }

        /**
         * @param vpcId The ID of the network-based SLB instance.
         * 
         * @return builder
         * 
         */
        public Builder vpcId(@Nullable Output<String> vpcId) {
            $.vpcId = vpcId;
            return this;
        }

        /**
         * @param vpcId The ID of the network-based SLB instance.
         * 
         * @return builder
         * 
         */
        public Builder vpcId(String vpcId) {
            return vpcId(Output.of(vpcId));
        }

        /**
         * @param zoneMappings The list of zones and vSwitch mappings. You must add at least two zones and a maximum of 10 zones. See `zone_mappings` below.
         * 
         * @return builder
         * 
         */
        public Builder zoneMappings(@Nullable Output<List<LoadBalancerZoneMappingArgs>> zoneMappings) {
            $.zoneMappings = zoneMappings;
            return this;
        }

        /**
         * @param zoneMappings The list of zones and vSwitch mappings. You must add at least two zones and a maximum of 10 zones. See `zone_mappings` below.
         * 
         * @return builder
         * 
         */
        public Builder zoneMappings(List<LoadBalancerZoneMappingArgs> zoneMappings) {
            return zoneMappings(Output.of(zoneMappings));
        }

        /**
         * @param zoneMappings The list of zones and vSwitch mappings. You must add at least two zones and a maximum of 10 zones. See `zone_mappings` below.
         * 
         * @return builder
         * 
         */
        public Builder zoneMappings(LoadBalancerZoneMappingArgs... zoneMappings) {
            return zoneMappings(List.of(zoneMappings));
        }

        public LoadBalancerState build() {
            return $;
        }
    }

}
