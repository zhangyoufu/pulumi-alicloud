// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.alb.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;

@CustomType
public final class GetHealthCheckTemplatesTemplate {
    /**
     * @return The HTTP status code that indicates a successful health check.
     * 
     */
    private List<String> healthCheckCodes;
    /**
     * @return The number of the port that is used for health checks.  Valid values: `0` to `65535`.  Default value:`  0 `. This default value indicates that the backend server is used for health checks.
     * 
     */
    private Integer healthCheckConnectPort;
    /**
     * @return The domain name that is used for health checks. Default value:  `$SERVER_IP`. The domain name must be 1 to 80 characters in length.
     * 
     */
    private String healthCheckHost;
    /**
     * @return The version of the HTTP protocol.  Valid values: `HTTP1.0` and `HTTP1.1`.  Default value: `HTTP1.1`.
     * 
     */
    private String healthCheckHttpVersion;
    /**
     * @return The time interval between two consecutive health checks.  Valid values: `1` to `50`. Unit: seconds.  Default value: `2`.
     * 
     */
    private Integer healthCheckInterval;
    /**
     * @return The health check method.  Valid values: `GET` and `HEAD`.  Default value: `HEAD`.
     * 
     */
    private String healthCheckMethod;
    /**
     * @return The URL that is used for health checks.  The URL must be 1 to 80 characters in length, and can contain letters, digits, hyphens (-), forward slashes (/), periods (.), percent signs (%!)(MISSING), question marks (?), number signs (#), and ampersands (&amp;). The URL can also contain the following extended characters: `  _ ; ~ ! ( )* [ ] @ $ ^ : &#39; , +. The URL must start with a forward slash (/) `.
     * 
     */
    private String healthCheckPath;
    /**
     * @return The protocol that is used for health checks.  Valid values: HTTP and TCP.  Default value: HTTP.
     * 
     */
    private String healthCheckProtocol;
    /**
     * @return The ID of the resource.
     * 
     */
    private String healthCheckTemplateId;
    /**
     * @return The name of the health check template.  The name must be 2 to 128 characters in length, and can contain letters, digits, periods (.), underscores (_), and hyphens (-). The name must start with a letter.
     * 
     */
    private String healthCheckTemplateName;
    /**
     * @return The timeout period of a health check response. If the backend Elastic Compute Service (ECS) instance does not send an expected response within the specified period of time, the health check fails.  Valid values: `1` to `300`. Unit: seconds.  Default value: `5`.
     * 
     */
    private Integer healthCheckTimeout;
    /**
     * @return The number of times that an unhealthy backend server must consecutively pass health checks before it is declared healthy (from fail to success). Valid values: `2` to `10`.  Default value: `3`. Unit: seconds.
     * 
     */
    private Integer healthyThreshold;
    /**
     * @return The ID of the Health Check Template.
     * 
     */
    private String id;
    /**
     * @return The number of times that an healthy backend server must consecutively fail health checks before it is declared unhealthy (from success to fail). Valid values: `2` to `10`.  Default value: `3`. Unit: seconds.
     * 
     */
    private Integer unhealthyThreshold;

    private GetHealthCheckTemplatesTemplate() {}
    /**
     * @return The HTTP status code that indicates a successful health check.
     * 
     */
    public List<String> healthCheckCodes() {
        return this.healthCheckCodes;
    }
    /**
     * @return The number of the port that is used for health checks.  Valid values: `0` to `65535`.  Default value:`  0 `. This default value indicates that the backend server is used for health checks.
     * 
     */
    public Integer healthCheckConnectPort() {
        return this.healthCheckConnectPort;
    }
    /**
     * @return The domain name that is used for health checks. Default value:  `$SERVER_IP`. The domain name must be 1 to 80 characters in length.
     * 
     */
    public String healthCheckHost() {
        return this.healthCheckHost;
    }
    /**
     * @return The version of the HTTP protocol.  Valid values: `HTTP1.0` and `HTTP1.1`.  Default value: `HTTP1.1`.
     * 
     */
    public String healthCheckHttpVersion() {
        return this.healthCheckHttpVersion;
    }
    /**
     * @return The time interval between two consecutive health checks.  Valid values: `1` to `50`. Unit: seconds.  Default value: `2`.
     * 
     */
    public Integer healthCheckInterval() {
        return this.healthCheckInterval;
    }
    /**
     * @return The health check method.  Valid values: `GET` and `HEAD`.  Default value: `HEAD`.
     * 
     */
    public String healthCheckMethod() {
        return this.healthCheckMethod;
    }
    /**
     * @return The URL that is used for health checks.  The URL must be 1 to 80 characters in length, and can contain letters, digits, hyphens (-), forward slashes (/), periods (.), percent signs (%!)(MISSING), question marks (?), number signs (#), and ampersands (&amp;). The URL can also contain the following extended characters: `  _ ; ~ ! ( )* [ ] @ $ ^ : &#39; , +. The URL must start with a forward slash (/) `.
     * 
     */
    public String healthCheckPath() {
        return this.healthCheckPath;
    }
    /**
     * @return The protocol that is used for health checks.  Valid values: HTTP and TCP.  Default value: HTTP.
     * 
     */
    public String healthCheckProtocol() {
        return this.healthCheckProtocol;
    }
    /**
     * @return The ID of the resource.
     * 
     */
    public String healthCheckTemplateId() {
        return this.healthCheckTemplateId;
    }
    /**
     * @return The name of the health check template.  The name must be 2 to 128 characters in length, and can contain letters, digits, periods (.), underscores (_), and hyphens (-). The name must start with a letter.
     * 
     */
    public String healthCheckTemplateName() {
        return this.healthCheckTemplateName;
    }
    /**
     * @return The timeout period of a health check response. If the backend Elastic Compute Service (ECS) instance does not send an expected response within the specified period of time, the health check fails.  Valid values: `1` to `300`. Unit: seconds.  Default value: `5`.
     * 
     */
    public Integer healthCheckTimeout() {
        return this.healthCheckTimeout;
    }
    /**
     * @return The number of times that an unhealthy backend server must consecutively pass health checks before it is declared healthy (from fail to success). Valid values: `2` to `10`.  Default value: `3`. Unit: seconds.
     * 
     */
    public Integer healthyThreshold() {
        return this.healthyThreshold;
    }
    /**
     * @return The ID of the Health Check Template.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return The number of times that an healthy backend server must consecutively fail health checks before it is declared unhealthy (from success to fail). Valid values: `2` to `10`.  Default value: `3`. Unit: seconds.
     * 
     */
    public Integer unhealthyThreshold() {
        return this.unhealthyThreshold;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetHealthCheckTemplatesTemplate defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private List<String> healthCheckCodes;
        private Integer healthCheckConnectPort;
        private String healthCheckHost;
        private String healthCheckHttpVersion;
        private Integer healthCheckInterval;
        private String healthCheckMethod;
        private String healthCheckPath;
        private String healthCheckProtocol;
        private String healthCheckTemplateId;
        private String healthCheckTemplateName;
        private Integer healthCheckTimeout;
        private Integer healthyThreshold;
        private String id;
        private Integer unhealthyThreshold;
        public Builder() {}
        public Builder(GetHealthCheckTemplatesTemplate defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.healthCheckCodes = defaults.healthCheckCodes;
    	      this.healthCheckConnectPort = defaults.healthCheckConnectPort;
    	      this.healthCheckHost = defaults.healthCheckHost;
    	      this.healthCheckHttpVersion = defaults.healthCheckHttpVersion;
    	      this.healthCheckInterval = defaults.healthCheckInterval;
    	      this.healthCheckMethod = defaults.healthCheckMethod;
    	      this.healthCheckPath = defaults.healthCheckPath;
    	      this.healthCheckProtocol = defaults.healthCheckProtocol;
    	      this.healthCheckTemplateId = defaults.healthCheckTemplateId;
    	      this.healthCheckTemplateName = defaults.healthCheckTemplateName;
    	      this.healthCheckTimeout = defaults.healthCheckTimeout;
    	      this.healthyThreshold = defaults.healthyThreshold;
    	      this.id = defaults.id;
    	      this.unhealthyThreshold = defaults.unhealthyThreshold;
        }

        @CustomType.Setter
        public Builder healthCheckCodes(List<String> healthCheckCodes) {
            if (healthCheckCodes == null) {
              throw new MissingRequiredPropertyException("GetHealthCheckTemplatesTemplate", "healthCheckCodes");
            }
            this.healthCheckCodes = healthCheckCodes;
            return this;
        }
        public Builder healthCheckCodes(String... healthCheckCodes) {
            return healthCheckCodes(List.of(healthCheckCodes));
        }
        @CustomType.Setter
        public Builder healthCheckConnectPort(Integer healthCheckConnectPort) {
            if (healthCheckConnectPort == null) {
              throw new MissingRequiredPropertyException("GetHealthCheckTemplatesTemplate", "healthCheckConnectPort");
            }
            this.healthCheckConnectPort = healthCheckConnectPort;
            return this;
        }
        @CustomType.Setter
        public Builder healthCheckHost(String healthCheckHost) {
            if (healthCheckHost == null) {
              throw new MissingRequiredPropertyException("GetHealthCheckTemplatesTemplate", "healthCheckHost");
            }
            this.healthCheckHost = healthCheckHost;
            return this;
        }
        @CustomType.Setter
        public Builder healthCheckHttpVersion(String healthCheckHttpVersion) {
            if (healthCheckHttpVersion == null) {
              throw new MissingRequiredPropertyException("GetHealthCheckTemplatesTemplate", "healthCheckHttpVersion");
            }
            this.healthCheckHttpVersion = healthCheckHttpVersion;
            return this;
        }
        @CustomType.Setter
        public Builder healthCheckInterval(Integer healthCheckInterval) {
            if (healthCheckInterval == null) {
              throw new MissingRequiredPropertyException("GetHealthCheckTemplatesTemplate", "healthCheckInterval");
            }
            this.healthCheckInterval = healthCheckInterval;
            return this;
        }
        @CustomType.Setter
        public Builder healthCheckMethod(String healthCheckMethod) {
            if (healthCheckMethod == null) {
              throw new MissingRequiredPropertyException("GetHealthCheckTemplatesTemplate", "healthCheckMethod");
            }
            this.healthCheckMethod = healthCheckMethod;
            return this;
        }
        @CustomType.Setter
        public Builder healthCheckPath(String healthCheckPath) {
            if (healthCheckPath == null) {
              throw new MissingRequiredPropertyException("GetHealthCheckTemplatesTemplate", "healthCheckPath");
            }
            this.healthCheckPath = healthCheckPath;
            return this;
        }
        @CustomType.Setter
        public Builder healthCheckProtocol(String healthCheckProtocol) {
            if (healthCheckProtocol == null) {
              throw new MissingRequiredPropertyException("GetHealthCheckTemplatesTemplate", "healthCheckProtocol");
            }
            this.healthCheckProtocol = healthCheckProtocol;
            return this;
        }
        @CustomType.Setter
        public Builder healthCheckTemplateId(String healthCheckTemplateId) {
            if (healthCheckTemplateId == null) {
              throw new MissingRequiredPropertyException("GetHealthCheckTemplatesTemplate", "healthCheckTemplateId");
            }
            this.healthCheckTemplateId = healthCheckTemplateId;
            return this;
        }
        @CustomType.Setter
        public Builder healthCheckTemplateName(String healthCheckTemplateName) {
            if (healthCheckTemplateName == null) {
              throw new MissingRequiredPropertyException("GetHealthCheckTemplatesTemplate", "healthCheckTemplateName");
            }
            this.healthCheckTemplateName = healthCheckTemplateName;
            return this;
        }
        @CustomType.Setter
        public Builder healthCheckTimeout(Integer healthCheckTimeout) {
            if (healthCheckTimeout == null) {
              throw new MissingRequiredPropertyException("GetHealthCheckTemplatesTemplate", "healthCheckTimeout");
            }
            this.healthCheckTimeout = healthCheckTimeout;
            return this;
        }
        @CustomType.Setter
        public Builder healthyThreshold(Integer healthyThreshold) {
            if (healthyThreshold == null) {
              throw new MissingRequiredPropertyException("GetHealthCheckTemplatesTemplate", "healthyThreshold");
            }
            this.healthyThreshold = healthyThreshold;
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            if (id == null) {
              throw new MissingRequiredPropertyException("GetHealthCheckTemplatesTemplate", "id");
            }
            this.id = id;
            return this;
        }
        @CustomType.Setter
        public Builder unhealthyThreshold(Integer unhealthyThreshold) {
            if (unhealthyThreshold == null) {
              throw new MissingRequiredPropertyException("GetHealthCheckTemplatesTemplate", "unhealthyThreshold");
            }
            this.unhealthyThreshold = unhealthyThreshold;
            return this;
        }
        public GetHealthCheckTemplatesTemplate build() {
            final var _resultValue = new GetHealthCheckTemplatesTemplate();
            _resultValue.healthCheckCodes = healthCheckCodes;
            _resultValue.healthCheckConnectPort = healthCheckConnectPort;
            _resultValue.healthCheckHost = healthCheckHost;
            _resultValue.healthCheckHttpVersion = healthCheckHttpVersion;
            _resultValue.healthCheckInterval = healthCheckInterval;
            _resultValue.healthCheckMethod = healthCheckMethod;
            _resultValue.healthCheckPath = healthCheckPath;
            _resultValue.healthCheckProtocol = healthCheckProtocol;
            _resultValue.healthCheckTemplateId = healthCheckTemplateId;
            _resultValue.healthCheckTemplateName = healthCheckTemplateName;
            _resultValue.healthCheckTimeout = healthCheckTimeout;
            _resultValue.healthyThreshold = healthyThreshold;
            _resultValue.id = id;
            _resultValue.unhealthyThreshold = unhealthyThreshold;
            return _resultValue;
        }
    }
}
