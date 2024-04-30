// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.sls.inputs;

import com.pulumi.alicloud.sls.inputs.AlertConfigurationSeverityConfigurationEvalConditionArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class AlertConfigurationSeverityConfigurationArgs extends com.pulumi.resources.ResourceArgs {

    public static final AlertConfigurationSeverityConfigurationArgs Empty = new AlertConfigurationSeverityConfigurationArgs();

    @Import(name="evalCondition")
    private @Nullable Output<AlertConfigurationSeverityConfigurationEvalConditionArgs> evalCondition;

    public Optional<Output<AlertConfigurationSeverityConfigurationEvalConditionArgs>> evalCondition() {
        return Optional.ofNullable(this.evalCondition);
    }

    @Import(name="severity")
    private @Nullable Output<Integer> severity;

    public Optional<Output<Integer>> severity() {
        return Optional.ofNullable(this.severity);
    }

    private AlertConfigurationSeverityConfigurationArgs() {}

    private AlertConfigurationSeverityConfigurationArgs(AlertConfigurationSeverityConfigurationArgs $) {
        this.evalCondition = $.evalCondition;
        this.severity = $.severity;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(AlertConfigurationSeverityConfigurationArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private AlertConfigurationSeverityConfigurationArgs $;

        public Builder() {
            $ = new AlertConfigurationSeverityConfigurationArgs();
        }

        public Builder(AlertConfigurationSeverityConfigurationArgs defaults) {
            $ = new AlertConfigurationSeverityConfigurationArgs(Objects.requireNonNull(defaults));
        }

        public Builder evalCondition(@Nullable Output<AlertConfigurationSeverityConfigurationEvalConditionArgs> evalCondition) {
            $.evalCondition = evalCondition;
            return this;
        }

        public Builder evalCondition(AlertConfigurationSeverityConfigurationEvalConditionArgs evalCondition) {
            return evalCondition(Output.of(evalCondition));
        }

        public Builder severity(@Nullable Output<Integer> severity) {
            $.severity = severity;
            return this;
        }

        public Builder severity(Integer severity) {
            return severity(Output.of(severity));
        }

        public AlertConfigurationSeverityConfigurationArgs build() {
            return $;
        }
    }

}
