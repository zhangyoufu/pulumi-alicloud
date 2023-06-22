// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.sae.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class IngressRule {
    /**
     * @return Target application ID.
     * 
     */
    private String appId;
    /**
     * @return Target application name.
     * 
     */
    private String appName;
    /**
     * @return The backend protocol.
     * 
     */
    private @Nullable String backendProtocol;
    /**
     * @return Application backend port.
     * 
     */
    private Integer containerPort;
    /**
     * @return Application domain name.
     * 
     */
    private String domain;
    /**
     * @return URL path.
     * 
     */
    private String path;
    /**
     * @return The rewrite path.
     * 
     */
    private @Nullable String rewritePath;

    private IngressRule() {}
    /**
     * @return Target application ID.
     * 
     */
    public String appId() {
        return this.appId;
    }
    /**
     * @return Target application name.
     * 
     */
    public String appName() {
        return this.appName;
    }
    /**
     * @return The backend protocol.
     * 
     */
    public Optional<String> backendProtocol() {
        return Optional.ofNullable(this.backendProtocol);
    }
    /**
     * @return Application backend port.
     * 
     */
    public Integer containerPort() {
        return this.containerPort;
    }
    /**
     * @return Application domain name.
     * 
     */
    public String domain() {
        return this.domain;
    }
    /**
     * @return URL path.
     * 
     */
    public String path() {
        return this.path;
    }
    /**
     * @return The rewrite path.
     * 
     */
    public Optional<String> rewritePath() {
        return Optional.ofNullable(this.rewritePath);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(IngressRule defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String appId;
        private String appName;
        private @Nullable String backendProtocol;
        private Integer containerPort;
        private String domain;
        private String path;
        private @Nullable String rewritePath;
        public Builder() {}
        public Builder(IngressRule defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.appId = defaults.appId;
    	      this.appName = defaults.appName;
    	      this.backendProtocol = defaults.backendProtocol;
    	      this.containerPort = defaults.containerPort;
    	      this.domain = defaults.domain;
    	      this.path = defaults.path;
    	      this.rewritePath = defaults.rewritePath;
        }

        @CustomType.Setter
        public Builder appId(String appId) {
            this.appId = Objects.requireNonNull(appId);
            return this;
        }
        @CustomType.Setter
        public Builder appName(String appName) {
            this.appName = Objects.requireNonNull(appName);
            return this;
        }
        @CustomType.Setter
        public Builder backendProtocol(@Nullable String backendProtocol) {
            this.backendProtocol = backendProtocol;
            return this;
        }
        @CustomType.Setter
        public Builder containerPort(Integer containerPort) {
            this.containerPort = Objects.requireNonNull(containerPort);
            return this;
        }
        @CustomType.Setter
        public Builder domain(String domain) {
            this.domain = Objects.requireNonNull(domain);
            return this;
        }
        @CustomType.Setter
        public Builder path(String path) {
            this.path = Objects.requireNonNull(path);
            return this;
        }
        @CustomType.Setter
        public Builder rewritePath(@Nullable String rewritePath) {
            this.rewritePath = rewritePath;
            return this;
        }
        public IngressRule build() {
            final var o = new IngressRule();
            o.appId = appId;
            o.appName = appName;
            o.backendProtocol = backendProtocol;
            o.containerPort = containerPort;
            o.domain = domain;
            o.path = path;
            o.rewritePath = rewritePath;
            return o;
        }
    }
}
