{{/* vim: set filetype=mustache: */}}
{{/*
Expand the name of the chart.
*/}}
{{- define "wallet-logs.name" -}}
{{- default .Chart.Name .Values.nameOverride | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Create a default fully qualified app name.
We truncate at 63 chars because some Kubernetes name fields are limited to this (by the DNS naming spec).
If release name contains chart name it will be used as a full name.
*/}}
{{- define "wallet-logs.fullname" -}}
{{- if .Values.fullnameOverride }}
{{- .Values.fullnameOverride | trunc 63 | trimSuffix "-" }}
{{- else }}
{{- $name := default .Chart.Name .Values.nameOverride }}
{{- if contains $name .Release.Name }}
{{- .Release.Name | trunc 63 | trimSuffix "-" }}
{{- else }}
{{- printf "%s-%s" .Release.Name $name | trunc 63 | trimSuffix "-" }}
{{- end }}
{{- end }}
{{- end }}

{{/*
Create chart name and version as used by the chart label.
*/}}
{{- define "wallet-logs.chart" -}}
{{- printf "%s-%s" .Chart.Name .Chart.Version | replace "+" "_" | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Common labels
*/}}
{{- define "wallet-logs.labels" -}}
helm.sh/chart: {{ include "wallet-logs.chart" . }}
{{ include "wallet-logs.selectorLabels" . }}
{{- if .Chart.AppVersion }}
app.kubernetes.io/version: {{ .Chart.AppVersion | quote }}
{{- end }}
app.kubernetes.io/managed-by: {{ .Release.Service }}
{{- end }}

{{/*
Selector labels
*/}}
{{- define "wallet-logs.selectorLabels" -}}
app.kubernetes.io/name: {{ include "wallet-logs.name" . }}
app.kubernetes.io/instance: {{ .Release.Name }}
{{- end }}

{{/*
Create the name of the service account to use
*/}}
{{- define "wallet-logs.serviceAccountName" -}}
{{- if .Values.serviceAccount.create }}
{{- default (include "wallet-logs.fullname" .) .Values.serviceAccount.name }}
{{- else }}
{{- default "default" .Values.serviceAccount.name }}
{{- end }}
{{- end }}

{{/*
Create tag name of the image
*/}}
{{- define "wallet-logs.imageTag" -}}
{{ .Values.image.tag | default .Chart.AppVersion }}
{{- end }}

{{/*
Create the name of the image repository
*/}}
{{- define "wallet-logs.imageRepository" -}}
{{ .Values.image.repository | default (printf "velmie/%s" .Chart.Name) }}
{{- end }}

{{/*
Create full image repository name including tag
*/}}
{{- define "wallet-logs.imageRepositoryWithTag" -}}
{{ include "wallet-logs.imageRepository" . }}:{{ include "wallet-logs.imageTag" . }}
{{- end }}

{{/*
Create full database migration image repository name
*/}}
{{- define "wallet-logs.dbMigrationImageRepositoryWithTag" -}}
{{ include "wallet-logs.imageRepository" . }}-db-migration:{{ include "wallet-logs.imageTag" . }}
{{- end }}