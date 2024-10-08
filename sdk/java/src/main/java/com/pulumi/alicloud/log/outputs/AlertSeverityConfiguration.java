// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.log.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Integer;
import java.lang.String;
import java.util.Map;
import java.util.Objects;

@CustomType
public final class AlertSeverityConfiguration {
    /**
     * @return Severity when this condition is met.
     * 
     */
    private Map<String,String> evalCondition;
    /**
     * @return Severity for new alert, including 2,4,6,8,10 for Report,Low,Medium,High,Critical.
     * 
     */
    private Integer severity;

    private AlertSeverityConfiguration() {}
    /**
     * @return Severity when this condition is met.
     * 
     */
    public Map<String,String> evalCondition() {
        return this.evalCondition;
    }
    /**
     * @return Severity for new alert, including 2,4,6,8,10 for Report,Low,Medium,High,Critical.
     * 
     */
    public Integer severity() {
        return this.severity;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(AlertSeverityConfiguration defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private Map<String,String> evalCondition;
        private Integer severity;
        public Builder() {}
        public Builder(AlertSeverityConfiguration defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.evalCondition = defaults.evalCondition;
    	      this.severity = defaults.severity;
        }

        @CustomType.Setter
        public Builder evalCondition(Map<String,String> evalCondition) {
            if (evalCondition == null) {
              throw new MissingRequiredPropertyException("AlertSeverityConfiguration", "evalCondition");
            }
            this.evalCondition = evalCondition;
            return this;
        }
        @CustomType.Setter
        public Builder severity(Integer severity) {
            if (severity == null) {
              throw new MissingRequiredPropertyException("AlertSeverityConfiguration", "severity");
            }
            this.severity = severity;
            return this;
        }
        public AlertSeverityConfiguration build() {
            final var _resultValue = new AlertSeverityConfiguration();
            _resultValue.evalCondition = evalCondition;
            _resultValue.severity = severity;
            return _resultValue;
        }
    }
}
