// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.sls.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class AlertConfigurationConditionConfigurationArgs extends com.pulumi.resources.ResourceArgs {

    public static final AlertConfigurationConditionConfigurationArgs Empty = new AlertConfigurationConditionConfigurationArgs();

    @Import(name="condition")
    private @Nullable Output<String> condition;

    public Optional<Output<String>> condition() {
        return Optional.ofNullable(this.condition);
    }

    @Import(name="countCondition")
    private @Nullable Output<String> countCondition;

    public Optional<Output<String>> countCondition() {
        return Optional.ofNullable(this.countCondition);
    }

    private AlertConfigurationConditionConfigurationArgs() {}

    private AlertConfigurationConditionConfigurationArgs(AlertConfigurationConditionConfigurationArgs $) {
        this.condition = $.condition;
        this.countCondition = $.countCondition;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(AlertConfigurationConditionConfigurationArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private AlertConfigurationConditionConfigurationArgs $;

        public Builder() {
            $ = new AlertConfigurationConditionConfigurationArgs();
        }

        public Builder(AlertConfigurationConditionConfigurationArgs defaults) {
            $ = new AlertConfigurationConditionConfigurationArgs(Objects.requireNonNull(defaults));
        }

        public Builder condition(@Nullable Output<String> condition) {
            $.condition = condition;
            return this;
        }

        public Builder condition(String condition) {
            return condition(Output.of(condition));
        }

        public Builder countCondition(@Nullable Output<String> countCondition) {
            $.countCondition = countCondition;
            return this;
        }

        public Builder countCondition(String countCondition) {
            return countCondition(Output.of(countCondition));
        }

        public AlertConfigurationConditionConfigurationArgs build() {
            return $;
        }
    }

}
