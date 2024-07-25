// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.eci.inputs;

import com.pulumi.alicloud.eci.inputs.ContainerGroupContainerReadinessProbeExecArgs;
import com.pulumi.alicloud.eci.inputs.ContainerGroupContainerReadinessProbeHttpGetArgs;
import com.pulumi.alicloud.eci.inputs.ContainerGroupContainerReadinessProbeTcpSocketArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ContainerGroupContainerReadinessProbeArgs extends com.pulumi.resources.ResourceArgs {

    public static final ContainerGroupContainerReadinessProbeArgs Empty = new ContainerGroupContainerReadinessProbeArgs();

    /**
     * Health check using command line method. See `exec` below.
     * 
     */
    @Import(name="execs")
    private @Nullable Output<List<ContainerGroupContainerReadinessProbeExecArgs>> execs;

    /**
     * @return Health check using command line method. See `exec` below.
     * 
     */
    public Optional<Output<List<ContainerGroupContainerReadinessProbeExecArgs>>> execs() {
        return Optional.ofNullable(this.execs);
    }

    /**
     * Threshold for the number of checks that are determined to have failed since the last successful check (must be consecutive failures), default is 3.
     * 
     */
    @Import(name="failureThreshold")
    private @Nullable Output<Integer> failureThreshold;

    /**
     * @return Threshold for the number of checks that are determined to have failed since the last successful check (must be consecutive failures), default is 3.
     * 
     */
    public Optional<Output<Integer>> failureThreshold() {
        return Optional.ofNullable(this.failureThreshold);
    }

    /**
     * Health check using HTTP request method. See `http_get` below.
     * 
     * &gt; **NOTE:** When you configure `readiness_probe`, you can select only one of the `exec`, `tcp_socket`, `http_get`.
     * 
     */
    @Import(name="httpGets")
    private @Nullable Output<List<ContainerGroupContainerReadinessProbeHttpGetArgs>> httpGets;

    /**
     * @return Health check using HTTP request method. See `http_get` below.
     * 
     * &gt; **NOTE:** When you configure `readiness_probe`, you can select only one of the `exec`, `tcp_socket`, `http_get`.
     * 
     */
    public Optional<Output<List<ContainerGroupContainerReadinessProbeHttpGetArgs>>> httpGets() {
        return Optional.ofNullable(this.httpGets);
    }

    /**
     * Check the time to start execution, calculated from the completion of container startup.
     * 
     */
    @Import(name="initialDelaySeconds")
    private @Nullable Output<Integer> initialDelaySeconds;

    /**
     * @return Check the time to start execution, calculated from the completion of container startup.
     * 
     */
    public Optional<Output<Integer>> initialDelaySeconds() {
        return Optional.ofNullable(this.initialDelaySeconds);
    }

    /**
     * Buffer time for the program to handle operations before closing.
     * 
     */
    @Import(name="periodSeconds")
    private @Nullable Output<Integer> periodSeconds;

    /**
     * @return Buffer time for the program to handle operations before closing.
     * 
     */
    public Optional<Output<Integer>> periodSeconds() {
        return Optional.ofNullable(this.periodSeconds);
    }

    /**
     * The check count threshold for re-identifying successful checks since the last failed check (must be consecutive successes), default is 1. Current must be 1.
     * 
     */
    @Import(name="successThreshold")
    private @Nullable Output<Integer> successThreshold;

    /**
     * @return The check count threshold for re-identifying successful checks since the last failed check (must be consecutive successes), default is 1. Current must be 1.
     * 
     */
    public Optional<Output<Integer>> successThreshold() {
        return Optional.ofNullable(this.successThreshold);
    }

    /**
     * Health check using TCP socket method. See `tcp_socket` below.
     * 
     */
    @Import(name="tcpSockets")
    private @Nullable Output<List<ContainerGroupContainerReadinessProbeTcpSocketArgs>> tcpSockets;

    /**
     * @return Health check using TCP socket method. See `tcp_socket` below.
     * 
     */
    public Optional<Output<List<ContainerGroupContainerReadinessProbeTcpSocketArgs>>> tcpSockets() {
        return Optional.ofNullable(this.tcpSockets);
    }

    /**
     * Check the timeout, the default is 1 second, the minimum is 1 second.
     * 
     */
    @Import(name="timeoutSeconds")
    private @Nullable Output<Integer> timeoutSeconds;

    /**
     * @return Check the timeout, the default is 1 second, the minimum is 1 second.
     * 
     */
    public Optional<Output<Integer>> timeoutSeconds() {
        return Optional.ofNullable(this.timeoutSeconds);
    }

    private ContainerGroupContainerReadinessProbeArgs() {}

    private ContainerGroupContainerReadinessProbeArgs(ContainerGroupContainerReadinessProbeArgs $) {
        this.execs = $.execs;
        this.failureThreshold = $.failureThreshold;
        this.httpGets = $.httpGets;
        this.initialDelaySeconds = $.initialDelaySeconds;
        this.periodSeconds = $.periodSeconds;
        this.successThreshold = $.successThreshold;
        this.tcpSockets = $.tcpSockets;
        this.timeoutSeconds = $.timeoutSeconds;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ContainerGroupContainerReadinessProbeArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ContainerGroupContainerReadinessProbeArgs $;

        public Builder() {
            $ = new ContainerGroupContainerReadinessProbeArgs();
        }

        public Builder(ContainerGroupContainerReadinessProbeArgs defaults) {
            $ = new ContainerGroupContainerReadinessProbeArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param execs Health check using command line method. See `exec` below.
         * 
         * @return builder
         * 
         */
        public Builder execs(@Nullable Output<List<ContainerGroupContainerReadinessProbeExecArgs>> execs) {
            $.execs = execs;
            return this;
        }

        /**
         * @param execs Health check using command line method. See `exec` below.
         * 
         * @return builder
         * 
         */
        public Builder execs(List<ContainerGroupContainerReadinessProbeExecArgs> execs) {
            return execs(Output.of(execs));
        }

        /**
         * @param execs Health check using command line method. See `exec` below.
         * 
         * @return builder
         * 
         */
        public Builder execs(ContainerGroupContainerReadinessProbeExecArgs... execs) {
            return execs(List.of(execs));
        }

        /**
         * @param failureThreshold Threshold for the number of checks that are determined to have failed since the last successful check (must be consecutive failures), default is 3.
         * 
         * @return builder
         * 
         */
        public Builder failureThreshold(@Nullable Output<Integer> failureThreshold) {
            $.failureThreshold = failureThreshold;
            return this;
        }

        /**
         * @param failureThreshold Threshold for the number of checks that are determined to have failed since the last successful check (must be consecutive failures), default is 3.
         * 
         * @return builder
         * 
         */
        public Builder failureThreshold(Integer failureThreshold) {
            return failureThreshold(Output.of(failureThreshold));
        }

        /**
         * @param httpGets Health check using HTTP request method. See `http_get` below.
         * 
         * &gt; **NOTE:** When you configure `readiness_probe`, you can select only one of the `exec`, `tcp_socket`, `http_get`.
         * 
         * @return builder
         * 
         */
        public Builder httpGets(@Nullable Output<List<ContainerGroupContainerReadinessProbeHttpGetArgs>> httpGets) {
            $.httpGets = httpGets;
            return this;
        }

        /**
         * @param httpGets Health check using HTTP request method. See `http_get` below.
         * 
         * &gt; **NOTE:** When you configure `readiness_probe`, you can select only one of the `exec`, `tcp_socket`, `http_get`.
         * 
         * @return builder
         * 
         */
        public Builder httpGets(List<ContainerGroupContainerReadinessProbeHttpGetArgs> httpGets) {
            return httpGets(Output.of(httpGets));
        }

        /**
         * @param httpGets Health check using HTTP request method. See `http_get` below.
         * 
         * &gt; **NOTE:** When you configure `readiness_probe`, you can select only one of the `exec`, `tcp_socket`, `http_get`.
         * 
         * @return builder
         * 
         */
        public Builder httpGets(ContainerGroupContainerReadinessProbeHttpGetArgs... httpGets) {
            return httpGets(List.of(httpGets));
        }

        /**
         * @param initialDelaySeconds Check the time to start execution, calculated from the completion of container startup.
         * 
         * @return builder
         * 
         */
        public Builder initialDelaySeconds(@Nullable Output<Integer> initialDelaySeconds) {
            $.initialDelaySeconds = initialDelaySeconds;
            return this;
        }

        /**
         * @param initialDelaySeconds Check the time to start execution, calculated from the completion of container startup.
         * 
         * @return builder
         * 
         */
        public Builder initialDelaySeconds(Integer initialDelaySeconds) {
            return initialDelaySeconds(Output.of(initialDelaySeconds));
        }

        /**
         * @param periodSeconds Buffer time for the program to handle operations before closing.
         * 
         * @return builder
         * 
         */
        public Builder periodSeconds(@Nullable Output<Integer> periodSeconds) {
            $.periodSeconds = periodSeconds;
            return this;
        }

        /**
         * @param periodSeconds Buffer time for the program to handle operations before closing.
         * 
         * @return builder
         * 
         */
        public Builder periodSeconds(Integer periodSeconds) {
            return periodSeconds(Output.of(periodSeconds));
        }

        /**
         * @param successThreshold The check count threshold for re-identifying successful checks since the last failed check (must be consecutive successes), default is 1. Current must be 1.
         * 
         * @return builder
         * 
         */
        public Builder successThreshold(@Nullable Output<Integer> successThreshold) {
            $.successThreshold = successThreshold;
            return this;
        }

        /**
         * @param successThreshold The check count threshold for re-identifying successful checks since the last failed check (must be consecutive successes), default is 1. Current must be 1.
         * 
         * @return builder
         * 
         */
        public Builder successThreshold(Integer successThreshold) {
            return successThreshold(Output.of(successThreshold));
        }

        /**
         * @param tcpSockets Health check using TCP socket method. See `tcp_socket` below.
         * 
         * @return builder
         * 
         */
        public Builder tcpSockets(@Nullable Output<List<ContainerGroupContainerReadinessProbeTcpSocketArgs>> tcpSockets) {
            $.tcpSockets = tcpSockets;
            return this;
        }

        /**
         * @param tcpSockets Health check using TCP socket method. See `tcp_socket` below.
         * 
         * @return builder
         * 
         */
        public Builder tcpSockets(List<ContainerGroupContainerReadinessProbeTcpSocketArgs> tcpSockets) {
            return tcpSockets(Output.of(tcpSockets));
        }

        /**
         * @param tcpSockets Health check using TCP socket method. See `tcp_socket` below.
         * 
         * @return builder
         * 
         */
        public Builder tcpSockets(ContainerGroupContainerReadinessProbeTcpSocketArgs... tcpSockets) {
            return tcpSockets(List.of(tcpSockets));
        }

        /**
         * @param timeoutSeconds Check the timeout, the default is 1 second, the minimum is 1 second.
         * 
         * @return builder
         * 
         */
        public Builder timeoutSeconds(@Nullable Output<Integer> timeoutSeconds) {
            $.timeoutSeconds = timeoutSeconds;
            return this;
        }

        /**
         * @param timeoutSeconds Check the timeout, the default is 1 second, the minimum is 1 second.
         * 
         * @return builder
         * 
         */
        public Builder timeoutSeconds(Integer timeoutSeconds) {
            return timeoutSeconds(Output.of(timeoutSeconds));
        }

        public ContainerGroupContainerReadinessProbeArgs build() {
            return $;
        }
    }

}
