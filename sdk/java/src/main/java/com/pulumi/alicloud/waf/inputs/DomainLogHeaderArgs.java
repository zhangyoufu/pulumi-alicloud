// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.waf.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class DomainLogHeaderArgs extends com.pulumi.resources.ResourceArgs {

    public static final DomainLogHeaderArgs Empty = new DomainLogHeaderArgs();

    @Import(name="key")
    private @Nullable Output<String> key;

    public Optional<Output<String>> key() {
        return Optional.ofNullable(this.key);
    }

    @Import(name="value")
    private @Nullable Output<String> value;

    public Optional<Output<String>> value() {
        return Optional.ofNullable(this.value);
    }

    private DomainLogHeaderArgs() {}

    private DomainLogHeaderArgs(DomainLogHeaderArgs $) {
        this.key = $.key;
        this.value = $.value;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(DomainLogHeaderArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private DomainLogHeaderArgs $;

        public Builder() {
            $ = new DomainLogHeaderArgs();
        }

        public Builder(DomainLogHeaderArgs defaults) {
            $ = new DomainLogHeaderArgs(Objects.requireNonNull(defaults));
        }

        public Builder key(@Nullable Output<String> key) {
            $.key = key;
            return this;
        }

        public Builder key(String key) {
            return key(Output.of(key));
        }

        public Builder value(@Nullable Output<String> value) {
            $.value = value;
            return this;
        }

        public Builder value(String value) {
            return value(Output.of(value));
        }

        public DomainLogHeaderArgs build() {
            return $;
        }
    }

}
