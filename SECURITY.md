# Security Policy

## Reporting a Vulnerability

If you discover a security vulnerability in the Glean CLI, please report it responsibly.

**Do not open a public GitHub issue.** Instead, email security@glean.com with:

- A description of the vulnerability
- Steps to reproduce
- The version of the CLI affected (`glean version`)
- Any relevant logs or screenshots

We will acknowledge your report within 2 business days and aim to release a fix within 14 days for critical issues.

## Supported Versions

Only the latest release is actively supported with security patches.

| Version | Supported |
|---------|-----------|
| Latest  | ✅        |
| Older   | ❌        |

## Scope

The following are in scope:

- Authentication and credential handling (`glean auth`, token storage)
- Data exfiltration via CLI flags or environment variables
- Command injection via `--json` or other input flags
- Insecure defaults in the install script

The following are out of scope:

- Vulnerabilities in the Glean backend API (report to security@glean.com separately)
- Social engineering attacks
- Denial of service

## Disclosure Policy

We follow a coordinated disclosure process. Once a fix is released, we will publish a security advisory on the GitHub repository crediting the reporter (unless they prefer to remain anonymous).
