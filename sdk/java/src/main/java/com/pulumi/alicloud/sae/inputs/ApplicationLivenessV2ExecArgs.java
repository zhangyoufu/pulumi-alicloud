// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.sae.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ApplicationLivenessV2ExecArgs extends com.pulumi.resources.ResourceArgs {

    public static final ApplicationLivenessV2ExecArgs Empty = new ApplicationLivenessV2ExecArgs();

    /**
     * Mirror start command. The command must be an executable object in the container. For example: sleep. Setting this command will cause the original startup command of the mirror to become invalid.
     * 
     */
    @Import(name="commands")
    private @Nullable Output<List<String>> commands;

    /**
     * @return Mirror start command. The command must be an executable object in the container. For example: sleep. Setting this command will cause the original startup command of the mirror to become invalid.
     * 
     */
    public Optional<Output<List<String>>> commands() {
        return Optional.ofNullable(this.commands);
    }

    private ApplicationLivenessV2ExecArgs() {}

    private ApplicationLivenessV2ExecArgs(ApplicationLivenessV2ExecArgs $) {
        this.commands = $.commands;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ApplicationLivenessV2ExecArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ApplicationLivenessV2ExecArgs $;

        public Builder() {
            $ = new ApplicationLivenessV2ExecArgs();
        }

        public Builder(ApplicationLivenessV2ExecArgs defaults) {
            $ = new ApplicationLivenessV2ExecArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param commands Mirror start command. The command must be an executable object in the container. For example: sleep. Setting this command will cause the original startup command of the mirror to become invalid.
         * 
         * @return builder
         * 
         */
        public Builder commands(@Nullable Output<List<String>> commands) {
            $.commands = commands;
            return this;
        }

        /**
         * @param commands Mirror start command. The command must be an executable object in the container. For example: sleep. Setting this command will cause the original startup command of the mirror to become invalid.
         * 
         * @return builder
         * 
         */
        public Builder commands(List<String> commands) {
            return commands(Output.of(commands));
        }

        /**
         * @param commands Mirror start command. The command must be an executable object in the container. For example: sleep. Setting this command will cause the original startup command of the mirror to become invalid.
         * 
         * @return builder
         * 
         */
        public Builder commands(String... commands) {
            return commands(List.of(commands));
        }

        public ApplicationLivenessV2ExecArgs build() {
            return $;
        }
    }

}
