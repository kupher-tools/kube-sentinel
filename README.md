# 🛡️ kube-sentinel

A Kubernetes **Mutating Admission Webhook** that enforces secure `securityContext` defaults for every Pod — protecting your workloads at the moment they are born.

---

## 📌 Description

**kube-sentinel** is a lightweight, high-performance admission controller that automatically injects secure defaults into every Pod's `securityContext`. It ensures all Pods in your cluster follow baseline security practices — even when developers forget or ignore them.

Built using `controller-runtime`, kube-sentinel acts as your cluster’s **first line of defense**, enforcing policies like `runAsNonRoot`, `readOnlyRootFilesystem`, and secure `seccompProfile` — without disrupting developer workflows.

---

## ✨ Key Features

✅ **Mutates Pods on Admission**  
Automatically enforces best practices on Pods as they’re created — no need for developers to specify security fields manually.

✅ **Enforces Security Defaults**
- `runAsNonRoot: true`
- `readOnlyRootFilesystem: true`
- `allowPrivilegeEscalation: false`
- `seccompProfile: RuntimeDefault`
- Default non-root UID/GID (optional)

✅ **Namespace Scoping**  
Supports namespace-based configuration or enforcement.

✅ **Auditable**  
Can label mutated Pods to support audit trails.

✅ **Works with Any Pod-Creating Resource**  
Also secures Pods created via Deployments, CronJobs, StatefulSets, etc.

✅ **Pluggable Architecture**  
Designed to extend beyond Pods — future support planned for other resources and validating webhooks.

---

## 🧱 Architecture

```text
+---------------------------+
|  Developer Applies a Pod |
+------------+--------------+
             |
             v
+-------------------------------+
| kube-apiserver (Admission)   |
| - Triggers MutatingWebhook   |
+-------------------------------+
             |
             v
+----------------------------+
|     kube-sentinel          |
| - Mutates Pod Spec         |
| - Injects Security Context |
+----------------------------+
             |
             v
+----------------------------+
|  Secured Pod is Deployed   |
+----------------------------+

