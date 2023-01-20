// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Ess.Outputs
{

    [OutputType]
    public sealed class EciScalingConfigurationContainer
    {
        /// <summary>
        /// The arguments passed to the commands.
        /// </summary>
        public readonly ImmutableArray<string> Args;
        /// <summary>
        /// The commands run by the init container.
        /// </summary>
        public readonly ImmutableArray<string> Commands;
        /// <summary>
        /// The amount of CPU resources allocated to the container group.
        /// </summary>
        public readonly double? Cpu;
        /// <summary>
        /// The structure of environmentVars.
        /// See Block_environment_var_in_init_container below for details.
        /// See Block_environment_var_in_container below for details.
        /// </summary>
        public readonly ImmutableArray<Outputs.EciScalingConfigurationContainerEnvironmentVar> EnvironmentVars;
        /// <summary>
        /// The number GPUs.
        /// </summary>
        public readonly int? Gpu;
        /// <summary>
        /// The image of the container.
        /// </summary>
        public readonly string? Image;
        /// <summary>
        /// The restart policy of the image.
        /// </summary>
        public readonly string? ImagePullPolicy;
        /// <summary>
        /// Commands that you want to run in containers when you use the CLI to perform liveness probes.
        /// </summary>
        public readonly ImmutableArray<string> LivenessProbeExecCommands;
        /// <summary>
        /// The minimum number of consecutive failures for the liveness probe to be considered failed after having been successful. Default value: 3.
        /// </summary>
        public readonly int? LivenessProbeFailureThreshold;
        /// <summary>
        /// The path to which HTTP GET requests are sent when you use HTTP requests to perform liveness probes.
        /// </summary>
        public readonly string? LivenessProbeHttpGetPath;
        /// <summary>
        /// The port to which HTTP GET requests are sent when you use HTTP requests to perform liveness probes.
        /// </summary>
        public readonly int? LivenessProbeHttpGetPort;
        /// <summary>
        /// The protocol type of HTTP GET requests when you use HTTP requests for liveness probes.Valid values:HTTP and HTTPS.
        /// </summary>
        public readonly string? LivenessProbeHttpGetScheme;
        /// <summary>
        /// The number of seconds after container has started before liveness probes are initiated.
        /// </summary>
        public readonly int? LivenessProbeInitialDelaySeconds;
        /// <summary>
        /// The interval at which the liveness probe is performed. Unit: seconds. Default value: 10. Minimum value: 1.
        /// </summary>
        public readonly int? LivenessProbePeriodSeconds;
        /// <summary>
        /// The minimum number of consecutive successes for the liveness probe to be considered successful after having failed. Default value: 1. Set the value to 1.
        /// </summary>
        public readonly int? LivenessProbeSuccessThreshold;
        /// <summary>
        /// The port detected by TCP sockets when you use TCP sockets to perform liveness probes.
        /// </summary>
        public readonly int? LivenessProbeTcpSocketPort;
        /// <summary>
        /// The timeout period for the liveness probe. Unit: seconds. Default value: 1. Minimum value: 1.
        /// </summary>
        public readonly int? LivenessProbeTimeoutSeconds;
        /// <summary>
        /// The amount of memory resources allocated to the container group.
        /// </summary>
        public readonly double? Memory;
        /// <summary>
        /// The name of the volume.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// The structure of port. See Block_port_in_init_container below
        /// for details.
        /// </summary>
        public readonly ImmutableArray<Outputs.EciScalingConfigurationContainerPort> Ports;
        /// <summary>
        /// Commands that you want to run in containers when you use the CLI to perform readiness probes.
        /// </summary>
        public readonly ImmutableArray<string> ReadinessProbeExecCommands;
        /// <summary>
        /// The minimum number of consecutive failures for the readiness probe to be considered failed after having been successful. Default value: 3.
        /// </summary>
        public readonly int? ReadinessProbeFailureThreshold;
        /// <summary>
        /// The path to which HTTP GET requests are sent when you use HTTP requests to perform readiness probes.
        /// </summary>
        public readonly string? ReadinessProbeHttpGetPath;
        /// <summary>
        /// The port to which HTTP GET requests are sent when you use HTTP requests to perform readiness probes.
        /// </summary>
        public readonly int? ReadinessProbeHttpGetPort;
        /// <summary>
        /// The protocol type of HTTP GET requests when you use HTTP requests for readiness probes. Valid values: HTTP and HTTPS.
        /// </summary>
        public readonly string? ReadinessProbeHttpGetScheme;
        /// <summary>
        /// The number of seconds after container N has started before readiness probes are initiated.
        /// </summary>
        public readonly int? ReadinessProbeInitialDelaySeconds;
        /// <summary>
        /// The interval at which the readiness probe is performed. Unit: seconds. Default value: 10. Minimum value: 1.
        /// </summary>
        public readonly int? ReadinessProbePeriodSeconds;
        /// <summary>
        /// The minimum number of consecutive successes for the readiness probe to be considered successful after having failed. Default value: 1. Set the value to 1.
        /// </summary>
        public readonly int? ReadinessProbeSuccessThreshold;
        /// <summary>
        /// The port detected by Transmission Control Protocol (TCP) sockets when you use TCP sockets to perform readiness probes.
        /// </summary>
        public readonly int? ReadinessProbeTcpSocketPort;
        /// <summary>
        /// The timeout period for the readiness probe. Unit: seconds. Default value: 1. Minimum value: 1.
        /// </summary>
        public readonly int? ReadinessProbeTimeoutSeconds;
        /// <summary>
        /// The structure of volumeMounts.
        /// See Block_volume_mount_in_init_container below for details.
        /// See Block_volume_mount_in_container below for details.
        /// </summary>
        public readonly ImmutableArray<Outputs.EciScalingConfigurationContainerVolumeMount> VolumeMounts;
        /// <summary>
        /// The working directory of the container.
        /// </summary>
        public readonly string? WorkingDir;

        [OutputConstructor]
        private EciScalingConfigurationContainer(
            ImmutableArray<string> args,

            ImmutableArray<string> commands,

            double? cpu,

            ImmutableArray<Outputs.EciScalingConfigurationContainerEnvironmentVar> environmentVars,

            int? gpu,

            string? image,

            string? imagePullPolicy,

            ImmutableArray<string> livenessProbeExecCommands,

            int? livenessProbeFailureThreshold,

            string? livenessProbeHttpGetPath,

            int? livenessProbeHttpGetPort,

            string? livenessProbeHttpGetScheme,

            int? livenessProbeInitialDelaySeconds,

            int? livenessProbePeriodSeconds,

            int? livenessProbeSuccessThreshold,

            int? livenessProbeTcpSocketPort,

            int? livenessProbeTimeoutSeconds,

            double? memory,

            string? name,

            ImmutableArray<Outputs.EciScalingConfigurationContainerPort> ports,

            ImmutableArray<string> readinessProbeExecCommands,

            int? readinessProbeFailureThreshold,

            string? readinessProbeHttpGetPath,

            int? readinessProbeHttpGetPort,

            string? readinessProbeHttpGetScheme,

            int? readinessProbeInitialDelaySeconds,

            int? readinessProbePeriodSeconds,

            int? readinessProbeSuccessThreshold,

            int? readinessProbeTcpSocketPort,

            int? readinessProbeTimeoutSeconds,

            ImmutableArray<Outputs.EciScalingConfigurationContainerVolumeMount> volumeMounts,

            string? workingDir)
        {
            Args = args;
            Commands = commands;
            Cpu = cpu;
            EnvironmentVars = environmentVars;
            Gpu = gpu;
            Image = image;
            ImagePullPolicy = imagePullPolicy;
            LivenessProbeExecCommands = livenessProbeExecCommands;
            LivenessProbeFailureThreshold = livenessProbeFailureThreshold;
            LivenessProbeHttpGetPath = livenessProbeHttpGetPath;
            LivenessProbeHttpGetPort = livenessProbeHttpGetPort;
            LivenessProbeHttpGetScheme = livenessProbeHttpGetScheme;
            LivenessProbeInitialDelaySeconds = livenessProbeInitialDelaySeconds;
            LivenessProbePeriodSeconds = livenessProbePeriodSeconds;
            LivenessProbeSuccessThreshold = livenessProbeSuccessThreshold;
            LivenessProbeTcpSocketPort = livenessProbeTcpSocketPort;
            LivenessProbeTimeoutSeconds = livenessProbeTimeoutSeconds;
            Memory = memory;
            Name = name;
            Ports = ports;
            ReadinessProbeExecCommands = readinessProbeExecCommands;
            ReadinessProbeFailureThreshold = readinessProbeFailureThreshold;
            ReadinessProbeHttpGetPath = readinessProbeHttpGetPath;
            ReadinessProbeHttpGetPort = readinessProbeHttpGetPort;
            ReadinessProbeHttpGetScheme = readinessProbeHttpGetScheme;
            ReadinessProbeInitialDelaySeconds = readinessProbeInitialDelaySeconds;
            ReadinessProbePeriodSeconds = readinessProbePeriodSeconds;
            ReadinessProbeSuccessThreshold = readinessProbeSuccessThreshold;
            ReadinessProbeTcpSocketPort = readinessProbeTcpSocketPort;
            ReadinessProbeTimeoutSeconds = readinessProbeTimeoutSeconds;
            VolumeMounts = volumeMounts;
            WorkingDir = workingDir;
        }
    }
}
