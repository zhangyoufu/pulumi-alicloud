// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.sae.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import javax.annotation.Nullable;

@CustomType
public final class ApplicationPostStartV2Exec {
    /**
     * @return Mirror start command. The command must be an executable object in the container. For example: sleep. Setting this command will cause the original startup command of the mirror to become invalid.
     * 
     */
    private @Nullable List<String> commands;

    private ApplicationPostStartV2Exec() {}
    /**
     * @return Mirror start command. The command must be an executable object in the container. For example: sleep. Setting this command will cause the original startup command of the mirror to become invalid.
     * 
     */
    public List<String> commands() {
        return this.commands == null ? List.of() : this.commands;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(ApplicationPostStartV2Exec defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable List<String> commands;
        public Builder() {}
        public Builder(ApplicationPostStartV2Exec defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.commands = defaults.commands;
        }

        @CustomType.Setter
        public Builder commands(@Nullable List<String> commands) {
            this.commands = commands;
            return this;
        }
        public Builder commands(String... commands) {
            return commands(List.of(commands));
        }
        public ApplicationPostStartV2Exec build() {
            final var o = new ApplicationPostStartV2Exec();
            o.commands = commands;
            return o;
        }
    }
}
