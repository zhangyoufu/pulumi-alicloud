// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.nlb;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Object;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class SecurityPolicyArgs extends com.pulumi.resources.ResourceArgs {

    public static final SecurityPolicyArgs Empty = new SecurityPolicyArgs();

    /**
     * The supported cipher suites, which are determined by the TLS protocol version. You can specify at most 32 cipher suites.
     * - TLS 1.0 and TLS 1.1 support the following cipher suites: `ECDHE-ECDSA-AES128-SHA`, `ECDHE-ECDSA-AES256-SHA`, `ECDHE-RSA-AES128-SHA`, `ECDHE-RSA-AES256-SHA`, `AES128-SHA`, `AES256-SHA`, `DES-CBC3-SHA`
     * - TLS 1.2 supports the following cipher suites: `ECDHE-ECDSA-AES128-SHA`, `ECDHE-ECDSA-AES256-SHA`, `ECDHE-RSA-AES128-SHA`, `ECDHE-RSA-AES256-SHA`, `AES128-SHA`, `AES256-SHA, DES-CBC3-SHA`, `ECDHE-ECDSA-AES128-GCM-SHA256`, `ECDHE-ECDSA-AES256-GCM-SHA384`, `ECDHE-ECDSA-AES128-SHA256`, `ECDHE-ECDSA-AES256-SHA384`, `ECDHE-RSA-AES128-GCM-SHA256`, `ECDHE-RSA-AES256-GCM-SHA384`, `ECDHE-RSA-AES128-SHA256`, `ECDHE-RSA-AES256-SHA384`, `AES128-GCM-SHA256`, `AES256-GCM-SHA384`, `AES128-SHA256`, `AES256-SHA256`
     * - TLS 1.3 supports the following cipher suites: `TLS_AES_128_GCM_SHA256`, `TLS_AES_256_GCM_SHA384`, `TLS_CHACHA20_POLY1305_SHA256`, `TLS_AES_128_CCM_SHA256`, `TLS_AES_128_CCM_8_SHA256`
     * 
     */
    @Import(name="ciphers", required=true)
    private Output<List<String>> ciphers;

    /**
     * @return The supported cipher suites, which are determined by the TLS protocol version. You can specify at most 32 cipher suites.
     * - TLS 1.0 and TLS 1.1 support the following cipher suites: `ECDHE-ECDSA-AES128-SHA`, `ECDHE-ECDSA-AES256-SHA`, `ECDHE-RSA-AES128-SHA`, `ECDHE-RSA-AES256-SHA`, `AES128-SHA`, `AES256-SHA`, `DES-CBC3-SHA`
     * - TLS 1.2 supports the following cipher suites: `ECDHE-ECDSA-AES128-SHA`, `ECDHE-ECDSA-AES256-SHA`, `ECDHE-RSA-AES128-SHA`, `ECDHE-RSA-AES256-SHA`, `AES128-SHA`, `AES256-SHA, DES-CBC3-SHA`, `ECDHE-ECDSA-AES128-GCM-SHA256`, `ECDHE-ECDSA-AES256-GCM-SHA384`, `ECDHE-ECDSA-AES128-SHA256`, `ECDHE-ECDSA-AES256-SHA384`, `ECDHE-RSA-AES128-GCM-SHA256`, `ECDHE-RSA-AES256-GCM-SHA384`, `ECDHE-RSA-AES128-SHA256`, `ECDHE-RSA-AES256-SHA384`, `AES128-GCM-SHA256`, `AES256-GCM-SHA384`, `AES128-SHA256`, `AES256-SHA256`
     * - TLS 1.3 supports the following cipher suites: `TLS_AES_128_GCM_SHA256`, `TLS_AES_256_GCM_SHA384`, `TLS_CHACHA20_POLY1305_SHA256`, `TLS_AES_128_CCM_SHA256`, `TLS_AES_128_CCM_8_SHA256`
     * 
     */
    public Output<List<String>> ciphers() {
        return this.ciphers;
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
     * The name of the security policy. The name must be 1 to 200 characters in length, and can contain letters, digits, periods (.), underscores (_), and hyphens (-).
     * 
     */
    @Import(name="securityPolicyName")
    private @Nullable Output<String> securityPolicyName;

    /**
     * @return The name of the security policy. The name must be 1 to 200 characters in length, and can contain letters, digits, periods (.), underscores (_), and hyphens (-).
     * 
     */
    public Optional<Output<String>> securityPolicyName() {
        return Optional.ofNullable(this.securityPolicyName);
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
     * The supported versions of the Transport Layer Security (TLS) protocol. Valid values: `TLSv1.0`, `TLSv1.1`, `TLSv1.2`, and `TLSv1.3`. You can specify at most four TLS versions.
     * 
     */
    @Import(name="tlsVersions", required=true)
    private Output<List<String>> tlsVersions;

    /**
     * @return The supported versions of the Transport Layer Security (TLS) protocol. Valid values: `TLSv1.0`, `TLSv1.1`, `TLSv1.2`, and `TLSv1.3`. You can specify at most four TLS versions.
     * 
     */
    public Output<List<String>> tlsVersions() {
        return this.tlsVersions;
    }

    private SecurityPolicyArgs() {}

    private SecurityPolicyArgs(SecurityPolicyArgs $) {
        this.ciphers = $.ciphers;
        this.resourceGroupId = $.resourceGroupId;
        this.securityPolicyName = $.securityPolicyName;
        this.tags = $.tags;
        this.tlsVersions = $.tlsVersions;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(SecurityPolicyArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private SecurityPolicyArgs $;

        public Builder() {
            $ = new SecurityPolicyArgs();
        }

        public Builder(SecurityPolicyArgs defaults) {
            $ = new SecurityPolicyArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param ciphers The supported cipher suites, which are determined by the TLS protocol version. You can specify at most 32 cipher suites.
         * - TLS 1.0 and TLS 1.1 support the following cipher suites: `ECDHE-ECDSA-AES128-SHA`, `ECDHE-ECDSA-AES256-SHA`, `ECDHE-RSA-AES128-SHA`, `ECDHE-RSA-AES256-SHA`, `AES128-SHA`, `AES256-SHA`, `DES-CBC3-SHA`
         * - TLS 1.2 supports the following cipher suites: `ECDHE-ECDSA-AES128-SHA`, `ECDHE-ECDSA-AES256-SHA`, `ECDHE-RSA-AES128-SHA`, `ECDHE-RSA-AES256-SHA`, `AES128-SHA`, `AES256-SHA, DES-CBC3-SHA`, `ECDHE-ECDSA-AES128-GCM-SHA256`, `ECDHE-ECDSA-AES256-GCM-SHA384`, `ECDHE-ECDSA-AES128-SHA256`, `ECDHE-ECDSA-AES256-SHA384`, `ECDHE-RSA-AES128-GCM-SHA256`, `ECDHE-RSA-AES256-GCM-SHA384`, `ECDHE-RSA-AES128-SHA256`, `ECDHE-RSA-AES256-SHA384`, `AES128-GCM-SHA256`, `AES256-GCM-SHA384`, `AES128-SHA256`, `AES256-SHA256`
         * - TLS 1.3 supports the following cipher suites: `TLS_AES_128_GCM_SHA256`, `TLS_AES_256_GCM_SHA384`, `TLS_CHACHA20_POLY1305_SHA256`, `TLS_AES_128_CCM_SHA256`, `TLS_AES_128_CCM_8_SHA256`
         * 
         * @return builder
         * 
         */
        public Builder ciphers(Output<List<String>> ciphers) {
            $.ciphers = ciphers;
            return this;
        }

        /**
         * @param ciphers The supported cipher suites, which are determined by the TLS protocol version. You can specify at most 32 cipher suites.
         * - TLS 1.0 and TLS 1.1 support the following cipher suites: `ECDHE-ECDSA-AES128-SHA`, `ECDHE-ECDSA-AES256-SHA`, `ECDHE-RSA-AES128-SHA`, `ECDHE-RSA-AES256-SHA`, `AES128-SHA`, `AES256-SHA`, `DES-CBC3-SHA`
         * - TLS 1.2 supports the following cipher suites: `ECDHE-ECDSA-AES128-SHA`, `ECDHE-ECDSA-AES256-SHA`, `ECDHE-RSA-AES128-SHA`, `ECDHE-RSA-AES256-SHA`, `AES128-SHA`, `AES256-SHA, DES-CBC3-SHA`, `ECDHE-ECDSA-AES128-GCM-SHA256`, `ECDHE-ECDSA-AES256-GCM-SHA384`, `ECDHE-ECDSA-AES128-SHA256`, `ECDHE-ECDSA-AES256-SHA384`, `ECDHE-RSA-AES128-GCM-SHA256`, `ECDHE-RSA-AES256-GCM-SHA384`, `ECDHE-RSA-AES128-SHA256`, `ECDHE-RSA-AES256-SHA384`, `AES128-GCM-SHA256`, `AES256-GCM-SHA384`, `AES128-SHA256`, `AES256-SHA256`
         * - TLS 1.3 supports the following cipher suites: `TLS_AES_128_GCM_SHA256`, `TLS_AES_256_GCM_SHA384`, `TLS_CHACHA20_POLY1305_SHA256`, `TLS_AES_128_CCM_SHA256`, `TLS_AES_128_CCM_8_SHA256`
         * 
         * @return builder
         * 
         */
        public Builder ciphers(List<String> ciphers) {
            return ciphers(Output.of(ciphers));
        }

        /**
         * @param ciphers The supported cipher suites, which are determined by the TLS protocol version. You can specify at most 32 cipher suites.
         * - TLS 1.0 and TLS 1.1 support the following cipher suites: `ECDHE-ECDSA-AES128-SHA`, `ECDHE-ECDSA-AES256-SHA`, `ECDHE-RSA-AES128-SHA`, `ECDHE-RSA-AES256-SHA`, `AES128-SHA`, `AES256-SHA`, `DES-CBC3-SHA`
         * - TLS 1.2 supports the following cipher suites: `ECDHE-ECDSA-AES128-SHA`, `ECDHE-ECDSA-AES256-SHA`, `ECDHE-RSA-AES128-SHA`, `ECDHE-RSA-AES256-SHA`, `AES128-SHA`, `AES256-SHA, DES-CBC3-SHA`, `ECDHE-ECDSA-AES128-GCM-SHA256`, `ECDHE-ECDSA-AES256-GCM-SHA384`, `ECDHE-ECDSA-AES128-SHA256`, `ECDHE-ECDSA-AES256-SHA384`, `ECDHE-RSA-AES128-GCM-SHA256`, `ECDHE-RSA-AES256-GCM-SHA384`, `ECDHE-RSA-AES128-SHA256`, `ECDHE-RSA-AES256-SHA384`, `AES128-GCM-SHA256`, `AES256-GCM-SHA384`, `AES128-SHA256`, `AES256-SHA256`
         * - TLS 1.3 supports the following cipher suites: `TLS_AES_128_GCM_SHA256`, `TLS_AES_256_GCM_SHA384`, `TLS_CHACHA20_POLY1305_SHA256`, `TLS_AES_128_CCM_SHA256`, `TLS_AES_128_CCM_8_SHA256`
         * 
         * @return builder
         * 
         */
        public Builder ciphers(String... ciphers) {
            return ciphers(List.of(ciphers));
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
         * @param securityPolicyName The name of the security policy. The name must be 1 to 200 characters in length, and can contain letters, digits, periods (.), underscores (_), and hyphens (-).
         * 
         * @return builder
         * 
         */
        public Builder securityPolicyName(@Nullable Output<String> securityPolicyName) {
            $.securityPolicyName = securityPolicyName;
            return this;
        }

        /**
         * @param securityPolicyName The name of the security policy. The name must be 1 to 200 characters in length, and can contain letters, digits, periods (.), underscores (_), and hyphens (-).
         * 
         * @return builder
         * 
         */
        public Builder securityPolicyName(String securityPolicyName) {
            return securityPolicyName(Output.of(securityPolicyName));
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
         * @param tlsVersions The supported versions of the Transport Layer Security (TLS) protocol. Valid values: `TLSv1.0`, `TLSv1.1`, `TLSv1.2`, and `TLSv1.3`. You can specify at most four TLS versions.
         * 
         * @return builder
         * 
         */
        public Builder tlsVersions(Output<List<String>> tlsVersions) {
            $.tlsVersions = tlsVersions;
            return this;
        }

        /**
         * @param tlsVersions The supported versions of the Transport Layer Security (TLS) protocol. Valid values: `TLSv1.0`, `TLSv1.1`, `TLSv1.2`, and `TLSv1.3`. You can specify at most four TLS versions.
         * 
         * @return builder
         * 
         */
        public Builder tlsVersions(List<String> tlsVersions) {
            return tlsVersions(Output.of(tlsVersions));
        }

        /**
         * @param tlsVersions The supported versions of the Transport Layer Security (TLS) protocol. Valid values: `TLSv1.0`, `TLSv1.1`, `TLSv1.2`, and `TLSv1.3`. You can specify at most four TLS versions.
         * 
         * @return builder
         * 
         */
        public Builder tlsVersions(String... tlsVersions) {
            return tlsVersions(List.of(tlsVersions));
        }

        public SecurityPolicyArgs build() {
            $.ciphers = Objects.requireNonNull($.ciphers, "expected parameter 'ciphers' to be non-null");
            $.tlsVersions = Objects.requireNonNull($.tlsVersions, "expected parameter 'tlsVersions' to be non-null");
            return $;
        }
    }

}
