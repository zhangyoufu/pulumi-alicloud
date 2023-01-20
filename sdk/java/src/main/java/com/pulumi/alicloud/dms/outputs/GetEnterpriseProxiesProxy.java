// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.dms.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetEnterpriseProxiesProxy {
    /**
     * @return The ID of the user who enabled the secure access proxy feature.
     * 
     */
    private String creatorId;
    /**
     * @return The nickname of the user who enabled the secure access proxy feature.
     * 
     */
    private String creatorName;
    /**
     * @return The port that was used by HTTPS clients to connect to the database instance.
     * 
     */
    private Integer httpsPort;
    /**
     * @return The ID of the Proxy.
     * 
     */
    private String id;
    /**
     * @return The ID of the database instance.
     * 
     */
    private String instanceId;
    /**
     * @return Indicates whether the internal endpoint is enabled. Default value: true.
     * 
     */
    private Boolean privateEnable;
    /**
     * @return The internal endpoint.
     * 
     */
    private String privateHost;
    /**
     * @return Database protocol connection port number.
     * 
     */
    private Integer protocolPort;
    /**
     * @return Database protocol type, for example, MYSQL.
     * 
     */
    private String protocolType;
    /**
     * @return The ID of the secure access proxy.
     * 
     */
    private String proxyId;
    /**
     * @return Indicates whether the public endpoint is enabled.
     * 
     */
    private Boolean publicEnable;
    /**
     * @return The public endpoint. A public endpoint is returned no matter whether the public endpoint is enabled or disabled. **Note:** When the public network address is in the **true** state, the returned public network address is a valid address with DNS resolution capability. When the public address is in the **false** state, the returned Public address is an invalid address without DNS resolution.
     * 
     */
    private String publicHost;

    private GetEnterpriseProxiesProxy() {}
    /**
     * @return The ID of the user who enabled the secure access proxy feature.
     * 
     */
    public String creatorId() {
        return this.creatorId;
    }
    /**
     * @return The nickname of the user who enabled the secure access proxy feature.
     * 
     */
    public String creatorName() {
        return this.creatorName;
    }
    /**
     * @return The port that was used by HTTPS clients to connect to the database instance.
     * 
     */
    public Integer httpsPort() {
        return this.httpsPort;
    }
    /**
     * @return The ID of the Proxy.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return The ID of the database instance.
     * 
     */
    public String instanceId() {
        return this.instanceId;
    }
    /**
     * @return Indicates whether the internal endpoint is enabled. Default value: true.
     * 
     */
    public Boolean privateEnable() {
        return this.privateEnable;
    }
    /**
     * @return The internal endpoint.
     * 
     */
    public String privateHost() {
        return this.privateHost;
    }
    /**
     * @return Database protocol connection port number.
     * 
     */
    public Integer protocolPort() {
        return this.protocolPort;
    }
    /**
     * @return Database protocol type, for example, MYSQL.
     * 
     */
    public String protocolType() {
        return this.protocolType;
    }
    /**
     * @return The ID of the secure access proxy.
     * 
     */
    public String proxyId() {
        return this.proxyId;
    }
    /**
     * @return Indicates whether the public endpoint is enabled.
     * 
     */
    public Boolean publicEnable() {
        return this.publicEnable;
    }
    /**
     * @return The public endpoint. A public endpoint is returned no matter whether the public endpoint is enabled or disabled. **Note:** When the public network address is in the **true** state, the returned public network address is a valid address with DNS resolution capability. When the public address is in the **false** state, the returned Public address is an invalid address without DNS resolution.
     * 
     */
    public String publicHost() {
        return this.publicHost;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetEnterpriseProxiesProxy defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String creatorId;
        private String creatorName;
        private Integer httpsPort;
        private String id;
        private String instanceId;
        private Boolean privateEnable;
        private String privateHost;
        private Integer protocolPort;
        private String protocolType;
        private String proxyId;
        private Boolean publicEnable;
        private String publicHost;
        public Builder() {}
        public Builder(GetEnterpriseProxiesProxy defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.creatorId = defaults.creatorId;
    	      this.creatorName = defaults.creatorName;
    	      this.httpsPort = defaults.httpsPort;
    	      this.id = defaults.id;
    	      this.instanceId = defaults.instanceId;
    	      this.privateEnable = defaults.privateEnable;
    	      this.privateHost = defaults.privateHost;
    	      this.protocolPort = defaults.protocolPort;
    	      this.protocolType = defaults.protocolType;
    	      this.proxyId = defaults.proxyId;
    	      this.publicEnable = defaults.publicEnable;
    	      this.publicHost = defaults.publicHost;
        }

        @CustomType.Setter
        public Builder creatorId(String creatorId) {
            this.creatorId = Objects.requireNonNull(creatorId);
            return this;
        }
        @CustomType.Setter
        public Builder creatorName(String creatorName) {
            this.creatorName = Objects.requireNonNull(creatorName);
            return this;
        }
        @CustomType.Setter
        public Builder httpsPort(Integer httpsPort) {
            this.httpsPort = Objects.requireNonNull(httpsPort);
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder instanceId(String instanceId) {
            this.instanceId = Objects.requireNonNull(instanceId);
            return this;
        }
        @CustomType.Setter
        public Builder privateEnable(Boolean privateEnable) {
            this.privateEnable = Objects.requireNonNull(privateEnable);
            return this;
        }
        @CustomType.Setter
        public Builder privateHost(String privateHost) {
            this.privateHost = Objects.requireNonNull(privateHost);
            return this;
        }
        @CustomType.Setter
        public Builder protocolPort(Integer protocolPort) {
            this.protocolPort = Objects.requireNonNull(protocolPort);
            return this;
        }
        @CustomType.Setter
        public Builder protocolType(String protocolType) {
            this.protocolType = Objects.requireNonNull(protocolType);
            return this;
        }
        @CustomType.Setter
        public Builder proxyId(String proxyId) {
            this.proxyId = Objects.requireNonNull(proxyId);
            return this;
        }
        @CustomType.Setter
        public Builder publicEnable(Boolean publicEnable) {
            this.publicEnable = Objects.requireNonNull(publicEnable);
            return this;
        }
        @CustomType.Setter
        public Builder publicHost(String publicHost) {
            this.publicHost = Objects.requireNonNull(publicHost);
            return this;
        }
        public GetEnterpriseProxiesProxy build() {
            final var o = new GetEnterpriseProxiesProxy();
            o.creatorId = creatorId;
            o.creatorName = creatorName;
            o.httpsPort = httpsPort;
            o.id = id;
            o.instanceId = instanceId;
            o.privateEnable = privateEnable;
            o.privateHost = privateHost;
            o.protocolPort = protocolPort;
            o.protocolType = protocolType;
            o.proxyId = proxyId;
            o.publicEnable = publicEnable;
            o.publicHost = publicHost;
            return o;
        }
    }
}
