// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.nlb;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ListenerArgs extends com.pulumi.resources.ResourceArgs {

    public static final ListenerArgs Empty = new ListenerArgs();

    /**
     * Specifies whether to enable Application-Layer Protocol Negotiation (ALPN).
     * 
     */
    @Import(name="alpnEnabled")
    private @Nullable Output<Boolean> alpnEnabled;

    /**
     * @return Specifies whether to enable Application-Layer Protocol Negotiation (ALPN).
     * 
     */
    public Optional<Output<Boolean>> alpnEnabled() {
        return Optional.ofNullable(this.alpnEnabled);
    }

    /**
     * The ALPN policy.
     * 
     */
    @Import(name="alpnPolicy")
    private @Nullable Output<String> alpnPolicy;

    /**
     * @return The ALPN policy.
     * 
     */
    public Optional<Output<String>> alpnPolicy() {
        return Optional.ofNullable(this.alpnPolicy);
    }

    /**
     * The list of certificate authority (CA) certificates. This parameter takes effect only for listeners that use SSL over TCP. **Note:** Only one CA certificate is supported.
     * 
     */
    @Import(name="caCertificateIds")
    private @Nullable Output<List<String>> caCertificateIds;

    /**
     * @return The list of certificate authority (CA) certificates. This parameter takes effect only for listeners that use SSL over TCP. **Note:** Only one CA certificate is supported.
     * 
     */
    public Optional<Output<List<String>>> caCertificateIds() {
        return Optional.ofNullable(this.caCertificateIds);
    }

    /**
     * Specifies whether to enable mutual authentication.
     * 
     */
    @Import(name="caEnabled")
    private @Nullable Output<Boolean> caEnabled;

    /**
     * @return Specifies whether to enable mutual authentication.
     * 
     */
    public Optional<Output<Boolean>> caEnabled() {
        return Optional.ofNullable(this.caEnabled);
    }

    /**
     * The list of server certificates. This parameter takes effect only for listeners that use SSL over TCP. **Note:** Only one server certificate is supported.
     * 
     */
    @Import(name="certificateIds")
    private @Nullable Output<List<String>> certificateIds;

    /**
     * @return The list of server certificates. This parameter takes effect only for listeners that use SSL over TCP. **Note:** Only one server certificate is supported.
     * 
     */
    public Optional<Output<List<String>>> certificateIds() {
        return Optional.ofNullable(this.certificateIds);
    }

    /**
     * The maximum number of connections that can be created per second on the NLB instance. Valid values: 0 to 1000000. 0 specifies that the number of connections is unlimited.
     * 
     */
    @Import(name="cps")
    private @Nullable Output<Integer> cps;

    /**
     * @return The maximum number of connections that can be created per second on the NLB instance. Valid values: 0 to 1000000. 0 specifies that the number of connections is unlimited.
     * 
     */
    public Optional<Output<Integer>> cps() {
        return Optional.ofNullable(this.cps);
    }

    /**
     * Full port listening end port. Valid values: `0` ~ `65535`. The value of the end port is less than the start port.
     * 
     */
    @Import(name="endPort")
    private @Nullable Output<Integer> endPort;

    /**
     * @return Full port listening end port. Valid values: `0` ~ `65535`. The value of the end port is less than the start port.
     * 
     */
    public Optional<Output<Integer>> endPort() {
        return Optional.ofNullable(this.endPort);
    }

    /**
     * The timeout period of an idle connection. Unit: seconds. Valid values: `1` to `900`. Default value: `900`.
     * 
     */
    @Import(name="idleTimeout")
    private @Nullable Output<Integer> idleTimeout;

    /**
     * @return The timeout period of an idle connection. Unit: seconds. Valid values: `1` to `900`. Default value: `900`.
     * 
     */
    public Optional<Output<Integer>> idleTimeout() {
        return Optional.ofNullable(this.idleTimeout);
    }

    /**
     * Custom listener name. The length is limited to 2 to 256 characters, supports Chinese and English letters, and can include numbers, commas (,), half-width periods (.), half-width semicolons (;), forward slashes (/), at(@), underscores (_), and dashes (-).
     * 
     */
    @Import(name="listenerDescription")
    private @Nullable Output<String> listenerDescription;

    /**
     * @return Custom listener name. The length is limited to 2 to 256 characters, supports Chinese and English letters, and can include numbers, commas (,), half-width periods (.), half-width semicolons (;), forward slashes (/), at(@), underscores (_), and dashes (-).
     * 
     */
    public Optional<Output<String>> listenerDescription() {
        return Optional.ofNullable(this.listenerDescription);
    }

    /**
     * Listening port. Valid values: 0 ~ 65535. `0`: indicates that full port listening is used. When set to `0`, you must configure `StartPort` and `EndPort`.
     * 
     */
    @Import(name="listenerPort", required=true)
    private Output<Integer> listenerPort;

    /**
     * @return Listening port. Valid values: 0 ~ 65535. `0`: indicates that full port listening is used. When set to `0`, you must configure `StartPort` and `EndPort`.
     * 
     */
    public Output<Integer> listenerPort() {
        return this.listenerPort;
    }

    /**
     * The listening protocol. Valid values: `TCP`, `UDP`, or `TCPSSL`.
     * 
     */
    @Import(name="listenerProtocol", required=true)
    private Output<String> listenerProtocol;

    /**
     * @return The listening protocol. Valid values: `TCP`, `UDP`, or `TCPSSL`.
     * 
     */
    public Output<String> listenerProtocol() {
        return this.listenerProtocol;
    }

    /**
     * The ID of the network-based server load balancer instance.
     * 
     */
    @Import(name="loadBalancerId", required=true)
    private Output<String> loadBalancerId;

    /**
     * @return The ID of the network-based server load balancer instance.
     * 
     */
    public Output<String> loadBalancerId() {
        return this.loadBalancerId;
    }

    /**
     * The maximum size of a TCP segment. Unit: bytes. Valid values: 0 to 1500. 0 specifies that the maximum segment size remains unchanged. **Note:** This parameter is supported only by listeners that use SSL over TCP.
     * 
     */
    @Import(name="mss")
    private @Nullable Output<Integer> mss;

    /**
     * @return The maximum size of a TCP segment. Unit: bytes. Valid values: 0 to 1500. 0 specifies that the maximum segment size remains unchanged. **Note:** This parameter is supported only by listeners that use SSL over TCP.
     * 
     */
    public Optional<Output<Integer>> mss() {
        return Optional.ofNullable(this.mss);
    }

    /**
     * Specifies whether to use the Proxy protocol to pass client IP addresses to backend servers.
     * 
     */
    @Import(name="proxyProtocolEnabled")
    private @Nullable Output<Boolean> proxyProtocolEnabled;

    /**
     * @return Specifies whether to use the Proxy protocol to pass client IP addresses to backend servers.
     * 
     */
    public Optional<Output<Boolean>> proxyProtocolEnabled() {
        return Optional.ofNullable(this.proxyProtocolEnabled);
    }

    /**
     * Specifies whether to enable fine-grained monitoring.
     * 
     */
    @Import(name="secSensorEnabled")
    private @Nullable Output<Boolean> secSensorEnabled;

    /**
     * @return Specifies whether to enable fine-grained monitoring.
     * 
     */
    public Optional<Output<Boolean>> secSensorEnabled() {
        return Optional.ofNullable(this.secSensorEnabled);
    }

    /**
     * The ID of the security policy. System security policies and custom security policies are supported.
     * System security policies valid values: `tls_cipher_policy_1_0` (default), `tls_cipher_policy_1_1,` `tls_cipher_policy_1_2`, `tls_cipher_policy_1_2_strict`, and `tls_cipher_policy_1_2_strict_with_1_3`.
     * Custom security policies can be created by resource `alicloud.nlb.SecurityPolicy`.
     * 
     */
    @Import(name="securityPolicyId")
    private @Nullable Output<String> securityPolicyId;

    /**
     * @return The ID of the security policy. System security policies and custom security policies are supported.
     * System security policies valid values: `tls_cipher_policy_1_0` (default), `tls_cipher_policy_1_1,` `tls_cipher_policy_1_2`, `tls_cipher_policy_1_2_strict`, and `tls_cipher_policy_1_2_strict_with_1_3`.
     * Custom security policies can be created by resource `alicloud.nlb.SecurityPolicy`.
     * 
     */
    public Optional<Output<String>> securityPolicyId() {
        return Optional.ofNullable(this.securityPolicyId);
    }

    /**
     * The ID of the server group.
     * 
     */
    @Import(name="serverGroupId", required=true)
    private Output<String> serverGroupId;

    /**
     * @return The ID of the server group.
     * 
     */
    public Output<String> serverGroupId() {
        return this.serverGroupId;
    }

    /**
     * Full Port listens to the starting port. Valid values: `0` ~ `65535`.
     * 
     */
    @Import(name="startPort")
    private @Nullable Output<Integer> startPort;

    /**
     * @return Full Port listens to the starting port. Valid values: `0` ~ `65535`.
     * 
     */
    public Optional<Output<Integer>> startPort() {
        return Optional.ofNullable(this.startPort);
    }

    /**
     * The status of the resource. Valid values: `Running`, `Stopped`.
     * 
     */
    @Import(name="status")
    private @Nullable Output<String> status;

    /**
     * @return The status of the resource. Valid values: `Running`, `Stopped`.
     * 
     */
    public Optional<Output<String>> status() {
        return Optional.ofNullable(this.status);
    }

    private ListenerArgs() {}

    private ListenerArgs(ListenerArgs $) {
        this.alpnEnabled = $.alpnEnabled;
        this.alpnPolicy = $.alpnPolicy;
        this.caCertificateIds = $.caCertificateIds;
        this.caEnabled = $.caEnabled;
        this.certificateIds = $.certificateIds;
        this.cps = $.cps;
        this.endPort = $.endPort;
        this.idleTimeout = $.idleTimeout;
        this.listenerDescription = $.listenerDescription;
        this.listenerPort = $.listenerPort;
        this.listenerProtocol = $.listenerProtocol;
        this.loadBalancerId = $.loadBalancerId;
        this.mss = $.mss;
        this.proxyProtocolEnabled = $.proxyProtocolEnabled;
        this.secSensorEnabled = $.secSensorEnabled;
        this.securityPolicyId = $.securityPolicyId;
        this.serverGroupId = $.serverGroupId;
        this.startPort = $.startPort;
        this.status = $.status;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ListenerArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ListenerArgs $;

        public Builder() {
            $ = new ListenerArgs();
        }

        public Builder(ListenerArgs defaults) {
            $ = new ListenerArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param alpnEnabled Specifies whether to enable Application-Layer Protocol Negotiation (ALPN).
         * 
         * @return builder
         * 
         */
        public Builder alpnEnabled(@Nullable Output<Boolean> alpnEnabled) {
            $.alpnEnabled = alpnEnabled;
            return this;
        }

        /**
         * @param alpnEnabled Specifies whether to enable Application-Layer Protocol Negotiation (ALPN).
         * 
         * @return builder
         * 
         */
        public Builder alpnEnabled(Boolean alpnEnabled) {
            return alpnEnabled(Output.of(alpnEnabled));
        }

        /**
         * @param alpnPolicy The ALPN policy.
         * 
         * @return builder
         * 
         */
        public Builder alpnPolicy(@Nullable Output<String> alpnPolicy) {
            $.alpnPolicy = alpnPolicy;
            return this;
        }

        /**
         * @param alpnPolicy The ALPN policy.
         * 
         * @return builder
         * 
         */
        public Builder alpnPolicy(String alpnPolicy) {
            return alpnPolicy(Output.of(alpnPolicy));
        }

        /**
         * @param caCertificateIds The list of certificate authority (CA) certificates. This parameter takes effect only for listeners that use SSL over TCP. **Note:** Only one CA certificate is supported.
         * 
         * @return builder
         * 
         */
        public Builder caCertificateIds(@Nullable Output<List<String>> caCertificateIds) {
            $.caCertificateIds = caCertificateIds;
            return this;
        }

        /**
         * @param caCertificateIds The list of certificate authority (CA) certificates. This parameter takes effect only for listeners that use SSL over TCP. **Note:** Only one CA certificate is supported.
         * 
         * @return builder
         * 
         */
        public Builder caCertificateIds(List<String> caCertificateIds) {
            return caCertificateIds(Output.of(caCertificateIds));
        }

        /**
         * @param caCertificateIds The list of certificate authority (CA) certificates. This parameter takes effect only for listeners that use SSL over TCP. **Note:** Only one CA certificate is supported.
         * 
         * @return builder
         * 
         */
        public Builder caCertificateIds(String... caCertificateIds) {
            return caCertificateIds(List.of(caCertificateIds));
        }

        /**
         * @param caEnabled Specifies whether to enable mutual authentication.
         * 
         * @return builder
         * 
         */
        public Builder caEnabled(@Nullable Output<Boolean> caEnabled) {
            $.caEnabled = caEnabled;
            return this;
        }

        /**
         * @param caEnabled Specifies whether to enable mutual authentication.
         * 
         * @return builder
         * 
         */
        public Builder caEnabled(Boolean caEnabled) {
            return caEnabled(Output.of(caEnabled));
        }

        /**
         * @param certificateIds The list of server certificates. This parameter takes effect only for listeners that use SSL over TCP. **Note:** Only one server certificate is supported.
         * 
         * @return builder
         * 
         */
        public Builder certificateIds(@Nullable Output<List<String>> certificateIds) {
            $.certificateIds = certificateIds;
            return this;
        }

        /**
         * @param certificateIds The list of server certificates. This parameter takes effect only for listeners that use SSL over TCP. **Note:** Only one server certificate is supported.
         * 
         * @return builder
         * 
         */
        public Builder certificateIds(List<String> certificateIds) {
            return certificateIds(Output.of(certificateIds));
        }

        /**
         * @param certificateIds The list of server certificates. This parameter takes effect only for listeners that use SSL over TCP. **Note:** Only one server certificate is supported.
         * 
         * @return builder
         * 
         */
        public Builder certificateIds(String... certificateIds) {
            return certificateIds(List.of(certificateIds));
        }

        /**
         * @param cps The maximum number of connections that can be created per second on the NLB instance. Valid values: 0 to 1000000. 0 specifies that the number of connections is unlimited.
         * 
         * @return builder
         * 
         */
        public Builder cps(@Nullable Output<Integer> cps) {
            $.cps = cps;
            return this;
        }

        /**
         * @param cps The maximum number of connections that can be created per second on the NLB instance. Valid values: 0 to 1000000. 0 specifies that the number of connections is unlimited.
         * 
         * @return builder
         * 
         */
        public Builder cps(Integer cps) {
            return cps(Output.of(cps));
        }

        /**
         * @param endPort Full port listening end port. Valid values: `0` ~ `65535`. The value of the end port is less than the start port.
         * 
         * @return builder
         * 
         */
        public Builder endPort(@Nullable Output<Integer> endPort) {
            $.endPort = endPort;
            return this;
        }

        /**
         * @param endPort Full port listening end port. Valid values: `0` ~ `65535`. The value of the end port is less than the start port.
         * 
         * @return builder
         * 
         */
        public Builder endPort(Integer endPort) {
            return endPort(Output.of(endPort));
        }

        /**
         * @param idleTimeout The timeout period of an idle connection. Unit: seconds. Valid values: `1` to `900`. Default value: `900`.
         * 
         * @return builder
         * 
         */
        public Builder idleTimeout(@Nullable Output<Integer> idleTimeout) {
            $.idleTimeout = idleTimeout;
            return this;
        }

        /**
         * @param idleTimeout The timeout period of an idle connection. Unit: seconds. Valid values: `1` to `900`. Default value: `900`.
         * 
         * @return builder
         * 
         */
        public Builder idleTimeout(Integer idleTimeout) {
            return idleTimeout(Output.of(idleTimeout));
        }

        /**
         * @param listenerDescription Custom listener name. The length is limited to 2 to 256 characters, supports Chinese and English letters, and can include numbers, commas (,), half-width periods (.), half-width semicolons (;), forward slashes (/), at(@), underscores (_), and dashes (-).
         * 
         * @return builder
         * 
         */
        public Builder listenerDescription(@Nullable Output<String> listenerDescription) {
            $.listenerDescription = listenerDescription;
            return this;
        }

        /**
         * @param listenerDescription Custom listener name. The length is limited to 2 to 256 characters, supports Chinese and English letters, and can include numbers, commas (,), half-width periods (.), half-width semicolons (;), forward slashes (/), at(@), underscores (_), and dashes (-).
         * 
         * @return builder
         * 
         */
        public Builder listenerDescription(String listenerDescription) {
            return listenerDescription(Output.of(listenerDescription));
        }

        /**
         * @param listenerPort Listening port. Valid values: 0 ~ 65535. `0`: indicates that full port listening is used. When set to `0`, you must configure `StartPort` and `EndPort`.
         * 
         * @return builder
         * 
         */
        public Builder listenerPort(Output<Integer> listenerPort) {
            $.listenerPort = listenerPort;
            return this;
        }

        /**
         * @param listenerPort Listening port. Valid values: 0 ~ 65535. `0`: indicates that full port listening is used. When set to `0`, you must configure `StartPort` and `EndPort`.
         * 
         * @return builder
         * 
         */
        public Builder listenerPort(Integer listenerPort) {
            return listenerPort(Output.of(listenerPort));
        }

        /**
         * @param listenerProtocol The listening protocol. Valid values: `TCP`, `UDP`, or `TCPSSL`.
         * 
         * @return builder
         * 
         */
        public Builder listenerProtocol(Output<String> listenerProtocol) {
            $.listenerProtocol = listenerProtocol;
            return this;
        }

        /**
         * @param listenerProtocol The listening protocol. Valid values: `TCP`, `UDP`, or `TCPSSL`.
         * 
         * @return builder
         * 
         */
        public Builder listenerProtocol(String listenerProtocol) {
            return listenerProtocol(Output.of(listenerProtocol));
        }

        /**
         * @param loadBalancerId The ID of the network-based server load balancer instance.
         * 
         * @return builder
         * 
         */
        public Builder loadBalancerId(Output<String> loadBalancerId) {
            $.loadBalancerId = loadBalancerId;
            return this;
        }

        /**
         * @param loadBalancerId The ID of the network-based server load balancer instance.
         * 
         * @return builder
         * 
         */
        public Builder loadBalancerId(String loadBalancerId) {
            return loadBalancerId(Output.of(loadBalancerId));
        }

        /**
         * @param mss The maximum size of a TCP segment. Unit: bytes. Valid values: 0 to 1500. 0 specifies that the maximum segment size remains unchanged. **Note:** This parameter is supported only by listeners that use SSL over TCP.
         * 
         * @return builder
         * 
         */
        public Builder mss(@Nullable Output<Integer> mss) {
            $.mss = mss;
            return this;
        }

        /**
         * @param mss The maximum size of a TCP segment. Unit: bytes. Valid values: 0 to 1500. 0 specifies that the maximum segment size remains unchanged. **Note:** This parameter is supported only by listeners that use SSL over TCP.
         * 
         * @return builder
         * 
         */
        public Builder mss(Integer mss) {
            return mss(Output.of(mss));
        }

        /**
         * @param proxyProtocolEnabled Specifies whether to use the Proxy protocol to pass client IP addresses to backend servers.
         * 
         * @return builder
         * 
         */
        public Builder proxyProtocolEnabled(@Nullable Output<Boolean> proxyProtocolEnabled) {
            $.proxyProtocolEnabled = proxyProtocolEnabled;
            return this;
        }

        /**
         * @param proxyProtocolEnabled Specifies whether to use the Proxy protocol to pass client IP addresses to backend servers.
         * 
         * @return builder
         * 
         */
        public Builder proxyProtocolEnabled(Boolean proxyProtocolEnabled) {
            return proxyProtocolEnabled(Output.of(proxyProtocolEnabled));
        }

        /**
         * @param secSensorEnabled Specifies whether to enable fine-grained monitoring.
         * 
         * @return builder
         * 
         */
        public Builder secSensorEnabled(@Nullable Output<Boolean> secSensorEnabled) {
            $.secSensorEnabled = secSensorEnabled;
            return this;
        }

        /**
         * @param secSensorEnabled Specifies whether to enable fine-grained monitoring.
         * 
         * @return builder
         * 
         */
        public Builder secSensorEnabled(Boolean secSensorEnabled) {
            return secSensorEnabled(Output.of(secSensorEnabled));
        }

        /**
         * @param securityPolicyId The ID of the security policy. System security policies and custom security policies are supported.
         * System security policies valid values: `tls_cipher_policy_1_0` (default), `tls_cipher_policy_1_1,` `tls_cipher_policy_1_2`, `tls_cipher_policy_1_2_strict`, and `tls_cipher_policy_1_2_strict_with_1_3`.
         * Custom security policies can be created by resource `alicloud.nlb.SecurityPolicy`.
         * 
         * @return builder
         * 
         */
        public Builder securityPolicyId(@Nullable Output<String> securityPolicyId) {
            $.securityPolicyId = securityPolicyId;
            return this;
        }

        /**
         * @param securityPolicyId The ID of the security policy. System security policies and custom security policies are supported.
         * System security policies valid values: `tls_cipher_policy_1_0` (default), `tls_cipher_policy_1_1,` `tls_cipher_policy_1_2`, `tls_cipher_policy_1_2_strict`, and `tls_cipher_policy_1_2_strict_with_1_3`.
         * Custom security policies can be created by resource `alicloud.nlb.SecurityPolicy`.
         * 
         * @return builder
         * 
         */
        public Builder securityPolicyId(String securityPolicyId) {
            return securityPolicyId(Output.of(securityPolicyId));
        }

        /**
         * @param serverGroupId The ID of the server group.
         * 
         * @return builder
         * 
         */
        public Builder serverGroupId(Output<String> serverGroupId) {
            $.serverGroupId = serverGroupId;
            return this;
        }

        /**
         * @param serverGroupId The ID of the server group.
         * 
         * @return builder
         * 
         */
        public Builder serverGroupId(String serverGroupId) {
            return serverGroupId(Output.of(serverGroupId));
        }

        /**
         * @param startPort Full Port listens to the starting port. Valid values: `0` ~ `65535`.
         * 
         * @return builder
         * 
         */
        public Builder startPort(@Nullable Output<Integer> startPort) {
            $.startPort = startPort;
            return this;
        }

        /**
         * @param startPort Full Port listens to the starting port. Valid values: `0` ~ `65535`.
         * 
         * @return builder
         * 
         */
        public Builder startPort(Integer startPort) {
            return startPort(Output.of(startPort));
        }

        /**
         * @param status The status of the resource. Valid values: `Running`, `Stopped`.
         * 
         * @return builder
         * 
         */
        public Builder status(@Nullable Output<String> status) {
            $.status = status;
            return this;
        }

        /**
         * @param status The status of the resource. Valid values: `Running`, `Stopped`.
         * 
         * @return builder
         * 
         */
        public Builder status(String status) {
            return status(Output.of(status));
        }

        public ListenerArgs build() {
            $.listenerPort = Objects.requireNonNull($.listenerPort, "expected parameter 'listenerPort' to be non-null");
            $.listenerProtocol = Objects.requireNonNull($.listenerProtocol, "expected parameter 'listenerProtocol' to be non-null");
            $.loadBalancerId = Objects.requireNonNull($.loadBalancerId, "expected parameter 'loadBalancerId' to be non-null");
            $.serverGroupId = Objects.requireNonNull($.serverGroupId, "expected parameter 'serverGroupId' to be non-null");
            return $;
        }
    }

}
