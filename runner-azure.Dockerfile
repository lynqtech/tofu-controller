ARG BASE_IMAGE

FROM mcr.microsoft.com/azure-cli:2.75.0 as azcli

FROM $BASE_IMAGE

ARG TARGETARCH
ARG TF_VERSION=1.5.7

# Switch to root to have permissions for operations
USER root

ADD https://releases.hashicorp.com/terraform/${TF_VERSION}/terraform_${TF_VERSION}_linux_${TARGETARCH}.zip /terraform_${TF_VERSION}_linux_${TARGETARCH}.zip
RUN unzip -q /terraform_${TF_VERSION}_linux_${TARGETARCH}.zip -d /usr/local/bin/ && \
    rm /terraform_${TF_VERSION}_linux_${TARGETARCH}.zip && \
    chmod +x /usr/local/bin/terraform

# Copy az cli
COPY --from=azcli /usr/bin/az /usr/bin/az

# Switch back to the non-root user after operations
USER 65532:65532

ENV GNUPGHOME=/tmp