// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.threatdetection.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetBaselineStrategiesStrategy {
    /**
     * @return The ID of the baseline check policy.
     * 
     */
    private String baselineStrategyId;
    /**
     * @return Policy name.
     * 
     */
    private String baselineStrategyName;
    /**
     * @return The type of policy. Value:-**common**: standard policy-**custom**: custom policy
     * 
     */
    private String customType;
    /**
     * @return The detection period of the policy.
     * 
     */
    private Integer cycleDays;
    /**
     * @return The detection period of the policy. Value:
     * * **0**: 0:00~06:00
     * * **6**: 6:00~12:00
     * * *12**: 12:00~18:00
     * * **18**: 18:00~24:00
     * 
     */
    private Integer cycleStartTime;
    /**
     * @return The baseline check policy execution end time.
     * 
     */
    private String endTime;
    /**
     * @return The ID of the baseline check policy.
     * 
     */
    private String id;
    private String riskSubTypeName;
    /**
     * @return The baseline check policy start time.
     * 
     */
    private String startTime;
    private String targetType;

    private GetBaselineStrategiesStrategy() {}
    /**
     * @return The ID of the baseline check policy.
     * 
     */
    public String baselineStrategyId() {
        return this.baselineStrategyId;
    }
    /**
     * @return Policy name.
     * 
     */
    public String baselineStrategyName() {
        return this.baselineStrategyName;
    }
    /**
     * @return The type of policy. Value:-**common**: standard policy-**custom**: custom policy
     * 
     */
    public String customType() {
        return this.customType;
    }
    /**
     * @return The detection period of the policy.
     * 
     */
    public Integer cycleDays() {
        return this.cycleDays;
    }
    /**
     * @return The detection period of the policy. Value:
     * * **0**: 0:00~06:00
     * * **6**: 6:00~12:00
     * * *12**: 12:00~18:00
     * * **18**: 18:00~24:00
     * 
     */
    public Integer cycleStartTime() {
        return this.cycleStartTime;
    }
    /**
     * @return The baseline check policy execution end time.
     * 
     */
    public String endTime() {
        return this.endTime;
    }
    /**
     * @return The ID of the baseline check policy.
     * 
     */
    public String id() {
        return this.id;
    }
    public String riskSubTypeName() {
        return this.riskSubTypeName;
    }
    /**
     * @return The baseline check policy start time.
     * 
     */
    public String startTime() {
        return this.startTime;
    }
    public String targetType() {
        return this.targetType;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetBaselineStrategiesStrategy defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String baselineStrategyId;
        private String baselineStrategyName;
        private String customType;
        private Integer cycleDays;
        private Integer cycleStartTime;
        private String endTime;
        private String id;
        private String riskSubTypeName;
        private String startTime;
        private String targetType;
        public Builder() {}
        public Builder(GetBaselineStrategiesStrategy defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.baselineStrategyId = defaults.baselineStrategyId;
    	      this.baselineStrategyName = defaults.baselineStrategyName;
    	      this.customType = defaults.customType;
    	      this.cycleDays = defaults.cycleDays;
    	      this.cycleStartTime = defaults.cycleStartTime;
    	      this.endTime = defaults.endTime;
    	      this.id = defaults.id;
    	      this.riskSubTypeName = defaults.riskSubTypeName;
    	      this.startTime = defaults.startTime;
    	      this.targetType = defaults.targetType;
        }

        @CustomType.Setter
        public Builder baselineStrategyId(String baselineStrategyId) {
            this.baselineStrategyId = Objects.requireNonNull(baselineStrategyId);
            return this;
        }
        @CustomType.Setter
        public Builder baselineStrategyName(String baselineStrategyName) {
            this.baselineStrategyName = Objects.requireNonNull(baselineStrategyName);
            return this;
        }
        @CustomType.Setter
        public Builder customType(String customType) {
            this.customType = Objects.requireNonNull(customType);
            return this;
        }
        @CustomType.Setter
        public Builder cycleDays(Integer cycleDays) {
            this.cycleDays = Objects.requireNonNull(cycleDays);
            return this;
        }
        @CustomType.Setter
        public Builder cycleStartTime(Integer cycleStartTime) {
            this.cycleStartTime = Objects.requireNonNull(cycleStartTime);
            return this;
        }
        @CustomType.Setter
        public Builder endTime(String endTime) {
            this.endTime = Objects.requireNonNull(endTime);
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder riskSubTypeName(String riskSubTypeName) {
            this.riskSubTypeName = Objects.requireNonNull(riskSubTypeName);
            return this;
        }
        @CustomType.Setter
        public Builder startTime(String startTime) {
            this.startTime = Objects.requireNonNull(startTime);
            return this;
        }
        @CustomType.Setter
        public Builder targetType(String targetType) {
            this.targetType = Objects.requireNonNull(targetType);
            return this;
        }
        public GetBaselineStrategiesStrategy build() {
            final var o = new GetBaselineStrategiesStrategy();
            o.baselineStrategyId = baselineStrategyId;
            o.baselineStrategyName = baselineStrategyName;
            o.customType = customType;
            o.cycleDays = cycleDays;
            o.cycleStartTime = cycleStartTime;
            o.endTime = endTime;
            o.id = id;
            o.riskSubTypeName = riskSubTypeName;
            o.startTime = startTime;
            o.targetType = targetType;
            return o;
        }
    }
}
