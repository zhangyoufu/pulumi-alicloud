// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Eci.Outputs
{

    [OutputType]
    public sealed class ContainerGroupInitContainer
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
        /// The amount of CPU resources allocated to the container.
        /// </summary>
        public readonly double? Cpu;
        /// <summary>
        /// The structure of environmentVars.
        /// </summary>
        public readonly ImmutableArray<Outputs.ContainerGroupInitContainerEnvironmentVar> EnvironmentVars;
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
        /// The amount of memory resources allocated to the container.
        /// </summary>
        public readonly double? Memory;
        /// <summary>
        /// The name of the mounted volume.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// The structure of port.
        /// </summary>
        public readonly ImmutableArray<Outputs.ContainerGroupInitContainerPort> Ports;
        public readonly bool? Ready;
        public readonly int? RestartCount;
        /// <summary>
        /// The structure of volumeMounts.
        /// </summary>
        public readonly ImmutableArray<Outputs.ContainerGroupInitContainerVolumeMount> VolumeMounts;
        /// <summary>
        /// The working directory of the container.
        /// </summary>
        public readonly string? WorkingDir;

        [OutputConstructor]
        private ContainerGroupInitContainer(
            ImmutableArray<string> args,

            ImmutableArray<string> commands,

            double? cpu,

            ImmutableArray<Outputs.ContainerGroupInitContainerEnvironmentVar> environmentVars,

            int? gpu,

            string? image,

            string? imagePullPolicy,

            double? memory,

            string? name,

            ImmutableArray<Outputs.ContainerGroupInitContainerPort> ports,

            bool? ready,

            int? restartCount,

            ImmutableArray<Outputs.ContainerGroupInitContainerVolumeMount> volumeMounts,

            string? workingDir)
        {
            Args = args;
            Commands = commands;
            Cpu = cpu;
            EnvironmentVars = environmentVars;
            Gpu = gpu;
            Image = image;
            ImagePullPolicy = imagePullPolicy;
            Memory = memory;
            Name = name;
            Ports = ports;
            Ready = ready;
            RestartCount = restartCount;
            VolumeMounts = volumeMounts;
            WorkingDir = workingDir;
        }
    }
}
