// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.alb.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ListenerAccessLogTracingConfigArgs extends com.pulumi.resources.ResourceArgs {

    public static final ListenerAccessLogTracingConfigArgs Empty = new ListenerAccessLogTracingConfigArgs();

    /**
     * Xtrace Function. Value: `True` Or `False` . Default Value: `False`.
     * 
     * &gt; **NOTE:** Only Instances outside the Security Group to Access the Log Switch `accesslogenabled` Open, in Order to Set This Parameter to the `True`.
     * 
     */
    @Import(name="tracingEnabled")
    private @Nullable Output<Boolean> tracingEnabled;

    /**
     * @return Xtrace Function. Value: `True` Or `False` . Default Value: `False`.
     * 
     * &gt; **NOTE:** Only Instances outside the Security Group to Access the Log Switch `accesslogenabled` Open, in Order to Set This Parameter to the `True`.
     * 
     */
    public Optional<Output<Boolean>> tracingEnabled() {
        return Optional.ofNullable(this.tracingEnabled);
    }

    /**
     * Xtrace Sampling Rate. Value: `1` to `10000`.
     * 
     * &gt; **NOTE:** This attribute is valid when `tracingenabled` is `true`.
     * 
     */
    @Import(name="tracingSample")
    private @Nullable Output<Integer> tracingSample;

    /**
     * @return Xtrace Sampling Rate. Value: `1` to `10000`.
     * 
     * &gt; **NOTE:** This attribute is valid when `tracingenabled` is `true`.
     * 
     */
    public Optional<Output<Integer>> tracingSample() {
        return Optional.ofNullable(this.tracingSample);
    }

    /**
     * Xtrace Type Value Is `Zipkin`.
     * 
     * &gt; **NOTE:** This attribute is valid when `tracingenabled` is `true`.
     * 
     */
    @Import(name="tracingType")
    private @Nullable Output<String> tracingType;

    /**
     * @return Xtrace Type Value Is `Zipkin`.
     * 
     * &gt; **NOTE:** This attribute is valid when `tracingenabled` is `true`.
     * 
     */
    public Optional<Output<String>> tracingType() {
        return Optional.ofNullable(this.tracingType);
    }

    private ListenerAccessLogTracingConfigArgs() {}

    private ListenerAccessLogTracingConfigArgs(ListenerAccessLogTracingConfigArgs $) {
        this.tracingEnabled = $.tracingEnabled;
        this.tracingSample = $.tracingSample;
        this.tracingType = $.tracingType;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ListenerAccessLogTracingConfigArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ListenerAccessLogTracingConfigArgs $;

        public Builder() {
            $ = new ListenerAccessLogTracingConfigArgs();
        }

        public Builder(ListenerAccessLogTracingConfigArgs defaults) {
            $ = new ListenerAccessLogTracingConfigArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param tracingEnabled Xtrace Function. Value: `True` Or `False` . Default Value: `False`.
         * 
         * &gt; **NOTE:** Only Instances outside the Security Group to Access the Log Switch `accesslogenabled` Open, in Order to Set This Parameter to the `True`.
         * 
         * @return builder
         * 
         */
        public Builder tracingEnabled(@Nullable Output<Boolean> tracingEnabled) {
            $.tracingEnabled = tracingEnabled;
            return this;
        }

        /**
         * @param tracingEnabled Xtrace Function. Value: `True` Or `False` . Default Value: `False`.
         * 
         * &gt; **NOTE:** Only Instances outside the Security Group to Access the Log Switch `accesslogenabled` Open, in Order to Set This Parameter to the `True`.
         * 
         * @return builder
         * 
         */
        public Builder tracingEnabled(Boolean tracingEnabled) {
            return tracingEnabled(Output.of(tracingEnabled));
        }

        /**
         * @param tracingSample Xtrace Sampling Rate. Value: `1` to `10000`.
         * 
         * &gt; **NOTE:** This attribute is valid when `tracingenabled` is `true`.
         * 
         * @return builder
         * 
         */
        public Builder tracingSample(@Nullable Output<Integer> tracingSample) {
            $.tracingSample = tracingSample;
            return this;
        }

        /**
         * @param tracingSample Xtrace Sampling Rate. Value: `1` to `10000`.
         * 
         * &gt; **NOTE:** This attribute is valid when `tracingenabled` is `true`.
         * 
         * @return builder
         * 
         */
        public Builder tracingSample(Integer tracingSample) {
            return tracingSample(Output.of(tracingSample));
        }

        /**
         * @param tracingType Xtrace Type Value Is `Zipkin`.
         * 
         * &gt; **NOTE:** This attribute is valid when `tracingenabled` is `true`.
         * 
         * @return builder
         * 
         */
        public Builder tracingType(@Nullable Output<String> tracingType) {
            $.tracingType = tracingType;
            return this;
        }

        /**
         * @param tracingType Xtrace Type Value Is `Zipkin`.
         * 
         * &gt; **NOTE:** This attribute is valid when `tracingenabled` is `true`.
         * 
         * @return builder
         * 
         */
        public Builder tracingType(String tracingType) {
            return tracingType(Output.of(tracingType));
        }

        public ListenerAccessLogTracingConfigArgs build() {
            return $;
        }
    }

}
