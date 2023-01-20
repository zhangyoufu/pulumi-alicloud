// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.eds;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class AdConnectorOfficeSiteArgs extends com.pulumi.resources.ResourceArgs {

    public static final AdConnectorOfficeSiteArgs Empty = new AdConnectorOfficeSiteArgs();

    /**
     * The name of the workspace. The name must be 2 to 255 characters in length. It must start with a letter and cannot start with `http://` or `https://`. It can contain digits, colons (:), underscores (_), and hyphens (-).
     * 
     */
    @Import(name="adConnectorOfficeSiteName", required=true)
    private Output<String> adConnectorOfficeSiteName;

    /**
     * @return The name of the workspace. The name must be 2 to 255 characters in length. It must start with a letter and cannot start with `http://` or `https://`. It can contain digits, colons (:), underscores (_), and hyphens (-).
     * 
     */
    public Output<String> adConnectorOfficeSiteName() {
        return this.adConnectorOfficeSiteName;
    }

    /**
     * The ad hostname.
     * 
     */
    @Import(name="adHostname")
    private @Nullable Output<String> adHostname;

    /**
     * @return The ad hostname.
     * 
     */
    public Optional<Output<String>> adHostname() {
        return Optional.ofNullable(this.adHostname);
    }

    /**
     * The maximum public bandwidth value. Valid values: 0 to 200. If you do not specify this parameter or you set this parameter to 0, Internet access is disabled.
     * 
     */
    @Import(name="bandwidth")
    private @Nullable Output<Integer> bandwidth;

    /**
     * @return The maximum public bandwidth value. Valid values: 0 to 200. If you do not specify this parameter or you set this parameter to 0, Internet access is disabled.
     * 
     */
    public Optional<Output<Integer>> bandwidth() {
        return Optional.ofNullable(this.bandwidth);
    }

    /**
     * The ID of the CEN instance.
     * 
     */
    @Import(name="cenId", required=true)
    private Output<String> cenId;

    /**
     * @return The ID of the CEN instance.
     * 
     */
    public Output<String> cenId() {
        return this.cenId;
    }

    /**
     * The cen owner id.
     * 
     */
    @Import(name="cenOwnerId")
    private @Nullable Output<String> cenOwnerId;

    /**
     * @return The cen owner id.
     * 
     */
    public Optional<Output<String>> cenOwnerId() {
        return Optional.ofNullable(this.cenOwnerId);
    }

    /**
     * Workspace Corresponds to the Security Office Network of IPv4 Segment.
     * 
     */
    @Import(name="cidrBlock", required=true)
    private Output<String> cidrBlock;

    /**
     * @return Workspace Corresponds to the Security Office Network of IPv4 Segment.
     * 
     */
    public Output<String> cidrBlock() {
        return this.cidrBlock;
    }

    /**
     * The method that you use to connect to cloud desktops. **Note:** The VPC connection method is provided by Alibaba Cloud PrivateLink. You are not charged for PrivateLink. When you set this parameter to VPC or Any, PrivateLink is automatically activated. Default value: `INTERNET`. Valid values:
     * 
     */
    @Import(name="desktopAccessType")
    private @Nullable Output<String> desktopAccessType;

    /**
     * @return The method that you use to connect to cloud desktops. **Note:** The VPC connection method is provided by Alibaba Cloud PrivateLink. You are not charged for PrivateLink. When you set this parameter to VPC or Any, PrivateLink is automatically activated. Default value: `INTERNET`. Valid values:
     * 
     */
    public Optional<Output<String>> desktopAccessType() {
        return Optional.ofNullable(this.desktopAccessType);
    }

    /**
     * The IP address N of the DNS server of the enterprise AD system. You can specify only one IP address.
     * 
     */
    @Import(name="dnsAddresses", required=true)
    private Output<List<String>> dnsAddresses;

    /**
     * @return The IP address N of the DNS server of the enterprise AD system. You can specify only one IP address.
     * 
     */
    public Output<List<String>> dnsAddresses() {
        return this.dnsAddresses;
    }

    /**
     * The domain name of the enterprise AD system. You can register each domain name only once.
     * 
     */
    @Import(name="domainName", required=true)
    private Output<String> domainName;

    /**
     * @return The domain name of the enterprise AD system. You can register each domain name only once.
     * 
     */
    public Output<String> domainName() {
        return this.domainName;
    }

    /**
     * The password of the domain administrator. The password can be up to 64 characters in length.
     * 
     */
    @Import(name="domainPassword")
    private @Nullable Output<String> domainPassword;

    /**
     * @return The password of the domain administrator. The password can be up to 64 characters in length.
     * 
     */
    public Optional<Output<String>> domainPassword() {
        return Optional.ofNullable(this.domainPassword);
    }

    /**
     * The username of the domain administrator. The username can be up to 64 characters in length.
     * 
     */
    @Import(name="domainUserName")
    private @Nullable Output<String> domainUserName;

    /**
     * @return The username of the domain administrator. The username can be up to 64 characters in length.
     * 
     */
    public Optional<Output<String>> domainUserName() {
        return Optional.ofNullable(this.domainUserName);
    }

    /**
     * Specifies whether to grant the permissions of the local administrator to the desktop users. Default value: true.
     * 
     */
    @Import(name="enableAdminAccess")
    private @Nullable Output<Boolean> enableAdminAccess;

    /**
     * @return Specifies whether to grant the permissions of the local administrator to the desktop users. Default value: true.
     * 
     */
    public Optional<Output<Boolean>> enableAdminAccess() {
        return Optional.ofNullable(this.enableAdminAccess);
    }

    /**
     * Specifies whether to enable Internet access.
     * 
     */
    @Import(name="enableInternetAccess")
    private @Nullable Output<Boolean> enableInternetAccess;

    /**
     * @return Specifies whether to enable Internet access.
     * 
     */
    public Optional<Output<Boolean>> enableInternetAccess() {
        return Optional.ofNullable(this.enableInternetAccess);
    }

    /**
     * Specifies whether to enable multi-factor authentication (MFA).
     * 
     */
    @Import(name="mfaEnabled")
    private @Nullable Output<Boolean> mfaEnabled;

    /**
     * @return Specifies whether to enable multi-factor authentication (MFA).
     * 
     */
    public Optional<Output<Boolean>> mfaEnabled() {
        return Optional.ofNullable(this.mfaEnabled);
    }

    /**
     * The protocol type. Valid values: `ASP`, `HDX`.
     * 
     */
    @Import(name="protocolType")
    private @Nullable Output<String> protocolType;

    /**
     * @return The protocol type. Valid values: `ASP`, `HDX`.
     * 
     */
    public Optional<Output<String>> protocolType() {
        return Optional.ofNullable(this.protocolType);
    }

    /**
     * The AD Connector specifications. Valid values: `1`, `2`.
     * 
     */
    @Import(name="specification")
    private @Nullable Output<Integer> specification;

    /**
     * @return The AD Connector specifications. Valid values: `1`, `2`.
     * 
     */
    public Optional<Output<Integer>> specification() {
        return Optional.ofNullable(this.specification);
    }

    /**
     * The DNS address N of the enterprise AD subdomain. If you specify a value for the `sub_domain_name` parameter but you do not specify a value for this parameter, the DNS address of the subdomain is the same as the DNS address of the parent domain.
     * 
     */
    @Import(name="subDomainDnsAddresses")
    private @Nullable Output<List<String>> subDomainDnsAddresses;

    /**
     * @return The DNS address N of the enterprise AD subdomain. If you specify a value for the `sub_domain_name` parameter but you do not specify a value for this parameter, the DNS address of the subdomain is the same as the DNS address of the parent domain.
     * 
     */
    public Optional<Output<List<String>>> subDomainDnsAddresses() {
        return Optional.ofNullable(this.subDomainDnsAddresses);
    }

    /**
     * The domain name of the enterprise AD subdomain.
     * 
     */
    @Import(name="subDomainName")
    private @Nullable Output<String> subDomainName;

    /**
     * @return The domain name of the enterprise AD subdomain.
     * 
     */
    public Optional<Output<String>> subDomainName() {
        return Optional.ofNullable(this.subDomainName);
    }

    /**
     * The verification code. If the CEN instance that you specify for the CenId parameter belongs to another Alibaba Cloud account, you must call the SendVerifyCode operation to obtain the verification code.
     * 
     */
    @Import(name="verifyCode")
    private @Nullable Output<String> verifyCode;

    /**
     * @return The verification code. If the CEN instance that you specify for the CenId parameter belongs to another Alibaba Cloud account, you must call the SendVerifyCode operation to obtain the verification code.
     * 
     */
    public Optional<Output<String>> verifyCode() {
        return Optional.ofNullable(this.verifyCode);
    }

    private AdConnectorOfficeSiteArgs() {}

    private AdConnectorOfficeSiteArgs(AdConnectorOfficeSiteArgs $) {
        this.adConnectorOfficeSiteName = $.adConnectorOfficeSiteName;
        this.adHostname = $.adHostname;
        this.bandwidth = $.bandwidth;
        this.cenId = $.cenId;
        this.cenOwnerId = $.cenOwnerId;
        this.cidrBlock = $.cidrBlock;
        this.desktopAccessType = $.desktopAccessType;
        this.dnsAddresses = $.dnsAddresses;
        this.domainName = $.domainName;
        this.domainPassword = $.domainPassword;
        this.domainUserName = $.domainUserName;
        this.enableAdminAccess = $.enableAdminAccess;
        this.enableInternetAccess = $.enableInternetAccess;
        this.mfaEnabled = $.mfaEnabled;
        this.protocolType = $.protocolType;
        this.specification = $.specification;
        this.subDomainDnsAddresses = $.subDomainDnsAddresses;
        this.subDomainName = $.subDomainName;
        this.verifyCode = $.verifyCode;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(AdConnectorOfficeSiteArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private AdConnectorOfficeSiteArgs $;

        public Builder() {
            $ = new AdConnectorOfficeSiteArgs();
        }

        public Builder(AdConnectorOfficeSiteArgs defaults) {
            $ = new AdConnectorOfficeSiteArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param adConnectorOfficeSiteName The name of the workspace. The name must be 2 to 255 characters in length. It must start with a letter and cannot start with `http://` or `https://`. It can contain digits, colons (:), underscores (_), and hyphens (-).
         * 
         * @return builder
         * 
         */
        public Builder adConnectorOfficeSiteName(Output<String> adConnectorOfficeSiteName) {
            $.adConnectorOfficeSiteName = adConnectorOfficeSiteName;
            return this;
        }

        /**
         * @param adConnectorOfficeSiteName The name of the workspace. The name must be 2 to 255 characters in length. It must start with a letter and cannot start with `http://` or `https://`. It can contain digits, colons (:), underscores (_), and hyphens (-).
         * 
         * @return builder
         * 
         */
        public Builder adConnectorOfficeSiteName(String adConnectorOfficeSiteName) {
            return adConnectorOfficeSiteName(Output.of(adConnectorOfficeSiteName));
        }

        /**
         * @param adHostname The ad hostname.
         * 
         * @return builder
         * 
         */
        public Builder adHostname(@Nullable Output<String> adHostname) {
            $.adHostname = adHostname;
            return this;
        }

        /**
         * @param adHostname The ad hostname.
         * 
         * @return builder
         * 
         */
        public Builder adHostname(String adHostname) {
            return adHostname(Output.of(adHostname));
        }

        /**
         * @param bandwidth The maximum public bandwidth value. Valid values: 0 to 200. If you do not specify this parameter or you set this parameter to 0, Internet access is disabled.
         * 
         * @return builder
         * 
         */
        public Builder bandwidth(@Nullable Output<Integer> bandwidth) {
            $.bandwidth = bandwidth;
            return this;
        }

        /**
         * @param bandwidth The maximum public bandwidth value. Valid values: 0 to 200. If you do not specify this parameter or you set this parameter to 0, Internet access is disabled.
         * 
         * @return builder
         * 
         */
        public Builder bandwidth(Integer bandwidth) {
            return bandwidth(Output.of(bandwidth));
        }

        /**
         * @param cenId The ID of the CEN instance.
         * 
         * @return builder
         * 
         */
        public Builder cenId(Output<String> cenId) {
            $.cenId = cenId;
            return this;
        }

        /**
         * @param cenId The ID of the CEN instance.
         * 
         * @return builder
         * 
         */
        public Builder cenId(String cenId) {
            return cenId(Output.of(cenId));
        }

        /**
         * @param cenOwnerId The cen owner id.
         * 
         * @return builder
         * 
         */
        public Builder cenOwnerId(@Nullable Output<String> cenOwnerId) {
            $.cenOwnerId = cenOwnerId;
            return this;
        }

        /**
         * @param cenOwnerId The cen owner id.
         * 
         * @return builder
         * 
         */
        public Builder cenOwnerId(String cenOwnerId) {
            return cenOwnerId(Output.of(cenOwnerId));
        }

        /**
         * @param cidrBlock Workspace Corresponds to the Security Office Network of IPv4 Segment.
         * 
         * @return builder
         * 
         */
        public Builder cidrBlock(Output<String> cidrBlock) {
            $.cidrBlock = cidrBlock;
            return this;
        }

        /**
         * @param cidrBlock Workspace Corresponds to the Security Office Network of IPv4 Segment.
         * 
         * @return builder
         * 
         */
        public Builder cidrBlock(String cidrBlock) {
            return cidrBlock(Output.of(cidrBlock));
        }

        /**
         * @param desktopAccessType The method that you use to connect to cloud desktops. **Note:** The VPC connection method is provided by Alibaba Cloud PrivateLink. You are not charged for PrivateLink. When you set this parameter to VPC or Any, PrivateLink is automatically activated. Default value: `INTERNET`. Valid values:
         * 
         * @return builder
         * 
         */
        public Builder desktopAccessType(@Nullable Output<String> desktopAccessType) {
            $.desktopAccessType = desktopAccessType;
            return this;
        }

        /**
         * @param desktopAccessType The method that you use to connect to cloud desktops. **Note:** The VPC connection method is provided by Alibaba Cloud PrivateLink. You are not charged for PrivateLink. When you set this parameter to VPC or Any, PrivateLink is automatically activated. Default value: `INTERNET`. Valid values:
         * 
         * @return builder
         * 
         */
        public Builder desktopAccessType(String desktopAccessType) {
            return desktopAccessType(Output.of(desktopAccessType));
        }

        /**
         * @param dnsAddresses The IP address N of the DNS server of the enterprise AD system. You can specify only one IP address.
         * 
         * @return builder
         * 
         */
        public Builder dnsAddresses(Output<List<String>> dnsAddresses) {
            $.dnsAddresses = dnsAddresses;
            return this;
        }

        /**
         * @param dnsAddresses The IP address N of the DNS server of the enterprise AD system. You can specify only one IP address.
         * 
         * @return builder
         * 
         */
        public Builder dnsAddresses(List<String> dnsAddresses) {
            return dnsAddresses(Output.of(dnsAddresses));
        }

        /**
         * @param dnsAddresses The IP address N of the DNS server of the enterprise AD system. You can specify only one IP address.
         * 
         * @return builder
         * 
         */
        public Builder dnsAddresses(String... dnsAddresses) {
            return dnsAddresses(List.of(dnsAddresses));
        }

        /**
         * @param domainName The domain name of the enterprise AD system. You can register each domain name only once.
         * 
         * @return builder
         * 
         */
        public Builder domainName(Output<String> domainName) {
            $.domainName = domainName;
            return this;
        }

        /**
         * @param domainName The domain name of the enterprise AD system. You can register each domain name only once.
         * 
         * @return builder
         * 
         */
        public Builder domainName(String domainName) {
            return domainName(Output.of(domainName));
        }

        /**
         * @param domainPassword The password of the domain administrator. The password can be up to 64 characters in length.
         * 
         * @return builder
         * 
         */
        public Builder domainPassword(@Nullable Output<String> domainPassword) {
            $.domainPassword = domainPassword;
            return this;
        }

        /**
         * @param domainPassword The password of the domain administrator. The password can be up to 64 characters in length.
         * 
         * @return builder
         * 
         */
        public Builder domainPassword(String domainPassword) {
            return domainPassword(Output.of(domainPassword));
        }

        /**
         * @param domainUserName The username of the domain administrator. The username can be up to 64 characters in length.
         * 
         * @return builder
         * 
         */
        public Builder domainUserName(@Nullable Output<String> domainUserName) {
            $.domainUserName = domainUserName;
            return this;
        }

        /**
         * @param domainUserName The username of the domain administrator. The username can be up to 64 characters in length.
         * 
         * @return builder
         * 
         */
        public Builder domainUserName(String domainUserName) {
            return domainUserName(Output.of(domainUserName));
        }

        /**
         * @param enableAdminAccess Specifies whether to grant the permissions of the local administrator to the desktop users. Default value: true.
         * 
         * @return builder
         * 
         */
        public Builder enableAdminAccess(@Nullable Output<Boolean> enableAdminAccess) {
            $.enableAdminAccess = enableAdminAccess;
            return this;
        }

        /**
         * @param enableAdminAccess Specifies whether to grant the permissions of the local administrator to the desktop users. Default value: true.
         * 
         * @return builder
         * 
         */
        public Builder enableAdminAccess(Boolean enableAdminAccess) {
            return enableAdminAccess(Output.of(enableAdminAccess));
        }

        /**
         * @param enableInternetAccess Specifies whether to enable Internet access.
         * 
         * @return builder
         * 
         */
        public Builder enableInternetAccess(@Nullable Output<Boolean> enableInternetAccess) {
            $.enableInternetAccess = enableInternetAccess;
            return this;
        }

        /**
         * @param enableInternetAccess Specifies whether to enable Internet access.
         * 
         * @return builder
         * 
         */
        public Builder enableInternetAccess(Boolean enableInternetAccess) {
            return enableInternetAccess(Output.of(enableInternetAccess));
        }

        /**
         * @param mfaEnabled Specifies whether to enable multi-factor authentication (MFA).
         * 
         * @return builder
         * 
         */
        public Builder mfaEnabled(@Nullable Output<Boolean> mfaEnabled) {
            $.mfaEnabled = mfaEnabled;
            return this;
        }

        /**
         * @param mfaEnabled Specifies whether to enable multi-factor authentication (MFA).
         * 
         * @return builder
         * 
         */
        public Builder mfaEnabled(Boolean mfaEnabled) {
            return mfaEnabled(Output.of(mfaEnabled));
        }

        /**
         * @param protocolType The protocol type. Valid values: `ASP`, `HDX`.
         * 
         * @return builder
         * 
         */
        public Builder protocolType(@Nullable Output<String> protocolType) {
            $.protocolType = protocolType;
            return this;
        }

        /**
         * @param protocolType The protocol type. Valid values: `ASP`, `HDX`.
         * 
         * @return builder
         * 
         */
        public Builder protocolType(String protocolType) {
            return protocolType(Output.of(protocolType));
        }

        /**
         * @param specification The AD Connector specifications. Valid values: `1`, `2`.
         * 
         * @return builder
         * 
         */
        public Builder specification(@Nullable Output<Integer> specification) {
            $.specification = specification;
            return this;
        }

        /**
         * @param specification The AD Connector specifications. Valid values: `1`, `2`.
         * 
         * @return builder
         * 
         */
        public Builder specification(Integer specification) {
            return specification(Output.of(specification));
        }

        /**
         * @param subDomainDnsAddresses The DNS address N of the enterprise AD subdomain. If you specify a value for the `sub_domain_name` parameter but you do not specify a value for this parameter, the DNS address of the subdomain is the same as the DNS address of the parent domain.
         * 
         * @return builder
         * 
         */
        public Builder subDomainDnsAddresses(@Nullable Output<List<String>> subDomainDnsAddresses) {
            $.subDomainDnsAddresses = subDomainDnsAddresses;
            return this;
        }

        /**
         * @param subDomainDnsAddresses The DNS address N of the enterprise AD subdomain. If you specify a value for the `sub_domain_name` parameter but you do not specify a value for this parameter, the DNS address of the subdomain is the same as the DNS address of the parent domain.
         * 
         * @return builder
         * 
         */
        public Builder subDomainDnsAddresses(List<String> subDomainDnsAddresses) {
            return subDomainDnsAddresses(Output.of(subDomainDnsAddresses));
        }

        /**
         * @param subDomainDnsAddresses The DNS address N of the enterprise AD subdomain. If you specify a value for the `sub_domain_name` parameter but you do not specify a value for this parameter, the DNS address of the subdomain is the same as the DNS address of the parent domain.
         * 
         * @return builder
         * 
         */
        public Builder subDomainDnsAddresses(String... subDomainDnsAddresses) {
            return subDomainDnsAddresses(List.of(subDomainDnsAddresses));
        }

        /**
         * @param subDomainName The domain name of the enterprise AD subdomain.
         * 
         * @return builder
         * 
         */
        public Builder subDomainName(@Nullable Output<String> subDomainName) {
            $.subDomainName = subDomainName;
            return this;
        }

        /**
         * @param subDomainName The domain name of the enterprise AD subdomain.
         * 
         * @return builder
         * 
         */
        public Builder subDomainName(String subDomainName) {
            return subDomainName(Output.of(subDomainName));
        }

        /**
         * @param verifyCode The verification code. If the CEN instance that you specify for the CenId parameter belongs to another Alibaba Cloud account, you must call the SendVerifyCode operation to obtain the verification code.
         * 
         * @return builder
         * 
         */
        public Builder verifyCode(@Nullable Output<String> verifyCode) {
            $.verifyCode = verifyCode;
            return this;
        }

        /**
         * @param verifyCode The verification code. If the CEN instance that you specify for the CenId parameter belongs to another Alibaba Cloud account, you must call the SendVerifyCode operation to obtain the verification code.
         * 
         * @return builder
         * 
         */
        public Builder verifyCode(String verifyCode) {
            return verifyCode(Output.of(verifyCode));
        }

        public AdConnectorOfficeSiteArgs build() {
            $.adConnectorOfficeSiteName = Objects.requireNonNull($.adConnectorOfficeSiteName, "expected parameter 'adConnectorOfficeSiteName' to be non-null");
            $.cenId = Objects.requireNonNull($.cenId, "expected parameter 'cenId' to be non-null");
            $.cidrBlock = Objects.requireNonNull($.cidrBlock, "expected parameter 'cidrBlock' to be non-null");
            $.dnsAddresses = Objects.requireNonNull($.dnsAddresses, "expected parameter 'dnsAddresses' to be non-null");
            $.domainName = Objects.requireNonNull($.domainName, "expected parameter 'domainName' to be non-null");
            return $;
        }
    }

}
