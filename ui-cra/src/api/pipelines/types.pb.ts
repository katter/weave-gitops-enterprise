/* eslint-disable */
// @ts-nocheck
/*
* This file is a generated Typescript file for GRPC Gateway, DO NOT MODIFY
*/
export type ClusterRef = {
  kind?: string
  name?: string
  namespace?: string
}

export type Target = {
  namespace?: string
  clusterRef?: ClusterRef
}

export type Environment = {
  name?: string
  targets?: Target[]
}

export type ObjectRef = {
  kind?: string
  name?: string
  namespace?: string
}

export type AppRef = {
  apiVersion?: string
  kind?: string
  name?: string
}

export type WorkloadStatus = {
  kind?: string
  name?: string
  version?: string
}

export type PipelineTargetStatus = {
  clusterRef?: ClusterRef
  namespace?: string
  workloads?: WorkloadStatus[]
}

export type PipelineStatusTargetStatusList = {
  targetsStatuses?: PipelineTargetStatus[]
}

export type PipelineStatus = {
  environments?: {[key: string]: PipelineStatusTargetStatusList}
}

export type Pipeline = {
  name?: string
  namespace?: string
  appRef?: AppRef
  environments?: Environment[]
  targets?: Target[]
  status?: PipelineStatus
}