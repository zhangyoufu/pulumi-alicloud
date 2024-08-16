// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.dcdn.inputs;

import com.pulumi.alicloud.dcdn.inputs.DomainSourceArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class DomainState extends com.pulumi.resources.ResourceArgs {

    public static final DomainState Empty = new DomainState();

    /**
     * The certificate ID. This parameter is required and valid only when `CertType` is set to `cas`. If you specify this parameter, an existing certificate is used.
     * 
     */
    @Import(name="certId")
    private @Nullable Output<String> certId;

    /**
     * @return The certificate ID. This parameter is required and valid only when `CertType` is set to `cas`. If you specify this parameter, an existing certificate is used.
     * 
     */
    public Optional<Output<String>> certId() {
        return Optional.ofNullable(this.certId);
    }

    /**
     * The name of the new certificate. You can specify only one certificate name. This parameter is optional and valid only when `CertType` is set to `upload`.
     * 
     */
    @Import(name="certName")
    private @Nullable Output<String> certName;

    /**
     * @return The name of the new certificate. You can specify only one certificate name. This parameter is optional and valid only when `CertType` is set to `upload`.
     * 
     */
    public Optional<Output<String>> certName() {
        return Optional.ofNullable(this.certName);
    }

    /**
     * The region of the SSL certificate. This parameter takes effect only when `CertType` is set to `cas`. Default value: **cn-hangzhou**. Valid values: **cn-hangzhou** and **ap-southeast-1**.
     * 
     */
    @Import(name="certRegion")
    private @Nullable Output<String> certRegion;

    /**
     * @return The region of the SSL certificate. This parameter takes effect only when `CertType` is set to `cas`. Default value: **cn-hangzhou**. Valid values: **cn-hangzhou** and **ap-southeast-1**.
     * 
     */
    public Optional<Output<String>> certRegion() {
        return Optional.ofNullable(this.certRegion);
    }

    /**
     * The certificate type.
     * 
     */
    @Import(name="certType")
    private @Nullable Output<String> certType;

    /**
     * @return The certificate type.
     * 
     */
    public Optional<Output<String>> certType() {
        return Optional.ofNullable(this.certType);
    }

    /**
     * The URL that is used for health checks.
     * 
     */
    @Import(name="checkUrl")
    private @Nullable Output<String> checkUrl;

    /**
     * @return The URL that is used for health checks.
     * 
     */
    public Optional<Output<String>> checkUrl() {
        return Optional.ofNullable(this.checkUrl);
    }

    /**
     * The CNAME domain name corresponding to the accelerated domain name.
     * 
     */
    @Import(name="cname")
    private @Nullable Output<String> cname;

    /**
     * @return The CNAME domain name corresponding to the accelerated domain name.
     * 
     */
    public Optional<Output<String>> cname() {
        return Optional.ofNullable(this.cname);
    }

    /**
     * The time when the accelerated domain name was created.
     * 
     */
    @Import(name="createTime")
    private @Nullable Output<String> createTime;

    /**
     * @return The time when the accelerated domain name was created.
     * 
     */
    public Optional<Output<String>> createTime() {
        return Optional.ofNullable(this.createTime);
    }

    /**
     * The accelerated domain name. You can specify multiple domain names and separate them with commas (,). You can specify up to 500 domain names in each request. The query results of multiple domain names are aggregated. If you do not specify this parameter, data of all accelerated domain names under your account is queried.
     * 
     */
    @Import(name="domainName")
    private @Nullable Output<String> domainName;

    /**
     * @return The accelerated domain name. You can specify multiple domain names and separate them with commas (,). You can specify up to 500 domain names in each request. The query results of multiple domain names are aggregated. If you do not specify this parameter, data of all accelerated domain names under your account is queried.
     * 
     */
    public Optional<Output<String>> domainName() {
        return Optional.ofNullable(this.domainName);
    }

    /**
     * Specifies whether the certificate is issued in canary releases. If you set this parameter to `staging`, the certificate is issued in canary releases. If you do not specify this parameter or set this parameter to other values, the certificate is officially issued.
     * 
     */
    @Import(name="env")
    private @Nullable Output<String> env;

    /**
     * @return Specifies whether the certificate is issued in canary releases. If you set this parameter to `staging`, the certificate is issued in canary releases. If you do not specify this parameter or set this parameter to other values, the certificate is officially issued.
     * 
     */
    public Optional<Output<String>> env() {
        return Optional.ofNullable(this.env);
    }

    /**
     * Computing service type. Valid values:
     * 
     */
    @Import(name="functionType")
    private @Nullable Output<String> functionType;

    /**
     * @return Computing service type. Valid values:
     * 
     */
    public Optional<Output<String>> functionType() {
        return Optional.ofNullable(this.functionType);
    }

    /**
     * The ID of the resource group. If you do not specify a value for this parameter, the system automatically assigns the ID of the default resource group.
     * 
     */
    @Import(name="resourceGroupId")
    private @Nullable Output<String> resourceGroupId;

    /**
     * @return The ID of the resource group. If you do not specify a value for this parameter, the system automatically assigns the ID of the default resource group.
     * 
     */
    public Optional<Output<String>> resourceGroupId() {
        return Optional.ofNullable(this.resourceGroupId);
    }

    /**
     * The Acceleration scen. Supported:
     * 
     */
    @Import(name="scene")
    private @Nullable Output<String> scene;

    /**
     * @return The Acceleration scen. Supported:
     * 
     */
    public Optional<Output<String>> scene() {
        return Optional.ofNullable(this.scene);
    }

    /**
     * The region where the acceleration service is deployed. Valid values:
     * 
     */
    @Import(name="scope")
    private @Nullable Output<String> scope;

    /**
     * @return The region where the acceleration service is deployed. Valid values:
     * 
     */
    public Optional<Output<String>> scope() {
        return Optional.ofNullable(this.scope);
    }

    /**
     * Source  See `sources` below.
     * 
     */
    @Import(name="sources")
    private @Nullable Output<List<DomainSourceArgs>> sources;

    /**
     * @return Source  See `sources` below.
     * 
     */
    public Optional<Output<List<DomainSourceArgs>>> sources() {
        return Optional.ofNullable(this.sources);
    }

    /**
     * The private key. Specify the private key only if you want to enable the SSL certificate.
     * 
     */
    @Import(name="sslPri")
    private @Nullable Output<String> sslPri;

    /**
     * @return The private key. Specify the private key only if you want to enable the SSL certificate.
     * 
     */
    public Optional<Output<String>> sslPri() {
        return Optional.ofNullable(this.sslPri);
    }

    /**
     * Specifies whether to enable the SSL certificate. Valid values:
     * 
     */
    @Import(name="sslProtocol")
    private @Nullable Output<String> sslProtocol;

    /**
     * @return Specifies whether to enable the SSL certificate. Valid values:
     * 
     */
    public Optional<Output<String>> sslProtocol() {
        return Optional.ofNullable(this.sslProtocol);
    }

    /**
     * The content of the SSL certificate. Specify the content of the SSL certificate only if you want to enable the SSL certificate.
     * 
     */
    @Import(name="sslPub")
    private @Nullable Output<String> sslPub;

    /**
     * @return The content of the SSL certificate. Specify the content of the SSL certificate only if you want to enable the SSL certificate.
     * 
     */
    public Optional<Output<String>> sslPub() {
        return Optional.ofNullable(this.sslPub);
    }

    /**
     * The status of the domain name. Valid values:
     * 
     */
    @Import(name="status")
    private @Nullable Output<String> status;

    /**
     * @return The status of the domain name. Valid values:
     * 
     */
    public Optional<Output<String>> status() {
        return Optional.ofNullable(this.status);
    }

    /**
     * The tag of the resource
     * 
     */
    @Import(name="tags")
    private @Nullable Output<Map<String,String>> tags;

    /**
     * @return The tag of the resource
     * 
     */
    public Optional<Output<Map<String,String>>> tags() {
        return Optional.ofNullable(this.tags);
    }

    /**
     * The top-level domain.
     * 
     */
    @Import(name="topLevelDomain")
    private @Nullable Output<String> topLevelDomain;

    /**
     * @return The top-level domain.
     * 
     */
    public Optional<Output<String>> topLevelDomain() {
        return Optional.ofNullable(this.topLevelDomain);
    }

    private DomainState() {}

    private DomainState(DomainState $) {
        this.certId = $.certId;
        this.certName = $.certName;
        this.certRegion = $.certRegion;
        this.certType = $.certType;
        this.checkUrl = $.checkUrl;
        this.cname = $.cname;
        this.createTime = $.createTime;
        this.domainName = $.domainName;
        this.env = $.env;
        this.functionType = $.functionType;
        this.resourceGroupId = $.resourceGroupId;
        this.scene = $.scene;
        this.scope = $.scope;
        this.sources = $.sources;
        this.sslPri = $.sslPri;
        this.sslProtocol = $.sslProtocol;
        this.sslPub = $.sslPub;
        this.status = $.status;
        this.tags = $.tags;
        this.topLevelDomain = $.topLevelDomain;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(DomainState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private DomainState $;

        public Builder() {
            $ = new DomainState();
        }

        public Builder(DomainState defaults) {
            $ = new DomainState(Objects.requireNonNull(defaults));
        }

        /**
         * @param certId The certificate ID. This parameter is required and valid only when `CertType` is set to `cas`. If you specify this parameter, an existing certificate is used.
         * 
         * @return builder
         * 
         */
        public Builder certId(@Nullable Output<String> certId) {
            $.certId = certId;
            return this;
        }

        /**
         * @param certId The certificate ID. This parameter is required and valid only when `CertType` is set to `cas`. If you specify this parameter, an existing certificate is used.
         * 
         * @return builder
         * 
         */
        public Builder certId(String certId) {
            return certId(Output.of(certId));
        }

        /**
         * @param certName The name of the new certificate. You can specify only one certificate name. This parameter is optional and valid only when `CertType` is set to `upload`.
         * 
         * @return builder
         * 
         */
        public Builder certName(@Nullable Output<String> certName) {
            $.certName = certName;
            return this;
        }

        /**
         * @param certName The name of the new certificate. You can specify only one certificate name. This parameter is optional and valid only when `CertType` is set to `upload`.
         * 
         * @return builder
         * 
         */
        public Builder certName(String certName) {
            return certName(Output.of(certName));
        }

        /**
         * @param certRegion The region of the SSL certificate. This parameter takes effect only when `CertType` is set to `cas`. Default value: **cn-hangzhou**. Valid values: **cn-hangzhou** and **ap-southeast-1**.
         * 
         * @return builder
         * 
         */
        public Builder certRegion(@Nullable Output<String> certRegion) {
            $.certRegion = certRegion;
            return this;
        }

        /**
         * @param certRegion The region of the SSL certificate. This parameter takes effect only when `CertType` is set to `cas`. Default value: **cn-hangzhou**. Valid values: **cn-hangzhou** and **ap-southeast-1**.
         * 
         * @return builder
         * 
         */
        public Builder certRegion(String certRegion) {
            return certRegion(Output.of(certRegion));
        }

        /**
         * @param certType The certificate type.
         * 
         * @return builder
         * 
         */
        public Builder certType(@Nullable Output<String> certType) {
            $.certType = certType;
            return this;
        }

        /**
         * @param certType The certificate type.
         * 
         * @return builder
         * 
         */
        public Builder certType(String certType) {
            return certType(Output.of(certType));
        }

        /**
         * @param checkUrl The URL that is used for health checks.
         * 
         * @return builder
         * 
         */
        public Builder checkUrl(@Nullable Output<String> checkUrl) {
            $.checkUrl = checkUrl;
            return this;
        }

        /**
         * @param checkUrl The URL that is used for health checks.
         * 
         * @return builder
         * 
         */
        public Builder checkUrl(String checkUrl) {
            return checkUrl(Output.of(checkUrl));
        }

        /**
         * @param cname The CNAME domain name corresponding to the accelerated domain name.
         * 
         * @return builder
         * 
         */
        public Builder cname(@Nullable Output<String> cname) {
            $.cname = cname;
            return this;
        }

        /**
         * @param cname The CNAME domain name corresponding to the accelerated domain name.
         * 
         * @return builder
         * 
         */
        public Builder cname(String cname) {
            return cname(Output.of(cname));
        }

        /**
         * @param createTime The time when the accelerated domain name was created.
         * 
         * @return builder
         * 
         */
        public Builder createTime(@Nullable Output<String> createTime) {
            $.createTime = createTime;
            return this;
        }

        /**
         * @param createTime The time when the accelerated domain name was created.
         * 
         * @return builder
         * 
         */
        public Builder createTime(String createTime) {
            return createTime(Output.of(createTime));
        }

        /**
         * @param domainName The accelerated domain name. You can specify multiple domain names and separate them with commas (,). You can specify up to 500 domain names in each request. The query results of multiple domain names are aggregated. If you do not specify this parameter, data of all accelerated domain names under your account is queried.
         * 
         * @return builder
         * 
         */
        public Builder domainName(@Nullable Output<String> domainName) {
            $.domainName = domainName;
            return this;
        }

        /**
         * @param domainName The accelerated domain name. You can specify multiple domain names and separate them with commas (,). You can specify up to 500 domain names in each request. The query results of multiple domain names are aggregated. If you do not specify this parameter, data of all accelerated domain names under your account is queried.
         * 
         * @return builder
         * 
         */
        public Builder domainName(String domainName) {
            return domainName(Output.of(domainName));
        }

        /**
         * @param env Specifies whether the certificate is issued in canary releases. If you set this parameter to `staging`, the certificate is issued in canary releases. If you do not specify this parameter or set this parameter to other values, the certificate is officially issued.
         * 
         * @return builder
         * 
         */
        public Builder env(@Nullable Output<String> env) {
            $.env = env;
            return this;
        }

        /**
         * @param env Specifies whether the certificate is issued in canary releases. If you set this parameter to `staging`, the certificate is issued in canary releases. If you do not specify this parameter or set this parameter to other values, the certificate is officially issued.
         * 
         * @return builder
         * 
         */
        public Builder env(String env) {
            return env(Output.of(env));
        }

        /**
         * @param functionType Computing service type. Valid values:
         * 
         * @return builder
         * 
         */
        public Builder functionType(@Nullable Output<String> functionType) {
            $.functionType = functionType;
            return this;
        }

        /**
         * @param functionType Computing service type. Valid values:
         * 
         * @return builder
         * 
         */
        public Builder functionType(String functionType) {
            return functionType(Output.of(functionType));
        }

        /**
         * @param resourceGroupId The ID of the resource group. If you do not specify a value for this parameter, the system automatically assigns the ID of the default resource group.
         * 
         * @return builder
         * 
         */
        public Builder resourceGroupId(@Nullable Output<String> resourceGroupId) {
            $.resourceGroupId = resourceGroupId;
            return this;
        }

        /**
         * @param resourceGroupId The ID of the resource group. If you do not specify a value for this parameter, the system automatically assigns the ID of the default resource group.
         * 
         * @return builder
         * 
         */
        public Builder resourceGroupId(String resourceGroupId) {
            return resourceGroupId(Output.of(resourceGroupId));
        }

        /**
         * @param scene The Acceleration scen. Supported:
         * 
         * @return builder
         * 
         */
        public Builder scene(@Nullable Output<String> scene) {
            $.scene = scene;
            return this;
        }

        /**
         * @param scene The Acceleration scen. Supported:
         * 
         * @return builder
         * 
         */
        public Builder scene(String scene) {
            return scene(Output.of(scene));
        }

        /**
         * @param scope The region where the acceleration service is deployed. Valid values:
         * 
         * @return builder
         * 
         */
        public Builder scope(@Nullable Output<String> scope) {
            $.scope = scope;
            return this;
        }

        /**
         * @param scope The region where the acceleration service is deployed. Valid values:
         * 
         * @return builder
         * 
         */
        public Builder scope(String scope) {
            return scope(Output.of(scope));
        }

        /**
         * @param sources Source  See `sources` below.
         * 
         * @return builder
         * 
         */
        public Builder sources(@Nullable Output<List<DomainSourceArgs>> sources) {
            $.sources = sources;
            return this;
        }

        /**
         * @param sources Source  See `sources` below.
         * 
         * @return builder
         * 
         */
        public Builder sources(List<DomainSourceArgs> sources) {
            return sources(Output.of(sources));
        }

        /**
         * @param sources Source  See `sources` below.
         * 
         * @return builder
         * 
         */
        public Builder sources(DomainSourceArgs... sources) {
            return sources(List.of(sources));
        }

        /**
         * @param sslPri The private key. Specify the private key only if you want to enable the SSL certificate.
         * 
         * @return builder
         * 
         */
        public Builder sslPri(@Nullable Output<String> sslPri) {
            $.sslPri = sslPri;
            return this;
        }

        /**
         * @param sslPri The private key. Specify the private key only if you want to enable the SSL certificate.
         * 
         * @return builder
         * 
         */
        public Builder sslPri(String sslPri) {
            return sslPri(Output.of(sslPri));
        }

        /**
         * @param sslProtocol Specifies whether to enable the SSL certificate. Valid values:
         * 
         * @return builder
         * 
         */
        public Builder sslProtocol(@Nullable Output<String> sslProtocol) {
            $.sslProtocol = sslProtocol;
            return this;
        }

        /**
         * @param sslProtocol Specifies whether to enable the SSL certificate. Valid values:
         * 
         * @return builder
         * 
         */
        public Builder sslProtocol(String sslProtocol) {
            return sslProtocol(Output.of(sslProtocol));
        }

        /**
         * @param sslPub The content of the SSL certificate. Specify the content of the SSL certificate only if you want to enable the SSL certificate.
         * 
         * @return builder
         * 
         */
        public Builder sslPub(@Nullable Output<String> sslPub) {
            $.sslPub = sslPub;
            return this;
        }

        /**
         * @param sslPub The content of the SSL certificate. Specify the content of the SSL certificate only if you want to enable the SSL certificate.
         * 
         * @return builder
         * 
         */
        public Builder sslPub(String sslPub) {
            return sslPub(Output.of(sslPub));
        }

        /**
         * @param status The status of the domain name. Valid values:
         * 
         * @return builder
         * 
         */
        public Builder status(@Nullable Output<String> status) {
            $.status = status;
            return this;
        }

        /**
         * @param status The status of the domain name. Valid values:
         * 
         * @return builder
         * 
         */
        public Builder status(String status) {
            return status(Output.of(status));
        }

        /**
         * @param tags The tag of the resource
         * 
         * @return builder
         * 
         */
        public Builder tags(@Nullable Output<Map<String,String>> tags) {
            $.tags = tags;
            return this;
        }

        /**
         * @param tags The tag of the resource
         * 
         * @return builder
         * 
         */
        public Builder tags(Map<String,String> tags) {
            return tags(Output.of(tags));
        }

        /**
         * @param topLevelDomain The top-level domain.
         * 
         * @return builder
         * 
         */
        public Builder topLevelDomain(@Nullable Output<String> topLevelDomain) {
            $.topLevelDomain = topLevelDomain;
            return this;
        }

        /**
         * @param topLevelDomain The top-level domain.
         * 
         * @return builder
         * 
         */
        public Builder topLevelDomain(String topLevelDomain) {
            return topLevelDomain(Output.of(topLevelDomain));
        }

        public DomainState build() {
            return $;
        }
    }

}
