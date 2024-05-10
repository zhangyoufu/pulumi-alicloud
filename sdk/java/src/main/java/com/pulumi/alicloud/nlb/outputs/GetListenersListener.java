// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.nlb.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;

@CustomType
public final class GetListenersListener {
    /**
     * @return ndicates whether Application-Layer Protocol Negotiation (ALPN) is enabled.
     * 
     */
    private Boolean alpnEnabled;
    /**
     * @return The ALPN policy.
     * 
     */
    private String alpnPolicy;
    /**
     * @return CA certificate list information. Currently, only one CA certificate can be added. **NOTE:** This parameter only takes effect for `TCPSSL` listeners.
     * 
     */
    private List<String> caCertificateIds;
    /**
     * @return Whether to start two-way authentication.
     * 
     */
    private Boolean caEnabled;
    /**
     * @return Server certificate list information. Currently, only one server certificate can be added. This parameter only takes effect for `TCPSSL` listeners.
     * 
     */
    private List<String> certificateIds;
    /**
     * @return The new connection speed limit for a network-based load balancing instance per second. Valid values: `0` ~ `1000000`. `0` indicates unlimited speed.
     * 
     */
    private Integer cps;
    /**
     * @return Full port listening end port. Valid values: `0` ~ `65535`. The value of the end port is less than the start port.
     * 
     */
    private String endPort;
    /**
     * @return The ID of the Nlb Listener.
     * 
     */
    private String id;
    /**
     * @return Connection idle timeout time. Unit: seconds. Valid values: `1` ~ `900`.
     * 
     */
    private Integer idleTimeout;
    /**
     * @return Custom listener name. The length is limited to 2 to 256 characters, supports Chinese and English letters, and can include numbers, commas (,), half-width periods (.), half-width semicolons (;), forward slashes (/), at({@literal @}), underscores (_), and dashes (-).
     * 
     */
    private String listenerDescription;
    /**
     * @return The ID of the listener.
     * 
     */
    private String listenerId;
    /**
     * @return Listening port. Valid values: `0` ~ `65535`. `0`: indicates that full port listening is used. When set to 0, you must configure `StartPort` and `EndPort`.
     * 
     */
    private Integer listenerPort;
    /**
     * @return The listening protocol. Valid values: `TCP`, `UDP`, or `TCPSSL`.
     * 
     */
    private String listenerProtocol;
    /**
     * @return The ID of the network-based server load balancer instance.
     * 
     */
    private String loadBalancerId;
    /**
     * @return The maximum segment size of the TCP message. Unit: Bytes. Valid values: `0` ~ `1500`. `0` indicates that the MSS value of the TCP message is not modified. only `TCP` and `TCPSSL` listeners support this field value.
     * 
     */
    private Integer mss;
    /**
     * @return Whether to enable the Proxy Protocol to carry the source address of the client to the backend server.
     * 
     */
    private Boolean proxyProtocolEnabled;
    /**
     * @return Indicates whether fine-grained monitoring is enabled.
     * 
     */
    private Boolean secSensorEnabled;
    /**
     * @return Security policy ID. Support system security policies and custom security policies. Valid values: `tls_cipher_policy_1_0`, `tls_cipher_policy_1_1`, `tls_cipher_policy_1_2`, `tls_cipher_policy_1_2_strict`, or `tls_cipher_policy_1_2_strict_with_1_3`. **Note:** This parameter only takes effect for `TCPSSL` listeners.
     * 
     */
    private String securityPolicyId;
    /**
     * @return The ID of the server group.
     * 
     */
    private String serverGroupId;
    /**
     * @return Full Port listens to the starting port. Valid values: `0` ~ `65535`.
     * 
     */
    private String startPort;
    /**
     * @return The status of the resource.
     * 
     */
    private String status;

    private GetListenersListener() {}
    /**
     * @return ndicates whether Application-Layer Protocol Negotiation (ALPN) is enabled.
     * 
     */
    public Boolean alpnEnabled() {
        return this.alpnEnabled;
    }
    /**
     * @return The ALPN policy.
     * 
     */
    public String alpnPolicy() {
        return this.alpnPolicy;
    }
    /**
     * @return CA certificate list information. Currently, only one CA certificate can be added. **NOTE:** This parameter only takes effect for `TCPSSL` listeners.
     * 
     */
    public List<String> caCertificateIds() {
        return this.caCertificateIds;
    }
    /**
     * @return Whether to start two-way authentication.
     * 
     */
    public Boolean caEnabled() {
        return this.caEnabled;
    }
    /**
     * @return Server certificate list information. Currently, only one server certificate can be added. This parameter only takes effect for `TCPSSL` listeners.
     * 
     */
    public List<String> certificateIds() {
        return this.certificateIds;
    }
    /**
     * @return The new connection speed limit for a network-based load balancing instance per second. Valid values: `0` ~ `1000000`. `0` indicates unlimited speed.
     * 
     */
    public Integer cps() {
        return this.cps;
    }
    /**
     * @return Full port listening end port. Valid values: `0` ~ `65535`. The value of the end port is less than the start port.
     * 
     */
    public String endPort() {
        return this.endPort;
    }
    /**
     * @return The ID of the Nlb Listener.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return Connection idle timeout time. Unit: seconds. Valid values: `1` ~ `900`.
     * 
     */
    public Integer idleTimeout() {
        return this.idleTimeout;
    }
    /**
     * @return Custom listener name. The length is limited to 2 to 256 characters, supports Chinese and English letters, and can include numbers, commas (,), half-width periods (.), half-width semicolons (;), forward slashes (/), at({@literal @}), underscores (_), and dashes (-).
     * 
     */
    public String listenerDescription() {
        return this.listenerDescription;
    }
    /**
     * @return The ID of the listener.
     * 
     */
    public String listenerId() {
        return this.listenerId;
    }
    /**
     * @return Listening port. Valid values: `0` ~ `65535`. `0`: indicates that full port listening is used. When set to 0, you must configure `StartPort` and `EndPort`.
     * 
     */
    public Integer listenerPort() {
        return this.listenerPort;
    }
    /**
     * @return The listening protocol. Valid values: `TCP`, `UDP`, or `TCPSSL`.
     * 
     */
    public String listenerProtocol() {
        return this.listenerProtocol;
    }
    /**
     * @return The ID of the network-based server load balancer instance.
     * 
     */
    public String loadBalancerId() {
        return this.loadBalancerId;
    }
    /**
     * @return The maximum segment size of the TCP message. Unit: Bytes. Valid values: `0` ~ `1500`. `0` indicates that the MSS value of the TCP message is not modified. only `TCP` and `TCPSSL` listeners support this field value.
     * 
     */
    public Integer mss() {
        return this.mss;
    }
    /**
     * @return Whether to enable the Proxy Protocol to carry the source address of the client to the backend server.
     * 
     */
    public Boolean proxyProtocolEnabled() {
        return this.proxyProtocolEnabled;
    }
    /**
     * @return Indicates whether fine-grained monitoring is enabled.
     * 
     */
    public Boolean secSensorEnabled() {
        return this.secSensorEnabled;
    }
    /**
     * @return Security policy ID. Support system security policies and custom security policies. Valid values: `tls_cipher_policy_1_0`, `tls_cipher_policy_1_1`, `tls_cipher_policy_1_2`, `tls_cipher_policy_1_2_strict`, or `tls_cipher_policy_1_2_strict_with_1_3`. **Note:** This parameter only takes effect for `TCPSSL` listeners.
     * 
     */
    public String securityPolicyId() {
        return this.securityPolicyId;
    }
    /**
     * @return The ID of the server group.
     * 
     */
    public String serverGroupId() {
        return this.serverGroupId;
    }
    /**
     * @return Full Port listens to the starting port. Valid values: `0` ~ `65535`.
     * 
     */
    public String startPort() {
        return this.startPort;
    }
    /**
     * @return The status of the resource.
     * 
     */
    public String status() {
        return this.status;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetListenersListener defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private Boolean alpnEnabled;
        private String alpnPolicy;
        private List<String> caCertificateIds;
        private Boolean caEnabled;
        private List<String> certificateIds;
        private Integer cps;
        private String endPort;
        private String id;
        private Integer idleTimeout;
        private String listenerDescription;
        private String listenerId;
        private Integer listenerPort;
        private String listenerProtocol;
        private String loadBalancerId;
        private Integer mss;
        private Boolean proxyProtocolEnabled;
        private Boolean secSensorEnabled;
        private String securityPolicyId;
        private String serverGroupId;
        private String startPort;
        private String status;
        public Builder() {}
        public Builder(GetListenersListener defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.alpnEnabled = defaults.alpnEnabled;
    	      this.alpnPolicy = defaults.alpnPolicy;
    	      this.caCertificateIds = defaults.caCertificateIds;
    	      this.caEnabled = defaults.caEnabled;
    	      this.certificateIds = defaults.certificateIds;
    	      this.cps = defaults.cps;
    	      this.endPort = defaults.endPort;
    	      this.id = defaults.id;
    	      this.idleTimeout = defaults.idleTimeout;
    	      this.listenerDescription = defaults.listenerDescription;
    	      this.listenerId = defaults.listenerId;
    	      this.listenerPort = defaults.listenerPort;
    	      this.listenerProtocol = defaults.listenerProtocol;
    	      this.loadBalancerId = defaults.loadBalancerId;
    	      this.mss = defaults.mss;
    	      this.proxyProtocolEnabled = defaults.proxyProtocolEnabled;
    	      this.secSensorEnabled = defaults.secSensorEnabled;
    	      this.securityPolicyId = defaults.securityPolicyId;
    	      this.serverGroupId = defaults.serverGroupId;
    	      this.startPort = defaults.startPort;
    	      this.status = defaults.status;
        }

        @CustomType.Setter
        public Builder alpnEnabled(Boolean alpnEnabled) {
            if (alpnEnabled == null) {
              throw new MissingRequiredPropertyException("GetListenersListener", "alpnEnabled");
            }
            this.alpnEnabled = alpnEnabled;
            return this;
        }
        @CustomType.Setter
        public Builder alpnPolicy(String alpnPolicy) {
            if (alpnPolicy == null) {
              throw new MissingRequiredPropertyException("GetListenersListener", "alpnPolicy");
            }
            this.alpnPolicy = alpnPolicy;
            return this;
        }
        @CustomType.Setter
        public Builder caCertificateIds(List<String> caCertificateIds) {
            if (caCertificateIds == null) {
              throw new MissingRequiredPropertyException("GetListenersListener", "caCertificateIds");
            }
            this.caCertificateIds = caCertificateIds;
            return this;
        }
        public Builder caCertificateIds(String... caCertificateIds) {
            return caCertificateIds(List.of(caCertificateIds));
        }
        @CustomType.Setter
        public Builder caEnabled(Boolean caEnabled) {
            if (caEnabled == null) {
              throw new MissingRequiredPropertyException("GetListenersListener", "caEnabled");
            }
            this.caEnabled = caEnabled;
            return this;
        }
        @CustomType.Setter
        public Builder certificateIds(List<String> certificateIds) {
            if (certificateIds == null) {
              throw new MissingRequiredPropertyException("GetListenersListener", "certificateIds");
            }
            this.certificateIds = certificateIds;
            return this;
        }
        public Builder certificateIds(String... certificateIds) {
            return certificateIds(List.of(certificateIds));
        }
        @CustomType.Setter
        public Builder cps(Integer cps) {
            if (cps == null) {
              throw new MissingRequiredPropertyException("GetListenersListener", "cps");
            }
            this.cps = cps;
            return this;
        }
        @CustomType.Setter
        public Builder endPort(String endPort) {
            if (endPort == null) {
              throw new MissingRequiredPropertyException("GetListenersListener", "endPort");
            }
            this.endPort = endPort;
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            if (id == null) {
              throw new MissingRequiredPropertyException("GetListenersListener", "id");
            }
            this.id = id;
            return this;
        }
        @CustomType.Setter
        public Builder idleTimeout(Integer idleTimeout) {
            if (idleTimeout == null) {
              throw new MissingRequiredPropertyException("GetListenersListener", "idleTimeout");
            }
            this.idleTimeout = idleTimeout;
            return this;
        }
        @CustomType.Setter
        public Builder listenerDescription(String listenerDescription) {
            if (listenerDescription == null) {
              throw new MissingRequiredPropertyException("GetListenersListener", "listenerDescription");
            }
            this.listenerDescription = listenerDescription;
            return this;
        }
        @CustomType.Setter
        public Builder listenerId(String listenerId) {
            if (listenerId == null) {
              throw new MissingRequiredPropertyException("GetListenersListener", "listenerId");
            }
            this.listenerId = listenerId;
            return this;
        }
        @CustomType.Setter
        public Builder listenerPort(Integer listenerPort) {
            if (listenerPort == null) {
              throw new MissingRequiredPropertyException("GetListenersListener", "listenerPort");
            }
            this.listenerPort = listenerPort;
            return this;
        }
        @CustomType.Setter
        public Builder listenerProtocol(String listenerProtocol) {
            if (listenerProtocol == null) {
              throw new MissingRequiredPropertyException("GetListenersListener", "listenerProtocol");
            }
            this.listenerProtocol = listenerProtocol;
            return this;
        }
        @CustomType.Setter
        public Builder loadBalancerId(String loadBalancerId) {
            if (loadBalancerId == null) {
              throw new MissingRequiredPropertyException("GetListenersListener", "loadBalancerId");
            }
            this.loadBalancerId = loadBalancerId;
            return this;
        }
        @CustomType.Setter
        public Builder mss(Integer mss) {
            if (mss == null) {
              throw new MissingRequiredPropertyException("GetListenersListener", "mss");
            }
            this.mss = mss;
            return this;
        }
        @CustomType.Setter
        public Builder proxyProtocolEnabled(Boolean proxyProtocolEnabled) {
            if (proxyProtocolEnabled == null) {
              throw new MissingRequiredPropertyException("GetListenersListener", "proxyProtocolEnabled");
            }
            this.proxyProtocolEnabled = proxyProtocolEnabled;
            return this;
        }
        @CustomType.Setter
        public Builder secSensorEnabled(Boolean secSensorEnabled) {
            if (secSensorEnabled == null) {
              throw new MissingRequiredPropertyException("GetListenersListener", "secSensorEnabled");
            }
            this.secSensorEnabled = secSensorEnabled;
            return this;
        }
        @CustomType.Setter
        public Builder securityPolicyId(String securityPolicyId) {
            if (securityPolicyId == null) {
              throw new MissingRequiredPropertyException("GetListenersListener", "securityPolicyId");
            }
            this.securityPolicyId = securityPolicyId;
            return this;
        }
        @CustomType.Setter
        public Builder serverGroupId(String serverGroupId) {
            if (serverGroupId == null) {
              throw new MissingRequiredPropertyException("GetListenersListener", "serverGroupId");
            }
            this.serverGroupId = serverGroupId;
            return this;
        }
        @CustomType.Setter
        public Builder startPort(String startPort) {
            if (startPort == null) {
              throw new MissingRequiredPropertyException("GetListenersListener", "startPort");
            }
            this.startPort = startPort;
            return this;
        }
        @CustomType.Setter
        public Builder status(String status) {
            if (status == null) {
              throw new MissingRequiredPropertyException("GetListenersListener", "status");
            }
            this.status = status;
            return this;
        }
        public GetListenersListener build() {
            final var _resultValue = new GetListenersListener();
            _resultValue.alpnEnabled = alpnEnabled;
            _resultValue.alpnPolicy = alpnPolicy;
            _resultValue.caCertificateIds = caCertificateIds;
            _resultValue.caEnabled = caEnabled;
            _resultValue.certificateIds = certificateIds;
            _resultValue.cps = cps;
            _resultValue.endPort = endPort;
            _resultValue.id = id;
            _resultValue.idleTimeout = idleTimeout;
            _resultValue.listenerDescription = listenerDescription;
            _resultValue.listenerId = listenerId;
            _resultValue.listenerPort = listenerPort;
            _resultValue.listenerProtocol = listenerProtocol;
            _resultValue.loadBalancerId = loadBalancerId;
            _resultValue.mss = mss;
            _resultValue.proxyProtocolEnabled = proxyProtocolEnabled;
            _resultValue.secSensorEnabled = secSensorEnabled;
            _resultValue.securityPolicyId = securityPolicyId;
            _resultValue.serverGroupId = serverGroupId;
            _resultValue.startPort = startPort;
            _resultValue.status = status;
            return _resultValue;
        }
    }
}
