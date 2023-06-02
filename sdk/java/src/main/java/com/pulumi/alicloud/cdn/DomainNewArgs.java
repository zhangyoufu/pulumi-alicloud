// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.cdn;

import com.pulumi.alicloud.cdn.inputs.DomainNewCertificateConfigArgs;
import com.pulumi.alicloud.cdn.inputs.DomainNewSourceArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Object;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class DomainNewArgs extends com.pulumi.resources.ResourceArgs {

    public static final DomainNewArgs Empty = new DomainNewArgs();

    /**
     * Cdn type of the accelerated domain. Valid values are `web`, `download`, `video`.
     * 
     */
    @Import(name="cdnType", required=true)
    private Output<String> cdnType;

    /**
     * @return Cdn type of the accelerated domain. Valid values are `web`, `download`, `video`.
     * 
     */
    public Output<String> cdnType() {
        return this.cdnType;
    }

    /**
     * Certificate configuration. See the following `Block CertificateConfig`.
     * 
     */
    @Import(name="certificateConfig")
    private @Nullable Output<DomainNewCertificateConfigArgs> certificateConfig;

    /**
     * @return Certificate configuration. See the following `Block CertificateConfig`.
     * 
     */
    public Optional<Output<DomainNewCertificateConfigArgs>> certificateConfig() {
        return Optional.ofNullable(this.certificateConfig);
    }

    /**
     * Health test URL.
     * 
     */
    @Import(name="checkUrl")
    private @Nullable Output<String> checkUrl;

    /**
     * @return Health test URL.
     * 
     */
    public Optional<Output<String>> checkUrl() {
        return Optional.ofNullable(this.checkUrl);
    }

    /**
     * Name of the accelerated domain. This name without suffix can have a string of 1 to 63 characters, must contain only alphanumeric characters or &#34;-&#34;, and must not begin or end with &#34;-&#34;, and &#34;-&#34; must not in the 3th and 4th character positions at the same time. Suffix `.sh` and `.tel` are not supported.
     * 
     */
    @Import(name="domainName", required=true)
    private Output<String> domainName;

    /**
     * @return Name of the accelerated domain. This name without suffix can have a string of 1 to 63 characters, must contain only alphanumeric characters or &#34;-&#34;, and must not begin or end with &#34;-&#34;, and &#34;-&#34; must not in the 3th and 4th character positions at the same time. Suffix `.sh` and `.tel` are not supported.
     * 
     */
    public Output<String> domainName() {
        return this.domainName;
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
     * Scope of the accelerated domain. Valid values are `domestic`, `overseas`, `global`. Default value is `domestic`. This parameter&#39;s setting is valid Only for the international users and domestic L3 and above users. Value:
     * - **domestic**: Mainland China only.
     * - **overseas**: Global (excluding Mainland China).
     * - **global**: global.
     *   The default value is **domestic**.
     * 
     */
    @Import(name="scope")
    private @Nullable Output<String> scope;

    /**
     * @return Scope of the accelerated domain. Valid values are `domestic`, `overseas`, `global`. Default value is `domestic`. This parameter&#39;s setting is valid Only for the international users and domestic L3 and above users. Value:
     * - **domestic**: Mainland China only.
     * - **overseas**: Global (excluding Mainland China).
     * - **global**: global.
     *   The default value is **domestic**.
     * 
     */
    public Optional<Output<String>> scope() {
        return Optional.ofNullable(this.scope);
    }

    /**
     * The source address list of the accelerated domain. Defaults to null. See the following `Block Sources`.
     * 
     */
    @Import(name="sources", required=true)
    private Output<List<DomainNewSourceArgs>> sources;

    /**
     * @return The source address list of the accelerated domain. Defaults to null. See the following `Block Sources`.
     * 
     */
    public Output<List<DomainNewSourceArgs>> sources() {
        return this.sources;
    }

    /**
     * The tag of the resource.
     * 
     */
    @Import(name="tags")
    private @Nullable Output<Map<String,Object>> tags;

    /**
     * @return The tag of the resource.
     * 
     */
    public Optional<Output<Map<String,Object>>> tags() {
        return Optional.ofNullable(this.tags);
    }

    private DomainNewArgs() {}

    private DomainNewArgs(DomainNewArgs $) {
        this.cdnType = $.cdnType;
        this.certificateConfig = $.certificateConfig;
        this.checkUrl = $.checkUrl;
        this.domainName = $.domainName;
        this.resourceGroupId = $.resourceGroupId;
        this.scope = $.scope;
        this.sources = $.sources;
        this.tags = $.tags;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(DomainNewArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private DomainNewArgs $;

        public Builder() {
            $ = new DomainNewArgs();
        }

        public Builder(DomainNewArgs defaults) {
            $ = new DomainNewArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param cdnType Cdn type of the accelerated domain. Valid values are `web`, `download`, `video`.
         * 
         * @return builder
         * 
         */
        public Builder cdnType(Output<String> cdnType) {
            $.cdnType = cdnType;
            return this;
        }

        /**
         * @param cdnType Cdn type of the accelerated domain. Valid values are `web`, `download`, `video`.
         * 
         * @return builder
         * 
         */
        public Builder cdnType(String cdnType) {
            return cdnType(Output.of(cdnType));
        }

        /**
         * @param certificateConfig Certificate configuration. See the following `Block CertificateConfig`.
         * 
         * @return builder
         * 
         */
        public Builder certificateConfig(@Nullable Output<DomainNewCertificateConfigArgs> certificateConfig) {
            $.certificateConfig = certificateConfig;
            return this;
        }

        /**
         * @param certificateConfig Certificate configuration. See the following `Block CertificateConfig`.
         * 
         * @return builder
         * 
         */
        public Builder certificateConfig(DomainNewCertificateConfigArgs certificateConfig) {
            return certificateConfig(Output.of(certificateConfig));
        }

        /**
         * @param checkUrl Health test URL.
         * 
         * @return builder
         * 
         */
        public Builder checkUrl(@Nullable Output<String> checkUrl) {
            $.checkUrl = checkUrl;
            return this;
        }

        /**
         * @param checkUrl Health test URL.
         * 
         * @return builder
         * 
         */
        public Builder checkUrl(String checkUrl) {
            return checkUrl(Output.of(checkUrl));
        }

        /**
         * @param domainName Name of the accelerated domain. This name without suffix can have a string of 1 to 63 characters, must contain only alphanumeric characters or &#34;-&#34;, and must not begin or end with &#34;-&#34;, and &#34;-&#34; must not in the 3th and 4th character positions at the same time. Suffix `.sh` and `.tel` are not supported.
         * 
         * @return builder
         * 
         */
        public Builder domainName(Output<String> domainName) {
            $.domainName = domainName;
            return this;
        }

        /**
         * @param domainName Name of the accelerated domain. This name without suffix can have a string of 1 to 63 characters, must contain only alphanumeric characters or &#34;-&#34;, and must not begin or end with &#34;-&#34;, and &#34;-&#34; must not in the 3th and 4th character positions at the same time. Suffix `.sh` and `.tel` are not supported.
         * 
         * @return builder
         * 
         */
        public Builder domainName(String domainName) {
            return domainName(Output.of(domainName));
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
         * @param scope Scope of the accelerated domain. Valid values are `domestic`, `overseas`, `global`. Default value is `domestic`. This parameter&#39;s setting is valid Only for the international users and domestic L3 and above users. Value:
         * - **domestic**: Mainland China only.
         * - **overseas**: Global (excluding Mainland China).
         * - **global**: global.
         *   The default value is **domestic**.
         * 
         * @return builder
         * 
         */
        public Builder scope(@Nullable Output<String> scope) {
            $.scope = scope;
            return this;
        }

        /**
         * @param scope Scope of the accelerated domain. Valid values are `domestic`, `overseas`, `global`. Default value is `domestic`. This parameter&#39;s setting is valid Only for the international users and domestic L3 and above users. Value:
         * - **domestic**: Mainland China only.
         * - **overseas**: Global (excluding Mainland China).
         * - **global**: global.
         *   The default value is **domestic**.
         * 
         * @return builder
         * 
         */
        public Builder scope(String scope) {
            return scope(Output.of(scope));
        }

        /**
         * @param sources The source address list of the accelerated domain. Defaults to null. See the following `Block Sources`.
         * 
         * @return builder
         * 
         */
        public Builder sources(Output<List<DomainNewSourceArgs>> sources) {
            $.sources = sources;
            return this;
        }

        /**
         * @param sources The source address list of the accelerated domain. Defaults to null. See the following `Block Sources`.
         * 
         * @return builder
         * 
         */
        public Builder sources(List<DomainNewSourceArgs> sources) {
            return sources(Output.of(sources));
        }

        /**
         * @param sources The source address list of the accelerated domain. Defaults to null. See the following `Block Sources`.
         * 
         * @return builder
         * 
         */
        public Builder sources(DomainNewSourceArgs... sources) {
            return sources(List.of(sources));
        }

        /**
         * @param tags The tag of the resource.
         * 
         * @return builder
         * 
         */
        public Builder tags(@Nullable Output<Map<String,Object>> tags) {
            $.tags = tags;
            return this;
        }

        /**
         * @param tags The tag of the resource.
         * 
         * @return builder
         * 
         */
        public Builder tags(Map<String,Object> tags) {
            return tags(Output.of(tags));
        }

        public DomainNewArgs build() {
            $.cdnType = Objects.requireNonNull($.cdnType, "expected parameter 'cdnType' to be non-null");
            $.domainName = Objects.requireNonNull($.domainName, "expected parameter 'domainName' to be non-null");
            $.sources = Objects.requireNonNull($.sources, "expected parameter 'sources' to be non-null");
            return $;
        }
    }

}
