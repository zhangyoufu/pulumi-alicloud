// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud;

import com.pulumi.alicloud.inputs.ProviderAssumeRoleArgs;
import com.pulumi.alicloud.inputs.ProviderEndpointArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.core.internal.Codegen;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ProviderArgs extends com.pulumi.resources.ResourceArgs {

    public static final ProviderArgs Empty = new ProviderArgs();

    /**
     * The access key for API operations. You can retrieve this from the &#39;Security Management&#39; section of the Alibaba Cloud
     * console.
     * 
     */
    @Import(name="accessKey")
    private @Nullable Output<String> accessKey;

    /**
     * @return The access key for API operations. You can retrieve this from the &#39;Security Management&#39; section of the Alibaba Cloud
     * console.
     * 
     */
    public Optional<Output<String>> accessKey() {
        return Optional.ofNullable(this.accessKey);
    }

    /**
     * The account ID for some service API operations. You can retrieve this from the &#39;Security Settings&#39; section of the
     * Alibaba Cloud console.
     * 
     */
    @Import(name="accountId")
    private @Nullable Output<String> accountId;

    /**
     * @return The account ID for some service API operations. You can retrieve this from the &#39;Security Settings&#39; section of the
     * Alibaba Cloud console.
     * 
     */
    public Optional<Output<String>> accountId() {
        return Optional.ofNullable(this.accountId);
    }

    @Import(name="assumeRole", json=true)
    private @Nullable Output<ProviderAssumeRoleArgs> assumeRole;

    public Optional<Output<ProviderAssumeRoleArgs>> assumeRole() {
        return Optional.ofNullable(this.assumeRole);
    }

    /**
     * The maximum timeout of the client connection server.
     * 
     */
    @Import(name="clientConnectTimeout", json=true)
    private @Nullable Output<Integer> clientConnectTimeout;

    /**
     * @return The maximum timeout of the client connection server.
     * 
     */
    public Optional<Output<Integer>> clientConnectTimeout() {
        return Optional.ofNullable(this.clientConnectTimeout);
    }

    /**
     * The maximum timeout of the client read request.
     * 
     */
    @Import(name="clientReadTimeout", json=true)
    private @Nullable Output<Integer> clientReadTimeout;

    /**
     * @return The maximum timeout of the client read request.
     * 
     */
    public Optional<Output<Integer>> clientReadTimeout() {
        return Optional.ofNullable(this.clientReadTimeout);
    }

    /**
     * Use this to mark a terraform configuration file source.
     * 
     */
    @Import(name="configurationSource")
    private @Nullable Output<String> configurationSource;

    /**
     * @return Use this to mark a terraform configuration file source.
     * 
     */
    public Optional<Output<String>> configurationSource() {
        return Optional.ofNullable(this.configurationSource);
    }

    /**
     * The URI of sidecar credentials service.
     * 
     */
    @Import(name="credentialsUri")
    private @Nullable Output<String> credentialsUri;

    /**
     * @return The URI of sidecar credentials service.
     * 
     */
    public Optional<Output<String>> credentialsUri() {
        return Optional.ofNullable(this.credentialsUri);
    }

    /**
     * The RAM Role Name attached on a ECS instance for API operations. You can retrieve this from the &#39;Access Control&#39; section
     * of the Alibaba Cloud console.
     * 
     */
    @Import(name="ecsRoleName")
    private @Nullable Output<String> ecsRoleName;

    /**
     * @return The RAM Role Name attached on a ECS instance for API operations. You can retrieve this from the &#39;Access Control&#39; section
     * of the Alibaba Cloud console.
     * 
     */
    public Optional<Output<String>> ecsRoleName() {
        return Optional.ofNullable(this.ecsRoleName);
    }

    @Import(name="endpoints", json=true)
    private @Nullable Output<List<ProviderEndpointArgs>> endpoints;

    public Optional<Output<List<ProviderEndpointArgs>>> endpoints() {
        return Optional.ofNullable(this.endpoints);
    }

    /**
     * @deprecated
     * Field &#39;fc&#39; has been deprecated from provider version 1.28.0. New field &#39;fc&#39; which in nested endpoints instead.
     * 
     */
    @Deprecated /* Field 'fc' has been deprecated from provider version 1.28.0. New field 'fc' which in nested endpoints instead. */
    @Import(name="fc")
    private @Nullable Output<String> fc;

    /**
     * @deprecated
     * Field &#39;fc&#39; has been deprecated from provider version 1.28.0. New field &#39;fc&#39; which in nested endpoints instead.
     * 
     */
    @Deprecated /* Field 'fc' has been deprecated from provider version 1.28.0. New field 'fc' which in nested endpoints instead. */
    public Optional<Output<String>> fc() {
        return Optional.ofNullable(this.fc);
    }

    /**
     * @deprecated
     * Field &#39;log_endpoint&#39; has been deprecated from provider version 1.28.0. New field &#39;log&#39; which in nested endpoints instead.
     * 
     */
    @Deprecated /* Field 'log_endpoint' has been deprecated from provider version 1.28.0. New field 'log' which in nested endpoints instead. */
    @Import(name="logEndpoint")
    private @Nullable Output<String> logEndpoint;

    /**
     * @deprecated
     * Field &#39;log_endpoint&#39; has been deprecated from provider version 1.28.0. New field &#39;log&#39; which in nested endpoints instead.
     * 
     */
    @Deprecated /* Field 'log_endpoint' has been deprecated from provider version 1.28.0. New field 'log' which in nested endpoints instead. */
    public Optional<Output<String>> logEndpoint() {
        return Optional.ofNullable(this.logEndpoint);
    }

    /**
     * The maximum retry timeout of the request.
     * 
     */
    @Import(name="maxRetryTimeout", json=true)
    private @Nullable Output<Integer> maxRetryTimeout;

    /**
     * @return The maximum retry timeout of the request.
     * 
     */
    public Optional<Output<Integer>> maxRetryTimeout() {
        return Optional.ofNullable(this.maxRetryTimeout);
    }

    /**
     * @deprecated
     * Field &#39;mns_endpoint&#39; has been deprecated from provider version 1.28.0. New field &#39;mns&#39; which in nested endpoints instead.
     * 
     */
    @Deprecated /* Field 'mns_endpoint' has been deprecated from provider version 1.28.0. New field 'mns' which in nested endpoints instead. */
    @Import(name="mnsEndpoint")
    private @Nullable Output<String> mnsEndpoint;

    /**
     * @deprecated
     * Field &#39;mns_endpoint&#39; has been deprecated from provider version 1.28.0. New field &#39;mns&#39; which in nested endpoints instead.
     * 
     */
    @Deprecated /* Field 'mns_endpoint' has been deprecated from provider version 1.28.0. New field 'mns' which in nested endpoints instead. */
    public Optional<Output<String>> mnsEndpoint() {
        return Optional.ofNullable(this.mnsEndpoint);
    }

    /**
     * @deprecated
     * Field &#39;ots_instance_name&#39; has been deprecated from provider version 1.10.0. New field &#39;instance_name&#39; of resource &#39;alicloud_ots_table&#39; instead.
     * 
     */
    @Deprecated /* Field 'ots_instance_name' has been deprecated from provider version 1.10.0. New field 'instance_name' of resource 'alicloud_ots_table' instead. */
    @Import(name="otsInstanceName")
    private @Nullable Output<String> otsInstanceName;

    /**
     * @deprecated
     * Field &#39;ots_instance_name&#39; has been deprecated from provider version 1.10.0. New field &#39;instance_name&#39; of resource &#39;alicloud_ots_table&#39; instead.
     * 
     */
    @Deprecated /* Field 'ots_instance_name' has been deprecated from provider version 1.10.0. New field 'instance_name' of resource 'alicloud_ots_table' instead. */
    public Optional<Output<String>> otsInstanceName() {
        return Optional.ofNullable(this.otsInstanceName);
    }

    /**
     * The profile for API operations. If not set, the default profile created with `aliyun configure` will be used.
     * 
     */
    @Import(name="profile")
    private @Nullable Output<String> profile;

    /**
     * @return The profile for API operations. If not set, the default profile created with `aliyun configure` will be used.
     * 
     */
    public Optional<Output<String>> profile() {
        return Optional.ofNullable(this.profile);
    }

    @Import(name="protocol")
    private @Nullable Output<String> protocol;

    public Optional<Output<String>> protocol() {
        return Optional.ofNullable(this.protocol);
    }

    /**
     * The region where Alibaba Cloud operations will take place. Examples are cn-beijing, cn-hangzhou, eu-central-1, etc.
     * 
     */
    @Import(name="region")
    private @Nullable Output<String> region;

    /**
     * @return The region where Alibaba Cloud operations will take place. Examples are cn-beijing, cn-hangzhou, eu-central-1, etc.
     * 
     */
    public Optional<Output<String>> region() {
        return Optional.ofNullable(this.region);
    }

    /**
     * The secret key for API operations. You can retrieve this from the &#39;Security Management&#39; section of the Alibaba Cloud
     * console.
     * 
     */
    @Import(name="secretKey")
    private @Nullable Output<String> secretKey;

    /**
     * @return The secret key for API operations. You can retrieve this from the &#39;Security Management&#39; section of the Alibaba Cloud
     * console.
     * 
     */
    public Optional<Output<String>> secretKey() {
        return Optional.ofNullable(this.secretKey);
    }

    /**
     * The security transport for the assume role invoking.
     * 
     */
    @Import(name="secureTransport")
    private @Nullable Output<String> secureTransport;

    /**
     * @return The security transport for the assume role invoking.
     * 
     */
    public Optional<Output<String>> secureTransport() {
        return Optional.ofNullable(this.secureTransport);
    }

    /**
     * security token. A security token is only required if you are using Security Token Service.
     * 
     */
    @Import(name="securityToken")
    private @Nullable Output<String> securityToken;

    /**
     * @return security token. A security token is only required if you are using Security Token Service.
     * 
     */
    public Optional<Output<String>> securityToken() {
        return Optional.ofNullable(this.securityToken);
    }

    @Import(name="securityTransport")
    private @Nullable Output<String> securityTransport;

    public Optional<Output<String>> securityTransport() {
        return Optional.ofNullable(this.securityTransport);
    }

    /**
     * The path to the shared credentials file. If not set this defaults to ~/.aliyun/config.json
     * 
     */
    @Import(name="sharedCredentialsFile")
    private @Nullable Output<String> sharedCredentialsFile;

    /**
     * @return The path to the shared credentials file. If not set this defaults to ~/.aliyun/config.json
     * 
     */
    public Optional<Output<String>> sharedCredentialsFile() {
        return Optional.ofNullable(this.sharedCredentialsFile);
    }

    /**
     * Skip static validation of region ID. Used by users of alternative AlibabaCloud-like APIs or users w/ access to regions
     * that are not public (yet).
     * 
     */
    @Import(name="skipRegionValidation", json=true)
    private @Nullable Output<Boolean> skipRegionValidation;

    /**
     * @return Skip static validation of region ID. Used by users of alternative AlibabaCloud-like APIs or users w/ access to regions
     * that are not public (yet).
     * 
     */
    public Optional<Output<Boolean>> skipRegionValidation() {
        return Optional.ofNullable(this.skipRegionValidation);
    }

    /**
     * The source ip for the assume role invoking.
     * 
     */
    @Import(name="sourceIp")
    private @Nullable Output<String> sourceIp;

    /**
     * @return The source ip for the assume role invoking.
     * 
     */
    public Optional<Output<String>> sourceIp() {
        return Optional.ofNullable(this.sourceIp);
    }

    private ProviderArgs() {}

    private ProviderArgs(ProviderArgs $) {
        this.accessKey = $.accessKey;
        this.accountId = $.accountId;
        this.assumeRole = $.assumeRole;
        this.clientConnectTimeout = $.clientConnectTimeout;
        this.clientReadTimeout = $.clientReadTimeout;
        this.configurationSource = $.configurationSource;
        this.credentialsUri = $.credentialsUri;
        this.ecsRoleName = $.ecsRoleName;
        this.endpoints = $.endpoints;
        this.fc = $.fc;
        this.logEndpoint = $.logEndpoint;
        this.maxRetryTimeout = $.maxRetryTimeout;
        this.mnsEndpoint = $.mnsEndpoint;
        this.otsInstanceName = $.otsInstanceName;
        this.profile = $.profile;
        this.protocol = $.protocol;
        this.region = $.region;
        this.secretKey = $.secretKey;
        this.secureTransport = $.secureTransport;
        this.securityToken = $.securityToken;
        this.securityTransport = $.securityTransport;
        this.sharedCredentialsFile = $.sharedCredentialsFile;
        this.skipRegionValidation = $.skipRegionValidation;
        this.sourceIp = $.sourceIp;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ProviderArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ProviderArgs $;

        public Builder() {
            $ = new ProviderArgs();
        }

        public Builder(ProviderArgs defaults) {
            $ = new ProviderArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param accessKey The access key for API operations. You can retrieve this from the &#39;Security Management&#39; section of the Alibaba Cloud
         * console.
         * 
         * @return builder
         * 
         */
        public Builder accessKey(@Nullable Output<String> accessKey) {
            $.accessKey = accessKey;
            return this;
        }

        /**
         * @param accessKey The access key for API operations. You can retrieve this from the &#39;Security Management&#39; section of the Alibaba Cloud
         * console.
         * 
         * @return builder
         * 
         */
        public Builder accessKey(String accessKey) {
            return accessKey(Output.of(accessKey));
        }

        /**
         * @param accountId The account ID for some service API operations. You can retrieve this from the &#39;Security Settings&#39; section of the
         * Alibaba Cloud console.
         * 
         * @return builder
         * 
         */
        public Builder accountId(@Nullable Output<String> accountId) {
            $.accountId = accountId;
            return this;
        }

        /**
         * @param accountId The account ID for some service API operations. You can retrieve this from the &#39;Security Settings&#39; section of the
         * Alibaba Cloud console.
         * 
         * @return builder
         * 
         */
        public Builder accountId(String accountId) {
            return accountId(Output.of(accountId));
        }

        public Builder assumeRole(@Nullable Output<ProviderAssumeRoleArgs> assumeRole) {
            $.assumeRole = assumeRole;
            return this;
        }

        public Builder assumeRole(ProviderAssumeRoleArgs assumeRole) {
            return assumeRole(Output.of(assumeRole));
        }

        /**
         * @param clientConnectTimeout The maximum timeout of the client connection server.
         * 
         * @return builder
         * 
         */
        public Builder clientConnectTimeout(@Nullable Output<Integer> clientConnectTimeout) {
            $.clientConnectTimeout = clientConnectTimeout;
            return this;
        }

        /**
         * @param clientConnectTimeout The maximum timeout of the client connection server.
         * 
         * @return builder
         * 
         */
        public Builder clientConnectTimeout(Integer clientConnectTimeout) {
            return clientConnectTimeout(Output.of(clientConnectTimeout));
        }

        /**
         * @param clientReadTimeout The maximum timeout of the client read request.
         * 
         * @return builder
         * 
         */
        public Builder clientReadTimeout(@Nullable Output<Integer> clientReadTimeout) {
            $.clientReadTimeout = clientReadTimeout;
            return this;
        }

        /**
         * @param clientReadTimeout The maximum timeout of the client read request.
         * 
         * @return builder
         * 
         */
        public Builder clientReadTimeout(Integer clientReadTimeout) {
            return clientReadTimeout(Output.of(clientReadTimeout));
        }

        /**
         * @param configurationSource Use this to mark a terraform configuration file source.
         * 
         * @return builder
         * 
         */
        public Builder configurationSource(@Nullable Output<String> configurationSource) {
            $.configurationSource = configurationSource;
            return this;
        }

        /**
         * @param configurationSource Use this to mark a terraform configuration file source.
         * 
         * @return builder
         * 
         */
        public Builder configurationSource(String configurationSource) {
            return configurationSource(Output.of(configurationSource));
        }

        /**
         * @param credentialsUri The URI of sidecar credentials service.
         * 
         * @return builder
         * 
         */
        public Builder credentialsUri(@Nullable Output<String> credentialsUri) {
            $.credentialsUri = credentialsUri;
            return this;
        }

        /**
         * @param credentialsUri The URI of sidecar credentials service.
         * 
         * @return builder
         * 
         */
        public Builder credentialsUri(String credentialsUri) {
            return credentialsUri(Output.of(credentialsUri));
        }

        /**
         * @param ecsRoleName The RAM Role Name attached on a ECS instance for API operations. You can retrieve this from the &#39;Access Control&#39; section
         * of the Alibaba Cloud console.
         * 
         * @return builder
         * 
         */
        public Builder ecsRoleName(@Nullable Output<String> ecsRoleName) {
            $.ecsRoleName = ecsRoleName;
            return this;
        }

        /**
         * @param ecsRoleName The RAM Role Name attached on a ECS instance for API operations. You can retrieve this from the &#39;Access Control&#39; section
         * of the Alibaba Cloud console.
         * 
         * @return builder
         * 
         */
        public Builder ecsRoleName(String ecsRoleName) {
            return ecsRoleName(Output.of(ecsRoleName));
        }

        public Builder endpoints(@Nullable Output<List<ProviderEndpointArgs>> endpoints) {
            $.endpoints = endpoints;
            return this;
        }

        public Builder endpoints(List<ProviderEndpointArgs> endpoints) {
            return endpoints(Output.of(endpoints));
        }

        public Builder endpoints(ProviderEndpointArgs... endpoints) {
            return endpoints(List.of(endpoints));
        }

        /**
         * @return builder
         * 
         * @deprecated
         * Field &#39;fc&#39; has been deprecated from provider version 1.28.0. New field &#39;fc&#39; which in nested endpoints instead.
         * 
         */
        @Deprecated /* Field 'fc' has been deprecated from provider version 1.28.0. New field 'fc' which in nested endpoints instead. */
        public Builder fc(@Nullable Output<String> fc) {
            $.fc = fc;
            return this;
        }

        /**
         * @return builder
         * 
         * @deprecated
         * Field &#39;fc&#39; has been deprecated from provider version 1.28.0. New field &#39;fc&#39; which in nested endpoints instead.
         * 
         */
        @Deprecated /* Field 'fc' has been deprecated from provider version 1.28.0. New field 'fc' which in nested endpoints instead. */
        public Builder fc(String fc) {
            return fc(Output.of(fc));
        }

        /**
         * @return builder
         * 
         * @deprecated
         * Field &#39;log_endpoint&#39; has been deprecated from provider version 1.28.0. New field &#39;log&#39; which in nested endpoints instead.
         * 
         */
        @Deprecated /* Field 'log_endpoint' has been deprecated from provider version 1.28.0. New field 'log' which in nested endpoints instead. */
        public Builder logEndpoint(@Nullable Output<String> logEndpoint) {
            $.logEndpoint = logEndpoint;
            return this;
        }

        /**
         * @return builder
         * 
         * @deprecated
         * Field &#39;log_endpoint&#39; has been deprecated from provider version 1.28.0. New field &#39;log&#39; which in nested endpoints instead.
         * 
         */
        @Deprecated /* Field 'log_endpoint' has been deprecated from provider version 1.28.0. New field 'log' which in nested endpoints instead. */
        public Builder logEndpoint(String logEndpoint) {
            return logEndpoint(Output.of(logEndpoint));
        }

        /**
         * @param maxRetryTimeout The maximum retry timeout of the request.
         * 
         * @return builder
         * 
         */
        public Builder maxRetryTimeout(@Nullable Output<Integer> maxRetryTimeout) {
            $.maxRetryTimeout = maxRetryTimeout;
            return this;
        }

        /**
         * @param maxRetryTimeout The maximum retry timeout of the request.
         * 
         * @return builder
         * 
         */
        public Builder maxRetryTimeout(Integer maxRetryTimeout) {
            return maxRetryTimeout(Output.of(maxRetryTimeout));
        }

        /**
         * @return builder
         * 
         * @deprecated
         * Field &#39;mns_endpoint&#39; has been deprecated from provider version 1.28.0. New field &#39;mns&#39; which in nested endpoints instead.
         * 
         */
        @Deprecated /* Field 'mns_endpoint' has been deprecated from provider version 1.28.0. New field 'mns' which in nested endpoints instead. */
        public Builder mnsEndpoint(@Nullable Output<String> mnsEndpoint) {
            $.mnsEndpoint = mnsEndpoint;
            return this;
        }

        /**
         * @return builder
         * 
         * @deprecated
         * Field &#39;mns_endpoint&#39; has been deprecated from provider version 1.28.0. New field &#39;mns&#39; which in nested endpoints instead.
         * 
         */
        @Deprecated /* Field 'mns_endpoint' has been deprecated from provider version 1.28.0. New field 'mns' which in nested endpoints instead. */
        public Builder mnsEndpoint(String mnsEndpoint) {
            return mnsEndpoint(Output.of(mnsEndpoint));
        }

        /**
         * @return builder
         * 
         * @deprecated
         * Field &#39;ots_instance_name&#39; has been deprecated from provider version 1.10.0. New field &#39;instance_name&#39; of resource &#39;alicloud_ots_table&#39; instead.
         * 
         */
        @Deprecated /* Field 'ots_instance_name' has been deprecated from provider version 1.10.0. New field 'instance_name' of resource 'alicloud_ots_table' instead. */
        public Builder otsInstanceName(@Nullable Output<String> otsInstanceName) {
            $.otsInstanceName = otsInstanceName;
            return this;
        }

        /**
         * @return builder
         * 
         * @deprecated
         * Field &#39;ots_instance_name&#39; has been deprecated from provider version 1.10.0. New field &#39;instance_name&#39; of resource &#39;alicloud_ots_table&#39; instead.
         * 
         */
        @Deprecated /* Field 'ots_instance_name' has been deprecated from provider version 1.10.0. New field 'instance_name' of resource 'alicloud_ots_table' instead. */
        public Builder otsInstanceName(String otsInstanceName) {
            return otsInstanceName(Output.of(otsInstanceName));
        }

        /**
         * @param profile The profile for API operations. If not set, the default profile created with `aliyun configure` will be used.
         * 
         * @return builder
         * 
         */
        public Builder profile(@Nullable Output<String> profile) {
            $.profile = profile;
            return this;
        }

        /**
         * @param profile The profile for API operations. If not set, the default profile created with `aliyun configure` will be used.
         * 
         * @return builder
         * 
         */
        public Builder profile(String profile) {
            return profile(Output.of(profile));
        }

        public Builder protocol(@Nullable Output<String> protocol) {
            $.protocol = protocol;
            return this;
        }

        public Builder protocol(String protocol) {
            return protocol(Output.of(protocol));
        }

        /**
         * @param region The region where Alibaba Cloud operations will take place. Examples are cn-beijing, cn-hangzhou, eu-central-1, etc.
         * 
         * @return builder
         * 
         */
        public Builder region(@Nullable Output<String> region) {
            $.region = region;
            return this;
        }

        /**
         * @param region The region where Alibaba Cloud operations will take place. Examples are cn-beijing, cn-hangzhou, eu-central-1, etc.
         * 
         * @return builder
         * 
         */
        public Builder region(String region) {
            return region(Output.of(region));
        }

        /**
         * @param secretKey The secret key for API operations. You can retrieve this from the &#39;Security Management&#39; section of the Alibaba Cloud
         * console.
         * 
         * @return builder
         * 
         */
        public Builder secretKey(@Nullable Output<String> secretKey) {
            $.secretKey = secretKey;
            return this;
        }

        /**
         * @param secretKey The secret key for API operations. You can retrieve this from the &#39;Security Management&#39; section of the Alibaba Cloud
         * console.
         * 
         * @return builder
         * 
         */
        public Builder secretKey(String secretKey) {
            return secretKey(Output.of(secretKey));
        }

        /**
         * @param secureTransport The security transport for the assume role invoking.
         * 
         * @return builder
         * 
         */
        public Builder secureTransport(@Nullable Output<String> secureTransport) {
            $.secureTransport = secureTransport;
            return this;
        }

        /**
         * @param secureTransport The security transport for the assume role invoking.
         * 
         * @return builder
         * 
         */
        public Builder secureTransport(String secureTransport) {
            return secureTransport(Output.of(secureTransport));
        }

        /**
         * @param securityToken security token. A security token is only required if you are using Security Token Service.
         * 
         * @return builder
         * 
         */
        public Builder securityToken(@Nullable Output<String> securityToken) {
            $.securityToken = securityToken;
            return this;
        }

        /**
         * @param securityToken security token. A security token is only required if you are using Security Token Service.
         * 
         * @return builder
         * 
         */
        public Builder securityToken(String securityToken) {
            return securityToken(Output.of(securityToken));
        }

        public Builder securityTransport(@Nullable Output<String> securityTransport) {
            $.securityTransport = securityTransport;
            return this;
        }

        public Builder securityTransport(String securityTransport) {
            return securityTransport(Output.of(securityTransport));
        }

        /**
         * @param sharedCredentialsFile The path to the shared credentials file. If not set this defaults to ~/.aliyun/config.json
         * 
         * @return builder
         * 
         */
        public Builder sharedCredentialsFile(@Nullable Output<String> sharedCredentialsFile) {
            $.sharedCredentialsFile = sharedCredentialsFile;
            return this;
        }

        /**
         * @param sharedCredentialsFile The path to the shared credentials file. If not set this defaults to ~/.aliyun/config.json
         * 
         * @return builder
         * 
         */
        public Builder sharedCredentialsFile(String sharedCredentialsFile) {
            return sharedCredentialsFile(Output.of(sharedCredentialsFile));
        }

        /**
         * @param skipRegionValidation Skip static validation of region ID. Used by users of alternative AlibabaCloud-like APIs or users w/ access to regions
         * that are not public (yet).
         * 
         * @return builder
         * 
         */
        public Builder skipRegionValidation(@Nullable Output<Boolean> skipRegionValidation) {
            $.skipRegionValidation = skipRegionValidation;
            return this;
        }

        /**
         * @param skipRegionValidation Skip static validation of region ID. Used by users of alternative AlibabaCloud-like APIs or users w/ access to regions
         * that are not public (yet).
         * 
         * @return builder
         * 
         */
        public Builder skipRegionValidation(Boolean skipRegionValidation) {
            return skipRegionValidation(Output.of(skipRegionValidation));
        }

        /**
         * @param sourceIp The source ip for the assume role invoking.
         * 
         * @return builder
         * 
         */
        public Builder sourceIp(@Nullable Output<String> sourceIp) {
            $.sourceIp = sourceIp;
            return this;
        }

        /**
         * @param sourceIp The source ip for the assume role invoking.
         * 
         * @return builder
         * 
         */
        public Builder sourceIp(String sourceIp) {
            return sourceIp(Output.of(sourceIp));
        }

        public ProviderArgs build() {
            $.ecsRoleName = Codegen.stringProp("ecsRoleName").output().arg($.ecsRoleName).env("ALICLOUD_ECS_ROLE_NAME").getNullable();
            $.profile = Codegen.stringProp("profile").output().arg($.profile).env("ALICLOUD_PROFILE").getNullable();
            $.region = Codegen.stringProp("region").output().arg($.region).env("ALICLOUD_REGION").getNullable();
            return $;
        }
    }

}
