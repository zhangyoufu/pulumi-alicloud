// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.amqp.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class BindingState extends com.pulumi.resources.ResourceArgs {

    public static final BindingState Empty = new BindingState();

    /**
     * X-match Attributes. Valid Values:
     * * &#34;x-match:all&#34;: Default Value, All the Message Header of Key-Value Pairs Stored in the Must Match.
     * * &#34;x-match:any&#34;: at Least One Pair of the Message Header of Key-Value Pairs Stored in the Must Match.
     * 
     * **NOTE:** This Parameter Applies Only to Headers Exchange Other Types of Exchange Is Invalid. Other Types of Exchange Here Can Either Be an Arbitrary Value.
     * 
     */
    @Import(name="argument")
    private @Nullable Output<String> argument;

    /**
     * @return X-match Attributes. Valid Values:
     * * &#34;x-match:all&#34;: Default Value, All the Message Header of Key-Value Pairs Stored in the Must Match.
     * * &#34;x-match:any&#34;: at Least One Pair of the Message Header of Key-Value Pairs Stored in the Must Match.
     * 
     * **NOTE:** This Parameter Applies Only to Headers Exchange Other Types of Exchange Is Invalid. Other Types of Exchange Here Can Either Be an Arbitrary Value.
     * 
     */
    public Optional<Output<String>> argument() {
        return Optional.ofNullable(this.argument);
    }

    /**
     * The Binding Key.
     * * For a non-topic source exchange: The binding key can contain only letters, digits, hyphens (-), underscores (_), periods (.), and at signs (@).
     *   The binding key must be 1 to 255 characters in length.
     * * For a topic source exchange: The binding key can contain letters, digits, hyphens (-), underscores (_), periods (.), and at signs (@).
     *   If the binding key contains a number sign (#), the binding key must start with a number sign (#) followed by a period (.) or end with a number sign (#) that follows a period (.).
     *   The binding key must be 1 to 255 characters in length.
     * 
     */
    @Import(name="bindingKey")
    private @Nullable Output<String> bindingKey;

    /**
     * @return The Binding Key.
     * * For a non-topic source exchange: The binding key can contain only letters, digits, hyphens (-), underscores (_), periods (.), and at signs (@).
     *   The binding key must be 1 to 255 characters in length.
     * * For a topic source exchange: The binding key can contain letters, digits, hyphens (-), underscores (_), periods (.), and at signs (@).
     *   If the binding key contains a number sign (#), the binding key must start with a number sign (#) followed by a period (.) or end with a number sign (#) that follows a period (.).
     *   The binding key must be 1 to 255 characters in length.
     * 
     */
    public Optional<Output<String>> bindingKey() {
        return Optional.ofNullable(this.bindingKey);
    }

    /**
     * The Target Binding Types. Valid values: `EXCHANGE`, `QUEUE`.
     * 
     */
    @Import(name="bindingType")
    private @Nullable Output<String> bindingType;

    /**
     * @return The Target Binding Types. Valid values: `EXCHANGE`, `QUEUE`.
     * 
     */
    public Optional<Output<String>> bindingType() {
        return Optional.ofNullable(this.bindingType);
    }

    /**
     * The Target Queue Or Exchange of the Name.
     * 
     */
    @Import(name="destinationName")
    private @Nullable Output<String> destinationName;

    /**
     * @return The Target Queue Or Exchange of the Name.
     * 
     */
    public Optional<Output<String>> destinationName() {
        return Optional.ofNullable(this.destinationName);
    }

    /**
     * Instance Id.
     * 
     */
    @Import(name="instanceId")
    private @Nullable Output<String> instanceId;

    /**
     * @return Instance Id.
     * 
     */
    public Optional<Output<String>> instanceId() {
        return Optional.ofNullable(this.instanceId);
    }

    /**
     * The Source Exchange Name.
     * 
     */
    @Import(name="sourceExchange")
    private @Nullable Output<String> sourceExchange;

    /**
     * @return The Source Exchange Name.
     * 
     */
    public Optional<Output<String>> sourceExchange() {
        return Optional.ofNullable(this.sourceExchange);
    }

    /**
     * Virtualhost Name.
     * 
     */
    @Import(name="virtualHostName")
    private @Nullable Output<String> virtualHostName;

    /**
     * @return Virtualhost Name.
     * 
     */
    public Optional<Output<String>> virtualHostName() {
        return Optional.ofNullable(this.virtualHostName);
    }

    private BindingState() {}

    private BindingState(BindingState $) {
        this.argument = $.argument;
        this.bindingKey = $.bindingKey;
        this.bindingType = $.bindingType;
        this.destinationName = $.destinationName;
        this.instanceId = $.instanceId;
        this.sourceExchange = $.sourceExchange;
        this.virtualHostName = $.virtualHostName;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(BindingState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private BindingState $;

        public Builder() {
            $ = new BindingState();
        }

        public Builder(BindingState defaults) {
            $ = new BindingState(Objects.requireNonNull(defaults));
        }

        /**
         * @param argument X-match Attributes. Valid Values:
         * * &#34;x-match:all&#34;: Default Value, All the Message Header of Key-Value Pairs Stored in the Must Match.
         * * &#34;x-match:any&#34;: at Least One Pair of the Message Header of Key-Value Pairs Stored in the Must Match.
         * 
         * **NOTE:** This Parameter Applies Only to Headers Exchange Other Types of Exchange Is Invalid. Other Types of Exchange Here Can Either Be an Arbitrary Value.
         * 
         * @return builder
         * 
         */
        public Builder argument(@Nullable Output<String> argument) {
            $.argument = argument;
            return this;
        }

        /**
         * @param argument X-match Attributes. Valid Values:
         * * &#34;x-match:all&#34;: Default Value, All the Message Header of Key-Value Pairs Stored in the Must Match.
         * * &#34;x-match:any&#34;: at Least One Pair of the Message Header of Key-Value Pairs Stored in the Must Match.
         * 
         * **NOTE:** This Parameter Applies Only to Headers Exchange Other Types of Exchange Is Invalid. Other Types of Exchange Here Can Either Be an Arbitrary Value.
         * 
         * @return builder
         * 
         */
        public Builder argument(String argument) {
            return argument(Output.of(argument));
        }

        /**
         * @param bindingKey The Binding Key.
         * * For a non-topic source exchange: The binding key can contain only letters, digits, hyphens (-), underscores (_), periods (.), and at signs (@).
         *   The binding key must be 1 to 255 characters in length.
         * * For a topic source exchange: The binding key can contain letters, digits, hyphens (-), underscores (_), periods (.), and at signs (@).
         *   If the binding key contains a number sign (#), the binding key must start with a number sign (#) followed by a period (.) or end with a number sign (#) that follows a period (.).
         *   The binding key must be 1 to 255 characters in length.
         * 
         * @return builder
         * 
         */
        public Builder bindingKey(@Nullable Output<String> bindingKey) {
            $.bindingKey = bindingKey;
            return this;
        }

        /**
         * @param bindingKey The Binding Key.
         * * For a non-topic source exchange: The binding key can contain only letters, digits, hyphens (-), underscores (_), periods (.), and at signs (@).
         *   The binding key must be 1 to 255 characters in length.
         * * For a topic source exchange: The binding key can contain letters, digits, hyphens (-), underscores (_), periods (.), and at signs (@).
         *   If the binding key contains a number sign (#), the binding key must start with a number sign (#) followed by a period (.) or end with a number sign (#) that follows a period (.).
         *   The binding key must be 1 to 255 characters in length.
         * 
         * @return builder
         * 
         */
        public Builder bindingKey(String bindingKey) {
            return bindingKey(Output.of(bindingKey));
        }

        /**
         * @param bindingType The Target Binding Types. Valid values: `EXCHANGE`, `QUEUE`.
         * 
         * @return builder
         * 
         */
        public Builder bindingType(@Nullable Output<String> bindingType) {
            $.bindingType = bindingType;
            return this;
        }

        /**
         * @param bindingType The Target Binding Types. Valid values: `EXCHANGE`, `QUEUE`.
         * 
         * @return builder
         * 
         */
        public Builder bindingType(String bindingType) {
            return bindingType(Output.of(bindingType));
        }

        /**
         * @param destinationName The Target Queue Or Exchange of the Name.
         * 
         * @return builder
         * 
         */
        public Builder destinationName(@Nullable Output<String> destinationName) {
            $.destinationName = destinationName;
            return this;
        }

        /**
         * @param destinationName The Target Queue Or Exchange of the Name.
         * 
         * @return builder
         * 
         */
        public Builder destinationName(String destinationName) {
            return destinationName(Output.of(destinationName));
        }

        /**
         * @param instanceId Instance Id.
         * 
         * @return builder
         * 
         */
        public Builder instanceId(@Nullable Output<String> instanceId) {
            $.instanceId = instanceId;
            return this;
        }

        /**
         * @param instanceId Instance Id.
         * 
         * @return builder
         * 
         */
        public Builder instanceId(String instanceId) {
            return instanceId(Output.of(instanceId));
        }

        /**
         * @param sourceExchange The Source Exchange Name.
         * 
         * @return builder
         * 
         */
        public Builder sourceExchange(@Nullable Output<String> sourceExchange) {
            $.sourceExchange = sourceExchange;
            return this;
        }

        /**
         * @param sourceExchange The Source Exchange Name.
         * 
         * @return builder
         * 
         */
        public Builder sourceExchange(String sourceExchange) {
            return sourceExchange(Output.of(sourceExchange));
        }

        /**
         * @param virtualHostName Virtualhost Name.
         * 
         * @return builder
         * 
         */
        public Builder virtualHostName(@Nullable Output<String> virtualHostName) {
            $.virtualHostName = virtualHostName;
            return this;
        }

        /**
         * @param virtualHostName Virtualhost Name.
         * 
         * @return builder
         * 
         */
        public Builder virtualHostName(String virtualHostName) {
            return virtualHostName(Output.of(virtualHostName));
        }

        public BindingState build() {
            return $;
        }
    }

}
